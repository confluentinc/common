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

// NetworkingV1DnsRecordSpec The desired state of the Dns Record
type NetworkingV1DnsRecordSpec struct {
	// The name of the DNS record.
	DisplayName *string `json:"display_name,omitempty"`
	// The fully qualified domain name of the DNS record.
	Domain *string `json:"domain,omitempty"`
	// The config of the DNS record.
	Config *NetworkingV1DnsRecordSpecConfigOneOf `json:"config,omitempty"`
	// The environment to which this belongs.
	Environment *ObjectReference `json:"environment,omitempty"`
	// The gateway to which this belongs.
	Gateway *EnvScopedObjectReference `json:"gateway,omitempty"`
}

// NewNetworkingV1DnsRecordSpec instantiates a new NetworkingV1DnsRecordSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1DnsRecordSpec() *NetworkingV1DnsRecordSpec {
	this := NetworkingV1DnsRecordSpec{}
	return &this
}

// NewNetworkingV1DnsRecordSpecWithDefaults instantiates a new NetworkingV1DnsRecordSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1DnsRecordSpecWithDefaults() *NetworkingV1DnsRecordSpec {
	this := NetworkingV1DnsRecordSpec{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkingV1DnsRecordSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1DnsRecordSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkingV1DnsRecordSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkingV1DnsRecordSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *NetworkingV1DnsRecordSpec) GetDomain() string {
	if o == nil || o.Domain == nil {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1DnsRecordSpec) GetDomainOk() (*string, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *NetworkingV1DnsRecordSpec) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *NetworkingV1DnsRecordSpec) SetDomain(v string) {
	o.Domain = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *NetworkingV1DnsRecordSpec) GetConfig() NetworkingV1DnsRecordSpecConfigOneOf {
	if o == nil || o.Config == nil {
		var ret NetworkingV1DnsRecordSpecConfigOneOf
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1DnsRecordSpec) GetConfigOk() (*NetworkingV1DnsRecordSpecConfigOneOf, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *NetworkingV1DnsRecordSpec) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given NetworkingV1DnsRecordSpecConfigOneOf and assigns it to the Config field.
func (o *NetworkingV1DnsRecordSpec) SetConfig(v NetworkingV1DnsRecordSpecConfigOneOf) {
	o.Config = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NetworkingV1DnsRecordSpec) GetEnvironment() ObjectReference {
	if o == nil || o.Environment == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1DnsRecordSpec) GetEnvironmentOk() (*ObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NetworkingV1DnsRecordSpec) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectReference and assigns it to the Environment field.
func (o *NetworkingV1DnsRecordSpec) SetEnvironment(v ObjectReference) {
	o.Environment = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *NetworkingV1DnsRecordSpec) GetGateway() EnvScopedObjectReference {
	if o == nil || o.Gateway == nil {
		var ret EnvScopedObjectReference
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1DnsRecordSpec) GetGatewayOk() (*EnvScopedObjectReference, bool) {
	if o == nil || o.Gateway == nil {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *NetworkingV1DnsRecordSpec) HasGateway() bool {
	if o != nil && o.Gateway != nil {
		return true
	}

	return false
}

// SetGateway gets a reference to the given EnvScopedObjectReference and assigns it to the Gateway field.
func (o *NetworkingV1DnsRecordSpec) SetGateway(v EnvScopedObjectReference) {
	o.Gateway = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1DnsRecordSpec) Redact() {
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Domain)
	o.recurseRedact(o.Config)
	o.recurseRedact(o.Environment)
	o.recurseRedact(o.Gateway)
}

func (o *NetworkingV1DnsRecordSpec) recurseRedact(v interface{}) {
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

func (o NetworkingV1DnsRecordSpec) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1DnsRecordSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Gateway != nil {
		toSerialize["gateway"] = o.Gateway
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1DnsRecordSpec struct {
	value *NetworkingV1DnsRecordSpec
	isSet bool
}

func (v NullableNetworkingV1DnsRecordSpec) Get() *NetworkingV1DnsRecordSpec {
	return v.value
}

func (v *NullableNetworkingV1DnsRecordSpec) Set(val *NetworkingV1DnsRecordSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1DnsRecordSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1DnsRecordSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1DnsRecordSpec(val *NetworkingV1DnsRecordSpec) *NullableNetworkingV1DnsRecordSpec {
	return &NullableNetworkingV1DnsRecordSpec{value: val, isSet: true}
}

func (v NullableNetworkingV1DnsRecordSpec) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1DnsRecordSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
