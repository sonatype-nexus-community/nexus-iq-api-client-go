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

// checks if the IdeUsersOverviewDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdeUsersOverviewDTO{}

// IdeUsersOverviewDTO struct for IdeUsersOverviewDTO
type IdeUsersOverviewDTO struct {
	UserCount *int64 `json:"userCount,omitempty"`
}

// NewIdeUsersOverviewDTO instantiates a new IdeUsersOverviewDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdeUsersOverviewDTO() *IdeUsersOverviewDTO {
	this := IdeUsersOverviewDTO{}
	return &this
}

// NewIdeUsersOverviewDTOWithDefaults instantiates a new IdeUsersOverviewDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdeUsersOverviewDTOWithDefaults() *IdeUsersOverviewDTO {
	this := IdeUsersOverviewDTO{}
	return &this
}

// GetUserCount returns the UserCount field value if set, zero value otherwise.
func (o *IdeUsersOverviewDTO) GetUserCount() int64 {
	if o == nil || IsNil(o.UserCount) {
		var ret int64
		return ret
	}
	return *o.UserCount
}

// GetUserCountOk returns a tuple with the UserCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdeUsersOverviewDTO) GetUserCountOk() (*int64, bool) {
	if o == nil || IsNil(o.UserCount) {
		return nil, false
	}
	return o.UserCount, true
}

// HasUserCount returns a boolean if a field has been set.
func (o *IdeUsersOverviewDTO) HasUserCount() bool {
	if o != nil && !IsNil(o.UserCount) {
		return true
	}

	return false
}

// SetUserCount gets a reference to the given int64 and assigns it to the UserCount field.
func (o *IdeUsersOverviewDTO) SetUserCount(v int64) {
	o.UserCount = &v
}

func (o IdeUsersOverviewDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdeUsersOverviewDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserCount) {
		toSerialize["userCount"] = o.UserCount
	}
	return toSerialize, nil
}

type NullableIdeUsersOverviewDTO struct {
	value *IdeUsersOverviewDTO
	isSet bool
}

func (v NullableIdeUsersOverviewDTO) Get() *IdeUsersOverviewDTO {
	return v.value
}

func (v *NullableIdeUsersOverviewDTO) Set(val *IdeUsersOverviewDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableIdeUsersOverviewDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableIdeUsersOverviewDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdeUsersOverviewDTO(val *IdeUsersOverviewDTO) *NullableIdeUsersOverviewDTO {
	return &NullableIdeUsersOverviewDTO{value: val, isSet: true}
}

func (v NullableIdeUsersOverviewDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdeUsersOverviewDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


