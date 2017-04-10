/*

Copyright 2017 Continusec Pty Ltd

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package api

import (
	"golang.org/x/net/context"

	"github.com/continusec/go-client/continusec"
	"github.com/continusec/objecthash"
	"github.com/continusec/vds-server/pb"
	"github.com/golang/protobuf/proto"
)

const (
	logTypeUser        = 0
	logTypeMapMutation = 1
	logTypeMapTreeHead = 2
)

const (
	operationRawAdd         = 1
	operationReadEntry      = 2
	operationReadHash       = 3
	operationProveInclusion = 4
)

var (
	operationForLogType = map[int]map[int]pb.Permission{
		logTypeUser: map[int]pb.Permission{
			operationRawAdd:         pb.Permission_PERM_LOG_RAW_ADD,
			operationReadEntry:      pb.Permission_PERM_LOG_READ_ENTRY,
			operationReadHash:       pb.Permission_PERM_LOG_READ_HASH,
			operationProveInclusion: pb.Permission_PERM_LOG_PROVE_INCLUSION,
		},
		logTypeMapMutation: map[int]pb.Permission{ // This is not a typo, we deliberately consider read entries of mutation log as separate and sensitive.
			operationReadEntry:      pb.Permission_PERM_MAP_MUTATION_READ_ENTRY,
			operationReadHash:       pb.Permission_PERM_MAP_MUTATION_READ_HASH,
			operationProveInclusion: pb.Permission_PERM_MAP_MUTATION_READ_HASH,
		},
		logTypeMapTreeHead: map[int]pb.Permission{
			operationReadEntry:      pb.Permission_PERM_MAP_MUTATION_READ_HASH,
			operationReadHash:       pb.Permission_PERM_MAP_MUTATION_READ_HASH,
			operationProveInclusion: pb.Permission_PERM_MAP_MUTATION_READ_HASH,
		},
	}
)

var (
	logHead = []byte("head")
)

type serverLog struct {
	account *serverAccount
	name    string
	logType int

	bucketKey []byte // lazily set
}

func (l *serverLog) verifyAccessForOperation(operation int) error {
	perm, ok := operationForLogType[l.logType][operation]
	if !ok {
		return ErrNotAuthorized
	}

	return l.verifyAccessForPermission(perm)
}

func (l *serverLog) verifyAccessForPermission(perm pb.Permission) error {
	return l.account.Service.AccessPolicy.VerifyAllowed(l.account.Account, l.account.APIKey, l.name, perm)
}

func (l *serverLog) bucket() ([]byte, error) {
	if l.bucketKey == nil {
		var err error
		l.bucketKey, err = objecthash.ObjectHash(map[string]interface{}{
			"account": l.account.Account,
			"name":    l.name,
			"type":    l.logType,
		})
		if err != nil {
			return nil, err
		}
	}
	return l.bucketKey, nil
}

// Create will send an API call to create a new log with the name specified when the
// verifiableLogImpl object was instantiated.
func (l *serverLog) Create() error {
	err := l.verifyAccessForPermission(pb.Permission_PERM_LOG_CREATE)
	if err != nil {
		return err
	}
	if l.logType != logTypeUser {
		return ErrInvalidRequest
	}
	if len(l.name) == 0 {
		return ErrInvalidRequest
	}
	promise, err := l.account.Service.Mutator.QueueMutation(&pb.Mutation{
		Account:   l.account.Account,
		Name:      l.name,
		Operation: pb.MutationType_MUT_LOG_CREATE,
	})
	if err != nil {
		return err
	}
	return promise.WaitUntilDone()
}

// Destroy will send an API call to delete this log - this operation removes it permanently,
// and renders the name unusable again within the same account, so please use with caution.
func (l *serverLog) Destroy() error {
	err := l.verifyAccessForPermission(pb.Permission_PERM_LOG_DELETE)
	if err != nil {
		return err
	}
	if l.logType != logTypeUser {
		return ErrInvalidRequest
	}
	promise, err := l.account.Service.Mutator.QueueMutation(&pb.Mutation{
		Account:   l.account.Account,
		Name:      l.name,
		Operation: pb.MutationType_MUT_LOG_DESTROY,
	})
	if err != nil {
		return err
	}
	return promise.WaitUntilDone()
}

// Add will send an API call to add the specified entry to the log. If the exact entry
// already exists in the log, it will not be added a second time.
// Returns an AddEntryResponse which includes the leaf hash, whether it is a duplicate or not. Note that the
// entry is sequenced in the underlying log in an asynchronous fashion, so the tree size
// will not immediately increase, and inclusion proof checks will not reflect the new entry
// until it is sequenced.
func (l *serverLog) Add(e continusec.UploadableEntry) (*continusec.AddEntryResponse, error) {
	err := l.verifyAccessForOperation(operationRawAdd)
	if err != nil {
		return nil, err
	}

	if l.logType != logTypeUser {
		return nil, ErrInvalidRequest
	}

	leafInput, dataToStore, err := e.DataForStorage()
	if err != nil {
		return nil, err
	}

	leafHash := continusec.LeafMerkleTreeHash(leafInput)

	_, err = l.account.Service.Mutator.QueueMutation(&pb.Mutation{
		Account:   l.account.Account,
		Name:      l.name,
		Operation: pb.MutationType_MUT_LOG_ADD,
		Mtl:       leafHash,
		Value:     dataToStore,
	})
	if err != nil {
		return nil, err
	}

	return &continusec.AddEntryResponse{EntryLeafHash: leafHash}, nil
}

func (l *serverLog) readIntoProto(kr KeyReader, key []byte, m proto.Message) error {
	bucket, err := l.bucket()
	if err != nil {
		return err
	}

	val, err := kr.Get(bucket, key)
	if err != nil {
		return err
	}

	return proto.Unmarshal(val, m)
}

// TreeHead returns tree root hash for the log at the given tree size. Specify continusec.Head
// to receive a root hash for the latest tree size.
func (l *serverLog) TreeHead(treeSize int64) (*continusec.LogTreeHead, error) {
	err := l.verifyAccessForOperation(operationReadHash)
	if err != nil {
		return nil, err
	}

	if treeSize < 0 {
		return nil, ErrInvalidTreeRange
	}

	var rv *continusec.LogTreeHead
	err = l.account.Service.Reader.ExecuteReadOnly(func(kr KeyReader) error {
		var lth pb.LogTreeHash
		err := l.readIntoProto(kr, logHead, &lth)
		switch err {
		case nil:
		// pass
		case ErrNoSuchKey:
			lth.Reset()
		default:
			return err
		}

		rv = &continusec.LogTreeHead{
			RootHash: lth.Hash,
			TreeSize: lth.Size,
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return rv, nil
}

// InclusionProof will return a proof the the specified MerkleTreeLeaf is included in the
// log. The proof consists of the index within the log that the entry is stored, and an
// audit path which returns the corresponding leaf nodes that can be applied to the input
// leaf hash to generate the root tree hash for the log.
//
// Most clients instead use VerifyInclusion which additionally verifies the returned proof.
func (l *serverLog) InclusionProof(treeSize int64, leaf continusec.MerkleTreeLeaf) (*continusec.LogInclusionProof, error) {
	err := l.verifyAccessForOperation(operationProveInclusion)
	if err != nil {
		return nil, err
	}

	return nil, ErrNotImplemented
}

// InclusionProofByIndex will return an inclusion proof for a specified tree size and leaf index.
// This is not used by typical clients, however it can be useful for certain audit operations and debugging tools.
// The LogInclusionProof returned by this method will not have the LeafHash filled in and as such will fail to verify.
//
// Typical clients will instead use VerifyInclusionProof().
func (l *serverLog) InclusionProofByIndex(treeSize, leafIndex int64) (*continusec.LogInclusionProof, error) {
	err := l.verifyAccessForOperation(operationProveInclusion)
	if err != nil {
		return nil, err
	}

	return nil, ErrNotImplemented
}

// ConsistencyProof returns an audit path which contains the set of Merkle Subtree hashes
// that demonstrate how the root hash is calculated for both the first and second tree sizes.
//
// Most clients instead use VerifyInclusionProof which additionally verifies the returned proof.
func (l *serverLog) ConsistencyProof(first, second int64) (*continusec.LogConsistencyProof, error) {
	err := l.verifyAccessForOperation(operationReadHash)
	if err != nil {
		return nil, err
	}

	return nil, ErrNotImplemented
}

// Entry returns the entry stored for the given index using the passed in factory to instantiate the entry.
// This is normally one of RawDataEntryFactory, JsonEntryFactory or RedactedJsonEntryFactory.
// If the entry was stored using one of the ObjectHash formats, then the data returned by a RawDataEntryFactory,
// then the object hash itself is returned as the contents. To get the data itself, use JsonEntryFactory.
func (l *serverLog) Entry(idx int64, factory continusec.VerifiableEntryFactory) (continusec.VerifiableEntry, error) {
	err := l.verifyAccessForOperation(operationReadEntry)
	if err != nil {
		return nil, err
	}

	return nil, ErrNotImplemented
}

// Entries batches requests to fetch entries from the server and returns a channel with the data
// for each entry. Close the context passed to terminate early if desired. If an error is
// encountered, the channel will be closed early before all items are returned.
//
// factory is normally one of one of RawDataEntryFactory, JsonEntryFactory or RedactedJsonEntryFactory.
func (l *serverLog) Entries(ctx context.Context, start, end int64, factory continusec.VerifiableEntryFactory) <-chan continusec.VerifiableEntry {
	err := l.verifyAccessForOperation(operationReadEntry)
	if err != nil {
		return nil
	}

	panic(ErrNotImplemented)
}

// Name returns the name of the log
func (l *serverLog) Name() string {
	switch l.logType {
	case logTypeUser:
		return l.name
	case logTypeMapMutation:
		return l.name + " (mutation log)"
	case logTypeMapTreeHead:
		return l.name + " (tree head log)"
	default:
		return ""
	}
}