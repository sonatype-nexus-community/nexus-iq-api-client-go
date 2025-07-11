# ApiCrossStageViolationDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationName** | Pointer to **string** |  | [optional] 
**ApplicationPublicId** | Pointer to **string** |  | [optional] 
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**ConstraintViolations** | Pointer to [**[]ApiConstraintViolationDTO**](ApiConstraintViolationDTO.md) |  | [optional] 
**DisplayName** | Pointer to [**ComponentDisplayName**](ComponentDisplayName.md) |  | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**FixTime** | Pointer to **time.Time** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**LegacyViolationTime** | Pointer to **time.Time** |  | [optional] 
**OpenTime** | Pointer to **time.Time** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyOwner** | Pointer to [**PolicyOwner**](PolicyOwner.md) |  | [optional] 
**PolicyThreatCategory** | Pointer to **string** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**ReachabilityStatus** | Pointer to **string** |  | [optional] 
**StageData** | Pointer to [**map[string]StageData**](StageData.md) |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 
**WaiveTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewApiCrossStageViolationDTOV2

`func NewApiCrossStageViolationDTOV2() *ApiCrossStageViolationDTOV2`

NewApiCrossStageViolationDTOV2 instantiates a new ApiCrossStageViolationDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCrossStageViolationDTOV2WithDefaults

`func NewApiCrossStageViolationDTOV2WithDefaults() *ApiCrossStageViolationDTOV2`

NewApiCrossStageViolationDTOV2WithDefaults instantiates a new ApiCrossStageViolationDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationName

`func (o *ApiCrossStageViolationDTOV2) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *ApiCrossStageViolationDTOV2) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *ApiCrossStageViolationDTOV2) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *ApiCrossStageViolationDTOV2) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationPublicId

`func (o *ApiCrossStageViolationDTOV2) GetApplicationPublicId() string`

GetApplicationPublicId returns the ApplicationPublicId field if non-nil, zero value otherwise.

### GetApplicationPublicIdOk

`func (o *ApiCrossStageViolationDTOV2) GetApplicationPublicIdOk() (*string, bool)`

GetApplicationPublicIdOk returns a tuple with the ApplicationPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPublicId

`func (o *ApiCrossStageViolationDTOV2) SetApplicationPublicId(v string)`

SetApplicationPublicId sets ApplicationPublicId field to given value.

### HasApplicationPublicId

`func (o *ApiCrossStageViolationDTOV2) HasApplicationPublicId() bool`

HasApplicationPublicId returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *ApiCrossStageViolationDTOV2) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiCrossStageViolationDTOV2) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiCrossStageViolationDTOV2) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiCrossStageViolationDTOV2) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetConstraintViolations

`func (o *ApiCrossStageViolationDTOV2) GetConstraintViolations() []ApiConstraintViolationDTO`

GetConstraintViolations returns the ConstraintViolations field if non-nil, zero value otherwise.

### GetConstraintViolationsOk

`func (o *ApiCrossStageViolationDTOV2) GetConstraintViolationsOk() (*[]ApiConstraintViolationDTO, bool)`

GetConstraintViolationsOk returns a tuple with the ConstraintViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintViolations

`func (o *ApiCrossStageViolationDTOV2) SetConstraintViolations(v []ApiConstraintViolationDTO)`

SetConstraintViolations sets ConstraintViolations field to given value.

### HasConstraintViolations

`func (o *ApiCrossStageViolationDTOV2) HasConstraintViolations() bool`

HasConstraintViolations returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiCrossStageViolationDTOV2) GetDisplayName() ComponentDisplayName`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiCrossStageViolationDTOV2) GetDisplayNameOk() (*ComponentDisplayName, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiCrossStageViolationDTOV2) SetDisplayName(v ComponentDisplayName)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiCrossStageViolationDTOV2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFilename

`func (o *ApiCrossStageViolationDTOV2) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ApiCrossStageViolationDTOV2) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ApiCrossStageViolationDTOV2) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *ApiCrossStageViolationDTOV2) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetFixTime

`func (o *ApiCrossStageViolationDTOV2) GetFixTime() time.Time`

GetFixTime returns the FixTime field if non-nil, zero value otherwise.

### GetFixTimeOk

`func (o *ApiCrossStageViolationDTOV2) GetFixTimeOk() (*time.Time, bool)`

GetFixTimeOk returns a tuple with the FixTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixTime

`func (o *ApiCrossStageViolationDTOV2) SetFixTime(v time.Time)`

SetFixTime sets FixTime field to given value.

### HasFixTime

`func (o *ApiCrossStageViolationDTOV2) HasFixTime() bool`

HasFixTime returns a boolean if a field has been set.

### GetHash

`func (o *ApiCrossStageViolationDTOV2) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiCrossStageViolationDTOV2) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiCrossStageViolationDTOV2) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiCrossStageViolationDTOV2) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetLegacyViolationTime

`func (o *ApiCrossStageViolationDTOV2) GetLegacyViolationTime() time.Time`

GetLegacyViolationTime returns the LegacyViolationTime field if non-nil, zero value otherwise.

### GetLegacyViolationTimeOk

`func (o *ApiCrossStageViolationDTOV2) GetLegacyViolationTimeOk() (*time.Time, bool)`

GetLegacyViolationTimeOk returns a tuple with the LegacyViolationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyViolationTime

