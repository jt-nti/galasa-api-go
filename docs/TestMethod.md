# TestMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassName** | Pointer to **string** |  | [optional] 
**MethodName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**EndTime** | Pointer to **string** |  | [optional] 
**RunLogStart** | Pointer to **int32** |  | [optional] 
**RunLogEnd** | Pointer to **int32** |  | [optional] 
**Befores** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Afters** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewTestMethod

`func NewTestMethod() *TestMethod`

NewTestMethod instantiates a new TestMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestMethodWithDefaults

`func NewTestMethodWithDefaults() *TestMethod`

NewTestMethodWithDefaults instantiates a new TestMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassName

`func (o *TestMethod) GetClassName() string`

GetClassName returns the ClassName field if non-nil, zero value otherwise.

### GetClassNameOk

`func (o *TestMethod) GetClassNameOk() (*string, bool)`

GetClassNameOk returns a tuple with the ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassName

`func (o *TestMethod) SetClassName(v string)`

SetClassName sets ClassName field to given value.

### HasClassName

`func (o *TestMethod) HasClassName() bool`

HasClassName returns a boolean if a field has been set.

### GetMethodName

`func (o *TestMethod) GetMethodName() string`

GetMethodName returns the MethodName field if non-nil, zero value otherwise.

### GetMethodNameOk

`func (o *TestMethod) GetMethodNameOk() (*string, bool)`

GetMethodNameOk returns a tuple with the MethodName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodName

`func (o *TestMethod) SetMethodName(v string)`

SetMethodName sets MethodName field to given value.

### HasMethodName

`func (o *TestMethod) HasMethodName() bool`

HasMethodName returns a boolean if a field has been set.

### GetType

`func (o *TestMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TestMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *TestMethod) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestMethod) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestMethod) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestMethod) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *TestMethod) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TestMethod) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TestMethod) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *TestMethod) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStartTime

`func (o *TestMethod) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TestMethod) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TestMethod) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TestMethod) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *TestMethod) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *TestMethod) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *TestMethod) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *TestMethod) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetRunLogStart

`func (o *TestMethod) GetRunLogStart() int32`

GetRunLogStart returns the RunLogStart field if non-nil, zero value otherwise.

### GetRunLogStartOk

`func (o *TestMethod) GetRunLogStartOk() (*int32, bool)`

GetRunLogStartOk returns a tuple with the RunLogStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunLogStart

`func (o *TestMethod) SetRunLogStart(v int32)`

SetRunLogStart sets RunLogStart field to given value.

### HasRunLogStart

`func (o *TestMethod) HasRunLogStart() bool`

HasRunLogStart returns a boolean if a field has been set.

### GetRunLogEnd

`func (o *TestMethod) GetRunLogEnd() int32`

GetRunLogEnd returns the RunLogEnd field if non-nil, zero value otherwise.

### GetRunLogEndOk

`func (o *TestMethod) GetRunLogEndOk() (*int32, bool)`

GetRunLogEndOk returns a tuple with the RunLogEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunLogEnd

`func (o *TestMethod) SetRunLogEnd(v int32)`

SetRunLogEnd sets RunLogEnd field to given value.

### HasRunLogEnd

`func (o *TestMethod) HasRunLogEnd() bool`

HasRunLogEnd returns a boolean if a field has been set.

### GetBefores

`func (o *TestMethod) GetBefores() []map[string]interface{}`

GetBefores returns the Befores field if non-nil, zero value otherwise.

### GetBeforesOk

`func (o *TestMethod) GetBeforesOk() (*[]map[string]interface{}, bool)`

GetBeforesOk returns a tuple with the Befores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefores

`func (o *TestMethod) SetBefores(v []map[string]interface{})`

SetBefores sets Befores field to given value.

### HasBefores

`func (o *TestMethod) HasBefores() bool`

HasBefores returns a boolean if a field has been set.

### GetAfters

`func (o *TestMethod) GetAfters() []map[string]interface{}`

GetAfters returns the Afters field if non-nil, zero value otherwise.

### GetAftersOk

`func (o *TestMethod) GetAftersOk() (*[]map[string]interface{}, bool)`

GetAftersOk returns a tuple with the Afters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfters

`func (o *TestMethod) SetAfters(v []map[string]interface{})`

SetAfters sets Afters field to given value.

### HasAfters

`func (o *TestMethod) HasAfters() bool`

HasAfters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


