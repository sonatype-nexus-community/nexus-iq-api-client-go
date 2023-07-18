/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 164
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiSourceControlConfigurationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSourceControlConfigurationDTO{}

// ApiSourceControlConfigurationDTO struct for ApiSourceControlConfigurationDTO
type ApiSourceControlConfigurationDTO struct {
	CloneDirectory *string `json:"cloneDirectory,omitempty"`
	CommitEmail *string `json:"commitEmail,omitempty"`
	CommitUsername *string `json:"commitUsername,omitempty"`
	DefaultBranchMonitoringIntervalHours *int32 `json:"defaultBranchMonitoringIntervalHours,omitempty"`
	DefaultBranchMonitoringStartTime *string `json:"defaultBranchMonitoringStartTime,omitempty"`
	GitExecutable *string `json:"gitExecutable,omitempty"`
	GitImplementation *string `json:"gitImplementation,omitempty"`
	GitTimeoutSeconds *int32 `json:"gitTimeoutSeconds,omitempty"`
	PrCommentPurgeWindow *int32 `json:"prCommentPurgeWindow,omitempty"`
	PrEventPurgeWindow *int32 `json:"prEventPurgeWindow,omitempty"`
	PullRequestMonitoringIntervalSeconds *int32 `json:"pullRequestMonitoringIntervalSeconds,omitempty"`
	UseUsernameInRepositoryCloneUrl *bool `json:"useUsernameInRepositoryCloneUrl,omitempty"`
}

// NewApiSourceControlConfigurationDTO instantiates a new ApiSourceControlConfigurationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSourceControlConfigurationDTO() *ApiSourceControlConfigurationDTO {
	this := ApiSourceControlConfigurationDTO{}
	return &this
}

// NewApiSourceControlConfigurationDTOWithDefaults instantiates a new ApiSourceControlConfigurationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSourceControlConfigurationDTOWithDefaults() *ApiSourceControlConfigurationDTO {
	this := ApiSourceControlConfigurationDTO{}
	return &this
}

// GetCloneDirectory returns the CloneDirectory field value if set, zero value otherwise.
func (o *ApiSourceControlConfigurationDTO) GetCloneDirectory() string {
	if o == nil || IsNil(o.CloneDirectory) {
		var ret string
		return ret
	}
	return *o.CloneDirectory
}

// GetCloneDirectoryOk returns a tuple with the CloneDirectory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlConfigurationDTO) GetCloneDirectoryOk() (*string, bool) {
	if o == nil || IsNil(o.CloneDirectory) {
		return nil, false
	}
	return o.CloneDirectory, true
}

// HasCloneDirectory returns a boolean if a field has been set.
func (o *ApiSourceControlConfigurationDTO) HasCloneDirectory() bool {
	if o != nil && !IsNil(o.CloneDirectory) {
		return true
	}

	return false
}

// SetCloneDirectory gets a reference to the given string and assigns it to the CloneDirectory field.
func (o *ApiSourceControlConfigurationDTO) SetCloneDirectory(v string) {
	o.CloneDirectory = &v
}

// GetCommitEmail returns the CommitEmail field value if set, zero value otherwise.
func (o *ApiSourceControlConfigurationDTO) GetCommitEmail() string {
	if o == nil || IsNil(o.CommitEmail) {
		var ret string
		return ret
	}
	return *o.CommitEmail
}

// GetCommitEmailOk returns a tuple with the CommitEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlConfigurationDTO) GetCommitEmailOk() (*string, bool) {
	if o == nil || IsNil(o.CommitEmail) {
		return nil, false
	}
	return o.CommitEmail, true
}

// HasCommitEmail returns a boolean if a field has been set.
func (o *ApiSourceControlConfigurationDTO) HasCommitEmail() bool {
	if o != nil && !IsNil(o.CommitEmail) {
		return true
	}

	return false
}

// SetCommitEmail gets a reference to the given string and assigns it to the CommitEmail field.
func (o *ApiSourceControlConfigurationDTO) SetCommitEmail(v string) {
	o.CommitEmail = &v
}

// GetCommitUsername returns the CommitUsername field value if set, zero value otherwise.
func (o *ApiSourceControlConfigurationDTO) GetCommitUsername() string {
	if o == nil || IsNil(o.CommitUsername) {
		var ret string
		return ret
	}
	return *o.CommitUsername
}

// GetCommitUsernameOk returns a tuple with the CommitUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlConfigurationDTO) GetCommitUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.CommitUsername) {
		return nil, false
	}
	return o.CommitUsername, true
}

// HasCommitUsername returns a boolean if a field has been set.
func (o *ApiSourceControlConfigurationDTO) HasCommitUsername() bool {
	if o != nil && !IsNil(o.CommitUsername) {
		return true
	}

	return false
}

