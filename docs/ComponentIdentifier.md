# ComponentIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Coordinates** | Pointer to **map[string]string** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 

## Methods

### NewComponentIdentifier

`func NewComponentIdentifier() *ComponentIdentifier`

NewComponentIdentifier instantiates a new ComponentIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentIdentifierWithDefaults

`func NewComponentIdentifierWithDefaults() *ComponentIdentifier`

NewComponentIdentifierWithDefaults instantiates a new ComponentIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCoordinates

`func (o *ComponentIdentifier) GetCoordinates() map[string]string`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *ComponentIdentifier) GetCoordinatesOk() (*map[string]string, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *ComponentIdentifier) SetCoordinates(v map[string]string)`

SetCoordinates sets Coordinates field to given value.

### HasCoordinates

`func (o *ComponentIdentifier) HasCoordinates() bool`

HasCoordinates returns a boolean if a field has been set.

### GetFormat

`func (o *ComponentIdentifier) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ComponentIdentifier) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ComponentIdentifier) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ComponentIdentifier) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


