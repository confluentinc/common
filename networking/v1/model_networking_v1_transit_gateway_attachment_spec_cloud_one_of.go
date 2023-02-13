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
Network API

Network API

API version: 0.0.1-alpha1
Contact: cire-traffic@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"encoding/json"
	"fmt"
)

// NetworkingV1TransitGatewayAttachmentSpecCloudOneOf - struct for NetworkingV1TransitGatewayAttachmentSpecCloudOneOf
type NetworkingV1TransitGatewayAttachmentSpecCloudOneOf struct {
	NetworkingV1AwsTransitGatewayAttachment *NetworkingV1AwsTransitGatewayAttachment
}

// NetworkingV1AwsTransitGatewayAttachmentAsNetworkingV1TransitGatewayAttachmentSpecCloudOneOf is a convenience function that returns NetworkingV1AwsTransitGatewayAttachment wrapped in NetworkingV1TransitGatewayAttachmentSpecCloudOneOf
func NetworkingV1AwsTransitGatewayAttachmentAsNetworkingV1TransitGatewayAttachmentSpecCloudOneOf(v *NetworkingV1AwsTransitGatewayAttachment) NetworkingV1TransitGatewayAttachmentSpecCloudOneOf {
	return NetworkingV1TransitGatewayAttachmentSpecCloudOneOf{ NetworkingV1AwsTransitGatewayAttachment: v}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *NetworkingV1TransitGatewayAttachmentSpecCloudOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'AwsTransitGatewayAttachment'
	if jsonDict["kind"] == "AwsTransitGatewayAttachment" {
		// try to unmarshal JSON data into NetworkingV1AwsTransitGatewayAttachment
		err = json.Unmarshal(data, &dst.NetworkingV1AwsTransitGatewayAttachment)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AwsTransitGatewayAttachment, return on the first match
		} else {
			dst.NetworkingV1AwsTransitGatewayAttachment = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1TransitGatewayAttachmentSpecCloudOneOf as NetworkingV1AwsTransitGatewayAttachment: %s", err.Error())
		}
	}

	// check if the discriminator value is 'networking.v1.AwsTransitGatewayAttachment'
	if jsonDict["kind"] == "networking.v1.AwsTransitGatewayAttachment" {
		// try to unmarshal JSON data into NetworkingV1AwsTransitGatewayAttachment
		err = json.Unmarshal(data, &dst.NetworkingV1AwsTransitGatewayAttachment)
		if err == nil {
			return nil // data stored in dst.NetworkingV1AwsTransitGatewayAttachment, return on the first match
		} else {
			dst.NetworkingV1AwsTransitGatewayAttachment = nil
			return fmt.Errorf("Failed to unmarshal NetworkingV1TransitGatewayAttachmentSpecCloudOneOf as NetworkingV1AwsTransitGatewayAttachment: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NetworkingV1TransitGatewayAttachmentSpecCloudOneOf) MarshalJSON() ([]byte, error) {
	if src.NetworkingV1AwsTransitGatewayAttachment != nil {
		return json.Marshal(&src.NetworkingV1AwsTransitGatewayAttachment)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NetworkingV1TransitGatewayAttachmentSpecCloudOneOf) GetActualInstance() (interface{}) {
	if obj.NetworkingV1AwsTransitGatewayAttachment != nil {
		return obj.NetworkingV1AwsTransitGatewayAttachment
	}

	// all schemas are nil
	return nil
}

type NullableNetworkingV1TransitGatewayAttachmentSpecCloudOneOf struct {
	value *NetworkingV1TransitGatewayAttachmentSpecCloudOneOf
	isSet bool
}

func (v NullableNetworkingV1TransitGatewayAttachmentSpecCloudOneOf) Get() *NetworkingV1TransitGatewayAttachmentSpecCloudOneOf {
	return v.value
}

func (v *NullableNetworkingV1TransitGatewayAttachmentSpecCloudOneOf) Set(val *NetworkingV1TransitGatewayAttachmentSpecCloudOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1TransitGatewayAttachmentSpecCloudOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1TransitGatewayAttachmentSpecCloudOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1TransitGatewayAttachmentSpecCloudOneOf(val *NetworkingV1TransitGatewayAttachmentSpecCloudOneOf) *NullableNetworkingV1TransitGatewayAttachmentSpecCloudOneOf {
	return &NullableNetworkingV1TransitGatewayAttachmentSpecCloudOneOf{value: val, isSet: true}
}

func (v NullableNetworkingV1TransitGatewayAttachmentSpecCloudOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1TransitGatewayAttachmentSpecCloudOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


