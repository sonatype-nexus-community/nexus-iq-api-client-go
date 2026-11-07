# ApiEcosystemConfigurationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowOverride** | Pointer to **bool** |  | [optional] 
**AllowOverrideInheritedFrom** | Pointer to [**OwnerRefDTO**](OwnerRefDTO.md) |  | [optional] 
**Axes** | Pointer to [**map[string]AxisValueDTO**](AxisValueDTO.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InheritanceCategoryIds** | Pointer to **[]string** |  | [optional] 
**InheritanceMode** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **string** |  | [optional] 
**PackageFormat** | Pointer to **string** |  | [optional] 
**PackageManagers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiEcosystemConfigurationDTO

`func NewApiEcosystemConfigurationDTO() *ApiEcosystemConfigurationDTO`

NewApiEcosystemConfigurationDTO instantiates a new ApiEcosystemConfigurationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEcosystemConfigurationDTOWithDefaults

`func NewApiEcosystemConfigurationDTOWithDefaults() *ApiEcosystemConfigurationDTO`

NewApiEcosystemConfigurationDTOWithDefaults instantiates a new ApiEcosystemConfigurationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowOverride

`func (o *ApiEcosystemConfigurationDTO) GetAllowOverride() bool`

GetAllowOverride returns the AllowOverride field if non-nil, zero value otherwise.

### GetAllowOverrideOk

`func (o *ApiEcosystemConfigurationDTO) GetAllowOverrideOk() (*bool, bool)`

GetAllowOverrideOk returns a tuple with the AllowOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverride

`func (o *ApiEcosystemConfigurationDTO) SetAllowOverride(v bool)`

SetAllowOverride sets AllowOverride field to given value.

### HasAllowOverride

`func (o *ApiEcosystemConfigurationDTO) HasAllowOverride() bool`

HasAllowOverride returns a boolean if a field has been set.

### GetAllowOverrideInheritedFrom

`func (o *ApiEcosystemConfigurationDTO) GetAllowOverrideInheritedFrom() OwnerRefDTO`

GetAllowOverrideInheritedFrom returns the AllowOverrideInheritedFrom field if non-nil, zero value otherwise.

### GetAllowOverrideInheritedFromOk

`func (o *ApiEcosystemConfigurationDTO) GetAllowOverrideInheritedFromOk() (*OwnerRefDTO, bool)`

GetAllowOverrideInheritedFromOk returns a tuple with the AllowOverrideInheritedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverrideInheritedFrom

`func (o *ApiEcosystemConfigurationDTO) SetAllowOverrideInheritedFrom(v OwnerRefDTO)`

SetAllowOverrideInheritedFrom sets AllowOverrideInheritedFrom field to given value.

### HasAllowOverrideInheritedFrom

`func (o *ApiEcosystemConfigurationDTO) HasAllowOverrideInheritedFrom() bool`

HasAllowOverrideInheritedFrom returns a boolean if a field has been set.

### GetAxes

`func (o *ApiEcosystemConfigurationDTO) GetAxes() map[string]AxisValueDTO`

GetAxes returns the Axes field if non-nil, zero value otherwise.

### GetAxesOk

`func (o *ApiEcosystemConfigurationDTO) GetAxesOk() (*map[string]AxisValueDTO, bool)`

GetAxesOk returns a tuple with the Axes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxes

`func (o *ApiEcosystemConfigurationDTO) SetAxes(v map[string]AxisValueDTO)`

SetAxes sets Axes field to given value.

### HasAxes

`func (o *ApiEcosystemConfigurationDTO) HasAxes() bool`

HasAxes returns a boolean if a field has been set.

### GetId

`func (o *ApiEcosystemConfigurationDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiEcosystemConfigurationDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiEcosystemConfigurationDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiEcosystemConfigurationDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceCategoryIds

`func (o *ApiEcosystemConfigurationDTO) GetInheritanceCategoryIds() []string`

GetInheritanceCategoryIds returns the InheritanceCategoryIds field if non-nil, zero value otherwise.

### GetInheritanceCategoryIdsOk

`func (o *ApiEcosystemConfigurationDTO) GetInheritanceCategoryIdsOk() (*[]string, bool)`

GetInheritanceCategoryIdsOk returns a tuple with the InheritanceCategoryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceCategoryIds

`func (o *ApiEcosystemConfigurationDTO) SetInheritanceCategoryIds(v []string)`

SetInheritanceCategoryIds sets InheritanceCategoryIds field to given value.

### HasInheritanceCategoryIds

`func (o *ApiEcosystemConfigurationDTO) HasInheritanceCategoryIds() bool`

HasInheritanceCategoryIds returns a boolean if a field has been set.

### GetInheritanceMode

`func (o *ApiEcosystemConfigurationDTO) GetInheritanceMode() string`

GetInheritanceMode returns the InheritanceMode field if non-nil, zero value otherwise.

### GetInheritanceModeOk

`func (o *ApiEcosystemConfigurationDTO) GetInheritanceModeOk() (*string, bool)`

GetInheritanceModeOk returns a tuple with the InheritanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceMode

`func (o *ApiEcosystemConfigurationDTO) SetInheritanceMode(v string)`

SetInheritanceMode sets InheritanceMode field to given value.

### HasInheritanceMode

`func (o *ApiEcosystemConfigurationDTO) HasInheritanceMode() bool`

HasInheritanceMode returns a boolean if a field has been set.

### GetOwnerId

`func (o *ApiEcosystemConfigurationDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ApiEcosystemConfigurationDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ApiEcosystemConfigurationDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ApiEcosystemConfigurationDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerType

`func (o *ApiEcosystemConfigurationDTO) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *ApiEcosystemConfigurationDTO) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *ApiEcosystemConfigurationDTO) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *ApiEcosystemConfigurationDTO) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetPackageFormat

`func (o *ApiEcosystemConfigurationDTO) GetPackageFormat() string`

GetPackageFormat returns the PackageFormat field if non-nil, zero value otherwise.

### GetPackageFormatOk

`func (o *ApiEcosystemConfigurationDTO) GetPackageFormatOk() (*string, bool)`

GetPackageFormatOk returns a tuple with the PackageFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageFormat

`func (o *ApiEcosystemConfigurationDTO) SetPackageFormat(v string)`

SetPackageFormat sets PackageFormat field to given value.

### HasPackageFormat

`func (o *ApiEcosystemConfigurationDTO) HasPackageFormat() bool`

HasPackageFormat returns a boolean if a field has been set.

### GetPackageManagers

`func (o *ApiEcosystemConfigurationDTO) GetPackageManagers() []string`

GetPackageManagers returns the PackageManagers field if non-nil, zero value otherwise.

### GetPackageManagersOk

`func (o *ApiEcosystemConfigurationDTO) GetPackageManagersOk() (*[]string, bool)`

GetPackageManagersOk returns a tuple with the PackageManagers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageManagers

`func (o *ApiEcosystemConfigurationDTO) SetPackageManagers(v []string)`

SetPackageManagers sets PackageManagers field to given value.

### HasPackageManagers

`func (o *ApiEcosystemConfigurationDTO) HasPackageManagers() bool`

HasPackageManagers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


