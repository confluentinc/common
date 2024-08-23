# Go API client for v1

This is the Flink Artifact Management API.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://confluent.slack.com/app_redirect?channel=flink-runtime-eng](https://confluent.slack.com/app_redirect?channel=flink-runtime-eng)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./v1"
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

All URIs are relative to *https://api.confluent.cloud*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FlinkArtifactsArtifactV1Api* | [**CreateArtifactV1FlinkArtifact**](docs/FlinkArtifactsArtifactV1Api.md#createartifactv1flinkartifact) | **Post** /artifact/v1/flink-artifacts | Create a new Flink Artifact.
*FlinkArtifactsArtifactV1Api* | [**DeleteArtifactV1FlinkArtifact**](docs/FlinkArtifactsArtifactV1Api.md#deleteartifactv1flinkartifact) | **Delete** /artifact/v1/flink-artifacts/{id} | Delete a Flink Artifact
*FlinkArtifactsArtifactV1Api* | [**GetArtifactV1FlinkArtifact**](docs/FlinkArtifactsArtifactV1Api.md#getartifactv1flinkartifact) | **Get** /artifact/v1/flink-artifacts/{id} | Read a Flink Artifact
*FlinkArtifactsArtifactV1Api* | [**ListArtifactV1FlinkArtifacts**](docs/FlinkArtifactsArtifactV1Api.md#listartifactv1flinkartifacts) | **Get** /artifact/v1/flink-artifacts | List of Flink Artifacts
*FlinkArtifactsArtifactV1Api* | [**UpdateArtifactV1FlinkArtifact**](docs/FlinkArtifactsArtifactV1Api.md#updateartifactv1flinkartifact) | **Patch** /artifact/v1/flink-artifacts/{id} | Update a Flink Artifact
*PresignedUrlsArtifactV1Api* | [**PresignedUploadUrlArtifactV1PresignedUrl**](docs/PresignedUrlsArtifactV1Api.md#presigneduploadurlartifactv1presignedurl) | **Post** /artifact/v1/presigned-upload-url | Request a presigned upload URL for a new Flink Artifact.


## Documentation For Models

 - [ArtifactV1FlinkArtifact](docs/ArtifactV1FlinkArtifact.md)
 - [ArtifactV1FlinkArtifactList](docs/ArtifactV1FlinkArtifactList.md)
 - [ArtifactV1FlinkArtifactUpdate](docs/ArtifactV1FlinkArtifactUpdate.md)
 - [ArtifactV1FlinkArtifactVersion](docs/ArtifactV1FlinkArtifactVersion.md)
 - [ArtifactV1FlinkArtifactVersionUploadSourceOneOf](docs/ArtifactV1FlinkArtifactVersionUploadSourceOneOf.md)
 - [ArtifactV1PresignedUrl](docs/ArtifactV1PresignedUrl.md)
 - [ArtifactV1PresignedUrlRequest](docs/ArtifactV1PresignedUrlRequest.md)
 - [ArtifactV1UploadSourcePresignedUrl](docs/ArtifactV1UploadSourcePresignedUrl.md)
 - [Error](docs/Error.md)
 - [ErrorSource](docs/ErrorSource.md)
 - [Failure](docs/Failure.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObjectUploadSourceOneOf](docs/InlineObjectUploadSourceOneOf.md)
 - [ListMeta](docs/ListMeta.md)
 - [ObjectMeta](docs/ObjectMeta.md)
 - [ObjectReference](docs/ObjectReference.md)


## Documentation For Authorization



### cloud-api-key

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


### confluent-sts-access-token


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
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

flink-runtime@confluent.io
