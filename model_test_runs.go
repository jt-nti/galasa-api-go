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

// checks if the TestRuns type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestRuns{}

// TestRuns struct for TestRuns
type TestRuns struct {
	Complete *bool `json:"complete,omitempty"`
	Runs []TestRun `json:"runs,omitempty"`
}

// NewTestRuns instantiates a new TestRuns object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestRuns() *TestRuns {
	this := TestRuns{}
	return &this
}

// NewTestRunsWithDefaults instantiates a new TestRuns object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestRunsWithDefaults() *TestRuns {
	this := TestRuns{}
	return &this
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *TestRuns) GetComplete() bool {
	if o == nil || IsNil(o.Complete) {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRuns) GetCompleteOk() (*bool, bool) {
	if o == nil || IsNil(o.Complete) {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *TestRuns) HasComplete() bool {
	if o != nil && !IsNil(o.Complete) {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *TestRuns) SetComplete(v bool) {
	o.Complete = &v
}

// GetRuns returns the Runs field value if set, zero value otherwise.
func (o *TestRuns) GetRuns() []TestRun {
	if o == nil || IsNil(o.Runs) {
		var ret []TestRun
		return ret
	}
	return o.Runs
}

// GetRunsOk returns a tuple with the Runs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestRuns) GetRunsOk() ([]TestRun, bool) {
	if o == nil || IsNil(o.Runs) {
		return nil, false
	}
	return o.Runs, true
}

// HasRuns returns a boolean if a field has been set.
func (o *TestRuns) HasRuns() bool {
	if o != nil && !IsNil(o.Runs) {
		return true
	}

	return false
}

// SetRuns gets a reference to the given []TestRun and assigns it to the Runs field.
func (o *TestRuns) SetRuns(v []TestRun) {
	o.Runs = v
}

func (o TestRuns) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestRuns) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Complete) {
		toSerialize["complete"] = o.Complete
	}
	if !IsNil(o.Runs) {
		toSerialize["runs"] = o.Runs
	}
	return toSerialize, nil
}

type NullableTestRuns struct {
	value *TestRuns
	isSet bool
}

func (v NullableTestRuns) Get() *TestRuns {
	return v.value
}

func (v *NullableTestRuns) Set(val *TestRuns) {
	v.value = val
	v.isSet = true
}

func (v NullableTestRuns) IsSet() bool {
	return v.isSet
}

func (v *NullableTestRuns) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestRuns(val *TestRuns) *NullableTestRuns {
	return &NullableTestRuns{value: val, isSet: true}
}

func (v NullableTestRuns) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestRuns) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


