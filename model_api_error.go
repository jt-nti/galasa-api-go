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

// checks if the APIError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIError{}

// APIError struct for APIError
type APIError struct {
	ErrorCode *int32 `json:"error_code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty"`
}

// NewAPIError instantiates a new APIError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIError() *APIError {
	this := APIError{}
	return &this
}

// NewAPIErrorWithDefaults instantiates a new APIError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIErrorWithDefaults() *APIError {
	this := APIError{}
	return &this
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *APIError) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode) {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIError) GetErrorCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *APIError) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *APIError) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *APIError) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIError) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *APIError) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *APIError) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o APIError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ErrorCode) {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return toSerialize, nil
}

type NullableAPIError struct {
	value *APIError
	isSet bool
}

func (v NullableAPIError) Get() *APIError {
	return v.value
}

func (v *NullableAPIError) Set(val *APIError) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIError) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIError(val *APIError) *NullableAPIError {
	return &NullableAPIError{value: val, isSet: true}
}

func (v NullableAPIError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


