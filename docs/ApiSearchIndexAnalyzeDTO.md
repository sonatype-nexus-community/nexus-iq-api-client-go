# ApiSearchIndexAnalyzeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveJobId** | Pointer to **string** |  | [optional] 
**AdvancedSearchEnabled** | Pointer to **bool** |  | [optional] 
**ApplicationCount** | Pointer to **int64** |  | [optional] 
**BuildingGenerationId** | Pointer to **string** |  | [optional] 
**ComponentCount** | Pointer to **int64** |  | [optional] 
**EstateCapturedAt** | Pointer to **time.Time** |  | [optional] 
**EtaHighMinutes** | Pointer to **int32** |  | [optional] 
**EtaLowMinutes** | Pointer to **int32** |  | [optional] 
**FailedChangeCount** | Pointer to **int64** |  | [optional] 
**HealthStatus** | Pointer to **string** |  | [optional] 
**LastCleanupAt** | Pointer to **time.Time** |  | [optional] 
**LastSuccessfulCutoverAt** | Pointer to **time.Time** |  | [optional] 
**NouxUnlockState** | Pointer to **string** |  | [optional] 
**PendingChangeCount** | Pointer to **int64** |  | [optional] 
**QueueLagSeconds** | Pointer to **int64** |  | [optional] 
**RecommendedOp** | Pointer to **string** |  | [optional] 
**ServingGenerationId** | Pointer to **string** |  | [optional] 
**ViolationCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewApiSearchIndexAnalyzeDTO

`func NewApiSearchIndexAnalyzeDTO() *ApiSearchIndexAnalyzeDTO`

NewApiSearchIndexAnalyzeDTO instantiates a new ApiSearchIndexAnalyzeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchIndexAnalyzeDTOWithDefaults

`func NewApiSearchIndexAnalyzeDTOWithDefaults() *ApiSearchIndexAnalyzeDTO`

NewApiSearchIndexAnalyzeDTOWithDefaults instantiates a new ApiSearchIndexAnalyzeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveJobId

`func (o *ApiSearchIndexAnalyzeDTO) GetActiveJobId() string`

GetActiveJobId returns the ActiveJobId field if non-nil, zero value otherwise.

### GetActiveJobIdOk

`func (o *ApiSearchIndexAnalyzeDTO) GetActiveJobIdOk() (*string, bool)`

GetActiveJobIdOk returns a tuple with the ActiveJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveJobId

`func (o *ApiSearchIndexAnalyzeDTO) SetActiveJobId(v string)`

SetActiveJobId sets ActiveJobId field to given value.

### HasActiveJobId

`func (o *ApiSearchIndexAnalyzeDTO) HasActiveJobId() bool`

HasActiveJobId returns a boolean if a field has been set.

### GetAdvancedSearchEnabled

`func (o *ApiSearchIndexAnalyzeDTO) GetAdvancedSearchEnabled() bool`

GetAdvancedSearchEnabled returns the AdvancedSearchEnabled field if non-nil, zero value otherwise.

### GetAdvancedSearchEnabledOk

`func (o *ApiSearchIndexAnalyzeDTO) GetAdvancedSearchEnabledOk() (*bool, bool)`

GetAdvancedSearchEnabledOk returns a tuple with the AdvancedSearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSearchEnabled

`func (o *ApiSearchIndexAnalyzeDTO) SetAdvancedSearchEnabled(v bool)`

SetAdvancedSearchEnabled sets AdvancedSearchEnabled field to given value.

### HasAdvancedSearchEnabled

`func (o *ApiSearchIndexAnalyzeDTO) HasAdvancedSearchEnabled() bool`

HasAdvancedSearchEnabled returns a boolean if a field has been set.

### GetApplicationCount

`func (o *ApiSearchIndexAnalyzeDTO) GetApplicationCount() int64`

GetApplicationCount returns the ApplicationCount field if non-nil, zero value otherwise.

### GetApplicationCountOk

`func (o *ApiSearchIndexAnalyzeDTO) GetApplicationCountOk() (*int64, bool)`

