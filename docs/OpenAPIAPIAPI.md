# \OpenAPIAPIAPI

All URIs are relative to *http://http:*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOpenApiSpec**](OpenAPIAPIAPI.md#GetOpenApiSpec) | **Get** /openapi | Retrieve the API server&#39;s OpenAPI specification



## GetOpenApiSpec

> map[string]interface{} GetOpenApiSpec(ctx).ClientApiVersion(clientApiVersion).Execute()

Retrieve the API server's OpenAPI specification



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
	resp, r, err := apiClient.OpenAPIAPIAPI.GetOpenApiSpec(context.Background()).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenAPIAPIAPI.GetOpenApiSpec``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenApiSpec`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `OpenAPIAPIAPI.GetOpenApiSpec`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenApiSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientApiVersion** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/yaml, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

