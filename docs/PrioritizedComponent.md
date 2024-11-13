# PrioritizedComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**ComponentHash** | Pointer to **string** |  | [optional] 
**ComponentIdentifier** | Pointer to [**ComponentIdentifier**](ComponentIdentifier.md) |  | [optional] 
**DependencyType** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**HasFailActionOnComponent** | Pointer to **bool** |  | [optional] 
**HighestReachableThreat** | Pointer to **int32** |  | [optional] 
**HighestThreat** | Pointer to **int32** |  | [optional] 
**HighestThreatPolicyConstraintName** | Pointer to **string** |  | [optional] 
**HighestThreatPolicyName** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**RemediationType** | Pointer to **string** |  | [optional] 
**RemediationVersion** | Pointer to **string** |  | [optional] 
**SecurityReachable** | Pointer to **bool** |  | [optional] 

## Methods

### NewPrioritizedComponent

`func NewPrioritizedComponent() *PrioritizedComponent`

NewPrioritizedComponent instantiates a new PrioritizedComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrioritizedComponentWithDefaults

`func NewPrioritizedComponentWithDefaults() *PrioritizedComponent`

NewPrioritizedComponentWithDefaults instantiates a new PrioritizedComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PrioritizedComponent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PrioritizedComponent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PrioritizedComponent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *PrioritizedComponent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetComponentHash

`func (o *PrioritizedComponent) GetComponentHash() string`

GetComponentHash returns the ComponentHash field if non-nil, zero value otherwise.

### GetComponentHashOk

`func (o *PrioritizedComponent) GetComponentHashOk() (*string, bool)`

GetComponentHashOk returns a tuple with the ComponentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentHash

`func (o *PrioritizedComponent) SetComponentHash(v string)`

SetComponentHash sets ComponentHash field to given value.

### HasComponentHash

`func (o *PrioritizedComponent) HasComponentHash() bool`

HasComponentHash returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *PrioritizedComponent) GetComponentIdentifier() ComponentIdentifier`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *PrioritizedComponent) GetComponentIdentifierOk() (*ComponentIdentifier, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *PrioritizedComponent) SetComponentIdentifier(v ComponentIdentifier)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *PrioritizedComponent) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetDependencyType

`func (o *PrioritizedComponent) GetDependencyType() string`

GetDependencyType returns the DependencyType field if non-nil, zero value otherwise.

### GetDependencyTypeOk

`func (o *PrioritizedComponent) GetDependencyTypeOk() (*string, bool)`

GetDependencyTypeOk returns a tuple with the DependencyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyType

`func (o *PrioritizedComponent) SetDependencyType(v string)`

SetDependencyType sets DependencyType field to given value.

### HasDependencyType

`func (o *PrioritizedComponent) HasDependencyType() bool`

HasDependencyType returns a boolean if a field has been set.

### GetDisplayName

`func (o *PrioritizedComponent) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PrioritizedComponent) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PrioritizedComponent) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PrioritizedComponent) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHasFailActionOnComponent

`func (o *PrioritizedComponent) GetHasFailActionOnComponent() bool`

GetHasFailActionOnComponent returns the HasFailActionOnComponent field if non-nil, zero value otherwise.

### GetHasFailActionOnComponentOk

`func (o *PrioritizedComponent) GetHasFailActionOnComponentOk() (*bool, bool)`

GetHasFailActionOnComponentOk returns a tuple with the HasFailActionOnComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFailActionOnComponent

`func (o *PrioritizedComponent) SetHasFailActionOnComponent(v bool)`

SetHasFailActionOnComponent sets HasFailActionOnComponent field to given value.

### HasHasFailActionOnComponent

`func (o *PrioritizedComponent) HasHasFailActionOnComponent() bool`

HasHasFailActionOnComponent returns a boolean if a field has been set.

### GetHighestReachableThreat

`func (o *PrioritizedComponent) GetHighestReachableThreat() int32`

GetHighestReachableThreat returns the HighestReachableThreat field if non-nil, zero value otherwise.

### GetHighestReachableThreatOk

`func (o *PrioritizedComponent) GetHighestReachableThreatOk() (*int32, bool)`

GetHighestReachableThreatOk returns a tuple with the HighestReachableThreat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestReachableThreat

`func (o *PrioritizedComponent) SetHighestReachableThreat(v int32)`

SetHighestReachableThreat sets HighestReachableThreat field to given value.

### HasHighestReachableThreat

`func (o *PrioritizedComponent) HasHighestReachableThreat() bool`

HasHighestReachableThreat returns a boolean if a field has been set.

### GetHighestThreat

