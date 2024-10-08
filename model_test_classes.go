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

// checks if the TestClasses type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TestClasses{}

// TestClasses struct for TestClasses
type TestClasses struct {
	Testclasses []TestClass `json:"testclasses,omitempty"`
}

// NewTestClasses instantiates a new TestClasses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestClasses() *TestClasses {
	this := TestClasses{}
	return &this
}

// NewTestClassesWithDefaults instantiates a new TestClasses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestClassesWithDefaults() *TestClasses {
	this := TestClasses{}
	return &this
}

// GetTestclasses returns the Testclasses field value if set, zero value otherwise.
func (o *TestClasses) GetTestclasses() []TestClass {
	if o == nil || IsNil(o.Testclasses) {
		var ret []TestClass
		return ret
	}
	return o.Testclasses
}

// GetTestclassesOk returns a tuple with the Testclasses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestClasses) GetTestclassesOk() ([]TestClass, bool) {
	if o == nil || IsNil(o.Testclasses) {
		return nil, false
	}
	return o.Testclasses, true
}

// HasTestclasses returns a boolean if a field has been set.
func (o *TestClasses) HasTestclasses() bool {
	if o != nil && !IsNil(o.Testclasses) {
		return true
	}

	return false
}

// SetTestclasses gets a reference to the given []TestClass and assigns it to the Testclasses field.
func (o *TestClasses) SetTestclasses(v []TestClass) {
	o.Testclasses = v
}

func (o TestClasses) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestClasses) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Testclasses) {
		toSerialize["testclasses"] = o.Testclasses
	}
	return toSerialize, nil
}

type NullableTestClasses struct {
	value *TestClasses
	isSet bool
}

func (v NullableTestClasses) Get() *TestClasses {
	return v.value
}

func (v *NullableTestClasses) Set(val *TestClasses) {
	v.value = val
	v.isSet = true
}

func (v NullableTestClasses) IsSet() bool {
	return v.isSet
}

func (v *NullableTestClasses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestClasses(val *TestClasses) *NullableTestClasses {
	return &NullableTestClasses{value: val, isSet: true}
}

func (v NullableTestClasses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestClasses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


