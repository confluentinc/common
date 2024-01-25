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
Confluent Schema Registry APIs

REST API for the Schema Registry

API version: 1.0.0
Contact: data-governance@confluent.io
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

// CreateDekRequest struct for CreateDekRequest
type CreateDekRequest struct {
	Subject              *string `json:"subject,omitempty"`
	Version              *int32  `json:"version,omitempty"`
	Algorithm            *string `json:"algorithm,omitempty"`
	EncryptedKeyMaterial *string `json:"encryptedKeyMaterial,omitempty"`
}

// NewCreateDekRequest instantiates a new CreateDekRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDekRequest() *CreateDekRequest {
	this := CreateDekRequest{}
	return &this
}

// NewCreateDekRequestWithDefaults instantiates a new CreateDekRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDekRequestWithDefaults() *CreateDekRequest {
	this := CreateDekRequest{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CreateDekRequest) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDekRequest) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CreateDekRequest) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CreateDekRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *CreateDekRequest) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDekRequest) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *CreateDekRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *CreateDekRequest) SetVersion(v int32) {
	o.Version = &v
}

// GetAlgorithm returns the Algorithm field value if set, zero value otherwise.
func (o *CreateDekRequest) GetAlgorithm() string {
	if o == nil || o.Algorithm == nil {
		var ret string
		return ret
	}
	return *o.Algorithm
}

// GetAlgorithmOk returns a tuple with the Algorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDekRequest) GetAlgorithmOk() (*string, bool) {
	if o == nil || o.Algorithm == nil {
		return nil, false
	}
	return o.Algorithm, true
}

// HasAlgorithm returns a boolean if a field has been set.
func (o *CreateDekRequest) HasAlgorithm() bool {
	if o != nil && o.Algorithm != nil {
		return true
	}

	return false
}

// SetAlgorithm gets a reference to the given string and assigns it to the Algorithm field.
func (o *CreateDekRequest) SetAlgorithm(v string) {
	o.Algorithm = &v
}

// GetEncryptedKeyMaterial returns the EncryptedKeyMaterial field value if set, zero value otherwise.
func (o *CreateDekRequest) GetEncryptedKeyMaterial() string {
	if o == nil || o.EncryptedKeyMaterial == nil {
		var ret string
		return ret
	}
	return *o.EncryptedKeyMaterial
}

// GetEncryptedKeyMaterialOk returns a tuple with the EncryptedKeyMaterial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDekRequest) GetEncryptedKeyMaterialOk() (*string, bool) {
	if o == nil || o.EncryptedKeyMaterial == nil {
		return nil, false
	}
	return o.EncryptedKeyMaterial, true
}

// HasEncryptedKeyMaterial returns a boolean if a field has been set.
func (o *CreateDekRequest) HasEncryptedKeyMaterial() bool {
	if o != nil && o.EncryptedKeyMaterial != nil {
		return true
	}

	return false
}

// SetEncryptedKeyMaterial gets a reference to the given string and assigns it to the EncryptedKeyMaterial field.
func (o *CreateDekRequest) SetEncryptedKeyMaterial(v string) {
	o.EncryptedKeyMaterial = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CreateDekRequest) Redact() {
	o.recurseRedact(o.Subject)
	o.recurseRedact(o.Version)
	o.recurseRedact(o.Algorithm)
	o.recurseRedact(o.EncryptedKeyMaterial)
}

func (o *CreateDekRequest) recurseRedact(v interface{}) {
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

func (o CreateDekRequest) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CreateDekRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Algorithm != nil {
		toSerialize["algorithm"] = o.Algorithm
	}
	if o.EncryptedKeyMaterial != nil {
		toSerialize["encryptedKeyMaterial"] = o.EncryptedKeyMaterial
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableCreateDekRequest struct {
	value *CreateDekRequest
	isSet bool
}

func (v NullableCreateDekRequest) Get() *CreateDekRequest {
	return v.value
}

func (v *NullableCreateDekRequest) Set(val *CreateDekRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDekRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDekRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDekRequest(val *CreateDekRequest) *NullableCreateDekRequest {
	return &NullableCreateDekRequest{value: val, isSet: true}
}

func (v NullableCreateDekRequest) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableCreateDekRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
