# GalasaPropertyMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** | Property Namespace. First character of the namespace must be in the &#39;a&#39;-&#39;z&#39; range, and following characters can be &#39;a&#39;-&#39;z&#39; or &#39;0&#39;-&#39;9&#39;  | [optional] 
**Name** | Pointer to **string** | Property Name. The first character of the namespace must be in the &#39;a&#39;-&#39;z&#39; or &#39;A&#39;-&#39;Z&#39; ranges, and following characters can be &#39;a&#39;-&#39;z&#39;, &#39;A&#39;-&#39;Z&#39;, &#39;0&#39;-&#39;9&#39;, &#39;.&#39; (period), &#39;-&#39; (dash) or &#39;_&#39; (underscore). The property name is made of &#39;parts&#39;. Each part must be separated by a period character. There must be at least two parts in the property name. The last character cannot be a period character.  | [optional] 

## Methods

### NewGalasaPropertyMetadata

`func NewGalasaPropertyMetadata() *GalasaPropertyMetadata`

NewGalasaPropertyMetadata instantiates a new GalasaPropertyMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGalasaPropertyMetadataWithDefaults

`func NewGalasaPropertyMetadataWithDefaults() *GalasaPropertyMetadata`

NewGalasaPropertyMetadataWithDefaults instantiates a new GalasaPropertyMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *GalasaPropertyMetadata) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GalasaPropertyMetadata) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GalasaPropertyMetadata) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GalasaPropertyMetadata) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetName

`func (o *GalasaPropertyMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GalasaPropertyMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GalasaPropertyMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GalasaPropertyMetadata) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