// SetCommitUsername gets a reference to the given string and assigns it to the CommitUsername field.
func (o *ApiSourceControlConfigurationDTO) SetCommitUsername(v string) {
	o.CommitUsername = &v
}

// GetDefaultBranchMonitoringIntervalHours returns the DefaultBranchMonitoringIntervalHours field value if set, zero value otherwise.
func (o *ApiSourceControlConfigurationDTO) GetDefaultBranchMonitoringIntervalHours() int32 {
	if o == nil || IsNil(o.DefaultBranchMonitoringIntervalHours) {
		var ret int32
		return ret
	}
	return *o.DefaultBranchMonitoringIntervalHours
}

// GetDefaultBranchMonitoringIntervalHoursOk returns a tuple with the DefaultBranchMonitoringIntervalHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlConfigurationDTO) GetDefaultBranchMonitoringIntervalHoursOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultBranchMonitoringIntervalHours) {
		return nil, false
	}
	return o.DefaultBranchMonitoringIntervalHours, true
}

// HasDefaultBranchMonitoringIntervalHours returns a boolean if a field has been set.
func (o *ApiSourceControlConfigurationDTO) HasDefaultBranchMonitoringIntervalHours() bool {
	if o != nil && !IsNil(o.DefaultBranchMonitoringIntervalHours) {
		return true
	}

	return false
}

// SetDefaultBranchMonitoringIntervalHours gets a reference to the given int32 and assigns it to the DefaultBranchMonitoringIntervalHours field.
func (o *ApiSourceControlConfigurationDTO) SetDefaultBranchMonitoringIntervalHours(v int32) {
	o.DefaultBranchMonitoringIntervalHours = &v
}

// GetDefaultBranchMonitoringStartTime returns the DefaultBranchMonitoringStartTime field value if set, zero value otherwise.
func (o *ApiSourceControlConfigurationDTO) GetDefaultBranchMonitoringStartTime() string {
	if o == nil || IsNil(o.DefaultBranchMonitoringStartTime) {
		var ret string
		return ret
	}
	return *o.DefaultBranchMonitoringStartTime
}

// GetDefaultBranchMonitoringStartTimeOk returns a tuple with the DefaultBranchMonitoringStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlConfigurationDTO) GetDefaultBranchMonitoringStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultBranchMonitoringStartTime) {
		return nil, false
	}
	return o.DefaultBranchMonitoringStartTime, true
}

// HasDefaultBranchMonitoringStartTime returns a boolean if a field has been set.
func (o *ApiSourceControlConfigurationDTO) HasDefaultBranchMonitoringStartTime() bool {
	if o != nil && !IsNil(o.DefaultBranchMonitoringStartTime) {
		return true
	}

	return false
}

// SetDefaultBranchMonitoringStartTime gets a reference to the given string and assigns it to the DefaultBranchMonitoringStartTime field.
func (o *ApiSourceControlConfigurationDTO) SetDefaultBranchMonitoringStartTime(v string) {
	o.DefaultBranchMonitoringStartTime = &v
}

// GetGitExecutable returns the GitExecutable field value if set, zero value otherwise.
func (o *ApiSourceControlConfigurationDTO) GetGitExecutable() string {
	if o == nil || IsNil(o.GitExecutable) {
		var ret string
		return ret
	}
	return *o.GitExecutable
}

// GetGitExecutableOk returns a tuple with the GitExecutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlConfigurationDTO) GetGitExecutableOk() (*string, bool) {
	if o == nil || IsNil(o.GitExecutable) {
		return nil, false
	}
	return o.GitExecutable, true
}

// HasGitExecutable returns a boolean if a field has been set.
func (o *ApiSourceControlConfigurationDTO) HasGitExecutable() bool {
	if o != nil && !IsNil(o.GitExecutable) {
		return true
	}

	return false
}

// SetGitExecutable gets a reference to the given string and assigns it to the GitExecutable field.
func (o *ApiSourceControlConfigurationDTO) SetGitExecutable(v string) {
	o.GitExecutable = &v
}

// GetGitImplementation returns the GitImplementation field value if set, zero value otherwise.
func (o *ApiSourceControlConfigurationDTO) GetGitImplementation() string {
	if o == nil || IsNil(o.GitImplementation) {
		var ret string
		return ret
	}
	return *o.GitImplementation
}

// GetGitImplementationOk returns a tuple with the GitImplementation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlConfigurationDTO) GetGitImplementationOk() (*string, bool) {
	if o == nil || IsNil(o.GitImplementation) {
		return nil, false
	}
	return o.GitImplementation, true
}

// HasGitImplementation returns a boolean if a field has been set.
func (o *ApiSourceControlConfigurationDTO) HasGitImplementation() bool {
	if o != nil && !IsNil(o.GitImplementation) {
		return true
	}

	return false
}

