# \AuthenticationAPIAPI

All URIs are relative to *http://http:*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](AuthenticationAPIAPI.md#CreateToken) | **Post** /auth/tokens | Provide a refresh token and get back a JWT for authenticating to a Galasa ecosystem.
[**DeleteToken**](AuthenticationAPIAPI.md#DeleteToken) | **Delete** /auth/tokens/{tokenId} | Revokes a token which is used for authenticating with the Galasa API server.
[**GetAuthenticate**](AuthenticationAPIAPI.md#GetAuthenticate) | **Get** /auth | Redirect to authenticate with an upstream OpenID Connect connector.
[**GetAuthenticateCallback**](AuthenticationAPIAPI.md#GetAuthenticateCallback) | **Get** /auth/callback | Callback endpoint used as part of an OAuth 2.0 authorization code flow, redirects to a callback URL with an authorization code as a query parameter.
[**GetTokens**](AuthenticationAPIAPI.md#GetTokens) | **Get** /auth/tokens | Get a list of tokens used for authenticating with the Galasa API server.
[**PostAuthenticate**](AuthenticationAPIAPI.md#PostAuthenticate) | **Post** /auth | Provide a refresh token and get back a JWT for authenticating to a Galasa ecosystem.
[**PostClients**](AuthenticationAPIAPI.md#PostClients) | **Post** /auth/clients | Create a new Dex client to authenticate with.



## CreateToken

> TokenResponse CreateToken(ctx).AuthProperties(authProperties).ClientApiVersion(clientApiVersion).Execute()

Provide a refresh token and get back a JWT for authenticating to a Galasa ecosystem.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authProperties := *openapiclient.NewAuthProperties() // AuthProperties | 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPIAPI.CreateToken(context.Background()).AuthProperties(authProperties).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIAPI.CreateToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateToken`: TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPIAPI.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authProperties** | [**AuthProperties**](AuthProperties.md) |  | 
 **clientApiVersion** | **string** |  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToken

> string DeleteToken(ctx, tokenId).ClientApiVersion(clientApiVersion).Execute()

Revokes a token which is used for authenticating with the Galasa API server.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	tokenId := "tokenId_example" // string | Token ID. An alphanumeric string used to identify a token. 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPIAPI.DeleteToken(context.Background(), tokenId).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIAPI.DeleteToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteToken`: string
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPIAPI.DeleteToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenId** | **string** | Token ID. An alphanumeric string used to identify a token.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientApiVersion** | **string** |  | 

### Return type

**string**

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticate

> GetAuthenticate(ctx).ClientId(clientId).CallbackUrl(callbackUrl).ClientApiVersion(clientApiVersion).Execute()

Redirect to authenticate with an upstream OpenID Connect connector.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientId := "clientId_example" // string | The ID of the client used to authenticate with the upstream connector.
	callbackUrl := "callbackUrl_example" // string | The URL to return to once the authentication process is complete. This should be a URL to a client web application. 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPIAPI.GetAuthenticate(context.Background()).ClientId(clientId).CallbackUrl(callbackUrl).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIAPI.GetAuthenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | The ID of the client used to authenticate with the upstream connector. | 
 **callbackUrl** | **string** | The URL to return to once the authentication process is complete. This should be a URL to a client web application.  | 
 **clientApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAuthenticateCallback

> GetAuthenticateCallback(ctx).Code(code).State(state).ClientApiVersion(clientApiVersion).Execute()

Callback endpoint used as part of an OAuth 2.0 authorization code flow, redirects to a callback URL with an authorization code as a query parameter.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	code := "code_example" // string | The authorization code for the current authentication request.
	state := "state_example" // string | The state parameter associated with the current authentication request, used to mitigate CSRF attacks.
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AuthenticationAPIAPI.GetAuthenticateCallback(context.Background()).Code(code).State(state).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIAPI.GetAuthenticateCallback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuthenticateCallbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | The authorization code for the current authentication request. | 
 **state** | **string** | The state parameter associated with the current authentication request, used to mitigate CSRF attacks. | 
 **clientApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokens

> AuthTokens GetTokens(ctx).ClientApiVersion(clientApiVersion).LoginId(loginId).Execute()

Get a list of tokens used for authenticating with the Galasa API server.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientApiVersion := "clientApiVersion_example" // string |  (optional)
	loginId := "loginId_example" // string | The loginId of the user whose details will be returned.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPIAPI.GetTokens(context.Background()).ClientApiVersion(clientApiVersion).LoginId(loginId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIAPI.GetTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTokens`: AuthTokens
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPIAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientApiVersion** | **string** |  | 
 **loginId** | **string** | The loginId of the user whose details will be returned.  | 

### Return type

[**AuthTokens**](AuthTokens.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAuthenticate

> TokenResponse PostAuthenticate(ctx).AuthProperties(authProperties).ClientApiVersion(clientApiVersion).Execute()

Provide a refresh token and get back a JWT for authenticating to a Galasa ecosystem.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	authProperties := *openapiclient.NewAuthProperties() // AuthProperties | 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPIAPI.PostAuthenticate(context.Background()).AuthProperties(authProperties).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIAPI.PostAuthenticate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAuthenticate`: TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPIAPI.PostAuthenticate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAuthenticateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authProperties** | [**AuthProperties**](AuthProperties.md) |  | 
 **clientApiVersion** | **string** |  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostClients

> DexClient PostClients(ctx).ClientApiVersion(clientApiVersion).Execute()

Create a new Dex client to authenticate with.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuthenticationAPIAPI.PostClients(context.Background()).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPIAPI.PostClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostClients`: DexClient
	fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPIAPI.PostClients`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientApiVersion** | **string** |  | 

### Return type

[**DexClient**](DexClient.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

