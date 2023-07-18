/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 164
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiPolicyListDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPolicyListDTO{}

// ApiPolicyListDTO struct for ApiPolicyListDTO
type ApiPolicyListDTO struct {
	Policies []ApiPolicyDTO `json:"policies,omitempty"`
}

// NewApiPolicyListDTO instantiates a new ApiPolicyListDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPolicyListDTO() *ApiPolicyListDTO {
	this := ApiPolicyListDTO{}
	return &this
}

// NewApiPolicyListDTOWithDefaults instantiates a new ApiPolicyListDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPolicyListDTOWithDefaults() *ApiPolicyListDTO {
	this := ApiPolicyListDTO{}
	return &this
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *ApiPolicyListDTO) GetPolicies() []ApiPolicyDTO {
	if o == nil || IsNil(o.Policies) {
		var ret []ApiPolicyDTO
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyListDTO) GetPoliciesOk() ([]ApiPolicyDTO, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *ApiPolicyListDTO) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []ApiPolicyDTO and assigns it to the Policies field.
func (o *ApiPolicyListDTO) SetPolicies(v []ApiPolicyDTO) {
	o.Policies = v
}

func (o ApiPolicyListDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPolicyListDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	return toSerialize, nil
}

type NullableApiPolicyListDTO struct {
	value *ApiPolicyListDTO
	isSet bool
}

func (v NullableApiPolicyListDTO) Get() *ApiPolicyListDTO {
	return v.value
}

func (v *NullableApiPolicyListDTO) Set(val *ApiPolicyListDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPolicyListDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPolicyListDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPolicyListDTO(val *ApiPolicyListDTO) *NullableApiPolicyListDTO {
	return &NullableApiPolicyListDTO{value: val, isSet: true}
}

func (v NullableApiPolicyListDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPolicyListDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


