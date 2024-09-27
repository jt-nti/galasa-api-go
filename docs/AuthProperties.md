# AuthProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The ID of a Galasa client that is being authenticated. | [optional] 
**RefreshToken** | Pointer to **string** | The refresh token to be exchanged for a JWT, mutually exclusive with \&quot;code\&quot;. | [optional] 
**Code** | Pointer to **string** | The authorization code to be exchanged for a JWT, mutually exclusive with \&quot;refresh_token\&quot;. | [optional] 
**Description** | Pointer to **string** | Optional. A single-line description to be associated with the Galasa personal access token being created for the first time. | [optional] 

## Methods

### NewAuthProperties

`func NewAuthProperties() *AuthProperties`

NewAuthProperties instantiates a new AuthProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthPropertiesWithDefaults

`func NewAuthPropertiesWithDefaults() *AuthProperties`

NewAuthPropertiesWithDefaults instantiates a new AuthProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AuthProperties) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AuthProperties) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AuthProperties) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AuthProperties) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetRefreshToken

`func (o *AuthProperties) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *AuthProperties) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *AuthProperties) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *AuthProperties) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetCode

`func (o *AuthProperties) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthProperties) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthProperties) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthProperties) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *AuthProperties) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthProperties) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthProperties) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthProperties) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


