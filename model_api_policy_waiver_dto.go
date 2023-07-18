/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 164
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the ApiPolicyWaiverDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPolicyWaiverDTO{}

// ApiPolicyWaiverDTO struct for ApiPolicyWaiverDTO
type ApiPolicyWaiverDTO struct {
	AssociatedPackageUrl *string `json:"associatedPackageUrl,omitempty"`
	Comment *string `json:"comment,omitempty"`
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	ComponentName *string `json:"componentName,omitempty"`
	ComponentUpgradeAvailable *bool `json:"componentUpgradeAvailable,omitempty"`
	ConstraintFacts []ConstraintFact `json:"constraintFacts,omitempty"`
	ConstraintFactsJson *string `json:"constraintFactsJson,omitempty"`
	CreateTime *time.Time `json:"createTime,omitempty"`
	CreatorId *string `json:"creatorId,omitempty"`
	CreatorName *string `json:"creatorName,omitempty"`
	DisplayName *ComponentDisplayName `json:"displayName,omitempty"`
	ExpiryTime *time.Time `json:"expiryTime,omitempty"`
	Hash *string `json:"hash,omitempty"`
	IsObsolete *bool `json:"isObsolete,omitempty"`
	MatcherStrategy *string `json:"matcherStrategy,omitempty"`
	PolicyId *string `json:"policyId,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
	PolicyViolationId *string `json:"policyViolationId,omitempty"`
	PolicyWaiverId *string `json:"policyWaiverId,omitempty"`
	ScopeOwnerId *string `json:"scopeOwnerId,omitempty"`
	ScopeOwnerName *string `json:"scopeOwnerName,omitempty"`
	ScopeOwnerType *string `json:"scopeOwnerType,omitempty"`
	ThreatLevel *int32 `json:"threatLevel,omitempty"`
	VulnerabilityId *string `json:"vulnerabilityId,omitempty"`
}

// NewApiPolicyWaiverDTO instantiates a new ApiPolicyWaiverDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPolicyWaiverDTO() *ApiPolicyWaiverDTO {
	this := ApiPolicyWaiverDTO{}
	return &this
}

// NewApiPolicyWaiverDTOWithDefaults instantiates a new ApiPolicyWaiverDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPolicyWaiverDTOWithDefaults() *ApiPolicyWaiverDTO {
	this := ApiPolicyWaiverDTO{}
	return &this
}

// GetAssociatedPackageUrl returns the AssociatedPackageUrl field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetAssociatedPackageUrl() string {
	if o == nil || IsNil(o.AssociatedPackageUrl) {
		var ret string
		return ret
	}
	return *o.AssociatedPackageUrl
}

// GetAssociatedPackageUrlOk returns a tuple with the AssociatedPackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetAssociatedPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AssociatedPackageUrl) {
		return nil, false
	}
	return o.AssociatedPackageUrl, true
}

// HasAssociatedPackageUrl returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasAssociatedPackageUrl() bool {
	if o != nil && !IsNil(o.AssociatedPackageUrl) {
		return true
	}

	return false
}

// SetAssociatedPackageUrl gets a reference to the given string and assigns it to the AssociatedPackageUrl field.
func (o *ApiPolicyWaiverDTO) SetAssociatedPackageUrl(v string) {
	o.AssociatedPackageUrl = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ApiPolicyWaiverDTO) SetComment(v string) {
	o.Comment = &v
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *ApiPolicyWaiverDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetComponentName() string {
	if o == nil || IsNil(o.ComponentName) {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetComponentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentName) {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasComponentName() bool {
	if o != nil && !IsNil(o.ComponentName) {
		return true
	}

	return false
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *ApiPolicyWaiverDTO) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetComponentUpgradeAvailable returns the ComponentUpgradeAvailable field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetComponentUpgradeAvailable() bool {
	if o == nil || IsNil(o.ComponentUpgradeAvailable) {
		var ret bool
		return ret
	}
	return *o.ComponentUpgradeAvailable
}

// GetComponentUpgradeAvailableOk returns a tuple with the ComponentUpgradeAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetComponentUpgradeAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.ComponentUpgradeAvailable) {
		return nil, false
	}
	return o.ComponentUpgradeAvailable, true
}

// HasComponentUpgradeAvailable returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasComponentUpgradeAvailable() bool {
	if o != nil && !IsNil(o.ComponentUpgradeAvailable) {
		return true
	}

	return false
}

// SetComponentUpgradeAvailable gets a reference to the given bool and assigns it to the ComponentUpgradeAvailable field.
func (o *ApiPolicyWaiverDTO) SetComponentUpgradeAvailable(v bool) {
	o.ComponentUpgradeAvailable = &v
}

// GetConstraintFacts returns the ConstraintFacts field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetConstraintFacts() []ConstraintFact {
	if o == nil || IsNil(o.ConstraintFacts) {
		var ret []ConstraintFact
		return ret
	}
	return o.ConstraintFacts
}

// GetConstraintFactsOk returns a tuple with the ConstraintFacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetConstraintFactsOk() ([]ConstraintFact, bool) {
	if o == nil || IsNil(o.ConstraintFacts) {
		return nil, false
	}
	return o.ConstraintFacts, true
}

// HasConstraintFacts returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasConstraintFacts() bool {
	if o != nil && !IsNil(o.ConstraintFacts) {
		return true
	}

	return false
}

// SetConstraintFacts gets a reference to the given []ConstraintFact and assigns it to the ConstraintFacts field.
func (o *ApiPolicyWaiverDTO) SetConstraintFacts(v []ConstraintFact) {
	o.ConstraintFacts = v
}

// GetConstraintFactsJson returns the ConstraintFactsJson field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetConstraintFactsJson() string {
	if o == nil || IsNil(o.ConstraintFactsJson) {
		var ret string
		return ret
	}
	return *o.ConstraintFactsJson
}

// GetConstraintFactsJsonOk returns a tuple with the ConstraintFactsJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetConstraintFactsJsonOk() (*string, bool) {
	if o == nil || IsNil(o.ConstraintFactsJson) {
		return nil, false
	}
	return o.ConstraintFactsJson, true
}

// HasConstraintFactsJson returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasConstraintFactsJson() bool {
	if o != nil && !IsNil(o.ConstraintFactsJson) {
		return true
	}

	return false
}

// SetConstraintFactsJson gets a reference to the given string and assigns it to the ConstraintFactsJson field.
func (o *ApiPolicyWaiverDTO) SetConstraintFactsJson(v string) {
	o.ConstraintFactsJson = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetCreateTime() time.Time {
	if o == nil || IsNil(o.CreateTime) {
		var ret time.Time
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetCreateTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given time.Time and assigns it to the CreateTime field.
func (o *ApiPolicyWaiverDTO) SetCreateTime(v time.Time) {
	o.CreateTime = &v
}

// GetCreatorId returns the CreatorId field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetCreatorId() string {
	if o == nil || IsNil(o.CreatorId) {
		var ret string
		return ret
	}
	return *o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetCreatorIdOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorId) {
		return nil, false
	}
	return o.CreatorId, true
}

// HasCreatorId returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasCreatorId() bool {
	if o != nil && !IsNil(o.CreatorId) {
		return true
	}

	return false
}

// SetCreatorId gets a reference to the given string and assigns it to the CreatorId field.
func (o *ApiPolicyWaiverDTO) SetCreatorId(v string) {
	o.CreatorId = &v
}

// GetCreatorName returns the CreatorName field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetCreatorName() string {
	if o == nil || IsNil(o.CreatorName) {
		var ret string
		return ret
	}
	return *o.CreatorName
}

// GetCreatorNameOk returns a tuple with the CreatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetCreatorNameOk() (*string, bool) {
	if o == nil || IsNil(o.CreatorName) {
		return nil, false
	}
	return o.CreatorName, true
}

// HasCreatorName returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasCreatorName() bool {
	if o != nil && !IsNil(o.CreatorName) {
		return true
	}

	return false
}

// SetCreatorName gets a reference to the given string and assigns it to the CreatorName field.
func (o *ApiPolicyWaiverDTO) SetCreatorName(v string) {
	o.CreatorName = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetDisplayName() ComponentDisplayName {
	if o == nil || IsNil(o.DisplayName) {
		var ret ComponentDisplayName
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetDisplayNameOk() (*ComponentDisplayName, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given ComponentDisplayName and assigns it to the DisplayName field.
func (o *ApiPolicyWaiverDTO) SetDisplayName(v ComponentDisplayName) {
	o.DisplayName = &v
}

// GetExpiryTime returns the ExpiryTime field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetExpiryTime() time.Time {
	if o == nil || IsNil(o.ExpiryTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpiryTime
}

// GetExpiryTimeOk returns a tuple with the ExpiryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetExpiryTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiryTime) {
		return nil, false
	}
	return o.ExpiryTime, true
}

// HasExpiryTime returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasExpiryTime() bool {
	if o != nil && !IsNil(o.ExpiryTime) {
		return true
	}

	return false
}

// SetExpiryTime gets a reference to the given time.Time and assigns it to the ExpiryTime field.
func (o *ApiPolicyWaiverDTO) SetExpiryTime(v time.Time) {
	o.ExpiryTime = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ApiPolicyWaiverDTO) SetHash(v string) {
	o.Hash = &v
}

// GetIsObsolete returns the IsObsolete field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetIsObsolete() bool {
	if o == nil || IsNil(o.IsObsolete) {
		var ret bool
		return ret
	}
	return *o.IsObsolete
}

// GetIsObsoleteOk returns a tuple with the IsObsolete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetIsObsoleteOk() (*bool, bool) {
	if o == nil || IsNil(o.IsObsolete) {
		return nil, false
	}
	return o.IsObsolete, true
}

// HasIsObsolete returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasIsObsolete() bool {
	if o != nil && !IsNil(o.IsObsolete) {
		return true
	}

	return false
}

// SetIsObsolete gets a reference to the given bool and assigns it to the IsObsolete field.
func (o *ApiPolicyWaiverDTO) SetIsObsolete(v bool) {
	o.IsObsolete = &v
}

// GetMatcherStrategy returns the MatcherStrategy field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetMatcherStrategy() string {
	if o == nil || IsNil(o.MatcherStrategy) {
		var ret string
		return ret
	}
	return *o.MatcherStrategy
}

// GetMatcherStrategyOk returns a tuple with the MatcherStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetMatcherStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.MatcherStrategy) {
		return nil, false
	}
	return o.MatcherStrategy, true
}

// HasMatcherStrategy returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasMatcherStrategy() bool {
	if o != nil && !IsNil(o.MatcherStrategy) {
		return true
	}

	return false
}

// SetMatcherStrategy gets a reference to the given string and assigns it to the MatcherStrategy field.
func (o *ApiPolicyWaiverDTO) SetMatcherStrategy(v string) {
	o.MatcherStrategy = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *ApiPolicyWaiverDTO) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName) {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetPolicyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyName) {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasPolicyName() bool {
	if o != nil && !IsNil(o.PolicyName) {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *ApiPolicyWaiverDTO) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetPolicyViolationId returns the PolicyViolationId field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetPolicyViolationId() string {
	if o == nil || IsNil(o.PolicyViolationId) {
		var ret string
		return ret
	}
	return *o.PolicyViolationId
}

// GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetPolicyViolationIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyViolationId) {
		return nil, false
	}
	return o.PolicyViolationId, true
}

// HasPolicyViolationId returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasPolicyViolationId() bool {
	if o != nil && !IsNil(o.PolicyViolationId) {
		return true
	}

	return false
}

// SetPolicyViolationId gets a reference to the given string and assigns it to the PolicyViolationId field.
func (o *ApiPolicyWaiverDTO) SetPolicyViolationId(v string) {
	o.PolicyViolationId = &v
}

// GetPolicyWaiverId returns the PolicyWaiverId field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetPolicyWaiverId() string {
	if o == nil || IsNil(o.PolicyWaiverId) {
		var ret string
		return ret
	}
	return *o.PolicyWaiverId
}

// GetPolicyWaiverIdOk returns a tuple with the PolicyWaiverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetPolicyWaiverIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyWaiverId) {
		return nil, false
	}
	return o.PolicyWaiverId, true
}

// HasPolicyWaiverId returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasPolicyWaiverId() bool {
	if o != nil && !IsNil(o.PolicyWaiverId) {
		return true
	}

	return false
}

// SetPolicyWaiverId gets a reference to the given string and assigns it to the PolicyWaiverId field.
func (o *ApiPolicyWaiverDTO) SetPolicyWaiverId(v string) {
	o.PolicyWaiverId = &v
}

// GetScopeOwnerId returns the ScopeOwnerId field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetScopeOwnerId() string {
	if o == nil || IsNil(o.ScopeOwnerId) {
		var ret string
		return ret
	}
	return *o.ScopeOwnerId
}

// GetScopeOwnerIdOk returns a tuple with the ScopeOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetScopeOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeOwnerId) {
		return nil, false
	}
	return o.ScopeOwnerId, true
}

// HasScopeOwnerId returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasScopeOwnerId() bool {
	if o != nil && !IsNil(o.ScopeOwnerId) {
		return true
	}

	return false
}

// SetScopeOwnerId gets a reference to the given string and assigns it to the ScopeOwnerId field.
func (o *ApiPolicyWaiverDTO) SetScopeOwnerId(v string) {
	o.ScopeOwnerId = &v
}

// GetScopeOwnerName returns the ScopeOwnerName field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetScopeOwnerName() string {
	if o == nil || IsNil(o.ScopeOwnerName) {
		var ret string
		return ret
	}
	return *o.ScopeOwnerName
}

// GetScopeOwnerNameOk returns a tuple with the ScopeOwnerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetScopeOwnerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeOwnerName) {
		return nil, false
	}
	return o.ScopeOwnerName, true
}

// HasScopeOwnerName returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasScopeOwnerName() bool {
	if o != nil && !IsNil(o.ScopeOwnerName) {
		return true
	}

	return false
}

// SetScopeOwnerName gets a reference to the given string and assigns it to the ScopeOwnerName field.
func (o *ApiPolicyWaiverDTO) SetScopeOwnerName(v string) {
	o.ScopeOwnerName = &v
}

// GetScopeOwnerType returns the ScopeOwnerType field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetScopeOwnerType() string {
	if o == nil || IsNil(o.ScopeOwnerType) {
		var ret string
		return ret
	}
	return *o.ScopeOwnerType
}

// GetScopeOwnerTypeOk returns a tuple with the ScopeOwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetScopeOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ScopeOwnerType) {
		return nil, false
	}
	return o.ScopeOwnerType, true
}

// HasScopeOwnerType returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasScopeOwnerType() bool {
	if o != nil && !IsNil(o.ScopeOwnerType) {
		return true
	}

	return false
}

// SetScopeOwnerType gets a reference to the given string and assigns it to the ScopeOwnerType field.
func (o *ApiPolicyWaiverDTO) SetScopeOwnerType(v string) {
	o.ScopeOwnerType = &v
}

// GetThreatLevel returns the ThreatLevel field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetThreatLevel() int32 {
	if o == nil || IsNil(o.ThreatLevel) {
		var ret int32
		return ret
	}
	return *o.ThreatLevel
}

// GetThreatLevelOk returns a tuple with the ThreatLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetThreatLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.ThreatLevel) {
		return nil, false
	}
	return o.ThreatLevel, true
}

// HasThreatLevel returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasThreatLevel() bool {
	if o != nil && !IsNil(o.ThreatLevel) {
		return true
	}

	return false
}

// SetThreatLevel gets a reference to the given int32 and assigns it to the ThreatLevel field.
func (o *ApiPolicyWaiverDTO) SetThreatLevel(v int32) {
	o.ThreatLevel = &v
}

// GetVulnerabilityId returns the VulnerabilityId field value if set, zero value otherwise.
func (o *ApiPolicyWaiverDTO) GetVulnerabilityId() string {
	if o == nil || IsNil(o.VulnerabilityId) {
		var ret string
		return ret
	}
	return *o.VulnerabilityId
}

// GetVulnerabilityIdOk returns a tuple with the VulnerabilityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPolicyWaiverDTO) GetVulnerabilityIdOk() (*string, bool) {
	if o == nil || IsNil(o.VulnerabilityId) {
		return nil, false
	}
	return o.VulnerabilityId, true
}

// HasVulnerabilityId returns a boolean if a field has been set.
func (o *ApiPolicyWaiverDTO) HasVulnerabilityId() bool {
	if o != nil && !IsNil(o.VulnerabilityId) {
		return true
	}

	return false
}

// SetVulnerabilityId gets a reference to the given string and assigns it to the VulnerabilityId field.
func (o *ApiPolicyWaiverDTO) SetVulnerabilityId(v string) {
	o.VulnerabilityId = &v
}

func (o ApiPolicyWaiverDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPolicyWaiverDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AssociatedPackageUrl) {
		toSerialize["associatedPackageUrl"] = o.AssociatedPackageUrl
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.ComponentName) {
		toSerialize["componentName"] = o.ComponentName
	}
	if !IsNil(o.ComponentUpgradeAvailable) {
		toSerialize["componentUpgradeAvailable"] = o.ComponentUpgradeAvailable
	}
	if !IsNil(o.ConstraintFacts) {
		toSerialize["constraintFacts"] = o.ConstraintFacts
	}
	if !IsNil(o.ConstraintFactsJson) {
		toSerialize["constraintFactsJson"] = o.ConstraintFactsJson
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
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.ExpiryTime) {
		toSerialize["expiryTime"] = o.ExpiryTime
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.IsObsolete) {
		toSerialize["isObsolete"] = o.IsObsolete
	}
	if !IsNil(o.MatcherStrategy) {
		toSerialize["matcherStrategy"] = o.MatcherStrategy
	}
	if !IsNil(o.PolicyId) {
		toSerialize["policyId"] = o.PolicyId
	}
	if !IsNil(o.PolicyName) {
		toSerialize["policyName"] = o.PolicyName
	}
	if !IsNil(o.PolicyViolationId) {
		toSerialize["policyViolationId"] = o.PolicyViolationId
	}
	if !IsNil(o.PolicyWaiverId) {
		toSerialize["policyWaiverId"] = o.PolicyWaiverId
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
	if !IsNil(o.ThreatLevel) {
		toSerialize["threatLevel"] = o.ThreatLevel
	}
	if !IsNil(o.VulnerabilityId) {
		toSerialize["vulnerabilityId"] = o.VulnerabilityId
	}
	return toSerialize, nil
}

type NullableApiPolicyWaiverDTO struct {
	value *ApiPolicyWaiverDTO
	isSet bool
}

func (v NullableApiPolicyWaiverDTO) Get() *ApiPolicyWaiverDTO {
	return v.value
}

func (v *NullableApiPolicyWaiverDTO) Set(val *ApiPolicyWaiverDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPolicyWaiverDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPolicyWaiverDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPolicyWaiverDTO(val *ApiPolicyWaiverDTO) *NullableApiPolicyWaiverDTO {
	return &NullableApiPolicyWaiverDTO{value: val, isSet: true}
}

func (v NullableApiPolicyWaiverDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPolicyWaiverDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


