# ApiPolicyWaiverDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedPackageUrl** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**ComponentName** | Pointer to **string** |  | [optional] 
**ComponentUpgradeAvailable** | Pointer to **bool** |  | [optional] 
**ConstraintFacts** | Pointer to [**[]ConstraintFact**](ConstraintFact.md) |  | [optional] 
**ConstraintFactsJson** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**CreatorName** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to [**ComponentDisplayName**](ComponentDisplayName.md) |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**IsObsolete** | Pointer to **bool** |  | [optional] 
**MatcherStrategy** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**PolicyWaiverId** | Pointer to **string** |  | [optional] 
**ReasonText** | Pointer to **string** |  | [optional] 
**ScopeOwnerId** | Pointer to **string** |  | [optional] 
**ScopeOwnerName** | Pointer to **string** |  | [optional] 
**ScopeOwnerType** | Pointer to **string** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 
**VulnerabilityId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiPolicyWaiverDTO

`func NewApiPolicyWaiverDTO() *ApiPolicyWaiverDTO`

NewApiPolicyWaiverDTO instantiates a new ApiPolicyWaiverDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyWaiverDTOWithDefaults

`func NewApiPolicyWaiverDTOWithDefaults() *ApiPolicyWaiverDTO`

NewApiPolicyWaiverDTOWithDefaults instantiates a new ApiPolicyWaiverDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedPackageUrl

`func (o *ApiPolicyWaiverDTO) GetAssociatedPackageUrl() string`

GetAssociatedPackageUrl returns the AssociatedPackageUrl field if non-nil, zero value otherwise.

### GetAssociatedPackageUrlOk

`func (o *ApiPolicyWaiverDTO) GetAssociatedPackageUrlOk() (*string, bool)`

GetAssociatedPackageUrlOk returns a tuple with the AssociatedPackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedPackageUrl

`func (o *ApiPolicyWaiverDTO) SetAssociatedPackageUrl(v string)`

SetAssociatedPackageUrl sets AssociatedPackageUrl field to given value.

### HasAssociatedPackageUrl

`func (o *ApiPolicyWaiverDTO) HasAssociatedPackageUrl() bool`

HasAssociatedPackageUrl returns a boolean if a field has been set.

### GetComment

`func (o *ApiPolicyWaiverDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiPolicyWaiverDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiPolicyWaiverDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiPolicyWaiverDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *ApiPolicyWaiverDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiPolicyWaiverDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiPolicyWaiverDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiPolicyWaiverDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetComponentName

`func (o *ApiPolicyWaiverDTO) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *ApiPolicyWaiverDTO) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *ApiPolicyWaiverDTO) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *ApiPolicyWaiverDTO) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetComponentUpgradeAvailable

`func (o *ApiPolicyWaiverDTO) GetComponentUpgradeAvailable() bool`

GetComponentUpgradeAvailable returns the ComponentUpgradeAvailable field if non-nil, zero value otherwise.

### GetComponentUpgradeAvailableOk

`func (o *ApiPolicyWaiverDTO) GetComponentUpgradeAvailableOk() (*bool, bool)`

GetComponentUpgradeAvailableOk returns a tuple with the ComponentUpgradeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentUpgradeAvailable

`func (o *ApiPolicyWaiverDTO) SetComponentUpgradeAvailable(v bool)`

SetComponentUpgradeAvailable sets ComponentUpgradeAvailable field to given value.

### HasComponentUpgradeAvailable

`func (o *ApiPolicyWaiverDTO) HasComponentUpgradeAvailable() bool`

HasComponentUpgradeAvailable returns a boolean if a field has been set.

### GetConstraintFacts

`func (o *ApiPolicyWaiverDTO) GetConstraintFacts() []ConstraintFact`

GetConstraintFacts returns the ConstraintFacts field if non-nil, zero value otherwise.

### GetConstraintFactsOk

`func (o *ApiPolicyWaiverDTO) GetConstraintFactsOk() (*[]ConstraintFact, bool)`

GetConstraintFactsOk returns a tuple with the ConstraintFacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintFacts

`func (o *ApiPolicyWaiverDTO) SetConstraintFacts(v []ConstraintFact)`

SetConstraintFacts sets ConstraintFacts field to given value.

### HasConstraintFacts

`func (o *ApiPolicyWaiverDTO) HasConstraintFacts() bool`

