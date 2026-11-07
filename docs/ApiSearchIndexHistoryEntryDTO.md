# ApiSearchIndexHistoryEntryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildingGenerationId** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **string** |  | [optional] 
**DocCount** | Pointer to **int64** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**IndexId** | Pointer to **string** |  | [optional] 
**IndexLocation** | Pointer to **string** |  | [optional] 
**IndexType** | Pointer to **string** |  | [optional] 
**JobId** | Pointer to **string** |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**Phase** | Pointer to **string** |  | [optional] 
**ProductVersion** | Pointer to **string** |  | [optional] 
**ProgressPercent** | Pointer to **int32** |  | [optional] 
**SchemaVersion** | Pointer to **int32** |  | [optional] 
**SizeBytes** | Pointer to **int64** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Trigger** | Pointer to **string** |  | [optional] 

## Methods

### NewApiSearchIndexHistoryEntryDTO

`func NewApiSearchIndexHistoryEntryDTO() *ApiSearchIndexHistoryEntryDTO`

NewApiSearchIndexHistoryEntryDTO instantiates a new ApiSearchIndexHistoryEntryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchIndexHistoryEntryDTOWithDefaults

`func NewApiSearchIndexHistoryEntryDTOWithDefaults() *ApiSearchIndexHistoryEntryDTO`

NewApiSearchIndexHistoryEntryDTOWithDefaults instantiates a new ApiSearchIndexHistoryEntryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildingGenerationId

`func (o *ApiSearchIndexHistoryEntryDTO) GetBuildingGenerationId() string`

GetBuildingGenerationId returns the BuildingGenerationId field if non-nil, zero value otherwise.

### GetBuildingGenerationIdOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetBuildingGenerationIdOk() (*string, bool)`

GetBuildingGenerationIdOk returns a tuple with the BuildingGenerationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildingGenerationId

`func (o *ApiSearchIndexHistoryEntryDTO) SetBuildingGenerationId(v string)`

SetBuildingGenerationId sets BuildingGenerationId field to given value.

### HasBuildingGenerationId

`func (o *ApiSearchIndexHistoryEntryDTO) HasBuildingGenerationId() bool`

HasBuildingGenerationId returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ApiSearchIndexHistoryEntryDTO) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ApiSearchIndexHistoryEntryDTO) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ApiSearchIndexHistoryEntryDTO) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDocCount

`func (o *ApiSearchIndexHistoryEntryDTO) GetDocCount() int64`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetDocCountOk() (*int64, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *ApiSearchIndexHistoryEntryDTO) SetDocCount(v int64)`

SetDocCount sets DocCount field to given value.

### HasDocCount

`func (o *ApiSearchIndexHistoryEntryDTO) HasDocCount() bool`

HasDocCount returns a boolean if a field has been set.

### GetErrorCode

`func (o *ApiSearchIndexHistoryEntryDTO) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ApiSearchIndexHistoryEntryDTO) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ApiSearchIndexHistoryEntryDTO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ApiSearchIndexHistoryEntryDTO) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ApiSearchIndexHistoryEntryDTO) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ApiSearchIndexHistoryEntryDTO) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetFinishedAt

`func (o *ApiSearchIndexHistoryEntryDTO) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ApiSearchIndexHistoryEntryDTO) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ApiSearchIndexHistoryEntryDTO) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetIndexId

`func (o *ApiSearchIndexHistoryEntryDTO) GetIndexId() string`

GetIndexId returns the IndexId field if non-nil, zero value otherwise.

### GetIndexIdOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetIndexIdOk() (*string, bool)`

GetIndexIdOk returns a tuple with the IndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexId

`func (o *ApiSearchIndexHistoryEntryDTO) SetIndexId(v string)`

SetIndexId sets IndexId field to given value.

### HasIndexId

`func (o *ApiSearchIndexHistoryEntryDTO) HasIndexId() bool`

HasIndexId returns a boolean if a field has been set.

### GetIndexLocation

`func (o *ApiSearchIndexHistoryEntryDTO) GetIndexLocation() string`

GetIndexLocation returns the IndexLocation field if non-nil, zero value otherwise.

### GetIndexLocationOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetIndexLocationOk() (*string, bool)`

GetIndexLocationOk returns a tuple with the IndexLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexLocation

`func (o *ApiSearchIndexHistoryEntryDTO) SetIndexLocation(v string)`

SetIndexLocation sets IndexLocation field to given value.

### HasIndexLocation

`func (o *ApiSearchIndexHistoryEntryDTO) HasIndexLocation() bool`

HasIndexLocation returns a boolean if a field has been set.

### GetIndexType

`func (o *ApiSearchIndexHistoryEntryDTO) GetIndexType() string`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetIndexTypeOk() (*string, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *ApiSearchIndexHistoryEntryDTO) SetIndexType(v string)`

SetIndexType sets IndexType field to given value.

### HasIndexType

`func (o *ApiSearchIndexHistoryEntryDTO) HasIndexType() bool`

HasIndexType returns a boolean if a field has been set.

### GetJobId

`func (o *ApiSearchIndexHistoryEntryDTO) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *ApiSearchIndexHistoryEntryDTO) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *ApiSearchIndexHistoryEntryDTO) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetJobType

`func (o *ApiSearchIndexHistoryEntryDTO) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ApiSearchIndexHistoryEntryDTO) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *ApiSearchIndexHistoryEntryDTO) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetPhase

`func (o *ApiSearchIndexHistoryEntryDTO) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *ApiSearchIndexHistoryEntryDTO) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *ApiSearchIndexHistoryEntryDTO) HasPhase() bool`

HasPhase returns a boolean if a field has been set.

### GetProductVersion

`func (o *ApiSearchIndexHistoryEntryDTO) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ApiSearchIndexHistoryEntryDTO) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ApiSearchIndexHistoryEntryDTO) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetProgressPercent

`func (o *ApiSearchIndexHistoryEntryDTO) GetProgressPercent() int32`

GetProgressPercent returns the ProgressPercent field if non-nil, zero value otherwise.

### GetProgressPercentOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetProgressPercentOk() (*int32, bool)`

GetProgressPercentOk returns a tuple with the ProgressPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressPercent

`func (o *ApiSearchIndexHistoryEntryDTO) SetProgressPercent(v int32)`

SetProgressPercent sets ProgressPercent field to given value.

### HasProgressPercent

`func (o *ApiSearchIndexHistoryEntryDTO) HasProgressPercent() bool`

HasProgressPercent returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ApiSearchIndexHistoryEntryDTO) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ApiSearchIndexHistoryEntryDTO) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ApiSearchIndexHistoryEntryDTO) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSizeBytes

`func (o *ApiSearchIndexHistoryEntryDTO) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *ApiSearchIndexHistoryEntryDTO) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *ApiSearchIndexHistoryEntryDTO) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.

### GetStartedAt

`func (o *ApiSearchIndexHistoryEntryDTO) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ApiSearchIndexHistoryEntryDTO) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ApiSearchIndexHistoryEntryDTO) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *ApiSearchIndexHistoryEntryDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiSearchIndexHistoryEntryDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiSearchIndexHistoryEntryDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrigger

`func (o *ApiSearchIndexHistoryEntryDTO) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *ApiSearchIndexHistoryEntryDTO) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *ApiSearchIndexHistoryEntryDTO) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *ApiSearchIndexHistoryEntryDTO) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


