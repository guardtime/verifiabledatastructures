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
	"bytes"
	"encoding/json"

	"github.com/continusec/objecthash"

	"github.com/golang/protobuf/proto"
)

// CreateJSONLeafData creates a JSON based objecthash for the given JSON bytes.
func CreateJSONLeafData(data []byte) (*LeafData, error) {
	var o interface{}
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, ErrInvalidJSON
	}

	bflh, err := objecthash.ObjectHashWithStdRedaction(o)
	if err != nil {
		return nil, err
	}

	return &LeafData{
		LeafInput: bflh,
		ExtraData: data,
		Format:    DataFormat_JSON,
	}, nil
}

// CreateJSONLeafDataFromProto creates a JSON based objecthash for the given proto.
func CreateJSONLeafDataFromProto(m proto.Message) (*LeafData, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return CreateJSONLeafData(b)
}

// CreateJSONLeafDataFromObject creates a JSON based objecthash for the given object.
// The object is first serialized then unmarshalled into a string map.
func CreateJSONLeafDataFromObject(o interface{}) (*LeafData, error) {
	data, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	return CreateJSONLeafData(data)
}

// ValidateJSONLeafData verifies that the LeafInput is equal to the objecthash of the ExtraData.
// It ignores the format field.
func ValidateJSONLeafData(entry *LeafData) error {
	var o interface{}
	err := json.Unmarshal(entry.ExtraData, &o)
	if err != nil {
		return ErrVerificationFailed
	}
	h, err := objecthash.ObjectHashWithStdRedaction(o)
	if err != nil {
		return ErrVerificationFailed
	}
	if !bytes.Equal(entry.LeafInput, h) {
		return ErrVerificationFailed
	}
	return nil
}

// ShedRedactedJSONFields will parse the JSON, shed any redacted fields,
// and replace othe redactable tuples with the value component only,
// then write back out JSON suitable for parsing into the eventual object
// it should be used in.
func ShedRedactedJSONFields(b []byte) ([]byte, error) {
	var contents interface{}
	err := json.Unmarshal(b, &contents)
	if err != nil {
		return nil, err
	}
	newContents, err := objecthash.UnredactableWithStdPrefix(contents)
	if err != nil {
		return nil, err
	}
	return json.Marshal(newContents)
}

// CreateRedactableJSONLeafData creates a LeafData node with fields suitable
// for redaction, ie it replaces all values with a <nonce, value> tuple.
func CreateRedactableJSONLeafData(data []byte) (*LeafData, error) {
	var o interface{}
	err := json.Unmarshal(data, &o)
	if err != nil {
		return nil, ErrInvalidJSON
	}

	o, err = objecthash.Redactable(o)
	if err != nil {
		return nil, err
	}

	ojb, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}

	bflh, err := objecthash.ObjectHash(o)
	if err != nil {
		return nil, err
	}

	return &LeafData{
		LeafInput: bflh,
		ExtraData: ojb,
		Format:    DataFormat_JSON,
	}, nil
}