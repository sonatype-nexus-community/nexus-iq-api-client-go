/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.182.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiFirewallReleaseQuarantineSummaryDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiFirewallReleaseQuarantineSummaryDTO{}

// ApiFirewallReleaseQuarantineSummaryDTO struct for ApiFirewallReleaseQuarantineSummaryDTO
type ApiFirewallReleaseQuarantineSummaryDTO struct {
	AutoReleaseQuarantineCountMTD *int64 `json:"autoReleaseQuarantineCountMTD,omitempty"`
	AutoReleaseQuarantineCountYTD *int64 `json:"autoReleaseQuarantineCountYTD,omitempty"`
}

// NewApiFirewallReleaseQuarantineSummaryDTO instantiates a new ApiFirewallReleaseQuarantineSummaryDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiFirewallReleaseQuarantineSummaryDTO() *ApiFirewallReleaseQuarantineSummaryDTO {
	this := ApiFirewallReleaseQuarantineSummaryDTO{}
	return &this
}

// NewApiFirewallReleaseQuarantineSummaryDTOWithDefaults instantiates a new ApiFirewallReleaseQuarantineSummaryDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiFirewallReleaseQuarantineSummaryDTOWithDefaults() *ApiFirewallReleaseQuarantineSummaryDTO {
	this := ApiFirewallReleaseQuarantineSummaryDTO{}
	return &this
}

// GetAutoReleaseQuarantineCountMTD returns the AutoReleaseQuarantineCountMTD field value if set, zero value otherwise.
func (o *ApiFirewallReleaseQuarantineSummaryDTO) GetAutoReleaseQuarantineCountMTD() int64 {
	if o == nil || IsNil(o.AutoReleaseQuarantineCountMTD) {
		var ret int64
		return ret
	}
	return *o.AutoReleaseQuarantineCountMTD
}

// GetAutoReleaseQuarantineCountMTDOk returns a tuple with the AutoReleaseQuarantineCountMTD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFirewallReleaseQuarantineSummaryDTO) GetAutoReleaseQuarantineCountMTDOk() (*int64, bool) {
	if o == nil || IsNil(o.AutoReleaseQuarantineCountMTD) {
		return nil, false
	}
	return o.AutoReleaseQuarantineCountMTD, true
}

// HasAutoReleaseQuarantineCountMTD returns a boolean if a field has been set.
func (o *ApiFirewallReleaseQuarantineSummaryDTO) HasAutoReleaseQuarantineCountMTD() bool {
	if o != nil && !IsNil(o.AutoReleaseQuarantineCountMTD) {
		return true
	}

	return false
}

// SetAutoReleaseQuarantineCountMTD gets a reference to the given int64 and assigns it to the AutoReleaseQuarantineCountMTD field.
func (o *ApiFirewallReleaseQuarantineSummaryDTO) SetAutoReleaseQuarantineCountMTD(v int64) {
	o.AutoReleaseQuarantineCountMTD = &v
}

// GetAutoReleaseQuarantineCountYTD returns the AutoReleaseQuarantineCountYTD field value if set, zero value otherwise.
func (o *ApiFirewallReleaseQuarantineSummaryDTO) GetAutoReleaseQuarantineCountYTD() int64 {
	if o == nil || IsNil(o.AutoReleaseQuarantineCountYTD) {
		var ret int64
		return ret
	}
	return *o.AutoReleaseQuarantineCountYTD
}

// GetAutoReleaseQuarantineCountYTDOk returns a tuple with the AutoReleaseQuarantineCountYTD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiFirewallReleaseQuarantineSummaryDTO) GetAutoReleaseQuarantineCountYTDOk() (*int64, bool) {
	if o == nil || IsNil(o.AutoReleaseQuarantineCountYTD) {
		return nil, false
	}
	return o.AutoReleaseQuarantineCountYTD, true
}

// HasAutoReleaseQuarantineCountYTD returns a boolean if a field has been set.
func (o *ApiFirewallReleaseQuarantineSummaryDTO) HasAutoReleaseQuarantineCountYTD() bool {
	if o != nil && !IsNil(o.AutoReleaseQuarantineCountYTD) {
		return true
	}

	return false
}

// SetAutoReleaseQuarantineCountYTD gets a reference to the given int64 and assigns it to the AutoReleaseQuarantineCountYTD field.
func (o *ApiFirewallReleaseQuarantineSummaryDTO) SetAutoReleaseQuarantineCountYTD(v int64) {
	o.AutoReleaseQuarantineCountYTD = &v
}

func (o ApiFirewallReleaseQuarantineSummaryDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiFirewallReleaseQuarantineSummaryDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoReleaseQuarantineCountMTD) {
		toSerialize["autoReleaseQuarantineCountMTD"] = o.AutoReleaseQuarantineCountMTD
	}
	if !IsNil(o.AutoReleaseQuarantineCountYTD) {
		toSerialize["autoReleaseQuarantineCountYTD"] = o.AutoReleaseQuarantineCountYTD
	}
	return toSerialize, nil
}

type NullableApiFirewallReleaseQuarantineSummaryDTO struct {
	value *ApiFirewallReleaseQuarantineSummaryDTO
	isSet bool
}

func (v NullableApiFirewallReleaseQuarantineSummaryDTO) Get() *ApiFirewallReleaseQuarantineSummaryDTO {
	return v.value
}

func (v *NullableApiFirewallReleaseQuarantineSummaryDTO) Set(val *ApiFirewallReleaseQuarantineSummaryDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiFirewallReleaseQuarantineSummaryDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiFirewallReleaseQuarantineSummaryDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiFirewallReleaseQuarantineSummaryDTO(val *ApiFirewallReleaseQuarantineSummaryDTO) *NullableApiFirewallReleaseQuarantineSummaryDTO {
	return &NullableApiFirewallReleaseQuarantineSummaryDTO{value: val, isSet: true}
}

func (v NullableApiFirewallReleaseQuarantineSummaryDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiFirewallReleaseQuarantineSummaryDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


