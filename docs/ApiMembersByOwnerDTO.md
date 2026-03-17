# ApiMembersByOwnerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]ApiMemberWithDetailsDTO**](ApiMemberWithDetailsDTO.md) |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerName** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **string** |  | [optional] 

## Methods

### NewApiMembersByOwnerDTO

`func NewApiMembersByOwnerDTO() *ApiMembersByOwnerDTO`

NewApiMembersByOwnerDTO instantiates a new ApiMembersByOwnerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMembersByOwnerDTOWithDefaults

`func NewApiMembersByOwnerDTOWithDefaults() *ApiMembersByOwnerDTO`

NewApiMembersByOwnerDTOWithDefaults instantiates a new ApiMembersByOwnerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ApiMembersByOwnerDTO) GetMembers() []ApiMemberWithDetailsDTO`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ApiMembersByOwnerDTO) GetMembersOk() (*[]ApiMemberWithDetailsDTO, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ApiMembersByOwnerDTO) SetMembers(v []ApiMemberWithDetailsDTO)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *ApiMembersByOwnerDTO) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetOwnerId

`func (o *ApiMembersByOwnerDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ApiMembersByOwnerDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ApiMembersByOwnerDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ApiMembersByOwnerDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerName

`func (o *ApiMembersByOwnerDTO) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *ApiMembersByOwnerDTO) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *ApiMembersByOwnerDTO) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *ApiMembersByOwnerDTO) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetOwnerType

`func (o *ApiMembersByOwnerDTO) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *ApiMembersByOwnerDTO) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *ApiMembersByOwnerDTO) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *ApiMembersByOwnerDTO) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


