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
Cluster Management for Apache Kafka API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha1
Contact: orchestrator-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

import (
	"encoding/json"
)

import (
	"reflect"
)

// CmkV2ClusterSpec The desired state of the Cluster
type CmkV2ClusterSpec struct {
	// The name of the cluster.
	DisplayName *string `json:"display_name,omitempty"`
	// The availability zone configuration of the cluster.
	Availability *string `json:"availability,omitempty"`
	// The cloud service provider in which the cluster is running.
	Cloud *string `json:"cloud,omitempty"`
	// The cloud service provider region where the cluster is running.
	Region *string `json:"region,omitempty"`
	// The configuration of the Kafka cluster.  Note: Clusters can be upgraded from Basic to Standard, but cannot be downgraded from Standard to Basic. 
	Config *CmkV2ClusterSpecConfigOneOf `json:"config,omitempty"`
	// The bootstrap endpoint used by Kafka clients to connect to the cluster.
	KafkaBootstrapEndpoint *string `json:"kafka_bootstrap_endpoint,omitempty"`
	// The cluster HTTP request URL.
	HttpEndpoint *string `json:"http_endpoint,omitempty"`
	// The environment to which this belongs.
	Environment *ObjectReference `json:"environment,omitempty"`
	// The network associated with this object.
	Network *ObjectReference `json:"network,omitempty"`
}

// NewCmkV2ClusterSpec instantiates a new CmkV2ClusterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCmkV2ClusterSpec() *CmkV2ClusterSpec {
	this := CmkV2ClusterSpec{}
	var availability string = "SINGLE_ZONE"
	this.Availability = &availability
	return &this
}

// NewCmkV2ClusterSpecWithDefaults instantiates a new CmkV2ClusterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCmkV2ClusterSpecWithDefaults() *CmkV2ClusterSpec {
	this := CmkV2ClusterSpec{}
	var availability string = "SINGLE_ZONE"
	this.Availability = &availability
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CmkV2ClusterSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CmkV2ClusterSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CmkV2ClusterSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *CmkV2ClusterSpec) GetAvailability() string {
	if o == nil || o.Availability == nil {
		var ret string
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpec) GetAvailabilityOk() (*string, bool) {
	if o == nil || o.Availability == nil {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *CmkV2ClusterSpec) HasAvailability() bool {
	if o != nil && o.Availability != nil {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given string and assigns it to the Availability field.
func (o *CmkV2ClusterSpec) SetAvailability(v string) {
	o.Availability = &v
}

// GetCloud returns the Cloud field value if set, zero value otherwise.
func (o *CmkV2ClusterSpec) GetCloud() string {
	if o == nil || o.Cloud == nil {
		var ret string
		return ret
	}
	return *o.Cloud
}

// GetCloudOk returns a tuple with the Cloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpec) GetCloudOk() (*string, bool) {
	if o == nil || o.Cloud == nil {
		return nil, false
	}
	return o.Cloud, true
}

// HasCloud returns a boolean if a field has been set.
func (o *CmkV2ClusterSpec) HasCloud() bool {
	if o != nil && o.Cloud != nil {
		return true
	}

	return false
}

// SetCloud gets a reference to the given string and assigns it to the Cloud field.
func (o *CmkV2ClusterSpec) SetCloud(v string) {
	o.Cloud = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CmkV2ClusterSpec) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpec) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CmkV2ClusterSpec) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CmkV2ClusterSpec) SetRegion(v string) {
	o.Region = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *CmkV2ClusterSpec) GetConfig() CmkV2ClusterSpecConfigOneOf {
	if o == nil || o.Config == nil {
		var ret CmkV2ClusterSpecConfigOneOf
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpec) GetConfigOk() (*CmkV2ClusterSpecConfigOneOf, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *CmkV2ClusterSpec) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given CmkV2ClusterSpecConfigOneOf and assigns it to the Config field.
func (o *CmkV2ClusterSpec) SetConfig(v CmkV2ClusterSpecConfigOneOf) {
	o.Config = &v
}

// GetKafkaBootstrapEndpoint returns the KafkaBootstrapEndpoint field value if set, zero value otherwise.
func (o *CmkV2ClusterSpec) GetKafkaBootstrapEndpoint() string {
	if o == nil || o.KafkaBootstrapEndpoint == nil {
		var ret string
		return ret
	}
	return *o.KafkaBootstrapEndpoint
}

// GetKafkaBootstrapEndpointOk returns a tuple with the KafkaBootstrapEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpec) GetKafkaBootstrapEndpointOk() (*string, bool) {
	if o == nil || o.KafkaBootstrapEndpoint == nil {
		return nil, false
	}
	return o.KafkaBootstrapEndpoint, true
}

