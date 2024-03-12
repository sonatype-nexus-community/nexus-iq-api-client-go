# ApiThirdPartyScanResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentsAffected** | Pointer to [**ApiEvaluationResultCounterDTO**](ApiEvaluationResultCounterDTO.md) |  | [optional] 
**EmbeddableReportHtmlUrl** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**GrandfatheredPolicyViolations** | Pointer to **int32** |  | [optional] 
**IsError** | Pointer to **bool** |  | [optional] 
**LegacyViolations** | Pointer to **int32** |  | [optional] 
**OpenPolicyViolations** | Pointer to [**ApiEvaluationResultCounterDTO**](ApiEvaluationResultCounterDTO.md) |  | [optional] 
**PolicyAction** | Pointer to **string** |  | [optional] 
**ReportDataUrl** | Pointer to **string** |  | [optional] 
**ReportHtmlUrl** | Pointer to **string** |  | [optional] 
**ReportPdfUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewApiThirdPartyScanResultDTO

`func NewApiThirdPartyScanResultDTO() *ApiThirdPartyScanResultDTO`

NewApiThirdPartyScanResultDTO instantiates a new ApiThirdPartyScanResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiThirdPartyScanResultDTOWithDefaults

`func NewApiThirdPartyScanResultDTOWithDefaults() *ApiThirdPartyScanResultDTO`

NewApiThirdPartyScanResultDTOWithDefaults instantiates a new ApiThirdPartyScanResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentsAffected

`func (o *ApiThirdPartyScanResultDTO) GetComponentsAffected() ApiEvaluationResultCounterDTO`

GetComponentsAffected returns the ComponentsAffected field if non-nil, zero value otherwise.

### GetComponentsAffectedOk

`func (o *ApiThirdPartyScanResultDTO) GetComponentsAffectedOk() (*ApiEvaluationResultCounterDTO, bool)`

GetComponentsAffectedOk returns a tuple with the ComponentsAffected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentsAffected

`func (o *ApiThirdPartyScanResultDTO) SetComponentsAffected(v ApiEvaluationResultCounterDTO)`

SetComponentsAffected sets ComponentsAffected field to given value.

### HasComponentsAffected

`func (o *ApiThirdPartyScanResultDTO) HasComponentsAffected() bool`

HasComponentsAffected returns a boolean if a field has been set.

### GetEmbeddableReportHtmlUrl

`func (o *ApiThirdPartyScanResultDTO) GetEmbeddableReportHtmlUrl() string`

GetEmbeddableReportHtmlUrl returns the EmbeddableReportHtmlUrl field if non-nil, zero value otherwise.

### GetEmbeddableReportHtmlUrlOk

`func (o *ApiThirdPartyScanResultDTO) GetEmbeddableReportHtmlUrlOk() (*string, bool)`

GetEmbeddableReportHtmlUrlOk returns a tuple with the EmbeddableReportHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddableReportHtmlUrl

`func (o *ApiThirdPartyScanResultDTO) SetEmbeddableReportHtmlUrl(v string)`

SetEmbeddableReportHtmlUrl sets EmbeddableReportHtmlUrl field to given value.

### HasEmbeddableReportHtmlUrl

`func (o *ApiThirdPartyScanResultDTO) HasEmbeddableReportHtmlUrl() bool`

HasEmbeddableReportHtmlUrl returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ApiThirdPartyScanResultDTO) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiThirdPartyScanResultDTO) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiThirdPartyScanResultDTO) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiThirdPartyScanResultDTO) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetGrandfatheredPolicyViolations

`func (o *ApiThirdPartyScanResultDTO) GetGrandfatheredPolicyViolations() int32`

GetGrandfatheredPolicyViolations returns the GrandfatheredPolicyViolations field if non-nil, zero value otherwise.

### GetGrandfatheredPolicyViolationsOk

`func (o *ApiThirdPartyScanResultDTO) GetGrandfatheredPolicyViolationsOk() (*int32, bool)`

GetGrandfatheredPolicyViolationsOk returns a tuple with the GrandfatheredPolicyViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandfatheredPolicyViolations

`func (o *ApiThirdPartyScanResultDTO) SetGrandfatheredPolicyViolations(v int32)`

SetGrandfatheredPolicyViolations sets GrandfatheredPolicyViolations field to given value.

### HasGrandfatheredPolicyViolations

`func (o *ApiThirdPartyScanResultDTO) HasGrandfatheredPolicyViolations() bool`

HasGrandfatheredPolicyViolations returns a boolean if a field has been set.

### GetIsError

