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

// ConnectV1ConnectorError struct for ConnectV1ConnectorError
type ConnectV1ConnectorError struct {
	Error *ConnectV1ConnectorErrorError `json:"error,omitempty"`
}

// NewConnectV1ConnectorError instantiates a new ConnectV1ConnectorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectV1ConnectorError() *ConnectV1ConnectorError {
	this := ConnectV1ConnectorError{}
	return &this
}

// NewConnectV1ConnectorErrorWithDefaults instantiates a new ConnectV1ConnectorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectV1ConnectorErrorWithDefaults() *ConnectV1ConnectorError {
	this := ConnectV1ConnectorError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ConnectV1ConnectorError) GetError() ConnectV1ConnectorErrorError {
	if o == nil || o.Error == nil {
		var ret ConnectV1ConnectorErrorError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1ConnectorError) GetErrorOk() (*ConnectV1ConnectorErrorError, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ConnectV1ConnectorError) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given ConnectV1ConnectorErrorError and assigns it to the Error field.
func (o *ConnectV1ConnectorError) SetError(v ConnectV1ConnectorErrorError) {
	o.Error = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConnectV1ConnectorError) Redact() {
	o.recurseRedact(o.Error)
}

func (o *ConnectV1ConnectorError) recurseRedact(v interface{}) {
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

func (o ConnectV1ConnectorError) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConnectV1ConnectorError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConnectV1ConnectorError struct {
	value *ConnectV1ConnectorError
	isSet bool
}

func (v NullableConnectV1ConnectorError) Get() *ConnectV1ConnectorError {
	return v.value
}

func (v *NullableConnectV1ConnectorError) Set(val *ConnectV1ConnectorError) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1ConnectorError) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1ConnectorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1ConnectorError(val *ConnectV1ConnectorError) *NullableConnectV1ConnectorError {
	return &NullableConnectV1ConnectorError{value: val, isSet: true}
}

func (v NullableConnectV1ConnectorError) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConnectV1ConnectorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}