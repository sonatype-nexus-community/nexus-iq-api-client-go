# AxisValueDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **[]string** |  | [optional] 
**Effective** | Pointer to **[]string** |  | [optional] 
**InheritedFrom** | Pointer to [**OwnerRefDTO**](OwnerRefDTO.md) |  | [optional] 

## Methods

### NewAxisValueDTO

`func NewAxisValueDTO() *AxisValueDTO`

NewAxisValueDTO instantiates a new AxisValueDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAxisValueDTOWithDefaults

`func NewAxisValueDTOWithDefaults() *AxisValueDTO`

NewAxisValueDTOWithDefaults instantiates a new AxisValueDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *AxisValueDTO) GetConfigured() []string`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *AxisValueDTO) GetConfiguredOk() (*[]string, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *AxisValueDTO) SetConfigured(v []string)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *AxisValueDTO) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetEffective

`func (o *AxisValueDTO) GetEffective() []string`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *AxisValueDTO) GetEffectiveOk() (*[]string, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *AxisValueDTO) SetEffective(v []string)`

SetEffective sets Effective field to given value.

### HasEffective

`func (o *AxisValueDTO) HasEffective() bool`

HasEffective returns a boolean if a field has been set.

### GetInheritedFrom

`func (o *AxisValueDTO) GetInheritedFrom() OwnerRefDTO`

GetInheritedFrom returns the InheritedFrom field if non-nil, zero value otherwise.

### GetInheritedFromOk

`func (o *AxisValueDTO) GetInheritedFromOk() (*OwnerRefDTO, bool)`

GetInheritedFromOk returns a tuple with the InheritedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFrom

`func (o *AxisValueDTO) SetInheritedFrom(v OwnerRefDTO)`

SetInheritedFrom sets InheritedFrom field to given value.

### HasInheritedFrom

`func (o *AxisValueDTO) HasInheritedFrom() bool`

HasInheritedFrom returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


