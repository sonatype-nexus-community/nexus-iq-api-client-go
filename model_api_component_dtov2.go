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

// checks if the ApiComponentDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentDTOV2{}

// ApiComponentDTOV2 struct for ApiComponentDTOV2
type ApiComponentDTOV2 struct {
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Hash NullableString `json:"hash,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
	Proprietary *bool `json:"proprietary,omitempty"`
	Sha256 *string `json:"sha256,omitempty"`
	ThirdParty *bool `json:"thirdParty,omitempty"`
}

// NewApiComponentDTOV2 instantiates a new ApiComponentDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentDTOV2() *ApiComponentDTOV2 {
	this := ApiComponentDTOV2{}
	return &this
}

// NewApiComponentDTOV2WithDefaults instantiates a new ApiComponentDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentDTOV2WithDefaults() *ApiComponentDTOV2 {
	this := ApiComponentDTOV2{}
	return &this
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiComponentDTOV2) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentDTOV2) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiComponentDTOV2) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiComponentDTOV2) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ApiComponentDTOV2) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentDTOV2) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ApiComponentDTOV2) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ApiComponentDTOV2) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHash returns the Hash field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiComponentDTOV2) GetHash() string {
	if o == nil || IsNil(o.Hash.Get()) {
		var ret string
		return ret
	}
	return *o.Hash.Get()
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiComponentDTOV2) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hash.Get(), o.Hash.IsSet()
}

// HasHash returns a boolean if a field has been set.
func (o *ApiComponentDTOV2) HasHash() bool {
	if o != nil && o.Hash.IsSet() {
		return true
	}

	return false
}

// SetHash gets a reference to the given NullableString and assigns it to the Hash field.
func (o *ApiComponentDTOV2) SetHash(v string) {
	o.Hash.Set(&v)
}
// SetHashNil sets the value for Hash to be an explicit nil
func (o *ApiComponentDTOV2) SetHashNil() {
	o.Hash.Set(nil)
}

// UnsetHash ensures that no value is present for Hash, not even an explicit nil
func (o *ApiComponentDTOV2) UnsetHash() {
	o.Hash.Unset()
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ApiComponentDTOV2) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentDTOV2) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ApiComponentDTOV2) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ApiComponentDTOV2) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

// GetProprietary returns the Proprietary field value if set, zero value otherwise.
func (o *ApiComponentDTOV2) GetProprietary() bool {
	if o == nil || IsNil(o.Proprietary) {
		var ret bool
		return ret
	}
	return *o.Proprietary
}

// GetProprietaryOk returns a tuple with the Proprietary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentDTOV2) GetProprietaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Proprietary) {
		return nil, false
	}
	return o.Proprietary, true
}

// HasProprietary returns a boolean if a field has been set.
func (o *ApiComponentDTOV2) HasProprietary() bool {
	if o != nil && !IsNil(o.Proprietary) {
		return true
	}

	return false
}

// SetProprietary gets a reference to the given bool and assigns it to the Proprietary field.
func (o *ApiComponentDTOV2) SetProprietary(v bool) {
	o.Proprietary = &v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *ApiComponentDTOV2) GetSha256() string {
	if o == nil || IsNil(o.Sha256) {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentDTOV2) GetSha256Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha256) {
		return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *ApiComponentDTOV2) HasSha256() bool {
	if o != nil && !IsNil(o.Sha256) {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *ApiComponentDTOV2) SetSha256(v string) {
	o.Sha256 = &v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *ApiComponentDTOV2) GetThirdParty() bool {
	if o == nil || IsNil(o.ThirdParty) {
		var ret bool
		return ret
	}
	return *o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentDTOV2) GetThirdPartyOk() (*bool, bool) {
	if o == nil || IsNil(o.ThirdParty) {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *ApiComponentDTOV2) HasThirdParty() bool {
	if o != nil && !IsNil(o.ThirdParty) {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given bool and assigns it to the ThirdParty field.
func (o *ApiComponentDTOV2) SetThirdParty(v bool) {
	o.ThirdParty = &v
}

func (o ApiComponentDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Hash.IsSet() {
		toSerialize["hash"] = o.Hash.Get()
	}
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
	}
	if !IsNil(o.Proprietary) {
		toSerialize["proprietary"] = o.Proprietary
	}
	if !IsNil(o.Sha256) {
		toSerialize["sha256"] = o.Sha256
	}
	if !IsNil(o.ThirdParty) {
		toSerialize["thirdParty"] = o.ThirdParty
	}
	return toSerialize, nil
}

type NullableApiComponentDTOV2 struct {
	value *ApiComponentDTOV2
	isSet bool
}

func (v NullableApiComponentDTOV2) Get() *ApiComponentDTOV2 {
	return v.value
}

func (v *NullableApiComponentDTOV2) Set(val *ApiComponentDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentDTOV2(val *ApiComponentDTOV2) *NullableApiComponentDTOV2 {
	return &NullableApiComponentDTOV2{value: val, isSet: true}
}

func (v NullableApiComponentDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


