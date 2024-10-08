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

// checks if the AuthProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthProperties{}

// AuthProperties struct for AuthProperties
type AuthProperties struct {
	// The ID of a Galasa client that is being authenticated.
	ClientId *string `json:"client_id,omitempty"`
	// The refresh token to be exchanged for a JWT, mutually exclusive with \"code\".
	RefreshToken *string `json:"refresh_token,omitempty"`
	// The authorization code to be exchanged for a JWT, mutually exclusive with \"refresh_token\".
	Code *string `json:"code,omitempty"`
	// Optional. A single-line description to be associated with the Galasa personal access token being created for the first time.
	Description *string `json:"description,omitempty"`
}

// NewAuthProperties instantiates a new AuthProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthProperties() *AuthProperties {
	this := AuthProperties{}
	return &this
}

// NewAuthPropertiesWithDefaults instantiates a new AuthProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthPropertiesWithDefaults() *AuthProperties {
	this := AuthProperties{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AuthProperties) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProperties) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AuthProperties) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AuthProperties) SetClientId(v string) {
	o.ClientId = &v
}

// GetRefreshToken returns the RefreshToken field value if set, zero value otherwise.
func (o *AuthProperties) GetRefreshToken() string {
	if o == nil || IsNil(o.RefreshToken) {
		var ret string
		return ret
	}
	return *o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProperties) GetRefreshTokenOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshToken) {
		return nil, false
	}
	return o.RefreshToken, true
}

// HasRefreshToken returns a boolean if a field has been set.
func (o *AuthProperties) HasRefreshToken() bool {
	if o != nil && !IsNil(o.RefreshToken) {
		return true
	}

	return false
}

// SetRefreshToken gets a reference to the given string and assigns it to the RefreshToken field.
func (o *AuthProperties) SetRefreshToken(v string) {
	o.RefreshToken = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AuthProperties) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProperties) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AuthProperties) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AuthProperties) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthProperties) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthProperties) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthProperties) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthProperties) SetDescription(v string) {
	o.Description = &v
}

func (o AuthProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["client_id"] = o.ClientId
	}
	if !IsNil(o.RefreshToken) {
		toSerialize["refresh_token"] = o.RefreshToken
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableAuthProperties struct {
	value *AuthProperties
	isSet bool
}

func (v NullableAuthProperties) Get() *AuthProperties {
	return v.value
}

func (v *NullableAuthProperties) Set(val *AuthProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthProperties(val *AuthProperties) *NullableAuthProperties {
	return &NullableAuthProperties{value: val, isSet: true}
}

func (v NullableAuthProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


