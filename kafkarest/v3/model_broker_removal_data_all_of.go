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

// BrokerRemovalDataAllOf struct for BrokerRemovalDataAllOf
type BrokerRemovalDataAllOf struct {
	ClusterId  string       `json:"cluster_id,omitempty"`
	BrokerId   int32        `json:"broker_id,omitempty"`
	BrokerTask Relationship `json:"broker_task,omitempty"`
	Broker     Relationship `json:"broker,omitempty"`
}

// NewBrokerRemovalDataAllOf instantiates a new BrokerRemovalDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrokerRemovalDataAllOf(clusterId string, brokerId int32, brokerTask Relationship, broker Relationship) *BrokerRemovalDataAllOf {
	this := BrokerRemovalDataAllOf{}
	this.ClusterId = clusterId
	this.BrokerId = brokerId
	this.BrokerTask = brokerTask
	this.Broker = broker
	return &this
}

// NewBrokerRemovalDataAllOfWithDefaults instantiates a new BrokerRemovalDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrokerRemovalDataAllOfWithDefaults() *BrokerRemovalDataAllOf {
	this := BrokerRemovalDataAllOf{}
	return &this
}

// GetClusterId returns the ClusterId field value
func (o *BrokerRemovalDataAllOf) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *BrokerRemovalDataAllOf) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *BrokerRemovalDataAllOf) SetClusterId(v string) {
	o.ClusterId = v
}

// GetBrokerId returns the BrokerId field value
func (o *BrokerRemovalDataAllOf) GetBrokerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BrokerId
}

// GetBrokerIdOk returns a tuple with the BrokerId field value
// and a boolean to check if the value has been set.
func (o *BrokerRemovalDataAllOf) GetBrokerIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerId, true
}

// SetBrokerId sets field value
func (o *BrokerRemovalDataAllOf) SetBrokerId(v int32) {
	o.BrokerId = v
}

// GetBrokerTask returns the BrokerTask field value
func (o *BrokerRemovalDataAllOf) GetBrokerTask() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.BrokerTask
}

// GetBrokerTaskOk returns a tuple with the BrokerTask field value
// and a boolean to check if the value has been set.
func (o *BrokerRemovalDataAllOf) GetBrokerTaskOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerTask, true
}

// SetBrokerTask sets field value
func (o *BrokerRemovalDataAllOf) SetBrokerTask(v Relationship) {
	o.BrokerTask = v
}

// GetBroker returns the Broker field value
func (o *BrokerRemovalDataAllOf) GetBroker() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Broker
}

// GetBrokerOk returns a tuple with the Broker field value
// and a boolean to check if the value has been set.
func (o *BrokerRemovalDataAllOf) GetBrokerOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Broker, true
}

// SetBroker sets field value
func (o *BrokerRemovalDataAllOf) SetBroker(v Relationship) {
	o.Broker = v
}

// Redact resets all sensitive fields to their zero value.
func (o *BrokerRemovalDataAllOf) Redact() {
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.BrokerId)
	o.recurseRedact(&o.BrokerTask)
	o.recurseRedact(&o.Broker)
}

func (o *BrokerRemovalDataAllOf) recurseRedact(v interface{}) {
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

func (o BrokerRemovalDataAllOf) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o BrokerRemovalDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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

type NullableBrokerRemovalDataAllOf struct {
	value *BrokerRemovalDataAllOf
	isSet bool
}

func (v NullableBrokerRemovalDataAllOf) Get() *BrokerRemovalDataAllOf {
	return v.value
}

func (v *NullableBrokerRemovalDataAllOf) Set(val *BrokerRemovalDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBrokerRemovalDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBrokerRemovalDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrokerRemovalDataAllOf(val *BrokerRemovalDataAllOf) *NullableBrokerRemovalDataAllOf {
	return &NullableBrokerRemovalDataAllOf{value: val, isSet: true}
}

func (v NullableBrokerRemovalDataAllOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableBrokerRemovalDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}