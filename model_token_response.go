/*
Galasa Ecosystem API

The Galasa Ecosystem REST API allows you to interact with a Galasa Ecosystem.

API version: 0.37.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package galasaapi

import (
	"encoding/json"
)

// checks if the TokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenResponse{}

// TokenResponse struct for TokenResponse
type TokenResponse struct {
	// A JSON Web Token to use when authenticating with a Galasa ecosystem.
	Jwt *string `json:"jwt,omitempty"`
	// A refresh token to allow the retrieval of new JWTs when authenticating with a Galasa ecosystem from a client tool.
	RefreshToken *string `json:"refresh_token,omitempty"`
}

// NewTokenResponse instantiates a new TokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenResponse() *TokenResponse {
	this := TokenResponse{}
	return &this
}

// NewTokenResponseWithDefaults instantiates a new TokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenResponseWithDefaults() *TokenResponse {
	this := TokenResponse{}
	return &this
}

// GetJwt returns the Jwt field value if set, zero value otherwise.
func (o *TokenResponse) GetJwt() string {
	if o == nil || IsNil(o.Jwt) {
		var ret string
		return ret
	}
	return *o.Jwt
}

// GetJwtOk returns a tuple with the Jwt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetJwtOk() (*string, bool) {
	if o == nil || IsNil(o.Jwt) {
		return nil, false
	}
	return o.Jwt, true
}

// HasJwt returns a boolean if a field has been set.
func (o *TokenResponse) HasJwt() bool {
	if o != nil && !IsNil(o.Jwt) {
		return true
	}

	return false
}

// SetJwt gets a reference to the given string and assigns it to the Jwt field.
func (o *TokenResponse) SetJwt(v string) {
	o.Jwt = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *TokenResponse) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *TokenResponse) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *TokenResponse) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

func (o TokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Jwt) {
		toSerialize["jwt"] = o.Jwt
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	return toSerialize, nil
}

type NullableTokenResponse struct {
	value *TokenResponse
	isSet bool
}

func (v NullableTokenResponse) Get() *TokenResponse {
	return v.value
}

func (v *NullableTokenResponse) Set(val *TokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenResponse(val *TokenResponse) *NullableTokenResponse {
	return &NullableTokenResponse{value: val, isSet: true}
}

func (v NullableTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


