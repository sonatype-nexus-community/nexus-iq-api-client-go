/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.177.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ConditionFact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionFact{}

// ConditionFact struct for ConditionFact
type ConditionFact struct {
	ConditionIndex *int32 `json:"conditionIndex,omitempty"`
	ConditionTypeId *string `json:"conditionTypeId,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Reference *TriggerReference `json:"reference,omitempty"`
	Summary *string `json:"summary,omitempty"`
	TriggerJson *string `json:"triggerJson,omitempty"`
}

// NewConditionFact instantiates a new ConditionFact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionFact() *ConditionFact {
	this := ConditionFact{}
	return &this
}

// NewConditionFactWithDefaults instantiates a new ConditionFact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionFactWithDefaults() *ConditionFact {
	this := ConditionFact{}
	return &this
}

// GetConditionIndex returns the ConditionIndex field value if set, zero value otherwise.
func (o *ConditionFact) GetConditionIndex() int32 {
	if o == nil || IsNil(o.ConditionIndex) {
		var ret int32
		return ret
	}
	return *o.ConditionIndex
}

// GetConditionIndexOk returns a tuple with the ConditionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionFact) GetConditionIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.ConditionIndex) {
		return nil, false
	}
	return o.ConditionIndex, true
}

// HasConditionIndex returns a boolean if a field has been set.
func (o *ConditionFact) HasConditionIndex() bool {
	if o != nil && !IsNil(o.ConditionIndex) {
		return true
	}

	return false
}

// SetConditionIndex gets a reference to the given int32 and assigns it to the ConditionIndex field.
func (o *ConditionFact) SetConditionIndex(v int32) {
	o.ConditionIndex = &v
}

// GetConditionTypeId returns the ConditionTypeId field value if set, zero value otherwise.
func (o *ConditionFact) GetConditionTypeId() string {
	if o == nil || IsNil(o.ConditionTypeId) {
		var ret string
		return ret
	}
	return *o.ConditionTypeId
}

// GetConditionTypeIdOk returns a tuple with the ConditionTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionFact) GetConditionTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConditionTypeId) {
		return nil, false
	}
	return o.ConditionTypeId, true
}

// HasConditionTypeId returns a boolean if a field has been set.
func (o *ConditionFact) HasConditionTypeId() bool {
	if o != nil && !IsNil(o.ConditionTypeId) {
		return true
	}

	return false
}

// SetConditionTypeId gets a reference to the given string and assigns it to the ConditionTypeId field.
func (o *ConditionFact) SetConditionTypeId(v string) {
	o.ConditionTypeId = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ConditionFact) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionFact) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ConditionFact) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ConditionFact) SetReason(v string) {
	o.Reason = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ConditionFact) GetReference() TriggerReference {
	if o == nil || IsNil(o.Reference) {
		var ret TriggerReference
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionFact) GetReferenceOk() (*TriggerReference, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ConditionFact) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given TriggerReference and assigns it to the Reference field.
func (o *ConditionFact) SetReference(v TriggerReference) {
	o.Reference = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *ConditionFact) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionFact) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *ConditionFact) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *ConditionFact) SetSummary(v string) {
	o.Summary = &v
}

// GetTriggerJson returns the TriggerJson field value if set, zero value otherwise.
func (o *ConditionFact) GetTriggerJson() string {
	if o == nil || IsNil(o.TriggerJson) {
		var ret string
		return ret
	}
	return *o.TriggerJson
}

// GetTriggerJsonOk returns a tuple with the TriggerJson field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionFact) GetTriggerJsonOk() (*string, bool) {
	if o == nil || IsNil(o.TriggerJson) {
		return nil, false
	}
	return o.TriggerJson, true
}

// HasTriggerJson returns a boolean if a field has been set.
func (o *ConditionFact) HasTriggerJson() bool {
	if o != nil && !IsNil(o.TriggerJson) {
		return true
	}

	return false
}

// SetTriggerJson gets a reference to the given string and assigns it to the TriggerJson field.
func (o *ConditionFact) SetTriggerJson(v string) {
	o.TriggerJson = &v
}

func (o ConditionFact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionFact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConditionIndex) {
		toSerialize["conditionIndex"] = o.ConditionIndex
	}
	if !IsNil(o.ConditionTypeId) {
		toSerialize["conditionTypeId"] = o.ConditionTypeId
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.TriggerJson) {
		toSerialize["triggerJson"] = o.TriggerJson
	}
	return toSerialize, nil
}

type NullableConditionFact struct {
	value *ConditionFact
	isSet bool
}

func (v NullableConditionFact) Get() *ConditionFact {
	return v.value
}

func (v *NullableConditionFact) Set(val *ConditionFact) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionFact) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionFact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionFact(val *ConditionFact) *NullableConditionFact {
	return &NullableConditionFact{value: val, isSet: true}
}

func (v NullableConditionFact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionFact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


