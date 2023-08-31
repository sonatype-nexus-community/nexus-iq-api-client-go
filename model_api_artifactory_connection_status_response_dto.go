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

// checks if the ApiArtifactoryConnectionStatusResponseDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiArtifactoryConnectionStatusResponseDTO{}

// ApiArtifactoryConnectionStatusResponseDTO struct for ApiArtifactoryConnectionStatusResponseDTO
type ApiArtifactoryConnectionStatusResponseDTO struct {
	AllowChange *bool `json:"allowChange,omitempty"`
	AllowOverride *bool `json:"allowOverride,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	InheritedFromOrgEnabled *bool `json:"inheritedFromOrgEnabled,omitempty"`
	InheritedFromOrganizationId *string `json:"inheritedFromOrganizationId,omitempty"`
	InheritedFromOrganizationName *string `json:"inheritedFromOrganizationName,omitempty"`
}

// NewApiArtifactoryConnectionStatusResponseDTO instantiates a new ApiArtifactoryConnectionStatusResponseDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiArtifactoryConnectionStatusResponseDTO() *ApiArtifactoryConnectionStatusResponseDTO {
	this := ApiArtifactoryConnectionStatusResponseDTO{}
	return &this
}

// NewApiArtifactoryConnectionStatusResponseDTOWithDefaults instantiates a new ApiArtifactoryConnectionStatusResponseDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiArtifactoryConnectionStatusResponseDTOWithDefaults() *ApiArtifactoryConnectionStatusResponseDTO {
	this := ApiArtifactoryConnectionStatusResponseDTO{}
	return &this
}

// GetAllowChange returns the AllowChange field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionStatusResponseDTO) GetAllowChange() bool {
	if o == nil || IsNil(o.AllowChange) {
		var ret bool
		return ret
	}
	return *o.AllowChange
}

// GetAllowChangeOk returns a tuple with the AllowChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionStatusResponseDTO) GetAllowChangeOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowChange) {
		return nil, false
	}
	return o.AllowChange, true
}

// HasAllowChange returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionStatusResponseDTO) HasAllowChange() bool {
	if o != nil && !IsNil(o.AllowChange) {
		return true
	}

	return false
}

// SetAllowChange gets a reference to the given bool and assigns it to the AllowChange field.
func (o *ApiArtifactoryConnectionStatusResponseDTO) SetAllowChange(v bool) {
	o.AllowChange = &v
}

// GetAllowOverride returns the AllowOverride field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionStatusResponseDTO) GetAllowOverride() bool {
	if o == nil || IsNil(o.AllowOverride) {
		var ret bool
		return ret
	}
	return *o.AllowOverride
}

// GetAllowOverrideOk returns a tuple with the AllowOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionStatusResponseDTO) GetAllowOverrideOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowOverride) {
		return nil, false
	}
	return o.AllowOverride, true
}

// HasAllowOverride returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionStatusResponseDTO) HasAllowOverride() bool {
	if o != nil && !IsNil(o.AllowOverride) {
		return true
	}

	return false
}

// SetAllowOverride gets a reference to the given bool and assigns it to the AllowOverride field.
func (o *ApiArtifactoryConnectionStatusResponseDTO) SetAllowOverride(v bool) {
	o.AllowOverride = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionStatusResponseDTO) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionStatusResponseDTO) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionStatusResponseDTO) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiArtifactoryConnectionStatusResponseDTO) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInheritedFromOrgEnabled returns the InheritedFromOrgEnabled field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionStatusResponseDTO) GetInheritedFromOrgEnabled() bool {
	if o == nil || IsNil(o.InheritedFromOrgEnabled) {
		var ret bool
		return ret
	}
	return *o.InheritedFromOrgEnabled
}

// GetInheritedFromOrgEnabledOk returns a tuple with the InheritedFromOrgEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionStatusResponseDTO) GetInheritedFromOrgEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.InheritedFromOrgEnabled) {
		return nil, false
	}
	return o.InheritedFromOrgEnabled, true
}

// HasInheritedFromOrgEnabled returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionStatusResponseDTO) HasInheritedFromOrgEnabled() bool {
	if o != nil && !IsNil(o.InheritedFromOrgEnabled) {
		return true
	}

	return false
}

// SetInheritedFromOrgEnabled gets a reference to the given bool and assigns it to the InheritedFromOrgEnabled field.
func (o *ApiArtifactoryConnectionStatusResponseDTO) SetInheritedFromOrgEnabled(v bool) {
	o.InheritedFromOrgEnabled = &v
}

// GetInheritedFromOrganizationId returns the InheritedFromOrganizationId field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionStatusResponseDTO) GetInheritedFromOrganizationId() string {
	if o == nil || IsNil(o.InheritedFromOrganizationId) {
		var ret string
		return ret
	}
	return *o.InheritedFromOrganizationId
}

// GetInheritedFromOrganizationIdOk returns a tuple with the InheritedFromOrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionStatusResponseDTO) GetInheritedFromOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.InheritedFromOrganizationId) {
		return nil, false
	}
	return o.InheritedFromOrganizationId, true
}

// HasInheritedFromOrganizationId returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionStatusResponseDTO) HasInheritedFromOrganizationId() bool {
	if o != nil && !IsNil(o.InheritedFromOrganizationId) {
		return true
	}

	return false
}

// SetInheritedFromOrganizationId gets a reference to the given string and assigns it to the InheritedFromOrganizationId field.
func (o *ApiArtifactoryConnectionStatusResponseDTO) SetInheritedFromOrganizationId(v string) {
	o.InheritedFromOrganizationId = &v
}

// GetInheritedFromOrganizationName returns the InheritedFromOrganizationName field value if set, zero value otherwise.
func (o *ApiArtifactoryConnectionStatusResponseDTO) GetInheritedFromOrganizationName() string {
	if o == nil || IsNil(o.InheritedFromOrganizationName) {
		var ret string
		return ret
	}
	return *o.InheritedFromOrganizationName
}

// GetInheritedFromOrganizationNameOk returns a tuple with the InheritedFromOrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiArtifactoryConnectionStatusResponseDTO) GetInheritedFromOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.InheritedFromOrganizationName) {
		return nil, false
	}
	return o.InheritedFromOrganizationName, true
}

// HasInheritedFromOrganizationName returns a boolean if a field has been set.
func (o *ApiArtifactoryConnectionStatusResponseDTO) HasInheritedFromOrganizationName() bool {
	if o != nil && !IsNil(o.InheritedFromOrganizationName) {
		return true
	}

	return false
}

// SetInheritedFromOrganizationName gets a reference to the given string and assigns it to the InheritedFromOrganizationName field.
func (o *ApiArtifactoryConnectionStatusResponseDTO) SetInheritedFromOrganizationName(v string) {
	o.InheritedFromOrganizationName = &v
}

func (o ApiArtifactoryConnectionStatusResponseDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiArtifactoryConnectionStatusResponseDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowChange) {
		toSerialize["allowChange"] = o.AllowChange
	}
	if !IsNil(o.AllowOverride) {
		toSerialize["allowOverride"] = o.AllowOverride
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.InheritedFromOrgEnabled) {
		toSerialize["inheritedFromOrgEnabled"] = o.InheritedFromOrgEnabled
	}
	if !IsNil(o.InheritedFromOrganizationId) {
		toSerialize["inheritedFromOrganizationId"] = o.InheritedFromOrganizationId
	}
	if !IsNil(o.InheritedFromOrganizationName) {
		toSerialize["inheritedFromOrganizationName"] = o.InheritedFromOrganizationName
	}
	return toSerialize, nil
}

type NullableApiArtifactoryConnectionStatusResponseDTO struct {
	value *ApiArtifactoryConnectionStatusResponseDTO
	isSet bool
}

func (v NullableApiArtifactoryConnectionStatusResponseDTO) Get() *ApiArtifactoryConnectionStatusResponseDTO {
	return v.value
}

func (v *NullableApiArtifactoryConnectionStatusResponseDTO) Set(val *ApiArtifactoryConnectionStatusResponseDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiArtifactoryConnectionStatusResponseDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiArtifactoryConnectionStatusResponseDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiArtifactoryConnectionStatusResponseDTO(val *ApiArtifactoryConnectionStatusResponseDTO) *NullableApiArtifactoryConnectionStatusResponseDTO {
	return &NullableApiArtifactoryConnectionStatusResponseDTO{value: val, isSet: true}
}

func (v NullableApiArtifactoryConnectionStatusResponseDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiArtifactoryConnectionStatusResponseDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


