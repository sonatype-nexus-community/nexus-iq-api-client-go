/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.182.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiRequestPolicyWaiverDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRequestPolicyWaiverDTO{}

// ApiRequestPolicyWaiverDTO struct for ApiRequestPolicyWaiverDTO
type ApiRequestPolicyWaiverDTO struct {
	AddWaiverLink *string `json:"addWaiverLink,omitempty"`
	Comment *string `json:"comment,omitempty"`
	PolicyViolationLink *string `json:"policyViolationLink,omitempty"`
}

// NewApiRequestPolicyWaiverDTO instantiates a new ApiRequestPolicyWaiverDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRequestPolicyWaiverDTO() *ApiRequestPolicyWaiverDTO {
	this := ApiRequestPolicyWaiverDTO{}
	return &this
}

// NewApiRequestPolicyWaiverDTOWithDefaults instantiates a new ApiRequestPolicyWaiverDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRequestPolicyWaiverDTOWithDefaults() *ApiRequestPolicyWaiverDTO {
	this := ApiRequestPolicyWaiverDTO{}
	return &this
}

// GetAddWaiverLink returns the AddWaiverLink field value if set, zero value otherwise.
func (o *ApiRequestPolicyWaiverDTO) GetAddWaiverLink() string {
	if o == nil || IsNil(o.AddWaiverLink) {
		var ret string
		return ret
	}
	return *o.AddWaiverLink
}

// GetAddWaiverLinkOk returns a tuple with the AddWaiverLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequestPolicyWaiverDTO) GetAddWaiverLinkOk() (*string, bool) {
	if o == nil || IsNil(o.AddWaiverLink) {
		return nil, false
	}
	return o.AddWaiverLink, true
}

// HasAddWaiverLink returns a boolean if a field has been set.
func (o *ApiRequestPolicyWaiverDTO) HasAddWaiverLink() bool {
	if o != nil && !IsNil(o.AddWaiverLink) {
		return true
	}

	return false
}

// SetAddWaiverLink gets a reference to the given string and assigns it to the AddWaiverLink field.
func (o *ApiRequestPolicyWaiverDTO) SetAddWaiverLink(v string) {
	o.AddWaiverLink = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApiRequestPolicyWaiverDTO) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequestPolicyWaiverDTO) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApiRequestPolicyWaiverDTO) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApiRequestPolicyWaiverDTO) SetComment(v string) {
	o.Comment = &v
}

// GetPolicyViolationLink returns the PolicyViolationLink field value if set, zero value otherwise.
func (o *ApiRequestPolicyWaiverDTO) GetPolicyViolationLink() string {
	if o == nil || IsNil(o.PolicyViolationLink) {
		var ret string
		return ret
	}
	return *o.PolicyViolationLink
}

// GetPolicyViolationLinkOk returns a tuple with the PolicyViolationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequestPolicyWaiverDTO) GetPolicyViolationLinkOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyViolationLink) {
		return nil, false
	}
	return o.PolicyViolationLink, true
}

// HasPolicyViolationLink returns a boolean if a field has been set.
func (o *ApiRequestPolicyWaiverDTO) HasPolicyViolationLink() bool {
	if o != nil && !IsNil(o.PolicyViolationLink) {
		return true
	}

	return false
}

// SetPolicyViolationLink gets a reference to the given string and assigns it to the PolicyViolationLink field.
func (o *ApiRequestPolicyWaiverDTO) SetPolicyViolationLink(v string) {
	o.PolicyViolationLink = &v
}

func (o ApiRequestPolicyWaiverDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRequestPolicyWaiverDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddWaiverLink) {
		toSerialize["addWaiverLink"] = o.AddWaiverLink
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.PolicyViolationLink) {
		toSerialize["policyViolationLink"] = o.PolicyViolationLink
	}
	return toSerialize, nil
}

type NullableApiRequestPolicyWaiverDTO struct {
	value *ApiRequestPolicyWaiverDTO
	isSet bool
}

func (v NullableApiRequestPolicyWaiverDTO) Get() *ApiRequestPolicyWaiverDTO {
	return v.value
}

func (v *NullableApiRequestPolicyWaiverDTO) Set(val *ApiRequestPolicyWaiverDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRequestPolicyWaiverDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRequestPolicyWaiverDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRequestPolicyWaiverDTO(val *ApiRequestPolicyWaiverDTO) *NullableApiRequestPolicyWaiverDTO {
	return &NullableApiRequestPolicyWaiverDTO{value: val, isSet: true}
}

func (v NullableApiRequestPolicyWaiverDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRequestPolicyWaiverDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


