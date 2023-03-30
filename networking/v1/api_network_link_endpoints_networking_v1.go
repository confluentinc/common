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

type NetworkLinkEndpointsNetworkingV1Api interface {

	/*
			CreateNetworkingV1NetworkLinkEndpoint Create a Network Link Endpoint

			[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to create a network link endpoint.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @return ApiCreateNetworkingV1NetworkLinkEndpointRequest
	*/
	CreateNetworkingV1NetworkLinkEndpoint(ctx _context.Context) ApiCreateNetworkingV1NetworkLinkEndpointRequest

	// CreateNetworkingV1NetworkLinkEndpointExecute executes the request
	//  @return NetworkingV1NetworkLinkEndpoint
	CreateNetworkingV1NetworkLinkEndpointExecute(r ApiCreateNetworkingV1NetworkLinkEndpointRequest) (NetworkingV1NetworkLinkEndpoint, *_nethttp.Response, error)

	/*
			DeleteNetworkingV1NetworkLinkEndpoint Delete a Network Link Endpoint

			[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to delete a network link endpoint.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param id The unique identifier for the network link endpoint.
			 @return ApiDeleteNetworkingV1NetworkLinkEndpointRequest
	*/
	DeleteNetworkingV1NetworkLinkEndpoint(ctx _context.Context, id string) ApiDeleteNetworkingV1NetworkLinkEndpointRequest

	// DeleteNetworkingV1NetworkLinkEndpointExecute executes the request
	DeleteNetworkingV1NetworkLinkEndpointExecute(r ApiDeleteNetworkingV1NetworkLinkEndpointRequest) (*_nethttp.Response, error)

	/*
			GetNetworkingV1NetworkLinkEndpoint Read a Network Link Endpoint

			[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to read a network link endpoint.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param id The unique identifier for the network link endpoint.
			 @return ApiGetNetworkingV1NetworkLinkEndpointRequest
	*/
	GetNetworkingV1NetworkLinkEndpoint(ctx _context.Context, id string) ApiGetNetworkingV1NetworkLinkEndpointRequest

	// GetNetworkingV1NetworkLinkEndpointExecute executes the request
	//  @return NetworkingV1NetworkLinkEndpoint
	GetNetworkingV1NetworkLinkEndpointExecute(r ApiGetNetworkingV1NetworkLinkEndpointRequest) (NetworkingV1NetworkLinkEndpoint, *_nethttp.Response, error)

	/*
			ListNetworkingV1NetworkLinkEndpoints List of Network Link Endpoints

			[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Retrieve a sorted, filtered, paginated list of all network link endpoints.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @return ApiListNetworkingV1NetworkLinkEndpointsRequest
	*/
	ListNetworkingV1NetworkLinkEndpoints(ctx _context.Context) ApiListNetworkingV1NetworkLinkEndpointsRequest

	// ListNetworkingV1NetworkLinkEndpointsExecute executes the request
	//  @return NetworkingV1NetworkLinkEndpointList
	ListNetworkingV1NetworkLinkEndpointsExecute(r ApiListNetworkingV1NetworkLinkEndpointsRequest) (NetworkingV1NetworkLinkEndpointList, *_nethttp.Response, error)

	/*
			UpdateNetworkingV1NetworkLinkEndpoint Update a Network Link Endpoint

			[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to update a network link endpoint.



			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param id The unique identifier for the network link endpoint.
			 @return ApiUpdateNetworkingV1NetworkLinkEndpointRequest
	*/
	UpdateNetworkingV1NetworkLinkEndpoint(ctx _context.Context, id string) ApiUpdateNetworkingV1NetworkLinkEndpointRequest

	// UpdateNetworkingV1NetworkLinkEndpointExecute executes the request
	//  @return NetworkingV1NetworkLinkEndpoint
	UpdateNetworkingV1NetworkLinkEndpointExecute(r ApiUpdateNetworkingV1NetworkLinkEndpointRequest) (NetworkingV1NetworkLinkEndpoint, *_nethttp.Response, error)
}

