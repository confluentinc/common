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
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type NetworksNetworkingV1Api interface {

	/*
			CreateNetworkingV1Network Create a Network

			[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to create a network.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @return ApiCreateNetworkingV1NetworkRequest
	*/
	CreateNetworkingV1Network(ctx _context.Context) ApiCreateNetworkingV1NetworkRequest

	// CreateNetworkingV1NetworkExecute executes the request
	//  @return NetworkingV1Network
	CreateNetworkingV1NetworkExecute(r ApiCreateNetworkingV1NetworkRequest) (NetworkingV1Network, *_nethttp.Response, error)

	/*
			DeleteNetworkingV1Network Delete a Network

			[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to delete a network.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param id The unique identifier for the network.
			 @return ApiDeleteNetworkingV1NetworkRequest
	*/
	DeleteNetworkingV1Network(ctx _context.Context, id string) ApiDeleteNetworkingV1NetworkRequest

	// DeleteNetworkingV1NetworkExecute executes the request
	DeleteNetworkingV1NetworkExecute(r ApiDeleteNetworkingV1NetworkRequest) (*_nethttp.Response, error)

	/*
			GetNetworkingV1Network Read a Network

			[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to read a network.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param id The unique identifier for the network.
			 @return ApiGetNetworkingV1NetworkRequest
	*/
	GetNetworkingV1Network(ctx _context.Context, id string) ApiGetNetworkingV1NetworkRequest

	// GetNetworkingV1NetworkExecute executes the request
	//  @return NetworkingV1Network
	GetNetworkingV1NetworkExecute(r ApiGetNetworkingV1NetworkRequest) (NetworkingV1Network, *_nethttp.Response, error)

	/*
			ListNetworkingV1Networks List of Networks

			[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Retrieve a sorted, filtered, paginated list of all networks.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @return ApiListNetworkingV1NetworksRequest
	*/
	ListNetworkingV1Networks(ctx _context.Context) ApiListNetworkingV1NetworksRequest

	// ListNetworkingV1NetworksExecute executes the request
	//  @return NetworkingV1NetworkList
	ListNetworkingV1NetworksExecute(r ApiListNetworkingV1NetworksRequest) (NetworkingV1NetworkList, *_nethttp.Response, error)

	/*
			UpdateNetworkingV1Network Update a Network

			[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to update a network.



			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param id The unique identifier for the network.
			 @return ApiUpdateNetworkingV1NetworkRequest
	*/
	UpdateNetworkingV1Network(ctx _context.Context, id string) ApiUpdateNetworkingV1NetworkRequest

	// UpdateNetworkingV1NetworkExecute executes the request
	//  @return NetworkingV1Network
	UpdateNetworkingV1NetworkExecute(r ApiUpdateNetworkingV1NetworkRequest) (NetworkingV1Network, *_nethttp.Response, error)
}

// NetworksNetworkingV1ApiService NetworksNetworkingV1Api service
type NetworksNetworkingV1ApiService service

type ApiCreateNetworkingV1NetworkRequest struct {
	ctx                 _context.Context
	ApiService          NetworksNetworkingV1Api
	networkingV1Network *NetworkingV1Network
}

func (r ApiCreateNetworkingV1NetworkRequest) NetworkingV1Network(networkingV1Network NetworkingV1Network) ApiCreateNetworkingV1NetworkRequest {
	r.networkingV1Network = &networkingV1Network
	return r
}

func (r ApiCreateNetworkingV1NetworkRequest) Execute() (NetworkingV1Network, *_nethttp.Response, error) {
	return r.ApiService.CreateNetworkingV1NetworkExecute(r)
}

