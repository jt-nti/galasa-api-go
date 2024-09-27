# GalasaProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**GalasaPropertyMetadata**](GalasaPropertyMetadata.md) |  | [optional] 
**Data** | Pointer to [**GalasaPropertyData**](GalasaPropertyData.md) |  | [optional] 

## Methods

### NewGalasaProperty

`func NewGalasaProperty() *GalasaProperty`

NewGalasaProperty instantiates a new GalasaProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGalasaPropertyWithDefaults

`func NewGalasaPropertyWithDefaults() *GalasaProperty`

NewGalasaPropertyWithDefaults instantiates a new GalasaProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *GalasaProperty) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GalasaProperty) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GalasaProperty) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *GalasaProperty) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *GalasaProperty) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GalasaProperty) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GalasaProperty) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GalasaProperty) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *GalasaProperty) GetMetadata() GalasaPropertyMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GalasaProperty) GetMetadataOk() (*GalasaPropertyMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GalasaProperty) SetMetadata(v GalasaPropertyMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GalasaProperty) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetData

`func (o *GalasaProperty) GetData() GalasaPropertyData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GalasaProperty) GetDataOk() (*GalasaPropertyData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GalasaProperty) SetData(v GalasaPropertyData)`

SetData sets Data field to given value.

### HasData

`func (o *GalasaProperty) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


