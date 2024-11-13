/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.184.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiRepositoryManagerListDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryManagerListDTO{}

// ApiRepositoryManagerListDTO struct for ApiRepositoryManagerListDTO
type ApiRepositoryManagerListDTO struct {
	RepositoryManagers []ApiRepositoryManagerDTO `json:"repositoryManagers,omitempty"`
}

// NewApiRepositoryManagerListDTO instantiates a new ApiRepositoryManagerListDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryManagerListDTO() *ApiRepositoryManagerListDTO {
	this := ApiRepositoryManagerListDTO{}
	return &this
}

// NewApiRepositoryManagerListDTOWithDefaults instantiates a new ApiRepositoryManagerListDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryManagerListDTOWithDefaults() *ApiRepositoryManagerListDTO {
	this := ApiRepositoryManagerListDTO{}
	return &this
}

// GetRepositoryManagers returns the RepositoryManagers field value if set, zero value otherwise.
func (o *ApiRepositoryManagerListDTO) GetRepositoryManagers() []ApiRepositoryManagerDTO {
	if o == nil || IsNil(o.RepositoryManagers) {
		var ret []ApiRepositoryManagerDTO
		return ret
	}
	return o.RepositoryManagers
}

// GetRepositoryManagersOk returns a tuple with the RepositoryManagers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryManagerListDTO) GetRepositoryManagersOk() ([]ApiRepositoryManagerDTO, bool) {
	if o == nil || IsNil(o.RepositoryManagers) {
		return nil, false
	}
	return o.RepositoryManagers, true
}

// HasRepositoryManagers returns a boolean if a field has been set.
func (o *ApiRepositoryManagerListDTO) HasRepositoryManagers() bool {
	if o != nil && !IsNil(o.RepositoryManagers) {
		return true
	}

	return false
}

// SetRepositoryManagers gets a reference to the given []ApiRepositoryManagerDTO and assigns it to the RepositoryManagers field.
func (o *ApiRepositoryManagerListDTO) SetRepositoryManagers(v []ApiRepositoryManagerDTO) {
	o.RepositoryManagers = v
}

func (o ApiRepositoryManagerListDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryManagerListDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RepositoryManagers) {
		toSerialize["repositoryManagers"] = o.RepositoryManagers
	}
	return toSerialize, nil
}

type NullableApiRepositoryManagerListDTO struct {
	value *ApiRepositoryManagerListDTO
	isSet bool
}

func (v NullableApiRepositoryManagerListDTO) Get() *ApiRepositoryManagerListDTO {
	return v.value
}

func (v *NullableApiRepositoryManagerListDTO) Set(val *ApiRepositoryManagerListDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryManagerListDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryManagerListDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryManagerListDTO(val *ApiRepositoryManagerListDTO) *NullableApiRepositoryManagerListDTO {
	return &NullableApiRepositoryManagerListDTO{value: val, isSet: true}
}

func (v NullableApiRepositoryManagerListDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryManagerListDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


