# \ConfigurationPropertyStoreAPIAPI

All URIs are relative to *http://http:*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCpsProperty**](ConfigurationPropertyStoreAPIAPI.md#CreateCpsProperty) | **Post** /cps/{namespace}/properties | Create a new CPS property
[**DeleteCpsProperty**](ConfigurationPropertyStoreAPIAPI.md#DeleteCpsProperty) | **Delete** /cps/{namespace}/properties/{propertyName} | Delete existing CPS property
[**GetAllCpsNamespaces**](ConfigurationPropertyStoreAPIAPI.md#GetAllCpsNamespaces) | **Get** /cps | Get all known CPS namespaces
[**GetCpsNamespaceCascadeProperty**](ConfigurationPropertyStoreAPIAPI.md#GetCpsNamespaceCascadeProperty) | **Get** /cps/namespace/{namespace}/prefix/{prefix}/suffix/{suffix} | Get cascade CPS property
[**GetCpsNamespaceProperties**](ConfigurationPropertyStoreAPIAPI.md#GetCpsNamespaceProperties) | **Get** /cps/namespace/{namespace} | Get all properties for a namepace
[**GetCpsNamespaces**](ConfigurationPropertyStoreAPIAPI.md#GetCpsNamespaces) | **Get** /cps/namespace | Get CPS Namespaces
[**GetCpsProperty**](ConfigurationPropertyStoreAPIAPI.md#GetCpsProperty) | **Get** /cps/{namespace}/properties/{propertyName} | Get single CPS property
[**PutCpsNamespaceProperty**](ConfigurationPropertyStoreAPIAPI.md#PutCpsNamespaceProperty) | **Put** /cps/namespace/{namespace}/property/{property} | Put new CPS Property
[**QueryCpsNamespaceProperties**](ConfigurationPropertyStoreAPIAPI.md#QueryCpsNamespaceProperties) | **Get** /cps/{namespace}/properties | Get all CPS namespace properties
[**UpdateCpsProperty**](ConfigurationPropertyStoreAPIAPI.md#UpdateCpsProperty) | **Put** /cps/{namespace}/properties/{propertyName} | Update existing CPS property



## CreateCpsProperty

> string CreateCpsProperty(ctx, namespace).GalasaProperty(galasaProperty).ClientApiVersion(clientApiVersion).Execute()

Create a new CPS property



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
	namespace := "namespace_example" // string | Property Namespace. First character of the namespace must be in the 'a'-'z' range, and following characters can be 'a'-'z' or '0'-'9' 
	galasaProperty := *openapiclient.NewGalasaProperty() // GalasaProperty | The value of the property being created.
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationPropertyStoreAPIAPI.CreateCpsProperty(context.Background(), namespace).GalasaProperty(galasaProperty).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPropertyStoreAPIAPI.CreateCpsProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCpsProperty`: string
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationPropertyStoreAPIAPI.CreateCpsProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Property Namespace. First character of the namespace must be in the &#39;a&#39;-&#39;z&#39; range, and following characters can be &#39;a&#39;-&#39;z&#39; or &#39;0&#39;-&#39;9&#39;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCpsPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **galasaProperty** | [**GalasaProperty**](GalasaProperty.md) | The value of the property being created. | 
 **clientApiVersion** | **string** |  | 

### Return type

**string**

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCpsProperty

> string DeleteCpsProperty(ctx, namespace, propertyName).ClientApiVersion(clientApiVersion).Execute()

Delete existing CPS property



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
	namespace := "namespace_example" // string | Property Namespace. First character of the namespace must be in the 'a'-'z' range, and following characters can be 'a'-'z' or '0'-'9' 
	propertyName := "propertyName_example" // string | Property Name. The first character of the namespace must be in the 'a'-'z' or 'A'-'Z' ranges, and following characters can be 'a'-'z', 'A'-'Z', '0'-'9', '.' (period), '-' (dash) or '_' (underscore) 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationPropertyStoreAPIAPI.DeleteCpsProperty(context.Background(), namespace, propertyName).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPropertyStoreAPIAPI.DeleteCpsProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCpsProperty`: string
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationPropertyStoreAPIAPI.DeleteCpsProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Property Namespace. First character of the namespace must be in the &#39;a&#39;-&#39;z&#39; range, and following characters can be &#39;a&#39;-&#39;z&#39; or &#39;0&#39;-&#39;9&#39;  | 
**propertyName** | **string** | Property Name. The first character of the namespace must be in the &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; ranges, and following characters can be &#39;a&#39;-&#39;z&#39;, &#39;A&#39;-&#39;Z&#39;, &#39;0&#39;-&#39;9&#39;, &#39;.&#39; (period), &#39;-&#39; (dash) or &#39;_&#39; (underscore)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCpsPropertyRequest struct via the builder pattern


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


## GetAllCpsNamespaces

> []Namespace GetAllCpsNamespaces(ctx).ClientApiVersion(clientApiVersion).Execute()

Get all known CPS namespaces



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
	resp, r, err := apiClient.ConfigurationPropertyStoreAPIAPI.GetAllCpsNamespaces(context.Background()).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPropertyStoreAPIAPI.GetAllCpsNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllCpsNamespaces`: []Namespace
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationPropertyStoreAPIAPI.GetAllCpsNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllCpsNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientApiVersion** | **string** |  | 

### Return type

[**[]Namespace**](Namespace.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCpsNamespaceCascadeProperty

> CpsProperty GetCpsNamespaceCascadeProperty(ctx, namespace, prefix, suffix).ClientApiVersion(clientApiVersion).Infixes(infixes).Execute()

Get cascade CPS property



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
	namespace := "namespace_example" // string | Property Namespace. The first character of the namespace must be in the 'a'-'z' or 'A'-'Z' ranges, and following characters can be 'a'-'z' or 'A'-'Z' or '0'-'9' 
	prefix := "prefix_example" // string | Property Prefix. The first character of the namespace must be in the 'a'-'z' or 'A'-'Z' ranges,   and following characters can be 'a'-'z', 'A'-'Z', '0'-'9', '.' (period), '-' (dash) or '_' (underscore) 
	suffix := "suffix_example" // string | Property suffix. The first character of the namespace must be in the 'a'-'z' or 'A'-'Z' ranges, and following characters can be 'a'-'z', 'A'-'Z', '0'-'9', '.' (period), '-' (dash) or '_' (underscore) 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)
	infixes := []string{"Inner_example"} // []string | Property infixes. The first character of the namespace must be in the 'a'-'z' or 'A'-'Z' ranges, and following characters can be 'a'-'z', 'A'-'Z', '0'-'9', '.' (period), '-' (dash) or '_' (underscore)  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationPropertyStoreAPIAPI.GetCpsNamespaceCascadeProperty(context.Background(), namespace, prefix, suffix).ClientApiVersion(clientApiVersion).Infixes(infixes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPropertyStoreAPIAPI.GetCpsNamespaceCascadeProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCpsNamespaceCascadeProperty`: CpsProperty
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationPropertyStoreAPIAPI.GetCpsNamespaceCascadeProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Property Namespace. The first character of the namespace must be in the &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; ranges, and following characters can be &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; or &#39;0&#39;-&#39;9&#39;  | 
**prefix** | **string** | Property Prefix. The first character of the namespace must be in the &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; ranges,   and following characters can be &#39;a&#39;-&#39;z&#39;, &#39;A&#39;-&#39;Z&#39;, &#39;0&#39;-&#39;9&#39;, &#39;.&#39; (period), &#39;-&#39; (dash) or &#39;_&#39; (underscore)  | 
**suffix** | **string** | Property suffix. The first character of the namespace must be in the &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; ranges, and following characters can be &#39;a&#39;-&#39;z&#39;, &#39;A&#39;-&#39;Z&#39;, &#39;0&#39;-&#39;9&#39;, &#39;.&#39; (period), &#39;-&#39; (dash) or &#39;_&#39; (underscore)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCpsNamespaceCascadePropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **clientApiVersion** | **string** |  | 
 **infixes** | **[]string** | Property infixes. The first character of the namespace must be in the &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; ranges, and following characters can be &#39;a&#39;-&#39;z&#39;, &#39;A&#39;-&#39;Z&#39;, &#39;0&#39;-&#39;9&#39;, &#39;.&#39; (period), &#39;-&#39; (dash) or &#39;_&#39; (underscore)  | 

### Return type

[**CpsProperty**](CpsProperty.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCpsNamespaceProperties

> []GalasaProperty GetCpsNamespaceProperties(ctx, namespace).ClientApiVersion(clientApiVersion).Execute()

Get all properties for a namepace

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
	namespace := "namespace_example" // string | Property Namespace. First character of the namespace must be in the 'a'-'z' range, and following characters can be 'a'-'z' or '0'-'9' 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationPropertyStoreAPIAPI.GetCpsNamespaceProperties(context.Background(), namespace).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPropertyStoreAPIAPI.GetCpsNamespaceProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCpsNamespaceProperties`: []GalasaProperty
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationPropertyStoreAPIAPI.GetCpsNamespaceProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Property Namespace. First character of the namespace must be in the &#39;a&#39;-&#39;z&#39; range, and following characters can be &#39;a&#39;-&#39;z&#39; or &#39;0&#39;-&#39;9&#39;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCpsNamespacePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientApiVersion** | **string** |  | 

### Return type

[**[]GalasaProperty**](GalasaProperty.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCpsNamespaces

> []string GetCpsNamespaces(ctx).ClientApiVersion(clientApiVersion).Execute()

Get CPS Namespaces

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
	resp, r, err := apiClient.ConfigurationPropertyStoreAPIAPI.GetCpsNamespaces(context.Background()).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPropertyStoreAPIAPI.GetCpsNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCpsNamespaces`: []string
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationPropertyStoreAPIAPI.GetCpsNamespaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCpsNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientApiVersion** | **string** |  | 

### Return type

**[]string**

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCpsProperty

> []GalasaProperty GetCpsProperty(ctx, namespace, propertyName).ClientApiVersion(clientApiVersion).Execute()

Get single CPS property



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
	namespace := "namespace_example" // string | Property Namespace. First character of the namespace must be in the 'a'-'z' range, and following characters can be 'a'-'z' or '0'-'9' 
	propertyName := "propertyName_example" // string | Property Name. The first character of the namespace must be in the 'a'-'z' or 'A'-'Z' ranges, and following characters can be 'a'-'z', 'A'-'Z', '0'-'9', '.' (period), '-' (dash) or '_' (underscore) 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationPropertyStoreAPIAPI.GetCpsProperty(context.Background(), namespace, propertyName).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPropertyStoreAPIAPI.GetCpsProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCpsProperty`: []GalasaProperty
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationPropertyStoreAPIAPI.GetCpsProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Property Namespace. First character of the namespace must be in the &#39;a&#39;-&#39;z&#39; range, and following characters can be &#39;a&#39;-&#39;z&#39; or &#39;0&#39;-&#39;9&#39;  | 
**propertyName** | **string** | Property Name. The first character of the namespace must be in the &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; ranges, and following characters can be &#39;a&#39;-&#39;z&#39;, &#39;A&#39;-&#39;Z&#39;, &#39;0&#39;-&#39;9&#39;, &#39;.&#39; (period), &#39;-&#39; (dash) or &#39;_&#39; (underscore)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCpsPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientApiVersion** | **string** |  | 

### Return type

[**[]GalasaProperty**](GalasaProperty.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutCpsNamespaceProperty

> CpsProperty PutCpsNamespaceProperty(ctx, namespace, property).CpsProperty(cpsProperty).ClientApiVersion(clientApiVersion).Execute()

Put new CPS Property



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
	namespace := "namespace_example" // string | Property Namespace. First character of the namespace must be in the 'a'-'z' range, and following characters can be 'a'-'z' or '0'-'9' 
	property := "property_example" // string | Property Name. The first character of the namespace must be in the 'a'-'z' or 'A'-'Z' ranges, and following characters can be 'a'-'z', 'A'-'Z', '0'-'9', '.' (period), '-' (dash) or '_' (underscore) 
	cpsProperty := *openapiclient.NewCpsProperty() // CpsProperty | 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationPropertyStoreAPIAPI.PutCpsNamespaceProperty(context.Background(), namespace, property).CpsProperty(cpsProperty).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPropertyStoreAPIAPI.PutCpsNamespaceProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutCpsNamespaceProperty`: CpsProperty
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationPropertyStoreAPIAPI.PutCpsNamespaceProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Property Namespace. First character of the namespace must be in the &#39;a&#39;-&#39;z&#39; range, and following characters can be &#39;a&#39;-&#39;z&#39; or &#39;0&#39;-&#39;9&#39;  | 
**property** | **string** | Property Name. The first character of the namespace must be in the &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; ranges, and following characters can be &#39;a&#39;-&#39;z&#39;, &#39;A&#39;-&#39;Z&#39;, &#39;0&#39;-&#39;9&#39;, &#39;.&#39; (period), &#39;-&#39; (dash) or &#39;_&#39; (underscore)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutCpsNamespacePropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cpsProperty** | [**CpsProperty**](CpsProperty.md) |  | 
 **clientApiVersion** | **string** |  | 

### Return type

[**CpsProperty**](CpsProperty.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryCpsNamespaceProperties

> []GalasaProperty QueryCpsNamespaceProperties(ctx, namespace).ClientApiVersion(clientApiVersion).Prefix(prefix).Suffix(suffix).Infix(infix).Execute()

Get all CPS namespace properties



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
	namespace := "namespace_example" // string | Property Namespace. First character of the namespace must be in the 'a'-'z' range, and following characters can be 'a'-'z' or '0'-'9' 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)
	prefix := "prefix_example" // string | Property Prefix. The first character must be in the 'a'-'z' or 'A'-'Z' ranges,   and following characters can be 'a'-'z', 'A'-'Z', '0'-'9', '.' (period), '-' (dash) or '_' (underscore)  (optional)
	suffix := "suffix_example" // string | Property suffix. The first character must be in the 'a'-'z' or 'A'-'Z' ranges, and following characters can be 'a'-'z', 'A'-'Z', '0'-'9', '.' (period), '-' (dash) or '_' (underscore)  (optional)
	infix := "infix_example" // string | Comma-separated list of Property infixes. The first character of each infix must be in the 'a'-'z' or 'A'-'Z' ranges, and following characters can be 'a'-'z', 'A'-'Z', '0'-'9', '.' (period), '-' (dash) or '_' (underscore)  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationPropertyStoreAPIAPI.QueryCpsNamespaceProperties(context.Background(), namespace).ClientApiVersion(clientApiVersion).Prefix(prefix).Suffix(suffix).Infix(infix).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPropertyStoreAPIAPI.QueryCpsNamespaceProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryCpsNamespaceProperties`: []GalasaProperty
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationPropertyStoreAPIAPI.QueryCpsNamespaceProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Property Namespace. First character of the namespace must be in the &#39;a&#39;-&#39;z&#39; range, and following characters can be &#39;a&#39;-&#39;z&#39; or &#39;0&#39;-&#39;9&#39;  | 

### Other Parameters

Other parameters are passed through a pointer to a apiQueryCpsNamespacePropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientApiVersion** | **string** |  | 
 **prefix** | **string** | Property Prefix. The first character must be in the &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; ranges,   and following characters can be &#39;a&#39;-&#39;z&#39;, &#39;A&#39;-&#39;Z&#39;, &#39;0&#39;-&#39;9&#39;, &#39;.&#39; (period), &#39;-&#39; (dash) or &#39;_&#39; (underscore)  | 
 **suffix** | **string** | Property suffix. The first character must be in the &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; ranges, and following characters can be &#39;a&#39;-&#39;z&#39;, &#39;A&#39;-&#39;Z&#39;, &#39;0&#39;-&#39;9&#39;, &#39;.&#39; (period), &#39;-&#39; (dash) or &#39;_&#39; (underscore)  | 
 **infix** | **string** | Comma-separated list of Property infixes. The first character of each infix must be in the &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; ranges, and following characters can be &#39;a&#39;-&#39;z&#39;, &#39;A&#39;-&#39;Z&#39;, &#39;0&#39;-&#39;9&#39;, &#39;.&#39; (period), &#39;-&#39; (dash) or &#39;_&#39; (underscore)  | 

### Return type

[**[]GalasaProperty**](GalasaProperty.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCpsProperty

> string UpdateCpsProperty(ctx, namespace, propertyName).GalasaProperty(galasaProperty).ClientApiVersion(clientApiVersion).Execute()

Update existing CPS property



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
	namespace := "namespace_example" // string | Property Namespace. First character of the namespace must be in the 'a'-'z' range, and following characters can be 'a'-'z' or '0'-'9' 
	propertyName := "propertyName_example" // string | Property Name. The first character of the namespace must be in the 'a'-'z' or 'A'-'Z' ranges, and following characters can be 'a'-'z', 'A'-'Z', '0'-'9', '.' (period), '-' (dash) or '_' (underscore) 
	galasaProperty := *openapiclient.NewGalasaProperty() // GalasaProperty | The value of the property being created.
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationPropertyStoreAPIAPI.UpdateCpsProperty(context.Background(), namespace, propertyName).GalasaProperty(galasaProperty).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationPropertyStoreAPIAPI.UpdateCpsProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCpsProperty`: string
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationPropertyStoreAPIAPI.UpdateCpsProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | Property Namespace. First character of the namespace must be in the &#39;a&#39;-&#39;z&#39; range, and following characters can be &#39;a&#39;-&#39;z&#39; or &#39;0&#39;-&#39;9&#39;  | 
**propertyName** | **string** | Property Name. The first character of the namespace must be in the &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; ranges, and following characters can be &#39;a&#39;-&#39;z&#39;, &#39;A&#39;-&#39;Z&#39;, &#39;0&#39;-&#39;9&#39;, &#39;.&#39; (period), &#39;-&#39; (dash) or &#39;_&#39; (underscore)  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCpsPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **galasaProperty** | [**GalasaProperty**](GalasaProperty.md) | The value of the property being created. | 
 **clientApiVersion** | **string** |  | 

### Return type

**string**

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

