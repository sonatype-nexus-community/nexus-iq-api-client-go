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

// checks if the ApiStagePolicyViolationComponentDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiStagePolicyViolationComponentDTO{}

// ApiStagePolicyViolationComponentDTO struct for ApiStagePolicyViolationComponentDTO
type ApiStagePolicyViolationComponentDTO struct {
	Action *string `json:"action,omitempty"`
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Hash *string `json:"hash,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
	PolicyId *string `json:"policyId,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
	PolicyViolationId *string `json:"policyViolationId,omitempty"`
	ThreatCategory *string `json:"threatCategory,omitempty"`
	ThreatLevel *int32 `json:"threatLevel,omitempty"`
}

// NewApiStagePolicyViolationComponentDTO instantiates a new ApiStagePolicyViolationComponentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStagePolicyViolationComponentDTO() *ApiStagePolicyViolationComponentDTO {
	this := ApiStagePolicyViolationComponentDTO{}
	return &this
}

// NewApiStagePolicyViolationComponentDTOWithDefaults instantiates a new ApiStagePolicyViolationComponentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStagePolicyViolationComponentDTOWithDefaults() *ApiStagePolicyViolationComponentDTO {
	this := ApiStagePolicyViolationComponentDTO{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ApiStagePolicyViolationComponentDTO) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStagePolicyViolationComponentDTO) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ApiStagePolicyViolationComponentDTO) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *ApiStagePolicyViolationComponentDTO) SetAction(v string) {
	o.Action = &v
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiStagePolicyViolationComponentDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStagePolicyViolationComponentDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiStagePolicyViolationComponentDTO) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiStagePolicyViolationComponentDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ApiStagePolicyViolationComponentDTO) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStagePolicyViolationComponentDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ApiStagePolicyViolationComponentDTO) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ApiStagePolicyViolationComponentDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ApiStagePolicyViolationComponentDTO) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStagePolicyViolationComponentDTO) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ApiStagePolicyViolationComponentDTO) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ApiStagePolicyViolationComponentDTO) SetHash(v string) {
	o.Hash = &v
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ApiStagePolicyViolationComponentDTO) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStagePolicyViolationComponentDTO) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ApiStagePolicyViolationComponentDTO) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ApiStagePolicyViolationComponentDTO) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *ApiStagePolicyViolationComponentDTO) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStagePolicyViolationComponentDTO) GetPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *ApiStagePolicyViolationComponentDTO) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *ApiStagePolicyViolationComponentDTO) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *ApiStagePolicyViolationComponentDTO) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName) {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStagePolicyViolationComponentDTO) GetPolicyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyName) {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *ApiStagePolicyViolationComponentDTO) HasPolicyName() bool {
	if o != nil && !IsNil(o.PolicyName) {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *ApiStagePolicyViolationComponentDTO) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetPolicyViolationId returns the PolicyViolationId field value if set, zero value otherwise.
func (o *ApiStagePolicyViolationComponentDTO) GetPolicyViolationId() string {
	if o == nil || IsNil(o.PolicyViolationId) {
		var ret string
		return ret
	}
	return *o.PolicyViolationId
}

// GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStagePolicyViolationComponentDTO) GetPolicyViolationIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyViolationId) {
		return nil, false
	}
	return o.PolicyViolationId, true
}

// HasPolicyViolationId returns a boolean if a field has been set.
func (o *ApiStagePolicyViolationComponentDTO) HasPolicyViolationId() bool {
	if o != nil && !IsNil(o.PolicyViolationId) {
		return true
	}

	return false
}

// SetPolicyViolationId gets a reference to the given string and assigns it to the PolicyViolationId field.
func (o *ApiStagePolicyViolationComponentDTO) SetPolicyViolationId(v string) {
	o.PolicyViolationId = &v
}

// GetThreatCategory returns the ThreatCategory field value if set, zero value otherwise.
func (o *ApiStagePolicyViolationComponentDTO) GetThreatCategory() string {
	if o == nil || IsNil(o.ThreatCategory) {
		var ret string
		return ret
	}
	return *o.ThreatCategory
}

// GetThreatCategoryOk returns a tuple with the ThreatCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStagePolicyViolationComponentDTO) GetThreatCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ThreatCategory) {
		return nil, false
	}
	return o.ThreatCategory, true
}

// HasThreatCategory returns a boolean if a field has been set.
func (o *ApiStagePolicyViolationComponentDTO) HasThreatCategory() bool {
	if o != nil && !IsNil(o.ThreatCategory) {
		return true
	}

	return false
}

// SetThreatCategory gets a reference to the given string and assigns it to the ThreatCategory field.
func (o *ApiStagePolicyViolationComponentDTO) SetThreatCategory(v string) {
	o.ThreatCategory = &v
}

// GetThreatLevel returns the ThreatLevel field value if set, zero value otherwise.
func (o *ApiStagePolicyViolationComponentDTO) GetThreatLevel() int32 {
	if o == nil || IsNil(o.ThreatLevel) {
		var ret int32
		return ret
	}
	return *o.ThreatLevel
}

// GetThreatLevelOk returns a tuple with the ThreatLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStagePolicyViolationComponentDTO) GetThreatLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.ThreatLevel) {
		return nil, false
	}
	return o.ThreatLevel, true
}

// HasThreatLevel returns a boolean if a field has been set.
func (o *ApiStagePolicyViolationComponentDTO) HasThreatLevel() bool {
	if o != nil && !IsNil(o.ThreatLevel) {
		return true
	}

	return false
}

// SetThreatLevel gets a reference to the given int32 and assigns it to the ThreatLevel field.
func (o *ApiStagePolicyViolationComponentDTO) SetThreatLevel(v int32) {
	o.ThreatLevel = &v
}

func (o ApiStagePolicyViolationComponentDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiStagePolicyViolationComponentDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
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
	if !IsNil(o.ThreatCategory) {
		toSerialize["threatCategory"] = o.ThreatCategory
	}
	if !IsNil(o.ThreatLevel) {
		toSerialize["threatLevel"] = o.ThreatLevel
	}
	return toSerialize, nil
}

type NullableApiStagePolicyViolationComponentDTO struct {
	value *ApiStagePolicyViolationComponentDTO
	isSet bool
}

func (v NullableApiStagePolicyViolationComponentDTO) Get() *ApiStagePolicyViolationComponentDTO {
	return v.value
}

func (v *NullableApiStagePolicyViolationComponentDTO) Set(val *ApiStagePolicyViolationComponentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStagePolicyViolationComponentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStagePolicyViolationComponentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStagePolicyViolationComponentDTO(val *ApiStagePolicyViolationComponentDTO) *NullableApiStagePolicyViolationComponentDTO {
	return &NullableApiStagePolicyViolationComponentDTO{value: val, isSet: true}
}

func (v NullableApiStagePolicyViolationComponentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStagePolicyViolationComponentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

