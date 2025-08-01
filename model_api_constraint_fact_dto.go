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

// checks if the ApiConstraintFactDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiConstraintFactDTO{}

// ApiConstraintFactDTO struct for ApiConstraintFactDTO
type ApiConstraintFactDTO struct {
	ConstraintId *string `json:"constraintId,omitempty"`
	ConstraintName *string `json:"constraintName,omitempty"`
	Reasons []ApiConditionFactReasonDTO `json:"reasons,omitempty"`
}

// NewApiConstraintFactDTO instantiates a new ApiConstraintFactDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiConstraintFactDTO() *ApiConstraintFactDTO {
	this := ApiConstraintFactDTO{}
	return &this
}

// NewApiConstraintFactDTOWithDefaults instantiates a new ApiConstraintFactDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiConstraintFactDTOWithDefaults() *ApiConstraintFactDTO {
	this := ApiConstraintFactDTO{}
	return &this
}

// GetConstraintId returns the ConstraintId field value if set, zero value otherwise.
func (o *ApiConstraintFactDTO) GetConstraintId() string {
	if o == nil || IsNil(o.ConstraintId) {
		var ret string
		return ret
	}
	return *o.ConstraintId
}

// GetConstraintIdOk returns a tuple with the ConstraintId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConstraintFactDTO) GetConstraintIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConstraintId) {
		return nil, false
	}
	return o.ConstraintId, true
}

// HasConstraintId returns a boolean if a field has been set.
func (o *ApiConstraintFactDTO) HasConstraintId() bool {
	if o != nil && !IsNil(o.ConstraintId) {
		return true
	}

	return false
}

// SetConstraintId gets a reference to the given string and assigns it to the ConstraintId field.
func (o *ApiConstraintFactDTO) SetConstraintId(v string) {
	o.ConstraintId = &v
}

// GetConstraintName returns the ConstraintName field value if set, zero value otherwise.
func (o *ApiConstraintFactDTO) GetConstraintName() string {
	if o == nil || IsNil(o.ConstraintName) {
		var ret string
		return ret
	}
	return *o.ConstraintName
}

// GetConstraintNameOk returns a tuple with the ConstraintName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConstraintFactDTO) GetConstraintNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConstraintName) {
		return nil, false
	}
	return o.ConstraintName, true
}

// HasConstraintName returns a boolean if a field has been set.
func (o *ApiConstraintFactDTO) HasConstraintName() bool {
	if o != nil && !IsNil(o.ConstraintName) {
		return true
	}

	return false
}

// SetConstraintName gets a reference to the given string and assigns it to the ConstraintName field.
func (o *ApiConstraintFactDTO) SetConstraintName(v string) {
	o.ConstraintName = &v
}

// GetReasons returns the Reasons field value if set, zero value otherwise.
func (o *ApiConstraintFactDTO) GetReasons() []ApiConditionFactReasonDTO {
	if o == nil || IsNil(o.Reasons) {
		var ret []ApiConditionFactReasonDTO
		return ret
	}
	return o.Reasons
}

// GetReasonsOk returns a tuple with the Reasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiConstraintFactDTO) GetReasonsOk() ([]ApiConditionFactReasonDTO, bool) {
	if o == nil || IsNil(o.Reasons) {
		return nil, false
	}
	return o.Reasons, true
}

// HasReasons returns a boolean if a field has been set.
func (o *ApiConstraintFactDTO) HasReasons() bool {
	if o != nil && !IsNil(o.Reasons) {
		return true
	}

	return false
}

// SetReasons gets a reference to the given []ApiConditionFactReasonDTO and assigns it to the Reasons field.
func (o *ApiConstraintFactDTO) SetReasons(v []ApiConditionFactReasonDTO) {
	o.Reasons = v
}

func (o ApiConstraintFactDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiConstraintFactDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConstraintId) {
		toSerialize["constraintId"] = o.ConstraintId
	}
	if !IsNil(o.ConstraintName) {
		toSerialize["constraintName"] = o.ConstraintName
	}
	if !IsNil(o.Reasons) {
		toSerialize["reasons"] = o.Reasons
	}
	return toSerialize, nil
}

type NullableApiConstraintFactDTO struct {
	value *ApiConstraintFactDTO
	isSet bool
}

func (v NullableApiConstraintFactDTO) Get() *ApiConstraintFactDTO {
	return v.value
}

func (v *NullableApiConstraintFactDTO) Set(val *ApiConstraintFactDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiConstraintFactDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiConstraintFactDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiConstraintFactDTO(val *ApiConstraintFactDTO) *NullableApiConstraintFactDTO {
	return &NullableApiConstraintFactDTO{value: val, isSet: true}
}

func (v NullableApiConstraintFactDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiConstraintFactDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


