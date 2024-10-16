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

// checks if the PolicyOwner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyOwner{}

// PolicyOwner struct for PolicyOwner
type PolicyOwner struct {
	OwnerId *string `json:"ownerId,omitempty"`
	OwnerName *string `json:"ownerName,omitempty"`
	OwnerPublicId *string `json:"ownerPublicId,omitempty"`
	OwnerType *string `json:"ownerType,omitempty"`
}

// NewPolicyOwner instantiates a new PolicyOwner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyOwner() *PolicyOwner {
	this := PolicyOwner{}
	return &this
}

// NewPolicyOwnerWithDefaults instantiates a new PolicyOwner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyOwnerWithDefaults() *PolicyOwner {
	this := PolicyOwner{}
	return &this
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *PolicyOwner) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyOwner) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *PolicyOwner) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *PolicyOwner) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *PolicyOwner) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyOwner) GetOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *PolicyOwner) HasOwnerName() bool {
	if o != nil && !IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *PolicyOwner) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetOwnerPublicId returns the OwnerPublicId field value if set, zero value otherwise.
func (o *PolicyOwner) GetOwnerPublicId() string {
	if o == nil || IsNil(o.OwnerPublicId) {
		var ret string
		return ret
	}
	return *o.OwnerPublicId
}

// GetOwnerPublicIdOk returns a tuple with the OwnerPublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyOwner) GetOwnerPublicIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerPublicId) {
		return nil, false
	}
	return o.OwnerPublicId, true
}

// HasOwnerPublicId returns a boolean if a field has been set.
func (o *PolicyOwner) HasOwnerPublicId() bool {
	if o != nil && !IsNil(o.OwnerPublicId) {
		return true
	}

	return false
}

// SetOwnerPublicId gets a reference to the given string and assigns it to the OwnerPublicId field.
func (o *PolicyOwner) SetOwnerPublicId(v string) {
	o.OwnerPublicId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *PolicyOwner) GetOwnerType() string {
	if o == nil || IsNil(o.OwnerType) {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyOwner) GetOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerType) {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *PolicyOwner) HasOwnerType() bool {
	if o != nil && !IsNil(o.OwnerType) {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *PolicyOwner) SetOwnerType(v string) {
	o.OwnerType = &v
}

func (o PolicyOwner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyOwner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !IsNil(o.OwnerPublicId) {
		toSerialize["ownerPublicId"] = o.OwnerPublicId
	}
	if !IsNil(o.OwnerType) {
		toSerialize["ownerType"] = o.OwnerType
	}
	return toSerialize, nil
}

type NullablePolicyOwner struct {
	value *PolicyOwner
	isSet bool
}

func (v NullablePolicyOwner) Get() *PolicyOwner {
	return v.value
}

func (v *NullablePolicyOwner) Set(val *PolicyOwner) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyOwner) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyOwner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyOwner(val *PolicyOwner) *NullablePolicyOwner {
	return &NullablePolicyOwner{value: val, isSet: true}
}

func (v NullablePolicyOwner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyOwner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


