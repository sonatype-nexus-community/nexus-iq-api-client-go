/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.193.0-01
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
	OwnerName *string `json:"ownerName,omitempty"`
	OwnerType *string `json:"ownerType,omitempty"`
	PathForward *bool `json:"pathForward,omitempty"`
	PublicId *string `json:"publicId,omitempty"`
	Reachability *bool `json:"reachability,omitempty"`
	ScopesOperatorAny *bool `json:"scopesOperatorAny,omitempty"`
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

// GetOwnerName returns the OwnerName field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetOwnerName() string {
	if o == nil || IsNil(o.OwnerName) {
		var ret string
		return ret
	}
	return *o.OwnerName
}

// GetOwnerNameOk returns a tuple with the OwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerName) {
		return nil, false
	}
	return o.OwnerName, true
}

// HasOwnerName returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasOwnerName() bool {
	if o != nil && !IsNil(o.OwnerName) {
		return true
	}

	return false
}

// SetOwnerName gets a reference to the given string and assigns it to the OwnerName field.
func (o *ApiAutoPolicyWaiverDTO) SetOwnerName(v string) {
	o.OwnerName = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetOwnerType() string {
	if o == nil || IsNil(o.OwnerType) {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerType) {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasOwnerType() bool {
	if o != nil && !IsNil(o.OwnerType) {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *ApiAutoPolicyWaiverDTO) SetOwnerType(v string) {
	o.OwnerType = &v
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

// GetPublicId returns the PublicId field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetPublicId() string {
	if o == nil || IsNil(o.PublicId) {
		var ret string
		return ret
	}
	return *o.PublicId
}

// GetPublicIdOk returns a tuple with the PublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetPublicIdOk() (*string, bool) {
	if o == nil || IsNil(o.PublicId) {
		return nil, false
	}
	return o.PublicId, true
}

// HasPublicId returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasPublicId() bool {
	if o != nil && !IsNil(o.PublicId) {
		return true
	}

	return false
}

// SetPublicId gets a reference to the given string and assigns it to the PublicId field.
func (o *ApiAutoPolicyWaiverDTO) SetPublicId(v string) {
	o.PublicId = &v
}

// GetReachability returns the Reachability field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetReachability() bool {
	if o == nil || IsNil(o.Reachability) {
		var ret bool
		return ret
	}
	return *o.Reachability
}

// GetReachabilityOk returns a tuple with the Reachability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetReachabilityOk() (*bool, bool) {
	if o == nil || IsNil(o.Reachability) {
		return nil, false
	}
	return o.Reachability, true
}

// HasReachability returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasReachability() bool {
	if o != nil && !IsNil(o.Reachability) {
		return true
	}

	return false
}

// SetReachability gets a reference to the given bool and assigns it to the Reachability field.
func (o *ApiAutoPolicyWaiverDTO) SetReachability(v bool) {
	o.Reachability = &v
}

// GetScopesOperatorAny returns the ScopesOperatorAny field value if set, zero value otherwise.
func (o *ApiAutoPolicyWaiverDTO) GetScopesOperatorAny() bool {
	if o == nil || IsNil(o.ScopesOperatorAny) {
		var ret bool
		return ret
	}
	return *o.ScopesOperatorAny
}

// GetScopesOperatorAnyOk returns a tuple with the ScopesOperatorAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAutoPolicyWaiverDTO) GetScopesOperatorAnyOk() (*bool, bool) {
	if o == nil || IsNil(o.ScopesOperatorAny) {
		return nil, false
	}
	return o.ScopesOperatorAny, true
}

// HasScopesOperatorAny returns a boolean if a field has been set.
func (o *ApiAutoPolicyWaiverDTO) HasScopesOperatorAny() bool {
	if o != nil && !IsNil(o.ScopesOperatorAny) {
		return true
	}

	return false
}

// SetScopesOperatorAny gets a reference to the given bool and assigns it to the ScopesOperatorAny field.
func (o *ApiAutoPolicyWaiverDTO) SetScopesOperatorAny(v bool) {
	o.ScopesOperatorAny = &v
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
	if !IsNil(o.OwnerName) {
		toSerialize["ownerName"] = o.OwnerName
	}
	if !IsNil(o.OwnerType) {
		toSerialize["ownerType"] = o.OwnerType
	}
	if !IsNil(o.PathForward) {
		toSerialize["pathForward"] = o.PathForward
	}
	if !IsNil(o.PublicId) {
		toSerialize["publicId"] = o.PublicId
	}
	if !IsNil(o.Reachability) {
		toSerialize["reachability"] = o.Reachability
	}
	if !IsNil(o.ScopesOperatorAny) {
		toSerialize["scopesOperatorAny"] = o.ScopesOperatorAny
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


