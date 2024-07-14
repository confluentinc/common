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
AI API

API for interacting with AI models from within Confluent Cloud.

API version: 0.0.1
Contact: api-team@confluent.io
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

type DocCompletionsDocsAiV1Api interface {

	/*
		QueryDocsAiV1DocCompletion Query a Doc Completion

		[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To AI API v1](https://img.shields.io/badge/-Request%20Access%20To%20AI%20API%20v1-%23bc8540)](mailto:ccloud-api-access+docs-ai-v1-early-access@confluent.io?subject=Request%20to%20join%20docs-ai/v1%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20docs-ai/v1%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)

	Query the Docs AI assistant, optionally with prior conversation history.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @return ApiQueryDocsAiV1DocCompletionRequest
	*/
	QueryDocsAiV1DocCompletion(ctx _context.Context) ApiQueryDocsAiV1DocCompletionRequest

	// QueryDocsAiV1DocCompletionExecute executes the request
	//  @return AiV1ChatCompletionsReply
	QueryDocsAiV1DocCompletionExecute(r ApiQueryDocsAiV1DocCompletionRequest) (AiV1ChatCompletionsReply, *_nethttp.Response, error)
}

// DocCompletionsDocsAiV1ApiService DocCompletionsDocsAiV1Api service
type DocCompletionsDocsAiV1ApiService service

type ApiQueryDocsAiV1DocCompletionRequest struct {
	ctx                        _context.Context
	ApiService                 DocCompletionsDocsAiV1Api
	aiV1ChatCompletionsRequest *AiV1ChatCompletionsRequest
}

func (r ApiQueryDocsAiV1DocCompletionRequest) AiV1ChatCompletionsRequest(aiV1ChatCompletionsRequest AiV1ChatCompletionsRequest) ApiQueryDocsAiV1DocCompletionRequest {
	r.aiV1ChatCompletionsRequest = &aiV1ChatCompletionsRequest
	return r
}

func (r ApiQueryDocsAiV1DocCompletionRequest) Execute() (AiV1ChatCompletionsReply, *_nethttp.Response, error) {
	return r.ApiService.QueryDocsAiV1DocCompletionExecute(r)
}

/*
QueryDocsAiV1DocCompletion Query a Doc Completion

[![Early Access](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To AI API v1](https://img.shields.io/badge/-Request%20Access%20To%20AI%20API%20v1-%23bc8540)](mailto:ccloud-api-access+docs-ai-v1-early-access@confluent.io?subject=Request%20to%20join%20docs-ai/v1%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20docs-ai/v1%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)

Query the Docs AI assistant, optionally with prior conversation history.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryDocsAiV1DocCompletionRequest
*/
func (a *DocCompletionsDocsAiV1ApiService) QueryDocsAiV1DocCompletion(ctx _context.Context) ApiQueryDocsAiV1DocCompletionRequest {
	return ApiQueryDocsAiV1DocCompletionRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return AiV1ChatCompletionsReply
func (a *DocCompletionsDocsAiV1ApiService) QueryDocsAiV1DocCompletionExecute(r ApiQueryDocsAiV1DocCompletionRequest) (AiV1ChatCompletionsReply, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AiV1ChatCompletionsReply
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DocCompletionsDocsAiV1ApiService.QueryDocsAiV1DocCompletion")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/docs-ai/v1/doc-completions"

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
	localVarPostBody = r.aiV1ChatCompletionsRequest
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