# SearchResultItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationCategoryColor** | Pointer to **string** |  | [optional] 
**ApplicationCategoryDescription** | Pointer to **string** |  | [optional] 
**ApplicationCategoryId** | Pointer to **string** |  | [optional] 
**ApplicationCategoryName** | Pointer to **string** |  | [optional] 
**ApplicationCategoryNames** | Pointer to **[]string** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**ApplicationLastEvaluationTimeEpochMs** | Pointer to **int64** |  | [optional] 
**ApplicationName** | Pointer to **string** |  | [optional] 
**ApplicationPublicId** | Pointer to **string** |  | [optional] 
**ApplicationStageSeverityCounts** | Pointer to **[]string** |  | [optional] 
**ApplicationVersion** | Pointer to **string** |  | [optional] 
**ApplicationViolationPolicyTypes** | Pointer to **[]string** |  | [optional] 
**ApplicationViolationStages** | Pointer to **[]string** |  | [optional] 
**ApplicationViolationStates** | Pointer to **[]string** |  | [optional] 
**ComponentEffectiveLicenseId** | Pointer to **string** |  | [optional] 
**ComponentEffectiveLicenseName** | Pointer to **string** |  | [optional] 
**ComponentHash** | Pointer to **string** |  | [optional] 
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**ComponentLabelColor** | Pointer to **string** |  | [optional] 
**ComponentLabelDescription** | Pointer to **string** |  | [optional] 
**ComponentLabelId** | Pointer to **string** |  | [optional] 
**ComponentLabelName** | Pointer to **string** |  | [optional] 
**ComponentLicenseThreatGroupName** | Pointer to **string** |  | [optional] 
**ComponentLicenseThreatLevel** | Pointer to **int32** |  | [optional] 
**ComponentName** | Pointer to **string** |  | [optional] 
**ComponentViolationPolicyTypes** | Pointer to **[]string** |  | [optional] 
**ComponentViolationStates** | Pointer to **[]string** |  | [optional] 
**ItemType** | Pointer to **string** |  | [optional] 
**NoteToReviewer** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**PolicyEvaluationStage** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyThreatCategory** | Pointer to **string** |  | [optional] 
**PolicyThreatLevel** | Pointer to **int32** |  | [optional] 
**PolicyViolationConstraintName** | Pointer to **string** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**PolicyViolationPolicyId** | Pointer to **string** |  | [optional] 
**PolicyViolationPolicyName** | Pointer to **string** |  | [optional] 
**PolicyViolationThreatCategory** | Pointer to **string** |  | [optional] 
**PolicyViolationThreatLevel** | Pointer to **int32** |  | [optional] 
**PolicyViolationWaiverStatus** | Pointer to **string** |  | [optional] 
**PolicyWaiverAuto** | Pointer to **bool** |  | [optional] 
**PolicyWaiverComment** | Pointer to **string** |  | [optional] 
**PolicyWaiverCreatedAt** | Pointer to **string** |  | [optional] 
**PolicyWaiverExpiresAt** | Pointer to **string** |  | [optional] 
**PolicyWaiverExpiryStatus** | Pointer to **string** |  | [optional] 
**PolicyWaiverId** | Pointer to **string** |  | [optional] 
**PolicyWaiverIsAuto** | Pointer to **bool** |  | [optional] 
**PolicyWaiverPolicyId** | Pointer to **string** |  | [optional] 
**PolicyWaiverPolicyName** | Pointer to **string** |  | [optional] 
**PolicyWaiverPolicyType** | Pointer to **string** |  | [optional] 
**PolicyWaiverReason** | Pointer to **string** |  | [optional] 
**PolicyWaiverRequestStatus** | Pointer to **string** |  | [optional] 
**PolicyWaiverScope** | Pointer to **string** |  | [optional] 
**PolicyWaiverScopeOwnerId** | Pointer to **string** |  | [optional] 
**PolicyWaiverScopeOwnerType** | Pointer to **string** |  | [optional] 
**PolicyWaiverThreatLevel** | Pointer to **int32** |  | [optional] 
**PolicyWaiverWaivedBy** | Pointer to **string** |  | [optional] 
**RejectionReason** | Pointer to **string** |  | [optional] 
**ReportId** | Pointer to **string** |  | [optional] 
**RequesterName** | Pointer to **string** |  | [optional] 
**ResultIndex** | Pointer to **int32** |  | [optional] 
**ReviewTime** | Pointer to **string** |  | [optional] 
**ReviewerName** | Pointer to **string** |  | [optional] 
**SbomSpecification** | Pointer to **string** |  | [optional] 
**VulnerabilityDescription** | Pointer to **string** |  | [optional] 
**VulnerabilityFirstSeenEpochMs** | Pointer to **int64** |  | [optional] 
**VulnerabilityId** | Pointer to **string** |  | [optional] 
**VulnerabilitySeverity** | Pointer to **float32** |  | [optional] 
**VulnerabilityStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewSearchResultItemDTO

`func NewSearchResultItemDTO() *SearchResultItemDTO`

NewSearchResultItemDTO instantiates a new SearchResultItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultItemDTOWithDefaults

`func NewSearchResultItemDTOWithDefaults() *SearchResultItemDTO`

NewSearchResultItemDTOWithDefaults instantiates a new SearchResultItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationCategoryColor

`func (o *SearchResultItemDTO) GetApplicationCategoryColor() string`

GetApplicationCategoryColor returns the ApplicationCategoryColor field if non-nil, zero value otherwise.

### GetApplicationCategoryColorOk

`func (o *SearchResultItemDTO) GetApplicationCategoryColorOk() (*string, bool)`

GetApplicationCategoryColorOk returns a tuple with the ApplicationCategoryColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategoryColor

`func (o *SearchResultItemDTO) SetApplicationCategoryColor(v string)`

SetApplicationCategoryColor sets ApplicationCategoryColor field to given value.

### HasApplicationCategoryColor

`func (o *SearchResultItemDTO) HasApplicationCategoryColor() bool`

HasApplicationCategoryColor returns a boolean if a field has been set.

### GetApplicationCategoryDescription

`func (o *SearchResultItemDTO) GetApplicationCategoryDescription() string`

GetApplicationCategoryDescription returns the ApplicationCategoryDescription field if non-nil, zero value otherwise.

### GetApplicationCategoryDescriptionOk

`func (o *SearchResultItemDTO) GetApplicationCategoryDescriptionOk() (*string, bool)`

GetApplicationCategoryDescriptionOk returns a tuple with the ApplicationCategoryDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategoryDescription

`func (o *SearchResultItemDTO) SetApplicationCategoryDescription(v string)`

SetApplicationCategoryDescription sets ApplicationCategoryDescription field to given value.

### HasApplicationCategoryDescription

`func (o *SearchResultItemDTO) HasApplicationCategoryDescription() bool`

HasApplicationCategoryDescription returns a boolean if a field has been set.

### GetApplicationCategoryId

`func (o *SearchResultItemDTO) GetApplicationCategoryId() string`

GetApplicationCategoryId returns the ApplicationCategoryId field if non-nil, zero value otherwise.

