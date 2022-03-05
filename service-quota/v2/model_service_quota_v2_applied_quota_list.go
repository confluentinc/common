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
Service Quota API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha0
Contact: api-framework-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

import (
	"reflect"
)

// ServiceQuotaV2AppliedQuotaList A `quota` object represents a quota configuration for a specific Confluent Cloud resource. Use this API to retrieve an individual quota or list of quotas for a given scope.   Related guide: [Service Quotas for Confluent Cloud](https://docs.confluent.io/cloud/current/quotas/index.html).  ## The Applied Quotas Model <SchemaDefinition schemaRef=\"#/components/schemas/service-quota.v2.AppliedQuota\" />
type ServiceQuotaV2AppliedQuotaList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version"`
	// Kind defines the object this REST resource represents.
	Kind     string   `json:"kind"`
	Metadata ListMeta `json:"metadata"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []ServiceQuotaV2AppliedQuota `json:"data"`
}

// NewServiceQuotaV2AppliedQuotaList instantiates a new ServiceQuotaV2AppliedQuotaList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceQuotaV2AppliedQuotaList(apiVersion string, kind string, metadata ListMeta, data []ServiceQuotaV2AppliedQuota) *ServiceQuotaV2AppliedQuotaList {
	this := ServiceQuotaV2AppliedQuotaList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewServiceQuotaV2AppliedQuotaListWithDefaults instantiates a new ServiceQuotaV2AppliedQuotaList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceQuotaV2AppliedQuotaListWithDefaults() *ServiceQuotaV2AppliedQuotaList {
	this := ServiceQuotaV2AppliedQuotaList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *ServiceQuotaV2AppliedQuotaList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV2AppliedQuotaList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *ServiceQuotaV2AppliedQuotaList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *ServiceQuotaV2AppliedQuotaList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV2AppliedQuotaList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *ServiceQuotaV2AppliedQuotaList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *ServiceQuotaV2AppliedQuotaList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV2AppliedQuotaList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ServiceQuotaV2AppliedQuotaList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *ServiceQuotaV2AppliedQuotaList) GetData() []ServiceQuotaV2AppliedQuota {
	if o == nil {
		var ret []ServiceQuotaV2AppliedQuota
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ServiceQuotaV2AppliedQuotaList) GetDataOk() (*[]ServiceQuotaV2AppliedQuota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ServiceQuotaV2AppliedQuotaList) SetData(v []ServiceQuotaV2AppliedQuota) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *ServiceQuotaV2AppliedQuotaList) Redact() {
	o.recurseRedact(&o.ApiVersion)
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *ServiceQuotaV2AppliedQuotaList) recurseRedact(v interface{}) {
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

func (o ServiceQuotaV2AppliedQuotaList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ServiceQuotaV2AppliedQuotaList) MarshalJSON() ([]byte, error) {
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

type NullableServiceQuotaV2AppliedQuotaList struct {
	value *ServiceQuotaV2AppliedQuotaList
	isSet bool
}

func (v NullableServiceQuotaV2AppliedQuotaList) Get() *ServiceQuotaV2AppliedQuotaList {
	return v.value
}

func (v *NullableServiceQuotaV2AppliedQuotaList) Set(val *ServiceQuotaV2AppliedQuotaList) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceQuotaV2AppliedQuotaList) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceQuotaV2AppliedQuotaList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceQuotaV2AppliedQuotaList(val *ServiceQuotaV2AppliedQuotaList) *NullableServiceQuotaV2AppliedQuotaList {
	return &NullableServiceQuotaV2AppliedQuotaList{value: val, isSet: true}
}

func (v NullableServiceQuotaV2AppliedQuotaList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceQuotaV2AppliedQuotaList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
