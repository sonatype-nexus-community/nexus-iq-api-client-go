# PolicyFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentFacts** | Pointer to [**[]ComponentFact**](ComponentFact.md) |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewPolicyFact

`func NewPolicyFact() *PolicyFact`

NewPolicyFact instantiates a new PolicyFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyFactWithDefaults

`func NewPolicyFactWithDefaults() *PolicyFact`

NewPolicyFactWithDefaults instantiates a new PolicyFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentFacts

`func (o *PolicyFact) GetComponentFacts() []ComponentFact`

GetComponentFacts returns the ComponentFacts field if non-nil, zero value otherwise.

### GetComponentFactsOk

`func (o *PolicyFact) GetComponentFactsOk() (*[]ComponentFact, bool)`

GetComponentFactsOk returns a tuple with the ComponentFacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentFacts

`func (o *PolicyFact) SetComponentFacts(v []ComponentFact)`

SetComponentFacts sets ComponentFacts field to given value.

### HasComponentFacts

`func (o *PolicyFact) HasComponentFacts() bool`

HasComponentFacts returns a boolean if a field has been set.

### GetPolicyId

`func (o *PolicyFact) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyFact) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyFact) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *PolicyFact) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *PolicyFact) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *PolicyFact) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *PolicyFact) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *PolicyFact) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *PolicyFact) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *PolicyFact) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *PolicyFact) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *PolicyFact) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetThreatLevel

`func (o *PolicyFact) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *PolicyFact) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *PolicyFact) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *PolicyFact) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


