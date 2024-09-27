# \RunsAPIAPI

All URIs are relative to *http://http:*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRunsGroup**](RunsAPIAPI.md#GetRunsGroup) | **Get** /runs/{groupId} | Get group runs
[**PostSubmitTestRuns**](RunsAPIAPI.md#PostSubmitTestRuns) | **Post** /runs/{groupId} | Submit test runs



## GetRunsGroup

> TestRuns GetRunsGroup(ctx, groupId).ClientApiVersion(clientApiVersion).Execute()

Get group runs



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
	groupId := "groupId_example" // string | Run Group ID
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunsAPIAPI.GetRunsGroup(context.Background(), groupId).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPIAPI.GetRunsGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRunsGroup`: TestRuns
	fmt.Fprintf(os.Stdout, "Response from `RunsAPIAPI.GetRunsGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Run Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunsGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientApiVersion** | **string** |  | 

### Return type

[**TestRuns**](TestRuns.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSubmitTestRuns

> TestRuns PostSubmitTestRuns(ctx, groupId).TestRunRequest(testRunRequest).ClientApiVersion(clientApiVersion).Execute()

Submit test runs



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
	groupId := "groupId_example" // string | Run Group ID
	testRunRequest := *openapiclient.NewTestRunRequest() // TestRunRequest | 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RunsAPIAPI.PostSubmitTestRuns(context.Background(), groupId).TestRunRequest(testRunRequest).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RunsAPIAPI.PostSubmitTestRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSubmitTestRuns`: TestRuns
	fmt.Fprintf(os.Stdout, "Response from `RunsAPIAPI.PostSubmitTestRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | Run Group ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSubmitTestRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testRunRequest** | [**TestRunRequest**](TestRunRequest.md) |  | 
 **clientApiVersion** | **string** |  | 

### Return type

[**TestRuns**](TestRuns.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

