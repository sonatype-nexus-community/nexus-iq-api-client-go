# Constraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]Condition**](Condition.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 

## Methods

### NewConstraint

`func NewConstraint() *Constraint`

NewConstraint instantiates a new Constraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConstraintWithDefaults

`func NewConstraintWithDefaults() *Constraint`

NewConstraintWithDefaults instantiates a new Constraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *Constraint) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Constraint) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Constraint) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Constraint) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetId

`func (o *Constraint) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Constraint) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Constraint) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Constraint) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Constraint) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Constraint) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Constraint) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Constraint) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperator

`func (o *Constraint) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Constraint) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Constraint) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Constraint) HasOperator() bool`

HasOperator returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


