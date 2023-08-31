/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.166.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiLicenseDataDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicenseDataDTO{}

// ApiLicenseDataDTO struct for ApiLicenseDataDTO
type ApiLicenseDataDTO struct {
	DeclaredLicenses []ApiLicenseDTO `json:"declaredLicenses,omitempty"`
	EffectiveLicenses []ApiLicenseDTO `json:"effectiveLicenses,omitempty"`
	ObservedLicenses []ApiLicenseDTO `json:"observedLicenses,omitempty"`
	OverriddenLicenses []ApiLicenseDTO `json:"overriddenLicenses,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewApiLicenseDataDTO instantiates a new ApiLicenseDataDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicenseDataDTO() *ApiLicenseDataDTO {
	this := ApiLicenseDataDTO{}
	return &this
}

// NewApiLicenseDataDTOWithDefaults instantiates a new ApiLicenseDataDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicenseDataDTOWithDefaults() *ApiLicenseDataDTO {
	this := ApiLicenseDataDTO{}
	return &this
}

// GetDeclaredLicenses returns the DeclaredLicenses field value if set, zero value otherwise.
func (o *ApiLicenseDataDTO) GetDeclaredLicenses() []ApiLicenseDTO {
	if o == nil || IsNil(o.DeclaredLicenses) {
		var ret []ApiLicenseDTO
		return ret
	}
	return o.DeclaredLicenses
}

// GetDeclaredLicensesOk returns a tuple with the DeclaredLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDataDTO) GetDeclaredLicensesOk() ([]ApiLicenseDTO, bool) {
	if o == nil || IsNil(o.DeclaredLicenses) {
		return nil, false
	}
	return o.DeclaredLicenses, true
}

// HasDeclaredLicenses returns a boolean if a field has been set.
func (o *ApiLicenseDataDTO) HasDeclaredLicenses() bool {
	if o != nil && !IsNil(o.DeclaredLicenses) {
		return true
	}

	return false
}

// SetDeclaredLicenses gets a reference to the given []ApiLicenseDTO and assigns it to the DeclaredLicenses field.
func (o *ApiLicenseDataDTO) SetDeclaredLicenses(v []ApiLicenseDTO) {
	o.DeclaredLicenses = v
}

// GetEffectiveLicenses returns the EffectiveLicenses field value if set, zero value otherwise.
func (o *ApiLicenseDataDTO) GetEffectiveLicenses() []ApiLicenseDTO {
	if o == nil || IsNil(o.EffectiveLicenses) {
		var ret []ApiLicenseDTO
		return ret
	}
	return o.EffectiveLicenses
}

// GetEffectiveLicensesOk returns a tuple with the EffectiveLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDataDTO) GetEffectiveLicensesOk() ([]ApiLicenseDTO, bool) {
	if o == nil || IsNil(o.EffectiveLicenses) {
		return nil, false
	}
	return o.EffectiveLicenses, true
}

// HasEffectiveLicenses returns a boolean if a field has been set.
func (o *ApiLicenseDataDTO) HasEffectiveLicenses() bool {
	if o != nil && !IsNil(o.EffectiveLicenses) {
		return true
	}

	return false
}

// SetEffectiveLicenses gets a reference to the given []ApiLicenseDTO and assigns it to the EffectiveLicenses field.
func (o *ApiLicenseDataDTO) SetEffectiveLicenses(v []ApiLicenseDTO) {
	o.EffectiveLicenses = v
}

// GetObservedLicenses returns the ObservedLicenses field value if set, zero value otherwise.
func (o *ApiLicenseDataDTO) GetObservedLicenses() []ApiLicenseDTO {
	if o == nil || IsNil(o.ObservedLicenses) {
		var ret []ApiLicenseDTO
		return ret
	}
	return o.ObservedLicenses
}

// GetObservedLicensesOk returns a tuple with the ObservedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDataDTO) GetObservedLicensesOk() ([]ApiLicenseDTO, bool) {
	if o == nil || IsNil(o.ObservedLicenses) {
		return nil, false
	}
	return o.ObservedLicenses, true
}

// HasObservedLicenses returns a boolean if a field has been set.
func (o *ApiLicenseDataDTO) HasObservedLicenses() bool {
	if o != nil && !IsNil(o.ObservedLicenses) {
		return true
	}

	return false
}

// SetObservedLicenses gets a reference to the given []ApiLicenseDTO and assigns it to the ObservedLicenses field.
func (o *ApiLicenseDataDTO) SetObservedLicenses(v []ApiLicenseDTO) {
	o.ObservedLicenses = v
}

// GetOverriddenLicenses returns the OverriddenLicenses field value if set, zero value otherwise.
func (o *ApiLicenseDataDTO) GetOverriddenLicenses() []ApiLicenseDTO {
	if o == nil || IsNil(o.OverriddenLicenses) {
		var ret []ApiLicenseDTO
		return ret
	}
	return o.OverriddenLicenses
}

// GetOverriddenLicensesOk returns a tuple with the OverriddenLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDataDTO) GetOverriddenLicensesOk() ([]ApiLicenseDTO, bool) {
	if o == nil || IsNil(o.OverriddenLicenses) {
		return nil, false
	}
	return o.OverriddenLicenses, true
}

// HasOverriddenLicenses returns a boolean if a field has been set.
func (o *ApiLicenseDataDTO) HasOverriddenLicenses() bool {
	if o != nil && !IsNil(o.OverriddenLicenses) {
		return true
	}

	return false
}

// SetOverriddenLicenses gets a reference to the given []ApiLicenseDTO and assigns it to the OverriddenLicenses field.
func (o *ApiLicenseDataDTO) SetOverriddenLicenses(v []ApiLicenseDTO) {
	o.OverriddenLicenses = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiLicenseDataDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseDataDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiLicenseDataDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiLicenseDataDTO) SetStatus(v string) {
	o.Status = &v
}

func (o ApiLicenseDataDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicenseDataDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeclaredLicenses) {
		toSerialize["declaredLicenses"] = o.DeclaredLicenses
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

type NullableApiLicenseDataDTO struct {
	value *ApiLicenseDataDTO
	isSet bool
}

func (v NullableApiLicenseDataDTO) Get() *ApiLicenseDataDTO {
	return v.value
}

func (v *NullableApiLicenseDataDTO) Set(val *ApiLicenseDataDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicenseDataDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicenseDataDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicenseDataDTO(val *ApiLicenseDataDTO) *NullableApiLicenseDataDTO {
	return &NullableApiLicenseDataDTO{value: val, isSet: true}
}

func (v NullableApiLicenseDataDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicenseDataDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


