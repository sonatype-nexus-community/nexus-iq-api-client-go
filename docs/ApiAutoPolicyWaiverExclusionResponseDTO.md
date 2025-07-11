# ApiAutoPolicyWaiverExclusionResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoPolicyWaiverExclusionId** | Pointer to **string** |  | [optional] 
**AutoPolicyWaiverId** | Pointer to **string** |  | [optional] 
**ComponentDisplayName** | Pointer to **string** |  | [optional] 
**ComponentIdentifier** | Pointer to [**ComponentIdentifier**](ComponentIdentifier.md) |  | [optional] 
**ComponentMatchStrategy** | Pointer to **string** |  | [optional] 
**ConstraintFacts** | Pointer to [**[]ConstraintFact**](ConstraintFact.md) |  | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**CreatorName** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerName** | Pointer to **string** |  | [optional] 
**OwnerPublicId** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**ScanId** | Pointer to **string** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 
**VulnerabilityIdentifiers** | Pointer to **string** |  | [optional] 

## Methods

### NewApiAutoPolicyWaiverExclusionResponseDTO

`func NewApiAutoPolicyWaiverExclusionResponseDTO() *ApiAutoPolicyWaiverExclusionResponseDTO`

NewApiAutoPolicyWaiverExclusionResponseDTO instantiates a new ApiAutoPolicyWaiverExclusionResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAutoPolicyWaiverExclusionResponseDTOWithDefaults

`func NewApiAutoPolicyWaiverExclusionResponseDTOWithDefaults() *ApiAutoPolicyWaiverExclusionResponseDTO`

NewApiAutoPolicyWaiverExclusionResponseDTOWithDefaults instantiates a new ApiAutoPolicyWaiverExclusionResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoPolicyWaiverExclusionId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetAutoPolicyWaiverExclusionId() string`

GetAutoPolicyWaiverExclusionId returns the AutoPolicyWaiverExclusionId field if non-nil, zero value otherwise.

### GetAutoPolicyWaiverExclusionIdOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetAutoPolicyWaiverExclusionIdOk() (*string, bool)`

GetAutoPolicyWaiverExclusionIdOk returns a tuple with the AutoPolicyWaiverExclusionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPolicyWaiverExclusionId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetAutoPolicyWaiverExclusionId(v string)`

SetAutoPolicyWaiverExclusionId sets AutoPolicyWaiverExclusionId field to given value.

### HasAutoPolicyWaiverExclusionId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasAutoPolicyWaiverExclusionId() bool`

HasAutoPolicyWaiverExclusionId returns a boolean if a field has been set.

### GetAutoPolicyWaiverId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetAutoPolicyWaiverId() string`

GetAutoPolicyWaiverId returns the AutoPolicyWaiverId field if non-nil, zero value otherwise.

### GetAutoPolicyWaiverIdOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetAutoPolicyWaiverIdOk() (*string, bool)`

GetAutoPolicyWaiverIdOk returns a tuple with the AutoPolicyWaiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPolicyWaiverId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetAutoPolicyWaiverId(v string)`

SetAutoPolicyWaiverId sets AutoPolicyWaiverId field to given value.

### HasAutoPolicyWaiverId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasAutoPolicyWaiverId() bool`

HasAutoPolicyWaiverId returns a boolean if a field has been set.

### GetComponentDisplayName

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetComponentDisplayName() string`

GetComponentDisplayName returns the ComponentDisplayName field if non-nil, zero value otherwise.

### GetComponentDisplayNameOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetComponentDisplayNameOk() (*string, bool)`

GetComponentDisplayNameOk returns a tuple with the ComponentDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentDisplayName

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetComponentDisplayName(v string)`

SetComponentDisplayName sets ComponentDisplayName field to given value.

### HasComponentDisplayName

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasComponentDisplayName() bool`

HasComponentDisplayName returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetComponentIdentifier() ComponentIdentifier`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetComponentIdentifierOk() (*ComponentIdentifier, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetComponentIdentifier(v ComponentIdentifier)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetComponentMatchStrategy

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetComponentMatchStrategy() string`

GetComponentMatchStrategy returns the ComponentMatchStrategy field if non-nil, zero value otherwise.

### GetComponentMatchStrategyOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetComponentMatchStrategyOk() (*string, bool)`

GetComponentMatchStrategyOk returns a tuple with the ComponentMatchStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentMatchStrategy

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetComponentMatchStrategy(v string)`

SetComponentMatchStrategy sets ComponentMatchStrategy field to given value.

### HasComponentMatchStrategy

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasComponentMatchStrategy() bool`

HasComponentMatchStrategy returns a boolean if a field has been set.

### GetConstraintFacts

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetConstraintFacts() []ConstraintFact`

GetConstraintFacts returns the ConstraintFacts field if non-nil, zero value otherwise.

### GetConstraintFactsOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetConstraintFactsOk() (*[]ConstraintFact, bool)`

GetConstraintFactsOk returns a tuple with the ConstraintFacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintFacts

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetConstraintFacts(v []ConstraintFact)`

SetConstraintFacts sets ConstraintFacts field to given value.

### HasConstraintFacts

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasConstraintFacts() bool`

HasConstraintFacts returns a boolean if a field has been set.

### GetCreateTime

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCreatorName

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetHash

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetOwnerId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerName

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetOwnerPublicId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetOwnerPublicId() string`

GetOwnerPublicId returns the OwnerPublicId field if non-nil, zero value otherwise.

### GetOwnerPublicIdOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetOwnerPublicIdOk() (*string, bool)`

GetOwnerPublicIdOk returns a tuple with the OwnerPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPublicId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetOwnerPublicId(v string)`

SetOwnerPublicId sets OwnerPublicId field to given value.

### HasOwnerPublicId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasOwnerPublicId() bool`

HasOwnerPublicId returns a boolean if a field has been set.

### GetOwnerType

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetScanId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.

### GetVulnerabilityIdentifiers

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetVulnerabilityIdentifiers() string`

GetVulnerabilityIdentifiers returns the VulnerabilityIdentifiers field if non-nil, zero value otherwise.

### GetVulnerabilityIdentifiersOk

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) GetVulnerabilityIdentifiersOk() (*string, bool)`

GetVulnerabilityIdentifiersOk returns a tuple with the VulnerabilityIdentifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityIdentifiers

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) SetVulnerabilityIdentifiers(v string)`

SetVulnerabilityIdentifiers sets VulnerabilityIdentifiers field to given value.

### HasVulnerabilityIdentifiers

`func (o *ApiAutoPolicyWaiverExclusionResponseDTO) HasVulnerabilityIdentifiers() bool`

HasVulnerabilityIdentifiers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


