/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.193.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the PolicyEvaluationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyEvaluationResult{}

// PolicyEvaluationResult struct for PolicyEvaluationResult
type PolicyEvaluationResult struct {
	AffectedComponentCount *int32 `json:"affectedComponentCount,omitempty"`
	Alerts []PolicyAlert `json:"alerts,omitempty"`
	CriticalComponentCount *int32 `json:"criticalComponentCount,omitempty"`
	CriticalPolicyViolationCount *int32 `json:"criticalPolicyViolationCount,omitempty"`
	CriticalSastPolicyViolationCount *int32 `json:"criticalSastPolicyViolationCount,omitempty"`
	GrandfatheredPolicyViolationCount *int32 `json:"grandfatheredPolicyViolationCount,omitempty"`
	LegacyViolationCount *int32 `json:"legacyViolationCount,omitempty"`
	ModerateComponentCount *int32 `json:"moderateComponentCount,omitempty"`
	ModeratePolicyViolationCount *int32 `json:"moderatePolicyViolationCount,omitempty"`
	ModerateSastPolicyViolationCount *int32 `json:"moderateSastPolicyViolationCount,omitempty"`
	SastAlerts []PolicyAlert `json:"sastAlerts,omitempty"`
	SevereComponentCount *int32 `json:"severeComponentCount,omitempty"`
	SeverePolicyViolationCount *int32 `json:"severePolicyViolationCount,omitempty"`
	SevereSastPolicyViolationCount *int32 `json:"severeSastPolicyViolationCount,omitempty"`
	TotalComponentCount *int32 `json:"totalComponentCount,omitempty"`
	TotalSastFindingCount *int32 `json:"totalSastFindingCount,omitempty"`
}

// NewPolicyEvaluationResult instantiates a new PolicyEvaluationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyEvaluationResult() *PolicyEvaluationResult {
	this := PolicyEvaluationResult{}
	return &this
}

// NewPolicyEvaluationResultWithDefaults instantiates a new PolicyEvaluationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyEvaluationResultWithDefaults() *PolicyEvaluationResult {
	this := PolicyEvaluationResult{}
	return &this
}

// GetAffectedComponentCount returns the AffectedComponentCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetAffectedComponentCount() int32 {
	if o == nil || IsNil(o.AffectedComponentCount) {
		var ret int32
		return ret
	}
	return *o.AffectedComponentCount
}

// GetAffectedComponentCountOk returns a tuple with the AffectedComponentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetAffectedComponentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.AffectedComponentCount) {
		return nil, false
	}
	return o.AffectedComponentCount, true
}

// HasAffectedComponentCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasAffectedComponentCount() bool {
	if o != nil && !IsNil(o.AffectedComponentCount) {
		return true
	}

	return false
}

// SetAffectedComponentCount gets a reference to the given int32 and assigns it to the AffectedComponentCount field.
func (o *PolicyEvaluationResult) SetAffectedComponentCount(v int32) {
	o.AffectedComponentCount = &v
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetAlerts() []PolicyAlert {
	if o == nil || IsNil(o.Alerts) {
		var ret []PolicyAlert
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetAlertsOk() ([]PolicyAlert, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []PolicyAlert and assigns it to the Alerts field.
func (o *PolicyEvaluationResult) SetAlerts(v []PolicyAlert) {
	o.Alerts = v
}

// GetCriticalComponentCount returns the CriticalComponentCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetCriticalComponentCount() int32 {
	if o == nil || IsNil(o.CriticalComponentCount) {
		var ret int32
		return ret
	}
	return *o.CriticalComponentCount
}

// GetCriticalComponentCountOk returns a tuple with the CriticalComponentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetCriticalComponentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CriticalComponentCount) {
		return nil, false
	}
	return o.CriticalComponentCount, true
}

// HasCriticalComponentCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasCriticalComponentCount() bool {
	if o != nil && !IsNil(o.CriticalComponentCount) {
		return true
	}

	return false
}

