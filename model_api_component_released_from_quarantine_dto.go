/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 165
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiComponentReleasedFromQuarantineDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentReleasedFromQuarantineDTO{}

// ApiComponentReleasedFromQuarantineDTO struct for ApiComponentReleasedFromQuarantineDTO
type ApiComponentReleasedFromQuarantineDTO struct {
	ComponentReleasedFromQuarantine *ApiRepositoryComponentPolicyViolationDTO `json:"componentReleasedFromQuarantine,omitempty"`
}

// NewApiComponentReleasedFromQuarantineDTO instantiates a new ApiComponentReleasedFromQuarantineDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentReleasedFromQuarantineDTO() *ApiComponentReleasedFromQuarantineDTO {
	this := ApiComponentReleasedFromQuarantineDTO{}
	return &this
}

// NewApiComponentReleasedFromQuarantineDTOWithDefaults instantiates a new ApiComponentReleasedFromQuarantineDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentReleasedFromQuarantineDTOWithDefaults() *ApiComponentReleasedFromQuarantineDTO {
	this := ApiComponentReleasedFromQuarantineDTO{}
	return &this
}

// GetComponentReleasedFromQuarantine returns the ComponentReleasedFromQuarantine field value if set, zero value otherwise.
func (o *ApiComponentReleasedFromQuarantineDTO) GetComponentReleasedFromQuarantine() ApiRepositoryComponentPolicyViolationDTO {
	if o == nil || IsNil(o.ComponentReleasedFromQuarantine) {
		var ret ApiRepositoryComponentPolicyViolationDTO
		return ret
	}
	return *o.ComponentReleasedFromQuarantine
}

// GetComponentReleasedFromQuarantineOk returns a tuple with the ComponentReleasedFromQuarantine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentReleasedFromQuarantineDTO) GetComponentReleasedFromQuarantineOk() (*ApiRepositoryComponentPolicyViolationDTO, bool) {
	if o == nil || IsNil(o.ComponentReleasedFromQuarantine) {
		return nil, false
	}
	return o.ComponentReleasedFromQuarantine, true
}

// HasComponentReleasedFromQuarantine returns a boolean if a field has been set.
func (o *ApiComponentReleasedFromQuarantineDTO) HasComponentReleasedFromQuarantine() bool {
	if o != nil && !IsNil(o.ComponentReleasedFromQuarantine) {
		return true
	}

	return false
}

// SetComponentReleasedFromQuarantine gets a reference to the given ApiRepositoryComponentPolicyViolationDTO and assigns it to the ComponentReleasedFromQuarantine field.
func (o *ApiComponentReleasedFromQuarantineDTO) SetComponentReleasedFromQuarantine(v ApiRepositoryComponentPolicyViolationDTO) {
	o.ComponentReleasedFromQuarantine = &v
}

func (o ApiComponentReleasedFromQuarantineDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentReleasedFromQuarantineDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentReleasedFromQuarantine) {
		toSerialize["componentReleasedFromQuarantine"] = o.ComponentReleasedFromQuarantine
	}
	return toSerialize, nil
}

type NullableApiComponentReleasedFromQuarantineDTO struct {
	value *ApiComponentReleasedFromQuarantineDTO
	isSet bool
}

func (v NullableApiComponentReleasedFromQuarantineDTO) Get() *ApiComponentReleasedFromQuarantineDTO {
	return v.value
}

func (v *NullableApiComponentReleasedFromQuarantineDTO) Set(val *ApiComponentReleasedFromQuarantineDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentReleasedFromQuarantineDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentReleasedFromQuarantineDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentReleasedFromQuarantineDTO(val *ApiComponentReleasedFromQuarantineDTO) *NullableApiComponentReleasedFromQuarantineDTO {
	return &NullableApiComponentReleasedFromQuarantineDTO{value: val, isSet: true}
}

func (v NullableApiComponentReleasedFromQuarantineDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentReleasedFromQuarantineDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

