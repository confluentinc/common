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
)

import (
	"reflect"
)

// NetworkingV1GcpPeering GCP VPC Peering.
type NetworkingV1GcpPeering struct {
	// Peering kind type.
	Kind string `json:"kind"`
	// The Google Cloud project ID associated with the VPC that you are peering with Confluent Cloud network.
	Project string `json:"project"`
	// The name of the VPC that you are peering with Confluent Cloud network.
	VpcNetwork string `json:"vpc_network"`
	// Enable customer route import. For more information, see [Importing custom routes](https://cloud.google.com/vpc/docs/vpc-peering#importing-exporting-routes).
	ImportCustomRoutes *bool `json:"import_custom_routes,omitempty"`
}

// NewNetworkingV1GcpPeering instantiates a new NetworkingV1GcpPeering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1GcpPeering(kind string, project string, vpcNetwork string) *NetworkingV1GcpPeering {
	this := NetworkingV1GcpPeering{}
	this.Kind = kind
	this.Project = project
	this.VpcNetwork = vpcNetwork
	var importCustomRoutes bool = false
	this.ImportCustomRoutes = &importCustomRoutes
	return &this
}

// NewNetworkingV1GcpPeeringWithDefaults instantiates a new NetworkingV1GcpPeering object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1GcpPeeringWithDefaults() *NetworkingV1GcpPeering {
	this := NetworkingV1GcpPeering{}
	var importCustomRoutes bool = false
	this.ImportCustomRoutes = &importCustomRoutes
	return &this
}

// GetKind returns the Kind field value
func (o *NetworkingV1GcpPeering) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1GcpPeering) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1GcpPeering) SetKind(v string) {
	o.Kind = v
}

// GetProject returns the Project field value
func (o *NetworkingV1GcpPeering) GetProject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1GcpPeering) GetProjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *NetworkingV1GcpPeering) SetProject(v string) {
	o.Project = v
}

// GetVpcNetwork returns the VpcNetwork field value
func (o *NetworkingV1GcpPeering) GetVpcNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VpcNetwork
}

// GetVpcNetworkOk returns a tuple with the VpcNetwork field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1GcpPeering) GetVpcNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpcNetwork, true
}

// SetVpcNetwork sets field value
func (o *NetworkingV1GcpPeering) SetVpcNetwork(v string) {
	o.VpcNetwork = v
}

// GetImportCustomRoutes returns the ImportCustomRoutes field value if set, zero value otherwise.
func (o *NetworkingV1GcpPeering) GetImportCustomRoutes() bool {
	if o == nil || o.ImportCustomRoutes == nil {
		var ret bool
		return ret
	}
	return *o.ImportCustomRoutes
}

// GetImportCustomRoutesOk returns a tuple with the ImportCustomRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkingV1GcpPeering) GetImportCustomRoutesOk() (*bool, bool) {
	if o == nil || o.ImportCustomRoutes == nil {
		return nil, false
	}
	return o.ImportCustomRoutes, true
}

// HasImportCustomRoutes returns a boolean if a field has been set.
func (o *NetworkingV1GcpPeering) HasImportCustomRoutes() bool {
	if o != nil && o.ImportCustomRoutes != nil {
		return true
	}

	return false
}

// SetImportCustomRoutes gets a reference to the given bool and assigns it to the ImportCustomRoutes field.
func (o *NetworkingV1GcpPeering) SetImportCustomRoutes(v bool) {
	o.ImportCustomRoutes = &v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1GcpPeering) Redact() {
	o.recurseRedact(&o.Kind)
	o.recurseRedact(&o.Project)
	o.recurseRedact(&o.VpcNetwork)
	o.recurseRedact(o.ImportCustomRoutes)
}

func (o *NetworkingV1GcpPeering) recurseRedact(v interface{}) {
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

func (o NetworkingV1GcpPeering) zeroField(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1GcpPeering) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["project"] = o.Project
	}
	if true {
		toSerialize["vpc_network"] = o.VpcNetwork
	}
	if o.ImportCustomRoutes != nil {
		toSerialize["import_custom_routes"] = o.ImportCustomRoutes
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1GcpPeering struct {
	value *NetworkingV1GcpPeering
	isSet bool
}

func (v NullableNetworkingV1GcpPeering) Get() *NetworkingV1GcpPeering {
	return v.value
}

func (v *NullableNetworkingV1GcpPeering) Set(val *NetworkingV1GcpPeering) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1GcpPeering) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1GcpPeering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1GcpPeering(val *NetworkingV1GcpPeering) *NullableNetworkingV1GcpPeering {
	return &NullableNetworkingV1GcpPeering{value: val, isSet: true}
}

func (v NullableNetworkingV1GcpPeering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1GcpPeering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
