/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.184.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApplicableLabels type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicableLabels{}

// ApplicableLabels struct for ApplicableLabels
type ApplicableLabels struct {
	LabelsByOwner []LabelsByOwner `json:"labelsByOwner,omitempty"`
}

// NewApplicableLabels instantiates a new ApplicableLabels object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicableLabels() *ApplicableLabels {
	this := ApplicableLabels{}
	return &this
}

// NewApplicableLabelsWithDefaults instantiates a new ApplicableLabels object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicableLabelsWithDefaults() *ApplicableLabels {
	this := ApplicableLabels{}
	return &this
}

// GetLabelsByOwner returns the LabelsByOwner field value if set, zero value otherwise.
func (o *ApplicableLabels) GetLabelsByOwner() []LabelsByOwner {
	if o == nil || IsNil(o.LabelsByOwner) {
		var ret []LabelsByOwner
		return ret
	}
	return o.LabelsByOwner
}

// GetLabelsByOwnerOk returns a tuple with the LabelsByOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicableLabels) GetLabelsByOwnerOk() ([]LabelsByOwner, bool) {
	if o == nil || IsNil(o.LabelsByOwner) {
		return nil, false
	}
	return o.LabelsByOwner, true
}

// HasLabelsByOwner returns a boolean if a field has been set.
func (o *ApplicableLabels) HasLabelsByOwner() bool {
	if o != nil && !IsNil(o.LabelsByOwner) {
		return true
	}

	return false
}

// SetLabelsByOwner gets a reference to the given []LabelsByOwner and assigns it to the LabelsByOwner field.
func (o *ApplicableLabels) SetLabelsByOwner(v []LabelsByOwner) {
	o.LabelsByOwner = v
}

func (o ApplicableLabels) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicableLabels) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LabelsByOwner) {
		toSerialize["labelsByOwner"] = o.LabelsByOwner
	}
	return toSerialize, nil
}

type NullableApplicableLabels struct {
	value *ApplicableLabels
	isSet bool
}

func (v NullableApplicableLabels) Get() *ApplicableLabels {
	return v.value
}

func (v *NullableApplicableLabels) Set(val *ApplicableLabels) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicableLabels) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicableLabels) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicableLabels(val *ApplicableLabels) *NullableApplicableLabels {
	return &NullableApplicableLabels{value: val, isSet: true}
}

func (v NullableApplicableLabels) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicableLabels) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


