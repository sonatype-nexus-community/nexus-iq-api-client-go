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

// checks if the ApiReportConstraintViolationDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReportConstraintViolationDTOV2{}

// ApiReportConstraintViolationDTOV2 struct for ApiReportConstraintViolationDTOV2
type ApiReportConstraintViolationDTOV2 struct {
	Conditions []ApiReportConstraintConditionDTOV2 `json:"conditions,omitempty"`
	ConstraintId *string `json:"constraintId,omitempty"`
	ConstraintName *string `json:"constraintName,omitempty"`
}

// NewApiReportConstraintViolationDTOV2 instantiates a new ApiReportConstraintViolationDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReportConstraintViolationDTOV2() *ApiReportConstraintViolationDTOV2 {
	this := ApiReportConstraintViolationDTOV2{}
	return &this
}

// NewApiReportConstraintViolationDTOV2WithDefaults instantiates a new ApiReportConstraintViolationDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReportConstraintViolationDTOV2WithDefaults() *ApiReportConstraintViolationDTOV2 {
	this := ApiReportConstraintViolationDTOV2{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *ApiReportConstraintViolationDTOV2) GetConditions() []ApiReportConstraintConditionDTOV2 {
	if o == nil || IsNil(o.Conditions) {
		var ret []ApiReportConstraintConditionDTOV2
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportConstraintViolationDTOV2) GetConditionsOk() ([]ApiReportConstraintConditionDTOV2, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *ApiReportConstraintViolationDTOV2) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []ApiReportConstraintConditionDTOV2 and assigns it to the Conditions field.
func (o *ApiReportConstraintViolationDTOV2) SetConditions(v []ApiReportConstraintConditionDTOV2) {
	o.Conditions = v
}

// GetConstraintId returns the ConstraintId field value if set, zero value otherwise.
func (o *ApiReportConstraintViolationDTOV2) GetConstraintId() string {
	if o == nil || IsNil(o.ConstraintId) {
		var ret string
		return ret
	}
	return *o.ConstraintId
}

// GetConstraintIdOk returns a tuple with the ConstraintId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportConstraintViolationDTOV2) GetConstraintIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConstraintId) {
		return nil, false
	}
	return o.ConstraintId, true
}

// HasConstraintId returns a boolean if a field has been set.
func (o *ApiReportConstraintViolationDTOV2) HasConstraintId() bool {
	if o != nil && !IsNil(o.ConstraintId) {
		return true
	}

	return false
}

// SetConstraintId gets a reference to the given string and assigns it to the ConstraintId field.
func (o *ApiReportConstraintViolationDTOV2) SetConstraintId(v string) {
	o.ConstraintId = &v
}

// GetConstraintName returns the ConstraintName field value if set, zero value otherwise.
func (o *ApiReportConstraintViolationDTOV2) GetConstraintName() string {
	if o == nil || IsNil(o.ConstraintName) {
		var ret string
		return ret
	}
	return *o.ConstraintName
}

// GetConstraintNameOk returns a tuple with the ConstraintName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportConstraintViolationDTOV2) GetConstraintNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConstraintName) {
		return nil, false
	}
	return o.ConstraintName, true
}

// HasConstraintName returns a boolean if a field has been set.
func (o *ApiReportConstraintViolationDTOV2) HasConstraintName() bool {
	if o != nil && !IsNil(o.ConstraintName) {
		return true
	}

	return false
}

// SetConstraintName gets a reference to the given string and assigns it to the ConstraintName field.
func (o *ApiReportConstraintViolationDTOV2) SetConstraintName(v string) {
	o.ConstraintName = &v
}

func (o ApiReportConstraintViolationDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReportConstraintViolationDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	if !IsNil(o.ConstraintId) {
		toSerialize["constraintId"] = o.ConstraintId
	}
	if !IsNil(o.ConstraintName) {
		toSerialize["constraintName"] = o.ConstraintName
	}
	return toSerialize, nil
}

type NullableApiReportConstraintViolationDTOV2 struct {
	value *ApiReportConstraintViolationDTOV2
	isSet bool
}

func (v NullableApiReportConstraintViolationDTOV2) Get() *ApiReportConstraintViolationDTOV2 {
	return v.value
}

func (v *NullableApiReportConstraintViolationDTOV2) Set(val *ApiReportConstraintViolationDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReportConstraintViolationDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReportConstraintViolationDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReportConstraintViolationDTOV2(val *ApiReportConstraintViolationDTOV2) *NullableApiReportConstraintViolationDTOV2 {
	return &NullableApiReportConstraintViolationDTOV2{value: val, isSet: true}
}

func (v NullableApiReportConstraintViolationDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReportConstraintViolationDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


