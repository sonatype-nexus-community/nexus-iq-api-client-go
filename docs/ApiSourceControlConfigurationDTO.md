# ApiSourceControlConfigurationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneDirectory** | Pointer to **string** |  | [optional] 
**CommitEmail** | Pointer to **string** |  | [optional] 
**CommitUsername** | Pointer to **string** |  | [optional] 
**DefaultBranchMonitoringIntervalHours** | Pointer to **int32** |  | [optional] 
**DefaultBranchMonitoringStartTime** | Pointer to **string** |  | [optional] 
**GitExecutable** | Pointer to **string** |  | [optional] 
**GitImplementation** | Pointer to **string** |  | [optional] 
**GitTimeoutSeconds** | Pointer to **int32** |  | [optional] 
**PrCommentPurgeWindow** | Pointer to **int32** |  | [optional] 
**PrEventPurgeWindow** | Pointer to **int32** |  | [optional] 
**PullRequestMonitoringIntervalSeconds** | Pointer to **int32** |  | [optional] 
**UseUsernameInRepositoryCloneUrl** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiSourceControlConfigurationDTO

`func NewApiSourceControlConfigurationDTO() *ApiSourceControlConfigurationDTO`

NewApiSourceControlConfigurationDTO instantiates a new ApiSourceControlConfigurationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSourceControlConfigurationDTOWithDefaults

`func NewApiSourceControlConfigurationDTOWithDefaults() *ApiSourceControlConfigurationDTO`

NewApiSourceControlConfigurationDTOWithDefaults instantiates a new ApiSourceControlConfigurationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneDirectory

`func (o *ApiSourceControlConfigurationDTO) GetCloneDirectory() string`

GetCloneDirectory returns the CloneDirectory field if non-nil, zero value otherwise.

### GetCloneDirectoryOk

`func (o *ApiSourceControlConfigurationDTO) GetCloneDirectoryOk() (*string, bool)`

GetCloneDirectoryOk returns a tuple with the CloneDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneDirectory

`func (o *ApiSourceControlConfigurationDTO) SetCloneDirectory(v string)`

SetCloneDirectory sets CloneDirectory field to given value.

### HasCloneDirectory

`func (o *ApiSourceControlConfigurationDTO) HasCloneDirectory() bool`

HasCloneDirectory returns a boolean if a field has been set.

### GetCommitEmail

`func (o *ApiSourceControlConfigurationDTO) GetCommitEmail() string`

GetCommitEmail returns the CommitEmail field if non-nil, zero value otherwise.

### GetCommitEmailOk

`func (o *ApiSourceControlConfigurationDTO) GetCommitEmailOk() (*string, bool)`

GetCommitEmailOk returns a tuple with the CommitEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitEmail

`func (o *ApiSourceControlConfigurationDTO) SetCommitEmail(v string)`

SetCommitEmail sets CommitEmail field to given value.

### HasCommitEmail

`func (o *ApiSourceControlConfigurationDTO) HasCommitEmail() bool`

HasCommitEmail returns a boolean if a field has been set.

### GetCommitUsername

`func (o *ApiSourceControlConfigurationDTO) GetCommitUsername() string`

GetCommitUsername returns the CommitUsername field if non-nil, zero value otherwise.

### GetCommitUsernameOk

`func (o *ApiSourceControlConfigurationDTO) GetCommitUsernameOk() (*string, bool)`

GetCommitUsernameOk returns a tuple with the CommitUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUsername

`func (o *ApiSourceControlConfigurationDTO) SetCommitUsername(v string)`

SetCommitUsername sets CommitUsername field to given value.

### HasCommitUsername

`func (o *ApiSourceControlConfigurationDTO) HasCommitUsername() bool`

HasCommitUsername returns a boolean if a field has been set.

### GetDefaultBranchMonitoringIntervalHours

`func (o *ApiSourceControlConfigurationDTO) GetDefaultBranchMonitoringIntervalHours() int32`

GetDefaultBranchMonitoringIntervalHours returns the DefaultBranchMonitoringIntervalHours field if non-nil, zero value otherwise.

### GetDefaultBranchMonitoringIntervalHoursOk

