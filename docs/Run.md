# Run

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | Pointer to **string** |  | [optional] 
**TestStructure** | Pointer to [**TestStructure**](TestStructure.md) |  | [optional] 
**Artifacts** | Pointer to [**[]Artifact**](Artifact.md) |  | [optional] 

## Methods

### NewRun

`func NewRun() *Run`

NewRun instantiates a new Run object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunWithDefaults

`func NewRunWithDefaults() *Run`

NewRunWithDefaults instantiates a new Run object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *Run) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *Run) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *Run) SetRunId(v string)`

SetRunId sets RunId field to given value.

### HasRunId

`func (o *Run) HasRunId() bool`

HasRunId returns a boolean if a field has been set.

### GetTestStructure

`func (o *Run) GetTestStructure() TestStructure`

GetTestStructure returns the TestStructure field if non-nil, zero value otherwise.

### GetTestStructureOk

`func (o *Run) GetTestStructureOk() (*TestStructure, bool)`

GetTestStructureOk returns a tuple with the TestStructure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestStructure

`func (o *Run) SetTestStructure(v TestStructure)`

SetTestStructure sets TestStructure field to given value.

### HasTestStructure

`func (o *Run) HasTestStructure() bool`

HasTestStructure returns a boolean if a field has been set.

### GetArtifacts

`func (o *Run) GetArtifacts() []Artifact`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *Run) GetArtifactsOk() (*[]Artifact, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *Run) SetArtifacts(v []Artifact)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *Run) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


