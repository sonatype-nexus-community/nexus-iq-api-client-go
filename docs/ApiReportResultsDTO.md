# ApiReportResultsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**CommitHash** | Pointer to **string** |  | [optional] 
**EmbeddableReportHtmlUrl** | Pointer to **string** |  | [optional] 
**EvaluationDate** | Pointer to **time.Time** |  | [optional] 
**IsForMonitoring** | Pointer to **bool** |  | [optional] 
**IsReevaluation** | Pointer to **bool** |  | [optional] 
**LatestReportHtmlUrl** | Pointer to **string** |  | [optional] 
**PolicyEvaluationId** | Pointer to **string** |  | [optional] 
**PolicyEvaluationResult** | Pointer to [**PolicyEvaluationResult**](PolicyEvaluationResult.md) |  | [optional] 
**ReportDataUrl** | Pointer to **string** |  | [optional] 
**ReportHtmlUrl** | Pointer to **string** |  | [optional] 
**ReportPdfUrl** | Pointer to **string** |  | [optional] 
**ScanId** | Pointer to **string** |  | [optional] 
**Stage** | Pointer to **string** |  | [optional] 

## Methods

### NewApiReportResultsDTO

`func NewApiReportResultsDTO() *ApiReportResultsDTO`

NewApiReportResultsDTO instantiates a new ApiReportResultsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiReportResultsDTOWithDefaults

`func NewApiReportResultsDTOWithDefaults() *ApiReportResultsDTO`

NewApiReportResultsDTOWithDefaults instantiates a new ApiReportResultsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApiReportResultsDTO) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApiReportResultsDTO) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApiReportResultsDTO) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApiReportResultsDTO) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetCommitHash

