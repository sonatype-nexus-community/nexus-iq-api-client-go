# ApiHostedRepositoryComponentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationPublicId** | Pointer to **string** |  | [optional] 
**ComponentCount** | Pointer to **int32** |  | [optional] 
**ComponentIdentifier** | Pointer to [**ComponentIdentifier**](ComponentIdentifier.md) |  | [optional] 
**CriticalViolationCount** | Pointer to **int32** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdentifiedName** | Pointer to **string** |  | [optional] 
**LastEvaluationTime** | Pointer to **int64** |  | [optional] 
**MatchStateId** | Pointer to **string** |  | [optional] 
**MaxThreatLevel** | Pointer to **int32** |  | [optional] 
**ModerateViolationCount** | Pointer to **int32** |  | [optional] 
**Pathname** | Pointer to **string** |  | [optional] 
**Quarantined** | Pointer to **bool** |  | [optional] 
**ScanId** | Pointer to **string** |  | [optional] 
**SevereViolationCount** | Pointer to **int32** |  | [optional] 
**StageTypeId** | Pointer to **string** |  | [optional] 
**ViolationCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiHostedRepositoryComponentDTO

`func NewApiHostedRepositoryComponentDTO() *ApiHostedRepositoryComponentDTO`

NewApiHostedRepositoryComponentDTO instantiates a new ApiHostedRepositoryComponentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHostedRepositoryComponentDTOWithDefaults

`func NewApiHostedRepositoryComponentDTOWithDefaults() *ApiHostedRepositoryComponentDTO`

NewApiHostedRepositoryComponentDTOWithDefaults instantiates a new ApiHostedRepositoryComponentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationPublicId

`func (o *ApiHostedRepositoryComponentDTO) GetApplicationPublicId() string`

GetApplicationPublicId returns the ApplicationPublicId field if non-nil, zero value otherwise.

### GetApplicationPublicIdOk

`func (o *ApiHostedRepositoryComponentDTO) GetApplicationPublicIdOk() (*string, bool)`

GetApplicationPublicIdOk returns a tuple with the ApplicationPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPublicId

`func (o *ApiHostedRepositoryComponentDTO) SetApplicationPublicId(v string)`

SetApplicationPublicId sets ApplicationPublicId field to given value.

### HasApplicationPublicId

`func (o *ApiHostedRepositoryComponentDTO) HasApplicationPublicId() bool`

HasApplicationPublicId returns a boolean if a field has been set.

### GetComponentCount

`func (o *ApiHostedRepositoryComponentDTO) GetComponentCount() int32`

GetComponentCount returns the ComponentCount field if non-nil, zero value otherwise.

### GetComponentCountOk

`func (o *ApiHostedRepositoryComponentDTO) GetComponentCountOk() (*int32, bool)`

GetComponentCountOk returns a tuple with the ComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCount

`func (o *ApiHostedRepositoryComponentDTO) SetComponentCount(v int32)`

SetComponentCount sets ComponentCount field to given value.

### HasComponentCount

`func (o *ApiHostedRepositoryComponentDTO) HasComponentCount() bool`

HasComponentCount returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *ApiHostedRepositoryComponentDTO) GetComponentIdentifier() ComponentIdentifier`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiHostedRepositoryComponentDTO) GetComponentIdentifierOk() (*ComponentIdentifier, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiHostedRepositoryComponentDTO) SetComponentIdentifier(v ComponentIdentifier)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiHostedRepositoryComponentDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetCriticalViolationCount

`func (o *ApiHostedRepositoryComponentDTO) GetCriticalViolationCount() int32`

GetCriticalViolationCount returns the CriticalViolationCount field if non-nil, zero value otherwise.

### GetCriticalViolationCountOk

`func (o *ApiHostedRepositoryComponentDTO) GetCriticalViolationCountOk() (*int32, bool)`

GetCriticalViolationCountOk returns a tuple with the CriticalViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalViolationCount

`func (o *ApiHostedRepositoryComponentDTO) SetCriticalViolationCount(v int32)`

SetCriticalViolationCount sets CriticalViolationCount field to given value.

### HasCriticalViolationCount

`func (o *ApiHostedRepositoryComponentDTO) HasCriticalViolationCount() bool`

HasCriticalViolationCount returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiHostedRepositoryComponentDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiHostedRepositoryComponentDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiHostedRepositoryComponentDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiHostedRepositoryComponentDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHash

`func (o *ApiHostedRepositoryComponentDTO) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiHostedRepositoryComponentDTO) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiHostedRepositoryComponentDTO) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiHostedRepositoryComponentDTO) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetId

`func (o *ApiHostedRepositoryComponentDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiHostedRepositoryComponentDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiHostedRepositoryComponentDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiHostedRepositoryComponentDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifiedName

`func (o *ApiHostedRepositoryComponentDTO) GetIdentifiedName() string`

GetIdentifiedName returns the IdentifiedName field if non-nil, zero value otherwise.

### GetIdentifiedNameOk

`func (o *ApiHostedRepositoryComponentDTO) GetIdentifiedNameOk() (*string, bool)`

GetIdentifiedNameOk returns a tuple with the IdentifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiedName

`func (o *ApiHostedRepositoryComponentDTO) SetIdentifiedName(v string)`

SetIdentifiedName sets IdentifiedName field to given value.

### HasIdentifiedName

`func (o *ApiHostedRepositoryComponentDTO) HasIdentifiedName() bool`

HasIdentifiedName returns a boolean if a field has been set.

### GetLastEvaluationTime

`func (o *ApiHostedRepositoryComponentDTO) GetLastEvaluationTime() int64`

GetLastEvaluationTime returns the LastEvaluationTime field if non-nil, zero value otherwise.

### GetLastEvaluationTimeOk

`func (o *ApiHostedRepositoryComponentDTO) GetLastEvaluationTimeOk() (*int64, bool)`

GetLastEvaluationTimeOk returns a tuple with the LastEvaluationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvaluationTime

`func (o *ApiHostedRepositoryComponentDTO) SetLastEvaluationTime(v int64)`

SetLastEvaluationTime sets LastEvaluationTime field to given value.

### HasLastEvaluationTime

`func (o *ApiHostedRepositoryComponentDTO) HasLastEvaluationTime() bool`

HasLastEvaluationTime returns a boolean if a field has been set.

### GetMatchStateId

`func (o *ApiHostedRepositoryComponentDTO) GetMatchStateId() string`

GetMatchStateId returns the MatchStateId field if non-nil, zero value otherwise.

### GetMatchStateIdOk

`func (o *ApiHostedRepositoryComponentDTO) GetMatchStateIdOk() (*string, bool)`

GetMatchStateIdOk returns a tuple with the MatchStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchStateId

`func (o *ApiHostedRepositoryComponentDTO) SetMatchStateId(v string)`

SetMatchStateId sets MatchStateId field to given value.

### HasMatchStateId

`func (o *ApiHostedRepositoryComponentDTO) HasMatchStateId() bool`

HasMatchStateId returns a boolean if a field has been set.

### GetMaxThreatLevel

`func (o *ApiHostedRepositoryComponentDTO) GetMaxThreatLevel() int32`

GetMaxThreatLevel returns the MaxThreatLevel field if non-nil, zero value otherwise.

### GetMaxThreatLevelOk

`func (o *ApiHostedRepositoryComponentDTO) GetMaxThreatLevelOk() (*int32, bool)`

GetMaxThreatLevelOk returns a tuple with the MaxThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThreatLevel

`func (o *ApiHostedRepositoryComponentDTO) SetMaxThreatLevel(v int32)`

SetMaxThreatLevel sets MaxThreatLevel field to given value.

### HasMaxThreatLevel

`func (o *ApiHostedRepositoryComponentDTO) HasMaxThreatLevel() bool`

HasMaxThreatLevel returns a boolean if a field has been set.

### GetModerateViolationCount

`func (o *ApiHostedRepositoryComponentDTO) GetModerateViolationCount() int32`

GetModerateViolationCount returns the ModerateViolationCount field if non-nil, zero value otherwise.

### GetModerateViolationCountOk

`func (o *ApiHostedRepositoryComponentDTO) GetModerateViolationCountOk() (*int32, bool)`

GetModerateViolationCountOk returns a tuple with the ModerateViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerateViolationCount

`func (o *ApiHostedRepositoryComponentDTO) SetModerateViolationCount(v int32)`

SetModerateViolationCount sets ModerateViolationCount field to given value.

### HasModerateViolationCount

`func (o *ApiHostedRepositoryComponentDTO) HasModerateViolationCount() bool`

HasModerateViolationCount returns a boolean if a field has been set.

### GetPathname

`func (o *ApiHostedRepositoryComponentDTO) GetPathname() string`

GetPathname returns the Pathname field if non-nil, zero value otherwise.

### GetPathnameOk

`func (o *ApiHostedRepositoryComponentDTO) GetPathnameOk() (*string, bool)`

GetPathnameOk returns a tuple with the Pathname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathname

`func (o *ApiHostedRepositoryComponentDTO) SetPathname(v string)`

SetPathname sets Pathname field to given value.

### HasPathname

`func (o *ApiHostedRepositoryComponentDTO) HasPathname() bool`

HasPathname returns a boolean if a field has been set.

### GetQuarantined

`func (o *ApiHostedRepositoryComponentDTO) GetQuarantined() bool`

GetQuarantined returns the Quarantined field if non-nil, zero value otherwise.

### GetQuarantinedOk

`func (o *ApiHostedRepositoryComponentDTO) GetQuarantinedOk() (*bool, bool)`

GetQuarantinedOk returns a tuple with the Quarantined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantined

`func (o *ApiHostedRepositoryComponentDTO) SetQuarantined(v bool)`

SetQuarantined sets Quarantined field to given value.

### HasQuarantined

`func (o *ApiHostedRepositoryComponentDTO) HasQuarantined() bool`

HasQuarantined returns a boolean if a field has been set.

### GetScanId

`func (o *ApiHostedRepositoryComponentDTO) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ApiHostedRepositoryComponentDTO) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ApiHostedRepositoryComponentDTO) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *ApiHostedRepositoryComponentDTO) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetSevereViolationCount

