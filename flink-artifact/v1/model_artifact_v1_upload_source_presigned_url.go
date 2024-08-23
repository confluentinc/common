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
Flink Artifact Management API

This is the Flink Artifact Management API.

API version: 0.0.1
Contact: flink-runtime@confluent.io
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

// ArtifactV1UploadSourcePresignedUrl Request a presigned upload URL for new Flink Artifact. Note that the URL policy expires in one hour. If the policy expires, you can request a new presigned upload URL.   ## The Presigned Urls Model <SchemaDefinition schemaRef=\"#/components/schemas/artifact.v1.PresignedUrl\" />
type ArtifactV1UploadSourcePresignedUrl struct {
	// APIVersion defines the schema version of this representation of a resource.
	ApiVersion *string `json:"api_version,omitempty"`
	// Kind defines the object this REST resource represents.
	Kind *string `json:"kind,omitempty"`
	// ID is the \"natural identifier\" for an object within its scope/namespace; it is normally unique across time but not space. That is, you can assume that the ID will not be reclaimed and reused after an object is deleted (\"time\"); however, it may collide with IDs for other object `kinds` or objects of the same `kind` within a different scope/namespace (\"space\").
	Id       *string     `json:"id,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	// Location of the Flink Artifact source.
	Location *string `json:"location,omitempty"`
	// Upload ID returned by the `/presigned-upload-url` API. This field returns an empty string in all responses.
	UploadId *string `json:"upload_id,omitempty"`
}

// NewArtifactV1UploadSourcePresignedUrl instantiates a new ArtifactV1UploadSourcePresignedUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactV1UploadSourcePresignedUrl() *ArtifactV1UploadSourcePresignedUrl {
	this := ArtifactV1UploadSourcePresignedUrl{}
	return &this
}

// NewArtifactV1UploadSourcePresignedUrlWithDefaults instantiates a new ArtifactV1UploadSourcePresignedUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactV1UploadSourcePresignedUrlWithDefaults() *ArtifactV1UploadSourcePresignedUrl {
	this := ArtifactV1UploadSourcePresignedUrl{}
	return &this
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ArtifactV1UploadSourcePresignedUrl) GetApiVersion() string {
	if o == nil || o.ApiVersion == nil {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1UploadSourcePresignedUrl) GetApiVersionOk() (*string, bool) {
	if o == nil || o.ApiVersion == nil {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ArtifactV1UploadSourcePresignedUrl) HasApiVersion() bool {
	if o != nil && o.ApiVersion != nil {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ArtifactV1UploadSourcePresignedUrl) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ArtifactV1UploadSourcePresignedUrl) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1UploadSourcePresignedUrl) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ArtifactV1UploadSourcePresignedUrl) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ArtifactV1UploadSourcePresignedUrl) SetKind(v string) {
	o.Kind = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ArtifactV1UploadSourcePresignedUrl) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1UploadSourcePresignedUrl) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ArtifactV1UploadSourcePresignedUrl) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ArtifactV1UploadSourcePresignedUrl) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ArtifactV1UploadSourcePresignedUrl) GetMetadata() ObjectMeta {
	if o == nil || o.Metadata == nil {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1UploadSourcePresignedUrl) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ArtifactV1UploadSourcePresignedUrl) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *ArtifactV1UploadSourcePresignedUrl) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ArtifactV1UploadSourcePresignedUrl) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1UploadSourcePresignedUrl) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ArtifactV1UploadSourcePresignedUrl) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *ArtifactV1UploadSourcePresignedUrl) SetLocation(v string) {
	o.Location = &v
}

// GetUploadId returns the UploadId field value if set, zero value otherwise.
func (o *ArtifactV1UploadSourcePresignedUrl) GetUploadId() string {
	if o == nil || o.UploadId == nil {
		var ret string
		return ret
	}
	return *o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactV1UploadSourcePresignedUrl) GetUploadIdOk() (*string, bool) {
	if o == nil || o.UploadId == nil {
		return nil, false
	}
	return o.UploadId, true
}

// HasUploadId returns a boolean if a field has been set.
func (o *ArtifactV1UploadSourcePresignedUrl) HasUploadId() bool {
	if o != nil && o.UploadId != nil {
		return true
	}

	return false
}

// SetUploadId gets a reference to the given string and assigns it to the UploadId field.
func (o *ArtifactV1UploadSourcePresignedUrl) SetUploadId(v string) {
	o.UploadId = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *ArtifactV1UploadSourcePresignedUrl) Redact() {
	o.recurseRedact(o.ApiVersion)
	o.recurseRedact(o.Kind)
	o.recurseRedact(o.Id)
	o.recurseRedact(o.Metadata)
	o.recurseRedact(o.Location)
	o.recurseRedact(o.UploadId)
}

func (o *ArtifactV1UploadSourcePresignedUrl) recurseRedact(v interface{}) {
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

func (o ArtifactV1UploadSourcePresignedUrl) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o ArtifactV1UploadSourcePresignedUrl) MarshalJSON() ([]byte, error) {
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
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.UploadId != nil {
		toSerialize["upload_id"] = o.UploadId
	}
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(toSerialize)
	return buffer.Bytes(), err
}

type NullableArtifactV1UploadSourcePresignedUrl struct {
	value *ArtifactV1UploadSourcePresignedUrl
	isSet bool
}

func (v NullableArtifactV1UploadSourcePresignedUrl) Get() *ArtifactV1UploadSourcePresignedUrl {
	return v.value
}

func (v *NullableArtifactV1UploadSourcePresignedUrl) Set(val *ArtifactV1UploadSourcePresignedUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactV1UploadSourcePresignedUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactV1UploadSourcePresignedUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactV1UploadSourcePresignedUrl(val *ArtifactV1UploadSourcePresignedUrl) *NullableArtifactV1UploadSourcePresignedUrl {
	return &NullableArtifactV1UploadSourcePresignedUrl{value: val, isSet: true}
}

func (v NullableArtifactV1UploadSourcePresignedUrl) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableArtifactV1UploadSourcePresignedUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}