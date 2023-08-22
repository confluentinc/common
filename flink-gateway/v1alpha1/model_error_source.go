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
SQL API v1alpha1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha1

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// ErrorSource If this error was caused by a particular part of the API request, the source will point to the query string parameter or request body property that caused it.
type ErrorSource struct {
	// A JSON Pointer [RFC6901] to the associated entity in the request document [e.g. \"/spec\" for a spec object, or \"/spec/title\" for a specific field].
	Pointer *string `json:"pointer,omitempty"`
	// A string indicating which query parameter caused the error.
	Parameter *string `json:"parameter,omitempty"`
}

// NewErrorSource instantiates a new ErrorSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorSource() *ErrorSource {
	this := ErrorSource{}
	return &this
}

// NewErrorSourceWithDefaults instantiates a new ErrorSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorSourceWithDefaults() *ErrorSource {
	this := ErrorSource{}
	return &this
}

// GetPointer returns the Pointer field value if set, zero value otherwise.
func (o *ErrorSource) GetPointer() string {
	if o == nil || o.Pointer == nil {
		var ret string
		return ret
	}
	return *o.Pointer
}

// GetPointerOk returns a tuple with the Pointer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSource) GetPointerOk() (*string, bool) {
	if o == nil || o.Pointer == nil {
		return nil, false
	}
	return o.Pointer, true
}

// HasPointer returns a boolean if a field has been set.
func (o *ErrorSource) HasPointer() bool {
	if o != nil && o.Pointer != nil {
		return true
	}

	return false
}

// SetPointer gets a reference to the given string and assigns it to the Pointer field.
func (o *ErrorSource) SetPointer(v string) {
	o.Pointer = &v
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *ErrorSource) GetParameter() string {
	if o == nil || o.Parameter == nil {
		var ret string
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSource) GetParameterOk() (*string, bool) {
	if o == nil || o.Parameter == nil {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *ErrorSource) HasParameter() bool {
	if o != nil && o.Parameter != nil {
		return true
	}

	return false
}

// SetParameter gets a reference to the given string and assigns it to the Parameter field.
func (o *ErrorSource) SetParameter(v string) {
	o.Parameter = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ErrorSource) Redact() {
	o.recurseRedact(o.Pointer)
	o.recurseRedact(o.Parameter)
}

func (o *ErrorSource) recurseRedact(v interface{}) {
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

func (o ErrorSource) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ErrorSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pointer != nil {
		toSerialize["pointer"] = o.Pointer
	}
	if o.Parameter != nil {
		toSerialize["parameter"] = o.Parameter
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableErrorSource struct {
	value *ErrorSource
	isSet bool
}

func (v NullableErrorSource) Get() *ErrorSource {
	return v.value
}

func (v *NullableErrorSource) Set(val *ErrorSource) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorSource) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorSource(val *ErrorSource) *NullableErrorSource {
	return &NullableErrorSource{value: val, isSet: true}
}

func (v NullableErrorSource) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableErrorSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
