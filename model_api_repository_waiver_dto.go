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

// checks if the ApiRepositoryWaiverDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryWaiverDTO{}

// ApiRepositoryWaiverDTO struct for ApiRepositoryWaiverDTO
type ApiRepositoryWaiverDTO struct {
	Repository *ApiRepositoryDTO `json:"repository,omitempty"`
	Stages []ApiPolicyViolationStageDTO `json:"stages,omitempty"`
}

// NewApiRepositoryWaiverDTO instantiates a new ApiRepositoryWaiverDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryWaiverDTO() *ApiRepositoryWaiverDTO {
	this := ApiRepositoryWaiverDTO{}
	return &this
}

// NewApiRepositoryWaiverDTOWithDefaults instantiates a new ApiRepositoryWaiverDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryWaiverDTOWithDefaults() *ApiRepositoryWaiverDTO {
	this := ApiRepositoryWaiverDTO{}
	return &this
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *ApiRepositoryWaiverDTO) GetRepository() ApiRepositoryDTO {
	if o == nil || IsNil(o.Repository) {
		var ret ApiRepositoryDTO
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryWaiverDTO) GetRepositoryOk() (*ApiRepositoryDTO, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *ApiRepositoryWaiverDTO) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given ApiRepositoryDTO and assigns it to the Repository field.
func (o *ApiRepositoryWaiverDTO) SetRepository(v ApiRepositoryDTO) {
	o.Repository = &v
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *ApiRepositoryWaiverDTO) GetStages() []ApiPolicyViolationStageDTO {
	if o == nil || IsNil(o.Stages) {
		var ret []ApiPolicyViolationStageDTO
		return ret
	}
	return o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryWaiverDTO) GetStagesOk() ([]ApiPolicyViolationStageDTO, bool) {
	if o == nil || IsNil(o.Stages) {
		return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *ApiRepositoryWaiverDTO) HasStages() bool {
	if o != nil && !IsNil(o.Stages) {
		return true
	}

	return false
}

// SetStages gets a reference to the given []ApiPolicyViolationStageDTO and assigns it to the Stages field.
func (o *ApiRepositoryWaiverDTO) SetStages(v []ApiPolicyViolationStageDTO) {
	o.Stages = v
}

func (o ApiRepositoryWaiverDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryWaiverDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.Stages) {
		toSerialize["stages"] = o.Stages
	}
	return toSerialize, nil
}

type NullableApiRepositoryWaiverDTO struct {
	value *ApiRepositoryWaiverDTO
	isSet bool
}

func (v NullableApiRepositoryWaiverDTO) Get() *ApiRepositoryWaiverDTO {
	return v.value
}

func (v *NullableApiRepositoryWaiverDTO) Set(val *ApiRepositoryWaiverDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryWaiverDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryWaiverDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryWaiverDTO(val *ApiRepositoryWaiverDTO) *NullableApiRepositoryWaiverDTO {
	return &NullableApiRepositoryWaiverDTO{value: val, isSet: true}
}

func (v NullableApiRepositoryWaiverDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryWaiverDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


