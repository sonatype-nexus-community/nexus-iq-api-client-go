# JiraNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueTypeId** | Pointer to **int64** |  | [optional] 
**ProjectKey** | Pointer to **string** |  | [optional] 
**StageIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewJiraNotification

`func NewJiraNotification() *JiraNotification`

NewJiraNotification instantiates a new JiraNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraNotificationWithDefaults

`func NewJiraNotificationWithDefaults() *JiraNotification`

NewJiraNotificationWithDefaults instantiates a new JiraNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueTypeId

`func (o *JiraNotification) GetIssueTypeId() int64`

GetIssueTypeId returns the IssueTypeId field if non-nil, zero value otherwise.

### GetIssueTypeIdOk

`func (o *JiraNotification) GetIssueTypeIdOk() (*int64, bool)`

GetIssueTypeIdOk returns a tuple with the IssueTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeId

`func (o *JiraNotification) SetIssueTypeId(v int64)`

SetIssueTypeId sets IssueTypeId field to given value.

### HasIssueTypeId

`func (o *JiraNotification) HasIssueTypeId() bool`

HasIssueTypeId returns a boolean if a field has been set.

### GetProjectKey

`func (o *JiraNotification) GetProjectKey() string`

GetProjectKey returns the ProjectKey field if non-nil, zero value otherwise.

### GetProjectKeyOk

`func (o *JiraNotification) GetProjectKeyOk() (*string, bool)`

GetProjectKeyOk returns a tuple with the ProjectKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectKey

`func (o *JiraNotification) SetProjectKey(v string)`

SetProjectKey sets ProjectKey field to given value.

### HasProjectKey

`func (o *JiraNotification) HasProjectKey() bool`

HasProjectKey returns a boolean if a field has been set.

### GetStageIds

`func (o *JiraNotification) GetStageIds() []string`

GetStageIds returns the StageIds field if non-nil, zero value otherwise.

### GetStageIdsOk

`func (o *JiraNotification) GetStageIdsOk() (*[]string, bool)`

GetStageIdsOk returns a tuple with the StageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageIds

`func (o *JiraNotification) SetStageIds(v []string)`

SetStageIds sets StageIds field to given value.

### HasStageIds

`func (o *JiraNotification) HasStageIds() bool`

HasStageIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


