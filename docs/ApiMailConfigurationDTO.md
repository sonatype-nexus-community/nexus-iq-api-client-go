# ApiMailConfigurationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **[]string** |  | [optional] 
**PasswordIsIncluded** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**SslEnabled** | Pointer to **bool** |  | [optional] 
**StartTlsEnabled** | Pointer to **bool** |  | [optional] 
**SystemEmail** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewApiMailConfigurationDTO

`func NewApiMailConfigurationDTO() *ApiMailConfigurationDTO`

NewApiMailConfigurationDTO instantiates a new ApiMailConfigurationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMailConfigurationDTOWithDefaults

`func NewApiMailConfigurationDTOWithDefaults() *ApiMailConfigurationDTO`

NewApiMailConfigurationDTOWithDefaults instantiates a new ApiMailConfigurationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *ApiMailConfigurationDTO) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ApiMailConfigurationDTO) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ApiMailConfigurationDTO) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ApiMailConfigurationDTO) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPassword

`func (o *ApiMailConfigurationDTO) GetPassword() []string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ApiMailConfigurationDTO) GetPasswordOk() (*[]string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ApiMailConfigurationDTO) SetPassword(v []string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ApiMailConfigurationDTO) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordIsIncluded

`func (o *ApiMailConfigurationDTO) GetPasswordIsIncluded() bool`

GetPasswordIsIncluded returns the PasswordIsIncluded field if non-nil, zero value otherwise.

### GetPasswordIsIncludedOk

`func (o *ApiMailConfigurationDTO) GetPasswordIsIncludedOk() (*bool, bool)`

GetPasswordIsIncludedOk returns a tuple with the PasswordIsIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordIsIncluded

`func (o *ApiMailConfigurationDTO) SetPasswordIsIncluded(v bool)`

SetPasswordIsIncluded sets PasswordIsIncluded field to given value.

### HasPasswordIsIncluded

`func (o *ApiMailConfigurationDTO) HasPasswordIsIncluded() bool`

HasPasswordIsIncluded returns a boolean if a field has been set.

### GetPort

`func (o *ApiMailConfigurationDTO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApiMailConfigurationDTO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApiMailConfigurationDTO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApiMailConfigurationDTO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSslEnabled

`func (o *ApiMailConfigurationDTO) GetSslEnabled() bool`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *ApiMailConfigurationDTO) GetSslEnabledOk() (*bool, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *ApiMailConfigurationDTO) SetSslEnabled(v bool)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *ApiMailConfigurationDTO) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### GetStartTlsEnabled

`func (o *ApiMailConfigurationDTO) GetStartTlsEnabled() bool`

GetStartTlsEnabled returns the StartTlsEnabled field if non-nil, zero value otherwise.

### GetStartTlsEnabledOk

`func (o *ApiMailConfigurationDTO) GetStartTlsEnabledOk() (*bool, bool)`

GetStartTlsEnabledOk returns a tuple with the StartTlsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTlsEnabled

`func (o *ApiMailConfigurationDTO) SetStartTlsEnabled(v bool)`

SetStartTlsEnabled sets StartTlsEnabled field to given value.

### HasStartTlsEnabled

`func (o *ApiMailConfigurationDTO) HasStartTlsEnabled() bool`

HasStartTlsEnabled returns a boolean if a field has been set.

### GetSystemEmail

`func (o *ApiMailConfigurationDTO) GetSystemEmail() string`

GetSystemEmail returns the SystemEmail field if non-nil, zero value otherwise.

### GetSystemEmailOk

`func (o *ApiMailConfigurationDTO) GetSystemEmailOk() (*string, bool)`

GetSystemEmailOk returns a tuple with the SystemEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemEmail

`func (o *ApiMailConfigurationDTO) SetSystemEmail(v string)`

SetSystemEmail sets SystemEmail field to given value.

### HasSystemEmail

`func (o *ApiMailConfigurationDTO) HasSystemEmail() bool`

HasSystemEmail returns a boolean if a field has been set.

### GetUsername

`func (o *ApiMailConfigurationDTO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiMailConfigurationDTO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiMailConfigurationDTO) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ApiMailConfigurationDTO) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