### GetApplicationCategoryIdOk

`func (o *SearchResultItemDTO) GetApplicationCategoryIdOk() (*string, bool)`

GetApplicationCategoryIdOk returns a tuple with the ApplicationCategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategoryId

`func (o *SearchResultItemDTO) SetApplicationCategoryId(v string)`

SetApplicationCategoryId sets ApplicationCategoryId field to given value.

### HasApplicationCategoryId

`func (o *SearchResultItemDTO) HasApplicationCategoryId() bool`

HasApplicationCategoryId returns a boolean if a field has been set.

### GetApplicationCategoryName

`func (o *SearchResultItemDTO) GetApplicationCategoryName() string`

GetApplicationCategoryName returns the ApplicationCategoryName field if non-nil, zero value otherwise.

### GetApplicationCategoryNameOk

`func (o *SearchResultItemDTO) GetApplicationCategoryNameOk() (*string, bool)`

GetApplicationCategoryNameOk returns a tuple with the ApplicationCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategoryName

`func (o *SearchResultItemDTO) SetApplicationCategoryName(v string)`

SetApplicationCategoryName sets ApplicationCategoryName field to given value.

### HasApplicationCategoryName

`func (o *SearchResultItemDTO) HasApplicationCategoryName() bool`

HasApplicationCategoryName returns a boolean if a field has been set.

### GetApplicationCategoryNames

`func (o *SearchResultItemDTO) GetApplicationCategoryNames() []string`

GetApplicationCategoryNames returns the ApplicationCategoryNames field if non-nil, zero value otherwise.

### GetApplicationCategoryNamesOk

`func (o *SearchResultItemDTO) GetApplicationCategoryNamesOk() (*[]string, bool)`

GetApplicationCategoryNamesOk returns a tuple with the ApplicationCategoryNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategoryNames

`func (o *SearchResultItemDTO) SetApplicationCategoryNames(v []string)`

SetApplicationCategoryNames sets ApplicationCategoryNames field to given value.

### HasApplicationCategoryNames

`func (o *SearchResultItemDTO) HasApplicationCategoryNames() bool`

HasApplicationCategoryNames returns a boolean if a field has been set.

### GetApplicationId

`func (o *SearchResultItemDTO) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *SearchResultItemDTO) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *SearchResultItemDTO) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *SearchResultItemDTO) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationLastEvaluationTimeEpochMs

`func (o *SearchResultItemDTO) GetApplicationLastEvaluationTimeEpochMs() int64`

GetApplicationLastEvaluationTimeEpochMs returns the ApplicationLastEvaluationTimeEpochMs field if non-nil, zero value otherwise.

### GetApplicationLastEvaluationTimeEpochMsOk

`func (o *SearchResultItemDTO) GetApplicationLastEvaluationTimeEpochMsOk() (*int64, bool)`

GetApplicationLastEvaluationTimeEpochMsOk returns a tuple with the ApplicationLastEvaluationTimeEpochMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationLastEvaluationTimeEpochMs

`func (o *SearchResultItemDTO) SetApplicationLastEvaluationTimeEpochMs(v int64)`

SetApplicationLastEvaluationTimeEpochMs sets ApplicationLastEvaluationTimeEpochMs field to given value.

### HasApplicationLastEvaluationTimeEpochMs

`func (o *SearchResultItemDTO) HasApplicationLastEvaluationTimeEpochMs() bool`

HasApplicationLastEvaluationTimeEpochMs returns a boolean if a field has been set.

### GetApplicationName

`func (o *SearchResultItemDTO) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *SearchResultItemDTO) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *SearchResultItemDTO) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *SearchResultItemDTO) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationPublicId

`func (o *SearchResultItemDTO) GetApplicationPublicId() string`

GetApplicationPublicId returns the ApplicationPublicId field if non-nil, zero value otherwise.

### GetApplicationPublicIdOk

`func (o *SearchResultItemDTO) GetApplicationPublicIdOk() (*string, bool)`

GetApplicationPublicIdOk returns a tuple with the ApplicationPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPublicId

`func (o *SearchResultItemDTO) SetApplicationPublicId(v string)`

SetApplicationPublicId sets ApplicationPublicId field to given value.

### HasApplicationPublicId

`func (o *SearchResultItemDTO) HasApplicationPublicId() bool`

HasApplicationPublicId returns a boolean if a field has been set.

### GetApplicationStageSeverityCounts

`func (o *SearchResultItemDTO) GetApplicationStageSeverityCounts() []string`

GetApplicationStageSeverityCounts returns the ApplicationStageSeverityCounts field if non-nil, zero value otherwise.

### GetApplicationStageSeverityCountsOk

`func (o *SearchResultItemDTO) GetApplicationStageSeverityCountsOk() (*[]string, bool)`

GetApplicationStageSeverityCountsOk returns a tuple with the ApplicationStageSeverityCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationStageSeverityCounts

`func (o *SearchResultItemDTO) SetApplicationStageSeverityCounts(v []string)`

SetApplicationStageSeverityCounts sets ApplicationStageSeverityCounts field to given value.

### HasApplicationStageSeverityCounts

`func (o *SearchResultItemDTO) HasApplicationStageSeverityCounts() bool`

HasApplicationStageSeverityCounts returns a boolean if a field has been set.

### GetApplicationVersion

