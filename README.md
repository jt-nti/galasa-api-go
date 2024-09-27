# Go API client for pkg/galasaapi

The Galasa Ecosystem REST API allows you to interact with a Galasa Ecosystem.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.37.0
- Package version: 1.0.0
- Generator version: 7.8.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://galasa.dev/support](https://galasa.dev/support)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import pkg/galasaapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `pkg/galasaapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), pkg/galasaapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `pkg/galasaapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), pkg/galasaapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `pkg/galasaapi.ContextOperationServerIndices` and `pkg/galasaapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), pkg/galasaapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), pkg/galasaapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://http:*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthenticationAPIAPI* | [**CreateToken**](docs/AuthenticationAPIAPI.md#createtoken) | **Post** /auth/tokens | Provide a refresh token and get back a JWT for authenticating to a Galasa ecosystem.
*AuthenticationAPIAPI* | [**DeleteToken**](docs/AuthenticationAPIAPI.md#deletetoken) | **Delete** /auth/tokens/{tokenId} | Revokes a token which is used for authenticating with the Galasa API server.
*AuthenticationAPIAPI* | [**GetAuthenticate**](docs/AuthenticationAPIAPI.md#getauthenticate) | **Get** /auth | Redirect to authenticate with an upstream OpenID Connect connector.
*AuthenticationAPIAPI* | [**GetAuthenticateCallback**](docs/AuthenticationAPIAPI.md#getauthenticatecallback) | **Get** /auth/callback | Callback endpoint used as part of an OAuth 2.0 authorization code flow, redirects to a callback URL with an authorization code as a query parameter.
*AuthenticationAPIAPI* | [**GetTokens**](docs/AuthenticationAPIAPI.md#gettokens) | **Get** /auth/tokens | Get a list of tokens used for authenticating with the Galasa API server.
*AuthenticationAPIAPI* | [**PostAuthenticate**](docs/AuthenticationAPIAPI.md#postauthenticate) | **Post** /auth | Provide a refresh token and get back a JWT for authenticating to a Galasa ecosystem.
*AuthenticationAPIAPI* | [**PostClients**](docs/AuthenticationAPIAPI.md#postclients) | **Post** /auth/clients | Create a new Dex client to authenticate with.
*BootstrapAPIAPI* | [**GetEcosystemBootstrap**](docs/BootstrapAPIAPI.md#getecosystembootstrap) | **Get** /bootstrap | Contact the Galasa ecosystem
*ConfigurationPropertyStoreAPIAPI* | [**CreateCpsProperty**](docs/ConfigurationPropertyStoreAPIAPI.md#createcpsproperty) | **Post** /cps/{namespace}/properties | Create a new CPS property
*ConfigurationPropertyStoreAPIAPI* | [**DeleteCpsProperty**](docs/ConfigurationPropertyStoreAPIAPI.md#deletecpsproperty) | **Delete** /cps/{namespace}/properties/{propertyName} | Delete existing CPS property
*ConfigurationPropertyStoreAPIAPI* | [**GetAllCpsNamespaces**](docs/ConfigurationPropertyStoreAPIAPI.md#getallcpsnamespaces) | **Get** /cps | Get all known CPS namespaces
*ConfigurationPropertyStoreAPIAPI* | [**GetCpsNamespaceCascadeProperty**](docs/ConfigurationPropertyStoreAPIAPI.md#getcpsnamespacecascadeproperty) | **Get** /cps/namespace/{namespace}/prefix/{prefix}/suffix/{suffix} | Get cascade CPS property
*ConfigurationPropertyStoreAPIAPI* | [**GetCpsNamespaceProperties**](docs/ConfigurationPropertyStoreAPIAPI.md#getcpsnamespaceproperties) | **Get** /cps/namespace/{namespace} | Get all properties for a namepace
*ConfigurationPropertyStoreAPIAPI* | [**GetCpsNamespaces**](docs/ConfigurationPropertyStoreAPIAPI.md#getcpsnamespaces) | **Get** /cps/namespace | Get CPS Namespaces
*ConfigurationPropertyStoreAPIAPI* | [**GetCpsProperty**](docs/ConfigurationPropertyStoreAPIAPI.md#getcpsproperty) | **Get** /cps/{namespace}/properties/{propertyName} | Get single CPS property
*ConfigurationPropertyStoreAPIAPI* | [**PutCpsNamespaceProperty**](docs/ConfigurationPropertyStoreAPIAPI.md#putcpsnamespaceproperty) | **Put** /cps/namespace/{namespace}/property/{property} | Put new CPS Property
*ConfigurationPropertyStoreAPIAPI* | [**QueryCpsNamespaceProperties**](docs/ConfigurationPropertyStoreAPIAPI.md#querycpsnamespaceproperties) | **Get** /cps/{namespace}/properties | Get all CPS namespace properties
*ConfigurationPropertyStoreAPIAPI* | [**UpdateCpsProperty**](docs/ConfigurationPropertyStoreAPIAPI.md#updatecpsproperty) | **Put** /cps/{namespace}/properties/{propertyName} | Update existing CPS property
*OpenAPIAPIAPI* | [**GetOpenApiSpec**](docs/OpenAPIAPIAPI.md#getopenapispec) | **Get** /openapi | Retrieve the API server&#39;s OpenAPI specification
*ResourcesAPIAPI* | [**SetEcosystemResources**](docs/ResourcesAPIAPI.md#setecosystemresources) | **Post** /resources/ | Upload Resources to the ecosystem
*ResultArchiveStoreAPIAPI* | [**DeleteRasRunById**](docs/ResultArchiveStoreAPIAPI.md#deleterasrunbyid) | **Delete** /ras/runs/{runid} | Delete run from the RAS using a given run ID
*ResultArchiveStoreAPIAPI* | [**GetRasRequestors**](docs/ResultArchiveStoreAPIAPI.md#getrasrequestors) | **Get** /ras/requestors | Get all known requestors
*ResultArchiveStoreAPIAPI* | [**GetRasResultNames**](docs/ResultArchiveStoreAPIAPI.md#getrasresultnames) | **Get** /ras/resultnames | Get all the known result names
*ResultArchiveStoreAPIAPI* | [**GetRasRunArtifactByPath**](docs/ResultArchiveStoreAPIAPI.md#getrasrunartifactbypath) | **Get** /ras/runs/{runid}/files/{artifactPath} | Download artifact for a given runid by artifactPath
*ResultArchiveStoreAPIAPI* | [**GetRasRunArtifactList**](docs/ResultArchiveStoreAPIAPI.md#getrasrunartifactlist) | **Get** /ras/runs/{runid}/artifacts | Get the available run artifacts which can be downloaded.
*ResultArchiveStoreAPIAPI* | [**GetRasRunById**](docs/ResultArchiveStoreAPIAPI.md#getrasrunbyid) | **Get** /ras/runs/{runid} | Get Run by ID
*ResultArchiveStoreAPIAPI* | [**GetRasRunLog**](docs/ResultArchiveStoreAPIAPI.md#getrasrunlog) | **Get** /ras/runs/{runid}/runlog | Get Run Log
*ResultArchiveStoreAPIAPI* | [**GetRasSearchRuns**](docs/ResultArchiveStoreAPIAPI.md#getrassearchruns) | **Get** /ras/runs | Get Runs from Query
*ResultArchiveStoreAPIAPI* | [**GetRasTestclasses**](docs/ResultArchiveStoreAPIAPI.md#getrastestclasses) | **Get** /ras/testclasses | Get all the known test classes
*ResultArchiveStoreAPIAPI* | [**PutRasRunStatusById**](docs/ResultArchiveStoreAPIAPI.md#putrasrunstatusbyid) | **Put** /ras/runs/{runid} | Update the status of a test run
*RunsAPIAPI* | [**GetRunsGroup**](docs/RunsAPIAPI.md#getrunsgroup) | **Get** /runs/{groupId} | Get group runs
*RunsAPIAPI* | [**PostSubmitTestRuns**](docs/RunsAPIAPI.md#postsubmittestruns) | **Post** /runs/{groupId} | Submit test runs
*UsersAPIAPI* | [**GetUserByLoginId**](docs/UsersAPIAPI.md#getuserbyloginid) | **Get** /users | Returns the details of the requesting user.


## Documentation For Models

 - [APIError](docs/APIError.md)
 - [Artifact](docs/Artifact.md)
 - [ArtifactIndexEntry](docs/ArtifactIndexEntry.md)
 - [AuthProperties](docs/AuthProperties.md)
 - [AuthToken](docs/AuthToken.md)
 - [AuthTokens](docs/AuthTokens.md)
 - [CpsProperty](docs/CpsProperty.md)
 - [DexClient](docs/DexClient.md)
 - [GalasaProperty](docs/GalasaProperty.md)
 - [GalasaPropertyData](docs/GalasaPropertyData.md)
 - [GalasaPropertyMetadata](docs/GalasaPropertyMetadata.md)
 - [Namespace](docs/Namespace.md)
 - [Requestors](docs/Requestors.md)
 - [Resources](docs/Resources.md)
 - [ResourcesDataInner](docs/ResourcesDataInner.md)
 - [ResultNames](docs/ResultNames.md)
 - [Run](docs/Run.md)
 - [RunResults](docs/RunResults.md)
 - [TestClass](docs/TestClass.md)
 - [TestClasses](docs/TestClasses.md)
 - [TestMethod](docs/TestMethod.md)
 - [TestRun](docs/TestRun.md)
 - [TestRunRequest](docs/TestRunRequest.md)
 - [TestRuns](docs/TestRuns.md)
 - [TestStructure](docs/TestStructure.md)
 - [TokenResponse](docs/TokenResponse.md)
 - [UpdateRunStatusRequest](docs/UpdateRunStatusRequest.md)
 - [User](docs/User.md)
 - [UserData](docs/UserData.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### JwtAuth

- **Type**: HTTP Bearer token authentication

Example

```go
auth := context.WithValue(context.Background(), pkg/galasaapi.ContextAccessToken, "BEARER_TOKEN_STRING")
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