`func (o *ApiThirdPartyScanResultDTO) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *ApiThirdPartyScanResultDTO) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *ApiThirdPartyScanResultDTO) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *ApiThirdPartyScanResultDTO) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetLegacyViolations

`func (o *ApiThirdPartyScanResultDTO) GetLegacyViolations() int32`

GetLegacyViolations returns the LegacyViolations field if non-nil, zero value otherwise.

### GetLegacyViolationsOk

`func (o *ApiThirdPartyScanResultDTO) GetLegacyViolationsOk() (*int32, bool)`

GetLegacyViolationsOk returns a tuple with the LegacyViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyViolations

`func (o *ApiThirdPartyScanResultDTO) SetLegacyViolations(v int32)`

SetLegacyViolations sets LegacyViolations field to given value.

### HasLegacyViolations

`func (o *ApiThirdPartyScanResultDTO) HasLegacyViolations() bool`

HasLegacyViolations returns a boolean if a field has been set.

### GetOpenPolicyViolations

`func (o *ApiThirdPartyScanResultDTO) GetOpenPolicyViolations() ApiEvaluationResultCounterDTO`

GetOpenPolicyViolations returns the OpenPolicyViolations field if non-nil, zero value otherwise.

### GetOpenPolicyViolationsOk

`func (o *ApiThirdPartyScanResultDTO) GetOpenPolicyViolationsOk() (*ApiEvaluationResultCounterDTO, bool)`

GetOpenPolicyViolationsOk returns a tuple with the OpenPolicyViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPolicyViolations

`func (o *ApiThirdPartyScanResultDTO) SetOpenPolicyViolations(v ApiEvaluationResultCounterDTO)`

SetOpenPolicyViolations sets OpenPolicyViolations field to given value.

### HasOpenPolicyViolations

`func (o *ApiThirdPartyScanResultDTO) HasOpenPolicyViolations() bool`

HasOpenPolicyViolations returns a boolean if a field has been set.

### GetPolicyAction

`func (o *ApiThirdPartyScanResultDTO) GetPolicyAction() string`

GetPolicyAction returns the PolicyAction field if non-nil, zero value otherwise.

### GetPolicyActionOk

`func (o *ApiThirdPartyScanResultDTO) GetPolicyActionOk() (*string, bool)`

GetPolicyActionOk returns a tuple with the PolicyAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyAction

`func (o *ApiThirdPartyScanResultDTO) SetPolicyAction(v string)`

SetPolicyAction sets PolicyAction field to given value.

### HasPolicyAction

`func (o *ApiThirdPartyScanResultDTO) HasPolicyAction() bool`

HasPolicyAction returns a boolean if a field has been set.

### GetReportDataUrl

`func (o *ApiThirdPartyScanResultDTO) GetReportDataUrl() string`

GetReportDataUrl returns the ReportDataUrl field if non-nil, zero value otherwise.

### GetReportDataUrlOk

`func (o *ApiThirdPartyScanResultDTO) GetReportDataUrlOk() (*string, bool)`

GetReportDataUrlOk returns a tuple with the ReportDataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDataUrl

`func (o *ApiThirdPartyScanResultDTO) SetReportDataUrl(v string)`

SetReportDataUrl sets ReportDataUrl field to given value.

### HasReportDataUrl

`func (o *ApiThirdPartyScanResultDTO) HasReportDataUrl() bool`

HasReportDataUrl returns a boolean if a field has been set.

### GetReportHtmlUrl

`func (o *ApiThirdPartyScanResultDTO) GetReportHtmlUrl() string`

GetReportHtmlUrl returns the ReportHtmlUrl field if non-nil, zero value otherwise.

### GetReportHtmlUrlOk

`func (o *ApiThirdPartyScanResultDTO) GetReportHtmlUrlOk() (*string, bool)`

GetReportHtmlUrlOk returns a tuple with the ReportHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportHtmlUrl

`func (o *ApiThirdPartyScanResultDTO) SetReportHtmlUrl(v string)`

SetReportHtmlUrl sets ReportHtmlUrl field to given value.

### HasReportHtmlUrl

`func (o *ApiThirdPartyScanResultDTO) HasReportHtmlUrl() bool`

HasReportHtmlUrl returns a boolean if a field has been set.

### GetReportPdfUrl

`func (o *ApiThirdPartyScanResultDTO) GetReportPdfUrl() string`

GetReportPdfUrl returns the ReportPdfUrl field if non-nil, zero value otherwise.

### GetReportPdfUrlOk

`func (o *ApiThirdPartyScanResultDTO) GetReportPdfUrlOk() (*string, bool)`

GetReportPdfUrlOk returns a tuple with the ReportPdfUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportPdfUrl

`func (o *ApiThirdPartyScanResultDTO) SetReportPdfUrl(v string)`

SetReportPdfUrl sets ReportPdfUrl field to given value.

### HasReportPdfUrl

`func (o *ApiThirdPartyScanResultDTO) HasReportPdfUrl() bool`

HasReportPdfUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


