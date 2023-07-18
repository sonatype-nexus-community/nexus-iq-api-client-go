# TagsByOwnerDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationCategories** | Pointer to [**[]ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md) |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerName** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **string** |  | [optional] 

## Methods

### NewTagsByOwnerDTO

`func NewTagsByOwnerDTO() *TagsByOwnerDTO`

NewTagsByOwnerDTO instantiates a new TagsByOwnerDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagsByOwnerDTOWithDefaults

`func NewTagsByOwnerDTOWithDefaults() *TagsByOwnerDTO`

NewTagsByOwnerDTOWithDefaults instantiates a new TagsByOwnerDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationCategories

`func (o *TagsByOwnerDTO) GetApplicationCategories() []ApiApplicationCategoryDTO`

GetApplicationCategories returns the ApplicationCategories field if non-nil, zero value otherwise.

### GetApplicationCategoriesOk

`func (o *TagsByOwnerDTO) GetApplicationCategoriesOk() (*[]ApiApplicationCategoryDTO, bool)`

GetApplicationCategoriesOk returns a tuple with the ApplicationCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCategories

`func (o *TagsByOwnerDTO) SetApplicationCategories(v []ApiApplicationCategoryDTO)`

SetApplicationCategories sets ApplicationCategories field to given value.

### HasApplicationCategories

`func (o *TagsByOwnerDTO) HasApplicationCategories() bool`

HasApplicationCategories returns a boolean if a field has been set.

### GetOwnerId

`func (o *TagsByOwnerDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *TagsByOwnerDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *TagsByOwnerDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *TagsByOwnerDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerName

`func (o *TagsByOwnerDTO) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *TagsByOwnerDTO) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *TagsByOwnerDTO) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *TagsByOwnerDTO) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetOwnerType

`func (o *TagsByOwnerDTO) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *TagsByOwnerDTO) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *TagsByOwnerDTO) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *TagsByOwnerDTO) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


