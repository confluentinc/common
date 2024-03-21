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
	"time"
)

import (
	"reflect"
)

// ConnectV1ConnectorOffsetsMetadata Metadata of the connector offset.
type ConnectV1ConnectorOffsetsMetadata struct {
	// The time at which the offsets were observed. The time is in UTC, ISO 8601 format.
	ObservedAt *time.Time `json:"observed_at,omitempty"`
}

// NewConnectV1ConnectorOffsetsMetadata instantiates a new ConnectV1ConnectorOffsetsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectV1ConnectorOffsetsMetadata() *ConnectV1ConnectorOffsetsMetadata {
	this := ConnectV1ConnectorOffsetsMetadata{}
	return &this
}

// NewConnectV1ConnectorOffsetsMetadataWithDefaults instantiates a new ConnectV1ConnectorOffsetsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectV1ConnectorOffsetsMetadataWithDefaults() *ConnectV1ConnectorOffsetsMetadata {
	this := ConnectV1ConnectorOffsetsMetadata{}
	return &this
}

// GetObservedAt returns the ObservedAt field value if set, zero value otherwise.
func (o *ConnectV1ConnectorOffsetsMetadata) GetObservedAt() time.Time {
	if o == nil || o.ObservedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ObservedAt
}

// GetObservedAtOk returns a tuple with the ObservedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1ConnectorOffsetsMetadata) GetObservedAtOk() (*time.Time, bool) {
	if o == nil || o.ObservedAt == nil {
		return nil, false
	}
	return o.ObservedAt, true
}

// HasObservedAt returns a boolean if a field has been set.
func (o *ConnectV1ConnectorOffsetsMetadata) HasObservedAt() bool {
	if o != nil && o.ObservedAt != nil {
		return true
	}

	return false
}

// SetObservedAt gets a reference to the given time.Time and assigns it to the ObservedAt field.
func (o *ConnectV1ConnectorOffsetsMetadata) SetObservedAt(v time.Time) {
	o.ObservedAt = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConnectV1ConnectorOffsetsMetadata) Redact() {
	o.recurseRedact(o.ObservedAt)
}

func (o *ConnectV1ConnectorOffsetsMetadata) recurseRedact(v interface{}) {
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

func (o ConnectV1ConnectorOffsetsMetadata) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConnectV1ConnectorOffsetsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObservedAt != nil {
		toSerialize["observed_at"] = o.ObservedAt
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConnectV1ConnectorOffsetsMetadata struct {
	value *ConnectV1ConnectorOffsetsMetadata
	isSet bool
}

func (v NullableConnectV1ConnectorOffsetsMetadata) Get() *ConnectV1ConnectorOffsetsMetadata {
	return v.value
}

func (v *NullableConnectV1ConnectorOffsetsMetadata) Set(val *ConnectV1ConnectorOffsetsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1ConnectorOffsetsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1ConnectorOffsetsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1ConnectorOffsetsMetadata(val *ConnectV1ConnectorOffsetsMetadata) *NullableConnectV1ConnectorOffsetsMetadata {
	return &NullableConnectV1ConnectorOffsetsMetadata{value: val, isSet: true}
}

func (v NullableConnectV1ConnectorOffsetsMetadata) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConnectV1ConnectorOffsetsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
