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

// checks if the GroupingByDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupingByDTO{}

// GroupingByDTO struct for GroupingByDTO
type GroupingByDTO struct {
	AdditionalInfo *string `json:"additionalInfo,omitempty"`
	GroupBy *string `json:"groupBy,omitempty"`
	GroupIdentifier *string `json:"groupIdentifier,omitempty"`
	SearchResultItemDTOS []SearchResultItemDTO `json:"searchResultItemDTOS,omitempty"`
}

// NewGroupingByDTO instantiates a new GroupingByDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupingByDTO() *GroupingByDTO {
	this := GroupingByDTO{}
	return &this
}

// NewGroupingByDTOWithDefaults instantiates a new GroupingByDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupingByDTOWithDefaults() *GroupingByDTO {
	this := GroupingByDTO{}
	return &this
}

// GetAdditionalInfo returns the AdditionalInfo field value if set, zero value otherwise.
func (o *GroupingByDTO) GetAdditionalInfo() string {
	if o == nil || IsNil(o.AdditionalInfo) {
		var ret string
		return ret
	}
	return *o.AdditionalInfo
}

// GetAdditionalInfoOk returns a tuple with the AdditionalInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupingByDTO) GetAdditionalInfoOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalInfo) {
		return nil, false
	}
	return o.AdditionalInfo, true
}

// HasAdditionalInfo returns a boolean if a field has been set.
func (o *GroupingByDTO) HasAdditionalInfo() bool {
	if o != nil && !IsNil(o.AdditionalInfo) {
		return true
	}

	return false
}

// SetAdditionalInfo gets a reference to the given string and assigns it to the AdditionalInfo field.
func (o *GroupingByDTO) SetAdditionalInfo(v string) {
	o.AdditionalInfo = &v
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *GroupingByDTO) GetGroupBy() string {
	if o == nil || IsNil(o.GroupBy) {
		var ret string
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupingByDTO) GetGroupByOk() (*string, bool) {
	if o == nil || IsNil(o.GroupBy) {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *GroupingByDTO) HasGroupBy() bool {
	if o != nil && !IsNil(o.GroupBy) {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given string and assigns it to the GroupBy field.
func (o *GroupingByDTO) SetGroupBy(v string) {
	o.GroupBy = &v
}

// GetGroupIdentifier returns the GroupIdentifier field value if set, zero value otherwise.
func (o *GroupingByDTO) GetGroupIdentifier() string {
	if o == nil || IsNil(o.GroupIdentifier) {
		var ret string
		return ret
	}
	return *o.GroupIdentifier
}

// GetGroupIdentifierOk returns a tuple with the GroupIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupingByDTO) GetGroupIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.GroupIdentifier) {
		return nil, false
	}
	return o.GroupIdentifier, true
}

// HasGroupIdentifier returns a boolean if a field has been set.
func (o *GroupingByDTO) HasGroupIdentifier() bool {
	if o != nil && !IsNil(o.GroupIdentifier) {
		return true
	}

	return false
}

// SetGroupIdentifier gets a reference to the given string and assigns it to the GroupIdentifier field.
func (o *GroupingByDTO) SetGroupIdentifier(v string) {
	o.GroupIdentifier = &v
}

// GetSearchResultItemDTOS returns the SearchResultItemDTOS field value if set, zero value otherwise.
func (o *GroupingByDTO) GetSearchResultItemDTOS() []SearchResultItemDTO {
	if o == nil || IsNil(o.SearchResultItemDTOS) {
		var ret []SearchResultItemDTO
		return ret
	}
	return o.SearchResultItemDTOS
}

// GetSearchResultItemDTOSOk returns a tuple with the SearchResultItemDTOS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupingByDTO) GetSearchResultItemDTOSOk() ([]SearchResultItemDTO, bool) {
	if o == nil || IsNil(o.SearchResultItemDTOS) {
		return nil, false
	}
	return o.SearchResultItemDTOS, true
}

// HasSearchResultItemDTOS returns a boolean if a field has been set.
func (o *GroupingByDTO) HasSearchResultItemDTOS() bool {
	if o != nil && !IsNil(o.SearchResultItemDTOS) {
		return true
	}

	return false
}

// SetSearchResultItemDTOS gets a reference to the given []SearchResultItemDTO and assigns it to the SearchResultItemDTOS field.
func (o *GroupingByDTO) SetSearchResultItemDTOS(v []SearchResultItemDTO) {
	o.SearchResultItemDTOS = v
}

func (o GroupingByDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupingByDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalInfo) {
		toSerialize["additionalInfo"] = o.AdditionalInfo
	}
	if !IsNil(o.GroupBy) {
		toSerialize["groupBy"] = o.GroupBy
	}
	if !IsNil(o.GroupIdentifier) {
		toSerialize["groupIdentifier"] = o.GroupIdentifier
	}
	if !IsNil(o.SearchResultItemDTOS) {
		toSerialize["searchResultItemDTOS"] = o.SearchResultItemDTOS
	}
	return toSerialize, nil
}

type NullableGroupingByDTO struct {
	value *GroupingByDTO
	isSet bool
}

func (v NullableGroupingByDTO) Get() *GroupingByDTO {
	return v.value
}

func (v *NullableGroupingByDTO) Set(val *GroupingByDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupingByDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupingByDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupingByDTO(val *GroupingByDTO) *NullableGroupingByDTO {
	return &NullableGroupingByDTO{value: val, isSet: true}
}

func (v NullableGroupingByDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupingByDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


