/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.166.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiLicenseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicenseDTO{}

// ApiLicenseDTO struct for ApiLicenseDTO
type ApiLicenseDTO struct {
	LicenseId *string `json:"licenseId,omitempty"`
	LicenseName *string `json:"licenseName,omitempty"`
}

// NewApiLicenseDTO instantiates a new ApiLicenseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicenseDTO() *ApiLicenseDTO {
	this := ApiLicenseDTO{}
	return &this
}

// NewApiLicenseDTOWithDefaults instantiates a new ApiLicenseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicenseDTOWithDefaults() *ApiLicenseDTO {
	this := ApiLicenseDTO{}
	return &this
}

// GetLicenseId returns the LicenseId field value if set, zero value otherwise.
func (o *ApiLicenseDTO) GetLicenseId() string {
	if o == nil || IsNil(o.LicenseId) {
		var ret string
		return ret
	}
	return *o.LicenseId
}

// GetLicenseIdOk returns a tuple with the LicenseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDTO) GetLicenseIdOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseId) {
		return nil, false
	}
	return o.LicenseId, true
}

// HasLicenseId returns a boolean if a field has been set.
func (o *ApiLicenseDTO) HasLicenseId() bool {
	if o != nil && !IsNil(o.LicenseId) {
		return true
	}

	return false
}

// SetLicenseId gets a reference to the given string and assigns it to the LicenseId field.
func (o *ApiLicenseDTO) SetLicenseId(v string) {
	o.LicenseId = &v
}

// GetLicenseName returns the LicenseName field value if set, zero value otherwise.
func (o *ApiLicenseDTO) GetLicenseName() string {
	if o == nil || IsNil(o.LicenseName) {
		var ret string
		return ret
	}
	return *o.LicenseName
}

// GetLicenseNameOk returns a tuple with the LicenseName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDTO) GetLicenseNameOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseName) {
		return nil, false
	}
	return o.LicenseName, true
}

// HasLicenseName returns a boolean if a field has been set.
func (o *ApiLicenseDTO) HasLicenseName() bool {
	if o != nil && !IsNil(o.LicenseName) {
		return true
	}

	return false
}

// SetLicenseName gets a reference to the given string and assigns it to the LicenseName field.
func (o *ApiLicenseDTO) SetLicenseName(v string) {
	o.LicenseName = &v
}

func (o ApiLicenseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicenseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LicenseId) {
		toSerialize["licenseId"] = o.LicenseId
	}
	if !IsNil(o.LicenseName) {
		toSerialize["licenseName"] = o.LicenseName
	}
	return toSerialize, nil
}

type NullableApiLicenseDTO struct {
	value *ApiLicenseDTO
	isSet bool
}

func (v NullableApiLicenseDTO) Get() *ApiLicenseDTO {
	return v.value
}

func (v *NullableApiLicenseDTO) Set(val *ApiLicenseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicenseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicenseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicenseDTO(val *ApiLicenseDTO) *NullableApiLicenseDTO {
	return &NullableApiLicenseDTO{value: val, isSet: true}
}

func (v NullableApiLicenseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicenseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


