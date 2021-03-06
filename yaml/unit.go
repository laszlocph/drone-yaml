// Copyright 2019 Drone IO, Inc.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//      http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yaml

import (
	"strconv"
	"strings"

	"github.com/docker/go-units"
)
// BytesSize stores a human-readable size in bytes,
// kibibytes, mebibytes, gibibytes, or tebibytes
// (eg. "44kiB", "17MiB").
type BytesSize int64

// UnmarshalYAML implements yaml unmarshalling.
func (b *BytesSize) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var intType int64
	if err := unmarshal(&intType); err == nil {
		*b = BytesSize(intType)
		return nil
	}

	var stringType string
	if err := unmarshal(&stringType); err != nil {
		return err
	}

	intType, err := units.RAMInBytes(stringType)
	if err == nil {
		*b = BytesSize(intType)
	}
	return err
}

// String returns a human-readable size in bytes,
// kibibytes, mebibytes, gibibytes, or tebibytes
// (eg. "44kiB", "17MiB").
func (b BytesSize) String() string {
	return units.BytesSize(float64(b))
}

// MilliSize will convert cpus to millicpus as int64.
// for instance "1" will be converted to 1000 and "100m" to 100
type MilliSize int64

// UnmarshalYAML implements yaml unmarshalling.
func (m *MilliSize) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var intType int64
	if err := unmarshal(&intType); err == nil {
		*m = MilliSize(intType * 1000)
		return nil
	}

	var stringType string
	if err := unmarshal(&stringType); err != nil {
		return err
	}
	if len(stringType) > 0 {
		lastChar := string(stringType[len(stringType)-1:])
		if lastChar == "m" {
			// convert to int64
			i, err := strconv.ParseInt(strings.TrimSuffix(stringType, "m"), 10, 64)
			if err != nil {
				return err
			}
			*m = MilliSize(i)
		}
	}
	return nil
}

// String returns a human-readable cpu millis,
// (eg. "1000", "10").
func (m MilliSize) String() string {
	return strconv.FormatInt(int64(m), 10)
}
