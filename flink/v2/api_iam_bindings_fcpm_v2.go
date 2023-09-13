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
Flink Compute Pool Management API

This is the Flink Compute Pool management API.

API version: 0.0.1
Contact: ksql-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v2

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

type IamBindingsFcpmV2Api interface {

	/*
			CreateFcpmV2IamBinding Create an Iam Binding

			[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to create an iam binding.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @return ApiCreateFcpmV2IamBindingRequest
	*/
	CreateFcpmV2IamBinding(ctx _context.Context) ApiCreateFcpmV2IamBindingRequest

	// CreateFcpmV2IamBindingExecute executes the request
	//  @return FcpmV2IamBinding
	CreateFcpmV2IamBindingExecute(r ApiCreateFcpmV2IamBindingRequest) (FcpmV2IamBinding, *_nethttp.Response, error)

	/*
			DeleteFcpmV2IamBinding Delete an Iam Binding

			[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

		Make a request to delete an iam binding.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param id The unique identifier for the iam binding.
			 @return ApiDeleteFcpmV2IamBindingRequest
	*/
	DeleteFcpmV2IamBinding(ctx _context.Context, id string) ApiDeleteFcpmV2IamBindingRequest

	// DeleteFcpmV2IamBindingExecute executes the request
	DeleteFcpmV2IamBindingExecute(r ApiDeleteFcpmV2IamBindingRequest) (*_nethttp.Response, error)

	/*
			ListFcpmV2IamBindings List of Iam Bindings

			[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

		Retrieve a sorted, filtered, paginated list of all iam bindings.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @return ApiListFcpmV2IamBindingsRequest
	*/
	ListFcpmV2IamBindings(ctx _context.Context) ApiListFcpmV2IamBindingsRequest

	// ListFcpmV2IamBindingsExecute executes the request
	//  @return FcpmV2IamBindingList
	ListFcpmV2IamBindingsExecute(r ApiListFcpmV2IamBindingsRequest) (FcpmV2IamBindingList, *_nethttp.Response, error)
}

// IamBindingsFcpmV2ApiService IamBindingsFcpmV2Api service
type IamBindingsFcpmV2ApiService service

type ApiCreateFcpmV2IamBindingRequest struct {
	ctx              _context.Context
	ApiService       IamBindingsFcpmV2Api
	fcpmV2IamBinding *FcpmV2IamBinding
}

func (r ApiCreateFcpmV2IamBindingRequest) FcpmV2IamBinding(fcpmV2IamBinding FcpmV2IamBinding) ApiCreateFcpmV2IamBindingRequest {
	r.fcpmV2IamBinding = &fcpmV2IamBinding
	return r
}

func (r ApiCreateFcpmV2IamBindingRequest) Execute() (FcpmV2IamBinding, *_nethttp.Response, error) {
	return r.ApiService.CreateFcpmV2IamBindingExecute(r)
}

