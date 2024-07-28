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
Stream Share APIs

# Introduction

API version: 0.1.0-alpha0
Contact: cdx@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
)

import (
	"reflect"
)

// CdxV1RedeemTokenResponse Share details for the consumer org or user
type CdxV1RedeemTokenResponse struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string     `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// The api key
	ApiKey *string `json:"api_key,omitempty"`
	// The api key secret
	Secret *string `json:"secret,omitempty"`
	// The kafka cluster bootstrap url
	KafkaBootstrapUrl *string `json:"kafka_bootstrap_url,omitempty"`
	// The api key for schema registry
	SchemaRegistryApiKey *string `json:"schema_registry_api_key,omitempty"`
	// The api key secret for schema registry
	SchemaRegistrySecret *string `json:"schema_registry_secret,omitempty"`
	// The schema registry endpoint url
	SchemaRegistryUrl *string `json:"schema_registry_url,omitempty"`
	// List of shared resources
	Resources *[]CdxV1RedeemTokenResponseResourcesOneOf `json:"resources,omitempty"`
}

// NewCdxV1RedeemTokenResponse instantiates a new CdxV1RedeemTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdxV1RedeemTokenResponse() *CdxV1RedeemTokenResponse {
	this := CdxV1RedeemTokenResponse{}
	return &this
}

// NewCdxV1RedeemTokenResponseWithDefaults instantiates a new CdxV1RedeemTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdxV1RedeemTokenResponseWithDefaults() *CdxV1RedeemTokenResponse {
	this := CdxV1RedeemTokenResponse{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *CdxV1RedeemTokenResponse) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1RedeemTokenResponse) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *CdxV1RedeemTokenResponse) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *CdxV1RedeemTokenResponse) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *CdxV1RedeemTokenResponse) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1RedeemTokenResponse) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *CdxV1RedeemTokenResponse) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *CdxV1RedeemTokenResponse) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CdxV1RedeemTokenResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1RedeemTokenResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CdxV1RedeemTokenResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CdxV1RedeemTokenResponse) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CdxV1RedeemTokenResponse) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1RedeemTokenResponse) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CdxV1RedeemTokenResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *CdxV1RedeemTokenResponse) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *CdxV1RedeemTokenResponse) GetApiKey() string {
	if o == nil || o.ApiKey == nil {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1RedeemTokenResponse) GetApiKeyOk() (*string, bool) {
	if o == nil || o.ApiKey == nil {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *CdxV1RedeemTokenResponse) HasApiKey() bool {
	if o != nil && o.ApiKey != nil {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *CdxV1RedeemTokenResponse) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *CdxV1RedeemTokenResponse) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1RedeemTokenResponse) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *CdxV1RedeemTokenResponse) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *CdxV1RedeemTokenResponse) SetSecret(v string) {
	o.Secret = &v
}

// GetKafkaBootstrapUrl returns the KafkaBootstrapUrl field value if set, zero value otherwise.
func (o *CdxV1RedeemTokenResponse) GetKafkaBootstrapUrl() string {
	if o == nil || o.KafkaBootstrapUrl == nil {
		var ret string
		return ret
	}
	return *o.KafkaBootstrapUrl
}

// GetKafkaBootstrapUrlOk returns a tuple with the KafkaBootstrapUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1RedeemTokenResponse) GetKafkaBootstrapUrlOk() (*string, bool) {
	if o == nil || o.KafkaBootstrapUrl == nil {
		return nil, false
	}
	return o.KafkaBootstrapUrl, true
}

// HasKafkaBootstrapUrl returns a boolean if a field has been set.
func (o *CdxV1RedeemTokenResponse) HasKafkaBootstrapUrl() bool {
	if o != nil && o.KafkaBootstrapUrl != nil {
		return true
	}

	return false
}

// SetKafkaBootstrapUrl gets a reference to the given string and assigns it to the KafkaBootstrapUrl field.
func (o *CdxV1RedeemTokenResponse) SetKafkaBootstrapUrl(v string) {
	o.KafkaBootstrapUrl = &v
}

// GetSchemaRegistryApiKey returns the SchemaRegistryApiKey field value if set, zero value otherwise.
func (o *CdxV1RedeemTokenResponse) GetSchemaRegistryApiKey() string {
	if o == nil || o.SchemaRegistryApiKey == nil {
		var ret string
		return ret
	}
	return *o.SchemaRegistryApiKey
}

// GetSchemaRegistryApiKeyOk returns a tuple with the SchemaRegistryApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1RedeemTokenResponse) GetSchemaRegistryApiKeyOk() (*string, bool) {
	if o == nil || o.SchemaRegistryApiKey == nil {
		return nil, false
	}
	return o.SchemaRegistryApiKey, true
}

// HasSchemaRegistryApiKey returns a boolean if a field has been set.
func (o *CdxV1RedeemTokenResponse) HasSchemaRegistryApiKey() bool {
	if o != nil && o.SchemaRegistryApiKey != nil {
		return true
	}

	return false
}

// SetSchemaRegistryApiKey gets a reference to the given string and assigns it to the SchemaRegistryApiKey field.
func (o *CdxV1RedeemTokenResponse) SetSchemaRegistryApiKey(v string) {
	o.SchemaRegistryApiKey = &v
}

// GetSchemaRegistrySecret returns the SchemaRegistrySecret field value if set, zero value otherwise.
func (o *CdxV1RedeemTokenResponse) GetSchemaRegistrySecret() string {
	if o == nil || o.SchemaRegistrySecret == nil {
		var ret string
		return ret
	}
	return *o.SchemaRegistrySecret
}

// GetSchemaRegistrySecretOk returns a tuple with the SchemaRegistrySecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1RedeemTokenResponse) GetSchemaRegistrySecretOk() (*string, bool) {
	if o == nil || o.SchemaRegistrySecret == nil {
		return nil, false
	}
	return o.SchemaRegistrySecret, true
}

// HasSchemaRegistrySecret returns a boolean if a field has been set.
func (o *CdxV1RedeemTokenResponse) HasSchemaRegistrySecret() bool {
	if o != nil && o.SchemaRegistrySecret != nil {
		return true
	}

	return false
}

// SetSchemaRegistrySecret gets a reference to the given string and assigns it to the SchemaRegistrySecret field.
func (o *CdxV1RedeemTokenResponse) SetSchemaRegistrySecret(v string) {
	o.SchemaRegistrySecret = &v
}

// GetSchemaRegistryUrl returns the SchemaRegistryUrl field value if set, zero value otherwise.
func (o *CdxV1RedeemTokenResponse) GetSchemaRegistryUrl() string {
	if o == nil || o.SchemaRegistryUrl == nil {
		var ret string
		return ret
	}
	return *o.SchemaRegistryUrl
}

// GetSchemaRegistryUrlOk returns a tuple with the SchemaRegistryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1RedeemTokenResponse) GetSchemaRegistryUrlOk() (*string, bool) {
	if o == nil || o.SchemaRegistryUrl == nil {
		return nil, false
	}
	return o.SchemaRegistryUrl, true
}

// HasSchemaRegistryUrl returns a boolean if a field has been set.
func (o *CdxV1RedeemTokenResponse) HasSchemaRegistryUrl() bool {
	if o != nil && o.SchemaRegistryUrl != nil {
		return true
	}

	return false
}

// SetSchemaRegistryUrl gets a reference to the given string and assigns it to the SchemaRegistryUrl field.
func (o *CdxV1RedeemTokenResponse) SetSchemaRegistryUrl(v string) {
	o.SchemaRegistryUrl = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *CdxV1RedeemTokenResponse) GetResources() []CdxV1RedeemTokenResponseResourcesOneOf {
	if o == nil || o.Resources == nil {
		var ret []CdxV1RedeemTokenResponseResourcesOneOf
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdxV1RedeemTokenResponse) GetResourcesOk() (*[]CdxV1RedeemTokenResponseResourcesOneOf, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *CdxV1RedeemTokenResponse) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []CdxV1RedeemTokenResponseResourcesOneOf and assigns it to the Resources field.
func (o *CdxV1RedeemTokenResponse) SetResources(v []CdxV1RedeemTokenResponseResourcesOneOf) {
	o.Resources = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CdxV1RedeemTokenResponse) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.ApiKey)
	o.Secret = nil
	o.recurseRedact(o.KafkaBootstrapUrl)
	o.recurseRedact(o.SchemaRegistryApiKey)
	o.SchemaRegistrySecret = nil
	o.recurseRedact(o.SchemaRegistryUrl)
	o.recurseRedact(o.Resources)
}

func (o *CdxV1RedeemTokenResponse) recurseRedact(v interface{}) {
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

func (o CdxV1RedeemTokenResponse) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o CdxV1RedeemTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVersion != nil {
		toSerialize["api_version"] = o.ApiVersion
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.ApiKey != nil {
		toSerialize["api_key"] = o.ApiKey
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.KafkaBootstrapUrl != nil {
		toSerialize["kafka_bootstrap_url"] = o.KafkaBootstrapUrl
	}
	if o.SchemaRegistryApiKey != nil {
		toSerialize["schema_registry_api_key"] = o.SchemaRegistryApiKey
	}
	if o.SchemaRegistrySecret != nil {
		toSerialize["schema_registry_secret"] = o.SchemaRegistrySecret
	}
	if o.SchemaRegistryUrl != nil {
		toSerialize["schema_registry_url"] = o.SchemaRegistryUrl
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableCdxV1RedeemTokenResponse struct {
	value *CdxV1RedeemTokenResponse
	isSet bool
}

func (v NullableCdxV1RedeemTokenResponse) Get() *CdxV1RedeemTokenResponse {
	return v.value
}

func (v *NullableCdxV1RedeemTokenResponse) Set(val *CdxV1RedeemTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCdxV1RedeemTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCdxV1RedeemTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdxV1RedeemTokenResponse(val *CdxV1RedeemTokenResponse) *NullableCdxV1RedeemTokenResponse {
	return &NullableCdxV1RedeemTokenResponse{value: val, isSet: true}
}

func (v NullableCdxV1RedeemTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdxV1RedeemTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}