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

// NetworkingV1DnsConfig The network DNS config
type NetworkingV1DnsConfig struct {
	// Network DNS resolution
	Resolution string `json:"resolution"`
}

// NewNetworkingV1DnsConfig instantiates a new NetworkingV1DnsConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1DnsConfig(resolution string) *NetworkingV1DnsConfig {
	this := NetworkingV1DnsConfig{}
	this.Resolution = resolution
	return &this
}

// NewNetworkingV1DnsConfigWithDefaults instantiates a new NetworkingV1DnsConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1DnsConfigWithDefaults() *NetworkingV1DnsConfig {
	this := NetworkingV1DnsConfig{}
	return &this
}

// GetResolution returns the Resolution field value
func (o *NetworkingV1DnsConfig) GetResolution() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resolution
}

// GetResolutionOk returns a tuple with the Resolution field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1DnsConfig) GetResolutionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Resolution, true
}

// SetResolution sets field value
func (o *NetworkingV1DnsConfig) SetResolution(v string) {
	o.Resolution = v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1DnsConfig) Redact() {
    o.recurseRedact(&o.Resolution)
}

func (o *NetworkingV1DnsConfig) recurseRedact(v interface{}) {
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

func (o NetworkingV1DnsConfig) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1DnsConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["resolution"] = o.Resolution
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1DnsConfig struct {
	value *NetworkingV1DnsConfig
	isSet bool
}

func (v NullableNetworkingV1DnsConfig) Get() *NetworkingV1DnsConfig {
	return v.value
}

func (v *NullableNetworkingV1DnsConfig) Set(val *NetworkingV1DnsConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1DnsConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1DnsConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1DnsConfig(val *NetworkingV1DnsConfig) *NullableNetworkingV1DnsConfig {
	return &NullableNetworkingV1DnsConfig{value: val, isSet: true}
}

func (v NullableNetworkingV1DnsConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1DnsConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


