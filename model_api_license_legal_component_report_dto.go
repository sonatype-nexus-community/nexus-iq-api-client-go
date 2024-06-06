/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.177.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiLicenseLegalComponentReportDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicenseLegalComponentReportDTO{}

// ApiLicenseLegalComponentReportDTO struct for ApiLicenseLegalComponentReportDTO
type ApiLicenseLegalComponentReportDTO struct {
	Component *ApiLicenseLegalComponentDTO `json:"component,omitempty"`
	LicenseLegalMetadata []ApiLicenseLegalMetadataDTO `json:"licenseLegalMetadata,omitempty"`
}

// NewApiLicenseLegalComponentReportDTO instantiates a new ApiLicenseLegalComponentReportDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicenseLegalComponentReportDTO() *ApiLicenseLegalComponentReportDTO {
	this := ApiLicenseLegalComponentReportDTO{}
	return &this
}

// NewApiLicenseLegalComponentReportDTOWithDefaults instantiates a new ApiLicenseLegalComponentReportDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicenseLegalComponentReportDTOWithDefaults() *ApiLicenseLegalComponentReportDTO {
	this := ApiLicenseLegalComponentReportDTO{}
	return &this
}

// GetComponent returns the Component field value if set, zero value otherwise.
func (o *ApiLicenseLegalComponentReportDTO) GetComponent() ApiLicenseLegalComponentDTO {
	if o == nil || IsNil(o.Component) {
		var ret ApiLicenseLegalComponentDTO
		return ret
	}
	return *o.Component
}

// GetComponentOk returns a tuple with the Component field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalComponentReportDTO) GetComponentOk() (*ApiLicenseLegalComponentDTO, bool) {
	if o == nil || IsNil(o.Component) {
		return nil, false
	}
	return o.Component, true
}

// HasComponent returns a boolean if a field has been set.
func (o *ApiLicenseLegalComponentReportDTO) HasComponent() bool {
	if o != nil && !IsNil(o.Component) {
		return true
	}

	return false
}

// SetComponent gets a reference to the given ApiLicenseLegalComponentDTO and assigns it to the Component field.
func (o *ApiLicenseLegalComponentReportDTO) SetComponent(v ApiLicenseLegalComponentDTO) {
	o.Component = &v
}

// GetLicenseLegalMetadata returns the LicenseLegalMetadata field value if set, zero value otherwise.
func (o *ApiLicenseLegalComponentReportDTO) GetLicenseLegalMetadata() []ApiLicenseLegalMetadataDTO {
	if o == nil || IsNil(o.LicenseLegalMetadata) {
		var ret []ApiLicenseLegalMetadataDTO
		return ret
	}
	return o.LicenseLegalMetadata
}

// GetLicenseLegalMetadataOk returns a tuple with the LicenseLegalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalComponentReportDTO) GetLicenseLegalMetadataOk() ([]ApiLicenseLegalMetadataDTO, bool) {
	if o == nil || IsNil(o.LicenseLegalMetadata) {
		return nil, false
	}
	return o.LicenseLegalMetadata, true
}

// HasLicenseLegalMetadata returns a boolean if a field has been set.
func (o *ApiLicenseLegalComponentReportDTO) HasLicenseLegalMetadata() bool {
	if o != nil && !IsNil(o.LicenseLegalMetadata) {
		return true
	}

	return false
}

// SetLicenseLegalMetadata gets a reference to the given []ApiLicenseLegalMetadataDTO and assigns it to the LicenseLegalMetadata field.
func (o *ApiLicenseLegalComponentReportDTO) SetLicenseLegalMetadata(v []ApiLicenseLegalMetadataDTO) {
	o.LicenseLegalMetadata = v
}

func (o ApiLicenseLegalComponentReportDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicenseLegalComponentReportDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Component) {
		toSerialize["component"] = o.Component
	}
	if !IsNil(o.LicenseLegalMetadata) {
		toSerialize["licenseLegalMetadata"] = o.LicenseLegalMetadata
	}
	return toSerialize, nil
}

type NullableApiLicenseLegalComponentReportDTO struct {
	value *ApiLicenseLegalComponentReportDTO
	isSet bool
}

func (v NullableApiLicenseLegalComponentReportDTO) Get() *ApiLicenseLegalComponentReportDTO {
	return v.value
}

func (v *NullableApiLicenseLegalComponentReportDTO) Set(val *ApiLicenseLegalComponentReportDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicenseLegalComponentReportDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicenseLegalComponentReportDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicenseLegalComponentReportDTO(val *ApiLicenseLegalComponentReportDTO) *NullableApiLicenseLegalComponentReportDTO {
	return &NullableApiLicenseLegalComponentReportDTO{value: val, isSet: true}
}

func (v NullableApiLicenseLegalComponentReportDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicenseLegalComponentReportDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


