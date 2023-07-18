# PolicyAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**[]Action**](Action.md) |  | [optional] 
**Trigger** | Pointer to [**PolicyFact**](PolicyFact.md) |  | [optional] 

## Methods

### NewPolicyAlert

`func NewPolicyAlert() *PolicyAlert`

NewPolicyAlert instantiates a new PolicyAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyAlertWithDefaults

`func NewPolicyAlertWithDefaults() *PolicyAlert`

NewPolicyAlertWithDefaults instantiates a new PolicyAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *PolicyAlert) GetActions() []Action`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *PolicyAlert) GetActionsOk() (*[]Action, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *PolicyAlert) SetActions(v []Action)`

SetActions sets Actions field to given value.

### HasActions

`func (o *PolicyAlert) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetTrigger

`func (o *PolicyAlert) GetTrigger() PolicyFact`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *PolicyAlert) GetTriggerOk() (*PolicyFact, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *PolicyAlert) SetTrigger(v PolicyFact)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *PolicyAlert) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


