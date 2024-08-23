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
)

import (
	"reflect"
)

// ConnectV1AlterOffsetRequest Request to alter the offset of a connector. The offsets parameter is options for DELETE type.
type ConnectV1AlterOffsetRequest struct {
	Type ConnectV1AlterOffsetRequestType `json:"type,omitempty"`
	// Array of offsets which are categorised into partitions.
	Offsets *[]map[string]interface{} `json:"offsets,omitempty"`
}

// NewConnectV1AlterOffsetRequest instantiates a new ConnectV1AlterOffsetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectV1AlterOffsetRequest(type_ ConnectV1AlterOffsetRequestType) *ConnectV1AlterOffsetRequest {
	this := ConnectV1AlterOffsetRequest{}
	this.Type = type_
	return &this
}

// NewConnectV1AlterOffsetRequestWithDefaults instantiates a new ConnectV1AlterOffsetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectV1AlterOffsetRequestWithDefaults() *ConnectV1AlterOffsetRequest {
	this := ConnectV1AlterOffsetRequest{}
	return &this
}

// GetType returns the Type field value
func (o *ConnectV1AlterOffsetRequest) GetType() ConnectV1AlterOffsetRequestType {
	if o == nil {
		var ret ConnectV1AlterOffsetRequestType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ConnectV1AlterOffsetRequest) GetTypeOk() (*ConnectV1AlterOffsetRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ConnectV1AlterOffsetRequest) SetType(v ConnectV1AlterOffsetRequestType) {
	o.Type = v
}

// GetOffsets returns the Offsets field value if set, zero value otherwise.
func (o *ConnectV1AlterOffsetRequest) GetOffsets() []map[string]interface{} {
	if o == nil || o.Offsets == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Offsets
}

// GetOffsetsOk returns a tuple with the Offsets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1AlterOffsetRequest) GetOffsetsOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Offsets == nil {
		return nil, false
	}
	return o.Offsets, true
}

// HasOffsets returns a boolean if a field has been set.
func (o *ConnectV1AlterOffsetRequest) HasOffsets() bool {
	if o != nil && o.Offsets != nil {
		return true
	}

	return false
}

// SetOffsets gets a reference to the given []map[string]interface{} and assigns it to the Offsets field.
func (o *ConnectV1AlterOffsetRequest) SetOffsets(v []map[string]interface{}) {
	o.Offsets = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConnectV1AlterOffsetRequest) Redact() {
	o.recurseRedact(&o.Type)
	o.recurseRedact(o.Offsets)
}

func (o *ConnectV1AlterOffsetRequest) recurseRedact(v interface{}) {
	type redactor interface {
		Redact()
	}
	if r, ok := v.(redactor); ok {
		r.Redact()
	} else {
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		switch val.Kind() {
		case reflect.Slice, reflect.Array:
			for i := 0; i < val.Len(); i++ {
				// support data types declared without pointers
				o.recurseRedact(val.Index(i).Interface())
				// ... and data types that were declared without but need pointers (for Redact)
				if val.Index(i).CanAddr() {
					o.recurseRedact(val.Index(i).Addr().Interface())
				}
			}
		}
	}
}

func (o ConnectV1AlterOffsetRequest) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConnectV1AlterOffsetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Offsets != nil {
		toSerialize["offsets"] = o.Offsets
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConnectV1AlterOffsetRequest struct {
	value *ConnectV1AlterOffsetRequest
	isSet bool
}

func (v NullableConnectV1AlterOffsetRequest) Get() *ConnectV1AlterOffsetRequest {
	return v.value
}

func (v *NullableConnectV1AlterOffsetRequest) Set(val *ConnectV1AlterOffsetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1AlterOffsetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1AlterOffsetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1AlterOffsetRequest(val *ConnectV1AlterOffsetRequest) *NullableConnectV1AlterOffsetRequest {
	return &NullableConnectV1AlterOffsetRequest{value: val, isSet: true}
}

func (v NullableConnectV1AlterOffsetRequest) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConnectV1AlterOffsetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}