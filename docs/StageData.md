# StageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionTypeId** | Pointer to **string** |  | [optional] 
**MostRecentEvaluationTime** | Pointer to **time.Time** |  | [optional] 
**MostRecentScanId** | Pointer to **string** |  | [optional] 

## Methods

### NewStageData

`func NewStageData() *StageData`

NewStageData instantiates a new StageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStageDataWithDefaults

`func NewStageDataWithDefaults() *StageData`

NewStageDataWithDefaults instantiates a new StageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionTypeId

`func (o *StageData) GetActionTypeId() string`

GetActionTypeId returns the ActionTypeId field if non-nil, zero value otherwise.

### GetActionTypeIdOk

`func (o *StageData) GetActionTypeIdOk() (*string, bool)`

GetActionTypeIdOk returns a tuple with the ActionTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTypeId

`func (o *StageData) SetActionTypeId(v string)`

SetActionTypeId sets ActionTypeId field to given value.

### HasActionTypeId

`func (o *StageData) HasActionTypeId() bool`

HasActionTypeId returns a boolean if a field has been set.

### GetMostRecentEvaluationTime

`func (o *StageData) GetMostRecentEvaluationTime() time.Time`

GetMostRecentEvaluationTime returns the MostRecentEvaluationTime field if non-nil, zero value otherwise.

### GetMostRecentEvaluationTimeOk

`func (o *StageData) GetMostRecentEvaluationTimeOk() (*time.Time, bool)`

GetMostRecentEvaluationTimeOk returns a tuple with the MostRecentEvaluationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentEvaluationTime

`func (o *StageData) SetMostRecentEvaluationTime(v time.Time)`

SetMostRecentEvaluationTime sets MostRecentEvaluationTime field to given value.

### HasMostRecentEvaluationTime

`func (o *StageData) HasMostRecentEvaluationTime() bool`

HasMostRecentEvaluationTime returns a boolean if a field has been set.

### GetMostRecentScanId

`func (o *StageData) GetMostRecentScanId() string`

GetMostRecentScanId returns the MostRecentScanId field if non-nil, zero value otherwise.

### GetMostRecentScanIdOk

`func (o *StageData) GetMostRecentScanIdOk() (*string, bool)`

GetMostRecentScanIdOk returns a tuple with the MostRecentScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMostRecentScanId

`func (o *StageData) SetMostRecentScanId(v string)`

SetMostRecentScanId sets MostRecentScanId field to given value.

### HasMostRecentScanId

`func (o *StageData) HasMostRecentScanId() bool`

HasMostRecentScanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


