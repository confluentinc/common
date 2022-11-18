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
Cluster Management for Schema Registry API

Cluster Management for Schema Registry API

API version: 0.0.1-alpha1
Contact: data-governance@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

import (
	"reflect"
)

// MultipleSearchFilter Filter a collection by a string search for one or more values
type MultipleSearchFilter struct {
	Items []string
}

// NewMultipleSearchFilter instantiates a new MultipleSearchFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleSearchFilter() *MultipleSearchFilter {
	this := MultipleSearchFilter{}
	return &this
}

// NewMultipleSearchFilterWithDefaults instantiates a new MultipleSearchFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleSearchFilterWithDefaults() *MultipleSearchFilter {
	this := MultipleSearchFilter{}
	return &this
}

// Redact resets all sensitive fields to their zero value.
func (o *MultipleSearchFilter) Redact() {
}

func (o *MultipleSearchFilter) recurseRedact(v interface{}) {
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

func (o MultipleSearchFilter) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o MultipleSearchFilter) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return json.Marshal(toSerialize)
}

func (o *MultipleSearchFilter) UnmarshalJSON(bytes []byte) (err error) {
	return json.Unmarshal(bytes, &o.Items)
}

type NullableMultipleSearchFilter struct {
	value *MultipleSearchFilter
	isSet bool
}

func (v NullableMultipleSearchFilter) Get() *MultipleSearchFilter {
	return v.value
}

func (v *NullableMultipleSearchFilter) Set(val *MultipleSearchFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleSearchFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleSearchFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleSearchFilter(val *MultipleSearchFilter) *NullableMultipleSearchFilter {
	return &NullableMultipleSearchFilter{value: val, isSet: true}
}

func (v NullableMultipleSearchFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleSearchFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


