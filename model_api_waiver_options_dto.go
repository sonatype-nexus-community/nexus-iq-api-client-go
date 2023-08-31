/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.166.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiWaiverOptionsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiWaiverOptionsDTO{}

// ApiWaiverOptionsDTO struct for ApiWaiverOptionsDTO
type ApiWaiverOptionsDTO struct {
	ApplyToAllComponents *bool `json:"applyToAllComponents,omitempty"`
	Comment *string `json:"comment,omitempty"`
	ExpiryTime *time.Time `json:"expiryTime,omitempty"`
	MatcherStrategy *string `json:"matcherStrategy,omitempty"`
}

// NewApiWaiverOptionsDTO instantiates a new ApiWaiverOptionsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiWaiverOptionsDTO() *ApiWaiverOptionsDTO {
	this := ApiWaiverOptionsDTO{}
	return &this
}

// NewApiWaiverOptionsDTOWithDefaults instantiates a new ApiWaiverOptionsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiWaiverOptionsDTOWithDefaults() *ApiWaiverOptionsDTO {
	this := ApiWaiverOptionsDTO{}
	return &this
}

// GetApplyToAllComponents returns the ApplyToAllComponents field value if set, zero value otherwise.
func (o *ApiWaiverOptionsDTO) GetApplyToAllComponents() bool {
	if o == nil || IsNil(o.ApplyToAllComponents) {
		var ret bool
		return ret
	}
	return *o.ApplyToAllComponents
}

// GetApplyToAllComponentsOk returns a tuple with the ApplyToAllComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWaiverOptionsDTO) GetApplyToAllComponentsOk() (*bool, bool) {
	if o == nil || IsNil(o.ApplyToAllComponents) {
		return nil, false
	}
	return o.ApplyToAllComponents, true
}

// HasApplyToAllComponents returns a boolean if a field has been set.
func (o *ApiWaiverOptionsDTO) HasApplyToAllComponents() bool {
	if o != nil && !IsNil(o.ApplyToAllComponents) {
		return true
	}

	return false
}

// SetApplyToAllComponents gets a reference to the given bool and assigns it to the ApplyToAllComponents field.
func (o *ApiWaiverOptionsDTO) SetApplyToAllComponents(v bool) {
	o.ApplyToAllComponents = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApiWaiverOptionsDTO) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWaiverOptionsDTO) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApiWaiverOptionsDTO) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApiWaiverOptionsDTO) SetComment(v string) {
	o.Comment = &v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *ApiWaiverOptionsDTO) GetExpiryTime() time.Time {
	if o == nil || IsNil(o.ExpiryTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWaiverOptionsDTO) GetExpiryTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiryTime) {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *ApiWaiverOptionsDTO) HasExpiryTime() bool {
	if o != nil && !IsNil(o.ExpiryTime) {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given time.Time and assigns it to the ExpiryTime field.
func (o *ApiWaiverOptionsDTO) SetExpiryTime(v time.Time) {
	o.ExpiryTime = &v
}

// GetMatcherStrategy returns the MatcherStrategy field value if set, zero value otherwise.
func (o *ApiWaiverOptionsDTO) GetMatcherStrategy() string {
	if o == nil || IsNil(o.MatcherStrategy) {
		var ret string
		return ret
	}
	return *o.MatcherStrategy
}

// GetMatcherStrategyOk returns a tuple with the MatcherStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWaiverOptionsDTO) GetMatcherStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.MatcherStrategy) {
		return nil, false
	}
	return o.MatcherStrategy, true
}

// HasMatcherStrategy returns a boolean if a field has been set.
func (o *ApiWaiverOptionsDTO) HasMatcherStrategy() bool {
	if o != nil && !IsNil(o.MatcherStrategy) {
		return true
	}

	return false
}

// SetMatcherStrategy gets a reference to the given string and assigns it to the MatcherStrategy field.
func (o *ApiWaiverOptionsDTO) SetMatcherStrategy(v string) {
	o.MatcherStrategy = &v
}

func (o ApiWaiverOptionsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiWaiverOptionsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplyToAllComponents) {
		toSerialize["applyToAllComponents"] = o.ApplyToAllComponents
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ExpiryTime) {
		toSerialize["expiryTime"] = o.ExpiryTime
	}
	if !IsNil(o.MatcherStrategy) {
		toSerialize["matcherStrategy"] = o.MatcherStrategy
	}
	return toSerialize, nil
}

type NullableApiWaiverOptionsDTO struct {
	value *ApiWaiverOptionsDTO
	isSet bool
}

func (v NullableApiWaiverOptionsDTO) Get() *ApiWaiverOptionsDTO {
	return v.value
}

func (v *NullableApiWaiverOptionsDTO) Set(val *ApiWaiverOptionsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiWaiverOptionsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiWaiverOptionsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiWaiverOptionsDTO(val *ApiWaiverOptionsDTO) *NullableApiWaiverOptionsDTO {
	return &NullableApiWaiverOptionsDTO{value: val, isSet: true}
}

func (v NullableApiWaiverOptionsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiWaiverOptionsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


