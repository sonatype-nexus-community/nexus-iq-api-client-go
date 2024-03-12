# PolicyEvaluationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedComponentCount** | Pointer to **int32** |  | [optional] 
**Alerts** | Pointer to [**[]PolicyAlert**](PolicyAlert.md) |  | [optional] 
**CriticalComponentCount** | Pointer to **int32** |  | [optional] 
**CriticalPolicyViolationCount** | Pointer to **int32** |  | [optional] 
**GrandfatheredPolicyViolationCount** | Pointer to **int32** |  | [optional] 
**LegacyViolationCount** | Pointer to **int32** |  | [optional] 
**ModerateComponentCount** | Pointer to **int32** |  | [optional] 
**ModeratePolicyViolationCount** | Pointer to **int32** |  | [optional] 
**SevereComponentCount** | Pointer to **int32** |  | [optional] 
**SeverePolicyViolationCount** | Pointer to **int32** |  | [optional] 
**TotalComponentCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewPolicyEvaluationResult

`func NewPolicyEvaluationResult() *PolicyEvaluationResult`

NewPolicyEvaluationResult instantiates a new PolicyEvaluationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyEvaluationResultWithDefaults

`func NewPolicyEvaluationResultWithDefaults() *PolicyEvaluationResult`

NewPolicyEvaluationResultWithDefaults instantiates a new PolicyEvaluationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedComponentCount

`func (o *PolicyEvaluationResult) GetAffectedComponentCount() int32`

GetAffectedComponentCount returns the AffectedComponentCount field if non-nil, zero value otherwise.

### GetAffectedComponentCountOk

`func (o *PolicyEvaluationResult) GetAffectedComponentCountOk() (*int32, bool)`

GetAffectedComponentCountOk returns a tuple with the AffectedComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedComponentCount

`func (o *PolicyEvaluationResult) SetAffectedComponentCount(v int32)`

SetAffectedComponentCount sets AffectedComponentCount field to given value.

### HasAffectedComponentCount

`func (o *PolicyEvaluationResult) HasAffectedComponentCount() bool`

HasAffectedComponentCount returns a boolean if a field has been set.

### GetAlerts

