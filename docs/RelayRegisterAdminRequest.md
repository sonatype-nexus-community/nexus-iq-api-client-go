# RelayRegisterAdminRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstallationId** | Pointer to **string** |  | [optional] 
**WebhookSecret** | Pointer to **string** |  | [optional] 

## Methods

### NewRelayRegisterAdminRequest

`func NewRelayRegisterAdminRequest() *RelayRegisterAdminRequest`

NewRelayRegisterAdminRequest instantiates a new RelayRegisterAdminRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelayRegisterAdminRequestWithDefaults

`func NewRelayRegisterAdminRequestWithDefaults() *RelayRegisterAdminRequest`

NewRelayRegisterAdminRequestWithDefaults instantiates a new RelayRegisterAdminRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstallationId

`func (o *RelayRegisterAdminRequest) GetInstallationId() string`

GetInstallationId returns the InstallationId field if non-nil, zero value otherwise.

### GetInstallationIdOk

`func (o *RelayRegisterAdminRequest) GetInstallationIdOk() (*string, bool)`

GetInstallationIdOk returns a tuple with the InstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationId

`func (o *RelayRegisterAdminRequest) SetInstallationId(v string)`

SetInstallationId sets InstallationId field to given value.

### HasInstallationId

`func (o *RelayRegisterAdminRequest) HasInstallationId() bool`

HasInstallationId returns a boolean if a field has been set.

### GetWebhookSecret

`func (o *RelayRegisterAdminRequest) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *RelayRegisterAdminRequest) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *RelayRegisterAdminRequest) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *RelayRegisterAdminRequest) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


