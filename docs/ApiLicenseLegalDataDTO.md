# ApiLicenseLegalDataDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributions** | Pointer to [**[]ComponentObligationAttributionDTO**](ComponentObligationAttributionDTO.md) |  | [optional] 
**ComponentCopyrightId** | Pointer to **string** |  | [optional] 
**ComponentCopyrightLastUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ComponentCopyrightLastUpdatedByUsername** | Pointer to **string** |  | [optional] 
**ComponentCopyrightScopeOwnerId** | Pointer to **string** |  | [optional] 
**ComponentLicensesId** | Pointer to **string** |  | [optional] 
**ComponentLicensesLastUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ComponentLicensesLastUpdatedByUsername** | Pointer to **string** |  | [optional] 
**ComponentLicensesScopeOwnerId** | Pointer to **string** |  | [optional] 
**ComponentNoticesId** | Pointer to **string** |  | [optional] 
**ComponentNoticesLastUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**ComponentNoticesLastUpdatedByUsername** | Pointer to **string** |  | [optional] 
**ComponentNoticesScopeOwnerId** | Pointer to **string** |  | [optional] 
**Copyrights** | Pointer to [**[]ApiLicenseLegalCopyrightDTO**](ApiLicenseLegalCopyrightDTO.md) |  | [optional] 
**DeclaredLicenses** | Pointer to **[]string** |  | [optional] 
**EffectiveLicenseStatus** | Pointer to **string** |  | [optional] 
**EffectiveLicenses** | Pointer to **[]string** |  | [optional] 
**HighestEffectiveLicenseThreatGroup** | Pointer to [**ApiLicenseThreatDTOV2**](ApiLicenseThreatDTOV2.md) |  | [optional] 
**LicenseFiles** | Pointer to [**[]ApiLicenseLegalFileDTO**](ApiLicenseLegalFileDTO.md) |  | [optional] 
**NoticeFiles** | Pointer to [**[]ApiLicenseLegalFileDTO**](ApiLicenseLegalFileDTO.md) |  | [optional] 
**Obligations** | Pointer to [**[]ApiLicenseLegalObligationDTO**](ApiLicenseLegalObligationDTO.md) |  | [optional] 
**ObservedLicenses** | Pointer to **[]string** |  | [optional] 
**SourceLinks** | Pointer to [**[]LegalSourceLinkDTO**](LegalSourceLinkDTO.md) |  | [optional] 

## Methods

### NewApiLicenseLegalDataDTO

`func NewApiLicenseLegalDataDTO() *ApiLicenseLegalDataDTO`

NewApiLicenseLegalDataDTO instantiates a new ApiLicenseLegalDataDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLicenseLegalDataDTOWithDefaults

`func NewApiLicenseLegalDataDTOWithDefaults() *ApiLicenseLegalDataDTO`

NewApiLicenseLegalDataDTOWithDefaults instantiates a new ApiLicenseLegalDataDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributions

`func (o *ApiLicenseLegalDataDTO) GetAttributions() []ComponentObligationAttributionDTO`

GetAttributions returns the Attributions field if non-nil, zero value otherwise.

### GetAttributionsOk

`func (o *ApiLicenseLegalDataDTO) GetAttributionsOk() (*[]ComponentObligationAttributionDTO, bool)`

GetAttributionsOk returns a tuple with the Attributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributions

`func (o *ApiLicenseLegalDataDTO) SetAttributions(v []ComponentObligationAttributionDTO)`

SetAttributions sets Attributions field to given value.

### HasAttributions

`func (o *ApiLicenseLegalDataDTO) HasAttributions() bool`

HasAttributions returns a boolean if a field has been set.

### GetComponentCopyrightId

`func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightId() string`

GetComponentCopyrightId returns the ComponentCopyrightId field if non-nil, zero value otherwise.

### GetComponentCopyrightIdOk

`func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightIdOk() (*string, bool)`

GetComponentCopyrightIdOk returns a tuple with the ComponentCopyrightId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCopyrightId

