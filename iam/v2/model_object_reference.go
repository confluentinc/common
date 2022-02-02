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
Identity & Access Management API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha0
Contact: paas-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

import (
	"reflect"
)

// ObjectReference ObjectReference provides information for you to locate the referred object
type ObjectReference struct {
	// ID of the referred resource
	Id string `json:"id"`
	// Environment of the referred resource, if env-scoped
	Environment *string `json:"environment,omitempty"`
	// API URL for accessing or modifying the referred object
	Related string `json:"related"`
	// CRN reference to the referred resource
	ResourceName string `json:"resource_name"`
	// API group and version of the referred resource
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind of the referred resource
	Kind *string `json:"kind,omitempty"`
}

// NewObjectReference instantiates a new ObjectReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectReference(id string, related string, resourceName string) *ObjectReference {
	this := ObjectReference{}
	this.Id = id
	this.Related = related
	this.ResourceName = resourceName
	return &this
}

// NewObjectReferenceWithDefaults instantiates a new ObjectReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectReferenceWithDefaults() *ObjectReference {
	this := ObjectReference{}
	return &this
}

// GetId returns the Id field value
func (o *ObjectReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObjectReference) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ObjectReference) SetId(v string) {
	o.Id = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *ObjectReference) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectReference) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *ObjectReference) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *ObjectReference) SetEnvironment(v string) {
	o.Environment = &v
}

// GetRelated returns the Related field value
func (o *ObjectReference) GetRelated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *ObjectReference) GetRelatedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value
func (o *ObjectReference) SetRelated(v string) {
	o.Related = v
}

// GetResourceName returns the ResourceName field value
func (o *ObjectReference) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *ObjectReference) GetResourceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *ObjectReference) SetResourceName(v string) {
	o.ResourceName = v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ObjectReference) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectReference) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ObjectReference) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ObjectReference) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ObjectReference) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectReference) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ObjectReference) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ObjectReference) SetKind(v string) {
	o.Kind = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ObjectReference) Redact() {
    o.recurseRedact(&o.Id)
    o.recurseRedact(o.Environment)
    o.recurseRedact(&o.Related)
    o.recurseRedact(&o.ResourceName)
    o.recurseRedact(o.ApiVersion)
    o.recurseRedact(o.Kind)
}

func (o *ObjectReference) recurseRedact(v interface{}) {
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

func (o ObjectReference) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o ObjectReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if true {
		toSerialize["related"] = o.Related
	}
	if true {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	return json.Marshal(toSerialize)
}

type NullableObjectReference struct {
	value *ObjectReference
	isSet bool
}

func (v NullableObjectReference) Get() *ObjectReference {
	return v.value
}

func (v *NullableObjectReference) Set(val *ObjectReference) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectReference) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectReference(val *ObjectReference) *NullableObjectReference {
	return &NullableObjectReference{value: val, isSet: true}
}

func (v NullableObjectReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


