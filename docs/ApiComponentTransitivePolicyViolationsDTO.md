# ApiComponentTransitivePolicyViolationsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**IsInnerSource** | Pointer to **bool** |  | [optional] 
**PackageUrl** | Pointer to **string** |  | [optional] 
**TransitivePolicyViolations** | Pointer to [**[]ApiStagePolicyViolationComponentDTO**](ApiStagePolicyViolationComponentDTO.md) |  | [optional] 

## Methods

### NewApiComponentTransitivePolicyViolationsDTO

`func NewApiComponentTransitivePolicyViolationsDTO() *ApiComponentTransitivePolicyViolationsDTO`

NewApiComponentTransitivePolicyViolationsDTO instantiates a new ApiComponentTransitivePolicyViolationsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiComponentTransitivePolicyViolationsDTOWithDefaults

`func NewApiComponentTransitivePolicyViolationsDTOWithDefaults() *ApiComponentTransitivePolicyViolationsDTO`

NewApiComponentTransitivePolicyViolationsDTOWithDefaults instantiates a new ApiComponentTransitivePolicyViolationsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentIdentifier

`func (o *ApiComponentTransitivePolicyViolationsDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiComponentTransitivePolicyViolationsDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiComponentTransitivePolicyViolationsDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiComponentTransitivePolicyViolationsDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiComponentTransitivePolicyViolationsDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiComponentTransitivePolicyViolationsDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiComponentTransitivePolicyViolationsDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiComponentTransitivePolicyViolationsDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHash

`func (o *ApiComponentTransitivePolicyViolationsDTO) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiComponentTransitivePolicyViolationsDTO) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiComponentTransitivePolicyViolationsDTO) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiComponentTransitivePolicyViolationsDTO) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetIsInnerSource

`func (o *ApiComponentTransitivePolicyViolationsDTO) GetIsInnerSource() bool`

GetIsInnerSource returns the IsInnerSource field if non-nil, zero value otherwise.

### GetIsInnerSourceOk

`func (o *ApiComponentTransitivePolicyViolationsDTO) GetIsInnerSourceOk() (*bool, bool)`

GetIsInnerSourceOk returns a tuple with the IsInnerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInnerSource

`func (o *ApiComponentTransitivePolicyViolationsDTO) SetIsInnerSource(v bool)`

SetIsInnerSource sets IsInnerSource field to given value.

### HasIsInnerSource

`func (o *ApiComponentTransitivePolicyViolationsDTO) HasIsInnerSource() bool`

HasIsInnerSource returns a boolean if a field has been set.

### GetPackageUrl

`func (o *ApiComponentTransitivePolicyViolationsDTO) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *ApiComponentTransitivePolicyViolationsDTO) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *ApiComponentTransitivePolicyViolationsDTO) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *ApiComponentTransitivePolicyViolationsDTO) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.

### GetTransitivePolicyViolations

`func (o *ApiComponentTransitivePolicyViolationsDTO) GetTransitivePolicyViolations() []ApiStagePolicyViolationComponentDTO`

GetTransitivePolicyViolations returns the TransitivePolicyViolations field if non-nil, zero value otherwise.

### GetTransitivePolicyViolationsOk

`func (o *ApiComponentTransitivePolicyViolationsDTO) GetTransitivePolicyViolationsOk() (*[]ApiStagePolicyViolationComponentDTO, bool)`

GetTransitivePolicyViolationsOk returns a tuple with the TransitivePolicyViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransitivePolicyViolations

`func (o *ApiComponentTransitivePolicyViolationsDTO) SetTransitivePolicyViolations(v []ApiStagePolicyViolationComponentDTO)`

SetTransitivePolicyViolations sets TransitivePolicyViolations field to given value.

### HasTransitivePolicyViolations

`func (o *ApiComponentTransitivePolicyViolationsDTO) HasTransitivePolicyViolations() bool`

HasTransitivePolicyViolations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


