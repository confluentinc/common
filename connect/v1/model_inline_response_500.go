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

// InlineResponse500 struct for InlineResponse500
type InlineResponse500 struct {
	ErrorCode *int32  `json:"error_code,omitempty"`
	Message   *string `json:"message,omitempty"`
}

// NewInlineResponse500 instantiates a new InlineResponse500 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse500() *InlineResponse500 {
	this := InlineResponse500{}
	return &this
}

// NewInlineResponse500WithDefaults instantiates a new InlineResponse500 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse500WithDefaults() *InlineResponse500 {
	this := InlineResponse500{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *InlineResponse500) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse500) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *InlineResponse500) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *InlineResponse500) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponse500) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse500) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponse500) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponse500) SetMessage(v string) {
	o.Message = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *InlineResponse500) Redact() {
	o.recurseRedact(o.ErrorCode)
	o.recurseRedact(o.Message)
}

func (o *InlineResponse500) recurseRedact(v interface{}) {
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

func (o InlineResponse500) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o InlineResponse500) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
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

type NullableInlineResponse500 struct {
	value *InlineResponse500
	isSet bool
}

func (v NullableInlineResponse500) Get() *InlineResponse500 {
	return v.value
}

func (v *NullableInlineResponse500) Set(val *InlineResponse500) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse500) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse500) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse500(val *InlineResponse500) *NullableInlineResponse500 {
	return &NullableInlineResponse500{value: val, isSet: true}
}

func (v NullableInlineResponse500) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableInlineResponse500) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
