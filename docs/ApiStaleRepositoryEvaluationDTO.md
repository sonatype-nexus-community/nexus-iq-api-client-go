# ApiStaleRepositoryEvaluationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repository** | Pointer to [**ApiRepositoryDTO**](ApiRepositoryDTO.md) |  | [optional] 
**Stages** | Pointer to [**[]ApiStaleEvaluationStageDTO**](ApiStaleEvaluationStageDTO.md) |  | [optional] 

## Methods

### NewApiStaleRepositoryEvaluationDTO

`func NewApiStaleRepositoryEvaluationDTO() *ApiStaleRepositoryEvaluationDTO`

NewApiStaleRepositoryEvaluationDTO instantiates a new ApiStaleRepositoryEvaluationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStaleRepositoryEvaluationDTOWithDefaults

`func NewApiStaleRepositoryEvaluationDTOWithDefaults() *ApiStaleRepositoryEvaluationDTO`

NewApiStaleRepositoryEvaluationDTOWithDefaults instantiates a new ApiStaleRepositoryEvaluationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepository

`func (o *ApiStaleRepositoryEvaluationDTO) GetRepository() ApiRepositoryDTO`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ApiStaleRepositoryEvaluationDTO) GetRepositoryOk() (*ApiRepositoryDTO, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ApiStaleRepositoryEvaluationDTO) SetRepository(v ApiRepositoryDTO)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ApiStaleRepositoryEvaluationDTO) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### GetStages

`func (o *ApiStaleRepositoryEvaluationDTO) GetStages() []ApiStaleEvaluationStageDTO`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *ApiStaleRepositoryEvaluationDTO) GetStagesOk() (*[]ApiStaleEvaluationStageDTO, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *ApiStaleRepositoryEvaluationDTO) SetStages(v []ApiStaleEvaluationStageDTO)`

SetStages sets Stages field to given value.

### HasStages

`func (o *ApiStaleRepositoryEvaluationDTO) HasStages() bool`

HasStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


