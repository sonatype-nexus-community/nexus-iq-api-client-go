# ConfigurationValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationComplete** | Pointer to [**ValidationResult**](ValidationResult.md) |  | [optional] 
**RepoPrivate** | Pointer to [**ValidationResult**](ValidationResult.md) |  | [optional] 
**SshConfiguration** | Pointer to [**ValidationResult**](ValidationResult.md) |  | [optional] 
**TokenPermissions** | Pointer to [**ValidationResult**](ValidationResult.md) |  | [optional] 

## Methods

### NewConfigurationValidationResult

`func NewConfigurationValidationResult() *ConfigurationValidationResult`

NewConfigurationValidationResult instantiates a new ConfigurationValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationValidationResultWithDefaults

`func NewConfigurationValidationResultWithDefaults() *ConfigurationValidationResult`

NewConfigurationValidationResultWithDefaults instantiates a new ConfigurationValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationComplete

`func (o *ConfigurationValidationResult) GetConfigurationComplete() ValidationResult`

GetConfigurationComplete returns the ConfigurationComplete field if non-nil, zero value otherwise.

### GetConfigurationCompleteOk

`func (o *ConfigurationValidationResult) GetConfigurationCompleteOk() (*ValidationResult, bool)`

GetConfigurationCompleteOk returns a tuple with the ConfigurationComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationComplete

`func (o *ConfigurationValidationResult) SetConfigurationComplete(v ValidationResult)`

SetConfigurationComplete sets ConfigurationComplete field to given value.

### HasConfigurationComplete

`func (o *ConfigurationValidationResult) HasConfigurationComplete() bool`

HasConfigurationComplete returns a boolean if a field has been set.

### GetRepoPrivate

`func (o *ConfigurationValidationResult) GetRepoPrivate() ValidationResult`

GetRepoPrivate returns the RepoPrivate field if non-nil, zero value otherwise.

### GetRepoPrivateOk

`func (o *ConfigurationValidationResult) GetRepoPrivateOk() (*ValidationResult, bool)`

GetRepoPrivateOk returns a tuple with the RepoPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoPrivate

`func (o *ConfigurationValidationResult) SetRepoPrivate(v ValidationResult)`

SetRepoPrivate sets RepoPrivate field to given value.

### HasRepoPrivate

`func (o *ConfigurationValidationResult) HasRepoPrivate() bool`

HasRepoPrivate returns a boolean if a field has been set.

### GetSshConfiguration

`func (o *ConfigurationValidationResult) GetSshConfiguration() ValidationResult`

GetSshConfiguration returns the SshConfiguration field if non-nil, zero value otherwise.

### GetSshConfigurationOk

`func (o *ConfigurationValidationResult) GetSshConfigurationOk() (*ValidationResult, bool)`

GetSshConfigurationOk returns a tuple with the SshConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshConfiguration

`func (o *ConfigurationValidationResult) SetSshConfiguration(v ValidationResult)`

SetSshConfiguration sets SshConfiguration field to given value.

### HasSshConfiguration

`func (o *ConfigurationValidationResult) HasSshConfiguration() bool`

HasSshConfiguration returns a boolean if a field has been set.

### GetTokenPermissions

`func (o *ConfigurationValidationResult) GetTokenPermissions() ValidationResult`

GetTokenPermissions returns the TokenPermissions field if non-nil, zero value otherwise.

### GetTokenPermissionsOk

`func (o *ConfigurationValidationResult) GetTokenPermissionsOk() (*ValidationResult, bool)`

GetTokenPermissionsOk returns a tuple with the TokenPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPermissions

`func (o *ConfigurationValidationResult) SetTokenPermissions(v ValidationResult)`

SetTokenPermissions sets TokenPermissions field to given value.

### HasTokenPermissions

`func (o *ConfigurationValidationResult) HasTokenPermissions() bool`

HasTokenPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