`func (o *ApiLicenseLegalDataDTO) SetComponentCopyrightId(v string)`

SetComponentCopyrightId sets ComponentCopyrightId field to given value.

### HasComponentCopyrightId

`func (o *ApiLicenseLegalDataDTO) HasComponentCopyrightId() bool`

HasComponentCopyrightId returns a boolean if a field has been set.

### GetComponentCopyrightLastUpdatedAt

`func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightLastUpdatedAt() time.Time`

GetComponentCopyrightLastUpdatedAt returns the ComponentCopyrightLastUpdatedAt field if non-nil, zero value otherwise.

### GetComponentCopyrightLastUpdatedAtOk

`func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightLastUpdatedAtOk() (*time.Time, bool)`

GetComponentCopyrightLastUpdatedAtOk returns a tuple with the ComponentCopyrightLastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCopyrightLastUpdatedAt

`func (o *ApiLicenseLegalDataDTO) SetComponentCopyrightLastUpdatedAt(v time.Time)`

SetComponentCopyrightLastUpdatedAt sets ComponentCopyrightLastUpdatedAt field to given value.

### HasComponentCopyrightLastUpdatedAt

`func (o *ApiLicenseLegalDataDTO) HasComponentCopyrightLastUpdatedAt() bool`

HasComponentCopyrightLastUpdatedAt returns a boolean if a field has been set.

### GetComponentCopyrightLastUpdatedByUsername

`func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightLastUpdatedByUsername() string`

GetComponentCopyrightLastUpdatedByUsername returns the ComponentCopyrightLastUpdatedByUsername field if non-nil, zero value otherwise.

### GetComponentCopyrightLastUpdatedByUsernameOk

`func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightLastUpdatedByUsernameOk() (*string, bool)`

GetComponentCopyrightLastUpdatedByUsernameOk returns a tuple with the ComponentCopyrightLastUpdatedByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCopyrightLastUpdatedByUsername

`func (o *ApiLicenseLegalDataDTO) SetComponentCopyrightLastUpdatedByUsername(v string)`

SetComponentCopyrightLastUpdatedByUsername sets ComponentCopyrightLastUpdatedByUsername field to given value.

### HasComponentCopyrightLastUpdatedByUsername

`func (o *ApiLicenseLegalDataDTO) HasComponentCopyrightLastUpdatedByUsername() bool`

HasComponentCopyrightLastUpdatedByUsername returns a boolean if a field has been set.

### GetComponentCopyrightScopeOwnerId

`func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightScopeOwnerId() string`

GetComponentCopyrightScopeOwnerId returns the ComponentCopyrightScopeOwnerId field if non-nil, zero value otherwise.

### GetComponentCopyrightScopeOwnerIdOk

`func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightScopeOwnerIdOk() (*string, bool)`

GetComponentCopyrightScopeOwnerIdOk returns a tuple with the ComponentCopyrightScopeOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCopyrightScopeOwnerId

`func (o *ApiLicenseLegalDataDTO) SetComponentCopyrightScopeOwnerId(v string)`

SetComponentCopyrightScopeOwnerId sets ComponentCopyrightScopeOwnerId field to given value.

### HasComponentCopyrightScopeOwnerId

`func (o *ApiLicenseLegalDataDTO) HasComponentCopyrightScopeOwnerId() bool`

HasComponentCopyrightScopeOwnerId returns a boolean if a field has been set.

### GetComponentLicensesId

`func (o *ApiLicenseLegalDataDTO) GetComponentLicensesId() string`

GetComponentLicensesId returns the ComponentLicensesId field if non-nil, zero value otherwise.

### GetComponentLicensesIdOk

`func (o *ApiLicenseLegalDataDTO) GetComponentLicensesIdOk() (*string, bool)`

GetComponentLicensesIdOk returns a tuple with the ComponentLicensesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLicensesId

`func (o *ApiLicenseLegalDataDTO) SetComponentLicensesId(v string)`

SetComponentLicensesId sets ComponentLicensesId field to given value.

### HasComponentLicensesId

