/*
Galasa Ecosystem API

The Galasa Ecosystem REST API allows you to interact with a Galasa Ecosystem.

API version: 0.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pkg/galasaapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// RunsAPIAPIService RunsAPIAPI service
type RunsAPIAPIService service

type ApiGetRunsGroupRequest struct {
	ctx context.Context
	ApiService *RunsAPIAPIService
	groupId string
	clientApiVersion *string
}

func (r ApiGetRunsGroupRequest) ClientApiVersion(clientApiVersion string) ApiGetRunsGroupRequest {
	r.clientApiVersion = &clientApiVersion
	return r
}

func (r ApiGetRunsGroupRequest) Execute() (*TestRuns, *http.Response, error) {
	return r.ApiService.GetRunsGroupExecute(r)
}

/*
GetRunsGroup Get group runs

Returns the details of a group of runs.

Requests to this endpoint require a valid bearer token in JWT format to be provided
in the 'Authorization' header (e.g. 'Authorization: Bearer <bearer-token>').


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Run Group ID
 @return ApiGetRunsGroupRequest
*/
func (a *RunsAPIAPIService) GetRunsGroup(ctx context.Context, groupId string) ApiGetRunsGroupRequest {
	return ApiGetRunsGroupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return TestRuns
func (a *RunsAPIAPIService) GetRunsGroupExecute(r ApiGetRunsGroupRequest) (*TestRuns, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TestRuns
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RunsAPIAPIService.GetRunsGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/runs/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "ClientApiVersion", r.clientApiVersion, "simple", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPostSubmitTestRunsRequest struct {
	ctx context.Context
	ApiService *RunsAPIAPIService
	groupId string
	testRunRequest *TestRunRequest
	clientApiVersion *string
}

func (r ApiPostSubmitTestRunsRequest) TestRunRequest(testRunRequest TestRunRequest) ApiPostSubmitTestRunsRequest {
	r.testRunRequest = &testRunRequest
	return r
}

func (r ApiPostSubmitTestRunsRequest) ClientApiVersion(clientApiVersion string) ApiPostSubmitTestRunsRequest {
	r.clientApiVersion = &clientApiVersion
	return r
}

func (r ApiPostSubmitTestRunsRequest) Execute() (*TestRuns, *http.Response, error) {
	return r.ApiService.PostSubmitTestRunsExecute(r)
}

/*
PostSubmitTestRuns Submit test runs

Submits a portfolio of test runs to the ecosystem.

Requests to this endpoint require a valid bearer token in JWT format to be provided
in the 'Authorization' header (e.g. 'Authorization: Bearer <bearer-token>').


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Run Group ID
 @return ApiPostSubmitTestRunsRequest
*/
func (a *RunsAPIAPIService) PostSubmitTestRuns(ctx context.Context, groupId string) ApiPostSubmitTestRunsRequest {
	return ApiPostSubmitTestRunsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return TestRuns
func (a *RunsAPIAPIService) PostSubmitTestRunsExecute(r ApiPostSubmitTestRunsRequest) (*TestRuns, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TestRuns
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RunsAPIAPIService.PostSubmitTestRuns")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/runs/{groupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.testRunRequest == nil {
		return localVarReturnValue, nil, reportError("testRunRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.clientApiVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "ClientApiVersion", r.clientApiVersion, "simple", "")
	}
	// body params
	localVarPostBody = r.testRunRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
