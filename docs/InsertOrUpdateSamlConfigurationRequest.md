# InsertOrUpdateSamlConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityProviderXml** | ***os.File** | Enter the SAML metadata XML of your IdP. Refer to the IdP documentation to obtain this metadata. | 
**SamlConfiguration** | [**ApiSamlConfigurationDTO**](ApiSamlConfigurationDTO.md) |  | 

## Methods

### NewInsertOrUpdateSamlConfigurationRequest

`func NewInsertOrUpdateSamlConfigurationRequest(identityProviderXml *os.File, samlConfiguration ApiSamlConfigurationDTO, ) *InsertOrUpdateSamlConfigurationRequest`

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

`func (o *InsertOrUpdateSamlConfigurationRequest) GetIdentityProviderXml() *os.File`

GetIdentityProviderXml returns the IdentityProviderXml field if non-nil, zero value otherwise.

### GetIdentityProviderXmlOk

`func (o *InsertOrUpdateSamlConfigurationRequest) GetIdentityProviderXmlOk() (**os.File, bool)`

GetIdentityProviderXmlOk returns a tuple with the IdentityProviderXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProviderXml

`func (o *InsertOrUpdateSamlConfigurationRequest) SetIdentityProviderXml(v *os.File)`

SetIdentityProviderXml sets IdentityProviderXml field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