`func (o *ApiLicenseLegalDataDTO) HasComponentLicensesId() bool`

HasComponentLicensesId returns a boolean if a field has been set.

### GetComponentLicensesLastUpdatedAt

`func (o *ApiLicenseLegalDataDTO) GetComponentLicensesLastUpdatedAt() time.Time`

GetComponentLicensesLastUpdatedAt returns the ComponentLicensesLastUpdatedAt field if non-nil, zero value otherwise.

### GetComponentLicensesLastUpdatedAtOk

`func (o *ApiLicenseLegalDataDTO) GetComponentLicensesLastUpdatedAtOk() (*time.Time, bool)`

GetComponentLicensesLastUpdatedAtOk returns a tuple with the ComponentLicensesLastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLicensesLastUpdatedAt

`func (o *ApiLicenseLegalDataDTO) SetComponentLicensesLastUpdatedAt(v time.Time)`

SetComponentLicensesLastUpdatedAt sets ComponentLicensesLastUpdatedAt field to given value.

### HasComponentLicensesLastUpdatedAt

`func (o *ApiLicenseLegalDataDTO) HasComponentLicensesLastUpdatedAt() bool`

HasComponentLicensesLastUpdatedAt returns a boolean if a field has been set.

### GetComponentLicensesLastUpdatedByUsername

`func (o *ApiLicenseLegalDataDTO) GetComponentLicensesLastUpdatedByUsername() string`

GetComponentLicensesLastUpdatedByUsername returns the ComponentLicensesLastUpdatedByUsername field if non-nil, zero value otherwise.

### GetComponentLicensesLastUpdatedByUsernameOk

`func (o *ApiLicenseLegalDataDTO) GetComponentLicensesLastUpdatedByUsernameOk() (*string, bool)`

GetComponentLicensesLastUpdatedByUsernameOk returns a tuple with the ComponentLicensesLastUpdatedByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLicensesLastUpdatedByUsername

`func (o *ApiLicenseLegalDataDTO) SetComponentLicensesLastUpdatedByUsername(v string)`

SetComponentLicensesLastUpdatedByUsername sets ComponentLicensesLastUpdatedByUsername field to given value.

### HasComponentLicensesLastUpdatedByUsername

`func (o *ApiLicenseLegalDataDTO) HasComponentLicensesLastUpdatedByUsername() bool`

HasComponentLicensesLastUpdatedByUsername returns a boolean if a field has been set.

### GetComponentLicensesScopeOwnerId

`func (o *ApiLicenseLegalDataDTO) GetComponentLicensesScopeOwnerId() string`

GetComponentLicensesScopeOwnerId returns the ComponentLicensesScopeOwnerId field if non-nil, zero value otherwise.

### GetComponentLicensesScopeOwnerIdOk

`func (o *ApiLicenseLegalDataDTO) GetComponentLicensesScopeOwnerIdOk() (*string, bool)`

GetComponentLicensesScopeOwnerIdOk returns a tuple with the ComponentLicensesScopeOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentLicensesScopeOwnerId

`func (o *ApiLicenseLegalDataDTO) SetComponentLicensesScopeOwnerId(v string)`

SetComponentLicensesScopeOwnerId sets ComponentLicensesScopeOwnerId field to given value.

### HasComponentLicensesScopeOwnerId

`func (o *ApiLicenseLegalDataDTO) HasComponentLicensesScopeOwnerId() bool`

HasComponentLicensesScopeOwnerId returns a boolean if a field has been set.

### GetComponentNoticesId

`func (o *ApiLicenseLegalDataDTO) GetComponentNoticesId() string`

GetComponentNoticesId returns the ComponentNoticesId field if non-nil, zero value otherwise.

### GetComponentNoticesIdOk

`func (o *ApiLicenseLegalDataDTO) GetComponentNoticesIdOk() (*string, bool)`

GetComponentNoticesIdOk returns a tuple with the ComponentNoticesId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentNoticesId

`func (o *ApiLicenseLegalDataDTO) SetComponentNoticesId(v string)`

SetComponentNoticesId sets ComponentNoticesId field to given value.