`func (o *ApiSourceControlConfigurationDTO) GetDefaultBranchMonitoringIntervalHoursOk() (*int32, bool)`

GetDefaultBranchMonitoringIntervalHoursOk returns a tuple with the DefaultBranchMonitoringIntervalHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchMonitoringIntervalHours

`func (o *ApiSourceControlConfigurationDTO) SetDefaultBranchMonitoringIntervalHours(v int32)`

SetDefaultBranchMonitoringIntervalHours sets DefaultBranchMonitoringIntervalHours field to given value.

### HasDefaultBranchMonitoringIntervalHours

`func (o *ApiSourceControlConfigurationDTO) HasDefaultBranchMonitoringIntervalHours() bool`

HasDefaultBranchMonitoringIntervalHours returns a boolean if a field has been set.

### GetDefaultBranchMonitoringStartTime

`func (o *ApiSourceControlConfigurationDTO) GetDefaultBranchMonitoringStartTime() string`

GetDefaultBranchMonitoringStartTime returns the DefaultBranchMonitoringStartTime field if non-nil, zero value otherwise.

### GetDefaultBranchMonitoringStartTimeOk

`func (o *ApiSourceControlConfigurationDTO) GetDefaultBranchMonitoringStartTimeOk() (*string, bool)`

GetDefaultBranchMonitoringStartTimeOk returns a tuple with the DefaultBranchMonitoringStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchMonitoringStartTime

`func (o *ApiSourceControlConfigurationDTO) SetDefaultBranchMonitoringStartTime(v string)`

SetDefaultBranchMonitoringStartTime sets DefaultBranchMonitoringStartTime field to given value.

### HasDefaultBranchMonitoringStartTime

`func (o *ApiSourceControlConfigurationDTO) HasDefaultBranchMonitoringStartTime() bool`

HasDefaultBranchMonitoringStartTime returns a boolean if a field has been set.

### GetGitExecutable

`func (o *ApiSourceControlConfigurationDTO) GetGitExecutable() string`

GetGitExecutable returns the GitExecutable field if non-nil, zero value otherwise.

### GetGitExecutableOk

`func (o *ApiSourceControlConfigurationDTO) GetGitExecutableOk() (*string, bool)`

GetGitExecutableOk returns a tuple with the GitExecutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitExecutable

`func (o *ApiSourceControlConfigurationDTO) SetGitExecutable(v string)`

SetGitExecutable sets GitExecutable field to given value.

### HasGitExecutable

`func (o *ApiSourceControlConfigurationDTO) HasGitExecutable() bool`

HasGitExecutable returns a boolean if a field has been set.

### GetGitImplementation

`func (o *ApiSourceControlConfigurationDTO) GetGitImplementation() string`

GetGitImplementation returns the GitImplementation field if non-nil, zero value otherwise.

### GetGitImplementationOk

`func (o *ApiSourceControlConfigurationDTO) GetGitImplementationOk() (*string, bool)`

GetGitImplementationOk returns a tuple with the GitImplementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitImplementation

`func (o *ApiSourceControlConfigurationDTO) SetGitImplementation(v string)`

SetGitImplementation sets GitImplementation field to given value.

### HasGitImplementation

`func (o *ApiSourceControlConfigurationDTO) HasGitImplementation() bool`

HasGitImplementation returns a boolean if a field has been set.

### GetGitTimeoutSeconds

`func (o *ApiSourceControlConfigurationDTO) GetGitTimeoutSeconds() int32`

GetGitTimeoutSeconds returns the GitTimeoutSeconds field if non-nil, zero value otherwise.

### GetGitTimeoutSecondsOk

`func (o *ApiSourceControlConfigurationDTO) GetGitTimeoutSecondsOk() (*int32, bool)`

GetGitTimeoutSecondsOk returns a tuple with the GitTimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTimeoutSeconds

`func (o *ApiSourceControlConfigurationDTO) SetGitTimeoutSeconds(v int32)`

SetGitTimeoutSeconds sets GitTimeoutSeconds field to given value.

### HasGitTimeoutSeconds

`func (o *ApiSourceControlConfigurationDTO) HasGitTimeoutSeconds() bool`

HasGitTimeoutSeconds returns a boolean if a field has been set.

