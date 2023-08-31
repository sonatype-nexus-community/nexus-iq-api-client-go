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

// checks if the LicenseObligationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseObligationDTO{}

// LicenseObligationDTO struct for LicenseObligationDTO
type LicenseObligationDTO struct {
	Name *string `json:"name,omitempty"`
	ObligationTexts []string `json:"obligationTexts,omitempty"`
}

// NewLicenseObligationDTO instantiates a new LicenseObligationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseObligationDTO() *LicenseObligationDTO {
	this := LicenseObligationDTO{}
	return &this
}

// NewLicenseObligationDTOWithDefaults instantiates a new LicenseObligationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseObligationDTOWithDefaults() *LicenseObligationDTO {
	this := LicenseObligationDTO{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LicenseObligationDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseObligationDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LicenseObligationDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LicenseObligationDTO) SetName(v string) {
	o.Name = &v
}

// GetObligationTexts returns the ObligationTexts field value if set, zero value otherwise.
func (o *LicenseObligationDTO) GetObligationTexts() []string {
	if o == nil || IsNil(o.ObligationTexts) {
		var ret []string
		return ret
	}
	return o.ObligationTexts
}

// GetObligationTextsOk returns a tuple with the ObligationTexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseObligationDTO) GetObligationTextsOk() ([]string, bool) {
	if o == nil || IsNil(o.ObligationTexts) {
		return nil, false
	}
	return o.ObligationTexts, true
}

// HasObligationTexts returns a boolean if a field has been set.
func (o *LicenseObligationDTO) HasObligationTexts() bool {
	if o != nil && !IsNil(o.ObligationTexts) {
		return true
	}

	return false
}

// SetObligationTexts gets a reference to the given []string and assigns it to the ObligationTexts field.
func (o *LicenseObligationDTO) SetObligationTexts(v []string) {
	o.ObligationTexts = v
}

func (o LicenseObligationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseObligationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ObligationTexts) {
		toSerialize["obligationTexts"] = o.ObligationTexts
	}
	return toSerialize, nil
}

type NullableLicenseObligationDTO struct {
	value *LicenseObligationDTO
	isSet bool
}

func (v NullableLicenseObligationDTO) Get() *LicenseObligationDTO {
	return v.value
}

func (v *NullableLicenseObligationDTO) Set(val *LicenseObligationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseObligationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseObligationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseObligationDTO(val *LicenseObligationDTO) *NullableLicenseObligationDTO {
	return &NullableLicenseObligationDTO{value: val, isSet: true}
}

func (v NullableLicenseObligationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseObligationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


