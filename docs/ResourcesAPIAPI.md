# \ResourcesAPIAPI

All URIs are relative to *http://http:*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SetEcosystemResources**](ResourcesAPIAPI.md#SetEcosystemResources) | **Post** /resources/ | Upload Resources to the ecosystem



## SetEcosystemResources

> SetEcosystemResources(ctx).Resources(resources).ClientApiVersion(clientApiVersion).Execute()

Upload Resources to the ecosystem



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
	resources := *openapiclient.NewResources() // Resources | 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourcesAPIAPI.SetEcosystemResources(context.Background()).Resources(resources).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourcesAPIAPI.SetEcosystemResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetEcosystemResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resources** | [**Resources**](Resources.md) |  | 
 **clientApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

