# ApiApplicationWaiverDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**ApiApplicationBaseDTO**](ApiApplicationBaseDTO.md) |  | [optional] 
**Stages** | Pointer to [**[]ApiPolicyViolationStageDTO**](ApiPolicyViolationStageDTO.md) |  | [optional] 

## Methods

### NewApiApplicationWaiverDTO

`func NewApiApplicationWaiverDTO() *ApiApplicationWaiverDTO`

NewApiApplicationWaiverDTO instantiates a new ApiApplicationWaiverDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiApplicationWaiverDTOWithDefaults

`func NewApiApplicationWaiverDTOWithDefaults() *ApiApplicationWaiverDTO`

NewApiApplicationWaiverDTOWithDefaults instantiates a new ApiApplicationWaiverDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ApiApplicationWaiverDTO) GetApplication() ApiApplicationBaseDTO`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApiApplicationWaiverDTO) GetApplicationOk() (*ApiApplicationBaseDTO, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApiApplicationWaiverDTO) SetApplication(v ApiApplicationBaseDTO)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApiApplicationWaiverDTO) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetStages

`func (o *ApiApplicationWaiverDTO) GetStages() []ApiPolicyViolationStageDTO`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *ApiApplicationWaiverDTO) GetStagesOk() (*[]ApiPolicyViolationStageDTO, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *ApiApplicationWaiverDTO) SetStages(v []ApiPolicyViolationStageDTO)`

SetStages sets Stages field to given value.

### HasStages

`func (o *ApiApplicationWaiverDTO) HasStages() bool`

HasStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


