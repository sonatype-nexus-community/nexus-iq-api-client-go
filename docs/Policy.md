# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to **map[string]string** |  | [optional] 
**Constraints** | Pointer to [**[]Constraint**](Constraint.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LegacyViolationAllowed** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Notifications** | Pointer to [**Notifications**](Notifications.md) |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**PolicyActionsOverrideAllowed** | Pointer to **bool** |  | [optional] 
**PolicyActionsOverrides** | Pointer to **map[string]map[string]string** |  | [optional] 
**PolicyNotificationsOverrideAllowed** | Pointer to **bool** |  | [optional] 
**PolicyNotificationsOverrides** | Pointer to [**map[string]Notifications**](Notifications.md) |  | [optional] 
**PolicyViolationGrandfatheringAllowed** | Pointer to **bool** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewPolicy

`func NewPolicy() *Policy`

NewPolicy instantiates a new Policy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyWithDefaults

`func NewPolicyWithDefaults() *Policy`

NewPolicyWithDefaults instantiates a new Policy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *Policy) GetActions() map[string]string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Policy) GetActionsOk() (*map[string]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Policy) SetActions(v map[string]string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *Policy) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetConstraints

`func (o *Policy) GetConstraints() []Constraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *Policy) GetConstraintsOk() (*[]Constraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *Policy) SetConstraints(v []Constraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *Policy) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetId

`func (o *Policy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Policy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Policy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Policy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLegacyViolationAllowed

`func (o *Policy) GetLegacyViolationAllowed() bool`

GetLegacyViolationAllowed returns the LegacyViolationAllowed field if non-nil, zero value otherwise.

### GetLegacyViolationAllowedOk

`func (o *Policy) GetLegacyViolationAllowedOk() (*bool, bool)`

GetLegacyViolationAllowedOk returns a tuple with the LegacyViolationAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyViolationAllowed

`func (o *Policy) SetLegacyViolationAllowed(v bool)`

SetLegacyViolationAllowed sets LegacyViolationAllowed field to given value.

### HasLegacyViolationAllowed

`func (o *Policy) HasLegacyViolationAllowed() bool`

HasLegacyViolationAllowed returns a boolean if a field has been set.

### GetName

`func (o *Policy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Policy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Policy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Policy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifications

`func (o *Policy) GetNotifications() Notifications`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *Policy) GetNotificationsOk() (*Notifications, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *Policy) SetNotifications(v Notifications)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *Policy) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetOwnerId

`func (o *Policy) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Policy) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Policy) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Policy) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPolicyActionsOverrideAllowed

`func (o *Policy) GetPolicyActionsOverrideAllowed() bool`

GetPolicyActionsOverrideAllowed returns the PolicyActionsOverrideAllowed field if non-nil, zero value otherwise.

### GetPolicyActionsOverrideAllowedOk

`func (o *Policy) GetPolicyActionsOverrideAllowedOk() (*bool, bool)`

GetPolicyActionsOverrideAllowedOk returns a tuple with the PolicyActionsOverrideAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyActionsOverrideAllowed

`func (o *Policy) SetPolicyActionsOverrideAllowed(v bool)`

SetPolicyActionsOverrideAllowed sets PolicyActionsOverrideAllowed field to given value.

### HasPolicyActionsOverrideAllowed

`func (o *Policy) HasPolicyActionsOverrideAllowed() bool`

HasPolicyActionsOverrideAllowed returns a boolean if a field has been set.

### GetPolicyActionsOverrides

`func (o *Policy) GetPolicyActionsOverrides() map[string]map[string]string`

GetPolicyActionsOverrides returns the PolicyActionsOverrides field if non-nil, zero value otherwise.

### GetPolicyActionsOverridesOk

`func (o *Policy) GetPolicyActionsOverridesOk() (*map[string]map[string]string, bool)`

GetPolicyActionsOverridesOk returns a tuple with the PolicyActionsOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyActionsOverrides

`func (o *Policy) SetPolicyActionsOverrides(v map[string]map[string]string)`

SetPolicyActionsOverrides sets PolicyActionsOverrides field to given value.

### HasPolicyActionsOverrides

`func (o *Policy) HasPolicyActionsOverrides() bool`

HasPolicyActionsOverrides returns a boolean if a field has been set.

### GetPolicyNotificationsOverrideAllowed

`func (o *Policy) GetPolicyNotificationsOverrideAllowed() bool`

GetPolicyNotificationsOverrideAllowed returns the PolicyNotificationsOverrideAllowed field if non-nil, zero value otherwise.

### GetPolicyNotificationsOverrideAllowedOk

`func (o *Policy) GetPolicyNotificationsOverrideAllowedOk() (*bool, bool)`

GetPolicyNotificationsOverrideAllowedOk returns a tuple with the PolicyNotificationsOverrideAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyNotificationsOverrideAllowed

`func (o *Policy) SetPolicyNotificationsOverrideAllowed(v bool)`

SetPolicyNotificationsOverrideAllowed sets PolicyNotificationsOverrideAllowed field to given value.

### HasPolicyNotificationsOverrideAllowed

`func (o *Policy) HasPolicyNotificationsOverrideAllowed() bool`

HasPolicyNotificationsOverrideAllowed returns a boolean if a field has been set.

### GetPolicyNotificationsOverrides

`func (o *Policy) GetPolicyNotificationsOverrides() map[string]Notifications`

GetPolicyNotificationsOverrides returns the PolicyNotificationsOverrides field if non-nil, zero value otherwise.

### GetPolicyNotificationsOverridesOk

`func (o *Policy) GetPolicyNotificationsOverridesOk() (*map[string]Notifications, bool)`

GetPolicyNotificationsOverridesOk returns a tuple with the PolicyNotificationsOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyNotificationsOverrides

`func (o *Policy) SetPolicyNotificationsOverrides(v map[string]Notifications)`

SetPolicyNotificationsOverrides sets PolicyNotificationsOverrides field to given value.

### HasPolicyNotificationsOverrides

`func (o *Policy) HasPolicyNotificationsOverrides() bool`

HasPolicyNotificationsOverrides returns a boolean if a field has been set.

### GetPolicyViolationGrandfatheringAllowed

`func (o *Policy) GetPolicyViolationGrandfatheringAllowed() bool`

GetPolicyViolationGrandfatheringAllowed returns the PolicyViolationGrandfatheringAllowed field if non-nil, zero value otherwise.

### GetPolicyViolationGrandfatheringAllowedOk

`func (o *Policy) GetPolicyViolationGrandfatheringAllowedOk() (*bool, bool)`

GetPolicyViolationGrandfatheringAllowedOk returns a tuple with the PolicyViolationGrandfatheringAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationGrandfatheringAllowed

`func (o *Policy) SetPolicyViolationGrandfatheringAllowed(v bool)`

SetPolicyViolationGrandfatheringAllowed sets PolicyViolationGrandfatheringAllowed field to given value.

### HasPolicyViolationGrandfatheringAllowed

`func (o *Policy) HasPolicyViolationGrandfatheringAllowed() bool`

HasPolicyViolationGrandfatheringAllowed returns a boolean if a field has been set.

### GetThreatLevel

`func (o *Policy) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *Policy) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *Policy) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *Policy) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