### HasComponentNoticesId

`func (o *ApiLicenseLegalDataDTO) HasComponentNoticesId() bool`

HasComponentNoticesId returns a boolean if a field has been set.

### GetComponentNoticesLastUpdatedAt

`func (o *ApiLicenseLegalDataDTO) GetComponentNoticesLastUpdatedAt() time.Time`

GetComponentNoticesLastUpdatedAt returns the ComponentNoticesLastUpdatedAt field if non-nil, zero value otherwise.

### GetComponentNoticesLastUpdatedAtOk

`func (o *ApiLicenseLegalDataDTO) GetComponentNoticesLastUpdatedAtOk() (*time.Time, bool)`

GetComponentNoticesLastUpdatedAtOk returns a tuple with the ComponentNoticesLastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentNoticesLastUpdatedAt

`func (o *ApiLicenseLegalDataDTO) SetComponentNoticesLastUpdatedAt(v time.Time)`

SetComponentNoticesLastUpdatedAt sets ComponentNoticesLastUpdatedAt field to given value.

### HasComponentNoticesLastUpdatedAt

`func (o *ApiLicenseLegalDataDTO) HasComponentNoticesLastUpdatedAt() bool`

HasComponentNoticesLastUpdatedAt returns a boolean if a field has been set.

### GetComponentNoticesLastUpdatedByUsername

`func (o *ApiLicenseLegalDataDTO) GetComponentNoticesLastUpdatedByUsername() string`

GetComponentNoticesLastUpdatedByUsername returns the ComponentNoticesLastUpdatedByUsername field if non-nil, zero value otherwise.

### GetComponentNoticesLastUpdatedByUsernameOk

`func (o *ApiLicenseLegalDataDTO) GetComponentNoticesLastUpdatedByUsernameOk() (*string, bool)`

GetComponentNoticesLastUpdatedByUsernameOk returns a tuple with the ComponentNoticesLastUpdatedByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentNoticesLastUpdatedByUsername

`func (o *ApiLicenseLegalDataDTO) SetComponentNoticesLastUpdatedByUsername(v string)`

SetComponentNoticesLastUpdatedByUsername sets ComponentNoticesLastUpdatedByUsername field to given value.

### HasComponentNoticesLastUpdatedByUsername

`func (o *ApiLicenseLegalDataDTO) HasComponentNoticesLastUpdatedByUsername() bool`

HasComponentNoticesLastUpdatedByUsername returns a boolean if a field has been set.

### GetComponentNoticesScopeOwnerId

`func (o *ApiLicenseLegalDataDTO) GetComponentNoticesScopeOwnerId() string`

GetComponentNoticesScopeOwnerId returns the ComponentNoticesScopeOwnerId field if non-nil, zero value otherwise.

### GetComponentNoticesScopeOwnerIdOk

`func (o *ApiLicenseLegalDataDTO) GetComponentNoticesScopeOwnerIdOk() (*string, bool)`

GetComponentNoticesScopeOwnerIdOk returns a tuple with the ComponentNoticesScopeOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentNoticesScopeOwnerId

`func (o *ApiLicenseLegalDataDTO) SetComponentNoticesScopeOwnerId(v string)`

SetComponentNoticesScopeOwnerId sets ComponentNoticesScopeOwnerId field to given value.

### HasComponentNoticesScopeOwnerId

`func (o *ApiLicenseLegalDataDTO) HasComponentNoticesScopeOwnerId() bool`

HasComponentNoticesScopeOwnerId returns a boolean if a field has been set.

### GetCopyrights

`func (o *ApiLicenseLegalDataDTO) GetCopyrights() []ApiLicenseLegalCopyrightDTO`

GetCopyrights returns the Copyrights field if non-nil, zero value otherwise.

### GetCopyrightsOk

`func (o *ApiLicenseLegalDataDTO) GetCopyrightsOk() (*[]ApiLicenseLegalCopyrightDTO, bool)`

GetCopyrightsOk returns a tuple with the Copyrights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrights

