# ApiConstraintViolationReasonDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to [**TriggerReference**](TriggerReference.md) |  | [optional] 

## Methods

### NewApiConstraintViolationReasonDTO

`func NewApiConstraintViolationReasonDTO() *ApiConstraintViolationReasonDTO`

NewApiConstraintViolationReasonDTO instantiates a new ApiConstraintViolationReasonDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiConstraintViolationReasonDTOWithDefaults

`func NewApiConstraintViolationReasonDTOWithDefaults() *ApiConstraintViolationReasonDTO`

NewApiConstraintViolationReasonDTOWithDefaults instantiates a new ApiConstraintViolationReasonDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ApiConstraintViolationReasonDTO) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiConstraintViolationReasonDTO) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiConstraintViolationReasonDTO) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ApiConstraintViolationReasonDTO) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReference

`func (o *ApiConstraintViolationReasonDTO) GetReference() TriggerReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ApiConstraintViolationReasonDTO) GetReferenceOk() (*TriggerReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ApiConstraintViolationReasonDTO) SetReference(v TriggerReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ApiConstraintViolationReasonDTO) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


