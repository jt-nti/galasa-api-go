# TestRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Test** | Pointer to **string** |  | [optional] 
**BundleName** | Pointer to **string** |  | [optional] 
**TestName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Queued** | Pointer to **string** |  | [optional] 
**Requestor** | Pointer to **string** |  | [optional] 
**Stream** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** |  | [optional] 
**Obr** | Pointer to **string** |  | [optional] 
**Rerun** | Pointer to **bool** |  | [optional] 
**RerunReason** | Pointer to **string** |  | [optional] 
**Local** | Pointer to **bool** |  | [optional] 
**Trace** | Pointer to **bool** |  | [optional] 
**RasRunId** | Pointer to **string** |  | [optional] 

## Methods

### NewTestRun

`func NewTestRun() *TestRun`

NewTestRun instantiates a new TestRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestRunWithDefaults

`func NewTestRunWithDefaults() *TestRun`

NewTestRunWithDefaults instantiates a new TestRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TestRun) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TestRun) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TestRun) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TestRun) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *TestRun) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestRun) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestRun) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TestRun) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *TestRun) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *TestRun) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *TestRun) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *TestRun) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetTest

`func (o *TestRun) GetTest() string`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *TestRun) GetTestOk() (*string, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *TestRun) SetTest(v string)`

SetTest sets Test field to given value.

### HasTest

`func (o *TestRun) HasTest() bool`

HasTest returns a boolean if a field has been set.

### GetBundleName

`func (o *TestRun) GetBundleName() string`

GetBundleName returns the BundleName field if non-nil, zero value otherwise.

### GetBundleNameOk

`func (o *TestRun) GetBundleNameOk() (*string, bool)`

GetBundleNameOk returns a tuple with the BundleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleName

`func (o *TestRun) SetBundleName(v string)`

SetBundleName sets BundleName field to given value.

### HasBundleName

`func (o *TestRun) HasBundleName() bool`

HasBundleName returns a boolean if a field has been set.

### GetTestName

`func (o *TestRun) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *TestRun) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *TestRun) SetTestName(v string)`

SetTestName sets TestName field to given value.

### HasTestName

`func (o *TestRun) HasTestName() bool`

HasTestName returns a boolean if a field has been set.

### GetStatus

`func (o *TestRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestRun) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestRun) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *TestRun) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TestRun) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TestRun) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *TestRun) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetQueued

`func (o *TestRun) GetQueued() string`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *TestRun) GetQueuedOk() (*string, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *TestRun) SetQueued(v string)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *TestRun) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetRequestor

`func (o *TestRun) GetRequestor() string`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *TestRun) GetRequestorOk() (*string, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *TestRun) SetRequestor(v string)`

SetRequestor sets Requestor field to given value.

### HasRequestor

`func (o *TestRun) HasRequestor() bool`

HasRequestor returns a boolean if a field has been set.

### GetStream

`func (o *TestRun) GetStream() string`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *TestRun) GetStreamOk() (*string, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *TestRun) SetStream(v string)`

SetStream sets Stream field to given value.

### HasStream

`func (o *TestRun) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetRepo

`func (o *TestRun) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *TestRun) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *TestRun) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *TestRun) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetObr

`func (o *TestRun) GetObr() string`

GetObr returns the Obr field if non-nil, zero value otherwise.

### GetObrOk

`func (o *TestRun) GetObrOk() (*string, bool)`

GetObrOk returns a tuple with the Obr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObr

`func (o *TestRun) SetObr(v string)`

SetObr sets Obr field to given value.

### HasObr

`func (o *TestRun) HasObr() bool`

HasObr returns a boolean if a field has been set.

### GetRerun

`func (o *TestRun) GetRerun() bool`

GetRerun returns the Rerun field if non-nil, zero value otherwise.

### GetRerunOk

`func (o *TestRun) GetRerunOk() (*bool, bool)`

GetRerunOk returns a tuple with the Rerun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerun

`func (o *TestRun) SetRerun(v bool)`

SetRerun sets Rerun field to given value.

### HasRerun

`func (o *TestRun) HasRerun() bool`

HasRerun returns a boolean if a field has been set.

### GetRerunReason

`func (o *TestRun) GetRerunReason() string`

GetRerunReason returns the RerunReason field if non-nil, zero value otherwise.

### GetRerunReasonOk

`func (o *TestRun) GetRerunReasonOk() (*string, bool)`

GetRerunReasonOk returns a tuple with the RerunReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerunReason

`func (o *TestRun) SetRerunReason(v string)`

SetRerunReason sets RerunReason field to given value.

### HasRerunReason

`func (o *TestRun) HasRerunReason() bool`

HasRerunReason returns a boolean if a field has been set.

### GetLocal

`func (o *TestRun) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *TestRun) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *TestRun) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *TestRun) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetTrace

`func (o *TestRun) GetTrace() bool`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *TestRun) GetTraceOk() (*bool, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *TestRun) SetTrace(v bool)`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *TestRun) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### GetRasRunId

`func (o *TestRun) GetRasRunId() string`

GetRasRunId returns the RasRunId field if non-nil, zero value otherwise.

### GetRasRunIdOk

`func (o *TestRun) GetRasRunIdOk() (*string, bool)`

GetRasRunIdOk returns a tuple with the RasRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRasRunId

`func (o *TestRun) SetRasRunId(v string)`

SetRasRunId sets RasRunId field to given value.

### HasRasRunId

`func (o *TestRun) HasRasRunId() bool`

HasRasRunId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


