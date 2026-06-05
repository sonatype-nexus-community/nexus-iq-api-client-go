# Notifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JiraNotifications** | Pointer to [**[]JiraNotification**](JiraNotification.md) |  | [optional] 
**RoleNotifications** | Pointer to [**[]RoleNotification**](RoleNotification.md) |  | [optional] 
**UserNotifications** | Pointer to [**[]UserNotification**](UserNotification.md) |  | [optional] 
**WebhookNotifications** | Pointer to [**[]WebhookNotification**](WebhookNotification.md) |  | [optional] 

## Methods

### NewNotifications

`func NewNotifications() *Notifications`

NewNotifications instantiates a new Notifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsWithDefaults

`func NewNotificationsWithDefaults() *Notifications`

NewNotificationsWithDefaults instantiates a new Notifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJiraNotifications

`func (o *Notifications) GetJiraNotifications() []JiraNotification`

GetJiraNotifications returns the JiraNotifications field if non-nil, zero value otherwise.

### GetJiraNotificationsOk

`func (o *Notifications) GetJiraNotificationsOk() (*[]JiraNotification, bool)`

GetJiraNotificationsOk returns a tuple with the JiraNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraNotifications

`func (o *Notifications) SetJiraNotifications(v []JiraNotification)`

SetJiraNotifications sets JiraNotifications field to given value.

### HasJiraNotifications

`func (o *Notifications) HasJiraNotifications() bool`

HasJiraNotifications returns a boolean if a field has been set.

### GetRoleNotifications

`func (o *Notifications) GetRoleNotifications() []RoleNotification`

GetRoleNotifications returns the RoleNotifications field if non-nil, zero value otherwise.

### GetRoleNotificationsOk

`func (o *Notifications) GetRoleNotificationsOk() (*[]RoleNotification, bool)`

GetRoleNotificationsOk returns a tuple with the RoleNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleNotifications

`func (o *Notifications) SetRoleNotifications(v []RoleNotification)`

SetRoleNotifications sets RoleNotifications field to given value.

### HasRoleNotifications

`func (o *Notifications) HasRoleNotifications() bool`

HasRoleNotifications returns a boolean if a field has been set.

### GetUserNotifications

`func (o *Notifications) GetUserNotifications() []UserNotification`

GetUserNotifications returns the UserNotifications field if non-nil, zero value otherwise.

### GetUserNotificationsOk

`func (o *Notifications) GetUserNotificationsOk() (*[]UserNotification, bool)`

GetUserNotificationsOk returns a tuple with the UserNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNotifications

`func (o *Notifications) SetUserNotifications(v []UserNotification)`

SetUserNotifications sets UserNotifications field to given value.

### HasUserNotifications

`func (o *Notifications) HasUserNotifications() bool`

HasUserNotifications returns a boolean if a field has been set.

### GetWebhookNotifications

`func (o *Notifications) GetWebhookNotifications() []WebhookNotification`

GetWebhookNotifications returns the WebhookNotifications field if non-nil, zero value otherwise.

### GetWebhookNotificationsOk

`func (o *Notifications) GetWebhookNotificationsOk() (*[]WebhookNotification, bool)`

GetWebhookNotificationsOk returns a tuple with the WebhookNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookNotifications

`func (o *Notifications) SetWebhookNotifications(v []WebhookNotification)`

SetWebhookNotifications sets WebhookNotifications field to given value.

### HasWebhookNotifications

`func (o *Notifications) HasWebhookNotifications() bool`

HasWebhookNotifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


