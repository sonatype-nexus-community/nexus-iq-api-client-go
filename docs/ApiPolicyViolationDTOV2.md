# ApiPolicyViolationDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConstraintViolations** | Pointer to [**[]ApiConstraintViolationDTO**](ApiConstraintViolationDTO.md) |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiPolicyViolationDTOV2

`func NewApiPolicyViolationDTOV2() *ApiPolicyViolationDTOV2`

NewApiPolicyViolationDTOV2 instantiates a new ApiPolicyViolationDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyViolationDTOV2WithDefaults

`func NewApiPolicyViolationDTOV2WithDefaults() *ApiPolicyViolationDTOV2`

NewApiPolicyViolationDTOV2WithDefaults instantiates a new ApiPolicyViolationDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraintViolations

`func (o *ApiPolicyViolationDTOV2) GetConstraintViolations() []ApiConstraintViolationDTO`

GetConstraintViolations returns the ConstraintViolations field if non-nil, zero value otherwise.

### GetConstraintViolationsOk

`func (o *ApiPolicyViolationDTOV2) GetConstraintViolationsOk() (*[]ApiConstraintViolationDTO, bool)`

GetConstraintViolationsOk returns a tuple with the ConstraintViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintViolations

`func (o *ApiPolicyViolationDTOV2) SetConstraintViolations(v []ApiConstraintViolationDTO)`

SetConstraintViolations sets ConstraintViolations field to given value.

### HasConstraintViolations

`func (o *ApiPolicyViolationDTOV2) HasConstraintViolations() bool`

HasConstraintViolations returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApiPolicyViolationDTOV2) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApiPolicyViolationDTOV2) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApiPolicyViolationDTOV2) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApiPolicyViolationDTOV2) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiPolicyViolationDTOV2) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiPolicyViolationDTOV2) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiPolicyViolationDTOV2) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiPolicyViolationDTOV2) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *ApiPolicyViolationDTOV2) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *ApiPolicyViolationDTOV2) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *ApiPolicyViolationDTOV2) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *ApiPolicyViolationDTOV2) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiPolicyViolationDTOV2) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiPolicyViolationDTOV2) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiPolicyViolationDTOV2) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiPolicyViolationDTOV2) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


