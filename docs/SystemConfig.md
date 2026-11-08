# SystemConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ADVANCED_REPORTING_INSIGHTS_ENABLED** | Pointer to **NullableBool** |  | [optional] 
**AdvancedSearchCSVExportDelimiter** | Pointer to **NullableString** |  | [optional] 
**AlpObservedLicenseDetectionEnabled** | Pointer to **NullableBool** |  | [optional] 
**ApiAccessAllowList** | Pointer to **[]string** |  | [optional] 
**BaseUrl** | Pointer to **NullableString** |  | [optional] 
**CsrfProtection** | Pointer to **NullableBool** |  | [optional] 
**EventBusMaxThreadPoolSize** | Pointer to **NullableInt32** |  | [optional] 
**ForceBaseUrl** | Pointer to **NullableBool** |  | [optional] 
**FrameAncestorsAllowlist** | Pointer to **[]string** |  | [optional] 
**HdsUrl** | Pointer to **NullableString** |  | [optional] 
**IntegrationsSupportedVersionCount** | Pointer to **NullableInt32** |  | [optional] 
**MaxAdvancedSearchClauseCount** | Pointer to **NullableInt32** |  | [optional] 
**PolicyMonitoringHour** | Pointer to **NullableInt32** |  | [optional] 
**PurgeScanFiles** | Pointer to **NullableString** |  | [optional] 
**SessionTimeout** | Pointer to **NullableInt32** |  | [optional] 
**SuccessMetricsStageId** | Pointer to **NullableString** |  | [optional] 
**UserAgentSuffix** | Pointer to **NullableString** |  | [optional] 
**WaivedComponentUpgradeMonitoringEnabled** | Pointer to **NullableBool** |  | [optional] 
**WebhookSecretPassphrase** | Pointer to **NullableString** |  | [optional] 

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

### GetADVANCED_REPORTING_INSIGHTS_ENABLED

`func (o *SystemConfig) GetADVANCED_REPORTING_INSIGHTS_ENABLED() bool`

GetADVANCED_REPORTING_INSIGHTS_ENABLED returns the ADVANCED_REPORTING_INSIGHTS_ENABLED field if non-nil, zero value otherwise.

### GetADVANCED_REPORTING_INSIGHTS_ENABLEDOk

`func (o *SystemConfig) GetADVANCED_REPORTING_INSIGHTS_ENABLEDOk() (*bool, bool)`

GetADVANCED_REPORTING_INSIGHTS_ENABLEDOk returns a tuple with the ADVANCED_REPORTING_INSIGHTS_ENABLED field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetADVANCED_REPORTING_INSIGHTS_ENABLED

`func (o *SystemConfig) SetADVANCED_REPORTING_INSIGHTS_ENABLED(v bool)`

SetADVANCED_REPORTING_INSIGHTS_ENABLED sets ADVANCED_REPORTING_INSIGHTS_ENABLED field to given value.

### HasADVANCED_REPORTING_INSIGHTS_ENABLED

`func (o *SystemConfig) HasADVANCED_REPORTING_INSIGHTS_ENABLED() bool`

HasADVANCED_REPORTING_INSIGHTS_ENABLED returns a boolean if a field has been set.

### SetADVANCED_REPORTING_INSIGHTS_ENABLEDNil

`func (o *SystemConfig) SetADVANCED_REPORTING_INSIGHTS_ENABLEDNil(b bool)`

 SetADVANCED_REPORTING_INSIGHTS_ENABLEDNil sets the value for ADVANCED_REPORTING_INSIGHTS_ENABLED to be an explicit nil

### UnsetADVANCED_REPORTING_INSIGHTS_ENABLED
`func (o *SystemConfig) UnsetADVANCED_REPORTING_INSIGHTS_ENABLED()`

UnsetADVANCED_REPORTING_INSIGHTS_ENABLED ensures that no value is present for ADVANCED_REPORTING_INSIGHTS_ENABLED, not even an explicit nil
### GetAdvancedSearchCSVExportDelimiter

`func (o *SystemConfig) GetAdvancedSearchCSVExportDelimiter() string`

GetAdvancedSearchCSVExportDelimiter returns the AdvancedSearchCSVExportDelimiter field if non-nil, zero value otherwise.

