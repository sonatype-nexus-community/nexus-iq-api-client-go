/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 164
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiLicenseThreatDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicenseThreatDTOV2{}

// ApiLicenseThreatDTOV2 struct for ApiLicenseThreatDTOV2
type ApiLicenseThreatDTOV2 struct {
	LicenseThreatGroupCategory *string `json:"licenseThreatGroupCategory,omitempty"`
	LicenseThreatGroupLevel *int32 `json:"licenseThreatGroupLevel,omitempty"`
	LicenseThreatGroupName *string `json:"licenseThreatGroupName,omitempty"`
}

// NewApiLicenseThreatDTOV2 instantiates a new ApiLicenseThreatDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicenseThreatDTOV2() *ApiLicenseThreatDTOV2 {
	this := ApiLicenseThreatDTOV2{}
	return &this
}

// NewApiLicenseThreatDTOV2WithDefaults instantiates a new ApiLicenseThreatDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicenseThreatDTOV2WithDefaults() *ApiLicenseThreatDTOV2 {
	this := ApiLicenseThreatDTOV2{}
	return &this
}

// GetLicenseThreatGroupCategory returns the LicenseThreatGroupCategory field value if set, zero value otherwise.
func (o *ApiLicenseThreatDTOV2) GetLicenseThreatGroupCategory() string {
	if o == nil || IsNil(o.LicenseThreatGroupCategory) {
		var ret string
		return ret
	}
	return *o.LicenseThreatGroupCategory
}

// GetLicenseThreatGroupCategoryOk returns a tuple with the LicenseThreatGroupCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseThreatDTOV2) GetLicenseThreatGroupCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseThreatGroupCategory) {
		return nil, false
	}
	return o.LicenseThreatGroupCategory, true
}

// HasLicenseThreatGroupCategory returns a boolean if a field has been set.
func (o *ApiLicenseThreatDTOV2) HasLicenseThreatGroupCategory() bool {
	if o != nil && !IsNil(o.LicenseThreatGroupCategory) {
		return true
	}

	return false
}

// SetLicenseThreatGroupCategory gets a reference to the given string and assigns it to the LicenseThreatGroupCategory field.
func (o *ApiLicenseThreatDTOV2) SetLicenseThreatGroupCategory(v string) {
	o.LicenseThreatGroupCategory = &v
}

// GetLicenseThreatGroupLevel returns the LicenseThreatGroupLevel field value if set, zero value otherwise.
func (o *ApiLicenseThreatDTOV2) GetLicenseThreatGroupLevel() int32 {
	if o == nil || IsNil(o.LicenseThreatGroupLevel) {
		var ret int32
		return ret
	}
	return *o.LicenseThreatGroupLevel
}

// GetLicenseThreatGroupLevelOk returns a tuple with the LicenseThreatGroupLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseThreatDTOV2) GetLicenseThreatGroupLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.LicenseThreatGroupLevel) {
		return nil, false
	}
	return o.LicenseThreatGroupLevel, true
}

// HasLicenseThreatGroupLevel returns a boolean if a field has been set.
func (o *ApiLicenseThreatDTOV2) HasLicenseThreatGroupLevel() bool {
	if o != nil && !IsNil(o.LicenseThreatGroupLevel) {
		return true
	}

	return false
}

// SetLicenseThreatGroupLevel gets a reference to the given int32 and assigns it to the LicenseThreatGroupLevel field.
func (o *ApiLicenseThreatDTOV2) SetLicenseThreatGroupLevel(v int32) {
	o.LicenseThreatGroupLevel = &v
}

// GetLicenseThreatGroupName returns the LicenseThreatGroupName field value if set, zero value otherwise.
func (o *ApiLicenseThreatDTOV2) GetLicenseThreatGroupName() string {
	if o == nil || IsNil(o.LicenseThreatGroupName) {
		var ret string
		return ret
	}
	return *o.LicenseThreatGroupName
}

// GetLicenseThreatGroupNameOk returns a tuple with the LicenseThreatGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseThreatDTOV2) GetLicenseThreatGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseThreatGroupName) {
		return nil, false
	}
	return o.LicenseThreatGroupName, true
}

// HasLicenseThreatGroupName returns a boolean if a field has been set.
func (o *ApiLicenseThreatDTOV2) HasLicenseThreatGroupName() bool {
	if o != nil && !IsNil(o.LicenseThreatGroupName) {
		return true
	}

	return false
}

// SetLicenseThreatGroupName gets a reference to the given string and assigns it to the LicenseThreatGroupName field.
func (o *ApiLicenseThreatDTOV2) SetLicenseThreatGroupName(v string) {
	o.LicenseThreatGroupName = &v
}

func (o ApiLicenseThreatDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicenseThreatDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LicenseThreatGroupCategory) {
		toSerialize["licenseThreatGroupCategory"] = o.LicenseThreatGroupCategory
	}
	if !IsNil(o.LicenseThreatGroupLevel) {
		toSerialize["licenseThreatGroupLevel"] = o.LicenseThreatGroupLevel
	}
	if !IsNil(o.LicenseThreatGroupName) {
		toSerialize["licenseThreatGroupName"] = o.LicenseThreatGroupName
	}
	return toSerialize, nil
}

type NullableApiLicenseThreatDTOV2 struct {
	value *ApiLicenseThreatDTOV2
	isSet bool
}

func (v NullableApiLicenseThreatDTOV2) Get() *ApiLicenseThreatDTOV2 {
	return v.value
}

func (v *NullableApiLicenseThreatDTOV2) Set(val *ApiLicenseThreatDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicenseThreatDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicenseThreatDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicenseThreatDTOV2(val *ApiLicenseThreatDTOV2) *NullableApiLicenseThreatDTOV2 {
	return &NullableApiLicenseThreatDTOV2{value: val, isSet: true}
}

func (v NullableApiLicenseThreatDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicenseThreatDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