GetApplicationCountOk returns a tuple with the ApplicationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCount

`func (o *ApiSearchIndexAnalyzeDTO) SetApplicationCount(v int64)`

SetApplicationCount sets ApplicationCount field to given value.

### HasApplicationCount

`func (o *ApiSearchIndexAnalyzeDTO) HasApplicationCount() bool`

HasApplicationCount returns a boolean if a field has been set.

### GetBuildingGenerationId

`func (o *ApiSearchIndexAnalyzeDTO) GetBuildingGenerationId() string`

GetBuildingGenerationId returns the BuildingGenerationId field if non-nil, zero value otherwise.

### GetBuildingGenerationIdOk

`func (o *ApiSearchIndexAnalyzeDTO) GetBuildingGenerationIdOk() (*string, bool)`

GetBuildingGenerationIdOk returns a tuple with the BuildingGenerationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingGenerationId

`func (o *ApiSearchIndexAnalyzeDTO) SetBuildingGenerationId(v string)`

SetBuildingGenerationId sets BuildingGenerationId field to given value.

### HasBuildingGenerationId

`func (o *ApiSearchIndexAnalyzeDTO) HasBuildingGenerationId() bool`

HasBuildingGenerationId returns a boolean if a field has been set.

### GetComponentCount

`func (o *ApiSearchIndexAnalyzeDTO) GetComponentCount() int64`

GetComponentCount returns the ComponentCount field if non-nil, zero value otherwise.

### GetComponentCountOk

`func (o *ApiSearchIndexAnalyzeDTO) GetComponentCountOk() (*int64, bool)`

GetComponentCountOk returns a tuple with the ComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentCount

`func (o *ApiSearchIndexAnalyzeDTO) SetComponentCount(v int64)`

SetComponentCount sets ComponentCount field to given value.

### HasComponentCount

`func (o *ApiSearchIndexAnalyzeDTO) HasComponentCount() bool`

HasComponentCount returns a boolean if a field has been set.

### GetEstateCapturedAt

`func (o *ApiSearchIndexAnalyzeDTO) GetEstateCapturedAt() time.Time`

GetEstateCapturedAt returns the EstateCapturedAt field if non-nil, zero value otherwise.

### GetEstateCapturedAtOk

`func (o *ApiSearchIndexAnalyzeDTO) GetEstateCapturedAtOk() (*time.Time, bool)`

GetEstateCapturedAtOk returns a tuple with the EstateCapturedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstateCapturedAt

`func (o *ApiSearchIndexAnalyzeDTO) SetEstateCapturedAt(v time.Time)`

SetEstateCapturedAt sets EstateCapturedAt field to given value.

### HasEstateCapturedAt

`func (o *ApiSearchIndexAnalyzeDTO) HasEstateCapturedAt() bool`

HasEstateCapturedAt returns a boolean if a field has been set.

### GetEtaHighMinutes

`func (o *ApiSearchIndexAnalyzeDTO) GetEtaHighMinutes() int32`

GetEtaHighMinutes returns the EtaHighMinutes field if non-nil, zero value otherwise.

### GetEtaHighMinutesOk

`func (o *ApiSearchIndexAnalyzeDTO) GetEtaHighMinutesOk() (*int32, bool)`

GetEtaHighMinutesOk returns a tuple with the EtaHighMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtaHighMinutes

`func (o *ApiSearchIndexAnalyzeDTO) SetEtaHighMinutes(v int32)`

SetEtaHighMinutes sets EtaHighMinutes field to given value.

### HasEtaHighMinutes

`func (o *ApiSearchIndexAnalyzeDTO) HasEtaHighMinutes() bool`

HasEtaHighMinutes returns a boolean if a field has been set.

### GetEtaLowMinutes

`func (o *ApiSearchIndexAnalyzeDTO) GetEtaLowMinutes() int32`

GetEtaLowMinutes returns the EtaLowMinutes field if non-nil, zero value otherwise.

### GetEtaLowMinutesOk

`func (o *ApiSearchIndexAnalyzeDTO) GetEtaLowMinutesOk() (*int32, bool)`

