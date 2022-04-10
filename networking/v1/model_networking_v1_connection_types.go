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
Network API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha1
Contact: cire-traffic@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// NetworkingV1ConnectionTypes The connection types requested for use with the network.
type NetworkingV1ConnectionTypes struct {
	Items []string
}

// NewNetworkingV1ConnectionTypes instantiates a new NetworkingV1ConnectionTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1ConnectionTypes() *NetworkingV1ConnectionTypes {
	this := NetworkingV1ConnectionTypes{}
	return &this
}

// NewNetworkingV1ConnectionTypesWithDefaults instantiates a new NetworkingV1ConnectionTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1ConnectionTypesWithDefaults() *NetworkingV1ConnectionTypes {
	this := NetworkingV1ConnectionTypes{}
	return &this
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1ConnectionTypes) Redact() {
}

func (o *NetworkingV1ConnectionTypes) recurseRedact(v interface{}) {
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

func (o NetworkingV1ConnectionTypes) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1ConnectionTypes) MarshalJSON() ([]byte, error) {
	toSerialize := make([]interface{}, len(o.Items))
	for i, item := range o.Items {
		toSerialize[i] = item
	}
	return json.Marshal(toSerialize)
}

func (o *NetworkingV1ConnectionTypes) UnmarshalJSON(bytes []byte) (err error) {
	return json.Unmarshal(bytes, &o.Items)
}

type NullableNetworkingV1ConnectionTypes struct {
	value *NetworkingV1ConnectionTypes
	isSet bool
}

func (v NullableNetworkingV1ConnectionTypes) Get() *NetworkingV1ConnectionTypes {
	return v.value
}

func (v *NullableNetworkingV1ConnectionTypes) Set(val *NetworkingV1ConnectionTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1ConnectionTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1ConnectionTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1ConnectionTypes(val *NetworkingV1ConnectionTypes) *NullableNetworkingV1ConnectionTypes {
	return &NullableNetworkingV1ConnectionTypes{value: val, isSet: true}
}

func (v NullableNetworkingV1ConnectionTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1ConnectionTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


