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
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// NetworkingV1AzureEgressPrivateLinkEndpoint Azure Private Endpoint.
type NetworkingV1AzureEgressPrivateLinkEndpoint struct {
	// AzureEgressPrivateLinkEndpoint kind.
	Kind string `json:"kind,omitempty"`
	// Resource ID of the Azure Private Link service.
	PrivateLinkServiceResourceId string `json:"private_link_service_resource_id,omitempty"`
	// Name of the subresource for the Private Endpoint to connect to.
	PrivateLinkSubresourceName *string `json:"private_link_subresource_name,omitempty"`
}

// NewNetworkingV1AzureEgressPrivateLinkEndpoint instantiates a new NetworkingV1AzureEgressPrivateLinkEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1AzureEgressPrivateLinkEndpoint(kind string, privateLinkServiceResourceId string) *NetworkingV1AzureEgressPrivateLinkEndpoint {
	this := NetworkingV1AzureEgressPrivateLinkEndpoint{}
	this.Kind = kind
	this.PrivateLinkServiceResourceId = privateLinkServiceResourceId
	return &this
}

// NewNetworkingV1AzureEgressPrivateLinkEndpointWithDefaults instantiates a new NetworkingV1AzureEgressPrivateLinkEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1AzureEgressPrivateLinkEndpointWithDefaults() *NetworkingV1AzureEgressPrivateLinkEndpoint {
	this := NetworkingV1AzureEgressPrivateLinkEndpoint{}
	return &this
}

// GetKind returns the Kind field value
func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) SetKind(v string) {
	o.Kind = v
}

// GetPrivateLinkServiceResourceId returns the PrivateLinkServiceResourceId field value
func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetPrivateLinkServiceResourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateLinkServiceResourceId
}

// GetPrivateLinkServiceResourceIdOk returns a tuple with the PrivateLinkServiceResourceId field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetPrivateLinkServiceResourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateLinkServiceResourceId, true
}

// SetPrivateLinkServiceResourceId sets field value
func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) SetPrivateLinkServiceResourceId(v string) {
	o.PrivateLinkServiceResourceId = v
}

// GetPrivateLinkSubresourceName returns the PrivateLinkSubresourceName field value if set, zero value otherwise.
func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetPrivateLinkSubresourceName() string {
	if o == nil || o.PrivateLinkSubresourceName == nil {
		var ret string
		return ret
	}
	return *o.PrivateLinkSubresourceName
}

// GetPrivateLinkSubresourceNameOk returns a tuple with the PrivateLinkSubresourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) GetPrivateLinkSubresourceNameOk() (*string, bool) {
	if o == nil || o.PrivateLinkSubresourceName == nil {
		return nil, false
	}
	return o.PrivateLinkSubresourceName, true
}

// HasPrivateLinkSubresourceName returns a boolean if a field has been set.
func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) HasPrivateLinkSubresourceName() bool {
	if o != nil && o.PrivateLinkSubresourceName != nil {
		return true
	}

	return false
}

// SetPrivateLinkSubresourceName gets a reference to the given string and assigns it to the PrivateLinkSubresourceName field.
func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) SetPrivateLinkSubresourceName(v string) {
	o.PrivateLinkSubresourceName = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.PrivateLinkServiceResourceId)
	o.recurseRedact(o.PrivateLinkSubresourceName)
}

func (o *NetworkingV1AzureEgressPrivateLinkEndpoint) recurseRedact(v interface{}) {
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

func (o NetworkingV1AzureEgressPrivateLinkEndpoint) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1AzureEgressPrivateLinkEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["private_link_service_resource_id"] = o.PrivateLinkServiceResourceId
	}
	if o.PrivateLinkSubresourceName != nil {
		toSerialize["private_link_subresource_name"] = o.PrivateLinkSubresourceName
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1AzureEgressPrivateLinkEndpoint struct {
	value *NetworkingV1AzureEgressPrivateLinkEndpoint
	isSet bool
}

func (v NullableNetworkingV1AzureEgressPrivateLinkEndpoint) Get() *NetworkingV1AzureEgressPrivateLinkEndpoint {
	return v.value
}

func (v *NullableNetworkingV1AzureEgressPrivateLinkEndpoint) Set(val *NetworkingV1AzureEgressPrivateLinkEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1AzureEgressPrivateLinkEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1AzureEgressPrivateLinkEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1AzureEgressPrivateLinkEndpoint(val *NetworkingV1AzureEgressPrivateLinkEndpoint) *NullableNetworkingV1AzureEgressPrivateLinkEndpoint {
	return &NullableNetworkingV1AzureEgressPrivateLinkEndpoint{value: val, isSet: true}
}

func (v NullableNetworkingV1AzureEgressPrivateLinkEndpoint) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1AzureEgressPrivateLinkEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
