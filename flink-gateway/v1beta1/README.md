# Go API client for v1beta1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://confluent.slack.com/app_redirect?channel=flink-ccloud](https://confluent.slack.com/app_redirect?channel=flink-ccloud)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./v1beta1"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://flink.region.provider.confluent.cloud*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*StatementExceptionsSqlV1beta1Api* | [**GetSqlv1beta1StatementExceptions**](docs/StatementExceptionsSqlV1beta1Api.md#getsqlv1beta1statementexceptions) | **Get** /sql/v1beta1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name}/exceptions | List of Statement Exceptions
*StatementResultSqlV1beta1Api* | [**GetSqlv1beta1StatementResult**](docs/StatementResultSqlV1beta1Api.md#getsqlv1beta1statementresult) | **Get** /sql/v1beta1/organizations/{organization_id}/environments/{environment_id}/statements/{name}/results | Read Statement Result
*StatementsSqlV1beta1Api* | [**CreateSqlv1beta1Statement**](docs/StatementsSqlV1beta1Api.md#createsqlv1beta1statement) | **Post** /sql/v1beta1/organizations/{organization_id}/environments/{environment_id}/statements | Create a Statement
*StatementsSqlV1beta1Api* | [**DeleteSqlv1beta1Statement**](docs/StatementsSqlV1beta1Api.md#deletesqlv1beta1statement) | **Delete** /sql/v1beta1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name} | Delete a Statement
*StatementsSqlV1beta1Api* | [**GetSqlv1beta1Statement**](docs/StatementsSqlV1beta1Api.md#getsqlv1beta1statement) | **Get** /sql/v1beta1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name} | Read a Statement
*StatementsSqlV1beta1Api* | [**ListSqlv1beta1Statements**](docs/StatementsSqlV1beta1Api.md#listsqlv1beta1statements) | **Get** /sql/v1beta1/organizations/{organization_id}/environments/{environment_id}/statements | List of Statements
*StatementsSqlV1beta1Api* | [**UpdateSqlv1beta1Statement**](docs/StatementsSqlV1beta1Api.md#updatesqlv1beta1statement) | **Put** /sql/v1beta1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name} | Update a Statement


## Documentation For Models

 - [ColumnDetails](docs/ColumnDetails.md)
 - [DataType](docs/DataType.md)
 - [Error](docs/Error.md)
 - [ErrorSource](docs/ErrorSource.md)
 - [ExceptionListMeta](docs/ExceptionListMeta.md)
 - [Failure](docs/Failure.md)
 - [ListMeta](docs/ListMeta.md)
 - [MultipleSearchFilter](docs/MultipleSearchFilter.md)
 - [ObjectMeta](docs/ObjectMeta.md)
 - [ResultListMeta](docs/ResultListMeta.md)
 - [RowFieldType](docs/RowFieldType.md)
 - [SqlV1beta1ResultSchema](docs/SqlV1beta1ResultSchema.md)
 - [SqlV1beta1Statement](docs/SqlV1beta1Statement.md)
 - [SqlV1beta1StatementException](docs/SqlV1beta1StatementException.md)
 - [SqlV1beta1StatementExceptionList](docs/SqlV1beta1StatementExceptionList.md)
 - [SqlV1beta1StatementList](docs/SqlV1beta1StatementList.md)
 - [SqlV1beta1StatementResult](docs/SqlV1beta1StatementResult.md)
 - [SqlV1beta1StatementResultResults](docs/SqlV1beta1StatementResultResults.md)
 - [SqlV1beta1StatementSpec](docs/SqlV1beta1StatementSpec.md)
 - [SqlV1beta1StatementStatus](docs/SqlV1beta1StatementStatus.md)


## Documentation For Authorization



### resource-api-key

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

flink-control-plane@confluent.io
