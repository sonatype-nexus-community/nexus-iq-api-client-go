# InsertOrUpdateSamlConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityProviderXml** | Pointer to **string** | Enter the SAML metadata XML of your IdP. Refer to the IdP documentation to obtain this metadata. | [optional] 
**SamlConfiguration** | Pointer to [**ApiSamlConfigurationDTO**](ApiSamlConfigurationDTO.md) |  | [optional] 

## Methods

### NewInsertOrUpdateSamlConfigurationRequest

`func NewInsertOrUpdateSamlConfigurationRequest() *InsertOrUpdateSamlConfigurationRequest`

NewInsertOrUpdateSamlConfigurationRequest instantiates a new InsertOrUpdateSamlConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsertOrUpdateSamlConfigurationRequestWithDefaults

`func NewInsertOrUpdateSamlConfigurationRequestWithDefaults() *InsertOrUpdateSamlConfigurationRequest`

NewInsertOrUpdateSamlConfigurationRequestWithDefaults instantiates a new InsertOrUpdateSamlConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityProviderXml

`func (o *InsertOrUpdateSamlConfigurationRequest) GetIdentityProviderXml() string`

GetIdentityProviderXml returns the IdentityProviderXml field if non-nil, zero value otherwise.

### GetIdentityProviderXmlOk

`func (o *InsertOrUpdateSamlConfigurationRequest) GetIdentityProviderXmlOk() (*string, bool)`

GetIdentityProviderXmlOk returns a tuple with the IdentityProviderXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderXml

`func (o *InsertOrUpdateSamlConfigurationRequest) SetIdentityProviderXml(v string)`

SetIdentityProviderXml sets IdentityProviderXml field to given value.

### HasIdentityProviderXml

`func (o *InsertOrUpdateSamlConfigurationRequest) HasIdentityProviderXml() bool`

HasIdentityProviderXml returns a boolean if a field has been set.

### GetSamlConfiguration

`func (o *InsertOrUpdateSamlConfigurationRequest) GetSamlConfiguration() ApiSamlConfigurationDTO`

GetSamlConfiguration returns the SamlConfiguration field if non-nil, zero value otherwise.

### GetSamlConfigurationOk

`func (o *InsertOrUpdateSamlConfigurationRequest) GetSamlConfigurationOk() (*ApiSamlConfigurationDTO, bool)`

GetSamlConfigurationOk returns a tuple with the SamlConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlConfiguration

`func (o *InsertOrUpdateSamlConfigurationRequest) SetSamlConfiguration(v ApiSamlConfigurationDTO)`

SetSamlConfiguration sets SamlConfiguration field to given value.

### HasSamlConfiguration

`func (o *InsertOrUpdateSamlConfigurationRequest) HasSamlConfiguration() bool`

HasSamlConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


