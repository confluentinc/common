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

type PrivateLinkAttachmentsNetworkingV1Api interface {

	/*
		CreateNetworkingV1PrivateLinkAttachment Create a Private Link Attachment

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Make a request to create a private link attachment.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return ApiCreateNetworkingV1PrivateLinkAttachmentRequest
	*/
	CreateNetworkingV1PrivateLinkAttachment(ctx _context.Context) ApiCreateNetworkingV1PrivateLinkAttachmentRequest

	// CreateNetworkingV1PrivateLinkAttachmentExecute executes the request
	//  @return NetworkingV1PrivateLinkAttachment
	CreateNetworkingV1PrivateLinkAttachmentExecute(r ApiCreateNetworkingV1PrivateLinkAttachmentRequest) (NetworkingV1PrivateLinkAttachment, *_nethttp.Response, error)

	/*
		DeleteNetworkingV1PrivateLinkAttachment Delete a Private Link Attachment

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Make a request to delete a private link attachment.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param id The unique identifier for the private link attachment.
		 @return ApiDeleteNetworkingV1PrivateLinkAttachmentRequest
	*/
	DeleteNetworkingV1PrivateLinkAttachment(ctx _context.Context, id string) ApiDeleteNetworkingV1PrivateLinkAttachmentRequest

	// DeleteNetworkingV1PrivateLinkAttachmentExecute executes the request
	DeleteNetworkingV1PrivateLinkAttachmentExecute(r ApiDeleteNetworkingV1PrivateLinkAttachmentRequest) (*_nethttp.Response, error)

	/*
		GetNetworkingV1PrivateLinkAttachment Read a Private Link Attachment

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Make a request to read a private link attachment.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param id The unique identifier for the private link attachment.
		 @return ApiGetNetworkingV1PrivateLinkAttachmentRequest
	*/
	GetNetworkingV1PrivateLinkAttachment(ctx _context.Context, id string) ApiGetNetworkingV1PrivateLinkAttachmentRequest

	// GetNetworkingV1PrivateLinkAttachmentExecute executes the request
	//  @return NetworkingV1PrivateLinkAttachment
	GetNetworkingV1PrivateLinkAttachmentExecute(r ApiGetNetworkingV1PrivateLinkAttachmentRequest) (NetworkingV1PrivateLinkAttachment, *_nethttp.Response, error)

	/*
		ListNetworkingV1PrivateLinkAttachments List of Private Link Attachments

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Retrieve a sorted, filtered, paginated list of all private link attachments.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return ApiListNetworkingV1PrivateLinkAttachmentsRequest
	*/
	ListNetworkingV1PrivateLinkAttachments(ctx _context.Context) ApiListNetworkingV1PrivateLinkAttachmentsRequest

	// ListNetworkingV1PrivateLinkAttachmentsExecute executes the request
	//  @return NetworkingV1PrivateLinkAttachmentList
	ListNetworkingV1PrivateLinkAttachmentsExecute(r ApiListNetworkingV1PrivateLinkAttachmentsRequest) (NetworkingV1PrivateLinkAttachmentList, *_nethttp.Response, error)

	/*
		UpdateNetworkingV1PrivateLinkAttachment Update a Private Link Attachment

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Make a request to update a private link attachment.



		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param id The unique identifier for the private link attachment.
		 @return ApiUpdateNetworkingV1PrivateLinkAttachmentRequest
	*/
	UpdateNetworkingV1PrivateLinkAttachment(ctx _context.Context, id string) ApiUpdateNetworkingV1PrivateLinkAttachmentRequest

	// UpdateNetworkingV1PrivateLinkAttachmentExecute executes the request
	//  @return NetworkingV1PrivateLinkAttachment
	UpdateNetworkingV1PrivateLinkAttachmentExecute(r ApiUpdateNetworkingV1PrivateLinkAttachmentRequest) (NetworkingV1PrivateLinkAttachment, *_nethttp.Response, error)
}

