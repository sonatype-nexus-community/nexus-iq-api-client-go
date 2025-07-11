# ApiPolicyWaiverRequestDTO

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
**DisplayName** | Pointer to [**ComponentDisplayName**](ComponentDisplayName.md) |  | [optional] 
**ExpireWhenRemediationAvailable** | Pointer to **bool** |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**IsObsolete** | Pointer to **bool** |  | [optional] 
**MatcherStrategy** | Pointer to **string** |  | [optional] 
**NoteToReviewer** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**PolicyWaiverReasonId** | Pointer to **string** |  | [optional] 
**PolicyWaiverRequestId** | Pointer to **string** |  | [optional] 
**ReasonText** | Pointer to **string** |  | [optional] 
**RejectionReason** | Pointer to **string** |  | [optional] 
**RequestTime** | Pointer to **time.Time** |  | [optional] 
**RequesterId** | Pointer to **string** |  | [optional] 
**RequesterName** | Pointer to **string** |  | [optional] 
**ReviewerId** | Pointer to **string** |  | [optional] 
**ReviewerName** | Pointer to **string** |  | [optional] 
**ScopeOwnerId** | Pointer to **string** |  | [optional] 
**ScopeOwnerName** | Pointer to **string** |  | [optional] 
**ScopeOwnerType** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 
**VulnerabilityId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiPolicyWaiverRequestDTO

`func NewApiPolicyWaiverRequestDTO() *ApiPolicyWaiverRequestDTO`

NewApiPolicyWaiverRequestDTO instantiates a new ApiPolicyWaiverRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyWaiverRequestDTOWithDefaults

`func NewApiPolicyWaiverRequestDTOWithDefaults() *ApiPolicyWaiverRequestDTO`

NewApiPolicyWaiverRequestDTOWithDefaults instantiates a new ApiPolicyWaiverRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedPackageUrl

`func (o *ApiPolicyWaiverRequestDTO) GetAssociatedPackageUrl() string`

GetAssociatedPackageUrl returns the AssociatedPackageUrl field if non-nil, zero value otherwise.

### GetAssociatedPackageUrlOk

`func (o *ApiPolicyWaiverRequestDTO) GetAssociatedPackageUrlOk() (*string, bool)`

GetAssociatedPackageUrlOk returns a tuple with the AssociatedPackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedPackageUrl

`func (o *ApiPolicyWaiverRequestDTO) SetAssociatedPackageUrl(v string)`

SetAssociatedPackageUrl sets AssociatedPackageUrl field to given value.

### HasAssociatedPackageUrl

`func (o *ApiPolicyWaiverRequestDTO) HasAssociatedPackageUrl() bool`

HasAssociatedPackageUrl returns a boolean if a field has been set.

### GetComment

`func (o *ApiPolicyWaiverRequestDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiPolicyWaiverRequestDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiPolicyWaiverRequestDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiPolicyWaiverRequestDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *ApiPolicyWaiverRequestDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiPolicyWaiverRequestDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiPolicyWaiverRequestDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiPolicyWaiverRequestDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetComponentName

`func (o *ApiPolicyWaiverRequestDTO) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *ApiPolicyWaiverRequestDTO) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *ApiPolicyWaiverRequestDTO) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *ApiPolicyWaiverRequestDTO) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetComponentUpgradeAvailable

`func (o *ApiPolicyWaiverRequestDTO) GetComponentUpgradeAvailable() bool`

GetComponentUpgradeAvailable returns the ComponentUpgradeAvailable field if non-nil, zero value otherwise.

### GetComponentUpgradeAvailableOk

`func (o *ApiPolicyWaiverRequestDTO) GetComponentUpgradeAvailableOk() (*bool, bool)`

GetComponentUpgradeAvailableOk returns a tuple with the ComponentUpgradeAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentUpgradeAvailable

`func (o *ApiPolicyWaiverRequestDTO) SetComponentUpgradeAvailable(v bool)`

