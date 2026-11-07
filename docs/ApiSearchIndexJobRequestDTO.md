# ApiSearchIndexJobRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocTypes** | Pointer to **[]string** | Doc types to derive from the database, for a FULL_REBUILD. The rest are copied from the index serving now and keep the mappings and shard count they were built with. Absent, null or empty derives every doc type, and so does naming all of them. Naming any reference doc type derives all six, because one pass derives them together, and application is always derived. | [optional] 
**JobType** | Pointer to **string** | Job to run: FULL_REBUILD or FIRST_TIME_INDEX. | [optional] 
**Trigger** | Pointer to **string** | What asked for the job: UNLOCK_WIZARD, HEALTH_UI, SUPPORT or SYSTEM. Defaults to HEALTH_UI. | [optional] 

## Methods

### NewApiSearchIndexJobRequestDTO

`func NewApiSearchIndexJobRequestDTO() *ApiSearchIndexJobRequestDTO`

NewApiSearchIndexJobRequestDTO instantiates a new ApiSearchIndexJobRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchIndexJobRequestDTOWithDefaults

`func NewApiSearchIndexJobRequestDTOWithDefaults() *ApiSearchIndexJobRequestDTO`

NewApiSearchIndexJobRequestDTOWithDefaults instantiates a new ApiSearchIndexJobRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocTypes

`func (o *ApiSearchIndexJobRequestDTO) GetDocTypes() []string`

GetDocTypes returns the DocTypes field if non-nil, zero value otherwise.

### GetDocTypesOk

`func (o *ApiSearchIndexJobRequestDTO) GetDocTypesOk() (*[]string, bool)`

GetDocTypesOk returns a tuple with the DocTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocTypes

`func (o *ApiSearchIndexJobRequestDTO) SetDocTypes(v []string)`

SetDocTypes sets DocTypes field to given value.

### HasDocTypes

`func (o *ApiSearchIndexJobRequestDTO) HasDocTypes() bool`

HasDocTypes returns a boolean if a field has been set.

### GetJobType

`func (o *ApiSearchIndexJobRequestDTO) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *ApiSearchIndexJobRequestDTO) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *ApiSearchIndexJobRequestDTO) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *ApiSearchIndexJobRequestDTO) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetTrigger

`func (o *ApiSearchIndexJobRequestDTO) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *ApiSearchIndexJobRequestDTO) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *ApiSearchIndexJobRequestDTO) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *ApiSearchIndexJobRequestDTO) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


