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
Organization Management API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha0
Contact: paas-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// OrgV2EnvironmentList `Environment` objects represent an isolated namespace for your Confluent resources for organizational purposes.  The API allows you to create, delete, and update your environments. You can retrieve individual environments as well as a list of all your environments.   Related guide: [Environments in Confluent Cloud](https://docs.confluent.io/cloud/current/access-management/environments.html).  ## The Environments Model <SchemaDefinition schemaRef=\"#/components/schemas/org.v2.Environment\" />  ## Quotas and Limits This resource is subject to the [following quotas](https://docs.confluent.io/cloud/current/quotas/overview.html):  | Quota | Description | | --- | --- | | `environments_per_org` | Environments in one Confluent Cloud organization |
type OrgV2EnvironmentList struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind     string   `json:"kind,omitempty"`
	Metadata ListMeta `json:"metadata,omitempty"`
	// A data property that contains an array of resource items. Each entry in the array is a separate resource.
	Data []OrgV2Environment `json:"data,omitempty"`
}

// NewOrgV2EnvironmentList instantiates a new OrgV2EnvironmentList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgV2EnvironmentList(apiVersion string, kind string, metadata ListMeta, data []OrgV2Environment) *OrgV2EnvironmentList {
	this := OrgV2EnvironmentList{}
	this.ApiVersion = apiVersion
	this.Kind = kind
	this.Metadata = metadata
	this.Data = data
	return &this
}

// NewOrgV2EnvironmentListWithDefaults instantiates a new OrgV2EnvironmentList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgV2EnvironmentListWithDefaults() *OrgV2EnvironmentList {
	this := OrgV2EnvironmentList{}
	return &this
}

// GetApiVersion returns the ApiVersion field value
func (o *OrgV2EnvironmentList) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *OrgV2EnvironmentList) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiVersion, true
}

// SetApiVersion sets field value
func (o *OrgV2EnvironmentList) SetApiVersion(v string) {
	o.ApiVersion = v
}

// GetKind returns the Kind field value
func (o *OrgV2EnvironmentList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *OrgV2EnvironmentList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *OrgV2EnvironmentList) SetKind(v string) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value
func (o *OrgV2EnvironmentList) GetMetadata() ListMeta {
	if o == nil {
		var ret ListMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *OrgV2EnvironmentList) GetMetadataOk() (*ListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *OrgV2EnvironmentList) SetMetadata(v ListMeta) {
	o.Metadata = v
}

// GetData returns the Data field value
func (o *OrgV2EnvironmentList) GetData() []OrgV2Environment {
	if o == nil {
		var ret []OrgV2Environment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *OrgV2EnvironmentList) GetDataOk() (*[]OrgV2Environment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *OrgV2EnvironmentList) SetData(v []OrgV2Environment) {
	o.Data = v
}

// Redact resets all sensitive fields to their zero value.
func (o *OrgV2EnvironmentList) Redact() {
	o.recurseRedact(&o.ApiVersion)
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Metadata)
	o.recurseRedact(&o.Data)
}

func (o *OrgV2EnvironmentList) recurseRedact(v interface{}) {
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

func (o OrgV2EnvironmentList) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o OrgV2EnvironmentList) MarshalJSON() ([]byte, error) {
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
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableOrgV2EnvironmentList struct {
	value *OrgV2EnvironmentList
	isSet bool
}

func (v NullableOrgV2EnvironmentList) Get() *OrgV2EnvironmentList {
	return v.value
}

func (v *NullableOrgV2EnvironmentList) Set(val *OrgV2EnvironmentList) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgV2EnvironmentList) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgV2EnvironmentList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgV2EnvironmentList(val *OrgV2EnvironmentList) *NullableOrgV2EnvironmentList {
	return &NullableOrgV2EnvironmentList{value: val, isSet: true}
}

func (v NullableOrgV2EnvironmentList) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableOrgV2EnvironmentList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}