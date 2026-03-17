# ApiRoleDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltIn** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PermissionCategories** | Pointer to [**[]ApiPermissionCategoryDTO**](ApiPermissionCategoryDTO.md) |  | [optional] 

## Methods

### NewApiRoleDTO

`func NewApiRoleDTO() *ApiRoleDTO`

NewApiRoleDTO instantiates a new ApiRoleDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRoleDTOWithDefaults

`func NewApiRoleDTOWithDefaults() *ApiRoleDTO`

NewApiRoleDTOWithDefaults instantiates a new ApiRoleDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltIn

`func (o *ApiRoleDTO) GetBuiltIn() bool`

GetBuiltIn returns the BuiltIn field if non-nil, zero value otherwise.

### GetBuiltInOk

`func (o *ApiRoleDTO) GetBuiltInOk() (*bool, bool)`

GetBuiltInOk returns a tuple with the BuiltIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltIn

`func (o *ApiRoleDTO) SetBuiltIn(v bool)`

SetBuiltIn sets BuiltIn field to given value.

### HasBuiltIn

`func (o *ApiRoleDTO) HasBuiltIn() bool`

HasBuiltIn returns a boolean if a field has been set.

### GetDescription

`func (o *ApiRoleDTO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiRoleDTO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiRoleDTO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiRoleDTO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ApiRoleDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiRoleDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiRoleDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiRoleDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApiRoleDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiRoleDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiRoleDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiRoleDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissionCategories

`func (o *ApiRoleDTO) GetPermissionCategories() []ApiPermissionCategoryDTO`

GetPermissionCategories returns the PermissionCategories field if non-nil, zero value otherwise.

### GetPermissionCategoriesOk

`func (o *ApiRoleDTO) GetPermissionCategoriesOk() (*[]ApiPermissionCategoryDTO, bool)`

GetPermissionCategoriesOk returns a tuple with the PermissionCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionCategories

`func (o *ApiRoleDTO) SetPermissionCategories(v []ApiPermissionCategoryDTO)`

SetPermissionCategories sets PermissionCategories field to given value.

### HasPermissionCategories

`func (o *ApiRoleDTO) HasPermissionCategories() bool`

HasPermissionCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