// HasKafkaBootstrapEndpoint returns a boolean if a field has been set.
func (o *CmkV2ClusterSpec) HasKafkaBootstrapEndpoint() bool {
	if o != nil && o.KafkaBootstrapEndpoint != nil {
		return true
	}

	return false
}

// SetKafkaBootstrapEndpoint gets a reference to the given string and assigns it to the KafkaBootstrapEndpoint field.
func (o *CmkV2ClusterSpec) SetKafkaBootstrapEndpoint(v string) {
	o.KafkaBootstrapEndpoint = &v
}

// GetHttpEndpoint returns the HttpEndpoint field value if set, zero value otherwise.
func (o *CmkV2ClusterSpec) GetHttpEndpoint() string {
	if o == nil || o.HttpEndpoint == nil {
		var ret string
		return ret
	}
	return *o.HttpEndpoint
}

// GetHttpEndpointOk returns a tuple with the HttpEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpec) GetHttpEndpointOk() (*string, bool) {
	if o == nil || o.HttpEndpoint == nil {
		return nil, false
	}
	return o.HttpEndpoint, true
}

// HasHttpEndpoint returns a boolean if a field has been set.
func (o *CmkV2ClusterSpec) HasHttpEndpoint() bool {
	if o != nil && o.HttpEndpoint != nil {
		return true
	}

	return false
}

// SetHttpEndpoint gets a reference to the given string and assigns it to the HttpEndpoint field.
func (o *CmkV2ClusterSpec) SetHttpEndpoint(v string) {
	o.HttpEndpoint = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CmkV2ClusterSpec) GetEnvironment() ObjectReference {
	if o == nil || o.Environment == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpec) GetEnvironmentOk() (*ObjectReference, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CmkV2ClusterSpec) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectReference and assigns it to the Environment field.
func (o *CmkV2ClusterSpec) SetEnvironment(v ObjectReference) {
	o.Environment = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *CmkV2ClusterSpec) GetNetwork() ObjectReference {
	if o == nil || o.Network == nil {
		var ret ObjectReference
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CmkV2ClusterSpec) GetNetworkOk() (*ObjectReference, bool) {
	if o == nil || o.Network == nil {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *CmkV2ClusterSpec) HasNetwork() bool {
	if o != nil && o.Network != nil {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ObjectReference and assigns it to the Network field.
func (o *CmkV2ClusterSpec) SetNetwork(v ObjectReference) {
	o.Network = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *CmkV2ClusterSpec) Redact() {
    o.recurseRedact(o.DisplayName)
    o.recurseRedact(o.Availability)
    o.recurseRedact(o.Cloud)
    o.recurseRedact(o.Region)
    o.recurseRedact(o.Config)
    o.recurseRedact(o.KafkaBootstrapEndpoint)
    o.recurseRedact(o.HttpEndpoint)
    o.recurseRedact(o.Environment)
    o.recurseRedact(o.Network)
}

func (o *CmkV2ClusterSpec) recurseRedact(v interface{}) {
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

func (o CmkV2ClusterSpec) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o CmkV2ClusterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Availability != nil {
		toSerialize["availability"] = o.Availability
	}
	if o.Cloud != nil {
		toSerialize["cloud"] = o.Cloud
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.KafkaBootstrapEndpoint != nil {
		toSerialize["kafka_bootstrap_endpoint"] = o.KafkaBootstrapEndpoint
	}
	if o.HttpEndpoint != nil {
		toSerialize["http_endpoint"] = o.HttpEndpoint
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Network != nil {
		toSerialize["network"] = o.Network
	}
	return json.Marshal(toSerialize)
}

type NullableCmkV2ClusterSpec struct {
	value *CmkV2ClusterSpec
	isSet bool
}

func (v NullableCmkV2ClusterSpec) Get() *CmkV2ClusterSpec {
	return v.value
}

func (v *NullableCmkV2ClusterSpec) Set(val *CmkV2ClusterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCmkV2ClusterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCmkV2ClusterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCmkV2ClusterSpec(val *CmkV2ClusterSpec) *NullableCmkV2ClusterSpec {
	return &NullableCmkV2ClusterSpec{value: val, isSet: true}
}

func (v NullableCmkV2ClusterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCmkV2ClusterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


