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
	"encoding/json"
)

import (
	"reflect"
)

// ListLinksResponseData struct for ListLinksResponseData
type ListLinksResponseData struct {
	Kind                 string           `json:"kind"`
	Metadata             ResourceMetadata `json:"metadata"`
	SourceClusterId      NullableString   `json:"source_cluster_id,omitempty"`
	DestinationClusterId NullableString   `json:"destination_cluster_id,omitempty"`
	LinkName             string           `json:"link_name"`
	LinkId               string           `json:"link_id"`
	TopicsNames          *[]string        `json:"topics_names,omitempty"`
}

// NewListLinksResponseData instantiates a new ListLinksResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLinksResponseData(kind string, metadata ResourceMetadata, linkName string, linkId string) *ListLinksResponseData {
	this := ListLinksResponseData{}
	this.Kind = kind
	this.Metadata = metadata
	this.LinkName = linkName
	this.LinkId = linkId
	return &this
}

// NewListLinksResponseDataWithDefaults instantiates a new ListLinksResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLinksResponseDataWithDefaults() *ListLinksResponseData {
	this := ListLinksResponseData{}
	return &this
}

// GetKind returns the Kind field value
func (o *ListLinksResponseData) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ListLinksResponseData) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ListLinksResponseData) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *ListLinksResponseData) GetMetadata() ResourceMetadata {
	if o == nil {
		var ret ResourceMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ListLinksResponseData) GetMetadataOk() (*ResourceMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ListLinksResponseData) SetMetadata(v ResourceMetadata) {
	o.Metadata = v
}

// GetSourceClusterId returns the SourceClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListLinksResponseData) GetSourceClusterId() string {
	if o == nil || o.SourceClusterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceClusterId.Get()
}

// GetSourceClusterIdOk returns a tuple with the SourceClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListLinksResponseData) GetSourceClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceClusterId.Get(), o.SourceClusterId.IsSet()
}

// HasSourceClusterId returns a boolean if a field has been set.
func (o *ListLinksResponseData) HasSourceClusterId() bool {
	if o != nil && o.SourceClusterId.IsSet() {
		return true
	}

	return false
}

// SetSourceClusterId gets a reference to the given NullableString and assigns it to the SourceClusterId field.
func (o *ListLinksResponseData) SetSourceClusterId(v string) {
	o.SourceClusterId.Set(&v)
}

// SetSourceClusterIdNil sets the value for SourceClusterId to be an explicit nil
func (o *ListLinksResponseData) SetSourceClusterIdNil() {
	o.SourceClusterId.Set(nil)
}

// UnsetSourceClusterId ensures that no value is present for SourceClusterId, not even an explicit nil
func (o *ListLinksResponseData) UnsetSourceClusterId() {
	o.SourceClusterId.Unset()
}

// GetDestinationClusterId returns the DestinationClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListLinksResponseData) GetDestinationClusterId() string {
	if o == nil || o.DestinationClusterId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DestinationClusterId.Get()
}

// GetDestinationClusterIdOk returns a tuple with the DestinationClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListLinksResponseData) GetDestinationClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestinationClusterId.Get(), o.DestinationClusterId.IsSet()
}

// HasDestinationClusterId returns a boolean if a field has been set.
func (o *ListLinksResponseData) HasDestinationClusterId() bool {
	if o != nil && o.DestinationClusterId.IsSet() {
		return true
	}

	return false
}

// SetDestinationClusterId gets a reference to the given NullableString and assigns it to the DestinationClusterId field.
func (o *ListLinksResponseData) SetDestinationClusterId(v string) {
	o.DestinationClusterId.Set(&v)
}

// SetDestinationClusterIdNil sets the value for DestinationClusterId to be an explicit nil
func (o *ListLinksResponseData) SetDestinationClusterIdNil() {
	o.DestinationClusterId.Set(nil)
}

// UnsetDestinationClusterId ensures that no value is present for DestinationClusterId, not even an explicit nil
func (o *ListLinksResponseData) UnsetDestinationClusterId() {
	o.DestinationClusterId.Unset()
}

// GetLinkName returns the LinkName field value
func (o *ListLinksResponseData) GetLinkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkName
}

// GetLinkNameOk returns a tuple with the LinkName field value
// and a boolean to check if the value has been set.
func (o *ListLinksResponseData) GetLinkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkName, true
}

// SetLinkName sets field value
func (o *ListLinksResponseData) SetLinkName(v string) {
	o.LinkName = v
}

// GetLinkId returns the LinkId field value
func (o *ListLinksResponseData) GetLinkId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value
// and a boolean to check if the value has been set.
func (o *ListLinksResponseData) GetLinkIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkId, true
}

// SetLinkId sets field value
func (o *ListLinksResponseData) SetLinkId(v string) {
	o.LinkId = v
}

// GetTopicsNames returns the TopicsNames field value if set, zero value otherwise.
func (o *ListLinksResponseData) GetTopicsNames() []string {
	if o == nil || o.TopicsNames == nil {
		var ret []string
		return ret
	}
	return *o.TopicsNames
}

// GetTopicsNamesOk returns a tuple with the TopicsNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLinksResponseData) GetTopicsNamesOk() (*[]string, bool) {
	if o == nil || o.TopicsNames == nil {
		return nil, false
	}
	return o.TopicsNames, true
}

// HasTopicsNames returns a boolean if a field has been set.
func (o *ListLinksResponseData) HasTopicsNames() bool {
	if o != nil && o.TopicsNames != nil {
		return true
	}

	return false
}

// SetTopicsNames gets a reference to the given []string and assigns it to the TopicsNames field.
func (o *ListLinksResponseData) SetTopicsNames(v []string) {
	o.TopicsNames = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ListLinksResponseData) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(o.SourceClusterId)
	o.recurseRedact(o.DestinationClusterId)
	o.recurseRedact(&o.LinkName)
	o.recurseRedact(&o.LinkId)
	o.recurseRedact(o.TopicsNames)
}

func (o *ListLinksResponseData) recurseRedact(v interface{}) {
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

func (o ListLinksResponseData) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ListLinksResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if o.SourceClusterId.IsSet() {
		toSerialize["source_cluster_id"] = o.SourceClusterId.Get()
	}
	if o.DestinationClusterId.IsSet() {
		toSerialize["destination_cluster_id"] = o.DestinationClusterId.Get()
	}
	if true {
		toSerialize["link_name"] = o.LinkName
	}
	if true {
		toSerialize["link_id"] = o.LinkId
	}
	if o.TopicsNames != nil {
		toSerialize["topics_names"] = o.TopicsNames
	}
	return json.Marshal(toSerialize)
}

type NullableListLinksResponseData struct {
	value *ListLinksResponseData
	isSet bool
}

func (v NullableListLinksResponseData) Get() *ListLinksResponseData {
	return v.value
}

func (v *NullableListLinksResponseData) Set(val *ListLinksResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableListLinksResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableListLinksResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLinksResponseData(val *ListLinksResponseData) *NullableListLinksResponseData {
	return &NullableListLinksResponseData{value: val, isSet: true}
}

func (v NullableListLinksResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLinksResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
