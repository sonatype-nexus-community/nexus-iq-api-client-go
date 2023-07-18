# ApiPolicyViolationStageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentPolicyViolations** | Pointer to [**[]ApiComponentPolicyViolationDTO**](ApiComponentPolicyViolationDTO.md) |  | [optional] 
**StageId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiPolicyViolationStageDTO

`func NewApiPolicyViolationStageDTO() *ApiPolicyViolationStageDTO`

NewApiPolicyViolationStageDTO instantiates a new ApiPolicyViolationStageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyViolationStageDTOWithDefaults

`func NewApiPolicyViolationStageDTOWithDefaults() *ApiPolicyViolationStageDTO`

NewApiPolicyViolationStageDTOWithDefaults instantiates a new ApiPolicyViolationStageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentPolicyViolations

`func (o *ApiPolicyViolationStageDTO) GetComponentPolicyViolations() []ApiComponentPolicyViolationDTO`

GetComponentPolicyViolations returns the ComponentPolicyViolations field if non-nil, zero value otherwise.

### GetComponentPolicyViolationsOk

`func (o *ApiPolicyViolationStageDTO) GetComponentPolicyViolationsOk() (*[]ApiComponentPolicyViolationDTO, bool)`

GetComponentPolicyViolationsOk returns a tuple with the ComponentPolicyViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPolicyViolations

`func (o *ApiPolicyViolationStageDTO) SetComponentPolicyViolations(v []ApiComponentPolicyViolationDTO)`

SetComponentPolicyViolations sets ComponentPolicyViolations field to given value.

### HasComponentPolicyViolations

`func (o *ApiPolicyViolationStageDTO) HasComponentPolicyViolations() bool`

HasComponentPolicyViolations returns a boolean if a field has been set.

### GetStageId

`func (o *ApiPolicyViolationStageDTO) GetStageId() string`

GetStageId returns the StageId field if non-nil, zero value otherwise.

### GetStageIdOk

`func (o *ApiPolicyViolationStageDTO) GetStageIdOk() (*string, bool)`

GetStageIdOk returns a tuple with the StageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageId

`func (o *ApiPolicyViolationStageDTO) SetStageId(v string)`

SetStageId sets StageId field to given value.

### HasStageId

`func (o *ApiPolicyViolationStageDTO) HasStageId() bool`

HasStageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


