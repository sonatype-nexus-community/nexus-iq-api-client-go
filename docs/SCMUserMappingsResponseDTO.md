# SCMUserMappingsResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inherited** | Pointer to **bool** |  | [optional] 
**OwnerInternalId** | Pointer to **string** |  | [optional] 
**UserMapping** | Pointer to [**SCMUserMappingsDTO**](SCMUserMappingsDTO.md) |  | [optional] 

## Methods

### NewSCMUserMappingsResponseDTO

`func NewSCMUserMappingsResponseDTO() *SCMUserMappingsResponseDTO`

NewSCMUserMappingsResponseDTO instantiates a new SCMUserMappingsResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCMUserMappingsResponseDTOWithDefaults

`func NewSCMUserMappingsResponseDTOWithDefaults() *SCMUserMappingsResponseDTO`

NewSCMUserMappingsResponseDTOWithDefaults instantiates a new SCMUserMappingsResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInherited

`func (o *SCMUserMappingsResponseDTO) GetInherited() bool`

GetInherited returns the Inherited field if non-nil, zero value otherwise.

### GetInheritedOk

`func (o *SCMUserMappingsResponseDTO) GetInheritedOk() (*bool, bool)`

GetInheritedOk returns a tuple with the Inherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInherited

`func (o *SCMUserMappingsResponseDTO) SetInherited(v bool)`

SetInherited sets Inherited field to given value.

### HasInherited

`func (o *SCMUserMappingsResponseDTO) HasInherited() bool`

HasInherited returns a boolean if a field has been set.

### GetOwnerInternalId

`func (o *SCMUserMappingsResponseDTO) GetOwnerInternalId() string`

GetOwnerInternalId returns the OwnerInternalId field if non-nil, zero value otherwise.

### GetOwnerInternalIdOk

`func (o *SCMUserMappingsResponseDTO) GetOwnerInternalIdOk() (*string, bool)`

GetOwnerInternalIdOk returns a tuple with the OwnerInternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerInternalId

`func (o *SCMUserMappingsResponseDTO) SetOwnerInternalId(v string)`

SetOwnerInternalId sets OwnerInternalId field to given value.

### HasOwnerInternalId

`func (o *SCMUserMappingsResponseDTO) HasOwnerInternalId() bool`

HasOwnerInternalId returns a boolean if a field has been set.

### GetUserMapping

`func (o *SCMUserMappingsResponseDTO) GetUserMapping() SCMUserMappingsDTO`

GetUserMapping returns the UserMapping field if non-nil, zero value otherwise.

### GetUserMappingOk

`func (o *SCMUserMappingsResponseDTO) GetUserMappingOk() (*SCMUserMappingsDTO, bool)`

GetUserMappingOk returns a tuple with the UserMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMapping

`func (o *SCMUserMappingsResponseDTO) SetUserMapping(v SCMUserMappingsDTO)`

SetUserMapping sets UserMapping field to given value.

### HasUserMapping

`func (o *SCMUserMappingsResponseDTO) HasUserMapping() bool`

HasUserMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


