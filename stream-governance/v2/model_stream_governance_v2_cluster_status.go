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
Cluster Management for Stream Governance API

Cluster Management for Stream Governance API

API version: 0.0.1-alpha1
Contact: data-governance@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

import (
	"reflect"
)

// StreamGovernanceV2ClusterStatus The status of the Cluster
type StreamGovernanceV2ClusterStatus struct {
	// The lifecyle phase of the cluster:   PROVISIONED:  cluster is provisioned;   PROVISIONING:  cluster provisioning is in progress;   FAILED:  provisioning failed
	Phase string `json:"phase"`
}

// NewStreamGovernanceV2ClusterStatus instantiates a new StreamGovernanceV2ClusterStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamGovernanceV2ClusterStatus(phase string) *StreamGovernanceV2ClusterStatus {
	this := StreamGovernanceV2ClusterStatus{}
	this.Phase = phase
	return &this
}

// NewStreamGovernanceV2ClusterStatusWithDefaults instantiates a new StreamGovernanceV2ClusterStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamGovernanceV2ClusterStatusWithDefaults() *StreamGovernanceV2ClusterStatus {
	this := StreamGovernanceV2ClusterStatus{}
	return &this
}

// GetPhase returns the Phase field value
func (o *StreamGovernanceV2ClusterStatus) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *StreamGovernanceV2ClusterStatus) GetPhaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *StreamGovernanceV2ClusterStatus) SetPhase(v string) {
	o.Phase = v
}

// Redact resets all sensitive fields to their zero value.
func (o *StreamGovernanceV2ClusterStatus) Redact() {
	o.recurseRedact(&o.Phase)
}

func (o *StreamGovernanceV2ClusterStatus) recurseRedact(v interface{}) {
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

func (o StreamGovernanceV2ClusterStatus) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o StreamGovernanceV2ClusterStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["phase"] = o.Phase
	}
	return json.Marshal(toSerialize)
}

type NullableStreamGovernanceV2ClusterStatus struct {
	value *StreamGovernanceV2ClusterStatus
	isSet bool
}

func (v NullableStreamGovernanceV2ClusterStatus) Get() *StreamGovernanceV2ClusterStatus {
	return v.value
}

func (v *NullableStreamGovernanceV2ClusterStatus) Set(val *StreamGovernanceV2ClusterStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamGovernanceV2ClusterStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamGovernanceV2ClusterStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamGovernanceV2ClusterStatus(val *StreamGovernanceV2ClusterStatus) *NullableStreamGovernanceV2ClusterStatus {
	return &NullableStreamGovernanceV2ClusterStatus{value: val, isSet: true}
}

func (v NullableStreamGovernanceV2ClusterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamGovernanceV2ClusterStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
