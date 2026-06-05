# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionIndex** | Pointer to **int32** |  | [optional] 
**ConditionTypeId** | Pointer to **string** |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewCondition

`func NewCondition() *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionIndex

`func (o *Condition) GetConditionIndex() int32`

GetConditionIndex returns the ConditionIndex field if non-nil, zero value otherwise.

### GetConditionIndexOk

`func (o *Condition) GetConditionIndexOk() (*int32, bool)`

GetConditionIndexOk returns a tuple with the ConditionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIndex

`func (o *Condition) SetConditionIndex(v int32)`

SetConditionIndex sets ConditionIndex field to given value.

### HasConditionIndex

`func (o *Condition) HasConditionIndex() bool`

HasConditionIndex returns a boolean if a field has been set.

### GetConditionTypeId

`func (o *Condition) GetConditionTypeId() string`

GetConditionTypeId returns the ConditionTypeId field if non-nil, zero value otherwise.

### GetConditionTypeIdOk

`func (o *Condition) GetConditionTypeIdOk() (*string, bool)`

GetConditionTypeIdOk returns a tuple with the ConditionTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionTypeId

`func (o *Condition) SetConditionTypeId(v string)`

SetConditionTypeId sets ConditionTypeId field to given value.

### HasConditionTypeId

`func (o *Condition) HasConditionTypeId() bool`

HasConditionTypeId returns a boolean if a field has been set.

### GetOperator

`func (o *Condition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Condition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Condition) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Condition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *Condition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Condition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Condition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Condition) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


