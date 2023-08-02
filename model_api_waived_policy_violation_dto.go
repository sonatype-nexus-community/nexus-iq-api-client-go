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

// checks if the ApiWaivedPolicyViolationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiWaivedPolicyViolationDTO{}

// ApiWaivedPolicyViolationDTO struct for ApiWaivedPolicyViolationDTO
type ApiWaivedPolicyViolationDTO struct {
	ConstraintViolations []ApiConstraintViolationDTO `json:"constraintViolations,omitempty"`
	PolicyId *string `json:"policyId,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
	PolicyViolationId *string `json:"policyViolationId,omitempty"`
	PolicyWaiver *ApiPolicyWaiverDTO `json:"policyWaiver,omitempty"`
	ThreatLevel *int32 `json:"threatLevel,omitempty"`
}

// NewApiWaivedPolicyViolationDTO instantiates a new ApiWaivedPolicyViolationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiWaivedPolicyViolationDTO() *ApiWaivedPolicyViolationDTO {
	this := ApiWaivedPolicyViolationDTO{}
	return &this
}

// NewApiWaivedPolicyViolationDTOWithDefaults instantiates a new ApiWaivedPolicyViolationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiWaivedPolicyViolationDTOWithDefaults() *ApiWaivedPolicyViolationDTO {
	this := ApiWaivedPolicyViolationDTO{}
	return &this
}

// GetConstraintViolations returns the ConstraintViolations field value if set, zero value otherwise.
func (o *ApiWaivedPolicyViolationDTO) GetConstraintViolations() []ApiConstraintViolationDTO {
	if o == nil || IsNil(o.ConstraintViolations) {
		var ret []ApiConstraintViolationDTO
		return ret
	}
	return o.ConstraintViolations
}

// GetConstraintViolationsOk returns a tuple with the ConstraintViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWaivedPolicyViolationDTO) GetConstraintViolationsOk() ([]ApiConstraintViolationDTO, bool) {
	if o == nil || IsNil(o.ConstraintViolations) {
		return nil, false
	}
	return o.ConstraintViolations, true
}

// HasConstraintViolations returns a boolean if a field has been set.
func (o *ApiWaivedPolicyViolationDTO) HasConstraintViolations() bool {
	if o != nil && !IsNil(o.ConstraintViolations) {
		return true
	}

	return false
}

// SetConstraintViolations gets a reference to the given []ApiConstraintViolationDTO and assigns it to the ConstraintViolations field.
func (o *ApiWaivedPolicyViolationDTO) SetConstraintViolations(v []ApiConstraintViolationDTO) {
	o.ConstraintViolations = v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *ApiWaivedPolicyViolationDTO) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWaivedPolicyViolationDTO) GetPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *ApiWaivedPolicyViolationDTO) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *ApiWaivedPolicyViolationDTO) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *ApiWaivedPolicyViolationDTO) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName) {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWaivedPolicyViolationDTO) GetPolicyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyName) {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *ApiWaivedPolicyViolationDTO) HasPolicyName() bool {
	if o != nil && !IsNil(o.PolicyName) {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *ApiWaivedPolicyViolationDTO) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetPolicyViolationId returns the PolicyViolationId field value if set, zero value otherwise.
func (o *ApiWaivedPolicyViolationDTO) GetPolicyViolationId() string {
	if o == nil || IsNil(o.PolicyViolationId) {
		var ret string
		return ret
	}
	return *o.PolicyViolationId
}

// GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWaivedPolicyViolationDTO) GetPolicyViolationIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyViolationId) {
		return nil, false
	}
	return o.PolicyViolationId, true
}

// HasPolicyViolationId returns a boolean if a field has been set.
func (o *ApiWaivedPolicyViolationDTO) HasPolicyViolationId() bool {
	if o != nil && !IsNil(o.PolicyViolationId) {
		return true
	}

	return false
}

// SetPolicyViolationId gets a reference to the given string and assigns it to the PolicyViolationId field.
func (o *ApiWaivedPolicyViolationDTO) SetPolicyViolationId(v string) {
	o.PolicyViolationId = &v
}

// GetPolicyWaiver returns the PolicyWaiver field value if set, zero value otherwise.
func (o *ApiWaivedPolicyViolationDTO) GetPolicyWaiver() ApiPolicyWaiverDTO {
	if o == nil || IsNil(o.PolicyWaiver) {
		var ret ApiPolicyWaiverDTO
		return ret
	}
	return *o.PolicyWaiver
}

// GetPolicyWaiverOk returns a tuple with the PolicyWaiver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWaivedPolicyViolationDTO) GetPolicyWaiverOk() (*ApiPolicyWaiverDTO, bool) {
	if o == nil || IsNil(o.PolicyWaiver) {
		return nil, false
	}
	return o.PolicyWaiver, true
}

// HasPolicyWaiver returns a boolean if a field has been set.
func (o *ApiWaivedPolicyViolationDTO) HasPolicyWaiver() bool {
	if o != nil && !IsNil(o.PolicyWaiver) {
		return true
	}

	return false
}

// SetPolicyWaiver gets a reference to the given ApiPolicyWaiverDTO and assigns it to the PolicyWaiver field.
func (o *ApiWaivedPolicyViolationDTO) SetPolicyWaiver(v ApiPolicyWaiverDTO) {
	o.PolicyWaiver = &v
}

// GetThreatLevel returns the ThreatLevel field value if set, zero value otherwise.
func (o *ApiWaivedPolicyViolationDTO) GetThreatLevel() int32 {
	if o == nil || IsNil(o.ThreatLevel) {
		var ret int32
		return ret
	}
	return *o.ThreatLevel
}

// GetThreatLevelOk returns a tuple with the ThreatLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWaivedPolicyViolationDTO) GetThreatLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.ThreatLevel) {
		return nil, false
	}
	return o.ThreatLevel, true
}

// HasThreatLevel returns a boolean if a field has been set.
func (o *ApiWaivedPolicyViolationDTO) HasThreatLevel() bool {
	if o != nil && !IsNil(o.ThreatLevel) {
		return true
	}

	return false
}

// SetThreatLevel gets a reference to the given int32 and assigns it to the ThreatLevel field.
func (o *ApiWaivedPolicyViolationDTO) SetThreatLevel(v int32) {
	o.ThreatLevel = &v
}

func (o ApiWaivedPolicyViolationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiWaivedPolicyViolationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConstraintViolations) {
		toSerialize["constraintViolations"] = o.ConstraintViolations
	}
	if !IsNil(o.PolicyId) {
		toSerialize["policyId"] = o.PolicyId
	}
	if !IsNil(o.PolicyName) {
		toSerialize["policyName"] = o.PolicyName
	}
	if !IsNil(o.PolicyViolationId) {
		toSerialize["policyViolationId"] = o.PolicyViolationId
	}
	if !IsNil(o.PolicyWaiver) {
		toSerialize["policyWaiver"] = o.PolicyWaiver
	}
	if !IsNil(o.ThreatLevel) {
		toSerialize["threatLevel"] = o.ThreatLevel
	}
	return toSerialize, nil
}

type NullableApiWaivedPolicyViolationDTO struct {
	value *ApiWaivedPolicyViolationDTO
	isSet bool
}

func (v NullableApiWaivedPolicyViolationDTO) Get() *ApiWaivedPolicyViolationDTO {
	return v.value
}

func (v *NullableApiWaivedPolicyViolationDTO) Set(val *ApiWaivedPolicyViolationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiWaivedPolicyViolationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiWaivedPolicyViolationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiWaivedPolicyViolationDTO(val *ApiWaivedPolicyViolationDTO) *NullableApiWaivedPolicyViolationDTO {
	return &NullableApiWaivedPolicyViolationDTO{value: val, isSet: true}
}

func (v NullableApiWaivedPolicyViolationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiWaivedPolicyViolationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

