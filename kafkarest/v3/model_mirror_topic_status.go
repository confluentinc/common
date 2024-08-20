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
REST Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0
Contact: kafka-clients-proxy-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// MirrorTopicStatus the model 'MirrorTopicStatus'
type MirrorTopicStatus string

// List of MirrorTopicStatus
const (
	ACTIVE                    MirrorTopicStatus = "ACTIVE"
	FAILED                    MirrorTopicStatus = "FAILED"
	LINK_FAILED               MirrorTopicStatus = "LINK_FAILED"
	LINK_PAUSED               MirrorTopicStatus = "LINK_PAUSED"
	PAUSED                    MirrorTopicStatus = "PAUSED"
	PENDING_STOPPED           MirrorTopicStatus = "PENDING_STOPPED"
	SOURCE_UNAVAILABLE        MirrorTopicStatus = "SOURCE_UNAVAILABLE"
	STOPPED                   MirrorTopicStatus = "STOPPED"
	PENDING_MIRROR            MirrorTopicStatus = "PENDING_MIRROR"
	PENDING_SYNCHRONIZE       MirrorTopicStatus = "PENDING_SYNCHRONIZE"
	PENDING_SETUP_FOR_RESTORE MirrorTopicStatus = "PENDING_SETUP_FOR_RESTORE"
	PENDING_RESTORE           MirrorTopicStatus = "PENDING_RESTORE"
)

var allowedMirrorTopicStatusEnumValues = []MirrorTopicStatus{
	"ACTIVE",
	"FAILED",
	"LINK_FAILED",
	"LINK_PAUSED",
	"PAUSED",
	"PENDING_STOPPED",
	"SOURCE_UNAVAILABLE",
	"STOPPED",
	"PENDING_MIRROR",
	"PENDING_SYNCHRONIZE",
	"PENDING_SETUP_FOR_RESTORE",
	"PENDING_RESTORE",
}

func (v *MirrorTopicStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MirrorTopicStatus(value)
	for _, existing := range allowedMirrorTopicStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MirrorTopicStatus", value)
}

// NewMirrorTopicStatusFromValue returns a pointer to a valid MirrorTopicStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMirrorTopicStatusFromValue(v string) (*MirrorTopicStatus, error) {
	ev := MirrorTopicStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MirrorTopicStatus: valid values are %v", v, allowedMirrorTopicStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MirrorTopicStatus) IsValid() bool {
	for _, existing := range allowedMirrorTopicStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MirrorTopicStatus value
func (v MirrorTopicStatus) Ptr() *MirrorTopicStatus {
	return &v
}

type NullableMirrorTopicStatus struct {
	value *MirrorTopicStatus
	isSet bool
}

func (v NullableMirrorTopicStatus) Get() *MirrorTopicStatus {
	return v.value
}

func (v *NullableMirrorTopicStatus) Set(val *MirrorTopicStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMirrorTopicStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMirrorTopicStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMirrorTopicStatus(val *MirrorTopicStatus) *NullableMirrorTopicStatus {
	return &NullableMirrorTopicStatus{value: val, isSet: true}
}

func (v NullableMirrorTopicStatus) MarshalJSON() ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v.value)
	return buffer.Bytes(), err
}

func (v *NullableMirrorTopicStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
