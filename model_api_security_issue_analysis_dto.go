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

// checks if the ApiSecurityIssueAnalysisDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSecurityIssueAnalysisDTO{}

// ApiSecurityIssueAnalysisDTO struct for ApiSecurityIssueAnalysisDTO
type ApiSecurityIssueAnalysisDTO struct {
	Detail *string `json:"detail,omitempty"`
	Justification *string `json:"justification,omitempty"`
	Response *string `json:"response,omitempty"`
	State *string `json:"state,omitempty"`
}

// NewApiSecurityIssueAnalysisDTO instantiates a new ApiSecurityIssueAnalysisDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSecurityIssueAnalysisDTO() *ApiSecurityIssueAnalysisDTO {
	this := ApiSecurityIssueAnalysisDTO{}
	return &this
}

// NewApiSecurityIssueAnalysisDTOWithDefaults instantiates a new ApiSecurityIssueAnalysisDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSecurityIssueAnalysisDTOWithDefaults() *ApiSecurityIssueAnalysisDTO {
	this := ApiSecurityIssueAnalysisDTO{}
	return &this
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *ApiSecurityIssueAnalysisDTO) GetDetail() string {
	if o == nil || IsNil(o.Detail) {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueAnalysisDTO) GetDetailOk() (*string, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ApiSecurityIssueAnalysisDTO) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *ApiSecurityIssueAnalysisDTO) SetDetail(v string) {
	o.Detail = &v
}

// GetJustification returns the Justification field value if set, zero value otherwise.
func (o *ApiSecurityIssueAnalysisDTO) GetJustification() string {
	if o == nil || IsNil(o.Justification) {
		var ret string
		return ret
	}
	return *o.Justification
}

// GetJustificationOk returns a tuple with the Justification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueAnalysisDTO) GetJustificationOk() (*string, bool) {
	if o == nil || IsNil(o.Justification) {
		return nil, false
	}
	return o.Justification, true
}

// HasJustification returns a boolean if a field has been set.
func (o *ApiSecurityIssueAnalysisDTO) HasJustification() bool {
	if o != nil && !IsNil(o.Justification) {
		return true
	}

	return false
}

// SetJustification gets a reference to the given string and assigns it to the Justification field.
func (o *ApiSecurityIssueAnalysisDTO) SetJustification(v string) {
	o.Justification = &v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *ApiSecurityIssueAnalysisDTO) GetResponse() string {
	if o == nil || IsNil(o.Response) {
		var ret string
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueAnalysisDTO) GetResponseOk() (*string, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *ApiSecurityIssueAnalysisDTO) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given string and assigns it to the Response field.
func (o *ApiSecurityIssueAnalysisDTO) SetResponse(v string) {
	o.Response = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ApiSecurityIssueAnalysisDTO) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueAnalysisDTO) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ApiSecurityIssueAnalysisDTO) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ApiSecurityIssueAnalysisDTO) SetState(v string) {
	o.State = &v
}

func (o ApiSecurityIssueAnalysisDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSecurityIssueAnalysisDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Detail) {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.Justification) {
		toSerialize["justification"] = o.Justification
	}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableApiSecurityIssueAnalysisDTO struct {
	value *ApiSecurityIssueAnalysisDTO
	isSet bool
}

func (v NullableApiSecurityIssueAnalysisDTO) Get() *ApiSecurityIssueAnalysisDTO {
	return v.value
}

func (v *NullableApiSecurityIssueAnalysisDTO) Set(val *ApiSecurityIssueAnalysisDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSecurityIssueAnalysisDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSecurityIssueAnalysisDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSecurityIssueAnalysisDTO(val *ApiSecurityIssueAnalysisDTO) *NullableApiSecurityIssueAnalysisDTO {
	return &NullableApiSecurityIssueAnalysisDTO{value: val, isSet: true}
}

func (v NullableApiSecurityIssueAnalysisDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSecurityIssueAnalysisDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


