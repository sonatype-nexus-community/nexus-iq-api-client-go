# ApiComponentEvaluationResultDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**EvaluationDate** | Pointer to **time.Time** |  | [optional] 
**IsError** | Pointer to **bool** |  | [optional] 
**Results** | Pointer to [**[]ApiComponentDetailsDTOV2**](ApiComponentDetailsDTOV2.md) |  | [optional] 
**SubmittedDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewApiComponentEvaluationResultDTOV2

`func NewApiComponentEvaluationResultDTOV2() *ApiComponentEvaluationResultDTOV2`

NewApiComponentEvaluationResultDTOV2 instantiates a new ApiComponentEvaluationResultDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiComponentEvaluationResultDTOV2WithDefaults

`func NewApiComponentEvaluationResultDTOV2WithDefaults() *ApiComponentEvaluationResultDTOV2`

NewApiComponentEvaluationResultDTOV2WithDefaults instantiates a new ApiComponentEvaluationResultDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApiComponentEvaluationResultDTOV2) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApiComponentEvaluationResultDTOV2) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApiComponentEvaluationResultDTOV2) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApiComponentEvaluationResultDTOV2) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ApiComponentEvaluationResultDTOV2) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiComponentEvaluationResultDTOV2) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiComponentEvaluationResultDTOV2) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiComponentEvaluationResultDTOV2) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ApiComponentEvaluationResultDTOV2) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ApiComponentEvaluationResultDTOV2) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetEvaluationDate

`func (o *ApiComponentEvaluationResultDTOV2) GetEvaluationDate() time.Time`

GetEvaluationDate returns the EvaluationDate field if non-nil, zero value otherwise.

### GetEvaluationDateOk

`func (o *ApiComponentEvaluationResultDTOV2) GetEvaluationDateOk() (*time.Time, bool)`

GetEvaluationDateOk returns a tuple with the EvaluationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDate

`func (o *ApiComponentEvaluationResultDTOV2) SetEvaluationDate(v time.Time)`

SetEvaluationDate sets EvaluationDate field to given value.

### HasEvaluationDate

`func (o *ApiComponentEvaluationResultDTOV2) HasEvaluationDate() bool`

HasEvaluationDate returns a boolean if a field has been set.

### GetIsError

`func (o *ApiComponentEvaluationResultDTOV2) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *ApiComponentEvaluationResultDTOV2) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *ApiComponentEvaluationResultDTOV2) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *ApiComponentEvaluationResultDTOV2) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetResults

`func (o *ApiComponentEvaluationResultDTOV2) GetResults() []ApiComponentDetailsDTOV2`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApiComponentEvaluationResultDTOV2) GetResultsOk() (*[]ApiComponentDetailsDTOV2, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApiComponentEvaluationResultDTOV2) SetResults(v []ApiComponentDetailsDTOV2)`

SetResults sets Results field to given value.

### HasResults

`func (o *ApiComponentEvaluationResultDTOV2) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSubmittedDate

`func (o *ApiComponentEvaluationResultDTOV2) GetSubmittedDate() time.Time`

GetSubmittedDate returns the SubmittedDate field if non-nil, zero value otherwise.

### GetSubmittedDateOk

`func (o *ApiComponentEvaluationResultDTOV2) GetSubmittedDateOk() (*time.Time, bool)`

GetSubmittedDateOk returns a tuple with the SubmittedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedDate

`func (o *ApiComponentEvaluationResultDTOV2) SetSubmittedDate(v time.Time)`

SetSubmittedDate sets SubmittedDate field to given value.

### HasSubmittedDate

`func (o *ApiComponentEvaluationResultDTOV2) HasSubmittedDate() bool`

HasSubmittedDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


