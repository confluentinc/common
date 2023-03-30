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

// NetworkingV1TransitGatewayAttachmentStatus The status of the Transit Gateway Attachment
type NetworkingV1TransitGatewayAttachmentStatus struct {
	// The lifecycle phase of the TGW attachment:    PROVISIONING: attachment provisioning is in progress;    PENDING_ACCEPT: attachment request is pending acceptance by the customer;    READY:  attachment is ready;    FAILED: attachment is in a failed state;    DEPROVISIONING: attachment deprovisioning is in progress;    DISCONNECTED: attachment was manually deleted directly in the cloud provider by the customer;    ERROR: invalid customer input during attachment creation.
	Phase string `json:"phase"`
	// Error code if TGW attachment is in a failed state. May be used for programmatic error checking.
	ErrorCode *string `json:"error_code,omitempty"`
	// Displayable error message if TGW attachment is in a failed state
	ErrorMessage *string `json:"error_message,omitempty"`
	// The cloud-specific TGW attachment details.
	Cloud *NetworkingV1TransitGatewayAttachmentStatusCloudOneOf `json:"cloud,omitempty"`
}

// NewNetworkingV1TransitGatewayAttachmentStatus instantiates a new NetworkingV1TransitGatewayAttachmentStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1TransitGatewayAttachmentStatus(phase string) *NetworkingV1TransitGatewayAttachmentStatus {
	this := NetworkingV1TransitGatewayAttachmentStatus{}
	this.Phase = phase
	return &this
}

// NewNetworkingV1TransitGatewayAttachmentStatusWithDefaults instantiates a new NetworkingV1TransitGatewayAttachmentStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1TransitGatewayAttachmentStatusWithDefaults() *NetworkingV1TransitGatewayAttachmentStatus {
	this := NetworkingV1TransitGatewayAttachmentStatus{}
	return &this
}

// GetPhase returns the Phase field value
func (o *NetworkingV1TransitGatewayAttachmentStatus) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1TransitGatewayAttachmentStatus) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *NetworkingV1TransitGatewayAttachmentStatus) SetPhase(v string) {
	o.Phase = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *NetworkingV1TransitGatewayAttachmentStatus) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1TransitGatewayAttachmentStatus) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *NetworkingV1TransitGatewayAttachmentStatus) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *NetworkingV1TransitGatewayAttachmentStatus) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *NetworkingV1TransitGatewayAttachmentStatus) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1TransitGatewayAttachmentStatus) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *NetworkingV1TransitGatewayAttachmentStatus) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *NetworkingV1TransitGatewayAttachmentStatus) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *NetworkingV1TransitGatewayAttachmentStatus) GetCloud() NetworkingV1TransitGatewayAttachmentStatusCloudOneOf {
	if o == nil || o.Cloud == nil {
		var ret NetworkingV1TransitGatewayAttachmentStatusCloudOneOf
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1TransitGatewayAttachmentStatus) GetCloudOk() (*NetworkingV1TransitGatewayAttachmentStatusCloudOneOf, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *NetworkingV1TransitGatewayAttachmentStatus) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given NetworkingV1TransitGatewayAttachmentStatusCloudOneOf and assigns it to the Cloud field.
func (o *NetworkingV1TransitGatewayAttachmentStatus) SetCloud(v NetworkingV1TransitGatewayAttachmentStatusCloudOneOf) {
	o.Cloud = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1TransitGatewayAttachmentStatus) Redact() {
	o.recurseRedact(&o.Phase)
	o.recurseRedact(o.ErrorCode)
	o.recurseRedact(o.ErrorMessage)
	o.recurseRedact(o.Cloud)
}

func (o *NetworkingV1TransitGatewayAttachmentStatus) recurseRedact(v interface{}) {
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

func (o NetworkingV1TransitGatewayAttachmentStatus) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1TransitGatewayAttachmentStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["phase"] = o.Phase
	}
	if o.ErrorCode != nil {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1TransitGatewayAttachmentStatus struct {
	value *NetworkingV1TransitGatewayAttachmentStatus
	isSet bool
}

func (v NullableNetworkingV1TransitGatewayAttachmentStatus) Get() *NetworkingV1TransitGatewayAttachmentStatus {
	return v.value
}

func (v *NullableNetworkingV1TransitGatewayAttachmentStatus) Set(val *NetworkingV1TransitGatewayAttachmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1TransitGatewayAttachmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1TransitGatewayAttachmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1TransitGatewayAttachmentStatus(val *NetworkingV1TransitGatewayAttachmentStatus) *NullableNetworkingV1TransitGatewayAttachmentStatus {
	return &NullableNetworkingV1TransitGatewayAttachmentStatus{value: val, isSet: true}
}

func (v NullableNetworkingV1TransitGatewayAttachmentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1TransitGatewayAttachmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
