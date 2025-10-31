# ApiReportPolicyDataDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**ApiApplicationBaseDTO**](ApiApplicationBaseDTO.md) |  | [optional] 
**CommitHash** | Pointer to **string** |  | [optional] 
**Components** | Pointer to [**[]ApiReportComponentPolicyViolationsDTOV2**](ApiReportComponentPolicyViolationsDTOV2.md) |  | [optional] 
**Counts** | Pointer to **map[string]int32** |  | [optional] 
**Initiator** | Pointer to **string** |  | [optional] 
**ReportTime** | Pointer to **float32** |  | [optional] 
**ReportTitle** | Pointer to **string** |  | [optional] 

## Methods

### NewApiReportPolicyDataDTOV2

`func NewApiReportPolicyDataDTOV2() *ApiReportPolicyDataDTOV2`

NewApiReportPolicyDataDTOV2 instantiates a new ApiReportPolicyDataDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiReportPolicyDataDTOV2WithDefaults

`func NewApiReportPolicyDataDTOV2WithDefaults() *ApiReportPolicyDataDTOV2`

NewApiReportPolicyDataDTOV2WithDefaults instantiates a new ApiReportPolicyDataDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ApiReportPolicyDataDTOV2) GetApplication() ApiApplicationBaseDTO`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApiReportPolicyDataDTOV2) GetApplicationOk() (*ApiApplicationBaseDTO, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApiReportPolicyDataDTOV2) SetApplication(v ApiApplicationBaseDTO)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApiReportPolicyDataDTOV2) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCommitHash

`func (o *ApiReportPolicyDataDTOV2) GetCommitHash() string`

GetCommitHash returns the CommitHash field if non-nil, zero value otherwise.

### GetCommitHashOk

`func (o *ApiReportPolicyDataDTOV2) GetCommitHashOk() (*string, bool)`

GetCommitHashOk returns a tuple with the CommitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitHash

`func (o *ApiReportPolicyDataDTOV2) SetCommitHash(v string)`

SetCommitHash sets CommitHash field to given value.

### HasCommitHash

`func (o *ApiReportPolicyDataDTOV2) HasCommitHash() bool`

HasCommitHash returns a boolean if a field has been set.

### GetComponents

`func (o *ApiReportPolicyDataDTOV2) GetComponents() []ApiReportComponentPolicyViolationsDTOV2`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ApiReportPolicyDataDTOV2) GetComponentsOk() (*[]ApiReportComponentPolicyViolationsDTOV2, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ApiReportPolicyDataDTOV2) SetComponents(v []ApiReportComponentPolicyViolationsDTOV2)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ApiReportPolicyDataDTOV2) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetCounts

`func (o *ApiReportPolicyDataDTOV2) GetCounts() map[string]int32`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *ApiReportPolicyDataDTOV2) GetCountsOk() (*map[string]int32, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *ApiReportPolicyDataDTOV2) SetCounts(v map[string]int32)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *ApiReportPolicyDataDTOV2) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### GetInitiator

`func (o *ApiReportPolicyDataDTOV2) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *ApiReportPolicyDataDTOV2) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *ApiReportPolicyDataDTOV2) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *ApiReportPolicyDataDTOV2) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetReportTime

`func (o *ApiReportPolicyDataDTOV2) GetReportTime() float32`

GetReportTime returns the ReportTime field if non-nil, zero value otherwise.

### GetReportTimeOk

`func (o *ApiReportPolicyDataDTOV2) GetReportTimeOk() (*float32, bool)`

GetReportTimeOk returns a tuple with the ReportTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTime

`func (o *ApiReportPolicyDataDTOV2) SetReportTime(v float32)`

SetReportTime sets ReportTime field to given value.

### HasReportTime

`func (o *ApiReportPolicyDataDTOV2) HasReportTime() bool`

HasReportTime returns a boolean if a field has been set.

### GetReportTitle

`func (o *ApiReportPolicyDataDTOV2) GetReportTitle() string`

GetReportTitle returns the ReportTitle field if non-nil, zero value otherwise.

### GetReportTitleOk

`func (o *ApiReportPolicyDataDTOV2) GetReportTitleOk() (*string, bool)`

GetReportTitleOk returns a tuple with the ReportTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportTitle

`func (o *ApiReportPolicyDataDTOV2) SetReportTitle(v string)`

SetReportTitle sets ReportTitle field to given value.

### HasReportTitle

`func (o *ApiReportPolicyDataDTOV2) HasReportTitle() bool`

HasReportTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


