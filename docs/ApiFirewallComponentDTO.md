# ApiFirewallComponentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentIdentifier** | Pointer to [**ComponentIdentifier**](ComponentIdentifier.md) |  | [optional] 
**DateCleared** | Pointer to **time.Time** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**MatchState** | Pointer to **string** |  | [optional] 
**Pathname** | Pointer to **string** |  | [optional] 
**QuarantineDate** | Pointer to **time.Time** |  | [optional] 
**QuarantinePolicyViolations** | Pointer to [**[]ApiPolicyViolationDTOV2**](ApiPolicyViolationDTOV2.md) |  | [optional] 
**Quarantined** | Pointer to **bool** |  | [optional] 
**Repository** | Pointer to **string** |  | [optional] 
**RepositoryId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiFirewallComponentDTO

`func NewApiFirewallComponentDTO() *ApiFirewallComponentDTO`

NewApiFirewallComponentDTO instantiates a new ApiFirewallComponentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFirewallComponentDTOWithDefaults

`func NewApiFirewallComponentDTOWithDefaults() *ApiFirewallComponentDTO`

NewApiFirewallComponentDTOWithDefaults instantiates a new ApiFirewallComponentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentIdentifier

`func (o *ApiFirewallComponentDTO) GetComponentIdentifier() ComponentIdentifier`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiFirewallComponentDTO) GetComponentIdentifierOk() (*ComponentIdentifier, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiFirewallComponentDTO) SetComponentIdentifier(v ComponentIdentifier)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiFirewallComponentDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetDateCleared

`func (o *ApiFirewallComponentDTO) GetDateCleared() time.Time`

GetDateCleared returns the DateCleared field if non-nil, zero value otherwise.

### GetDateClearedOk

`func (o *ApiFirewallComponentDTO) GetDateClearedOk() (*time.Time, bool)`

GetDateClearedOk returns a tuple with the DateCleared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCleared

`func (o *ApiFirewallComponentDTO) SetDateCleared(v time.Time)`

SetDateCleared sets DateCleared field to given value.

### HasDateCleared

`func (o *ApiFirewallComponentDTO) HasDateCleared() bool`

HasDateCleared returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiFirewallComponentDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiFirewallComponentDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiFirewallComponentDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiFirewallComponentDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHash

`func (o *ApiFirewallComponentDTO) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiFirewallComponentDTO) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiFirewallComponentDTO) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiFirewallComponentDTO) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetMatchState

`func (o *ApiFirewallComponentDTO) GetMatchState() string`

GetMatchState returns the MatchState field if non-nil, zero value otherwise.

### GetMatchStateOk

`func (o *ApiFirewallComponentDTO) GetMatchStateOk() (*string, bool)`

GetMatchStateOk returns a tuple with the MatchState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchState

`func (o *ApiFirewallComponentDTO) SetMatchState(v string)`

SetMatchState sets MatchState field to given value.

### HasMatchState

`func (o *ApiFirewallComponentDTO) HasMatchState() bool`

HasMatchState returns a boolean if a field has been set.

### GetPathname

`func (o *ApiFirewallComponentDTO) GetPathname() string`

GetPathname returns the Pathname field if non-nil, zero value otherwise.

### GetPathnameOk

`func (o *ApiFirewallComponentDTO) GetPathnameOk() (*string, bool)`

GetPathnameOk returns a tuple with the Pathname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathname

`func (o *ApiFirewallComponentDTO) SetPathname(v string)`

SetPathname sets Pathname field to given value.

### HasPathname

`func (o *ApiFirewallComponentDTO) HasPathname() bool`

HasPathname returns a boolean if a field has been set.

### GetQuarantineDate

`func (o *ApiFirewallComponentDTO) GetQuarantineDate() time.Time`

GetQuarantineDate returns the QuarantineDate field if non-nil, zero value otherwise.

### GetQuarantineDateOk

`func (o *ApiFirewallComponentDTO) GetQuarantineDateOk() (*time.Time, bool)`

GetQuarantineDateOk returns a tuple with the QuarantineDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineDate

`func (o *ApiFirewallComponentDTO) SetQuarantineDate(v time.Time)`

SetQuarantineDate sets QuarantineDate field to given value.

### HasQuarantineDate

`func (o *ApiFirewallComponentDTO) HasQuarantineDate() bool`

HasQuarantineDate returns a boolean if a field has been set.

### GetQuarantinePolicyViolations

`func (o *ApiFirewallComponentDTO) GetQuarantinePolicyViolations() []ApiPolicyViolationDTOV2`

GetQuarantinePolicyViolations returns the QuarantinePolicyViolations field if non-nil, zero value otherwise.

### GetQuarantinePolicyViolationsOk

`func (o *ApiFirewallComponentDTO) GetQuarantinePolicyViolationsOk() (*[]ApiPolicyViolationDTOV2, bool)`

GetQuarantinePolicyViolationsOk returns a tuple with the QuarantinePolicyViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantinePolicyViolations

`func (o *ApiFirewallComponentDTO) SetQuarantinePolicyViolations(v []ApiPolicyViolationDTOV2)`

SetQuarantinePolicyViolations sets QuarantinePolicyViolations field to given value.

### HasQuarantinePolicyViolations

`func (o *ApiFirewallComponentDTO) HasQuarantinePolicyViolations() bool`

HasQuarantinePolicyViolations returns a boolean if a field has been set.

### GetQuarantined

`func (o *ApiFirewallComponentDTO) GetQuarantined() bool`

GetQuarantined returns the Quarantined field if non-nil, zero value otherwise.

### GetQuarantinedOk

`func (o *ApiFirewallComponentDTO) GetQuarantinedOk() (*bool, bool)`

GetQuarantinedOk returns a tuple with the Quarantined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantined

`func (o *ApiFirewallComponentDTO) SetQuarantined(v bool)`

SetQuarantined sets Quarantined field to given value.

### HasQuarantined

`func (o *ApiFirewallComponentDTO) HasQuarantined() bool`

HasQuarantined returns a boolean if a field has been set.

### GetRepository

`func (o *ApiFirewallComponentDTO) GetRepository() string`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ApiFirewallComponentDTO) GetRepositoryOk() (*string, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ApiFirewallComponentDTO) SetRepository(v string)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ApiFirewallComponentDTO) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetRepositoryId

`func (o *ApiFirewallComponentDTO) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *ApiFirewallComponentDTO) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *ApiFirewallComponentDTO) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *ApiFirewallComponentDTO) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


