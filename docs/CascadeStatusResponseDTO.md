# CascadeStatusResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Evaluated** | Pointer to [**[]CascadeComponentProgressDTO**](CascadeComponentProgressDTO.md) |  | [optional] 
**Failed** | Pointer to [**[]CascadeComponentProgressDTO**](CascadeComponentProgressDTO.md) |  | [optional] 
**Pending** | Pointer to [**[]CascadeComponentProgressDTO**](CascadeComponentProgressDTO.md) |  | [optional] 
**ReferenceComponentHash** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCascadeStatusResponseDTO

`func NewCascadeStatusResponseDTO() *CascadeStatusResponseDTO`

NewCascadeStatusResponseDTO instantiates a new CascadeStatusResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCascadeStatusResponseDTOWithDefaults

`func NewCascadeStatusResponseDTOWithDefaults() *CascadeStatusResponseDTO`

NewCascadeStatusResponseDTOWithDefaults instantiates a new CascadeStatusResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvaluated

`func (o *CascadeStatusResponseDTO) GetEvaluated() []CascadeComponentProgressDTO`

GetEvaluated returns the Evaluated field if non-nil, zero value otherwise.

### GetEvaluatedOk

`func (o *CascadeStatusResponseDTO) GetEvaluatedOk() (*[]CascadeComponentProgressDTO, bool)`

GetEvaluatedOk returns a tuple with the Evaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluated

`func (o *CascadeStatusResponseDTO) SetEvaluated(v []CascadeComponentProgressDTO)`

SetEvaluated sets Evaluated field to given value.

### HasEvaluated

`func (o *CascadeStatusResponseDTO) HasEvaluated() bool`

HasEvaluated returns a boolean if a field has been set.

### GetFailed

`func (o *CascadeStatusResponseDTO) GetFailed() []CascadeComponentProgressDTO`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *CascadeStatusResponseDTO) GetFailedOk() (*[]CascadeComponentProgressDTO, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *CascadeStatusResponseDTO) SetFailed(v []CascadeComponentProgressDTO)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *CascadeStatusResponseDTO) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetPending

`func (o *CascadeStatusResponseDTO) GetPending() []CascadeComponentProgressDTO`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *CascadeStatusResponseDTO) GetPendingOk() (*[]CascadeComponentProgressDTO, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *CascadeStatusResponseDTO) SetPending(v []CascadeComponentProgressDTO)`

SetPending sets Pending field to given value.

### HasPending

`func (o *CascadeStatusResponseDTO) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetReferenceComponentHash

`func (o *CascadeStatusResponseDTO) GetReferenceComponentHash() string`

GetReferenceComponentHash returns the ReferenceComponentHash field if non-nil, zero value otherwise.

### GetReferenceComponentHashOk

`func (o *CascadeStatusResponseDTO) GetReferenceComponentHashOk() (*string, bool)`

GetReferenceComponentHashOk returns a tuple with the ReferenceComponentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceComponentHash

`func (o *CascadeStatusResponseDTO) SetReferenceComponentHash(v string)`

SetReferenceComponentHash sets ReferenceComponentHash field to given value.

### HasReferenceComponentHash

`func (o *CascadeStatusResponseDTO) HasReferenceComponentHash() bool`

HasReferenceComponentHash returns a boolean if a field has been set.

### GetStatus

`func (o *CascadeStatusResponseDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CascadeStatusResponseDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CascadeStatusResponseDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CascadeStatusResponseDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