`func (o *ApiLicenseLegalDataDTO) SetCopyrights(v []ApiLicenseLegalCopyrightDTO)`

SetCopyrights sets Copyrights field to given value.

### HasCopyrights

`func (o *ApiLicenseLegalDataDTO) HasCopyrights() bool`

HasCopyrights returns a boolean if a field has been set.

### GetDeclaredLicenses

`func (o *ApiLicenseLegalDataDTO) GetDeclaredLicenses() []string`

GetDeclaredLicenses returns the DeclaredLicenses field if non-nil, zero value otherwise.

### GetDeclaredLicensesOk

`func (o *ApiLicenseLegalDataDTO) GetDeclaredLicensesOk() (*[]string, bool)`

GetDeclaredLicensesOk returns a tuple with the DeclaredLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredLicenses

`func (o *ApiLicenseLegalDataDTO) SetDeclaredLicenses(v []string)`

SetDeclaredLicenses sets DeclaredLicenses field to given value.

### HasDeclaredLicenses

`func (o *ApiLicenseLegalDataDTO) HasDeclaredLicenses() bool`

HasDeclaredLicenses returns a boolean if a field has been set.

### GetEffectiveLicenseStatus

`func (o *ApiLicenseLegalDataDTO) GetEffectiveLicenseStatus() string`

GetEffectiveLicenseStatus returns the EffectiveLicenseStatus field if non-nil, zero value otherwise.

### GetEffectiveLicenseStatusOk

`func (o *ApiLicenseLegalDataDTO) GetEffectiveLicenseStatusOk() (*string, bool)`

GetEffectiveLicenseStatusOk returns a tuple with the EffectiveLicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveLicenseStatus

`func (o *ApiLicenseLegalDataDTO) SetEffectiveLicenseStatus(v string)`

SetEffectiveLicenseStatus sets EffectiveLicenseStatus field to given value.

### HasEffectiveLicenseStatus

`func (o *ApiLicenseLegalDataDTO) HasEffectiveLicenseStatus() bool`

HasEffectiveLicenseStatus returns a boolean if a field has been set.

### GetEffectiveLicenses

`func (o *ApiLicenseLegalDataDTO) GetEffectiveLicenses() []string`

GetEffectiveLicenses returns the EffectiveLicenses field if non-nil, zero value otherwise.

### GetEffectiveLicensesOk

`func (o *ApiLicenseLegalDataDTO) GetEffectiveLicensesOk() (*[]string, bool)`

GetEffectiveLicensesOk returns a tuple with the EffectiveLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveLicenses

`func (o *ApiLicenseLegalDataDTO) SetEffectiveLicenses(v []string)`

SetEffectiveLicenses sets EffectiveLicenses field to given value.

### HasEffectiveLicenses

`func (o *ApiLicenseLegalDataDTO) HasEffectiveLicenses() bool`

HasEffectiveLicenses returns a boolean if a field has been set.

### GetHighestEffectiveLicenseThreatGroup

`func (o *ApiLicenseLegalDataDTO) GetHighestEffectiveLicenseThreatGroup() ApiLicenseThreatDTOV2`

GetHighestEffectiveLicenseThreatGroup returns the HighestEffectiveLicenseThreatGroup field if non-nil, zero value otherwise.

### GetHighestEffectiveLicenseThreatGroupOk

`func (o *ApiLicenseLegalDataDTO) GetHighestEffectiveLicenseThreatGroupOk() (*ApiLicenseThreatDTOV2, bool)`

GetHighestEffectiveLicenseThreatGroupOk returns a tuple with the HighestEffectiveLicenseThreatGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestEffectiveLicenseThreatGroup

`func (o *ApiLicenseLegalDataDTO) SetHighestEffectiveLicenseThreatGroup(v ApiLicenseThreatDTOV2)`

SetHighestEffectiveLicenseThreatGroup sets HighestEffectiveLicenseThreatGroup field to given value.

### HasHighestEffectiveLicenseThreatGroup

`func (o *ApiLicenseLegalDataDTO) HasHighestEffectiveLicenseThreatGroup() bool`

