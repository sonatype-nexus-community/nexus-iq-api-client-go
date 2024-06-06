/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.177.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiApplicationEvaluationCommitDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiApplicationEvaluationCommitDTO{}

// ApiApplicationEvaluationCommitDTO struct for ApiApplicationEvaluationCommitDTO
type ApiApplicationEvaluationCommitDTO struct {
	CommitHash *string `json:"commitHash,omitempty"`
	ReportUrl *string `json:"reportUrl,omitempty"`
	ScanId *string `json:"scanId,omitempty"`
	ScanTime *time.Time `json:"scanTime,omitempty"`
}

// NewApiApplicationEvaluationCommitDTO instantiates a new ApiApplicationEvaluationCommitDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiApplicationEvaluationCommitDTO() *ApiApplicationEvaluationCommitDTO {
	this := ApiApplicationEvaluationCommitDTO{}
	return &this
}

// NewApiApplicationEvaluationCommitDTOWithDefaults instantiates a new ApiApplicationEvaluationCommitDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiApplicationEvaluationCommitDTOWithDefaults() *ApiApplicationEvaluationCommitDTO {
	this := ApiApplicationEvaluationCommitDTO{}
	return &this
}

// GetCommitHash returns the CommitHash field value if set, zero value otherwise.
func (o *ApiApplicationEvaluationCommitDTO) GetCommitHash() string {
	if o == nil || IsNil(o.CommitHash) {
		var ret string
		return ret
	}
	return *o.CommitHash
}

// GetCommitHashOk returns a tuple with the CommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationEvaluationCommitDTO) GetCommitHashOk() (*string, bool) {
	if o == nil || IsNil(o.CommitHash) {
		return nil, false
	}
	return o.CommitHash, true
}

// HasCommitHash returns a boolean if a field has been set.
func (o *ApiApplicationEvaluationCommitDTO) HasCommitHash() bool {
	if o != nil && !IsNil(o.CommitHash) {
		return true
	}

	return false
}

// SetCommitHash gets a reference to the given string and assigns it to the CommitHash field.
func (o *ApiApplicationEvaluationCommitDTO) SetCommitHash(v string) {
	o.CommitHash = &v
}

// GetReportUrl returns the ReportUrl field value if set, zero value otherwise.
func (o *ApiApplicationEvaluationCommitDTO) GetReportUrl() string {
	if o == nil || IsNil(o.ReportUrl) {
		var ret string
		return ret
	}
	return *o.ReportUrl
}

// GetReportUrlOk returns a tuple with the ReportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationEvaluationCommitDTO) GetReportUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportUrl) {
		return nil, false
	}
	return o.ReportUrl, true
}

// HasReportUrl returns a boolean if a field has been set.
func (o *ApiApplicationEvaluationCommitDTO) HasReportUrl() bool {
	if o != nil && !IsNil(o.ReportUrl) {
		return true
	}

	return false
}

// SetReportUrl gets a reference to the given string and assigns it to the ReportUrl field.
func (o *ApiApplicationEvaluationCommitDTO) SetReportUrl(v string) {
	o.ReportUrl = &v
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *ApiApplicationEvaluationCommitDTO) GetScanId() string {
	if o == nil || IsNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationEvaluationCommitDTO) GetScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *ApiApplicationEvaluationCommitDTO) HasScanId() bool {
	if o != nil && !IsNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *ApiApplicationEvaluationCommitDTO) SetScanId(v string) {
	o.ScanId = &v
}

// GetScanTime returns the ScanTime field value if set, zero value otherwise.
func (o *ApiApplicationEvaluationCommitDTO) GetScanTime() time.Time {
	if o == nil || IsNil(o.ScanTime) {
		var ret time.Time
		return ret
	}
	return *o.ScanTime
}

// GetScanTimeOk returns a tuple with the ScanTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationEvaluationCommitDTO) GetScanTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScanTime) {
		return nil, false
	}
	return o.ScanTime, true
}

// HasScanTime returns a boolean if a field has been set.
func (o *ApiApplicationEvaluationCommitDTO) HasScanTime() bool {
	if o != nil && !IsNil(o.ScanTime) {
		return true
	}

	return false
}

// SetScanTime gets a reference to the given time.Time and assigns it to the ScanTime field.
func (o *ApiApplicationEvaluationCommitDTO) SetScanTime(v time.Time) {
	o.ScanTime = &v
}

func (o ApiApplicationEvaluationCommitDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiApplicationEvaluationCommitDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommitHash) {
		toSerialize["commitHash"] = o.CommitHash
	}
	if !IsNil(o.ReportUrl) {
		toSerialize["reportUrl"] = o.ReportUrl
	}
	if !IsNil(o.ScanId) {
		toSerialize["scanId"] = o.ScanId
	}
	if !IsNil(o.ScanTime) {
		toSerialize["scanTime"] = o.ScanTime
	}
	return toSerialize, nil
}

type NullableApiApplicationEvaluationCommitDTO struct {
	value *ApiApplicationEvaluationCommitDTO
	isSet bool
}

func (v NullableApiApplicationEvaluationCommitDTO) Get() *ApiApplicationEvaluationCommitDTO {
	return v.value
}

func (v *NullableApiApplicationEvaluationCommitDTO) Set(val *ApiApplicationEvaluationCommitDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiApplicationEvaluationCommitDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiApplicationEvaluationCommitDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiApplicationEvaluationCommitDTO(val *ApiApplicationEvaluationCommitDTO) *NullableApiApplicationEvaluationCommitDTO {
	return &NullableApiApplicationEvaluationCommitDTO{value: val, isSet: true}
}

func (v NullableApiApplicationEvaluationCommitDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiApplicationEvaluationCommitDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


