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

// checks if the ApiUserListDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUserListDTO{}

// ApiUserListDTO struct for ApiUserListDTO
type ApiUserListDTO struct {
	Users []ApiUserDTO `json:"users,omitempty"`
}

// NewApiUserListDTO instantiates a new ApiUserListDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUserListDTO() *ApiUserListDTO {
	this := ApiUserListDTO{}
	return &this
}

// NewApiUserListDTOWithDefaults instantiates a new ApiUserListDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUserListDTOWithDefaults() *ApiUserListDTO {
	this := ApiUserListDTO{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ApiUserListDTO) GetUsers() []ApiUserDTO {
	if o == nil || IsNil(o.Users) {
		var ret []ApiUserDTO
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserListDTO) GetUsersOk() ([]ApiUserDTO, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ApiUserListDTO) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []ApiUserDTO and assigns it to the Users field.
func (o *ApiUserListDTO) SetUsers(v []ApiUserDTO) {
	o.Users = v
}

func (o ApiUserListDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUserListDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableApiUserListDTO struct {
	value *ApiUserListDTO
	isSet bool
}

func (v NullableApiUserListDTO) Get() *ApiUserListDTO {
	return v.value
}

func (v *NullableApiUserListDTO) Set(val *ApiUserListDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUserListDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUserListDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUserListDTO(val *ApiUserListDTO) *NullableApiUserListDTO {
	return &NullableApiUserListDTO{value: val, isSet: true}
}

func (v NullableApiUserListDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUserListDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