`func (o *SearchResultItemDTO) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *SearchResultItemDTO) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *SearchResultItemDTO) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *SearchResultItemDTO) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### GetApplicationViolationPolicyTypes

`func (o *SearchResultItemDTO) GetApplicationViolationPolicyTypes() []string`

GetApplicationViolationPolicyTypes returns the ApplicationViolationPolicyTypes field if non-nil, zero value otherwise.

### GetApplicationViolationPolicyTypesOk

`func (o *SearchResultItemDTO) GetApplicationViolationPolicyTypesOk() (*[]string, bool)`

GetApplicationViolationPolicyTypesOk returns a tuple with the ApplicationViolationPolicyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationViolationPolicyTypes

`func (o *SearchResultItemDTO) SetApplicationViolationPolicyTypes(v []string)`

SetApplicationViolationPolicyTypes sets ApplicationViolationPolicyTypes field to given value.

### HasApplicationViolationPolicyTypes

`func (o *SearchResultItemDTO) HasApplicationViolationPolicyTypes() bool`

HasApplicationViolationPolicyTypes returns a boolean if a field has been set.

### GetApplicationViolationStages

`func (o *SearchResultItemDTO) GetApplicationViolationStages() []string`

GetApplicationViolationStages returns the ApplicationViolationStages field if non-nil, zero value otherwise.

### GetApplicationViolationStagesOk

`func (o *SearchResultItemDTO) GetApplicationViolationStagesOk() (*[]string, bool)`

GetApplicationViolationStagesOk returns a tuple with the ApplicationViolationStages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationViolationStages

`func (o *SearchResultItemDTO) SetApplicationViolationStages(v []string)`

SetApplicationViolationStages sets ApplicationViolationStages field to given value.

### HasApplicationViolationStages

`func (o *SearchResultItemDTO) HasApplicationViolationStages() bool`

HasApplicationViolationStages returns a boolean if a field has been set.

### GetApplicationViolationStates

`func (o *SearchResultItemDTO) GetApplicationViolationStates() []string`

GetApplicationViolationStates returns the ApplicationViolationStates field if non-nil, zero value otherwise.

### GetApplicationViolationStatesOk

`func (o *SearchResultItemDTO) GetApplicationViolationStatesOk() (*[]string, bool)`

GetApplicationViolationStatesOk returns a tuple with the ApplicationViolationStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationViolationStates

`func (o *SearchResultItemDTO) SetApplicationViolationStates(v []string)`

SetApplicationViolationStates sets ApplicationViolationStates field to given value.

### HasApplicationViolationStates

`func (o *SearchResultItemDTO) HasApplicationViolationStates() bool`

HasApplicationViolationStates returns a boolean if a field has been set.

### GetComponentEffectiveLicenseId

`func (o *SearchResultItemDTO) GetComponentEffectiveLicenseId() string`

GetComponentEffectiveLicenseId returns the ComponentEffectiveLicenseId field if non-nil, zero value otherwise.

### GetComponentEffectiveLicenseIdOk

`func (o *SearchResultItemDTO) GetComponentEffectiveLicenseIdOk() (*string, bool)`

GetComponentEffectiveLicenseIdOk returns a tuple with the ComponentEffectiveLicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentEffectiveLicenseId

`func (o *SearchResultItemDTO) SetComponentEffectiveLicenseId(v string)`

SetComponentEffectiveLicenseId sets ComponentEffectiveLicenseId field to given value.

### HasComponentEffectiveLicenseId

`func (o *SearchResultItemDTO) HasComponentEffectiveLicenseId() bool`

HasComponentEffectiveLicenseId returns a boolean if a field has been set.

### GetComponentEffectiveLicenseName

`func (o *SearchResultItemDTO) GetComponentEffectiveLicenseName() string`

GetComponentEffectiveLicenseName returns the ComponentEffectiveLicenseName field if non-nil, zero value otherwise.

### GetComponentEffectiveLicenseNameOk

`func (o *SearchResultItemDTO) GetComponentEffectiveLicenseNameOk() (*string, bool)`

GetComponentEffectiveLicenseNameOk returns a tuple with the ComponentEffectiveLicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentEffectiveLicenseName

`func (o *SearchResultItemDTO) SetComponentEffectiveLicenseName(v string)`

SetComponentEffectiveLicenseName sets ComponentEffectiveLicenseName field to given value.

### HasComponentEffectiveLicenseName

`func (o *SearchResultItemDTO) HasComponentEffectiveLicenseName() bool`

HasComponentEffectiveLicenseName returns a boolean if a field has been set.

### GetComponentHash

`func (o *SearchResultItemDTO) GetComponentHash() string`

GetComponentHash returns the ComponentHash field if non-nil, zero value otherwise.

### GetComponentHashOk

`func (o *SearchResultItemDTO) GetComponentHashOk() (*string, bool)`

GetComponentHashOk returns a tuple with the ComponentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentHash

`func (o *SearchResultItemDTO) SetComponentHash(v string)`

SetComponentHash sets ComponentHash field to given value.

### HasComponentHash

`func (o *SearchResultItemDTO) HasComponentHash() bool`

HasComponentHash returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *SearchResultItemDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *SearchResultItemDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *SearchResultItemDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *SearchResultItemDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetComponentLabelColor

`func (o *SearchResultItemDTO) GetComponentLabelColor() string`

GetComponentLabelColor returns the ComponentLabelColor field if non-nil, zero value otherwise.

### GetComponentLabelColorOk

`func (o *SearchResultItemDTO) GetComponentLabelColorOk() (*string, bool)`

GetComponentLabelColorOk returns a tuple with the ComponentLabelColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLabelColor

`func (o *SearchResultItemDTO) SetComponentLabelColor(v string)`

SetComponentLabelColor sets ComponentLabelColor field to given value.

### HasComponentLabelColor

`func (o *SearchResultItemDTO) HasComponentLabelColor() bool`

HasComponentLabelColor returns a boolean if a field has been set.

### GetComponentLabelDescription

`func (o *SearchResultItemDTO) GetComponentLabelDescription() string`

GetComponentLabelDescription returns the ComponentLabelDescription field if non-nil, zero value otherwise.

### GetComponentLabelDescriptionOk

`func (o *SearchResultItemDTO) GetComponentLabelDescriptionOk() (*string, bool)`

GetComponentLabelDescriptionOk returns a tuple with the ComponentLabelDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLabelDescription

`func (o *SearchResultItemDTO) SetComponentLabelDescription(v string)`

SetComponentLabelDescription sets ComponentLabelDescription field to given value.

### HasComponentLabelDescription

`func (o *SearchResultItemDTO) HasComponentLabelDescription() bool`

HasComponentLabelDescription returns a boolean if a field has been set.

### GetComponentLabelId

`func (o *SearchResultItemDTO) GetComponentLabelId() string`

GetComponentLabelId returns the ComponentLabelId field if non-nil, zero value otherwise.

### GetComponentLabelIdOk

`func (o *SearchResultItemDTO) GetComponentLabelIdOk() (*string, bool)`

GetComponentLabelIdOk returns a tuple with the ComponentLabelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLabelId

`func (o *SearchResultItemDTO) SetComponentLabelId(v string)`

SetComponentLabelId sets ComponentLabelId field to given value.

### HasComponentLabelId

`func (o *SearchResultItemDTO) HasComponentLabelId() bool`

HasComponentLabelId returns a boolean if a field has been set.

### GetComponentLabelName

`func (o *SearchResultItemDTO) GetComponentLabelName() string`

GetComponentLabelName returns the ComponentLabelName field if non-nil, zero value otherwise.

### GetComponentLabelNameOk

`func (o *SearchResultItemDTO) GetComponentLabelNameOk() (*string, bool)`

GetComponentLabelNameOk returns a tuple with the ComponentLabelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLabelName

`func (o *SearchResultItemDTO) SetComponentLabelName(v string)`

SetComponentLabelName sets ComponentLabelName field to given value.

### HasComponentLabelName

`func (o *SearchResultItemDTO) HasComponentLabelName() bool`

HasComponentLabelName returns a boolean if a field has been set.

### GetComponentLicenseThreatGroupName

`func (o *SearchResultItemDTO) GetComponentLicenseThreatGroupName() string`

GetComponentLicenseThreatGroupName returns the ComponentLicenseThreatGroupName field if non-nil, zero value otherwise.

### GetComponentLicenseThreatGroupNameOk

`func (o *SearchResultItemDTO) GetComponentLicenseThreatGroupNameOk() (*string, bool)`

GetComponentLicenseThreatGroupNameOk returns a tuple with the ComponentLicenseThreatGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLicenseThreatGroupName

`func (o *SearchResultItemDTO) SetComponentLicenseThreatGroupName(v string)`

SetComponentLicenseThreatGroupName sets ComponentLicenseThreatGroupName field to given value.

### HasComponentLicenseThreatGroupName

`func (o *SearchResultItemDTO) HasComponentLicenseThreatGroupName() bool`

HasComponentLicenseThreatGroupName returns a boolean if a field has been set.

### GetComponentLicenseThreatLevel

`func (o *SearchResultItemDTO) GetComponentLicenseThreatLevel() int32`

GetComponentLicenseThreatLevel returns the ComponentLicenseThreatLevel field if non-nil, zero value otherwise.

### GetComponentLicenseThreatLevelOk

`func (o *SearchResultItemDTO) GetComponentLicenseThreatLevelOk() (*int32, bool)`

GetComponentLicenseThreatLevelOk returns a tuple with the ComponentLicenseThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLicenseThreatLevel

`func (o *SearchResultItemDTO) SetComponentLicenseThreatLevel(v int32)`

SetComponentLicenseThreatLevel sets ComponentLicenseThreatLevel field to given value.

### HasComponentLicenseThreatLevel

`func (o *SearchResultItemDTO) HasComponentLicenseThreatLevel() bool`

HasComponentLicenseThreatLevel returns a boolean if a field has been set.

### GetComponentName

`func (o *SearchResultItemDTO) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *SearchResultItemDTO) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *SearchResultItemDTO) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *SearchResultItemDTO) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetComponentViolationPolicyTypes

