/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.177.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiComponentProjectDataDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiComponentProjectDataDTO{}

// ApiComponentProjectDataDTO struct for ApiComponentProjectDataDTO
type ApiComponentProjectDataDTO struct {
	FirstReleaseDate *time.Time `json:"firstReleaseDate,omitempty"`
	LastReleaseDate *time.Time `json:"lastReleaseDate,omitempty"`
	ProjectMetadata *ApiComponentProjectMetadataDTO `json:"projectMetadata,omitempty"`
	SourceControlManagement *ApiComponentProjectScmDTO `json:"sourceControlManagement,omitempty"`
}

// NewApiComponentProjectDataDTO instantiates a new ApiComponentProjectDataDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiComponentProjectDataDTO() *ApiComponentProjectDataDTO {
	this := ApiComponentProjectDataDTO{}
	return &this
}

// NewApiComponentProjectDataDTOWithDefaults instantiates a new ApiComponentProjectDataDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiComponentProjectDataDTOWithDefaults() *ApiComponentProjectDataDTO {
	this := ApiComponentProjectDataDTO{}
	return &this
}

// GetFirstReleaseDate returns the FirstReleaseDate field value if set, zero value otherwise.
func (o *ApiComponentProjectDataDTO) GetFirstReleaseDate() time.Time {
	if o == nil || IsNil(o.FirstReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.FirstReleaseDate
}

// GetFirstReleaseDateOk returns a tuple with the FirstReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectDataDTO) GetFirstReleaseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FirstReleaseDate) {
		return nil, false
	}
	return o.FirstReleaseDate, true
}

// HasFirstReleaseDate returns a boolean if a field has been set.
func (o *ApiComponentProjectDataDTO) HasFirstReleaseDate() bool {
	if o != nil && !IsNil(o.FirstReleaseDate) {
		return true
	}

	return false
}

// SetFirstReleaseDate gets a reference to the given time.Time and assigns it to the FirstReleaseDate field.
func (o *ApiComponentProjectDataDTO) SetFirstReleaseDate(v time.Time) {
	o.FirstReleaseDate = &v
}

// GetLastReleaseDate returns the LastReleaseDate field value if set, zero value otherwise.
func (o *ApiComponentProjectDataDTO) GetLastReleaseDate() time.Time {
	if o == nil || IsNil(o.LastReleaseDate) {
		var ret time.Time
		return ret
	}
	return *o.LastReleaseDate
}

// GetLastReleaseDateOk returns a tuple with the LastReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectDataDTO) GetLastReleaseDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastReleaseDate) {
		return nil, false
	}
	return o.LastReleaseDate, true
}

// HasLastReleaseDate returns a boolean if a field has been set.
func (o *ApiComponentProjectDataDTO) HasLastReleaseDate() bool {
	if o != nil && !IsNil(o.LastReleaseDate) {
		return true
	}

	return false
}

// SetLastReleaseDate gets a reference to the given time.Time and assigns it to the LastReleaseDate field.
func (o *ApiComponentProjectDataDTO) SetLastReleaseDate(v time.Time) {
	o.LastReleaseDate = &v
}

// GetProjectMetadata returns the ProjectMetadata field value if set, zero value otherwise.
func (o *ApiComponentProjectDataDTO) GetProjectMetadata() ApiComponentProjectMetadataDTO {
	if o == nil || IsNil(o.ProjectMetadata) {
		var ret ApiComponentProjectMetadataDTO
		return ret
	}
	return *o.ProjectMetadata
}

// GetProjectMetadataOk returns a tuple with the ProjectMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectDataDTO) GetProjectMetadataOk() (*ApiComponentProjectMetadataDTO, bool) {
	if o == nil || IsNil(o.ProjectMetadata) {
		return nil, false
	}
	return o.ProjectMetadata, true
}

// HasProjectMetadata returns a boolean if a field has been set.
func (o *ApiComponentProjectDataDTO) HasProjectMetadata() bool {
	if o != nil && !IsNil(o.ProjectMetadata) {
		return true
	}

	return false
}

// SetProjectMetadata gets a reference to the given ApiComponentProjectMetadataDTO and assigns it to the ProjectMetadata field.
func (o *ApiComponentProjectDataDTO) SetProjectMetadata(v ApiComponentProjectMetadataDTO) {
	o.ProjectMetadata = &v
}

// GetSourceControlManagement returns the SourceControlManagement field value if set, zero value otherwise.
func (o *ApiComponentProjectDataDTO) GetSourceControlManagement() ApiComponentProjectScmDTO {
	if o == nil || IsNil(o.SourceControlManagement) {
		var ret ApiComponentProjectScmDTO
		return ret
	}
	return *o.SourceControlManagement
}

// GetSourceControlManagementOk returns a tuple with the SourceControlManagement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiComponentProjectDataDTO) GetSourceControlManagementOk() (*ApiComponentProjectScmDTO, bool) {
	if o == nil || IsNil(o.SourceControlManagement) {
		return nil, false
	}
	return o.SourceControlManagement, true
}

// HasSourceControlManagement returns a boolean if a field has been set.
func (o *ApiComponentProjectDataDTO) HasSourceControlManagement() bool {
	if o != nil && !IsNil(o.SourceControlManagement) {
		return true
	}

	return false
}

// SetSourceControlManagement gets a reference to the given ApiComponentProjectScmDTO and assigns it to the SourceControlManagement field.
func (o *ApiComponentProjectDataDTO) SetSourceControlManagement(v ApiComponentProjectScmDTO) {
	o.SourceControlManagement = &v
}

func (o ApiComponentProjectDataDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiComponentProjectDataDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstReleaseDate) {
		toSerialize["firstReleaseDate"] = o.FirstReleaseDate
	}
	if !IsNil(o.LastReleaseDate) {
		toSerialize["lastReleaseDate"] = o.LastReleaseDate
	}
	if !IsNil(o.ProjectMetadata) {
		toSerialize["projectMetadata"] = o.ProjectMetadata
	}
	if !IsNil(o.SourceControlManagement) {
		toSerialize["sourceControlManagement"] = o.SourceControlManagement
	}
	return toSerialize, nil
}

type NullableApiComponentProjectDataDTO struct {
	value *ApiComponentProjectDataDTO
	isSet bool
}

func (v NullableApiComponentProjectDataDTO) Get() *ApiComponentProjectDataDTO {
	return v.value
}

func (v *NullableApiComponentProjectDataDTO) Set(val *ApiComponentProjectDataDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiComponentProjectDataDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiComponentProjectDataDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiComponentProjectDataDTO(val *ApiComponentProjectDataDTO) *NullableApiComponentProjectDataDTO {
	return &NullableApiComponentProjectDataDTO{value: val, isSet: true}
}

func (v NullableApiComponentProjectDataDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiComponentProjectDataDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


