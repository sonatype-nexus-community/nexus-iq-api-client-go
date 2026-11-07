# ApiSearchIndexAnalyzeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveIndexId** | Pointer to **string** |  | [optional] 
**ActiveJobId** | Pointer to **string** |  | [optional] 
**AdvancedSearchEnabled** | Pointer to **bool** |  | [optional] 
**ApplicationCount** | Pointer to **int64** |  | [optional] 
**BuildingGenerationId** | Pointer to **string** |  | [optional] 
**ChurnRatio** | Pointer to **float64** |  | [optional] 
**ComponentCount** | Pointer to **int64** |  | [optional] 
**EstateCapturedAt** | Pointer to **time.Time** |  | [optional] 
**EstateStale** | Pointer to **bool** |  | [optional] 
**EtaHighMinutes** | Pointer to **int32** |  | [optional] 
**EtaLowMinutes** | Pointer to **int32** |  | [optional] 
**FailedChangeCount** | Pointer to **int64** |  | [optional] 
**FailedChangeState** | Pointer to **string** |  | [optional] 
**HealthStatus** | Pointer to **string** |  | [optional] 
**IndexBackend** | Pointer to **string** |  | [optional] 
**IndexIncompatible** | Pointer to **bool** |  | [optional] 
**IndexName** | Pointer to **string** |  | [optional] 
**IndexPath** | Pointer to **string** |  | [optional] 
**IndexStorageRef** | Pointer to **string** |  | [optional] 
**IndexVersion** | Pointer to **int32** |  | [optional] 
**IndexedApplicationCount** | Pointer to **int64** |  | [optional] 
**LastAnalysisCompletedAt** | Pointer to **time.Time** |  | [optional] 
**LastIncrementalDocumentsAdded** | Pointer to **int64** |  | [optional] 
**LastIncrementalDocumentsDeleted** | Pointer to **int64** |  | [optional] 
**LastIncrementalFinishedAt** | Pointer to **time.Time** |  | [optional] 
**LastIndexCompletedAt** | Pointer to **time.Time** |  | [optional] 
**LastIndexId** | Pointer to **string** |  | [optional] 
**LastIndexJobId** | Pointer to **string** |  | [optional] 
**LastIndexJobStatus** | Pointer to **string** |  | [optional] 
**LastIndexJobType** | Pointer to **string** |  | [optional] 
**LastIndexStartedAt** | Pointer to **time.Time** |  | [optional] 
**LastRunState** | Pointer to **string** |  | [optional] 
**LastSuccessfulCutoverAt** | Pointer to **time.Time** |  | [optional] 
**NouxUnlockState** | Pointer to **string** |  | [optional] 
**OwnerApplicationCount** | Pointer to **int64** |  | [optional] 
**PendingChangeCount** | Pointer to **int64** |  | [optional] 
**PendingChangeCountByType** | Pointer to **map[string]int64** |  | [optional] 
**PendingChangeState** | Pointer to **string** |  | [optional] 
**QueueLagSeconds** | Pointer to **int64** |  | [optional] 
**QueueLagState** | Pointer to **string** |  | [optional] 
**RebuildRecommended** | Pointer to **bool** |  | [optional] 
**RebuildRequired** | Pointer to **bool** |  | [optional] 
**RecommendedOp** | Pointer to **string** |  | [optional] 
**SchemaUpgradeRequired** | Pointer to **bool** |  | [optional] 
**SearchableDocumentCount** | Pointer to **int64** |  | [optional] 
**ServingDocCount** | Pointer to **int64** |  | [optional] 
**ServingGenerationId** | Pointer to **string** |  | [optional] 
**ServingIndexId** | Pointer to **string** |  | [optional] 
**ServingSchemaVersion** | Pointer to **int32** |  | [optional] 
**ServingSince** | Pointer to **time.Time** |  | [optional] 
**ServingSizeBytes** | Pointer to **int64** |  | [optional] 
**StampedAtEpochMs** | Pointer to **int64** |  | [optional] 
**StampedIndexVersion** | Pointer to **int32** |  | [optional] 
**StampedIqVersion** | Pointer to **string** |  | [optional] 
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

### GetActiveIndexId