// NetworkLinkEndpointsNetworkingV1ApiService NetworkLinkEndpointsNetworkingV1Api service
type NetworkLinkEndpointsNetworkingV1ApiService service

type ApiCreateNetworkingV1NetworkLinkEndpointRequest struct {
	ctx                             _context.Context
	ApiService                      NetworkLinkEndpointsNetworkingV1Api
	networkingV1NetworkLinkEndpoint *NetworkingV1NetworkLinkEndpoint
}

func (r ApiCreateNetworkingV1NetworkLinkEndpointRequest) NetworkingV1NetworkLinkEndpoint(networkingV1NetworkLinkEndpoint NetworkingV1NetworkLinkEndpoint) ApiCreateNetworkingV1NetworkLinkEndpointRequest {
	r.networkingV1NetworkLinkEndpoint = &networkingV1NetworkLinkEndpoint
	return r
}

func (r ApiCreateNetworkingV1NetworkLinkEndpointRequest) Execute() (NetworkingV1NetworkLinkEndpoint, *_nethttp.Response, error) {
	return r.ApiService.CreateNetworkingV1NetworkLinkEndpointExecute(r)
}

/*
CreateNetworkingV1NetworkLinkEndpoint Create a Network Link Endpoint

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to create a network link endpoint.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateNetworkingV1NetworkLinkEndpointRequest
*/
func (a *NetworkLinkEndpointsNetworkingV1ApiService) CreateNetworkingV1NetworkLinkEndpoint(ctx _context.Context) ApiCreateNetworkingV1NetworkLinkEndpointRequest {
	return ApiCreateNetworkingV1NetworkLinkEndpointRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return NetworkingV1NetworkLinkEndpoint
func (a *NetworkLinkEndpointsNetworkingV1ApiService) CreateNetworkingV1NetworkLinkEndpointExecute(r ApiCreateNetworkingV1NetworkLinkEndpointRequest) (NetworkingV1NetworkLinkEndpoint, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1NetworkLinkEndpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkLinkEndpointsNetworkingV1ApiService.CreateNetworkingV1NetworkLinkEndpoint")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/network-link-endpoints"

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
	localVarPostBody = r.networkingV1NetworkLinkEndpoint
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

type ApiDeleteNetworkingV1NetworkLinkEndpointRequest struct {
	ctx         _context.Context
	ApiService  NetworkLinkEndpointsNetworkingV1Api
	environment *string
	id          string
}

// Scope the operation to the given environment.
func (r ApiDeleteNetworkingV1NetworkLinkEndpointRequest) Environment(environment string) ApiDeleteNetworkingV1NetworkLinkEndpointRequest {
	r.environment = &environment
	return r
}

func (r ApiDeleteNetworkingV1NetworkLinkEndpointRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteNetworkingV1NetworkLinkEndpointExecute(r)
}

/*
DeleteNetworkingV1NetworkLinkEndpoint Delete a Network Link Endpoint

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to delete a network link endpoint.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the network link endpoint.
 @return ApiDeleteNetworkingV1NetworkLinkEndpointRequest
*/
func (a *NetworkLinkEndpointsNetworkingV1ApiService) DeleteNetworkingV1NetworkLinkEndpoint(ctx _context.Context, id string) ApiDeleteNetworkingV1NetworkLinkEndpointRequest {
	return ApiDeleteNetworkingV1NetworkLinkEndpointRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *NetworkLinkEndpointsNetworkingV1ApiService) DeleteNetworkingV1NetworkLinkEndpointExecute(r ApiDeleteNetworkingV1NetworkLinkEndpointRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkLinkEndpointsNetworkingV1ApiService.DeleteNetworkingV1NetworkLinkEndpoint")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/network-link-endpoints/{id}"
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

type ApiGetNetworkingV1NetworkLinkEndpointRequest struct {
	ctx         _context.Context
	ApiService  NetworkLinkEndpointsNetworkingV1Api
	environment *string
	id          string
}

// Scope the operation to the given environment.
func (r ApiGetNetworkingV1NetworkLinkEndpointRequest) Environment(environment string) ApiGetNetworkingV1NetworkLinkEndpointRequest {
	r.environment = &environment
	return r
}

func (r ApiGetNetworkingV1NetworkLinkEndpointRequest) Execute() (NetworkingV1NetworkLinkEndpoint, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkingV1NetworkLinkEndpointExecute(r)
}

/*
GetNetworkingV1NetworkLinkEndpoint Read a Network Link Endpoint

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to read a network link endpoint.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the network link endpoint.
 @return ApiGetNetworkingV1NetworkLinkEndpointRequest
*/
func (a *NetworkLinkEndpointsNetworkingV1ApiService) GetNetworkingV1NetworkLinkEndpoint(ctx _context.Context, id string) ApiGetNetworkingV1NetworkLinkEndpointRequest {
	return ApiGetNetworkingV1NetworkLinkEndpointRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return NetworkingV1NetworkLinkEndpoint
func (a *NetworkLinkEndpointsNetworkingV1ApiService) GetNetworkingV1NetworkLinkEndpointExecute(r ApiGetNetworkingV1NetworkLinkEndpointRequest) (NetworkingV1NetworkLinkEndpoint, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1NetworkLinkEndpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkLinkEndpointsNetworkingV1ApiService.GetNetworkingV1NetworkLinkEndpoint")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/network-link-endpoints/{id}"
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

type ApiListNetworkingV1NetworkLinkEndpointsRequest struct {
	ctx                    _context.Context
	ApiService             NetworkLinkEndpointsNetworkingV1Api
	environment            *string
	specDisplayName        *MultipleSearchFilter
	statusPhase            *MultipleSearchFilter
	specNetwork            *MultipleSearchFilter
	specNetworkLinkService *MultipleSearchFilter
	pageSize               *int32
	pageToken              *string
}

// Filter the results by exact match for environment.
func (r ApiListNetworkingV1NetworkLinkEndpointsRequest) Environment(environment string) ApiListNetworkingV1NetworkLinkEndpointsRequest {
	r.environment = &environment
	return r
}

// Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1NetworkLinkEndpointsRequest) SpecDisplayName(specDisplayName MultipleSearchFilter) ApiListNetworkingV1NetworkLinkEndpointsRequest {
	r.specDisplayName = &specDisplayName
	return r
}

// Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1NetworkLinkEndpointsRequest) StatusPhase(statusPhase MultipleSearchFilter) ApiListNetworkingV1NetworkLinkEndpointsRequest {
	r.statusPhase = &statusPhase
	return r
}

// Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1NetworkLinkEndpointsRequest) SpecNetwork(specNetwork MultipleSearchFilter) ApiListNetworkingV1NetworkLinkEndpointsRequest {
	r.specNetwork = &specNetwork
	return r
}

// Filter the results by exact match for spec.network_link_service. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1NetworkLinkEndpointsRequest) SpecNetworkLinkService(specNetworkLinkService MultipleSearchFilter) ApiListNetworkingV1NetworkLinkEndpointsRequest {
	r.specNetworkLinkService = &specNetworkLinkService
	return r
}

// A pagination size for collection requests.
func (r ApiListNetworkingV1NetworkLinkEndpointsRequest) PageSize(pageSize int32) ApiListNetworkingV1NetworkLinkEndpointsRequest {
	r.pageSize = &pageSize
	return r
}

// An opaque pagination token for collection requests.
func (r ApiListNetworkingV1NetworkLinkEndpointsRequest) PageToken(pageToken string) ApiListNetworkingV1NetworkLinkEndpointsRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListNetworkingV1NetworkLinkEndpointsRequest) Execute() (NetworkingV1NetworkLinkEndpointList, *_nethttp.Response, error) {
	return r.ApiService.ListNetworkingV1NetworkLinkEndpointsExecute(r)
}

