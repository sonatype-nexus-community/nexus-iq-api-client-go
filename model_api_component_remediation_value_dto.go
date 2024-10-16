/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.182.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiComponentRemediationValueDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentRemediationValueDTO{}

// ApiComponentRemediationValueDTO struct for ApiComponentRemediationValueDTO
type ApiComponentRemediationValueDTO struct {
	SuggestedVersionChange *ApiSuggestedVersionChangeOptionDTO `json:"suggestedVersionChange,omitempty"`
	VersionChanges []ApiVersionChangeOptionDTO `json:"versionChanges,omitempty"`
}

// NewApiComponentRemediationValueDTO instantiates a new ApiComponentRemediationValueDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentRemediationValueDTO() *ApiComponentRemediationValueDTO {
	this := ApiComponentRemediationValueDTO{}
	return &this
}

// NewApiComponentRemediationValueDTOWithDefaults instantiates a new ApiComponentRemediationValueDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentRemediationValueDTOWithDefaults() *ApiComponentRemediationValueDTO {
	this := ApiComponentRemediationValueDTO{}
	return &this
}

// GetSuggestedVersionChange returns the SuggestedVersionChange field value if set, zero value otherwise.
func (o *ApiComponentRemediationValueDTO) GetSuggestedVersionChange() ApiSuggestedVersionChangeOptionDTO {
	if o == nil || IsNil(o.SuggestedVersionChange) {
		var ret ApiSuggestedVersionChangeOptionDTO
		return ret
	}
	return *o.SuggestedVersionChange
}

// GetSuggestedVersionChangeOk returns a tuple with the SuggestedVersionChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentRemediationValueDTO) GetSuggestedVersionChangeOk() (*ApiSuggestedVersionChangeOptionDTO, bool) {
	if o == nil || IsNil(o.SuggestedVersionChange) {
		return nil, false
	}
	return o.SuggestedVersionChange, true
}

// HasSuggestedVersionChange returns a boolean if a field has been set.
func (o *ApiComponentRemediationValueDTO) HasSuggestedVersionChange() bool {
	if o != nil && !IsNil(o.SuggestedVersionChange) {
		return true
	}

	return false
}

// SetSuggestedVersionChange gets a reference to the given ApiSuggestedVersionChangeOptionDTO and assigns it to the SuggestedVersionChange field.
func (o *ApiComponentRemediationValueDTO) SetSuggestedVersionChange(v ApiSuggestedVersionChangeOptionDTO) {
	o.SuggestedVersionChange = &v
}

// GetVersionChanges returns the VersionChanges field value if set, zero value otherwise.
func (o *ApiComponentRemediationValueDTO) GetVersionChanges() []ApiVersionChangeOptionDTO {
	if o == nil || IsNil(o.VersionChanges) {
		var ret []ApiVersionChangeOptionDTO
		return ret
	}
	return o.VersionChanges
}

// GetVersionChangesOk returns a tuple with the VersionChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentRemediationValueDTO) GetVersionChangesOk() ([]ApiVersionChangeOptionDTO, bool) {
	if o == nil || IsNil(o.VersionChanges) {
		return nil, false
	}
	return o.VersionChanges, true
}

// HasVersionChanges returns a boolean if a field has been set.
func (o *ApiComponentRemediationValueDTO) HasVersionChanges() bool {
	if o != nil && !IsNil(o.VersionChanges) {
		return true
	}

	return false
}

// SetVersionChanges gets a reference to the given []ApiVersionChangeOptionDTO and assigns it to the VersionChanges field.
func (o *ApiComponentRemediationValueDTO) SetVersionChanges(v []ApiVersionChangeOptionDTO) {
	o.VersionChanges = v
}

func (o ApiComponentRemediationValueDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentRemediationValueDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuggestedVersionChange) {
		toSerialize["suggestedVersionChange"] = o.SuggestedVersionChange
	}
	if !IsNil(o.VersionChanges) {
		toSerialize["versionChanges"] = o.VersionChanges
	}
	return toSerialize, nil
}

type NullableApiComponentRemediationValueDTO struct {
	value *ApiComponentRemediationValueDTO
	isSet bool
}

func (v NullableApiComponentRemediationValueDTO) Get() *ApiComponentRemediationValueDTO {
	return v.value
}

func (v *NullableApiComponentRemediationValueDTO) Set(val *ApiComponentRemediationValueDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentRemediationValueDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentRemediationValueDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentRemediationValueDTO(val *ApiComponentRemediationValueDTO) *NullableApiComponentRemediationValueDTO {
	return &NullableApiComponentRemediationValueDTO{value: val, isSet: true}
}

func (v NullableApiComponentRemediationValueDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentRemediationValueDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


