# ApiSearchIndexGenerationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CreatedByJobId** | Pointer to **string** |  | [optional] 
**CreatedByTrigger** | Pointer to **string** |  | [optional] 
**CreatedByUserId** | Pointer to **string** |  | [optional] 
**DocCount** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IndexId** | Pointer to **string** |  | [optional] 
**JobFinishedAt** | Pointer to **time.Time** |  | [optional] 
**JobStartedAt** | Pointer to **time.Time** |  | [optional] 
**ProductVersion** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**SchemaVersion** | Pointer to **int32** |  | [optional] 
**ServingSince** | Pointer to **time.Time** |  | [optional] 
**SizeBytes** | Pointer to **int64** |  | [optional] 

## Methods

### NewApiSearchIndexGenerationDTO

`func NewApiSearchIndexGenerationDTO() *ApiSearchIndexGenerationDTO`

NewApiSearchIndexGenerationDTO instantiates a new ApiSearchIndexGenerationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchIndexGenerationDTOWithDefaults

`func NewApiSearchIndexGenerationDTOWithDefaults() *ApiSearchIndexGenerationDTO`

NewApiSearchIndexGenerationDTOWithDefaults instantiates a new ApiSearchIndexGenerationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ApiSearchIndexGenerationDTO) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiSearchIndexGenerationDTO) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiSearchIndexGenerationDTO) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiSearchIndexGenerationDTO) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreatedByJobId

`func (o *ApiSearchIndexGenerationDTO) GetCreatedByJobId() string`

GetCreatedByJobId returns the CreatedByJobId field if non-nil, zero value otherwise.

### GetCreatedByJobIdOk

`func (o *ApiSearchIndexGenerationDTO) GetCreatedByJobIdOk() (*string, bool)`

GetCreatedByJobIdOk returns a tuple with the CreatedByJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByJobId

`func (o *ApiSearchIndexGenerationDTO) SetCreatedByJobId(v string)`

SetCreatedByJobId sets CreatedByJobId field to given value.

### HasCreatedByJobId

`func (o *ApiSearchIndexGenerationDTO) HasCreatedByJobId() bool`

HasCreatedByJobId returns a boolean if a field has been set.

### GetCreatedByTrigger

`func (o *ApiSearchIndexGenerationDTO) GetCreatedByTrigger() string`

GetCreatedByTrigger returns the CreatedByTrigger field if non-nil, zero value otherwise.

### GetCreatedByTriggerOk

`func (o *ApiSearchIndexGenerationDTO) GetCreatedByTriggerOk() (*string, bool)`

GetCreatedByTriggerOk returns a tuple with the CreatedByTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByTrigger

`func (o *ApiSearchIndexGenerationDTO) SetCreatedByTrigger(v string)`

SetCreatedByTrigger sets CreatedByTrigger field to given value.

### HasCreatedByTrigger

`func (o *ApiSearchIndexGenerationDTO) HasCreatedByTrigger() bool`

HasCreatedByTrigger returns a boolean if a field has been set.

### GetCreatedByUserId