`func (o *SearchResultItemDTO) GetComponentViolationPolicyTypes() []string`

GetComponentViolationPolicyTypes returns the ComponentViolationPolicyTypes field if non-nil, zero value otherwise.

### GetComponentViolationPolicyTypesOk

`func (o *SearchResultItemDTO) GetComponentViolationPolicyTypesOk() (*[]string, bool)`

GetComponentViolationPolicyTypesOk returns a tuple with the ComponentViolationPolicyTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentViolationPolicyTypes

`func (o *SearchResultItemDTO) SetComponentViolationPolicyTypes(v []string)`

SetComponentViolationPolicyTypes sets ComponentViolationPolicyTypes field to given value.

### HasComponentViolationPolicyTypes

`func (o *SearchResultItemDTO) HasComponentViolationPolicyTypes() bool`

HasComponentViolationPolicyTypes returns a boolean if a field has been set.

### GetComponentViolationStates

`func (o *SearchResultItemDTO) GetComponentViolationStates() []string`

GetComponentViolationStates returns the ComponentViolationStates field if non-nil, zero value otherwise.

### GetComponentViolationStatesOk

`func (o *SearchResultItemDTO) GetComponentViolationStatesOk() (*[]string, bool)`

GetComponentViolationStatesOk returns a tuple with the ComponentViolationStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentViolationStates

`func (o *SearchResultItemDTO) SetComponentViolationStates(v []string)`

SetComponentViolationStates sets ComponentViolationStates field to given value.

### HasComponentViolationStates

`func (o *SearchResultItemDTO) HasComponentViolationStates() bool`

HasComponentViolationStates returns a boolean if a field has been set.

### GetItemType

`func (o *SearchResultItemDTO) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *SearchResultItemDTO) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *SearchResultItemDTO) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *SearchResultItemDTO) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetNoteToReviewer

`func (o *SearchResultItemDTO) GetNoteToReviewer() string`

GetNoteToReviewer returns the NoteToReviewer field if non-nil, zero value otherwise.

### GetNoteToReviewerOk

`func (o *SearchResultItemDTO) GetNoteToReviewerOk() (*string, bool)`

GetNoteToReviewerOk returns a tuple with the NoteToReviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteToReviewer

`func (o *SearchResultItemDTO) SetNoteToReviewer(v string)`

SetNoteToReviewer sets NoteToReviewer field to given value.

### HasNoteToReviewer

`func (o *SearchResultItemDTO) HasNoteToReviewer() bool`

HasNoteToReviewer returns a boolean if a field has been set.

### GetOrganizationId

`func (o *SearchResultItemDTO) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *SearchResultItemDTO) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *SearchResultItemDTO) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *SearchResultItemDTO) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetOrganizationName

`func (o *SearchResultItemDTO) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *SearchResultItemDTO) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *SearchResultItemDTO) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *SearchResultItemDTO) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### GetPolicyEvaluationStage

`func (o *SearchResultItemDTO) GetPolicyEvaluationStage() string`

GetPolicyEvaluationStage returns the PolicyEvaluationStage field if non-nil, zero value otherwise.

### GetPolicyEvaluationStageOk

`func (o *SearchResultItemDTO) GetPolicyEvaluationStageOk() (*string, bool)`

GetPolicyEvaluationStageOk returns a tuple with the PolicyEvaluationStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEvaluationStage

`func (o *SearchResultItemDTO) SetPolicyEvaluationStage(v string)`

SetPolicyEvaluationStage sets PolicyEvaluationStage field to given value.

### HasPolicyEvaluationStage

`func (o *SearchResultItemDTO) HasPolicyEvaluationStage() bool`

HasPolicyEvaluationStage returns a boolean if a field has been set.

### GetPolicyId

`func (o *SearchResultItemDTO) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *SearchResultItemDTO) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *SearchResultItemDTO) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *SearchResultItemDTO) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *SearchResultItemDTO) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *SearchResultItemDTO) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *SearchResultItemDTO) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *SearchResultItemDTO) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyThreatCategory

`func (o *SearchResultItemDTO) GetPolicyThreatCategory() string`

GetPolicyThreatCategory returns the PolicyThreatCategory field if non-nil, zero value otherwise.

### GetPolicyThreatCategoryOk

`func (o *SearchResultItemDTO) GetPolicyThreatCategoryOk() (*string, bool)`

GetPolicyThreatCategoryOk returns a tuple with the PolicyThreatCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyThreatCategory

`func (o *SearchResultItemDTO) SetPolicyThreatCategory(v string)`

SetPolicyThreatCategory sets PolicyThreatCategory field to given value.

### HasPolicyThreatCategory

`func (o *SearchResultItemDTO) HasPolicyThreatCategory() bool`

HasPolicyThreatCategory returns a boolean if a field has been set.

### GetPolicyThreatLevel

`func (o *SearchResultItemDTO) GetPolicyThreatLevel() int32`

GetPolicyThreatLevel returns the PolicyThreatLevel field if non-nil, zero value otherwise.

### GetPolicyThreatLevelOk

`func (o *SearchResultItemDTO) GetPolicyThreatLevelOk() (*int32, bool)`

GetPolicyThreatLevelOk returns a tuple with the PolicyThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyThreatLevel

`func (o *SearchResultItemDTO) SetPolicyThreatLevel(v int32)`

SetPolicyThreatLevel sets PolicyThreatLevel field to given value.

### HasPolicyThreatLevel

`func (o *SearchResultItemDTO) HasPolicyThreatLevel() bool`

HasPolicyThreatLevel returns a boolean if a field has been set.

### GetPolicyViolationConstraintName

`func (o *SearchResultItemDTO) GetPolicyViolationConstraintName() string`