`func (o *ApiSearchIndexAnalyzeDTO) GetActiveIndexId() string`

GetActiveIndexId returns the ActiveIndexId field if non-nil, zero value otherwise.

### GetActiveIndexIdOk

`func (o *ApiSearchIndexAnalyzeDTO) GetActiveIndexIdOk() (*string, bool)`

GetActiveIndexIdOk returns a tuple with the ActiveIndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveIndexId

`func (o *ApiSearchIndexAnalyzeDTO) SetActiveIndexId(v string)`

SetActiveIndexId sets ActiveIndexId field to given value.

### HasActiveIndexId

`func (o *ApiSearchIndexAnalyzeDTO) HasActiveIndexId() bool`

HasActiveIndexId returns a boolean if a field has been set.

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

### GetChurnRatio

`func (o *ApiSearchIndexAnalyzeDTO) GetChurnRatio() float64`

GetChurnRatio returns the ChurnRatio field if non-nil, zero value otherwise.

### GetChurnRatioOk

`func (o *ApiSearchIndexAnalyzeDTO) GetChurnRatioOk() (*float64, bool)`

GetChurnRatioOk returns a tuple with the ChurnRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChurnRatio

`func (o *ApiSearchIndexAnalyzeDTO) SetChurnRatio(v float64)`

SetChurnRatio sets ChurnRatio field to given value.

### HasChurnRatio

`func (o *ApiSearchIndexAnalyzeDTO) HasChurnRatio() bool`

HasChurnRatio returns a boolean if a field has been set.

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

### GetEstateStale

`func (o *ApiSearchIndexAnalyzeDTO) GetEstateStale() bool`

GetEstateStale returns the EstateStale field if non-nil, zero value otherwise.

### GetEstateStaleOk

`func (o *ApiSearchIndexAnalyzeDTO) GetEstateStaleOk() (*bool, bool)`

GetEstateStaleOk returns a tuple with the EstateStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstateStale

`func (o *ApiSearchIndexAnalyzeDTO) SetEstateStale(v bool)`

SetEstateStale sets EstateStale field to given value.

### HasEstateStale

`func (o *ApiSearchIndexAnalyzeDTO) HasEstateStale() bool`

HasEstateStale returns a boolean if a field has been set.

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

### GetFailedChangeState

`func (o *ApiSearchIndexAnalyzeDTO) GetFailedChangeState() string`

GetFailedChangeState returns the FailedChangeState field if non-nil, zero value otherwise.

### GetFailedChangeStateOk

`func (o *ApiSearchIndexAnalyzeDTO) GetFailedChangeStateOk() (*string, bool)`

GetFailedChangeStateOk returns a tuple with the FailedChangeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedChangeState

`func (o *ApiSearchIndexAnalyzeDTO) SetFailedChangeState(v string)`

SetFailedChangeState sets FailedChangeState field to given value.

### HasFailedChangeState

`func (o *ApiSearchIndexAnalyzeDTO) HasFailedChangeState() bool`

HasFailedChangeState returns a boolean if a field has been set.

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

### GetIndexBackend

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexBackend() string`

GetIndexBackend returns the IndexBackend field if non-nil, zero value otherwise.

### GetIndexBackendOk

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexBackendOk() (*string, bool)`

GetIndexBackendOk returns a tuple with the IndexBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexBackend

`func (o *ApiSearchIndexAnalyzeDTO) SetIndexBackend(v string)`

SetIndexBackend sets IndexBackend field to given value.

### HasIndexBackend

`func (o *ApiSearchIndexAnalyzeDTO) HasIndexBackend() bool`

HasIndexBackend returns a boolean if a field has been set.

### GetIndexIncompatible

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexIncompatible() bool`

GetIndexIncompatible returns the IndexIncompatible field if non-nil, zero value otherwise.

### GetIndexIncompatibleOk

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexIncompatibleOk() (*bool, bool)`

GetIndexIncompatibleOk returns a tuple with the IndexIncompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexIncompatible

`func (o *ApiSearchIndexAnalyzeDTO) SetIndexIncompatible(v bool)`

SetIndexIncompatible sets IndexIncompatible field to given value.

### HasIndexIncompatible