`func (o *ApiCrossStageViolationDTOV2) SetLegacyViolationTime(v time.Time)`

SetLegacyViolationTime sets LegacyViolationTime field to given value.

### HasLegacyViolationTime

`func (o *ApiCrossStageViolationDTOV2) HasLegacyViolationTime() bool`

HasLegacyViolationTime returns a boolean if a field has been set.

### GetOpenTime

`func (o *ApiCrossStageViolationDTOV2) GetOpenTime() time.Time`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *ApiCrossStageViolationDTOV2) GetOpenTimeOk() (*time.Time, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *ApiCrossStageViolationDTOV2) SetOpenTime(v time.Time)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *ApiCrossStageViolationDTOV2) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetOrganizationName

`func (o *ApiCrossStageViolationDTOV2) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *ApiCrossStageViolationDTOV2) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *ApiCrossStageViolationDTOV2) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *ApiCrossStageViolationDTOV2) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApiCrossStageViolationDTOV2) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApiCrossStageViolationDTOV2) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApiCrossStageViolationDTOV2) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApiCrossStageViolationDTOV2) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiCrossStageViolationDTOV2) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiCrossStageViolationDTOV2) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiCrossStageViolationDTOV2) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiCrossStageViolationDTOV2) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyOwner

`func (o *ApiCrossStageViolationDTOV2) GetPolicyOwner() PolicyOwner`

GetPolicyOwner returns the PolicyOwner field if non-nil, zero value otherwise.

### GetPolicyOwnerOk

`func (o *ApiCrossStageViolationDTOV2) GetPolicyOwnerOk() (*PolicyOwner, bool)`

GetPolicyOwnerOk returns a tuple with the PolicyOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyOwner

`func (o *ApiCrossStageViolationDTOV2) SetPolicyOwner(v PolicyOwner)`

SetPolicyOwner sets PolicyOwner field to given value.

### HasPolicyOwner

`func (o *ApiCrossStageViolationDTOV2) HasPolicyOwner() bool`

HasPolicyOwner returns a boolean if a field has been set.

### GetPolicyThreatCategory

`func (o *ApiCrossStageViolationDTOV2) GetPolicyThreatCategory() string`

GetPolicyThreatCategory returns the PolicyThreatCategory field if non-nil, zero value otherwise.

### GetPolicyThreatCategoryOk

`func (o *ApiCrossStageViolationDTOV2) GetPolicyThreatCategoryOk() (*string, bool)`

GetPolicyThreatCategoryOk returns a tuple with the PolicyThreatCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyThreatCategory

`func (o *ApiCrossStageViolationDTOV2) SetPolicyThreatCategory(v string)`

SetPolicyThreatCategory sets PolicyThreatCategory field to given value.

### HasPolicyThreatCategory

`func (o *ApiCrossStageViolationDTOV2) HasPolicyThreatCategory() bool`

HasPolicyThreatCategory returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *ApiCrossStageViolationDTOV2) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *ApiCrossStageViolationDTOV2) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *ApiCrossStageViolationDTOV2) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *ApiCrossStageViolationDTOV2) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetReachabilityStatus

`func (o *ApiCrossStageViolationDTOV2) GetReachabilityStatus() string`

GetReachabilityStatus returns the ReachabilityStatus field if non-nil, zero value otherwise.

### GetReachabilityStatusOk

`func (o *ApiCrossStageViolationDTOV2) GetReachabilityStatusOk() (*string, bool)`

GetReachabilityStatusOk returns a tuple with the ReachabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachabilityStatus

`func (o *ApiCrossStageViolationDTOV2) SetReachabilityStatus(v string)`

SetReachabilityStatus sets ReachabilityStatus field to given value.

### HasReachabilityStatus

`func (o *ApiCrossStageViolationDTOV2) HasReachabilityStatus() bool`

HasReachabilityStatus returns a boolean if a field has been set.

### GetStageData

`func (o *ApiCrossStageViolationDTOV2) GetStageData() map[string]StageData`

GetStageData returns the StageData field if non-nil, zero value otherwise.

### GetStageDataOk

`func (o *ApiCrossStageViolationDTOV2) GetStageDataOk() (*map[string]StageData, bool)`

GetStageDataOk returns a tuple with the StageData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageData

`func (o *ApiCrossStageViolationDTOV2) SetStageData(v map[string]StageData)`

SetStageData sets StageData field to given value.

### HasStageData

`func (o *ApiCrossStageViolationDTOV2) HasStageData() bool`

HasStageData returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiCrossStageViolationDTOV2) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiCrossStageViolationDTOV2) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiCrossStageViolationDTOV2) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiCrossStageViolationDTOV2) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.

### GetWaiveTime

`func (o *ApiCrossStageViolationDTOV2) GetWaiveTime() time.Time`

GetWaiveTime returns the WaiveTime field if non-nil, zero value otherwise.

### GetWaiveTimeOk

`func (o *ApiCrossStageViolationDTOV2) GetWaiveTimeOk() (*time.Time, bool)`

GetWaiveTimeOk returns a tuple with the WaiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiveTime

`func (o *ApiCrossStageViolationDTOV2) SetWaiveTime(v time.Time)`

SetWaiveTime sets WaiveTime field to given value.

### HasWaiveTime

`func (o *ApiCrossStageViolationDTOV2) HasWaiveTime() bool`

HasWaiveTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


