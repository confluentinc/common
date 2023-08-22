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
SQL API v1beta1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1beta1

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

type StatementExceptionsSqlV1beta1Api interface {

	/*
			GetSqlv1beta1StatementExceptions List of Statement Exceptions

			[![Open Preview](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To SQL API v1beta1](https://img.shields.io/badge/-Request%20Access%20To%20Flink%20Gateway%20API-%23bc8540)](mailto:ccloud-api-access+sql-v1beta1-early-access@confluent.io?subject=Request%20to%20join%20sql/v1beta1%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20sql/v1beta1%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)
		Retrieve a list of the 10 most recent statement exceptions.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param organizationId The unique identifier for the organization.
			 @param environmentId The unique identifier for the environment.
			 @param statementName The unique identifier for the statement.
			 @return ApiGetSqlv1beta1StatementExceptionsRequest
	*/
	GetSqlv1beta1StatementExceptions(ctx _context.Context, organizationId string, environmentId string, statementName string) ApiGetSqlv1beta1StatementExceptionsRequest

	// GetSqlv1beta1StatementExceptionsExecute executes the request
	//  @return SqlV1beta1StatementExceptionList
	GetSqlv1beta1StatementExceptionsExecute(r ApiGetSqlv1beta1StatementExceptionsRequest) (SqlV1beta1StatementExceptionList, *_nethttp.Response, error)
}

// StatementExceptionsSqlV1beta1ApiService StatementExceptionsSqlV1beta1Api service
type StatementExceptionsSqlV1beta1ApiService service

type ApiGetSqlv1beta1StatementExceptionsRequest struct {
	ctx            _context.Context
	ApiService     StatementExceptionsSqlV1beta1Api
	organizationId string
	environmentId  string
	statementName  string
}

func (r ApiGetSqlv1beta1StatementExceptionsRequest) Execute() (SqlV1beta1StatementExceptionList, *_nethttp.Response, error) {
	return r.ApiService.GetSqlv1beta1StatementExceptionsExecute(r)
}

/*
GetSqlv1beta1StatementExceptions List of Statement Exceptions

[![Open Preview](https://img.shields.io/badge/Lifecycle%20Stage-Early%20Access-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy) [![Request Access To SQL API v1beta1](https://img.shields.io/badge/-Request%20Access%20To%20Flink%20Gateway%20API-%23bc8540)](mailto:ccloud-api-access+sql-v1beta1-early-access@confluent.io?subject=Request%20to%20join%20sql/v1beta1%20API%20Early%20Access&body=I%E2%80%99d%20like%20to%20join%20the%20Confluent%20Cloud%20API%20Early%20Access%20for%20sql/v1beta1%20to%20provide%20early%20feedback%21%20My%20Cloud%20Organization%20ID%20is%20%3Cretrieve%20from%20https%3A//confluent.cloud/settings/billing/payment%3E.)
Retrieve a list of the 10 most recent statement exceptions.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique identifier for the organization.
	@param environmentId The unique identifier for the environment.
	@param statementName The unique identifier for the statement.
	@return ApiGetSqlv1beta1StatementExceptionsRequest
*/
func (a *StatementExceptionsSqlV1beta1ApiService) GetSqlv1beta1StatementExceptions(ctx _context.Context, organizationId string, environmentId string, statementName string) ApiGetSqlv1beta1StatementExceptionsRequest {
	return ApiGetSqlv1beta1StatementExceptionsRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		environmentId:  environmentId,
		statementName:  statementName,
	}
}

// Execute executes the request
//
//	@return SqlV1beta1StatementExceptionList
func (a *StatementExceptionsSqlV1beta1ApiService) GetSqlv1beta1StatementExceptionsExecute(r ApiGetSqlv1beta1StatementExceptionsRequest) (SqlV1beta1StatementExceptionList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SqlV1beta1StatementExceptionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatementExceptionsSqlV1beta1ApiService.GetSqlv1beta1StatementExceptions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1beta1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name}/exceptions"
	localVarPath = strings.Replace(localVarPath, "{"+"organization_id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"statement_name"+"}", _neturl.PathEscape(parameterToString(r.statementName, "")), -1)

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