### GetAdvancedSearchCSVExportDelimiterOk

`func (o *SystemConfig) GetAdvancedSearchCSVExportDelimiterOk() (*string, bool)`

GetAdvancedSearchCSVExportDelimiterOk returns a tuple with the AdvancedSearchCSVExportDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedSearchCSVExportDelimiter

`func (o *SystemConfig) SetAdvancedSearchCSVExportDelimiter(v string)`

SetAdvancedSearchCSVExportDelimiter sets AdvancedSearchCSVExportDelimiter field to given value.

### HasAdvancedSearchCSVExportDelimiter

`func (o *SystemConfig) HasAdvancedSearchCSVExportDelimiter() bool`

HasAdvancedSearchCSVExportDelimiter returns a boolean if a field has been set.

### SetAdvancedSearchCSVExportDelimiterNil

`func (o *SystemConfig) SetAdvancedSearchCSVExportDelimiterNil(b bool)`

 SetAdvancedSearchCSVExportDelimiterNil sets the value for AdvancedSearchCSVExportDelimiter to be an explicit nil

### UnsetAdvancedSearchCSVExportDelimiter
`func (o *SystemConfig) UnsetAdvancedSearchCSVExportDelimiter()`

UnsetAdvancedSearchCSVExportDelimiter ensures that no value is present for AdvancedSearchCSVExportDelimiter, not even an explicit nil
### GetAlpObservedLicenseDetectionEnabled

`func (o *SystemConfig) GetAlpObservedLicenseDetectionEnabled() bool`

GetAlpObservedLicenseDetectionEnabled returns the AlpObservedLicenseDetectionEnabled field if non-nil, zero value otherwise.

### GetAlpObservedLicenseDetectionEnabledOk

`func (o *SystemConfig) GetAlpObservedLicenseDetectionEnabledOk() (*bool, bool)`

GetAlpObservedLicenseDetectionEnabledOk returns a tuple with the AlpObservedLicenseDetectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpObservedLicenseDetectionEnabled

`func (o *SystemConfig) SetAlpObservedLicenseDetectionEnabled(v bool)`

SetAlpObservedLicenseDetectionEnabled sets AlpObservedLicenseDetectionEnabled field to given value.

### HasAlpObservedLicenseDetectionEnabled

`func (o *SystemConfig) HasAlpObservedLicenseDetectionEnabled() bool`

HasAlpObservedLicenseDetectionEnabled returns a boolean if a field has been set.

### SetAlpObservedLicenseDetectionEnabledNil

`func (o *SystemConfig) SetAlpObservedLicenseDetectionEnabledNil(b bool)`

 SetAlpObservedLicenseDetectionEnabledNil sets the value for AlpObservedLicenseDetectionEnabled to be an explicit nil

### UnsetAlpObservedLicenseDetectionEnabled
`func (o *SystemConfig) UnsetAlpObservedLicenseDetectionEnabled()`

UnsetAlpObservedLicenseDetectionEnabled ensures that no value is present for AlpObservedLicenseDetectionEnabled, not even an explicit nil
### GetApiAccessAllowList

`func (o *SystemConfig) GetApiAccessAllowList() []string`

GetApiAccessAllowList returns the ApiAccessAllowList field if non-nil, zero value otherwise.

### GetApiAccessAllowListOk

`func (o *SystemConfig) GetApiAccessAllowListOk() (*[]string, bool)`

GetApiAccessAllowListOk returns a tuple with the ApiAccessAllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAccessAllowList

`func (o *SystemConfig) SetApiAccessAllowList(v []string)`

SetApiAccessAllowList sets ApiAccessAllowList field to given value.

### HasApiAccessAllowList

`func (o *SystemConfig) HasApiAccessAllowList() bool`

HasApiAccessAllowList returns a boolean if a field has been set.

### SetApiAccessAllowListNil

`func (o *SystemConfig) SetApiAccessAllowListNil(b bool)`

 SetApiAccessAllowListNil sets the value for ApiAccessAllowList to be an explicit nil

### UnsetApiAccessAllowList
`func (o *SystemConfig) UnsetApiAccessAllowList()`

UnsetApiAccessAllowList ensures that no value is present for ApiAccessAllowList, not even an explicit nil
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
### GetCsrfProtection

