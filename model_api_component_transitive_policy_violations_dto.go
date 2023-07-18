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

// checks if the ApiComponentTransitivePolicyViolationsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentTransitivePolicyViolationsDTO{}

// ApiComponentTransitivePolicyViolationsDTO struct for ApiComponentTransitivePolicyViolationsDTO
type ApiComponentTransitivePolicyViolationsDTO struct {
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Hash *string `json:"hash,omitempty"`
	IsInnerSource *bool `json:"isInnerSource,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
	TransitivePolicyViolations []ApiStagePolicyViolationComponentDTO `json:"transitivePolicyViolations,omitempty"`
}

// NewApiComponentTransitivePolicyViolationsDTO instantiates a new ApiComponentTransitivePolicyViolationsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentTransitivePolicyViolationsDTO() *ApiComponentTransitivePolicyViolationsDTO {
	this := ApiComponentTransitivePolicyViolationsDTO{}
	return &this
}

// NewApiComponentTransitivePolicyViolationsDTOWithDefaults instantiates a new ApiComponentTransitivePolicyViolationsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentTransitivePolicyViolationsDTOWithDefaults() *ApiComponentTransitivePolicyViolationsDTO {
	this := ApiComponentTransitivePolicyViolationsDTO{}
	return &this
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiComponentTransitivePolicyViolationsDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentTransitivePolicyViolationsDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiComponentTransitivePolicyViolationsDTO) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiComponentTransitivePolicyViolationsDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ApiComponentTransitivePolicyViolationsDTO) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentTransitivePolicyViolationsDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ApiComponentTransitivePolicyViolationsDTO) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ApiComponentTransitivePolicyViolationsDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ApiComponentTransitivePolicyViolationsDTO) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentTransitivePolicyViolationsDTO) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ApiComponentTransitivePolicyViolationsDTO) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ApiComponentTransitivePolicyViolationsDTO) SetHash(v string) {
	o.Hash = &v
}

// GetIsInnerSource returns the IsInnerSource field value if set, zero value otherwise.
func (o *ApiComponentTransitivePolicyViolationsDTO) GetIsInnerSource() bool {
	if o == nil || IsNil(o.IsInnerSource) {
		var ret bool
		return ret
	}
	return *o.IsInnerSource
}

// GetIsInnerSourceOk returns a tuple with the IsInnerSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentTransitivePolicyViolationsDTO) GetIsInnerSourceOk() (*bool, bool) {
	if o == nil || IsNil(o.IsInnerSource) {
		return nil, false
	}
	return o.IsInnerSource, true
}

// HasIsInnerSource returns a boolean if a field has been set.
func (o *ApiComponentTransitivePolicyViolationsDTO) HasIsInnerSource() bool {
	if o != nil && !IsNil(o.IsInnerSource) {
		return true
	}

	return false
}

// SetIsInnerSource gets a reference to the given bool and assigns it to the IsInnerSource field.
func (o *ApiComponentTransitivePolicyViolationsDTO) SetIsInnerSource(v bool) {
	o.IsInnerSource = &v
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ApiComponentTransitivePolicyViolationsDTO) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentTransitivePolicyViolationsDTO) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ApiComponentTransitivePolicyViolationsDTO) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ApiComponentTransitivePolicyViolationsDTO) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

// GetTransitivePolicyViolations returns the TransitivePolicyViolations field value if set, zero value otherwise.
func (o *ApiComponentTransitivePolicyViolationsDTO) GetTransitivePolicyViolations() []ApiStagePolicyViolationComponentDTO {
	if o == nil || IsNil(o.TransitivePolicyViolations) {
		var ret []ApiStagePolicyViolationComponentDTO
		return ret
	}
	return o.TransitivePolicyViolations
}

// GetTransitivePolicyViolationsOk returns a tuple with the TransitivePolicyViolations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentTransitivePolicyViolationsDTO) GetTransitivePolicyViolationsOk() ([]ApiStagePolicyViolationComponentDTO, bool) {
	if o == nil || IsNil(o.TransitivePolicyViolations) {
		return nil, false
	}
	return o.TransitivePolicyViolations, true
}

// HasTransitivePolicyViolations returns a boolean if a field has been set.
func (o *ApiComponentTransitivePolicyViolationsDTO) HasTransitivePolicyViolations() bool {
	if o != nil && !IsNil(o.TransitivePolicyViolations) {
		return true
	}

	return false
}

// SetTransitivePolicyViolations gets a reference to the given []ApiStagePolicyViolationComponentDTO and assigns it to the TransitivePolicyViolations field.
func (o *ApiComponentTransitivePolicyViolationsDTO) SetTransitivePolicyViolations(v []ApiStagePolicyViolationComponentDTO) {
	o.TransitivePolicyViolations = v
}

func (o ApiComponentTransitivePolicyViolationsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentTransitivePolicyViolationsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.IsInnerSource) {
		toSerialize["isInnerSource"] = o.IsInnerSource
	}
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
	}
	if !IsNil(o.TransitivePolicyViolations) {
		toSerialize["transitivePolicyViolations"] = o.TransitivePolicyViolations
	}
	return toSerialize, nil
}

type NullableApiComponentTransitivePolicyViolationsDTO struct {
	value *ApiComponentTransitivePolicyViolationsDTO
	isSet bool
}

func (v NullableApiComponentTransitivePolicyViolationsDTO) Get() *ApiComponentTransitivePolicyViolationsDTO {
	return v.value
}

func (v *NullableApiComponentTransitivePolicyViolationsDTO) Set(val *ApiComponentTransitivePolicyViolationsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentTransitivePolicyViolationsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentTransitivePolicyViolationsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentTransitivePolicyViolationsDTO(val *ApiComponentTransitivePolicyViolationsDTO) *NullableApiComponentTransitivePolicyViolationsDTO {
	return &NullableApiComponentTransitivePolicyViolationsDTO{value: val, isSet: true}
}

func (v NullableApiComponentTransitivePolicyViolationsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentTransitivePolicyViolationsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


