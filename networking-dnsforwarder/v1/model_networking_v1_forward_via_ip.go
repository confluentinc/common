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

// NetworkingV1ForwardViaIp DNS Forwarder Configured via DNS Server IPs.
type NetworkingV1ForwardViaIp struct {
	// DNS Forwarder Configured via DNS Server IPs kind type.
	Kind string `json:"kind,omitempty"`
	// List of IP addresses of the DNS server
	DnsServerIps []string `json:"dns_server_ips,omitempty"`
}

// NewNetworkingV1ForwardViaIp instantiates a new NetworkingV1ForwardViaIp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1ForwardViaIp(kind string, dnsServerIps []string) *NetworkingV1ForwardViaIp {
	this := NetworkingV1ForwardViaIp{}
	this.Kind = kind
	this.DnsServerIps = dnsServerIps
	return &this
}

// NewNetworkingV1ForwardViaIpWithDefaults instantiates a new NetworkingV1ForwardViaIp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1ForwardViaIpWithDefaults() *NetworkingV1ForwardViaIp {
	this := NetworkingV1ForwardViaIp{}
	return &this
}

// GetKind returns the Kind field value
func (o *NetworkingV1ForwardViaIp) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1ForwardViaIp) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1ForwardViaIp) SetKind(v string) {
	o.Kind = v
}

// GetDnsServerIps returns the DnsServerIps field value
func (o *NetworkingV1ForwardViaIp) GetDnsServerIps() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DnsServerIps
}

// GetDnsServerIpsOk returns a tuple with the DnsServerIps field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1ForwardViaIp) GetDnsServerIpsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnsServerIps, true
}

// SetDnsServerIps sets field value
func (o *NetworkingV1ForwardViaIp) SetDnsServerIps(v []string) {
	o.DnsServerIps = v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1ForwardViaIp) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.DnsServerIps)
}

func (o *NetworkingV1ForwardViaIp) recurseRedact(v interface{}) {
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

func (o NetworkingV1ForwardViaIp) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1ForwardViaIp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["dns_server_ips"] = o.DnsServerIps
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1ForwardViaIp struct {
	value *NetworkingV1ForwardViaIp
	isSet bool
}

func (v NullableNetworkingV1ForwardViaIp) Get() *NetworkingV1ForwardViaIp {
	return v.value
}

func (v *NullableNetworkingV1ForwardViaIp) Set(val *NetworkingV1ForwardViaIp) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1ForwardViaIp) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1ForwardViaIp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1ForwardViaIp(val *NetworkingV1ForwardViaIp) *NullableNetworkingV1ForwardViaIp {
	return &NullableNetworkingV1ForwardViaIp{value: val, isSet: true}
}

func (v NullableNetworkingV1ForwardViaIp) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1ForwardViaIp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
