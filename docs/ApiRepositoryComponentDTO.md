# ApiRepositoryComponentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**PackageUrl** | Pointer to **string** |  | [optional] 
**QuarantineId** | Pointer to **string** |  | [optional] 
**QuarantineReleaseTime** | Pointer to **time.Time** |  | [optional] 
**QuarantineTime** | Pointer to **time.Time** |  | [optional] 
**Sha256** | Pointer to **string** |  | [optional] 
**ThirdParty** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiRepositoryComponentDTO

`func NewApiRepositoryComponentDTO() *ApiRepositoryComponentDTO`

NewApiRepositoryComponentDTO instantiates a new ApiRepositoryComponentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRepositoryComponentDTOWithDefaults

`func NewApiRepositoryComponentDTOWithDefaults() *ApiRepositoryComponentDTO`

NewApiRepositoryComponentDTOWithDefaults instantiates a new ApiRepositoryComponentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentIdentifier

`func (o *ApiRepositoryComponentDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiRepositoryComponentDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiRepositoryComponentDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiRepositoryComponentDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiRepositoryComponentDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiRepositoryComponentDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiRepositoryComponentDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiRepositoryComponentDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHash

`func (o *ApiRepositoryComponentDTO) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiRepositoryComponentDTO) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiRepositoryComponentDTO) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiRepositoryComponentDTO) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetPackageUrl

`func (o *ApiRepositoryComponentDTO) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *ApiRepositoryComponentDTO) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *ApiRepositoryComponentDTO) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *ApiRepositoryComponentDTO) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.

### GetQuarantineId

`func (o *ApiRepositoryComponentDTO) GetQuarantineId() string`

GetQuarantineId returns the QuarantineId field if non-nil, zero value otherwise.

### GetQuarantineIdOk

`func (o *ApiRepositoryComponentDTO) GetQuarantineIdOk() (*string, bool)`

GetQuarantineIdOk returns a tuple with the QuarantineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineId

`func (o *ApiRepositoryComponentDTO) SetQuarantineId(v string)`

SetQuarantineId sets QuarantineId field to given value.

### HasQuarantineId

`func (o *ApiRepositoryComponentDTO) HasQuarantineId() bool`

HasQuarantineId returns a boolean if a field has been set.

### GetQuarantineReleaseTime

`func (o *ApiRepositoryComponentDTO) GetQuarantineReleaseTime() time.Time`

GetQuarantineReleaseTime returns the QuarantineReleaseTime field if non-nil, zero value otherwise.

### GetQuarantineReleaseTimeOk

`func (o *ApiRepositoryComponentDTO) GetQuarantineReleaseTimeOk() (*time.Time, bool)`

GetQuarantineReleaseTimeOk returns a tuple with the QuarantineReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineReleaseTime

`func (o *ApiRepositoryComponentDTO) SetQuarantineReleaseTime(v time.Time)`

SetQuarantineReleaseTime sets QuarantineReleaseTime field to given value.

### HasQuarantineReleaseTime

`func (o *ApiRepositoryComponentDTO) HasQuarantineReleaseTime() bool`

HasQuarantineReleaseTime returns a boolean if a field has been set.

### GetQuarantineTime

`func (o *ApiRepositoryComponentDTO) GetQuarantineTime() time.Time`

GetQuarantineTime returns the QuarantineTime field if non-nil, zero value otherwise.

### GetQuarantineTimeOk

`func (o *ApiRepositoryComponentDTO) GetQuarantineTimeOk() (*time.Time, bool)`

GetQuarantineTimeOk returns a tuple with the QuarantineTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineTime

`func (o *ApiRepositoryComponentDTO) SetQuarantineTime(v time.Time)`

SetQuarantineTime sets QuarantineTime field to given value.

### HasQuarantineTime

`func (o *ApiRepositoryComponentDTO) HasQuarantineTime() bool`

HasQuarantineTime returns a boolean if a field has been set.

### GetSha256

`func (o *ApiRepositoryComponentDTO) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *ApiRepositoryComponentDTO) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *ApiRepositoryComponentDTO) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *ApiRepositoryComponentDTO) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetThirdParty

`func (o *ApiRepositoryComponentDTO) GetThirdParty() bool`

GetThirdParty returns the ThirdParty field if non-nil, zero value otherwise.

### GetThirdPartyOk

`func (o *ApiRepositoryComponentDTO) GetThirdPartyOk() (*bool, bool)`

GetThirdPartyOk returns a tuple with the ThirdParty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThirdParty

`func (o *ApiRepositoryComponentDTO) SetThirdParty(v bool)`

SetThirdParty sets ThirdParty field to given value.

### HasThirdParty

`func (o *ApiRepositoryComponentDTO) HasThirdParty() bool`

HasThirdParty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


