/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 164
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the StageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StageData{}

// StageData struct for StageData
type StageData struct {
	ActionTypeId *string `json:"actionTypeId,omitempty"`
	MostRecentEvaluationTime *time.Time `json:"mostRecentEvaluationTime,omitempty"`
	MostRecentScanId *string `json:"mostRecentScanId,omitempty"`
}

// NewStageData instantiates a new StageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStageData() *StageData {
	this := StageData{}
	return &this
}

// NewStageDataWithDefaults instantiates a new StageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStageDataWithDefaults() *StageData {
	this := StageData{}
	return &this
}

// GetActionTypeId returns the ActionTypeId field value if set, zero value otherwise.
func (o *StageData) GetActionTypeId() string {
	if o == nil || IsNil(o.ActionTypeId) {
		var ret string
		return ret
	}
	return *o.ActionTypeId
}

// GetActionTypeIdOk returns a tuple with the ActionTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageData) GetActionTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ActionTypeId) {
		return nil, false
	}
	return o.ActionTypeId, true
}

// HasActionTypeId returns a boolean if a field has been set.
func (o *StageData) HasActionTypeId() bool {
	if o != nil && !IsNil(o.ActionTypeId) {
		return true
	}

	return false
}

// SetActionTypeId gets a reference to the given string and assigns it to the ActionTypeId field.
func (o *StageData) SetActionTypeId(v string) {
	o.ActionTypeId = &v
}

// GetMostRecentEvaluationTime returns the MostRecentEvaluationTime field value if set, zero value otherwise.
func (o *StageData) GetMostRecentEvaluationTime() time.Time {
	if o == nil || IsNil(o.MostRecentEvaluationTime) {
		var ret time.Time
		return ret
	}
	return *o.MostRecentEvaluationTime
}

// GetMostRecentEvaluationTimeOk returns a tuple with the MostRecentEvaluationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageData) GetMostRecentEvaluationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MostRecentEvaluationTime) {
		return nil, false
	}
	return o.MostRecentEvaluationTime, true
}

// HasMostRecentEvaluationTime returns a boolean if a field has been set.
func (o *StageData) HasMostRecentEvaluationTime() bool {
	if o != nil && !IsNil(o.MostRecentEvaluationTime) {
		return true
	}

	return false
}

// SetMostRecentEvaluationTime gets a reference to the given time.Time and assigns it to the MostRecentEvaluationTime field.
func (o *StageData) SetMostRecentEvaluationTime(v time.Time) {
	o.MostRecentEvaluationTime = &v
}

// GetMostRecentScanId returns the MostRecentScanId field value if set, zero value otherwise.
func (o *StageData) GetMostRecentScanId() string {
	if o == nil || IsNil(o.MostRecentScanId) {
		var ret string
		return ret
	}
	return *o.MostRecentScanId
}

// GetMostRecentScanIdOk returns a tuple with the MostRecentScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StageData) GetMostRecentScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.MostRecentScanId) {
		return nil, false
	}
	return o.MostRecentScanId, true
}

// HasMostRecentScanId returns a boolean if a field has been set.
func (o *StageData) HasMostRecentScanId() bool {
	if o != nil && !IsNil(o.MostRecentScanId) {
		return true
	}

	return false
}

// SetMostRecentScanId gets a reference to the given string and assigns it to the MostRecentScanId field.
func (o *StageData) SetMostRecentScanId(v string) {
	o.MostRecentScanId = &v
}

func (o StageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionTypeId) {
		toSerialize["actionTypeId"] = o.ActionTypeId
	}
	if !IsNil(o.MostRecentEvaluationTime) {
		toSerialize["mostRecentEvaluationTime"] = o.MostRecentEvaluationTime
	}
	if !IsNil(o.MostRecentScanId) {
		toSerialize["mostRecentScanId"] = o.MostRecentScanId
	}
	return toSerialize, nil
}

type NullableStageData struct {
	value *StageData
	isSet bool
}

func (v NullableStageData) Get() *StageData {
	return v.value
}

func (v *NullableStageData) Set(val *StageData) {
	v.value = val
	v.isSet = true
}

func (v NullableStageData) IsSet() bool {
	return v.isSet
}

func (v *NullableStageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStageData(val *StageData) *NullableStageData {
	return &NullableStageData{value: val, isSet: true}
}

func (v NullableStageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


