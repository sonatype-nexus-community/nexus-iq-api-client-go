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

// checks if the ApiStaleEvaluationsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiStaleEvaluationsDTO{}

// ApiStaleEvaluationsDTO struct for ApiStaleEvaluationsDTO
type ApiStaleEvaluationsDTO struct {
	Applications []ApiStaleApplicationEvaluationDTO `json:"applications,omitempty"`
	Repositories []ApiStaleRepositoryEvaluationDTO `json:"repositories,omitempty"`
}

// NewApiStaleEvaluationsDTO instantiates a new ApiStaleEvaluationsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStaleEvaluationsDTO() *ApiStaleEvaluationsDTO {
	this := ApiStaleEvaluationsDTO{}
	return &this
}

// NewApiStaleEvaluationsDTOWithDefaults instantiates a new ApiStaleEvaluationsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStaleEvaluationsDTOWithDefaults() *ApiStaleEvaluationsDTO {
	this := ApiStaleEvaluationsDTO{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *ApiStaleEvaluationsDTO) GetApplications() []ApiStaleApplicationEvaluationDTO {
	if o == nil || IsNil(o.Applications) {
		var ret []ApiStaleApplicationEvaluationDTO
		return ret
	}
	return o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleEvaluationsDTO) GetApplicationsOk() ([]ApiStaleApplicationEvaluationDTO, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *ApiStaleEvaluationsDTO) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given []ApiStaleApplicationEvaluationDTO and assigns it to the Applications field.
func (o *ApiStaleEvaluationsDTO) SetApplications(v []ApiStaleApplicationEvaluationDTO) {
	o.Applications = v
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *ApiStaleEvaluationsDTO) GetRepositories() []ApiStaleRepositoryEvaluationDTO {
	if o == nil || IsNil(o.Repositories) {
		var ret []ApiStaleRepositoryEvaluationDTO
		return ret
	}
	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleEvaluationsDTO) GetRepositoriesOk() ([]ApiStaleRepositoryEvaluationDTO, bool) {
	if o == nil || IsNil(o.Repositories) {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *ApiStaleEvaluationsDTO) HasRepositories() bool {
	if o != nil && !IsNil(o.Repositories) {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given []ApiStaleRepositoryEvaluationDTO and assigns it to the Repositories field.
func (o *ApiStaleEvaluationsDTO) SetRepositories(v []ApiStaleRepositoryEvaluationDTO) {
	o.Repositories = v
}

func (o ApiStaleEvaluationsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiStaleEvaluationsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	if !IsNil(o.Repositories) {
		toSerialize["repositories"] = o.Repositories
	}
	return toSerialize, nil
}

type NullableApiStaleEvaluationsDTO struct {
	value *ApiStaleEvaluationsDTO
	isSet bool
}

func (v NullableApiStaleEvaluationsDTO) Get() *ApiStaleEvaluationsDTO {
	return v.value
}

func (v *NullableApiStaleEvaluationsDTO) Set(val *ApiStaleEvaluationsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStaleEvaluationsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStaleEvaluationsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStaleEvaluationsDTO(val *ApiStaleEvaluationsDTO) *NullableApiStaleEvaluationsDTO {
	return &NullableApiStaleEvaluationsDTO{value: val, isSet: true}
}

func (v NullableApiStaleEvaluationsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStaleEvaluationsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

