/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.184.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiAutoPolicyWaiverStatusDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAutoPolicyWaiverStatusDTO{}

// ApiAutoPolicyWaiverStatusDTO struct for ApiAutoPolicyWaiverStatusDTO
type ApiAutoPolicyWaiverStatusDTO struct {
	AutoPolicyWaiverId *string `json:"autoPolicyWaiverId,omitempty"`
	AutoPolicyWaiverOwnerId *string `json:"autoPolicyWaiverOwnerId,omitempty"`
	AutoPolicyWaiverOwnerName *string `json:"autoPolicyWaiverOwnerName,omitempty"`
	IsAutoWaiverEnabled *bool `json:"isAutoWaiverEnabled,omitempty"`
	IsInherited *bool `json:"isInherited,omitempty"`
}

// NewApiAutoPolicyWaiverStatusDTO instantiates a new ApiAutoPolicyWaiverStatusDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAutoPolicyWaiverStatusDTO() *ApiAutoPolicyWaiverStatusDTO {
	this := ApiAutoPolicyWaiverStatusDTO{}
	return &this
}

// NewApiAutoPolicyWaiverStatusDTOWithDefaults instantiates a new ApiAutoPolicyWaiverStatusDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAutoPolicyWaiverStatusDTOWithDefaults() *ApiAutoPolicyWaiverStatusDTO {
	this := ApiAutoPolicyWaiverStatusDTO{}
	return &this
}

// GetAutoPolicyWaiverId returns the AutoPolicyWaiverId field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverId() string {
	if o == nil || IsNil(o.AutoPolicyWaiverId) {
		var ret string
		return ret
	}
	return *o.AutoPolicyWaiverId
}

// GetAutoPolicyWaiverIdOk returns a tuple with the AutoPolicyWaiverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverIdOk() (*string, bool) {
	if o == nil || IsNil(o.AutoPolicyWaiverId) {
		return nil, false
	}
	return o.AutoPolicyWaiverId, true
}

// HasAutoPolicyWaiverId returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverStatusDTO) HasAutoPolicyWaiverId() bool {
	if o != nil && !IsNil(o.AutoPolicyWaiverId) {
		return true
	}

	return false
}

// SetAutoPolicyWaiverId gets a reference to the given string and assigns it to the AutoPolicyWaiverId field.
func (o *ApiAutoPolicyWaiverStatusDTO) SetAutoPolicyWaiverId(v string) {
	o.AutoPolicyWaiverId = &v
}

// GetAutoPolicyWaiverOwnerId returns the AutoPolicyWaiverOwnerId field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverOwnerId() string {
	if o == nil || IsNil(o.AutoPolicyWaiverOwnerId) {
		var ret string
		return ret
	}
	return *o.AutoPolicyWaiverOwnerId
}

// GetAutoPolicyWaiverOwnerIdOk returns a tuple with the AutoPolicyWaiverOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.AutoPolicyWaiverOwnerId) {
		return nil, false
	}
	return o.AutoPolicyWaiverOwnerId, true
}

// HasAutoPolicyWaiverOwnerId returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverStatusDTO) HasAutoPolicyWaiverOwnerId() bool {
	if o != nil && !IsNil(o.AutoPolicyWaiverOwnerId) {
		return true
	}

	return false
}

// SetAutoPolicyWaiverOwnerId gets a reference to the given string and assigns it to the AutoPolicyWaiverOwnerId field.
func (o *ApiAutoPolicyWaiverStatusDTO) SetAutoPolicyWaiverOwnerId(v string) {
	o.AutoPolicyWaiverOwnerId = &v
}

// GetAutoPolicyWaiverOwnerName returns the AutoPolicyWaiverOwnerName field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverOwnerName() string {
	if o == nil || IsNil(o.AutoPolicyWaiverOwnerName) {
		var ret string
		return ret
	}
	return *o.AutoPolicyWaiverOwnerName
}

// GetAutoPolicyWaiverOwnerNameOk returns a tuple with the AutoPolicyWaiverOwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.AutoPolicyWaiverOwnerName) {
		return nil, false
	}
	return o.AutoPolicyWaiverOwnerName, true
}

