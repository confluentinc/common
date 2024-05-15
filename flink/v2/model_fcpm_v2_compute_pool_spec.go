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
Flink Compute Pool Management API

This is the Flink Compute Pool management API.

API version: 0.0.1
Contact: ksql-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// FcpmV2ComputePoolSpec The desired state of the Compute Pool
type FcpmV2ComputePoolSpec struct {
	// The name of the Flink compute pool.
	DisplayName *string `json:"display_name,omitempty"`
	// The cloud service provider that runs the compute pool.
	Cloud *string `json:"cloud,omitempty"`
	// Flink compute pools in the region provided will be able to use this identity pool
	Region *string `json:"region,omitempty"`
	// Maximum number of Confluent Flink Units (CFUs) that the Flink compute pool should auto-scale to.
	MaxCfu *int32 `json:"max_cfu,omitempty"`
	// The environment to which this belongs.
	Environment *GlobalObjectReference `json:"environment,omitempty"`
	// The network to which this belongs.
	Network *EnvScopedObjectReference `json:"network,omitempty"`
}

// NewFcpmV2ComputePoolSpec instantiates a new FcpmV2ComputePoolSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFcpmV2ComputePoolSpec() *FcpmV2ComputePoolSpec {
	this := FcpmV2ComputePoolSpec{}
	return &this
}

// NewFcpmV2ComputePoolSpecWithDefaults instantiates a new FcpmV2ComputePoolSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFcpmV2ComputePoolSpecWithDefaults() *FcpmV2ComputePoolSpec {
	this := FcpmV2ComputePoolSpec{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *FcpmV2ComputePoolSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpmV2ComputePoolSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *FcpmV2ComputePoolSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *FcpmV2ComputePoolSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *FcpmV2ComputePoolSpec) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpmV2ComputePoolSpec) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *FcpmV2ComputePoolSpec) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *FcpmV2ComputePoolSpec) SetCloud(v string) {
	o.Cloud = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *FcpmV2ComputePoolSpec) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpmV2ComputePoolSpec) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *FcpmV2ComputePoolSpec) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *FcpmV2ComputePoolSpec) SetRegion(v string) {
	o.Region = &v
}

// GetMaxCfu returns the MaxCfu field value if set, zero value otherwise.
func (o *FcpmV2ComputePoolSpec) GetMaxCfu() int32 {
	if o == nil || o.MaxCfu == nil {
		var ret int32
		return ret
	}
	return *o.MaxCfu
}

// GetMaxCfuOk returns a tuple with the MaxCfu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpmV2ComputePoolSpec) GetMaxCfuOk() (*int32, bool) {
	if o == nil || o.MaxCfu == nil {
		return nil, false
	}
	return o.MaxCfu, true
}

// HasMaxCfu returns a boolean if a field has been set.
func (o *FcpmV2ComputePoolSpec) HasMaxCfu() bool {
	if o != nil && o.MaxCfu != nil {
		return true
	}

	return false
}

// SetMaxCfu gets a reference to the given int32 and assigns it to the MaxCfu field.
func (o *FcpmV2ComputePoolSpec) SetMaxCfu(v int32) {
	o.MaxCfu = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *FcpmV2ComputePoolSpec) GetEnvironment() GlobalObjectReference {
	if o == nil || o.Environment == nil {
		var ret GlobalObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpmV2ComputePoolSpec) GetEnvironmentOk() (*GlobalObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *FcpmV2ComputePoolSpec) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given GlobalObjectReference and assigns it to the Environment field.
func (o *FcpmV2ComputePoolSpec) SetEnvironment(v GlobalObjectReference) {
	o.Environment = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *FcpmV2ComputePoolSpec) GetNetwork() EnvScopedObjectReference {
	if o == nil || o.Network == nil {
		var ret EnvScopedObjectReference
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FcpmV2ComputePoolSpec) GetNetworkOk() (*EnvScopedObjectReference, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *FcpmV2ComputePoolSpec) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given EnvScopedObjectReference and assigns it to the Network field.
func (o *FcpmV2ComputePoolSpec) SetNetwork(v EnvScopedObjectReference) {
	o.Network = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *FcpmV2ComputePoolSpec) Redact() {
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Cloud)
	o.recurseRedact(o.Region)
	o.recurseRedact(o.MaxCfu)
	o.recurseRedact(o.Environment)
	o.recurseRedact(o.Network)
}

func (o *FcpmV2ComputePoolSpec) recurseRedact(v interface{}) {
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

func (o FcpmV2ComputePoolSpec) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o FcpmV2ComputePoolSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.MaxCfu != nil {
		toSerialize["max_cfu"] = o.MaxCfu
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableFcpmV2ComputePoolSpec struct {
	value *FcpmV2ComputePoolSpec
	isSet bool
}

func (v NullableFcpmV2ComputePoolSpec) Get() *FcpmV2ComputePoolSpec {
	return v.value
}

func (v *NullableFcpmV2ComputePoolSpec) Set(val *FcpmV2ComputePoolSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableFcpmV2ComputePoolSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableFcpmV2ComputePoolSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFcpmV2ComputePoolSpec(val *FcpmV2ComputePoolSpec) *NullableFcpmV2ComputePoolSpec {
	return &NullableFcpmV2ComputePoolSpec{value: val, isSet: true}
}

func (v NullableFcpmV2ComputePoolSpec) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableFcpmV2ComputePoolSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
