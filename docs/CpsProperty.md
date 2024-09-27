# CpsProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The fully-qualified property name, including the &#39;namespace&#39;. | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewCpsProperty

`func NewCpsProperty() *CpsProperty`

NewCpsProperty instantiates a new CpsProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCpsPropertyWithDefaults

`func NewCpsPropertyWithDefaults() *CpsProperty`

NewCpsPropertyWithDefaults instantiates a new CpsProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CpsProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CpsProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CpsProperty) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CpsProperty) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *CpsProperty) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CpsProperty) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CpsProperty) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CpsProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


