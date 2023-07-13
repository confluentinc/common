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
Confluent Schema Registry APIs

REST API for the Schema Registry

API version: 1.0.0
Contact: data-governance@confluent.io
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

// ConfigUpdateRequest Config update request
type ConfigUpdateRequest struct {
	// Compatibility Level
	Compatibility *string `json:"compatibility,omitempty"`
}

// NewConfigUpdateRequest instantiates a new ConfigUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigUpdateRequest() *ConfigUpdateRequest {
	this := ConfigUpdateRequest{}
	return &this
}

// NewConfigUpdateRequestWithDefaults instantiates a new ConfigUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigUpdateRequestWithDefaults() *ConfigUpdateRequest {
	this := ConfigUpdateRequest{}
	return &this
}

// GetCompatibility returns the Compatibility field value if set, zero value otherwise.
func (o *ConfigUpdateRequest) GetCompatibility() string {
	if o == nil || o.Compatibility == nil {
		var ret string
		return ret
	}
	return *o.Compatibility
}

// GetCompatibilityOk returns a tuple with the Compatibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigUpdateRequest) GetCompatibilityOk() (*string, bool) {
	if o == nil || o.Compatibility == nil {
		return nil, false
	}
	return o.Compatibility, true
}

// HasCompatibility returns a boolean if a field has been set.
func (o *ConfigUpdateRequest) HasCompatibility() bool {
	if o != nil && o.Compatibility != nil {
		return true
	}

	return false
}

// SetCompatibility gets a reference to the given string and assigns it to the Compatibility field.
func (o *ConfigUpdateRequest) SetCompatibility(v string) {
	o.Compatibility = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConfigUpdateRequest) Redact() {
	o.recurseRedact(o.Compatibility)
}

func (o *ConfigUpdateRequest) recurseRedact(v interface{}) {
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

func (o ConfigUpdateRequest) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConfigUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Compatibility != nil {
		toSerialize["compatibility"] = o.Compatibility
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConfigUpdateRequest struct {
	value *ConfigUpdateRequest
	isSet bool
}

func (v NullableConfigUpdateRequest) Get() *ConfigUpdateRequest {
	return v.value
}

func (v *NullableConfigUpdateRequest) Set(val *ConfigUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigUpdateRequest(val *ConfigUpdateRequest) *NullableConfigUpdateRequest {
	return &NullableConfigUpdateRequest{value: val, isSet: true}
}

func (v NullableConfigUpdateRequest) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConfigUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
