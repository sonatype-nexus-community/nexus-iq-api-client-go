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

// checks if the ApiDependencyTreeNodeDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDependencyTreeNodeDTO{}

// ApiDependencyTreeNodeDTO struct for ApiDependencyTreeNodeDTO
type ApiDependencyTreeNodeDTO struct {
	Children []ApiDependencyTreeNodeDTO `json:"children,omitempty"`
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	Direct *bool `json:"direct,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
}

// NewApiDependencyTreeNodeDTO instantiates a new ApiDependencyTreeNodeDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDependencyTreeNodeDTO() *ApiDependencyTreeNodeDTO {
	this := ApiDependencyTreeNodeDTO{}
	return &this
}

// NewApiDependencyTreeNodeDTOWithDefaults instantiates a new ApiDependencyTreeNodeDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDependencyTreeNodeDTOWithDefaults() *ApiDependencyTreeNodeDTO {
	this := ApiDependencyTreeNodeDTO{}
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *ApiDependencyTreeNodeDTO) GetChildren() []ApiDependencyTreeNodeDTO {
	if o == nil || IsNil(o.Children) {
		var ret []ApiDependencyTreeNodeDTO
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDependencyTreeNodeDTO) GetChildrenOk() ([]ApiDependencyTreeNodeDTO, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ApiDependencyTreeNodeDTO) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []ApiDependencyTreeNodeDTO and assigns it to the Children field.
func (o *ApiDependencyTreeNodeDTO) SetChildren(v []ApiDependencyTreeNodeDTO) {
	o.Children = v
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiDependencyTreeNodeDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDependencyTreeNodeDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiDependencyTreeNodeDTO) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiDependencyTreeNodeDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetDirect returns the Direct field value if set, zero value otherwise.
func (o *ApiDependencyTreeNodeDTO) GetDirect() bool {
	if o == nil || IsNil(o.Direct) {
		var ret bool
		return ret
	}
	return *o.Direct
}

// GetDirectOk returns a tuple with the Direct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDependencyTreeNodeDTO) GetDirectOk() (*bool, bool) {
	if o == nil || IsNil(o.Direct) {
		return nil, false
	}
	return o.Direct, true
}

// HasDirect returns a boolean if a field has been set.
func (o *ApiDependencyTreeNodeDTO) HasDirect() bool {
	if o != nil && !IsNil(o.Direct) {
		return true
	}

	return false
}

// SetDirect gets a reference to the given bool and assigns it to the Direct field.
func (o *ApiDependencyTreeNodeDTO) SetDirect(v bool) {
	o.Direct = &v
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ApiDependencyTreeNodeDTO) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDependencyTreeNodeDTO) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ApiDependencyTreeNodeDTO) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ApiDependencyTreeNodeDTO) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

func (o ApiDependencyTreeNodeDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDependencyTreeNodeDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.Direct) {
		toSerialize["direct"] = o.Direct
	}
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
	}
	return toSerialize, nil
}

type NullableApiDependencyTreeNodeDTO struct {
	value *ApiDependencyTreeNodeDTO
	isSet bool
}

func (v NullableApiDependencyTreeNodeDTO) Get() *ApiDependencyTreeNodeDTO {
	return v.value
}

func (v *NullableApiDependencyTreeNodeDTO) Set(val *ApiDependencyTreeNodeDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDependencyTreeNodeDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDependencyTreeNodeDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDependencyTreeNodeDTO(val *ApiDependencyTreeNodeDTO) *NullableApiDependencyTreeNodeDTO {
	return &NullableApiDependencyTreeNodeDTO{value: val, isSet: true}
}

func (v NullableApiDependencyTreeNodeDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDependencyTreeNodeDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