GetPolicyViolationConstraintName returns the PolicyViolationConstraintName field if non-nil, zero value otherwise.

### GetPolicyViolationConstraintNameOk

`func (o *SearchResultItemDTO) GetPolicyViolationConstraintNameOk() (*string, bool)`

GetPolicyViolationConstraintNameOk returns a tuple with the PolicyViolationConstraintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationConstraintName

`func (o *SearchResultItemDTO) SetPolicyViolationConstraintName(v string)`

SetPolicyViolationConstraintName sets PolicyViolationConstraintName field to given value.

### HasPolicyViolationConstraintName

`func (o *SearchResultItemDTO) HasPolicyViolationConstraintName() bool`

HasPolicyViolationConstraintName returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *SearchResultItemDTO) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *SearchResultItemDTO) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *SearchResultItemDTO) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *SearchResultItemDTO) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetPolicyViolationPolicyId

`func (o *SearchResultItemDTO) GetPolicyViolationPolicyId() string`

GetPolicyViolationPolicyId returns the PolicyViolationPolicyId field if non-nil, zero value otherwise.

### GetPolicyViolationPolicyIdOk

`func (o *SearchResultItemDTO) GetPolicyViolationPolicyIdOk() (*string, bool)`

GetPolicyViolationPolicyIdOk returns a tuple with the PolicyViolationPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationPolicyId

`func (o *SearchResultItemDTO) SetPolicyViolationPolicyId(v string)`

SetPolicyViolationPolicyId sets PolicyViolationPolicyId field to given value.

### HasPolicyViolationPolicyId

`func (o *SearchResultItemDTO) HasPolicyViolationPolicyId() bool`

HasPolicyViolationPolicyId returns a boolean if a field has been set.

### GetPolicyViolationPolicyName

`func (o *SearchResultItemDTO) GetPolicyViolationPolicyName() string`

GetPolicyViolationPolicyName returns the PolicyViolationPolicyName field if non-nil, zero value otherwise.

### GetPolicyViolationPolicyNameOk

`func (o *SearchResultItemDTO) GetPolicyViolationPolicyNameOk() (*string, bool)`

GetPolicyViolationPolicyNameOk returns a tuple with the PolicyViolationPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationPolicyName

`func (o *SearchResultItemDTO) SetPolicyViolationPolicyName(v string)`

SetPolicyViolationPolicyName sets PolicyViolationPolicyName field to given value.

### HasPolicyViolationPolicyName

`func (o *SearchResultItemDTO) HasPolicyViolationPolicyName() bool`

HasPolicyViolationPolicyName returns a boolean if a field has been set.

### GetPolicyViolationThreatCategory

`func (o *SearchResultItemDTO) GetPolicyViolationThreatCategory() string`

GetPolicyViolationThreatCategory returns the PolicyViolationThreatCategory field if non-nil, zero value otherwise.

### GetPolicyViolationThreatCategoryOk

`func (o *SearchResultItemDTO) GetPolicyViolationThreatCategoryOk() (*string, bool)`

GetPolicyViolationThreatCategoryOk returns a tuple with the PolicyViolationThreatCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationThreatCategory

`func (o *SearchResultItemDTO) SetPolicyViolationThreatCategory(v string)`

SetPolicyViolationThreatCategory sets PolicyViolationThreatCategory field to given value.

### HasPolicyViolationThreatCategory

`func (o *SearchResultItemDTO) HasPolicyViolationThreatCategory() bool`

HasPolicyViolationThreatCategory returns a boolean if a field has been set.

### GetPolicyViolationThreatLevel

`func (o *SearchResultItemDTO) GetPolicyViolationThreatLevel() int32`

GetPolicyViolationThreatLevel returns the PolicyViolationThreatLevel field if non-nil, zero value otherwise.

### GetPolicyViolationThreatLevelOk

`func (o *SearchResultItemDTO) GetPolicyViolationThreatLevelOk() (*int32, bool)`

GetPolicyViolationThreatLevelOk returns a tuple with the PolicyViolationThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationThreatLevel

`func (o *SearchResultItemDTO) SetPolicyViolationThreatLevel(v int32)`

SetPolicyViolationThreatLevel sets PolicyViolationThreatLevel field to given value.

### HasPolicyViolationThreatLevel

`func (o *SearchResultItemDTO) HasPolicyViolationThreatLevel() bool`

HasPolicyViolationThreatLevel returns a boolean if a field has been set.

### GetPolicyViolationWaiverStatus

`func (o *SearchResultItemDTO) GetPolicyViolationWaiverStatus() string`

GetPolicyViolationWaiverStatus returns the PolicyViolationWaiverStatus field if non-nil, zero value otherwise.

### GetPolicyViolationWaiverStatusOk

`func (o *SearchResultItemDTO) GetPolicyViolationWaiverStatusOk() (*string, bool)`

GetPolicyViolationWaiverStatusOk returns a tuple with the PolicyViolationWaiverStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationWaiverStatus

`func (o *SearchResultItemDTO) SetPolicyViolationWaiverStatus(v string)`

SetPolicyViolationWaiverStatus sets PolicyViolationWaiverStatus field to given value.

### HasPolicyViolationWaiverStatus

`func (o *SearchResultItemDTO) HasPolicyViolationWaiverStatus() bool`

HasPolicyViolationWaiverStatus returns a boolean if a field has been set.

### GetPolicyWaiverAuto

`func (o *SearchResultItemDTO) GetPolicyWaiverAuto() bool`

GetPolicyWaiverAuto returns the PolicyWaiverAuto field if non-nil, zero value otherwise.

### GetPolicyWaiverAutoOk

`func (o *SearchResultItemDTO) GetPolicyWaiverAutoOk() (*bool, bool)`

GetPolicyWaiverAutoOk returns a tuple with the PolicyWaiverAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverAuto

`func (o *SearchResultItemDTO) SetPolicyWaiverAuto(v bool)`

SetPolicyWaiverAuto sets PolicyWaiverAuto field to given value.

### HasPolicyWaiverAuto

`func (o *SearchResultItemDTO) HasPolicyWaiverAuto() bool`

HasPolicyWaiverAuto returns a boolean if a field has been set.

### GetPolicyWaiverComment

`func (o *SearchResultItemDTO) GetPolicyWaiverComment() string`

GetPolicyWaiverComment returns the PolicyWaiverComment field if non-nil, zero value otherwise.

### GetPolicyWaiverCommentOk

`func (o *SearchResultItemDTO) GetPolicyWaiverCommentOk() (*string, bool)`

GetPolicyWaiverCommentOk returns a tuple with the PolicyWaiverComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverComment

`func (o *SearchResultItemDTO) SetPolicyWaiverComment(v string)`

SetPolicyWaiverComment sets PolicyWaiverComment field to given value.

### HasPolicyWaiverComment

`func (o *SearchResultItemDTO) HasPolicyWaiverComment() bool`

HasPolicyWaiverComment returns a boolean if a field has been set.

### GetPolicyWaiverCreatedAt

`func (o *SearchResultItemDTO) GetPolicyWaiverCreatedAt() string`

