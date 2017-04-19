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

package verifiabledatastructures

import (
	"encoding/json"

	"github.com/continusec/objecthash"
)

const (
	operationRawAdd         = 1
	operationReadEntry      = 2
	operationReadHash       = 3
	operationProveInclusion = 4
)

var (
	operationForLogType = map[LogType]map[int]Permission{
		LogType_STRUCT_TYPE_LOG: map[int]Permission{
			operationRawAdd:         Permission_PERM_LOG_RAW_ADD,
			operationReadEntry:      Permission_PERM_LOG_READ_ENTRY,
			operationReadHash:       Permission_PERM_LOG_READ_HASH,
			operationProveInclusion: Permission_PERM_LOG_PROVE_INCLUSION,
		},
		LogType_STRUCT_TYPE_MUTATION_LOG: map[int]Permission{ // This is not a typo, we deliberately consider read entries of mutation log as separate and sensitive.
			operationReadEntry:      Permission_PERM_MAP_MUTATION_READ_ENTRY,
			operationReadHash:       Permission_PERM_MAP_MUTATION_READ_HASH,
			operationProveInclusion: Permission_PERM_MAP_MUTATION_READ_HASH,
		},
		LogType_STRUCT_TYPE_TREEHEAD_LOG: map[int]Permission{
			operationReadEntry:      Permission_PERM_MAP_MUTATION_READ_HASH,
			operationReadHash:       Permission_PERM_MAP_MUTATION_READ_HASH,
			operationProveInclusion: Permission_PERM_MAP_MUTATION_READ_HASH,
		},
	}
)

func (s *LocalService) verifyAccessForMap(vmap *MapRef, perm Permission) (*AccessModifier, error) {
	return s.AccessPolicy.VerifyAllowed(vmap.Account.Id, vmap.Account.ApiKey, vmap.Name, perm)
}

func (s *LocalService) verifyAccessForLog(log *LogRef, perm Permission) (*AccessModifier, error) {
	return s.AccessPolicy.VerifyAllowed(log.Account.Id, log.Account.ApiKey, log.Name, perm)
}

func (s *LocalService) verifyAccessForLogOperation(log *LogRef, op int) (*AccessModifier, error) {
	perm, ok := operationForLogType[log.LogType][op]
	if !ok {
		return nil, ErrNotAuthorized
	}

	return s.verifyAccessForLog(log, perm)
}

func filterLeafData(ld *LeafData, am *AccessModifier) (*LeafData, error) {
	if am.FieldFilter == AllFields {
		return ld, nil
	}
	// Largely just for show, but saves us working on the wrong data type
	switch ld.Format {
	case DataFormat_JSON:
		var o interface{}
		err := json.Unmarshal(ld.ExtraData, &o)
		if err != nil {
			return nil, ErrInvalidJSON
		}
		o, err = objecthash.Filtered(o, am.FieldFilter)
		if err != nil {
			return nil, err
		}
		rv, err := json.Marshal(o)
		if err != nil {
			return nil, err
		}
		return &LeafData{
			LeafInput: ld.LeafInput,
			ExtraData: rv, // redacted form
		}, nil
	default:
		return nil, ErrNotImplemented
	}
}