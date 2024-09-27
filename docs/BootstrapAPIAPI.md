# \BootstrapAPIAPI

All URIs are relative to *http://http:*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEcosystemBootstrap**](BootstrapAPIAPI.md#GetEcosystemBootstrap) | **Get** /bootstrap | Contact the Galasa ecosystem



## GetEcosystemBootstrap

> string GetEcosystemBootstrap(ctx).ClientApiVersion(clientApiVersion).Execute()

Contact the Galasa ecosystem

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
	resp, r, err := apiClient.BootstrapAPIAPI.GetEcosystemBootstrap(context.Background()).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BootstrapAPIAPI.GetEcosystemBootstrap``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcosystemBootstrap`: string
	fmt.Fprintf(os.Stdout, "Response from `BootstrapAPIAPI.GetEcosystemBootstrap`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEcosystemBootstrapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientApiVersion** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

