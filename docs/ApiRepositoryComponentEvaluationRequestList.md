# ApiRepositoryComponentEvaluationRequestList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]ApiRepositoryComponentEvaluationRequest**](ApiRepositoryComponentEvaluationRequest.md) |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 

## Methods

### NewApiRepositoryComponentEvaluationRequestList

`func NewApiRepositoryComponentEvaluationRequestList() *ApiRepositoryComponentEvaluationRequestList`

NewApiRepositoryComponentEvaluationRequestList instantiates a new ApiRepositoryComponentEvaluationRequestList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRepositoryComponentEvaluationRequestListWithDefaults

`func NewApiRepositoryComponentEvaluationRequestListWithDefaults() *ApiRepositoryComponentEvaluationRequestList`

NewApiRepositoryComponentEvaluationRequestListWithDefaults instantiates a new ApiRepositoryComponentEvaluationRequestList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *ApiRepositoryComponentEvaluationRequestList) GetComponents() []ApiRepositoryComponentEvaluationRequest`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ApiRepositoryComponentEvaluationRequestList) GetComponentsOk() (*[]ApiRepositoryComponentEvaluationRequest, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ApiRepositoryComponentEvaluationRequestList) SetComponents(v []ApiRepositoryComponentEvaluationRequest)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ApiRepositoryComponentEvaluationRequestList) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetFormat

`func (o *ApiRepositoryComponentEvaluationRequestList) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ApiRepositoryComponentEvaluationRequestList) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ApiRepositoryComponentEvaluationRequestList) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ApiRepositoryComponentEvaluationRequestList) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


