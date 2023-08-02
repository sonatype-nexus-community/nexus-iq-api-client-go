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

// checks if the ComponentIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentIdentifier{}

// ComponentIdentifier struct for ComponentIdentifier
type ComponentIdentifier struct {
	Coordinates *map[string]string `json:"coordinates,omitempty"`
	Format *string `json:"format,omitempty"`
}

// NewComponentIdentifier instantiates a new ComponentIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentIdentifier() *ComponentIdentifier {
	this := ComponentIdentifier{}
	return &this
}

// NewComponentIdentifierWithDefaults instantiates a new ComponentIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentIdentifierWithDefaults() *ComponentIdentifier {
	this := ComponentIdentifier{}
	return &this
}

// GetCoordinates returns the Coordinates field value if set, zero value otherwise.
func (o *ComponentIdentifier) GetCoordinates() map[string]string {
	if o == nil || IsNil(o.Coordinates) {
		var ret map[string]string
		return ret
	}
	return *o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentIdentifier) GetCoordinatesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Coordinates) {
		return nil, false
	}
	return o.Coordinates, true
}

// HasCoordinates returns a boolean if a field has been set.
func (o *ComponentIdentifier) HasCoordinates() bool {
	if o != nil && !IsNil(o.Coordinates) {
		return true
	}

	return false
}

// SetCoordinates gets a reference to the given map[string]string and assigns it to the Coordinates field.
func (o *ComponentIdentifier) SetCoordinates(v map[string]string) {
	o.Coordinates = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ComponentIdentifier) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentIdentifier) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ComponentIdentifier) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ComponentIdentifier) SetFormat(v string) {
	o.Format = &v
}

func (o ComponentIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Coordinates) {
		toSerialize["coordinates"] = o.Coordinates
	}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	return toSerialize, nil
}

type NullableComponentIdentifier struct {
	value *ComponentIdentifier
	isSet bool
}

func (v NullableComponentIdentifier) Get() *ComponentIdentifier {
	return v.value
}

func (v *NullableComponentIdentifier) Set(val *ComponentIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentIdentifier(val *ComponentIdentifier) *NullableComponentIdentifier {
	return &NullableComponentIdentifier{value: val, isSet: true}
}

func (v NullableComponentIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

