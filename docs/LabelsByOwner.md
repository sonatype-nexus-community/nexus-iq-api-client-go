# LabelsByOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to [**[]ApiLabelDTO**](ApiLabelDTO.md) |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerName** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **string** |  | [optional] 

## Methods

### NewLabelsByOwner

`func NewLabelsByOwner() *LabelsByOwner`

NewLabelsByOwner instantiates a new LabelsByOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelsByOwnerWithDefaults

`func NewLabelsByOwnerWithDefaults() *LabelsByOwner`

NewLabelsByOwnerWithDefaults instantiates a new LabelsByOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *LabelsByOwner) GetLabels() []ApiLabelDTO`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *LabelsByOwner) GetLabelsOk() (*[]ApiLabelDTO, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *LabelsByOwner) SetLabels(v []ApiLabelDTO)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *LabelsByOwner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOwnerId

`func (o *LabelsByOwner) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *LabelsByOwner) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *LabelsByOwner) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *LabelsByOwner) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwnerName

`func (o *LabelsByOwner) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *LabelsByOwner) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *LabelsByOwner) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *LabelsByOwner) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetOwnerType

`func (o *LabelsByOwner) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *LabelsByOwner) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *LabelsByOwner) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *LabelsByOwner) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


