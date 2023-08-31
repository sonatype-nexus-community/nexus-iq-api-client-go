/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.166.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the InnerSourceData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InnerSourceData{}

// InnerSourceData struct for InnerSourceData
type InnerSourceData struct {
	InnerSourceComponentPurl *string `json:"innerSourceComponentPurl,omitempty"`
	OwnerApplicationId *string `json:"ownerApplicationId,omitempty"`
	OwnerApplicationName *string `json:"ownerApplicationName,omitempty"`
}

// NewInnerSourceData instantiates a new InnerSourceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInnerSourceData() *InnerSourceData {
	this := InnerSourceData{}
	return &this
}

// NewInnerSourceDataWithDefaults instantiates a new InnerSourceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInnerSourceDataWithDefaults() *InnerSourceData {
	this := InnerSourceData{}
	return &this
}

// GetInnerSourceComponentPurl returns the InnerSourceComponentPurl field value if set, zero value otherwise.
func (o *InnerSourceData) GetInnerSourceComponentPurl() string {
	if o == nil || IsNil(o.InnerSourceComponentPurl) {
		var ret string
		return ret
	}
	return *o.InnerSourceComponentPurl
}

// GetInnerSourceComponentPurlOk returns a tuple with the InnerSourceComponentPurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InnerSourceData) GetInnerSourceComponentPurlOk() (*string, bool) {
	if o == nil || IsNil(o.InnerSourceComponentPurl) {
		return nil, false
	}
	return o.InnerSourceComponentPurl, true
}

// HasInnerSourceComponentPurl returns a boolean if a field has been set.
func (o *InnerSourceData) HasInnerSourceComponentPurl() bool {
	if o != nil && !IsNil(o.InnerSourceComponentPurl) {
		return true
	}

	return false
}

// SetInnerSourceComponentPurl gets a reference to the given string and assigns it to the InnerSourceComponentPurl field.
func (o *InnerSourceData) SetInnerSourceComponentPurl(v string) {
	o.InnerSourceComponentPurl = &v
}

// GetOwnerApplicationId returns the OwnerApplicationId field value if set, zero value otherwise.
func (o *InnerSourceData) GetOwnerApplicationId() string {
	if o == nil || IsNil(o.OwnerApplicationId) {
		var ret string
		return ret
	}
	return *o.OwnerApplicationId
}

// GetOwnerApplicationIdOk returns a tuple with the OwnerApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InnerSourceData) GetOwnerApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerApplicationId) {
		return nil, false
	}
	return o.OwnerApplicationId, true
}

// HasOwnerApplicationId returns a boolean if a field has been set.
func (o *InnerSourceData) HasOwnerApplicationId() bool {
	if o != nil && !IsNil(o.OwnerApplicationId) {
		return true
	}

	return false
}

// SetOwnerApplicationId gets a reference to the given string and assigns it to the OwnerApplicationId field.
func (o *InnerSourceData) SetOwnerApplicationId(v string) {
	o.OwnerApplicationId = &v
}

// GetOwnerApplicationName returns the OwnerApplicationName field value if set, zero value otherwise.
func (o *InnerSourceData) GetOwnerApplicationName() string {
	if o == nil || IsNil(o.OwnerApplicationName) {
		var ret string
		return ret
	}
	return *o.OwnerApplicationName
}

// GetOwnerApplicationNameOk returns a tuple with the OwnerApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InnerSourceData) GetOwnerApplicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerApplicationName) {
		return nil, false
	}
	return o.OwnerApplicationName, true
}

// HasOwnerApplicationName returns a boolean if a field has been set.
func (o *InnerSourceData) HasOwnerApplicationName() bool {
	if o != nil && !IsNil(o.OwnerApplicationName) {
		return true
	}

	return false
}

// SetOwnerApplicationName gets a reference to the given string and assigns it to the OwnerApplicationName field.
func (o *InnerSourceData) SetOwnerApplicationName(v string) {
	o.OwnerApplicationName = &v
}

func (o InnerSourceData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InnerSourceData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InnerSourceComponentPurl) {
		toSerialize["innerSourceComponentPurl"] = o.InnerSourceComponentPurl
	}
	if !IsNil(o.OwnerApplicationId) {
		toSerialize["ownerApplicationId"] = o.OwnerApplicationId
	}
	if !IsNil(o.OwnerApplicationName) {
		toSerialize["ownerApplicationName"] = o.OwnerApplicationName
	}
	return toSerialize, nil
}

type NullableInnerSourceData struct {
	value *InnerSourceData
	isSet bool
}

func (v NullableInnerSourceData) Get() *InnerSourceData {
	return v.value
}

func (v *NullableInnerSourceData) Set(val *InnerSourceData) {
	v.value = val
	v.isSet = true
}

func (v NullableInnerSourceData) IsSet() bool {
	return v.isSet
}

func (v *NullableInnerSourceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInnerSourceData(val *InnerSourceData) *NullableInnerSourceData {
	return &NullableInnerSourceData{value: val, isSet: true}
}

func (v NullableInnerSourceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInnerSourceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


