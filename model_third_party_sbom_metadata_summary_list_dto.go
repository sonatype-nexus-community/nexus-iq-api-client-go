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

// checks if the ThirdPartySbomMetadataSummaryListDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThirdPartySbomMetadataSummaryListDTO{}

// ThirdPartySbomMetadataSummaryListDTO struct for ThirdPartySbomMetadataSummaryListDTO
type ThirdPartySbomMetadataSummaryListDTO struct {
	Results []ThirdPartySbomMetadataSummaryDTO `json:"results,omitempty"`
	TotalResultsCount *int32 `json:"totalResultsCount,omitempty"`
}

// NewThirdPartySbomMetadataSummaryListDTO instantiates a new ThirdPartySbomMetadataSummaryListDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThirdPartySbomMetadataSummaryListDTO() *ThirdPartySbomMetadataSummaryListDTO {
	this := ThirdPartySbomMetadataSummaryListDTO{}
	return &this
}

// NewThirdPartySbomMetadataSummaryListDTOWithDefaults instantiates a new ThirdPartySbomMetadataSummaryListDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThirdPartySbomMetadataSummaryListDTOWithDefaults() *ThirdPartySbomMetadataSummaryListDTO {
	this := ThirdPartySbomMetadataSummaryListDTO{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ThirdPartySbomMetadataSummaryListDTO) GetResults() []ThirdPartySbomMetadataSummaryDTO {
	if o == nil || IsNil(o.Results) {
		var ret []ThirdPartySbomMetadataSummaryDTO
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartySbomMetadataSummaryListDTO) GetResultsOk() ([]ThirdPartySbomMetadataSummaryDTO, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ThirdPartySbomMetadataSummaryListDTO) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ThirdPartySbomMetadataSummaryDTO and assigns it to the Results field.
func (o *ThirdPartySbomMetadataSummaryListDTO) SetResults(v []ThirdPartySbomMetadataSummaryDTO) {
	o.Results = v
}

// GetTotalResultsCount returns the TotalResultsCount field value if set, zero value otherwise.
func (o *ThirdPartySbomMetadataSummaryListDTO) GetTotalResultsCount() int32 {
	if o == nil || IsNil(o.TotalResultsCount) {
		var ret int32
		return ret
	}
	return *o.TotalResultsCount
}

// GetTotalResultsCountOk returns a tuple with the TotalResultsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThirdPartySbomMetadataSummaryListDTO) GetTotalResultsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalResultsCount) {
		return nil, false
	}
	return o.TotalResultsCount, true
}

// HasTotalResultsCount returns a boolean if a field has been set.
func (o *ThirdPartySbomMetadataSummaryListDTO) HasTotalResultsCount() bool {
	if o != nil && !IsNil(o.TotalResultsCount) {
		return true
	}

	return false
}

// SetTotalResultsCount gets a reference to the given int32 and assigns it to the TotalResultsCount field.
func (o *ThirdPartySbomMetadataSummaryListDTO) SetTotalResultsCount(v int32) {
	o.TotalResultsCount = &v
}

func (o ThirdPartySbomMetadataSummaryListDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThirdPartySbomMetadataSummaryListDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.TotalResultsCount) {
		toSerialize["totalResultsCount"] = o.TotalResultsCount
	}
	return toSerialize, nil
}

type NullableThirdPartySbomMetadataSummaryListDTO struct {
	value *ThirdPartySbomMetadataSummaryListDTO
	isSet bool
}

func (v NullableThirdPartySbomMetadataSummaryListDTO) Get() *ThirdPartySbomMetadataSummaryListDTO {
	return v.value
}

func (v *NullableThirdPartySbomMetadataSummaryListDTO) Set(val *ThirdPartySbomMetadataSummaryListDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableThirdPartySbomMetadataSummaryListDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableThirdPartySbomMetadataSummaryListDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThirdPartySbomMetadataSummaryListDTO(val *ThirdPartySbomMetadataSummaryListDTO) *NullableThirdPartySbomMetadataSummaryListDTO {
	return &NullableThirdPartySbomMetadataSummaryListDTO{value: val, isSet: true}
}

func (v NullableThirdPartySbomMetadataSummaryListDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThirdPartySbomMetadataSummaryListDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


