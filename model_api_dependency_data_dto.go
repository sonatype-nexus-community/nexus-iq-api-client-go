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

// checks if the ApiDependencyDataDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDependencyDataDTO{}

// ApiDependencyDataDTO struct for ApiDependencyDataDTO
type ApiDependencyDataDTO struct {
	DirectDependency *bool `json:"directDependency,omitempty"`
	InnerSource *bool `json:"innerSource,omitempty"`
	InnerSourceData []InnerSourceData `json:"innerSourceData,omitempty"`
	ParentComponentPurls []string `json:"parentComponentPurls,omitempty"`
}

// NewApiDependencyDataDTO instantiates a new ApiDependencyDataDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDependencyDataDTO() *ApiDependencyDataDTO {
	this := ApiDependencyDataDTO{}
	return &this
}

// NewApiDependencyDataDTOWithDefaults instantiates a new ApiDependencyDataDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDependencyDataDTOWithDefaults() *ApiDependencyDataDTO {
	this := ApiDependencyDataDTO{}
	return &this
}

// GetDirectDependency returns the DirectDependency field value if set, zero value otherwise.
func (o *ApiDependencyDataDTO) GetDirectDependency() bool {
	if o == nil || IsNil(o.DirectDependency) {
		var ret bool
		return ret
	}
	return *o.DirectDependency
}

// GetDirectDependencyOk returns a tuple with the DirectDependency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDependencyDataDTO) GetDirectDependencyOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectDependency) {
		return nil, false
	}
	return o.DirectDependency, true
}

// HasDirectDependency returns a boolean if a field has been set.
func (o *ApiDependencyDataDTO) HasDirectDependency() bool {
	if o != nil && !IsNil(o.DirectDependency) {
		return true
	}

	return false
}

// SetDirectDependency gets a reference to the given bool and assigns it to the DirectDependency field.
func (o *ApiDependencyDataDTO) SetDirectDependency(v bool) {
	o.DirectDependency = &v
}

// GetInnerSource returns the InnerSource field value if set, zero value otherwise.
func (o *ApiDependencyDataDTO) GetInnerSource() bool {
	if o == nil || IsNil(o.InnerSource) {
		var ret bool
		return ret
	}
	return *o.InnerSource
}

// GetInnerSourceOk returns a tuple with the InnerSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDependencyDataDTO) GetInnerSourceOk() (*bool, bool) {
	if o == nil || IsNil(o.InnerSource) {
		return nil, false
	}
	return o.InnerSource, true
}

// HasInnerSource returns a boolean if a field has been set.
func (o *ApiDependencyDataDTO) HasInnerSource() bool {
	if o != nil && !IsNil(o.InnerSource) {
		return true
	}

	return false
}

// SetInnerSource gets a reference to the given bool and assigns it to the InnerSource field.
func (o *ApiDependencyDataDTO) SetInnerSource(v bool) {
	o.InnerSource = &v
}

// GetInnerSourceData returns the InnerSourceData field value if set, zero value otherwise.
func (o *ApiDependencyDataDTO) GetInnerSourceData() []InnerSourceData {
	if o == nil || IsNil(o.InnerSourceData) {
		var ret []InnerSourceData
		return ret
	}
	return o.InnerSourceData
}

// GetInnerSourceDataOk returns a tuple with the InnerSourceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDependencyDataDTO) GetInnerSourceDataOk() ([]InnerSourceData, bool) {
	if o == nil || IsNil(o.InnerSourceData) {
		return nil, false
	}
	return o.InnerSourceData, true
}

// HasInnerSourceData returns a boolean if a field has been set.
func (o *ApiDependencyDataDTO) HasInnerSourceData() bool {
	if o != nil && !IsNil(o.InnerSourceData) {
		return true
	}

	return false
}

// SetInnerSourceData gets a reference to the given []InnerSourceData and assigns it to the InnerSourceData field.
func (o *ApiDependencyDataDTO) SetInnerSourceData(v []InnerSourceData) {
	o.InnerSourceData = v
}

// GetParentComponentPurls returns the ParentComponentPurls field value if set, zero value otherwise.
func (o *ApiDependencyDataDTO) GetParentComponentPurls() []string {
	if o == nil || IsNil(o.ParentComponentPurls) {
		var ret []string
		return ret
	}
	return o.ParentComponentPurls
}

// GetParentComponentPurlsOk returns a tuple with the ParentComponentPurls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDependencyDataDTO) GetParentComponentPurlsOk() ([]string, bool) {
	if o == nil || IsNil(o.ParentComponentPurls) {
		return nil, false
	}
	return o.ParentComponentPurls, true
}

// HasParentComponentPurls returns a boolean if a field has been set.
func (o *ApiDependencyDataDTO) HasParentComponentPurls() bool {
	if o != nil && !IsNil(o.ParentComponentPurls) {
		return true
	}

	return false
}

// SetParentComponentPurls gets a reference to the given []string and assigns it to the ParentComponentPurls field.
func (o *ApiDependencyDataDTO) SetParentComponentPurls(v []string) {
	o.ParentComponentPurls = v
}

func (o ApiDependencyDataDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDependencyDataDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DirectDependency) {
		toSerialize["directDependency"] = o.DirectDependency
	}
	if !IsNil(o.InnerSource) {
		toSerialize["innerSource"] = o.InnerSource
	}
	if !IsNil(o.InnerSourceData) {
		toSerialize["innerSourceData"] = o.InnerSourceData
	}
	if !IsNil(o.ParentComponentPurls) {
		toSerialize["parentComponentPurls"] = o.ParentComponentPurls
	}
	return toSerialize, nil
}

type NullableApiDependencyDataDTO struct {
	value *ApiDependencyDataDTO
	isSet bool
}

func (v NullableApiDependencyDataDTO) Get() *ApiDependencyDataDTO {
	return v.value
}

func (v *NullableApiDependencyDataDTO) Set(val *ApiDependencyDataDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDependencyDataDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDependencyDataDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDependencyDataDTO(val *ApiDependencyDataDTO) *NullableApiDependencyDataDTO {
	return &NullableApiDependencyDataDTO{value: val, isSet: true}
}

func (v NullableApiDependencyDataDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDependencyDataDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


