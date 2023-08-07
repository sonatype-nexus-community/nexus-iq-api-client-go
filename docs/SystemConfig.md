# SystemConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | Pointer to **NullableString** |  | [optional] 
**ForceBaseUrl** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewSystemConfig

`func NewSystemConfig() *SystemConfig`

NewSystemConfig instantiates a new SystemConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemConfigWithDefaults

`func NewSystemConfigWithDefaults() *SystemConfig`

NewSystemConfigWithDefaults instantiates a new SystemConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *SystemConfig) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *SystemConfig) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *SystemConfig) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.

### HasBaseUrl

`func (o *SystemConfig) HasBaseUrl() bool`

HasBaseUrl returns a boolean if a field has been set.

### SetBaseUrlNil

`func (o *SystemConfig) SetBaseUrlNil(b bool)`

 SetBaseUrlNil sets the value for BaseUrl to be an explicit nil

### UnsetBaseUrl
`func (o *SystemConfig) UnsetBaseUrl()`

UnsetBaseUrl ensures that no value is present for BaseUrl, not even an explicit nil
### GetForceBaseUrl

`func (o *SystemConfig) GetForceBaseUrl() bool`

GetForceBaseUrl returns the ForceBaseUrl field if non-nil, zero value otherwise.

### GetForceBaseUrlOk

`func (o *SystemConfig) GetForceBaseUrlOk() (*bool, bool)`

GetForceBaseUrlOk returns a tuple with the ForceBaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceBaseUrl

`func (o *SystemConfig) SetForceBaseUrl(v bool)`

SetForceBaseUrl sets ForceBaseUrl field to given value.

### HasForceBaseUrl

`func (o *SystemConfig) HasForceBaseUrl() bool`

HasForceBaseUrl returns a boolean if a field has been set.

### SetForceBaseUrlNil

`func (o *SystemConfig) SetForceBaseUrlNil(b bool)`

 SetForceBaseUrlNil sets the value for ForceBaseUrl to be an explicit nil

### UnsetForceBaseUrl
`func (o *SystemConfig) UnsetForceBaseUrl()`

UnsetForceBaseUrl ensures that no value is present for ForceBaseUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