SetComponentUpgradeAvailable sets ComponentUpgradeAvailable field to given value.

### HasComponentUpgradeAvailable

`func (o *ApiPolicyWaiverRequestDTO) HasComponentUpgradeAvailable() bool`

HasComponentUpgradeAvailable returns a boolean if a field has been set.

### GetConstraintFacts

`func (o *ApiPolicyWaiverRequestDTO) GetConstraintFacts() []ConstraintFact`

GetConstraintFacts returns the ConstraintFacts field if non-nil, zero value otherwise.

### GetConstraintFactsOk

`func (o *ApiPolicyWaiverRequestDTO) GetConstraintFactsOk() (*[]ConstraintFact, bool)`

GetConstraintFactsOk returns a tuple with the ConstraintFacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintFacts

`func (o *ApiPolicyWaiverRequestDTO) SetConstraintFacts(v []ConstraintFact)`

SetConstraintFacts sets ConstraintFacts field to given value.

### HasConstraintFacts

`func (o *ApiPolicyWaiverRequestDTO) HasConstraintFacts() bool`

HasConstraintFacts returns a boolean if a field has been set.

### GetConstraintFactsJson

`func (o *ApiPolicyWaiverRequestDTO) GetConstraintFactsJson() string`

GetConstraintFactsJson returns the ConstraintFactsJson field if non-nil, zero value otherwise.

### GetConstraintFactsJsonOk

`func (o *ApiPolicyWaiverRequestDTO) GetConstraintFactsJsonOk() (*string, bool)`

GetConstraintFactsJsonOk returns a tuple with the ConstraintFactsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintFactsJson

`func (o *ApiPolicyWaiverRequestDTO) SetConstraintFactsJson(v string)`

SetConstraintFactsJson sets ConstraintFactsJson field to given value.

### HasConstraintFactsJson

`func (o *ApiPolicyWaiverRequestDTO) HasConstraintFactsJson() bool`

HasConstraintFactsJson returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiPolicyWaiverRequestDTO) GetDisplayName() ComponentDisplayName`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiPolicyWaiverRequestDTO) GetDisplayNameOk() (*ComponentDisplayName, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiPolicyWaiverRequestDTO) SetDisplayName(v ComponentDisplayName)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiPolicyWaiverRequestDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExpireWhenRemediationAvailable

`func (o *ApiPolicyWaiverRequestDTO) GetExpireWhenRemediationAvailable() bool`

GetExpireWhenRemediationAvailable returns the ExpireWhenRemediationAvailable field if non-nil, zero value otherwise.

### GetExpireWhenRemediationAvailableOk

`func (o *ApiPolicyWaiverRequestDTO) GetExpireWhenRemediationAvailableOk() (*bool, bool)`

GetExpireWhenRemediationAvailableOk returns a tuple with the ExpireWhenRemediationAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWhenRemediationAvailable

`func (o *ApiPolicyWaiverRequestDTO) SetExpireWhenRemediationAvailable(v bool)`

SetExpireWhenRemediationAvailable sets ExpireWhenRemediationAvailable field to given value.

### HasExpireWhenRemediationAvailable

`func (o *ApiPolicyWaiverRequestDTO) HasExpireWhenRemediationAvailable() bool`

HasExpireWhenRemediationAvailable returns a boolean if a field has been set.

### GetExpiryTime

`func (o *ApiPolicyWaiverRequestDTO) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ApiPolicyWaiverRequestDTO) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ApiPolicyWaiverRequestDTO) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *ApiPolicyWaiverRequestDTO) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetHash

