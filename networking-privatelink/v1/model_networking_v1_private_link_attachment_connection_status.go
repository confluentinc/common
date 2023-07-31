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

// NetworkingV1PrivateLinkAttachmentConnectionStatus The status of the Private Link Attachment Connection
type NetworkingV1PrivateLinkAttachmentConnectionStatus struct {
	// The lifecycle phase of the PrivateLink attachment:    PROVISIONING: PrivateLink attachment connection provisioning is in progress;    READY: PrivateLink attachment connection is ready;    FAILED: PrivateLink attachment connection is in a failed state;    DEPROVISIONING: PrivateLink attachment connection deprovisioning is in progress;
	Phase string `json:"phase,omitempty"`
	// Error code if PrivateLink attachment connection is in a failed state. May be used for programmatic error checking.
	ErrorCode *string `json:"error_code,omitempty"`
	// Displayable error message if PrivateLink attachment connection is in a failed state.
	ErrorMessage *string `json:"error_message,omitempty"`
	// The cloud specific status of the PrivateLink attachment connection.
	Cloud *NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf `json:"cloud,omitempty"`
}

// NewNetworkingV1PrivateLinkAttachmentConnectionStatus instantiates a new NetworkingV1PrivateLinkAttachmentConnectionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1PrivateLinkAttachmentConnectionStatus(phase string) *NetworkingV1PrivateLinkAttachmentConnectionStatus {
	this := NetworkingV1PrivateLinkAttachmentConnectionStatus{}
	this.Phase = phase
	return &this
}

// NewNetworkingV1PrivateLinkAttachmentConnectionStatusWithDefaults instantiates a new NetworkingV1PrivateLinkAttachmentConnectionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1PrivateLinkAttachmentConnectionStatusWithDefaults() *NetworkingV1PrivateLinkAttachmentConnectionStatus {
	this := NetworkingV1PrivateLinkAttachmentConnectionStatus{}
	return &this
}

// GetPhase returns the Phase field value
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) SetPhase(v string) {
	o.Phase = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetCloud() NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf {
	if o == nil || o.Cloud == nil {
		var ret NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) GetCloudOk() (*NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf and assigns it to the Cloud field.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) SetCloud(v NetworkingV1PrivateLinkAttachmentConnectionStatusCloudOneOf) {
	o.Cloud = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) Redact() {
	o.recurseRedact(&o.Phase)
	o.recurseRedact(o.ErrorCode)
	o.recurseRedact(o.ErrorMessage)
	o.recurseRedact(o.Cloud)
}

func (o *NetworkingV1PrivateLinkAttachmentConnectionStatus) recurseRedact(v interface{}) {
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

func (o NetworkingV1PrivateLinkAttachmentConnectionStatus) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1PrivateLinkAttachmentConnectionStatus) MarshalJSON() ([]byte, error) {
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
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableNetworkingV1PrivateLinkAttachmentConnectionStatus struct {
	value *NetworkingV1PrivateLinkAttachmentConnectionStatus
	isSet bool
}

func (v NullableNetworkingV1PrivateLinkAttachmentConnectionStatus) Get() *NetworkingV1PrivateLinkAttachmentConnectionStatus {
	return v.value
}

func (v *NullableNetworkingV1PrivateLinkAttachmentConnectionStatus) Set(val *NetworkingV1PrivateLinkAttachmentConnectionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1PrivateLinkAttachmentConnectionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1PrivateLinkAttachmentConnectionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1PrivateLinkAttachmentConnectionStatus(val *NetworkingV1PrivateLinkAttachmentConnectionStatus) *NullableNetworkingV1PrivateLinkAttachmentConnectionStatus {
	return &NullableNetworkingV1PrivateLinkAttachmentConnectionStatus{value: val, isSet: true}
}

func (v NullableNetworkingV1PrivateLinkAttachmentConnectionStatus) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1PrivateLinkAttachmentConnectionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
