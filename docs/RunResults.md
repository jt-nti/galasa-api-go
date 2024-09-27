# RunResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageNumber** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**NumPages** | Pointer to **int32** |  | [optional] 
**AmountOfRuns** | Pointer to **int32** |  | [optional] 
**NextCursor** | Pointer to **string** |  | [optional] 
**Runs** | Pointer to [**[]Run**](Run.md) |  | [optional] 

## Methods

### NewRunResults

`func NewRunResults() *RunResults`

NewRunResults instantiates a new RunResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunResultsWithDefaults

`func NewRunResultsWithDefaults() *RunResults`

NewRunResultsWithDefaults instantiates a new RunResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageNumber

`func (o *RunResults) GetPageNumber() int32`

GetPageNumber returns the PageNumber field if non-nil, zero value otherwise.

### GetPageNumberOk

`func (o *RunResults) GetPageNumberOk() (*int32, bool)`

GetPageNumberOk returns a tuple with the PageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageNumber

`func (o *RunResults) SetPageNumber(v int32)`

SetPageNumber sets PageNumber field to given value.

### HasPageNumber

`func (o *RunResults) HasPageNumber() bool`

HasPageNumber returns a boolean if a field has been set.

### GetPageSize

`func (o *RunResults) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *RunResults) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *RunResults) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *RunResults) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetNumPages

`func (o *RunResults) GetNumPages() int32`

GetNumPages returns the NumPages field if non-nil, zero value otherwise.

### GetNumPagesOk

`func (o *RunResults) GetNumPagesOk() (*int32, bool)`

GetNumPagesOk returns a tuple with the NumPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumPages

`func (o *RunResults) SetNumPages(v int32)`

SetNumPages sets NumPages field to given value.

### HasNumPages

`func (o *RunResults) HasNumPages() bool`

HasNumPages returns a boolean if a field has been set.

### GetAmountOfRuns

`func (o *RunResults) GetAmountOfRuns() int32`

GetAmountOfRuns returns the AmountOfRuns field if non-nil, zero value otherwise.

### GetAmountOfRunsOk

`func (o *RunResults) GetAmountOfRunsOk() (*int32, bool)`

GetAmountOfRunsOk returns a tuple with the AmountOfRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountOfRuns

`func (o *RunResults) SetAmountOfRuns(v int32)`

SetAmountOfRuns sets AmountOfRuns field to given value.

### HasAmountOfRuns

`func (o *RunResults) HasAmountOfRuns() bool`

HasAmountOfRuns returns a boolean if a field has been set.

### GetNextCursor

`func (o *RunResults) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *RunResults) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *RunResults) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *RunResults) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### GetRuns

`func (o *RunResults) GetRuns() []Run`

GetRuns returns the Runs field if non-nil, zero value otherwise.

### GetRunsOk

`func (o *RunResults) GetRunsOk() (*[]Run, bool)`

GetRunsOk returns a tuple with the Runs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuns

`func (o *RunResults) SetRuns(v []Run)`

SetRuns sets Runs field to given value.

### HasRuns

`func (o *RunResults) HasRuns() bool`

HasRuns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


