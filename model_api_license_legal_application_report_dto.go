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

// checks if the ApiLicenseLegalApplicationReportDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicenseLegalApplicationReportDTO{}

// ApiLicenseLegalApplicationReportDTO struct for ApiLicenseLegalApplicationReportDTO
type ApiLicenseLegalApplicationReportDTO struct {
	Components []ApiLicenseLegalComponentDTO `json:"components,omitempty"`
	LicenseLegalMetadata []ApiLicenseLegalMetadataDTO `json:"licenseLegalMetadata,omitempty"`
}

// NewApiLicenseLegalApplicationReportDTO instantiates a new ApiLicenseLegalApplicationReportDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicenseLegalApplicationReportDTO() *ApiLicenseLegalApplicationReportDTO {
	this := ApiLicenseLegalApplicationReportDTO{}
	return &this
}

// NewApiLicenseLegalApplicationReportDTOWithDefaults instantiates a new ApiLicenseLegalApplicationReportDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicenseLegalApplicationReportDTOWithDefaults() *ApiLicenseLegalApplicationReportDTO {
	this := ApiLicenseLegalApplicationReportDTO{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ApiLicenseLegalApplicationReportDTO) GetComponents() []ApiLicenseLegalComponentDTO {
	if o == nil || IsNil(o.Components) {
		var ret []ApiLicenseLegalComponentDTO
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalApplicationReportDTO) GetComponentsOk() ([]ApiLicenseLegalComponentDTO, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ApiLicenseLegalApplicationReportDTO) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []ApiLicenseLegalComponentDTO and assigns it to the Components field.
func (o *ApiLicenseLegalApplicationReportDTO) SetComponents(v []ApiLicenseLegalComponentDTO) {
	o.Components = v
}

// GetLicenseLegalMetadata returns the LicenseLegalMetadata field value if set, zero value otherwise.
func (o *ApiLicenseLegalApplicationReportDTO) GetLicenseLegalMetadata() []ApiLicenseLegalMetadataDTO {
	if o == nil || IsNil(o.LicenseLegalMetadata) {
		var ret []ApiLicenseLegalMetadataDTO
		return ret
	}
	return o.LicenseLegalMetadata
}

// GetLicenseLegalMetadataOk returns a tuple with the LicenseLegalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalApplicationReportDTO) GetLicenseLegalMetadataOk() ([]ApiLicenseLegalMetadataDTO, bool) {
	if o == nil || IsNil(o.LicenseLegalMetadata) {
		return nil, false
	}
	return o.LicenseLegalMetadata, true
}

// HasLicenseLegalMetadata returns a boolean if a field has been set.
func (o *ApiLicenseLegalApplicationReportDTO) HasLicenseLegalMetadata() bool {
	if o != nil && !IsNil(o.LicenseLegalMetadata) {
		return true
	}

	return false
}

// SetLicenseLegalMetadata gets a reference to the given []ApiLicenseLegalMetadataDTO and assigns it to the LicenseLegalMetadata field.
func (o *ApiLicenseLegalApplicationReportDTO) SetLicenseLegalMetadata(v []ApiLicenseLegalMetadataDTO) {
	o.LicenseLegalMetadata = v
}

func (o ApiLicenseLegalApplicationReportDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicenseLegalApplicationReportDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !IsNil(o.LicenseLegalMetadata) {
		toSerialize["licenseLegalMetadata"] = o.LicenseLegalMetadata
	}
	return toSerialize, nil
}

type NullableApiLicenseLegalApplicationReportDTO struct {
	value *ApiLicenseLegalApplicationReportDTO
	isSet bool
}

func (v NullableApiLicenseLegalApplicationReportDTO) Get() *ApiLicenseLegalApplicationReportDTO {
	return v.value
}

func (v *NullableApiLicenseLegalApplicationReportDTO) Set(val *ApiLicenseLegalApplicationReportDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicenseLegalApplicationReportDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicenseLegalApplicationReportDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicenseLegalApplicationReportDTO(val *ApiLicenseLegalApplicationReportDTO) *NullableApiLicenseLegalApplicationReportDTO {
	return &NullableApiLicenseLegalApplicationReportDTO{value: val, isSet: true}
}

func (v NullableApiLicenseLegalApplicationReportDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicenseLegalApplicationReportDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


