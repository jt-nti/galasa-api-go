# TestStructure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunName** | Pointer to **string** |  | [optional] 
**Bundle** | Pointer to **string** |  | [optional] 
**TestName** | Pointer to **string** |  | [optional] 
**TestShortName** | Pointer to **string** |  | [optional] 
**Requestor** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**Queued** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**Methods** | Pointer to [**[]TestMethod**](TestMethod.md) |  | [optional] 

## Methods

### NewTestStructure

`func NewTestStructure() *TestStructure`

NewTestStructure instantiates a new TestStructure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStructureWithDefaults

`func NewTestStructureWithDefaults() *TestStructure`

NewTestStructureWithDefaults instantiates a new TestStructure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunName

`func (o *TestStructure) GetRunName() string`

GetRunName returns the RunName field if non-nil, zero value otherwise.

### GetRunNameOk

`func (o *TestStructure) GetRunNameOk() (*string, bool)`

GetRunNameOk returns a tuple with the RunName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunName

`func (o *TestStructure) SetRunName(v string)`

SetRunName sets RunName field to given value.

### HasRunName

`func (o *TestStructure) HasRunName() bool`

HasRunName returns a boolean if a field has been set.

### GetBundle

`func (o *TestStructure) GetBundle() string`

GetBundle returns the Bundle field if non-nil, zero value otherwise.

### GetBundleOk

`func (o *TestStructure) GetBundleOk() (*string, bool)`

GetBundleOk returns a tuple with the Bundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundle

`func (o *TestStructure) SetBundle(v string)`

SetBundle sets Bundle field to given value.

### HasBundle

`func (o *TestStructure) HasBundle() bool`

HasBundle returns a boolean if a field has been set.

### GetTestName

`func (o *TestStructure) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *TestStructure) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *TestStructure) SetTestName(v string)`

SetTestName sets TestName field to given value.

### HasTestName

`func (o *TestStructure) HasTestName() bool`

HasTestName returns a boolean if a field has been set.

### GetTestShortName

`func (o *TestStructure) GetTestShortName() string`

GetTestShortName returns the TestShortName field if non-nil, zero value otherwise.

### GetTestShortNameOk

`func (o *TestStructure) GetTestShortNameOk() (*string, bool)`

GetTestShortNameOk returns a tuple with the TestShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestShortName

`func (o *TestStructure) SetTestShortName(v string)`

SetTestShortName sets TestShortName field to given value.

### HasTestShortName

`func (o *TestStructure) HasTestShortName() bool`

HasTestShortName returns a boolean if a field has been set.

### GetRequestor

`func (o *TestStructure) GetRequestor() string`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *TestStructure) GetRequestorOk() (*string, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *TestStructure) SetRequestor(v string)`

SetRequestor sets Requestor field to given value.

### HasRequestor

`func (o *TestStructure) HasRequestor() bool`

HasRequestor returns a boolean if a field has been set.

### GetStatus

`func (o *TestStructure) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestStructure) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestStructure) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestStructure) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *TestStructure) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TestStructure) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TestStructure) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *TestStructure) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetQueued

`func (o *TestStructure) GetQueued() string`

GetQueued returns the Queued field if non-nil, zero value otherwise.

### GetQueuedOk

`func (o *TestStructure) GetQueuedOk() (*string, bool)`

GetQueuedOk returns a tuple with the Queued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueued

`func (o *TestStructure) SetQueued(v string)`

SetQueued sets Queued field to given value.

### HasQueued

`func (o *TestStructure) HasQueued() bool`

HasQueued returns a boolean if a field has been set.

### GetStartTime

`func (o *TestStructure) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TestStructure) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TestStructure) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TestStructure) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *TestStructure) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TestStructure) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TestStructure) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TestStructure) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetMethods

`func (o *TestStructure) GetMethods() []TestMethod`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *TestStructure) GetMethodsOk() (*[]TestMethod, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *TestStructure) SetMethods(v []TestMethod)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *TestStructure) HasMethods() bool`

HasMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


