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

type PrivateLinkAccessesNetworkingV1Api interface {

	/*
	CreateNetworkingV1PrivateLinkAccess Create a Private Link Access

	[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to create a private link access.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiCreateNetworkingV1PrivateLinkAccessRequest
	*/
	CreateNetworkingV1PrivateLinkAccess(ctx _context.Context) ApiCreateNetworkingV1PrivateLinkAccessRequest

	// CreateNetworkingV1PrivateLinkAccessExecute executes the request
	//  @return NetworkingV1PrivateLinkAccess
	CreateNetworkingV1PrivateLinkAccessExecute(r ApiCreateNetworkingV1PrivateLinkAccessRequest) (NetworkingV1PrivateLinkAccess, *_nethttp.Response, error)

	/*
	DeleteNetworkingV1PrivateLinkAccess Delete a Private Link Access

	[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to delete a private link access.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param id The unique identifier for the private link access.
	 @return ApiDeleteNetworkingV1PrivateLinkAccessRequest
	*/
	DeleteNetworkingV1PrivateLinkAccess(ctx _context.Context, id string) ApiDeleteNetworkingV1PrivateLinkAccessRequest

	// DeleteNetworkingV1PrivateLinkAccessExecute executes the request
	DeleteNetworkingV1PrivateLinkAccessExecute(r ApiDeleteNetworkingV1PrivateLinkAccessRequest) (*_nethttp.Response, error)

	/*
	GetNetworkingV1PrivateLinkAccess Read a Private Link Access

	[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to read a private link access.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param id The unique identifier for the private link access.
	 @return ApiGetNetworkingV1PrivateLinkAccessRequest
	*/
	GetNetworkingV1PrivateLinkAccess(ctx _context.Context, id string) ApiGetNetworkingV1PrivateLinkAccessRequest

	// GetNetworkingV1PrivateLinkAccessExecute executes the request
	//  @return NetworkingV1PrivateLinkAccess
	GetNetworkingV1PrivateLinkAccessExecute(r ApiGetNetworkingV1PrivateLinkAccessRequest) (NetworkingV1PrivateLinkAccess, *_nethttp.Response, error)

	/*
	ListNetworkingV1PrivateLinkAccesses List of Private Link Accesses

	[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Retrieve a sorted, filtered, paginated list of all private link accesses.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiListNetworkingV1PrivateLinkAccessesRequest
	*/
	ListNetworkingV1PrivateLinkAccesses(ctx _context.Context) ApiListNetworkingV1PrivateLinkAccessesRequest

	// ListNetworkingV1PrivateLinkAccessesExecute executes the request
	//  @return NetworkingV1PrivateLinkAccessList
	ListNetworkingV1PrivateLinkAccessesExecute(r ApiListNetworkingV1PrivateLinkAccessesRequest) (NetworkingV1PrivateLinkAccessList, *_nethttp.Response, error)

	/*
	UpdateNetworkingV1PrivateLinkAccess Update a Private Link Access

	[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to update a private link access.



	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param id The unique identifier for the private link access.
	 @return ApiUpdateNetworkingV1PrivateLinkAccessRequest
	*/
	UpdateNetworkingV1PrivateLinkAccess(ctx _context.Context, id string) ApiUpdateNetworkingV1PrivateLinkAccessRequest

	// UpdateNetworkingV1PrivateLinkAccessExecute executes the request
	//  @return NetworkingV1PrivateLinkAccess
	UpdateNetworkingV1PrivateLinkAccessExecute(r ApiUpdateNetworkingV1PrivateLinkAccessRequest) (NetworkingV1PrivateLinkAccess, *_nethttp.Response, error)
}

// PrivateLinkAccessesNetworkingV1ApiService PrivateLinkAccessesNetworkingV1Api service
type PrivateLinkAccessesNetworkingV1ApiService service

type ApiCreateNetworkingV1PrivateLinkAccessRequest struct {
	ctx _context.Context
	ApiService PrivateLinkAccessesNetworkingV1Api
	networkingV1PrivateLinkAccess *NetworkingV1PrivateLinkAccess
}

func (r ApiCreateNetworkingV1PrivateLinkAccessRequest) NetworkingV1PrivateLinkAccess(networkingV1PrivateLinkAccess NetworkingV1PrivateLinkAccess) ApiCreateNetworkingV1PrivateLinkAccessRequest {
	r.networkingV1PrivateLinkAccess = &networkingV1PrivateLinkAccess
	return r
}

