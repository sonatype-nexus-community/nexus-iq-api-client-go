/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.193.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiRepositoryListDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryListDTO{}

// ApiRepositoryListDTO struct for ApiRepositoryListDTO
type ApiRepositoryListDTO struct {
	Repositories []ApiRepositoryDTO `json:"repositories,omitempty"`
}

// NewApiRepositoryListDTO instantiates a new ApiRepositoryListDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryListDTO() *ApiRepositoryListDTO {
	this := ApiRepositoryListDTO{}
	return &this
}

// NewApiRepositoryListDTOWithDefaults instantiates a new ApiRepositoryListDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryListDTOWithDefaults() *ApiRepositoryListDTO {
	this := ApiRepositoryListDTO{}
	return &this
}

// GetRepositories returns the Repositories field value if set, zero value otherwise.
func (o *ApiRepositoryListDTO) GetRepositories() []ApiRepositoryDTO {
	if o == nil || IsNil(o.Repositories) {
		var ret []ApiRepositoryDTO
		return ret
	}
	return o.Repositories
}

// GetRepositoriesOk returns a tuple with the Repositories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryListDTO) GetRepositoriesOk() ([]ApiRepositoryDTO, bool) {
	if o == nil || IsNil(o.Repositories) {
		return nil, false
	}
	return o.Repositories, true
}

// HasRepositories returns a boolean if a field has been set.
func (o *ApiRepositoryListDTO) HasRepositories() bool {
	if o != nil && !IsNil(o.Repositories) {
		return true
	}

	return false
}

// SetRepositories gets a reference to the given []ApiRepositoryDTO and assigns it to the Repositories field.
func (o *ApiRepositoryListDTO) SetRepositories(v []ApiRepositoryDTO) {
	o.Repositories = v
}

func (o ApiRepositoryListDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryListDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Repositories) {
		toSerialize["repositories"] = o.Repositories
	}
	return toSerialize, nil
}

type NullableApiRepositoryListDTO struct {
	value *ApiRepositoryListDTO
	isSet bool
}

func (v NullableApiRepositoryListDTO) Get() *ApiRepositoryListDTO {
	return v.value
}

func (v *NullableApiRepositoryListDTO) Set(val *ApiRepositoryListDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryListDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryListDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryListDTO(val *ApiRepositoryListDTO) *NullableApiRepositoryListDTO {
	return &NullableApiRepositoryListDTO{value: val, isSet: true}
}

func (v NullableApiRepositoryListDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryListDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


