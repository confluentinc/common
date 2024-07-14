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
Billing API

Confluent Cloud Billing API 

API version: 0.0.1-alpha0
Contact: monetization-eng@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// BillingV1Environment The details of the environment for a given resource. 
type BillingV1Environment struct {
	// ID of the environment.
	Id *string `json:"id,omitempty"`
}

// NewBillingV1Environment instantiates a new BillingV1Environment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingV1Environment() *BillingV1Environment {
	this := BillingV1Environment{}
	return &this
}

// NewBillingV1EnvironmentWithDefaults instantiates a new BillingV1Environment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingV1EnvironmentWithDefaults() *BillingV1Environment {
	this := BillingV1Environment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BillingV1Environment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingV1Environment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BillingV1Environment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BillingV1Environment) SetId(v string) {
	o.Id = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *BillingV1Environment) Redact() {
    o.recurseRedact(o.Id)
}

func (o *BillingV1Environment) recurseRedact(v interface{}) {
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

func (o BillingV1Environment) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o BillingV1Environment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableBillingV1Environment struct {
	value *BillingV1Environment
	isSet bool
}

func (v NullableBillingV1Environment) Get() *BillingV1Environment {
	return v.value
}

func (v *NullableBillingV1Environment) Set(val *BillingV1Environment) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingV1Environment) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingV1Environment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingV1Environment(val *BillingV1Environment) *NullableBillingV1Environment {
	return &NullableBillingV1Environment{value: val, isSet: true}
}

func (v NullableBillingV1Environment) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableBillingV1Environment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

