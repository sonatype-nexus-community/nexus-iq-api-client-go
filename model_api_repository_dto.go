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

// checks if the ApiRepositoryDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryDTO{}

// ApiRepositoryDTO struct for ApiRepositoryDTO
type ApiRepositoryDTO struct {
	Format *string `json:"format,omitempty"`
	PublicId *string `json:"publicId,omitempty"`
	RepositoryId *string `json:"repositoryId,omitempty"`
}

// NewApiRepositoryDTO instantiates a new ApiRepositoryDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryDTO() *ApiRepositoryDTO {
	this := ApiRepositoryDTO{}
	return &this
}

// NewApiRepositoryDTOWithDefaults instantiates a new ApiRepositoryDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryDTOWithDefaults() *ApiRepositoryDTO {
	this := ApiRepositoryDTO{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ApiRepositoryDTO) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryDTO) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ApiRepositoryDTO) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ApiRepositoryDTO) SetFormat(v string) {
	o.Format = &v
}

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *ApiRepositoryDTO) GetPublicId() string {
	if o == nil || IsNil(o.PublicId) {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryDTO) GetPublicIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublicId) {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *ApiRepositoryDTO) HasPublicId() bool {
	if o != nil && !IsNil(o.PublicId) {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *ApiRepositoryDTO) SetPublicId(v string) {
	o.PublicId = &v
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *ApiRepositoryDTO) GetRepositoryId() string {
	if o == nil || IsNil(o.RepositoryId) {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryDTO) GetRepositoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryId) {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *ApiRepositoryDTO) HasRepositoryId() bool {
	if o != nil && !IsNil(o.RepositoryId) {
		return true
	}

	return false
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *ApiRepositoryDTO) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

func (o ApiRepositoryDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.PublicId) {
		toSerialize["publicId"] = o.PublicId
	}
	if !IsNil(o.RepositoryId) {
		toSerialize["repositoryId"] = o.RepositoryId
	}
	return toSerialize, nil
}

type NullableApiRepositoryDTO struct {
	value *ApiRepositoryDTO
	isSet bool
}

func (v NullableApiRepositoryDTO) Get() *ApiRepositoryDTO {
	return v.value
}

func (v *NullableApiRepositoryDTO) Set(val *ApiRepositoryDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryDTO(val *ApiRepositoryDTO) *NullableApiRepositoryDTO {
	return &NullableApiRepositoryDTO{value: val, isSet: true}
}

func (v NullableApiRepositoryDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