/*
ListNetworkingV1NetworkLinkEndpoints List of Network Link Endpoints

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Retrieve a sorted, filtered, paginated list of all network link endpoints.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListNetworkingV1NetworkLinkEndpointsRequest
*/
func (a *NetworkLinkEndpointsNetworkingV1ApiService) ListNetworkingV1NetworkLinkEndpoints(ctx _context.Context) ApiListNetworkingV1NetworkLinkEndpointsRequest {
	return ApiListNetworkingV1NetworkLinkEndpointsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return NetworkingV1NetworkLinkEndpointList
func (a *NetworkLinkEndpointsNetworkingV1ApiService) ListNetworkingV1NetworkLinkEndpointsExecute(r ApiListNetworkingV1NetworkLinkEndpointsRequest) (NetworkingV1NetworkLinkEndpointList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1NetworkLinkEndpointList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkLinkEndpointsNetworkingV1ApiService.ListNetworkingV1NetworkLinkEndpoints")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/network-link-endpoints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.environment == nil {
		return localVarReturnValue, nil, reportError("environment is required and must be specified")
	}

	if r.specDisplayName != nil {
		localVarQueryParams.Add("spec.display_name", parameterToString(*r.specDisplayName, ""))
	}
	if r.statusPhase != nil {
		localVarQueryParams.Add("status.phase", parameterToString(*r.statusPhase, ""))
	}
	localVarQueryParams.Add("environment", parameterToString(*r.environment, ""))
	if r.specNetwork != nil {
		localVarQueryParams.Add("spec.network", parameterToString(*r.specNetwork, ""))
	}
	if r.specNetworkLinkService != nil {
		localVarQueryParams.Add("spec.network_link_service", parameterToString(*r.specNetworkLinkService, ""))
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

type ApiUpdateNetworkingV1NetworkLinkEndpointRequest struct {
	ctx                                   _context.Context
	ApiService                            NetworkLinkEndpointsNetworkingV1Api
	id                                    string
	networkingV1NetworkLinkEndpointUpdate *NetworkingV1NetworkLinkEndpointUpdate
}

func (r ApiUpdateNetworkingV1NetworkLinkEndpointRequest) NetworkingV1NetworkLinkEndpointUpdate(networkingV1NetworkLinkEndpointUpdate NetworkingV1NetworkLinkEndpointUpdate) ApiUpdateNetworkingV1NetworkLinkEndpointRequest {
	r.networkingV1NetworkLinkEndpointUpdate = &networkingV1NetworkLinkEndpointUpdate
	return r
}

func (r ApiUpdateNetworkingV1NetworkLinkEndpointRequest) Execute() (NetworkingV1NetworkLinkEndpoint, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkingV1NetworkLinkEndpointExecute(r)
}

/*
UpdateNetworkingV1NetworkLinkEndpoint Update a Network Link Endpoint

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to update a network link endpoint.



 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the network link endpoint.
 @return ApiUpdateNetworkingV1NetworkLinkEndpointRequest
*/
func (a *NetworkLinkEndpointsNetworkingV1ApiService) UpdateNetworkingV1NetworkLinkEndpoint(ctx _context.Context, id string) ApiUpdateNetworkingV1NetworkLinkEndpointRequest {
	return ApiUpdateNetworkingV1NetworkLinkEndpointRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return NetworkingV1NetworkLinkEndpoint
func (a *NetworkLinkEndpointsNetworkingV1ApiService) UpdateNetworkingV1NetworkLinkEndpointExecute(r ApiUpdateNetworkingV1NetworkLinkEndpointRequest) (NetworkingV1NetworkLinkEndpoint, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1NetworkLinkEndpoint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NetworkLinkEndpointsNetworkingV1ApiService.UpdateNetworkingV1NetworkLinkEndpoint")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/network-link-endpoints/{id}"
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
	localVarPostBody = r.networkingV1NetworkLinkEndpointUpdate
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
