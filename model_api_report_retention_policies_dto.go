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

// checks if the ApiReportRetentionPoliciesDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReportRetentionPoliciesDTO{}

// ApiReportRetentionPoliciesDTO struct for ApiReportRetentionPoliciesDTO
type ApiReportRetentionPoliciesDTO struct {
	Stages *map[string]ApiReportRetentionPolicyDTO `json:"stages,omitempty"`
}

// NewApiReportRetentionPoliciesDTO instantiates a new ApiReportRetentionPoliciesDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReportRetentionPoliciesDTO() *ApiReportRetentionPoliciesDTO {
	this := ApiReportRetentionPoliciesDTO{}
	return &this
}

// NewApiReportRetentionPoliciesDTOWithDefaults instantiates a new ApiReportRetentionPoliciesDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReportRetentionPoliciesDTOWithDefaults() *ApiReportRetentionPoliciesDTO {
	this := ApiReportRetentionPoliciesDTO{}
	return &this
}

// GetStages returns the Stages field value if set, zero value otherwise.
func (o *ApiReportRetentionPoliciesDTO) GetStages() map[string]ApiReportRetentionPolicyDTO {
	if o == nil || IsNil(o.Stages) {
		var ret map[string]ApiReportRetentionPolicyDTO
		return ret
	}
	return *o.Stages
}

// GetStagesOk returns a tuple with the Stages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportRetentionPoliciesDTO) GetStagesOk() (*map[string]ApiReportRetentionPolicyDTO, bool) {
	if o == nil || IsNil(o.Stages) {
		return nil, false
	}
	return o.Stages, true
}

// HasStages returns a boolean if a field has been set.
func (o *ApiReportRetentionPoliciesDTO) HasStages() bool {
	if o != nil && !IsNil(o.Stages) {
		return true
	}

	return false
}

// SetStages gets a reference to the given map[string]ApiReportRetentionPolicyDTO and assigns it to the Stages field.
func (o *ApiReportRetentionPoliciesDTO) SetStages(v map[string]ApiReportRetentionPolicyDTO) {
	o.Stages = &v
}

func (o ApiReportRetentionPoliciesDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReportRetentionPoliciesDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Stages) {
		toSerialize["stages"] = o.Stages
	}
	return toSerialize, nil
}

type NullableApiReportRetentionPoliciesDTO struct {
	value *ApiReportRetentionPoliciesDTO
	isSet bool
}

func (v NullableApiReportRetentionPoliciesDTO) Get() *ApiReportRetentionPoliciesDTO {
	return v.value
}

func (v *NullableApiReportRetentionPoliciesDTO) Set(val *ApiReportRetentionPoliciesDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReportRetentionPoliciesDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReportRetentionPoliciesDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReportRetentionPoliciesDTO(val *ApiReportRetentionPoliciesDTO) *NullableApiReportRetentionPoliciesDTO {
	return &NullableApiReportRetentionPoliciesDTO{value: val, isSet: true}
}

func (v NullableApiReportRetentionPoliciesDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReportRetentionPoliciesDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


