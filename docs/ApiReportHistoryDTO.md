# ApiReportHistoryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**Reports** | Pointer to [**[]ApiReportResultsDTO**](ApiReportResultsDTO.md) |  | [optional] 

## Methods

### NewApiReportHistoryDTO

`func NewApiReportHistoryDTO() *ApiReportHistoryDTO`

NewApiReportHistoryDTO instantiates a new ApiReportHistoryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiReportHistoryDTOWithDefaults

`func NewApiReportHistoryDTOWithDefaults() *ApiReportHistoryDTO`

NewApiReportHistoryDTOWithDefaults instantiates a new ApiReportHistoryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApiReportHistoryDTO) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApiReportHistoryDTO) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApiReportHistoryDTO) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApiReportHistoryDTO) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetReports

`func (o *ApiReportHistoryDTO) GetReports() []ApiReportResultsDTO`

GetReports returns the Reports field if non-nil, zero value otherwise.

### GetReportsOk

`func (o *ApiReportHistoryDTO) GetReportsOk() (*[]ApiReportResultsDTO, bool)`

GetReportsOk returns a tuple with the Reports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReports

`func (o *ApiReportHistoryDTO) SetReports(v []ApiReportResultsDTO)`

SetReports sets Reports field to given value.

### HasReports

`func (o *ApiReportHistoryDTO) HasReports() bool`

HasReports returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


