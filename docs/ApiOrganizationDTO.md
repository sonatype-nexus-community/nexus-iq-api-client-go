# ApiOrganizationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentOrganizationId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]ApiTagDTO**](ApiTagDTO.md) |  | [optional] 

## Methods

### NewApiOrganizationDTO

`func NewApiOrganizationDTO() *ApiOrganizationDTO`

NewApiOrganizationDTO instantiates a new ApiOrganizationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOrganizationDTOWithDefaults

`func NewApiOrganizationDTOWithDefaults() *ApiOrganizationDTO`

NewApiOrganizationDTOWithDefaults instantiates a new ApiOrganizationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiOrganizationDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiOrganizationDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiOrganizationDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiOrganizationDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApiOrganizationDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiOrganizationDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiOrganizationDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiOrganizationDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentOrganizationId

`func (o *ApiOrganizationDTO) GetParentOrganizationId() string`

GetParentOrganizationId returns the ParentOrganizationId field if non-nil, zero value otherwise.

### GetParentOrganizationIdOk

`func (o *ApiOrganizationDTO) GetParentOrganizationIdOk() (*string, bool)`

GetParentOrganizationIdOk returns a tuple with the ParentOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOrganizationId

`func (o *ApiOrganizationDTO) SetParentOrganizationId(v string)`

SetParentOrganizationId sets ParentOrganizationId field to given value.

### HasParentOrganizationId

`func (o *ApiOrganizationDTO) HasParentOrganizationId() bool`

HasParentOrganizationId returns a boolean if a field has been set.

### GetTags

`func (o *ApiOrganizationDTO) GetTags() []ApiTagDTO`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ApiOrganizationDTO) GetTagsOk() (*[]ApiTagDTO, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ApiOrganizationDTO) SetTags(v []ApiTagDTO)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ApiOrganizationDTO) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


