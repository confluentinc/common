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
	"fmt"
)

// ArtifactV1FlinkArtifactVersionUploadSourceOneOf - struct for ArtifactV1FlinkArtifactVersionUploadSourceOneOf
type ArtifactV1FlinkArtifactVersionUploadSourceOneOf struct {
	ArtifactV1UploadSourcePresignedUrl *ArtifactV1UploadSourcePresignedUrl
}

// ArtifactV1UploadSourcePresignedUrlAsArtifactV1FlinkArtifactVersionUploadSourceOneOf is a convenience function that returns ArtifactV1UploadSourcePresignedUrl wrapped in ArtifactV1FlinkArtifactVersionUploadSourceOneOf
func ArtifactV1UploadSourcePresignedUrlAsArtifactV1FlinkArtifactVersionUploadSourceOneOf(v *ArtifactV1UploadSourcePresignedUrl) ArtifactV1FlinkArtifactVersionUploadSourceOneOf {
	return ArtifactV1FlinkArtifactVersionUploadSourceOneOf{ArtifactV1UploadSourcePresignedUrl: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ArtifactV1FlinkArtifactVersionUploadSourceOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'PRESIGNED_URL_LOCATION'
	if jsonDict["location"] == "PRESIGNED_URL_LOCATION" {
		// try to unmarshal JSON data into ArtifactV1UploadSourcePresignedUrl
		err = json.Unmarshal(data, &dst.ArtifactV1UploadSourcePresignedUrl)
		if err == nil {
			return nil // data stored in dst.ArtifactV1UploadSourcePresignedUrl, return on the first match
		} else {
			dst.ArtifactV1UploadSourcePresignedUrl = nil
			return fmt.Errorf("Failed to unmarshal ArtifactV1FlinkArtifactVersionUploadSourceOneOf as ArtifactV1UploadSourcePresignedUrl: %s", err.Error())
		}
	}

	// check if the discriminator value is 'artifact.v1.UploadSource.PresignedUrl'
	if jsonDict["location"] == "artifact.v1.UploadSource.PresignedUrl" {
		// try to unmarshal JSON data into ArtifactV1UploadSourcePresignedUrl
		err = json.Unmarshal(data, &dst.ArtifactV1UploadSourcePresignedUrl)
		if err == nil {
			return nil // data stored in dst.ArtifactV1UploadSourcePresignedUrl, return on the first match
		} else {
			dst.ArtifactV1UploadSourcePresignedUrl = nil
			return fmt.Errorf("Failed to unmarshal ArtifactV1FlinkArtifactVersionUploadSourceOneOf as ArtifactV1UploadSourcePresignedUrl: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ArtifactV1FlinkArtifactVersionUploadSourceOneOf) MarshalJSON() ([]byte, error) {
	if src.ArtifactV1UploadSourcePresignedUrl != nil {
		buffer := &bytes.Buffer{}
		encoder := json.NewEncoder(buffer)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(&src.ArtifactV1UploadSourcePresignedUrl)
		return buffer.Bytes(), err
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ArtifactV1FlinkArtifactVersionUploadSourceOneOf) GetActualInstance() interface{} {
	if obj.ArtifactV1UploadSourcePresignedUrl != nil {
		return obj.ArtifactV1UploadSourcePresignedUrl
	}

	// all schemas are nil
	return nil
}

type NullableArtifactV1FlinkArtifactVersionUploadSourceOneOf struct {
	value *ArtifactV1FlinkArtifactVersionUploadSourceOneOf
	isSet bool
}

func (v NullableArtifactV1FlinkArtifactVersionUploadSourceOneOf) Get() *ArtifactV1FlinkArtifactVersionUploadSourceOneOf {
	return v.value
}

func (v *NullableArtifactV1FlinkArtifactVersionUploadSourceOneOf) Set(val *ArtifactV1FlinkArtifactVersionUploadSourceOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactV1FlinkArtifactVersionUploadSourceOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactV1FlinkArtifactVersionUploadSourceOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactV1FlinkArtifactVersionUploadSourceOneOf(val *ArtifactV1FlinkArtifactVersionUploadSourceOneOf) *NullableArtifactV1FlinkArtifactVersionUploadSourceOneOf {
	return &NullableArtifactV1FlinkArtifactVersionUploadSourceOneOf{value: val, isSet: true}
}

func (v NullableArtifactV1FlinkArtifactVersionUploadSourceOneOf) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableArtifactV1FlinkArtifactVersionUploadSourceOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}