`func (o *PrioritizedComponent) GetHighestThreat() int32`

GetHighestThreat returns the HighestThreat field if non-nil, zero value otherwise.

### GetHighestThreatOk

`func (o *PrioritizedComponent) GetHighestThreatOk() (*int32, bool)`

GetHighestThreatOk returns a tuple with the HighestThreat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestThreat

`func (o *PrioritizedComponent) SetHighestThreat(v int32)`

SetHighestThreat sets HighestThreat field to given value.

### HasHighestThreat

`func (o *PrioritizedComponent) HasHighestThreat() bool`

HasHighestThreat returns a boolean if a field has been set.

### GetHighestThreatPolicyConstraintName

`func (o *PrioritizedComponent) GetHighestThreatPolicyConstraintName() string`

GetHighestThreatPolicyConstraintName returns the HighestThreatPolicyConstraintName field if non-nil, zero value otherwise.

### GetHighestThreatPolicyConstraintNameOk

`func (o *PrioritizedComponent) GetHighestThreatPolicyConstraintNameOk() (*string, bool)`

GetHighestThreatPolicyConstraintNameOk returns a tuple with the HighestThreatPolicyConstraintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestThreatPolicyConstraintName

`func (o *PrioritizedComponent) SetHighestThreatPolicyConstraintName(v string)`

SetHighestThreatPolicyConstraintName sets HighestThreatPolicyConstraintName field to given value.

### HasHighestThreatPolicyConstraintName

`func (o *PrioritizedComponent) HasHighestThreatPolicyConstraintName() bool`

HasHighestThreatPolicyConstraintName returns a boolean if a field has been set.

### GetHighestThreatPolicyName

`func (o *PrioritizedComponent) GetHighestThreatPolicyName() string`

GetHighestThreatPolicyName returns the HighestThreatPolicyName field if non-nil, zero value otherwise.

### GetHighestThreatPolicyNameOk

`func (o *PrioritizedComponent) GetHighestThreatPolicyNameOk() (*string, bool)`

GetHighestThreatPolicyNameOk returns a tuple with the HighestThreatPolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestThreatPolicyName

`func (o *PrioritizedComponent) SetHighestThreatPolicyName(v string)`

SetHighestThreatPolicyName sets HighestThreatPolicyName field to given value.

### HasHighestThreatPolicyName

`func (o *PrioritizedComponent) HasHighestThreatPolicyName() bool`

HasHighestThreatPolicyName returns a boolean if a field has been set.

### GetPriority

`func (o *PrioritizedComponent) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *PrioritizedComponent) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *PrioritizedComponent) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *PrioritizedComponent) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetRemediationType

`func (o *PrioritizedComponent) GetRemediationType() string`

GetRemediationType returns the RemediationType field if non-nil, zero value otherwise.

### GetRemediationTypeOk

`func (o *PrioritizedComponent) GetRemediationTypeOk() (*string, bool)`

GetRemediationTypeOk returns a tuple with the RemediationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationType

`func (o *PrioritizedComponent) SetRemediationType(v string)`

SetRemediationType sets RemediationType field to given value.

### HasRemediationType

`func (o *PrioritizedComponent) HasRemediationType() bool`

HasRemediationType returns a boolean if a field has been set.

### GetRemediationVersion

`func (o *PrioritizedComponent) GetRemediationVersion() string`

GetRemediationVersion returns the RemediationVersion field if non-nil, zero value otherwise.

### GetRemediationVersionOk

`func (o *PrioritizedComponent) GetRemediationVersionOk() (*string, bool)`

GetRemediationVersionOk returns a tuple with the RemediationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationVersion

`func (o *PrioritizedComponent) SetRemediationVersion(v string)`

SetRemediationVersion sets RemediationVersion field to given value.

### HasRemediationVersion

`func (o *PrioritizedComponent) HasRemediationVersion() bool`

HasRemediationVersion returns a boolean if a field has been set.

### GetSecurityReachable

`func (o *PrioritizedComponent) GetSecurityReachable() bool`

GetSecurityReachable returns the SecurityReachable field if non-nil, zero value otherwise.

### GetSecurityReachableOk

`func (o *PrioritizedComponent) GetSecurityReachableOk() (*bool, bool)`

GetSecurityReachableOk returns a tuple with the SecurityReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityReachable

`func (o *PrioritizedComponent) SetSecurityReachable(v bool)`

SetSecurityReachable sets SecurityReachable field to given value.

### HasSecurityReachable

`func (o *PrioritizedComponent) HasSecurityReachable() bool`

HasSecurityReachable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


