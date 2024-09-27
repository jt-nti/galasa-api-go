# UpdateRunStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Result** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateRunStatusRequest

`func NewUpdateRunStatusRequest() *UpdateRunStatusRequest`

NewUpdateRunStatusRequest instantiates a new UpdateRunStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRunStatusRequestWithDefaults

`func NewUpdateRunStatusRequestWithDefaults() *UpdateRunStatusRequest`

NewUpdateRunStatusRequestWithDefaults instantiates a new UpdateRunStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateRunStatusRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateRunStatusRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateRunStatusRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateRunStatusRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *UpdateRunStatusRequest) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateRunStatusRequest) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateRunStatusRequest) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *UpdateRunStatusRequest) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


