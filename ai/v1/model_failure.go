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
AI API

API for interacting with AI models from within Confluent Cloud.

API version: 0.0.1
Contact: api-team@confluent.io
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

// Failure Provides information about problems encountered while performing an operation.
type Failure struct {
	// List of errors which caused this operation to fail
	Errors []Error `json:"errors,omitempty"`
}

// NewFailure instantiates a new Failure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailure(errors []Error) *Failure {
	this := Failure{}
	this.Errors = errors
	return &this
}

// NewFailureWithDefaults instantiates a new Failure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailureWithDefaults() *Failure {
	this := Failure{}
	return &this
}

// GetErrors returns the Errors field value
func (o *Failure) GetErrors() []Error {
	if o == nil {
		var ret []Error
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *Failure) GetErrorsOk() (*[]Error, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *Failure) SetErrors(v []Error) {
	o.Errors = v
}

// Redact resets all sensitive fields to their zero value.
func (o *Failure) Redact() {
	o.recurseRedact(&o.Errors)
}

func (o *Failure) recurseRedact(v interface{}) {
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

func (o Failure) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o Failure) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["errors"] = o.Errors
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableFailure struct {
	value *Failure
	isSet bool
}

func (v NullableFailure) Get() *Failure {
	return v.value
}

func (v *NullableFailure) Set(val *Failure) {
	v.value = val
	v.isSet = true
}

func (v NullableFailure) IsSet() bool {
	return v.isSet
}

func (v *NullableFailure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailure(val *Failure) *NullableFailure {
	return &NullableFailure{value: val, isSet: true}
}

func (v NullableFailure) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableFailure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
