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

// checks if the ApiSecurityDataDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSecurityDataDTO{}

// ApiSecurityDataDTO struct for ApiSecurityDataDTO
type ApiSecurityDataDTO struct {
	SecurityIssues []ApiSecurityIssueDTO `json:"securityIssues,omitempty"`
}

// NewApiSecurityDataDTO instantiates a new ApiSecurityDataDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSecurityDataDTO() *ApiSecurityDataDTO {
	this := ApiSecurityDataDTO{}
	return &this
}

// NewApiSecurityDataDTOWithDefaults instantiates a new ApiSecurityDataDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSecurityDataDTOWithDefaults() *ApiSecurityDataDTO {
	this := ApiSecurityDataDTO{}
	return &this
}

// GetSecurityIssues returns the SecurityIssues field value if set, zero value otherwise.
func (o *ApiSecurityDataDTO) GetSecurityIssues() []ApiSecurityIssueDTO {
	if o == nil || IsNil(o.SecurityIssues) {
		var ret []ApiSecurityIssueDTO
		return ret
	}
	return o.SecurityIssues
}

// GetSecurityIssuesOk returns a tuple with the SecurityIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityDataDTO) GetSecurityIssuesOk() ([]ApiSecurityIssueDTO, bool) {
	if o == nil || IsNil(o.SecurityIssues) {
		return nil, false
	}
	return o.SecurityIssues, true
}

// HasSecurityIssues returns a boolean if a field has been set.
func (o *ApiSecurityDataDTO) HasSecurityIssues() bool {
	if o != nil && !IsNil(o.SecurityIssues) {
		return true
	}

	return false
}

// SetSecurityIssues gets a reference to the given []ApiSecurityIssueDTO and assigns it to the SecurityIssues field.
func (o *ApiSecurityDataDTO) SetSecurityIssues(v []ApiSecurityIssueDTO) {
	o.SecurityIssues = v
}

func (o ApiSecurityDataDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSecurityDataDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SecurityIssues) {
		toSerialize["securityIssues"] = o.SecurityIssues
	}
	return toSerialize, nil
}

type NullableApiSecurityDataDTO struct {
	value *ApiSecurityDataDTO
	isSet bool
}

func (v NullableApiSecurityDataDTO) Get() *ApiSecurityDataDTO {
	return v.value
}

func (v *NullableApiSecurityDataDTO) Set(val *ApiSecurityDataDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSecurityDataDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSecurityDataDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSecurityDataDTO(val *ApiSecurityDataDTO) *NullableApiSecurityDataDTO {
	return &NullableApiSecurityDataDTO{value: val, isSet: true}
}

func (v NullableApiSecurityDataDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSecurityDataDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


