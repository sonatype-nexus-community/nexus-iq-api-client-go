# ApiRepositoryComponentPolicyViolationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to [**ApiRepositoryComponentDTO**](ApiRepositoryComponentDTO.md) |  | [optional] 
**PolicyViolations** | Pointer to [**[]ApiPolicyViolationDTOV2**](ApiPolicyViolationDTOV2.md) |  | [optional] 
**WaivedPolicyViolations** | Pointer to [**[]ApiWaivedPolicyViolationDTO**](ApiWaivedPolicyViolationDTO.md) |  | [optional] 

## Methods

### NewApiRepositoryComponentPolicyViolationDTO

`func NewApiRepositoryComponentPolicyViolationDTO() *ApiRepositoryComponentPolicyViolationDTO`

NewApiRepositoryComponentPolicyViolationDTO instantiates a new ApiRepositoryComponentPolicyViolationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRepositoryComponentPolicyViolationDTOWithDefaults

`func NewApiRepositoryComponentPolicyViolationDTOWithDefaults() *ApiRepositoryComponentPolicyViolationDTO`

NewApiRepositoryComponentPolicyViolationDTOWithDefaults instantiates a new ApiRepositoryComponentPolicyViolationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ApiRepositoryComponentPolicyViolationDTO) GetComponent() ApiRepositoryComponentDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ApiRepositoryComponentPolicyViolationDTO) GetComponentOk() (*ApiRepositoryComponentDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ApiRepositoryComponentPolicyViolationDTO) SetComponent(v ApiRepositoryComponentDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ApiRepositoryComponentPolicyViolationDTO) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetPolicyViolations

`func (o *ApiRepositoryComponentPolicyViolationDTO) GetPolicyViolations() []ApiPolicyViolationDTOV2`

GetPolicyViolations returns the PolicyViolations field if non-nil, zero value otherwise.

### GetPolicyViolationsOk

`func (o *ApiRepositoryComponentPolicyViolationDTO) GetPolicyViolationsOk() (*[]ApiPolicyViolationDTOV2, bool)`

GetPolicyViolationsOk returns a tuple with the PolicyViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolations

`func (o *ApiRepositoryComponentPolicyViolationDTO) SetPolicyViolations(v []ApiPolicyViolationDTOV2)`

SetPolicyViolations sets PolicyViolations field to given value.

### HasPolicyViolations

`func (o *ApiRepositoryComponentPolicyViolationDTO) HasPolicyViolations() bool`

HasPolicyViolations returns a boolean if a field has been set.

### GetWaivedPolicyViolations

`func (o *ApiRepositoryComponentPolicyViolationDTO) GetWaivedPolicyViolations() []ApiWaivedPolicyViolationDTO`

GetWaivedPolicyViolations returns the WaivedPolicyViolations field if non-nil, zero value otherwise.

### GetWaivedPolicyViolationsOk

`func (o *ApiRepositoryComponentPolicyViolationDTO) GetWaivedPolicyViolationsOk() (*[]ApiWaivedPolicyViolationDTO, bool)`

GetWaivedPolicyViolationsOk returns a tuple with the WaivedPolicyViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaivedPolicyViolations

`func (o *ApiRepositoryComponentPolicyViolationDTO) SetWaivedPolicyViolations(v []ApiWaivedPolicyViolationDTO)`

SetWaivedPolicyViolations sets WaivedPolicyViolations field to given value.

### HasWaivedPolicyViolations

`func (o *ApiRepositoryComponentPolicyViolationDTO) HasWaivedPolicyViolations() bool`

HasWaivedPolicyViolations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


