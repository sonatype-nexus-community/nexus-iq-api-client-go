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

// checks if the ApiSearchResultDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSearchResultDTOV2{}

// ApiSearchResultDTOV2 struct for ApiSearchResultDTOV2
type ApiSearchResultDTOV2 struct {
	ApplicationId *string `json:"applicationId,omitempty"`
	ApplicationName *string `json:"applicationName,omitempty"`
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	DependencyData *ApiDependencyDataDTO `json:"dependencyData,omitempty"`
	Hash *string `json:"hash,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
	ReportHtmlUrl *string `json:"reportHtmlUrl,omitempty"`
	ReportUrl *string `json:"reportUrl,omitempty"`
	ThreatLevel *int32 `json:"threatLevel,omitempty"`
}

// NewApiSearchResultDTOV2 instantiates a new ApiSearchResultDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSearchResultDTOV2() *ApiSearchResultDTOV2 {
	this := ApiSearchResultDTOV2{}
	return &this
}

// NewApiSearchResultDTOV2WithDefaults instantiates a new ApiSearchResultDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSearchResultDTOV2WithDefaults() *ApiSearchResultDTOV2 {
	this := ApiSearchResultDTOV2{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ApiSearchResultDTOV2) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchResultDTOV2) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ApiSearchResultDTOV2) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ApiSearchResultDTOV2) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *ApiSearchResultDTOV2) GetApplicationName() string {
	if o == nil || IsNil(o.ApplicationName) {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchResultDTOV2) GetApplicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationName) {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *ApiSearchResultDTOV2) HasApplicationName() bool {
	if o != nil && !IsNil(o.ApplicationName) {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *ApiSearchResultDTOV2) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiSearchResultDTOV2) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchResultDTOV2) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiSearchResultDTOV2) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiSearchResultDTOV2) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetDependencyData returns the DependencyData field value if set, zero value otherwise.
func (o *ApiSearchResultDTOV2) GetDependencyData() ApiDependencyDataDTO {
	if o == nil || IsNil(o.DependencyData) {
		var ret ApiDependencyDataDTO
		return ret
	}
	return *o.DependencyData
}

// GetDependencyDataOk returns a tuple with the DependencyData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchResultDTOV2) GetDependencyDataOk() (*ApiDependencyDataDTO, bool) {
	if o == nil || IsNil(o.DependencyData) {
		return nil, false
	}
	return o.DependencyData, true
}

// HasDependencyData returns a boolean if a field has been set.
func (o *ApiSearchResultDTOV2) HasDependencyData() bool {
	if o != nil && !IsNil(o.DependencyData) {
		return true
	}

	return false
}

// SetDependencyData gets a reference to the given ApiDependencyDataDTO and assigns it to the DependencyData field.
func (o *ApiSearchResultDTOV2) SetDependencyData(v ApiDependencyDataDTO) {
	o.DependencyData = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ApiSearchResultDTOV2) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchResultDTOV2) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ApiSearchResultDTOV2) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ApiSearchResultDTOV2) SetHash(v string) {
	o.Hash = &v
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ApiSearchResultDTOV2) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchResultDTOV2) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ApiSearchResultDTOV2) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ApiSearchResultDTOV2) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

// GetReportHtmlUrl returns the ReportHtmlUrl field value if set, zero value otherwise.
func (o *ApiSearchResultDTOV2) GetReportHtmlUrl() string {
	if o == nil || IsNil(o.ReportHtmlUrl) {
		var ret string
		return ret
	}
	return *o.ReportHtmlUrl
}

// GetReportHtmlUrlOk returns a tuple with the ReportHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchResultDTOV2) GetReportHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportHtmlUrl) {
		return nil, false
	}
	return o.ReportHtmlUrl, true
}

// HasReportHtmlUrl returns a boolean if a field has been set.
func (o *ApiSearchResultDTOV2) HasReportHtmlUrl() bool {
	if o != nil && !IsNil(o.ReportHtmlUrl) {
		return true
	}

	return false
}

// SetReportHtmlUrl gets a reference to the given string and assigns it to the ReportHtmlUrl field.
func (o *ApiSearchResultDTOV2) SetReportHtmlUrl(v string) {
	o.ReportHtmlUrl = &v
}

// GetReportUrl returns the ReportUrl field value if set, zero value otherwise.
func (o *ApiSearchResultDTOV2) GetReportUrl() string {
	if o == nil || IsNil(o.ReportUrl) {
		var ret string
		return ret
	}
	return *o.ReportUrl
}

// GetReportUrlOk returns a tuple with the ReportUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchResultDTOV2) GetReportUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportUrl) {
		return nil, false
	}
	return o.ReportUrl, true
}

// HasReportUrl returns a boolean if a field has been set.
func (o *ApiSearchResultDTOV2) HasReportUrl() bool {
	if o != nil && !IsNil(o.ReportUrl) {
		return true
	}

	return false
}

// SetReportUrl gets a reference to the given string and assigns it to the ReportUrl field.
func (o *ApiSearchResultDTOV2) SetReportUrl(v string) {
	o.ReportUrl = &v
}

// GetThreatLevel returns the ThreatLevel field value if set, zero value otherwise.
func (o *ApiSearchResultDTOV2) GetThreatLevel() int32 {
	if o == nil || IsNil(o.ThreatLevel) {
		var ret int32
		return ret
	}
	return *o.ThreatLevel
}

// GetThreatLevelOk returns a tuple with the ThreatLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSearchResultDTOV2) GetThreatLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.ThreatLevel) {
		return nil, false
	}
	return o.ThreatLevel, true
}

// HasThreatLevel returns a boolean if a field has been set.
func (o *ApiSearchResultDTOV2) HasThreatLevel() bool {
	if o != nil && !IsNil(o.ThreatLevel) {
		return true
	}

	return false
}

// SetThreatLevel gets a reference to the given int32 and assigns it to the ThreatLevel field.
func (o *ApiSearchResultDTOV2) SetThreatLevel(v int32) {
	o.ThreatLevel = &v
}

func (o ApiSearchResultDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSearchResultDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.ApplicationName) {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.DependencyData) {
		toSerialize["dependencyData"] = o.DependencyData
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
	}
	if !IsNil(o.ReportHtmlUrl) {
		toSerialize["reportHtmlUrl"] = o.ReportHtmlUrl
	}
	if !IsNil(o.ReportUrl) {
		toSerialize["reportUrl"] = o.ReportUrl
	}
	if !IsNil(o.ThreatLevel) {
		toSerialize["threatLevel"] = o.ThreatLevel
	}
	return toSerialize, nil
}

type NullableApiSearchResultDTOV2 struct {
	value *ApiSearchResultDTOV2
	isSet bool
}

func (v NullableApiSearchResultDTOV2) Get() *ApiSearchResultDTOV2 {
	return v.value
}

func (v *NullableApiSearchResultDTOV2) Set(val *ApiSearchResultDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSearchResultDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSearchResultDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSearchResultDTOV2(val *ApiSearchResultDTOV2) *NullableApiSearchResultDTOV2 {
	return &NullableApiSearchResultDTOV2{value: val, isSet: true}
}

func (v NullableApiSearchResultDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSearchResultDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