`func (o *PolicyEvaluationResult) GetAlerts() []PolicyAlert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *PolicyEvaluationResult) GetAlertsOk() (*[]PolicyAlert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *PolicyEvaluationResult) SetAlerts(v []PolicyAlert)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *PolicyEvaluationResult) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetCriticalComponentCount

`func (o *PolicyEvaluationResult) GetCriticalComponentCount() int32`

GetCriticalComponentCount returns the CriticalComponentCount field if non-nil, zero value otherwise.

### GetCriticalComponentCountOk

`func (o *PolicyEvaluationResult) GetCriticalComponentCountOk() (*int32, bool)`

GetCriticalComponentCountOk returns a tuple with the CriticalComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalComponentCount

`func (o *PolicyEvaluationResult) SetCriticalComponentCount(v int32)`

SetCriticalComponentCount sets CriticalComponentCount field to given value.

### HasCriticalComponentCount

`func (o *PolicyEvaluationResult) HasCriticalComponentCount() bool`

HasCriticalComponentCount returns a boolean if a field has been set.

### GetCriticalPolicyViolationCount

`func (o *PolicyEvaluationResult) GetCriticalPolicyViolationCount() int32`

GetCriticalPolicyViolationCount returns the CriticalPolicyViolationCount field if non-nil, zero value otherwise.

### GetCriticalPolicyViolationCountOk

`func (o *PolicyEvaluationResult) GetCriticalPolicyViolationCountOk() (*int32, bool)`

GetCriticalPolicyViolationCountOk returns a tuple with the CriticalPolicyViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalPolicyViolationCount

`func (o *PolicyEvaluationResult) SetCriticalPolicyViolationCount(v int32)`

SetCriticalPolicyViolationCount sets CriticalPolicyViolationCount field to given value.

### HasCriticalPolicyViolationCount

`func (o *PolicyEvaluationResult) HasCriticalPolicyViolationCount() bool`

HasCriticalPolicyViolationCount returns a boolean if a field has been set.

### GetGrandfatheredPolicyViolationCount

`func (o *PolicyEvaluationResult) GetGrandfatheredPolicyViolationCount() int32`

GetGrandfatheredPolicyViolationCount returns the GrandfatheredPolicyViolationCount field if non-nil, zero value otherwise.

### GetGrandfatheredPolicyViolationCountOk

`func (o *PolicyEvaluationResult) GetGrandfatheredPolicyViolationCountOk() (*int32, bool)`

GetGrandfatheredPolicyViolationCountOk returns a tuple with the GrandfatheredPolicyViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandfatheredPolicyViolationCount

`func (o *PolicyEvaluationResult) SetGrandfatheredPolicyViolationCount(v int32)`

SetGrandfatheredPolicyViolationCount sets GrandfatheredPolicyViolationCount field to given value.

### HasGrandfatheredPolicyViolationCount

`func (o *PolicyEvaluationResult) HasGrandfatheredPolicyViolationCount() bool`

HasGrandfatheredPolicyViolationCount returns a boolean if a field has been set.

### GetLegacyViolationCount

`func (o *PolicyEvaluationResult) GetLegacyViolationCount() int32`

GetLegacyViolationCount returns the LegacyViolationCount field if non-nil, zero value otherwise.

### GetLegacyViolationCountOk

`func (o *PolicyEvaluationResult) GetLegacyViolationCountOk() (*int32, bool)`

GetLegacyViolationCountOk returns a tuple with the LegacyViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyViolationCount

`func (o *PolicyEvaluationResult) SetLegacyViolationCount(v int32)`

SetLegacyViolationCount sets LegacyViolationCount field to given value.

### HasLegacyViolationCount

`func (o *PolicyEvaluationResult) HasLegacyViolationCount() bool`

HasLegacyViolationCount returns a boolean if a field has been set.

### GetModerateComponentCount

`func (o *PolicyEvaluationResult) GetModerateComponentCount() int32`

GetModerateComponentCount returns the ModerateComponentCount field if non-nil, zero value otherwise.

### GetModerateComponentCountOk

`func (o *PolicyEvaluationResult) GetModerateComponentCountOk() (*int32, bool)`

GetModerateComponentCountOk returns a tuple with the ModerateComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerateComponentCount

`func (o *PolicyEvaluationResult) SetModerateComponentCount(v int32)`

SetModerateComponentCount sets ModerateComponentCount field to given value.

### HasModerateComponentCount

`func (o *PolicyEvaluationResult) HasModerateComponentCount() bool`

HasModerateComponentCount returns a boolean if a field has been set.

### GetModeratePolicyViolationCount

`func (o *PolicyEvaluationResult) GetModeratePolicyViolationCount() int32`

GetModeratePolicyViolationCount returns the ModeratePolicyViolationCount field if non-nil, zero value otherwise.

### GetModeratePolicyViolationCountOk

`func (o *PolicyEvaluationResult) GetModeratePolicyViolationCountOk() (*int32, bool)`

GetModeratePolicyViolationCountOk returns a tuple with the ModeratePolicyViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeratePolicyViolationCount

`func (o *PolicyEvaluationResult) SetModeratePolicyViolationCount(v int32)`

SetModeratePolicyViolationCount sets ModeratePolicyViolationCount field to given value.

### HasModeratePolicyViolationCount

`func (o *PolicyEvaluationResult) HasModeratePolicyViolationCount() bool`

HasModeratePolicyViolationCount returns a boolean if a field has been set.

### GetSevereComponentCount

`func (o *PolicyEvaluationResult) GetSevereComponentCount() int32`

GetSevereComponentCount returns the SevereComponentCount field if non-nil, zero value otherwise.

### GetSevereComponentCountOk

`func (o *PolicyEvaluationResult) GetSevereComponentCountOk() (*int32, bool)`

GetSevereComponentCountOk returns a tuple with the SevereComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevereComponentCount

`func (o *PolicyEvaluationResult) SetSevereComponentCount(v int32)`

SetSevereComponentCount sets SevereComponentCount field to given value.

### HasSevereComponentCount

`func (o *PolicyEvaluationResult) HasSevereComponentCount() bool`

HasSevereComponentCount returns a boolean if a field has been set.

### GetSeverePolicyViolationCount

`func (o *PolicyEvaluationResult) GetSeverePolicyViolationCount() int32`

GetSeverePolicyViolationCount returns the SeverePolicyViolationCount field if non-nil, zero value otherwise.

### GetSeverePolicyViolationCountOk

`func (o *PolicyEvaluationResult) GetSeverePolicyViolationCountOk() (*int32, bool)`

GetSeverePolicyViolationCountOk returns a tuple with the SeverePolicyViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverePolicyViolationCount

`func (o *PolicyEvaluationResult) SetSeverePolicyViolationCount(v int32)`

SetSeverePolicyViolationCount sets SeverePolicyViolationCount field to given value.

### HasSeverePolicyViolationCount

`func (o *PolicyEvaluationResult) HasSeverePolicyViolationCount() bool`

HasSeverePolicyViolationCount returns a boolean if a field has been set.

### GetTotalComponentCount

`func (o *PolicyEvaluationResult) GetTotalComponentCount() int32`

GetTotalComponentCount returns the TotalComponentCount field if non-nil, zero value otherwise.

### GetTotalComponentCountOk

`func (o *PolicyEvaluationResult) GetTotalComponentCountOk() (*int32, bool)`

GetTotalComponentCountOk returns a tuple with the TotalComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalComponentCount

`func (o *PolicyEvaluationResult) SetTotalComponentCount(v int32)`

SetTotalComponentCount sets TotalComponentCount field to given value.

### HasTotalComponentCount

`func (o *PolicyEvaluationResult) HasTotalComponentCount() bool`

HasTotalComponentCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


