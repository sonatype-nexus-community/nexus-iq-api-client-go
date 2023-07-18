# ApiPullRequestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExceptionThrown** | Pointer to **bool** |  | [optional] 
**Reasoning** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **time.Time** |  | [optional] 
**Successful** | Pointer to **bool** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**TotalTime** | Pointer to **int64** |  | [optional] 

## Methods

### NewApiPullRequestResult

`func NewApiPullRequestResult() *ApiPullRequestResult`

NewApiPullRequestResult instantiates a new ApiPullRequestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPullRequestResultWithDefaults

`func NewApiPullRequestResultWithDefaults() *ApiPullRequestResult`

NewApiPullRequestResultWithDefaults instantiates a new ApiPullRequestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExceptionThrown

`func (o *ApiPullRequestResult) GetExceptionThrown() bool`

GetExceptionThrown returns the ExceptionThrown field if non-nil, zero value otherwise.

### GetExceptionThrownOk

`func (o *ApiPullRequestResult) GetExceptionThrownOk() (*bool, bool)`

GetExceptionThrownOk returns a tuple with the ExceptionThrown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionThrown

`func (o *ApiPullRequestResult) SetExceptionThrown(v bool)`

SetExceptionThrown sets ExceptionThrown field to given value.

### HasExceptionThrown

`func (o *ApiPullRequestResult) HasExceptionThrown() bool`

HasExceptionThrown returns a boolean if a field has been set.

### GetReasoning

`func (o *ApiPullRequestResult) GetReasoning() string`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *ApiPullRequestResult) GetReasoningOk() (*string, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *ApiPullRequestResult) SetReasoning(v string)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *ApiPullRequestResult) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### GetStartTime

`func (o *ApiPullRequestResult) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ApiPullRequestResult) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ApiPullRequestResult) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ApiPullRequestResult) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetSuccessful

`func (o *ApiPullRequestResult) GetSuccessful() bool`

GetSuccessful returns the Successful field if non-nil, zero value otherwise.

### GetSuccessfulOk

`func (o *ApiPullRequestResult) GetSuccessfulOk() (*bool, bool)`

GetSuccessfulOk returns a tuple with the Successful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessful

`func (o *ApiPullRequestResult) SetSuccessful(v bool)`

SetSuccessful sets Successful field to given value.

### HasSuccessful

`func (o *ApiPullRequestResult) HasSuccessful() bool`

HasSuccessful returns a boolean if a field has been set.

### GetTitle

`func (o *ApiPullRequestResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiPullRequestResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiPullRequestResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiPullRequestResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTotalTime

`func (o *ApiPullRequestResult) GetTotalTime() int64`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *ApiPullRequestResult) GetTotalTimeOk() (*int64, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *ApiPullRequestResult) SetTotalTime(v int64)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *ApiPullRequestResult) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


