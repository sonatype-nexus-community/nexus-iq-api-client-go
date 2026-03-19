# ApiCiConfigurationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvancedProperties** | Pointer to **[]string** |  | [optional] 
**Download** | Pointer to [**DownloadConfig**](DownloadConfig.md) |  | [optional] 
**EnableDebugLogging** | Pointer to **bool** |  | [optional] 
**FailBuildOnNetworkError** | Pointer to **bool** |  | [optional] 
**FailBuildOnPolicyWarnings** | Pointer to **bool** |  | [optional] 
**FailBuildOnReachabilityErrors** | Pointer to **bool** |  | [optional] 
**FailBuildOnScanningErrors** | Pointer to **bool** |  | [optional] 
**ModuleExcludes** | Pointer to **[]string** |  | [optional] 
**ParameterPriority** | Pointer to **string** |  | [optional] 
**Proxy** | Pointer to [**ProxyConfig**](ProxyConfig.md) |  | [optional] 
**Reachability** | Pointer to [**ReachabilityConfig**](ReachabilityConfig.md) |  | [optional] 
**ResultFile** | Pointer to **string** |  | [optional] 
**SarifFile** | Pointer to **string** |  | [optional] 
**ScanPatterns** | Pointer to **[]string** |  | [optional] 
**UnstableBuildOnPolicyWarnings** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiCiConfigurationDto

`func NewApiCiConfigurationDto() *ApiCiConfigurationDto`

NewApiCiConfigurationDto instantiates a new ApiCiConfigurationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCiConfigurationDtoWithDefaults

`func NewApiCiConfigurationDtoWithDefaults() *ApiCiConfigurationDto`

NewApiCiConfigurationDtoWithDefaults instantiates a new ApiCiConfigurationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvancedProperties

`func (o *ApiCiConfigurationDto) GetAdvancedProperties() []string`

GetAdvancedProperties returns the AdvancedProperties field if non-nil, zero value otherwise.

### GetAdvancedPropertiesOk

`func (o *ApiCiConfigurationDto) GetAdvancedPropertiesOk() (*[]string, bool)`

GetAdvancedPropertiesOk returns a tuple with the AdvancedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedProperties

`func (o *ApiCiConfigurationDto) SetAdvancedProperties(v []string)`

SetAdvancedProperties sets AdvancedProperties field to given value.

### HasAdvancedProperties

`func (o *ApiCiConfigurationDto) HasAdvancedProperties() bool`

HasAdvancedProperties returns a boolean if a field has been set.

### GetDownload

`func (o *ApiCiConfigurationDto) GetDownload() DownloadConfig`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ApiCiConfigurationDto) GetDownloadOk() (*DownloadConfig, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ApiCiConfigurationDto) SetDownload(v DownloadConfig)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *ApiCiConfigurationDto) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetEnableDebugLogging

`func (o *ApiCiConfigurationDto) GetEnableDebugLogging() bool`

GetEnableDebugLogging returns the EnableDebugLogging field if non-nil, zero value otherwise.

### GetEnableDebugLoggingOk

`func (o *ApiCiConfigurationDto) GetEnableDebugLoggingOk() (*bool, bool)`

GetEnableDebugLoggingOk returns a tuple with the EnableDebugLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableDebugLogging

`func (o *ApiCiConfigurationDto) SetEnableDebugLogging(v bool)`

SetEnableDebugLogging sets EnableDebugLogging field to given value.

### HasEnableDebugLogging

`func (o *ApiCiConfigurationDto) HasEnableDebugLogging() bool`

HasEnableDebugLogging returns a boolean if a field has been set.

### GetFailBuildOnNetworkError

`func (o *ApiCiConfigurationDto) GetFailBuildOnNetworkError() bool`

GetFailBuildOnNetworkError returns the FailBuildOnNetworkError field if non-nil, zero value otherwise.

### GetFailBuildOnNetworkErrorOk

`func (o *ApiCiConfigurationDto) GetFailBuildOnNetworkErrorOk() (*bool, bool)`

GetFailBuildOnNetworkErrorOk returns a tuple with the FailBuildOnNetworkError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailBuildOnNetworkError

`func (o *ApiCiConfigurationDto) SetFailBuildOnNetworkError(v bool)`

SetFailBuildOnNetworkError sets FailBuildOnNetworkError field to given value.

