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
Cluster Management for Stream Governance API

Cluster Management for Stream Governance API

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

// StreamGovernanceV2RegionSpec The desired state of the Region
type StreamGovernanceV2RegionSpec struct {
	// The display name.
	DisplayName *string `json:"display_name,omitempty"`
	// The cloud service provider that hosts the region.
	Cloud *string `json:"cloud,omitempty"`
	// The region name.
	RegionName *string `json:"region_name,omitempty"`
	// List of Stream Governance packages allowing placement in this region.
	Packages *[]string `json:"packages,omitempty"`
}

// NewStreamGovernanceV2RegionSpec instantiates a new StreamGovernanceV2RegionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamGovernanceV2RegionSpec() *StreamGovernanceV2RegionSpec {
	this := StreamGovernanceV2RegionSpec{}
	return &this
}

// NewStreamGovernanceV2RegionSpecWithDefaults instantiates a new StreamGovernanceV2RegionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamGovernanceV2RegionSpecWithDefaults() *StreamGovernanceV2RegionSpec {
	this := StreamGovernanceV2RegionSpec{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *StreamGovernanceV2RegionSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGovernanceV2RegionSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *StreamGovernanceV2RegionSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *StreamGovernanceV2RegionSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *StreamGovernanceV2RegionSpec) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGovernanceV2RegionSpec) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *StreamGovernanceV2RegionSpec) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *StreamGovernanceV2RegionSpec) SetCloud(v string) {
	o.Cloud = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *StreamGovernanceV2RegionSpec) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGovernanceV2RegionSpec) GetRegionNameOk() (*string, bool) {
	if o == nil || o.RegionName == nil {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *StreamGovernanceV2RegionSpec) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *StreamGovernanceV2RegionSpec) SetRegionName(v string) {
	o.RegionName = &v
}

// GetPackages returns the Packages field value if set, zero value otherwise.
func (o *StreamGovernanceV2RegionSpec) GetPackages() []string {
	if o == nil || o.Packages == nil {
		var ret []string
		return ret
	}
	return *o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGovernanceV2RegionSpec) GetPackagesOk() (*[]string, bool) {
	if o == nil || o.Packages == nil {
		return nil, false
	}
	return o.Packages, true
}

// HasPackages returns a boolean if a field has been set.
func (o *StreamGovernanceV2RegionSpec) HasPackages() bool {
	if o != nil && o.Packages != nil {
		return true
	}

	return false
}

// SetPackages gets a reference to the given []string and assigns it to the Packages field.
func (o *StreamGovernanceV2RegionSpec) SetPackages(v []string) {
	o.Packages = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *StreamGovernanceV2RegionSpec) Redact() {
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Cloud)
	o.recurseRedact(o.RegionName)
	o.recurseRedact(o.Packages)
}

func (o *StreamGovernanceV2RegionSpec) recurseRedact(v interface{}) {
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

func (o StreamGovernanceV2RegionSpec) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o StreamGovernanceV2RegionSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.RegionName != nil {
		toSerialize["region_name"] = o.RegionName
	}
	if o.Packages != nil {
		toSerialize["packages"] = o.Packages
	}
	return json.Marshal(toSerialize)
}

type NullableStreamGovernanceV2RegionSpec struct {
	value *StreamGovernanceV2RegionSpec
	isSet bool
}

func (v NullableStreamGovernanceV2RegionSpec) Get() *StreamGovernanceV2RegionSpec {
	return v.value
}

func (v *NullableStreamGovernanceV2RegionSpec) Set(val *StreamGovernanceV2RegionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamGovernanceV2RegionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamGovernanceV2RegionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamGovernanceV2RegionSpec(val *StreamGovernanceV2RegionSpec) *NullableStreamGovernanceV2RegionSpec {
	return &NullableStreamGovernanceV2RegionSpec{value: val, isSet: true}
}

func (v NullableStreamGovernanceV2RegionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamGovernanceV2RegionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