### GetPrCommentPurgeWindow

`func (o *ApiSourceControlConfigurationDTO) GetPrCommentPurgeWindow() int32`

GetPrCommentPurgeWindow returns the PrCommentPurgeWindow field if non-nil, zero value otherwise.

### GetPrCommentPurgeWindowOk

`func (o *ApiSourceControlConfigurationDTO) GetPrCommentPurgeWindowOk() (*int32, bool)`

GetPrCommentPurgeWindowOk returns a tuple with the PrCommentPurgeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrCommentPurgeWindow

`func (o *ApiSourceControlConfigurationDTO) SetPrCommentPurgeWindow(v int32)`

SetPrCommentPurgeWindow sets PrCommentPurgeWindow field to given value.

### HasPrCommentPurgeWindow

`func (o *ApiSourceControlConfigurationDTO) HasPrCommentPurgeWindow() bool`

HasPrCommentPurgeWindow returns a boolean if a field has been set.

### GetPrEventPurgeWindow

`func (o *ApiSourceControlConfigurationDTO) GetPrEventPurgeWindow() int32`

GetPrEventPurgeWindow returns the PrEventPurgeWindow field if non-nil, zero value otherwise.

### GetPrEventPurgeWindowOk

`func (o *ApiSourceControlConfigurationDTO) GetPrEventPurgeWindowOk() (*int32, bool)`

GetPrEventPurgeWindowOk returns a tuple with the PrEventPurgeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrEventPurgeWindow

`func (o *ApiSourceControlConfigurationDTO) SetPrEventPurgeWindow(v int32)`

SetPrEventPurgeWindow sets PrEventPurgeWindow field to given value.

### HasPrEventPurgeWindow

`func (o *ApiSourceControlConfigurationDTO) HasPrEventPurgeWindow() bool`

HasPrEventPurgeWindow returns a boolean if a field has been set.

### GetPullRequestMonitoringIntervalSeconds

`func (o *ApiSourceControlConfigurationDTO) GetPullRequestMonitoringIntervalSeconds() int32`

GetPullRequestMonitoringIntervalSeconds returns the PullRequestMonitoringIntervalSeconds field if non-nil, zero value otherwise.

### GetPullRequestMonitoringIntervalSecondsOk

`func (o *ApiSourceControlConfigurationDTO) GetPullRequestMonitoringIntervalSecondsOk() (*int32, bool)`

GetPullRequestMonitoringIntervalSecondsOk returns a tuple with the PullRequestMonitoringIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestMonitoringIntervalSeconds

`func (o *ApiSourceControlConfigurationDTO) SetPullRequestMonitoringIntervalSeconds(v int32)`

SetPullRequestMonitoringIntervalSeconds sets PullRequestMonitoringIntervalSeconds field to given value.

### HasPullRequestMonitoringIntervalSeconds

`func (o *ApiSourceControlConfigurationDTO) HasPullRequestMonitoringIntervalSeconds() bool`

HasPullRequestMonitoringIntervalSeconds returns a boolean if a field has been set.

### GetUseUsernameInRepositoryCloneUrl

`func (o *ApiSourceControlConfigurationDTO) GetUseUsernameInRepositoryCloneUrl() bool`

GetUseUsernameInRepositoryCloneUrl returns the UseUsernameInRepositoryCloneUrl field if non-nil, zero value otherwise.

### GetUseUsernameInRepositoryCloneUrlOk

`func (o *ApiSourceControlConfigurationDTO) GetUseUsernameInRepositoryCloneUrlOk() (*bool, bool)`

GetUseUsernameInRepositoryCloneUrlOk returns a tuple with the UseUsernameInRepositoryCloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseUsernameInRepositoryCloneUrl

`func (o *ApiSourceControlConfigurationDTO) SetUseUsernameInRepositoryCloneUrl(v bool)`

SetUseUsernameInRepositoryCloneUrl sets UseUsernameInRepositoryCloneUrl field to given value.

### HasUseUsernameInRepositoryCloneUrl

`func (o *ApiSourceControlConfigurationDTO) HasUseUsernameInRepositoryCloneUrl() bool`

HasUseUsernameInRepositoryCloneUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


