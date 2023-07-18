# ApiLicenseLegalComponentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**LicenseLegalData** | Pointer to [**ApiLicenseLegalDataDTO**](ApiLicenseLegalDataDTO.md) |  | [optional] 
**PackageUrl** | Pointer to **string** |  | [optional] 
**StageScans** | Pointer to [**[]ApiLicenseLegalStageScanDTO**](ApiLicenseLegalStageScanDTO.md) |  | [optional] 

## Methods

### NewApiLicenseLegalComponentDTO

`func NewApiLicenseLegalComponentDTO() *ApiLicenseLegalComponentDTO`

NewApiLicenseLegalComponentDTO instantiates a new ApiLicenseLegalComponentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLicenseLegalComponentDTOWithDefaults

`func NewApiLicenseLegalComponentDTOWithDefaults() *ApiLicenseLegalComponentDTO`

NewApiLicenseLegalComponentDTOWithDefaults instantiates a new ApiLicenseLegalComponentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentIdentifier

`func (o *ApiLicenseLegalComponentDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiLicenseLegalComponentDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiLicenseLegalComponentDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiLicenseLegalComponentDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiLicenseLegalComponentDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiLicenseLegalComponentDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiLicenseLegalComponentDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiLicenseLegalComponentDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHash

`func (o *ApiLicenseLegalComponentDTO) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiLicenseLegalComponentDTO) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiLicenseLegalComponentDTO) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiLicenseLegalComponentDTO) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetLicenseLegalData

`func (o *ApiLicenseLegalComponentDTO) GetLicenseLegalData() ApiLicenseLegalDataDTO`

GetLicenseLegalData returns the LicenseLegalData field if non-nil, zero value otherwise.

### GetLicenseLegalDataOk

`func (o *ApiLicenseLegalComponentDTO) GetLicenseLegalDataOk() (*ApiLicenseLegalDataDTO, bool)`

GetLicenseLegalDataOk returns a tuple with the LicenseLegalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseLegalData

`func (o *ApiLicenseLegalComponentDTO) SetLicenseLegalData(v ApiLicenseLegalDataDTO)`

SetLicenseLegalData sets LicenseLegalData field to given value.

### HasLicenseLegalData

`func (o *ApiLicenseLegalComponentDTO) HasLicenseLegalData() bool`

HasLicenseLegalData returns a boolean if a field has been set.

### GetPackageUrl

`func (o *ApiLicenseLegalComponentDTO) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *ApiLicenseLegalComponentDTO) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *ApiLicenseLegalComponentDTO) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *ApiLicenseLegalComponentDTO) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.

### GetStageScans

`func (o *ApiLicenseLegalComponentDTO) GetStageScans() []ApiLicenseLegalStageScanDTO`

GetStageScans returns the StageScans field if non-nil, zero value otherwise.

### GetStageScansOk

`func (o *ApiLicenseLegalComponentDTO) GetStageScansOk() (*[]ApiLicenseLegalStageScanDTO, bool)`

GetStageScansOk returns a tuple with the StageScans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageScans

`func (o *ApiLicenseLegalComponentDTO) SetStageScans(v []ApiLicenseLegalStageScanDTO)`

SetStageScans sets StageScans field to given value.

### HasStageScans

`func (o *ApiLicenseLegalComponentDTO) HasStageScans() bool`

HasStageScans returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


