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

// NetworkingV1AzurePrivateLinkAccess Azure PrivateLink access configuration.
type NetworkingV1AzurePrivateLinkAccess struct {
	// PrivateLink kind type.
	Kind string `json:"kind"`
	// Azure subscription to allow for PrivateLink access.
	Subscription string `json:"subscription"`
}

// NewNetworkingV1AzurePrivateLinkAccess instantiates a new NetworkingV1AzurePrivateLinkAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1AzurePrivateLinkAccess(kind string, subscription string) *NetworkingV1AzurePrivateLinkAccess {
	this := NetworkingV1AzurePrivateLinkAccess{}
	this.Kind = kind
	this.Subscription = subscription
	return &this
}

// NewNetworkingV1AzurePrivateLinkAccessWithDefaults instantiates a new NetworkingV1AzurePrivateLinkAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1AzurePrivateLinkAccessWithDefaults() *NetworkingV1AzurePrivateLinkAccess {
	this := NetworkingV1AzurePrivateLinkAccess{}
	return &this
}

// GetKind returns the Kind field value
func (o *NetworkingV1AzurePrivateLinkAccess) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzurePrivateLinkAccess) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1AzurePrivateLinkAccess) SetKind(v string) {
	o.Kind = v
}

// GetSubscription returns the Subscription field value
func (o *NetworkingV1AzurePrivateLinkAccess) GetSubscription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzurePrivateLinkAccess) GetSubscriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subscription, true
}

// SetSubscription sets field value
func (o *NetworkingV1AzurePrivateLinkAccess) SetSubscription(v string) {
	o.Subscription = v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1AzurePrivateLinkAccess) Redact() {
    o.recurseRedact(&o.Kind)
    o.recurseRedact(&o.Subscription)
}

func (o *NetworkingV1AzurePrivateLinkAccess) recurseRedact(v interface{}) {
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

func (o NetworkingV1AzurePrivateLinkAccess) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1AzurePrivateLinkAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["subscription"] = o.Subscription
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1AzurePrivateLinkAccess struct {
	value *NetworkingV1AzurePrivateLinkAccess
	isSet bool
}

func (v NullableNetworkingV1AzurePrivateLinkAccess) Get() *NetworkingV1AzurePrivateLinkAccess {
	return v.value
}

func (v *NullableNetworkingV1AzurePrivateLinkAccess) Set(val *NetworkingV1AzurePrivateLinkAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1AzurePrivateLinkAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1AzurePrivateLinkAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1AzurePrivateLinkAccess(val *NetworkingV1AzurePrivateLinkAccess) *NullableNetworkingV1AzurePrivateLinkAccess {
	return &NullableNetworkingV1AzurePrivateLinkAccess{value: val, isSet: true}
}

func (v NullableNetworkingV1AzurePrivateLinkAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1AzurePrivateLinkAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


