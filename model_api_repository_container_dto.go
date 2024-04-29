/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.175.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiRepositoryContainerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryContainerDTO{}

// ApiRepositoryContainerDTO struct for ApiRepositoryContainerDTO
type ApiRepositoryContainerDTO struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewApiRepositoryContainerDTO instantiates a new ApiRepositoryContainerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryContainerDTO() *ApiRepositoryContainerDTO {
	this := ApiRepositoryContainerDTO{}
	return &this
}

// NewApiRepositoryContainerDTOWithDefaults instantiates a new ApiRepositoryContainerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryContainerDTOWithDefaults() *ApiRepositoryContainerDTO {
	this := ApiRepositoryContainerDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiRepositoryContainerDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryContainerDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiRepositoryContainerDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiRepositoryContainerDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiRepositoryContainerDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryContainerDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiRepositoryContainerDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiRepositoryContainerDTO) SetName(v string) {
	o.Name = &v
}

func (o ApiRepositoryContainerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryContainerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableApiRepositoryContainerDTO struct {
	value *ApiRepositoryContainerDTO
	isSet bool
}

func (v NullableApiRepositoryContainerDTO) Get() *ApiRepositoryContainerDTO {
	return v.value
}

func (v *NullableApiRepositoryContainerDTO) Set(val *ApiRepositoryContainerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryContainerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryContainerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryContainerDTO(val *ApiRepositoryContainerDTO) *NullableApiRepositoryContainerDTO {
	return &NullableApiRepositoryContainerDTO{value: val, isSet: true}
}

func (v NullableApiRepositoryContainerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryContainerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

