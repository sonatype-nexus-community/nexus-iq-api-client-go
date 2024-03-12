# ApiRepositoryListDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repositories** | Pointer to [**[]ApiRepositoryDTO**](ApiRepositoryDTO.md) |  | [optional] 

## Methods

### NewApiRepositoryListDTO

`func NewApiRepositoryListDTO() *ApiRepositoryListDTO`

NewApiRepositoryListDTO instantiates a new ApiRepositoryListDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRepositoryListDTOWithDefaults

`func NewApiRepositoryListDTOWithDefaults() *ApiRepositoryListDTO`

NewApiRepositoryListDTOWithDefaults instantiates a new ApiRepositoryListDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositories

`func (o *ApiRepositoryListDTO) GetRepositories() []ApiRepositoryDTO`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *ApiRepositoryListDTO) GetRepositoriesOk() (*[]ApiRepositoryDTO, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *ApiRepositoryListDTO) SetRepositories(v []ApiRepositoryDTO)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *ApiRepositoryListDTO) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


