# ComponentFact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentIdentifier** | Pointer to [**ComponentIdentifier**](ComponentIdentifier.md) |  | [optional] 
**ConstraintFacts** | Pointer to [**[]ConstraintFact**](ConstraintFact.md) |  | [optional] 
**DisplayName** | Pointer to [**ComponentDisplayName**](ComponentDisplayName.md) |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**Pathnames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewComponentFact

`func NewComponentFact() *ComponentFact`

NewComponentFact instantiates a new ComponentFact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentFactWithDefaults

`func NewComponentFactWithDefaults() *ComponentFact`

NewComponentFactWithDefaults instantiates a new ComponentFact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentIdentifier

`func (o *ComponentFact) GetComponentIdentifier() ComponentIdentifier`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ComponentFact) GetComponentIdentifierOk() (*ComponentIdentifier, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ComponentFact) SetComponentIdentifier(v ComponentIdentifier)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ComponentFact) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetConstraintFacts

`func (o *ComponentFact) GetConstraintFacts() []ConstraintFact`

GetConstraintFacts returns the ConstraintFacts field if non-nil, zero value otherwise.

### GetConstraintFactsOk

`func (o *ComponentFact) GetConstraintFactsOk() (*[]ConstraintFact, bool)`

GetConstraintFactsOk returns a tuple with the ConstraintFacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintFacts

`func (o *ComponentFact) SetConstraintFacts(v []ConstraintFact)`

SetConstraintFacts sets ConstraintFacts field to given value.

### HasConstraintFacts

`func (o *ComponentFact) HasConstraintFacts() bool`

HasConstraintFacts returns a boolean if a field has been set.

### GetDisplayName

`func (o *ComponentFact) GetDisplayName() ComponentDisplayName`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ComponentFact) GetDisplayNameOk() (*ComponentDisplayName, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ComponentFact) SetDisplayName(v ComponentDisplayName)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ComponentFact) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHash

`func (o *ComponentFact) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ComponentFact) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ComponentFact) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ComponentFact) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetPathnames

`func (o *ComponentFact) GetPathnames() []string`

GetPathnames returns the Pathnames field if non-nil, zero value otherwise.

### GetPathnamesOk

`func (o *ComponentFact) GetPathnamesOk() (*[]string, bool)`

GetPathnamesOk returns a tuple with the Pathnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathnames

`func (o *ComponentFact) SetPathnames(v []string)`

SetPathnames sets Pathnames field to given value.

### HasPathnames

`func (o *ComponentFact) HasPathnames() bool`

HasPathnames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