`func (o *SystemConfig) GetCsrfProtection() bool`

GetCsrfProtection returns the CsrfProtection field if non-nil, zero value otherwise.

### GetCsrfProtectionOk

`func (o *SystemConfig) GetCsrfProtectionOk() (*bool, bool)`

GetCsrfProtectionOk returns a tuple with the CsrfProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsrfProtection

`func (o *SystemConfig) SetCsrfProtection(v bool)`

SetCsrfProtection sets CsrfProtection field to given value.

### HasCsrfProtection

`func (o *SystemConfig) HasCsrfProtection() bool`

HasCsrfProtection returns a boolean if a field has been set.

### SetCsrfProtectionNil

`func (o *SystemConfig) SetCsrfProtectionNil(b bool)`

 SetCsrfProtectionNil sets the value for CsrfProtection to be an explicit nil

### UnsetCsrfProtection
`func (o *SystemConfig) UnsetCsrfProtection()`

UnsetCsrfProtection ensures that no value is present for CsrfProtection, not even an explicit nil
### GetEventBusMaxThreadPoolSize

`func (o *SystemConfig) GetEventBusMaxThreadPoolSize() int32`

GetEventBusMaxThreadPoolSize returns the EventBusMaxThreadPoolSize field if non-nil, zero value otherwise.

### GetEventBusMaxThreadPoolSizeOk

`func (o *SystemConfig) GetEventBusMaxThreadPoolSizeOk() (*int32, bool)`

GetEventBusMaxThreadPoolSizeOk returns a tuple with the EventBusMaxThreadPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventBusMaxThreadPoolSize

`func (o *SystemConfig) SetEventBusMaxThreadPoolSize(v int32)`

SetEventBusMaxThreadPoolSize sets EventBusMaxThreadPoolSize field to given value.

### HasEventBusMaxThreadPoolSize

`func (o *SystemConfig) HasEventBusMaxThreadPoolSize() bool`

HasEventBusMaxThreadPoolSize returns a boolean if a field has been set.

### SetEventBusMaxThreadPoolSizeNil

`func (o *SystemConfig) SetEventBusMaxThreadPoolSizeNil(b bool)`

 SetEventBusMaxThreadPoolSizeNil sets the value for EventBusMaxThreadPoolSize to be an explicit nil

### UnsetEventBusMaxThreadPoolSize
`func (o *SystemConfig) UnsetEventBusMaxThreadPoolSize()`

UnsetEventBusMaxThreadPoolSize ensures that no value is present for EventBusMaxThreadPoolSize, not even an explicit nil
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
### GetFrameAncestorsAllowlist

`func (o *SystemConfig) GetFrameAncestorsAllowlist() []string`

GetFrameAncestorsAllowlist returns the FrameAncestorsAllowlist field if non-nil, zero value otherwise.

### GetFrameAncestorsAllowlistOk

`func (o *SystemConfig) GetFrameAncestorsAllowlistOk() (*[]string, bool)`

GetFrameAncestorsAllowlistOk returns a tuple with the FrameAncestorsAllowlist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameAncestorsAllowlist

`func (o *SystemConfig) SetFrameAncestorsAllowlist(v []string)`

SetFrameAncestorsAllowlist sets FrameAncestorsAllowlist field to given value.

### HasFrameAncestorsAllowlist

`func (o *SystemConfig) HasFrameAncestorsAllowlist() bool`

HasFrameAncestorsAllowlist returns a boolean if a field has been set.

### SetFrameAncestorsAllowlistNil

`func (o *SystemConfig) SetFrameAncestorsAllowlistNil(b bool)`

 SetFrameAncestorsAllowlistNil sets the value for FrameAncestorsAllowlist to be an explicit nil

### UnsetFrameAncestorsAllowlist
`func (o *SystemConfig) UnsetFrameAncestorsAllowlist()`

UnsetFrameAncestorsAllowlist ensures that no value is present for FrameAncestorsAllowlist, not even an explicit nil
### GetHdsUrl

`func (o *SystemConfig) GetHdsUrl() string`

GetHdsUrl returns the HdsUrl field if non-nil, zero value otherwise.

### GetHdsUrlOk

`func (o *SystemConfig) GetHdsUrlOk() (*string, bool)`

GetHdsUrlOk returns a tuple with the HdsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHdsUrl

