/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.184.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the TagsByOwnerDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagsByOwnerDTO{}

// TagsByOwnerDTO struct for TagsByOwnerDTO
type TagsByOwnerDTO struct {
	ApplicationCategories []ApiApplicationCategoryDTO `json:"applicationCategories,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	OwnerName *string `json:"ownerName,omitempty"`
	OwnerType *string `json:"ownerType,omitempty"`
}

// NewTagsByOwnerDTO instantiates a new TagsByOwnerDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagsByOwnerDTO() *TagsByOwnerDTO {
	this := TagsByOwnerDTO{}
	return &this
}

// NewTagsByOwnerDTOWithDefaults instantiates a new TagsByOwnerDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagsByOwnerDTOWithDefaults() *TagsByOwnerDTO {
	this := TagsByOwnerDTO{}
	return &this
}

// GetApplicationCategories returns the ApplicationCategories field value if set, zero value otherwise.
func (o *TagsByOwnerDTO) GetApplicationCategories() []ApiApplicationCategoryDTO {
	if o == nil || IsNil(o.ApplicationCategories) {
		var ret []ApiApplicationCategoryDTO
		return ret
	}
	return o.ApplicationCategories
}

// GetApplicationCategoriesOk returns a tuple with the ApplicationCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsByOwnerDTO) GetApplicationCategoriesOk() ([]ApiApplicationCategoryDTO, bool) {
	if o == nil || IsNil(o.ApplicationCategories) {
		return nil, false
	}
	return o.ApplicationCategories, true
}

// HasApplicationCategories returns a boolean if a field has been set.
func (o *TagsByOwnerDTO) HasApplicationCategories() bool {
	if o != nil && !IsNil(o.ApplicationCategories) {
		return true
	}

	return false
}

// SetApplicationCategories gets a reference to the given []ApiApplicationCategoryDTO and assigns it to the ApplicationCategories field.
func (o *TagsByOwnerDTO) SetApplicationCategories(v []ApiApplicationCategoryDTO) {
	o.ApplicationCategories = v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *TagsByOwnerDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsByOwnerDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *TagsByOwnerDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *TagsByOwnerDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *TagsByOwnerDTO) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsByOwnerDTO) GetOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *TagsByOwnerDTO) HasOwnerName() bool {
	if o != nil && !IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *TagsByOwnerDTO) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *TagsByOwnerDTO) GetOwnerType() string {
	if o == nil || IsNil(o.OwnerType) {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsByOwnerDTO) GetOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerType) {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *TagsByOwnerDTO) HasOwnerType() bool {
	if o != nil && !IsNil(o.OwnerType) {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *TagsByOwnerDTO) SetOwnerType(v string) {
	o.OwnerType = &v
}

func (o TagsByOwnerDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagsByOwnerDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationCategories) {
		toSerialize["applicationCategories"] = o.ApplicationCategories
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !IsNil(o.OwnerType) {
		toSerialize["ownerType"] = o.OwnerType
	}
	return toSerialize, nil
}

type NullableTagsByOwnerDTO struct {
	value *TagsByOwnerDTO
	isSet bool
}

func (v NullableTagsByOwnerDTO) Get() *TagsByOwnerDTO {
	return v.value
}

func (v *NullableTagsByOwnerDTO) Set(val *TagsByOwnerDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTagsByOwnerDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTagsByOwnerDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagsByOwnerDTO(val *TagsByOwnerDTO) *NullableTagsByOwnerDTO {
	return &NullableTagsByOwnerDTO{value: val, isSet: true}
}

func (v NullableTagsByOwnerDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagsByOwnerDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


