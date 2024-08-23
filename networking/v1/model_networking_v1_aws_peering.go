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

// NetworkingV1AwsPeering AWS VPC Peering.
type NetworkingV1AwsPeering struct {
	// Peering kind type.
	Kind string `json:"kind,omitempty"`
	// The AWS account ID associated with the VPC you are peering with Confluent Cloud network.
	Account string `json:"account,omitempty"`
	// The VPC ID you are peering with Confluent Cloud network.
	Vpc string `json:"vpc,omitempty"`
	// The [CIDR blocks](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing) of the VPC you are peering with Confluent Cloud network. This is used by Confluent Cloud network to route traffic back to your network. The CIDR block must be a private range and cannot overlap with the Confluent Cloud CIDR block.
	Routes []string `json:"routes,omitempty"`
	// The region of the VPC you are peering with Confluent Cloud network.
	CustomerRegion string `json:"customer_region,omitempty"`
}

// NewNetworkingV1AwsPeering instantiates a new NetworkingV1AwsPeering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1AwsPeering(kind string, account string, vpc string, routes []string, customerRegion string) *NetworkingV1AwsPeering {
	this := NetworkingV1AwsPeering{}
	this.Kind = kind
	this.Account = account
	this.Vpc = vpc
	this.Routes = routes
	this.CustomerRegion = customerRegion
	return &this
}

// NewNetworkingV1AwsPeeringWithDefaults instantiates a new NetworkingV1AwsPeering object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1AwsPeeringWithDefaults() *NetworkingV1AwsPeering {
	this := NetworkingV1AwsPeering{}
	return &this
}

// GetKind returns the Kind field value
func (o *NetworkingV1AwsPeering) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsPeering) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1AwsPeering) SetKind(v string) {
	o.Kind = v
}

// GetAccount returns the Account field value
func (o *NetworkingV1AwsPeering) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsPeering) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *NetworkingV1AwsPeering) SetAccount(v string) {
	o.Account = v
}

// GetVpc returns the Vpc field value
func (o *NetworkingV1AwsPeering) GetVpc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vpc
}

// GetVpcOk returns a tuple with the Vpc field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsPeering) GetVpcOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vpc, true
}

// SetVpc sets field value
func (o *NetworkingV1AwsPeering) SetVpc(v string) {
	o.Vpc = v
}

// GetRoutes returns the Routes field value
func (o *NetworkingV1AwsPeering) GetRoutes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Routes
}

// GetRoutesOk returns a tuple with the Routes field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsPeering) GetRoutesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Routes, true
}

// SetRoutes sets field value
func (o *NetworkingV1AwsPeering) SetRoutes(v []string) {
	o.Routes = v
}

// GetCustomerRegion returns the CustomerRegion field value
func (o *NetworkingV1AwsPeering) GetCustomerRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerRegion
}

// GetCustomerRegionOk returns a tuple with the CustomerRegion field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AwsPeering) GetCustomerRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerRegion, true
}

// SetCustomerRegion sets field value
func (o *NetworkingV1AwsPeering) SetCustomerRegion(v string) {
	o.CustomerRegion = v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1AwsPeering) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Account)
	o.recurseRedact(&o.Vpc)
	o.recurseRedact(&o.Routes)
	o.recurseRedact(&o.CustomerRegion)
}

func (o *NetworkingV1AwsPeering) recurseRedact(v interface{}) {
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

func (o NetworkingV1AwsPeering) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1AwsPeering) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["account"] = o.Account
	}
	if true {
		toSerialize["vpc"] = o.Vpc
	}
	if true {
		toSerialize["routes"] = o.Routes
	}
	if true {
		toSerialize["customer_region"] = o.CustomerRegion
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1AwsPeering struct {
	value *NetworkingV1AwsPeering
	isSet bool
}

func (v NullableNetworkingV1AwsPeering) Get() *NetworkingV1AwsPeering {
	return v.value
}

func (v *NullableNetworkingV1AwsPeering) Set(val *NetworkingV1AwsPeering) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1AwsPeering) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1AwsPeering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1AwsPeering(val *NetworkingV1AwsPeering) *NullableNetworkingV1AwsPeering {
	return &NullableNetworkingV1AwsPeering{value: val, isSet: true}
}

func (v NullableNetworkingV1AwsPeering) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1AwsPeering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}