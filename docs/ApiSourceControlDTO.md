# ApiSourceControlDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseBranch** | Pointer to **string** |  | [optional] 
**CommitStatusEnabled** | Pointer to **bool** |  | [optional] 
**EnablePullRequests** | Pointer to **bool** |  | [optional] 
**EnableStatusChecks** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ManualPullRequestsEnabled** | Pointer to **bool** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**PullRequestCommentingEnabled** | Pointer to **bool** |  | [optional] 
**RemediationPullRequestsEnabled** | Pointer to **bool** |  | [optional] 
**RepositoryUrl** | Pointer to **string** |  | [optional] 
**SourceControlEvaluationsEnabled** | Pointer to **bool** |  | [optional] 
**SourceControlScanTarget** | Pointer to **string** |  | [optional] 
**SshEnabled** | Pointer to **bool** |  | [optional] 
**StatusChecksEnabled** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewApiSourceControlDTO

`func NewApiSourceControlDTO() *ApiSourceControlDTO`

NewApiSourceControlDTO instantiates a new ApiSourceControlDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSourceControlDTOWithDefaults

`func NewApiSourceControlDTOWithDefaults() *ApiSourceControlDTO`

NewApiSourceControlDTOWithDefaults instantiates a new ApiSourceControlDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseBranch

`func (o *ApiSourceControlDTO) GetBaseBranch() string`

GetBaseBranch returns the BaseBranch field if non-nil, zero value otherwise.

### GetBaseBranchOk

`func (o *ApiSourceControlDTO) GetBaseBranchOk() (*string, bool)`

GetBaseBranchOk returns a tuple with the BaseBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBranch

`func (o *ApiSourceControlDTO) SetBaseBranch(v string)`

SetBaseBranch sets BaseBranch field to given value.

### HasBaseBranch

`func (o *ApiSourceControlDTO) HasBaseBranch() bool`

HasBaseBranch returns a boolean if a field has been set.

### GetCommitStatusEnabled

`func (o *ApiSourceControlDTO) GetCommitStatusEnabled() bool`

GetCommitStatusEnabled returns the CommitStatusEnabled field if non-nil, zero value otherwise.

### GetCommitStatusEnabledOk

`func (o *ApiSourceControlDTO) GetCommitStatusEnabledOk() (*bool, bool)`

GetCommitStatusEnabledOk returns a tuple with the CommitStatusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitStatusEnabled

`func (o *ApiSourceControlDTO) SetCommitStatusEnabled(v bool)`

SetCommitStatusEnabled sets CommitStatusEnabled field to given value.

### HasCommitStatusEnabled

`func (o *ApiSourceControlDTO) HasCommitStatusEnabled() bool`

HasCommitStatusEnabled returns a boolean if a field has been set.

### GetEnablePullRequests

`func (o *ApiSourceControlDTO) GetEnablePullRequests() bool`

GetEnablePullRequests returns the EnablePullRequests field if non-nil, zero value otherwise.

### GetEnablePullRequestsOk

`func (o *ApiSourceControlDTO) GetEnablePullRequestsOk() (*bool, bool)`

GetEnablePullRequestsOk returns a tuple with the EnablePullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePullRequests

`func (o *ApiSourceControlDTO) SetEnablePullRequests(v bool)`

SetEnablePullRequests sets EnablePullRequests field to given value.

### HasEnablePullRequests

`func (o *ApiSourceControlDTO) HasEnablePullRequests() bool`

HasEnablePullRequests returns a boolean if a field has been set.

### GetEnableStatusChecks

`func (o *ApiSourceControlDTO) GetEnableStatusChecks() bool`

GetEnableStatusChecks returns the EnableStatusChecks field if non-nil, zero value otherwise.

### GetEnableStatusChecksOk

`func (o *ApiSourceControlDTO) GetEnableStatusChecksOk() (*bool, bool)`

GetEnableStatusChecksOk returns a tuple with the EnableStatusChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableStatusChecks

`func (o *ApiSourceControlDTO) SetEnableStatusChecks(v bool)`

SetEnableStatusChecks sets EnableStatusChecks field to given value.

### HasEnableStatusChecks

`func (o *ApiSourceControlDTO) HasEnableStatusChecks() bool`

HasEnableStatusChecks returns a boolean if a field has been set.

### GetId

