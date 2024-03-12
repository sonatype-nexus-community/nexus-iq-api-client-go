# Swid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Patch** | Pointer to **bool** |  | [optional] 
**TagId** | Pointer to **string** |  | [optional] 
**TagVersion** | Pointer to **int32** |  | [optional] 
**Text** | Pointer to [**AttachmentText**](AttachmentText.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewSwid

`func NewSwid() *Swid`

NewSwid instantiates a new Swid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwidWithDefaults

`func NewSwidWithDefaults() *Swid`

NewSwidWithDefaults instantiates a new Swid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Swid) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Swid) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Swid) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Swid) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPatch

`func (o *Swid) GetPatch() bool`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *Swid) GetPatchOk() (*bool, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *Swid) SetPatch(v bool)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *Swid) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetTagId

`func (o *Swid) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *Swid) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *Swid) SetTagId(v string)`

SetTagId sets TagId field to given value.

### HasTagId

`func (o *Swid) HasTagId() bool`

HasTagId returns a boolean if a field has been set.

### GetTagVersion

`func (o *Swid) GetTagVersion() int32`

GetTagVersion returns the TagVersion field if non-nil, zero value otherwise.

### GetTagVersionOk

`func (o *Swid) GetTagVersionOk() (*int32, bool)`

GetTagVersionOk returns a tuple with the TagVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagVersion

`func (o *Swid) SetTagVersion(v int32)`

SetTagVersion sets TagVersion field to given value.

### HasTagVersion

`func (o *Swid) HasTagVersion() bool`

HasTagVersion returns a boolean if a field has been set.

### GetText

`func (o *Swid) GetText() AttachmentText`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Swid) GetTextOk() (*AttachmentText, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Swid) SetText(v AttachmentText)`

SetText sets Text field to given value.

### HasText

`func (o *Swid) HasText() bool`

HasText returns a boolean if a field has been set.

### GetVersion

`func (o *Swid) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Swid) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Swid) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Swid) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