`func (o *SystemConfig) SetHdsUrl(v string)`

SetHdsUrl sets HdsUrl field to given value.

### HasHdsUrl

`func (o *SystemConfig) HasHdsUrl() bool`

HasHdsUrl returns a boolean if a field has been set.

### SetHdsUrlNil

`func (o *SystemConfig) SetHdsUrlNil(b bool)`

 SetHdsUrlNil sets the value for HdsUrl to be an explicit nil

### UnsetHdsUrl
`func (o *SystemConfig) UnsetHdsUrl()`

UnsetHdsUrl ensures that no value is present for HdsUrl, not even an explicit nil
### GetIntegrationsSupportedVersionCount

`func (o *SystemConfig) GetIntegrationsSupportedVersionCount() int32`

GetIntegrationsSupportedVersionCount returns the IntegrationsSupportedVersionCount field if non-nil, zero value otherwise.

### GetIntegrationsSupportedVersionCountOk

`func (o *SystemConfig) GetIntegrationsSupportedVersionCountOk() (*int32, bool)`

GetIntegrationsSupportedVersionCountOk returns a tuple with the IntegrationsSupportedVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationsSupportedVersionCount

`func (o *SystemConfig) SetIntegrationsSupportedVersionCount(v int32)`

SetIntegrationsSupportedVersionCount sets IntegrationsSupportedVersionCount field to given value.

### HasIntegrationsSupportedVersionCount

`func (o *SystemConfig) HasIntegrationsSupportedVersionCount() bool`

HasIntegrationsSupportedVersionCount returns a boolean if a field has been set.

### SetIntegrationsSupportedVersionCountNil

`func (o *SystemConfig) SetIntegrationsSupportedVersionCountNil(b bool)`

 SetIntegrationsSupportedVersionCountNil sets the value for IntegrationsSupportedVersionCount to be an explicit nil

### UnsetIntegrationsSupportedVersionCount
`func (o *SystemConfig) UnsetIntegrationsSupportedVersionCount()`

UnsetIntegrationsSupportedVersionCount ensures that no value is present for IntegrationsSupportedVersionCount, not even an explicit nil
### GetMaxAdvancedSearchClauseCount

`func (o *SystemConfig) GetMaxAdvancedSearchClauseCount() int32`

GetMaxAdvancedSearchClauseCount returns the MaxAdvancedSearchClauseCount field if non-nil, zero value otherwise.

### GetMaxAdvancedSearchClauseCountOk

`func (o *SystemConfig) GetMaxAdvancedSearchClauseCountOk() (*int32, bool)`

GetMaxAdvancedSearchClauseCountOk returns a tuple with the MaxAdvancedSearchClauseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAdvancedSearchClauseCount

`func (o *SystemConfig) SetMaxAdvancedSearchClauseCount(v int32)`

SetMaxAdvancedSearchClauseCount sets MaxAdvancedSearchClauseCount field to given value.

### HasMaxAdvancedSearchClauseCount

`func (o *SystemConfig) HasMaxAdvancedSearchClauseCount() bool`

HasMaxAdvancedSearchClauseCount returns a boolean if a field has been set.

### SetMaxAdvancedSearchClauseCountNil

`func (o *SystemConfig) SetMaxAdvancedSearchClauseCountNil(b bool)`

 SetMaxAdvancedSearchClauseCountNil sets the value for MaxAdvancedSearchClauseCount to be an explicit nil

### UnsetMaxAdvancedSearchClauseCount
`func (o *SystemConfig) UnsetMaxAdvancedSearchClauseCount()`

UnsetMaxAdvancedSearchClauseCount ensures that no value is present for MaxAdvancedSearchClauseCount, not even an explicit nil
### GetPolicyMonitoringHour

`func (o *SystemConfig) GetPolicyMonitoringHour() int32`

GetPolicyMonitoringHour returns the PolicyMonitoringHour field if non-nil, zero value otherwise.

### GetPolicyMonitoringHourOk

`func (o *SystemConfig) GetPolicyMonitoringHourOk() (*int32, bool)`

GetPolicyMonitoringHourOk returns a tuple with the PolicyMonitoringHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyMonitoringHour

`func (o *SystemConfig) SetPolicyMonitoringHour(v int32)`