`func (o *ApiSearchIndexAnalyzeDTO) HasIndexIncompatible() bool`

HasIndexIncompatible returns a boolean if a field has been set.

### GetIndexName

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *ApiSearchIndexAnalyzeDTO) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.

### HasIndexName

`func (o *ApiSearchIndexAnalyzeDTO) HasIndexName() bool`

HasIndexName returns a boolean if a field has been set.

### GetIndexPath

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexPath() string`

GetIndexPath returns the IndexPath field if non-nil, zero value otherwise.

### GetIndexPathOk

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexPathOk() (*string, bool)`

GetIndexPathOk returns a tuple with the IndexPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexPath

`func (o *ApiSearchIndexAnalyzeDTO) SetIndexPath(v string)`

SetIndexPath sets IndexPath field to given value.

### HasIndexPath

`func (o *ApiSearchIndexAnalyzeDTO) HasIndexPath() bool`

HasIndexPath returns a boolean if a field has been set.

### GetIndexStorageRef

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexStorageRef() string`

GetIndexStorageRef returns the IndexStorageRef field if non-nil, zero value otherwise.

### GetIndexStorageRefOk

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexStorageRefOk() (*string, bool)`

GetIndexStorageRefOk returns a tuple with the IndexStorageRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexStorageRef

`func (o *ApiSearchIndexAnalyzeDTO) SetIndexStorageRef(v string)`

SetIndexStorageRef sets IndexStorageRef field to given value.

### HasIndexStorageRef

`func (o *ApiSearchIndexAnalyzeDTO) HasIndexStorageRef() bool`

HasIndexStorageRef returns a boolean if a field has been set.

### GetIndexVersion

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexVersion() int32`

GetIndexVersion returns the IndexVersion field if non-nil, zero value otherwise.

### GetIndexVersionOk

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexVersionOk() (*int32, bool)`

GetIndexVersionOk returns a tuple with the IndexVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexVersion

`func (o *ApiSearchIndexAnalyzeDTO) SetIndexVersion(v int32)`

SetIndexVersion sets IndexVersion field to given value.

### HasIndexVersion

`func (o *ApiSearchIndexAnalyzeDTO) HasIndexVersion() bool`

HasIndexVersion returns a boolean if a field has been set.

### GetIndexedApplicationCount

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexedApplicationCount() int64`

GetIndexedApplicationCount returns the IndexedApplicationCount field if non-nil, zero value otherwise.

### GetIndexedApplicationCountOk

`func (o *ApiSearchIndexAnalyzeDTO) GetIndexedApplicationCountOk() (*int64, bool)`

GetIndexedApplicationCountOk returns a tuple with the IndexedApplicationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexedApplicationCount

`func (o *ApiSearchIndexAnalyzeDTO) SetIndexedApplicationCount(v int64)`

SetIndexedApplicationCount sets IndexedApplicationCount field to given value.

### HasIndexedApplicationCount

`func (o *ApiSearchIndexAnalyzeDTO) HasIndexedApplicationCount() bool`

HasIndexedApplicationCount returns a boolean if a field has been set.

### GetLastAnalysisCompletedAt

`func (o *ApiSearchIndexAnalyzeDTO) GetLastAnalysisCompletedAt() time.Time`

GetLastAnalysisCompletedAt returns the LastAnalysisCompletedAt field if non-nil, zero value otherwise.

### GetLastAnalysisCompletedAtOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastAnalysisCompletedAtOk() (*time.Time, bool)`

GetLastAnalysisCompletedAtOk returns a tuple with the LastAnalysisCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAnalysisCompletedAt

`func (o *ApiSearchIndexAnalyzeDTO) SetLastAnalysisCompletedAt(v time.Time)`

SetLastAnalysisCompletedAt sets LastAnalysisCompletedAt field to given value.

### HasLastAnalysisCompletedAt

`func (o *ApiSearchIndexAnalyzeDTO) HasLastAnalysisCompletedAt() bool`

HasLastAnalysisCompletedAt returns a boolean if a field has been set.

### GetLastIncrementalDocumentsAdded

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIncrementalDocumentsAdded() int64`

GetLastIncrementalDocumentsAdded returns the LastIncrementalDocumentsAdded field if non-nil, zero value otherwise.

### GetLastIncrementalDocumentsAddedOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIncrementalDocumentsAddedOk() (*int64, bool)`

