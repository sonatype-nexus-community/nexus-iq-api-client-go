# ApiDependencyTreeNodeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]ApiDependencyTreeNodeDTO**](ApiDependencyTreeNodeDTO.md) |  | [optional] 
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**Direct** | Pointer to **bool** |  | [optional] 
**PackageUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewApiDependencyTreeNodeDTO

`func NewApiDependencyTreeNodeDTO() *ApiDependencyTreeNodeDTO`

NewApiDependencyTreeNodeDTO instantiates a new ApiDependencyTreeNodeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDependencyTreeNodeDTOWithDefaults

`func NewApiDependencyTreeNodeDTOWithDefaults() *ApiDependencyTreeNodeDTO`

NewApiDependencyTreeNodeDTOWithDefaults instantiates a new ApiDependencyTreeNodeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *ApiDependencyTreeNodeDTO) GetChildren() []ApiDependencyTreeNodeDTO`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ApiDependencyTreeNodeDTO) GetChildrenOk() (*[]ApiDependencyTreeNodeDTO, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ApiDependencyTreeNodeDTO) SetChildren(v []ApiDependencyTreeNodeDTO)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ApiDependencyTreeNodeDTO) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *ApiDependencyTreeNodeDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiDependencyTreeNodeDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiDependencyTreeNodeDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiDependencyTreeNodeDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetDirect

`func (o *ApiDependencyTreeNodeDTO) GetDirect() bool`

GetDirect returns the Direct field if non-nil, zero value otherwise.

### GetDirectOk

`func (o *ApiDependencyTreeNodeDTO) GetDirectOk() (*bool, bool)`

GetDirectOk returns a tuple with the Direct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirect

`func (o *ApiDependencyTreeNodeDTO) SetDirect(v bool)`

SetDirect sets Direct field to given value.

### HasDirect

`func (o *ApiDependencyTreeNodeDTO) HasDirect() bool`

HasDirect returns a boolean if a field has been set.

### GetPackageUrl

`func (o *ApiDependencyTreeNodeDTO) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *ApiDependencyTreeNodeDTO) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *ApiDependencyTreeNodeDTO) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *ApiDependencyTreeNodeDTO) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