`func (o *ApiReportResultsDTO) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ApiReportResultsDTO) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ApiReportResultsDTO) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *ApiReportResultsDTO) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetEmbeddableReportHtmlUrl

`func (o *ApiReportResultsDTO) GetEmbeddableReportHtmlUrl() string`

GetEmbeddableReportHtmlUrl returns the EmbeddableReportHtmlUrl field if non-nil, zero value otherwise.

### GetEmbeddableReportHtmlUrlOk

`func (o *ApiReportResultsDTO) GetEmbeddableReportHtmlUrlOk() (*string, bool)`

GetEmbeddableReportHtmlUrlOk returns a tuple with the EmbeddableReportHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddableReportHtmlUrl

`func (o *ApiReportResultsDTO) SetEmbeddableReportHtmlUrl(v string)`

SetEmbeddableReportHtmlUrl sets EmbeddableReportHtmlUrl field to given value.

### HasEmbeddableReportHtmlUrl

`func (o *ApiReportResultsDTO) HasEmbeddableReportHtmlUrl() bool`

HasEmbeddableReportHtmlUrl returns a boolean if a field has been set.

### GetEvaluationDate

`func (o *ApiReportResultsDTO) GetEvaluationDate() time.Time`

GetEvaluationDate returns the EvaluationDate field if non-nil, zero value otherwise.

### GetEvaluationDateOk

`func (o *ApiReportResultsDTO) GetEvaluationDateOk() (*time.Time, bool)`

GetEvaluationDateOk returns a tuple with the EvaluationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDate

`func (o *ApiReportResultsDTO) SetEvaluationDate(v time.Time)`

SetEvaluationDate sets EvaluationDate field to given value.

### HasEvaluationDate

`func (o *ApiReportResultsDTO) HasEvaluationDate() bool`

HasEvaluationDate returns a boolean if a field has been set.

### GetIsForMonitoring

`func (o *ApiReportResultsDTO) GetIsForMonitoring() bool`

GetIsForMonitoring returns the IsForMonitoring field if non-nil, zero value otherwise.

### GetIsForMonitoringOk

`func (o *ApiReportResultsDTO) GetIsForMonitoringOk() (*bool, bool)`

GetIsForMonitoringOk returns a tuple with the IsForMonitoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsForMonitoring

`func (o *ApiReportResultsDTO) SetIsForMonitoring(v bool)`

SetIsForMonitoring sets IsForMonitoring field to given value.

### HasIsForMonitoring

`func (o *ApiReportResultsDTO) HasIsForMonitoring() bool`

HasIsForMonitoring returns a boolean if a field has been set.

### GetIsReevaluation

`func (o *ApiReportResultsDTO) GetIsReevaluation() bool`

GetIsReevaluation returns the IsReevaluation field if non-nil, zero value otherwise.

### GetIsReevaluationOk

`func (o *ApiReportResultsDTO) GetIsReevaluationOk() (*bool, bool)`

GetIsReevaluationOk returns a tuple with the IsReevaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReevaluation

`func (o *ApiReportResultsDTO) SetIsReevaluation(v bool)`

SetIsReevaluation sets IsReevaluation field to given value.

### HasIsReevaluation

`func (o *ApiReportResultsDTO) HasIsReevaluation() bool`

HasIsReevaluation returns a boolean if a field has been set.

### GetLatestReportHtmlUrl

`func (o *ApiReportResultsDTO) GetLatestReportHtmlUrl() string`

GetLatestReportHtmlUrl returns the LatestReportHtmlUrl field if non-nil, zero value otherwise.

### GetLatestReportHtmlUrlOk

`func (o *ApiReportResultsDTO) GetLatestReportHtmlUrlOk() (*string, bool)`

GetLatestReportHtmlUrlOk returns a tuple with the LatestReportHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestReportHtmlUrl

`func (o *ApiReportResultsDTO) SetLatestReportHtmlUrl(v string)`

SetLatestReportHtmlUrl sets LatestReportHtmlUrl field to given value.

### HasLatestReportHtmlUrl

`func (o *ApiReportResultsDTO) HasLatestReportHtmlUrl() bool`

HasLatestReportHtmlUrl returns a boolean if a field has been set.

### GetPolicyEvaluationId

`func (o *ApiReportResultsDTO) GetPolicyEvaluationId() string`

GetPolicyEvaluationId returns the PolicyEvaluationId field if non-nil, zero value otherwise.

### GetPolicyEvaluationIdOk

`func (o *ApiReportResultsDTO) GetPolicyEvaluationIdOk() (*string, bool)`

GetPolicyEvaluationIdOk returns a tuple with the PolicyEvaluationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEvaluationId

`func (o *ApiReportResultsDTO) SetPolicyEvaluationId(v string)`

SetPolicyEvaluationId sets PolicyEvaluationId field to given value.

### HasPolicyEvaluationId

`func (o *ApiReportResultsDTO) HasPolicyEvaluationId() bool`

HasPolicyEvaluationId returns a boolean if a field has been set.

### GetPolicyEvaluationResult

`func (o *ApiReportResultsDTO) GetPolicyEvaluationResult() PolicyEvaluationResult`

GetPolicyEvaluationResult returns the PolicyEvaluationResult field if non-nil, zero value otherwise.

### GetPolicyEvaluationResultOk

`func (o *ApiReportResultsDTO) GetPolicyEvaluationResultOk() (*PolicyEvaluationResult, bool)`

GetPolicyEvaluationResultOk returns a tuple with the PolicyEvaluationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEvaluationResult

`func (o *ApiReportResultsDTO) SetPolicyEvaluationResult(v PolicyEvaluationResult)`

SetPolicyEvaluationResult sets PolicyEvaluationResult field to given value.

### HasPolicyEvaluationResult

`func (o *ApiReportResultsDTO) HasPolicyEvaluationResult() bool`

HasPolicyEvaluationResult returns a boolean if a field has been set.

### GetReportDataUrl

`func (o *ApiReportResultsDTO) GetReportDataUrl() string`

GetReportDataUrl returns the ReportDataUrl field if non-nil, zero value otherwise.

### GetReportDataUrlOk

`func (o *ApiReportResultsDTO) GetReportDataUrlOk() (*string, bool)`

GetReportDataUrlOk returns a tuple with the ReportDataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDataUrl

`func (o *ApiReportResultsDTO) SetReportDataUrl(v string)`

SetReportDataUrl sets ReportDataUrl field to given value.

### HasReportDataUrl

`func (o *ApiReportResultsDTO) HasReportDataUrl() bool`

HasReportDataUrl returns a boolean if a field has been set.

### GetReportHtmlUrl

`func (o *ApiReportResultsDTO) GetReportHtmlUrl() string`

GetReportHtmlUrl returns the ReportHtmlUrl field if non-nil, zero value otherwise.

### GetReportHtmlUrlOk

`func (o *ApiReportResultsDTO) GetReportHtmlUrlOk() (*string, bool)`

GetReportHtmlUrlOk returns a tuple with the ReportHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportHtmlUrl

`func (o *ApiReportResultsDTO) SetReportHtmlUrl(v string)`

SetReportHtmlUrl sets ReportHtmlUrl field to given value.

### HasReportHtmlUrl

`func (o *ApiReportResultsDTO) HasReportHtmlUrl() bool`

HasReportHtmlUrl returns a boolean if a field has been set.

### GetReportPdfUrl

`func (o *ApiReportResultsDTO) GetReportPdfUrl() string`

GetReportPdfUrl returns the ReportPdfUrl field if non-nil, zero value otherwise.

### GetReportPdfUrlOk

`func (o *ApiReportResultsDTO) GetReportPdfUrlOk() (*string, bool)`

GetReportPdfUrlOk returns a tuple with the ReportPdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportPdfUrl

`func (o *ApiReportResultsDTO) SetReportPdfUrl(v string)`

SetReportPdfUrl sets ReportPdfUrl field to given value.

### HasReportPdfUrl

`func (o *ApiReportResultsDTO) HasReportPdfUrl() bool`

HasReportPdfUrl returns a boolean if a field has been set.

### GetScanId

`func (o *ApiReportResultsDTO) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ApiReportResultsDTO) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ApiReportResultsDTO) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *ApiReportResultsDTO) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetStage

`func (o *ApiReportResultsDTO) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *ApiReportResultsDTO) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *ApiReportResultsDTO) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *ApiReportResultsDTO) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


