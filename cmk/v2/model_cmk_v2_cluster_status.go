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
Cluster Management for Apache Kafka API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha1
Contact: orchestrator-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

import (
	"reflect"
)

// CmkV2ClusterStatus The status of the Cluster
type CmkV2ClusterStatus struct {
	// The lifecyle phase of the cluster:   PROVISIONED:  cluster is provisioned;   PROVISIONING:  cluster provisioning is in progress;   FAILED:  provisioning failed 
	Phase string `json:"phase"`
	// The number of Confluent Kafka Units (CKUs) the Dedicated cluster currently has. 
	Cku *int32 `json:"cku,omitempty"`
}

// NewCmkV2ClusterStatus instantiates a new CmkV2ClusterStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmkV2ClusterStatus(phase string) *CmkV2ClusterStatus {
	this := CmkV2ClusterStatus{}
	this.Phase = phase
	return &this
}

// NewCmkV2ClusterStatusWithDefaults instantiates a new CmkV2ClusterStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmkV2ClusterStatusWithDefaults() *CmkV2ClusterStatus {
	this := CmkV2ClusterStatus{}
	return &this
}

// GetPhase returns the Phase field value
func (o *CmkV2ClusterStatus) GetPhase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Phase
}

// GetPhaseOk returns a tuple with the Phase field value
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterStatus) GetPhaseOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Phase, true
}

// SetPhase sets field value
func (o *CmkV2ClusterStatus) SetPhase(v string) {
	o.Phase = v
}

// GetCku returns the Cku field value if set, zero value otherwise.
func (o *CmkV2ClusterStatus) GetCku() int32 {
	if o == nil || o.Cku == nil {
		var ret int32
		return ret
	}
	return *o.Cku
}

// GetCkuOk returns a tuple with the Cku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterStatus) GetCkuOk() (*int32, bool) {
	if o == nil || o.Cku == nil {
		return nil, false
	}
	return o.Cku, true
}

// HasCku returns a boolean if a field has been set.
func (o *CmkV2ClusterStatus) HasCku() bool {
	if o != nil && o.Cku != nil {
		return true
	}

	return false
}

// SetCku gets a reference to the given int32 and assigns it to the Cku field.
func (o *CmkV2ClusterStatus) SetCku(v int32) {
	o.Cku = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CmkV2ClusterStatus) Redact() {
    o.recurseRedact(&o.Phase)
    o.recurseRedact(o.Cku)
}

func (o *CmkV2ClusterStatus) recurseRedact(v interface{}) {
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

func (o CmkV2ClusterStatus) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o CmkV2ClusterStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["phase"] = o.Phase
	}
	if o.Cku != nil {
		toSerialize["cku"] = o.Cku
	}
	return json.Marshal(toSerialize)
}

type NullableCmkV2ClusterStatus struct {
	value *CmkV2ClusterStatus
	isSet bool
}

func (v NullableCmkV2ClusterStatus) Get() *CmkV2ClusterStatus {
	return v.value
}

func (v *NullableCmkV2ClusterStatus) Set(val *CmkV2ClusterStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCmkV2ClusterStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCmkV2ClusterStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmkV2ClusterStatus(val *CmkV2ClusterStatus) *NullableCmkV2ClusterStatus {
	return &NullableCmkV2ClusterStatus{value: val, isSet: true}
}

func (v NullableCmkV2ClusterStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmkV2ClusterStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


