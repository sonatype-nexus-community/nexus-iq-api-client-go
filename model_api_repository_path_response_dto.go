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

// checks if the ApiRepositoryPathResponseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryPathResponseDTO{}

// ApiRepositoryPathResponseDTO struct for ApiRepositoryPathResponseDTO
type ApiRepositoryPathResponseDTO struct {
	PathVersions []ApiRepositoryPathVersions `json:"pathVersions,omitempty"`
}

// NewApiRepositoryPathResponseDTO instantiates a new ApiRepositoryPathResponseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryPathResponseDTO() *ApiRepositoryPathResponseDTO {
	this := ApiRepositoryPathResponseDTO{}
	return &this
}

// NewApiRepositoryPathResponseDTOWithDefaults instantiates a new ApiRepositoryPathResponseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryPathResponseDTOWithDefaults() *ApiRepositoryPathResponseDTO {
	this := ApiRepositoryPathResponseDTO{}
	return &this
}

// GetPathVersions returns the PathVersions field value if set, zero value otherwise.
func (o *ApiRepositoryPathResponseDTO) GetPathVersions() []ApiRepositoryPathVersions {
	if o == nil || IsNil(o.PathVersions) {
		var ret []ApiRepositoryPathVersions
		return ret
	}
	return o.PathVersions
}

// GetPathVersionsOk returns a tuple with the PathVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryPathResponseDTO) GetPathVersionsOk() ([]ApiRepositoryPathVersions, bool) {
	if o == nil || IsNil(o.PathVersions) {
		return nil, false
	}
	return o.PathVersions, true
}

// HasPathVersions returns a boolean if a field has been set.
func (o *ApiRepositoryPathResponseDTO) HasPathVersions() bool {
	if o != nil && !IsNil(o.PathVersions) {
		return true
	}

	return false
}

// SetPathVersions gets a reference to the given []ApiRepositoryPathVersions and assigns it to the PathVersions field.
func (o *ApiRepositoryPathResponseDTO) SetPathVersions(v []ApiRepositoryPathVersions) {
	o.PathVersions = v
}

func (o ApiRepositoryPathResponseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryPathResponseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PathVersions) {
		toSerialize["pathVersions"] = o.PathVersions
	}
	return toSerialize, nil
}

type NullableApiRepositoryPathResponseDTO struct {
	value *ApiRepositoryPathResponseDTO
	isSet bool
}

func (v NullableApiRepositoryPathResponseDTO) Get() *ApiRepositoryPathResponseDTO {
	return v.value
}

func (v *NullableApiRepositoryPathResponseDTO) Set(val *ApiRepositoryPathResponseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryPathResponseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryPathResponseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryPathResponseDTO(val *ApiRepositoryPathResponseDTO) *NullableApiRepositoryPathResponseDTO {
	return &NullableApiRepositoryPathResponseDTO{value: val, isSet: true}
}

func (v NullableApiRepositoryPathResponseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryPathResponseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

