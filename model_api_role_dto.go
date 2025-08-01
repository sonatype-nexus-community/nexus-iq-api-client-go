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

// checks if the ApiRoleDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRoleDTO{}

// ApiRoleDTO struct for ApiRoleDTO
type ApiRoleDTO struct {
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewApiRoleDTO instantiates a new ApiRoleDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRoleDTO() *ApiRoleDTO {
	this := ApiRoleDTO{}
	return &this
}

// NewApiRoleDTOWithDefaults instantiates a new ApiRoleDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRoleDTOWithDefaults() *ApiRoleDTO {
	this := ApiRoleDTO{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiRoleDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiRoleDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiRoleDTO) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiRoleDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiRoleDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiRoleDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiRoleDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRoleDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiRoleDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiRoleDTO) SetName(v string) {
	o.Name = &v
}

func (o ApiRoleDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRoleDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableApiRoleDTO struct {
	value *ApiRoleDTO
	isSet bool
}

func (v NullableApiRoleDTO) Get() *ApiRoleDTO {
	return v.value
}

func (v *NullableApiRoleDTO) Set(val *ApiRoleDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRoleDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRoleDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRoleDTO(val *ApiRoleDTO) *NullableApiRoleDTO {
	return &NullableApiRoleDTO{value: val, isSet: true}
}

func (v NullableApiRoleDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRoleDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