`func (o *ApiHostedRepositoryComponentDTO) GetSevereViolationCount() int32`

GetSevereViolationCount returns the SevereViolationCount field if non-nil, zero value otherwise.

### GetSevereViolationCountOk

`func (o *ApiHostedRepositoryComponentDTO) GetSevereViolationCountOk() (*int32, bool)`

GetSevereViolationCountOk returns a tuple with the SevereViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevereViolationCount

`func (o *ApiHostedRepositoryComponentDTO) SetSevereViolationCount(v int32)`

SetSevereViolationCount sets SevereViolationCount field to given value.

### HasSevereViolationCount

`func (o *ApiHostedRepositoryComponentDTO) HasSevereViolationCount() bool`

HasSevereViolationCount returns a boolean if a field has been set.

### GetStageTypeId

`func (o *ApiHostedRepositoryComponentDTO) GetStageTypeId() string`

GetStageTypeId returns the StageTypeId field if non-nil, zero value otherwise.

### GetStageTypeIdOk

`func (o *ApiHostedRepositoryComponentDTO) GetStageTypeIdOk() (*string, bool)`

GetStageTypeIdOk returns a tuple with the StageTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageTypeId

`func (o *ApiHostedRepositoryComponentDTO) SetStageTypeId(v string)`

SetStageTypeId sets StageTypeId field to given value.

### HasStageTypeId

`func (o *ApiHostedRepositoryComponentDTO) HasStageTypeId() bool`

HasStageTypeId returns a boolean if a field has been set.

### GetViolationCount

`func (o *ApiHostedRepositoryComponentDTO) GetViolationCount() int32`

GetViolationCount returns the ViolationCount field if non-nil, zero value otherwise.

### GetViolationCountOk

`func (o *ApiHostedRepositoryComponentDTO) GetViolationCountOk() (*int32, bool)`

GetViolationCountOk returns a tuple with the ViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCount

`func (o *ApiHostedRepositoryComponentDTO) SetViolationCount(v int32)`

SetViolationCount sets ViolationCount field to given value.

### HasViolationCount

`func (o *ApiHostedRepositoryComponentDTO) HasViolationCount() bool`

HasViolationCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


