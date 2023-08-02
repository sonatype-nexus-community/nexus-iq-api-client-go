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

// checks if the ParameterizedHeader type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ParameterizedHeader{}

// ParameterizedHeader struct for ParameterizedHeader
type ParameterizedHeader struct {
	Parameters *map[string]string `json:"parameters,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewParameterizedHeader instantiates a new ParameterizedHeader object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParameterizedHeader() *ParameterizedHeader {
	this := ParameterizedHeader{}
	return &this
}

// NewParameterizedHeaderWithDefaults instantiates a new ParameterizedHeader object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParameterizedHeaderWithDefaults() *ParameterizedHeader {
	this := ParameterizedHeader{}
	return &this
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ParameterizedHeader) GetParameters() map[string]string {
	if o == nil || IsNil(o.Parameters) {
		var ret map[string]string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterizedHeader) GetParametersOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Parameters) {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ParameterizedHeader) HasParameters() bool {
	if o != nil && !IsNil(o.Parameters) {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]string and assigns it to the Parameters field.
func (o *ParameterizedHeader) SetParameters(v map[string]string) {
	o.Parameters = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ParameterizedHeader) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ParameterizedHeader) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ParameterizedHeader) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ParameterizedHeader) SetValue(v string) {
	o.Value = &v
}

func (o ParameterizedHeader) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ParameterizedHeader) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parameters) {
		toSerialize["parameters"] = o.Parameters
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableParameterizedHeader struct {
	value *ParameterizedHeader
	isSet bool
}

func (v NullableParameterizedHeader) Get() *ParameterizedHeader {
	return v.value
}

func (v *NullableParameterizedHeader) Set(val *ParameterizedHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableParameterizedHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableParameterizedHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParameterizedHeader(val *ParameterizedHeader) *NullableParameterizedHeader {
	return &NullableParameterizedHeader{value: val, isSet: true}
}

func (v NullableParameterizedHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParameterizedHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

