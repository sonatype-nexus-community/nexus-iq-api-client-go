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

// checks if the ApiFirewallReleaseQuarantineConfigDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiFirewallReleaseQuarantineConfigDTO{}

// ApiFirewallReleaseQuarantineConfigDTO struct for ApiFirewallReleaseQuarantineConfigDTO
type ApiFirewallReleaseQuarantineConfigDTO struct {
	AutoReleaseQuarantineEnabled *bool `json:"autoReleaseQuarantineEnabled,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewApiFirewallReleaseQuarantineConfigDTO instantiates a new ApiFirewallReleaseQuarantineConfigDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiFirewallReleaseQuarantineConfigDTO() *ApiFirewallReleaseQuarantineConfigDTO {
	this := ApiFirewallReleaseQuarantineConfigDTO{}
	return &this
}

// NewApiFirewallReleaseQuarantineConfigDTOWithDefaults instantiates a new ApiFirewallReleaseQuarantineConfigDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiFirewallReleaseQuarantineConfigDTOWithDefaults() *ApiFirewallReleaseQuarantineConfigDTO {
	this := ApiFirewallReleaseQuarantineConfigDTO{}
	return &this
}

// GetAutoReleaseQuarantineEnabled returns the AutoReleaseQuarantineEnabled field value if set, zero value otherwise.
func (o *ApiFirewallReleaseQuarantineConfigDTO) GetAutoReleaseQuarantineEnabled() bool {
	if o == nil || IsNil(o.AutoReleaseQuarantineEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoReleaseQuarantineEnabled
}

// GetAutoReleaseQuarantineEnabledOk returns a tuple with the AutoReleaseQuarantineEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFirewallReleaseQuarantineConfigDTO) GetAutoReleaseQuarantineEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoReleaseQuarantineEnabled) {
		return nil, false
	}
	return o.AutoReleaseQuarantineEnabled, true
}

// HasAutoReleaseQuarantineEnabled returns a boolean if a field has been set.
func (o *ApiFirewallReleaseQuarantineConfigDTO) HasAutoReleaseQuarantineEnabled() bool {
	if o != nil && !IsNil(o.AutoReleaseQuarantineEnabled) {
		return true
	}

	return false
}

// SetAutoReleaseQuarantineEnabled gets a reference to the given bool and assigns it to the AutoReleaseQuarantineEnabled field.
func (o *ApiFirewallReleaseQuarantineConfigDTO) SetAutoReleaseQuarantineEnabled(v bool) {
	o.AutoReleaseQuarantineEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiFirewallReleaseQuarantineConfigDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFirewallReleaseQuarantineConfigDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiFirewallReleaseQuarantineConfigDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiFirewallReleaseQuarantineConfigDTO) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiFirewallReleaseQuarantineConfigDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFirewallReleaseQuarantineConfigDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiFirewallReleaseQuarantineConfigDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiFirewallReleaseQuarantineConfigDTO) SetName(v string) {
	o.Name = &v
}

func (o ApiFirewallReleaseQuarantineConfigDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiFirewallReleaseQuarantineConfigDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoReleaseQuarantineEnabled) {
		toSerialize["autoReleaseQuarantineEnabled"] = o.AutoReleaseQuarantineEnabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableApiFirewallReleaseQuarantineConfigDTO struct {
	value *ApiFirewallReleaseQuarantineConfigDTO
	isSet bool
}

func (v NullableApiFirewallReleaseQuarantineConfigDTO) Get() *ApiFirewallReleaseQuarantineConfigDTO {
	return v.value
}

func (v *NullableApiFirewallReleaseQuarantineConfigDTO) Set(val *ApiFirewallReleaseQuarantineConfigDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiFirewallReleaseQuarantineConfigDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiFirewallReleaseQuarantineConfigDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiFirewallReleaseQuarantineConfigDTO(val *ApiFirewallReleaseQuarantineConfigDTO) *NullableApiFirewallReleaseQuarantineConfigDTO {
	return &NullableApiFirewallReleaseQuarantineConfigDTO{value: val, isSet: true}
}

func (v NullableApiFirewallReleaseQuarantineConfigDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiFirewallReleaseQuarantineConfigDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