SetPolicyMonitoringHour sets PolicyMonitoringHour field to given value.

### HasPolicyMonitoringHour

`func (o *SystemConfig) HasPolicyMonitoringHour() bool`

HasPolicyMonitoringHour returns a boolean if a field has been set.

### SetPolicyMonitoringHourNil

`func (o *SystemConfig) SetPolicyMonitoringHourNil(b bool)`

 SetPolicyMonitoringHourNil sets the value for PolicyMonitoringHour to be an explicit nil

### UnsetPolicyMonitoringHour
`func (o *SystemConfig) UnsetPolicyMonitoringHour()`

UnsetPolicyMonitoringHour ensures that no value is present for PolicyMonitoringHour, not even an explicit nil
### GetPurgeScanFiles

`func (o *SystemConfig) GetPurgeScanFiles() string`

GetPurgeScanFiles returns the PurgeScanFiles field if non-nil, zero value otherwise.

### GetPurgeScanFilesOk

`func (o *SystemConfig) GetPurgeScanFilesOk() (*string, bool)`

GetPurgeScanFilesOk returns a tuple with the PurgeScanFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurgeScanFiles

`func (o *SystemConfig) SetPurgeScanFiles(v string)`

SetPurgeScanFiles sets PurgeScanFiles field to given value.

### HasPurgeScanFiles

`func (o *SystemConfig) HasPurgeScanFiles() bool`

HasPurgeScanFiles returns a boolean if a field has been set.

### SetPurgeScanFilesNil

`func (o *SystemConfig) SetPurgeScanFilesNil(b bool)`

 SetPurgeScanFilesNil sets the value for PurgeScanFiles to be an explicit nil

### UnsetPurgeScanFiles
`func (o *SystemConfig) UnsetPurgeScanFiles()`

UnsetPurgeScanFiles ensures that no value is present for PurgeScanFiles, not even an explicit nil
### GetSessionTimeout

`func (o *SystemConfig) GetSessionTimeout() int32`

GetSessionTimeout returns the SessionTimeout field if non-nil, zero value otherwise.

### GetSessionTimeoutOk

`func (o *SystemConfig) GetSessionTimeoutOk() (*int32, bool)`

GetSessionTimeoutOk returns a tuple with the SessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeout

`func (o *SystemConfig) SetSessionTimeout(v int32)`

SetSessionTimeout sets SessionTimeout field to given value.

### HasSessionTimeout

`func (o *SystemConfig) HasSessionTimeout() bool`

HasSessionTimeout returns a boolean if a field has been set.

### SetSessionTimeoutNil

`func (o *SystemConfig) SetSessionTimeoutNil(b bool)`

 SetSessionTimeoutNil sets the value for SessionTimeout to be an explicit nil

### UnsetSessionTimeout
`func (o *SystemConfig) UnsetSessionTimeout()`

UnsetSessionTimeout ensures that no value is present for SessionTimeout, not even an explicit nil
### GetSuccessMetricsStageId

`func (o *SystemConfig) GetSuccessMetricsStageId() string`

GetSuccessMetricsStageId returns the SuccessMetricsStageId field if non-nil, zero value otherwise.

### GetSuccessMetricsStageIdOk

`func (o *SystemConfig) GetSuccessMetricsStageIdOk() (*string, bool)`

GetSuccessMetricsStageIdOk returns a tuple with the SuccessMetricsStageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessMetricsStageId

`func (o *SystemConfig) SetSuccessMetricsStageId(v string)`

SetSuccessMetricsStageId sets SuccessMetricsStageId field to given value.

### HasSuccessMetricsStageId

`func (o *SystemConfig) HasSuccessMetricsStageId() bool`

HasSuccessMetricsStageId returns a boolean if a field has been set.

### SetSuccessMetricsStageIdNil

`func (o *SystemConfig) SetSuccessMetricsStageIdNil(b bool)`

 SetSuccessMetricsStageIdNil sets the value for SuccessMetricsStageId to be an explicit nil

### UnsetSuccessMetricsStageId
`func (o *SystemConfig) UnsetSuccessMetricsStageId()`

UnsetSuccessMetricsStageId ensures that no value is present for SuccessMetricsStageId, not even an explicit nil
### GetUserAgentSuffix

