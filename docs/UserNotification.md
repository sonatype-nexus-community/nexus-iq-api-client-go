# UserNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** |  | [optional] 
**StageIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUserNotification

`func NewUserNotification() *UserNotification`

NewUserNotification instantiates a new UserNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNotificationWithDefaults

`func NewUserNotificationWithDefaults() *UserNotification`

NewUserNotificationWithDefaults instantiates a new UserNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *UserNotification) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UserNotification) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UserNotification) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *UserNotification) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetStageIds

`func (o *UserNotification) GetStageIds() []string`

GetStageIds returns the StageIds field if non-nil, zero value otherwise.

### GetStageIdsOk

`func (o *UserNotification) GetStageIdsOk() (*[]string, bool)`

GetStageIdsOk returns a tuple with the StageIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageIds

`func (o *UserNotification) SetStageIds(v []string)`

SetStageIds sets StageIds field to given value.

### HasStageIds

`func (o *UserNotification) HasStageIds() bool`

HasStageIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


