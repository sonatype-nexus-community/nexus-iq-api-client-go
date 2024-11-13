# SbomComponentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentIdentifier** | Pointer to [**ComponentIdentifier**](ComponentIdentifier.md) |  | [optional] 
**DependencyType** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**FileCoordinateId** | Pointer to **string** |  | [optional] 
**Filenames** | Pointer to **[]string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Licenses** | Pointer to [**[]License**](License.md) |  | [optional] 
**MatchStateId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PackageUrl** | Pointer to **string** |  | [optional] 
**PercentageAnnotated** | Pointer to **float64** |  | [optional] 
**PolicyViolationCount** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**VulnerabilitySeverityCriticalCount** | Pointer to **int32** |  | [optional] 
**VulnerabilitySeverityHighCount** | Pointer to **int32** |  | [optional] 
**VulnerabilitySeverityLowCount** | Pointer to **int32** |  | [optional] 
**VulnerabilitySeverityMediumCount** | Pointer to **int32** |  | [optional] 
**VulnerabilitySeverityNoneCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewSbomComponentDTO

`func NewSbomComponentDTO() *SbomComponentDTO`

NewSbomComponentDTO instantiates a new SbomComponentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSbomComponentDTOWithDefaults

`func NewSbomComponentDTOWithDefaults() *SbomComponentDTO`

NewSbomComponentDTOWithDefaults instantiates a new SbomComponentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentIdentifier

`func (o *SbomComponentDTO) GetComponentIdentifier() ComponentIdentifier`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *SbomComponentDTO) GetComponentIdentifierOk() (*ComponentIdentifier, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *SbomComponentDTO) SetComponentIdentifier(v ComponentIdentifier)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *SbomComponentDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetDependencyType

`func (o *SbomComponentDTO) GetDependencyType() string`

GetDependencyType returns the DependencyType field if non-nil, zero value otherwise.

### GetDependencyTypeOk

`func (o *SbomComponentDTO) GetDependencyTypeOk() (*string, bool)`

GetDependencyTypeOk returns a tuple with the DependencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyType

`func (o *SbomComponentDTO) SetDependencyType(v string)`

SetDependencyType sets DependencyType field to given value.

### HasDependencyType

`func (o *SbomComponentDTO) HasDependencyType() bool`

HasDependencyType returns a boolean if a field has been set.

### GetDisplayName

`func (o *SbomComponentDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SbomComponentDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SbomComponentDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SbomComponentDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFileCoordinateId

`func (o *SbomComponentDTO) GetFileCoordinateId() string`

GetFileCoordinateId returns the FileCoordinateId field if non-nil, zero value otherwise.

### GetFileCoordinateIdOk

`func (o *SbomComponentDTO) GetFileCoordinateIdOk() (*string, bool)`

GetFileCoordinateIdOk returns a tuple with the FileCoordinateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCoordinateId

`func (o *SbomComponentDTO) SetFileCoordinateId(v string)`

SetFileCoordinateId sets FileCoordinateId field to given value.

### HasFileCoordinateId

`func (o *SbomComponentDTO) HasFileCoordinateId() bool`

HasFileCoordinateId returns a boolean if a field has been set.

### GetFilenames

`func (o *SbomComponentDTO) GetFilenames() []string`

GetFilenames returns the Filenames field if non-nil, zero value otherwise.

### GetFilenamesOk

`func (o *SbomComponentDTO) GetFilenamesOk() (*[]string, bool)`

GetFilenamesOk returns a tuple with the Filenames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilenames

`func (o *SbomComponentDTO) SetFilenames(v []string)`

SetFilenames sets Filenames field to given value.

### HasFilenames

`func (o *SbomComponentDTO) HasFilenames() bool`

HasFilenames returns a boolean if a field has been set.

### GetHash

`func (o *SbomComponentDTO) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SbomComponentDTO) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SbomComponentDTO) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *SbomComponentDTO) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetLicenses

`func (o *SbomComponentDTO) GetLicenses() []License`

GetLicenses returns the Licenses field if non-nil, zero value otherwise.

