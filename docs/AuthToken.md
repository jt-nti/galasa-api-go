# AuthToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenId** | Pointer to **string** | The ID of the token&#39;s record in Galasa&#39;s tokens database. | [optional] 
**CreationTime** | Pointer to **string** | The creation date and time of the auth token (e.g. \&quot;2024-04-01T12:25:00.123Z\&quot;). | [optional] 
**Owner** | Pointer to [**User**](User.md) |  | [optional] 
**Description** | Pointer to **string** | A description of the token, provided by the user when creating it. | [optional] 

## Methods

### NewAuthToken

`func NewAuthToken() *AuthToken`

NewAuthToken instantiates a new AuthToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTokenWithDefaults

`func NewAuthTokenWithDefaults() *AuthToken`

NewAuthTokenWithDefaults instantiates a new AuthToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenId

`func (o *AuthToken) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AuthToken) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AuthToken) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *AuthToken) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetCreationTime

`func (o *AuthToken) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *AuthToken) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *AuthToken) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *AuthToken) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetOwner

`func (o *AuthToken) GetOwner() User`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AuthToken) GetOwnerOk() (*User, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AuthToken) SetOwner(v User)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AuthToken) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDescription

`func (o *AuthToken) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthToken) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthToken) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthToken) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


