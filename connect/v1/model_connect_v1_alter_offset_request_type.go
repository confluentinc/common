// Copyright 2021 Confluent Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
Kafka Connect APIs

REST API for managing connectors

API version: 1.0
Contact: connect@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// ConnectV1AlterOffsetRequestType The type of alter operation. PATCH will update the offset to the provided values. The update will only happen for the partitions provided in the request.  DELETE will delete the offset for the provided partitions and reset them back to the base state. It is as if, a fresh new connector was created.  For sink connectors PATCH/DELETE will move the offsets to the provided point in the  topic partition. If the offset provided is not present in the topic partition it will by default reset to the earliest offset in the topic partition.  For source connectors, post PATCH/DELETE the connector will attempt to read from the  position defined in the altered offsets.
type ConnectV1AlterOffsetRequestType string

// List of connect.v1.AlterOffsetRequestType
const (
	PATCH  ConnectV1AlterOffsetRequestType = "PATCH"
	DELETE ConnectV1AlterOffsetRequestType = "DELETE"
)

var allowedConnectV1AlterOffsetRequestTypeEnumValues = []ConnectV1AlterOffsetRequestType{
	"PATCH",
	"DELETE",
}

func (v *ConnectV1AlterOffsetRequestType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectV1AlterOffsetRequestType(value)
	for _, existing := range allowedConnectV1AlterOffsetRequestTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectV1AlterOffsetRequestType", value)
}

// NewConnectV1AlterOffsetRequestTypeFromValue returns a pointer to a valid ConnectV1AlterOffsetRequestType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectV1AlterOffsetRequestTypeFromValue(v string) (*ConnectV1AlterOffsetRequestType, error) {
	ev := ConnectV1AlterOffsetRequestType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectV1AlterOffsetRequestType: valid values are %v", v, allowedConnectV1AlterOffsetRequestTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectV1AlterOffsetRequestType) IsValid() bool {
	for _, existing := range allowedConnectV1AlterOffsetRequestTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to connect.v1.AlterOffsetRequestType value
func (v ConnectV1AlterOffsetRequestType) Ptr() *ConnectV1AlterOffsetRequestType {
	return &v
}

type NullableConnectV1AlterOffsetRequestType struct {
	value *ConnectV1AlterOffsetRequestType
	isSet bool
}

func (v NullableConnectV1AlterOffsetRequestType) Get() *ConnectV1AlterOffsetRequestType {
	return v.value
}

func (v *NullableConnectV1AlterOffsetRequestType) Set(val *ConnectV1AlterOffsetRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1AlterOffsetRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1AlterOffsetRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1AlterOffsetRequestType(val *ConnectV1AlterOffsetRequestType) *NullableConnectV1AlterOffsetRequestType {
	return &NullableConnectV1AlterOffsetRequestType{value: val, isSet: true}
}

func (v NullableConnectV1AlterOffsetRequestType) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConnectV1AlterOffsetRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}