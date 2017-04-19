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
	"time"

	"golang.org/x/net/context"

	"github.com/golang/protobuf/proto"
)

func makeJSONMutationEntry(req *MapSetValueRequest) (*MapMutation, error) {
	// is there a better way to clone?
	b, err := proto.Marshal(req.Mutation)
	if err != nil {
		return nil, err
	}
	var rv MapMutation
	err = proto.Unmarshal(b, &rv)
	if err != nil {
		return nil, err
	}
	rv.Timestamp = time.Now().Format(time.RFC3339Nano)
	return &rv, nil
}

// MapSetValue sets a value in a map
func (s *LocalService) MapSetValue(ctx context.Context, req *MapSetValueRequest) (*MapSetValueResponse, error) {
	_, err := s.verifyAccessForMap(req.Map, Permission_PERM_MAP_SET_VALUE)
	if err != nil {
		return nil, err
	}

	// Since this whacks in timestamp (deliberately, so that mutation log receives unique mutations),
	// we must keep this rather than regenerate another object.
	mm, err := makeJSONMutationEntry(req)
	if err != nil {
		return nil, ErrInvalidRequest
	}

	mutData, err := CreateJSONLeafDataFromProto(mm)
	if err != nil {
		return nil, err
	}

	ns, err := mapBucket(req.Map)
	if err != nil {
		return nil, ErrInvalidRequest
	}

	err = s.Mutator.QueueMutation(ns, &Mutation{
		LogAddEntry: &LogAddEntryRequest{
			Log: &LogRef{
				Account: req.Map.Account,
				Name:    req.Map.Name,
				LogType: LogType_STRUCT_TYPE_MUTATION_LOG,
			},
			Value: mutData,
		},
	})
	if err != nil {
		return nil, err
	}
	return &MapSetValueResponse{
		LeafHash: LeafMerkleTreeHash(mutData.LeafInput),
	}, nil
}