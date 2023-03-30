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
Network API

Network API

API version: 0.0.1-alpha1
Contact: cire-traffic@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// EnvScopedObjectReference ObjectReference provides information for you to locate the referred object
type EnvScopedObjectReference struct {
	// ID of the referred resource
	Id string `json:"id"`
	// Environment of the referred resource, if env-scoped
	Environment *string `json:"environment,omitempty"`
	// API URL for accessing or modifying the referred object
	Related string `json:"related"`
	// CRN reference to the referred resource
	ResourceName string `json:"resource_name"`
}

// NewEnvScopedObjectReference instantiates a new EnvScopedObjectReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvScopedObjectReference(id string, related string, resourceName string) *EnvScopedObjectReference {
	this := EnvScopedObjectReference{}
	this.Id = id
	this.Related = related
	this.ResourceName = resourceName
	return &this
}

// NewEnvScopedObjectReferenceWithDefaults instantiates a new EnvScopedObjectReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvScopedObjectReferenceWithDefaults() *EnvScopedObjectReference {
	this := EnvScopedObjectReference{}
	return &this
}

// GetId returns the Id field value
func (o *EnvScopedObjectReference) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EnvScopedObjectReference) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EnvScopedObjectReference) SetId(v string) {
	o.Id = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *EnvScopedObjectReference) GetEnvironment() string {
	if o == nil || o.Environment == nil {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvScopedObjectReference) GetEnvironmentOk() (*string, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *EnvScopedObjectReference) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *EnvScopedObjectReference) SetEnvironment(v string) {
	o.Environment = &v
}

// GetRelated returns the Related field value
func (o *EnvScopedObjectReference) GetRelated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Related
}

// GetRelatedOk returns a tuple with the Related field value
// and a boolean to check if the value has been set.
func (o *EnvScopedObjectReference) GetRelatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Related, true
}

// SetRelated sets field value
func (o *EnvScopedObjectReference) SetRelated(v string) {
	o.Related = v
}

// GetResourceName returns the ResourceName field value
func (o *EnvScopedObjectReference) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *EnvScopedObjectReference) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *EnvScopedObjectReference) SetResourceName(v string) {
	o.ResourceName = v
}

// Redact resets all sensitive fields to their zero value.
func (o *EnvScopedObjectReference) Redact() {
	o.recurseRedact(&o.Id)
	o.recurseRedact(o.Environment)
	o.recurseRedact(&o.Related)
	o.recurseRedact(&o.ResourceName)
}

func (o *EnvScopedObjectReference) recurseRedact(v interface{}) {
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

func (o EnvScopedObjectReference) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o EnvScopedObjectReference) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableEnvScopedObjectReference struct {
	value *EnvScopedObjectReference
	isSet bool
}

func (v NullableEnvScopedObjectReference) Get() *EnvScopedObjectReference {
	return v.value
}

func (v *NullableEnvScopedObjectReference) Set(val *EnvScopedObjectReference) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvScopedObjectReference) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvScopedObjectReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvScopedObjectReference(val *EnvScopedObjectReference) *NullableEnvScopedObjectReference {
	return &NullableEnvScopedObjectReference{value: val, isSet: true}
}

func (v NullableEnvScopedObjectReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvScopedObjectReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