GetEtaLowMinutesOk returns a tuple with the EtaLowMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtaLowMinutes

`func (o *ApiSearchIndexAnalyzeDTO) SetEtaLowMinutes(v int32)`

SetEtaLowMinutes sets EtaLowMinutes field to given value.

### HasEtaLowMinutes

`func (o *ApiSearchIndexAnalyzeDTO) HasEtaLowMinutes() bool`

HasEtaLowMinutes returns a boolean if a field has been set.

### GetFailedChangeCount

`func (o *ApiSearchIndexAnalyzeDTO) GetFailedChangeCount() int64`

GetFailedChangeCount returns the FailedChangeCount field if non-nil, zero value otherwise.

### GetFailedChangeCountOk

`func (o *ApiSearchIndexAnalyzeDTO) GetFailedChangeCountOk() (*int64, bool)`

GetFailedChangeCountOk returns a tuple with the FailedChangeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedChangeCount

`func (o *ApiSearchIndexAnalyzeDTO) SetFailedChangeCount(v int64)`

SetFailedChangeCount sets FailedChangeCount field to given value.

### HasFailedChangeCount

`func (o *ApiSearchIndexAnalyzeDTO) HasFailedChangeCount() bool`

HasFailedChangeCount returns a boolean if a field has been set.

### GetHealthStatus

`func (o *ApiSearchIndexAnalyzeDTO) GetHealthStatus() string`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *ApiSearchIndexAnalyzeDTO) GetHealthStatusOk() (*string, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *ApiSearchIndexAnalyzeDTO) SetHealthStatus(v string)`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *ApiSearchIndexAnalyzeDTO) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.

### GetLastCleanupAt

`func (o *ApiSearchIndexAnalyzeDTO) GetLastCleanupAt() time.Time`

GetLastCleanupAt returns the LastCleanupAt field if non-nil, zero value otherwise.

### GetLastCleanupAtOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastCleanupAtOk() (*time.Time, bool)`

GetLastCleanupAtOk returns a tuple with the LastCleanupAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCleanupAt

`func (o *ApiSearchIndexAnalyzeDTO) SetLastCleanupAt(v time.Time)`

SetLastCleanupAt sets LastCleanupAt field to given value.

### HasLastCleanupAt

`func (o *ApiSearchIndexAnalyzeDTO) HasLastCleanupAt() bool`

HasLastCleanupAt returns a boolean if a field has been set.

### GetLastSuccessfulCutoverAt

`func (o *ApiSearchIndexAnalyzeDTO) GetLastSuccessfulCutoverAt() time.Time`

GetLastSuccessfulCutoverAt returns the LastSuccessfulCutoverAt field if non-nil, zero value otherwise.

### GetLastSuccessfulCutoverAtOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastSuccessfulCutoverAtOk() (*time.Time, bool)`

GetLastSuccessfulCutoverAtOk returns a tuple with the LastSuccessfulCutoverAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulCutoverAt

`func (o *ApiSearchIndexAnalyzeDTO) SetLastSuccessfulCutoverAt(v time.Time)`

SetLastSuccessfulCutoverAt sets LastSuccessfulCutoverAt field to given value.

### HasLastSuccessfulCutoverAt

`func (o *ApiSearchIndexAnalyzeDTO) HasLastSuccessfulCutoverAt() bool`

HasLastSuccessfulCutoverAt returns a boolean if a field has been set.

### GetNouxUnlockState

`func (o *ApiSearchIndexAnalyzeDTO) GetNouxUnlockState() string`

GetNouxUnlockState returns the NouxUnlockState field if non-nil, zero value otherwise.

### GetNouxUnlockStateOk

`func (o *ApiSearchIndexAnalyzeDTO) GetNouxUnlockStateOk() (*string, bool)`

GetNouxUnlockStateOk returns a tuple with the NouxUnlockState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNouxUnlockState

`func (o *ApiSearchIndexAnalyzeDTO) SetNouxUnlockState(v string)`

SetNouxUnlockState sets NouxUnlockState field to given value.

### HasNouxUnlockState

