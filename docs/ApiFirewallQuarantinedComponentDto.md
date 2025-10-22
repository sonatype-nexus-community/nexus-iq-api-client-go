# ApiFirewallQuarantinedComponentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentIdentifier** | Pointer to [**ComponentIdentifier**](ComponentIdentifier.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**MatchState** | Pointer to **string** |  | [optional] 
**Pathname** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**QuarantineDate** | Pointer to **string** |  | [optional] 
**Quarantined** | Pointer to **bool** |  | [optional] 
**RepositoryId** | Pointer to **string** |  | [optional] 
**RepositoryName** | Pointer to **string** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiFirewallQuarantinedComponentDto

`func NewApiFirewallQuarantinedComponentDto() *ApiFirewallQuarantinedComponentDto`

NewApiFirewallQuarantinedComponentDto instantiates a new ApiFirewallQuarantinedComponentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFirewallQuarantinedComponentDtoWithDefaults

`func NewApiFirewallQuarantinedComponentDtoWithDefaults() *ApiFirewallQuarantinedComponentDto`

NewApiFirewallQuarantinedComponentDtoWithDefaults instantiates a new ApiFirewallQuarantinedComponentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentIdentifier

`func (o *ApiFirewallQuarantinedComponentDto) GetComponentIdentifier() ComponentIdentifier`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiFirewallQuarantinedComponentDto) GetComponentIdentifierOk() (*ComponentIdentifier, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiFirewallQuarantinedComponentDto) SetComponentIdentifier(v ComponentIdentifier)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiFirewallQuarantinedComponentDto) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiFirewallQuarantinedComponentDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiFirewallQuarantinedComponentDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiFirewallQuarantinedComponentDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiFirewallQuarantinedComponentDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHash

`func (o *ApiFirewallQuarantinedComponentDto) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiFirewallQuarantinedComponentDto) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiFirewallQuarantinedComponentDto) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiFirewallQuarantinedComponentDto) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetMatchState

`func (o *ApiFirewallQuarantinedComponentDto) GetMatchState() string`

GetMatchState returns the MatchState field if non-nil, zero value otherwise.

### GetMatchStateOk

`func (o *ApiFirewallQuarantinedComponentDto) GetMatchStateOk() (*string, bool)`

GetMatchStateOk returns a tuple with the MatchState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchState

`func (o *ApiFirewallQuarantinedComponentDto) SetMatchState(v string)`

SetMatchState sets MatchState field to given value.

### HasMatchState

`func (o *ApiFirewallQuarantinedComponentDto) HasMatchState() bool`

HasMatchState returns a boolean if a field has been set.

### GetPathname

`func (o *ApiFirewallQuarantinedComponentDto) GetPathname() string`

GetPathname returns the Pathname field if non-nil, zero value otherwise.

### GetPathnameOk

`func (o *ApiFirewallQuarantinedComponentDto) GetPathnameOk() (*string, bool)`

GetPathnameOk returns a tuple with the Pathname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathname

`func (o *ApiFirewallQuarantinedComponentDto) SetPathname(v string)`

SetPathname sets Pathname field to given value.

### HasPathname

`func (o *ApiFirewallQuarantinedComponentDto) HasPathname() bool`

HasPathname returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiFirewallQuarantinedComponentDto) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiFirewallQuarantinedComponentDto) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiFirewallQuarantinedComponentDto) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiFirewallQuarantinedComponentDto) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetQuarantineDate

`func (o *ApiFirewallQuarantinedComponentDto) GetQuarantineDate() string`

GetQuarantineDate returns the QuarantineDate field if non-nil, zero value otherwise.

### GetQuarantineDateOk

`func (o *ApiFirewallQuarantinedComponentDto) GetQuarantineDateOk() (*string, bool)`

GetQuarantineDateOk returns a tuple with the QuarantineDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineDate

`func (o *ApiFirewallQuarantinedComponentDto) SetQuarantineDate(v string)`

SetQuarantineDate sets QuarantineDate field to given value.

### HasQuarantineDate

`func (o *ApiFirewallQuarantinedComponentDto) HasQuarantineDate() bool`

HasQuarantineDate returns a boolean if a field has been set.

### GetQuarantined

`func (o *ApiFirewallQuarantinedComponentDto) GetQuarantined() bool`

GetQuarantined returns the Quarantined field if non-nil, zero value otherwise.

### GetQuarantinedOk

`func (o *ApiFirewallQuarantinedComponentDto) GetQuarantinedOk() (*bool, bool)`

GetQuarantinedOk returns a tuple with the Quarantined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantined

`func (o *ApiFirewallQuarantinedComponentDto) SetQuarantined(v bool)`

SetQuarantined sets Quarantined field to given value.

### HasQuarantined

`func (o *ApiFirewallQuarantinedComponentDto) HasQuarantined() bool`

HasQuarantined returns a boolean if a field has been set.

### GetRepositoryId

`func (o *ApiFirewallQuarantinedComponentDto) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *ApiFirewallQuarantinedComponentDto) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *ApiFirewallQuarantinedComponentDto) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *ApiFirewallQuarantinedComponentDto) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetRepositoryName

`func (o *ApiFirewallQuarantinedComponentDto) GetRepositoryName() string`

GetRepositoryName returns the RepositoryName field if non-nil, zero value otherwise.

### GetRepositoryNameOk

`func (o *ApiFirewallQuarantinedComponentDto) GetRepositoryNameOk() (*string, bool)`

GetRepositoryNameOk returns a tuple with the RepositoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryName

`func (o *ApiFirewallQuarantinedComponentDto) SetRepositoryName(v string)`

SetRepositoryName sets RepositoryName field to given value.

### HasRepositoryName

`func (o *ApiFirewallQuarantinedComponentDto) HasRepositoryName() bool`

HasRepositoryName returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiFirewallQuarantinedComponentDto) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiFirewallQuarantinedComponentDto) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiFirewallQuarantinedComponentDto) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiFirewallQuarantinedComponentDto) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