HasConstraintFacts returns a boolean if a field has been set.

### GetConstraintFactsJson

`func (o *ApiPolicyWaiverDTO) GetConstraintFactsJson() string`

GetConstraintFactsJson returns the ConstraintFactsJson field if non-nil, zero value otherwise.

### GetConstraintFactsJsonOk

`func (o *ApiPolicyWaiverDTO) GetConstraintFactsJsonOk() (*string, bool)`

GetConstraintFactsJsonOk returns a tuple with the ConstraintFactsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintFactsJson

`func (o *ApiPolicyWaiverDTO) SetConstraintFactsJson(v string)`

SetConstraintFactsJson sets ConstraintFactsJson field to given value.

### HasConstraintFactsJson

`func (o *ApiPolicyWaiverDTO) HasConstraintFactsJson() bool`

HasConstraintFactsJson returns a boolean if a field has been set.

### GetCreateTime

`func (o *ApiPolicyWaiverDTO) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ApiPolicyWaiverDTO) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ApiPolicyWaiverDTO) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ApiPolicyWaiverDTO) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *ApiPolicyWaiverDTO) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ApiPolicyWaiverDTO) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ApiPolicyWaiverDTO) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *ApiPolicyWaiverDTO) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCreatorName

`func (o *ApiPolicyWaiverDTO) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *ApiPolicyWaiverDTO) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *ApiPolicyWaiverDTO) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *ApiPolicyWaiverDTO) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiPolicyWaiverDTO) GetDisplayName() ComponentDisplayName`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiPolicyWaiverDTO) GetDisplayNameOk() (*ComponentDisplayName, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiPolicyWaiverDTO) SetDisplayName(v ComponentDisplayName)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiPolicyWaiverDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExpiryTime

`func (o *ApiPolicyWaiverDTO) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ApiPolicyWaiverDTO) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ApiPolicyWaiverDTO) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *ApiPolicyWaiverDTO) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetHash

