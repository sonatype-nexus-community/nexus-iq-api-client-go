/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.184.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiLicenseDataDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicenseDataDTOV2{}

// ApiLicenseDataDTOV2 struct for ApiLicenseDataDTOV2
type ApiLicenseDataDTOV2 struct {
	DeclaredLicenses []ApiLicenseDTO `json:"declaredLicenses,omitempty"`
	EffectiveLicenseThreats []ApiLicenseThreatDTOV2 `json:"effectiveLicenseThreats,omitempty"`
	EffectiveLicenses []ApiLicenseDTO `json:"effectiveLicenses,omitempty"`
	ObservedLicenses []ApiLicenseDTO `json:"observedLicenses,omitempty"`
	OverriddenLicenses []ApiLicenseDTO `json:"overriddenLicenses,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewApiLicenseDataDTOV2 instantiates a new ApiLicenseDataDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicenseDataDTOV2() *ApiLicenseDataDTOV2 {
	this := ApiLicenseDataDTOV2{}
	return &this
}

// NewApiLicenseDataDTOV2WithDefaults instantiates a new ApiLicenseDataDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicenseDataDTOV2WithDefaults() *ApiLicenseDataDTOV2 {
	this := ApiLicenseDataDTOV2{}
	return &this
}

// GetDeclaredLicenses returns the DeclaredLicenses field value if set, zero value otherwise.
func (o *ApiLicenseDataDTOV2) GetDeclaredLicenses() []ApiLicenseDTO {
	if o == nil || IsNil(o.DeclaredLicenses) {
		var ret []ApiLicenseDTO
		return ret
	}
	return o.DeclaredLicenses
}

// GetDeclaredLicensesOk returns a tuple with the DeclaredLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDataDTOV2) GetDeclaredLicensesOk() ([]ApiLicenseDTO, bool) {
	if o == nil || IsNil(o.DeclaredLicenses) {
		return nil, false
	}
	return o.DeclaredLicenses, true
}

// HasDeclaredLicenses returns a boolean if a field has been set.
func (o *ApiLicenseDataDTOV2) HasDeclaredLicenses() bool {
	if o != nil && !IsNil(o.DeclaredLicenses) {
		return true
	}

	return false
}

// SetDeclaredLicenses gets a reference to the given []ApiLicenseDTO and assigns it to the DeclaredLicenses field.
func (o *ApiLicenseDataDTOV2) SetDeclaredLicenses(v []ApiLicenseDTO) {
	o.DeclaredLicenses = v
}

// GetEffectiveLicenseThreats returns the EffectiveLicenseThreats field value if set, zero value otherwise.
func (o *ApiLicenseDataDTOV2) GetEffectiveLicenseThreats() []ApiLicenseThreatDTOV2 {
	if o == nil || IsNil(o.EffectiveLicenseThreats) {
		var ret []ApiLicenseThreatDTOV2
		return ret
	}
	return o.EffectiveLicenseThreats
}

// GetEffectiveLicenseThreatsOk returns a tuple with the EffectiveLicenseThreats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDataDTOV2) GetEffectiveLicenseThreatsOk() ([]ApiLicenseThreatDTOV2, bool) {
	if o == nil || IsNil(o.EffectiveLicenseThreats) {
		return nil, false
	}
	return o.EffectiveLicenseThreats, true
}

// HasEffectiveLicenseThreats returns a boolean if a field has been set.
func (o *ApiLicenseDataDTOV2) HasEffectiveLicenseThreats() bool {
	if o != nil && !IsNil(o.EffectiveLicenseThreats) {
		return true
	}

	return false
}

// SetEffectiveLicenseThreats gets a reference to the given []ApiLicenseThreatDTOV2 and assigns it to the EffectiveLicenseThreats field.
func (o *ApiLicenseDataDTOV2) SetEffectiveLicenseThreats(v []ApiLicenseThreatDTOV2) {
	o.EffectiveLicenseThreats = v
}

// GetEffectiveLicenses returns the EffectiveLicenses field value if set, zero value otherwise.
func (o *ApiLicenseDataDTOV2) GetEffectiveLicenses() []ApiLicenseDTO {
	if o == nil || IsNil(o.EffectiveLicenses) {
		var ret []ApiLicenseDTO
		return ret
	}
	return o.EffectiveLicenses
}

// GetEffectiveLicensesOk returns a tuple with the EffectiveLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDataDTOV2) GetEffectiveLicensesOk() ([]ApiLicenseDTO, bool) {
	if o == nil || IsNil(o.EffectiveLicenses) {
		return nil, false
	}
	return o.EffectiveLicenses, true
}

// HasEffectiveLicenses returns a boolean if a field has been set.
func (o *ApiLicenseDataDTOV2) HasEffectiveLicenses() bool {
	if o != nil && !IsNil(o.EffectiveLicenses) {
		return true
	}

	return false
}

// SetEffectiveLicenses gets a reference to the given []ApiLicenseDTO and assigns it to the EffectiveLicenses field.
func (o *ApiLicenseDataDTOV2) SetEffectiveLicenses(v []ApiLicenseDTO) {
	o.EffectiveLicenses = v
}

// GetObservedLicenses returns the ObservedLicenses field value if set, zero value otherwise.
func (o *ApiLicenseDataDTOV2) GetObservedLicenses() []ApiLicenseDTO {
	if o == nil || IsNil(o.ObservedLicenses) {
		var ret []ApiLicenseDTO
		return ret
	}
	return o.ObservedLicenses
}

// GetObservedLicensesOk returns a tuple with the ObservedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDataDTOV2) GetObservedLicensesOk() ([]ApiLicenseDTO, bool) {
	if o == nil || IsNil(o.ObservedLicenses) {
		return nil, false
	}
	return o.ObservedLicenses, true
}

// HasObservedLicenses returns a boolean if a field has been set.
func (o *ApiLicenseDataDTOV2) HasObservedLicenses() bool {
	if o != nil && !IsNil(o.ObservedLicenses) {
		return true
	}

	return false
}

// SetObservedLicenses gets a reference to the given []ApiLicenseDTO and assigns it to the ObservedLicenses field.
func (o *ApiLicenseDataDTOV2) SetObservedLicenses(v []ApiLicenseDTO) {
	o.ObservedLicenses = v
}

// GetOverriddenLicenses returns the OverriddenLicenses field value if set, zero value otherwise.
func (o *ApiLicenseDataDTOV2) GetOverriddenLicenses() []ApiLicenseDTO {
	if o == nil || IsNil(o.OverriddenLicenses) {
		var ret []ApiLicenseDTO
		return ret
	}
	return o.OverriddenLicenses
}

// GetOverriddenLicensesOk returns a tuple with the OverriddenLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDataDTOV2) GetOverriddenLicensesOk() ([]ApiLicenseDTO, bool) {
	if o == nil || IsNil(o.OverriddenLicenses) {
		return nil, false
	}
	return o.OverriddenLicenses, true
}

// HasOverriddenLicenses returns a boolean if a field has been set.
func (o *ApiLicenseDataDTOV2) HasOverriddenLicenses() bool {
	if o != nil && !IsNil(o.OverriddenLicenses) {
		return true
	}

	return false
}

// SetOverriddenLicenses gets a reference to the given []ApiLicenseDTO and assigns it to the OverriddenLicenses field.
func (o *ApiLicenseDataDTOV2) SetOverriddenLicenses(v []ApiLicenseDTO) {
	o.OverriddenLicenses = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiLicenseDataDTOV2) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDataDTOV2) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiLicenseDataDTOV2) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiLicenseDataDTOV2) SetStatus(v string) {
	o.Status = &v
}

func (o ApiLicenseDataDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicenseDataDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeclaredLicenses) {
		toSerialize["declaredLicenses"] = o.DeclaredLicenses
	}
	if !IsNil(o.EffectiveLicenseThreats) {
		toSerialize["effectiveLicenseThreats"] = o.EffectiveLicenseThreats
	}
	if !IsNil(o.EffectiveLicenses) {
		toSerialize["effectiveLicenses"] = o.EffectiveLicenses
	}
	if !IsNil(o.ObservedLicenses) {
		toSerialize["observedLicenses"] = o.ObservedLicenses
	}
	if !IsNil(o.OverriddenLicenses) {
		toSerialize["overriddenLicenses"] = o.OverriddenLicenses
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableApiLicenseDataDTOV2 struct {
	value *ApiLicenseDataDTOV2
	isSet bool
}

func (v NullableApiLicenseDataDTOV2) Get() *ApiLicenseDataDTOV2 {
	return v.value
}

func (v *NullableApiLicenseDataDTOV2) Set(val *ApiLicenseDataDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicenseDataDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicenseDataDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicenseDataDTOV2(val *ApiLicenseDataDTOV2) *NullableApiLicenseDataDTOV2 {
	return &NullableApiLicenseDataDTOV2{value: val, isSet: true}
}

func (v NullableApiLicenseDataDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicenseDataDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