// SetGitImplementation gets a reference to the given string and assigns it to the GitImplementation field.
func (o *ApiSourceControlConfigurationDTO) SetGitImplementation(v string) {
	o.GitImplementation = &v
}

// GetGitTimeoutSeconds returns the GitTimeoutSeconds field value if set, zero value otherwise.
func (o *ApiSourceControlConfigurationDTO) GetGitTimeoutSeconds() int32 {
	if o == nil || IsNil(o.GitTimeoutSeconds) {
		var ret int32
		return ret
	}
	return *o.GitTimeoutSeconds
}

// GetGitTimeoutSecondsOk returns a tuple with the GitTimeoutSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlConfigurationDTO) GetGitTimeoutSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.GitTimeoutSeconds) {
		return nil, false
	}
	return o.GitTimeoutSeconds, true
}

// HasGitTimeoutSeconds returns a boolean if a field has been set.
func (o *ApiSourceControlConfigurationDTO) HasGitTimeoutSeconds() bool {
	if o != nil && !IsNil(o.GitTimeoutSeconds) {
		return true
	}

	return false
}

// SetGitTimeoutSeconds gets a reference to the given int32 and assigns it to the GitTimeoutSeconds field.
func (o *ApiSourceControlConfigurationDTO) SetGitTimeoutSeconds(v int32) {
	o.GitTimeoutSeconds = &v
}

// GetPrCommentPurgeWindow returns the PrCommentPurgeWindow field value if set, zero value otherwise.
func (o *ApiSourceControlConfigurationDTO) GetPrCommentPurgeWindow() int32 {
	if o == nil || IsNil(o.PrCommentPurgeWindow) {
		var ret int32
		return ret
	}
	return *o.PrCommentPurgeWindow
}

// GetPrCommentPurgeWindowOk returns a tuple with the PrCommentPurgeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlConfigurationDTO) GetPrCommentPurgeWindowOk() (*int32, bool) {
	if o == nil || IsNil(o.PrCommentPurgeWindow) {
		return nil, false
	}
	return o.PrCommentPurgeWindow, true
}

// HasPrCommentPurgeWindow returns a boolean if a field has been set.
func (o *ApiSourceControlConfigurationDTO) HasPrCommentPurgeWindow() bool {
	if o != nil && !IsNil(o.PrCommentPurgeWindow) {
		return true
	}

	return false
}

// SetPrCommentPurgeWindow gets a reference to the given int32 and assigns it to the PrCommentPurgeWindow field.
func (o *ApiSourceControlConfigurationDTO) SetPrCommentPurgeWindow(v int32) {
	o.PrCommentPurgeWindow = &v
}

// GetPrEventPurgeWindow returns the PrEventPurgeWindow field value if set, zero value otherwise.
func (o *ApiSourceControlConfigurationDTO) GetPrEventPurgeWindow() int32 {
	if o == nil || IsNil(o.PrEventPurgeWindow) {
		var ret int32
		return ret
	}
	return *o.PrEventPurgeWindow
}

// GetPrEventPurgeWindowOk returns a tuple with the PrEventPurgeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlConfigurationDTO) GetPrEventPurgeWindowOk() (*int32, bool) {
	if o == nil || IsNil(o.PrEventPurgeWindow) {
		return nil, false
	}
	return o.PrEventPurgeWindow, true
}

// HasPrEventPurgeWindow returns a boolean if a field has been set.
func (o *ApiSourceControlConfigurationDTO) HasPrEventPurgeWindow() bool {
	if o != nil && !IsNil(o.PrEventPurgeWindow) {
		return true
	}

	return false
}

// SetPrEventPurgeWindow gets a reference to the given int32 and assigns it to the PrEventPurgeWindow field.
func (o *ApiSourceControlConfigurationDTO) SetPrEventPurgeWindow(v int32) {
	o.PrEventPurgeWindow = &v
}

// GetPullRequestMonitoringIntervalSeconds returns the PullRequestMonitoringIntervalSeconds field value if set, zero value otherwise.
func (o *ApiSourceControlConfigurationDTO) GetPullRequestMonitoringIntervalSeconds() int32 {
	if o == nil || IsNil(o.PullRequestMonitoringIntervalSeconds) {
		var ret int32
		return ret
	}
	return *o.PullRequestMonitoringIntervalSeconds
}

// GetPullRequestMonitoringIntervalSecondsOk returns a tuple with the PullRequestMonitoringIntervalSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlConfigurationDTO) GetPullRequestMonitoringIntervalSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.PullRequestMonitoringIntervalSeconds) {
		return nil, false
	}
	return o.PullRequestMonitoringIntervalSeconds, true
}