`func (o *ApiPolicyWaiverRequestDTO) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiPolicyWaiverRequestDTO) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiPolicyWaiverRequestDTO) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiPolicyWaiverRequestDTO) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetIsObsolete

`func (o *ApiPolicyWaiverRequestDTO) GetIsObsolete() bool`

GetIsObsolete returns the IsObsolete field if non-nil, zero value otherwise.

### GetIsObsoleteOk

`func (o *ApiPolicyWaiverRequestDTO) GetIsObsoleteOk() (*bool, bool)`

GetIsObsoleteOk returns a tuple with the IsObsolete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsObsolete

`func (o *ApiPolicyWaiverRequestDTO) SetIsObsolete(v bool)`

SetIsObsolete sets IsObsolete field to given value.

### HasIsObsolete

`func (o *ApiPolicyWaiverRequestDTO) HasIsObsolete() bool`

HasIsObsolete returns a boolean if a field has been set.

### GetMatcherStrategy

`func (o *ApiPolicyWaiverRequestDTO) GetMatcherStrategy() string`

GetMatcherStrategy returns the MatcherStrategy field if non-nil, zero value otherwise.

### GetMatcherStrategyOk

`func (o *ApiPolicyWaiverRequestDTO) GetMatcherStrategyOk() (*string, bool)`

GetMatcherStrategyOk returns a tuple with the MatcherStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcherStrategy

`func (o *ApiPolicyWaiverRequestDTO) SetMatcherStrategy(v string)`

SetMatcherStrategy sets MatcherStrategy field to given value.

### HasMatcherStrategy

`func (o *ApiPolicyWaiverRequestDTO) HasMatcherStrategy() bool`

HasMatcherStrategy returns a boolean if a field has been set.

### GetNoteToReviewer

`func (o *ApiPolicyWaiverRequestDTO) GetNoteToReviewer() string`

GetNoteToReviewer returns the NoteToReviewer field if non-nil, zero value otherwise.

### GetNoteToReviewerOk

`func (o *ApiPolicyWaiverRequestDTO) GetNoteToReviewerOk() (*string, bool)`

GetNoteToReviewerOk returns a tuple with the NoteToReviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteToReviewer

`func (o *ApiPolicyWaiverRequestDTO) SetNoteToReviewer(v string)`

SetNoteToReviewer sets NoteToReviewer field to given value.

### HasNoteToReviewer

`func (o *ApiPolicyWaiverRequestDTO) HasNoteToReviewer() bool`

HasNoteToReviewer returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApiPolicyWaiverRequestDTO) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApiPolicyWaiverRequestDTO) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApiPolicyWaiverRequestDTO) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApiPolicyWaiverRequestDTO) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiPolicyWaiverRequestDTO) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiPolicyWaiverRequestDTO) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiPolicyWaiverRequestDTO) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiPolicyWaiverRequestDTO) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *ApiPolicyWaiverRequestDTO) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *ApiPolicyWaiverRequestDTO) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *ApiPolicyWaiverRequestDTO) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *ApiPolicyWaiverRequestDTO) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetPolicyWaiverReasonId

`func (o *ApiPolicyWaiverRequestDTO) GetPolicyWaiverReasonId() string`

GetPolicyWaiverReasonId returns the PolicyWaiverReasonId field if non-nil, zero value otherwise.

### GetPolicyWaiverReasonIdOk

`func (o *ApiPolicyWaiverRequestDTO) GetPolicyWaiverReasonIdOk() (*string, bool)`

GetPolicyWaiverReasonIdOk returns a tuple with the PolicyWaiverReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverReasonId

`func (o *ApiPolicyWaiverRequestDTO) SetPolicyWaiverReasonId(v string)`

SetPolicyWaiverReasonId sets PolicyWaiverReasonId field to given value.

### HasPolicyWaiverReasonId

`func (o *ApiPolicyWaiverRequestDTO) HasPolicyWaiverReasonId() bool`

HasPolicyWaiverReasonId returns a boolean if a field has been set.

### GetPolicyWaiverRequestId

`func (o *ApiPolicyWaiverRequestDTO) GetPolicyWaiverRequestId() string`

GetPolicyWaiverRequestId returns the PolicyWaiverRequestId field if non-nil, zero value otherwise.

### GetPolicyWaiverRequestIdOk

`func (o *ApiPolicyWaiverRequestDTO) GetPolicyWaiverRequestIdOk() (*string, bool)`

GetPolicyWaiverRequestIdOk returns a tuple with the PolicyWaiverRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverRequestId

`func (o *ApiPolicyWaiverRequestDTO) SetPolicyWaiverRequestId(v string)`

SetPolicyWaiverRequestId sets PolicyWaiverRequestId field to given value.

### HasPolicyWaiverRequestId

`func (o *ApiPolicyWaiverRequestDTO) HasPolicyWaiverRequestId() bool`

HasPolicyWaiverRequestId returns a boolean if a field has been set.

### GetReasonText

`func (o *ApiPolicyWaiverRequestDTO) GetReasonText() string`

GetReasonText returns the ReasonText field if non-nil, zero value otherwise.

### GetReasonTextOk

`func (o *ApiPolicyWaiverRequestDTO) GetReasonTextOk() (*string, bool)`

GetReasonTextOk returns a tuple with the ReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonText

`func (o *ApiPolicyWaiverRequestDTO) SetReasonText(v string)`

SetReasonText sets ReasonText field to given value.

### HasReasonText

`func (o *ApiPolicyWaiverRequestDTO) HasReasonText() bool`

HasReasonText returns a boolean if a field has been set.

### GetRejectionReason

`func (o *ApiPolicyWaiverRequestDTO) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *ApiPolicyWaiverRequestDTO) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *ApiPolicyWaiverRequestDTO) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *ApiPolicyWaiverRequestDTO) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetRequestTime

