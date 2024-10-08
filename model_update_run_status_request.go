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

// checks if the UpdateRunStatusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateRunStatusRequest{}

// UpdateRunStatusRequest struct for UpdateRunStatusRequest
type UpdateRunStatusRequest struct {
	Status *string `json:"status,omitempty"`
	Result *string `json:"result,omitempty"`
}

// NewUpdateRunStatusRequest instantiates a new UpdateRunStatusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateRunStatusRequest() *UpdateRunStatusRequest {
	this := UpdateRunStatusRequest{}
	return &this
}

// NewUpdateRunStatusRequestWithDefaults instantiates a new UpdateRunStatusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateRunStatusRequestWithDefaults() *UpdateRunStatusRequest {
	this := UpdateRunStatusRequest{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateRunStatusRequest) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRunStatusRequest) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateRunStatusRequest) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *UpdateRunStatusRequest) SetStatus(v string) {
	o.Status = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *UpdateRunStatusRequest) GetResult() string {
	if o == nil || IsNil(o.Result) {
		var ret string
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateRunStatusRequest) GetResultOk() (*string, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *UpdateRunStatusRequest) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given string and assigns it to the Result field.
func (o *UpdateRunStatusRequest) SetResult(v string) {
	o.Result = &v
}

func (o UpdateRunStatusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateRunStatusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableUpdateRunStatusRequest struct {
	value *UpdateRunStatusRequest
	isSet bool
}

func (v NullableUpdateRunStatusRequest) Get() *UpdateRunStatusRequest {
	return v.value
}

func (v *NullableUpdateRunStatusRequest) Set(val *UpdateRunStatusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRunStatusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRunStatusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRunStatusRequest(val *UpdateRunStatusRequest) *NullableUpdateRunStatusRequest {
	return &NullableUpdateRunStatusRequest{value: val, isSet: true}
}

func (v NullableUpdateRunStatusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRunStatusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


