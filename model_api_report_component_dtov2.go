/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.166.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiReportComponentDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReportComponentDTOV2{}

// ApiReportComponentDTOV2 struct for ApiReportComponentDTOV2
type ApiReportComponentDTOV2 struct {
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	DependencyData *ApiDependencyDataDTO `json:"dependencyData,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Hash *string `json:"hash,omitempty"`
	IdentificationSource *string `json:"identificationSource,omitempty"`
	LicenseData *ApiLicenseDataDTOV2 `json:"licenseData,omitempty"`
	MatchState *string `json:"matchState,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
	Pathnames []string `json:"pathnames,omitempty"`
	Proprietary *bool `json:"proprietary,omitempty"`
	SecurityData *ApiSecurityDataDTO `json:"securityData,omitempty"`
	Sha256 *string `json:"sha256,omitempty"`
	ThirdParty *bool `json:"thirdParty,omitempty"`
}

// NewApiReportComponentDTOV2 instantiates a new ApiReportComponentDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReportComponentDTOV2() *ApiReportComponentDTOV2 {
	this := ApiReportComponentDTOV2{}
	return &this
}

// NewApiReportComponentDTOV2WithDefaults instantiates a new ApiReportComponentDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReportComponentDTOV2WithDefaults() *ApiReportComponentDTOV2 {
	this := ApiReportComponentDTOV2{}
	return &this
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiReportComponentDTOV2) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetDependencyData returns the DependencyData field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetDependencyData() ApiDependencyDataDTO {
	if o == nil || IsNil(o.DependencyData) {
		var ret ApiDependencyDataDTO
		return ret
	}
	return *o.DependencyData
}

// GetDependencyDataOk returns a tuple with the DependencyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetDependencyDataOk() (*ApiDependencyDataDTO, bool) {
	if o == nil || IsNil(o.DependencyData) {
		return nil, false
	}
	return o.DependencyData, true
}

// HasDependencyData returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasDependencyData() bool {
	if o != nil && !IsNil(o.DependencyData) {
		return true
	}

	return false
}

// SetDependencyData gets a reference to the given ApiDependencyDataDTO and assigns it to the DependencyData field.
func (o *ApiReportComponentDTOV2) SetDependencyData(v ApiDependencyDataDTO) {
	o.DependencyData = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ApiReportComponentDTOV2) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ApiReportComponentDTOV2) SetHash(v string) {
	o.Hash = &v
}

// GetIdentificationSource returns the IdentificationSource field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetIdentificationSource() string {
	if o == nil || IsNil(o.IdentificationSource) {
		var ret string
		return ret
	}
	return *o.IdentificationSource
}

// GetIdentificationSourceOk returns a tuple with the IdentificationSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetIdentificationSourceOk() (*string, bool) {
	if o == nil || IsNil(o.IdentificationSource) {
		return nil, false
	}
	return o.IdentificationSource, true
}

// HasIdentificationSource returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasIdentificationSource() bool {
	if o != nil && !IsNil(o.IdentificationSource) {
		return true
	}

	return false
}

// SetIdentificationSource gets a reference to the given string and assigns it to the IdentificationSource field.
func (o *ApiReportComponentDTOV2) SetIdentificationSource(v string) {
	o.IdentificationSource = &v
}

// GetLicenseData returns the LicenseData field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetLicenseData() ApiLicenseDataDTOV2 {
	if o == nil || IsNil(o.LicenseData) {
		var ret ApiLicenseDataDTOV2
		return ret
	}
	return *o.LicenseData
}

// GetLicenseDataOk returns a tuple with the LicenseData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetLicenseDataOk() (*ApiLicenseDataDTOV2, bool) {
	if o == nil || IsNil(o.LicenseData) {
		return nil, false
	}
	return o.LicenseData, true
}

// HasLicenseData returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasLicenseData() bool {
	if o != nil && !IsNil(o.LicenseData) {
		return true
	}

	return false
}

// SetLicenseData gets a reference to the given ApiLicenseDataDTOV2 and assigns it to the LicenseData field.
func (o *ApiReportComponentDTOV2) SetLicenseData(v ApiLicenseDataDTOV2) {
	o.LicenseData = &v
}

// GetMatchState returns the MatchState field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetMatchState() string {
	if o == nil || IsNil(o.MatchState) {
		var ret string
		return ret
	}
	return *o.MatchState
}

// GetMatchStateOk returns a tuple with the MatchState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetMatchStateOk() (*string, bool) {
	if o == nil || IsNil(o.MatchState) {
		return nil, false
	}
	return o.MatchState, true
}

// HasMatchState returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasMatchState() bool {
	if o != nil && !IsNil(o.MatchState) {
		return true
	}

	return false
}

// SetMatchState gets a reference to the given string and assigns it to the MatchState field.
func (o *ApiReportComponentDTOV2) SetMatchState(v string) {
	o.MatchState = &v
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ApiReportComponentDTOV2) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

// GetPathnames returns the Pathnames field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetPathnames() []string {
	if o == nil || IsNil(o.Pathnames) {
		var ret []string
		return ret
	}
	return o.Pathnames
}

// GetPathnamesOk returns a tuple with the Pathnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetPathnamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Pathnames) {
		return nil, false
	}
	return o.Pathnames, true
}

