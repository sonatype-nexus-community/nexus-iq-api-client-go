# ApiEnhancedPolicyViolationDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to [**ApiComponentDTOV2**](ApiComponentDTOV2.md) |  | [optional] 
**ConstraintViolations** | Pointer to [**[]ApiConstraintViolationDTO**](ApiConstraintViolationDTO.md) |  | [optional] 
**OpenTime** | Pointer to **time.Time** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**ReportId** | Pointer to **string** |  | [optional] 
**ReportUrl** | Pointer to **string** |  | [optional] 
**StageId** | Pointer to **string** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiEnhancedPolicyViolationDTOV2

`func NewApiEnhancedPolicyViolationDTOV2() *ApiEnhancedPolicyViolationDTOV2`

NewApiEnhancedPolicyViolationDTOV2 instantiates a new ApiEnhancedPolicyViolationDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiEnhancedPolicyViolationDTOV2WithDefaults

`func NewApiEnhancedPolicyViolationDTOV2WithDefaults() *ApiEnhancedPolicyViolationDTOV2`

NewApiEnhancedPolicyViolationDTOV2WithDefaults instantiates a new ApiEnhancedPolicyViolationDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ApiEnhancedPolicyViolationDTOV2) GetComponent() ApiComponentDTOV2`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ApiEnhancedPolicyViolationDTOV2) GetComponentOk() (*ApiComponentDTOV2, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ApiEnhancedPolicyViolationDTOV2) SetComponent(v ApiComponentDTOV2)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ApiEnhancedPolicyViolationDTOV2) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetConstraintViolations

`func (o *ApiEnhancedPolicyViolationDTOV2) GetConstraintViolations() []ApiConstraintViolationDTO`

GetConstraintViolations returns the ConstraintViolations field if non-nil, zero value otherwise.

### GetConstraintViolationsOk

`func (o *ApiEnhancedPolicyViolationDTOV2) GetConstraintViolationsOk() (*[]ApiConstraintViolationDTO, bool)`

GetConstraintViolationsOk returns a tuple with the ConstraintViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintViolations

`func (o *ApiEnhancedPolicyViolationDTOV2) SetConstraintViolations(v []ApiConstraintViolationDTO)`

SetConstraintViolations sets ConstraintViolations field to given value.

### HasConstraintViolations

`func (o *ApiEnhancedPolicyViolationDTOV2) HasConstraintViolations() bool`

HasConstraintViolations returns a boolean if a field has been set.

### GetOpenTime

`func (o *ApiEnhancedPolicyViolationDTOV2) GetOpenTime() time.Time`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *ApiEnhancedPolicyViolationDTOV2) GetOpenTimeOk() (*time.Time, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *ApiEnhancedPolicyViolationDTOV2) SetOpenTime(v time.Time)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *ApiEnhancedPolicyViolationDTOV2) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApiEnhancedPolicyViolationDTOV2) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApiEnhancedPolicyViolationDTOV2) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApiEnhancedPolicyViolationDTOV2) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApiEnhancedPolicyViolationDTOV2) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiEnhancedPolicyViolationDTOV2) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiEnhancedPolicyViolationDTOV2) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiEnhancedPolicyViolationDTOV2) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiEnhancedPolicyViolationDTOV2) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *ApiEnhancedPolicyViolationDTOV2) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *ApiEnhancedPolicyViolationDTOV2) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *ApiEnhancedPolicyViolationDTOV2) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *ApiEnhancedPolicyViolationDTOV2) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetReportId

`func (o *ApiEnhancedPolicyViolationDTOV2) GetReportId() string`

GetReportId returns the ReportId field if non-nil, zero value otherwise.

### GetReportIdOk

`func (o *ApiEnhancedPolicyViolationDTOV2) GetReportIdOk() (*string, bool)`

GetReportIdOk returns a tuple with the ReportId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportId

`func (o *ApiEnhancedPolicyViolationDTOV2) SetReportId(v string)`

SetReportId sets ReportId field to given value.

### HasReportId

`func (o *ApiEnhancedPolicyViolationDTOV2) HasReportId() bool`

HasReportId returns a boolean if a field has been set.

### GetReportUrl

`func (o *ApiEnhancedPolicyViolationDTOV2) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *ApiEnhancedPolicyViolationDTOV2) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *ApiEnhancedPolicyViolationDTOV2) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *ApiEnhancedPolicyViolationDTOV2) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.

### GetStageId

`func (o *ApiEnhancedPolicyViolationDTOV2) GetStageId() string`

GetStageId returns the StageId field if non-nil, zero value otherwise.

### GetStageIdOk

`func (o *ApiEnhancedPolicyViolationDTOV2) GetStageIdOk() (*string, bool)`

GetStageIdOk returns a tuple with the StageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageId

`func (o *ApiEnhancedPolicyViolationDTOV2) SetStageId(v string)`

SetStageId sets StageId field to given value.

### HasStageId

`func (o *ApiEnhancedPolicyViolationDTOV2) HasStageId() bool`

HasStageId returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiEnhancedPolicyViolationDTOV2) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiEnhancedPolicyViolationDTOV2) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiEnhancedPolicyViolationDTOV2) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiEnhancedPolicyViolationDTOV2) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


