# ApiPolicyWaiverRequestsApplicableToViolationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveWaiverRequests** | Pointer to [**[]ApiPolicyWaiverRequestDTO**](ApiPolicyWaiverRequestDTO.md) |  | [optional] 
**ExpiredWaiverRequests** | Pointer to [**[]ApiPolicyWaiverRequestDTO**](ApiPolicyWaiverRequestDTO.md) |  | [optional] 

## Methods

### NewApiPolicyWaiverRequestsApplicableToViolationDTO

`func NewApiPolicyWaiverRequestsApplicableToViolationDTO() *ApiPolicyWaiverRequestsApplicableToViolationDTO`

NewApiPolicyWaiverRequestsApplicableToViolationDTO instantiates a new ApiPolicyWaiverRequestsApplicableToViolationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyWaiverRequestsApplicableToViolationDTOWithDefaults

`func NewApiPolicyWaiverRequestsApplicableToViolationDTOWithDefaults() *ApiPolicyWaiverRequestsApplicableToViolationDTO`

NewApiPolicyWaiverRequestsApplicableToViolationDTOWithDefaults instantiates a new ApiPolicyWaiverRequestsApplicableToViolationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveWaiverRequests

`func (o *ApiPolicyWaiverRequestsApplicableToViolationDTO) GetActiveWaiverRequests() []ApiPolicyWaiverRequestDTO`

GetActiveWaiverRequests returns the ActiveWaiverRequests field if non-nil, zero value otherwise.

### GetActiveWaiverRequestsOk

`func (o *ApiPolicyWaiverRequestsApplicableToViolationDTO) GetActiveWaiverRequestsOk() (*[]ApiPolicyWaiverRequestDTO, bool)`

GetActiveWaiverRequestsOk returns a tuple with the ActiveWaiverRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveWaiverRequests

`func (o *ApiPolicyWaiverRequestsApplicableToViolationDTO) SetActiveWaiverRequests(v []ApiPolicyWaiverRequestDTO)`

SetActiveWaiverRequests sets ActiveWaiverRequests field to given value.

### HasActiveWaiverRequests

`func (o *ApiPolicyWaiverRequestsApplicableToViolationDTO) HasActiveWaiverRequests() bool`

HasActiveWaiverRequests returns a boolean if a field has been set.

### GetExpiredWaiverRequests

`func (o *ApiPolicyWaiverRequestsApplicableToViolationDTO) GetExpiredWaiverRequests() []ApiPolicyWaiverRequestDTO`

GetExpiredWaiverRequests returns the ExpiredWaiverRequests field if non-nil, zero value otherwise.

### GetExpiredWaiverRequestsOk

`func (o *ApiPolicyWaiverRequestsApplicableToViolationDTO) GetExpiredWaiverRequestsOk() (*[]ApiPolicyWaiverRequestDTO, bool)`

GetExpiredWaiverRequestsOk returns a tuple with the ExpiredWaiverRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredWaiverRequests

`func (o *ApiPolicyWaiverRequestsApplicableToViolationDTO) SetExpiredWaiverRequests(v []ApiPolicyWaiverRequestDTO)`

SetExpiredWaiverRequests sets ExpiredWaiverRequests field to given value.

### HasExpiredWaiverRequests

`func (o *ApiPolicyWaiverRequestsApplicableToViolationDTO) HasExpiredWaiverRequests() bool`

HasExpiredWaiverRequests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


