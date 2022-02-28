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
Quotas Service API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1-alpha0
Contact: api-framework-team@confluent.io
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

type ScopesQuotasV2Api interface {

	/*
	GetQuotasV2Scope Read a Scope

	[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To Quotas v2](https://img.shields.io/badge/-Request%20Access%20To%20Quotas%20v2-%23bc8540)](mailto:ccloud-api-access+quotas-v2-early-access@confluent.io?subject=Request%20to%20join%20quotas/v2%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20quotas/v2%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)

Make a request to read a scope.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param id The unique identifier for the scope.
	 @return ApiGetQuotasV2ScopeRequest
	*/
	GetQuotasV2Scope(ctx _context.Context, id string) ApiGetQuotasV2ScopeRequest

	// GetQuotasV2ScopeExecute executes the request
	//  @return QuotasV2Scope
	GetQuotasV2ScopeExecute(r ApiGetQuotasV2ScopeRequest) (QuotasV2Scope, *_nethttp.Response, error)

	/*
	ListQuotasV2Scopes List of Scopes

	[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To Quotas v2](https://img.shields.io/badge/-Request%20Access%20To%20Quotas%20v2-%23bc8540)](mailto:ccloud-api-access+quotas-v2-early-access@confluent.io?subject=Request%20to%20join%20quotas/v2%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20quotas/v2%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)

Retrieve a sorted, filtered, paginated list of all scopes.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @return ApiListQuotasV2ScopesRequest
	*/
	ListQuotasV2Scopes(ctx _context.Context) ApiListQuotasV2ScopesRequest

	// ListQuotasV2ScopesExecute executes the request
	//  @return QuotasV2ScopeList
	ListQuotasV2ScopesExecute(r ApiListQuotasV2ScopesRequest) (QuotasV2ScopeList, *_nethttp.Response, error)
}

// ScopesQuotasV2ApiService ScopesQuotasV2Api service
type ScopesQuotasV2ApiService service

type ApiGetQuotasV2ScopeRequest struct {
	ctx _context.Context
	ApiService ScopesQuotasV2Api
	id string
}


func (r ApiGetQuotasV2ScopeRequest) Execute() (QuotasV2Scope, *_nethttp.Response, error) {
	return r.ApiService.GetQuotasV2ScopeExecute(r)
}

/*
GetQuotasV2Scope Read a Scope

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To Quotas v2](https://img.shields.io/badge/-Request%20Access%20To%20Quotas%20v2-%23bc8540)](mailto:ccloud-api-access+quotas-v2-early-access@confluent.io?subject=Request%20to%20join%20quotas/v2%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20quotas/v2%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)

Make a request to read a scope.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the scope.
 @return ApiGetQuotasV2ScopeRequest
*/
func (a *ScopesQuotasV2ApiService) GetQuotasV2Scope(ctx _context.Context, id string) ApiGetQuotasV2ScopeRequest {
	return ApiGetQuotasV2ScopeRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return QuotasV2Scope
func (a *ScopesQuotasV2ApiService) GetQuotasV2ScopeExecute(r ApiGetQuotasV2ScopeRequest) (QuotasV2Scope, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  QuotasV2Scope
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScopesQuotasV2ApiService.GetQuotasV2Scope")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/quotas/v2/scopes/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiListQuotasV2ScopesRequest struct {
	ctx _context.Context
	ApiService ScopesQuotasV2Api
	pageSize *int32
	pageToken *string
}

// A pagination size for collection requests.
func (r ApiListQuotasV2ScopesRequest) PageSize(pageSize int32) ApiListQuotasV2ScopesRequest {
	r.pageSize = &pageSize
	return r
}
// An opaque pagination token for collection requests.
func (r ApiListQuotasV2ScopesRequest) PageToken(pageToken string) ApiListQuotasV2ScopesRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListQuotasV2ScopesRequest) Execute() (QuotasV2ScopeList, *_nethttp.Response, error) {
	return r.ApiService.ListQuotasV2ScopesExecute(r)
}

/*
ListQuotasV2Scopes List of Scopes

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To Quotas v2](https://img.shields.io/badge/-Request%20Access%20To%20Quotas%20v2-%23bc8540)](mailto:ccloud-api-access+quotas-v2-early-access@confluent.io?subject=Request%20to%20join%20quotas/v2%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20quotas/v2%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)

Retrieve a sorted, filtered, paginated list of all scopes.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListQuotasV2ScopesRequest
*/
func (a *ScopesQuotasV2ApiService) ListQuotasV2Scopes(ctx _context.Context) ApiListQuotasV2ScopesRequest {
	return ApiListQuotasV2ScopesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return QuotasV2ScopeList
func (a *ScopesQuotasV2ApiService) ListQuotasV2ScopesExecute(r ApiListQuotasV2ScopesRequest) (QuotasV2ScopeList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  QuotasV2ScopeList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScopesQuotasV2ApiService.ListQuotasV2Scopes")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/quotas/v2/scopes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
