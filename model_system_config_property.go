/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.184.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"fmt"
)

// SystemConfigProperty the model 'SystemConfigProperty'
type SystemConfigProperty string

// List of SystemConfigProperty
const (
	BASE_URL SystemConfigProperty = "baseUrl"
	FORCE_BASE_URL SystemConfigProperty = "forceBaseUrl"
)

// All allowed values of SystemConfigProperty enum
var AllowedSystemConfigPropertyEnumValues = []SystemConfigProperty{
	"baseUrl",
	"forceBaseUrl",
}

func (v *SystemConfigProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SystemConfigProperty(value)
	for _, existing := range AllowedSystemConfigPropertyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SystemConfigProperty", value)
}

// NewSystemConfigPropertyFromValue returns a pointer to a valid SystemConfigProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSystemConfigPropertyFromValue(v string) (*SystemConfigProperty, error) {
	ev := SystemConfigProperty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SystemConfigProperty: valid values are %v", v, AllowedSystemConfigPropertyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SystemConfigProperty) IsValid() bool {
	for _, existing := range AllowedSystemConfigPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SystemConfigProperty value
func (v SystemConfigProperty) Ptr() *SystemConfigProperty {
	return &v
}

type NullableSystemConfigProperty struct {
	value *SystemConfigProperty
	isSet bool
}

func (v NullableSystemConfigProperty) Get() *SystemConfigProperty {
	return v.value
}

func (v *NullableSystemConfigProperty) Set(val *SystemConfigProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemConfigProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemConfigProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemConfigProperty(val *SystemConfigProperty) *NullableSystemConfigProperty {
	return &NullableSystemConfigProperty{value: val, isSet: true}
}

func (v NullableSystemConfigProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemConfigProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

