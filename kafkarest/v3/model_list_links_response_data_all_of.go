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

// ListLinksResponseDataAllOf struct for ListLinksResponseDataAllOf
type ListLinksResponseDataAllOf struct {
	SourceClusterId      NullableString `json:"source_cluster_id,omitempty"`
	DestinationClusterId NullableString `json:"destination_cluster_id,omitempty"`
	RemoteClusterId      NullableString `json:"remote_cluster_id,omitempty"`
	LinkName             string         `json:"link_name,omitempty"`
	// Deprecated
	LinkId           *string        `json:"link_id,omitempty"`
	ClusterLinkId    string         `json:"cluster_link_id,omitempty"`
	TopicNames       []string       `json:"topic_names,omitempty"`
	LinkError        *string        `json:"link_error,omitempty"`
	LinkErrorMessage NullableString `json:"link_error_message,omitempty"`
	LinkState        *string        `json:"link_state,omitempty"`
	Tasks            []LinkTask     `json:"tasks,omitempty"`
}

// NewListLinksResponseDataAllOf instantiates a new ListLinksResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLinksResponseDataAllOf(linkName string, clusterLinkId string, topicNames []string) *ListLinksResponseDataAllOf {
	this := ListLinksResponseDataAllOf{}
	this.LinkName = linkName
	this.ClusterLinkId = clusterLinkId
	this.TopicNames = topicNames
	return &this
}

// NewListLinksResponseDataAllOfWithDefaults instantiates a new ListLinksResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLinksResponseDataAllOfWithDefaults() *ListLinksResponseDataAllOf {
	this := ListLinksResponseDataAllOf{}
	return &this
}

// GetSourceClusterId returns the SourceClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListLinksResponseDataAllOf) GetSourceClusterId() string {
	if o == nil || o.SourceClusterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceClusterId.Get()
}

// GetSourceClusterIdOk returns a tuple with the SourceClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListLinksResponseDataAllOf) GetSourceClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceClusterId.Get(), o.SourceClusterId.IsSet()
}

// HasSourceClusterId returns a boolean if a field has been set.
func (o *ListLinksResponseDataAllOf) HasSourceClusterId() bool {
	if o != nil && o.SourceClusterId.IsSet() {
		return true
	}

	return false
}

// SetSourceClusterId gets a reference to the given NullableString and assigns it to the SourceClusterId field.
func (o *ListLinksResponseDataAllOf) SetSourceClusterId(v string) {
	o.SourceClusterId.Set(&v)
}

// SetSourceClusterIdNil sets the value for SourceClusterId to be an explicit nil
func (o *ListLinksResponseDataAllOf) SetSourceClusterIdNil() {
	o.SourceClusterId.Set(nil)
}

// UnsetSourceClusterId ensures that no value is present for SourceClusterId, not even an explicit nil
func (o *ListLinksResponseDataAllOf) UnsetSourceClusterId() {
	o.SourceClusterId.Unset()
}

// GetDestinationClusterId returns the DestinationClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListLinksResponseDataAllOf) GetDestinationClusterId() string {
	if o == nil || o.DestinationClusterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DestinationClusterId.Get()
}

// GetDestinationClusterIdOk returns a tuple with the DestinationClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListLinksResponseDataAllOf) GetDestinationClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestinationClusterId.Get(), o.DestinationClusterId.IsSet()
}

// HasDestinationClusterId returns a boolean if a field has been set.
func (o *ListLinksResponseDataAllOf) HasDestinationClusterId() bool {
	if o != nil && o.DestinationClusterId.IsSet() {
		return true
	}

	return false
}

// SetDestinationClusterId gets a reference to the given NullableString and assigns it to the DestinationClusterId field.
func (o *ListLinksResponseDataAllOf) SetDestinationClusterId(v string) {
	o.DestinationClusterId.Set(&v)
}

// SetDestinationClusterIdNil sets the value for DestinationClusterId to be an explicit nil
func (o *ListLinksResponseDataAllOf) SetDestinationClusterIdNil() {
	o.DestinationClusterId.Set(nil)
}

// UnsetDestinationClusterId ensures that no value is present for DestinationClusterId, not even an explicit nil
func (o *ListLinksResponseDataAllOf) UnsetDestinationClusterId() {
	o.DestinationClusterId.Unset()
}

// GetRemoteClusterId returns the RemoteClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListLinksResponseDataAllOf) GetRemoteClusterId() string {
	if o == nil || o.RemoteClusterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.RemoteClusterId.Get()
}

// GetRemoteClusterIdOk returns a tuple with the RemoteClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListLinksResponseDataAllOf) GetRemoteClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemoteClusterId.Get(), o.RemoteClusterId.IsSet()
}

