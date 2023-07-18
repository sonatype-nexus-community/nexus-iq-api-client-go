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

// checks if the ApiSourceControlEvaluationRequestDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSourceControlEvaluationRequestDTO{}

// ApiSourceControlEvaluationRequestDTO struct for ApiSourceControlEvaluationRequestDTO
type ApiSourceControlEvaluationRequestDTO struct {
	BranchName *string `json:"branchName,omitempty"`
	ScanTargets []string `json:"scanTargets,omitempty"`
	StageId *string `json:"stageId,omitempty"`
}

// NewApiSourceControlEvaluationRequestDTO instantiates a new ApiSourceControlEvaluationRequestDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSourceControlEvaluationRequestDTO() *ApiSourceControlEvaluationRequestDTO {
	this := ApiSourceControlEvaluationRequestDTO{}
	return &this
}

// NewApiSourceControlEvaluationRequestDTOWithDefaults instantiates a new ApiSourceControlEvaluationRequestDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSourceControlEvaluationRequestDTOWithDefaults() *ApiSourceControlEvaluationRequestDTO {
	this := ApiSourceControlEvaluationRequestDTO{}
	return &this
}

// GetBranchName returns the BranchName field value if set, zero value otherwise.
func (o *ApiSourceControlEvaluationRequestDTO) GetBranchName() string {
	if o == nil || IsNil(o.BranchName) {
		var ret string
		return ret
	}
	return *o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlEvaluationRequestDTO) GetBranchNameOk() (*string, bool) {
	if o == nil || IsNil(o.BranchName) {
		return nil, false
	}
	return o.BranchName, true
}

// HasBranchName returns a boolean if a field has been set.
func (o *ApiSourceControlEvaluationRequestDTO) HasBranchName() bool {
	if o != nil && !IsNil(o.BranchName) {
		return true
	}

	return false
}

// SetBranchName gets a reference to the given string and assigns it to the BranchName field.
func (o *ApiSourceControlEvaluationRequestDTO) SetBranchName(v string) {
	o.BranchName = &v
}

// GetScanTargets returns the ScanTargets field value if set, zero value otherwise.
func (o *ApiSourceControlEvaluationRequestDTO) GetScanTargets() []string {
	if o == nil || IsNil(o.ScanTargets) {
		var ret []string
		return ret
	}
	return o.ScanTargets
}

// GetScanTargetsOk returns a tuple with the ScanTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlEvaluationRequestDTO) GetScanTargetsOk() ([]string, bool) {
	if o == nil || IsNil(o.ScanTargets) {
		return nil, false
	}
	return o.ScanTargets, true
}

// HasScanTargets returns a boolean if a field has been set.
func (o *ApiSourceControlEvaluationRequestDTO) HasScanTargets() bool {
	if o != nil && !IsNil(o.ScanTargets) {
		return true
	}

	return false
}

// SetScanTargets gets a reference to the given []string and assigns it to the ScanTargets field.
func (o *ApiSourceControlEvaluationRequestDTO) SetScanTargets(v []string) {
	o.ScanTargets = v
}

// GetStageId returns the StageId field value if set, zero value otherwise.
func (o *ApiSourceControlEvaluationRequestDTO) GetStageId() string {
	if o == nil || IsNil(o.StageId) {
		var ret string
		return ret
	}
	return *o.StageId
}

// GetStageIdOk returns a tuple with the StageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlEvaluationRequestDTO) GetStageIdOk() (*string, bool) {
	if o == nil || IsNil(o.StageId) {
		return nil, false
	}
	return o.StageId, true
}

// HasStageId returns a boolean if a field has been set.
func (o *ApiSourceControlEvaluationRequestDTO) HasStageId() bool {
	if o != nil && !IsNil(o.StageId) {
		return true
	}

	return false
}

// SetStageId gets a reference to the given string and assigns it to the StageId field.
func (o *ApiSourceControlEvaluationRequestDTO) SetStageId(v string) {
	o.StageId = &v
}

func (o ApiSourceControlEvaluationRequestDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSourceControlEvaluationRequestDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BranchName) {
		toSerialize["branchName"] = o.BranchName
	}
	if !IsNil(o.ScanTargets) {
		toSerialize["scanTargets"] = o.ScanTargets
	}
	if !IsNil(o.StageId) {
		toSerialize["stageId"] = o.StageId
	}
	return toSerialize, nil
}

type NullableApiSourceControlEvaluationRequestDTO struct {
	value *ApiSourceControlEvaluationRequestDTO
	isSet bool
}

func (v NullableApiSourceControlEvaluationRequestDTO) Get() *ApiSourceControlEvaluationRequestDTO {
	return v.value
}

func (v *NullableApiSourceControlEvaluationRequestDTO) Set(val *ApiSourceControlEvaluationRequestDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSourceControlEvaluationRequestDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSourceControlEvaluationRequestDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSourceControlEvaluationRequestDTO(val *ApiSourceControlEvaluationRequestDTO) *NullableApiSourceControlEvaluationRequestDTO {
	return &NullableApiSourceControlEvaluationRequestDTO{value: val, isSet: true}
}

func (v NullableApiSourceControlEvaluationRequestDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSourceControlEvaluationRequestDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