// SetCriticalComponentCount gets a reference to the given int32 and assigns it to the CriticalComponentCount field.
func (o *PolicyEvaluationResult) SetCriticalComponentCount(v int32) {
	o.CriticalComponentCount = &v
}

// GetCriticalPolicyViolationCount returns the CriticalPolicyViolationCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetCriticalPolicyViolationCount() int32 {
	if o == nil || IsNil(o.CriticalPolicyViolationCount) {
		var ret int32
		return ret
	}
	return *o.CriticalPolicyViolationCount
}

// GetCriticalPolicyViolationCountOk returns a tuple with the CriticalPolicyViolationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetCriticalPolicyViolationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CriticalPolicyViolationCount) {
		return nil, false
	}
	return o.CriticalPolicyViolationCount, true
}

// HasCriticalPolicyViolationCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasCriticalPolicyViolationCount() bool {
	if o != nil && !IsNil(o.CriticalPolicyViolationCount) {
		return true
	}

	return false
}

// SetCriticalPolicyViolationCount gets a reference to the given int32 and assigns it to the CriticalPolicyViolationCount field.
func (o *PolicyEvaluationResult) SetCriticalPolicyViolationCount(v int32) {
	o.CriticalPolicyViolationCount = &v
}

// GetCriticalSastPolicyViolationCount returns the CriticalSastPolicyViolationCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetCriticalSastPolicyViolationCount() int32 {
	if o == nil || IsNil(o.CriticalSastPolicyViolationCount) {
		var ret int32
		return ret
	}
	return *o.CriticalSastPolicyViolationCount
}

// GetCriticalSastPolicyViolationCountOk returns a tuple with the CriticalSastPolicyViolationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetCriticalSastPolicyViolationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CriticalSastPolicyViolationCount) {
		return nil, false
	}
	return o.CriticalSastPolicyViolationCount, true
}

// HasCriticalSastPolicyViolationCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasCriticalSastPolicyViolationCount() bool {
	if o != nil && !IsNil(o.CriticalSastPolicyViolationCount) {
		return true
	}

	return false
}

// SetCriticalSastPolicyViolationCount gets a reference to the given int32 and assigns it to the CriticalSastPolicyViolationCount field.
func (o *PolicyEvaluationResult) SetCriticalSastPolicyViolationCount(v int32) {
	o.CriticalSastPolicyViolationCount = &v
}

// GetGrandfatheredPolicyViolationCount returns the GrandfatheredPolicyViolationCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetGrandfatheredPolicyViolationCount() int32 {
	if o == nil || IsNil(o.GrandfatheredPolicyViolationCount) {
		var ret int32
		return ret
	}
	return *o.GrandfatheredPolicyViolationCount
}

// GetGrandfatheredPolicyViolationCountOk returns a tuple with the GrandfatheredPolicyViolationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetGrandfatheredPolicyViolationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.GrandfatheredPolicyViolationCount) {
		return nil, false
	}
	return o.GrandfatheredPolicyViolationCount, true
}

// HasGrandfatheredPolicyViolationCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasGrandfatheredPolicyViolationCount() bool {
	if o != nil && !IsNil(o.GrandfatheredPolicyViolationCount) {
		return true
	}

	return false
}

// SetGrandfatheredPolicyViolationCount gets a reference to the given int32 and assigns it to the GrandfatheredPolicyViolationCount field.
func (o *PolicyEvaluationResult) SetGrandfatheredPolicyViolationCount(v int32) {
	o.GrandfatheredPolicyViolationCount = &v
}

// GetLegacyViolationCount returns the LegacyViolationCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetLegacyViolationCount() int32 {
	if o == nil || IsNil(o.LegacyViolationCount) {
		var ret int32
		return ret
	}
	return *o.LegacyViolationCount
}

// GetLegacyViolationCountOk returns a tuple with the LegacyViolationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetLegacyViolationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.LegacyViolationCount) {
		return nil, false
	}
	return o.LegacyViolationCount, true
}

// HasLegacyViolationCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasLegacyViolationCount() bool {
	if o != nil && !IsNil(o.LegacyViolationCount) {
		return true
	}

	return false
}

// SetLegacyViolationCount gets a reference to the given int32 and assigns it to the LegacyViolationCount field.
func (o *PolicyEvaluationResult) SetLegacyViolationCount(v int32) {
	o.LegacyViolationCount = &v
}

// GetModerateComponentCount returns the ModerateComponentCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetModerateComponentCount() int32 {
	if o == nil || IsNil(o.ModerateComponentCount) {
		var ret int32
		return ret
	}
	return *o.ModerateComponentCount
}

// GetModerateComponentCountOk returns a tuple with the ModerateComponentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetModerateComponentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ModerateComponentCount) {
		return nil, false
	}
	return o.ModerateComponentCount, true
}

// HasModerateComponentCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasModerateComponentCount() bool {
	if o != nil && !IsNil(o.ModerateComponentCount) {
		return true
	}

	return false
}

// SetModerateComponentCount gets a reference to the given int32 and assigns it to the ModerateComponentCount field.
func (o *PolicyEvaluationResult) SetModerateComponentCount(v int32) {
	o.ModerateComponentCount = &v
}

// GetModeratePolicyViolationCount returns the ModeratePolicyViolationCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetModeratePolicyViolationCount() int32 {
	if o == nil || IsNil(o.ModeratePolicyViolationCount) {
		var ret int32
		return ret
	}
	return *o.ModeratePolicyViolationCount
}

// GetModeratePolicyViolationCountOk returns a tuple with the ModeratePolicyViolationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetModeratePolicyViolationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ModeratePolicyViolationCount) {
		return nil, false
	}
	return o.ModeratePolicyViolationCount, true
}

// HasModeratePolicyViolationCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasModeratePolicyViolationCount() bool {
	if o != nil && !IsNil(o.ModeratePolicyViolationCount) {
		return true
	}

	return false
}

// SetModeratePolicyViolationCount gets a reference to the given int32 and assigns it to the ModeratePolicyViolationCount field.
func (o *PolicyEvaluationResult) SetModeratePolicyViolationCount(v int32) {
	o.ModeratePolicyViolationCount = &v
}

// GetModerateSastPolicyViolationCount returns the ModerateSastPolicyViolationCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetModerateSastPolicyViolationCount() int32 {
	if o == nil || IsNil(o.ModerateSastPolicyViolationCount) {
		var ret int32
		return ret
	}
	return *o.ModerateSastPolicyViolationCount
}

// GetModerateSastPolicyViolationCountOk returns a tuple with the ModerateSastPolicyViolationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetModerateSastPolicyViolationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.ModerateSastPolicyViolationCount) {
		return nil, false
	}
	return o.ModerateSastPolicyViolationCount, true
}

// HasModerateSastPolicyViolationCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasModerateSastPolicyViolationCount() bool {
	if o != nil && !IsNil(o.ModerateSastPolicyViolationCount) {
		return true
	}

	return false
}

// SetModerateSastPolicyViolationCount gets a reference to the given int32 and assigns it to the ModerateSastPolicyViolationCount field.
func (o *PolicyEvaluationResult) SetModerateSastPolicyViolationCount(v int32) {
	o.ModerateSastPolicyViolationCount = &v
}

// GetSastAlerts returns the SastAlerts field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetSastAlerts() []PolicyAlert {
	if o == nil || IsNil(o.SastAlerts) {
		var ret []PolicyAlert
		return ret
	}
	return o.SastAlerts
}

// GetSastAlertsOk returns a tuple with the SastAlerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetSastAlertsOk() ([]PolicyAlert, bool) {
	if o == nil || IsNil(o.SastAlerts) {
		return nil, false
	}
	return o.SastAlerts, true
}

// HasSastAlerts returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasSastAlerts() bool {
	if o != nil && !IsNil(o.SastAlerts) {
		return true
	}

	return false
}

