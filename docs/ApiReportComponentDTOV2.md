# ApiReportComponentDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**Cpe** | Pointer to **string** |  | [optional] 
**DependencyData** | Pointer to [**ApiDependencyDataDTO**](ApiDependencyDataDTO.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Filenames** | Pointer to **[]string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**IdentificationSource** | Pointer to **string** |  | [optional] 
**LicenseData** | Pointer to [**ApiLicenseDataDTOV2**](ApiLicenseDataDTOV2.md) |  | [optional] 
**MatchState** | Pointer to **string** |  | [optional] 
**PackageUrl** | Pointer to **string** |  | [optional] 
**Pathnames** | Pointer to **[]string** |  | [optional] 
**Proprietary** | Pointer to **bool** |  | [optional] 
**SecurityData** | Pointer to [**ApiSecurityDataDTO**](ApiSecurityDataDTO.md) |  | [optional] 
**Sha256** | Pointer to **string** |  | [optional] 
**Swid** | Pointer to [**Swid**](Swid.md) |  | [optional] 
**ThirdParty** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiReportComponentDTOV2

`func NewApiReportComponentDTOV2() *ApiReportComponentDTOV2`

NewApiReportComponentDTOV2 instantiates a new ApiReportComponentDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiReportComponentDTOV2WithDefaults

`func NewApiReportComponentDTOV2WithDefaults() *ApiReportComponentDTOV2`

NewApiReportComponentDTOV2WithDefaults instantiates a new ApiReportComponentDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentIdentifier

`func (o *ApiReportComponentDTOV2) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiReportComponentDTOV2) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiReportComponentDTOV2) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiReportComponentDTOV2) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetCpe

`func (o *ApiReportComponentDTOV2) GetCpe() string`

GetCpe returns the Cpe field if non-nil, zero value otherwise.

### GetCpeOk

`func (o *ApiReportComponentDTOV2) GetCpeOk() (*string, bool)`

GetCpeOk returns a tuple with the Cpe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpe

`func (o *ApiReportComponentDTOV2) SetCpe(v string)`

SetCpe sets Cpe field to given value.

### HasCpe

`func (o *ApiReportComponentDTOV2) HasCpe() bool`

HasCpe returns a boolean if a field has been set.

### GetDependencyData

`func (o *ApiReportComponentDTOV2) GetDependencyData() ApiDependencyDataDTO`

GetDependencyData returns the DependencyData field if non-nil, zero value otherwise.

### GetDependencyDataOk

`func (o *ApiReportComponentDTOV2) GetDependencyDataOk() (*ApiDependencyDataDTO, bool)`

GetDependencyDataOk returns a tuple with the DependencyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyData

`func (o *ApiReportComponentDTOV2) SetDependencyData(v ApiDependencyDataDTO)`

SetDependencyData sets DependencyData field to given value.

### HasDependencyData

`func (o *ApiReportComponentDTOV2) HasDependencyData() bool`

HasDependencyData returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiReportComponentDTOV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiReportComponentDTOV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiReportComponentDTOV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiReportComponentDTOV2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFilenames

`func (o *ApiReportComponentDTOV2) GetFilenames() []string`

GetFilenames returns the Filenames field if non-nil, zero value otherwise.

### GetFilenamesOk

`func (o *ApiReportComponentDTOV2) GetFilenamesOk() (*[]string, bool)`

GetFilenamesOk returns a tuple with the Filenames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenames

`func (o *ApiReportComponentDTOV2) SetFilenames(v []string)`

SetFilenames sets Filenames field to given value.

### HasFilenames

`func (o *ApiReportComponentDTOV2) HasFilenames() bool`

HasFilenames returns a boolean if a field has been set.

### GetHash

`func (o *ApiReportComponentDTOV2) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiReportComponentDTOV2) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiReportComponentDTOV2) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiReportComponentDTOV2) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetIdentificationSource

`func (o *ApiReportComponentDTOV2) GetIdentificationSource() string`

GetIdentificationSource returns the IdentificationSource field if non-nil, zero value otherwise.

### GetIdentificationSourceOk

`func (o *ApiReportComponentDTOV2) GetIdentificationSourceOk() (*string, bool)`

GetIdentificationSourceOk returns a tuple with the IdentificationSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentificationSource

`func (o *ApiReportComponentDTOV2) SetIdentificationSource(v string)`

SetIdentificationSource sets IdentificationSource field to given value.

### HasIdentificationSource

`func (o *ApiReportComponentDTOV2) HasIdentificationSource() bool`

HasIdentificationSource returns a boolean if a field has been set.

### GetLicenseData

