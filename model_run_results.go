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

// checks if the RunResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RunResults{}

// RunResults struct for RunResults
type RunResults struct {
	PageNumber *int32 `json:"pageNumber,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	NumPages *int32 `json:"numPages,omitempty"`
	AmountOfRuns *int32 `json:"amountOfRuns,omitempty"`
	NextCursor *string `json:"nextCursor,omitempty"`
	Runs []Run `json:"runs,omitempty"`
}

// NewRunResults instantiates a new RunResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRunResults() *RunResults {
	this := RunResults{}
	return &this
}

// NewRunResultsWithDefaults instantiates a new RunResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRunResultsWithDefaults() *RunResults {
	this := RunResults{}
	return &this
}

// GetPageNumber returns the PageNumber field value if set, zero value otherwise.
func (o *RunResults) GetPageNumber() int32 {
	if o == nil || IsNil(o.PageNumber) {
		var ret int32
		return ret
	}
	return *o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunResults) GetPageNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PageNumber) {
		return nil, false
	}
	return o.PageNumber, true
}

// HasPageNumber returns a boolean if a field has been set.
func (o *RunResults) HasPageNumber() bool {
	if o != nil && !IsNil(o.PageNumber) {
		return true
	}

	return false
}

// SetPageNumber gets a reference to the given int32 and assigns it to the PageNumber field.
func (o *RunResults) SetPageNumber(v int32) {
	o.PageNumber = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *RunResults) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunResults) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *RunResults) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *RunResults) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetNumPages returns the NumPages field value if set, zero value otherwise.
func (o *RunResults) GetNumPages() int32 {
	if o == nil || IsNil(o.NumPages) {
		var ret int32
		return ret
	}
	return *o.NumPages
}

// GetNumPagesOk returns a tuple with the NumPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunResults) GetNumPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.NumPages) {
		return nil, false
	}
	return o.NumPages, true
}

// HasNumPages returns a boolean if a field has been set.
func (o *RunResults) HasNumPages() bool {
	if o != nil && !IsNil(o.NumPages) {
		return true
	}

	return false
}

// SetNumPages gets a reference to the given int32 and assigns it to the NumPages field.
func (o *RunResults) SetNumPages(v int32) {
	o.NumPages = &v
}

// GetAmountOfRuns returns the AmountOfRuns field value if set, zero value otherwise.
func (o *RunResults) GetAmountOfRuns() int32 {
	if o == nil || IsNil(o.AmountOfRuns) {
		var ret int32
		return ret
	}
	return *o.AmountOfRuns
}

// GetAmountOfRunsOk returns a tuple with the AmountOfRuns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunResults) GetAmountOfRunsOk() (*int32, bool) {
	if o == nil || IsNil(o.AmountOfRuns) {
		return nil, false
	}
	return o.AmountOfRuns, true
}

// HasAmountOfRuns returns a boolean if a field has been set.
func (o *RunResults) HasAmountOfRuns() bool {
	if o != nil && !IsNil(o.AmountOfRuns) {
		return true
	}

	return false
}

// SetAmountOfRuns gets a reference to the given int32 and assigns it to the AmountOfRuns field.
func (o *RunResults) SetAmountOfRuns(v int32) {
	o.AmountOfRuns = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *RunResults) GetNextCursor() string {
	if o == nil || IsNil(o.NextCursor) {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunResults) GetNextCursorOk() (*string, bool) {
	if o == nil || IsNil(o.NextCursor) {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *RunResults) HasNextCursor() bool {
	if o != nil && !IsNil(o.NextCursor) {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *RunResults) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetRuns returns the Runs field value if set, zero value otherwise.
func (o *RunResults) GetRuns() []Run {
	if o == nil || IsNil(o.Runs) {
		var ret []Run
		return ret
	}
	return o.Runs
}

// GetRunsOk returns a tuple with the Runs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RunResults) GetRunsOk() ([]Run, bool) {
	if o == nil || IsNil(o.Runs) {
		return nil, false
	}
	return o.Runs, true
}

// HasRuns returns a boolean if a field has been set.
func (o *RunResults) HasRuns() bool {
	if o != nil && !IsNil(o.Runs) {
		return true
	}

	return false
}

// SetRuns gets a reference to the given []Run and assigns it to the Runs field.
func (o *RunResults) SetRuns(v []Run) {
	o.Runs = v
}

func (o RunResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RunResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PageNumber) {
		toSerialize["pageNumber"] = o.PageNumber
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.NumPages) {
		toSerialize["numPages"] = o.NumPages
	}
	if !IsNil(o.AmountOfRuns) {
		toSerialize["amountOfRuns"] = o.AmountOfRuns
	}
	if !IsNil(o.NextCursor) {
		toSerialize["nextCursor"] = o.NextCursor
	}
	if !IsNil(o.Runs) {
		toSerialize["runs"] = o.Runs
	}
	return toSerialize, nil
}

type NullableRunResults struct {
	value *RunResults
	isSet bool
}

func (v NullableRunResults) Get() *RunResults {
	return v.value
}

func (v *NullableRunResults) Set(val *RunResults) {
	v.value = val
	v.isSet = true
}

func (v NullableRunResults) IsSet() bool {
	return v.isSet
}

func (v *NullableRunResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRunResults(val *RunResults) *NullableRunResults {
	return &NullableRunResults{value: val, isSet: true}
}

func (v NullableRunResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRunResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


