// Copyright 2020 Harald Albrecht.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy
// of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// +build !go1.14

package testbasher

import (
	"encoding/json"
	"io"
)

// Decoder just wraps a json.Decoder on Golang up to 1.13, without an input
// data memento.
type Decoder struct {
	// wrapped JSON decoder
	*json.Decoder
}

// NewDecoder returns a new JSON decoder, reading from the specified reader.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		Decoder: json.NewDecoder(r),
	}
}
