# ApiSearchIndexJobRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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


