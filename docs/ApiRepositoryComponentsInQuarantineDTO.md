# ApiRepositoryComponentsInQuarantineDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]ApiRepositoryComponentPolicyViolationDTO**](ApiRepositoryComponentPolicyViolationDTO.md) |  | [optional] 
**Repository** | Pointer to [**ApiRepositoryDTO**](ApiRepositoryDTO.md) |  | [optional] 

## Methods

### NewApiRepositoryComponentsInQuarantineDTO

`func NewApiRepositoryComponentsInQuarantineDTO() *ApiRepositoryComponentsInQuarantineDTO`

NewApiRepositoryComponentsInQuarantineDTO instantiates a new ApiRepositoryComponentsInQuarantineDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRepositoryComponentsInQuarantineDTOWithDefaults

`func NewApiRepositoryComponentsInQuarantineDTOWithDefaults() *ApiRepositoryComponentsInQuarantineDTO`

NewApiRepositoryComponentsInQuarantineDTOWithDefaults instantiates a new ApiRepositoryComponentsInQuarantineDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *ApiRepositoryComponentsInQuarantineDTO) GetComponents() []ApiRepositoryComponentPolicyViolationDTO`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ApiRepositoryComponentsInQuarantineDTO) GetComponentsOk() (*[]ApiRepositoryComponentPolicyViolationDTO, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ApiRepositoryComponentsInQuarantineDTO) SetComponents(v []ApiRepositoryComponentPolicyViolationDTO)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ApiRepositoryComponentsInQuarantineDTO) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetRepository

`func (o *ApiRepositoryComponentsInQuarantineDTO) GetRepository() ApiRepositoryDTO`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ApiRepositoryComponentsInQuarantineDTO) GetRepositoryOk() (*ApiRepositoryDTO, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ApiRepositoryComponentsInQuarantineDTO) SetRepository(v ApiRepositoryDTO)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ApiRepositoryComponentsInQuarantineDTO) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


