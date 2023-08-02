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

// checks if the ApiEvaluationResultCounterDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiEvaluationResultCounterDTO{}

// ApiEvaluationResultCounterDTO struct for ApiEvaluationResultCounterDTO
type ApiEvaluationResultCounterDTO struct {
	Critical *int32 `json:"critical,omitempty"`
	Moderate *int32 `json:"moderate,omitempty"`
	Severe *int32 `json:"severe,omitempty"`
}

// NewApiEvaluationResultCounterDTO instantiates a new ApiEvaluationResultCounterDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiEvaluationResultCounterDTO() *ApiEvaluationResultCounterDTO {
	this := ApiEvaluationResultCounterDTO{}
	return &this
}

// NewApiEvaluationResultCounterDTOWithDefaults instantiates a new ApiEvaluationResultCounterDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiEvaluationResultCounterDTOWithDefaults() *ApiEvaluationResultCounterDTO {
	this := ApiEvaluationResultCounterDTO{}
	return &this
}

// GetCritical returns the Critical field value if set, zero value otherwise.
func (o *ApiEvaluationResultCounterDTO) GetCritical() int32 {
	if o == nil || IsNil(o.Critical) {
		var ret int32
		return ret
	}
	return *o.Critical
}

// GetCriticalOk returns a tuple with the Critical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEvaluationResultCounterDTO) GetCriticalOk() (*int32, bool) {
	if o == nil || IsNil(o.Critical) {
		return nil, false
	}
	return o.Critical, true
}

// HasCritical returns a boolean if a field has been set.
func (o *ApiEvaluationResultCounterDTO) HasCritical() bool {
	if o != nil && !IsNil(o.Critical) {
		return true
	}

	return false
}

// SetCritical gets a reference to the given int32 and assigns it to the Critical field.
func (o *ApiEvaluationResultCounterDTO) SetCritical(v int32) {
	o.Critical = &v
}

// GetModerate returns the Moderate field value if set, zero value otherwise.
func (o *ApiEvaluationResultCounterDTO) GetModerate() int32 {
	if o == nil || IsNil(o.Moderate) {
		var ret int32
		return ret
	}
	return *o.Moderate
}

// GetModerateOk returns a tuple with the Moderate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEvaluationResultCounterDTO) GetModerateOk() (*int32, bool) {
	if o == nil || IsNil(o.Moderate) {
		return nil, false
	}
	return o.Moderate, true
}

// HasModerate returns a boolean if a field has been set.
func (o *ApiEvaluationResultCounterDTO) HasModerate() bool {
	if o != nil && !IsNil(o.Moderate) {
		return true
	}

	return false
}

// SetModerate gets a reference to the given int32 and assigns it to the Moderate field.
func (o *ApiEvaluationResultCounterDTO) SetModerate(v int32) {
	o.Moderate = &v
}

// GetSevere returns the Severe field value if set, zero value otherwise.
func (o *ApiEvaluationResultCounterDTO) GetSevere() int32 {
	if o == nil || IsNil(o.Severe) {
		var ret int32
		return ret
	}
	return *o.Severe
}

// GetSevereOk returns a tuple with the Severe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiEvaluationResultCounterDTO) GetSevereOk() (*int32, bool) {
	if o == nil || IsNil(o.Severe) {
		return nil, false
	}
	return o.Severe, true
}

// HasSevere returns a boolean if a field has been set.
func (o *ApiEvaluationResultCounterDTO) HasSevere() bool {
	if o != nil && !IsNil(o.Severe) {
		return true
	}

	return false
}

// SetSevere gets a reference to the given int32 and assigns it to the Severe field.
func (o *ApiEvaluationResultCounterDTO) SetSevere(v int32) {
	o.Severe = &v
}

func (o ApiEvaluationResultCounterDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiEvaluationResultCounterDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Critical) {
		toSerialize["critical"] = o.Critical
	}
	if !IsNil(o.Moderate) {
		toSerialize["moderate"] = o.Moderate
	}
	if !IsNil(o.Severe) {
		toSerialize["severe"] = o.Severe
	}
	return toSerialize, nil
}

type NullableApiEvaluationResultCounterDTO struct {
	value *ApiEvaluationResultCounterDTO
	isSet bool
}

func (v NullableApiEvaluationResultCounterDTO) Get() *ApiEvaluationResultCounterDTO {
	return v.value
}

func (v *NullableApiEvaluationResultCounterDTO) Set(val *ApiEvaluationResultCounterDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiEvaluationResultCounterDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiEvaluationResultCounterDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiEvaluationResultCounterDTO(val *ApiEvaluationResultCounterDTO) *NullableApiEvaluationResultCounterDTO {
	return &NullableApiEvaluationResultCounterDTO{value: val, isSet: true}
}

func (v NullableApiEvaluationResultCounterDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiEvaluationResultCounterDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

