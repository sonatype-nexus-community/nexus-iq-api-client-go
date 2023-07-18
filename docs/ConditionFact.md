# ConditionFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionIndex** | Pointer to **int32** |  | [optional] 
**ConditionTypeId** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to [**TriggerReference**](TriggerReference.md) |  | [optional] 
**Summary** | Pointer to **string** |  | [optional] 
**TriggerJson** | Pointer to **string** |  | [optional] 

## Methods

### NewConditionFact

`func NewConditionFact() *ConditionFact`

NewConditionFact instantiates a new ConditionFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionFactWithDefaults

`func NewConditionFactWithDefaults() *ConditionFact`

NewConditionFactWithDefaults instantiates a new ConditionFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionIndex

`func (o *ConditionFact) GetConditionIndex() int32`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *ConditionFact) GetConditionIndexOk() (*int32, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIndex

`func (o *ConditionFact) SetConditionIndex(v int32)`

SetConditionIndex sets ConditionIndex field to given value.

### HasConditionIndex

`func (o *ConditionFact) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.

### GetConditionTypeId

`func (o *ConditionFact) GetConditionTypeId() string`

GetConditionTypeId returns the ConditionTypeId field if non-nil, zero value otherwise.

### GetConditionTypeIdOk

`func (o *ConditionFact) GetConditionTypeIdOk() (*string, bool)`

GetConditionTypeIdOk returns a tuple with the ConditionTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionTypeId

`func (o *ConditionFact) SetConditionTypeId(v string)`

SetConditionTypeId sets ConditionTypeId field to given value.

### HasConditionTypeId

`func (o *ConditionFact) HasConditionTypeId() bool`

HasConditionTypeId returns a boolean if a field has been set.

### GetReason

`func (o *ConditionFact) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ConditionFact) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ConditionFact) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ConditionFact) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReference

`func (o *ConditionFact) GetReference() TriggerReference`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ConditionFact) GetReferenceOk() (*TriggerReference, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ConditionFact) SetReference(v TriggerReference)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ConditionFact) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSummary

`func (o *ConditionFact) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ConditionFact) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ConditionFact) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ConditionFact) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTriggerJson

`func (o *ConditionFact) GetTriggerJson() string`

GetTriggerJson returns the TriggerJson field if non-nil, zero value otherwise.

### GetTriggerJsonOk

`func (o *ConditionFact) GetTriggerJsonOk() (*string, bool)`

GetTriggerJsonOk returns a tuple with the TriggerJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerJson

`func (o *ConditionFact) SetTriggerJson(v string)`

SetTriggerJson sets TriggerJson field to given value.

### HasTriggerJson

`func (o *ConditionFact) HasTriggerJson() bool`

HasTriggerJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