GetLastIncrementalDocumentsAddedOk returns a tuple with the LastIncrementalDocumentsAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIncrementalDocumentsAdded

`func (o *ApiSearchIndexAnalyzeDTO) SetLastIncrementalDocumentsAdded(v int64)`

SetLastIncrementalDocumentsAdded sets LastIncrementalDocumentsAdded field to given value.

### HasLastIncrementalDocumentsAdded

`func (o *ApiSearchIndexAnalyzeDTO) HasLastIncrementalDocumentsAdded() bool`

HasLastIncrementalDocumentsAdded returns a boolean if a field has been set.

### GetLastIncrementalDocumentsDeleted

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIncrementalDocumentsDeleted() int64`

GetLastIncrementalDocumentsDeleted returns the LastIncrementalDocumentsDeleted field if non-nil, zero value otherwise.

### GetLastIncrementalDocumentsDeletedOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIncrementalDocumentsDeletedOk() (*int64, bool)`

GetLastIncrementalDocumentsDeletedOk returns a tuple with the LastIncrementalDocumentsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIncrementalDocumentsDeleted

`func (o *ApiSearchIndexAnalyzeDTO) SetLastIncrementalDocumentsDeleted(v int64)`

SetLastIncrementalDocumentsDeleted sets LastIncrementalDocumentsDeleted field to given value.

### HasLastIncrementalDocumentsDeleted

`func (o *ApiSearchIndexAnalyzeDTO) HasLastIncrementalDocumentsDeleted() bool`

HasLastIncrementalDocumentsDeleted returns a boolean if a field has been set.

### GetLastIncrementalFinishedAt

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIncrementalFinishedAt() time.Time`

GetLastIncrementalFinishedAt returns the LastIncrementalFinishedAt field if non-nil, zero value otherwise.

### GetLastIncrementalFinishedAtOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIncrementalFinishedAtOk() (*time.Time, bool)`

GetLastIncrementalFinishedAtOk returns a tuple with the LastIncrementalFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIncrementalFinishedAt

`func (o *ApiSearchIndexAnalyzeDTO) SetLastIncrementalFinishedAt(v time.Time)`

SetLastIncrementalFinishedAt sets LastIncrementalFinishedAt field to given value.

### HasLastIncrementalFinishedAt

`func (o *ApiSearchIndexAnalyzeDTO) HasLastIncrementalFinishedAt() bool`

HasLastIncrementalFinishedAt returns a boolean if a field has been set.

### GetLastIndexCompletedAt

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIndexCompletedAt() time.Time`

GetLastIndexCompletedAt returns the LastIndexCompletedAt field if non-nil, zero value otherwise.

### GetLastIndexCompletedAtOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIndexCompletedAtOk() (*time.Time, bool)`

GetLastIndexCompletedAtOk returns a tuple with the LastIndexCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexCompletedAt

`func (o *ApiSearchIndexAnalyzeDTO) SetLastIndexCompletedAt(v time.Time)`

SetLastIndexCompletedAt sets LastIndexCompletedAt field to given value.

### HasLastIndexCompletedAt

`func (o *ApiSearchIndexAnalyzeDTO) HasLastIndexCompletedAt() bool`

HasLastIndexCompletedAt returns a boolean if a field has been set.

### GetLastIndexId

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIndexId() string`

GetLastIndexId returns the LastIndexId field if non-nil, zero value otherwise.

### GetLastIndexIdOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIndexIdOk() (*string, bool)`

GetLastIndexIdOk returns a tuple with the LastIndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexId

`func (o *ApiSearchIndexAnalyzeDTO) SetLastIndexId(v string)`

SetLastIndexId sets LastIndexId field to given value.

### HasLastIndexId

`func (o *ApiSearchIndexAnalyzeDTO) HasLastIndexId() bool`

HasLastIndexId returns a boolean if a field has been set.

### GetLastIndexJobId

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIndexJobId() string`

GetLastIndexJobId returns the LastIndexJobId field if non-nil, zero value otherwise.

### GetLastIndexJobIdOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIndexJobIdOk() (*string, bool)`

