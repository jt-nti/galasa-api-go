# \UsersAPIAPI

All URIs are relative to *http://http:*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserByLoginId**](UsersAPIAPI.md#GetUserByLoginId) | **Get** /users | Returns the details of the requesting user.



## GetUserByLoginId

> []UserData GetUserByLoginId(ctx).LoginId(loginId).ClientApiVersion(clientApiVersion).Execute()

Returns the details of the requesting user.

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
	loginId := "loginId_example" // string | The loginId of the user whose details will be returned.
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPIAPI.GetUserByLoginId(context.Background()).LoginId(loginId).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPIAPI.GetUserByLoginId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserByLoginId`: []UserData
	fmt.Fprintf(os.Stdout, "Response from `UsersAPIAPI.GetUserByLoginId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByLoginIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loginId** | **string** | The loginId of the user whose details will be returned. | 
 **clientApiVersion** | **string** |  | 

### Return type

[**[]UserData**](UserData.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

