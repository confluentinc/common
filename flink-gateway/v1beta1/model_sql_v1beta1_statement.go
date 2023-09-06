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
SQL API v1beta1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1beta1

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// SqlV1beta1Statement `Statement` represents a core resource used to model SQL statements for execution. A statement generalizes DDL, DML, DQL, etc., but doesn’t attempt to handle session management or any higher-level functionality. The API allows you to list, create, read, and delete your statements. ## The Statements Model <SchemaDefinition schemaRef=\"#/components/schemas/sql.v1beta1.Statement\" />
type SqlV1beta1Statement struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind     *string     `json:"kind,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// The user provided name of the resource, unique within this environment.
	Name *string `json:"name,omitempty"`
	// The unique identifier for the organization.
	OrganizationId *string `json:"organization_id,omitempty"`
	// The unique identifier for the environment.
	EnvironmentId *string                    `json:"environment_id,omitempty"`
	Spec          *SqlV1beta1StatementSpec   `json:"spec,omitempty"`
	Status        *SqlV1beta1StatementStatus `json:"status,omitempty"`
}

// NewSqlV1beta1Statement instantiates a new SqlV1beta1Statement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqlV1beta1Statement() *SqlV1beta1Statement {
	this := SqlV1beta1Statement{}
	return &this
}

// NewSqlV1beta1StatementWithDefaults instantiates a new SqlV1beta1Statement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqlV1beta1StatementWithDefaults() *SqlV1beta1Statement {
	this := SqlV1beta1Statement{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *SqlV1beta1Statement) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1beta1Statement) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *SqlV1beta1Statement) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *SqlV1beta1Statement) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *SqlV1beta1Statement) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1beta1Statement) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *SqlV1beta1Statement) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *SqlV1beta1Statement) SetKind(v string) {
	o.Kind = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *SqlV1beta1Statement) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1beta1Statement) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *SqlV1beta1Statement) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *SqlV1beta1Statement) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SqlV1beta1Statement) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1beta1Statement) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SqlV1beta1Statement) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SqlV1beta1Statement) SetName(v string) {
	o.Name = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *SqlV1beta1Statement) GetOrganizationId() string {
	if o == nil || o.OrganizationId == nil {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1beta1Statement) GetOrganizationIdOk() (*string, bool) {
	if o == nil || o.OrganizationId == nil {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *SqlV1beta1Statement) HasOrganizationId() bool {
	if o != nil && o.OrganizationId != nil {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *SqlV1beta1Statement) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *SqlV1beta1Statement) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1beta1Statement) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || o.EnvironmentId == nil {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *SqlV1beta1Statement) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId != nil {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *SqlV1beta1Statement) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *SqlV1beta1Statement) GetSpec() SqlV1beta1StatementSpec {
	if o == nil || o.Spec == nil {
		var ret SqlV1beta1StatementSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1beta1Statement) GetSpecOk() (*SqlV1beta1StatementSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *SqlV1beta1Statement) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given SqlV1beta1StatementSpec and assigns it to the Spec field.
func (o *SqlV1beta1Statement) SetSpec(v SqlV1beta1StatementSpec) {
	o.Spec = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SqlV1beta1Statement) GetStatus() SqlV1beta1StatementStatus {
	if o == nil || o.Status == nil {
		var ret SqlV1beta1StatementStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1beta1Statement) GetStatusOk() (*SqlV1beta1StatementStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SqlV1beta1Statement) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SqlV1beta1StatementStatus and assigns it to the Status field.
func (o *SqlV1beta1Statement) SetStatus(v SqlV1beta1StatementStatus) {
	o.Status = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *SqlV1beta1Statement) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.Name)
	o.recurseRedact(o.OrganizationId)
	o.recurseRedact(o.EnvironmentId)
	o.recurseRedact(o.Spec)
	o.recurseRedact(o.Status)
}

func (o *SqlV1beta1Statement) recurseRedact(v interface{}) {
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

func (o SqlV1beta1Statement) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o SqlV1beta1Statement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrganizationId != nil {
		toSerialize["organization_id"] = o.OrganizationId
	}
	if o.EnvironmentId != nil {
		toSerialize["environment_id"] = o.EnvironmentId
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

type NullableSqlV1beta1Statement struct {
	value *SqlV1beta1Statement
	isSet bool
}

func (v NullableSqlV1beta1Statement) Get() *SqlV1beta1Statement {
	return v.value
}

func (v *NullableSqlV1beta1Statement) Set(val *SqlV1beta1Statement) {
	v.value = val
	v.isSet = true
}

func (v NullableSqlV1beta1Statement) IsSet() bool {
	return v.isSet
}

func (v *NullableSqlV1beta1Statement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqlV1beta1Statement(val *SqlV1beta1Statement) *NullableSqlV1beta1Statement {
	return &NullableSqlV1beta1Statement{value: val, isSet: true}
}

func (v NullableSqlV1beta1Statement) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableSqlV1beta1Statement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