GetLastIndexJobIdOk returns a tuple with the LastIndexJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexJobId

`func (o *ApiSearchIndexAnalyzeDTO) SetLastIndexJobId(v string)`

SetLastIndexJobId sets LastIndexJobId field to given value.

### HasLastIndexJobId

`func (o *ApiSearchIndexAnalyzeDTO) HasLastIndexJobId() bool`

HasLastIndexJobId returns a boolean if a field has been set.

### GetLastIndexJobStatus

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIndexJobStatus() string`

GetLastIndexJobStatus returns the LastIndexJobStatus field if non-nil, zero value otherwise.

### GetLastIndexJobStatusOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIndexJobStatusOk() (*string, bool)`

GetLastIndexJobStatusOk returns a tuple with the LastIndexJobStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexJobStatus

`func (o *ApiSearchIndexAnalyzeDTO) SetLastIndexJobStatus(v string)`

SetLastIndexJobStatus sets LastIndexJobStatus field to given value.

### HasLastIndexJobStatus

`func (o *ApiSearchIndexAnalyzeDTO) HasLastIndexJobStatus() bool`

HasLastIndexJobStatus returns a boolean if a field has been set.

### GetLastIndexJobType

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIndexJobType() string`

GetLastIndexJobType returns the LastIndexJobType field if non-nil, zero value otherwise.

### GetLastIndexJobTypeOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIndexJobTypeOk() (*string, bool)`

GetLastIndexJobTypeOk returns a tuple with the LastIndexJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexJobType

`func (o *ApiSearchIndexAnalyzeDTO) SetLastIndexJobType(v string)`

SetLastIndexJobType sets LastIndexJobType field to given value.

### HasLastIndexJobType

`func (o *ApiSearchIndexAnalyzeDTO) HasLastIndexJobType() bool`

HasLastIndexJobType returns a boolean if a field has been set.

### GetLastIndexStartedAt

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIndexStartedAt() time.Time`

GetLastIndexStartedAt returns the LastIndexStartedAt field if non-nil, zero value otherwise.

### GetLastIndexStartedAtOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastIndexStartedAtOk() (*time.Time, bool)`

GetLastIndexStartedAtOk returns a tuple with the LastIndexStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastIndexStartedAt

`func (o *ApiSearchIndexAnalyzeDTO) SetLastIndexStartedAt(v time.Time)`

SetLastIndexStartedAt sets LastIndexStartedAt field to given value.

### HasLastIndexStartedAt

`func (o *ApiSearchIndexAnalyzeDTO) HasLastIndexStartedAt() bool`

HasLastIndexStartedAt returns a boolean if a field has been set.

### GetLastRunState

`func (o *ApiSearchIndexAnalyzeDTO) GetLastRunState() string`

GetLastRunState returns the LastRunState field if non-nil, zero value otherwise.

### GetLastRunStateOk

`func (o *ApiSearchIndexAnalyzeDTO) GetLastRunStateOk() (*string, bool)`

GetLastRunStateOk returns a tuple with the LastRunState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunState

`func (o *ApiSearchIndexAnalyzeDTO) SetLastRunState(v string)`

SetLastRunState sets LastRunState field to given value.

### HasLastRunState

`func (o *ApiSearchIndexAnalyzeDTO) HasLastRunState() bool`

HasLastRunState returns a boolean if a field has been set.

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

### GetOwnerApplicationCount

`func (o *ApiSearchIndexAnalyzeDTO) GetOwnerApplicationCount() int64`

GetOwnerApplicationCount returns the OwnerApplicationCount field if non-nil, zero value otherwise.

### GetOwnerApplicationCountOk

`func (o *ApiSearchIndexAnalyzeDTO) GetOwnerApplicationCountOk() (*int64, bool)`

GetOwnerApplicationCountOk returns a tuple with the OwnerApplicationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerApplicationCount

`func (o *ApiSearchIndexAnalyzeDTO) SetOwnerApplicationCount(v int64)`

SetOwnerApplicationCount sets OwnerApplicationCount field to given value.

### HasOwnerApplicationCount

`func (o *ApiSearchIndexAnalyzeDTO) HasOwnerApplicationCount() bool`

