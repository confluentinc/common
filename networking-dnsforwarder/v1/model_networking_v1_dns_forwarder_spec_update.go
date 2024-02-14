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

// NetworkingV1DnsForwarderSpecUpdate The desired state of the Dns Forwarder
type NetworkingV1DnsForwarderSpecUpdate struct {
	// The name of the DNS forwarder
	DisplayName *string `json:"display_name,omitempty"`
	// List of domains for the DNS forwarder to use
	Domains *[]string `json:"domains,omitempty"`
	// The specific details of different kinds of configuration for DNS Forwarder.
	Config *NetworkingV1DnsForwarderSpecUpdateConfigOneOf `json:"config,omitempty"`
	// The environment to which this belongs.
	Environment *ObjectReference `json:"environment,omitempty"`
}

// NewNetworkingV1DnsForwarderSpecUpdate instantiates a new NetworkingV1DnsForwarderSpecUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1DnsForwarderSpecUpdate() *NetworkingV1DnsForwarderSpecUpdate {
	this := NetworkingV1DnsForwarderSpecUpdate{}
	return &this
}

// NewNetworkingV1DnsForwarderSpecUpdateWithDefaults instantiates a new NetworkingV1DnsForwarderSpecUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1DnsForwarderSpecUpdateWithDefaults() *NetworkingV1DnsForwarderSpecUpdate {
	this := NetworkingV1DnsForwarderSpecUpdate{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkingV1DnsForwarderSpecUpdate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1DnsForwarderSpecUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkingV1DnsForwarderSpecUpdate) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkingV1DnsForwarderSpecUpdate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *NetworkingV1DnsForwarderSpecUpdate) GetDomains() []string {
	if o == nil || o.Domains == nil {
		var ret []string
		return ret
	}
	return *o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1DnsForwarderSpecUpdate) GetDomainsOk() (*[]string, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *NetworkingV1DnsForwarderSpecUpdate) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []string and assigns it to the Domains field.
func (o *NetworkingV1DnsForwarderSpecUpdate) SetDomains(v []string) {
	o.Domains = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *NetworkingV1DnsForwarderSpecUpdate) GetConfig() NetworkingV1DnsForwarderSpecUpdateConfigOneOf {
	if o == nil || o.Config == nil {
		var ret NetworkingV1DnsForwarderSpecUpdateConfigOneOf
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1DnsForwarderSpecUpdate) GetConfigOk() (*NetworkingV1DnsForwarderSpecUpdateConfigOneOf, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *NetworkingV1DnsForwarderSpecUpdate) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given NetworkingV1DnsForwarderSpecUpdateConfigOneOf and assigns it to the Config field.
func (o *NetworkingV1DnsForwarderSpecUpdate) SetConfig(v NetworkingV1DnsForwarderSpecUpdateConfigOneOf) {
	o.Config = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NetworkingV1DnsForwarderSpecUpdate) GetEnvironment() ObjectReference {
	if o == nil || o.Environment == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1DnsForwarderSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NetworkingV1DnsForwarderSpecUpdate) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectReference and assigns it to the Environment field.
func (o *NetworkingV1DnsForwarderSpecUpdate) SetEnvironment(v ObjectReference) {
	o.Environment = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1DnsForwarderSpecUpdate) Redact() {
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Domains)
	o.recurseRedact(o.Config)
	o.recurseRedact(o.Environment)
}

func (o *NetworkingV1DnsForwarderSpecUpdate) recurseRedact(v interface{}) {
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

func (o NetworkingV1DnsForwarderSpecUpdate) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1DnsForwarderSpecUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1DnsForwarderSpecUpdate struct {
	value *NetworkingV1DnsForwarderSpecUpdate
	isSet bool
}

func (v NullableNetworkingV1DnsForwarderSpecUpdate) Get() *NetworkingV1DnsForwarderSpecUpdate {
	return v.value
}

func (v *NullableNetworkingV1DnsForwarderSpecUpdate) Set(val *NetworkingV1DnsForwarderSpecUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1DnsForwarderSpecUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1DnsForwarderSpecUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1DnsForwarderSpecUpdate(val *NetworkingV1DnsForwarderSpecUpdate) *NullableNetworkingV1DnsForwarderSpecUpdate {
	return &NullableNetworkingV1DnsForwarderSpecUpdate{value: val, isSet: true}
}

func (v NullableNetworkingV1DnsForwarderSpecUpdate) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1DnsForwarderSpecUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
