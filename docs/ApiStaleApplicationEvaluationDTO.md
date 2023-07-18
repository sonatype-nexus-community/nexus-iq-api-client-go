# ApiStaleApplicationEvaluationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**ApiApplicationBaseDTO**](ApiApplicationBaseDTO.md) |  | [optional] 
**Stages** | Pointer to [**[]ApiStaleEvaluationStageDTO**](ApiStaleEvaluationStageDTO.md) |  | [optional] 

## Methods

### NewApiStaleApplicationEvaluationDTO

`func NewApiStaleApplicationEvaluationDTO() *ApiStaleApplicationEvaluationDTO`

NewApiStaleApplicationEvaluationDTO instantiates a new ApiStaleApplicationEvaluationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStaleApplicationEvaluationDTOWithDefaults

`func NewApiStaleApplicationEvaluationDTOWithDefaults() *ApiStaleApplicationEvaluationDTO`

NewApiStaleApplicationEvaluationDTOWithDefaults instantiates a new ApiStaleApplicationEvaluationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ApiStaleApplicationEvaluationDTO) GetApplication() ApiApplicationBaseDTO`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApiStaleApplicationEvaluationDTO) GetApplicationOk() (*ApiApplicationBaseDTO, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApiStaleApplicationEvaluationDTO) SetApplication(v ApiApplicationBaseDTO)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApiStaleApplicationEvaluationDTO) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetStages

`func (o *ApiStaleApplicationEvaluationDTO) GetStages() []ApiStaleEvaluationStageDTO`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *ApiStaleApplicationEvaluationDTO) GetStagesOk() (*[]ApiStaleEvaluationStageDTO, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *ApiStaleApplicationEvaluationDTO) SetStages(v []ApiStaleEvaluationStageDTO)`

SetStages sets Stages field to given value.

### HasStages

`func (o *ApiStaleApplicationEvaluationDTO) HasStages() bool`

HasStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


