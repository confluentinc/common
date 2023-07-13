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
)

// Linger please
var (
	_ _context.Context
)

type RegionsFcpmV2Api interface {

	/*
		ListFcpmV2Regions List of Regions

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Retrieve a sorted, filtered, paginated list of all regions.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return ApiListFcpmV2RegionsRequest
	*/
	ListFcpmV2Regions(ctx _context.Context) ApiListFcpmV2RegionsRequest

	// ListFcpmV2RegionsExecute executes the request
	//  @return FcpmV2RegionList
	ListFcpmV2RegionsExecute(r ApiListFcpmV2RegionsRequest) (FcpmV2RegionList, *_nethttp.Response, error)
}

// RegionsFcpmV2ApiService RegionsFcpmV2Api service
type RegionsFcpmV2ApiService service

type ApiListFcpmV2RegionsRequest struct {
	ctx        _context.Context
	ApiService RegionsFcpmV2Api
	cloud      *string
	regionName *string
	pageSize   *int32
	pageToken  *string
}

// Filter the results by exact match for cloud.
func (r ApiListFcpmV2RegionsRequest) Cloud(cloud string) ApiListFcpmV2RegionsRequest {
	r.cloud = &cloud
	return r
}

// Filter the results by exact match for region_name.
func (r ApiListFcpmV2RegionsRequest) RegionName(regionName string) ApiListFcpmV2RegionsRequest {
	r.regionName = &regionName
	return r
}

// A pagination size for collection requests.
func (r ApiListFcpmV2RegionsRequest) PageSize(pageSize int32) ApiListFcpmV2RegionsRequest {
	r.pageSize = &pageSize
	return r
}

// An opaque pagination token for collection requests.
func (r ApiListFcpmV2RegionsRequest) PageToken(pageToken string) ApiListFcpmV2RegionsRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListFcpmV2RegionsRequest) Execute() (FcpmV2RegionList, *_nethttp.Response, error) {
	return r.ApiService.ListFcpmV2RegionsExecute(r)
}

/*
ListFcpmV2Regions List of Regions

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Retrieve a sorted, filtered, paginated list of all regions.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListFcpmV2RegionsRequest
*/
func (a *RegionsFcpmV2ApiService) ListFcpmV2Regions(ctx _context.Context) ApiListFcpmV2RegionsRequest {
	return ApiListFcpmV2RegionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return FcpmV2RegionList
func (a *RegionsFcpmV2ApiService) ListFcpmV2RegionsExecute(r ApiListFcpmV2RegionsRequest) (FcpmV2RegionList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  FcpmV2RegionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegionsFcpmV2ApiService.ListFcpmV2Regions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/fcpm/v2/regions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.cloud != nil {
		localVarQueryParams.Add("cloud", parameterToString(*r.cloud, ""))
	}
	if r.regionName != nil {
		localVarQueryParams.Add("region_name", parameterToString(*r.regionName, ""))
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