GetPolicyWaiverCreatedAt returns the PolicyWaiverCreatedAt field if non-nil, zero value otherwise.

### GetPolicyWaiverCreatedAtOk

`func (o *SearchResultItemDTO) GetPolicyWaiverCreatedAtOk() (*string, bool)`

GetPolicyWaiverCreatedAtOk returns a tuple with the PolicyWaiverCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverCreatedAt

`func (o *SearchResultItemDTO) SetPolicyWaiverCreatedAt(v string)`

SetPolicyWaiverCreatedAt sets PolicyWaiverCreatedAt field to given value.

### HasPolicyWaiverCreatedAt

`func (o *SearchResultItemDTO) HasPolicyWaiverCreatedAt() bool`

HasPolicyWaiverCreatedAt returns a boolean if a field has been set.

### GetPolicyWaiverExpiresAt

`func (o *SearchResultItemDTO) GetPolicyWaiverExpiresAt() string`

GetPolicyWaiverExpiresAt returns the PolicyWaiverExpiresAt field if non-nil, zero value otherwise.

### GetPolicyWaiverExpiresAtOk

`func (o *SearchResultItemDTO) GetPolicyWaiverExpiresAtOk() (*string, bool)`

GetPolicyWaiverExpiresAtOk returns a tuple with the PolicyWaiverExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverExpiresAt

`func (o *SearchResultItemDTO) SetPolicyWaiverExpiresAt(v string)`

SetPolicyWaiverExpiresAt sets PolicyWaiverExpiresAt field to given value.

### HasPolicyWaiverExpiresAt

`func (o *SearchResultItemDTO) HasPolicyWaiverExpiresAt() bool`

HasPolicyWaiverExpiresAt returns a boolean if a field has been set.

### GetPolicyWaiverExpiryStatus

`func (o *SearchResultItemDTO) GetPolicyWaiverExpiryStatus() string`

GetPolicyWaiverExpiryStatus returns the PolicyWaiverExpiryStatus field if non-nil, zero value otherwise.

### GetPolicyWaiverExpiryStatusOk

`func (o *SearchResultItemDTO) GetPolicyWaiverExpiryStatusOk() (*string, bool)`

GetPolicyWaiverExpiryStatusOk returns a tuple with the PolicyWaiverExpiryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverExpiryStatus

`func (o *SearchResultItemDTO) SetPolicyWaiverExpiryStatus(v string)`

SetPolicyWaiverExpiryStatus sets PolicyWaiverExpiryStatus field to given value.

### HasPolicyWaiverExpiryStatus

`func (o *SearchResultItemDTO) HasPolicyWaiverExpiryStatus() bool`

HasPolicyWaiverExpiryStatus returns a boolean if a field has been set.

### GetPolicyWaiverId

`func (o *SearchResultItemDTO) GetPolicyWaiverId() string`

GetPolicyWaiverId returns the PolicyWaiverId field if non-nil, zero value otherwise.

### GetPolicyWaiverIdOk

`func (o *SearchResultItemDTO) GetPolicyWaiverIdOk() (*string, bool)`

GetPolicyWaiverIdOk returns a tuple with the PolicyWaiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverId

`func (o *SearchResultItemDTO) SetPolicyWaiverId(v string)`

SetPolicyWaiverId sets PolicyWaiverId field to given value.

### HasPolicyWaiverId

`func (o *SearchResultItemDTO) HasPolicyWaiverId() bool`

HasPolicyWaiverId returns a boolean if a field has been set.

### GetPolicyWaiverIsAuto

`func (o *SearchResultItemDTO) GetPolicyWaiverIsAuto() bool`

GetPolicyWaiverIsAuto returns the PolicyWaiverIsAuto field if non-nil, zero value otherwise.

### GetPolicyWaiverIsAutoOk

`func (o *SearchResultItemDTO) GetPolicyWaiverIsAutoOk() (*bool, bool)`

GetPolicyWaiverIsAutoOk returns a tuple with the PolicyWaiverIsAuto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverIsAuto

`func (o *SearchResultItemDTO) SetPolicyWaiverIsAuto(v bool)`

SetPolicyWaiverIsAuto sets PolicyWaiverIsAuto field to given value.

### HasPolicyWaiverIsAuto

`func (o *SearchResultItemDTO) HasPolicyWaiverIsAuto() bool`

HasPolicyWaiverIsAuto returns a boolean if a field has been set.

### GetPolicyWaiverPolicyId

`func (o *SearchResultItemDTO) GetPolicyWaiverPolicyId() string`

GetPolicyWaiverPolicyId returns the PolicyWaiverPolicyId field if non-nil, zero value otherwise.

### GetPolicyWaiverPolicyIdOk

`func (o *SearchResultItemDTO) GetPolicyWaiverPolicyIdOk() (*string, bool)`

GetPolicyWaiverPolicyIdOk returns a tuple with the PolicyWaiverPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverPolicyId

`func (o *SearchResultItemDTO) SetPolicyWaiverPolicyId(v string)`

SetPolicyWaiverPolicyId sets PolicyWaiverPolicyId field to given value.

### HasPolicyWaiverPolicyId

`func (o *SearchResultItemDTO) HasPolicyWaiverPolicyId() bool`

HasPolicyWaiverPolicyId returns a boolean if a field has been set.

### GetPolicyWaiverPolicyName

`func (o *SearchResultItemDTO) GetPolicyWaiverPolicyName() string`

GetPolicyWaiverPolicyName returns the PolicyWaiverPolicyName field if non-nil, zero value otherwise.

### GetPolicyWaiverPolicyNameOk

`func (o *SearchResultItemDTO) GetPolicyWaiverPolicyNameOk() (*string, bool)`

GetPolicyWaiverPolicyNameOk returns a tuple with the PolicyWaiverPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverPolicyName

`func (o *SearchResultItemDTO) SetPolicyWaiverPolicyName(v string)`

SetPolicyWaiverPolicyName sets PolicyWaiverPolicyName field to given value.

### HasPolicyWaiverPolicyName

`func (o *SearchResultItemDTO) HasPolicyWaiverPolicyName() bool`

HasPolicyWaiverPolicyName returns a boolean if a field has been set.

### GetPolicyWaiverPolicyType

`func (o *SearchResultItemDTO) GetPolicyWaiverPolicyType() string`

GetPolicyWaiverPolicyType returns the PolicyWaiverPolicyType field if non-nil, zero value otherwise.

### GetPolicyWaiverPolicyTypeOk

`func (o *SearchResultItemDTO) GetPolicyWaiverPolicyTypeOk() (*string, bool)`

GetPolicyWaiverPolicyTypeOk returns a tuple with the PolicyWaiverPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverPolicyType

`func (o *SearchResultItemDTO) SetPolicyWaiverPolicyType(v string)`

SetPolicyWaiverPolicyType sets PolicyWaiverPolicyType field to given value.

### HasPolicyWaiverPolicyType

`func (o *SearchResultItemDTO) HasPolicyWaiverPolicyType() bool`

HasPolicyWaiverPolicyType returns a boolean if a field has been set.