`func (o *ApiPolicyWaiverDTO) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiPolicyWaiverDTO) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiPolicyWaiverDTO) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiPolicyWaiverDTO) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetIsObsolete

`func (o *ApiPolicyWaiverDTO) GetIsObsolete() bool`

GetIsObsolete returns the IsObsolete field if non-nil, zero value otherwise.

### GetIsObsoleteOk

`func (o *ApiPolicyWaiverDTO) GetIsObsoleteOk() (*bool, bool)`

GetIsObsoleteOk returns a tuple with the IsObsolete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsObsolete

`func (o *ApiPolicyWaiverDTO) SetIsObsolete(v bool)`

SetIsObsolete sets IsObsolete field to given value.

### HasIsObsolete

`func (o *ApiPolicyWaiverDTO) HasIsObsolete() bool`

HasIsObsolete returns a boolean if a field has been set.

### GetMatcherStrategy

`func (o *ApiPolicyWaiverDTO) GetMatcherStrategy() string`

GetMatcherStrategy returns the MatcherStrategy field if non-nil, zero value otherwise.

### GetMatcherStrategyOk

`func (o *ApiPolicyWaiverDTO) GetMatcherStrategyOk() (*string, bool)`

GetMatcherStrategyOk returns a tuple with the MatcherStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcherStrategy

`func (o *ApiPolicyWaiverDTO) SetMatcherStrategy(v string)`

SetMatcherStrategy sets MatcherStrategy field to given value.

### HasMatcherStrategy

`func (o *ApiPolicyWaiverDTO) HasMatcherStrategy() bool`

HasMatcherStrategy returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApiPolicyWaiverDTO) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApiPolicyWaiverDTO) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApiPolicyWaiverDTO) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApiPolicyWaiverDTO) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiPolicyWaiverDTO) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiPolicyWaiverDTO) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiPolicyWaiverDTO) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiPolicyWaiverDTO) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *ApiPolicyWaiverDTO) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *ApiPolicyWaiverDTO) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *ApiPolicyWaiverDTO) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *ApiPolicyWaiverDTO) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetPolicyWaiverId

`func (o *ApiPolicyWaiverDTO) GetPolicyWaiverId() string`

GetPolicyWaiverId returns the PolicyWaiverId field if non-nil, zero value otherwise.

### GetPolicyWaiverIdOk

`func (o *ApiPolicyWaiverDTO) GetPolicyWaiverIdOk() (*string, bool)`

GetPolicyWaiverIdOk returns a tuple with the PolicyWaiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverId

`func (o *ApiPolicyWaiverDTO) SetPolicyWaiverId(v string)`

SetPolicyWaiverId sets PolicyWaiverId field to given value.

### HasPolicyWaiverId

`func (o *ApiPolicyWaiverDTO) HasPolicyWaiverId() bool`

HasPolicyWaiverId returns a boolean if a field has been set.

### GetReasonText

`func (o *ApiPolicyWaiverDTO) GetReasonText() string`

GetReasonText returns the ReasonText field if non-nil, zero value otherwise.

### GetReasonTextOk

`func (o *ApiPolicyWaiverDTO) GetReasonTextOk() (*string, bool)`

GetReasonTextOk returns a tuple with the ReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonText

`func (o *ApiPolicyWaiverDTO) SetReasonText(v string)`

SetReasonText sets ReasonText field to given value.

### HasReasonText

`func (o *ApiPolicyWaiverDTO) HasReasonText() bool`

HasReasonText returns a boolean if a field has been set.

### GetScopeOwnerId

`func (o *ApiPolicyWaiverDTO) GetScopeOwnerId() string`

GetScopeOwnerId returns the ScopeOwnerId field if non-nil, zero value otherwise.

### GetScopeOwnerIdOk

`func (o *ApiPolicyWaiverDTO) GetScopeOwnerIdOk() (*string, bool)`

GetScopeOwnerIdOk returns a tuple with the ScopeOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOwnerId

`func (o *ApiPolicyWaiverDTO) SetScopeOwnerId(v string)`

SetScopeOwnerId sets ScopeOwnerId field to given value.

### HasScopeOwnerId

`func (o *ApiPolicyWaiverDTO) HasScopeOwnerId() bool`

HasScopeOwnerId returns a boolean if a field has been set.

### GetScopeOwnerName

`func (o *ApiPolicyWaiverDTO) GetScopeOwnerName() string`

GetScopeOwnerName returns the ScopeOwnerName field if non-nil, zero value otherwise.

### GetScopeOwnerNameOk

`func (o *ApiPolicyWaiverDTO) GetScopeOwnerNameOk() (*string, bool)`

GetScopeOwnerNameOk returns a tuple with the ScopeOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOwnerName

`func (o *ApiPolicyWaiverDTO) SetScopeOwnerName(v string)`

SetScopeOwnerName sets ScopeOwnerName field to given value.

### HasScopeOwnerName

`func (o *ApiPolicyWaiverDTO) HasScopeOwnerName() bool`

HasScopeOwnerName returns a boolean if a field has been set.

### GetScopeOwnerType

`func (o *ApiPolicyWaiverDTO) GetScopeOwnerType() string`

GetScopeOwnerType returns the ScopeOwnerType field if non-nil, zero value otherwise.

### GetScopeOwnerTypeOk

`func (o *ApiPolicyWaiverDTO) GetScopeOwnerTypeOk() (*string, bool)`

GetScopeOwnerTypeOk returns a tuple with the ScopeOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOwnerType

`func (o *ApiPolicyWaiverDTO) SetScopeOwnerType(v string)`

SetScopeOwnerType sets ScopeOwnerType field to given value.

### HasScopeOwnerType

`func (o *ApiPolicyWaiverDTO) HasScopeOwnerType() bool`

HasScopeOwnerType returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiPolicyWaiverDTO) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiPolicyWaiverDTO) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiPolicyWaiverDTO) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiPolicyWaiverDTO) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.

### GetVulnerabilityId

`func (o *ApiPolicyWaiverDTO) GetVulnerabilityId() string`

GetVulnerabilityId returns the VulnerabilityId field if non-nil, zero value otherwise.

### GetVulnerabilityIdOk

`func (o *ApiPolicyWaiverDTO) GetVulnerabilityIdOk() (*string, bool)`

GetVulnerabilityIdOk returns a tuple with the VulnerabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityId

`func (o *ApiPolicyWaiverDTO) SetVulnerabilityId(v string)`

SetVulnerabilityId sets VulnerabilityId field to given value.

### HasVulnerabilityId

`func (o *ApiPolicyWaiverDTO) HasVulnerabilityId() bool`

HasVulnerabilityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


