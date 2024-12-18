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

// checks if the ApiAutoPolicyWaiverDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAutoPolicyWaiverDTO{}

// ApiAutoPolicyWaiverDTO struct for ApiAutoPolicyWaiverDTO
type ApiAutoPolicyWaiverDTO struct {
	AutoPolicyWaiverId *string `json:"autoPolicyWaiverId,omitempty"`
	CreateTime *time.Time `json:"createTime,omitempty"`
	CreatorId *string `json:"creatorId,omitempty"`
	CreatorName *string `json:"creatorName,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	PathForward *bool `json:"pathForward,omitempty"`
	Reachable *bool `json:"reachable,omitempty"`
	ThreatLevel *int32 `json:"threatLevel,omitempty"`
}

// NewApiAutoPolicyWaiverDTO instantiates a new ApiAutoPolicyWaiverDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAutoPolicyWaiverDTO() *ApiAutoPolicyWaiverDTO {
	this := ApiAutoPolicyWaiverDTO{}
	return &this
}

// NewApiAutoPolicyWaiverDTOWithDefaults instantiates a new ApiAutoPolicyWaiverDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAutoPolicyWaiverDTOWithDefaults() *ApiAutoPolicyWaiverDTO {
	this := ApiAutoPolicyWaiverDTO{}
	return &this
}

// GetAutoPolicyWaiverId returns the AutoPolicyWaiverId field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetAutoPolicyWaiverId() string {
	if o == nil || IsNil(o.AutoPolicyWaiverId) {
		var ret string
		return ret
	}
	return *o.AutoPolicyWaiverId
}

// GetAutoPolicyWaiverIdOk returns a tuple with the AutoPolicyWaiverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetAutoPolicyWaiverIdOk() (*string, bool) {
	if o == nil || IsNil(o.AutoPolicyWaiverId) {
		return nil, false
	}
	return o.AutoPolicyWaiverId, true
}

// HasAutoPolicyWaiverId returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasAutoPolicyWaiverId() bool {
	if o != nil && !IsNil(o.AutoPolicyWaiverId) {
		return true
	}

	return false
}

// SetAutoPolicyWaiverId gets a reference to the given string and assigns it to the AutoPolicyWaiverId field.
func (o *ApiAutoPolicyWaiverDTO) SetAutoPolicyWaiverId(v string) {
	o.AutoPolicyWaiverId = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *ApiAutoPolicyWaiverDTO) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *ApiAutoPolicyWaiverDTO) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetCreatorName returns the CreatorName field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetCreatorName() string {
	if o == nil || IsNil(o.CreatorName) {
		var ret string
		return ret
	}
	return *o.CreatorName
}

// GetCreatorNameOk returns a tuple with the CreatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetCreatorNameOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorName) {
		return nil, false
	}
	return o.CreatorName, true
}

// HasCreatorName returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasCreatorName() bool {
	if o != nil && !IsNil(o.CreatorName) {
		return true
	}

	return false
}

// SetCreatorName gets a reference to the given string and assigns it to the CreatorName field.
func (o *ApiAutoPolicyWaiverDTO) SetCreatorName(v string) {
	o.CreatorName = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ApiAutoPolicyWaiverDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPathForward returns the PathForward field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetPathForward() bool {
	if o == nil || IsNil(o.PathForward) {
		var ret bool
		return ret
	}
	return *o.PathForward
}

// GetPathForwardOk returns a tuple with the PathForward field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetPathForwardOk() (*bool, bool) {
	if o == nil || IsNil(o.PathForward) {
		return nil, false
	}
	return o.PathForward, true
}

// HasPathForward returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasPathForward() bool {
	if o != nil && !IsNil(o.PathForward) {
		return true
	}

	return false
}

// SetPathForward gets a reference to the given bool and assigns it to the PathForward field.
func (o *ApiAutoPolicyWaiverDTO) SetPathForward(v bool) {
	o.PathForward = &v
}

// GetReachable returns the Reachable field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetReachable() bool {
	if o == nil || IsNil(o.Reachable) {
		var ret bool
		return ret
	}
	return *o.Reachable
}

// GetReachableOk returns a tuple with the Reachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetReachableOk() (*bool, bool) {
	if o == nil || IsNil(o.Reachable) {
		return nil, false
	}
	return o.Reachable, true
}

// HasReachable returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasReachable() bool {
	if o != nil && !IsNil(o.Reachable) {
		return true
	}

	return false
}

// SetReachable gets a reference to the given bool and assigns it to the Reachable field.
func (o *ApiAutoPolicyWaiverDTO) SetReachable(v bool) {
	o.Reachable = &v
}

// GetThreatLevel returns the ThreatLevel field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetThreatLevel() int32 {
	if o == nil || IsNil(o.ThreatLevel) {
		var ret int32
		return ret
	}
	return *o.ThreatLevel
}

// GetThreatLevelOk returns a tuple with the ThreatLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetThreatLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.ThreatLevel) {
		return nil, false
	}
	return o.ThreatLevel, true
}

// HasThreatLevel returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasThreatLevel() bool {
	if o != nil && !IsNil(o.ThreatLevel) {
		return true
	}

	return false
}

// SetThreatLevel gets a reference to the given int32 and assigns it to the ThreatLevel field.
func (o *ApiAutoPolicyWaiverDTO) SetThreatLevel(v int32) {
	o.ThreatLevel = &v
}

func (o ApiAutoPolicyWaiverDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAutoPolicyWaiverDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoPolicyWaiverId) {
		toSerialize["autoPolicyWaiverId"] = o.AutoPolicyWaiverId
	}
	if !IsNil(o.CreateTime) {
		toSerialize["createTime"] = o.CreateTime
	}
	if !IsNil(o.CreatorId) {
		toSerialize["creatorId"] = o.CreatorId
	}
	if !IsNil(o.CreatorName) {
		toSerialize["creatorName"] = o.CreatorName
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.PathForward) {
		toSerialize["pathForward"] = o.PathForward
	}
	if !IsNil(o.Reachable) {
		toSerialize["reachable"] = o.Reachable
	}
	if !IsNil(o.ThreatLevel) {
		toSerialize["threatLevel"] = o.ThreatLevel
	}
	return toSerialize, nil
}

type NullableApiAutoPolicyWaiverDTO struct {
	value *ApiAutoPolicyWaiverDTO
	isSet bool
}

func (v NullableApiAutoPolicyWaiverDTO) Get() *ApiAutoPolicyWaiverDTO {
	return v.value
}

func (v *NullableApiAutoPolicyWaiverDTO) Set(val *ApiAutoPolicyWaiverDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAutoPolicyWaiverDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAutoPolicyWaiverDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAutoPolicyWaiverDTO(val *ApiAutoPolicyWaiverDTO) *NullableApiAutoPolicyWaiverDTO {
	return &NullableApiAutoPolicyWaiverDTO{value: val, isSet: true}
}

func (v NullableApiAutoPolicyWaiverDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAutoPolicyWaiverDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