// HasRemoteClusterId returns a boolean if a field has been set.
func (o *ListLinksResponseDataAllOf) HasRemoteClusterId() bool {
	if o != nil && o.RemoteClusterId.IsSet() {
		return true
	}

	return false
}

// SetRemoteClusterId gets a reference to the given NullableString and assigns it to the RemoteClusterId field.
func (o *ListLinksResponseDataAllOf) SetRemoteClusterId(v string) {
	o.RemoteClusterId.Set(&v)
}

// SetRemoteClusterIdNil sets the value for RemoteClusterId to be an explicit nil
func (o *ListLinksResponseDataAllOf) SetRemoteClusterIdNil() {
	o.RemoteClusterId.Set(nil)
}

// UnsetRemoteClusterId ensures that no value is present for RemoteClusterId, not even an explicit nil
func (o *ListLinksResponseDataAllOf) UnsetRemoteClusterId() {
	o.RemoteClusterId.Unset()
}

// GetLinkName returns the LinkName field value
func (o *ListLinksResponseDataAllOf) GetLinkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkName
}

// GetLinkNameOk returns a tuple with the LinkName field value
// and a boolean to check if the value has been set.
func (o *ListLinksResponseDataAllOf) GetLinkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkName, true
}

// SetLinkName sets field value
func (o *ListLinksResponseDataAllOf) SetLinkName(v string) {
	o.LinkName = v
}

// GetLinkId returns the LinkId field value if set, zero value otherwise.
// Deprecated
func (o *ListLinksResponseDataAllOf) GetLinkId() string {
	if o == nil || o.LinkId == nil {
		var ret string
		return ret
	}
	return *o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ListLinksResponseDataAllOf) GetLinkIdOk() (*string, bool) {
	if o == nil || o.LinkId == nil {
		return nil, false
	}
	return o.LinkId, true
}

// HasLinkId returns a boolean if a field has been set.
func (o *ListLinksResponseDataAllOf) HasLinkId() bool {
	if o != nil && o.LinkId != nil {
		return true
	}

	return false
}

// SetLinkId gets a reference to the given string and assigns it to the LinkId field.
// Deprecated
func (o *ListLinksResponseDataAllOf) SetLinkId(v string) {
	o.LinkId = &v
}

// GetClusterLinkId returns the ClusterLinkId field value
func (o *ListLinksResponseDataAllOf) GetClusterLinkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterLinkId
}

// GetClusterLinkIdOk returns a tuple with the ClusterLinkId field value
// and a boolean to check if the value has been set.
func (o *ListLinksResponseDataAllOf) GetClusterLinkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterLinkId, true
}

// SetClusterLinkId sets field value
func (o *ListLinksResponseDataAllOf) SetClusterLinkId(v string) {
	o.ClusterLinkId = v
}

// GetTopicNames returns the TopicNames field value
func (o *ListLinksResponseDataAllOf) GetTopicNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TopicNames
}

// GetTopicNamesOk returns a tuple with the TopicNames field value
// and a boolean to check if the value has been set.
func (o *ListLinksResponseDataAllOf) GetTopicNamesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopicNames, true
}

// SetTopicNames sets field value
func (o *ListLinksResponseDataAllOf) SetTopicNames(v []string) {
	o.TopicNames = v
}

// GetLinkError returns the LinkError field value if set, zero value otherwise.
func (o *ListLinksResponseDataAllOf) GetLinkError() string {
	if o == nil || o.LinkError == nil {
		var ret string
		return ret
	}
	return *o.LinkError
}

// GetLinkErrorOk returns a tuple with the LinkError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLinksResponseDataAllOf) GetLinkErrorOk() (*string, bool) {
	if o == nil || o.LinkError == nil {
		return nil, false
	}
	return o.LinkError, true
}

// HasLinkError returns a boolean if a field has been set.
func (o *ListLinksResponseDataAllOf) HasLinkError() bool {
	if o != nil && o.LinkError != nil {
		return true
	}

	return false
}

// SetLinkError gets a reference to the given string and assigns it to the LinkError field.
func (o *ListLinksResponseDataAllOf) SetLinkError(v string) {
	o.LinkError = &v
}

// GetLinkErrorMessage returns the LinkErrorMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListLinksResponseDataAllOf) GetLinkErrorMessage() string {
	if o == nil || o.LinkErrorMessage.Get() == nil {
		var ret string
		return ret
	}
	return *o.LinkErrorMessage.Get()
}

// GetLinkErrorMessageOk returns a tuple with the LinkErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListLinksResponseDataAllOf) GetLinkErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LinkErrorMessage.Get(), o.LinkErrorMessage.IsSet()
}

