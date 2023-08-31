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

// checks if the ApiDataRetentionPoliciesDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDataRetentionPoliciesDTO{}

// ApiDataRetentionPoliciesDTO struct for ApiDataRetentionPoliciesDTO
type ApiDataRetentionPoliciesDTO struct {
	ApplicationReports *ApiReportRetentionPoliciesDTO `json:"applicationReports,omitempty"`
	SuccessMetrics *ApiSuccessMetricsRetentionPolicyDTO `json:"successMetrics,omitempty"`
}

// NewApiDataRetentionPoliciesDTO instantiates a new ApiDataRetentionPoliciesDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDataRetentionPoliciesDTO() *ApiDataRetentionPoliciesDTO {
	this := ApiDataRetentionPoliciesDTO{}
	return &this
}

// NewApiDataRetentionPoliciesDTOWithDefaults instantiates a new ApiDataRetentionPoliciesDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDataRetentionPoliciesDTOWithDefaults() *ApiDataRetentionPoliciesDTO {
	this := ApiDataRetentionPoliciesDTO{}
	return &this
}

// GetApplicationReports returns the ApplicationReports field value if set, zero value otherwise.
func (o *ApiDataRetentionPoliciesDTO) GetApplicationReports() ApiReportRetentionPoliciesDTO {
	if o == nil || IsNil(o.ApplicationReports) {
		var ret ApiReportRetentionPoliciesDTO
		return ret
	}
	return *o.ApplicationReports
}

// GetApplicationReportsOk returns a tuple with the ApplicationReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDataRetentionPoliciesDTO) GetApplicationReportsOk() (*ApiReportRetentionPoliciesDTO, bool) {
	if o == nil || IsNil(o.ApplicationReports) {
		return nil, false
	}
	return o.ApplicationReports, true
}

// HasApplicationReports returns a boolean if a field has been set.
func (o *ApiDataRetentionPoliciesDTO) HasApplicationReports() bool {
	if o != nil && !IsNil(o.ApplicationReports) {
		return true
	}

	return false
}

// SetApplicationReports gets a reference to the given ApiReportRetentionPoliciesDTO and assigns it to the ApplicationReports field.
func (o *ApiDataRetentionPoliciesDTO) SetApplicationReports(v ApiReportRetentionPoliciesDTO) {
	o.ApplicationReports = &v
}

// GetSuccessMetrics returns the SuccessMetrics field value if set, zero value otherwise.
func (o *ApiDataRetentionPoliciesDTO) GetSuccessMetrics() ApiSuccessMetricsRetentionPolicyDTO {
	if o == nil || IsNil(o.SuccessMetrics) {
		var ret ApiSuccessMetricsRetentionPolicyDTO
		return ret
	}
	return *o.SuccessMetrics
}

// GetSuccessMetricsOk returns a tuple with the SuccessMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDataRetentionPoliciesDTO) GetSuccessMetricsOk() (*ApiSuccessMetricsRetentionPolicyDTO, bool) {
	if o == nil || IsNil(o.SuccessMetrics) {
		return nil, false
	}
	return o.SuccessMetrics, true
}

// HasSuccessMetrics returns a boolean if a field has been set.
func (o *ApiDataRetentionPoliciesDTO) HasSuccessMetrics() bool {
	if o != nil && !IsNil(o.SuccessMetrics) {
		return true
	}

	return false
}

// SetSuccessMetrics gets a reference to the given ApiSuccessMetricsRetentionPolicyDTO and assigns it to the SuccessMetrics field.
func (o *ApiDataRetentionPoliciesDTO) SetSuccessMetrics(v ApiSuccessMetricsRetentionPolicyDTO) {
	o.SuccessMetrics = &v
}

func (o ApiDataRetentionPoliciesDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDataRetentionPoliciesDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationReports) {
		toSerialize["applicationReports"] = o.ApplicationReports
	}
	if !IsNil(o.SuccessMetrics) {
		toSerialize["successMetrics"] = o.SuccessMetrics
	}
	return toSerialize, nil
}

type NullableApiDataRetentionPoliciesDTO struct {
	value *ApiDataRetentionPoliciesDTO
	isSet bool
}

func (v NullableApiDataRetentionPoliciesDTO) Get() *ApiDataRetentionPoliciesDTO {
	return v.value
}

func (v *NullableApiDataRetentionPoliciesDTO) Set(val *ApiDataRetentionPoliciesDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDataRetentionPoliciesDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDataRetentionPoliciesDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDataRetentionPoliciesDTO(val *ApiDataRetentionPoliciesDTO) *NullableApiDataRetentionPoliciesDTO {
	return &NullableApiDataRetentionPoliciesDTO{value: val, isSet: true}
}

func (v NullableApiDataRetentionPoliciesDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDataRetentionPoliciesDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


