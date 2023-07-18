# FormDataMultiPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BodyParts** | Pointer to [**[]BodyPart**](BodyPart.md) |  | [optional] 
**ContentDisposition** | Pointer to [**ContentDisposition**](ContentDisposition.md) |  | [optional] 
**Entity** | Pointer to **map[string]interface{}** |  | [optional] 
**Fields** | Pointer to [**map[string][]FormDataBodyPart**](array.md) |  | [optional] 
**Headers** | Pointer to **map[string][]string** |  | [optional] 
**MediaType** | Pointer to [**BodyPartMediaType**](BodyPartMediaType.md) |  | [optional] 
**MessageBodyWorkers** | Pointer to **map[string]interface{}** |  | [optional] 
**ParameterizedHeaders** | Pointer to [**map[string][]ParameterizedHeader**](array.md) |  | [optional] 
**Parent** | Pointer to [**MultiPart**](MultiPart.md) |  | [optional] 
**Providers** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewFormDataMultiPart

`func NewFormDataMultiPart() *FormDataMultiPart`

NewFormDataMultiPart instantiates a new FormDataMultiPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormDataMultiPartWithDefaults

`func NewFormDataMultiPartWithDefaults() *FormDataMultiPart`

NewFormDataMultiPartWithDefaults instantiates a new FormDataMultiPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBodyParts

`func (o *FormDataMultiPart) GetBodyParts() []BodyPart`

GetBodyParts returns the BodyParts field if non-nil, zero value otherwise.

### GetBodyPartsOk

`func (o *FormDataMultiPart) GetBodyPartsOk() (*[]BodyPart, bool)`

GetBodyPartsOk returns a tuple with the BodyParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyParts

`func (o *FormDataMultiPart) SetBodyParts(v []BodyPart)`

SetBodyParts sets BodyParts field to given value.

### HasBodyParts

`func (o *FormDataMultiPart) HasBodyParts() bool`

HasBodyParts returns a boolean if a field has been set.

### GetContentDisposition

`func (o *FormDataMultiPart) GetContentDisposition() ContentDisposition`

GetContentDisposition returns the ContentDisposition field if non-nil, zero value otherwise.

### GetContentDispositionOk

`func (o *FormDataMultiPart) GetContentDispositionOk() (*ContentDisposition, bool)`

GetContentDispositionOk returns a tuple with the ContentDisposition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentDisposition

`func (o *FormDataMultiPart) SetContentDisposition(v ContentDisposition)`

SetContentDisposition sets ContentDisposition field to given value.

### HasContentDisposition

`func (o *FormDataMultiPart) HasContentDisposition() bool`

HasContentDisposition returns a boolean if a field has been set.

### GetEntity

`func (o *FormDataMultiPart) GetEntity() map[string]interface{}`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *FormDataMultiPart) GetEntityOk() (*map[string]interface{}, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *FormDataMultiPart) SetEntity(v map[string]interface{})`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *FormDataMultiPart) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetFields

`func (o *FormDataMultiPart) GetFields() map[string][]FormDataBodyPart`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FormDataMultiPart) GetFieldsOk() (*map[string][]FormDataBodyPart, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FormDataMultiPart) SetFields(v map[string][]FormDataBodyPart)`

SetFields sets Fields field to given value.

### HasFields

`func (o *FormDataMultiPart) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHeaders

`func (o *FormDataMultiPart) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *FormDataMultiPart) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *FormDataMultiPart) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *FormDataMultiPart) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetMediaType

`func (o *FormDataMultiPart) GetMediaType() BodyPartMediaType`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *FormDataMultiPart) GetMediaTypeOk() (*BodyPartMediaType, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *FormDataMultiPart) SetMediaType(v BodyPartMediaType)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *FormDataMultiPart) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMessageBodyWorkers

`func (o *FormDataMultiPart) GetMessageBodyWorkers() map[string]interface{}`

GetMessageBodyWorkers returns the MessageBodyWorkers field if non-nil, zero value otherwise.

### GetMessageBodyWorkersOk

`func (o *FormDataMultiPart) GetMessageBodyWorkersOk() (*map[string]interface{}, bool)`

GetMessageBodyWorkersOk returns a tuple with the MessageBodyWorkers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBodyWorkers

`func (o *FormDataMultiPart) SetMessageBodyWorkers(v map[string]interface{})`

SetMessageBodyWorkers sets MessageBodyWorkers field to given value.

### HasMessageBodyWorkers

`func (o *FormDataMultiPart) HasMessageBodyWorkers() bool`

HasMessageBodyWorkers returns a boolean if a field has been set.

### GetParameterizedHeaders

`func (o *FormDataMultiPart) GetParameterizedHeaders() map[string][]ParameterizedHeader`

GetParameterizedHeaders returns the ParameterizedHeaders field if non-nil, zero value otherwise.

### GetParameterizedHeadersOk

`func (o *FormDataMultiPart) GetParameterizedHeadersOk() (*map[string][]ParameterizedHeader, bool)`

GetParameterizedHeadersOk returns a tuple with the ParameterizedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterizedHeaders

`func (o *FormDataMultiPart) SetParameterizedHeaders(v map[string][]ParameterizedHeader)`

SetParameterizedHeaders sets ParameterizedHeaders field to given value.

### HasParameterizedHeaders

`func (o *FormDataMultiPart) HasParameterizedHeaders() bool`

HasParameterizedHeaders returns a boolean if a field has been set.

### GetParent

`func (o *FormDataMultiPart) GetParent() MultiPart`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FormDataMultiPart) GetParentOk() (*MultiPart, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FormDataMultiPart) SetParent(v MultiPart)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FormDataMultiPart) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProviders

`func (o *FormDataMultiPart) GetProviders() map[string]interface{}`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *FormDataMultiPart) GetProvidersOk() (*map[string]interface{}, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *FormDataMultiPart) SetProviders(v map[string]interface{})`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *FormDataMultiPart) HasProviders() bool`

HasProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