// PrivateLinkAttachmentsNetworkingV1ApiService PrivateLinkAttachmentsNetworkingV1Api service
type PrivateLinkAttachmentsNetworkingV1ApiService service

type ApiCreateNetworkingV1PrivateLinkAttachmentRequest struct {
	ctx                               _context.Context
	ApiService                        PrivateLinkAttachmentsNetworkingV1Api
	networkingV1PrivateLinkAttachment *NetworkingV1PrivateLinkAttachment
}

func (r ApiCreateNetworkingV1PrivateLinkAttachmentRequest) NetworkingV1PrivateLinkAttachment(networkingV1PrivateLinkAttachment NetworkingV1PrivateLinkAttachment) ApiCreateNetworkingV1PrivateLinkAttachmentRequest {
	r.networkingV1PrivateLinkAttachment = &networkingV1PrivateLinkAttachment
	return r
}

func (r ApiCreateNetworkingV1PrivateLinkAttachmentRequest) Execute() (NetworkingV1PrivateLinkAttachment, *_nethttp.Response, error) {
	return r.ApiService.CreateNetworkingV1PrivateLinkAttachmentExecute(r)
}

/*
CreateNetworkingV1PrivateLinkAttachment Create a Private Link Attachment

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Make a request to create a private link attachment.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateNetworkingV1PrivateLinkAttachmentRequest
*/
func (a *PrivateLinkAttachmentsNetworkingV1ApiService) CreateNetworkingV1PrivateLinkAttachment(ctx _context.Context) ApiCreateNetworkingV1PrivateLinkAttachmentRequest {
	return ApiCreateNetworkingV1PrivateLinkAttachmentRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return NetworkingV1PrivateLinkAttachment
func (a *PrivateLinkAttachmentsNetworkingV1ApiService) CreateNetworkingV1PrivateLinkAttachmentExecute(r ApiCreateNetworkingV1PrivateLinkAttachmentRequest) (NetworkingV1PrivateLinkAttachment, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1PrivateLinkAttachment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateLinkAttachmentsNetworkingV1ApiService.CreateNetworkingV1PrivateLinkAttachment")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/private-link-attachments"

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
	localVarPostBody = r.networkingV1PrivateLinkAttachment
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

type ApiDeleteNetworkingV1PrivateLinkAttachmentRequest struct {
	ctx         _context.Context
	ApiService  PrivateLinkAttachmentsNetworkingV1Api
	environment *string
	id          string
}

// Scope the operation to the given environment.
func (r ApiDeleteNetworkingV1PrivateLinkAttachmentRequest) Environment(environment string) ApiDeleteNetworkingV1PrivateLinkAttachmentRequest {
	r.environment = &environment
	return r
}

func (r ApiDeleteNetworkingV1PrivateLinkAttachmentRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteNetworkingV1PrivateLinkAttachmentExecute(r)
}

/*
DeleteNetworkingV1PrivateLinkAttachment Delete a Private Link Attachment

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Make a request to delete a private link attachment.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the private link attachment.
 @return ApiDeleteNetworkingV1PrivateLinkAttachmentRequest
*/
func (a *PrivateLinkAttachmentsNetworkingV1ApiService) DeleteNetworkingV1PrivateLinkAttachment(ctx _context.Context, id string) ApiDeleteNetworkingV1PrivateLinkAttachmentRequest {
	return ApiDeleteNetworkingV1PrivateLinkAttachmentRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *PrivateLinkAttachmentsNetworkingV1ApiService) DeleteNetworkingV1PrivateLinkAttachmentExecute(r ApiDeleteNetworkingV1PrivateLinkAttachmentRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateLinkAttachmentsNetworkingV1ApiService.DeleteNetworkingV1PrivateLinkAttachment")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/private-link-attachments/{id}"
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

type ApiGetNetworkingV1PrivateLinkAttachmentRequest struct {
	ctx         _context.Context
	ApiService  PrivateLinkAttachmentsNetworkingV1Api
	environment *string
	id          string
}

// Scope the operation to the given environment.
func (r ApiGetNetworkingV1PrivateLinkAttachmentRequest) Environment(environment string) ApiGetNetworkingV1PrivateLinkAttachmentRequest {
	r.environment = &environment
	return r
}

func (r ApiGetNetworkingV1PrivateLinkAttachmentRequest) Execute() (NetworkingV1PrivateLinkAttachment, *_nethttp.Response, error) {
	return r.ApiService.GetNetworkingV1PrivateLinkAttachmentExecute(r)
}

/*
GetNetworkingV1PrivateLinkAttachment Read a Private Link Attachment

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Make a request to read a private link attachment.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the private link attachment.
 @return ApiGetNetworkingV1PrivateLinkAttachmentRequest
*/
func (a *PrivateLinkAttachmentsNetworkingV1ApiService) GetNetworkingV1PrivateLinkAttachment(ctx _context.Context, id string) ApiGetNetworkingV1PrivateLinkAttachmentRequest {
	return ApiGetNetworkingV1PrivateLinkAttachmentRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return NetworkingV1PrivateLinkAttachment
func (a *PrivateLinkAttachmentsNetworkingV1ApiService) GetNetworkingV1PrivateLinkAttachmentExecute(r ApiGetNetworkingV1PrivateLinkAttachmentRequest) (NetworkingV1PrivateLinkAttachment, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1PrivateLinkAttachment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateLinkAttachmentsNetworkingV1ApiService.GetNetworkingV1PrivateLinkAttachment")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/private-link-attachments/{id}"
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

type ApiListNetworkingV1PrivateLinkAttachmentsRequest struct {
	ctx             _context.Context
	ApiService      PrivateLinkAttachmentsNetworkingV1Api
	environment     *string
	specDisplayName *MultipleSearchFilter
	specCloud       *MultipleSearchFilter
	specRegion      *MultipleSearchFilter
	statusPhase     *MultipleSearchFilter
	pageSize        *int32
	pageToken       *string
}

// Filter the results by exact match for environment.
func (r ApiListNetworkingV1PrivateLinkAttachmentsRequest) Environment(environment string) ApiListNetworkingV1PrivateLinkAttachmentsRequest {
	r.environment = &environment
	return r
}

// Filter the results by exact match for spec.display_name. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1PrivateLinkAttachmentsRequest) SpecDisplayName(specDisplayName MultipleSearchFilter) ApiListNetworkingV1PrivateLinkAttachmentsRequest {
	r.specDisplayName = &specDisplayName
	return r
}

// Filter the results by exact match for spec.cloud. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1PrivateLinkAttachmentsRequest) SpecCloud(specCloud MultipleSearchFilter) ApiListNetworkingV1PrivateLinkAttachmentsRequest {
	r.specCloud = &specCloud
	return r
}

// Filter the results by exact match for spec.region. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1PrivateLinkAttachmentsRequest) SpecRegion(specRegion MultipleSearchFilter) ApiListNetworkingV1PrivateLinkAttachmentsRequest {
	r.specRegion = &specRegion
	return r
}

// Filter the results by exact match for status.phase. Pass multiple times to see results matching any of the values.
func (r ApiListNetworkingV1PrivateLinkAttachmentsRequest) StatusPhase(statusPhase MultipleSearchFilter) ApiListNetworkingV1PrivateLinkAttachmentsRequest {
	r.statusPhase = &statusPhase
	return r
}

// A pagination size for collection requests.
func (r ApiListNetworkingV1PrivateLinkAttachmentsRequest) PageSize(pageSize int32) ApiListNetworkingV1PrivateLinkAttachmentsRequest {
	r.pageSize = &pageSize
	return r
}

// An opaque pagination token for collection requests.
func (r ApiListNetworkingV1PrivateLinkAttachmentsRequest) PageToken(pageToken string) ApiListNetworkingV1PrivateLinkAttachmentsRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListNetworkingV1PrivateLinkAttachmentsRequest) Execute() (NetworkingV1PrivateLinkAttachmentList, *_nethttp.Response, error) {
	return r.ApiService.ListNetworkingV1PrivateLinkAttachmentsExecute(r)
}

/*
ListNetworkingV1PrivateLinkAttachments List of Private Link Attachments

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Retrieve a sorted, filtered, paginated list of all private link attachments.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListNetworkingV1PrivateLinkAttachmentsRequest
*/
func (a *PrivateLinkAttachmentsNetworkingV1ApiService) ListNetworkingV1PrivateLinkAttachments(ctx _context.Context) ApiListNetworkingV1PrivateLinkAttachmentsRequest {
	return ApiListNetworkingV1PrivateLinkAttachmentsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return NetworkingV1PrivateLinkAttachmentList
func (a *PrivateLinkAttachmentsNetworkingV1ApiService) ListNetworkingV1PrivateLinkAttachmentsExecute(r ApiListNetworkingV1PrivateLinkAttachmentsRequest) (NetworkingV1PrivateLinkAttachmentList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1PrivateLinkAttachmentList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateLinkAttachmentsNetworkingV1ApiService.ListNetworkingV1PrivateLinkAttachments")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/private-link-attachments"

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
	if r.statusPhase != nil {
		localVarQueryParams.Add("status.phase", parameterToString(*r.statusPhase, ""))
	}
	localVarQueryParams.Add("environment", parameterToString(*r.environment, ""))
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

type ApiUpdateNetworkingV1PrivateLinkAttachmentRequest struct {
	ctx                                     _context.Context
	ApiService                              PrivateLinkAttachmentsNetworkingV1Api
	id                                      string
	networkingV1PrivateLinkAttachmentUpdate *NetworkingV1PrivateLinkAttachmentUpdate
}

func (r ApiUpdateNetworkingV1PrivateLinkAttachmentRequest) NetworkingV1PrivateLinkAttachmentUpdate(networkingV1PrivateLinkAttachmentUpdate NetworkingV1PrivateLinkAttachmentUpdate) ApiUpdateNetworkingV1PrivateLinkAttachmentRequest {
	r.networkingV1PrivateLinkAttachmentUpdate = &networkingV1PrivateLinkAttachmentUpdate
	return r
}

func (r ApiUpdateNetworkingV1PrivateLinkAttachmentRequest) Execute() (NetworkingV1PrivateLinkAttachment, *_nethttp.Response, error) {
	return r.ApiService.UpdateNetworkingV1PrivateLinkAttachmentExecute(r)
}

/*
UpdateNetworkingV1PrivateLinkAttachment Update a Private Link Attachment

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Make a request to update a private link attachment.



 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The unique identifier for the private link attachment.
 @return ApiUpdateNetworkingV1PrivateLinkAttachmentRequest
*/
func (a *PrivateLinkAttachmentsNetworkingV1ApiService) UpdateNetworkingV1PrivateLinkAttachment(ctx _context.Context, id string) ApiUpdateNetworkingV1PrivateLinkAttachmentRequest {
	return ApiUpdateNetworkingV1PrivateLinkAttachmentRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return NetworkingV1PrivateLinkAttachment
func (a *PrivateLinkAttachmentsNetworkingV1ApiService) UpdateNetworkingV1PrivateLinkAttachmentExecute(r ApiUpdateNetworkingV1PrivateLinkAttachmentRequest) (NetworkingV1PrivateLinkAttachment, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  NetworkingV1PrivateLinkAttachment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PrivateLinkAttachmentsNetworkingV1ApiService.UpdateNetworkingV1PrivateLinkAttachment")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/networking/v1/private-link-attachments/{id}"
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
	localVarPostBody = r.networkingV1PrivateLinkAttachmentUpdate
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
