# ApiComponentDetailsDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogDate** | Pointer to **time.Time** |  | [optional] 
**Component** | Pointer to [**ApiComponentDTOV2**](ApiComponentDTOV2.md) |  | [optional] 
**HygieneRating** | Pointer to **NullableString** |  | [optional] 
**IntegrityRating** | Pointer to **NullableString** |  | [optional] 
**LicenseData** | Pointer to [**ApiLicenseDataDTO**](ApiLicenseDataDTO.md) |  | [optional] 
**MatchState** | Pointer to **string** |  | [optional] 
**PolicyData** | Pointer to [**ApiComponentPolicyViolationListDTOV2**](ApiComponentPolicyViolationListDTOV2.md) |  | [optional] 
**ProjectData** | Pointer to [**ApiComponentProjectDataDTO**](ApiComponentProjectDataDTO.md) |  | [optional] 
**RelativePopularity** | Pointer to **NullableInt32** |  | [optional] 
**SecurityData** | Pointer to [**ApiSecurityDataDTO**](ApiSecurityDataDTO.md) |  | [optional] 

## Methods

### NewApiComponentDetailsDTOV2

`func NewApiComponentDetailsDTOV2() *ApiComponentDetailsDTOV2`

NewApiComponentDetailsDTOV2 instantiates a new ApiComponentDetailsDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiComponentDetailsDTOV2WithDefaults

`func NewApiComponentDetailsDTOV2WithDefaults() *ApiComponentDetailsDTOV2`

NewApiComponentDetailsDTOV2WithDefaults instantiates a new ApiComponentDetailsDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogDate

`func (o *ApiComponentDetailsDTOV2) GetCatalogDate() time.Time`

GetCatalogDate returns the CatalogDate field if non-nil, zero value otherwise.

### GetCatalogDateOk

`func (o *ApiComponentDetailsDTOV2) GetCatalogDateOk() (*time.Time, bool)`

GetCatalogDateOk returns a tuple with the CatalogDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogDate

`func (o *ApiComponentDetailsDTOV2) SetCatalogDate(v time.Time)`

SetCatalogDate sets CatalogDate field to given value.

### HasCatalogDate

`func (o *ApiComponentDetailsDTOV2) HasCatalogDate() bool`

HasCatalogDate returns a boolean if a field has been set.

### GetComponent

`func (o *ApiComponentDetailsDTOV2) GetComponent() ApiComponentDTOV2`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ApiComponentDetailsDTOV2) GetComponentOk() (*ApiComponentDTOV2, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ApiComponentDetailsDTOV2) SetComponent(v ApiComponentDTOV2)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ApiComponentDetailsDTOV2) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetHygieneRating

`func (o *ApiComponentDetailsDTOV2) GetHygieneRating() string`

GetHygieneRating returns the HygieneRating field if non-nil, zero value otherwise.

### GetHygieneRatingOk

`func (o *ApiComponentDetailsDTOV2) GetHygieneRatingOk() (*string, bool)`

GetHygieneRatingOk returns a tuple with the HygieneRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHygieneRating

`func (o *ApiComponentDetailsDTOV2) SetHygieneRating(v string)`

SetHygieneRating sets HygieneRating field to given value.

### HasHygieneRating

`func (o *ApiComponentDetailsDTOV2) HasHygieneRating() bool`

HasHygieneRating returns a boolean if a field has been set.

### SetHygieneRatingNil

`func (o *ApiComponentDetailsDTOV2) SetHygieneRatingNil(b bool)`

 SetHygieneRatingNil sets the value for HygieneRating to be an explicit nil

### UnsetHygieneRating
`func (o *ApiComponentDetailsDTOV2) UnsetHygieneRating()`

UnsetHygieneRating ensures that no value is present for HygieneRating, not even an explicit nil
### GetIntegrityRating

`func (o *ApiComponentDetailsDTOV2) GetIntegrityRating() string`

GetIntegrityRating returns the IntegrityRating field if non-nil, zero value otherwise.

### GetIntegrityRatingOk

`func (o *ApiComponentDetailsDTOV2) GetIntegrityRatingOk() (*string, bool)`

GetIntegrityRatingOk returns a tuple with the IntegrityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityRating

`func (o *ApiComponentDetailsDTOV2) SetIntegrityRating(v string)`

SetIntegrityRating sets IntegrityRating field to given value.

### HasIntegrityRating

`func (o *ApiComponentDetailsDTOV2) HasIntegrityRating() bool`

HasIntegrityRating returns a boolean if a field has been set.

### SetIntegrityRatingNil

`func (o *ApiComponentDetailsDTOV2) SetIntegrityRatingNil(b bool)`

 SetIntegrityRatingNil sets the value for IntegrityRating to be an explicit nil

### UnsetIntegrityRating
`func (o *ApiComponentDetailsDTOV2) UnsetIntegrityRating()`

