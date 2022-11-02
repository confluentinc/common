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

// StreamGovernanceV2ClusterSpec The desired state of the Cluster
type StreamGovernanceV2ClusterSpec struct {
	// The cluster name.
	DisplayName *string `json:"display_name,omitempty"`
	// The billing package.  Note: Clusters can be upgraded from ESSENTIALS to ADVANCED, but cannot be downgraded from ADVANCED to ESSENTIALS.
	Package *string `json:"package,omitempty"`
	// The cluster HTTP request URL.
	HttpEndpoint *string `json:"http_endpoint,omitempty"`
	// The environment to which this belongs.
	Environment *GlobalObjectReference `json:"environment,omitempty"`
	// The region to which this belongs.
	Region *GlobalObjectReference `json:"region,omitempty"`
}

// NewStreamGovernanceV2ClusterSpec instantiates a new StreamGovernanceV2ClusterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamGovernanceV2ClusterSpec() *StreamGovernanceV2ClusterSpec {
	this := StreamGovernanceV2ClusterSpec{}
	return &this
}

// NewStreamGovernanceV2ClusterSpecWithDefaults instantiates a new StreamGovernanceV2ClusterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamGovernanceV2ClusterSpecWithDefaults() *StreamGovernanceV2ClusterSpec {
	this := StreamGovernanceV2ClusterSpec{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *StreamGovernanceV2ClusterSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGovernanceV2ClusterSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *StreamGovernanceV2ClusterSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *StreamGovernanceV2ClusterSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *StreamGovernanceV2ClusterSpec) GetPackage() string {
	if o == nil || o.Package == nil {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGovernanceV2ClusterSpec) GetPackageOk() (*string, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *StreamGovernanceV2ClusterSpec) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *StreamGovernanceV2ClusterSpec) SetPackage(v string) {
	o.Package = &v
}

// GetHttpEndpoint returns the HttpEndpoint field value if set, zero value otherwise.
func (o *StreamGovernanceV2ClusterSpec) GetHttpEndpoint() string {
	if o == nil || o.HttpEndpoint == nil {
		var ret string
		return ret
	}
	return *o.HttpEndpoint
}

// GetHttpEndpointOk returns a tuple with the HttpEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGovernanceV2ClusterSpec) GetHttpEndpointOk() (*string, bool) {
	if o == nil || o.HttpEndpoint == nil {
		return nil, false
	}
	return o.HttpEndpoint, true
}

// HasHttpEndpoint returns a boolean if a field has been set.
func (o *StreamGovernanceV2ClusterSpec) HasHttpEndpoint() bool {
	if o != nil && o.HttpEndpoint != nil {
		return true
	}

	return false
}

// SetHttpEndpoint gets a reference to the given string and assigns it to the HttpEndpoint field.
func (o *StreamGovernanceV2ClusterSpec) SetHttpEndpoint(v string) {
	o.HttpEndpoint = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *StreamGovernanceV2ClusterSpec) GetEnvironment() GlobalObjectReference {
	if o == nil || o.Environment == nil {
		var ret GlobalObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGovernanceV2ClusterSpec) GetEnvironmentOk() (*GlobalObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *StreamGovernanceV2ClusterSpec) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given GlobalObjectReference and assigns it to the Environment field.
func (o *StreamGovernanceV2ClusterSpec) SetEnvironment(v GlobalObjectReference) {
	o.Environment = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *StreamGovernanceV2ClusterSpec) GetRegion() GlobalObjectReference {
	if o == nil || o.Region == nil {
		var ret GlobalObjectReference
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGovernanceV2ClusterSpec) GetRegionOk() (*GlobalObjectReference, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StreamGovernanceV2ClusterSpec) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given GlobalObjectReference and assigns it to the Region field.
func (o *StreamGovernanceV2ClusterSpec) SetRegion(v GlobalObjectReference) {
	o.Region = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *StreamGovernanceV2ClusterSpec) Redact() {
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Package)
	o.recurseRedact(o.HttpEndpoint)
	o.recurseRedact(o.Environment)
	o.recurseRedact(o.Region)
}

func (o *StreamGovernanceV2ClusterSpec) recurseRedact(v interface{}) {
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

func (o StreamGovernanceV2ClusterSpec) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o StreamGovernanceV2ClusterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Package != nil {
		toSerialize["package"] = o.Package
	}
	if o.HttpEndpoint != nil {
		toSerialize["http_endpoint"] = o.HttpEndpoint
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableStreamGovernanceV2ClusterSpec struct {
	value *StreamGovernanceV2ClusterSpec
	isSet bool
}

func (v NullableStreamGovernanceV2ClusterSpec) Get() *StreamGovernanceV2ClusterSpec {
	return v.value
}

func (v *NullableStreamGovernanceV2ClusterSpec) Set(val *StreamGovernanceV2ClusterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamGovernanceV2ClusterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamGovernanceV2ClusterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamGovernanceV2ClusterSpec(val *StreamGovernanceV2ClusterSpec) *NullableStreamGovernanceV2ClusterSpec {
	return &NullableStreamGovernanceV2ClusterSpec{value: val, isSet: true}
}

func (v NullableStreamGovernanceV2ClusterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamGovernanceV2ClusterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
