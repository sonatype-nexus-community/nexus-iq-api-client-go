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

// checks if the ApiConditionFactReasonDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiConditionFactReasonDTO{}

// ApiConditionFactReasonDTO struct for ApiConditionFactReasonDTO
type ApiConditionFactReasonDTO struct {
	Reason *string `json:"reason,omitempty"`
}

// NewApiConditionFactReasonDTO instantiates a new ApiConditionFactReasonDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiConditionFactReasonDTO() *ApiConditionFactReasonDTO {
	this := ApiConditionFactReasonDTO{}
	return &this
}

// NewApiConditionFactReasonDTOWithDefaults instantiates a new ApiConditionFactReasonDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiConditionFactReasonDTOWithDefaults() *ApiConditionFactReasonDTO {
	this := ApiConditionFactReasonDTO{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ApiConditionFactReasonDTO) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConditionFactReasonDTO) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ApiConditionFactReasonDTO) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ApiConditionFactReasonDTO) SetReason(v string) {
	o.Reason = &v
}

func (o ApiConditionFactReasonDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiConditionFactReasonDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableApiConditionFactReasonDTO struct {
	value *ApiConditionFactReasonDTO
	isSet bool
}

func (v NullableApiConditionFactReasonDTO) Get() *ApiConditionFactReasonDTO {
	return v.value
}

func (v *NullableApiConditionFactReasonDTO) Set(val *ApiConditionFactReasonDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiConditionFactReasonDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiConditionFactReasonDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiConditionFactReasonDTO(val *ApiConditionFactReasonDTO) *NullableApiConditionFactReasonDTO {
	return &NullableApiConditionFactReasonDTO{value: val, isSet: true}
}

func (v NullableApiConditionFactReasonDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiConditionFactReasonDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


