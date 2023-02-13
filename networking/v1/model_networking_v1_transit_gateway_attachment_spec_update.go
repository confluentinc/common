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

// NetworkingV1TransitGatewayAttachmentSpecUpdate The desired state of the Transit Gateway Attachment
type NetworkingV1TransitGatewayAttachmentSpecUpdate struct {
	// The name of the TGW attachment
	DisplayName *string `json:"display_name,omitempty"`
	// The environment to which this belongs.
	Environment *ObjectReference `json:"environment,omitempty"`
}

// NewNetworkingV1TransitGatewayAttachmentSpecUpdate instantiates a new NetworkingV1TransitGatewayAttachmentSpecUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1TransitGatewayAttachmentSpecUpdate() *NetworkingV1TransitGatewayAttachmentSpecUpdate {
	this := NetworkingV1TransitGatewayAttachmentSpecUpdate{}
	return &this
}

// NewNetworkingV1TransitGatewayAttachmentSpecUpdateWithDefaults instantiates a new NetworkingV1TransitGatewayAttachmentSpecUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1TransitGatewayAttachmentSpecUpdateWithDefaults() *NetworkingV1TransitGatewayAttachmentSpecUpdate {
	this := NetworkingV1TransitGatewayAttachmentSpecUpdate{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) GetEnvironment() ObjectReference {
	if o == nil || o.Environment == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) GetEnvironmentOk() (*ObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectReference and assigns it to the Environment field.
func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) SetEnvironment(v ObjectReference) {
	o.Environment = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) Redact() {
    o.recurseRedact(o.DisplayName)
    o.recurseRedact(o.Environment)
}

func (o *NetworkingV1TransitGatewayAttachmentSpecUpdate) recurseRedact(v interface{}) {
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

func (o NetworkingV1TransitGatewayAttachmentSpecUpdate) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1TransitGatewayAttachmentSpecUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1TransitGatewayAttachmentSpecUpdate struct {
	value *NetworkingV1TransitGatewayAttachmentSpecUpdate
	isSet bool
}

func (v NullableNetworkingV1TransitGatewayAttachmentSpecUpdate) Get() *NetworkingV1TransitGatewayAttachmentSpecUpdate {
	return v.value
}

func (v *NullableNetworkingV1TransitGatewayAttachmentSpecUpdate) Set(val *NetworkingV1TransitGatewayAttachmentSpecUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1TransitGatewayAttachmentSpecUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1TransitGatewayAttachmentSpecUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1TransitGatewayAttachmentSpecUpdate(val *NetworkingV1TransitGatewayAttachmentSpecUpdate) *NullableNetworkingV1TransitGatewayAttachmentSpecUpdate {
	return &NullableNetworkingV1TransitGatewayAttachmentSpecUpdate{value: val, isSet: true}
}

func (v NullableNetworkingV1TransitGatewayAttachmentSpecUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1TransitGatewayAttachmentSpecUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