`func (o *ApiSourceControlDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiSourceControlDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiSourceControlDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiSourceControlDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManualPullRequestsEnabled

`func (o *ApiSourceControlDTO) GetManualPullRequestsEnabled() bool`

GetManualPullRequestsEnabled returns the ManualPullRequestsEnabled field if non-nil, zero value otherwise.

### GetManualPullRequestsEnabledOk

`func (o *ApiSourceControlDTO) GetManualPullRequestsEnabledOk() (*bool, bool)`

GetManualPullRequestsEnabledOk returns a tuple with the ManualPullRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualPullRequestsEnabled

`func (o *ApiSourceControlDTO) SetManualPullRequestsEnabled(v bool)`

SetManualPullRequestsEnabled sets ManualPullRequestsEnabled field to given value.

### HasManualPullRequestsEnabled

`func (o *ApiSourceControlDTO) HasManualPullRequestsEnabled() bool`

HasManualPullRequestsEnabled returns a boolean if a field has been set.

### GetOwnerId

`func (o *ApiSourceControlDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ApiSourceControlDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ApiSourceControlDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ApiSourceControlDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetProvider

`func (o *ApiSourceControlDTO) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApiSourceControlDTO) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApiSourceControlDTO) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ApiSourceControlDTO) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPullRequestCommentingEnabled

`func (o *ApiSourceControlDTO) GetPullRequestCommentingEnabled() bool`

GetPullRequestCommentingEnabled returns the PullRequestCommentingEnabled field if non-nil, zero value otherwise.

### GetPullRequestCommentingEnabledOk

`func (o *ApiSourceControlDTO) GetPullRequestCommentingEnabledOk() (*bool, bool)`

GetPullRequestCommentingEnabledOk returns a tuple with the PullRequestCommentingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestCommentingEnabled

`func (o *ApiSourceControlDTO) SetPullRequestCommentingEnabled(v bool)`

SetPullRequestCommentingEnabled sets PullRequestCommentingEnabled field to given value.

### HasPullRequestCommentingEnabled

`func (o *ApiSourceControlDTO) HasPullRequestCommentingEnabled() bool`

HasPullRequestCommentingEnabled returns a boolean if a field has been set.

### GetRemediationPullRequestsEnabled

`func (o *ApiSourceControlDTO) GetRemediationPullRequestsEnabled() bool`

GetRemediationPullRequestsEnabled returns the RemediationPullRequestsEnabled field if non-nil, zero value otherwise.

### GetRemediationPullRequestsEnabledOk

`func (o *ApiSourceControlDTO) GetRemediationPullRequestsEnabledOk() (*bool, bool)`

GetRemediationPullRequestsEnabledOk returns a tuple with the RemediationPullRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationPullRequestsEnabled

`func (o *ApiSourceControlDTO) SetRemediationPullRequestsEnabled(v bool)`

SetRemediationPullRequestsEnabled sets RemediationPullRequestsEnabled field to given value.

### HasRemediationPullRequestsEnabled

`func (o *ApiSourceControlDTO) HasRemediationPullRequestsEnabled() bool`

HasRemediationPullRequestsEnabled returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *ApiSourceControlDTO) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *ApiSourceControlDTO) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *ApiSourceControlDTO) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *ApiSourceControlDTO) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.

### GetSourceControlEvaluationsEnabled

`func (o *ApiSourceControlDTO) GetSourceControlEvaluationsEnabled() bool`

GetSourceControlEvaluationsEnabled returns the SourceControlEvaluationsEnabled field if non-nil, zero value otherwise.

### GetSourceControlEvaluationsEnabledOk

`func (o *ApiSourceControlDTO) GetSourceControlEvaluationsEnabledOk() (*bool, bool)`

GetSourceControlEvaluationsEnabledOk returns a tuple with the SourceControlEvaluationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceControlEvaluationsEnabled

`func (o *ApiSourceControlDTO) SetSourceControlEvaluationsEnabled(v bool)`

SetSourceControlEvaluationsEnabled sets SourceControlEvaluationsEnabled field to given value.

### HasSourceControlEvaluationsEnabled

`func (o *ApiSourceControlDTO) HasSourceControlEvaluationsEnabled() bool`

HasSourceControlEvaluationsEnabled returns a boolean if a field has been set.

### GetSourceControlScanTarget

`func (o *ApiSourceControlDTO) GetSourceControlScanTarget() string`

GetSourceControlScanTarget returns the SourceControlScanTarget field if non-nil, zero value otherwise.

### GetSourceControlScanTargetOk

`func (o *ApiSourceControlDTO) GetSourceControlScanTargetOk() (*string, bool)`

GetSourceControlScanTargetOk returns a tuple with the SourceControlScanTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceControlScanTarget

`func (o *ApiSourceControlDTO) SetSourceControlScanTarget(v string)`

SetSourceControlScanTarget sets SourceControlScanTarget field to given value.

### HasSourceControlScanTarget

`func (o *ApiSourceControlDTO) HasSourceControlScanTarget() bool`

HasSourceControlScanTarget returns a boolean if a field has been set.

### GetSshEnabled

`func (o *ApiSourceControlDTO) GetSshEnabled() bool`

GetSshEnabled returns the SshEnabled field if non-nil, zero value otherwise.

### GetSshEnabledOk

`func (o *ApiSourceControlDTO) GetSshEnabledOk() (*bool, bool)`

GetSshEnabledOk returns a tuple with the SshEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshEnabled

`func (o *ApiSourceControlDTO) SetSshEnabled(v bool)`

SetSshEnabled sets SshEnabled field to given value.

### HasSshEnabled

`func (o *ApiSourceControlDTO) HasSshEnabled() bool`

HasSshEnabled returns a boolean if a field has been set.

### GetStatusChecksEnabled

`func (o *ApiSourceControlDTO) GetStatusChecksEnabled() bool`

GetStatusChecksEnabled returns the StatusChecksEnabled field if non-nil, zero value otherwise.

### GetStatusChecksEnabledOk

`func (o *ApiSourceControlDTO) GetStatusChecksEnabledOk() (*bool, bool)`

GetStatusChecksEnabledOk returns a tuple with the StatusChecksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChecksEnabled

`func (o *ApiSourceControlDTO) SetStatusChecksEnabled(v bool)`

SetStatusChecksEnabled sets StatusChecksEnabled field to given value.

### HasStatusChecksEnabled

`func (o *ApiSourceControlDTO) HasStatusChecksEnabled() bool`

HasStatusChecksEnabled returns a boolean if a field has been set.

### GetToken

`func (o *ApiSourceControlDTO) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApiSourceControlDTO) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApiSourceControlDTO) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ApiSourceControlDTO) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsername

`func (o *ApiSourceControlDTO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiSourceControlDTO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiSourceControlDTO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiSourceControlDTO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


