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

// checks if the ApiLicensedSolutionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicensedSolutionDTO{}

// ApiLicensedSolutionDTO struct for ApiLicensedSolutionDTO
type ApiLicensedSolutionDTO struct {
	Id *string `json:"id,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewApiLicensedSolutionDTO instantiates a new ApiLicensedSolutionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicensedSolutionDTO() *ApiLicensedSolutionDTO {
	this := ApiLicensedSolutionDTO{}
	return &this
}

// NewApiLicensedSolutionDTOWithDefaults instantiates a new ApiLicensedSolutionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicensedSolutionDTOWithDefaults() *ApiLicensedSolutionDTO {
	this := ApiLicensedSolutionDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiLicensedSolutionDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicensedSolutionDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiLicensedSolutionDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiLicensedSolutionDTO) SetId(v string) {
	o.Id = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ApiLicensedSolutionDTO) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicensedSolutionDTO) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ApiLicensedSolutionDTO) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ApiLicensedSolutionDTO) SetUrl(v string) {
	o.Url = &v
}

func (o ApiLicensedSolutionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicensedSolutionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableApiLicensedSolutionDTO struct {
	value *ApiLicensedSolutionDTO
	isSet bool
}

func (v NullableApiLicensedSolutionDTO) Get() *ApiLicensedSolutionDTO {
	return v.value
}

func (v *NullableApiLicensedSolutionDTO) Set(val *ApiLicensedSolutionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicensedSolutionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicensedSolutionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicensedSolutionDTO(val *ApiLicensedSolutionDTO) *NullableApiLicensedSolutionDTO {
	return &NullableApiLicensedSolutionDTO{value: val, isSet: true}
}

func (v NullableApiLicensedSolutionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicensedSolutionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


