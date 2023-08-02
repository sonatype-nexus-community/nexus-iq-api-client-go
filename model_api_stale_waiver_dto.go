/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 165
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiStaleWaiverDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiStaleWaiverDTO{}

// ApiStaleWaiverDTO struct for ApiStaleWaiverDTO
type ApiStaleWaiverDTO struct {
	Comment *string `json:"comment,omitempty"`
	ConstraintFacts []ApiConstraintFactDTO `json:"constraintFacts,omitempty"`
	CreateTime *time.Time `json:"createTime,omitempty"`
	CreatorId *string `json:"creatorId,omitempty"`
	CreatorName *string `json:"creatorName,omitempty"`
	ExpiryTime *time.Time `json:"expiryTime,omitempty"`
	PolicyId *string `json:"policyId,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
	ScopeOwnerId *string `json:"scopeOwnerId,omitempty"`
	ScopeOwnerName *string `json:"scopeOwnerName,omitempty"`
	ScopeOwnerType *string `json:"scopeOwnerType,omitempty"`
	StaleEvaluations *ApiStaleEvaluationsDTO `json:"staleEvaluations,omitempty"`
	WaiverId *string `json:"waiverId,omitempty"`
}

// NewApiStaleWaiverDTO instantiates a new ApiStaleWaiverDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStaleWaiverDTO() *ApiStaleWaiverDTO {
	this := ApiStaleWaiverDTO{}
	return &this
}

// NewApiStaleWaiverDTOWithDefaults instantiates a new ApiStaleWaiverDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStaleWaiverDTOWithDefaults() *ApiStaleWaiverDTO {
	this := ApiStaleWaiverDTO{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApiStaleWaiverDTO) SetComment(v string) {
	o.Comment = &v
}

// GetConstraintFacts returns the ConstraintFacts field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetConstraintFacts() []ApiConstraintFactDTO {
	if o == nil || IsNil(o.ConstraintFacts) {
		var ret []ApiConstraintFactDTO
		return ret
	}
	return o.ConstraintFacts
}

// GetConstraintFactsOk returns a tuple with the ConstraintFacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetConstraintFactsOk() ([]ApiConstraintFactDTO, bool) {
	if o == nil || IsNil(o.ConstraintFacts) {
		return nil, false
	}
	return o.ConstraintFacts, true
}

// HasConstraintFacts returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasConstraintFacts() bool {
	if o != nil && !IsNil(o.ConstraintFacts) {
		return true
	}

	return false
}

// SetConstraintFacts gets a reference to the given []ApiConstraintFactDTO and assigns it to the ConstraintFacts field.
func (o *ApiStaleWaiverDTO) SetConstraintFacts(v []ApiConstraintFactDTO) {
	o.ConstraintFacts = v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *ApiStaleWaiverDTO) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *ApiStaleWaiverDTO) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetCreatorName returns the CreatorName field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetCreatorName() string {
	if o == nil || IsNil(o.CreatorName) {
		var ret string
		return ret
	}
	return *o.CreatorName
}

// GetCreatorNameOk returns a tuple with the CreatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetCreatorNameOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorName) {
		return nil, false
	}
	return o.CreatorName, true
}

// HasCreatorName returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasCreatorName() bool {
	if o != nil && !IsNil(o.CreatorName) {
		return true
	}

	return false
}

// SetCreatorName gets a reference to the given string and assigns it to the CreatorName field.
func (o *ApiStaleWaiverDTO) SetCreatorName(v string) {
	o.CreatorName = &v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetExpiryTime() time.Time {
	if o == nil || IsNil(o.ExpiryTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetExpiryTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiryTime) {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasExpiryTime() bool {
	if o != nil && !IsNil(o.ExpiryTime) {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given time.Time and assigns it to the ExpiryTime field.
func (o *ApiStaleWaiverDTO) SetExpiryTime(v time.Time) {
	o.ExpiryTime = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *ApiStaleWaiverDTO) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName) {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetPolicyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyName) {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasPolicyName() bool {
	if o != nil && !IsNil(o.PolicyName) {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *ApiStaleWaiverDTO) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetScopeOwnerId returns the ScopeOwnerId field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetScopeOwnerId() string {
	if o == nil || IsNil(o.ScopeOwnerId) {
		var ret string
		return ret
	}
	return *o.ScopeOwnerId
}

// GetScopeOwnerIdOk returns a tuple with the ScopeOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetScopeOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeOwnerId) {
		return nil, false
	}
	return o.ScopeOwnerId, true
}

// HasScopeOwnerId returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasScopeOwnerId() bool {
	if o != nil && !IsNil(o.ScopeOwnerId) {
		return true
	}

	return false
}

// SetScopeOwnerId gets a reference to the given string and assigns it to the ScopeOwnerId field.
func (o *ApiStaleWaiverDTO) SetScopeOwnerId(v string) {
	o.ScopeOwnerId = &v
}

// GetScopeOwnerName returns the ScopeOwnerName field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetScopeOwnerName() string {
	if o == nil || IsNil(o.ScopeOwnerName) {
		var ret string
		return ret
	}
	return *o.ScopeOwnerName
}

// GetScopeOwnerNameOk returns a tuple with the ScopeOwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetScopeOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeOwnerName) {
		return nil, false
	}
	return o.ScopeOwnerName, true
}

// HasScopeOwnerName returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasScopeOwnerName() bool {
	if o != nil && !IsNil(o.ScopeOwnerName) {
		return true
	}

	return false
}

// SetScopeOwnerName gets a reference to the given string and assigns it to the ScopeOwnerName field.
func (o *ApiStaleWaiverDTO) SetScopeOwnerName(v string) {
	o.ScopeOwnerName = &v
}

// GetScopeOwnerType returns the ScopeOwnerType field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetScopeOwnerType() string {
	if o == nil || IsNil(o.ScopeOwnerType) {
		var ret string
		return ret
	}
	return *o.ScopeOwnerType
}

// GetScopeOwnerTypeOk returns a tuple with the ScopeOwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetScopeOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeOwnerType) {
		return nil, false
	}
	return o.ScopeOwnerType, true
}

// HasScopeOwnerType returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasScopeOwnerType() bool {
	if o != nil && !IsNil(o.ScopeOwnerType) {
		return true
	}

	return false
}

// SetScopeOwnerType gets a reference to the given string and assigns it to the ScopeOwnerType field.
func (o *ApiStaleWaiverDTO) SetScopeOwnerType(v string) {
	o.ScopeOwnerType = &v
}

// GetStaleEvaluations returns the StaleEvaluations field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetStaleEvaluations() ApiStaleEvaluationsDTO {
	if o == nil || IsNil(o.StaleEvaluations) {
		var ret ApiStaleEvaluationsDTO
		return ret
	}
	return *o.StaleEvaluations
}

// GetStaleEvaluationsOk returns a tuple with the StaleEvaluations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetStaleEvaluationsOk() (*ApiStaleEvaluationsDTO, bool) {
	if o == nil || IsNil(o.StaleEvaluations) {
		return nil, false
	}
	return o.StaleEvaluations, true
}

// HasStaleEvaluations returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasStaleEvaluations() bool {
	if o != nil && !IsNil(o.StaleEvaluations) {
		return true
	}

	return false
}

// SetStaleEvaluations gets a reference to the given ApiStaleEvaluationsDTO and assigns it to the StaleEvaluations field.
func (o *ApiStaleWaiverDTO) SetStaleEvaluations(v ApiStaleEvaluationsDTO) {
	o.StaleEvaluations = &v
}

// GetWaiverId returns the WaiverId field value if set, zero value otherwise.
func (o *ApiStaleWaiverDTO) GetWaiverId() string {
	if o == nil || IsNil(o.WaiverId) {
		var ret string
		return ret
	}
	return *o.WaiverId
}

// GetWaiverIdOk returns a tuple with the WaiverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStaleWaiverDTO) GetWaiverIdOk() (*string, bool) {
	if o == nil || IsNil(o.WaiverId) {
		return nil, false
	}
	return o.WaiverId, true
}

// HasWaiverId returns a boolean if a field has been set.
func (o *ApiStaleWaiverDTO) HasWaiverId() bool {
	if o != nil && !IsNil(o.WaiverId) {
		return true
	}

	return false
}

// SetWaiverId gets a reference to the given string and assigns it to the WaiverId field.
func (o *ApiStaleWaiverDTO) SetWaiverId(v string) {
	o.WaiverId = &v
}

func (o ApiStaleWaiverDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiStaleWaiverDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ConstraintFacts) {
		toSerialize["constraintFacts"] = o.ConstraintFacts
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
	if !IsNil(o.ExpiryTime) {
		toSerialize["expiryTime"] = o.ExpiryTime
	}
	if !IsNil(o.PolicyId) {
		toSerialize["policyId"] = o.PolicyId
	}
	if !IsNil(o.PolicyName) {
		toSerialize["policyName"] = o.PolicyName
	}
	if !IsNil(o.ScopeOwnerId) {
		toSerialize["scopeOwnerId"] = o.ScopeOwnerId
	}
	if !IsNil(o.ScopeOwnerName) {
		toSerialize["scopeOwnerName"] = o.ScopeOwnerName
	}
	if !IsNil(o.ScopeOwnerType) {
		toSerialize["scopeOwnerType"] = o.ScopeOwnerType
	}
	if !IsNil(o.StaleEvaluations) {
		toSerialize["staleEvaluations"] = o.StaleEvaluations
	}
	if !IsNil(o.WaiverId) {
		toSerialize["waiverId"] = o.WaiverId
	}
	return toSerialize, nil
}

type NullableApiStaleWaiverDTO struct {
	value *ApiStaleWaiverDTO
	isSet bool
}

func (v NullableApiStaleWaiverDTO) Get() *ApiStaleWaiverDTO {
	return v.value
}

func (v *NullableApiStaleWaiverDTO) Set(val *ApiStaleWaiverDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStaleWaiverDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStaleWaiverDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStaleWaiverDTO(val *ApiStaleWaiverDTO) *NullableApiStaleWaiverDTO {
	return &NullableApiStaleWaiverDTO{value: val, isSet: true}
}

func (v NullableApiStaleWaiverDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStaleWaiverDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

