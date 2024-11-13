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

// checks if the ApiComponentDetailsResultDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentDetailsResultDTOV2{}

// ApiComponentDetailsResultDTOV2 struct for ApiComponentDetailsResultDTOV2
type ApiComponentDetailsResultDTOV2 struct {
	ComponentDetails []ApiComponentDetailsDTOV2 `json:"componentDetails,omitempty"`
}

// NewApiComponentDetailsResultDTOV2 instantiates a new ApiComponentDetailsResultDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentDetailsResultDTOV2() *ApiComponentDetailsResultDTOV2 {
	this := ApiComponentDetailsResultDTOV2{}
	return &this
}

// NewApiComponentDetailsResultDTOV2WithDefaults instantiates a new ApiComponentDetailsResultDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentDetailsResultDTOV2WithDefaults() *ApiComponentDetailsResultDTOV2 {
	this := ApiComponentDetailsResultDTOV2{}
	return &this
}

// GetComponentDetails returns the ComponentDetails field value if set, zero value otherwise.
func (o *ApiComponentDetailsResultDTOV2) GetComponentDetails() []ApiComponentDetailsDTOV2 {
	if o == nil || IsNil(o.ComponentDetails) {
		var ret []ApiComponentDetailsDTOV2
		return ret
	}
	return o.ComponentDetails
}

// GetComponentDetailsOk returns a tuple with the ComponentDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentDetailsResultDTOV2) GetComponentDetailsOk() ([]ApiComponentDetailsDTOV2, bool) {
	if o == nil || IsNil(o.ComponentDetails) {
		return nil, false
	}
	return o.ComponentDetails, true
}

// HasComponentDetails returns a boolean if a field has been set.
func (o *ApiComponentDetailsResultDTOV2) HasComponentDetails() bool {
	if o != nil && !IsNil(o.ComponentDetails) {
		return true
	}

	return false
}

// SetComponentDetails gets a reference to the given []ApiComponentDetailsDTOV2 and assigns it to the ComponentDetails field.
func (o *ApiComponentDetailsResultDTOV2) SetComponentDetails(v []ApiComponentDetailsDTOV2) {
	o.ComponentDetails = v
}

func (o ApiComponentDetailsResultDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentDetailsResultDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentDetails) {
		toSerialize["componentDetails"] = o.ComponentDetails
	}
	return toSerialize, nil
}

type NullableApiComponentDetailsResultDTOV2 struct {
	value *ApiComponentDetailsResultDTOV2
	isSet bool
}

func (v NullableApiComponentDetailsResultDTOV2) Get() *ApiComponentDetailsResultDTOV2 {
	return v.value
}

func (v *NullableApiComponentDetailsResultDTOV2) Set(val *ApiComponentDetailsResultDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentDetailsResultDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentDetailsResultDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentDetailsResultDTOV2(val *ApiComponentDetailsResultDTOV2) *NullableApiComponentDetailsResultDTOV2 {
	return &NullableApiComponentDetailsResultDTOV2{value: val, isSet: true}
}

func (v NullableApiComponentDetailsResultDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentDetailsResultDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


