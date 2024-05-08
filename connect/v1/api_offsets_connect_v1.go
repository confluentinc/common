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
Kafka Connect APIs

REST API for managing connectors

API version: 1.0
Contact: connect@confluent.io
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

type OffsetsConnectV1Api interface {

	/*
		AlterConnectv1ConnectorOffsetsRequest Request to Alter the Connector Offsets

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Request to alter the offsets of a connector. This supports the ability to PATCH/DELETE the offsets of a connector.
	Note, you will see momentary downtime as this will internally stop the connector, while the offsets are being altered.
	You can only make one alter offsets request at a time for a connector.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param connectorName The unique name of the connector.
		 @param environmentId The unique identifier of the environment this resource belongs to.
		 @param kafkaClusterId The unique identifier for the Kafka cluster.
		 @return ApiAlterConnectv1ConnectorOffsetsRequestRequest
	*/
	AlterConnectv1ConnectorOffsetsRequest(ctx _context.Context, connectorName string, environmentId string, kafkaClusterId string) ApiAlterConnectv1ConnectorOffsetsRequestRequest

	// AlterConnectv1ConnectorOffsetsRequestExecute executes the request
	//  @return ConnectV1AlterOffsetRequestInfo
	AlterConnectv1ConnectorOffsetsRequestExecute(r ApiAlterConnectv1ConnectorOffsetsRequestRequest) (ConnectV1AlterOffsetRequestInfo, *_nethttp.Response, error)

	/*
		GetConnectv1ConnectorOffsets Get a Connector Offsets

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Get the current offsets for the connector. The offsets provide information on the point in the source system,
	from which the connector is pulling in data. The offsets of a connector are continuously observed periodically and are queryable via this API.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param connectorName The unique name of the connector.
		 @param environmentId The unique identifier of the environment this resource belongs to.
		 @param kafkaClusterId The unique identifier for the Kafka cluster.
		 @return ApiGetConnectv1ConnectorOffsetsRequest
	*/
	GetConnectv1ConnectorOffsets(ctx _context.Context, connectorName string, environmentId string, kafkaClusterId string) ApiGetConnectv1ConnectorOffsetsRequest

	// GetConnectv1ConnectorOffsetsExecute executes the request
	//  @return ConnectV1ConnectorOffsets
	GetConnectv1ConnectorOffsetsExecute(r ApiGetConnectv1ConnectorOffsetsRequest) (ConnectV1ConnectorOffsets, *_nethttp.Response, error)

	/*
		GetConnectv1ConnectorOffsetsRequestStatus Get the Status of Altered Offset Request

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

	Get the status of the previous alter offset request.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param connectorName The unique name of the connector.
		 @param environmentId The unique identifier of the environment this resource belongs to.
		 @param kafkaClusterId The unique identifier for the Kafka cluster.
		 @return ApiGetConnectv1ConnectorOffsetsRequestStatusRequest
	*/
	GetConnectv1ConnectorOffsetsRequestStatus(ctx _context.Context, connectorName string, environmentId string, kafkaClusterId string) ApiGetConnectv1ConnectorOffsetsRequestStatusRequest

	// GetConnectv1ConnectorOffsetsRequestStatusExecute executes the request
	//  @return ConnectV1AlterOffsetStatus
	GetConnectv1ConnectorOffsetsRequestStatusExecute(r ApiGetConnectv1ConnectorOffsetsRequestStatusRequest) (ConnectV1AlterOffsetStatus, *_nethttp.Response, error)
}

// OffsetsConnectV1ApiService OffsetsConnectV1Api service
type OffsetsConnectV1ApiService service

type ApiAlterConnectv1ConnectorOffsetsRequestRequest struct {
	ctx                         _context.Context
	ApiService                  OffsetsConnectV1Api
	connectorName               string
	environmentId               string
	kafkaClusterId              string
	connectV1AlterOffsetRequest *ConnectV1AlterOffsetRequest
}

func (r ApiAlterConnectv1ConnectorOffsetsRequestRequest) ConnectV1AlterOffsetRequest(connectV1AlterOffsetRequest ConnectV1AlterOffsetRequest) ApiAlterConnectv1ConnectorOffsetsRequestRequest {
	r.connectV1AlterOffsetRequest = &connectV1AlterOffsetRequest
	return r
}

func (r ApiAlterConnectv1ConnectorOffsetsRequestRequest) Execute() (ConnectV1AlterOffsetRequestInfo, *_nethttp.Response, error) {
	return r.ApiService.AlterConnectv1ConnectorOffsetsRequestExecute(r)
}

/*
AlterConnectv1ConnectorOffsetsRequest Request to Alter the Connector Offsets

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Request to alter the offsets of a connector. This supports the ability to PATCH/DELETE the offsets of a connector.
Note, you will see momentary downtime as this will internally stop the connector, while the offsets are being altered.
You can only make one alter offsets request at a time for a connector.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param connectorName The unique name of the connector.
	@param environmentId The unique identifier of the environment this resource belongs to.
	@param kafkaClusterId The unique identifier for the Kafka cluster.
	@return ApiAlterConnectv1ConnectorOffsetsRequestRequest
*/
func (a *OffsetsConnectV1ApiService) AlterConnectv1ConnectorOffsetsRequest(ctx _context.Context, connectorName string, environmentId string, kafkaClusterId string) ApiAlterConnectv1ConnectorOffsetsRequestRequest {
	return ApiAlterConnectv1ConnectorOffsetsRequestRequest{
		ApiService:     a,
		ctx:            ctx,
		connectorName:  connectorName,
		environmentId:  environmentId,
		kafkaClusterId: kafkaClusterId,
	}
}