### GetLicensesOk

`func (o *SbomComponentDTO) GetLicensesOk() (*[]License, bool)`

GetLicensesOk returns a tuple with the Licenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenses

`func (o *SbomComponentDTO) SetLicenses(v []License)`

SetLicenses sets Licenses field to given value.

### HasLicenses

`func (o *SbomComponentDTO) HasLicenses() bool`

HasLicenses returns a boolean if a field has been set.

### GetMatchStateId

`func (o *SbomComponentDTO) GetMatchStateId() string`

GetMatchStateId returns the MatchStateId field if non-nil, zero value otherwise.

### GetMatchStateIdOk

`func (o *SbomComponentDTO) GetMatchStateIdOk() (*string, bool)`

GetMatchStateIdOk returns a tuple with the MatchStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchStateId

`func (o *SbomComponentDTO) SetMatchStateId(v string)`

SetMatchStateId sets MatchStateId field to given value.

### HasMatchStateId

`func (o *SbomComponentDTO) HasMatchStateId() bool`

HasMatchStateId returns a boolean if a field has been set.

### GetName

`func (o *SbomComponentDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SbomComponentDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SbomComponentDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SbomComponentDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPackageUrl

`func (o *SbomComponentDTO) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *SbomComponentDTO) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *SbomComponentDTO) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *SbomComponentDTO) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.

### GetPercentageAnnotated

`func (o *SbomComponentDTO) GetPercentageAnnotated() float64`

GetPercentageAnnotated returns the PercentageAnnotated field if non-nil, zero value otherwise.

### GetPercentageAnnotatedOk

`func (o *SbomComponentDTO) GetPercentageAnnotatedOk() (*float64, bool)`

GetPercentageAnnotatedOk returns a tuple with the PercentageAnnotated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageAnnotated

`func (o *SbomComponentDTO) SetPercentageAnnotated(v float64)`

SetPercentageAnnotated sets PercentageAnnotated field to given value.

### HasPercentageAnnotated

`func (o *SbomComponentDTO) HasPercentageAnnotated() bool`

HasPercentageAnnotated returns a boolean if a field has been set.

### GetPolicyViolationCount

`func (o *SbomComponentDTO) GetPolicyViolationCount() int32`

GetPolicyViolationCount returns the PolicyViolationCount field if non-nil, zero value otherwise.

### GetPolicyViolationCountOk

`func (o *SbomComponentDTO) GetPolicyViolationCountOk() (*int32, bool)`

GetPolicyViolationCountOk returns a tuple with the PolicyViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationCount

`func (o *SbomComponentDTO) SetPolicyViolationCount(v int32)`

SetPolicyViolationCount sets PolicyViolationCount field to given value.

### HasPolicyViolationCount

`func (o *SbomComponentDTO) HasPolicyViolationCount() bool`

HasPolicyViolationCount returns a boolean if a field has been set.

### GetVersion

