# ApiRoleMemberMappingDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]ApiMemberDTO**](ApiMemberDTO.md) |  | [optional] 
**RoleId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiRoleMemberMappingDTO

`func NewApiRoleMemberMappingDTO() *ApiRoleMemberMappingDTO`

NewApiRoleMemberMappingDTO instantiates a new ApiRoleMemberMappingDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRoleMemberMappingDTOWithDefaults

`func NewApiRoleMemberMappingDTOWithDefaults() *ApiRoleMemberMappingDTO`

NewApiRoleMemberMappingDTOWithDefaults instantiates a new ApiRoleMemberMappingDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ApiRoleMemberMappingDTO) GetMembers() []ApiMemberDTO`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ApiRoleMemberMappingDTO) GetMembersOk() (*[]ApiMemberDTO, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ApiRoleMemberMappingDTO) SetMembers(v []ApiMemberDTO)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ApiRoleMemberMappingDTO) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetRoleId

`func (o *ApiRoleMemberMappingDTO) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *ApiRoleMemberMappingDTO) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *ApiRoleMemberMappingDTO) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.

### HasRoleId

`func (o *ApiRoleMemberMappingDTO) HasRoleId() bool`

HasRoleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


