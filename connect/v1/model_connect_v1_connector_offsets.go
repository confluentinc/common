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

// ConnectV1ConnectorOffsets Offsets for a connector
type ConnectV1ConnectorOffsets struct {
	// The name of the connector.
	Name *string `json:"name,omitempty"`
	// The ID of the connector.
	Id       *string                            `json:"id,omitempty"`
	Offsets  *ConnectV1Offsets                  `json:"offsets,omitempty"`
	Metadata *ConnectV1ConnectorOffsetsMetadata `json:"metadata,omitempty"`
}

// NewConnectV1ConnectorOffsets instantiates a new ConnectV1ConnectorOffsets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectV1ConnectorOffsets() *ConnectV1ConnectorOffsets {
	this := ConnectV1ConnectorOffsets{}
	return &this
}

// NewConnectV1ConnectorOffsetsWithDefaults instantiates a new ConnectV1ConnectorOffsets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectV1ConnectorOffsetsWithDefaults() *ConnectV1ConnectorOffsets {
	this := ConnectV1ConnectorOffsets{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ConnectV1ConnectorOffsets) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1ConnectorOffsets) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ConnectV1ConnectorOffsets) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ConnectV1ConnectorOffsets) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectV1ConnectorOffsets) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1ConnectorOffsets) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectV1ConnectorOffsets) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectV1ConnectorOffsets) SetId(v string) {
	o.Id = &v
}

// GetOffsets returns the Offsets field value if set, zero value otherwise.
func (o *ConnectV1ConnectorOffsets) GetOffsets() ConnectV1Offsets {
	if o == nil || o.Offsets == nil {
		var ret ConnectV1Offsets
		return ret
	}
	return *o.Offsets
}

// GetOffsetsOk returns a tuple with the Offsets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1ConnectorOffsets) GetOffsetsOk() (*ConnectV1Offsets, bool) {
	if o == nil || o.Offsets == nil {
		return nil, false
	}
	return o.Offsets, true
}

// HasOffsets returns a boolean if a field has been set.
func (o *ConnectV1ConnectorOffsets) HasOffsets() bool {
	if o != nil && o.Offsets != nil {
		return true
	}

	return false
}

// SetOffsets gets a reference to the given ConnectV1Offsets and assigns it to the Offsets field.
func (o *ConnectV1ConnectorOffsets) SetOffsets(v ConnectV1Offsets) {
	o.Offsets = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ConnectV1ConnectorOffsets) GetMetadata() ConnectV1ConnectorOffsetsMetadata {
	if o == nil || o.Metadata == nil {
		var ret ConnectV1ConnectorOffsetsMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1ConnectorOffsets) GetMetadataOk() (*ConnectV1ConnectorOffsetsMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ConnectV1ConnectorOffsets) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConnectV1ConnectorOffsetsMetadata and assigns it to the Metadata field.
func (o *ConnectV1ConnectorOffsets) SetMetadata(v ConnectV1ConnectorOffsetsMetadata) {
	o.Metadata = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConnectV1ConnectorOffsets) Redact() {
	o.recurseRedact(o.Name)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Offsets)
	o.recurseRedact(o.Metadata)
}

func (o *ConnectV1ConnectorOffsets) recurseRedact(v interface{}) {
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

func (o ConnectV1ConnectorOffsets) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConnectV1ConnectorOffsets) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Offsets != nil {
		toSerialize["offsets"] = o.Offsets
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConnectV1ConnectorOffsets struct {
	value *ConnectV1ConnectorOffsets
	isSet bool
}

func (v NullableConnectV1ConnectorOffsets) Get() *ConnectV1ConnectorOffsets {
	return v.value
}

func (v *NullableConnectV1ConnectorOffsets) Set(val *ConnectV1ConnectorOffsets) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1ConnectorOffsets) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1ConnectorOffsets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1ConnectorOffsets(val *ConnectV1ConnectorOffsets) *NullableConnectV1ConnectorOffsets {
	return &NullableConnectV1ConnectorOffsets{value: val, isSet: true}
}

func (v NullableConnectV1ConnectorOffsets) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConnectV1ConnectorOffsets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
