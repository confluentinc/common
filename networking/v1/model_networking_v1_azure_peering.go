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

// NetworkingV1AzurePeering Azure VNet Peering.
type NetworkingV1AzurePeering struct {
	// Peering kind type.
	Kind string `json:"kind"`
	// The Azure Tenant ID in which your Azure Subscription exists. Represents an organization in Azure Active Directory. You can find your Azure Tenant ID in the Azure Portal under [Azure Active Directory](https://portal.azure.com/#blade/Microsoft_AAD_IAM/ActiveDirectoryMenuBlade/Overview). Must be a valid **32 character UUID string**. 
	Tenant string `json:"tenant"`
	// The resource ID of the VNet that you are peering with Confluent Cloud. You can find the name of your Azure VNet in the [Azure Portal on the Overview tab of your Azure Virtual Network](https://portal.azure.com/#blade/HubsExtension/BrowseResource/resourceType/Microsoft.Network%2FvirtualNetworks).
	Vnet string `json:"vnet"`
	// The region of the VNet you are peering with Confluent Cloud network.
	CustomerRegion string `json:"customer_region"`
}

// NewNetworkingV1AzurePeering instantiates a new NetworkingV1AzurePeering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkingV1AzurePeering(kind string, tenant string, vnet string, customerRegion string) *NetworkingV1AzurePeering {
	this := NetworkingV1AzurePeering{}
	this.Kind = kind
	this.Tenant = tenant
	this.Vnet = vnet
	this.CustomerRegion = customerRegion
	return &this
}

// NewNetworkingV1AzurePeeringWithDefaults instantiates a new NetworkingV1AzurePeering object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkingV1AzurePeeringWithDefaults() *NetworkingV1AzurePeering {
	this := NetworkingV1AzurePeering{}
	return &this
}

// GetKind returns the Kind field value
func (o *NetworkingV1AzurePeering) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzurePeering) GetKindOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *NetworkingV1AzurePeering) SetKind(v string) {
	o.Kind = v
}

// GetTenant returns the Tenant field value
func (o *NetworkingV1AzurePeering) GetTenant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzurePeering) GetTenantOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tenant, true
}

// SetTenant sets field value
func (o *NetworkingV1AzurePeering) SetTenant(v string) {
	o.Tenant = v
}

// GetVnet returns the Vnet field value
func (o *NetworkingV1AzurePeering) GetVnet() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vnet
}

// GetVnetOk returns a tuple with the Vnet field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzurePeering) GetVnetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vnet, true
}

// SetVnet sets field value
func (o *NetworkingV1AzurePeering) SetVnet(v string) {
	o.Vnet = v
}

// GetCustomerRegion returns the CustomerRegion field value
func (o *NetworkingV1AzurePeering) GetCustomerRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerRegion
}

// GetCustomerRegionOk returns a tuple with the CustomerRegion field value
// and a boolean to check if the value has been set.
func (o *NetworkingV1AzurePeering) GetCustomerRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CustomerRegion, true
}

// SetCustomerRegion sets field value
func (o *NetworkingV1AzurePeering) SetCustomerRegion(v string) {
	o.CustomerRegion = v
}

// Redact resets all sensitive fields to their zero value.
func (o *NetworkingV1AzurePeering) Redact() {
    o.recurseRedact(&o.Kind)
    o.recurseRedact(&o.Tenant)
    o.recurseRedact(&o.Vnet)
    o.recurseRedact(&o.CustomerRegion)
}

func (o *NetworkingV1AzurePeering) recurseRedact(v interface{}) {
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

func (o NetworkingV1AzurePeering) zeroField(v interface{}) {
    p := reflect.ValueOf(v).Elem()
    p.Set(reflect.Zero(p.Type()))
}

func (o NetworkingV1AzurePeering) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["tenant"] = o.Tenant
	}
	if true {
		toSerialize["vnet"] = o.Vnet
	}
	if true {
		toSerialize["customer_region"] = o.CustomerRegion
	}
	return json.Marshal(toSerialize)
}

type NullableNetworkingV1AzurePeering struct {
	value *NetworkingV1AzurePeering
	isSet bool
}

func (v NullableNetworkingV1AzurePeering) Get() *NetworkingV1AzurePeering {
	return v.value
}

func (v *NullableNetworkingV1AzurePeering) Set(val *NetworkingV1AzurePeering) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkingV1AzurePeering) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkingV1AzurePeering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkingV1AzurePeering(val *NetworkingV1AzurePeering) *NullableNetworkingV1AzurePeering {
	return &NullableNetworkingV1AzurePeering{value: val, isSet: true}
}

func (v NullableNetworkingV1AzurePeering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkingV1AzurePeering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


