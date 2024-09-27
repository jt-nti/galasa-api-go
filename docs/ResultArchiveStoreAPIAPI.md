# \ResultArchiveStoreAPIAPI

All URIs are relative to *http://http:*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRasRunById**](ResultArchiveStoreAPIAPI.md#DeleteRasRunById) | **Delete** /ras/runs/{runid} | Delete run from the RAS using a given run ID
[**GetRasRequestors**](ResultArchiveStoreAPIAPI.md#GetRasRequestors) | **Get** /ras/requestors | Get all known requestors
[**GetRasResultNames**](ResultArchiveStoreAPIAPI.md#GetRasResultNames) | **Get** /ras/resultnames | Get all the known result names
[**GetRasRunArtifactByPath**](ResultArchiveStoreAPIAPI.md#GetRasRunArtifactByPath) | **Get** /ras/runs/{runid}/files/{artifactPath} | Download artifact for a given runid by artifactPath
[**GetRasRunArtifactList**](ResultArchiveStoreAPIAPI.md#GetRasRunArtifactList) | **Get** /ras/runs/{runid}/artifacts | Get the available run artifacts which can be downloaded.
[**GetRasRunById**](ResultArchiveStoreAPIAPI.md#GetRasRunById) | **Get** /ras/runs/{runid} | Get Run by ID
[**GetRasRunLog**](ResultArchiveStoreAPIAPI.md#GetRasRunLog) | **Get** /ras/runs/{runid}/runlog | Get Run Log
[**GetRasSearchRuns**](ResultArchiveStoreAPIAPI.md#GetRasSearchRuns) | **Get** /ras/runs | Get Runs from Query
[**GetRasTestclasses**](ResultArchiveStoreAPIAPI.md#GetRasTestclasses) | **Get** /ras/testclasses | Get all the known test classes
[**PutRasRunStatusById**](ResultArchiveStoreAPIAPI.md#PutRasRunStatusById) | **Put** /ras/runs/{runid} | Update the status of a test run



## DeleteRasRunById

> DeleteRasRunById(ctx, runid).ClientApiVersion(clientApiVersion).Execute()

Delete run from the RAS using a given run ID



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
	runid := "runid_example" // string | Run Id which starts \"cdb-\". This is NOT the short run-name (e.g. C1234)
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResultArchiveStoreAPIAPI.DeleteRasRunById(context.Background(), runid).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultArchiveStoreAPIAPI.DeleteRasRunById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runid** | **string** | Run Id which starts \&quot;cdb-\&quot;. This is NOT the short run-name (e.g. C1234) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRasRunByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientApiVersion** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRasRequestors

> Requestors GetRasRequestors(ctx).ClientApiVersion(clientApiVersion).Sort(sort).Execute()

Get all known requestors



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
	sort := "sort_example" // string | provides sorting, requestor:asc or requestor:desc (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultArchiveStoreAPIAPI.GetRasRequestors(context.Background()).ClientApiVersion(clientApiVersion).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultArchiveStoreAPIAPI.GetRasRequestors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRasRequestors`: Requestors
	fmt.Fprintf(os.Stdout, "Response from `ResultArchiveStoreAPIAPI.GetRasRequestors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRasRequestorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientApiVersion** | **string** |  | 
 **sort** | **string** | provides sorting, requestor:asc or requestor:desc | 

### Return type

[**Requestors**](Requestors.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRasResultNames

> ResultNames GetRasResultNames(ctx).ClientApiVersion(clientApiVersion).Sort(sort).Execute()

Get all the known result names



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
	sort := "sort_example" // string | provides sorting, results:asc or results:desc (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultArchiveStoreAPIAPI.GetRasResultNames(context.Background()).ClientApiVersion(clientApiVersion).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultArchiveStoreAPIAPI.GetRasResultNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRasResultNames`: ResultNames
	fmt.Fprintf(os.Stdout, "Response from `ResultArchiveStoreAPIAPI.GetRasResultNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRasResultNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientApiVersion** | **string** |  | 
 **sort** | **string** | provides sorting, results:asc or results:desc | 

### Return type

[**ResultNames**](ResultNames.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRasRunArtifactByPath

> *os.File GetRasRunArtifactByPath(ctx, runid, artifactPath).ClientApiVersion(clientApiVersion).Execute()

Download artifact for a given runid by artifactPath



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
	runid := "runid_example" // string | Run Id
	artifactPath := "artifactPath_example" // string | Run Artifact path
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultArchiveStoreAPIAPI.GetRasRunArtifactByPath(context.Background(), runid, artifactPath).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultArchiveStoreAPIAPI.GetRasRunArtifactByPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRasRunArtifactByPath`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `ResultArchiveStoreAPIAPI.GetRasRunArtifactByPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runid** | **string** | Run Id | 
**artifactPath** | **string** | Run Artifact path | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRasRunArtifactByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clientApiVersion** | **string** |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRasRunArtifactList

> []ArtifactIndexEntry GetRasRunArtifactList(ctx, runid).ClientApiVersion(clientApiVersion).Execute()

Get the available run artifacts which can be downloaded.



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
	runid := "runid_example" // string | Run Id
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultArchiveStoreAPIAPI.GetRasRunArtifactList(context.Background(), runid).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultArchiveStoreAPIAPI.GetRasRunArtifactList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRasRunArtifactList`: []ArtifactIndexEntry
	fmt.Fprintf(os.Stdout, "Response from `ResultArchiveStoreAPIAPI.GetRasRunArtifactList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runid** | **string** | Run Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRasRunArtifactListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientApiVersion** | **string** |  | 

### Return type

[**[]ArtifactIndexEntry**](ArtifactIndexEntry.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRasRunById

> Run GetRasRunById(ctx, runid).ClientApiVersion(clientApiVersion).Execute()

Get Run by ID



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
	runid := "runid_example" // string | Run Id which starts \"cdb-\". This is NOT the short run-name (e.g. C1234)
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultArchiveStoreAPIAPI.GetRasRunById(context.Background(), runid).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultArchiveStoreAPIAPI.GetRasRunById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRasRunById`: Run
	fmt.Fprintf(os.Stdout, "Response from `ResultArchiveStoreAPIAPI.GetRasRunById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runid** | **string** | Run Id which starts \&quot;cdb-\&quot;. This is NOT the short run-name (e.g. C1234) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRasRunByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clientApiVersion** | **string** |  | 

### Return type

[**Run**](Run.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRasRunLog

> string GetRasRunLog(ctx, runid).ClientApiVersion(clientApiVersion).Execute()

Get Run Log



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
	runid := "runid_example" // string | Run Id
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultArchiveStoreAPIAPI.GetRasRunLog(context.Background(), runid).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultArchiveStoreAPIAPI.GetRasRunLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRasRunLog`: string
	fmt.Fprintf(os.Stdout, "Response from `ResultArchiveStoreAPIAPI.GetRasRunLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runid** | **string** | Run Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRasRunLogRequest struct via the builder pattern


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


## GetRasSearchRuns

> RunResults GetRasSearchRuns(ctx).Sort(sort).ClientApiVersion(clientApiVersion).Result(result).Status(status).Bundle(bundle).Requestor(requestor).From(from).To(to).Testname(testname).Page(page).Size(size).RunId(runId).Runname(runname).IncludeCursor(includeCursor).Cursor(cursor).Execute()

Get Runs from Query



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	sort := "from:asc" // string | Sorts the returned runs based on the sort field. Supports sorting fields 'from', 'to', 'result' and 'testclass'.  If omitted, runs will be sorted in descending order based on their 'queued' time, which is equivalent to specifying 'from:desc' (i.e. latest queued run first, oldest last).  When sorting with 'to' or 'result', runs that have not yet finished will not be included in responses from this endpoint.  Use '{FIELD-NAME}:asc' to sort in ascending order. Use '{FIELD-NAME}:desc' to sort in descending order. 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)
	result := "Passed" // string | Result Status for the run. Commonly queried values: 'EnvFail','Passed','Failed' Multiple values can be used in the query for example: 'EnvFail,Passed,Failed'. Values are case insensitive. For example 'PASSED' and 'passED' will both be valid.  (optional)
	status := "running" // string | Test run lifecycle status. Current possibles: 'finished','building','generating','running','rundone','up','started','provstart','ending'. These are not case sensitive. Multiple values can be used in the query for example: 'finished,running,started'. Values are case insensitive. For example 'FINISHED' and 'finiSHed' will both be valid.  (optional)
	bundle := "dev.galasa.inttests" // string | The name of the OSGi bundle that the desired test run(s) were loaded with.  (optional)
	requestor := "MyAutomationJobName" // string | Name of the test requestor / submitter (optional)
	from := time.Now() // time.Time | Retrieve runs that started at a time after this date and time.  The only scenario in which from can be omitted is when a runname has been supplied  (optional)
	to := time.Now() // time.Time | Retrieve runs that ended at a date and time prior to this date and time value. If you specify this parameter, only test runs which have completed will be returned. Tests currently in-flight will not be visible.  (optional)
	testname := "dev.galasa.inttests.simbank.local.mvp.SimBankLocalJava11UbuntuMvp" // string | The full test name (package + short test name) (optional)
	page := int32(2) // int32 | Deprecated (since 0.37.0) - Use the 'cursor' query parameter instead. Causes a specific page in the available results to be returned. The first page is page 1. If omitted, then page 1 is returned.  (optional)
	size := int32(20) // int32 | The number of test results returned within each page. If omitted, the default value is 100.  (optional)
	runId := "cdb-a4ddb7fd1dab8d6025e4f3894010d20d" // string | The ID for a specific test run as seen by the RAS. This number is unique across the system, so using this field you can expect one or zero test runs in the first page.  (optional)
	runname := "U1578" // string | The name of the test run for which details will be returned. It will normally be unique, but this is not guaranteed, so you may see multiple results for the same runname under some rare circumstances.  (optional)
	includeCursor := "true" // string | A boolean flag to enable cursor-based pagination and return the next page cursor in the response. If omitted, it will default to false.  (optional)
	cursor := "g2wAAAACaAJkAA5zdGFydG" // string | The cursor representing the page of runs to be retrieved. This is a unique value that is specific to a query and is included in responses, allowing you to navigate through pages of runs. If omitted, the first page of runs for the given query will be returned and the response will display the cursor for the next page of runs.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultArchiveStoreAPIAPI.GetRasSearchRuns(context.Background()).Sort(sort).ClientApiVersion(clientApiVersion).Result(result).Status(status).Bundle(bundle).Requestor(requestor).From(from).To(to).Testname(testname).Page(page).Size(size).RunId(runId).Runname(runname).IncludeCursor(includeCursor).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultArchiveStoreAPIAPI.GetRasSearchRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRasSearchRuns`: RunResults
	fmt.Fprintf(os.Stdout, "Response from `ResultArchiveStoreAPIAPI.GetRasSearchRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRasSearchRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sorts the returned runs based on the sort field. Supports sorting fields &#39;from&#39;, &#39;to&#39;, &#39;result&#39; and &#39;testclass&#39;.  If omitted, runs will be sorted in descending order based on their &#39;queued&#39; time, which is equivalent to specifying &#39;from:desc&#39; (i.e. latest queued run first, oldest last).  When sorting with &#39;to&#39; or &#39;result&#39;, runs that have not yet finished will not be included in responses from this endpoint.  Use &#39;{FIELD-NAME}:asc&#39; to sort in ascending order. Use &#39;{FIELD-NAME}:desc&#39; to sort in descending order.  | 
 **clientApiVersion** | **string** |  | 
 **result** | **string** | Result Status for the run. Commonly queried values: &#39;EnvFail&#39;,&#39;Passed&#39;,&#39;Failed&#39; Multiple values can be used in the query for example: &#39;EnvFail,Passed,Failed&#39;. Values are case insensitive. For example &#39;PASSED&#39; and &#39;passED&#39; will both be valid.  | 
 **status** | **string** | Test run lifecycle status. Current possibles: &#39;finished&#39;,&#39;building&#39;,&#39;generating&#39;,&#39;running&#39;,&#39;rundone&#39;,&#39;up&#39;,&#39;started&#39;,&#39;provstart&#39;,&#39;ending&#39;. These are not case sensitive. Multiple values can be used in the query for example: &#39;finished,running,started&#39;. Values are case insensitive. For example &#39;FINISHED&#39; and &#39;finiSHed&#39; will both be valid.  | 
 **bundle** | **string** | The name of the OSGi bundle that the desired test run(s) were loaded with.  | 
 **requestor** | **string** | Name of the test requestor / submitter | 
 **from** | **time.Time** | Retrieve runs that started at a time after this date and time.  The only scenario in which from can be omitted is when a runname has been supplied  | 
 **to** | **time.Time** | Retrieve runs that ended at a date and time prior to this date and time value. If you specify this parameter, only test runs which have completed will be returned. Tests currently in-flight will not be visible.  | 
 **testname** | **string** | The full test name (package + short test name) | 
 **page** | **int32** | Deprecated (since 0.37.0) - Use the &#39;cursor&#39; query parameter instead. Causes a specific page in the available results to be returned. The first page is page 1. If omitted, then page 1 is returned.  | 
 **size** | **int32** | The number of test results returned within each page. If omitted, the default value is 100.  | 
 **runId** | **string** | The ID for a specific test run as seen by the RAS. This number is unique across the system, so using this field you can expect one or zero test runs in the first page.  | 
 **runname** | **string** | The name of the test run for which details will be returned. It will normally be unique, but this is not guaranteed, so you may see multiple results for the same runname under some rare circumstances.  | 
 **includeCursor** | **string** | A boolean flag to enable cursor-based pagination and return the next page cursor in the response. If omitted, it will default to false.  | 
 **cursor** | **string** | The cursor representing the page of runs to be retrieved. This is a unique value that is specific to a query and is included in responses, allowing you to navigate through pages of runs. If omitted, the first page of runs for the given query will be returned and the response will display the cursor for the next page of runs.  | 

### Return type

[**RunResults**](RunResults.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRasTestclasses

> TestClasses GetRasTestclasses(ctx).ClientApiVersion(clientApiVersion).Sort(sort).Execute()

Get all the known test classes



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
	sort := "sort_example" // string | Provide Sorting (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultArchiveStoreAPIAPI.GetRasTestclasses(context.Background()).ClientApiVersion(clientApiVersion).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultArchiveStoreAPIAPI.GetRasTestclasses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRasTestclasses`: TestClasses
	fmt.Fprintf(os.Stdout, "Response from `ResultArchiveStoreAPIAPI.GetRasTestclasses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRasTestclassesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientApiVersion** | **string** |  | 
 **sort** | **string** | Provide Sorting | 

### Return type

[**TestClasses**](TestClasses.md)

### Authorization

[JwtAuth](../README.md#JwtAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutRasRunStatusById

> string PutRasRunStatusById(ctx, runid).UpdateRunStatusRequest(updateRunStatusRequest).ClientApiVersion(clientApiVersion).Execute()

Update the status of a test run



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
	runid := "runid_example" // string | Run Id which starts \"cdb-\". This is NOT the short run-name (e.g. C1234)
	updateRunStatusRequest := *openapiclient.NewUpdateRunStatusRequest() // UpdateRunStatusRequest | 
	clientApiVersion := "clientApiVersion_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResultArchiveStoreAPIAPI.PutRasRunStatusById(context.Background(), runid).UpdateRunStatusRequest(updateRunStatusRequest).ClientApiVersion(clientApiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResultArchiveStoreAPIAPI.PutRasRunStatusById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutRasRunStatusById`: string
	fmt.Fprintf(os.Stdout, "Response from `ResultArchiveStoreAPIAPI.PutRasRunStatusById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runid** | **string** | Run Id which starts \&quot;cdb-\&quot;. This is NOT the short run-name (e.g. C1234) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutRasRunStatusByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRunStatusRequest** | [**UpdateRunStatusRequest**](UpdateRunStatusRequest.md) |  | 
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

