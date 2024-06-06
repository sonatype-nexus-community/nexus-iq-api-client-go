/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.177.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiComponentRemediationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentRemediationDTO{}

// ApiComponentRemediationDTO struct for ApiComponentRemediationDTO
type ApiComponentRemediationDTO struct {
	Remediation *ApiComponentRemediationValueDTO `json:"remediation,omitempty"`
}

// NewApiComponentRemediationDTO instantiates a new ApiComponentRemediationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentRemediationDTO() *ApiComponentRemediationDTO {
	this := ApiComponentRemediationDTO{}
	return &this
}

// NewApiComponentRemediationDTOWithDefaults instantiates a new ApiComponentRemediationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentRemediationDTOWithDefaults() *ApiComponentRemediationDTO {
	this := ApiComponentRemediationDTO{}
	return &this
}

// GetRemediation returns the Remediation field value if set, zero value otherwise.
func (o *ApiComponentRemediationDTO) GetRemediation() ApiComponentRemediationValueDTO {
	if o == nil || IsNil(o.Remediation) {
		var ret ApiComponentRemediationValueDTO
		return ret
	}
	return *o.Remediation
}

// GetRemediationOk returns a tuple with the Remediation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentRemediationDTO) GetRemediationOk() (*ApiComponentRemediationValueDTO, bool) {
	if o == nil || IsNil(o.Remediation) {
		return nil, false
	}
	return o.Remediation, true
}

// HasRemediation returns a boolean if a field has been set.
func (o *ApiComponentRemediationDTO) HasRemediation() bool {
	if o != nil && !IsNil(o.Remediation) {
		return true
	}

	return false
}

// SetRemediation gets a reference to the given ApiComponentRemediationValueDTO and assigns it to the Remediation field.
func (o *ApiComponentRemediationDTO) SetRemediation(v ApiComponentRemediationValueDTO) {
	o.Remediation = &v
}

func (o ApiComponentRemediationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentRemediationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Remediation) {
		toSerialize["remediation"] = o.Remediation
	}
	return toSerialize, nil
}

type NullableApiComponentRemediationDTO struct {
	value *ApiComponentRemediationDTO
	isSet bool
}

func (v NullableApiComponentRemediationDTO) Get() *ApiComponentRemediationDTO {
	return v.value
}

func (v *NullableApiComponentRemediationDTO) Set(val *ApiComponentRemediationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentRemediationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentRemediationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentRemediationDTO(val *ApiComponentRemediationDTO) *NullableApiComponentRemediationDTO {
	return &NullableApiComponentRemediationDTO{value: val, isSet: true}
}

func (v NullableApiComponentRemediationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentRemediationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