`func (o *ApiSearchIndexAnalyzeDTO) HasNouxUnlockState() bool`

HasNouxUnlockState returns a boolean if a field has been set.

### GetPendingChangeCount

`func (o *ApiSearchIndexAnalyzeDTO) GetPendingChangeCount() int64`

GetPendingChangeCount returns the PendingChangeCount field if non-nil, zero value otherwise.

### GetPendingChangeCountOk

`func (o *ApiSearchIndexAnalyzeDTO) GetPendingChangeCountOk() (*int64, bool)`

GetPendingChangeCountOk returns a tuple with the PendingChangeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingChangeCount

`func (o *ApiSearchIndexAnalyzeDTO) SetPendingChangeCount(v int64)`

SetPendingChangeCount sets PendingChangeCount field to given value.

### HasPendingChangeCount

`func (o *ApiSearchIndexAnalyzeDTO) HasPendingChangeCount() bool`

HasPendingChangeCount returns a boolean if a field has been set.

### GetQueueLagSeconds

`func (o *ApiSearchIndexAnalyzeDTO) GetQueueLagSeconds() int64`

GetQueueLagSeconds returns the QueueLagSeconds field if non-nil, zero value otherwise.

### GetQueueLagSecondsOk

`func (o *ApiSearchIndexAnalyzeDTO) GetQueueLagSecondsOk() (*int64, bool)`

GetQueueLagSecondsOk returns a tuple with the QueueLagSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueLagSeconds

`func (o *ApiSearchIndexAnalyzeDTO) SetQueueLagSeconds(v int64)`

SetQueueLagSeconds sets QueueLagSeconds field to given value.

### HasQueueLagSeconds

`func (o *ApiSearchIndexAnalyzeDTO) HasQueueLagSeconds() bool`

HasQueueLagSeconds returns a boolean if a field has been set.

### GetRecommendedOp

`func (o *ApiSearchIndexAnalyzeDTO) GetRecommendedOp() string`

GetRecommendedOp returns the RecommendedOp field if non-nil, zero value otherwise.

### GetRecommendedOpOk

`func (o *ApiSearchIndexAnalyzeDTO) GetRecommendedOpOk() (*string, bool)`

GetRecommendedOpOk returns a tuple with the RecommendedOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedOp

`func (o *ApiSearchIndexAnalyzeDTO) SetRecommendedOp(v string)`

SetRecommendedOp sets RecommendedOp field to given value.

### HasRecommendedOp

`func (o *ApiSearchIndexAnalyzeDTO) HasRecommendedOp() bool`

HasRecommendedOp returns a boolean if a field has been set.

### GetServingGenerationId

`func (o *ApiSearchIndexAnalyzeDTO) GetServingGenerationId() string`

GetServingGenerationId returns the ServingGenerationId field if non-nil, zero value otherwise.

### GetServingGenerationIdOk

`func (o *ApiSearchIndexAnalyzeDTO) GetServingGenerationIdOk() (*string, bool)`

GetServingGenerationIdOk returns a tuple with the ServingGenerationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingGenerationId

`func (o *ApiSearchIndexAnalyzeDTO) SetServingGenerationId(v string)`

SetServingGenerationId sets ServingGenerationId field to given value.

### HasServingGenerationId

`func (o *ApiSearchIndexAnalyzeDTO) HasServingGenerationId() bool`

HasServingGenerationId returns a boolean if a field has been set.

### GetViolationCount

`func (o *ApiSearchIndexAnalyzeDTO) GetViolationCount() int64`

GetViolationCount returns the ViolationCount field if non-nil, zero value otherwise.

### GetViolationCountOk

`func (o *ApiSearchIndexAnalyzeDTO) GetViolationCountOk() (*int64, bool)`

GetViolationCountOk returns a tuple with the ViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationCount

`func (o *ApiSearchIndexAnalyzeDTO) SetViolationCount(v int64)`

SetViolationCount sets ViolationCount field to given value.

### HasViolationCount

`func (o *ApiSearchIndexAnalyzeDTO) HasViolationCount() bool`

HasViolationCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


