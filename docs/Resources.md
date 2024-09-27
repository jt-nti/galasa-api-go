# Resources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]ResourcesDataInner**](ResourcesDataInner.md) | An array of Galasa Resources that should contain one or more entrys matching any of the schemas outlined | [optional] 

## Methods

### NewResources

`func NewResources() *Resources`

NewResources instantiates a new Resources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcesWithDefaults

`func NewResourcesWithDefaults() *Resources`

NewResourcesWithDefaults instantiates a new Resources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Resources) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Resources) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Resources) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Resources) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetData

`func (o *Resources) GetData() []ResourcesDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Resources) GetDataOk() (*[]ResourcesDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Resources) SetData(v []ResourcesDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *Resources) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


