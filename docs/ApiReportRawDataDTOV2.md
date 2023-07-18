# ApiReportRawDataDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]ApiReportComponentDTOV2**](ApiReportComponentDTOV2.md) |  | [optional] 
**MatchSummary** | Pointer to [**ApiMatchStateSummaryDTOV2**](ApiMatchStateSummaryDTOV2.md) |  | [optional] 

## Methods

### NewApiReportRawDataDTOV2

`func NewApiReportRawDataDTOV2() *ApiReportRawDataDTOV2`

NewApiReportRawDataDTOV2 instantiates a new ApiReportRawDataDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiReportRawDataDTOV2WithDefaults

`func NewApiReportRawDataDTOV2WithDefaults() *ApiReportRawDataDTOV2`

NewApiReportRawDataDTOV2WithDefaults instantiates a new ApiReportRawDataDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *ApiReportRawDataDTOV2) GetComponents() []ApiReportComponentDTOV2`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ApiReportRawDataDTOV2) GetComponentsOk() (*[]ApiReportComponentDTOV2, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ApiReportRawDataDTOV2) SetComponents(v []ApiReportComponentDTOV2)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ApiReportRawDataDTOV2) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetMatchSummary

`func (o *ApiReportRawDataDTOV2) GetMatchSummary() ApiMatchStateSummaryDTOV2`

GetMatchSummary returns the MatchSummary field if non-nil, zero value otherwise.

### GetMatchSummaryOk

`func (o *ApiReportRawDataDTOV2) GetMatchSummaryOk() (*ApiMatchStateSummaryDTOV2, bool)`

GetMatchSummaryOk returns a tuple with the MatchSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchSummary

`func (o *ApiReportRawDataDTOV2) SetMatchSummary(v ApiMatchStateSummaryDTOV2)`

SetMatchSummary sets MatchSummary field to given value.

### HasMatchSummary

`func (o *ApiReportRawDataDTOV2) HasMatchSummary() bool`

HasMatchSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


