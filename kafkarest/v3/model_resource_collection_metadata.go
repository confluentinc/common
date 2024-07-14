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
REST Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0
Contact: kafka-clients-proxy-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// ResourceCollectionMetadata struct for ResourceCollectionMetadata
type ResourceCollectionMetadata struct {
	Self string         `json:"self,omitempty"`
	Next NullableString `json:"next,omitempty"`
}

// NewResourceCollectionMetadata instantiates a new ResourceCollectionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceCollectionMetadata(self string) *ResourceCollectionMetadata {
	this := ResourceCollectionMetadata{}
	this.Self = self
	return &this
}

// NewResourceCollectionMetadataWithDefaults instantiates a new ResourceCollectionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceCollectionMetadataWithDefaults() *ResourceCollectionMetadata {
	this := ResourceCollectionMetadata{}
	return &this
}

// GetSelf returns the Self field value
func (o *ResourceCollectionMetadata) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ResourceCollectionMetadata) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *ResourceCollectionMetadata) SetSelf(v string) {
	o.Self = v
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceCollectionMetadata) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceCollectionMetadata) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *ResourceCollectionMetadata) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *ResourceCollectionMetadata) SetNext(v string) {
	o.Next.Set(&v)
}

// SetNextNil sets the value for Next to be an explicit nil
func (o *ResourceCollectionMetadata) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *ResourceCollectionMetadata) UnsetNext() {
	o.Next.Unset()
}

// Redact resets all sensitive fields to their zero value.
func (o *ResourceCollectionMetadata) Redact() {
	o.recurseRedact(&o.Self)
	o.recurseRedact(o.Next)
}

func (o *ResourceCollectionMetadata) recurseRedact(v interface{}) {
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

func (o ResourceCollectionMetadata) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ResourceCollectionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableResourceCollectionMetadata struct {
	value *ResourceCollectionMetadata
	isSet bool
}

func (v NullableResourceCollectionMetadata) Get() *ResourceCollectionMetadata {
	return v.value
}

func (v *NullableResourceCollectionMetadata) Set(val *ResourceCollectionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceCollectionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceCollectionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceCollectionMetadata(val *ResourceCollectionMetadata) *NullableResourceCollectionMetadata {
	return &NullableResourceCollectionMetadata{value: val, isSet: true}
}

func (v NullableResourceCollectionMetadata) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableResourceCollectionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}