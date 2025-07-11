# SCMUserMappingsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mappings** | Pointer to [**[]UserMapping**](UserMapping.md) |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewSCMUserMappingsDTO

`func NewSCMUserMappingsDTO() *SCMUserMappingsDTO`

NewSCMUserMappingsDTO instantiates a new SCMUserMappingsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCMUserMappingsDTOWithDefaults

`func NewSCMUserMappingsDTOWithDefaults() *SCMUserMappingsDTO`

NewSCMUserMappingsDTOWithDefaults instantiates a new SCMUserMappingsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappings

`func (o *SCMUserMappingsDTO) GetMappings() []UserMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *SCMUserMappingsDTO) GetMappingsOk() (*[]UserMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *SCMUserMappingsDTO) SetMappings(v []UserMapping)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *SCMUserMappingsDTO) HasMappings() bool`

HasMappings returns a boolean if a field has been set.

### GetRole

`func (o *SCMUserMappingsDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SCMUserMappingsDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SCMUserMappingsDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SCMUserMappingsDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