// Execute executes the request
//
//	@return ConnectV1AlterOffsetRequestInfo
func (a *OffsetsConnectV1ApiService) AlterConnectv1ConnectorOffsetsRequestExecute(r ApiAlterConnectv1ConnectorOffsetsRequestRequest) (ConnectV1AlterOffsetRequestInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectV1AlterOffsetRequestInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OffsetsConnectV1ApiService.AlterConnectv1ConnectorOffsetsRequest")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/offsets/request"
	localVarPath = strings.Replace(localVarPath, "{"+"connector_name"+"}", _neturl.PathEscape(parameterToString(r.connectorName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"kafka_cluster_id"+"}", _neturl.PathEscape(parameterToString(r.kafkaClusterId, "")), -1)

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
	localVarPostBody = r.connectV1AlterOffsetRequest
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
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ConnectV1ConnectorError
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

type ApiGetConnectv1ConnectorOffsetsRequest struct {
	ctx            _context.Context
	ApiService     OffsetsConnectV1Api
	connectorName  string
	environmentId  string
	kafkaClusterId string
}

func (r ApiGetConnectv1ConnectorOffsetsRequest) Execute() (ConnectV1ConnectorOffsets, *_nethttp.Response, error) {
	return r.ApiService.GetConnectv1ConnectorOffsetsExecute(r)
}

/*
GetConnectv1ConnectorOffsets Get a Connector Offsets

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Get the current offsets for the connector. The offsets provide information on the point in the source system,
from which the connector is pulling in data. The offsets of a connector are continuously observed periodically and are queryable via this API.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param connectorName The unique name of the connector.
	@param environmentId The unique identifier of the environment this resource belongs to.
	@param kafkaClusterId The unique identifier for the Kafka cluster.
	@return ApiGetConnectv1ConnectorOffsetsRequest
*/
func (a *OffsetsConnectV1ApiService) GetConnectv1ConnectorOffsets(ctx _context.Context, connectorName string, environmentId string, kafkaClusterId string) ApiGetConnectv1ConnectorOffsetsRequest {
	return ApiGetConnectv1ConnectorOffsetsRequest{
		ApiService:     a,
		ctx:            ctx,
		connectorName:  connectorName,
		environmentId:  environmentId,
		kafkaClusterId: kafkaClusterId,
	}
}

// Execute executes the request
//
//	@return ConnectV1ConnectorOffsets
func (a *OffsetsConnectV1ApiService) GetConnectv1ConnectorOffsetsExecute(r ApiGetConnectv1ConnectorOffsetsRequest) (ConnectV1ConnectorOffsets, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectV1ConnectorOffsets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OffsetsConnectV1ApiService.GetConnectv1ConnectorOffsets")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/offsets"
	localVarPath = strings.Replace(localVarPath, "{"+"connector_name"+"}", _neturl.PathEscape(parameterToString(r.connectorName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"kafka_cluster_id"+"}", _neturl.PathEscape(parameterToString(r.kafkaClusterId, "")), -1)

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
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ConnectV1ConnectorError
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

type ApiGetConnectv1ConnectorOffsetsRequestStatusRequest struct {
	ctx            _context.Context
	ApiService     OffsetsConnectV1Api
	connectorName  string
	environmentId  string
	kafkaClusterId string
}

func (r ApiGetConnectv1ConnectorOffsetsRequestStatusRequest) Execute() (ConnectV1AlterOffsetStatus, *_nethttp.Response, error) {
	return r.ApiService.GetConnectv1ConnectorOffsetsRequestStatusExecute(r)
}

/*
GetConnectv1ConnectorOffsetsRequestStatus Get the Status of Altered Offset Request

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)

Get the status of the previous alter offset request.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param connectorName The unique name of the connector.
	@param environmentId The unique identifier of the environment this resource belongs to.
	@param kafkaClusterId The unique identifier for the Kafka cluster.
	@return ApiGetConnectv1ConnectorOffsetsRequestStatusRequest
*/
func (a *OffsetsConnectV1ApiService) GetConnectv1ConnectorOffsetsRequestStatus(ctx _context.Context, connectorName string, environmentId string, kafkaClusterId string) ApiGetConnectv1ConnectorOffsetsRequestStatusRequest {
	return ApiGetConnectv1ConnectorOffsetsRequestStatusRequest{
		ApiService:     a,
		ctx:            ctx,
		connectorName:  connectorName,
		environmentId:  environmentId,
		kafkaClusterId: kafkaClusterId,
	}
}

// Execute executes the request
//
//	@return ConnectV1AlterOffsetStatus
func (a *OffsetsConnectV1ApiService) GetConnectv1ConnectorOffsetsRequestStatusExecute(r ApiGetConnectv1ConnectorOffsetsRequestStatusRequest) (ConnectV1AlterOffsetStatus, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ConnectV1AlterOffsetStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OffsetsConnectV1ApiService.GetConnectv1ConnectorOffsetsRequestStatus")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connect/v1/environments/{environment_id}/clusters/{kafka_cluster_id}/connectors/{connector_name}/offsets/request/status"
	localVarPath = strings.Replace(localVarPath, "{"+"connector_name"+"}", _neturl.PathEscape(parameterToString(r.connectorName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"kafka_cluster_id"+"}", _neturl.PathEscape(parameterToString(r.kafkaClusterId, "")), -1)

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
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ConnectV1ConnectorError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ConnectV1ConnectorError
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