`func (o *SystemConfig) GetUserAgentSuffix() string`

GetUserAgentSuffix returns the UserAgentSuffix field if non-nil, zero value otherwise.

### GetUserAgentSuffixOk

`func (o *SystemConfig) GetUserAgentSuffixOk() (*string, bool)`

GetUserAgentSuffixOk returns a tuple with the UserAgentSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentSuffix

`func (o *SystemConfig) SetUserAgentSuffix(v string)`

SetUserAgentSuffix sets UserAgentSuffix field to given value.

### HasUserAgentSuffix

`func (o *SystemConfig) HasUserAgentSuffix() bool`

HasUserAgentSuffix returns a boolean if a field has been set.

### SetUserAgentSuffixNil

`func (o *SystemConfig) SetUserAgentSuffixNil(b bool)`

 SetUserAgentSuffixNil sets the value for UserAgentSuffix to be an explicit nil

### UnsetUserAgentSuffix
`func (o *SystemConfig) UnsetUserAgentSuffix()`

UnsetUserAgentSuffix ensures that no value is present for UserAgentSuffix, not even an explicit nil
### GetWaivedComponentUpgradeMonitoringEnabled

`func (o *SystemConfig) GetWaivedComponentUpgradeMonitoringEnabled() bool`

GetWaivedComponentUpgradeMonitoringEnabled returns the WaivedComponentUpgradeMonitoringEnabled field if non-nil, zero value otherwise.

### GetWaivedComponentUpgradeMonitoringEnabledOk

`func (o *SystemConfig) GetWaivedComponentUpgradeMonitoringEnabledOk() (*bool, bool)`

GetWaivedComponentUpgradeMonitoringEnabledOk returns a tuple with the WaivedComponentUpgradeMonitoringEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaivedComponentUpgradeMonitoringEnabled

`func (o *SystemConfig) SetWaivedComponentUpgradeMonitoringEnabled(v bool)`

SetWaivedComponentUpgradeMonitoringEnabled sets WaivedComponentUpgradeMonitoringEnabled field to given value.

### HasWaivedComponentUpgradeMonitoringEnabled

`func (o *SystemConfig) HasWaivedComponentUpgradeMonitoringEnabled() bool`

HasWaivedComponentUpgradeMonitoringEnabled returns a boolean if a field has been set.

### SetWaivedComponentUpgradeMonitoringEnabledNil

`func (o *SystemConfig) SetWaivedComponentUpgradeMonitoringEnabledNil(b bool)`

 SetWaivedComponentUpgradeMonitoringEnabledNil sets the value for WaivedComponentUpgradeMonitoringEnabled to be an explicit nil

### UnsetWaivedComponentUpgradeMonitoringEnabled
`func (o *SystemConfig) UnsetWaivedComponentUpgradeMonitoringEnabled()`

UnsetWaivedComponentUpgradeMonitoringEnabled ensures that no value is present for WaivedComponentUpgradeMonitoringEnabled, not even an explicit nil
### GetWebhookSecretPassphrase

`func (o *SystemConfig) GetWebhookSecretPassphrase() string`

GetWebhookSecretPassphrase returns the WebhookSecretPassphrase field if non-nil, zero value otherwise.

### GetWebhookSecretPassphraseOk

`func (o *SystemConfig) GetWebhookSecretPassphraseOk() (*string, bool)`

GetWebhookSecretPassphraseOk returns a tuple with the WebhookSecretPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecretPassphrase

`func (o *SystemConfig) SetWebhookSecretPassphrase(v string)`

SetWebhookSecretPassphrase sets WebhookSecretPassphrase field to given value.

### HasWebhookSecretPassphrase

`func (o *SystemConfig) HasWebhookSecretPassphrase() bool`

HasWebhookSecretPassphrase returns a boolean if a field has been set.

### SetWebhookSecretPassphraseNil

`func (o *SystemConfig) SetWebhookSecretPassphraseNil(b bool)`

 SetWebhookSecretPassphraseNil sets the value for WebhookSecretPassphrase to be an explicit nil

### UnsetWebhookSecretPassphrase
`func (o *SystemConfig) UnsetWebhookSecretPassphrase()`

UnsetWebhookSecretPassphrase ensures that no value is present for WebhookSecretPassphrase, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


