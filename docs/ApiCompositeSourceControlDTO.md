# ApiCompositeSourceControlDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseBranch** | Pointer to [**ApiCompositeValueDTOString**](ApiCompositeValueDTOString.md) |  | [optional] 
**ClosePrAfterDays** | Pointer to [**ApiCompositeValueDTOInteger**](ApiCompositeValueDTOInteger.md) |  | [optional] 
**ClosePrAfterDaysOpenEnabled** | Pointer to [**ApiCompositeValueDTOBoolean**](ApiCompositeValueDTOBoolean.md) |  | [optional] 
**ClosePrOnFailedChecksEnabled** | Pointer to [**ApiCompositeValueDTOBoolean**](ApiCompositeValueDTOBoolean.md) |  | [optional] 
**CommitStatusEnabled** | Pointer to [**ApiCompositeValueDTOBoolean**](ApiCompositeValueDTOBoolean.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InnerSourceAutomatedUpdatesEnabled** | Pointer to [**ApiCompositeValueDTOBoolean**](ApiCompositeValueDTOBoolean.md) |  | [optional] 
**ManualPullRequestsEnabled** | Pointer to [**ApiCompositeValueDTOBoolean**](ApiCompositeValueDTOBoolean.md) |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**ApiCompositeValueDTOString**](ApiCompositeValueDTOString.md) |  | [optional] 
**PullRequestCommentingEnabled** | Pointer to [**ApiCompositeValueDTOBoolean**](ApiCompositeValueDTOBoolean.md) |  | [optional] 
**RemediationPullRequestsEnabled** | Pointer to [**ApiCompositeValueDTOBoolean**](ApiCompositeValueDTOBoolean.md) |  | [optional] 
**RepositoryUrl** | Pointer to **string** |  | [optional] 
**SourceControlEvaluationsEnabled** | Pointer to [**ApiCompositeValueDTOBoolean**](ApiCompositeValueDTOBoolean.md) |  | [optional] 
**SourceControlScanTarget** | Pointer to [**ApiCompositeValueDTOString**](ApiCompositeValueDTOString.md) |  | [optional] 
**SshEnabled** | Pointer to [**ApiCompositeValueDTOBoolean**](ApiCompositeValueDTOBoolean.md) |  | [optional] 
**StatusChecksEnabled** | Pointer to [**ApiCompositeValueDTOBoolean**](ApiCompositeValueDTOBoolean.md) |  | [optional] 
**Token** | Pointer to [**ApiCompositeValueDTOString**](ApiCompositeValueDTOString.md) |  | [optional] 
**Username** | Pointer to [**ApiCompositeValueDTOString**](ApiCompositeValueDTOString.md) |  | [optional] 

## Methods

### NewApiCompositeSourceControlDTO

`func NewApiCompositeSourceControlDTO() *ApiCompositeSourceControlDTO`

NewApiCompositeSourceControlDTO instantiates a new ApiCompositeSourceControlDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCompositeSourceControlDTOWithDefaults

`func NewApiCompositeSourceControlDTOWithDefaults() *ApiCompositeSourceControlDTO`

NewApiCompositeSourceControlDTOWithDefaults instantiates a new ApiCompositeSourceControlDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseBranch

`func (o *ApiCompositeSourceControlDTO) GetBaseBranch() ApiCompositeValueDTOString`

GetBaseBranch returns the BaseBranch field if non-nil, zero value otherwise.

### GetBaseBranchOk

`func (o *ApiCompositeSourceControlDTO) GetBaseBranchOk() (*ApiCompositeValueDTOString, bool)`

GetBaseBranchOk returns a tuple with the BaseBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseBranch

`func (o *ApiCompositeSourceControlDTO) SetBaseBranch(v ApiCompositeValueDTOString)`

SetBaseBranch sets BaseBranch field to given value.

### HasBaseBranch

`func (o *ApiCompositeSourceControlDTO) HasBaseBranch() bool`

HasBaseBranch returns a boolean if a field has been set.

### GetClosePrAfterDays

`func (o *ApiCompositeSourceControlDTO) GetClosePrAfterDays() ApiCompositeValueDTOInteger`

GetClosePrAfterDays returns the ClosePrAfterDays field if non-nil, zero value otherwise.

### GetClosePrAfterDaysOk

`func (o *ApiCompositeSourceControlDTO) GetClosePrAfterDaysOk() (*ApiCompositeValueDTOInteger, bool)`

GetClosePrAfterDaysOk returns a tuple with the ClosePrAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrAfterDays

`func (o *ApiCompositeSourceControlDTO) SetClosePrAfterDays(v ApiCompositeValueDTOInteger)`

