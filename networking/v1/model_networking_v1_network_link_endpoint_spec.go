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

Network API

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

// NetworkingV1NetworkLinkEndpointSpec The desired state of the Network Link Endpoint
type NetworkingV1NetworkLinkEndpointSpec struct {
	// The name of the network link endpoint
	DisplayName *string `json:"display_name,omitempty"`
	// The description of the network link endpoint
	Description *string `json:"description,omitempty"`
	// The environment to which this belongs.
	Environment *GlobalObjectReference `json:"environment,omitempty"`
	// The network to which this belongs.
	Network *EnvScopedObjectReference `json:"network,omitempty"`
	// The network_link_service to which this belongs.
	NetworkLinkService *EnvScopedObjectReference `json:"network_link_service,omitempty"`
}

// NewNetworkingV1NetworkLinkEndpointSpec instantiates a new NetworkingV1NetworkLinkEndpointSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1NetworkLinkEndpointSpec() *NetworkingV1NetworkLinkEndpointSpec {
	this := NetworkingV1NetworkLinkEndpointSpec{}
	return &this
}

// NewNetworkingV1NetworkLinkEndpointSpecWithDefaults instantiates a new NetworkingV1NetworkLinkEndpointSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1NetworkLinkEndpointSpecWithDefaults() *NetworkingV1NetworkLinkEndpointSpec {
	this := NetworkingV1NetworkLinkEndpointSpec{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkEndpointSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkEndpointSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkEndpointSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkingV1NetworkLinkEndpointSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkEndpointSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkEndpointSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkEndpointSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworkingV1NetworkLinkEndpointSpec) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkEndpointSpec) GetEnvironment() GlobalObjectReference {
	if o == nil || o.Environment == nil {
		var ret GlobalObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkEndpointSpec) GetEnvironmentOk() (*GlobalObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkEndpointSpec) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given GlobalObjectReference and assigns it to the Environment field.
func (o *NetworkingV1NetworkLinkEndpointSpec) SetEnvironment(v GlobalObjectReference) {
	o.Environment = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkEndpointSpec) GetNetwork() EnvScopedObjectReference {
	if o == nil || o.Network == nil {
		var ret EnvScopedObjectReference
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkEndpointSpec) GetNetworkOk() (*EnvScopedObjectReference, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkEndpointSpec) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given EnvScopedObjectReference and assigns it to the Network field.
func (o *NetworkingV1NetworkLinkEndpointSpec) SetNetwork(v EnvScopedObjectReference) {
	o.Network = &v
}

// GetNetworkLinkService returns the NetworkLinkService field value if set, zero value otherwise.
func (o *NetworkingV1NetworkLinkEndpointSpec) GetNetworkLinkService() EnvScopedObjectReference {
	if o == nil || o.NetworkLinkService == nil {
		var ret EnvScopedObjectReference
		return ret
	}
	return *o.NetworkLinkService
}

// GetNetworkLinkServiceOk returns a tuple with the NetworkLinkService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1NetworkLinkEndpointSpec) GetNetworkLinkServiceOk() (*EnvScopedObjectReference, bool) {
	if o == nil || o.NetworkLinkService == nil {
		return nil, false
	}
	return o.NetworkLinkService, true
}

// HasNetworkLinkService returns a boolean if a field has been set.
func (o *NetworkingV1NetworkLinkEndpointSpec) HasNetworkLinkService() bool {
	if o != nil && o.NetworkLinkService != nil {
		return true
	}

	return false
}

// SetNetworkLinkService gets a reference to the given EnvScopedObjectReference and assigns it to the NetworkLinkService field.
func (o *NetworkingV1NetworkLinkEndpointSpec) SetNetworkLinkService(v EnvScopedObjectReference) {
	o.NetworkLinkService = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1NetworkLinkEndpointSpec) Redact() {
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Description)
	o.recurseRedact(o.Environment)
	o.recurseRedact(o.Network)
	o.recurseRedact(o.NetworkLinkService)
}

func (o *NetworkingV1NetworkLinkEndpointSpec) recurseRedact(v interface{}) {
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

func (o NetworkingV1NetworkLinkEndpointSpec) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1NetworkLinkEndpointSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	if o.NetworkLinkService != nil {
		toSerialize["network_link_service"] = o.NetworkLinkService
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1NetworkLinkEndpointSpec struct {
	value *NetworkingV1NetworkLinkEndpointSpec
	isSet bool
}

func (v NullableNetworkingV1NetworkLinkEndpointSpec) Get() *NetworkingV1NetworkLinkEndpointSpec {
	return v.value
}

func (v *NullableNetworkingV1NetworkLinkEndpointSpec) Set(val *NetworkingV1NetworkLinkEndpointSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1NetworkLinkEndpointSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1NetworkLinkEndpointSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1NetworkLinkEndpointSpec(val *NetworkingV1NetworkLinkEndpointSpec) *NullableNetworkingV1NetworkLinkEndpointSpec {
	return &NullableNetworkingV1NetworkLinkEndpointSpec{value: val, isSet: true}
}

func (v NullableNetworkingV1NetworkLinkEndpointSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1NetworkLinkEndpointSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
