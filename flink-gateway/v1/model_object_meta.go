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
SQL API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	"encoding/json"
	"time"
)

import (
	"reflect"
)

// ObjectMeta The metadata of the statement.
type ObjectMeta struct {
	// Self is a Uniform Resource Locator (URL) at which an object can be addressed. This URL encodes the service location, API version, and other particulars necessary to locate the resource at a point in time
	Self string `json:"self,omitempty"`
	// The date and time at which this object was created. It is represented in RFC3339 format and is in UTC.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The date and time at which this object was last updated. It is represented in RFC3339 format and is in UTC.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// A system generated globally unique identifier for this resource.
	Uid *string `json:"uid,omitempty"`
	// A system generated string that uniquely identifies the version of this resource.
	ResourceVersion *string `json:"resource_version,omitempty"`
	// A map of key-value pairs that describe the resource.
	Labels *map[string]string `json:"labels,omitempty"`
}

// NewObjectMeta instantiates a new ObjectMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectMeta(self string) *ObjectMeta {
	this := ObjectMeta{}
	this.Self = self
	return &this
}

// NewObjectMetaWithDefaults instantiates a new ObjectMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectMetaWithDefaults() *ObjectMeta {
	this := ObjectMeta{}
	return &this
}

// GetSelf returns the Self field value
func (o *ObjectMeta) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *ObjectMeta) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *ObjectMeta) SetSelf(v string) {
	o.Self = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ObjectMeta) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMeta) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ObjectMeta) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ObjectMeta) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ObjectMeta) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMeta) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ObjectMeta) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ObjectMeta) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *ObjectMeta) GetUid() string {
	if o == nil || o.Uid == nil {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMeta) GetUidOk() (*string, bool) {
	if o == nil || o.Uid == nil {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *ObjectMeta) HasUid() bool {
	if o != nil && o.Uid != nil {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *ObjectMeta) SetUid(v string) {
	o.Uid = &v
}

// GetResourceVersion returns the ResourceVersion field value if set, zero value otherwise.
func (o *ObjectMeta) GetResourceVersion() string {
	if o == nil || o.ResourceVersion == nil {
		var ret string
		return ret
	}
	return *o.ResourceVersion
}

// GetResourceVersionOk returns a tuple with the ResourceVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMeta) GetResourceVersionOk() (*string, bool) {
	if o == nil || o.ResourceVersion == nil {
		return nil, false
	}
	return o.ResourceVersion, true
}

// HasResourceVersion returns a boolean if a field has been set.
func (o *ObjectMeta) HasResourceVersion() bool {
	if o != nil && o.ResourceVersion != nil {
		return true
	}

	return false
}

// SetResourceVersion gets a reference to the given string and assigns it to the ResourceVersion field.
func (o *ObjectMeta) SetResourceVersion(v string) {
	o.ResourceVersion = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ObjectMeta) GetLabels() map[string]string {
	if o == nil || o.Labels == nil {
		var ret map[string]string
		return ret
	}
	return *o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectMeta) GetLabelsOk() (*map[string]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ObjectMeta) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the Labels field.
func (o *ObjectMeta) SetLabels(v map[string]string) {
	o.Labels = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ObjectMeta) Redact() {
	o.recurseRedact(&o.Self)
	o.recurseRedact(o.CreatedAt)
	o.recurseRedact(o.UpdatedAt)
	o.recurseRedact(o.Uid)
	o.recurseRedact(o.ResourceVersion)
	o.recurseRedact(o.Labels)
}

func (o *ObjectMeta) recurseRedact(v interface{}) {
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

func (o ObjectMeta) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ObjectMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Uid != nil {
		toSerialize["uid"] = o.Uid
	}
	if o.ResourceVersion != nil {
		toSerialize["resource_version"] = o.ResourceVersion
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableObjectMeta struct {
	value *ObjectMeta
	isSet bool
}

func (v NullableObjectMeta) Get() *ObjectMeta {
	return v.value
}

func (v *NullableObjectMeta) Set(val *ObjectMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectMeta(val *ObjectMeta) *NullableObjectMeta {
	return &NullableObjectMeta{value: val, isSet: true}
}

func (v NullableObjectMeta) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableObjectMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
