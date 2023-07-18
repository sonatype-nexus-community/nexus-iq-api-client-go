# ApiReportComponentPolicyViolationsDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**DependencyData** | Pointer to [**ApiDependencyDataDTO**](ApiDependencyDataDTO.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**MatchState** | Pointer to **string** |  | [optional] 
**PackageUrl** | Pointer to **string** |  | [optional] 
**Pathnames** | Pointer to **[]string** |  | [optional] 
**Proprietary** | Pointer to **bool** |  | [optional] 
**Sha256** | Pointer to **string** |  | [optional] 
**ThirdParty** | Pointer to **bool** |  | [optional] 
**Violations** | Pointer to [**[]ApiReportPolicyViolationDTOV2**](ApiReportPolicyViolationDTOV2.md) |  | [optional] 

## Methods

### NewApiReportComponentPolicyViolationsDTOV2

`func NewApiReportComponentPolicyViolationsDTOV2() *ApiReportComponentPolicyViolationsDTOV2`

NewApiReportComponentPolicyViolationsDTOV2 instantiates a new ApiReportComponentPolicyViolationsDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiReportComponentPolicyViolationsDTOV2WithDefaults

`func NewApiReportComponentPolicyViolationsDTOV2WithDefaults() *ApiReportComponentPolicyViolationsDTOV2`

NewApiReportComponentPolicyViolationsDTOV2WithDefaults instantiates a new ApiReportComponentPolicyViolationsDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentIdentifier

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiReportComponentPolicyViolationsDTOV2) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiReportComponentPolicyViolationsDTOV2) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetDependencyData

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetDependencyData() ApiDependencyDataDTO`

GetDependencyData returns the DependencyData field if non-nil, zero value otherwise.

### GetDependencyDataOk

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetDependencyDataOk() (*ApiDependencyDataDTO, bool)`

GetDependencyDataOk returns a tuple with the DependencyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyData

`func (o *ApiReportComponentPolicyViolationsDTOV2) SetDependencyData(v ApiDependencyDataDTO)`

SetDependencyData sets DependencyData field to given value.

### HasDependencyData

`func (o *ApiReportComponentPolicyViolationsDTOV2) HasDependencyData() bool`

HasDependencyData returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiReportComponentPolicyViolationsDTOV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiReportComponentPolicyViolationsDTOV2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHash

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiReportComponentPolicyViolationsDTOV2) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiReportComponentPolicyViolationsDTOV2) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetMatchState

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetMatchState() string`

GetMatchState returns the MatchState field if non-nil, zero value otherwise.

### GetMatchStateOk

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetMatchStateOk() (*string, bool)`

GetMatchStateOk returns a tuple with the MatchState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchState

`func (o *ApiReportComponentPolicyViolationsDTOV2) SetMatchState(v string)`

SetMatchState sets MatchState field to given value.

### HasMatchState

`func (o *ApiReportComponentPolicyViolationsDTOV2) HasMatchState() bool`

HasMatchState returns a boolean if a field has been set.

### GetPackageUrl

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *ApiReportComponentPolicyViolationsDTOV2) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *ApiReportComponentPolicyViolationsDTOV2) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.

### GetPathnames

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetPathnames() []string`

GetPathnames returns the Pathnames field if non-nil, zero value otherwise.

### GetPathnamesOk

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetPathnamesOk() (*[]string, bool)`

GetPathnamesOk returns a tuple with the Pathnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathnames

`func (o *ApiReportComponentPolicyViolationsDTOV2) SetPathnames(v []string)`

SetPathnames sets Pathnames field to given value.

### HasPathnames

`func (o *ApiReportComponentPolicyViolationsDTOV2) HasPathnames() bool`

HasPathnames returns a boolean if a field has been set.

### GetProprietary

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetProprietary() bool`

GetProprietary returns the Proprietary field if non-nil, zero value otherwise.

### GetProprietaryOk

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetProprietaryOk() (*bool, bool)`

GetProprietaryOk returns a tuple with the Proprietary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietary

`func (o *ApiReportComponentPolicyViolationsDTOV2) SetProprietary(v bool)`

SetProprietary sets Proprietary field to given value.

### HasProprietary

`func (o *ApiReportComponentPolicyViolationsDTOV2) HasProprietary() bool`

HasProprietary returns a boolean if a field has been set.

### GetSha256

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ApiReportComponentPolicyViolationsDTOV2) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ApiReportComponentPolicyViolationsDTOV2) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetThirdParty

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetThirdParty() bool`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetThirdPartyOk() (*bool, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *ApiReportComponentPolicyViolationsDTOV2) SetThirdParty(v bool)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *ApiReportComponentPolicyViolationsDTOV2) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.

### GetViolations

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetViolations() []ApiReportPolicyViolationDTOV2`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *ApiReportComponentPolicyViolationsDTOV2) GetViolationsOk() (*[]ApiReportPolicyViolationDTOV2, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *ApiReportComponentPolicyViolationsDTOV2) SetViolations(v []ApiReportPolicyViolationDTOV2)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *ApiReportComponentPolicyViolationsDTOV2) HasViolations() bool`

HasViolations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