HasOwnerApplicationCount returns a boolean if a field has been set.

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

### GetPendingChangeCountByType

`func (o *ApiSearchIndexAnalyzeDTO) GetPendingChangeCountByType() map[string]int64`

GetPendingChangeCountByType returns the PendingChangeCountByType field if non-nil, zero value otherwise.

### GetPendingChangeCountByTypeOk

`func (o *ApiSearchIndexAnalyzeDTO) GetPendingChangeCountByTypeOk() (*map[string]int64, bool)`

GetPendingChangeCountByTypeOk returns a tuple with the PendingChangeCountByType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingChangeCountByType

`func (o *ApiSearchIndexAnalyzeDTO) SetPendingChangeCountByType(v map[string]int64)`

SetPendingChangeCountByType sets PendingChangeCountByType field to given value.

### HasPendingChangeCountByType

`func (o *ApiSearchIndexAnalyzeDTO) HasPendingChangeCountByType() bool`

HasPendingChangeCountByType returns a boolean if a field has been set.

### GetPendingChangeState

`func (o *ApiSearchIndexAnalyzeDTO) GetPendingChangeState() string`

GetPendingChangeState returns the PendingChangeState field if non-nil, zero value otherwise.

### GetPendingChangeStateOk

`func (o *ApiSearchIndexAnalyzeDTO) GetPendingChangeStateOk() (*string, bool)`

GetPendingChangeStateOk returns a tuple with the PendingChangeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingChangeState

`func (o *ApiSearchIndexAnalyzeDTO) SetPendingChangeState(v string)`

SetPendingChangeState sets PendingChangeState field to given value.

### HasPendingChangeState

`func (o *ApiSearchIndexAnalyzeDTO) HasPendingChangeState() bool`

HasPendingChangeState returns a boolean if a field has been set.

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

### GetQueueLagState

`func (o *ApiSearchIndexAnalyzeDTO) GetQueueLagState() string`

GetQueueLagState returns the QueueLagState field if non-nil, zero value otherwise.

### GetQueueLagStateOk

`func (o *ApiSearchIndexAnalyzeDTO) GetQueueLagStateOk() (*string, bool)`

GetQueueLagStateOk returns a tuple with the QueueLagState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueLagState

`func (o *ApiSearchIndexAnalyzeDTO) SetQueueLagState(v string)`

SetQueueLagState sets QueueLagState field to given value.

### HasQueueLagState

`func (o *ApiSearchIndexAnalyzeDTO) HasQueueLagState() bool`

HasQueueLagState returns a boolean if a field has been set.

### GetRebuildRecommended

`func (o *ApiSearchIndexAnalyzeDTO) GetRebuildRecommended() bool`

GetRebuildRecommended returns the RebuildRecommended field if non-nil, zero value otherwise.

### GetRebuildRecommendedOk

`func (o *ApiSearchIndexAnalyzeDTO) GetRebuildRecommendedOk() (*bool, bool)`

GetRebuildRecommendedOk returns a tuple with the RebuildRecommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebuildRecommended

`func (o *ApiSearchIndexAnalyzeDTO) SetRebuildRecommended(v bool)`

SetRebuildRecommended sets RebuildRecommended field to given value.

### HasRebuildRecommended

`func (o *ApiSearchIndexAnalyzeDTO) HasRebuildRecommended() bool`

HasRebuildRecommended returns a boolean if a field has been set.

### GetRebuildRequired

`func (o *ApiSearchIndexAnalyzeDTO) GetRebuildRequired() bool`

GetRebuildRequired returns the RebuildRequired field if non-nil, zero value otherwise.

### GetRebuildRequiredOk

`func (o *ApiSearchIndexAnalyzeDTO) GetRebuildRequiredOk() (*bool, bool)`

GetRebuildRequiredOk returns a tuple with the RebuildRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebuildRequired

`func (o *ApiSearchIndexAnalyzeDTO) SetRebuildRequired(v bool)`

SetRebuildRequired sets RebuildRequired field to given value.

### HasRebuildRequired

`func (o *ApiSearchIndexAnalyzeDTO) HasRebuildRequired() bool`

HasRebuildRequired returns a boolean if a field has been set.

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

