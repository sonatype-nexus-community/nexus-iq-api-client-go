# SsoConfigurationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oauth2Configuration** | Pointer to [**OAuth2ConfigurationDTO**](OAuth2ConfigurationDTO.md) |  | [optional] 
**OidcConfiguration** | Pointer to [**OidcConfigurationDTO**](OidcConfigurationDTO.md) |  | [optional] 

## Methods

### NewSsoConfigurationDTO

`func NewSsoConfigurationDTO() *SsoConfigurationDTO`

NewSsoConfigurationDTO instantiates a new SsoConfigurationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoConfigurationDTOWithDefaults

`func NewSsoConfigurationDTOWithDefaults() *SsoConfigurationDTO`

NewSsoConfigurationDTOWithDefaults instantiates a new SsoConfigurationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOauth2Configuration

`func (o *SsoConfigurationDTO) GetOauth2Configuration() OAuth2ConfigurationDTO`

GetOauth2Configuration returns the Oauth2Configuration field if non-nil, zero value otherwise.

### GetOauth2ConfigurationOk

`func (o *SsoConfigurationDTO) GetOauth2ConfigurationOk() (*OAuth2ConfigurationDTO, bool)`

GetOauth2ConfigurationOk returns a tuple with the Oauth2Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2Configuration

`func (o *SsoConfigurationDTO) SetOauth2Configuration(v OAuth2ConfigurationDTO)`

SetOauth2Configuration sets Oauth2Configuration field to given value.

### HasOauth2Configuration

`func (o *SsoConfigurationDTO) HasOauth2Configuration() bool`

HasOauth2Configuration returns a boolean if a field has been set.

### GetOidcConfiguration

`func (o *SsoConfigurationDTO) GetOidcConfiguration() OidcConfigurationDTO`

GetOidcConfiguration returns the OidcConfiguration field if non-nil, zero value otherwise.

### GetOidcConfigurationOk

`func (o *SsoConfigurationDTO) GetOidcConfigurationOk() (*OidcConfigurationDTO, bool)`

GetOidcConfigurationOk returns a tuple with the OidcConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcConfiguration

`func (o *SsoConfigurationDTO) SetOidcConfiguration(v OidcConfigurationDTO)`

SetOidcConfiguration sets OidcConfiguration field to given value.

### HasOidcConfiguration

`func (o *SsoConfigurationDTO) HasOidcConfiguration() bool`

HasOidcConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


