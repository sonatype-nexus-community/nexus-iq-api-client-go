# ApiStaleEvaluationStageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastEvaluationDate** | Pointer to **time.Time** |  | [optional] 
**StageId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiStaleEvaluationStageDTO

`func NewApiStaleEvaluationStageDTO() *ApiStaleEvaluationStageDTO`

NewApiStaleEvaluationStageDTO instantiates a new ApiStaleEvaluationStageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStaleEvaluationStageDTOWithDefaults

`func NewApiStaleEvaluationStageDTOWithDefaults() *ApiStaleEvaluationStageDTO`

NewApiStaleEvaluationStageDTOWithDefaults instantiates a new ApiStaleEvaluationStageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastEvaluationDate

`func (o *ApiStaleEvaluationStageDTO) GetLastEvaluationDate() time.Time`

GetLastEvaluationDate returns the LastEvaluationDate field if non-nil, zero value otherwise.

### GetLastEvaluationDateOk

`func (o *ApiStaleEvaluationStageDTO) GetLastEvaluationDateOk() (*time.Time, bool)`

GetLastEvaluationDateOk returns a tuple with the LastEvaluationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEvaluationDate

`func (o *ApiStaleEvaluationStageDTO) SetLastEvaluationDate(v time.Time)`

SetLastEvaluationDate sets LastEvaluationDate field to given value.

### HasLastEvaluationDate

`func (o *ApiStaleEvaluationStageDTO) HasLastEvaluationDate() bool`

HasLastEvaluationDate returns a boolean if a field has been set.

### GetStageId

`func (o *ApiStaleEvaluationStageDTO) GetStageId() string`

GetStageId returns the StageId field if non-nil, zero value otherwise.

### GetStageIdOk

`func (o *ApiStaleEvaluationStageDTO) GetStageIdOk() (*string, bool)`

GetStageIdOk returns a tuple with the StageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageId

`func (o *ApiStaleEvaluationStageDTO) SetStageId(v string)`

SetStageId sets StageId field to given value.

### HasStageId

`func (o *ApiStaleEvaluationStageDTO) HasStageId() bool`

HasStageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


