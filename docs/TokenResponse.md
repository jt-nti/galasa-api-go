# TokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwt** | Pointer to **string** | A JSON Web Token to use when authenticating with a Galasa ecosystem. | [optional] 
**RefreshToken** | Pointer to **string** | A refresh token to allow the retrieval of new JWTs when authenticating with a Galasa ecosystem from a client tool. | [optional] 

## Methods

### NewTokenResponse

`func NewTokenResponse() *TokenResponse`

NewTokenResponse instantiates a new TokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenResponseWithDefaults

`func NewTokenResponseWithDefaults() *TokenResponse`

NewTokenResponseWithDefaults instantiates a new TokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwt

`func (o *TokenResponse) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *TokenResponse) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *TokenResponse) SetJwt(v string)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *TokenResponse) HasJwt() bool`

HasJwt returns a boolean if a field has been set.

### GetRefreshToken

`func (o *TokenResponse) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *TokenResponse) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *TokenResponse) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *TokenResponse) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