`func (o *ApiPolicyWaiverRequestDTO) GetRequestTime() time.Time`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *ApiPolicyWaiverRequestDTO) GetRequestTimeOk() (*time.Time, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *ApiPolicyWaiverRequestDTO) SetRequestTime(v time.Time)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *ApiPolicyWaiverRequestDTO) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetRequesterId

`func (o *ApiPolicyWaiverRequestDTO) GetRequesterId() string`

GetRequesterId returns the RequesterId field if non-nil, zero value otherwise.

### GetRequesterIdOk

`func (o *ApiPolicyWaiverRequestDTO) GetRequesterIdOk() (*string, bool)`

GetRequesterIdOk returns a tuple with the RequesterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterId

`func (o *ApiPolicyWaiverRequestDTO) SetRequesterId(v string)`

SetRequesterId sets RequesterId field to given value.

### HasRequesterId

`func (o *ApiPolicyWaiverRequestDTO) HasRequesterId() bool`

HasRequesterId returns a boolean if a field has been set.

### GetRequesterName

`func (o *ApiPolicyWaiverRequestDTO) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *ApiPolicyWaiverRequestDTO) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *ApiPolicyWaiverRequestDTO) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *ApiPolicyWaiverRequestDTO) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### GetReviewerId

`func (o *ApiPolicyWaiverRequestDTO) GetReviewerId() string`

GetReviewerId returns the ReviewerId field if non-nil, zero value otherwise.

### GetReviewerIdOk

`func (o *ApiPolicyWaiverRequestDTO) GetReviewerIdOk() (*string, bool)`

GetReviewerIdOk returns a tuple with the ReviewerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerId

`func (o *ApiPolicyWaiverRequestDTO) SetReviewerId(v string)`

SetReviewerId sets ReviewerId field to given value.

### HasReviewerId

`func (o *ApiPolicyWaiverRequestDTO) HasReviewerId() bool`

HasReviewerId returns a boolean if a field has been set.

### GetReviewerName

`func (o *ApiPolicyWaiverRequestDTO) GetReviewerName() string`

GetReviewerName returns the ReviewerName field if non-nil, zero value otherwise.

### GetReviewerNameOk

`func (o *ApiPolicyWaiverRequestDTO) GetReviewerNameOk() (*string, bool)`

GetReviewerNameOk returns a tuple with the ReviewerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerName

`func (o *ApiPolicyWaiverRequestDTO) SetReviewerName(v string)`

SetReviewerName sets ReviewerName field to given value.

### HasReviewerName

`func (o *ApiPolicyWaiverRequestDTO) HasReviewerName() bool`

HasReviewerName returns a boolean if a field has been set.

### GetScopeOwnerId

`func (o *ApiPolicyWaiverRequestDTO) GetScopeOwnerId() string`

GetScopeOwnerId returns the ScopeOwnerId field if non-nil, zero value otherwise.

### GetScopeOwnerIdOk

`func (o *ApiPolicyWaiverRequestDTO) GetScopeOwnerIdOk() (*string, bool)`

GetScopeOwnerIdOk returns a tuple with the ScopeOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOwnerId

`func (o *ApiPolicyWaiverRequestDTO) SetScopeOwnerId(v string)`

SetScopeOwnerId sets ScopeOwnerId field to given value.

### HasScopeOwnerId

`func (o *ApiPolicyWaiverRequestDTO) HasScopeOwnerId() bool`

HasScopeOwnerId returns a boolean if a field has been set.

### GetScopeOwnerName

`func (o *ApiPolicyWaiverRequestDTO) GetScopeOwnerName() string`

GetScopeOwnerName returns the ScopeOwnerName field if non-nil, zero value otherwise.

### GetScopeOwnerNameOk

`func (o *ApiPolicyWaiverRequestDTO) GetScopeOwnerNameOk() (*string, bool)`

GetScopeOwnerNameOk returns a tuple with the ScopeOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOwnerName

`func (o *ApiPolicyWaiverRequestDTO) SetScopeOwnerName(v string)`

SetScopeOwnerName sets ScopeOwnerName field to given value.

### HasScopeOwnerName

`func (o *ApiPolicyWaiverRequestDTO) HasScopeOwnerName() bool`

HasScopeOwnerName returns a boolean if a field has been set.

### GetScopeOwnerType

`func (o *ApiPolicyWaiverRequestDTO) GetScopeOwnerType() string`

GetScopeOwnerType returns the ScopeOwnerType field if non-nil, zero value otherwise.

### GetScopeOwnerTypeOk

`func (o *ApiPolicyWaiverRequestDTO) GetScopeOwnerTypeOk() (*string, bool)`

GetScopeOwnerTypeOk returns a tuple with the ScopeOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOwnerType

`func (o *ApiPolicyWaiverRequestDTO) SetScopeOwnerType(v string)`

SetScopeOwnerType sets ScopeOwnerType field to given value.

### HasScopeOwnerType

`func (o *ApiPolicyWaiverRequestDTO) HasScopeOwnerType() bool`

HasScopeOwnerType returns a boolean if a field has been set.

### GetStatus

`func (o *ApiPolicyWaiverRequestDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiPolicyWaiverRequestDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiPolicyWaiverRequestDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiPolicyWaiverRequestDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiPolicyWaiverRequestDTO) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiPolicyWaiverRequestDTO) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiPolicyWaiverRequestDTO) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiPolicyWaiverRequestDTO) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.

### GetVulnerabilityId

`func (o *ApiPolicyWaiverRequestDTO) GetVulnerabilityId() string`

GetVulnerabilityId returns the VulnerabilityId field if non-nil, zero value otherwise.

### GetVulnerabilityIdOk

`func (o *ApiPolicyWaiverRequestDTO) GetVulnerabilityIdOk() (*string, bool)`

GetVulnerabilityIdOk returns a tuple with the VulnerabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityId

`func (o *ApiPolicyWaiverRequestDTO) SetVulnerabilityId(v string)`

SetVulnerabilityId sets VulnerabilityId field to given value.

### HasVulnerabilityId

`func (o *ApiPolicyWaiverRequestDTO) HasVulnerabilityId() bool`

HasVulnerabilityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


