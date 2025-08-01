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

// checks if the SystemConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SystemConfig{}

// SystemConfig struct for SystemConfig
type SystemConfig struct {
	BaseUrl NullableString `json:"baseUrl,omitempty"`
	ForceBaseUrl NullableBool `json:"forceBaseUrl,omitempty"`
}

// NewSystemConfig instantiates a new SystemConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemConfig() *SystemConfig {
	this := SystemConfig{}
	return &this
}

// NewSystemConfigWithDefaults instantiates a new SystemConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemConfigWithDefaults() *SystemConfig {
	this := SystemConfig{}
	return &this
}

// GetBaseUrl returns the BaseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfig) GetBaseUrl() string {
	if o == nil || IsNil(o.BaseUrl.Get()) {
		var ret string
		return ret
	}
	return *o.BaseUrl.Get()
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfig) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseUrl.Get(), o.BaseUrl.IsSet()
}

// HasBaseUrl returns a boolean if a field has been set.
func (o *SystemConfig) HasBaseUrl() bool {
	if o != nil && o.BaseUrl.IsSet() {
		return true
	}

	return false
}

// SetBaseUrl gets a reference to the given NullableString and assigns it to the BaseUrl field.
func (o *SystemConfig) SetBaseUrl(v string) {
	o.BaseUrl.Set(&v)
}
// SetBaseUrlNil sets the value for BaseUrl to be an explicit nil
func (o *SystemConfig) SetBaseUrlNil() {
	o.BaseUrl.Set(nil)
}

// UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
func (o *SystemConfig) UnsetBaseUrl() {
	o.BaseUrl.Unset()
}

// GetForceBaseUrl returns the ForceBaseUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SystemConfig) GetForceBaseUrl() bool {
	if o == nil || IsNil(o.ForceBaseUrl.Get()) {
		var ret bool
		return ret
	}
	return *o.ForceBaseUrl.Get()
}

// GetForceBaseUrlOk returns a tuple with the ForceBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SystemConfig) GetForceBaseUrlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForceBaseUrl.Get(), o.ForceBaseUrl.IsSet()
}

// HasForceBaseUrl returns a boolean if a field has been set.
func (o *SystemConfig) HasForceBaseUrl() bool {
	if o != nil && o.ForceBaseUrl.IsSet() {
		return true
	}

	return false
}

// SetForceBaseUrl gets a reference to the given NullableBool and assigns it to the ForceBaseUrl field.
func (o *SystemConfig) SetForceBaseUrl(v bool) {
	o.ForceBaseUrl.Set(&v)
}
// SetForceBaseUrlNil sets the value for ForceBaseUrl to be an explicit nil
func (o *SystemConfig) SetForceBaseUrlNil() {
	o.ForceBaseUrl.Set(nil)
}

// UnsetForceBaseUrl ensures that no value is present for ForceBaseUrl, not even an explicit nil
func (o *SystemConfig) UnsetForceBaseUrl() {
	o.ForceBaseUrl.Unset()
}

func (o SystemConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseUrl.IsSet() {
		toSerialize["baseUrl"] = o.BaseUrl.Get()
	}
	if o.ForceBaseUrl.IsSet() {
		toSerialize["forceBaseUrl"] = o.ForceBaseUrl.Get()
	}
	return toSerialize, nil
}

type NullableSystemConfig struct {
	value *SystemConfig
	isSet bool
}

func (v NullableSystemConfig) Get() *SystemConfig {
	return v.value
}

func (v *NullableSystemConfig) Set(val *SystemConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemConfig(val *SystemConfig) *NullableSystemConfig {
	return &NullableSystemConfig{value: val, isSet: true}
}

func (v NullableSystemConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


