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

// ClusterDataAllOf struct for ClusterDataAllOf
type ClusterDataAllOf struct {
	ClusterId              string        `json:"cluster_id,omitempty"`
	Controller             *Relationship `json:"controller,omitempty"`
	Acls                   Relationship  `json:"acls,omitempty"`
	Brokers                Relationship  `json:"brokers,omitempty"`
	BrokerConfigs          Relationship  `json:"broker_configs,omitempty"`
	ConsumerGroups         Relationship  `json:"consumer_groups,omitempty"`
	Topics                 Relationship  `json:"topics,omitempty"`
	PartitionReassignments Relationship  `json:"partition_reassignments,omitempty"`
}

// NewClusterDataAllOf instantiates a new ClusterDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterDataAllOf(clusterId string, acls Relationship, brokers Relationship, brokerConfigs Relationship, consumerGroups Relationship, topics Relationship, partitionReassignments Relationship) *ClusterDataAllOf {
	this := ClusterDataAllOf{}
	this.ClusterId = clusterId
	this.Acls = acls
	this.Brokers = brokers
	this.BrokerConfigs = brokerConfigs
	this.ConsumerGroups = consumerGroups
	this.Topics = topics
	this.PartitionReassignments = partitionReassignments
	return &this
}

// NewClusterDataAllOfWithDefaults instantiates a new ClusterDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterDataAllOfWithDefaults() *ClusterDataAllOf {
	this := ClusterDataAllOf{}
	return &this
}

// GetClusterId returns the ClusterId field value
func (o *ClusterDataAllOf) GetClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterId
}

// GetClusterIdOk returns a tuple with the ClusterId field value
// and a boolean to check if the value has been set.
func (o *ClusterDataAllOf) GetClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterId, true
}

// SetClusterId sets field value
func (o *ClusterDataAllOf) SetClusterId(v string) {
	o.ClusterId = v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *ClusterDataAllOf) GetController() Relationship {
	if o == nil || o.Controller == nil {
		var ret Relationship
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterDataAllOf) GetControllerOk() (*Relationship, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *ClusterDataAllOf) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given Relationship and assigns it to the Controller field.
func (o *ClusterDataAllOf) SetController(v Relationship) {
	o.Controller = &v
}

// GetAcls returns the Acls field value
func (o *ClusterDataAllOf) GetAcls() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Acls
}

// GetAclsOk returns a tuple with the Acls field value
// and a boolean to check if the value has been set.
func (o *ClusterDataAllOf) GetAclsOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Acls, true
}

// SetAcls sets field value
func (o *ClusterDataAllOf) SetAcls(v Relationship) {
	o.Acls = v
}

// GetBrokers returns the Brokers field value
func (o *ClusterDataAllOf) GetBrokers() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Brokers
}

// GetBrokersOk returns a tuple with the Brokers field value
// and a boolean to check if the value has been set.
func (o *ClusterDataAllOf) GetBrokersOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Brokers, true
}

// SetBrokers sets field value
func (o *ClusterDataAllOf) SetBrokers(v Relationship) {
	o.Brokers = v
}

// GetBrokerConfigs returns the BrokerConfigs field value
func (o *ClusterDataAllOf) GetBrokerConfigs() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.BrokerConfigs
}

// GetBrokerConfigsOk returns a tuple with the BrokerConfigs field value
// and a boolean to check if the value has been set.
func (o *ClusterDataAllOf) GetBrokerConfigsOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BrokerConfigs, true
}

// SetBrokerConfigs sets field value
func (o *ClusterDataAllOf) SetBrokerConfigs(v Relationship) {
	o.BrokerConfigs = v
}

// GetConsumerGroups returns the ConsumerGroups field value
func (o *ClusterDataAllOf) GetConsumerGroups() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.ConsumerGroups
}

// GetConsumerGroupsOk returns a tuple with the ConsumerGroups field value
// and a boolean to check if the value has been set.
func (o *ClusterDataAllOf) GetConsumerGroupsOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsumerGroups, true
}

// SetConsumerGroups sets field value
func (o *ClusterDataAllOf) SetConsumerGroups(v Relationship) {
	o.ConsumerGroups = v
}

// GetTopics returns the Topics field value
func (o *ClusterDataAllOf) GetTopics() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value
// and a boolean to check if the value has been set.
func (o *ClusterDataAllOf) GetTopicsOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Topics, true
}

// SetTopics sets field value
func (o *ClusterDataAllOf) SetTopics(v Relationship) {
	o.Topics = v
}

// GetPartitionReassignments returns the PartitionReassignments field value
func (o *ClusterDataAllOf) GetPartitionReassignments() Relationship {
	if o == nil {
		var ret Relationship
		return ret
	}

	return o.PartitionReassignments
}

// GetPartitionReassignmentsOk returns a tuple with the PartitionReassignments field value
// and a boolean to check if the value has been set.
func (o *ClusterDataAllOf) GetPartitionReassignmentsOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartitionReassignments, true
}

// SetPartitionReassignments sets field value
func (o *ClusterDataAllOf) SetPartitionReassignments(v Relationship) {
	o.PartitionReassignments = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ClusterDataAllOf) Redact() {
	o.recurseRedact(&o.ClusterId)
	o.recurseRedact(o.Controller)
	o.recurseRedact(&o.Acls)
	o.recurseRedact(&o.Brokers)
	o.recurseRedact(&o.BrokerConfigs)
	o.recurseRedact(&o.ConsumerGroups)
	o.recurseRedact(&o.Topics)
	o.recurseRedact(&o.PartitionReassignments)
}

func (o *ClusterDataAllOf) recurseRedact(v interface{}) {
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

func (o ClusterDataAllOf) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ClusterDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster_id"] = o.ClusterId
	}
	if o.Controller != nil {
		toSerialize["controller"] = o.Controller
	}
	if true {
		toSerialize["acls"] = o.Acls
	}
	if true {
		toSerialize["brokers"] = o.Brokers
	}
	if true {
		toSerialize["broker_configs"] = o.BrokerConfigs
	}
	if true {
		toSerialize["consumer_groups"] = o.ConsumerGroups
	}
	if true {
		toSerialize["topics"] = o.Topics
	}
	if true {
		toSerialize["partition_reassignments"] = o.PartitionReassignments
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableClusterDataAllOf struct {
	value *ClusterDataAllOf
	isSet bool
}

func (v NullableClusterDataAllOf) Get() *ClusterDataAllOf {
	return v.value
}

func (v *NullableClusterDataAllOf) Set(val *ClusterDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDataAllOf(val *ClusterDataAllOf) *NullableClusterDataAllOf {
	return &NullableClusterDataAllOf{value: val, isSet: true}
}

func (v NullableClusterDataAllOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableClusterDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}