### HasFailBuildOnNetworkError

`func (o *ApiCiConfigurationDto) HasFailBuildOnNetworkError() bool`

HasFailBuildOnNetworkError returns a boolean if a field has been set.

### GetFailBuildOnPolicyWarnings

`func (o *ApiCiConfigurationDto) GetFailBuildOnPolicyWarnings() bool`

GetFailBuildOnPolicyWarnings returns the FailBuildOnPolicyWarnings field if non-nil, zero value otherwise.

### GetFailBuildOnPolicyWarningsOk

`func (o *ApiCiConfigurationDto) GetFailBuildOnPolicyWarningsOk() (*bool, bool)`

GetFailBuildOnPolicyWarningsOk returns a tuple with the FailBuildOnPolicyWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailBuildOnPolicyWarnings

`func (o *ApiCiConfigurationDto) SetFailBuildOnPolicyWarnings(v bool)`

SetFailBuildOnPolicyWarnings sets FailBuildOnPolicyWarnings field to given value.

### HasFailBuildOnPolicyWarnings

`func (o *ApiCiConfigurationDto) HasFailBuildOnPolicyWarnings() bool`

HasFailBuildOnPolicyWarnings returns a boolean if a field has been set.

### GetFailBuildOnReachabilityErrors

`func (o *ApiCiConfigurationDto) GetFailBuildOnReachabilityErrors() bool`

GetFailBuildOnReachabilityErrors returns the FailBuildOnReachabilityErrors field if non-nil, zero value otherwise.

### GetFailBuildOnReachabilityErrorsOk

`func (o *ApiCiConfigurationDto) GetFailBuildOnReachabilityErrorsOk() (*bool, bool)`

GetFailBuildOnReachabilityErrorsOk returns a tuple with the FailBuildOnReachabilityErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailBuildOnReachabilityErrors

`func (o *ApiCiConfigurationDto) SetFailBuildOnReachabilityErrors(v bool)`

SetFailBuildOnReachabilityErrors sets FailBuildOnReachabilityErrors field to given value.

### HasFailBuildOnReachabilityErrors

`func (o *ApiCiConfigurationDto) HasFailBuildOnReachabilityErrors() bool`

HasFailBuildOnReachabilityErrors returns a boolean if a field has been set.

### GetFailBuildOnScanningErrors

`func (o *ApiCiConfigurationDto) GetFailBuildOnScanningErrors() bool`

GetFailBuildOnScanningErrors returns the FailBuildOnScanningErrors field if non-nil, zero value otherwise.

### GetFailBuildOnScanningErrorsOk

`func (o *ApiCiConfigurationDto) GetFailBuildOnScanningErrorsOk() (*bool, bool)`

GetFailBuildOnScanningErrorsOk returns a tuple with the FailBuildOnScanningErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailBuildOnScanningErrors

`func (o *ApiCiConfigurationDto) SetFailBuildOnScanningErrors(v bool)`

SetFailBuildOnScanningErrors sets FailBuildOnScanningErrors field to given value.

### HasFailBuildOnScanningErrors

`func (o *ApiCiConfigurationDto) HasFailBuildOnScanningErrors() bool`

HasFailBuildOnScanningErrors returns a boolean if a field has been set.

### GetModuleExcludes

`func (o *ApiCiConfigurationDto) GetModuleExcludes() []string`

GetModuleExcludes returns the ModuleExcludes field if non-nil, zero value otherwise.

### GetModuleExcludesOk

`func (o *ApiCiConfigurationDto) GetModuleExcludesOk() (*[]string, bool)`

GetModuleExcludesOk returns a tuple with the ModuleExcludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleExcludes

`func (o *ApiCiConfigurationDto) SetModuleExcludes(v []string)`

SetModuleExcludes sets ModuleExcludes field to given value.

### HasModuleExcludes

`func (o *ApiCiConfigurationDto) HasModuleExcludes() bool`

HasModuleExcludes returns a boolean if a field has been set.

### GetParameterPriority

`func (o *ApiCiConfigurationDto) GetParameterPriority() string`

GetParameterPriority returns the ParameterPriority field if non-nil, zero value otherwise.

### GetParameterPriorityOk

`func (o *ApiCiConfigurationDto) GetParameterPriorityOk() (*string, bool)`