### GetPolicyWaiverReason

`func (o *SearchResultItemDTO) GetPolicyWaiverReason() string`

GetPolicyWaiverReason returns the PolicyWaiverReason field if non-nil, zero value otherwise.

### GetPolicyWaiverReasonOk

`func (o *SearchResultItemDTO) GetPolicyWaiverReasonOk() (*string, bool)`

GetPolicyWaiverReasonOk returns a tuple with the PolicyWaiverReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverReason

`func (o *SearchResultItemDTO) SetPolicyWaiverReason(v string)`

SetPolicyWaiverReason sets PolicyWaiverReason field to given value.

### HasPolicyWaiverReason

`func (o *SearchResultItemDTO) HasPolicyWaiverReason() bool`

HasPolicyWaiverReason returns a boolean if a field has been set.

### GetPolicyWaiverRequestStatus

`func (o *SearchResultItemDTO) GetPolicyWaiverRequestStatus() string`

GetPolicyWaiverRequestStatus returns the PolicyWaiverRequestStatus field if non-nil, zero value otherwise.

### GetPolicyWaiverRequestStatusOk

`func (o *SearchResultItemDTO) GetPolicyWaiverRequestStatusOk() (*string, bool)`

GetPolicyWaiverRequestStatusOk returns a tuple with the PolicyWaiverRequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverRequestStatus

`func (o *SearchResultItemDTO) SetPolicyWaiverRequestStatus(v string)`

SetPolicyWaiverRequestStatus sets PolicyWaiverRequestStatus field to given value.

### HasPolicyWaiverRequestStatus

`func (o *SearchResultItemDTO) HasPolicyWaiverRequestStatus() bool`

HasPolicyWaiverRequestStatus returns a boolean if a field has been set.

### GetPolicyWaiverScope

`func (o *SearchResultItemDTO) GetPolicyWaiverScope() string`

GetPolicyWaiverScope returns the PolicyWaiverScope field if non-nil, zero value otherwise.

### GetPolicyWaiverScopeOk

`func (o *SearchResultItemDTO) GetPolicyWaiverScopeOk() (*string, bool)`

GetPolicyWaiverScopeOk returns a tuple with the PolicyWaiverScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverScope

`func (o *SearchResultItemDTO) SetPolicyWaiverScope(v string)`

SetPolicyWaiverScope sets PolicyWaiverScope field to given value.

### HasPolicyWaiverScope

`func (o *SearchResultItemDTO) HasPolicyWaiverScope() bool`

HasPolicyWaiverScope returns a boolean if a field has been set.

### GetPolicyWaiverScopeOwnerId

`func (o *SearchResultItemDTO) GetPolicyWaiverScopeOwnerId() string`

GetPolicyWaiverScopeOwnerId returns the PolicyWaiverScopeOwnerId field if non-nil, zero value otherwise.

### GetPolicyWaiverScopeOwnerIdOk

`func (o *SearchResultItemDTO) GetPolicyWaiverScopeOwnerIdOk() (*string, bool)`

GetPolicyWaiverScopeOwnerIdOk returns a tuple with the PolicyWaiverScopeOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverScopeOwnerId

`func (o *SearchResultItemDTO) SetPolicyWaiverScopeOwnerId(v string)`

SetPolicyWaiverScopeOwnerId sets PolicyWaiverScopeOwnerId field to given value.

### HasPolicyWaiverScopeOwnerId

`func (o *SearchResultItemDTO) HasPolicyWaiverScopeOwnerId() bool`

HasPolicyWaiverScopeOwnerId returns a boolean if a field has been set.

### GetPolicyWaiverScopeOwnerType

`func (o *SearchResultItemDTO) GetPolicyWaiverScopeOwnerType() string`

GetPolicyWaiverScopeOwnerType returns the PolicyWaiverScopeOwnerType field if non-nil, zero value otherwise.

### GetPolicyWaiverScopeOwnerTypeOk

`func (o *SearchResultItemDTO) GetPolicyWaiverScopeOwnerTypeOk() (*string, bool)`

GetPolicyWaiverScopeOwnerTypeOk returns a tuple with the PolicyWaiverScopeOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverScopeOwnerType

`func (o *SearchResultItemDTO) SetPolicyWaiverScopeOwnerType(v string)`

SetPolicyWaiverScopeOwnerType sets PolicyWaiverScopeOwnerType field to given value.

### HasPolicyWaiverScopeOwnerType

`func (o *SearchResultItemDTO) HasPolicyWaiverScopeOwnerType() bool`

HasPolicyWaiverScopeOwnerType returns a boolean if a field has been set.

### GetPolicyWaiverThreatLevel

`func (o *SearchResultItemDTO) GetPolicyWaiverThreatLevel() int32`

GetPolicyWaiverThreatLevel returns the PolicyWaiverThreatLevel field if non-nil, zero value otherwise.

### GetPolicyWaiverThreatLevelOk

`func (o *SearchResultItemDTO) GetPolicyWaiverThreatLevelOk() (*int32, bool)`

GetPolicyWaiverThreatLevelOk returns a tuple with the PolicyWaiverThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverThreatLevel

`func (o *SearchResultItemDTO) SetPolicyWaiverThreatLevel(v int32)`

SetPolicyWaiverThreatLevel sets PolicyWaiverThreatLevel field to given value.

### HasPolicyWaiverThreatLevel

`func (o *SearchResultItemDTO) HasPolicyWaiverThreatLevel() bool`

HasPolicyWaiverThreatLevel returns a boolean if a field has been set.

### GetPolicyWaiverWaivedBy

`func (o *SearchResultItemDTO) GetPolicyWaiverWaivedBy() string`

GetPolicyWaiverWaivedBy returns the PolicyWaiverWaivedBy field if non-nil, zero value otherwise.

### GetPolicyWaiverWaivedByOk

`func (o *SearchResultItemDTO) GetPolicyWaiverWaivedByOk() (*string, bool)`

GetPolicyWaiverWaivedByOk returns a tuple with the PolicyWaiverWaivedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverWaivedBy

`func (o *SearchResultItemDTO) SetPolicyWaiverWaivedBy(v string)`

SetPolicyWaiverWaivedBy sets PolicyWaiverWaivedBy field to given value.

### HasPolicyWaiverWaivedBy

`func (o *SearchResultItemDTO) HasPolicyWaiverWaivedBy() bool`

HasPolicyWaiverWaivedBy returns a boolean if a field has been set.

### GetRejectionReason

`func (o *SearchResultItemDTO) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *SearchResultItemDTO) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *SearchResultItemDTO) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *SearchResultItemDTO) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetReportId

`func (o *SearchResultItemDTO) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *SearchResultItemDTO) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *SearchResultItemDTO) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *SearchResultItemDTO) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetRequesterName

`func (o *SearchResultItemDTO) GetRequesterName() string`

GetRequesterName returns the RequesterName field if non-nil, zero value otherwise.

### GetRequesterNameOk

`func (o *SearchResultItemDTO) GetRequesterNameOk() (*string, bool)`

GetRequesterNameOk returns a tuple with the RequesterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterName

