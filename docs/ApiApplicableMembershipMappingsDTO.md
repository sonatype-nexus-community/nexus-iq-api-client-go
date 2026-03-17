# ApiApplicableMembershipMappingsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupSearchEnabled** | Pointer to **bool** |  | [optional] 
**MembersByRole** | Pointer to [**[]ApiRoleWithMembersByOwnerDTO**](ApiRoleWithMembersByOwnerDTO.md) |  | [optional] 

## Methods

### NewApiApplicableMembershipMappingsDTO

`func NewApiApplicableMembershipMappingsDTO() *ApiApplicableMembershipMappingsDTO`

NewApiApplicableMembershipMappingsDTO instantiates a new ApiApplicableMembershipMappingsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiApplicableMembershipMappingsDTOWithDefaults

`func NewApiApplicableMembershipMappingsDTOWithDefaults() *ApiApplicableMembershipMappingsDTO`

NewApiApplicableMembershipMappingsDTOWithDefaults instantiates a new ApiApplicableMembershipMappingsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupSearchEnabled

`func (o *ApiApplicableMembershipMappingsDTO) GetGroupSearchEnabled() bool`

GetGroupSearchEnabled returns the GroupSearchEnabled field if non-nil, zero value otherwise.

### GetGroupSearchEnabledOk

`func (o *ApiApplicableMembershipMappingsDTO) GetGroupSearchEnabledOk() (*bool, bool)`

GetGroupSearchEnabledOk returns a tuple with the GroupSearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSearchEnabled

`func (o *ApiApplicableMembershipMappingsDTO) SetGroupSearchEnabled(v bool)`

SetGroupSearchEnabled sets GroupSearchEnabled field to given value.

### HasGroupSearchEnabled

`func (o *ApiApplicableMembershipMappingsDTO) HasGroupSearchEnabled() bool`

HasGroupSearchEnabled returns a boolean if a field has been set.

### GetMembersByRole

`func (o *ApiApplicableMembershipMappingsDTO) GetMembersByRole() []ApiRoleWithMembersByOwnerDTO`

GetMembersByRole returns the MembersByRole field if non-nil, zero value otherwise.

### GetMembersByRoleOk

`func (o *ApiApplicableMembershipMappingsDTO) GetMembersByRoleOk() (*[]ApiRoleWithMembersByOwnerDTO, bool)`

GetMembersByRoleOk returns a tuple with the MembersByRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembersByRole

`func (o *ApiApplicableMembershipMappingsDTO) SetMembersByRole(v []ApiRoleWithMembersByOwnerDTO)`

SetMembersByRole sets MembersByRole field to given value.

### HasMembersByRole

`func (o *ApiApplicableMembershipMappingsDTO) HasMembersByRole() bool`

HasMembersByRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


