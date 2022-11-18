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

// SrcmV2ClusterSpecUpdate The desired state of the Cluster
type SrcmV2ClusterSpecUpdate struct {
	// The billing package.  Note: Clusters can be upgraded from ESSENTIALS to ADVANCED, but cannot be downgraded from ADVANCED to ESSENTIALS. 
	Package *string `json:"package,omitempty"`
	// The environment to which this belongs.
	Environment *GlobalObjectReference `json:"environment,omitempty"`
}

// NewSrcmV2ClusterSpecUpdate instantiates a new SrcmV2ClusterSpecUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSrcmV2ClusterSpecUpdate() *SrcmV2ClusterSpecUpdate {
	this := SrcmV2ClusterSpecUpdate{}
	return &this
}

// NewSrcmV2ClusterSpecUpdateWithDefaults instantiates a new SrcmV2ClusterSpecUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSrcmV2ClusterSpecUpdateWithDefaults() *SrcmV2ClusterSpecUpdate {
	this := SrcmV2ClusterSpecUpdate{}
	return &this
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *SrcmV2ClusterSpecUpdate) GetPackage() string {
	if o == nil || o.Package == nil {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SrcmV2ClusterSpecUpdate) GetPackageOk() (*string, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *SrcmV2ClusterSpecUpdate) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *SrcmV2ClusterSpecUpdate) SetPackage(v string) {
	o.Package = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SrcmV2ClusterSpecUpdate) GetEnvironment() GlobalObjectReference {
	if o == nil || o.Environment == nil {
		var ret GlobalObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SrcmV2ClusterSpecUpdate) GetEnvironmentOk() (*GlobalObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SrcmV2ClusterSpecUpdate) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given GlobalObjectReference and assigns it to the Environment field.
func (o *SrcmV2ClusterSpecUpdate) SetEnvironment(v GlobalObjectReference) {
	o.Environment = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *SrcmV2ClusterSpecUpdate) Redact() {
    o.recurseRedact(o.Package)
    o.recurseRedact(o.Environment)
}

func (o *SrcmV2ClusterSpecUpdate) recurseRedact(v interface{}) {
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

func (o SrcmV2ClusterSpecUpdate) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o SrcmV2ClusterSpecUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Package != nil {
		toSerialize["package"] = o.Package
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	return json.Marshal(toSerialize)
}

type NullableSrcmV2ClusterSpecUpdate struct {
	value *SrcmV2ClusterSpecUpdate
	isSet bool
}

func (v NullableSrcmV2ClusterSpecUpdate) Get() *SrcmV2ClusterSpecUpdate {
	return v.value
}

func (v *NullableSrcmV2ClusterSpecUpdate) Set(val *SrcmV2ClusterSpecUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSrcmV2ClusterSpecUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSrcmV2ClusterSpecUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSrcmV2ClusterSpecUpdate(val *SrcmV2ClusterSpecUpdate) *NullableSrcmV2ClusterSpecUpdate {
	return &NullableSrcmV2ClusterSpecUpdate{value: val, isSet: true}
}

func (v NullableSrcmV2ClusterSpecUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSrcmV2ClusterSpecUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


