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
Kafka Connect APIs

REST API for managing connectors

API version: 1.0
Contact: connect@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// InlineResponse2003Definition The definition for a config in the connector plugin, which includes the name, type, importance, etc.
type InlineResponse2003Definition struct {
	// The name of the configuration
	Name *string `json:"name,omitempty"`
	// The config types
	Type *string `json:"type,omitempty"`
	// Whether this configuration is required
	Required *bool `json:"required,omitempty"`
	// Default value for this configuration
	DefaultValue *string `json:"default_value,omitempty"`
	// The importance level for a configuration
	Importance *string `json:"importance,omitempty"`
	// The documentation for the configuration
	Documentation *string `json:"documentation,omitempty"`
	// The UI group to which the configuration belongs to
	Group *string `json:"group,omitempty"`
	// The width of a configuration value
	Width *string `json:"width,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	// Other configurations on which this configuration is dependent
	Dependents *[]string `json:"dependents,omitempty"`
	// The order of configuration in specified group
	Order *int32 `json:"order,omitempty"`
	Alias *string `json:"alias,omitempty"`
}

// NewInlineResponse2003Definition instantiates a new InlineResponse2003Definition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003Definition() *InlineResponse2003Definition {
	this := InlineResponse2003Definition{}
	return &this
}

// NewInlineResponse2003DefinitionWithDefaults instantiates a new InlineResponse2003Definition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003DefinitionWithDefaults() *InlineResponse2003Definition {
	this := InlineResponse2003Definition{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InlineResponse2003Definition) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Definition) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InlineResponse2003Definition) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InlineResponse2003Definition) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InlineResponse2003Definition) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Definition) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InlineResponse2003Definition) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *InlineResponse2003Definition) SetType(v string) {
	o.Type = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *InlineResponse2003Definition) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Definition) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *InlineResponse2003Definition) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *InlineResponse2003Definition) SetRequired(v bool) {
	o.Required = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *InlineResponse2003Definition) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Definition) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *InlineResponse2003Definition) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *InlineResponse2003Definition) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetImportance returns the Importance field value if set, zero value otherwise.
func (o *InlineResponse2003Definition) GetImportance() string {
	if o == nil || o.Importance == nil {
		var ret string
		return ret
	}
	return *o.Importance
}

// GetImportanceOk returns a tuple with the Importance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Definition) GetImportanceOk() (*string, bool) {
	if o == nil || o.Importance == nil {
		return nil, false
	}
	return o.Importance, true
}

// HasImportance returns a boolean if a field has been set.
func (o *InlineResponse2003Definition) HasImportance() bool {
	if o != nil && o.Importance != nil {
		return true
	}

	return false
}

// SetImportance gets a reference to the given string and assigns it to the Importance field.
func (o *InlineResponse2003Definition) SetImportance(v string) {
	o.Importance = &v
}

// GetDocumentation returns the Documentation field value if set, zero value otherwise.
func (o *InlineResponse2003Definition) GetDocumentation() string {
	if o == nil || o.Documentation == nil {
		var ret string
		return ret
	}
	return *o.Documentation
}

// GetDocumentationOk returns a tuple with the Documentation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Definition) GetDocumentationOk() (*string, bool) {
	if o == nil || o.Documentation == nil {
		return nil, false
	}
	return o.Documentation, true
}

// HasDocumentation returns a boolean if a field has been set.
func (o *InlineResponse2003Definition) HasDocumentation() bool {
	if o != nil && o.Documentation != nil {
		return true
	}

	return false
}

// SetDocumentation gets a reference to the given string and assigns it to the Documentation field.
func (o *InlineResponse2003Definition) SetDocumentation(v string) {
	o.Documentation = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *InlineResponse2003Definition) GetGroup() string {
	if o == nil || o.Group == nil {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Definition) GetGroupOk() (*string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *InlineResponse2003Definition) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *InlineResponse2003Definition) SetGroup(v string) {
	o.Group = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *InlineResponse2003Definition) GetWidth() string {
	if o == nil || o.Width == nil {
		var ret string
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Definition) GetWidthOk() (*string, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *InlineResponse2003Definition) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given string and assigns it to the Width field.
func (o *InlineResponse2003Definition) SetWidth(v string) {
	o.Width = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *InlineResponse2003Definition) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Definition) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *InlineResponse2003Definition) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *InlineResponse2003Definition) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDependents returns the Dependents field value if set, zero value otherwise.
func (o *InlineResponse2003Definition) GetDependents() []string {
	if o == nil || o.Dependents == nil {
		var ret []string
		return ret
	}
	return *o.Dependents
}

// GetDependentsOk returns a tuple with the Dependents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Definition) GetDependentsOk() (*[]string, bool) {
	if o == nil || o.Dependents == nil {
		return nil, false
	}
	return o.Dependents, true
}

// HasDependents returns a boolean if a field has been set.
func (o *InlineResponse2003Definition) HasDependents() bool {
	if o != nil && o.Dependents != nil {
		return true
	}

	return false
}

// SetDependents gets a reference to the given []string and assigns it to the Dependents field.
func (o *InlineResponse2003Definition) SetDependents(v []string) {
	o.Dependents = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *InlineResponse2003Definition) GetOrder() int32 {
	if o == nil || o.Order == nil {
		var ret int32
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Definition) GetOrderOk() (*int32, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *InlineResponse2003Definition) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given int32 and assigns it to the Order field.
func (o *InlineResponse2003Definition) SetOrder(v int32) {
	o.Order = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *InlineResponse2003Definition) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Definition) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *InlineResponse2003Definition) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *InlineResponse2003Definition) SetAlias(v string) {
	o.Alias = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *InlineResponse2003Definition) Redact() {
    o.recurseRedact(o.Name)
    o.recurseRedact(o.Type)
    o.recurseRedact(o.Required)
    o.recurseRedact(o.DefaultValue)
    o.recurseRedact(o.Importance)
    o.recurseRedact(o.Documentation)
    o.recurseRedact(o.Group)
    o.recurseRedact(o.Width)
    o.recurseRedact(o.DisplayName)
    o.recurseRedact(o.Dependents)
    o.recurseRedact(o.Order)
    o.recurseRedact(o.Alias)
}

func (o *InlineResponse2003Definition) recurseRedact(v interface{}) {
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

func (o InlineResponse2003Definition) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o InlineResponse2003Definition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.DefaultValue != nil {
		toSerialize["default_value"] = o.DefaultValue
	}
	if o.Importance != nil {
		toSerialize["importance"] = o.Importance
	}
	if o.Documentation != nil {
		toSerialize["documentation"] = o.Documentation
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Dependents != nil {
		toSerialize["dependents"] = o.Dependents
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2003Definition struct {
	value *InlineResponse2003Definition
	isSet bool
}

func (v NullableInlineResponse2003Definition) Get() *InlineResponse2003Definition {
	return v.value
}

func (v *NullableInlineResponse2003Definition) Set(val *InlineResponse2003Definition) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2003Definition) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2003Definition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2003Definition(val *InlineResponse2003Definition) *NullableInlineResponse2003Definition {
	return &NullableInlineResponse2003Definition{value: val, isSet: true}
}

func (v NullableInlineResponse2003Definition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2003Definition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