UnsetIntegrityRating ensures that no value is present for IntegrityRating, not even an explicit nil
### GetLicenseData

`func (o *ApiComponentDetailsDTOV2) GetLicenseData() ApiLicenseDataDTO`

GetLicenseData returns the LicenseData field if non-nil, zero value otherwise.

### GetLicenseDataOk

`func (o *ApiComponentDetailsDTOV2) GetLicenseDataOk() (*ApiLicenseDataDTO, bool)`

GetLicenseDataOk returns a tuple with the LicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseData

`func (o *ApiComponentDetailsDTOV2) SetLicenseData(v ApiLicenseDataDTO)`

SetLicenseData sets LicenseData field to given value.

### HasLicenseData

`func (o *ApiComponentDetailsDTOV2) HasLicenseData() bool`

HasLicenseData returns a boolean if a field has been set.

### GetMatchState

`func (o *ApiComponentDetailsDTOV2) GetMatchState() string`

GetMatchState returns the MatchState field if non-nil, zero value otherwise.

### GetMatchStateOk

`func (o *ApiComponentDetailsDTOV2) GetMatchStateOk() (*string, bool)`

GetMatchStateOk returns a tuple with the MatchState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchState

`func (o *ApiComponentDetailsDTOV2) SetMatchState(v string)`

SetMatchState sets MatchState field to given value.

### HasMatchState

`func (o *ApiComponentDetailsDTOV2) HasMatchState() bool`

HasMatchState returns a boolean if a field has been set.

### GetPolicyData

`func (o *ApiComponentDetailsDTOV2) GetPolicyData() ApiComponentPolicyViolationListDTOV2`

GetPolicyData returns the PolicyData field if non-nil, zero value otherwise.

### GetPolicyDataOk

`func (o *ApiComponentDetailsDTOV2) GetPolicyDataOk() (*ApiComponentPolicyViolationListDTOV2, bool)`

GetPolicyDataOk returns a tuple with the PolicyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyData

`func (o *ApiComponentDetailsDTOV2) SetPolicyData(v ApiComponentPolicyViolationListDTOV2)`

SetPolicyData sets PolicyData field to given value.

### HasPolicyData

`func (o *ApiComponentDetailsDTOV2) HasPolicyData() bool`

HasPolicyData returns a boolean if a field has been set.

### GetProjectData

`func (o *ApiComponentDetailsDTOV2) GetProjectData() ApiComponentProjectDataDTO`

GetProjectData returns the ProjectData field if non-nil, zero value otherwise.

### GetProjectDataOk

`func (o *ApiComponentDetailsDTOV2) GetProjectDataOk() (*ApiComponentProjectDataDTO, bool)`

GetProjectDataOk returns a tuple with the ProjectData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectData

`func (o *ApiComponentDetailsDTOV2) SetProjectData(v ApiComponentProjectDataDTO)`

SetProjectData sets ProjectData field to given value.

### HasProjectData

`func (o *ApiComponentDetailsDTOV2) HasProjectData() bool`

HasProjectData returns a boolean if a field has been set.

### GetRelativePopularity

`func (o *ApiComponentDetailsDTOV2) GetRelativePopularity() int32`

GetRelativePopularity returns the RelativePopularity field if non-nil, zero value otherwise.

### GetRelativePopularityOk

`func (o *ApiComponentDetailsDTOV2) GetRelativePopularityOk() (*int32, bool)`

GetRelativePopularityOk returns a tuple with the RelativePopularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePopularity

`func (o *ApiComponentDetailsDTOV2) SetRelativePopularity(v int32)`

SetRelativePopularity sets RelativePopularity field to given value.

### HasRelativePopularity

`func (o *ApiComponentDetailsDTOV2) HasRelativePopularity() bool`

HasRelativePopularity returns a boolean if a field has been set.

### SetRelativePopularityNil

`func (o *ApiComponentDetailsDTOV2) SetRelativePopularityNil(b bool)`

 SetRelativePopularityNil sets the value for RelativePopularity to be an explicit nil

### UnsetRelativePopularity
`func (o *ApiComponentDetailsDTOV2) UnsetRelativePopularity()`

UnsetRelativePopularity ensures that no value is present for RelativePopularity, not even an explicit nil
### GetSecurityData

`func (o *ApiComponentDetailsDTOV2) GetSecurityData() ApiSecurityDataDTO`

GetSecurityData returns the SecurityData field if non-nil, zero value otherwise.

### GetSecurityDataOk

`func (o *ApiComponentDetailsDTOV2) GetSecurityDataOk() (*ApiSecurityDataDTO, bool)`

GetSecurityDataOk returns a tuple with the SecurityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityData

`func (o *ApiComponentDetailsDTOV2) SetSecurityData(v ApiSecurityDataDTO)`

SetSecurityData sets SecurityData field to given value.

### HasSecurityData

`func (o *ApiComponentDetailsDTOV2) HasSecurityData() bool`

HasSecurityData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


