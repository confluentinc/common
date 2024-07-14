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
REST Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0
Contact: kafka-clients-proxy-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// BrokerRemovalData struct for BrokerRemovalData
type BrokerRemovalData struct {
	Kind       string           `json:"kind,omitempty"`
	Metadata   ResourceMetadata `json:"metadata,omitempty"`
	ClusterId  string           `json:"cluster_id,omitempty"`
	BrokerId   int32            `json:"broker_id,omitempty"`
	BrokerTask Relationship     `json:"broker_task,omitempty"`
	Broker     Relationship     `json:"broker,omitempty"`
}

// NewBrokerRemovalData instantiates a new BrokerRemovalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerRemovalData(kind string, metadata ResourceMetadata, clusterId string, brokerId int32, brokerTask Relationship, broker Relationship) *BrokerRemovalData {
	this := BrokerRemovalData{}
	this.Kind = kind
	this.Metadata = metadata
	this.ClusterId = clusterId
	this.BrokerId = brokerId
	this.BrokerTask = brokerTask
	this.Broker = broker
	return &this
}

// NewBrokerRemovalDataWithDefaults instantiates a new BrokerRemovalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerRemovalDataWithDefaults() *BrokerRemovalData {
	this := BrokerRemovalData{}
	return &this
}

// GetKind returns the Kind field value
func (o *BrokerRemovalData) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *BrokerRemovalData) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *BrokerRemovalData) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *BrokerRemovalData) GetMetadata() ResourceMetadata {
	if o == nil {
		var ret ResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *BrokerRemovalData) GetMetadataOk() (*ResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *BrokerRemovalData) SetMetadata(v ResourceMetadata) {
	o.Metadata = v
}

// GetClusterId returns the ClusterId field value
func (o *BrokerRemovalData) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *BrokerRemovalData) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *BrokerRemovalData) SetClusterId(v string) {
	o.ClusterId = v
}

// GetBrokerId returns the BrokerId field value
func (o *BrokerRemovalData) GetBrokerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value
// and a boolean to check if the value has been set.
func (o *BrokerRemovalData) GetBrokerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerId, true
}

// SetBrokerId sets field value
func (o *BrokerRemovalData) SetBrokerId(v int32) {
	o.BrokerId = v
}

// GetBrokerTask returns the BrokerTask field value
func (o *BrokerRemovalData) GetBrokerTask() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.BrokerTask
}

// GetBrokerTaskOk returns a tuple with the BrokerTask field value
// and a boolean to check if the value has been set.
func (o *BrokerRemovalData) GetBrokerTaskOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerTask, true
}

// SetBrokerTask sets field value
func (o *BrokerRemovalData) SetBrokerTask(v Relationship) {
	o.BrokerTask = v
}

// GetBroker returns the Broker field value
func (o *BrokerRemovalData) GetBroker() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Broker
}

// GetBrokerOk returns a tuple with the Broker field value
// and a boolean to check if the value has been set.
func (o *BrokerRemovalData) GetBrokerOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Broker, true
}

// SetBroker sets field value
func (o *BrokerRemovalData) SetBroker(v Relationship) {
	o.Broker = v
}

// Redact resets all sensitive fields to their zero value.
func (o *BrokerRemovalData) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.BrokerId)
	o.recurseRedact(&o.BrokerTask)
	o.recurseRedact(&o.Broker)
}

func (o *BrokerRemovalData) recurseRedact(v interface{}) {
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

func (o BrokerRemovalData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o BrokerRemovalData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if true {
		toSerialize["broker_id"] = o.BrokerId
	}
	if true {
		toSerialize["broker_task"] = o.BrokerTask
	}
	if true {
		toSerialize["broker"] = o.Broker
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableBrokerRemovalData struct {
	value *BrokerRemovalData
	isSet bool
}

func (v NullableBrokerRemovalData) Get() *BrokerRemovalData {
	return v.value
}

func (v *NullableBrokerRemovalData) Set(val *BrokerRemovalData) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerRemovalData) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerRemovalData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerRemovalData(val *BrokerRemovalData) *NullableBrokerRemovalData {
	return &NullableBrokerRemovalData{value: val, isSet: true}
}

func (v NullableBrokerRemovalData) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableBrokerRemovalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}