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
	"golang.org/x/net/context"
)

// LogAddEntry adds an entry to a log
func (s *LocalService) LogAddEntry(ctx context.Context, req *LogAddEntryRequest) (*LogAddEntryResponse, error) {
	_, err := s.verifyAccessForLogOperation(req.Log, operationRawAdd)
	if err != nil {
		return nil, err
	}

	if req.Log.LogType != LogType_STRUCT_TYPE_LOG {
		return nil, ErrInvalidRequest
	}

	ns, err := logBucket(req.Log)
	if err != nil {
		return nil, ErrInvalidRequest
	}

	err = s.Mutator.QueueMutation(ns, &Mutation{
		LogAddEntry: req,
	})
	if err != nil {
		return nil, err
	}

	return &LogAddEntryResponse{
		LeafHash: LeafMerkleTreeHash(req.Value.LeafInput),
	}, nil
}