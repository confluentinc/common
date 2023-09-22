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
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

type IPAddressesNetworkingV1Api interface {

	/*
		ListNetworkingV1IpAddresses List of IP Addresses

		[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

	Related guide: [Use Static IP addresses on Confluent Cloud](https://docs.confluent.io/cloud/current/networking/static-egress-ip-addresses.html)

	Retrieve a sorted, filtered, paginated list of all IP Addresses.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return ApiListNetworkingV1IpAddressesRequest
	*/
	ListNetworkingV1IpAddresses(ctx _context.Context) ApiListNetworkingV1IpAddressesRequest

	// ListNetworkingV1IpAddressesExecute executes the request
	//  @return NetworkingV1IpAddressList
	ListNetworkingV1IpAddressesExecute(r ApiListNetworkingV1IpAddressesRequest) (NetworkingV1IpAddressList, *_nethttp.Response, error)
}

// IPAddressesNetworkingV1ApiService IPAddressesNetworkingV1Api service
type IPAddressesNetworkingV1ApiService service

type ApiListNetworkingV1IpAddressesRequest struct {
	ctx         _context.Context
	ApiService  IPAddressesNetworkingV1Api
	cloud       *MultipleSearchFilter
	region      *MultipleSearchFilter
	services    *MultipleSearchFilter
	addressType *MultipleSearchFilter
	pageSize    *int32
	pageToken   *string
}

// Filter the results by exact match for cloud. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1IpAddressesRequest) Cloud(cloud MultipleSearchFilter) ApiListNetworkingV1IpAddressesRequest {
	r.cloud = &cloud
	return r
}

// Filter the results by exact match for region. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1IpAddressesRequest) Region(region MultipleSearchFilter) ApiListNetworkingV1IpAddressesRequest {
	r.region = &region
	return r
}

// Filter the results by exact match for services. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1IpAddressesRequest) Services(services MultipleSearchFilter) ApiListNetworkingV1IpAddressesRequest {
	r.services = &services
	return r
}

// Filter the results by exact match for address_type. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1IpAddressesRequest) AddressType(addressType MultipleSearchFilter) ApiListNetworkingV1IpAddressesRequest {
	r.addressType = &addressType
	return r
}

// A pagination size for collection requests.
func (r ApiListNetworkingV1IpAddressesRequest) PageSize(pageSize int32) ApiListNetworkingV1IpAddressesRequest {
	r.pageSize = &pageSize
	return r
}

// An opaque pagination token for collection requests.
func (r ApiListNetworkingV1IpAddressesRequest) PageToken(pageToken string) ApiListNetworkingV1IpAddressesRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListNetworkingV1IpAddressesRequest) Execute() (NetworkingV1IpAddressList, *_nethttp.Response, error) {
	return r.ApiService.ListNetworkingV1IpAddressesExecute(r)
}

/*
ListNetworkingV1IpAddresses List of IP Addresses

[![General Availability](https://img.shields.io/badge/Lifecycle%20Stage-General%20Availability-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Related guide: [Use Static IP addresses on Confluent Cloud](https://docs.confluent.io/cloud/current/networking/static-egress-ip-addresses.html)

Retrieve a sorted, filtered, paginated list of all IP Addresses.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListNetworkingV1IpAddressesRequest
*/
func (a *IPAddressesNetworkingV1ApiService) ListNetworkingV1IpAddresses(ctx _context.Context) ApiListNetworkingV1IpAddressesRequest {
	return ApiListNetworkingV1IpAddressesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return NetworkingV1IpAddressList
func (a *IPAddressesNetworkingV1ApiService) ListNetworkingV1IpAddressesExecute(r ApiListNetworkingV1IpAddressesRequest) (NetworkingV1IpAddressList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1IpAddressList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IPAddressesNetworkingV1ApiService.ListNetworkingV1IpAddresses")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/ip-addresses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.cloud != nil {
		localVarQueryParams.Add("cloud", parameterToString(*r.cloud, ""))
	}
	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.services != nil {
		localVarQueryParams.Add("services", parameterToString(*r.services, ""))
	}
	if r.addressType != nil {
		localVarQueryParams.Add("address_type", parameterToString(*r.addressType, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.pageToken != nil {
		localVarQueryParams.Add("page_token", parameterToString(*r.pageToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
