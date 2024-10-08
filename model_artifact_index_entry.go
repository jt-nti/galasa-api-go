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

// checks if the ArtifactIndexEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtifactIndexEntry{}

// ArtifactIndexEntry struct for ArtifactIndexEntry
type ArtifactIndexEntry struct {
	RunId *string `json:"runId,omitempty"`
	Path *string `json:"path,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewArtifactIndexEntry instantiates a new ArtifactIndexEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactIndexEntry() *ArtifactIndexEntry {
	this := ArtifactIndexEntry{}
	return &this
}

// NewArtifactIndexEntryWithDefaults instantiates a new ArtifactIndexEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactIndexEntryWithDefaults() *ArtifactIndexEntry {
	this := ArtifactIndexEntry{}
	return &this
}

// GetRunId returns the RunId field value if set, zero value otherwise.
func (o *ArtifactIndexEntry) GetRunId() string {
	if o == nil || IsNil(o.RunId) {
		var ret string
		return ret
	}
	return *o.RunId
}

// GetRunIdOk returns a tuple with the RunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactIndexEntry) GetRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.RunId) {
		return nil, false
	}
	return o.RunId, true
}

// HasRunId returns a boolean if a field has been set.
func (o *ArtifactIndexEntry) HasRunId() bool {
	if o != nil && !IsNil(o.RunId) {
		return true
	}

	return false
}

// SetRunId gets a reference to the given string and assigns it to the RunId field.
func (o *ArtifactIndexEntry) SetRunId(v string) {
	o.RunId = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ArtifactIndexEntry) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactIndexEntry) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ArtifactIndexEntry) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ArtifactIndexEntry) SetPath(v string) {
	o.Path = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ArtifactIndexEntry) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArtifactIndexEntry) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ArtifactIndexEntry) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ArtifactIndexEntry) SetUrl(v string) {
	o.Url = &v
}

func (o ArtifactIndexEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtifactIndexEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RunId) {
		toSerialize["runId"] = o.RunId
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableArtifactIndexEntry struct {
	value *ArtifactIndexEntry
	isSet bool
}

func (v NullableArtifactIndexEntry) Get() *ArtifactIndexEntry {
	return v.value
}

func (v *NullableArtifactIndexEntry) Set(val *ArtifactIndexEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactIndexEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactIndexEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactIndexEntry(val *ArtifactIndexEntry) *NullableArtifactIndexEntry {
	return &NullableArtifactIndexEntry{value: val, isSet: true}
}

func (v NullableArtifactIndexEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactIndexEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


