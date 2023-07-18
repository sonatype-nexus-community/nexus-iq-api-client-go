# ApiComponentPolicyViolationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to [**ApiComponentDTOV2**](ApiComponentDTOV2.md) |  | [optional] 
**WaivedPolicyViolations** | Pointer to [**[]ApiWaivedPolicyViolationDTO**](ApiWaivedPolicyViolationDTO.md) |  | [optional] 

## Methods

### NewApiComponentPolicyViolationDTO

`func NewApiComponentPolicyViolationDTO() *ApiComponentPolicyViolationDTO`

NewApiComponentPolicyViolationDTO instantiates a new ApiComponentPolicyViolationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiComponentPolicyViolationDTOWithDefaults

`func NewApiComponentPolicyViolationDTOWithDefaults() *ApiComponentPolicyViolationDTO`

NewApiComponentPolicyViolationDTOWithDefaults instantiates a new ApiComponentPolicyViolationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ApiComponentPolicyViolationDTO) GetComponent() ApiComponentDTOV2`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ApiComponentPolicyViolationDTO) GetComponentOk() (*ApiComponentDTOV2, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ApiComponentPolicyViolationDTO) SetComponent(v ApiComponentDTOV2)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ApiComponentPolicyViolationDTO) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetWaivedPolicyViolations

`func (o *ApiComponentPolicyViolationDTO) GetWaivedPolicyViolations() []ApiWaivedPolicyViolationDTO`

GetWaivedPolicyViolations returns the WaivedPolicyViolations field if non-nil, zero value otherwise.

### GetWaivedPolicyViolationsOk

`func (o *ApiComponentPolicyViolationDTO) GetWaivedPolicyViolationsOk() (*[]ApiWaivedPolicyViolationDTO, bool)`

GetWaivedPolicyViolationsOk returns a tuple with the WaivedPolicyViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaivedPolicyViolations

`func (o *ApiComponentPolicyViolationDTO) SetWaivedPolicyViolations(v []ApiWaivedPolicyViolationDTO)`

SetWaivedPolicyViolations sets WaivedPolicyViolations field to given value.

### HasWaivedPolicyViolations

`func (o *ApiComponentPolicyViolationDTO) HasWaivedPolicyViolations() bool`

HasWaivedPolicyViolations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