// HasPullRequestMonitoringIntervalSeconds returns a boolean if a field has been set.
func (o *ApiSourceControlConfigurationDTO) HasPullRequestMonitoringIntervalSeconds() bool {
	if o != nil && !IsNil(o.PullRequestMonitoringIntervalSeconds) {
		return true
	}

	return false
}

// SetPullRequestMonitoringIntervalSeconds gets a reference to the given int32 and assigns it to the PullRequestMonitoringIntervalSeconds field.
func (o *ApiSourceControlConfigurationDTO) SetPullRequestMonitoringIntervalSeconds(v int32) {
	o.PullRequestMonitoringIntervalSeconds = &v
}

// GetUseUsernameInRepositoryCloneUrl returns the UseUsernameInRepositoryCloneUrl field value if set, zero value otherwise.
func (o *ApiSourceControlConfigurationDTO) GetUseUsernameInRepositoryCloneUrl() bool {
	if o == nil || IsNil(o.UseUsernameInRepositoryCloneUrl) {
		var ret bool
		return ret
	}
	return *o.UseUsernameInRepositoryCloneUrl
}

// GetUseUsernameInRepositoryCloneUrlOk returns a tuple with the UseUsernameInRepositoryCloneUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSourceControlConfigurationDTO) GetUseUsernameInRepositoryCloneUrlOk() (*bool, bool) {
	if o == nil || IsNil(o.UseUsernameInRepositoryCloneUrl) {
		return nil, false
	}
	return o.UseUsernameInRepositoryCloneUrl, true
}

// HasUseUsernameInRepositoryCloneUrl returns a boolean if a field has been set.
func (o *ApiSourceControlConfigurationDTO) HasUseUsernameInRepositoryCloneUrl() bool {
	if o != nil && !IsNil(o.UseUsernameInRepositoryCloneUrl) {
		return true
	}

	return false
}

// SetUseUsernameInRepositoryCloneUrl gets a reference to the given bool and assigns it to the UseUsernameInRepositoryCloneUrl field.
func (o *ApiSourceControlConfigurationDTO) SetUseUsernameInRepositoryCloneUrl(v bool) {
	o.UseUsernameInRepositoryCloneUrl = &v
}

func (o ApiSourceControlConfigurationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSourceControlConfigurationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloneDirectory) {
		toSerialize["cloneDirectory"] = o.CloneDirectory
	}
	if !IsNil(o.CommitEmail) {
		toSerialize["commitEmail"] = o.CommitEmail
	}
	if !IsNil(o.CommitUsername) {
		toSerialize["commitUsername"] = o.CommitUsername
	}
	if !IsNil(o.DefaultBranchMonitoringIntervalHours) {
		toSerialize["defaultBranchMonitoringIntervalHours"] = o.DefaultBranchMonitoringIntervalHours
	}
	if !IsNil(o.DefaultBranchMonitoringStartTime) {
		toSerialize["defaultBranchMonitoringStartTime"] = o.DefaultBranchMonitoringStartTime
	}
	if !IsNil(o.GitExecutable) {
		toSerialize["gitExecutable"] = o.GitExecutable
	}
	if !IsNil(o.GitImplementation) {
		toSerialize["gitImplementation"] = o.GitImplementation
	}
	if !IsNil(o.GitTimeoutSeconds) {
		toSerialize["gitTimeoutSeconds"] = o.GitTimeoutSeconds
	}
	if !IsNil(o.PrCommentPurgeWindow) {
		toSerialize["prCommentPurgeWindow"] = o.PrCommentPurgeWindow
	}
	if !IsNil(o.PrEventPurgeWindow) {
		toSerialize["prEventPurgeWindow"] = o.PrEventPurgeWindow
	}
	if !IsNil(o.PullRequestMonitoringIntervalSeconds) {
		toSerialize["pullRequestMonitoringIntervalSeconds"] = o.PullRequestMonitoringIntervalSeconds
	}
	if !IsNil(o.UseUsernameInRepositoryCloneUrl) {
		toSerialize["useUsernameInRepositoryCloneUrl"] = o.UseUsernameInRepositoryCloneUrl
	}
	return toSerialize, nil
}

type NullableApiSourceControlConfigurationDTO struct {
	value *ApiSourceControlConfigurationDTO
	isSet bool
}

func (v NullableApiSourceControlConfigurationDTO) Get() *ApiSourceControlConfigurationDTO {
	return v.value
}

func (v *NullableApiSourceControlConfigurationDTO) Set(val *ApiSourceControlConfigurationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSourceControlConfigurationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSourceControlConfigurationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSourceControlConfigurationDTO(val *ApiSourceControlConfigurationDTO) *NullableApiSourceControlConfigurationDTO {
	return &NullableApiSourceControlConfigurationDTO{value: val, isSet: true}
}

func (v NullableApiSourceControlConfigurationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSourceControlConfigurationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


