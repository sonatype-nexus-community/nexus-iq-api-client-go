# ApiSearchIndexSetupDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**Config**](Config.md) |  | [optional] 
**Connection** | Pointer to [**Connection**](Connection.md) |  | [optional] 
**Index** | Pointer to [**Index**](Index.md) |  | [optional] 
**NexusOneEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiSearchIndexSetupDTO

`func NewApiSearchIndexSetupDTO() *ApiSearchIndexSetupDTO`

NewApiSearchIndexSetupDTO instantiates a new ApiSearchIndexSetupDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchIndexSetupDTOWithDefaults

`func NewApiSearchIndexSetupDTOWithDefaults() *ApiSearchIndexSetupDTO`

NewApiSearchIndexSetupDTOWithDefaults instantiates a new ApiSearchIndexSetupDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ApiSearchIndexSetupDTO) GetConfig() Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ApiSearchIndexSetupDTO) GetConfigOk() (*Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ApiSearchIndexSetupDTO) SetConfig(v Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ApiSearchIndexSetupDTO) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConnection

`func (o *ApiSearchIndexSetupDTO) GetConnection() Connection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *ApiSearchIndexSetupDTO) GetConnectionOk() (*Connection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *ApiSearchIndexSetupDTO) SetConnection(v Connection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *ApiSearchIndexSetupDTO) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetIndex

`func (o *ApiSearchIndexSetupDTO) GetIndex() Index`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *ApiSearchIndexSetupDTO) GetIndexOk() (*Index, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *ApiSearchIndexSetupDTO) SetIndex(v Index)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *ApiSearchIndexSetupDTO) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetNexusOneEnabled

`func (o *ApiSearchIndexSetupDTO) GetNexusOneEnabled() bool`

GetNexusOneEnabled returns the NexusOneEnabled field if non-nil, zero value otherwise.

### GetNexusOneEnabledOk

`func (o *ApiSearchIndexSetupDTO) GetNexusOneEnabledOk() (*bool, bool)`

GetNexusOneEnabledOk returns a tuple with the NexusOneEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNexusOneEnabled

`func (o *ApiSearchIndexSetupDTO) SetNexusOneEnabled(v bool)`

SetNexusOneEnabled sets NexusOneEnabled field to given value.

### HasNexusOneEnabled

`func (o *ApiSearchIndexSetupDTO) HasNexusOneEnabled() bool`

HasNexusOneEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


