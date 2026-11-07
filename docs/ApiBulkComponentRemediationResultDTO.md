# ApiBulkComponentRemediationResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to [**ApiComponentDTOV2**](ApiComponentDTOV2.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Remediation** | Pointer to [**ApiComponentRemediationValueDTO**](ApiComponentRemediationValueDTO.md) |  | [optional] 

## Methods

### NewApiBulkComponentRemediationResultDTO

`func NewApiBulkComponentRemediationResultDTO() *ApiBulkComponentRemediationResultDTO`

NewApiBulkComponentRemediationResultDTO instantiates a new ApiBulkComponentRemediationResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBulkComponentRemediationResultDTOWithDefaults

`func NewApiBulkComponentRemediationResultDTOWithDefaults() *ApiBulkComponentRemediationResultDTO`

NewApiBulkComponentRemediationResultDTOWithDefaults instantiates a new ApiBulkComponentRemediationResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ApiBulkComponentRemediationResultDTO) GetComponent() ApiComponentDTOV2`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ApiBulkComponentRemediationResultDTO) GetComponentOk() (*ApiComponentDTOV2, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ApiBulkComponentRemediationResultDTO) SetComponent(v ApiComponentDTOV2)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ApiBulkComponentRemediationResultDTO) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetError

`func (o *ApiBulkComponentRemediationResultDTO) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiBulkComponentRemediationResultDTO) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiBulkComponentRemediationResultDTO) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ApiBulkComponentRemediationResultDTO) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRemediation

`func (o *ApiBulkComponentRemediationResultDTO) GetRemediation() ApiComponentRemediationValueDTO`

GetRemediation returns the Remediation field if non-nil, zero value otherwise.

### GetRemediationOk

`func (o *ApiBulkComponentRemediationResultDTO) GetRemediationOk() (*ApiComponentRemediationValueDTO, bool)`

GetRemediationOk returns a tuple with the Remediation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediation

`func (o *ApiBulkComponentRemediationResultDTO) SetRemediation(v ApiComponentRemediationValueDTO)`

SetRemediation sets Remediation field to given value.

### HasRemediation

`func (o *ApiBulkComponentRemediationResultDTO) HasRemediation() bool`

HasRemediation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


