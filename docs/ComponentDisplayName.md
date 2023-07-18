# ComponentDisplayName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Parts** | Pointer to [**[]ComponentDisplayNamePart**](ComponentDisplayNamePart.md) |  | [optional] 

## Methods

### NewComponentDisplayName

`func NewComponentDisplayName() *ComponentDisplayName`

NewComponentDisplayName instantiates a new ComponentDisplayName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentDisplayNameWithDefaults

`func NewComponentDisplayNameWithDefaults() *ComponentDisplayName`

NewComponentDisplayNameWithDefaults instantiates a new ComponentDisplayName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ComponentDisplayName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentDisplayName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentDisplayName) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ComponentDisplayName) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParts

`func (o *ComponentDisplayName) GetParts() []ComponentDisplayNamePart`

GetParts returns the Parts field if non-nil, zero value otherwise.

### GetPartsOk

`func (o *ComponentDisplayName) GetPartsOk() (*[]ComponentDisplayNamePart, bool)`

GetPartsOk returns a tuple with the Parts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParts

`func (o *ComponentDisplayName) SetParts(v []ComponentDisplayNamePart)`

SetParts sets Parts field to given value.

### HasParts

`func (o *ComponentDisplayName) HasParts() bool`

HasParts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