// HasPathnames returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasPathnames() bool {
	if o != nil && !IsNil(o.Pathnames) {
		return true
	}

	return false
}

// SetPathnames gets a reference to the given []string and assigns it to the Pathnames field.
func (o *ApiReportComponentDTOV2) SetPathnames(v []string) {
	o.Pathnames = v
}

// GetProprietary returns the Proprietary field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetProprietary() bool {
	if o == nil || IsNil(o.Proprietary) {
		var ret bool
		return ret
	}
	return *o.Proprietary
}

// GetProprietaryOk returns a tuple with the Proprietary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetProprietaryOk() (*bool, bool) {
	if o == nil || IsNil(o.Proprietary) {
		return nil, false
	}
	return o.Proprietary, true
}

// HasProprietary returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasProprietary() bool {
	if o != nil && !IsNil(o.Proprietary) {
		return true
	}

	return false
}

// SetProprietary gets a reference to the given bool and assigns it to the Proprietary field.
func (o *ApiReportComponentDTOV2) SetProprietary(v bool) {
	o.Proprietary = &v
}

// GetSecurityData returns the SecurityData field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetSecurityData() ApiSecurityDataDTO {
	if o == nil || IsNil(o.SecurityData) {
		var ret ApiSecurityDataDTO
		return ret
	}
	return *o.SecurityData
}

// GetSecurityDataOk returns a tuple with the SecurityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetSecurityDataOk() (*ApiSecurityDataDTO, bool) {
	if o == nil || IsNil(o.SecurityData) {
		return nil, false
	}
	return o.SecurityData, true
}

// HasSecurityData returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasSecurityData() bool {
	if o != nil && !IsNil(o.SecurityData) {
		return true
	}

	return false
}

// SetSecurityData gets a reference to the given ApiSecurityDataDTO and assigns it to the SecurityData field.
func (o *ApiReportComponentDTOV2) SetSecurityData(v ApiSecurityDataDTO) {
	o.SecurityData = &v
}

// GetSha256 returns the Sha256 field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetSha256() string {
	if o == nil || IsNil(o.Sha256) {
		var ret string
		return ret
	}
	return *o.Sha256
}

// GetSha256Ok returns a tuple with the Sha256 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetSha256Ok() (*string, bool) {
	if o == nil || IsNil(o.Sha256) {
		return nil, false
	}
	return o.Sha256, true
}

// HasSha256 returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasSha256() bool {
	if o != nil && !IsNil(o.Sha256) {
		return true
	}

	return false
}

// SetSha256 gets a reference to the given string and assigns it to the Sha256 field.
func (o *ApiReportComponentDTOV2) SetSha256(v string) {
	o.Sha256 = &v
}

// GetThirdParty returns the ThirdParty field value if set, zero value otherwise.
func (o *ApiReportComponentDTOV2) GetThirdParty() bool {
	if o == nil || IsNil(o.ThirdParty) {
		var ret bool
		return ret
	}
	return *o.ThirdParty
}

// GetThirdPartyOk returns a tuple with the ThirdParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportComponentDTOV2) GetThirdPartyOk() (*bool, bool) {
	if o == nil || IsNil(o.ThirdParty) {
		return nil, false
	}
	return o.ThirdParty, true
}

// HasThirdParty returns a boolean if a field has been set.
func (o *ApiReportComponentDTOV2) HasThirdParty() bool {
	if o != nil && !IsNil(o.ThirdParty) {
		return true
	}

	return false
}

// SetThirdParty gets a reference to the given bool and assigns it to the ThirdParty field.
func (o *ApiReportComponentDTOV2) SetThirdParty(v bool) {
	o.ThirdParty = &v
}

func (o ApiReportComponentDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReportComponentDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.DependencyData) {
		toSerialize["dependencyData"] = o.DependencyData
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.IdentificationSource) {
		toSerialize["identificationSource"] = o.IdentificationSource
	}
	if !IsNil(o.LicenseData) {
		toSerialize["licenseData"] = o.LicenseData
	}
	if !IsNil(o.MatchState) {
		toSerialize["matchState"] = o.MatchState
	}
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
	}
	if !IsNil(o.Pathnames) {
		toSerialize["pathnames"] = o.Pathnames
	}
	if !IsNil(o.Proprietary) {
		toSerialize["proprietary"] = o.Proprietary
	}
	if !IsNil(o.SecurityData) {
		toSerialize["securityData"] = o.SecurityData
	}
	if !IsNil(o.Sha256) {
		toSerialize["sha256"] = o.Sha256
	}
	if !IsNil(o.ThirdParty) {
		toSerialize["thirdParty"] = o.ThirdParty
	}
	return toSerialize, nil
}

type NullableApiReportComponentDTOV2 struct {
	value *ApiReportComponentDTOV2
	isSet bool
}

func (v NullableApiReportComponentDTOV2) Get() *ApiReportComponentDTOV2 {
	return v.value
}

func (v *NullableApiReportComponentDTOV2) Set(val *ApiReportComponentDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReportComponentDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReportComponentDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReportComponentDTOV2(val *ApiReportComponentDTOV2) *NullableApiReportComponentDTOV2 {
	return &NullableApiReportComponentDTOV2{value: val, isSet: true}
}

func (v NullableApiReportComponentDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReportComponentDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


