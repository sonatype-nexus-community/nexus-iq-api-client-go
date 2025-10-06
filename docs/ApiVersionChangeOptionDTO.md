# ApiVersionChangeOptionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ApiComponentChangeActionDTO**](ApiComponentChangeActionDTO.md) |  | [optional] 
**DirectDependency** | Pointer to **bool** |  | [optional] 
**DirectDependencyData** | Pointer to [**[]ApiComponentChangeActionDTO**](ApiComponentChangeActionDTO.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewApiVersionChangeOptionDTO

`func NewApiVersionChangeOptionDTO() *ApiVersionChangeOptionDTO`

NewApiVersionChangeOptionDTO instantiates a new ApiVersionChangeOptionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiVersionChangeOptionDTOWithDefaults

`func NewApiVersionChangeOptionDTOWithDefaults() *ApiVersionChangeOptionDTO`

NewApiVersionChangeOptionDTOWithDefaults instantiates a new ApiVersionChangeOptionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApiVersionChangeOptionDTO) GetData() ApiComponentChangeActionDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiVersionChangeOptionDTO) GetDataOk() (*ApiComponentChangeActionDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiVersionChangeOptionDTO) SetData(v ApiComponentChangeActionDTO)`

SetData sets Data field to given value.

### HasData

`func (o *ApiVersionChangeOptionDTO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDirectDependency

`func (o *ApiVersionChangeOptionDTO) GetDirectDependency() bool`

GetDirectDependency returns the DirectDependency field if non-nil, zero value otherwise.

### GetDirectDependencyOk

`func (o *ApiVersionChangeOptionDTO) GetDirectDependencyOk() (*bool, bool)`

GetDirectDependencyOk returns a tuple with the DirectDependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDependency

`func (o *ApiVersionChangeOptionDTO) SetDirectDependency(v bool)`

SetDirectDependency sets DirectDependency field to given value.

### HasDirectDependency

`func (o *ApiVersionChangeOptionDTO) HasDirectDependency() bool`

HasDirectDependency returns a boolean if a field has been set.

### GetDirectDependencyData

`func (o *ApiVersionChangeOptionDTO) GetDirectDependencyData() []ApiComponentChangeActionDTO`

GetDirectDependencyData returns the DirectDependencyData field if non-nil, zero value otherwise.

### GetDirectDependencyDataOk

`func (o *ApiVersionChangeOptionDTO) GetDirectDependencyDataOk() (*[]ApiComponentChangeActionDTO, bool)`

GetDirectDependencyDataOk returns a tuple with the DirectDependencyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDependencyData

`func (o *ApiVersionChangeOptionDTO) SetDirectDependencyData(v []ApiComponentChangeActionDTO)`

SetDirectDependencyData sets DirectDependencyData field to given value.

### HasDirectDependencyData

`func (o *ApiVersionChangeOptionDTO) HasDirectDependencyData() bool`

HasDirectDependencyData returns a boolean if a field has been set.

### GetType

`func (o *ApiVersionChangeOptionDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiVersionChangeOptionDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiVersionChangeOptionDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiVersionChangeOptionDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


