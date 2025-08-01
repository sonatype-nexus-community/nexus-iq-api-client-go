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

// checks if the ApiSbomStatusDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSbomStatusDTO{}

// ApiSbomStatusDTO struct for ApiSbomStatusDTO
type ApiSbomStatusDTO struct {
	ApplicationId *string `json:"applicationId,omitempty"`
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	ErrorMessage *string `json:"errorMessage,omitempty"`
	IsError *bool `json:"isError,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewApiSbomStatusDTO instantiates a new ApiSbomStatusDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSbomStatusDTO() *ApiSbomStatusDTO {
	this := ApiSbomStatusDTO{}
	return &this
}

// NewApiSbomStatusDTOWithDefaults instantiates a new ApiSbomStatusDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSbomStatusDTOWithDefaults() *ApiSbomStatusDTO {
	this := ApiSbomStatusDTO{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ApiSbomStatusDTO) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSbomStatusDTO) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ApiSbomStatusDTO) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ApiSbomStatusDTO) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetDownloadUrl returns the DownloadUrl field value if set, zero value otherwise.
func (o *ApiSbomStatusDTO) GetDownloadUrl() string {
	if o == nil || IsNil(o.DownloadUrl) {
		var ret string
		return ret
	}
	return *o.DownloadUrl
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSbomStatusDTO) GetDownloadUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DownloadUrl) {
		return nil, false
	}
	return o.DownloadUrl, true
}

// HasDownloadUrl returns a boolean if a field has been set.
func (o *ApiSbomStatusDTO) HasDownloadUrl() bool {
	if o != nil && !IsNil(o.DownloadUrl) {
		return true
	}

	return false
}

// SetDownloadUrl gets a reference to the given string and assigns it to the DownloadUrl field.
func (o *ApiSbomStatusDTO) SetDownloadUrl(v string) {
	o.DownloadUrl = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ApiSbomStatusDTO) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSbomStatusDTO) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ApiSbomStatusDTO) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ApiSbomStatusDTO) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetIsError returns the IsError field value if set, zero value otherwise.
func (o *ApiSbomStatusDTO) GetIsError() bool {
	if o == nil || IsNil(o.IsError) {
		var ret bool
		return ret
	}
	return *o.IsError
}

// GetIsErrorOk returns a tuple with the IsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSbomStatusDTO) GetIsErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.IsError) {
		return nil, false
	}
	return o.IsError, true
}

// HasIsError returns a boolean if a field has been set.
func (o *ApiSbomStatusDTO) HasIsError() bool {
	if o != nil && !IsNil(o.IsError) {
		return true
	}

	return false
}

// SetIsError gets a reference to the given bool and assigns it to the IsError field.
func (o *ApiSbomStatusDTO) SetIsError(v bool) {
	o.IsError = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApiSbomStatusDTO) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSbomStatusDTO) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApiSbomStatusDTO) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApiSbomStatusDTO) SetVersion(v string) {
	o.Version = &v
}

func (o ApiSbomStatusDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSbomStatusDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.DownloadUrl) {
		toSerialize["downloadUrl"] = o.DownloadUrl
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if !IsNil(o.IsError) {
		toSerialize["isError"] = o.IsError
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableApiSbomStatusDTO struct {
	value *ApiSbomStatusDTO
	isSet bool
}

func (v NullableApiSbomStatusDTO) Get() *ApiSbomStatusDTO {
	return v.value
}

func (v *NullableApiSbomStatusDTO) Set(val *ApiSbomStatusDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSbomStatusDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSbomStatusDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSbomStatusDTO(val *ApiSbomStatusDTO) *NullableApiSbomStatusDTO {
	return &NullableApiSbomStatusDTO{value: val, isSet: true}
}

func (v NullableApiSbomStatusDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSbomStatusDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


