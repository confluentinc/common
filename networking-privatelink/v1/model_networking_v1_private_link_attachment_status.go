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

// NetworkingV1PrivateLinkAttachmentStatus The status of the Private Link Attachment
type NetworkingV1PrivateLinkAttachmentStatus struct {
	// The lifecycle phase of the PrivateLink attachment:    PROVISIONING: PrivateLink attachment provisioning is in progress;    WAITING_FOR_CONNECTIONS: PrivateLink attachment is waiting for connections;    READY: PrivateLink attachment is ready;    FAILED: PrivateLink attachment is in a failed state;    EXPIRED: PrivateLink attachment has timed out waiting for connections, can only be deleted;    DEPROVISIONING: PrivateLink attachment deprovisioning is in progress;
	Phase string `json:"phase,omitempty"`
	// Error code if PrivateLink attachment is in a failed state. May be used for programmatic error checking.
	ErrorCode *string `json:"error_code,omitempty"`
	// Displayable error message if PrivateLink attachment is in a failed state.
	ErrorMessage *string `json:"error_message,omitempty"`
	// The root DNS domain for the PrivateLink attachment.
	DnsDomain *string `json:"dns_domain,omitempty"`
	// The cloud specific status of the PrivateLink attachment. These will be populated when the PrivateLink attachment reaches the WAITING_FOR_CONNECTIONS state.
	Cloud *NetworkingV1PrivateLinkAttachmentStatusCloudOneOf `json:"cloud,omitempty"`
}

// NewNetworkingV1PrivateLinkAttachmentStatus instantiates a new NetworkingV1PrivateLinkAttachmentStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1PrivateLinkAttachmentStatus(phase string) *NetworkingV1PrivateLinkAttachmentStatus {
	this := NetworkingV1PrivateLinkAttachmentStatus{}
	this.Phase = phase
	return &this
}

// NewNetworkingV1PrivateLinkAttachmentStatusWithDefaults instantiates a new NetworkingV1PrivateLinkAttachmentStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1PrivateLinkAttachmentStatusWithDefaults() *NetworkingV1PrivateLinkAttachmentStatus {
	this := NetworkingV1PrivateLinkAttachmentStatus{}
	return &this
}

// GetPhase returns the Phase field value
func (o *NetworkingV1PrivateLinkAttachmentStatus) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentStatus) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *NetworkingV1PrivateLinkAttachmentStatus) SetPhase(v string) {
	o.Phase = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *NetworkingV1PrivateLinkAttachmentStatus) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentStatus) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *NetworkingV1PrivateLinkAttachmentStatus) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *NetworkingV1PrivateLinkAttachmentStatus) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *NetworkingV1PrivateLinkAttachmentStatus) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentStatus) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *NetworkingV1PrivateLinkAttachmentStatus) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *NetworkingV1PrivateLinkAttachmentStatus) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetDnsDomain returns the DnsDomain field value if set, zero value otherwise.
func (o *NetworkingV1PrivateLinkAttachmentStatus) GetDnsDomain() string {
	if o == nil || o.DnsDomain == nil {
		var ret string
		return ret
	}
	return *o.DnsDomain
}

// GetDnsDomainOk returns a tuple with the DnsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentStatus) GetDnsDomainOk() (*string, bool) {
	if o == nil || o.DnsDomain == nil {
		return nil, false
	}
	return o.DnsDomain, true
}

// HasDnsDomain returns a boolean if a field has been set.
func (o *NetworkingV1PrivateLinkAttachmentStatus) HasDnsDomain() bool {
	if o != nil && o.DnsDomain != nil {
		return true
	}

	return false
}

// SetDnsDomain gets a reference to the given string and assigns it to the DnsDomain field.
func (o *NetworkingV1PrivateLinkAttachmentStatus) SetDnsDomain(v string) {
	o.DnsDomain = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *NetworkingV1PrivateLinkAttachmentStatus) GetCloud() NetworkingV1PrivateLinkAttachmentStatusCloudOneOf {
	if o == nil || o.Cloud == nil {
		var ret NetworkingV1PrivateLinkAttachmentStatusCloudOneOf
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1PrivateLinkAttachmentStatus) GetCloudOk() (*NetworkingV1PrivateLinkAttachmentStatusCloudOneOf, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *NetworkingV1PrivateLinkAttachmentStatus) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given NetworkingV1PrivateLinkAttachmentStatusCloudOneOf and assigns it to the Cloud field.
func (o *NetworkingV1PrivateLinkAttachmentStatus) SetCloud(v NetworkingV1PrivateLinkAttachmentStatusCloudOneOf) {
	o.Cloud = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1PrivateLinkAttachmentStatus) Redact() {
	o.recurseRedact(&o.Phase)
	o.recurseRedact(o.ErrorCode)
	o.recurseRedact(o.ErrorMessage)
	o.recurseRedact(o.DnsDomain)
	o.recurseRedact(o.Cloud)
}

func (o *NetworkingV1PrivateLinkAttachmentStatus) recurseRedact(v interface{}) {
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

func (o NetworkingV1PrivateLinkAttachmentStatus) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1PrivateLinkAttachmentStatus) MarshalJSON() ([]byte, error) {
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
	if o.DnsDomain != nil {
		toSerialize["dns_domain"] = o.DnsDomain
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

type NullableNetworkingV1PrivateLinkAttachmentStatus struct {
	value *NetworkingV1PrivateLinkAttachmentStatus
	isSet bool
}

func (v NullableNetworkingV1PrivateLinkAttachmentStatus) Get() *NetworkingV1PrivateLinkAttachmentStatus {
	return v.value
}

func (v *NullableNetworkingV1PrivateLinkAttachmentStatus) Set(val *NetworkingV1PrivateLinkAttachmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1PrivateLinkAttachmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1PrivateLinkAttachmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1PrivateLinkAttachmentStatus(val *NetworkingV1PrivateLinkAttachmentStatus) *NullableNetworkingV1PrivateLinkAttachmentStatus {
	return &NullableNetworkingV1PrivateLinkAttachmentStatus{value: val, isSet: true}
}

func (v NullableNetworkingV1PrivateLinkAttachmentStatus) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableNetworkingV1PrivateLinkAttachmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