SetClosePrAfterDays sets ClosePrAfterDays field to given value.

### HasClosePrAfterDays

`func (o *ApiCompositeSourceControlDTO) HasClosePrAfterDays() bool`

HasClosePrAfterDays returns a boolean if a field has been set.

### GetClosePrAfterDaysOpenEnabled

`func (o *ApiCompositeSourceControlDTO) GetClosePrAfterDaysOpenEnabled() ApiCompositeValueDTOBoolean`

GetClosePrAfterDaysOpenEnabled returns the ClosePrAfterDaysOpenEnabled field if non-nil, zero value otherwise.

### GetClosePrAfterDaysOpenEnabledOk

`func (o *ApiCompositeSourceControlDTO) GetClosePrAfterDaysOpenEnabledOk() (*ApiCompositeValueDTOBoolean, bool)`

GetClosePrAfterDaysOpenEnabledOk returns a tuple with the ClosePrAfterDaysOpenEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrAfterDaysOpenEnabled

`func (o *ApiCompositeSourceControlDTO) SetClosePrAfterDaysOpenEnabled(v ApiCompositeValueDTOBoolean)`

SetClosePrAfterDaysOpenEnabled sets ClosePrAfterDaysOpenEnabled field to given value.

### HasClosePrAfterDaysOpenEnabled

`func (o *ApiCompositeSourceControlDTO) HasClosePrAfterDaysOpenEnabled() bool`

HasClosePrAfterDaysOpenEnabled returns a boolean if a field has been set.

### GetClosePrOnFailedChecksEnabled

`func (o *ApiCompositeSourceControlDTO) GetClosePrOnFailedChecksEnabled() ApiCompositeValueDTOBoolean`

GetClosePrOnFailedChecksEnabled returns the ClosePrOnFailedChecksEnabled field if non-nil, zero value otherwise.

### GetClosePrOnFailedChecksEnabledOk

`func (o *ApiCompositeSourceControlDTO) GetClosePrOnFailedChecksEnabledOk() (*ApiCompositeValueDTOBoolean, bool)`

GetClosePrOnFailedChecksEnabledOk returns a tuple with the ClosePrOnFailedChecksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosePrOnFailedChecksEnabled

`func (o *ApiCompositeSourceControlDTO) SetClosePrOnFailedChecksEnabled(v ApiCompositeValueDTOBoolean)`

SetClosePrOnFailedChecksEnabled sets ClosePrOnFailedChecksEnabled field to given value.

### HasClosePrOnFailedChecksEnabled

`func (o *ApiCompositeSourceControlDTO) HasClosePrOnFailedChecksEnabled() bool`

HasClosePrOnFailedChecksEnabled returns a boolean if a field has been set.

### GetCommitStatusEnabled

`func (o *ApiCompositeSourceControlDTO) GetCommitStatusEnabled() ApiCompositeValueDTOBoolean`

GetCommitStatusEnabled returns the CommitStatusEnabled field if non-nil, zero value otherwise.

### GetCommitStatusEnabledOk

`func (o *ApiCompositeSourceControlDTO) GetCommitStatusEnabledOk() (*ApiCompositeValueDTOBoolean, bool)`

GetCommitStatusEnabledOk returns a tuple with the CommitStatusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitStatusEnabled

`func (o *ApiCompositeSourceControlDTO) SetCommitStatusEnabled(v ApiCompositeValueDTOBoolean)`

SetCommitStatusEnabled sets CommitStatusEnabled field to given value.

### HasCommitStatusEnabled

`func (o *ApiCompositeSourceControlDTO) HasCommitStatusEnabled() bool`

HasCommitStatusEnabled returns a boolean if a field has been set.

### GetId

`func (o *ApiCompositeSourceControlDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiCompositeSourceControlDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiCompositeSourceControlDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiCompositeSourceControlDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInnerSourceAutomatedUpdatesEnabled

`func (o *ApiCompositeSourceControlDTO) GetInnerSourceAutomatedUpdatesEnabled() ApiCompositeValueDTOBoolean`

GetInnerSourceAutomatedUpdatesEnabled returns the InnerSourceAutomatedUpdatesEnabled field if non-nil, zero value otherwise.

### GetInnerSourceAutomatedUpdatesEnabledOk

`func (o *ApiCompositeSourceControlDTO) GetInnerSourceAutomatedUpdatesEnabledOk() (*ApiCompositeValueDTOBoolean, bool)`

GetInnerSourceAutomatedUpdatesEnabledOk returns a tuple with the InnerSourceAutomatedUpdatesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerSourceAutomatedUpdatesEnabled

`func (o *ApiCompositeSourceControlDTO) SetInnerSourceAutomatedUpdatesEnabled(v ApiCompositeValueDTOBoolean)`

SetInnerSourceAutomatedUpdatesEnabled sets InnerSourceAutomatedUpdatesEnabled field to given value.

### HasInnerSourceAutomatedUpdatesEnabled

`func (o *ApiCompositeSourceControlDTO) HasInnerSourceAutomatedUpdatesEnabled() bool`

HasInnerSourceAutomatedUpdatesEnabled returns a boolean if a field has been set.

### GetManualPullRequestsEnabled

`func (o *ApiCompositeSourceControlDTO) GetManualPullRequestsEnabled() ApiCompositeValueDTOBoolean`

GetManualPullRequestsEnabled returns the ManualPullRequestsEnabled field if non-nil, zero value otherwise.

### GetManualPullRequestsEnabledOk

`func (o *ApiCompositeSourceControlDTO) GetManualPullRequestsEnabledOk() (*ApiCompositeValueDTOBoolean, bool)`

GetManualPullRequestsEnabledOk returns a tuple with the ManualPullRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualPullRequestsEnabled

`func (o *ApiCompositeSourceControlDTO) SetManualPullRequestsEnabled(v ApiCompositeValueDTOBoolean)`

SetManualPullRequestsEnabled sets ManualPullRequestsEnabled field to given value.

### HasManualPullRequestsEnabled

`func (o *ApiCompositeSourceControlDTO) HasManualPullRequestsEnabled() bool`

HasManualPullRequestsEnabled returns a boolean if a field has been set.

### GetOwnerId

`func (o *ApiCompositeSourceControlDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ApiCompositeSourceControlDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ApiCompositeSourceControlDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ApiCompositeSourceControlDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetProvider

`func (o *ApiCompositeSourceControlDTO) GetProvider() ApiCompositeValueDTOString`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApiCompositeSourceControlDTO) GetProviderOk() (*ApiCompositeValueDTOString, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApiCompositeSourceControlDTO) SetProvider(v ApiCompositeValueDTOString)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ApiCompositeSourceControlDTO) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetPullRequestCommentingEnabled

`func (o *ApiCompositeSourceControlDTO) GetPullRequestCommentingEnabled() ApiCompositeValueDTOBoolean`

GetPullRequestCommentingEnabled returns the PullRequestCommentingEnabled field if non-nil, zero value otherwise.

### GetPullRequestCommentingEnabledOk

`func (o *ApiCompositeSourceControlDTO) GetPullRequestCommentingEnabledOk() (*ApiCompositeValueDTOBoolean, bool)`

GetPullRequestCommentingEnabledOk returns a tuple with the PullRequestCommentingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequestCommentingEnabled

`func (o *ApiCompositeSourceControlDTO) SetPullRequestCommentingEnabled(v ApiCompositeValueDTOBoolean)`

SetPullRequestCommentingEnabled sets PullRequestCommentingEnabled field to given value.

### HasPullRequestCommentingEnabled

`func (o *ApiCompositeSourceControlDTO) HasPullRequestCommentingEnabled() bool`

HasPullRequestCommentingEnabled returns a boolean if a field has been set.

### GetRemediationPullRequestsEnabled

`func (o *ApiCompositeSourceControlDTO) GetRemediationPullRequestsEnabled() ApiCompositeValueDTOBoolean`

GetRemediationPullRequestsEnabled returns the RemediationPullRequestsEnabled field if non-nil, zero value otherwise.

### GetRemediationPullRequestsEnabledOk

`func (o *ApiCompositeSourceControlDTO) GetRemediationPullRequestsEnabledOk() (*ApiCompositeValueDTOBoolean, bool)`

GetRemediationPullRequestsEnabledOk returns a tuple with the RemediationPullRequestsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationPullRequestsEnabled

`func (o *ApiCompositeSourceControlDTO) SetRemediationPullRequestsEnabled(v ApiCompositeValueDTOBoolean)`

SetRemediationPullRequestsEnabled sets RemediationPullRequestsEnabled field to given value.

### HasRemediationPullRequestsEnabled

`func (o *ApiCompositeSourceControlDTO) HasRemediationPullRequestsEnabled() bool`

HasRemediationPullRequestsEnabled returns a boolean if a field has been set.

### GetRepositoryUrl

