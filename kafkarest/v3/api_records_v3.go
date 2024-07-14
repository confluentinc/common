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
REST Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.0
Contact: kafka-clients-proxy-team@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v3

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

type RecordsV3Api interface {

	/*
			ProduceRecord Produce Records

			[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

		Produce records to the given topic, returning delivery reports for each
		record produced. This API can be used in streaming mode by setting
		"Transfer-Encoding: chunked" header. For as long as the connection is
		kept open, the server will keep accepting records. Records are streamed
		to and from the server as Concatenated JSON. For each record sent to the
		server, the server will asynchronously send back a delivery report, in
		the same order, each with its own error_code. An error_code of 200
		indicates success. The HTTP status code will be HTTP 200 OK as long as
		the connection is successfully established. To identify records that
		have encountered an error, check the error_code of each delivery report.

		This API currently does not support Schema Registry integration. Sending
		schemas is not supported. Only BINARY, JSON, and STRING formats are
		supported.

			 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			 @param clusterId The Kafka cluster ID.
			 @param topicName The topic name.
			 @return ApiProduceRecordRequest
	*/
	ProduceRecord(ctx _context.Context, clusterId string, topicName string) ApiProduceRecordRequest

	// ProduceRecordExecute executes the request
	//  @return ProduceResponse
	ProduceRecordExecute(r ApiProduceRecordRequest) (ProduceResponse, *_nethttp.Response, error)
}

// RecordsV3ApiService RecordsV3Api service
type RecordsV3ApiService service

type ApiProduceRecordRequest struct {
	ctx            _context.Context
	ApiService     RecordsV3Api
	clusterId      string
	topicName      string
	produceRequest *ProduceRequest
}

// A single record to be produced to Kafka. To produce multiple records in the same request, simply concatenate the records. The delivery reports are concatenated in the same order as the records are sent.
func (r ApiProduceRecordRequest) ProduceRequest(produceRequest ProduceRequest) ApiProduceRecordRequest {
	r.produceRequest = &produceRequest
	return r
}

func (r ApiProduceRecordRequest) Execute() (ProduceResponse, *_nethttp.Response, error) {
	return r.ApiService.ProduceRecordExecute(r)
}

/*
ProduceRecord Produce Records

[![Generally Available](https://img.shields.io/badge/Lifecycle%20Stage-Generally%20Available-%2345c6e8)](#section/Versioning/API-Lifecycle-Policy)

Produce records to the given topic, returning delivery reports for each
record produced. This API can be used in streaming mode by setting
"Transfer-Encoding: chunked" header. For as long as the connection is
kept open, the server will keep accepting records. Records are streamed
to and from the server as Concatenated JSON. For each record sent to the
server, the server will asynchronously send back a delivery report, in
the same order, each with its own error_code. An error_code of 200
indicates success. The HTTP status code will be HTTP 200 OK as long as
the connection is successfully established. To identify records that
have encountered an error, check the error_code of each delivery report.

This API currently does not support Schema Registry integration. Sending
schemas is not supported. Only BINARY, JSON, and STRING formats are
supported.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterId The Kafka cluster ID.
 @param topicName The topic name.
 @return ApiProduceRecordRequest
*/
func (a *RecordsV3ApiService) ProduceRecord(ctx _context.Context, clusterId string, topicName string) ApiProduceRecordRequest {
	return ApiProduceRecordRequest{
		ApiService: a,
		ctx:        ctx,
		clusterId:  clusterId,
		topicName:  topicName,
	}
}

// Execute executes the request
//  @return ProduceResponse
func (a *RecordsV3ApiService) ProduceRecordExecute(r ApiProduceRecordRequest) (ProduceResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ProduceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RecordsV3ApiService.ProduceRecord")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kafka/v3/clusters/{cluster_id}/topics/{topic_name}/records"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", _neturl.PathEscape(parameterToString(r.clusterId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"topic_name"+"}", _neturl.PathEscape(parameterToString(r.topicName, "")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.produceRequest
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
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 413 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 415 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode >= 500 {
			var v Error
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