GetParameterPriorityOk returns a tuple with the ParameterPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterPriority

`func (o *ApiCiConfigurationDto) SetParameterPriority(v string)`

SetParameterPriority sets ParameterPriority field to given value.

### HasParameterPriority

`func (o *ApiCiConfigurationDto) HasParameterPriority() bool`

HasParameterPriority returns a boolean if a field has been set.

### GetProxy

`func (o *ApiCiConfigurationDto) GetProxy() ProxyConfig`

GetProxy returns the Proxy field if non-nil, zero value otherwise.

### GetProxyOk

`func (o *ApiCiConfigurationDto) GetProxyOk() (*ProxyConfig, bool)`

GetProxyOk returns a tuple with the Proxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxy

`func (o *ApiCiConfigurationDto) SetProxy(v ProxyConfig)`

SetProxy sets Proxy field to given value.

### HasProxy

`func (o *ApiCiConfigurationDto) HasProxy() bool`

HasProxy returns a boolean if a field has been set.

### GetReachability

`func (o *ApiCiConfigurationDto) GetReachability() ReachabilityConfig`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *ApiCiConfigurationDto) GetReachabilityOk() (*ReachabilityConfig, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *ApiCiConfigurationDto) SetReachability(v ReachabilityConfig)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *ApiCiConfigurationDto) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetResultFile

`func (o *ApiCiConfigurationDto) GetResultFile() string`

GetResultFile returns the ResultFile field if non-nil, zero value otherwise.

### GetResultFileOk

`func (o *ApiCiConfigurationDto) GetResultFileOk() (*string, bool)`

GetResultFileOk returns a tuple with the ResultFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultFile

`func (o *ApiCiConfigurationDto) SetResultFile(v string)`

SetResultFile sets ResultFile field to given value.

### HasResultFile

`func (o *ApiCiConfigurationDto) HasResultFile() bool`

HasResultFile returns a boolean if a field has been set.

### GetSarifFile

`func (o *ApiCiConfigurationDto) GetSarifFile() string`

GetSarifFile returns the SarifFile field if non-nil, zero value otherwise.

### GetSarifFileOk

`func (o *ApiCiConfigurationDto) GetSarifFileOk() (*string, bool)`

GetSarifFileOk returns a tuple with the SarifFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSarifFile

`func (o *ApiCiConfigurationDto) SetSarifFile(v string)`

SetSarifFile sets SarifFile field to given value.

### HasSarifFile

`func (o *ApiCiConfigurationDto) HasSarifFile() bool`

HasSarifFile returns a boolean if a field has been set.

### GetScanPatterns

`func (o *ApiCiConfigurationDto) GetScanPatterns() []string`

GetScanPatterns returns the ScanPatterns field if non-nil, zero value otherwise.

### GetScanPatternsOk

`func (o *ApiCiConfigurationDto) GetScanPatternsOk() (*[]string, bool)`

GetScanPatternsOk returns a tuple with the ScanPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanPatterns

`func (o *ApiCiConfigurationDto) SetScanPatterns(v []string)`

SetScanPatterns sets ScanPatterns field to given value.

### HasScanPatterns

`func (o *ApiCiConfigurationDto) HasScanPatterns() bool`

HasScanPatterns returns a boolean if a field has been set.

### GetUnstableBuildOnPolicyWarnings

`func (o *ApiCiConfigurationDto) GetUnstableBuildOnPolicyWarnings() bool`

GetUnstableBuildOnPolicyWarnings returns the UnstableBuildOnPolicyWarnings field if non-nil, zero value otherwise.

### GetUnstableBuildOnPolicyWarningsOk

`func (o *ApiCiConfigurationDto) GetUnstableBuildOnPolicyWarningsOk() (*bool, bool)`

GetUnstableBuildOnPolicyWarningsOk returns a tuple with the UnstableBuildOnPolicyWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnstableBuildOnPolicyWarnings

`func (o *ApiCiConfigurationDto) SetUnstableBuildOnPolicyWarnings(v bool)`

SetUnstableBuildOnPolicyWarnings sets UnstableBuildOnPolicyWarnings field to given value.

### HasUnstableBuildOnPolicyWarnings

`func (o *ApiCiConfigurationDto) HasUnstableBuildOnPolicyWarnings() bool`

HasUnstableBuildOnPolicyWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


