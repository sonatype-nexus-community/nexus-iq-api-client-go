# ApiComponentChangeActionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to [**ApiComponentDTOV2**](ApiComponentDTOV2.md) |  | [optional] 

## Methods

### NewApiComponentChangeActionDTO

`func NewApiComponentChangeActionDTO() *ApiComponentChangeActionDTO`

NewApiComponentChangeActionDTO instantiates a new ApiComponentChangeActionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiComponentChangeActionDTOWithDefaults

`func NewApiComponentChangeActionDTOWithDefaults() *ApiComponentChangeActionDTO`

NewApiComponentChangeActionDTOWithDefaults instantiates a new ApiComponentChangeActionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ApiComponentChangeActionDTO) GetComponent() ApiComponentDTOV2`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ApiComponentChangeActionDTO) GetComponentOk() (*ApiComponentDTOV2, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ApiComponentChangeActionDTO) SetComponent(v ApiComponentDTOV2)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ApiComponentChangeActionDTO) HasComponent() bool`

HasComponent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