// SetSastAlerts gets a reference to the given []PolicyAlert and assigns it to the SastAlerts field.
func (o *PolicyEvaluationResult) SetSastAlerts(v []PolicyAlert) {
	o.SastAlerts = v
}

// GetSevereComponentCount returns the SevereComponentCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetSevereComponentCount() int32 {
	if o == nil || IsNil(o.SevereComponentCount) {
		var ret int32
		return ret
	}
	return *o.SevereComponentCount
}

// GetSevereComponentCountOk returns a tuple with the SevereComponentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetSevereComponentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SevereComponentCount) {
		return nil, false
	}
	return o.SevereComponentCount, true
}

// HasSevereComponentCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasSevereComponentCount() bool {
	if o != nil && !IsNil(o.SevereComponentCount) {
		return true
	}

	return false
}

// SetSevereComponentCount gets a reference to the given int32 and assigns it to the SevereComponentCount field.
func (o *PolicyEvaluationResult) SetSevereComponentCount(v int32) {
	o.SevereComponentCount = &v
}

// GetSeverePolicyViolationCount returns the SeverePolicyViolationCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetSeverePolicyViolationCount() int32 {
	if o == nil || IsNil(o.SeverePolicyViolationCount) {
		var ret int32
		return ret
	}
	return *o.SeverePolicyViolationCount
}

// GetSeverePolicyViolationCountOk returns a tuple with the SeverePolicyViolationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetSeverePolicyViolationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SeverePolicyViolationCount) {
		return nil, false
	}
	return o.SeverePolicyViolationCount, true
}

// HasSeverePolicyViolationCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasSeverePolicyViolationCount() bool {
	if o != nil && !IsNil(o.SeverePolicyViolationCount) {
		return true
	}

	return false
}

// SetSeverePolicyViolationCount gets a reference to the given int32 and assigns it to the SeverePolicyViolationCount field.
func (o *PolicyEvaluationResult) SetSeverePolicyViolationCount(v int32) {
	o.SeverePolicyViolationCount = &v
}

// GetSevereSastPolicyViolationCount returns the SevereSastPolicyViolationCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetSevereSastPolicyViolationCount() int32 {
	if o == nil || IsNil(o.SevereSastPolicyViolationCount) {
		var ret int32
		return ret
	}
	return *o.SevereSastPolicyViolationCount
}

// GetSevereSastPolicyViolationCountOk returns a tuple with the SevereSastPolicyViolationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetSevereSastPolicyViolationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SevereSastPolicyViolationCount) {
		return nil, false
	}
	return o.SevereSastPolicyViolationCount, true
}

// HasSevereSastPolicyViolationCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasSevereSastPolicyViolationCount() bool {
	if o != nil && !IsNil(o.SevereSastPolicyViolationCount) {
		return true
	}

	return false
}

// SetSevereSastPolicyViolationCount gets a reference to the given int32 and assigns it to the SevereSastPolicyViolationCount field.
func (o *PolicyEvaluationResult) SetSevereSastPolicyViolationCount(v int32) {
	o.SevereSastPolicyViolationCount = &v
}

// GetTotalComponentCount returns the TotalComponentCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetTotalComponentCount() int32 {
	if o == nil || IsNil(o.TotalComponentCount) {
		var ret int32
		return ret
	}
	return *o.TotalComponentCount
}

// GetTotalComponentCountOk returns a tuple with the TotalComponentCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetTotalComponentCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalComponentCount) {
		return nil, false
	}
	return o.TotalComponentCount, true
}

// HasTotalComponentCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasTotalComponentCount() bool {
	if o != nil && !IsNil(o.TotalComponentCount) {
		return true
	}

	return false
}

// SetTotalComponentCount gets a reference to the given int32 and assigns it to the TotalComponentCount field.
func (o *PolicyEvaluationResult) SetTotalComponentCount(v int32) {
	o.TotalComponentCount = &v
}

