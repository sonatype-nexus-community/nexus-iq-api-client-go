/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.166.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiComponentsInQuarantineDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentsInQuarantineDTO{}

// ApiComponentsInQuarantineDTO struct for ApiComponentsInQuarantineDTO
type ApiComponentsInQuarantineDTO struct {
	ComponentsInQuarantine []ApiRepositoryComponentsInQuarantineDTO `json:"componentsInQuarantine,omitempty"`
}

// NewApiComponentsInQuarantineDTO instantiates a new ApiComponentsInQuarantineDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentsInQuarantineDTO() *ApiComponentsInQuarantineDTO {
	this := ApiComponentsInQuarantineDTO{}
	return &this
}

// NewApiComponentsInQuarantineDTOWithDefaults instantiates a new ApiComponentsInQuarantineDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentsInQuarantineDTOWithDefaults() *ApiComponentsInQuarantineDTO {
	this := ApiComponentsInQuarantineDTO{}
	return &this
}

// GetComponentsInQuarantine returns the ComponentsInQuarantine field value if set, zero value otherwise.
func (o *ApiComponentsInQuarantineDTO) GetComponentsInQuarantine() []ApiRepositoryComponentsInQuarantineDTO {
	if o == nil || IsNil(o.ComponentsInQuarantine) {
		var ret []ApiRepositoryComponentsInQuarantineDTO
		return ret
	}
	return o.ComponentsInQuarantine
}

// GetComponentsInQuarantineOk returns a tuple with the ComponentsInQuarantine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentsInQuarantineDTO) GetComponentsInQuarantineOk() ([]ApiRepositoryComponentsInQuarantineDTO, bool) {
	if o == nil || IsNil(o.ComponentsInQuarantine) {
		return nil, false
	}
	return o.ComponentsInQuarantine, true
}

// HasComponentsInQuarantine returns a boolean if a field has been set.
func (o *ApiComponentsInQuarantineDTO) HasComponentsInQuarantine() bool {
	if o != nil && !IsNil(o.ComponentsInQuarantine) {
		return true
	}

	return false
}

// SetComponentsInQuarantine gets a reference to the given []ApiRepositoryComponentsInQuarantineDTO and assigns it to the ComponentsInQuarantine field.
func (o *ApiComponentsInQuarantineDTO) SetComponentsInQuarantine(v []ApiRepositoryComponentsInQuarantineDTO) {
	o.ComponentsInQuarantine = v
}

func (o ApiComponentsInQuarantineDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentsInQuarantineDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentsInQuarantine) {
		toSerialize["componentsInQuarantine"] = o.ComponentsInQuarantine
	}
	return toSerialize, nil
}

type NullableApiComponentsInQuarantineDTO struct {
	value *ApiComponentsInQuarantineDTO
	isSet bool
}

func (v NullableApiComponentsInQuarantineDTO) Get() *ApiComponentsInQuarantineDTO {
	return v.value
}

func (v *NullableApiComponentsInQuarantineDTO) Set(val *ApiComponentsInQuarantineDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentsInQuarantineDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentsInQuarantineDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentsInQuarantineDTO(val *ApiComponentsInQuarantineDTO) *NullableApiComponentsInQuarantineDTO {
	return &NullableApiComponentsInQuarantineDTO{value: val, isSet: true}
}

func (v NullableApiComponentsInQuarantineDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentsInQuarantineDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


