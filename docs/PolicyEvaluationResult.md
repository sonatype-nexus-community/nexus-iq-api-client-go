# PolicyEvaluationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedComponentCount** | Pointer to **int32** |  | [optional] 
**Alerts** | Pointer to [**[]PolicyAlert**](PolicyAlert.md) |  | [optional] 
**CriticalComponentCount** | Pointer to **int32** |  | [optional] 
**CriticalPolicyViolationCount** | Pointer to **int32** |  | [optional] 
**CriticalSastPolicyViolationCount** | Pointer to **int32** |  | [optional] 
**GrandfatheredPolicyViolationCount** | Pointer to **int32** |  | [optional] 
**LegacyViolationCount** | Pointer to **int32** |  | [optional] 
**ModerateComponentCount** | Pointer to **int32** |  | [optional] 
**ModeratePolicyViolationCount** | Pointer to **int32** |  | [optional] 
**ModerateSastPolicyViolationCount** | Pointer to **int32** |  | [optional] 
**SastAlerts** | Pointer to [**[]PolicyAlert**](PolicyAlert.md) |  | [optional] 
**SevereComponentCount** | Pointer to **int32** |  | [optional] 
**SeverePolicyViolationCount** | Pointer to **int32** |  | [optional] 
**SevereSastPolicyViolationCount** | Pointer to **int32** |  | [optional] 
**TotalComponentCount** | Pointer to **int32** |  | [optional] 
**TotalSastFindingCount** | Pointer to **int32** |  | [optional] 

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

### GetCriticalSastPolicyViolationCount

`func (o *PolicyEvaluationResult) GetCriticalSastPolicyViolationCount() int32`

GetCriticalSastPolicyViolationCount returns the CriticalSastPolicyViolationCount field if non-nil, zero value otherwise.

### GetCriticalSastPolicyViolationCountOk

`func (o *PolicyEvaluationResult) GetCriticalSastPolicyViolationCountOk() (*int32, bool)`

GetCriticalSastPolicyViolationCountOk returns a tuple with the CriticalSastPolicyViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalSastPolicyViolationCount

`func (o *PolicyEvaluationResult) SetCriticalSastPolicyViolationCount(v int32)`

SetCriticalSastPolicyViolationCount sets CriticalSastPolicyViolationCount field to given value.

### HasCriticalSastPolicyViolationCount

`func (o *PolicyEvaluationResult) HasCriticalSastPolicyViolationCount() bool`

HasCriticalSastPolicyViolationCount returns a boolean if a field has been set.

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

### GetModerateSastPolicyViolationCount

`func (o *PolicyEvaluationResult) GetModerateSastPolicyViolationCount() int32`

GetModerateSastPolicyViolationCount returns the ModerateSastPolicyViolationCount field if non-nil, zero value otherwise.

### GetModerateSastPolicyViolationCountOk

`func (o *PolicyEvaluationResult) GetModerateSastPolicyViolationCountOk() (*int32, bool)`

GetModerateSastPolicyViolationCountOk returns a tuple with the ModerateSastPolicyViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerateSastPolicyViolationCount

`func (o *PolicyEvaluationResult) SetModerateSastPolicyViolationCount(v int32)`

SetModerateSastPolicyViolationCount sets ModerateSastPolicyViolationCount field to given value.

### HasModerateSastPolicyViolationCount

`func (o *PolicyEvaluationResult) HasModerateSastPolicyViolationCount() bool`

HasModerateSastPolicyViolationCount returns a boolean if a field has been set.

### GetSastAlerts

`func (o *PolicyEvaluationResult) GetSastAlerts() []PolicyAlert`

GetSastAlerts returns the SastAlerts field if non-nil, zero value otherwise.

### GetSastAlertsOk

`func (o *PolicyEvaluationResult) GetSastAlertsOk() (*[]PolicyAlert, bool)`

GetSastAlertsOk returns a tuple with the SastAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSastAlerts

`func (o *PolicyEvaluationResult) SetSastAlerts(v []PolicyAlert)`

SetSastAlerts sets SastAlerts field to given value.

### HasSastAlerts

`func (o *PolicyEvaluationResult) HasSastAlerts() bool`

HasSastAlerts returns a boolean if a field has been set.

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

### GetSevereSastPolicyViolationCount

`func (o *PolicyEvaluationResult) GetSevereSastPolicyViolationCount() int32`

GetSevereSastPolicyViolationCount returns the SevereSastPolicyViolationCount field if non-nil, zero value otherwise.

### GetSevereSastPolicyViolationCountOk

`func (o *PolicyEvaluationResult) GetSevereSastPolicyViolationCountOk() (*int32, bool)`

GetSevereSastPolicyViolationCountOk returns a tuple with the SevereSastPolicyViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSevereSastPolicyViolationCount

`func (o *PolicyEvaluationResult) SetSevereSastPolicyViolationCount(v int32)`

SetSevereSastPolicyViolationCount sets SevereSastPolicyViolationCount field to given value.

### HasSevereSastPolicyViolationCount

`func (o *PolicyEvaluationResult) HasSevereSastPolicyViolationCount() bool`

HasSevereSastPolicyViolationCount returns a boolean if a field has been set.

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

### GetTotalSastFindingCount

`func (o *PolicyEvaluationResult) GetTotalSastFindingCount() int32`

GetTotalSastFindingCount returns the TotalSastFindingCount field if non-nil, zero value otherwise.

### GetTotalSastFindingCountOk

`func (o *PolicyEvaluationResult) GetTotalSastFindingCountOk() (*int32, bool)`

GetTotalSastFindingCountOk returns a tuple with the TotalSastFindingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSastFindingCount

`func (o *PolicyEvaluationResult) SetTotalSastFindingCount(v int32)`

SetTotalSastFindingCount sets TotalSastFindingCount field to given value.

### HasTotalSastFindingCount

`func (o *PolicyEvaluationResult) HasTotalSastFindingCount() bool`

HasTotalSastFindingCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


