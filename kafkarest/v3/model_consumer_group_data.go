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

// ConsumerGroupData struct for ConsumerGroupData
type ConsumerGroupData struct {
	Kind              string           `json:"kind,omitempty"`
	Metadata          ResourceMetadata `json:"metadata,omitempty"`
	ClusterId         string           `json:"cluster_id,omitempty"`
	ConsumerGroupId   string           `json:"consumer_group_id,omitempty"`
	IsSimple          bool             `json:"is_simple,omitempty"`
	PartitionAssignor string           `json:"partition_assignor,omitempty"`
	State             string           `json:"state,omitempty"`
	Coordinator       Relationship     `json:"coordinator,omitempty"`
	Consumer          *Relationship    `json:"consumer,omitempty"`
	LagSummary        Relationship     `json:"lag_summary,omitempty"`
}

// NewConsumerGroupData instantiates a new ConsumerGroupData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsumerGroupData(kind string, metadata ResourceMetadata, clusterId string, consumerGroupId string, isSimple bool, partitionAssignor string, state string, coordinator Relationship, lagSummary Relationship) *ConsumerGroupData {
	this := ConsumerGroupData{}
	this.Kind = kind
	this.Metadata = metadata
	this.ClusterId = clusterId
	this.ConsumerGroupId = consumerGroupId
	this.IsSimple = isSimple
	this.PartitionAssignor = partitionAssignor
	this.State = state
	this.Coordinator = coordinator
	this.LagSummary = lagSummary
	return &this
}

// NewConsumerGroupDataWithDefaults instantiates a new ConsumerGroupData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsumerGroupDataWithDefaults() *ConsumerGroupData {
	this := ConsumerGroupData{}
	return &this
}

// GetKind returns the Kind field value
func (o *ConsumerGroupData) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupData) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ConsumerGroupData) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *ConsumerGroupData) GetMetadata() ResourceMetadata {
	if o == nil {
		var ret ResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupData) GetMetadataOk() (*ResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ConsumerGroupData) SetMetadata(v ResourceMetadata) {
	o.Metadata = v
}

// GetClusterId returns the ClusterId field value
func (o *ConsumerGroupData) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupData) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *ConsumerGroupData) SetClusterId(v string) {
	o.ClusterId = v
}

// GetConsumerGroupId returns the ConsumerGroupId field value
func (o *ConsumerGroupData) GetConsumerGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsumerGroupId
}

// GetConsumerGroupIdOk returns a tuple with the ConsumerGroupId field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupData) GetConsumerGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerGroupId, true
}

// SetConsumerGroupId sets field value
func (o *ConsumerGroupData) SetConsumerGroupId(v string) {
	o.ConsumerGroupId = v
}

// GetIsSimple returns the IsSimple field value
func (o *ConsumerGroupData) GetIsSimple() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsSimple
}

// GetIsSimpleOk returns a tuple with the IsSimple field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupData) GetIsSimpleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsSimple, true
}

// SetIsSimple sets field value
func (o *ConsumerGroupData) SetIsSimple(v bool) {
	o.IsSimple = v
}

// GetPartitionAssignor returns the PartitionAssignor field value
func (o *ConsumerGroupData) GetPartitionAssignor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartitionAssignor
}

// GetPartitionAssignorOk returns a tuple with the PartitionAssignor field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupData) GetPartitionAssignorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionAssignor, true
}

// SetPartitionAssignor sets field value
func (o *ConsumerGroupData) SetPartitionAssignor(v string) {
	o.PartitionAssignor = v
}

// GetState returns the State field value
func (o *ConsumerGroupData) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupData) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ConsumerGroupData) SetState(v string) {
	o.State = v
}

// GetCoordinator returns the Coordinator field value
func (o *ConsumerGroupData) GetCoordinator() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Coordinator
}

// GetCoordinatorOk returns a tuple with the Coordinator field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupData) GetCoordinatorOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coordinator, true
}

// SetCoordinator sets field value
func (o *ConsumerGroupData) SetCoordinator(v Relationship) {
	o.Coordinator = v
}

// GetConsumer returns the Consumer field value if set, zero value otherwise.
func (o *ConsumerGroupData) GetConsumer() Relationship {
	if o == nil || o.Consumer == nil {
		var ret Relationship
		return ret
	}
	return *o.Consumer
}

// GetConsumerOk returns a tuple with the Consumer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsumerGroupData) GetConsumerOk() (*Relationship, bool) {
	if o == nil || o.Consumer == nil {
		return nil, false
	}
	return o.Consumer, true
}

// HasConsumer returns a boolean if a field has been set.
func (o *ConsumerGroupData) HasConsumer() bool {
	if o != nil && o.Consumer != nil {
		return true
	}

	return false
}

// SetConsumer gets a reference to the given Relationship and assigns it to the Consumer field.
func (o *ConsumerGroupData) SetConsumer(v Relationship) {
	o.Consumer = &v
}

// GetLagSummary returns the LagSummary field value
func (o *ConsumerGroupData) GetLagSummary() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.LagSummary
}

// GetLagSummaryOk returns a tuple with the LagSummary field value
// and a boolean to check if the value has been set.
func (o *ConsumerGroupData) GetLagSummaryOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LagSummary, true
}

// SetLagSummary sets field value
func (o *ConsumerGroupData) SetLagSummary(v Relationship) {
	o.LagSummary = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ConsumerGroupData) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(&o.ConsumerGroupId)
	o.recurseRedact(&o.IsSimple)
	o.recurseRedact(&o.PartitionAssignor)
	o.recurseRedact(&o.State)
	o.recurseRedact(&o.Coordinator)
	o.recurseRedact(o.Consumer)
	o.recurseRedact(&o.LagSummary)
}

func (o *ConsumerGroupData) recurseRedact(v interface{}) {
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

func (o ConsumerGroupData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ConsumerGroupData) MarshalJSON() ([]byte, error) {
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
		toSerialize["consumer_group_id"] = o.ConsumerGroupId
	}
	if true {
		toSerialize["is_simple"] = o.IsSimple
	}
	if true {
		toSerialize["partition_assignor"] = o.PartitionAssignor
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["coordinator"] = o.Coordinator
	}
	if o.Consumer != nil {
		toSerialize["consumer"] = o.Consumer
	}
	if true {
		toSerialize["lag_summary"] = o.LagSummary
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableConsumerGroupData struct {
	value *ConsumerGroupData
	isSet bool
}

func (v NullableConsumerGroupData) Get() *ConsumerGroupData {
	return v.value
}

func (v *NullableConsumerGroupData) Set(val *ConsumerGroupData) {
	v.value = val
	v.isSet = true
}

func (v NullableConsumerGroupData) IsSet() bool {
	return v.isSet
}

func (v *NullableConsumerGroupData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsumerGroupData(val *ConsumerGroupData) *NullableConsumerGroupData {
	return &NullableConsumerGroupData{value: val, isSet: true}
}

func (v NullableConsumerGroupData) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableConsumerGroupData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}