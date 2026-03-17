# ApiPermissionCategoryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to [**[]ApiPermissionDTO**](ApiPermissionDTO.md) |  | [optional] 

## Methods

### NewApiPermissionCategoryDTO

`func NewApiPermissionCategoryDTO() *ApiPermissionCategoryDTO`

NewApiPermissionCategoryDTO instantiates a new ApiPermissionCategoryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPermissionCategoryDTOWithDefaults

`func NewApiPermissionCategoryDTOWithDefaults() *ApiPermissionCategoryDTO`

NewApiPermissionCategoryDTOWithDefaults instantiates a new ApiPermissionCategoryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ApiPermissionCategoryDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiPermissionCategoryDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiPermissionCategoryDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiPermissionCategoryDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPermissions

`func (o *ApiPermissionCategoryDTO) GetPermissions() []ApiPermissionDTO`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ApiPermissionCategoryDTO) GetPermissionsOk() (*[]ApiPermissionDTO, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ApiPermissionCategoryDTO) SetPermissions(v []ApiPermissionDTO)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ApiPermissionCategoryDTO) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