### GetSchemaUpgradeRequired

`func (o *ApiSearchIndexAnalyzeDTO) GetSchemaUpgradeRequired() bool`

GetSchemaUpgradeRequired returns the SchemaUpgradeRequired field if non-nil, zero value otherwise.

### GetSchemaUpgradeRequiredOk

`func (o *ApiSearchIndexAnalyzeDTO) GetSchemaUpgradeRequiredOk() (*bool, bool)`

GetSchemaUpgradeRequiredOk returns a tuple with the SchemaUpgradeRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaUpgradeRequired

`func (o *ApiSearchIndexAnalyzeDTO) SetSchemaUpgradeRequired(v bool)`

SetSchemaUpgradeRequired sets SchemaUpgradeRequired field to given value.

### HasSchemaUpgradeRequired

`func (o *ApiSearchIndexAnalyzeDTO) HasSchemaUpgradeRequired() bool`

HasSchemaUpgradeRequired returns a boolean if a field has been set.

### GetSearchableDocumentCount

`func (o *ApiSearchIndexAnalyzeDTO) GetSearchableDocumentCount() int64`

GetSearchableDocumentCount returns the SearchableDocumentCount field if non-nil, zero value otherwise.

### GetSearchableDocumentCountOk

`func (o *ApiSearchIndexAnalyzeDTO) GetSearchableDocumentCountOk() (*int64, bool)`

GetSearchableDocumentCountOk returns a tuple with the SearchableDocumentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchableDocumentCount

`func (o *ApiSearchIndexAnalyzeDTO) SetSearchableDocumentCount(v int64)`

SetSearchableDocumentCount sets SearchableDocumentCount field to given value.

### HasSearchableDocumentCount

`func (o *ApiSearchIndexAnalyzeDTO) HasSearchableDocumentCount() bool`

HasSearchableDocumentCount returns a boolean if a field has been set.

### GetServingDocCount

`func (o *ApiSearchIndexAnalyzeDTO) GetServingDocCount() int64`

GetServingDocCount returns the ServingDocCount field if non-nil, zero value otherwise.

### GetServingDocCountOk

`func (o *ApiSearchIndexAnalyzeDTO) GetServingDocCountOk() (*int64, bool)`

GetServingDocCountOk returns a tuple with the ServingDocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingDocCount

`func (o *ApiSearchIndexAnalyzeDTO) SetServingDocCount(v int64)`

SetServingDocCount sets ServingDocCount field to given value.

### HasServingDocCount

`func (o *ApiSearchIndexAnalyzeDTO) HasServingDocCount() bool`

HasServingDocCount returns a boolean if a field has been set.

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

### GetServingIndexId

`func (o *ApiSearchIndexAnalyzeDTO) GetServingIndexId() string`

GetServingIndexId returns the ServingIndexId field if non-nil, zero value otherwise.

### GetServingIndexIdOk

`func (o *ApiSearchIndexAnalyzeDTO) GetServingIndexIdOk() (*string, bool)`

GetServingIndexIdOk returns a tuple with the ServingIndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingIndexId

`func (o *ApiSearchIndexAnalyzeDTO) SetServingIndexId(v string)`

SetServingIndexId sets ServingIndexId field to given value.

### HasServingIndexId

`func (o *ApiSearchIndexAnalyzeDTO) HasServingIndexId() bool`

HasServingIndexId returns a boolean if a field has been set.

### GetServingSchemaVersion

`func (o *ApiSearchIndexAnalyzeDTO) GetServingSchemaVersion() int32`

GetServingSchemaVersion returns the ServingSchemaVersion field if non-nil, zero value otherwise.

### GetServingSchemaVersionOk

`func (o *ApiSearchIndexAnalyzeDTO) GetServingSchemaVersionOk() (*int32, bool)`

GetServingSchemaVersionOk returns a tuple with the ServingSchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingSchemaVersion

`func (o *ApiSearchIndexAnalyzeDTO) SetServingSchemaVersion(v int32)`

SetServingSchemaVersion sets ServingSchemaVersion field to given value.

### HasServingSchemaVersion

`func (o *ApiSearchIndexAnalyzeDTO) HasServingSchemaVersion() bool`

