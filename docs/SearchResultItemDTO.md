# SearchResultItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationCategoryColor** | Pointer to **string** |  | [optional] 
**ApplicationCategoryDescription** | Pointer to **string** |  | [optional] 
**ApplicationCategoryId** | Pointer to **string** |  | [optional] 
**ApplicationCategoryName** | Pointer to **string** |  | [optional] 
**ApplicationId** | Pointer to **string** |  | [optional] 
**ApplicationName** | Pointer to **string** |  | [optional] 
**ApplicationPublicId** | Pointer to **string** |  | [optional] 
**ApplicationVersion** | Pointer to **string** |  | [optional] 
**ComponentHash** | Pointer to **string** |  | [optional] 
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**ComponentLabelColor** | Pointer to **string** |  | [optional] 
**ComponentLabelDescription** | Pointer to **string** |  | [optional] 
**ComponentLabelId** | Pointer to **string** |  | [optional] 
**ComponentLabelName** | Pointer to **string** |  | [optional] 
**ComponentName** | Pointer to **string** |  | [optional] 
**ItemType** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**OrganizationName** | Pointer to **string** |  | [optional] 
**PolicyEvaluationStage** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyThreatCategory** | Pointer to **string** |  | [optional] 
**PolicyThreatLevel** | Pointer to **int32** |  | [optional] 
**ReportId** | Pointer to **string** |  | [optional] 
**ResultIndex** | Pointer to **int32** |  | [optional] 
**SbomSpecification** | Pointer to **string** |  | [optional] 
**VulnerabilityDescription** | Pointer to **string** |  | [optional] 
**VulnerabilityId** | Pointer to **string** |  | [optional] 
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


