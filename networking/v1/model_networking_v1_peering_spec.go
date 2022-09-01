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

// NetworkingV1PeeringSpec The desired state of the Peering
type NetworkingV1PeeringSpec struct {
	// The name of the peering
	DisplayName *string `json:"display_name,omitempty"`
	// The cloud-specific peering details.
	Cloud *NetworkingV1PeeringSpecCloudOneOf `json:"cloud,omitempty"`
	// The environment to which this belongs.
	Environment *ObjectReference `json:"environment,omitempty"`
	// The network to which this belongs.
	Network *ObjectReference `json:"network,omitempty"`
}

// NewNetworkingV1PeeringSpec instantiates a new NetworkingV1PeeringSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1PeeringSpec() *NetworkingV1PeeringSpec {
	this := NetworkingV1PeeringSpec{}
	return &this
}

// NewNetworkingV1PeeringSpecWithDefaults instantiates a new NetworkingV1PeeringSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1PeeringSpecWithDefaults() *NetworkingV1PeeringSpec {
	this := NetworkingV1PeeringSpec{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkingV1PeeringSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PeeringSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkingV1PeeringSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkingV1PeeringSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *NetworkingV1PeeringSpec) GetCloud() NetworkingV1PeeringSpecCloudOneOf {
	if o == nil || o.Cloud == nil {
		var ret NetworkingV1PeeringSpecCloudOneOf
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PeeringSpec) GetCloudOk() (*NetworkingV1PeeringSpecCloudOneOf, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *NetworkingV1PeeringSpec) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given NetworkingV1PeeringSpecCloudOneOf and assigns it to the Cloud field.
func (o *NetworkingV1PeeringSpec) SetCloud(v NetworkingV1PeeringSpecCloudOneOf) {
	o.Cloud = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NetworkingV1PeeringSpec) GetEnvironment() ObjectReference {
	if o == nil || o.Environment == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PeeringSpec) GetEnvironmentOk() (*ObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NetworkingV1PeeringSpec) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectReference and assigns it to the Environment field.
func (o *NetworkingV1PeeringSpec) SetEnvironment(v ObjectReference) {
	o.Environment = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *NetworkingV1PeeringSpec) GetNetwork() ObjectReference {
	if o == nil || o.Network == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PeeringSpec) GetNetworkOk() (*ObjectReference, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *NetworkingV1PeeringSpec) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ObjectReference and assigns it to the Network field.
func (o *NetworkingV1PeeringSpec) SetNetwork(v ObjectReference) {
	o.Network = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1PeeringSpec) Redact() {
    o.recurseRedact(o.DisplayName)
    o.recurseRedact(o.Cloud)
    o.recurseRedact(o.Environment)
    o.recurseRedact(o.Network)
}

func (o *NetworkingV1PeeringSpec) recurseRedact(v interface{}) {
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

func (o NetworkingV1PeeringSpec) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1PeeringSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1PeeringSpec struct {
	value *NetworkingV1PeeringSpec
	isSet bool
}

func (v NullableNetworkingV1PeeringSpec) Get() *NetworkingV1PeeringSpec {
	return v.value
}

func (v *NullableNetworkingV1PeeringSpec) Set(val *NetworkingV1PeeringSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1PeeringSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1PeeringSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1PeeringSpec(val *NetworkingV1PeeringSpec) *NullableNetworkingV1PeeringSpec {
	return &NullableNetworkingV1PeeringSpec{value: val, isSet: true}
}

func (v NullableNetworkingV1PeeringSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1PeeringSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


