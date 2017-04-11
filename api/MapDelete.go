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

	"github.com/continusec/verifiabledatastructures/pb"
)

func (s *LocalService) MapDelete(ctx context.Context, req *pb.MapDeleteRequest) (*pb.MapDeleteResponse, error) {
	err := s.verifyAccessForMap(req.Map, pb.Permission_PERM_MAP_DELETE)
	if err != nil {
		return nil, err
	}
	if len(req.Map.Name) == 0 {
		return nil, ErrInvalidRequest
	}
	promise, err := s.Mutator.QueueMutation(&pb.Mutation{
		Account:   req.Map.Account.Id,
		Name:      req.Map.Name,
		Operation: pb.MutationType_MUT_MAP_DESTROY,
	})
	if err != nil {
		return nil, err
	}
	err = promise.WaitUntilDone()
	if err != nil {
		return nil, err
	}
	return &pb.MapDeleteResponse{}, nil
}
