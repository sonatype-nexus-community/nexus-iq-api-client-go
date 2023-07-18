# ApiPolicyWaiversApplicableToViolationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveWaivers** | Pointer to [**[]ApiPolicyWaiverDTO**](ApiPolicyWaiverDTO.md) |  | [optional] 
**ExpiredWaivers** | Pointer to [**[]ApiPolicyWaiverDTO**](ApiPolicyWaiverDTO.md) |  | [optional] 

## Methods

### NewApiPolicyWaiversApplicableToViolationDTO

`func NewApiPolicyWaiversApplicableToViolationDTO() *ApiPolicyWaiversApplicableToViolationDTO`

NewApiPolicyWaiversApplicableToViolationDTO instantiates a new ApiPolicyWaiversApplicableToViolationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyWaiversApplicableToViolationDTOWithDefaults

`func NewApiPolicyWaiversApplicableToViolationDTOWithDefaults() *ApiPolicyWaiversApplicableToViolationDTO`

NewApiPolicyWaiversApplicableToViolationDTOWithDefaults instantiates a new ApiPolicyWaiversApplicableToViolationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveWaivers

`func (o *ApiPolicyWaiversApplicableToViolationDTO) GetActiveWaivers() []ApiPolicyWaiverDTO`

GetActiveWaivers returns the ActiveWaivers field if non-nil, zero value otherwise.

### GetActiveWaiversOk

`func (o *ApiPolicyWaiversApplicableToViolationDTO) GetActiveWaiversOk() (*[]ApiPolicyWaiverDTO, bool)`

GetActiveWaiversOk returns a tuple with the ActiveWaivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveWaivers

`func (o *ApiPolicyWaiversApplicableToViolationDTO) SetActiveWaivers(v []ApiPolicyWaiverDTO)`

SetActiveWaivers sets ActiveWaivers field to given value.

### HasActiveWaivers

`func (o *ApiPolicyWaiversApplicableToViolationDTO) HasActiveWaivers() bool`

HasActiveWaivers returns a boolean if a field has been set.

### GetExpiredWaivers

`func (o *ApiPolicyWaiversApplicableToViolationDTO) GetExpiredWaivers() []ApiPolicyWaiverDTO`

GetExpiredWaivers returns the ExpiredWaivers field if non-nil, zero value otherwise.

### GetExpiredWaiversOk

`func (o *ApiPolicyWaiversApplicableToViolationDTO) GetExpiredWaiversOk() (*[]ApiPolicyWaiverDTO, bool)`

GetExpiredWaiversOk returns a tuple with the ExpiredWaivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredWaivers

`func (o *ApiPolicyWaiversApplicableToViolationDTO) SetExpiredWaivers(v []ApiPolicyWaiverDTO)`

SetExpiredWaivers sets ExpiredWaivers field to given value.

### HasExpiredWaivers

`func (o *ApiPolicyWaiversApplicableToViolationDTO) HasExpiredWaivers() bool`

HasExpiredWaivers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


