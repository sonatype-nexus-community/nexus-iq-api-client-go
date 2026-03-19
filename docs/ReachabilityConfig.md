# ReachabilityConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailOnError** | Pointer to **bool** |  | [optional] 
**JavaAnalysis** | Pointer to [**JavaAnalysisConfig**](JavaAnalysisConfig.md) |  | [optional] 
**JavaScriptAnalysis** | Pointer to [**JavaScriptAnalysisConfig**](JavaScriptAnalysisConfig.md) |  | [optional] 

## Methods

### NewReachabilityConfig

`func NewReachabilityConfig() *ReachabilityConfig`

NewReachabilityConfig instantiates a new ReachabilityConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReachabilityConfigWithDefaults

`func NewReachabilityConfigWithDefaults() *ReachabilityConfig`

NewReachabilityConfigWithDefaults instantiates a new ReachabilityConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailOnError

`func (o *ReachabilityConfig) GetFailOnError() bool`

GetFailOnError returns the FailOnError field if non-nil, zero value otherwise.

### GetFailOnErrorOk

`func (o *ReachabilityConfig) GetFailOnErrorOk() (*bool, bool)`

GetFailOnErrorOk returns a tuple with the FailOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailOnError

`func (o *ReachabilityConfig) SetFailOnError(v bool)`

SetFailOnError sets FailOnError field to given value.

### HasFailOnError

`func (o *ReachabilityConfig) HasFailOnError() bool`

HasFailOnError returns a boolean if a field has been set.

### GetJavaAnalysis

`func (o *ReachabilityConfig) GetJavaAnalysis() JavaAnalysisConfig`

GetJavaAnalysis returns the JavaAnalysis field if non-nil, zero value otherwise.

### GetJavaAnalysisOk

`func (o *ReachabilityConfig) GetJavaAnalysisOk() (*JavaAnalysisConfig, bool)`

GetJavaAnalysisOk returns a tuple with the JavaAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaAnalysis

`func (o *ReachabilityConfig) SetJavaAnalysis(v JavaAnalysisConfig)`

SetJavaAnalysis sets JavaAnalysis field to given value.

### HasJavaAnalysis

`func (o *ReachabilityConfig) HasJavaAnalysis() bool`

HasJavaAnalysis returns a boolean if a field has been set.

### GetJavaScriptAnalysis

`func (o *ReachabilityConfig) GetJavaScriptAnalysis() JavaScriptAnalysisConfig`

GetJavaScriptAnalysis returns the JavaScriptAnalysis field if non-nil, zero value otherwise.

### GetJavaScriptAnalysisOk

`func (o *ReachabilityConfig) GetJavaScriptAnalysisOk() (*JavaScriptAnalysisConfig, bool)`

GetJavaScriptAnalysisOk returns a tuple with the JavaScriptAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJavaScriptAnalysis

`func (o *ReachabilityConfig) SetJavaScriptAnalysis(v JavaScriptAnalysisConfig)`

SetJavaScriptAnalysis sets JavaScriptAnalysis field to given value.

### HasJavaScriptAnalysis

`func (o *ReachabilityConfig) HasJavaScriptAnalysis() bool`

HasJavaScriptAnalysis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


