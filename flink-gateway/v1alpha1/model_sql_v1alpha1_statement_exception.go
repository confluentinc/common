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
SQL API v1alpha1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha1

import (
	"bytes"
	"encoding/json"
	"time"
)

import (
	"reflect"
)

// SqlV1alpha1StatementException struct for SqlV1alpha1StatementException
type SqlV1alpha1StatementException struct {
	// Name of the SQL statement exception.
	Name *string `json:"name,omitempty"`
	// Stack trace of the statement exception.
	Stacktrace *string `json:"stacktrace,omitempty"`
	// The date and time at which the exception occurred. It is represented in RFC3339 format and is in UTC.
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewSqlV1alpha1StatementException instantiates a new SqlV1alpha1StatementException object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSqlV1alpha1StatementException() *SqlV1alpha1StatementException {
	this := SqlV1alpha1StatementException{}
	return &this
}

// NewSqlV1alpha1StatementExceptionWithDefaults instantiates a new SqlV1alpha1StatementException object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSqlV1alpha1StatementExceptionWithDefaults() *SqlV1alpha1StatementException {
	this := SqlV1alpha1StatementException{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SqlV1alpha1StatementException) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1alpha1StatementException) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SqlV1alpha1StatementException) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SqlV1alpha1StatementException) SetName(v string) {
	o.Name = &v
}

// GetStacktrace returns the Stacktrace field value if set, zero value otherwise.
func (o *SqlV1alpha1StatementException) GetStacktrace() string {
	if o == nil || o.Stacktrace == nil {
		var ret string
		return ret
	}
	return *o.Stacktrace
}

// GetStacktraceOk returns a tuple with the Stacktrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1alpha1StatementException) GetStacktraceOk() (*string, bool) {
	if o == nil || o.Stacktrace == nil {
		return nil, false
	}
	return o.Stacktrace, true
}

// HasStacktrace returns a boolean if a field has been set.
func (o *SqlV1alpha1StatementException) HasStacktrace() bool {
	if o != nil && o.Stacktrace != nil {
		return true
	}

	return false
}

// SetStacktrace gets a reference to the given string and assigns it to the Stacktrace field.
func (o *SqlV1alpha1StatementException) SetStacktrace(v string) {
	o.Stacktrace = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *SqlV1alpha1StatementException) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SqlV1alpha1StatementException) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *SqlV1alpha1StatementException) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *SqlV1alpha1StatementException) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *SqlV1alpha1StatementException) Redact() {
	o.recurseRedact(o.Name)
	o.recurseRedact(o.Stacktrace)
	o.recurseRedact(o.Timestamp)
}

func (o *SqlV1alpha1StatementException) recurseRedact(v interface{}) {
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

func (o SqlV1alpha1StatementException) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o SqlV1alpha1StatementException) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Stacktrace != nil {
		toSerialize["stacktrace"] = o.Stacktrace
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableSqlV1alpha1StatementException struct {
	value *SqlV1alpha1StatementException
	isSet bool
}

func (v NullableSqlV1alpha1StatementException) Get() *SqlV1alpha1StatementException {
	return v.value
}

func (v *NullableSqlV1alpha1StatementException) Set(val *SqlV1alpha1StatementException) {
	v.value = val
	v.isSet = true
}

func (v NullableSqlV1alpha1StatementException) IsSet() bool {
	return v.isSet
}

func (v *NullableSqlV1alpha1StatementException) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSqlV1alpha1StatementException(val *SqlV1alpha1StatementException) *NullableSqlV1alpha1StatementException {
	return &NullableSqlV1alpha1StatementException{value: val, isSet: true}
}

func (v NullableSqlV1alpha1StatementException) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableSqlV1alpha1StatementException) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
