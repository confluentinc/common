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
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// CmkV2ClusterUpdate `Clusters` objects represent Apache Kafka Clusters on Confluent Cloud.  The API allows you to list, create, read, update, and delete your Kafka clusters.   Related guide: [Confluent Cloud Cluster Management for Apache Kafka APIs](https://docs.confluent.io/cloud/current/clusters/cluster-api.html).  ## The Clusters Model <SchemaDefinition schemaRef=\"#/components/schemas/cmk.v2.Cluster\" />  ## Quotas and Limits This resource is subject to the [following quotas](https://docs.confluent.io/cloud/current/quotas/overview.html):  | Quota | Description | | --- | --- | | `kafka_clusters_per_environment` | Number of clusters in one Confluent Cloud environment |
type CmkV2ClusterUpdate struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string                 `json:"id,omitempty"`
	Metadata *ObjectMeta             `json:"metadata,omitempty"`
	Spec     *CmkV2ClusterSpecUpdate `json:"spec,omitempty"`
	Status   *CmkV2ClusterStatus     `json:"status,omitempty"`
}

// NewCmkV2ClusterUpdate instantiates a new CmkV2ClusterUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmkV2ClusterUpdate() *CmkV2ClusterUpdate {
	this := CmkV2ClusterUpdate{}
	return &this
}

// NewCmkV2ClusterUpdateWithDefaults instantiates a new CmkV2ClusterUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmkV2ClusterUpdateWithDefaults() *CmkV2ClusterUpdate {
	this := CmkV2ClusterUpdate{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *CmkV2ClusterUpdate) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterUpdate) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *CmkV2ClusterUpdate) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *CmkV2ClusterUpdate) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CmkV2ClusterUpdate) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterUpdate) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CmkV2ClusterUpdate) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CmkV2ClusterUpdate) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CmkV2ClusterUpdate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterUpdate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CmkV2ClusterUpdate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CmkV2ClusterUpdate) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CmkV2ClusterUpdate) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterUpdate) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CmkV2ClusterUpdate) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *CmkV2ClusterUpdate) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *CmkV2ClusterUpdate) GetSpec() CmkV2ClusterSpecUpdate {
	if o == nil || o.Spec == nil {
		var ret CmkV2ClusterSpecUpdate
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterUpdate) GetSpecOk() (*CmkV2ClusterSpecUpdate, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *CmkV2ClusterUpdate) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given CmkV2ClusterSpecUpdate and assigns it to the Spec field.
func (o *CmkV2ClusterUpdate) SetSpec(v CmkV2ClusterSpecUpdate) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CmkV2ClusterUpdate) GetStatus() CmkV2ClusterStatus {
	if o == nil || o.Status == nil {
		var ret CmkV2ClusterStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterUpdate) GetStatusOk() (*CmkV2ClusterStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CmkV2ClusterUpdate) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given CmkV2ClusterStatus and assigns it to the Status field.
func (o *CmkV2ClusterUpdate) SetStatus(v CmkV2ClusterStatus) {
	o.Status = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CmkV2ClusterUpdate) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.Spec)
	o.recurseRedact(o.Status)
}

func (o *CmkV2ClusterUpdate) recurseRedact(v interface{}) {
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

func (o CmkV2ClusterUpdate) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CmkV2ClusterUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableCmkV2ClusterUpdate struct {
	value *CmkV2ClusterUpdate
	isSet bool
}

func (v NullableCmkV2ClusterUpdate) Get() *CmkV2ClusterUpdate {
	return v.value
}

func (v *NullableCmkV2ClusterUpdate) Set(val *CmkV2ClusterUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableCmkV2ClusterUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableCmkV2ClusterUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmkV2ClusterUpdate(val *CmkV2ClusterUpdate) *NullableCmkV2ClusterUpdate {
	return &NullableCmkV2ClusterUpdate{value: val, isSet: true}
}

func (v NullableCmkV2ClusterUpdate) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableCmkV2ClusterUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
