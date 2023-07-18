# ValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 

## Methods

### NewValidationResult

`func NewValidationResult() *ValidationResult`

NewValidationResult instantiates a new ValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationResultWithDefaults

`func NewValidationResultWithDefaults() *ValidationResult`

NewValidationResultWithDefaults instantiates a new ValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ValidationResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidationResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidationResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidationResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetValid

`func (o *ValidationResult) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ValidationResult) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ValidationResult) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *ValidationResult) HasValid() bool`

HasValid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