`func (o *ApiCompositeSourceControlDTO) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *ApiCompositeSourceControlDTO) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *ApiCompositeSourceControlDTO) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.

### HasRepositoryUrl

`func (o *ApiCompositeSourceControlDTO) HasRepositoryUrl() bool`

HasRepositoryUrl returns a boolean if a field has been set.

### GetSourceControlEvaluationsEnabled

`func (o *ApiCompositeSourceControlDTO) GetSourceControlEvaluationsEnabled() ApiCompositeValueDTOBoolean`

GetSourceControlEvaluationsEnabled returns the SourceControlEvaluationsEnabled field if non-nil, zero value otherwise.

### GetSourceControlEvaluationsEnabledOk

`func (o *ApiCompositeSourceControlDTO) GetSourceControlEvaluationsEnabledOk() (*ApiCompositeValueDTOBoolean, bool)`

GetSourceControlEvaluationsEnabledOk returns a tuple with the SourceControlEvaluationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceControlEvaluationsEnabled

`func (o *ApiCompositeSourceControlDTO) SetSourceControlEvaluationsEnabled(v ApiCompositeValueDTOBoolean)`

SetSourceControlEvaluationsEnabled sets SourceControlEvaluationsEnabled field to given value.

### HasSourceControlEvaluationsEnabled

`func (o *ApiCompositeSourceControlDTO) HasSourceControlEvaluationsEnabled() bool`

HasSourceControlEvaluationsEnabled returns a boolean if a field has been set.

### GetSourceControlScanTarget

`func (o *ApiCompositeSourceControlDTO) GetSourceControlScanTarget() ApiCompositeValueDTOString`

GetSourceControlScanTarget returns the SourceControlScanTarget field if non-nil, zero value otherwise.

### GetSourceControlScanTargetOk

`func (o *ApiCompositeSourceControlDTO) GetSourceControlScanTargetOk() (*ApiCompositeValueDTOString, bool)`

GetSourceControlScanTargetOk returns a tuple with the SourceControlScanTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceControlScanTarget

`func (o *ApiCompositeSourceControlDTO) SetSourceControlScanTarget(v ApiCompositeValueDTOString)`

SetSourceControlScanTarget sets SourceControlScanTarget field to given value.

### HasSourceControlScanTarget

`func (o *ApiCompositeSourceControlDTO) HasSourceControlScanTarget() bool`

HasSourceControlScanTarget returns a boolean if a field has been set.

### GetSshEnabled

`func (o *ApiCompositeSourceControlDTO) GetSshEnabled() ApiCompositeValueDTOBoolean`

GetSshEnabled returns the SshEnabled field if non-nil, zero value otherwise.

### GetSshEnabledOk

`func (o *ApiCompositeSourceControlDTO) GetSshEnabledOk() (*ApiCompositeValueDTOBoolean, bool)`

GetSshEnabledOk returns a tuple with the SshEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshEnabled

`func (o *ApiCompositeSourceControlDTO) SetSshEnabled(v ApiCompositeValueDTOBoolean)`

SetSshEnabled sets SshEnabled field to given value.

### HasSshEnabled

`func (o *ApiCompositeSourceControlDTO) HasSshEnabled() bool`

HasSshEnabled returns a boolean if a field has been set.

### GetStatusChecksEnabled

`func (o *ApiCompositeSourceControlDTO) GetStatusChecksEnabled() ApiCompositeValueDTOBoolean`

GetStatusChecksEnabled returns the StatusChecksEnabled field if non-nil, zero value otherwise.

### GetStatusChecksEnabledOk

`func (o *ApiCompositeSourceControlDTO) GetStatusChecksEnabledOk() (*ApiCompositeValueDTOBoolean, bool)`

GetStatusChecksEnabledOk returns a tuple with the StatusChecksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChecksEnabled

`func (o *ApiCompositeSourceControlDTO) SetStatusChecksEnabled(v ApiCompositeValueDTOBoolean)`

SetStatusChecksEnabled sets StatusChecksEnabled field to given value.

### HasStatusChecksEnabled

`func (o *ApiCompositeSourceControlDTO) HasStatusChecksEnabled() bool`

HasStatusChecksEnabled returns a boolean if a field has been set.

### GetToken

`func (o *ApiCompositeSourceControlDTO) GetToken() ApiCompositeValueDTOString`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ApiCompositeSourceControlDTO) GetTokenOk() (*ApiCompositeValueDTOString, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ApiCompositeSourceControlDTO) SetToken(v ApiCompositeValueDTOString)`

SetToken sets Token field to given value.

### HasToken

`func (o *ApiCompositeSourceControlDTO) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUsername

`func (o *ApiCompositeSourceControlDTO) GetUsername() ApiCompositeValueDTOString`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiCompositeSourceControlDTO) GetUsernameOk() (*ApiCompositeValueDTOString, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiCompositeSourceControlDTO) SetUsername(v ApiCompositeValueDTOString)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiCompositeSourceControlDTO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


