# ApiStaleEvaluationsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]ApiStaleApplicationEvaluationDTO**](ApiStaleApplicationEvaluationDTO.md) |  | [optional] 
**Repositories** | Pointer to [**[]ApiStaleRepositoryEvaluationDTO**](ApiStaleRepositoryEvaluationDTO.md) |  | [optional] 

## Methods

### NewApiStaleEvaluationsDTO

`func NewApiStaleEvaluationsDTO() *ApiStaleEvaluationsDTO`

NewApiStaleEvaluationsDTO instantiates a new ApiStaleEvaluationsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStaleEvaluationsDTOWithDefaults

`func NewApiStaleEvaluationsDTOWithDefaults() *ApiStaleEvaluationsDTO`

NewApiStaleEvaluationsDTOWithDefaults instantiates a new ApiStaleEvaluationsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *ApiStaleEvaluationsDTO) GetApplications() []ApiStaleApplicationEvaluationDTO`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ApiStaleEvaluationsDTO) GetApplicationsOk() (*[]ApiStaleApplicationEvaluationDTO, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ApiStaleEvaluationsDTO) SetApplications(v []ApiStaleApplicationEvaluationDTO)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ApiStaleEvaluationsDTO) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetRepositories

`func (o *ApiStaleEvaluationsDTO) GetRepositories() []ApiStaleRepositoryEvaluationDTO`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *ApiStaleEvaluationsDTO) GetRepositoriesOk() (*[]ApiStaleRepositoryEvaluationDTO, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *ApiStaleEvaluationsDTO) SetRepositories(v []ApiStaleRepositoryEvaluationDTO)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *ApiStaleEvaluationsDTO) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