`func (o *SearchResultItemDTO) SetRequesterName(v string)`

SetRequesterName sets RequesterName field to given value.

### HasRequesterName

`func (o *SearchResultItemDTO) HasRequesterName() bool`

HasRequesterName returns a boolean if a field has been set.

### GetResultIndex

`func (o *SearchResultItemDTO) GetResultIndex() int32`

GetResultIndex returns the ResultIndex field if non-nil, zero value otherwise.

### GetResultIndexOk

`func (o *SearchResultItemDTO) GetResultIndexOk() (*int32, bool)`

GetResultIndexOk returns a tuple with the ResultIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultIndex

`func (o *SearchResultItemDTO) SetResultIndex(v int32)`

SetResultIndex sets ResultIndex field to given value.

### HasResultIndex

`func (o *SearchResultItemDTO) HasResultIndex() bool`

HasResultIndex returns a boolean if a field has been set.

### GetReviewTime

`func (o *SearchResultItemDTO) GetReviewTime() string`

GetReviewTime returns the ReviewTime field if non-nil, zero value otherwise.

### GetReviewTimeOk

`func (o *SearchResultItemDTO) GetReviewTimeOk() (*string, bool)`

GetReviewTimeOk returns a tuple with the ReviewTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewTime

`func (o *SearchResultItemDTO) SetReviewTime(v string)`

SetReviewTime sets ReviewTime field to given value.

### HasReviewTime

`func (o *SearchResultItemDTO) HasReviewTime() bool`

HasReviewTime returns a boolean if a field has been set.

### GetReviewerName

`func (o *SearchResultItemDTO) GetReviewerName() string`

GetReviewerName returns the ReviewerName field if non-nil, zero value otherwise.

### GetReviewerNameOk

`func (o *SearchResultItemDTO) GetReviewerNameOk() (*string, bool)`

GetReviewerNameOk returns a tuple with the ReviewerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerName

`func (o *SearchResultItemDTO) SetReviewerName(v string)`

SetReviewerName sets ReviewerName field to given value.

### HasReviewerName

`func (o *SearchResultItemDTO) HasReviewerName() bool`

HasReviewerName returns a boolean if a field has been set.

### GetSbomSpecification

`func (o *SearchResultItemDTO) GetSbomSpecification() string`

GetSbomSpecification returns the SbomSpecification field if non-nil, zero value otherwise.

### GetSbomSpecificationOk

`func (o *SearchResultItemDTO) GetSbomSpecificationOk() (*string, bool)`

GetSbomSpecificationOk returns a tuple with the SbomSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSbomSpecification

`func (o *SearchResultItemDTO) SetSbomSpecification(v string)`

SetSbomSpecification sets SbomSpecification field to given value.

### HasSbomSpecification

`func (o *SearchResultItemDTO) HasSbomSpecification() bool`

HasSbomSpecification returns a boolean if a field has been set.

### GetVulnerabilityDescription

`func (o *SearchResultItemDTO) GetVulnerabilityDescription() string`

GetVulnerabilityDescription returns the VulnerabilityDescription field if non-nil, zero value otherwise.

### GetVulnerabilityDescriptionOk

`func (o *SearchResultItemDTO) GetVulnerabilityDescriptionOk() (*string, bool)`

GetVulnerabilityDescriptionOk returns a tuple with the VulnerabilityDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityDescription

`func (o *SearchResultItemDTO) SetVulnerabilityDescription(v string)`

SetVulnerabilityDescription sets VulnerabilityDescription field to given value.

### HasVulnerabilityDescription

`func (o *SearchResultItemDTO) HasVulnerabilityDescription() bool`

HasVulnerabilityDescription returns a boolean if a field has been set.

### GetVulnerabilityFirstSeenEpochMs

`func (o *SearchResultItemDTO) GetVulnerabilityFirstSeenEpochMs() int64`

GetVulnerabilityFirstSeenEpochMs returns the VulnerabilityFirstSeenEpochMs field if non-nil, zero value otherwise.

### GetVulnerabilityFirstSeenEpochMsOk

`func (o *SearchResultItemDTO) GetVulnerabilityFirstSeenEpochMsOk() (*int64, bool)`

GetVulnerabilityFirstSeenEpochMsOk returns a tuple with the VulnerabilityFirstSeenEpochMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityFirstSeenEpochMs

`func (o *SearchResultItemDTO) SetVulnerabilityFirstSeenEpochMs(v int64)`

SetVulnerabilityFirstSeenEpochMs sets VulnerabilityFirstSeenEpochMs field to given value.

### HasVulnerabilityFirstSeenEpochMs

`func (o *SearchResultItemDTO) HasVulnerabilityFirstSeenEpochMs() bool`

HasVulnerabilityFirstSeenEpochMs returns a boolean if a field has been set.

### GetVulnerabilityId

`func (o *SearchResultItemDTO) GetVulnerabilityId() string`

GetVulnerabilityId returns the VulnerabilityId field if non-nil, zero value otherwise.

### GetVulnerabilityIdOk

`func (o *SearchResultItemDTO) GetVulnerabilityIdOk() (*string, bool)`

GetVulnerabilityIdOk returns a tuple with the VulnerabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityId

`func (o *SearchResultItemDTO) SetVulnerabilityId(v string)`

SetVulnerabilityId sets VulnerabilityId field to given value.

### HasVulnerabilityId

`func (o *SearchResultItemDTO) HasVulnerabilityId() bool`

HasVulnerabilityId returns a boolean if a field has been set.

### GetVulnerabilitySeverity

`func (o *SearchResultItemDTO) GetVulnerabilitySeverity() float32`

GetVulnerabilitySeverity returns the VulnerabilitySeverity field if non-nil, zero value otherwise.

### GetVulnerabilitySeverityOk

`func (o *SearchResultItemDTO) GetVulnerabilitySeverityOk() (*float32, bool)`

GetVulnerabilitySeverityOk returns a tuple with the VulnerabilitySeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilitySeverity

`func (o *SearchResultItemDTO) SetVulnerabilitySeverity(v float32)`

SetVulnerabilitySeverity sets VulnerabilitySeverity field to given value.

### HasVulnerabilitySeverity

`func (o *SearchResultItemDTO) HasVulnerabilitySeverity() bool`

HasVulnerabilitySeverity returns a boolean if a field has been set.

### GetVulnerabilityStatus

`func (o *SearchResultItemDTO) GetVulnerabilityStatus() string`

GetVulnerabilityStatus returns the VulnerabilityStatus field if non-nil, zero value otherwise.

### GetVulnerabilityStatusOk

`func (o *SearchResultItemDTO) GetVulnerabilityStatusOk() (*string, bool)`

GetVulnerabilityStatusOk returns a tuple with the VulnerabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityStatus

`func (o *SearchResultItemDTO) SetVulnerabilityStatus(v string)`

SetVulnerabilityStatus sets VulnerabilityStatus field to given value.

### HasVulnerabilityStatus

`func (o *SearchResultItemDTO) HasVulnerabilityStatus() bool`

HasVulnerabilityStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


