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

// checks if the ApiReportRawDataDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReportRawDataDTOV2{}

// ApiReportRawDataDTOV2 struct for ApiReportRawDataDTOV2
type ApiReportRawDataDTOV2 struct {
	Components []ApiReportComponentDTOV2 `json:"components,omitempty"`
	MatchSummary *ApiMatchStateSummaryDTOV2 `json:"matchSummary,omitempty"`
}

// NewApiReportRawDataDTOV2 instantiates a new ApiReportRawDataDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReportRawDataDTOV2() *ApiReportRawDataDTOV2 {
	this := ApiReportRawDataDTOV2{}
	return &this
}

// NewApiReportRawDataDTOV2WithDefaults instantiates a new ApiReportRawDataDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReportRawDataDTOV2WithDefaults() *ApiReportRawDataDTOV2 {
	this := ApiReportRawDataDTOV2{}
	return &this
}

// GetComponents returns the Components field value if set, zero value otherwise.
func (o *ApiReportRawDataDTOV2) GetComponents() []ApiReportComponentDTOV2 {
	if o == nil || IsNil(o.Components) {
		var ret []ApiReportComponentDTOV2
		return ret
	}
	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportRawDataDTOV2) GetComponentsOk() ([]ApiReportComponentDTOV2, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// HasComponents returns a boolean if a field has been set.
func (o *ApiReportRawDataDTOV2) HasComponents() bool {
	if o != nil && !IsNil(o.Components) {
		return true
	}

	return false
}

// SetComponents gets a reference to the given []ApiReportComponentDTOV2 and assigns it to the Components field.
func (o *ApiReportRawDataDTOV2) SetComponents(v []ApiReportComponentDTOV2) {
	o.Components = v
}

// GetMatchSummary returns the MatchSummary field value if set, zero value otherwise.
func (o *ApiReportRawDataDTOV2) GetMatchSummary() ApiMatchStateSummaryDTOV2 {
	if o == nil || IsNil(o.MatchSummary) {
		var ret ApiMatchStateSummaryDTOV2
		return ret
	}
	return *o.MatchSummary
}

// GetMatchSummaryOk returns a tuple with the MatchSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportRawDataDTOV2) GetMatchSummaryOk() (*ApiMatchStateSummaryDTOV2, bool) {
	if o == nil || IsNil(o.MatchSummary) {
		return nil, false
	}
	return o.MatchSummary, true
}

// HasMatchSummary returns a boolean if a field has been set.
func (o *ApiReportRawDataDTOV2) HasMatchSummary() bool {
	if o != nil && !IsNil(o.MatchSummary) {
		return true
	}

	return false
}

// SetMatchSummary gets a reference to the given ApiMatchStateSummaryDTOV2 and assigns it to the MatchSummary field.
func (o *ApiReportRawDataDTOV2) SetMatchSummary(v ApiMatchStateSummaryDTOV2) {
	o.MatchSummary = &v
}

func (o ApiReportRawDataDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReportRawDataDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Components) {
		toSerialize["components"] = o.Components
	}
	if !IsNil(o.MatchSummary) {
		toSerialize["matchSummary"] = o.MatchSummary
	}
	return toSerialize, nil
}

type NullableApiReportRawDataDTOV2 struct {
	value *ApiReportRawDataDTOV2
	isSet bool
}

func (v NullableApiReportRawDataDTOV2) Get() *ApiReportRawDataDTOV2 {
	return v.value
}

func (v *NullableApiReportRawDataDTOV2) Set(val *ApiReportRawDataDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReportRawDataDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReportRawDataDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReportRawDataDTOV2(val *ApiReportRawDataDTOV2) *NullableApiReportRawDataDTOV2 {
	return &NullableApiReportRawDataDTOV2{value: val, isSet: true}
}

func (v NullableApiReportRawDataDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReportRawDataDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