HasServingSchemaVersion returns a boolean if a field has been set.

### GetServingSince

`func (o *ApiSearchIndexAnalyzeDTO) GetServingSince() time.Time`

GetServingSince returns the ServingSince field if non-nil, zero value otherwise.

### GetServingSinceOk

`func (o *ApiSearchIndexAnalyzeDTO) GetServingSinceOk() (*time.Time, bool)`

GetServingSinceOk returns a tuple with the ServingSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingSince

`func (o *ApiSearchIndexAnalyzeDTO) SetServingSince(v time.Time)`

SetServingSince sets ServingSince field to given value.

### HasServingSince

`func (o *ApiSearchIndexAnalyzeDTO) HasServingSince() bool`

HasServingSince returns a boolean if a field has been set.

### GetServingSizeBytes

`func (o *ApiSearchIndexAnalyzeDTO) GetServingSizeBytes() int64`

GetServingSizeBytes returns the ServingSizeBytes field if non-nil, zero value otherwise.

### GetServingSizeBytesOk

`func (o *ApiSearchIndexAnalyzeDTO) GetServingSizeBytesOk() (*int64, bool)`

GetServingSizeBytesOk returns a tuple with the ServingSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingSizeBytes

`func (o *ApiSearchIndexAnalyzeDTO) SetServingSizeBytes(v int64)`

SetServingSizeBytes sets ServingSizeBytes field to given value.

### HasServingSizeBytes

`func (o *ApiSearchIndexAnalyzeDTO) HasServingSizeBytes() bool`

HasServingSizeBytes returns a boolean if a field has been set.

### GetStampedAtEpochMs

`func (o *ApiSearchIndexAnalyzeDTO) GetStampedAtEpochMs() int64`

GetStampedAtEpochMs returns the StampedAtEpochMs field if non-nil, zero value otherwise.

### GetStampedAtEpochMsOk

`func (o *ApiSearchIndexAnalyzeDTO) GetStampedAtEpochMsOk() (*int64, bool)`

GetStampedAtEpochMsOk returns a tuple with the StampedAtEpochMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampedAtEpochMs

`func (o *ApiSearchIndexAnalyzeDTO) SetStampedAtEpochMs(v int64)`

SetStampedAtEpochMs sets StampedAtEpochMs field to given value.

### HasStampedAtEpochMs

`func (o *ApiSearchIndexAnalyzeDTO) HasStampedAtEpochMs() bool`

HasStampedAtEpochMs returns a boolean if a field has been set.

### GetStampedIndexVersion

`func (o *ApiSearchIndexAnalyzeDTO) GetStampedIndexVersion() int32`

GetStampedIndexVersion returns the StampedIndexVersion field if non-nil, zero value otherwise.

### GetStampedIndexVersionOk

`func (o *ApiSearchIndexAnalyzeDTO) GetStampedIndexVersionOk() (*int32, bool)`

GetStampedIndexVersionOk returns a tuple with the StampedIndexVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampedIndexVersion

`func (o *ApiSearchIndexAnalyzeDTO) SetStampedIndexVersion(v int32)`

SetStampedIndexVersion sets StampedIndexVersion field to given value.

### HasStampedIndexVersion

`func (o *ApiSearchIndexAnalyzeDTO) HasStampedIndexVersion() bool`

HasStampedIndexVersion returns a boolean if a field has been set.

### GetStampedIqVersion

`func (o *ApiSearchIndexAnalyzeDTO) GetStampedIqVersion() string`

GetStampedIqVersion returns the StampedIqVersion field if non-nil, zero value otherwise.

### GetStampedIqVersionOk

`func (o *ApiSearchIndexAnalyzeDTO) GetStampedIqVersionOk() (*string, bool)`

GetStampedIqVersionOk returns a tuple with the StampedIqVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStampedIqVersion

`func (o *ApiSearchIndexAnalyzeDTO) SetStampedIqVersion(v string)`

SetStampedIqVersion sets StampedIqVersion field to given value.

### HasStampedIqVersion

`func (o *ApiSearchIndexAnalyzeDTO) HasStampedIqVersion() bool`

HasStampedIqVersion returns a boolean if a field has been set.

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


