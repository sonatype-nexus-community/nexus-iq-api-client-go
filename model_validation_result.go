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

// checks if the ValidationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationResult{}

// ValidationResult struct for ValidationResult
type ValidationResult struct {
	Message *string `json:"message,omitempty"`
	Valid *bool `json:"valid,omitempty"`
}

// NewValidationResult instantiates a new ValidationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationResult() *ValidationResult {
	this := ValidationResult{}
	return &this
}

// NewValidationResultWithDefaults instantiates a new ValidationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationResultWithDefaults() *ValidationResult {
	this := ValidationResult{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ValidationResult) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationResult) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ValidationResult) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ValidationResult) SetMessage(v string) {
	o.Message = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *ValidationResult) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationResult) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *ValidationResult) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *ValidationResult) SetValid(v bool) {
	o.Valid = &v
}

func (o ValidationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	return toSerialize, nil
}

type NullableValidationResult struct {
	value *ValidationResult
	isSet bool
}

func (v NullableValidationResult) Get() *ValidationResult {
	return v.value
}

func (v *NullableValidationResult) Set(val *ValidationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationResult(val *ValidationResult) *NullableValidationResult {
	return &NullableValidationResult{value: val, isSet: true}
}

func (v NullableValidationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

