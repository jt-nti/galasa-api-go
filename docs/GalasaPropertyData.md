# GalasaPropertyData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **string** | The value of the property, maximum length of approximately 1,500,000 characters. It cannot be a blank value.  | [optional] 

## Methods

### NewGalasaPropertyData

`func NewGalasaPropertyData() *GalasaPropertyData`

NewGalasaPropertyData instantiates a new GalasaPropertyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGalasaPropertyDataWithDefaults

`func NewGalasaPropertyDataWithDefaults() *GalasaPropertyData`

NewGalasaPropertyDataWithDefaults instantiates a new GalasaPropertyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *GalasaPropertyData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GalasaPropertyData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GalasaPropertyData) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GalasaPropertyData) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