HasHighestEffectiveLicenseThreatGroup returns a boolean if a field has been set.

### GetLicenseFiles

`func (o *ApiLicenseLegalDataDTO) GetLicenseFiles() []ApiLicenseLegalFileDTO`

GetLicenseFiles returns the LicenseFiles field if non-nil, zero value otherwise.

### GetLicenseFilesOk

`func (o *ApiLicenseLegalDataDTO) GetLicenseFilesOk() (*[]ApiLicenseLegalFileDTO, bool)`

GetLicenseFilesOk returns a tuple with the LicenseFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseFiles

`func (o *ApiLicenseLegalDataDTO) SetLicenseFiles(v []ApiLicenseLegalFileDTO)`

SetLicenseFiles sets LicenseFiles field to given value.

### HasLicenseFiles

`func (o *ApiLicenseLegalDataDTO) HasLicenseFiles() bool`

HasLicenseFiles returns a boolean if a field has been set.

### GetNoticeFiles

`func (o *ApiLicenseLegalDataDTO) GetNoticeFiles() []ApiLicenseLegalFileDTO`

GetNoticeFiles returns the NoticeFiles field if non-nil, zero value otherwise.

### GetNoticeFilesOk

`func (o *ApiLicenseLegalDataDTO) GetNoticeFilesOk() (*[]ApiLicenseLegalFileDTO, bool)`

GetNoticeFilesOk returns a tuple with the NoticeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoticeFiles

`func (o *ApiLicenseLegalDataDTO) SetNoticeFiles(v []ApiLicenseLegalFileDTO)`

SetNoticeFiles sets NoticeFiles field to given value.

### HasNoticeFiles

`func (o *ApiLicenseLegalDataDTO) HasNoticeFiles() bool`

HasNoticeFiles returns a boolean if a field has been set.

### GetObligations

`func (o *ApiLicenseLegalDataDTO) GetObligations() []ApiLicenseLegalObligationDTO`

GetObligations returns the Obligations field if non-nil, zero value otherwise.

### GetObligationsOk

`func (o *ApiLicenseLegalDataDTO) GetObligationsOk() (*[]ApiLicenseLegalObligationDTO, bool)`

GetObligationsOk returns a tuple with the Obligations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObligations

`func (o *ApiLicenseLegalDataDTO) SetObligations(v []ApiLicenseLegalObligationDTO)`

SetObligations sets Obligations field to given value.

### HasObligations

`func (o *ApiLicenseLegalDataDTO) HasObligations() bool`

HasObligations returns a boolean if a field has been set.

### GetObservedLicenses

`func (o *ApiLicenseLegalDataDTO) GetObservedLicenses() []string`

GetObservedLicenses returns the ObservedLicenses field if non-nil, zero value otherwise.

### GetObservedLicensesOk

`func (o *ApiLicenseLegalDataDTO) GetObservedLicensesOk() (*[]string, bool)`

GetObservedLicensesOk returns a tuple with the ObservedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedLicenses

`func (o *ApiLicenseLegalDataDTO) SetObservedLicenses(v []string)`

SetObservedLicenses sets ObservedLicenses field to given value.

### HasObservedLicenses

`func (o *ApiLicenseLegalDataDTO) HasObservedLicenses() bool`

HasObservedLicenses returns a boolean if a field has been set.

### GetSourceLinks

`func (o *ApiLicenseLegalDataDTO) GetSourceLinks() []LegalSourceLinkDTO`

GetSourceLinks returns the SourceLinks field if non-nil, zero value otherwise.

### GetSourceLinksOk

`func (o *ApiLicenseLegalDataDTO) GetSourceLinksOk() (*[]LegalSourceLinkDTO, bool)`

GetSourceLinksOk returns a tuple with the SourceLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLinks

`func (o *ApiLicenseLegalDataDTO) SetSourceLinks(v []LegalSourceLinkDTO)`

SetSourceLinks sets SourceLinks field to given value.

### HasSourceLinks

`func (o *ApiLicenseLegalDataDTO) HasSourceLinks() bool`

HasSourceLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


