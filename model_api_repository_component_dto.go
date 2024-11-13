/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.184.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiRepositoryComponentDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryComponentDTO{}

// ApiRepositoryComponentDTO struct for ApiRepositoryComponentDTO
type ApiRepositoryComponentDTO struct {
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Hash *string `json:"hash,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
	QuarantineId *string `json:"quarantineId,omitempty"`
	QuarantineReleaseTime *time.Time `json:"quarantineReleaseTime,omitempty"`
	QuarantineTime *time.Time `json:"quarantineTime,omitempty"`
	Sha256 *string `json:"sha256,omitempty"`
	ThirdParty *bool `json:"thirdParty,omitempty"`
}

// NewApiRepositoryComponentDTO instantiates a new ApiRepositoryComponentDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryComponentDTO() *ApiRepositoryComponentDTO {
	this := ApiRepositoryComponentDTO{}
	return &this
}

// NewApiRepositoryComponentDTOWithDefaults instantiates a new ApiRepositoryComponentDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryComponentDTOWithDefaults() *ApiRepositoryComponentDTO {
	this := ApiRepositoryComponentDTO{}
	return &this
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiRepositoryComponentDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiRepositoryComponentDTO) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiRepositoryComponentDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ApiRepositoryComponentDTO) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentDTO) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ApiRepositoryComponentDTO) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ApiRepositoryComponentDTO) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ApiRepositoryComponentDTO) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentDTO) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ApiRepositoryComponentDTO) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ApiRepositoryComponentDTO) SetHash(v string) {
	o.Hash = &v
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ApiRepositoryComponentDTO) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentDTO) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ApiRepositoryComponentDTO) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ApiRepositoryComponentDTO) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

// GetQuarantineId returns the QuarantineId field value if set, zero value otherwise.
func (o *ApiRepositoryComponentDTO) GetQuarantineId() string {
	if o == nil || IsNil(o.QuarantineId) {
		var ret string
		return ret
	}
	return *o.QuarantineId
}

// GetQuarantineIdOk returns a tuple with the QuarantineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentDTO) GetQuarantineIdOk() (*string, bool) {
	if o == nil || IsNil(o.QuarantineId) {
		return nil, false
	}
	return o.QuarantineId, true
}

// HasQuarantineId returns a boolean if a field has been set.
func (o *ApiRepositoryComponentDTO) HasQuarantineId() bool {
	if o != nil && !IsNil(o.QuarantineId) {
		return true
	}

	return false
}

// SetQuarantineId gets a reference to the given string and assigns it to the QuarantineId field.
func (o *ApiRepositoryComponentDTO) SetQuarantineId(v string) {
	o.QuarantineId = &v
}

// GetQuarantineReleaseTime returns the QuarantineReleaseTime field value if set, zero value otherwise.
func (o *ApiRepositoryComponentDTO) GetQuarantineReleaseTime() time.Time {
	if o == nil || IsNil(o.QuarantineReleaseTime) {
		var ret time.Time
		return ret
	}
	return *o.QuarantineReleaseTime
}

// GetQuarantineReleaseTimeOk returns a tuple with the QuarantineReleaseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentDTO) GetQuarantineReleaseTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.QuarantineReleaseTime) {
		return nil, false
	}
	return o.QuarantineReleaseTime, true
}

// HasQuarantineReleaseTime returns a boolean if a field has been set.
func (o *ApiRepositoryComponentDTO) HasQuarantineReleaseTime() bool {
	if o != nil && !IsNil(o.QuarantineReleaseTime) {
		return true
	}

	return false
}

// SetQuarantineReleaseTime gets a reference to the given time.Time and assigns it to the QuarantineReleaseTime field.
func (o *ApiRepositoryComponentDTO) SetQuarantineReleaseTime(v time.Time) {
	o.QuarantineReleaseTime = &v
}

// GetQuarantineTime returns the QuarantineTime field value if set, zero value otherwise.
func (o *ApiRepositoryComponentDTO) GetQuarantineTime() time.Time {
	if o == nil || IsNil(o.QuarantineTime) {
		var ret time.Time
		return ret
	}
	return *o.QuarantineTime
}

// GetQuarantineTimeOk returns a tuple with the QuarantineTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentDTO) GetQuarantineTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.QuarantineTime) {
		return nil, false
	}
	return o.QuarantineTime, true
}

// HasQuarantineTime returns a boolean if a field has been set.
func (o *ApiRepositoryComponentDTO) HasQuarantineTime() bool {
	if o != nil && !IsNil(o.QuarantineTime) {
		return true
	}

	return false
}

// SetQuarantineTime gets a reference to the given time.Time and assigns it to the QuarantineTime field.
func (o *ApiRepositoryComponentDTO) SetQuarantineTime(v time.Time) {
	o.QuarantineTime = &v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *ApiRepositoryComponentDTO) GetSha256() string {
	if o == nil || IsNil(o.Sha256) {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentDTO) GetSha256Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha256) {
		return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *ApiRepositoryComponentDTO) HasSha256() bool {
	if o != nil && !IsNil(o.Sha256) {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *ApiRepositoryComponentDTO) SetSha256(v string) {
	o.Sha256 = &v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *ApiRepositoryComponentDTO) GetThirdParty() bool {
	if o == nil || IsNil(o.ThirdParty) {
		var ret bool
		return ret
	}
	return *o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentDTO) GetThirdPartyOk() (*bool, bool) {
	if o == nil || IsNil(o.ThirdParty) {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *ApiRepositoryComponentDTO) HasThirdParty() bool {
	if o != nil && !IsNil(o.ThirdParty) {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given bool and assigns it to the ThirdParty field.
func (o *ApiRepositoryComponentDTO) SetThirdParty(v bool) {
	o.ThirdParty = &v
}

func (o ApiRepositoryComponentDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryComponentDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
	}
	if !IsNil(o.QuarantineId) {
		toSerialize["quarantineId"] = o.QuarantineId
	}
	if !IsNil(o.QuarantineReleaseTime) {
		toSerialize["quarantineReleaseTime"] = o.QuarantineReleaseTime
	}
	if !IsNil(o.QuarantineTime) {
		toSerialize["quarantineTime"] = o.QuarantineTime
	}
	if !IsNil(o.Sha256) {
		toSerialize["sha256"] = o.Sha256
	}
	if !IsNil(o.ThirdParty) {
		toSerialize["thirdParty"] = o.ThirdParty
	}
	return toSerialize, nil
}

type NullableApiRepositoryComponentDTO struct {
	value *ApiRepositoryComponentDTO
	isSet bool
}

func (v NullableApiRepositoryComponentDTO) Get() *ApiRepositoryComponentDTO {
	return v.value
}

func (v *NullableApiRepositoryComponentDTO) Set(val *ApiRepositoryComponentDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryComponentDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryComponentDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryComponentDTO(val *ApiRepositoryComponentDTO) *NullableApiRepositoryComponentDTO {
	return &NullableApiRepositoryComponentDTO{value: val, isSet: true}
}

func (v NullableApiRepositoryComponentDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryComponentDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