// HasLinkErrorMessage returns a boolean if a field has been set.
func (o *ListLinksResponseDataAllOf) HasLinkErrorMessage() bool {
	if o != nil && o.LinkErrorMessage.IsSet() {
		return true
	}

	return false
}

// SetLinkErrorMessage gets a reference to the given NullableString and assigns it to the LinkErrorMessage field.
func (o *ListLinksResponseDataAllOf) SetLinkErrorMessage(v string) {
	o.LinkErrorMessage.Set(&v)
}

// SetLinkErrorMessageNil sets the value for LinkErrorMessage to be an explicit nil
func (o *ListLinksResponseDataAllOf) SetLinkErrorMessageNil() {
	o.LinkErrorMessage.Set(nil)
}

// UnsetLinkErrorMessage ensures that no value is present for LinkErrorMessage, not even an explicit nil
func (o *ListLinksResponseDataAllOf) UnsetLinkErrorMessage() {
	o.LinkErrorMessage.Unset()
}

// GetLinkState returns the LinkState field value if set, zero value otherwise.
func (o *ListLinksResponseDataAllOf) GetLinkState() string {
	if o == nil || o.LinkState == nil {
		var ret string
		return ret
	}
	return *o.LinkState
}

// GetLinkStateOk returns a tuple with the LinkState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLinksResponseDataAllOf) GetLinkStateOk() (*string, bool) {
	if o == nil || o.LinkState == nil {
		return nil, false
	}
	return o.LinkState, true
}

// HasLinkState returns a boolean if a field has been set.
func (o *ListLinksResponseDataAllOf) HasLinkState() bool {
	if o != nil && o.LinkState != nil {
		return true
	}

	return false
}

// SetLinkState gets a reference to the given string and assigns it to the LinkState field.
func (o *ListLinksResponseDataAllOf) SetLinkState(v string) {
	o.LinkState = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListLinksResponseDataAllOf) GetTasks() []LinkTask {
	if o == nil {
		var ret []LinkTask
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListLinksResponseDataAllOf) GetTasksOk() (*[]LinkTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return &o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *ListLinksResponseDataAllOf) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []LinkTask and assigns it to the Tasks field.
func (o *ListLinksResponseDataAllOf) SetTasks(v []LinkTask) {
	o.Tasks = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ListLinksResponseDataAllOf) Redact() {
	o.recurseRedact(o.SourceClusterId)
	o.recurseRedact(o.DestinationClusterId)
	o.recurseRedact(o.RemoteClusterId)
	o.recurseRedact(&o.LinkName)
	o.recurseRedact(o.LinkId)
	o.recurseRedact(&o.ClusterLinkId)
	o.recurseRedact(&o.TopicNames)
	o.recurseRedact(o.LinkError)
	o.recurseRedact(o.LinkErrorMessage)
	o.recurseRedact(o.LinkState)
	o.recurseRedact(o.Tasks)
}

func (o *ListLinksResponseDataAllOf) recurseRedact(v interface{}) {
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

func (o ListLinksResponseDataAllOf) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ListLinksResponseDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SourceClusterId.IsSet() {
		toSerialize["source_cluster_id"] = o.SourceClusterId.Get()
	}
	if o.DestinationClusterId.IsSet() {
		toSerialize["destination_cluster_id"] = o.DestinationClusterId.Get()
	}
	if o.RemoteClusterId.IsSet() {
		toSerialize["remote_cluster_id"] = o.RemoteClusterId.Get()
	}
	if true {
		toSerialize["link_name"] = o.LinkName
	}
	if o.LinkId != nil {
		toSerialize["link_id"] = o.LinkId
	}
	if true {
		toSerialize["cluster_link_id"] = o.ClusterLinkId
	}
	if true {
		toSerialize["topic_names"] = o.TopicNames
	}
	if o.LinkError != nil {
		toSerialize["link_error"] = o.LinkError
	}
	if o.LinkErrorMessage.IsSet() {
		toSerialize["link_error_message"] = o.LinkErrorMessage.Get()
	}
	if o.LinkState != nil {
		toSerialize["link_state"] = o.LinkState
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableListLinksResponseDataAllOf struct {
	value *ListLinksResponseDataAllOf
	isSet bool
}

func (v NullableListLinksResponseDataAllOf) Get() *ListLinksResponseDataAllOf {
	return v.value
}

func (v *NullableListLinksResponseDataAllOf) Set(val *ListLinksResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListLinksResponseDataAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListLinksResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLinksResponseDataAllOf(val *ListLinksResponseDataAllOf) *NullableListLinksResponseDataAllOf {
	return &NullableListLinksResponseDataAllOf{value: val, isSet: true}
}

func (v NullableListLinksResponseDataAllOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableListLinksResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}