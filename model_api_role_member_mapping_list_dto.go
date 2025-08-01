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

// checks if the ApiRoleMemberMappingListDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRoleMemberMappingListDTO{}

// ApiRoleMemberMappingListDTO struct for ApiRoleMemberMappingListDTO
type ApiRoleMemberMappingListDTO struct {
	MemberMappings []ApiRoleMemberMappingDTO `json:"memberMappings,omitempty"`
}

// NewApiRoleMemberMappingListDTO instantiates a new ApiRoleMemberMappingListDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRoleMemberMappingListDTO() *ApiRoleMemberMappingListDTO {
	this := ApiRoleMemberMappingListDTO{}
	return &this
}

// NewApiRoleMemberMappingListDTOWithDefaults instantiates a new ApiRoleMemberMappingListDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRoleMemberMappingListDTOWithDefaults() *ApiRoleMemberMappingListDTO {
	this := ApiRoleMemberMappingListDTO{}
	return &this
}

// GetMemberMappings returns the MemberMappings field value if set, zero value otherwise.
func (o *ApiRoleMemberMappingListDTO) GetMemberMappings() []ApiRoleMemberMappingDTO {
	if o == nil || IsNil(o.MemberMappings) {
		var ret []ApiRoleMemberMappingDTO
		return ret
	}
	return o.MemberMappings
}

// GetMemberMappingsOk returns a tuple with the MemberMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleMemberMappingListDTO) GetMemberMappingsOk() ([]ApiRoleMemberMappingDTO, bool) {
	if o == nil || IsNil(o.MemberMappings) {
		return nil, false
	}
	return o.MemberMappings, true
}

// HasMemberMappings returns a boolean if a field has been set.
func (o *ApiRoleMemberMappingListDTO) HasMemberMappings() bool {
	if o != nil && !IsNil(o.MemberMappings) {
		return true
	}

	return false
}

// SetMemberMappings gets a reference to the given []ApiRoleMemberMappingDTO and assigns it to the MemberMappings field.
func (o *ApiRoleMemberMappingListDTO) SetMemberMappings(v []ApiRoleMemberMappingDTO) {
	o.MemberMappings = v
}

func (o ApiRoleMemberMappingListDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRoleMemberMappingListDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MemberMappings) {
		toSerialize["memberMappings"] = o.MemberMappings
	}
	return toSerialize, nil
}

type NullableApiRoleMemberMappingListDTO struct {
	value *ApiRoleMemberMappingListDTO
	isSet bool
}

func (v NullableApiRoleMemberMappingListDTO) Get() *ApiRoleMemberMappingListDTO {
	return v.value
}

func (v *NullableApiRoleMemberMappingListDTO) Set(val *ApiRoleMemberMappingListDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRoleMemberMappingListDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRoleMemberMappingListDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRoleMemberMappingListDTO(val *ApiRoleMemberMappingListDTO) *NullableApiRoleMemberMappingListDTO {
	return &NullableApiRoleMemberMappingListDTO{value: val, isSet: true}
}

func (v NullableApiRoleMemberMappingListDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRoleMemberMappingListDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


