# ApiProxyServerConfigurationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludeHosts** | Pointer to **[]string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **[]string** |  | [optional] 
**PasswordIsIncluded** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewApiProxyServerConfigurationDTO

`func NewApiProxyServerConfigurationDTO() *ApiProxyServerConfigurationDTO`

NewApiProxyServerConfigurationDTO instantiates a new ApiProxyServerConfigurationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiProxyServerConfigurationDTOWithDefaults

`func NewApiProxyServerConfigurationDTOWithDefaults() *ApiProxyServerConfigurationDTO`

NewApiProxyServerConfigurationDTOWithDefaults instantiates a new ApiProxyServerConfigurationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludeHosts

`func (o *ApiProxyServerConfigurationDTO) GetExcludeHosts() []string`

GetExcludeHosts returns the ExcludeHosts field if non-nil, zero value otherwise.

### GetExcludeHostsOk

`func (o *ApiProxyServerConfigurationDTO) GetExcludeHostsOk() (*[]string, bool)`

GetExcludeHostsOk returns a tuple with the ExcludeHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeHosts

`func (o *ApiProxyServerConfigurationDTO) SetExcludeHosts(v []string)`

SetExcludeHosts sets ExcludeHosts field to given value.

### HasExcludeHosts

`func (o *ApiProxyServerConfigurationDTO) HasExcludeHosts() bool`

HasExcludeHosts returns a boolean if a field has been set.

### GetHostname

`func (o *ApiProxyServerConfigurationDTO) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApiProxyServerConfigurationDTO) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApiProxyServerConfigurationDTO) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApiProxyServerConfigurationDTO) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPassword

`func (o *ApiProxyServerConfigurationDTO) GetPassword() []string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiProxyServerConfigurationDTO) GetPasswordOk() (*[]string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiProxyServerConfigurationDTO) SetPassword(v []string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiProxyServerConfigurationDTO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordIsIncluded

`func (o *ApiProxyServerConfigurationDTO) GetPasswordIsIncluded() bool`

GetPasswordIsIncluded returns the PasswordIsIncluded field if non-nil, zero value otherwise.

### GetPasswordIsIncludedOk

`func (o *ApiProxyServerConfigurationDTO) GetPasswordIsIncludedOk() (*bool, bool)`

GetPasswordIsIncludedOk returns a tuple with the PasswordIsIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordIsIncluded

`func (o *ApiProxyServerConfigurationDTO) SetPasswordIsIncluded(v bool)`

SetPasswordIsIncluded sets PasswordIsIncluded field to given value.

### HasPasswordIsIncluded

`func (o *ApiProxyServerConfigurationDTO) HasPasswordIsIncluded() bool`

HasPasswordIsIncluded returns a boolean if a field has been set.

### GetPort

`func (o *ApiProxyServerConfigurationDTO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApiProxyServerConfigurationDTO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApiProxyServerConfigurationDTO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApiProxyServerConfigurationDTO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *ApiProxyServerConfigurationDTO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiProxyServerConfigurationDTO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiProxyServerConfigurationDTO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiProxyServerConfigurationDTO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


