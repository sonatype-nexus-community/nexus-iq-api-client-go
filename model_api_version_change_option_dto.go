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

// checks if the ApiVersionChangeOptionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiVersionChangeOptionDTO{}

// ApiVersionChangeOptionDTO struct for ApiVersionChangeOptionDTO
type ApiVersionChangeOptionDTO struct {
	Data *ApiComponentChangeActionDTO `json:"data,omitempty"`
	DirectDependency *bool `json:"directDependency,omitempty"`
	DirectDependencyData []ApiComponentChangeActionDTO `json:"directDependencyData,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewApiVersionChangeOptionDTO instantiates a new ApiVersionChangeOptionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiVersionChangeOptionDTO() *ApiVersionChangeOptionDTO {
	this := ApiVersionChangeOptionDTO{}
	return &this
}

// NewApiVersionChangeOptionDTOWithDefaults instantiates a new ApiVersionChangeOptionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiVersionChangeOptionDTOWithDefaults() *ApiVersionChangeOptionDTO {
	this := ApiVersionChangeOptionDTO{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApiVersionChangeOptionDTO) GetData() ApiComponentChangeActionDTO {
	if o == nil || IsNil(o.Data) {
		var ret ApiComponentChangeActionDTO
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVersionChangeOptionDTO) GetDataOk() (*ApiComponentChangeActionDTO, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApiVersionChangeOptionDTO) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ApiComponentChangeActionDTO and assigns it to the Data field.
func (o *ApiVersionChangeOptionDTO) SetData(v ApiComponentChangeActionDTO) {
	o.Data = &v
}

// GetDirectDependency returns the DirectDependency field value if set, zero value otherwise.
func (o *ApiVersionChangeOptionDTO) GetDirectDependency() bool {
	if o == nil || IsNil(o.DirectDependency) {
		var ret bool
		return ret
	}
	return *o.DirectDependency
}

// GetDirectDependencyOk returns a tuple with the DirectDependency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVersionChangeOptionDTO) GetDirectDependencyOk() (*bool, bool) {
	if o == nil || IsNil(o.DirectDependency) {
		return nil, false
	}
	return o.DirectDependency, true
}

// HasDirectDependency returns a boolean if a field has been set.
func (o *ApiVersionChangeOptionDTO) HasDirectDependency() bool {
	if o != nil && !IsNil(o.DirectDependency) {
		return true
	}

	return false
}

// SetDirectDependency gets a reference to the given bool and assigns it to the DirectDependency field.
func (o *ApiVersionChangeOptionDTO) SetDirectDependency(v bool) {
	o.DirectDependency = &v
}

// GetDirectDependencyData returns the DirectDependencyData field value if set, zero value otherwise.
func (o *ApiVersionChangeOptionDTO) GetDirectDependencyData() []ApiComponentChangeActionDTO {
	if o == nil || IsNil(o.DirectDependencyData) {
		var ret []ApiComponentChangeActionDTO
		return ret
	}
	return o.DirectDependencyData
}

// GetDirectDependencyDataOk returns a tuple with the DirectDependencyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVersionChangeOptionDTO) GetDirectDependencyDataOk() ([]ApiComponentChangeActionDTO, bool) {
	if o == nil || IsNil(o.DirectDependencyData) {
		return nil, false
	}
	return o.DirectDependencyData, true
}

// HasDirectDependencyData returns a boolean if a field has been set.
func (o *ApiVersionChangeOptionDTO) HasDirectDependencyData() bool {
	if o != nil && !IsNil(o.DirectDependencyData) {
		return true
	}

	return false
}

// SetDirectDependencyData gets a reference to the given []ApiComponentChangeActionDTO and assigns it to the DirectDependencyData field.
func (o *ApiVersionChangeOptionDTO) SetDirectDependencyData(v []ApiComponentChangeActionDTO) {
	o.DirectDependencyData = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiVersionChangeOptionDTO) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiVersionChangeOptionDTO) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiVersionChangeOptionDTO) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiVersionChangeOptionDTO) SetType(v string) {
	o.Type = &v
}

func (o ApiVersionChangeOptionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiVersionChangeOptionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.DirectDependency) {
		toSerialize["directDependency"] = o.DirectDependency
	}
	if !IsNil(o.DirectDependencyData) {
		toSerialize["directDependencyData"] = o.DirectDependencyData
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableApiVersionChangeOptionDTO struct {
	value *ApiVersionChangeOptionDTO
	isSet bool
}

func (v NullableApiVersionChangeOptionDTO) Get() *ApiVersionChangeOptionDTO {
	return v.value
}

func (v *NullableApiVersionChangeOptionDTO) Set(val *ApiVersionChangeOptionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiVersionChangeOptionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiVersionChangeOptionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiVersionChangeOptionDTO(val *ApiVersionChangeOptionDTO) *NullableApiVersionChangeOptionDTO {
	return &NullableApiVersionChangeOptionDTO{value: val, isSet: true}
}

func (v NullableApiVersionChangeOptionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiVersionChangeOptionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


