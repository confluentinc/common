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

// ConnectV1ConnectorErrorError Connector Error with error code and message.
type ConnectV1ConnectorErrorError struct {
	// Error code for the type of error
	Code *int32 `json:"code,omitempty"`
	// Human readable error message
	Message *string `json:"message,omitempty"`
}

// NewConnectV1ConnectorErrorError instantiates a new ConnectV1ConnectorErrorError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectV1ConnectorErrorError() *ConnectV1ConnectorErrorError {
	this := ConnectV1ConnectorErrorError{}
	return &this
}

// NewConnectV1ConnectorErrorErrorWithDefaults instantiates a new ConnectV1ConnectorErrorError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectV1ConnectorErrorErrorWithDefaults() *ConnectV1ConnectorErrorError {
	this := ConnectV1ConnectorErrorError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ConnectV1ConnectorErrorError) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1ConnectorErrorError) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ConnectV1ConnectorErrorError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ConnectV1ConnectorErrorError) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ConnectV1ConnectorErrorError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectV1ConnectorErrorError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ConnectV1ConnectorErrorError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ConnectV1ConnectorErrorError) SetMessage(v string) {
	o.Message = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConnectV1ConnectorErrorError) Redact() {
	o.recurseRedact(o.Code)
	o.recurseRedact(o.Message)
}

func (o *ConnectV1ConnectorErrorError) recurseRedact(v interface{}) {
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

func (o ConnectV1ConnectorErrorError) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConnectV1ConnectorErrorError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConnectV1ConnectorErrorError struct {
	value *ConnectV1ConnectorErrorError
	isSet bool
}

func (v NullableConnectV1ConnectorErrorError) Get() *ConnectV1ConnectorErrorError {
	return v.value
}

func (v *NullableConnectV1ConnectorErrorError) Set(val *ConnectV1ConnectorErrorError) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectV1ConnectorErrorError) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectV1ConnectorErrorError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectV1ConnectorErrorError(val *ConnectV1ConnectorErrorError) *NullableConnectV1ConnectorErrorError {
	return &NullableConnectV1ConnectorErrorError{value: val, isSet: true}
}

func (v NullableConnectV1ConnectorErrorError) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConnectV1ConnectorErrorError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}