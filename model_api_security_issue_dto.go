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

// checks if the ApiSecurityIssueDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSecurityIssueDTO{}

// ApiSecurityIssueDTO struct for ApiSecurityIssueDTO
type ApiSecurityIssueDTO struct {
	Analysis *ApiSecurityIssueAnalysisDTO `json:"analysis,omitempty"`
	CvssVector *string `json:"cvssVector,omitempty"`
	CvssVectorSource *string `json:"cvssVectorSource,omitempty"`
	Cwe *string `json:"cwe,omitempty"`
	Reference *string `json:"reference,omitempty"`
	Severity *float32 `json:"severity,omitempty"`
	Source *string `json:"source,omitempty"`
	Status *string `json:"status,omitempty"`
	ThreatCategory *string `json:"threatCategory,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewApiSecurityIssueDTO instantiates a new ApiSecurityIssueDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSecurityIssueDTO() *ApiSecurityIssueDTO {
	this := ApiSecurityIssueDTO{}
	return &this
}

// NewApiSecurityIssueDTOWithDefaults instantiates a new ApiSecurityIssueDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSecurityIssueDTOWithDefaults() *ApiSecurityIssueDTO {
	this := ApiSecurityIssueDTO{}
	return &this
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *ApiSecurityIssueDTO) GetAnalysis() ApiSecurityIssueAnalysisDTO {
	if o == nil || IsNil(o.Analysis) {
		var ret ApiSecurityIssueAnalysisDTO
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueDTO) GetAnalysisOk() (*ApiSecurityIssueAnalysisDTO, bool) {
	if o == nil || IsNil(o.Analysis) {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *ApiSecurityIssueDTO) HasAnalysis() bool {
	if o != nil && !IsNil(o.Analysis) {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given ApiSecurityIssueAnalysisDTO and assigns it to the Analysis field.
func (o *ApiSecurityIssueDTO) SetAnalysis(v ApiSecurityIssueAnalysisDTO) {
	o.Analysis = &v
}

// GetCvssVector returns the CvssVector field value if set, zero value otherwise.
func (o *ApiSecurityIssueDTO) GetCvssVector() string {
	if o == nil || IsNil(o.CvssVector) {
		var ret string
		return ret
	}
	return *o.CvssVector
}

// GetCvssVectorOk returns a tuple with the CvssVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueDTO) GetCvssVectorOk() (*string, bool) {
	if o == nil || IsNil(o.CvssVector) {
		return nil, false
	}
	return o.CvssVector, true
}

// HasCvssVector returns a boolean if a field has been set.
func (o *ApiSecurityIssueDTO) HasCvssVector() bool {
	if o != nil && !IsNil(o.CvssVector) {
		return true
	}

	return false
}

// SetCvssVector gets a reference to the given string and assigns it to the CvssVector field.
func (o *ApiSecurityIssueDTO) SetCvssVector(v string) {
	o.CvssVector = &v
}

// GetCvssVectorSource returns the CvssVectorSource field value if set, zero value otherwise.
func (o *ApiSecurityIssueDTO) GetCvssVectorSource() string {
	if o == nil || IsNil(o.CvssVectorSource) {
		var ret string
		return ret
	}
	return *o.CvssVectorSource
}

// GetCvssVectorSourceOk returns a tuple with the CvssVectorSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueDTO) GetCvssVectorSourceOk() (*string, bool) {
	if o == nil || IsNil(o.CvssVectorSource) {
		return nil, false
	}
	return o.CvssVectorSource, true
}

// HasCvssVectorSource returns a boolean if a field has been set.
func (o *ApiSecurityIssueDTO) HasCvssVectorSource() bool {
	if o != nil && !IsNil(o.CvssVectorSource) {
		return true
	}

	return false
}

// SetCvssVectorSource gets a reference to the given string and assigns it to the CvssVectorSource field.
func (o *ApiSecurityIssueDTO) SetCvssVectorSource(v string) {
	o.CvssVectorSource = &v
}

// GetCwe returns the Cwe field value if set, zero value otherwise.
func (o *ApiSecurityIssueDTO) GetCwe() string {
	if o == nil || IsNil(o.Cwe) {
		var ret string
		return ret
	}
	return *o.Cwe
}

// GetCweOk returns a tuple with the Cwe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueDTO) GetCweOk() (*string, bool) {
	if o == nil || IsNil(o.Cwe) {
		return nil, false
	}
	return o.Cwe, true
}

// HasCwe returns a boolean if a field has been set.
func (o *ApiSecurityIssueDTO) HasCwe() bool {
	if o != nil && !IsNil(o.Cwe) {
		return true
	}

	return false
}

// SetCwe gets a reference to the given string and assigns it to the Cwe field.
func (o *ApiSecurityIssueDTO) SetCwe(v string) {
	o.Cwe = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *ApiSecurityIssueDTO) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueDTO) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *ApiSecurityIssueDTO) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *ApiSecurityIssueDTO) SetReference(v string) {
	o.Reference = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *ApiSecurityIssueDTO) GetSeverity() float32 {
	if o == nil || IsNil(o.Severity) {
		var ret float32
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueDTO) GetSeverityOk() (*float32, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *ApiSecurityIssueDTO) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given float32 and assigns it to the Severity field.
func (o *ApiSecurityIssueDTO) SetSeverity(v float32) {
	o.Severity = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ApiSecurityIssueDTO) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueDTO) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ApiSecurityIssueDTO) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ApiSecurityIssueDTO) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiSecurityIssueDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiSecurityIssueDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiSecurityIssueDTO) SetStatus(v string) {
	o.Status = &v
}

// GetThreatCategory returns the ThreatCategory field value if set, zero value otherwise.
func (o *ApiSecurityIssueDTO) GetThreatCategory() string {
	if o == nil || IsNil(o.ThreatCategory) {
		var ret string
		return ret
	}
	return *o.ThreatCategory
}

// GetThreatCategoryOk returns a tuple with the ThreatCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueDTO) GetThreatCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.ThreatCategory) {
		return nil, false
	}
	return o.ThreatCategory, true
}

// HasThreatCategory returns a boolean if a field has been set.
func (o *ApiSecurityIssueDTO) HasThreatCategory() bool {
	if o != nil && !IsNil(o.ThreatCategory) {
		return true
	}

	return false
}

// SetThreatCategory gets a reference to the given string and assigns it to the ThreatCategory field.
func (o *ApiSecurityIssueDTO) SetThreatCategory(v string) {
	o.ThreatCategory = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ApiSecurityIssueDTO) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSecurityIssueDTO) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ApiSecurityIssueDTO) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ApiSecurityIssueDTO) SetUrl(v string) {
	o.Url = &v
}

func (o ApiSecurityIssueDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSecurityIssueDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Analysis) {
		toSerialize["analysis"] = o.Analysis
	}
	if !IsNil(o.CvssVector) {
		toSerialize["cvssVector"] = o.CvssVector
	}
	if !IsNil(o.CvssVectorSource) {
		toSerialize["cvssVectorSource"] = o.CvssVectorSource
	}
	if !IsNil(o.Cwe) {
		toSerialize["cwe"] = o.Cwe
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ThreatCategory) {
		toSerialize["threatCategory"] = o.ThreatCategory
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableApiSecurityIssueDTO struct {
	value *ApiSecurityIssueDTO
	isSet bool
}

func (v NullableApiSecurityIssueDTO) Get() *ApiSecurityIssueDTO {
	return v.value
}

func (v *NullableApiSecurityIssueDTO) Set(val *ApiSecurityIssueDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSecurityIssueDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSecurityIssueDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSecurityIssueDTO(val *ApiSecurityIssueDTO) *NullableApiSecurityIssueDTO {
	return &NullableApiSecurityIssueDTO{value: val, isSet: true}
}

func (v NullableApiSecurityIssueDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSecurityIssueDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


