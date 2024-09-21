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
Certificate Authority Management API

mTLS Public API spec

API version: 0.0.1-alpha1
Contact: cloud-authn-platform-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"bytes"
	"encoding/json"
)

import (
	"reflect"
)

// IamV2CreateCertRequest This contains the json schema used to create a Certificate Authority
type IamV2CreateCertRequest struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string     `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// The human-readable name of the certificate authority.
	DisplayName *string `json:"display_name,omitempty"`
	// A description of the certificate authority.
	Description *string `json:"description,omitempty"`
	// The Base64 encoded string containing the signing certificate chain used to validate client certs.
	CertificateChain *string `json:"certificate_chain,omitempty"`
	// The name of the certificate file.
	CertificateChainFilename *string `json:"certificate_chain_filename,omitempty"`
	// The url from which to fetch the CRL for the certificate authority if crl_source is URL.
	CrlUrl *string `json:"crl_url,omitempty"`
	// The Base64 encoded string containing the CRL for this certificate authority. Defaults to this over `crl_url` if available.
	CrlChain *string `json:"crl_chain,omitempty"`
}

// NewIamV2CreateCertRequest instantiates a new IamV2CreateCertRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamV2CreateCertRequest() *IamV2CreateCertRequest {
	this := IamV2CreateCertRequest{}
	return &this
}

// NewIamV2CreateCertRequestWithDefaults instantiates a new IamV2CreateCertRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamV2CreateCertRequestWithDefaults() *IamV2CreateCertRequest {
	this := IamV2CreateCertRequest{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *IamV2CreateCertRequest) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CreateCertRequest) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *IamV2CreateCertRequest) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *IamV2CreateCertRequest) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *IamV2CreateCertRequest) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CreateCertRequest) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *IamV2CreateCertRequest) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *IamV2CreateCertRequest) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IamV2CreateCertRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CreateCertRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IamV2CreateCertRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IamV2CreateCertRequest) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IamV2CreateCertRequest) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CreateCertRequest) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IamV2CreateCertRequest) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *IamV2CreateCertRequest) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *IamV2CreateCertRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CreateCertRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *IamV2CreateCertRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *IamV2CreateCertRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamV2CreateCertRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CreateCertRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamV2CreateCertRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamV2CreateCertRequest) SetDescription(v string) {
	o.Description = &v
}

// GetCertificateChain returns the CertificateChain field value if set, zero value otherwise.
func (o *IamV2CreateCertRequest) GetCertificateChain() string {
	if o == nil || o.CertificateChain == nil {
		var ret string
		return ret
	}
	return *o.CertificateChain
}

// GetCertificateChainOk returns a tuple with the CertificateChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CreateCertRequest) GetCertificateChainOk() (*string, bool) {
	if o == nil || o.CertificateChain == nil {
		return nil, false
	}
	return o.CertificateChain, true
}

// HasCertificateChain returns a boolean if a field has been set.
func (o *IamV2CreateCertRequest) HasCertificateChain() bool {
	if o != nil && o.CertificateChain != nil {
		return true
	}

	return false
}

// SetCertificateChain gets a reference to the given string and assigns it to the CertificateChain field.
func (o *IamV2CreateCertRequest) SetCertificateChain(v string) {
	o.CertificateChain = &v
}

// GetCertificateChainFilename returns the CertificateChainFilename field value if set, zero value otherwise.
func (o *IamV2CreateCertRequest) GetCertificateChainFilename() string {
	if o == nil || o.CertificateChainFilename == nil {
		var ret string
		return ret
	}
	return *o.CertificateChainFilename
}

// GetCertificateChainFilenameOk returns a tuple with the CertificateChainFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CreateCertRequest) GetCertificateChainFilenameOk() (*string, bool) {
	if o == nil || o.CertificateChainFilename == nil {
		return nil, false
	}
	return o.CertificateChainFilename, true
}

// HasCertificateChainFilename returns a boolean if a field has been set.
func (o *IamV2CreateCertRequest) HasCertificateChainFilename() bool {
	if o != nil && o.CertificateChainFilename != nil {
		return true
	}

	return false
}

// SetCertificateChainFilename gets a reference to the given string and assigns it to the CertificateChainFilename field.
func (o *IamV2CreateCertRequest) SetCertificateChainFilename(v string) {
	o.CertificateChainFilename = &v
}

// GetCrlUrl returns the CrlUrl field value if set, zero value otherwise.
func (o *IamV2CreateCertRequest) GetCrlUrl() string {
	if o == nil || o.CrlUrl == nil {
		var ret string
		return ret
	}
	return *o.CrlUrl
}

// GetCrlUrlOk returns a tuple with the CrlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CreateCertRequest) GetCrlUrlOk() (*string, bool) {
	if o == nil || o.CrlUrl == nil {
		return nil, false
	}
	return o.CrlUrl, true
}

// HasCrlUrl returns a boolean if a field has been set.
func (o *IamV2CreateCertRequest) HasCrlUrl() bool {
	if o != nil && o.CrlUrl != nil {
		return true
	}

	return false
}

// SetCrlUrl gets a reference to the given string and assigns it to the CrlUrl field.
func (o *IamV2CreateCertRequest) SetCrlUrl(v string) {
	o.CrlUrl = &v
}

// GetCrlChain returns the CrlChain field value if set, zero value otherwise.
func (o *IamV2CreateCertRequest) GetCrlChain() string {
	if o == nil || o.CrlChain == nil {
		var ret string
		return ret
	}
	return *o.CrlChain
}

// GetCrlChainOk returns a tuple with the CrlChain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamV2CreateCertRequest) GetCrlChainOk() (*string, bool) {
	if o == nil || o.CrlChain == nil {
		return nil, false
	}
	return o.CrlChain, true
}

// HasCrlChain returns a boolean if a field has been set.
func (o *IamV2CreateCertRequest) HasCrlChain() bool {
	if o != nil && o.CrlChain != nil {
		return true
	}

	return false
}

// SetCrlChain gets a reference to the given string and assigns it to the CrlChain field.
func (o *IamV2CreateCertRequest) SetCrlChain(v string) {
	o.CrlChain = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *IamV2CreateCertRequest) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.DisplayName)
	o.recurseRedact(o.Description)
	o.recurseRedact(o.CertificateChain)
	o.recurseRedact(o.CertificateChainFilename)
	o.recurseRedact(o.CrlUrl)
	o.recurseRedact(o.CrlChain)
}

func (o *IamV2CreateCertRequest) recurseRedact(v interface{}) {
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

func (o IamV2CreateCertRequest) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o IamV2CreateCertRequest) MarshalJSON() ([]byte, error) {
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
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.CertificateChain != nil {
		toSerialize["certificate_chain"] = o.CertificateChain
	}
	if o.CertificateChainFilename != nil {
		toSerialize["certificate_chain_filename"] = o.CertificateChainFilename
	}
	if o.CrlUrl != nil {
		toSerialize["crl_url"] = o.CrlUrl
	}
	if o.CrlChain != nil {
		toSerialize["crl_chain"] = o.CrlChain
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableIamV2CreateCertRequest struct {
	value *IamV2CreateCertRequest
	isSet bool
}

func (v NullableIamV2CreateCertRequest) Get() *IamV2CreateCertRequest {
	return v.value
}

func (v *NullableIamV2CreateCertRequest) Set(val *IamV2CreateCertRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIamV2CreateCertRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIamV2CreateCertRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamV2CreateCertRequest(val *IamV2CreateCertRequest) *NullableIamV2CreateCertRequest {
	return &NullableIamV2CreateCertRequest{value: val, isSet: true}
}

func (v NullableIamV2CreateCertRequest) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableIamV2CreateCertRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
