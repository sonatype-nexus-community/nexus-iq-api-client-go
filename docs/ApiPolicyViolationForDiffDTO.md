# ApiPolicyViolationForDiffDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to [**ApiComponentDTOV2**](ApiComponentDTOV2.md) |  | [optional] 
**ConstraintViolations** | Pointer to [**[]ApiConstraintViolationDTO**](ApiConstraintViolationDTO.md) |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiPolicyViolationForDiffDTO

`func NewApiPolicyViolationForDiffDTO() *ApiPolicyViolationForDiffDTO`

NewApiPolicyViolationForDiffDTO instantiates a new ApiPolicyViolationForDiffDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyViolationForDiffDTOWithDefaults

`func NewApiPolicyViolationForDiffDTOWithDefaults() *ApiPolicyViolationForDiffDTO`

NewApiPolicyViolationForDiffDTOWithDefaults instantiates a new ApiPolicyViolationForDiffDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ApiPolicyViolationForDiffDTO) GetComponent() ApiComponentDTOV2`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ApiPolicyViolationForDiffDTO) GetComponentOk() (*ApiComponentDTOV2, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ApiPolicyViolationForDiffDTO) SetComponent(v ApiComponentDTOV2)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ApiPolicyViolationForDiffDTO) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetConstraintViolations

`func (o *ApiPolicyViolationForDiffDTO) GetConstraintViolations() []ApiConstraintViolationDTO`

GetConstraintViolations returns the ConstraintViolations field if non-nil, zero value otherwise.

### GetConstraintViolationsOk

`func (o *ApiPolicyViolationForDiffDTO) GetConstraintViolationsOk() (*[]ApiConstraintViolationDTO, bool)`

GetConstraintViolationsOk returns a tuple with the ConstraintViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintViolations

`func (o *ApiPolicyViolationForDiffDTO) SetConstraintViolations(v []ApiConstraintViolationDTO)`

SetConstraintViolations sets ConstraintViolations field to given value.

### HasConstraintViolations

`func (o *ApiPolicyViolationForDiffDTO) HasConstraintViolations() bool`

HasConstraintViolations returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApiPolicyViolationForDiffDTO) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApiPolicyViolationForDiffDTO) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApiPolicyViolationForDiffDTO) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApiPolicyViolationForDiffDTO) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiPolicyViolationForDiffDTO) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiPolicyViolationForDiffDTO) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiPolicyViolationForDiffDTO) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiPolicyViolationForDiffDTO) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *ApiPolicyViolationForDiffDTO) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *ApiPolicyViolationForDiffDTO) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *ApiPolicyViolationForDiffDTO) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *ApiPolicyViolationForDiffDTO) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiPolicyViolationForDiffDTO) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiPolicyViolationForDiffDTO) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiPolicyViolationForDiffDTO) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiPolicyViolationForDiffDTO) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


