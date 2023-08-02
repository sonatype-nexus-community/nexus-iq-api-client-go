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

// checks if the ApiReverseProxyAuthenticationConfigurationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReverseProxyAuthenticationConfigurationDTO{}

// ApiReverseProxyAuthenticationConfigurationDTO struct for ApiReverseProxyAuthenticationConfigurationDTO
type ApiReverseProxyAuthenticationConfigurationDTO struct {
	CsrfProtectionDisabled *bool `json:"csrfProtectionDisabled,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	UsernameHeader *string `json:"usernameHeader,omitempty"`
}

// NewApiReverseProxyAuthenticationConfigurationDTO instantiates a new ApiReverseProxyAuthenticationConfigurationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReverseProxyAuthenticationConfigurationDTO() *ApiReverseProxyAuthenticationConfigurationDTO {
	this := ApiReverseProxyAuthenticationConfigurationDTO{}
	return &this
}

// NewApiReverseProxyAuthenticationConfigurationDTOWithDefaults instantiates a new ApiReverseProxyAuthenticationConfigurationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReverseProxyAuthenticationConfigurationDTOWithDefaults() *ApiReverseProxyAuthenticationConfigurationDTO {
	this := ApiReverseProxyAuthenticationConfigurationDTO{}
	return &this
}

// GetCsrfProtectionDisabled returns the CsrfProtectionDisabled field value if set, zero value otherwise.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) GetCsrfProtectionDisabled() bool {
	if o == nil || IsNil(o.CsrfProtectionDisabled) {
		var ret bool
		return ret
	}
	return *o.CsrfProtectionDisabled
}

// GetCsrfProtectionDisabledOk returns a tuple with the CsrfProtectionDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) GetCsrfProtectionDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CsrfProtectionDisabled) {
		return nil, false
	}
	return o.CsrfProtectionDisabled, true
}

// HasCsrfProtectionDisabled returns a boolean if a field has been set.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) HasCsrfProtectionDisabled() bool {
	if o != nil && !IsNil(o.CsrfProtectionDisabled) {
		return true
	}

	return false
}

// SetCsrfProtectionDisabled gets a reference to the given bool and assigns it to the CsrfProtectionDisabled field.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) SetCsrfProtectionDisabled(v bool) {
	o.CsrfProtectionDisabled = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) GetLogoutUrl() string {
	if o == nil || IsNil(o.LogoutUrl) {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) GetLogoutUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LogoutUrl) {
		return nil, false
	}
	return o.LogoutUrl, true
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) HasLogoutUrl() bool {
	if o != nil && !IsNil(o.LogoutUrl) {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

// GetUsernameHeader returns the UsernameHeader field value if set, zero value otherwise.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) GetUsernameHeader() string {
	if o == nil || IsNil(o.UsernameHeader) {
		var ret string
		return ret
	}
	return *o.UsernameHeader
}

// GetUsernameHeaderOk returns a tuple with the UsernameHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) GetUsernameHeaderOk() (*string, bool) {
	if o == nil || IsNil(o.UsernameHeader) {
		return nil, false
	}
	return o.UsernameHeader, true
}

// HasUsernameHeader returns a boolean if a field has been set.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) HasUsernameHeader() bool {
	if o != nil && !IsNil(o.UsernameHeader) {
		return true
	}

	return false
}

// SetUsernameHeader gets a reference to the given string and assigns it to the UsernameHeader field.
func (o *ApiReverseProxyAuthenticationConfigurationDTO) SetUsernameHeader(v string) {
	o.UsernameHeader = &v
}

func (o ApiReverseProxyAuthenticationConfigurationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReverseProxyAuthenticationConfigurationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CsrfProtectionDisabled) {
		toSerialize["csrfProtectionDisabled"] = o.CsrfProtectionDisabled
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.LogoutUrl) {
		toSerialize["logoutUrl"] = o.LogoutUrl
	}
	if !IsNil(o.UsernameHeader) {
		toSerialize["usernameHeader"] = o.UsernameHeader
	}
	return toSerialize, nil
}

type NullableApiReverseProxyAuthenticationConfigurationDTO struct {
	value *ApiReverseProxyAuthenticationConfigurationDTO
	isSet bool
}

func (v NullableApiReverseProxyAuthenticationConfigurationDTO) Get() *ApiReverseProxyAuthenticationConfigurationDTO {
	return v.value
}

func (v *NullableApiReverseProxyAuthenticationConfigurationDTO) Set(val *ApiReverseProxyAuthenticationConfigurationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReverseProxyAuthenticationConfigurationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReverseProxyAuthenticationConfigurationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReverseProxyAuthenticationConfigurationDTO(val *ApiReverseProxyAuthenticationConfigurationDTO) *NullableApiReverseProxyAuthenticationConfigurationDTO {
	return &NullableApiReverseProxyAuthenticationConfigurationDTO{value: val, isSet: true}
}

func (v NullableApiReverseProxyAuthenticationConfigurationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReverseProxyAuthenticationConfigurationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

