# ApiCiConfigurationResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**ApiCiConfigurationDto**](ApiCiConfigurationDto.md) |  | [optional] 
**Source** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewApiCiConfigurationResponseDto

`func NewApiCiConfigurationResponseDto() *ApiCiConfigurationResponseDto`

NewApiCiConfigurationResponseDto instantiates a new ApiCiConfigurationResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCiConfigurationResponseDtoWithDefaults

`func NewApiCiConfigurationResponseDtoWithDefaults() *ApiCiConfigurationResponseDto`

NewApiCiConfigurationResponseDtoWithDefaults instantiates a new ApiCiConfigurationResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApiCiConfigurationResponseDto) GetData() ApiCiConfigurationDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApiCiConfigurationResponseDto) GetDataOk() (*ApiCiConfigurationDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApiCiConfigurationResponseDto) SetData(v ApiCiConfigurationDto)`

SetData sets Data field to given value.

### HasData

`func (o *ApiCiConfigurationResponseDto) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSource

`func (o *ApiCiConfigurationResponseDto) GetSource() map[string]string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiCiConfigurationResponseDto) GetSourceOk() (*map[string]string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiCiConfigurationResponseDto) SetSource(v map[string]string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiCiConfigurationResponseDto) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