`func (o *ApiSearchIndexGenerationDTO) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *ApiSearchIndexGenerationDTO) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *ApiSearchIndexGenerationDTO) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.

### HasCreatedByUserId

`func (o *ApiSearchIndexGenerationDTO) HasCreatedByUserId() bool`

HasCreatedByUserId returns a boolean if a field has been set.

### GetDocCount

`func (o *ApiSearchIndexGenerationDTO) GetDocCount() int64`

GetDocCount returns the DocCount field if non-nil, zero value otherwise.

### GetDocCountOk

`func (o *ApiSearchIndexGenerationDTO) GetDocCountOk() (*int64, bool)`

GetDocCountOk returns a tuple with the DocCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocCount

`func (o *ApiSearchIndexGenerationDTO) SetDocCount(v int64)`

SetDocCount sets DocCount field to given value.

### HasDocCount

`func (o *ApiSearchIndexGenerationDTO) HasDocCount() bool`

HasDocCount returns a boolean if a field has been set.

### GetId

`func (o *ApiSearchIndexGenerationDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiSearchIndexGenerationDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiSearchIndexGenerationDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiSearchIndexGenerationDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndexId

`func (o *ApiSearchIndexGenerationDTO) GetIndexId() string`

GetIndexId returns the IndexId field if non-nil, zero value otherwise.

### GetIndexIdOk

`func (o *ApiSearchIndexGenerationDTO) GetIndexIdOk() (*string, bool)`

GetIndexIdOk returns a tuple with the IndexId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexId

`func (o *ApiSearchIndexGenerationDTO) SetIndexId(v string)`

SetIndexId sets IndexId field to given value.

### HasIndexId

`func (o *ApiSearchIndexGenerationDTO) HasIndexId() bool`

HasIndexId returns a boolean if a field has been set.

### GetJobFinishedAt

`func (o *ApiSearchIndexGenerationDTO) GetJobFinishedAt() time.Time`

GetJobFinishedAt returns the JobFinishedAt field if non-nil, zero value otherwise.

### GetJobFinishedAtOk

`func (o *ApiSearchIndexGenerationDTO) GetJobFinishedAtOk() (*time.Time, bool)`

GetJobFinishedAtOk returns a tuple with the JobFinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFinishedAt

`func (o *ApiSearchIndexGenerationDTO) SetJobFinishedAt(v time.Time)`

SetJobFinishedAt sets JobFinishedAt field to given value.

### HasJobFinishedAt

`func (o *ApiSearchIndexGenerationDTO) HasJobFinishedAt() bool`

HasJobFinishedAt returns a boolean if a field has been set.

### GetJobStartedAt

`func (o *ApiSearchIndexGenerationDTO) GetJobStartedAt() time.Time`

GetJobStartedAt returns the JobStartedAt field if non-nil, zero value otherwise.

### GetJobStartedAtOk

`func (o *ApiSearchIndexGenerationDTO) GetJobStartedAtOk() (*time.Time, bool)`

GetJobStartedAtOk returns a tuple with the JobStartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobStartedAt

`func (o *ApiSearchIndexGenerationDTO) SetJobStartedAt(v time.Time)`

SetJobStartedAt sets JobStartedAt field to given value.

### HasJobStartedAt

`func (o *ApiSearchIndexGenerationDTO) HasJobStartedAt() bool`

HasJobStartedAt returns a boolean if a field has been set.

### GetProductVersion

`func (o *ApiSearchIndexGenerationDTO) GetProductVersion() string`

GetProductVersion returns the ProductVersion field if non-nil, zero value otherwise.

### GetProductVersionOk

`func (o *ApiSearchIndexGenerationDTO) GetProductVersionOk() (*string, bool)`

GetProductVersionOk returns a tuple with the ProductVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVersion

`func (o *ApiSearchIndexGenerationDTO) SetProductVersion(v string)`

SetProductVersion sets ProductVersion field to given value.

### HasProductVersion

`func (o *ApiSearchIndexGenerationDTO) HasProductVersion() bool`

HasProductVersion returns a boolean if a field has been set.

### GetRole

`func (o *ApiSearchIndexGenerationDTO) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ApiSearchIndexGenerationDTO) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ApiSearchIndexGenerationDTO) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ApiSearchIndexGenerationDTO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetSchemaVersion

`func (o *ApiSearchIndexGenerationDTO) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ApiSearchIndexGenerationDTO) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ApiSearchIndexGenerationDTO) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *ApiSearchIndexGenerationDTO) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetServingSince

`func (o *ApiSearchIndexGenerationDTO) GetServingSince() time.Time`

GetServingSince returns the ServingSince field if non-nil, zero value otherwise.

### GetServingSinceOk

`func (o *ApiSearchIndexGenerationDTO) GetServingSinceOk() (*time.Time, bool)`

GetServingSinceOk returns a tuple with the ServingSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingSince

`func (o *ApiSearchIndexGenerationDTO) SetServingSince(v time.Time)`

SetServingSince sets ServingSince field to given value.

### HasServingSince

`func (o *ApiSearchIndexGenerationDTO) HasServingSince() bool`

HasServingSince returns a boolean if a field has been set.

### GetSizeBytes

`func (o *ApiSearchIndexGenerationDTO) GetSizeBytes() int64`

GetSizeBytes returns the SizeBytes field if non-nil, zero value otherwise.

### GetSizeBytesOk

`func (o *ApiSearchIndexGenerationDTO) GetSizeBytesOk() (*int64, bool)`

GetSizeBytesOk returns a tuple with the SizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeBytes

`func (o *ApiSearchIndexGenerationDTO) SetSizeBytes(v int64)`

SetSizeBytes sets SizeBytes field to given value.

### HasSizeBytes

`func (o *ApiSearchIndexGenerationDTO) HasSizeBytes() bool`

HasSizeBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