// HasAutoPolicyWaiverOwnerName returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverStatusDTO) HasAutoPolicyWaiverOwnerName() bool {
	if o != nil && !IsNil(o.AutoPolicyWaiverOwnerName) {
		return true
	}

	return false
}

// SetAutoPolicyWaiverOwnerName gets a reference to the given string and assigns it to the AutoPolicyWaiverOwnerName field.
func (o *ApiAutoPolicyWaiverStatusDTO) SetAutoPolicyWaiverOwnerName(v string) {
	o.AutoPolicyWaiverOwnerName = &v
}

// GetIsAutoWaiverEnabled returns the IsAutoWaiverEnabled field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverStatusDTO) GetIsAutoWaiverEnabled() bool {
	if o == nil || IsNil(o.IsAutoWaiverEnabled) {
		var ret bool
		return ret
	}
	return *o.IsAutoWaiverEnabled
}

// GetIsAutoWaiverEnabledOk returns a tuple with the IsAutoWaiverEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverStatusDTO) GetIsAutoWaiverEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutoWaiverEnabled) {
		return nil, false
	}
	return o.IsAutoWaiverEnabled, true
}

// HasIsAutoWaiverEnabled returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverStatusDTO) HasIsAutoWaiverEnabled() bool {
	if o != nil && !IsNil(o.IsAutoWaiverEnabled) {
		return true
	}

	return false
}

// SetIsAutoWaiverEnabled gets a reference to the given bool and assigns it to the IsAutoWaiverEnabled field.
func (o *ApiAutoPolicyWaiverStatusDTO) SetIsAutoWaiverEnabled(v bool) {
	o.IsAutoWaiverEnabled = &v
}

// GetIsInherited returns the IsInherited field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverStatusDTO) GetIsInherited() bool {
	if o == nil || IsNil(o.IsInherited) {
		var ret bool
		return ret
	}
	return *o.IsInherited
}

// GetIsInheritedOk returns a tuple with the IsInherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverStatusDTO) GetIsInheritedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInherited) {
		return nil, false
	}
	return o.IsInherited, true
}

// HasIsInherited returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverStatusDTO) HasIsInherited() bool {
	if o != nil && !IsNil(o.IsInherited) {
		return true
	}

	return false
}

// SetIsInherited gets a reference to the given bool and assigns it to the IsInherited field.
func (o *ApiAutoPolicyWaiverStatusDTO) SetIsInherited(v bool) {
	o.IsInherited = &v
}

func (o ApiAutoPolicyWaiverStatusDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAutoPolicyWaiverStatusDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoPolicyWaiverId) {
		toSerialize["autoPolicyWaiverId"] = o.AutoPolicyWaiverId
	}
	if !IsNil(o.AutoPolicyWaiverOwnerId) {
		toSerialize["autoPolicyWaiverOwnerId"] = o.AutoPolicyWaiverOwnerId
	}
	if !IsNil(o.AutoPolicyWaiverOwnerName) {
		toSerialize["autoPolicyWaiverOwnerName"] = o.AutoPolicyWaiverOwnerName
	}
	if !IsNil(o.IsAutoWaiverEnabled) {
		toSerialize["isAutoWaiverEnabled"] = o.IsAutoWaiverEnabled
	}
	if !IsNil(o.IsInherited) {
		toSerialize["isInherited"] = o.IsInherited
	}
	return toSerialize, nil
}

type NullableApiAutoPolicyWaiverStatusDTO struct {
	value *ApiAutoPolicyWaiverStatusDTO
	isSet bool
}

func (v NullableApiAutoPolicyWaiverStatusDTO) Get() *ApiAutoPolicyWaiverStatusDTO {
	return v.value
}

func (v *NullableApiAutoPolicyWaiverStatusDTO) Set(val *ApiAutoPolicyWaiverStatusDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAutoPolicyWaiverStatusDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAutoPolicyWaiverStatusDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAutoPolicyWaiverStatusDTO(val *ApiAutoPolicyWaiverStatusDTO) *NullableApiAutoPolicyWaiverStatusDTO {
	return &NullableApiAutoPolicyWaiverStatusDTO{value: val, isSet: true}
}

func (v NullableApiAutoPolicyWaiverStatusDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAutoPolicyWaiverStatusDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


