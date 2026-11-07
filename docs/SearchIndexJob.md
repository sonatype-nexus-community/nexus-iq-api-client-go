# SearchIndexJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**ActiveSlot** | Pointer to **string** |  | [optional] 
**BuildingIndexName** | Pointer to **string** |  | [optional] 
**CancelRequestedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedByUserId** | Pointer to **string** |  | [optional] 
**DocCount** | Pointer to **int64** |  | [optional] 
**DocTypeScope** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**EtaFinishAt** | Pointer to **time.Time** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IqVersion** | Pointer to **string** |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**ProgressPercent** | Pointer to **int32** |  | [optional] 
**RecommendedOp** | Pointer to **string** |  | [optional] 
**SchemaVersion** | Pointer to **int32** |  | [optional] 
**SizeBytes** | Pointer to **int64** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Trigger** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSearchIndexJob

`func NewSearchIndexJob() *SearchIndexJob`

NewSearchIndexJob instantiates a new SearchIndexJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexJobWithDefaults

`func NewSearchIndexJobWithDefaults() *SearchIndexJob`

NewSearchIndexJobWithDefaults instantiates a new SearchIndexJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *SearchIndexJob) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *SearchIndexJob) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *SearchIndexJob) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *SearchIndexJob) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetActiveSlot

`func (o *SearchIndexJob) GetActiveSlot() string`

GetActiveSlot returns the ActiveSlot field if non-nil, zero value otherwise.

### GetActiveSlotOk

`func (o *SearchIndexJob) GetActiveSlotOk() (*string, bool)`

GetActiveSlotOk returns a tuple with the ActiveSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSlot

`func (o *SearchIndexJob) SetActiveSlot(v string)`

SetActiveSlot sets ActiveSlot field to given value.

### HasActiveSlot

`func (o *SearchIndexJob) HasActiveSlot() bool`

HasActiveSlot returns a boolean if a field has been set.

### GetBuildingIndexName

`func (o *SearchIndexJob) GetBuildingIndexName() string`

GetBuildingIndexName returns the BuildingIndexName field if non-nil, zero value otherwise.

### GetBuildingIndexNameOk

`func (o *SearchIndexJob) GetBuildingIndexNameOk() (*string, bool)`

GetBuildingIndexNameOk returns a tuple with the BuildingIndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingIndexName

`func (o *SearchIndexJob) SetBuildingIndexName(v string)`

SetBuildingIndexName sets BuildingIndexName field to given value.

### HasBuildingIndexName

`func (o *SearchIndexJob) HasBuildingIndexName() bool`

HasBuildingIndexName returns a boolean if a field has been set.

### GetCancelRequestedAt

`func (o *SearchIndexJob) GetCancelRequestedAt() time.Time`

GetCancelRequestedAt returns the CancelRequestedAt field if non-nil, zero value otherwise.

### GetCancelRequestedAtOk

`func (o *SearchIndexJob) GetCancelRequestedAtOk() (*time.Time, bool)`

GetCancelRequestedAtOk returns a tuple with the CancelRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelRequestedAt

`func (o *SearchIndexJob) SetCancelRequestedAt(v time.Time)`

SetCancelRequestedAt sets CancelRequestedAt field to given value.

### HasCancelRequestedAt

`func (o *SearchIndexJob) HasCancelRequestedAt() bool`

HasCancelRequestedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SearchIndexJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SearchIndexJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SearchIndexJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SearchIndexJob) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *SearchIndexJob) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *SearchIndexJob) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *SearchIndexJob) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *SearchIndexJob) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDocCount

`func (o *SearchIndexJob) GetDocCount() int64`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *SearchIndexJob) GetDocCountOk() (*int64, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *SearchIndexJob) SetDocCount(v int64)`

SetDocCount sets DocCount field to given value.

### HasDocCount

`func (o *SearchIndexJob) HasDocCount() bool`

HasDocCount returns a boolean if a field has been set.

### GetDocTypeScope

`func (o *SearchIndexJob) GetDocTypeScope() string`

GetDocTypeScope returns the DocTypeScope field if non-nil, zero value otherwise.

### GetDocTypeScopeOk

`func (o *SearchIndexJob) GetDocTypeScopeOk() (*string, bool)`

GetDocTypeScopeOk returns a tuple with the DocTypeScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocTypeScope

`func (o *SearchIndexJob) SetDocTypeScope(v string)`

SetDocTypeScope sets DocTypeScope field to given value.

### HasDocTypeScope

`func (o *SearchIndexJob) HasDocTypeScope() bool`

HasDocTypeScope returns a boolean if a field has been set.

### GetErrorCode

`func (o *SearchIndexJob) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SearchIndexJob) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SearchIndexJob) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *SearchIndexJob) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *SearchIndexJob) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *SearchIndexJob) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *SearchIndexJob) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *SearchIndexJob) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetEtaFinishAt

