# ApiDataRetentionPoliciesDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationReports** | Pointer to [**ApiReportRetentionPoliciesDTO**](ApiReportRetentionPoliciesDTO.md) |  | [optional] 
**SuccessMetrics** | Pointer to [**ApiSuccessMetricsRetentionPolicyDTO**](ApiSuccessMetricsRetentionPolicyDTO.md) |  | [optional] 

## Methods

### NewApiDataRetentionPoliciesDTO

`func NewApiDataRetentionPoliciesDTO() *ApiDataRetentionPoliciesDTO`

NewApiDataRetentionPoliciesDTO instantiates a new ApiDataRetentionPoliciesDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDataRetentionPoliciesDTOWithDefaults

`func NewApiDataRetentionPoliciesDTOWithDefaults() *ApiDataRetentionPoliciesDTO`

NewApiDataRetentionPoliciesDTOWithDefaults instantiates a new ApiDataRetentionPoliciesDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationReports

`func (o *ApiDataRetentionPoliciesDTO) GetApplicationReports() ApiReportRetentionPoliciesDTO`

GetApplicationReports returns the ApplicationReports field if non-nil, zero value otherwise.

### GetApplicationReportsOk

`func (o *ApiDataRetentionPoliciesDTO) GetApplicationReportsOk() (*ApiReportRetentionPoliciesDTO, bool)`

GetApplicationReportsOk returns a tuple with the ApplicationReports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationReports

`func (o *ApiDataRetentionPoliciesDTO) SetApplicationReports(v ApiReportRetentionPoliciesDTO)`

SetApplicationReports sets ApplicationReports field to given value.

### HasApplicationReports

`func (o *ApiDataRetentionPoliciesDTO) HasApplicationReports() bool`

HasApplicationReports returns a boolean if a field has been set.

### GetSuccessMetrics

`func (o *ApiDataRetentionPoliciesDTO) GetSuccessMetrics() ApiSuccessMetricsRetentionPolicyDTO`

GetSuccessMetrics returns the SuccessMetrics field if non-nil, zero value otherwise.

### GetSuccessMetricsOk

`func (o *ApiDataRetentionPoliciesDTO) GetSuccessMetricsOk() (*ApiSuccessMetricsRetentionPolicyDTO, bool)`

GetSuccessMetricsOk returns a tuple with the SuccessMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessMetrics

`func (o *ApiDataRetentionPoliciesDTO) SetSuccessMetrics(v ApiSuccessMetricsRetentionPolicyDTO)`

SetSuccessMetrics sets SuccessMetrics field to given value.

### HasSuccessMetrics

`func (o *ApiDataRetentionPoliciesDTO) HasSuccessMetrics() bool`

HasSuccessMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