// GetTotalSastFindingCount returns the TotalSastFindingCount field value if set, zero value otherwise.
func (o *PolicyEvaluationResult) GetTotalSastFindingCount() int32 {
	if o == nil || IsNil(o.TotalSastFindingCount) {
		var ret int32
		return ret
	}
	return *o.TotalSastFindingCount
}

// GetTotalSastFindingCountOk returns a tuple with the TotalSastFindingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyEvaluationResult) GetTotalSastFindingCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalSastFindingCount) {
		return nil, false
	}
	return o.TotalSastFindingCount, true
}

// HasTotalSastFindingCount returns a boolean if a field has been set.
func (o *PolicyEvaluationResult) HasTotalSastFindingCount() bool {
	if o != nil && !IsNil(o.TotalSastFindingCount) {
		return true
	}

	return false
}

// SetTotalSastFindingCount gets a reference to the given int32 and assigns it to the TotalSastFindingCount field.
func (o *PolicyEvaluationResult) SetTotalSastFindingCount(v int32) {
	o.TotalSastFindingCount = &v
}

func (o PolicyEvaluationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyEvaluationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffectedComponentCount) {
		toSerialize["affectedComponentCount"] = o.AffectedComponentCount
	}
	if !IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !IsNil(o.CriticalComponentCount) {
		toSerialize["criticalComponentCount"] = o.CriticalComponentCount
	}
	if !IsNil(o.CriticalPolicyViolationCount) {
		toSerialize["criticalPolicyViolationCount"] = o.CriticalPolicyViolationCount
	}
	if !IsNil(o.CriticalSastPolicyViolationCount) {
		toSerialize["criticalSastPolicyViolationCount"] = o.CriticalSastPolicyViolationCount
	}
	if !IsNil(o.GrandfatheredPolicyViolationCount) {
		toSerialize["grandfatheredPolicyViolationCount"] = o.GrandfatheredPolicyViolationCount
	}
	if !IsNil(o.LegacyViolationCount) {
		toSerialize["legacyViolationCount"] = o.LegacyViolationCount
	}
	if !IsNil(o.ModerateComponentCount) {
		toSerialize["moderateComponentCount"] = o.ModerateComponentCount
	}
	if !IsNil(o.ModeratePolicyViolationCount) {
		toSerialize["moderatePolicyViolationCount"] = o.ModeratePolicyViolationCount
	}
	if !IsNil(o.ModerateSastPolicyViolationCount) {
		toSerialize["moderateSastPolicyViolationCount"] = o.ModerateSastPolicyViolationCount
	}
	if !IsNil(o.SastAlerts) {
		toSerialize["sastAlerts"] = o.SastAlerts
	}
	if !IsNil(o.SevereComponentCount) {
		toSerialize["severeComponentCount"] = o.SevereComponentCount
	}
	if !IsNil(o.SeverePolicyViolationCount) {
		toSerialize["severePolicyViolationCount"] = o.SeverePolicyViolationCount
	}
	if !IsNil(o.SevereSastPolicyViolationCount) {
		toSerialize["severeSastPolicyViolationCount"] = o.SevereSastPolicyViolationCount
	}
	if !IsNil(o.TotalComponentCount) {
		toSerialize["totalComponentCount"] = o.TotalComponentCount
	}
	if !IsNil(o.TotalSastFindingCount) {
		toSerialize["totalSastFindingCount"] = o.TotalSastFindingCount
	}
	return toSerialize, nil
}

type NullablePolicyEvaluationResult struct {
	value *PolicyEvaluationResult
	isSet bool
}

func (v NullablePolicyEvaluationResult) Get() *PolicyEvaluationResult {
	return v.value
}

func (v *NullablePolicyEvaluationResult) Set(val *PolicyEvaluationResult) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyEvaluationResult) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyEvaluationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyEvaluationResult(val *PolicyEvaluationResult) *NullablePolicyEvaluationResult {
	return &NullablePolicyEvaluationResult{value: val, isSet: true}
}

func (v NullablePolicyEvaluationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyEvaluationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


