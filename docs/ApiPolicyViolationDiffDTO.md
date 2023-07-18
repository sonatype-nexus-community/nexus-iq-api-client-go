# ApiPolicyViolationDiffDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedViolations** | Pointer to [**[]ApiPolicyViolationForDiffDTO**](ApiPolicyViolationForDiffDTO.md) |  | [optional] 
**Application** | Pointer to [**ApiApplicationDTO**](ApiApplicationDTO.md) |  | [optional] 
**DiffTime** | Pointer to **time.Time** |  | [optional] 
**FromCommit** | Pointer to [**ApiApplicationEvaluationCommitDTO**](ApiApplicationEvaluationCommitDTO.md) |  | [optional] 
**RemovedViolations** | Pointer to [**[]ApiPolicyViolationForDiffDTO**](ApiPolicyViolationForDiffDTO.md) |  | [optional] 
**SameViolations** | Pointer to [**[]ApiPolicyViolationForDiffDTO**](ApiPolicyViolationForDiffDTO.md) |  | [optional] 
**ToCommit** | Pointer to [**ApiApplicationEvaluationCommitDTO**](ApiApplicationEvaluationCommitDTO.md) |  | [optional] 

## Methods

### NewApiPolicyViolationDiffDTO

`func NewApiPolicyViolationDiffDTO() *ApiPolicyViolationDiffDTO`

NewApiPolicyViolationDiffDTO instantiates a new ApiPolicyViolationDiffDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyViolationDiffDTOWithDefaults

`func NewApiPolicyViolationDiffDTOWithDefaults() *ApiPolicyViolationDiffDTO`

NewApiPolicyViolationDiffDTOWithDefaults instantiates a new ApiPolicyViolationDiffDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedViolations

`func (o *ApiPolicyViolationDiffDTO) GetAddedViolations() []ApiPolicyViolationForDiffDTO`

GetAddedViolations returns the AddedViolations field if non-nil, zero value otherwise.

### GetAddedViolationsOk

`func (o *ApiPolicyViolationDiffDTO) GetAddedViolationsOk() (*[]ApiPolicyViolationForDiffDTO, bool)`

GetAddedViolationsOk returns a tuple with the AddedViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedViolations

`func (o *ApiPolicyViolationDiffDTO) SetAddedViolations(v []ApiPolicyViolationForDiffDTO)`

SetAddedViolations sets AddedViolations field to given value.

### HasAddedViolations

`func (o *ApiPolicyViolationDiffDTO) HasAddedViolations() bool`

HasAddedViolations returns a boolean if a field has been set.

### GetApplication

`func (o *ApiPolicyViolationDiffDTO) GetApplication() ApiApplicationDTO`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApiPolicyViolationDiffDTO) GetApplicationOk() (*ApiApplicationDTO, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApiPolicyViolationDiffDTO) SetApplication(v ApiApplicationDTO)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApiPolicyViolationDiffDTO) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetDiffTime

`func (o *ApiPolicyViolationDiffDTO) GetDiffTime() time.Time`

GetDiffTime returns the DiffTime field if non-nil, zero value otherwise.

### GetDiffTimeOk

`func (o *ApiPolicyViolationDiffDTO) GetDiffTimeOk() (*time.Time, bool)`

GetDiffTimeOk returns a tuple with the DiffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffTime

`func (o *ApiPolicyViolationDiffDTO) SetDiffTime(v time.Time)`

SetDiffTime sets DiffTime field to given value.

### HasDiffTime

`func (o *ApiPolicyViolationDiffDTO) HasDiffTime() bool`

HasDiffTime returns a boolean if a field has been set.

### GetFromCommit

`func (o *ApiPolicyViolationDiffDTO) GetFromCommit() ApiApplicationEvaluationCommitDTO`

GetFromCommit returns the FromCommit field if non-nil, zero value otherwise.

### GetFromCommitOk

`func (o *ApiPolicyViolationDiffDTO) GetFromCommitOk() (*ApiApplicationEvaluationCommitDTO, bool)`

GetFromCommitOk returns a tuple with the FromCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCommit

`func (o *ApiPolicyViolationDiffDTO) SetFromCommit(v ApiApplicationEvaluationCommitDTO)`

SetFromCommit sets FromCommit field to given value.

### HasFromCommit

`func (o *ApiPolicyViolationDiffDTO) HasFromCommit() bool`

HasFromCommit returns a boolean if a field has been set.

### GetRemovedViolations

`func (o *ApiPolicyViolationDiffDTO) GetRemovedViolations() []ApiPolicyViolationForDiffDTO`

GetRemovedViolations returns the RemovedViolations field if non-nil, zero value otherwise.

### GetRemovedViolationsOk

`func (o *ApiPolicyViolationDiffDTO) GetRemovedViolationsOk() (*[]ApiPolicyViolationForDiffDTO, bool)`

GetRemovedViolationsOk returns a tuple with the RemovedViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedViolations

`func (o *ApiPolicyViolationDiffDTO) SetRemovedViolations(v []ApiPolicyViolationForDiffDTO)`

SetRemovedViolations sets RemovedViolations field to given value.

### HasRemovedViolations

`func (o *ApiPolicyViolationDiffDTO) HasRemovedViolations() bool`

HasRemovedViolations returns a boolean if a field has been set.

### GetSameViolations

`func (o *ApiPolicyViolationDiffDTO) GetSameViolations() []ApiPolicyViolationForDiffDTO`

GetSameViolations returns the SameViolations field if non-nil, zero value otherwise.

### GetSameViolationsOk

`func (o *ApiPolicyViolationDiffDTO) GetSameViolationsOk() (*[]ApiPolicyViolationForDiffDTO, bool)`

GetSameViolationsOk returns a tuple with the SameViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSameViolations

`func (o *ApiPolicyViolationDiffDTO) SetSameViolations(v []ApiPolicyViolationForDiffDTO)`

SetSameViolations sets SameViolations field to given value.

### HasSameViolations

`func (o *ApiPolicyViolationDiffDTO) HasSameViolations() bool`

HasSameViolations returns a boolean if a field has been set.

### GetToCommit

`func (o *ApiPolicyViolationDiffDTO) GetToCommit() ApiApplicationEvaluationCommitDTO`

GetToCommit returns the ToCommit field if non-nil, zero value otherwise.

### GetToCommitOk

`func (o *ApiPolicyViolationDiffDTO) GetToCommitOk() (*ApiApplicationEvaluationCommitDTO, bool)`

GetToCommitOk returns a tuple with the ToCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCommit

`func (o *ApiPolicyViolationDiffDTO) SetToCommit(v ApiApplicationEvaluationCommitDTO)`

SetToCommit sets ToCommit field to given value.

### HasToCommit

`func (o *ApiPolicyViolationDiffDTO) HasToCommit() bool`

HasToCommit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