/*
CreateFcpmV2IamBinding Create an Iam Binding

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Make a request to create an iam binding.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateFcpmV2IamBindingRequest
*/
func (a *IamBindingsFcpmV2ApiService) CreateFcpmV2IamBinding(ctx _context.Context) ApiCreateFcpmV2IamBindingRequest {
	return ApiCreateFcpmV2IamBindingRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FcpmV2IamBinding
func (a *IamBindingsFcpmV2ApiService) CreateFcpmV2IamBindingExecute(r ApiCreateFcpmV2IamBindingRequest) (FcpmV2IamBinding, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FcpmV2IamBinding
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamBindingsFcpmV2ApiService.CreateFcpmV2IamBinding")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fcpm/v2/iam-bindings"

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
	localVarPostBody = r.fcpmV2IamBinding
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

type ApiDeleteFcpmV2IamBindingRequest struct {
	ctx          _context.Context
	ApiService   IamBindingsFcpmV2Api
	environment  *string
	id           string
	identityPool *string
}

// Scope the operation to the given environment.
func (r ApiDeleteFcpmV2IamBindingRequest) Environment(environment string) ApiDeleteFcpmV2IamBindingRequest {
	r.environment = &environment
	return r
}

// Scope the operation to the given identity_pool.
func (r ApiDeleteFcpmV2IamBindingRequest) IdentityPool(identityPool string) ApiDeleteFcpmV2IamBindingRequest {
	r.identityPool = &identityPool
	return r
}

func (r ApiDeleteFcpmV2IamBindingRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteFcpmV2IamBindingExecute(r)
}

/*
DeleteFcpmV2IamBinding Delete an Iam Binding

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Make a request to delete an iam binding.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The unique identifier for the iam binding.
	@return ApiDeleteFcpmV2IamBindingRequest
*/
func (a *IamBindingsFcpmV2ApiService) DeleteFcpmV2IamBinding(ctx _context.Context, id string) ApiDeleteFcpmV2IamBindingRequest {
	return ApiDeleteFcpmV2IamBindingRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *IamBindingsFcpmV2ApiService) DeleteFcpmV2IamBindingExecute(r ApiDeleteFcpmV2IamBindingRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamBindingsFcpmV2ApiService.DeleteFcpmV2IamBinding")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fcpm/v2/iam-bindings/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.environment == nil {
		return nil, reportError("environment is required and must be specified")
	}

	localVarQueryParams.Add("environment", parameterToString(*r.environment, ""))
	if r.identityPool != nil {
		localVarQueryParams.Add("identity_pool", parameterToString(*r.identityPool, ""))
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

type ApiListFcpmV2IamBindingsRequest struct {
	ctx          _context.Context
	ApiService   IamBindingsFcpmV2Api
	environment  *string
	region       *string
	cloud        *string
	identityPool *string
	pageSize     *int32
	pageToken    *string
}

// Filter the results by exact match for environment.
func (r ApiListFcpmV2IamBindingsRequest) Environment(environment string) ApiListFcpmV2IamBindingsRequest {
	r.environment = &environment
	return r
}

// Filter the results by exact match for region.
func (r ApiListFcpmV2IamBindingsRequest) Region(region string) ApiListFcpmV2IamBindingsRequest {
	r.region = &region
	return r
}

// Filter the results by exact match for cloud.
func (r ApiListFcpmV2IamBindingsRequest) Cloud(cloud string) ApiListFcpmV2IamBindingsRequest {
	r.cloud = &cloud
	return r
}

// Filter the results by exact match for identity_pool.
func (r ApiListFcpmV2IamBindingsRequest) IdentityPool(identityPool string) ApiListFcpmV2IamBindingsRequest {
	r.identityPool = &identityPool
	return r
}

// A pagination size for collection requests.
func (r ApiListFcpmV2IamBindingsRequest) PageSize(pageSize int32) ApiListFcpmV2IamBindingsRequest {
	r.pageSize = &pageSize
	return r
}

// An opaque pagination token for collection requests.
func (r ApiListFcpmV2IamBindingsRequest) PageToken(pageToken string) ApiListFcpmV2IamBindingsRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListFcpmV2IamBindingsRequest) Execute() (FcpmV2IamBindingList, *_nethttp.Response, error) {
	return r.ApiService.ListFcpmV2IamBindingsExecute(r)
}

/*
ListFcpmV2IamBindings List of Iam Bindings

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Retrieve a sorted, filtered, paginated list of all iam bindings.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListFcpmV2IamBindingsRequest
*/
func (a *IamBindingsFcpmV2ApiService) ListFcpmV2IamBindings(ctx _context.Context) ApiListFcpmV2IamBindingsRequest {
	return ApiListFcpmV2IamBindingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return FcpmV2IamBindingList
func (a *IamBindingsFcpmV2ApiService) ListFcpmV2IamBindingsExecute(r ApiListFcpmV2IamBindingsRequest) (FcpmV2IamBindingList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FcpmV2IamBindingList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IamBindingsFcpmV2ApiService.ListFcpmV2IamBindings")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fcpm/v2/iam-bindings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.environment == nil {
		return localVarReturnValue, nil, reportError("environment is required and must be specified")
	}

	if r.region != nil {
		localVarQueryParams.Add("region", parameterToString(*r.region, ""))
	}
	if r.cloud != nil {
		localVarQueryParams.Add("cloud", parameterToString(*r.cloud, ""))
	}
	localVarQueryParams.Add("environment", parameterToString(*r.environment, ""))
	if r.identityPool != nil {
		localVarQueryParams.Add("identity_pool", parameterToString(*r.identityPool, ""))
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