/*
CreateNetworkingV1Network Create a Network

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to create a network.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateNetworkingV1NetworkRequest
*/
func (a *NetworksNetworkingV1ApiService) CreateNetworkingV1Network(ctx _context.Context) ApiCreateNetworkingV1NetworkRequest {
	return ApiCreateNetworkingV1NetworkRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return NetworkingV1Network
func (a *NetworksNetworkingV1ApiService) CreateNetworkingV1NetworkExecute(r ApiCreateNetworkingV1NetworkRequest) (NetworkingV1Network, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1Network
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworksNetworkingV1ApiService.CreateNetworkingV1Network")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/networks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.networkingV1Network
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
		if localVarHTTPResponse.StatusCode == 402 {
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
		if localVarHTTPResponse.StatusCode == 409 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
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

type ApiDeleteNetworkingV1NetworkRequest struct {
	ctx         _context.Context
	ApiService  NetworksNetworkingV1Api
	environment *string
	id          string
}

// Scope the operation to the given environment.
func (r ApiDeleteNetworkingV1NetworkRequest) Environment(environment string) ApiDeleteNetworkingV1NetworkRequest {
	r.environment = &environment
	return r
}

func (r ApiDeleteNetworkingV1NetworkRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteNetworkingV1NetworkExecute(r)
}

/*
DeleteNetworkingV1Network Delete a Network

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to delete a network.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the network.
 @return ApiDeleteNetworkingV1NetworkRequest
*/
func (a *NetworksNetworkingV1ApiService) DeleteNetworkingV1Network(ctx _context.Context, id string) ApiDeleteNetworkingV1NetworkRequest {
	return ApiDeleteNetworkingV1NetworkRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *NetworksNetworkingV1ApiService) DeleteNetworkingV1NetworkExecute(r ApiDeleteNetworkingV1NetworkRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworksNetworkingV1ApiService.DeleteNetworkingV1Network")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/networks/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.environment == nil {
		return nil, reportError("environment is required and must be specified")
	}

	localVarQueryParams.Add("environment", parameterToString(*r.environment, ""))
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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetNetworkingV1NetworkRequest struct {
	ctx         _context.Context
	ApiService  NetworksNetworkingV1Api
	environment *string
	id          string
}

// Scope the operation to the given environment.
func (r ApiGetNetworkingV1NetworkRequest) Environment(environment string) ApiGetNetworkingV1NetworkRequest {
	r.environment = &environment
	return r
}

func (r ApiGetNetworkingV1NetworkRequest) Execute() (NetworkingV1Network, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkingV1NetworkExecute(r)
}

/*
GetNetworkingV1Network Read a Network

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to read a network.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the network.
 @return ApiGetNetworkingV1NetworkRequest
*/
func (a *NetworksNetworkingV1ApiService) GetNetworkingV1Network(ctx _context.Context, id string) ApiGetNetworkingV1NetworkRequest {
	return ApiGetNetworkingV1NetworkRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return NetworkingV1Network
func (a *NetworksNetworkingV1ApiService) GetNetworkingV1NetworkExecute(r ApiGetNetworkingV1NetworkRequest) (NetworkingV1Network, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1Network
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworksNetworkingV1ApiService.GetNetworkingV1Network")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/networks/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.environment == nil {
		return localVarReturnValue, nil, reportError("environment is required and must be specified")
	}

	localVarQueryParams.Add("environment", parameterToString(*r.environment, ""))
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiListNetworkingV1NetworksRequest struct {
	ctx             _context.Context
	ApiService      NetworksNetworkingV1Api
	environment     *string
	specDisplayName *MultipleSearchFilter
	specCloud       *MultipleSearchFilter
	specRegion      *MultipleSearchFilter
	specCidr        *MultipleSearchFilter
	statusPhase     *MultipleSearchFilter
	connectionType  *MultipleSearchFilter
	pageSize        *int32
	pageToken       *string
}

// Filter the results by exact match for environment.
func (r ApiListNetworkingV1NetworksRequest) Environment(environment string) ApiListNetworkingV1NetworksRequest {
	r.environment = &environment
	return r
}

// Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1NetworksRequest) SpecDisplayName(specDisplayName MultipleSearchFilter) ApiListNetworkingV1NetworksRequest {
	r.specDisplayName = &specDisplayName
	return r
}

// Filter the results by exact match for spec.cloud. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1NetworksRequest) SpecCloud(specCloud MultipleSearchFilter) ApiListNetworkingV1NetworksRequest {
	r.specCloud = &specCloud
	return r
}

// Filter the results by exact match for spec.region. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1NetworksRequest) SpecRegion(specRegion MultipleSearchFilter) ApiListNetworkingV1NetworksRequest {
	r.specRegion = &specRegion
	return r
}

// Filter the results by exact match for spec.cidr. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1NetworksRequest) SpecCidr(specCidr MultipleSearchFilter) ApiListNetworkingV1NetworksRequest {
	r.specCidr = &specCidr
	return r
}

// Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1NetworksRequest) StatusPhase(statusPhase MultipleSearchFilter) ApiListNetworkingV1NetworksRequest {
	r.statusPhase = &statusPhase
	return r
}

// Filter the results by exact match for connection_type. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1NetworksRequest) ConnectionType(connectionType MultipleSearchFilter) ApiListNetworkingV1NetworksRequest {
	r.connectionType = &connectionType
	return r
}

// A pagination size for collection requests.
func (r ApiListNetworkingV1NetworksRequest) PageSize(pageSize int32) ApiListNetworkingV1NetworksRequest {
	r.pageSize = &pageSize
	return r
}

// An opaque pagination token for collection requests.
func (r ApiListNetworkingV1NetworksRequest) PageToken(pageToken string) ApiListNetworkingV1NetworksRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListNetworkingV1NetworksRequest) Execute() (NetworkingV1NetworkList, *_nethttp.Response, error) {
	return r.ApiService.ListNetworkingV1NetworksExecute(r)
}

/*
ListNetworkingV1Networks List of Networks

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Retrieve a sorted, filtered, paginated list of all networks.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListNetworkingV1NetworksRequest
*/
func (a *NetworksNetworkingV1ApiService) ListNetworkingV1Networks(ctx _context.Context) ApiListNetworkingV1NetworksRequest {
	return ApiListNetworkingV1NetworksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return NetworkingV1NetworkList
func (a *NetworksNetworkingV1ApiService) ListNetworkingV1NetworksExecute(r ApiListNetworkingV1NetworksRequest) (NetworkingV1NetworkList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1NetworkList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworksNetworkingV1ApiService.ListNetworkingV1Networks")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/networks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.environment == nil {
		return localVarReturnValue, nil, reportError("environment is required and must be specified")
	}

	if r.specDisplayName != nil {
		localVarQueryParams.Add("spec.display_name", parameterToString(*r.specDisplayName, ""))
	}
	if r.specCloud != nil {
		localVarQueryParams.Add("spec.cloud", parameterToString(*r.specCloud, ""))
	}
	if r.specRegion != nil {
		localVarQueryParams.Add("spec.region", parameterToString(*r.specRegion, ""))
	}
	if r.specCidr != nil {
		localVarQueryParams.Add("spec.cidr", parameterToString(*r.specCidr, ""))
	}
	if r.statusPhase != nil {
		localVarQueryParams.Add("status.phase", parameterToString(*r.statusPhase, ""))
	}
	localVarQueryParams.Add("environment", parameterToString(*r.environment, ""))
	if r.connectionType != nil {
		localVarQueryParams.Add("connection_type", parameterToString(*r.connectionType, ""))
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

type ApiUpdateNetworkingV1NetworkRequest struct {
	ctx                       _context.Context
	ApiService                NetworksNetworkingV1Api
	id                        string
	networkingV1NetworkUpdate *NetworkingV1NetworkUpdate
}

func (r ApiUpdateNetworkingV1NetworkRequest) NetworkingV1NetworkUpdate(networkingV1NetworkUpdate NetworkingV1NetworkUpdate) ApiUpdateNetworkingV1NetworkRequest {
	r.networkingV1NetworkUpdate = &networkingV1NetworkUpdate
	return r
}

func (r ApiUpdateNetworkingV1NetworkRequest) Execute() (NetworkingV1Network, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkingV1NetworkExecute(r)
}

/*
UpdateNetworkingV1Network Update a Network

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to update a network.



 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the network.
 @return ApiUpdateNetworkingV1NetworkRequest
*/
func (a *NetworksNetworkingV1ApiService) UpdateNetworkingV1Network(ctx _context.Context, id string) ApiUpdateNetworkingV1NetworkRequest {
	return ApiUpdateNetworkingV1NetworkRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return NetworkingV1Network
func (a *NetworksNetworkingV1ApiService) UpdateNetworkingV1NetworkExecute(r ApiUpdateNetworkingV1NetworkRequest) (NetworkingV1Network, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1Network
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworksNetworkingV1ApiService.UpdateNetworkingV1Network")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/networks/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.networkingV1NetworkUpdate
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
		if localVarHTTPResponse.StatusCode == 402 {
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
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
