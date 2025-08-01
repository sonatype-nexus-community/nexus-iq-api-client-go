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

// checks if the ApiLicenseOverrideDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicenseOverrideDTO{}

// ApiLicenseOverrideDTO struct for ApiLicenseOverrideDTO
type ApiLicenseOverrideDTO struct {
	Comment *string `json:"comment,omitempty"`
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	Id *string `json:"id,omitempty"`
	LicenseIds []string `json:"licenseIds,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewApiLicenseOverrideDTO instantiates a new ApiLicenseOverrideDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicenseOverrideDTO() *ApiLicenseOverrideDTO {
	this := ApiLicenseOverrideDTO{}
	return &this
}

// NewApiLicenseOverrideDTOWithDefaults instantiates a new ApiLicenseOverrideDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicenseOverrideDTOWithDefaults() *ApiLicenseOverrideDTO {
	this := ApiLicenseOverrideDTO{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApiLicenseOverrideDTO) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseOverrideDTO) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApiLicenseOverrideDTO) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApiLicenseOverrideDTO) SetComment(v string) {
	o.Comment = &v
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiLicenseOverrideDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseOverrideDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiLicenseOverrideDTO) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiLicenseOverrideDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiLicenseOverrideDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseOverrideDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiLicenseOverrideDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiLicenseOverrideDTO) SetId(v string) {
	o.Id = &v
}

// GetLicenseIds returns the LicenseIds field value if set, zero value otherwise.
func (o *ApiLicenseOverrideDTO) GetLicenseIds() []string {
	if o == nil || IsNil(o.LicenseIds) {
		var ret []string
		return ret
	}
	return o.LicenseIds
}

// GetLicenseIdsOk returns a tuple with the LicenseIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseOverrideDTO) GetLicenseIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.LicenseIds) {
		return nil, false
	}
	return o.LicenseIds, true
}

// HasLicenseIds returns a boolean if a field has been set.
func (o *ApiLicenseOverrideDTO) HasLicenseIds() bool {
	if o != nil && !IsNil(o.LicenseIds) {
		return true
	}

	return false
}

// SetLicenseIds gets a reference to the given []string and assigns it to the LicenseIds field.
func (o *ApiLicenseOverrideDTO) SetLicenseIds(v []string) {
	o.LicenseIds = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ApiLicenseOverrideDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseOverrideDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ApiLicenseOverrideDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ApiLicenseOverrideDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiLicenseOverrideDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseOverrideDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiLicenseOverrideDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiLicenseOverrideDTO) SetStatus(v string) {
	o.Status = &v
}

func (o ApiLicenseOverrideDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicenseOverrideDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LicenseIds) {
		toSerialize["licenseIds"] = o.LicenseIds
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableApiLicenseOverrideDTO struct {
	value *ApiLicenseOverrideDTO
	isSet bool
}

func (v NullableApiLicenseOverrideDTO) Get() *ApiLicenseOverrideDTO {
	return v.value
}

func (v *NullableApiLicenseOverrideDTO) Set(val *ApiLicenseOverrideDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicenseOverrideDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicenseOverrideDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicenseOverrideDTO(val *ApiLicenseOverrideDTO) *NullableApiLicenseOverrideDTO {
	return &NullableApiLicenseOverrideDTO{value: val, isSet: true}
}

func (v NullableApiLicenseOverrideDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicenseOverrideDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


