# ApiBulkWaiversDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiWaiverOptionsDTO** | [**ApiWaiverOptionsDTO**](ApiWaiverOptionsDTO.md) |  | 
**ViolationIds** | **[]string** | List of repository policy violation IDs to waive. Maximum 1000 violations per request. Duplicate IDs are automatically deduplicated. Supports both quarantine (FAIL) and non-quarantine (WARN) violations. Already-waived violations are skipped without error. | 

## Methods

### NewApiBulkWaiversDTO

`func NewApiBulkWaiversDTO(apiWaiverOptionsDTO ApiWaiverOptionsDTO, violationIds []string, ) *ApiBulkWaiversDTO`

NewApiBulkWaiversDTO instantiates a new ApiBulkWaiversDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBulkWaiversDTOWithDefaults

`func NewApiBulkWaiversDTOWithDefaults() *ApiBulkWaiversDTO`

NewApiBulkWaiversDTOWithDefaults instantiates a new ApiBulkWaiversDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiWaiverOptionsDTO

`func (o *ApiBulkWaiversDTO) GetApiWaiverOptionsDTO() ApiWaiverOptionsDTO`

GetApiWaiverOptionsDTO returns the ApiWaiverOptionsDTO field if non-nil, zero value otherwise.

### GetApiWaiverOptionsDTOOk

`func (o *ApiBulkWaiversDTO) GetApiWaiverOptionsDTOOk() (*ApiWaiverOptionsDTO, bool)`

GetApiWaiverOptionsDTOOk returns a tuple with the ApiWaiverOptionsDTO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiWaiverOptionsDTO

`func (o *ApiBulkWaiversDTO) SetApiWaiverOptionsDTO(v ApiWaiverOptionsDTO)`

SetApiWaiverOptionsDTO sets ApiWaiverOptionsDTO field to given value.


### GetViolationIds

`func (o *ApiBulkWaiversDTO) GetViolationIds() []string`

GetViolationIds returns the ViolationIds field if non-nil, zero value otherwise.

### GetViolationIdsOk

`func (o *ApiBulkWaiversDTO) GetViolationIdsOk() (*[]string, bool)`

GetViolationIdsOk returns a tuple with the ViolationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolationIds

`func (o *ApiBulkWaiversDTO) SetViolationIds(v []string)`

SetViolationIds sets ViolationIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