`func (o *SbomComponentDTO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SbomComponentDTO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SbomComponentDTO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SbomComponentDTO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVulnerabilitySeverityCriticalCount

`func (o *SbomComponentDTO) GetVulnerabilitySeverityCriticalCount() int32`

GetVulnerabilitySeverityCriticalCount returns the VulnerabilitySeverityCriticalCount field if non-nil, zero value otherwise.

### GetVulnerabilitySeverityCriticalCountOk

`func (o *SbomComponentDTO) GetVulnerabilitySeverityCriticalCountOk() (*int32, bool)`

GetVulnerabilitySeverityCriticalCountOk returns a tuple with the VulnerabilitySeverityCriticalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilitySeverityCriticalCount

`func (o *SbomComponentDTO) SetVulnerabilitySeverityCriticalCount(v int32)`

SetVulnerabilitySeverityCriticalCount sets VulnerabilitySeverityCriticalCount field to given value.

### HasVulnerabilitySeverityCriticalCount

`func (o *SbomComponentDTO) HasVulnerabilitySeverityCriticalCount() bool`

HasVulnerabilitySeverityCriticalCount returns a boolean if a field has been set.

### GetVulnerabilitySeverityHighCount

`func (o *SbomComponentDTO) GetVulnerabilitySeverityHighCount() int32`

GetVulnerabilitySeverityHighCount returns the VulnerabilitySeverityHighCount field if non-nil, zero value otherwise.

### GetVulnerabilitySeverityHighCountOk

`func (o *SbomComponentDTO) GetVulnerabilitySeverityHighCountOk() (*int32, bool)`

GetVulnerabilitySeverityHighCountOk returns a tuple with the VulnerabilitySeverityHighCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilitySeverityHighCount

`func (o *SbomComponentDTO) SetVulnerabilitySeverityHighCount(v int32)`

SetVulnerabilitySeverityHighCount sets VulnerabilitySeverityHighCount field to given value.

### HasVulnerabilitySeverityHighCount

`func (o *SbomComponentDTO) HasVulnerabilitySeverityHighCount() bool`

HasVulnerabilitySeverityHighCount returns a boolean if a field has been set.

### GetVulnerabilitySeverityLowCount

`func (o *SbomComponentDTO) GetVulnerabilitySeverityLowCount() int32`

GetVulnerabilitySeverityLowCount returns the VulnerabilitySeverityLowCount field if non-nil, zero value otherwise.

### GetVulnerabilitySeverityLowCountOk

`func (o *SbomComponentDTO) GetVulnerabilitySeverityLowCountOk() (*int32, bool)`

GetVulnerabilitySeverityLowCountOk returns a tuple with the VulnerabilitySeverityLowCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilitySeverityLowCount

`func (o *SbomComponentDTO) SetVulnerabilitySeverityLowCount(v int32)`

SetVulnerabilitySeverityLowCount sets VulnerabilitySeverityLowCount field to given value.

### HasVulnerabilitySeverityLowCount

`func (o *SbomComponentDTO) HasVulnerabilitySeverityLowCount() bool`

HasVulnerabilitySeverityLowCount returns a boolean if a field has been set.

### GetVulnerabilitySeverityMediumCount

`func (o *SbomComponentDTO) GetVulnerabilitySeverityMediumCount() int32`

GetVulnerabilitySeverityMediumCount returns the VulnerabilitySeverityMediumCount field if non-nil, zero value otherwise.

### GetVulnerabilitySeverityMediumCountOk

`func (o *SbomComponentDTO) GetVulnerabilitySeverityMediumCountOk() (*int32, bool)`

GetVulnerabilitySeverityMediumCountOk returns a tuple with the VulnerabilitySeverityMediumCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilitySeverityMediumCount

`func (o *SbomComponentDTO) SetVulnerabilitySeverityMediumCount(v int32)`

SetVulnerabilitySeverityMediumCount sets VulnerabilitySeverityMediumCount field to given value.

### HasVulnerabilitySeverityMediumCount

`func (o *SbomComponentDTO) HasVulnerabilitySeverityMediumCount() bool`

HasVulnerabilitySeverityMediumCount returns a boolean if a field has been set.

### GetVulnerabilitySeverityNoneCount

`func (o *SbomComponentDTO) GetVulnerabilitySeverityNoneCount() int32`

GetVulnerabilitySeverityNoneCount returns the VulnerabilitySeverityNoneCount field if non-nil, zero value otherwise.

### GetVulnerabilitySeverityNoneCountOk

`func (o *SbomComponentDTO) GetVulnerabilitySeverityNoneCountOk() (*int32, bool)`

GetVulnerabilitySeverityNoneCountOk returns a tuple with the VulnerabilitySeverityNoneCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilitySeverityNoneCount

`func (o *SbomComponentDTO) SetVulnerabilitySeverityNoneCount(v int32)`

SetVulnerabilitySeverityNoneCount sets VulnerabilitySeverityNoneCount field to given value.

### HasVulnerabilitySeverityNoneCount

`func (o *SbomComponentDTO) HasVulnerabilitySeverityNoneCount() bool`

HasVulnerabilitySeverityNoneCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


