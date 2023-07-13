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

// ExporterStatusResponse Exporter status get request
type ExporterStatusResponse struct {
	// Name of exporter.
	Name *string `json:"name,omitempty"`
	// State of the exporter. Could be STARTING, RUNNING or PAUSED
	State *string `json:"state,omitempty"`
	// Offset of the exporter
	Offset *int32 `json:"offset,omitempty"`
	// Timestamp of the exporter
	Ts *int32 `json:"ts,omitempty"`
	// Error trace of the exporter
	Trace *string `json:"trace,omitempty"`
}

// NewExporterStatusResponse instantiates a new ExporterStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExporterStatusResponse() *ExporterStatusResponse {
	this := ExporterStatusResponse{}
	return &this
}

// NewExporterStatusResponseWithDefaults instantiates a new ExporterStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExporterStatusResponseWithDefaults() *ExporterStatusResponse {
	this := ExporterStatusResponse{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExporterStatusResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExporterStatusResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExporterStatusResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ExporterStatusResponse) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ExporterStatusResponse) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExporterStatusResponse) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ExporterStatusResponse) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ExporterStatusResponse) SetState(v string) {
	o.State = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ExporterStatusResponse) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExporterStatusResponse) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ExporterStatusResponse) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *ExporterStatusResponse) SetOffset(v int32) {
	o.Offset = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *ExporterStatusResponse) GetTs() int32 {
	if o == nil || o.Ts == nil {
		var ret int32
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExporterStatusResponse) GetTsOk() (*int32, bool) {
	if o == nil || o.Ts == nil {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *ExporterStatusResponse) HasTs() bool {
	if o != nil && o.Ts != nil {
		return true
	}

	return false
}

// SetTs gets a reference to the given int32 and assigns it to the Ts field.
func (o *ExporterStatusResponse) SetTs(v int32) {
	o.Ts = &v
}

// GetTrace returns the Trace field value if set, zero value otherwise.
func (o *ExporterStatusResponse) GetTrace() string {
	if o == nil || o.Trace == nil {
		var ret string
		return ret
	}
	return *o.Trace
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExporterStatusResponse) GetTraceOk() (*string, bool) {
	if o == nil || o.Trace == nil {
		return nil, false
	}
	return o.Trace, true
}

// HasTrace returns a boolean if a field has been set.
func (o *ExporterStatusResponse) HasTrace() bool {
	if o != nil && o.Trace != nil {
		return true
	}

	return false
}

// SetTrace gets a reference to the given string and assigns it to the Trace field.
func (o *ExporterStatusResponse) SetTrace(v string) {
	o.Trace = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ExporterStatusResponse) Redact() {
	o.recurseRedact(o.Name)
	o.recurseRedact(o.State)
	o.recurseRedact(o.Offset)
	o.recurseRedact(o.Ts)
	o.recurseRedact(o.Trace)
}

func (o *ExporterStatusResponse) recurseRedact(v interface{}) {
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

func (o ExporterStatusResponse) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ExporterStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Ts != nil {
		toSerialize["ts"] = o.Ts
	}
	if o.Trace != nil {
		toSerialize["trace"] = o.Trace
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableExporterStatusResponse struct {
	value *ExporterStatusResponse
	isSet bool
}

func (v NullableExporterStatusResponse) Get() *ExporterStatusResponse {
	return v.value
}

func (v *NullableExporterStatusResponse) Set(val *ExporterStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExporterStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExporterStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExporterStatusResponse(val *ExporterStatusResponse) *NullableExporterStatusResponse {
	return &NullableExporterStatusResponse{value: val, isSet: true}
}

func (v NullableExporterStatusResponse) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableExporterStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
