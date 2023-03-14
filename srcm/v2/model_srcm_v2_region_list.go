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
Cluster Management for Schema Registry API

Cluster Management for Schema Registry API

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

// SrcmV2RegionList `Region` objects represent cloud provider regions available when placing Schema Registry clusters. The API allows you to list Schema Registry regions.   Related guide: [Confluent Cloud Schema Registry Regions](https://docs.confluent.io/cloud/current/stream-governance/clusters-regions-api.html#schema-registry-regions).  ## The Regions Model <SchemaDefinition schemaRef=\"#/components/schemas/srcm.v2.Region\" />
type SrcmV2RegionList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version"`
	// Kind defines the object this REST resource represents.
	Kind     string   `json:"kind"`
	Metadata ListMeta `json:"metadata"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []SrcmV2Region `json:"data"`
}

// NewSrcmV2RegionList instantiates a new SrcmV2RegionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSrcmV2RegionList(apiVersion string, kind string, metadata ListMeta, data []SrcmV2Region) *SrcmV2RegionList {
	this := SrcmV2RegionList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewSrcmV2RegionListWithDefaults instantiates a new SrcmV2RegionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSrcmV2RegionListWithDefaults() *SrcmV2RegionList {
	this := SrcmV2RegionList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *SrcmV2RegionList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *SrcmV2RegionList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *SrcmV2RegionList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *SrcmV2RegionList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *SrcmV2RegionList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *SrcmV2RegionList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *SrcmV2RegionList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *SrcmV2RegionList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *SrcmV2RegionList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *SrcmV2RegionList) GetData() []SrcmV2Region {
	if o == nil {
		var ret []SrcmV2Region
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SrcmV2RegionList) GetDataOk() (*[]SrcmV2Region, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SrcmV2RegionList) SetData(v []SrcmV2Region) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *SrcmV2RegionList) Redact() {
	o.recurseRedact(&o.ApiVersion)
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *SrcmV2RegionList) recurseRedact(v interface{}) {
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

func (o SrcmV2RegionList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o SrcmV2RegionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["api_version"] = o.ApiVersion
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSrcmV2RegionList struct {
	value *SrcmV2RegionList
	isSet bool
}

func (v NullableSrcmV2RegionList) Get() *SrcmV2RegionList {
	return v.value
}

func (v *NullableSrcmV2RegionList) Set(val *SrcmV2RegionList) {
	v.value = val
	v.isSet = true
}

func (v NullableSrcmV2RegionList) IsSet() bool {
	return v.isSet
}

func (v *NullableSrcmV2RegionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSrcmV2RegionList(val *SrcmV2RegionList) *NullableSrcmV2RegionList {
	return &NullableSrcmV2RegionList{value: val, isSet: true}
}

func (v NullableSrcmV2RegionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSrcmV2RegionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