`func (o *ApiReportComponentDTOV2) GetLicenseData() ApiLicenseDataDTOV2`

GetLicenseData returns the LicenseData field if non-nil, zero value otherwise.

### GetLicenseDataOk

`func (o *ApiReportComponentDTOV2) GetLicenseDataOk() (*ApiLicenseDataDTOV2, bool)`

GetLicenseDataOk returns a tuple with the LicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseData

`func (o *ApiReportComponentDTOV2) SetLicenseData(v ApiLicenseDataDTOV2)`

SetLicenseData sets LicenseData field to given value.

### HasLicenseData

`func (o *ApiReportComponentDTOV2) HasLicenseData() bool`

HasLicenseData returns a boolean if a field has been set.

### GetMatchState

`func (o *ApiReportComponentDTOV2) GetMatchState() string`

GetMatchState returns the MatchState field if non-nil, zero value otherwise.

### GetMatchStateOk

`func (o *ApiReportComponentDTOV2) GetMatchStateOk() (*string, bool)`

GetMatchStateOk returns a tuple with the MatchState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchState

`func (o *ApiReportComponentDTOV2) SetMatchState(v string)`

SetMatchState sets MatchState field to given value.

### HasMatchState

`func (o *ApiReportComponentDTOV2) HasMatchState() bool`

HasMatchState returns a boolean if a field has been set.

### GetPackageUrl

`func (o *ApiReportComponentDTOV2) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *ApiReportComponentDTOV2) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *ApiReportComponentDTOV2) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *ApiReportComponentDTOV2) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.

### GetPathnames

`func (o *ApiReportComponentDTOV2) GetPathnames() []string`

GetPathnames returns the Pathnames field if non-nil, zero value otherwise.

### GetPathnamesOk

`func (o *ApiReportComponentDTOV2) GetPathnamesOk() (*[]string, bool)`

GetPathnamesOk returns a tuple with the Pathnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathnames

`func (o *ApiReportComponentDTOV2) SetPathnames(v []string)`

SetPathnames sets Pathnames field to given value.

### HasPathnames

`func (o *ApiReportComponentDTOV2) HasPathnames() bool`

HasPathnames returns a boolean if a field has been set.

### GetProprietary

`func (o *ApiReportComponentDTOV2) GetProprietary() bool`

GetProprietary returns the Proprietary field if non-nil, zero value otherwise.

### GetProprietaryOk

`func (o *ApiReportComponentDTOV2) GetProprietaryOk() (*bool, bool)`

GetProprietaryOk returns a tuple with the Proprietary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProprietary

`func (o *ApiReportComponentDTOV2) SetProprietary(v bool)`

SetProprietary sets Proprietary field to given value.

### HasProprietary

`func (o *ApiReportComponentDTOV2) HasProprietary() bool`

HasProprietary returns a boolean if a field has been set.

### GetSecurityData

`func (o *ApiReportComponentDTOV2) GetSecurityData() ApiSecurityDataDTO`

GetSecurityData returns the SecurityData field if non-nil, zero value otherwise.

### GetSecurityDataOk

`func (o *ApiReportComponentDTOV2) GetSecurityDataOk() (*ApiSecurityDataDTO, bool)`

GetSecurityDataOk returns a tuple with the SecurityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityData

`func (o *ApiReportComponentDTOV2) SetSecurityData(v ApiSecurityDataDTO)`

SetSecurityData sets SecurityData field to given value.

### HasSecurityData

`func (o *ApiReportComponentDTOV2) HasSecurityData() bool`

HasSecurityData returns a boolean if a field has been set.

### GetSha256

`func (o *ApiReportComponentDTOV2) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ApiReportComponentDTOV2) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ApiReportComponentDTOV2) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ApiReportComponentDTOV2) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetSwid

`func (o *ApiReportComponentDTOV2) GetSwid() Swid`

GetSwid returns the Swid field if non-nil, zero value otherwise.

### GetSwidOk

`func (o *ApiReportComponentDTOV2) GetSwidOk() (*Swid, bool)`

GetSwidOk returns a tuple with the Swid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwid

`func (o *ApiReportComponentDTOV2) SetSwid(v Swid)`

SetSwid sets Swid field to given value.

### HasSwid

`func (o *ApiReportComponentDTOV2) HasSwid() bool`

HasSwid returns a boolean if a field has been set.

### GetThirdParty

`func (o *ApiReportComponentDTOV2) GetThirdParty() bool`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *ApiReportComponentDTOV2) GetThirdPartyOk() (*bool, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *ApiReportComponentDTOV2) SetThirdParty(v bool)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *ApiReportComponentDTOV2) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


