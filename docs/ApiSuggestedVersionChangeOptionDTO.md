# ApiSuggestedVersionChangeOptionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ApiComponentChangeActionDTO**](ApiComponentChangeActionDTO.md) |  | [optional] 
**DirectDependency** | Pointer to **bool** |  | [optional] 
**DirectDependencyData** | Pointer to [**[]ApiComponentChangeActionDTO**](ApiComponentChangeActionDTO.md) |  | [optional] 
**IsGolden** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewApiSuggestedVersionChangeOptionDTO

`func NewApiSuggestedVersionChangeOptionDTO() *ApiSuggestedVersionChangeOptionDTO`

NewApiSuggestedVersionChangeOptionDTO instantiates a new ApiSuggestedVersionChangeOptionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSuggestedVersionChangeOptionDTOWithDefaults

`func NewApiSuggestedVersionChangeOptionDTOWithDefaults() *ApiSuggestedVersionChangeOptionDTO`

NewApiSuggestedVersionChangeOptionDTOWithDefaults instantiates a new ApiSuggestedVersionChangeOptionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApiSuggestedVersionChangeOptionDTO) GetData() ApiComponentChangeActionDTO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiSuggestedVersionChangeOptionDTO) GetDataOk() (*ApiComponentChangeActionDTO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiSuggestedVersionChangeOptionDTO) SetData(v ApiComponentChangeActionDTO)`

SetData sets Data field to given value.

### HasData

`func (o *ApiSuggestedVersionChangeOptionDTO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDirectDependency

`func (o *ApiSuggestedVersionChangeOptionDTO) GetDirectDependency() bool`

GetDirectDependency returns the DirectDependency field if non-nil, zero value otherwise.

### GetDirectDependencyOk

`func (o *ApiSuggestedVersionChangeOptionDTO) GetDirectDependencyOk() (*bool, bool)`

GetDirectDependencyOk returns a tuple with the DirectDependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDependency

`func (o *ApiSuggestedVersionChangeOptionDTO) SetDirectDependency(v bool)`

SetDirectDependency sets DirectDependency field to given value.

### HasDirectDependency

`func (o *ApiSuggestedVersionChangeOptionDTO) HasDirectDependency() bool`

HasDirectDependency returns a boolean if a field has been set.

### GetDirectDependencyData

`func (o *ApiSuggestedVersionChangeOptionDTO) GetDirectDependencyData() []ApiComponentChangeActionDTO`

GetDirectDependencyData returns the DirectDependencyData field if non-nil, zero value otherwise.

### GetDirectDependencyDataOk

`func (o *ApiSuggestedVersionChangeOptionDTO) GetDirectDependencyDataOk() (*[]ApiComponentChangeActionDTO, bool)`

GetDirectDependencyDataOk returns a tuple with the DirectDependencyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDependencyData

`func (o *ApiSuggestedVersionChangeOptionDTO) SetDirectDependencyData(v []ApiComponentChangeActionDTO)`

SetDirectDependencyData sets DirectDependencyData field to given value.

### HasDirectDependencyData

`func (o *ApiSuggestedVersionChangeOptionDTO) HasDirectDependencyData() bool`

HasDirectDependencyData returns a boolean if a field has been set.

### GetIsGolden

`func (o *ApiSuggestedVersionChangeOptionDTO) GetIsGolden() bool`

GetIsGolden returns the IsGolden field if non-nil, zero value otherwise.

### GetIsGoldenOk

`func (o *ApiSuggestedVersionChangeOptionDTO) GetIsGoldenOk() (*bool, bool)`

GetIsGoldenOk returns a tuple with the IsGolden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGolden

`func (o *ApiSuggestedVersionChangeOptionDTO) SetIsGolden(v bool)`

SetIsGolden sets IsGolden field to given value.

### HasIsGolden

`func (o *ApiSuggestedVersionChangeOptionDTO) HasIsGolden() bool`

HasIsGolden returns a boolean if a field has been set.

### GetType

`func (o *ApiSuggestedVersionChangeOptionDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiSuggestedVersionChangeOptionDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiSuggestedVersionChangeOptionDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiSuggestedVersionChangeOptionDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