func (r ApiCreateNetworkingV1PrivateLinkAccessRequest) Execute() (NetworkingV1PrivateLinkAccess, *_nethttp.Response, error) {
	return r.ApiService.CreateNetworkingV1PrivateLinkAccessExecute(r)
}

/*
CreateNetworkingV1PrivateLinkAccess Create a Private Link Access

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to create a private link access.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateNetworkingV1PrivateLinkAccessRequest
*/
func (a *PrivateLinkAccessesNetworkingV1ApiService) CreateNetworkingV1PrivateLinkAccess(ctx _context.Context) ApiCreateNetworkingV1PrivateLinkAccessRequest {
	return ApiCreateNetworkingV1PrivateLinkAccessRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NetworkingV1PrivateLinkAccess
func (a *PrivateLinkAccessesNetworkingV1ApiService) CreateNetworkingV1PrivateLinkAccessExecute(r ApiCreateNetworkingV1PrivateLinkAccessRequest) (NetworkingV1PrivateLinkAccess, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1PrivateLinkAccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateLinkAccessesNetworkingV1ApiService.CreateNetworkingV1PrivateLinkAccess")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/private-link-accesses"

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
	localVarPostBody = r.networkingV1PrivateLinkAccess
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

type ApiDeleteNetworkingV1PrivateLinkAccessRequest struct {
	ctx _context.Context
	ApiService PrivateLinkAccessesNetworkingV1Api
	environment *string
	id string
}

// Scope the operation to the given environment.
func (r ApiDeleteNetworkingV1PrivateLinkAccessRequest) Environment(environment string) ApiDeleteNetworkingV1PrivateLinkAccessRequest {
	r.environment = &environment
	return r
}

func (r ApiDeleteNetworkingV1PrivateLinkAccessRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteNetworkingV1PrivateLinkAccessExecute(r)
}

/*
DeleteNetworkingV1PrivateLinkAccess Delete a Private Link Access

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to delete a private link access.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the private link access.
 @return ApiDeleteNetworkingV1PrivateLinkAccessRequest
*/
func (a *PrivateLinkAccessesNetworkingV1ApiService) DeleteNetworkingV1PrivateLinkAccess(ctx _context.Context, id string) ApiDeleteNetworkingV1PrivateLinkAccessRequest {
	return ApiDeleteNetworkingV1PrivateLinkAccessRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *PrivateLinkAccessesNetworkingV1ApiService) DeleteNetworkingV1PrivateLinkAccessExecute(r ApiDeleteNetworkingV1PrivateLinkAccessRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateLinkAccessesNetworkingV1ApiService.DeleteNetworkingV1PrivateLinkAccess")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/private-link-accesses/{id}"
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

type ApiGetNetworkingV1PrivateLinkAccessRequest struct {
	ctx _context.Context
	ApiService PrivateLinkAccessesNetworkingV1Api
	environment *string
	id string
}

// Scope the operation to the given environment.
func (r ApiGetNetworkingV1PrivateLinkAccessRequest) Environment(environment string) ApiGetNetworkingV1PrivateLinkAccessRequest {
	r.environment = &environment
	return r
}

func (r ApiGetNetworkingV1PrivateLinkAccessRequest) Execute() (NetworkingV1PrivateLinkAccess, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkingV1PrivateLinkAccessExecute(r)
}

/*
GetNetworkingV1PrivateLinkAccess Read a Private Link Access

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to read a private link access.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the private link access.
 @return ApiGetNetworkingV1PrivateLinkAccessRequest
*/
func (a *PrivateLinkAccessesNetworkingV1ApiService) GetNetworkingV1PrivateLinkAccess(ctx _context.Context, id string) ApiGetNetworkingV1PrivateLinkAccessRequest {
	return ApiGetNetworkingV1PrivateLinkAccessRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NetworkingV1PrivateLinkAccess
func (a *PrivateLinkAccessesNetworkingV1ApiService) GetNetworkingV1PrivateLinkAccessExecute(r ApiGetNetworkingV1PrivateLinkAccessRequest) (NetworkingV1PrivateLinkAccess, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1PrivateLinkAccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateLinkAccessesNetworkingV1ApiService.GetNetworkingV1PrivateLinkAccess")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/private-link-accesses/{id}"
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

type ApiListNetworkingV1PrivateLinkAccessesRequest struct {
	ctx _context.Context
	ApiService PrivateLinkAccessesNetworkingV1Api
	environment *string
	specDisplayName *MultipleSearchFilter
	statusPhase *MultipleSearchFilter
	specNetwork *MultipleSearchFilter
	pageSize *int32
	pageToken *string
}

// Filter the results by exact match for environment.
func (r ApiListNetworkingV1PrivateLinkAccessesRequest) Environment(environment string) ApiListNetworkingV1PrivateLinkAccessesRequest {
	r.environment = &environment
	return r
}
// Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1PrivateLinkAccessesRequest) SpecDisplayName(specDisplayName MultipleSearchFilter) ApiListNetworkingV1PrivateLinkAccessesRequest {
	r.specDisplayName = &specDisplayName
	return r
}
// Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1PrivateLinkAccessesRequest) StatusPhase(statusPhase MultipleSearchFilter) ApiListNetworkingV1PrivateLinkAccessesRequest {
	r.statusPhase = &statusPhase
	return r
}
// Filter the results by exact match for spec.network. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1PrivateLinkAccessesRequest) SpecNetwork(specNetwork MultipleSearchFilter) ApiListNetworkingV1PrivateLinkAccessesRequest {
	r.specNetwork = &specNetwork
	return r
}
// A pagination size for collection requests.
func (r ApiListNetworkingV1PrivateLinkAccessesRequest) PageSize(pageSize int32) ApiListNetworkingV1PrivateLinkAccessesRequest {
	r.pageSize = &pageSize
	return r
}
// An opaque pagination token for collection requests.
func (r ApiListNetworkingV1PrivateLinkAccessesRequest) PageToken(pageToken string) ApiListNetworkingV1PrivateLinkAccessesRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListNetworkingV1PrivateLinkAccessesRequest) Execute() (NetworkingV1PrivateLinkAccessList, *_nethttp.Response, error) {
	return r.ApiService.ListNetworkingV1PrivateLinkAccessesExecute(r)
}

/*
ListNetworkingV1PrivateLinkAccesses List of Private Link Accesses

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Retrieve a sorted, filtered, paginated list of all private link accesses.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListNetworkingV1PrivateLinkAccessesRequest
*/
func (a *PrivateLinkAccessesNetworkingV1ApiService) ListNetworkingV1PrivateLinkAccesses(ctx _context.Context) ApiListNetworkingV1PrivateLinkAccessesRequest {
	return ApiListNetworkingV1PrivateLinkAccessesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NetworkingV1PrivateLinkAccessList
func (a *PrivateLinkAccessesNetworkingV1ApiService) ListNetworkingV1PrivateLinkAccessesExecute(r ApiListNetworkingV1PrivateLinkAccessesRequest) (NetworkingV1PrivateLinkAccessList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1PrivateLinkAccessList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateLinkAccessesNetworkingV1ApiService.ListNetworkingV1PrivateLinkAccesses")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/private-link-accesses"

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

type ApiUpdateNetworkingV1PrivateLinkAccessRequest struct {
	ctx _context.Context
	ApiService PrivateLinkAccessesNetworkingV1Api
	id string
	networkingV1PrivateLinkAccessUpdate *NetworkingV1PrivateLinkAccessUpdate
}

func (r ApiUpdateNetworkingV1PrivateLinkAccessRequest) NetworkingV1PrivateLinkAccessUpdate(networkingV1PrivateLinkAccessUpdate NetworkingV1PrivateLinkAccessUpdate) ApiUpdateNetworkingV1PrivateLinkAccessRequest {
	r.networkingV1PrivateLinkAccessUpdate = &networkingV1PrivateLinkAccessUpdate
	return r
}

func (r ApiUpdateNetworkingV1PrivateLinkAccessRequest) Execute() (NetworkingV1PrivateLinkAccess, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkingV1PrivateLinkAccessExecute(r)
}

/*
UpdateNetworkingV1PrivateLinkAccess Update a Private Link Access

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Make a request to update a private link access.



 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the private link access.
 @return ApiUpdateNetworkingV1PrivateLinkAccessRequest
*/
func (a *PrivateLinkAccessesNetworkingV1ApiService) UpdateNetworkingV1PrivateLinkAccess(ctx _context.Context, id string) ApiUpdateNetworkingV1PrivateLinkAccessRequest {
	return ApiUpdateNetworkingV1PrivateLinkAccessRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return NetworkingV1PrivateLinkAccess
func (a *PrivateLinkAccessesNetworkingV1ApiService) UpdateNetworkingV1PrivateLinkAccessExecute(r ApiUpdateNetworkingV1PrivateLinkAccessRequest) (NetworkingV1PrivateLinkAccess, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1PrivateLinkAccess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateLinkAccessesNetworkingV1ApiService.UpdateNetworkingV1PrivateLinkAccess")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/private-link-accesses/{id}"
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
	localVarPostBody = r.networkingV1PrivateLinkAccessUpdate
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
