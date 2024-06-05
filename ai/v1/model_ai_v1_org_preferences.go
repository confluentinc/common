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
AI API

API for interacting with AI models from within Confluent Cloud.

API version: 0.0.1
Contact: api-team@confluent.io
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

// AiV1OrgPreferences Enable the Confluent AI Assistant for your organization. This operation is only available to users with the `OrganizationAdmin` role. By default, this setting is set to `True`.   ## The Org Preferences Model <SchemaDefinition schemaRef=\"#/components/schemas/ai.v1.OrgPreferences\" />
type AiV1OrgPreferences struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// Enable ai-assist for the organization
	AiAssistantEnabled *bool `json:"ai_assistant_enabled,omitempty"`
}

// NewAiV1OrgPreferences instantiates a new AiV1OrgPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAiV1OrgPreferences() *AiV1OrgPreferences {
	this := AiV1OrgPreferences{}
	return &this
}

// NewAiV1OrgPreferencesWithDefaults instantiates a new AiV1OrgPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAiV1OrgPreferencesWithDefaults() *AiV1OrgPreferences {
	this := AiV1OrgPreferences{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *AiV1OrgPreferences) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiV1OrgPreferences) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *AiV1OrgPreferences) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *AiV1OrgPreferences) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *AiV1OrgPreferences) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiV1OrgPreferences) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *AiV1OrgPreferences) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *AiV1OrgPreferences) SetKind(v string) {
	o.Kind = &v
}

// GetAiAssistantEnabled returns the AiAssistantEnabled field value if set, zero value otherwise.
func (o *AiV1OrgPreferences) GetAiAssistantEnabled() bool {
	if o == nil || o.AiAssistantEnabled == nil {
		var ret bool
		return ret
	}
	return *o.AiAssistantEnabled
}

// GetAiAssistantEnabledOk returns a tuple with the AiAssistantEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AiV1OrgPreferences) GetAiAssistantEnabledOk() (*bool, bool) {
	if o == nil || o.AiAssistantEnabled == nil {
		return nil, false
	}
	return o.AiAssistantEnabled, true
}

// HasAiAssistantEnabled returns a boolean if a field has been set.
func (o *AiV1OrgPreferences) HasAiAssistantEnabled() bool {
	if o != nil && o.AiAssistantEnabled != nil {
		return true
	}

	return false
}

// SetAiAssistantEnabled gets a reference to the given bool and assigns it to the AiAssistantEnabled field.
func (o *AiV1OrgPreferences) SetAiAssistantEnabled(v bool) {
	o.AiAssistantEnabled = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *AiV1OrgPreferences) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.AiAssistantEnabled)
}

func (o *AiV1OrgPreferences) recurseRedact(v interface{}) {
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

func (o AiV1OrgPreferences) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o AiV1OrgPreferences) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.AiAssistantEnabled != nil {
		toSerialize["ai_assistant_enabled"] = o.AiAssistantEnabled
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableAiV1OrgPreferences struct {
	value *AiV1OrgPreferences
	isSet bool
}

func (v NullableAiV1OrgPreferences) Get() *AiV1OrgPreferences {
	return v.value
}

func (v *NullableAiV1OrgPreferences) Set(val *AiV1OrgPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableAiV1OrgPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableAiV1OrgPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAiV1OrgPreferences(val *AiV1OrgPreferences) *NullableAiV1OrgPreferences {
	return &NullableAiV1OrgPreferences{value: val, isSet: true}
}

func (v NullableAiV1OrgPreferences) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableAiV1OrgPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