`func (o *SearchIndexJob) GetEtaFinishAt() time.Time`

GetEtaFinishAt returns the EtaFinishAt field if non-nil, zero value otherwise.

### GetEtaFinishAtOk

`func (o *SearchIndexJob) GetEtaFinishAtOk() (*time.Time, bool)`

GetEtaFinishAtOk returns a tuple with the EtaFinishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtaFinishAt

`func (o *SearchIndexJob) SetEtaFinishAt(v time.Time)`

SetEtaFinishAt sets EtaFinishAt field to given value.

### HasEtaFinishAt

`func (o *SearchIndexJob) HasEtaFinishAt() bool`

HasEtaFinishAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *SearchIndexJob) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *SearchIndexJob) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *SearchIndexJob) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *SearchIndexJob) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetId

`func (o *SearchIndexJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchIndexJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchIndexJob) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchIndexJob) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIqVersion

`func (o *SearchIndexJob) GetIqVersion() string`

GetIqVersion returns the IqVersion field if non-nil, zero value otherwise.

### GetIqVersionOk

`func (o *SearchIndexJob) GetIqVersionOk() (*string, bool)`

GetIqVersionOk returns a tuple with the IqVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIqVersion

`func (o *SearchIndexJob) SetIqVersion(v string)`

SetIqVersion sets IqVersion field to given value.

### HasIqVersion

`func (o *SearchIndexJob) HasIqVersion() bool`

HasIqVersion returns a boolean if a field has been set.

### GetJobType

`func (o *SearchIndexJob) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *SearchIndexJob) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *SearchIndexJob) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *SearchIndexJob) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetPhase

`func (o *SearchIndexJob) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *SearchIndexJob) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *SearchIndexJob) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *SearchIndexJob) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetProgressPercent

`func (o *SearchIndexJob) GetProgressPercent() int32`

GetProgressPercent returns the ProgressPercent field if non-nil, zero value otherwise.

### GetProgressPercentOk

`func (o *SearchIndexJob) GetProgressPercentOk() (*int32, bool)`

GetProgressPercentOk returns a tuple with the ProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercent

`func (o *SearchIndexJob) SetProgressPercent(v int32)`

SetProgressPercent sets ProgressPercent field to given value.

### HasProgressPercent

`func (o *SearchIndexJob) HasProgressPercent() bool`

HasProgressPercent returns a boolean if a field has been set.

### GetRecommendedOp

`func (o *SearchIndexJob) GetRecommendedOp() string`

GetRecommendedOp returns the RecommendedOp field if non-nil, zero value otherwise.

### GetRecommendedOpOk

`func (o *SearchIndexJob) GetRecommendedOpOk() (*string, bool)`

GetRecommendedOpOk returns a tuple with the RecommendedOp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedOp

`func (o *SearchIndexJob) SetRecommendedOp(v string)`

SetRecommendedOp sets RecommendedOp field to given value.

### HasRecommendedOp

`func (o *SearchIndexJob) HasRecommendedOp() bool`

HasRecommendedOp returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *SearchIndexJob) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *SearchIndexJob) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *SearchIndexJob) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *SearchIndexJob) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSizeBytes

`func (o *SearchIndexJob) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *SearchIndexJob) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *SearchIndexJob) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *SearchIndexJob) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetStartedAt

`func (o *SearchIndexJob) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SearchIndexJob) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SearchIndexJob) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *SearchIndexJob) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *SearchIndexJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SearchIndexJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SearchIndexJob) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SearchIndexJob) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrigger

`func (o *SearchIndexJob) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *SearchIndexJob) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *SearchIndexJob) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *SearchIndexJob) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SearchIndexJob) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SearchIndexJob) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SearchIndexJob) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SearchIndexJob) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


