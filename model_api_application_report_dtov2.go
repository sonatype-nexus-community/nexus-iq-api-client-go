/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.193.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiApplicationReportDTOV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiApplicationReportDTOV2{}

// ApiApplicationReportDTOV2 struct for ApiApplicationReportDTOV2
type ApiApplicationReportDTOV2 struct {
	ApplicationId *string `json:"applicationId,omitempty"`
	EmbeddableReportHtmlUrl *string `json:"embeddableReportHtmlUrl,omitempty"`
	EvaluationDate *time.Time `json:"evaluationDate,omitempty"`
	LatestReportHtmlUrl *string `json:"latestReportHtmlUrl,omitempty"`
	ReportDataUrl *string `json:"reportDataUrl,omitempty"`
	ReportHtmlUrl *string `json:"reportHtmlUrl,omitempty"`
	ReportPdfUrl *string `json:"reportPdfUrl,omitempty"`
	Stage *string `json:"stage,omitempty"`
}

// NewApiApplicationReportDTOV2 instantiates a new ApiApplicationReportDTOV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiApplicationReportDTOV2() *ApiApplicationReportDTOV2 {
	this := ApiApplicationReportDTOV2{}
	return &this
}

// NewApiApplicationReportDTOV2WithDefaults instantiates a new ApiApplicationReportDTOV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiApplicationReportDTOV2WithDefaults() *ApiApplicationReportDTOV2 {
	this := ApiApplicationReportDTOV2{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ApiApplicationReportDTOV2) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationReportDTOV2) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ApiApplicationReportDTOV2) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ApiApplicationReportDTOV2) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetEmbeddableReportHtmlUrl returns the EmbeddableReportHtmlUrl field value if set, zero value otherwise.
func (o *ApiApplicationReportDTOV2) GetEmbeddableReportHtmlUrl() string {
	if o == nil || IsNil(o.EmbeddableReportHtmlUrl) {
		var ret string
		return ret
	}
	return *o.EmbeddableReportHtmlUrl
}

// GetEmbeddableReportHtmlUrlOk returns a tuple with the EmbeddableReportHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationReportDTOV2) GetEmbeddableReportHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EmbeddableReportHtmlUrl) {
		return nil, false
	}
	return o.EmbeddableReportHtmlUrl, true
}

// HasEmbeddableReportHtmlUrl returns a boolean if a field has been set.
func (o *ApiApplicationReportDTOV2) HasEmbeddableReportHtmlUrl() bool {
	if o != nil && !IsNil(o.EmbeddableReportHtmlUrl) {
		return true
	}

	return false
}

// SetEmbeddableReportHtmlUrl gets a reference to the given string and assigns it to the EmbeddableReportHtmlUrl field.
func (o *ApiApplicationReportDTOV2) SetEmbeddableReportHtmlUrl(v string) {
	o.EmbeddableReportHtmlUrl = &v
}

// GetEvaluationDate returns the EvaluationDate field value if set, zero value otherwise.
func (o *ApiApplicationReportDTOV2) GetEvaluationDate() time.Time {
	if o == nil || IsNil(o.EvaluationDate) {
		var ret time.Time
		return ret
	}
	return *o.EvaluationDate
}

// GetEvaluationDateOk returns a tuple with the EvaluationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationReportDTOV2) GetEvaluationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EvaluationDate) {
		return nil, false
	}
	return o.EvaluationDate, true
}

// HasEvaluationDate returns a boolean if a field has been set.
func (o *ApiApplicationReportDTOV2) HasEvaluationDate() bool {
	if o != nil && !IsNil(o.EvaluationDate) {
		return true
	}

	return false
}

// SetEvaluationDate gets a reference to the given time.Time and assigns it to the EvaluationDate field.
func (o *ApiApplicationReportDTOV2) SetEvaluationDate(v time.Time) {
	o.EvaluationDate = &v
}

// GetLatestReportHtmlUrl returns the LatestReportHtmlUrl field value if set, zero value otherwise.
func (o *ApiApplicationReportDTOV2) GetLatestReportHtmlUrl() string {
	if o == nil || IsNil(o.LatestReportHtmlUrl) {
		var ret string
		return ret
	}
	return *o.LatestReportHtmlUrl
}

// GetLatestReportHtmlUrlOk returns a tuple with the LatestReportHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationReportDTOV2) GetLatestReportHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LatestReportHtmlUrl) {
		return nil, false
	}
	return o.LatestReportHtmlUrl, true
}

// HasLatestReportHtmlUrl returns a boolean if a field has been set.
func (o *ApiApplicationReportDTOV2) HasLatestReportHtmlUrl() bool {
	if o != nil && !IsNil(o.LatestReportHtmlUrl) {
		return true
	}

	return false
}

// SetLatestReportHtmlUrl gets a reference to the given string and assigns it to the LatestReportHtmlUrl field.
func (o *ApiApplicationReportDTOV2) SetLatestReportHtmlUrl(v string) {
	o.LatestReportHtmlUrl = &v
}

// GetReportDataUrl returns the ReportDataUrl field value if set, zero value otherwise.
func (o *ApiApplicationReportDTOV2) GetReportDataUrl() string {
	if o == nil || IsNil(o.ReportDataUrl) {
		var ret string
		return ret
	}
	return *o.ReportDataUrl
}

// GetReportDataUrlOk returns a tuple with the ReportDataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationReportDTOV2) GetReportDataUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportDataUrl) {
		return nil, false
	}
	return o.ReportDataUrl, true
}

// HasReportDataUrl returns a boolean if a field has been set.
func (o *ApiApplicationReportDTOV2) HasReportDataUrl() bool {
	if o != nil && !IsNil(o.ReportDataUrl) {
		return true
	}

	return false
}

// SetReportDataUrl gets a reference to the given string and assigns it to the ReportDataUrl field.
func (o *ApiApplicationReportDTOV2) SetReportDataUrl(v string) {
	o.ReportDataUrl = &v
}

// GetReportHtmlUrl returns the ReportHtmlUrl field value if set, zero value otherwise.
func (o *ApiApplicationReportDTOV2) GetReportHtmlUrl() string {
	if o == nil || IsNil(o.ReportHtmlUrl) {
		var ret string
		return ret
	}
	return *o.ReportHtmlUrl
}

// GetReportHtmlUrlOk returns a tuple with the ReportHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationReportDTOV2) GetReportHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportHtmlUrl) {
		return nil, false
	}
	return o.ReportHtmlUrl, true
}

// HasReportHtmlUrl returns a boolean if a field has been set.
func (o *ApiApplicationReportDTOV2) HasReportHtmlUrl() bool {
	if o != nil && !IsNil(o.ReportHtmlUrl) {
		return true
	}

	return false
}

// SetReportHtmlUrl gets a reference to the given string and assigns it to the ReportHtmlUrl field.
func (o *ApiApplicationReportDTOV2) SetReportHtmlUrl(v string) {
	o.ReportHtmlUrl = &v
}

// GetReportPdfUrl returns the ReportPdfUrl field value if set, zero value otherwise.
func (o *ApiApplicationReportDTOV2) GetReportPdfUrl() string {
	if o == nil || IsNil(o.ReportPdfUrl) {
		var ret string
		return ret
	}
	return *o.ReportPdfUrl
}

// GetReportPdfUrlOk returns a tuple with the ReportPdfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationReportDTOV2) GetReportPdfUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportPdfUrl) {
		return nil, false
	}
	return o.ReportPdfUrl, true
}

// HasReportPdfUrl returns a boolean if a field has been set.
func (o *ApiApplicationReportDTOV2) HasReportPdfUrl() bool {
	if o != nil && !IsNil(o.ReportPdfUrl) {
		return true
	}

	return false
}

// SetReportPdfUrl gets a reference to the given string and assigns it to the ReportPdfUrl field.
func (o *ApiApplicationReportDTOV2) SetReportPdfUrl(v string) {
	o.ReportPdfUrl = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *ApiApplicationReportDTOV2) GetStage() string {
	if o == nil || IsNil(o.Stage) {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiApplicationReportDTOV2) GetStageOk() (*string, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *ApiApplicationReportDTOV2) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *ApiApplicationReportDTOV2) SetStage(v string) {
	o.Stage = &v
}

func (o ApiApplicationReportDTOV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiApplicationReportDTOV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.EmbeddableReportHtmlUrl) {
		toSerialize["embeddableReportHtmlUrl"] = o.EmbeddableReportHtmlUrl
	}
	if !IsNil(o.EvaluationDate) {
		toSerialize["evaluationDate"] = o.EvaluationDate
	}
	if !IsNil(o.LatestReportHtmlUrl) {
		toSerialize["latestReportHtmlUrl"] = o.LatestReportHtmlUrl
	}
	if !IsNil(o.ReportDataUrl) {
		toSerialize["reportDataUrl"] = o.ReportDataUrl
	}
	if !IsNil(o.ReportHtmlUrl) {
		toSerialize["reportHtmlUrl"] = o.ReportHtmlUrl
	}
	if !IsNil(o.ReportPdfUrl) {
		toSerialize["reportPdfUrl"] = o.ReportPdfUrl
	}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	return toSerialize, nil
}

type NullableApiApplicationReportDTOV2 struct {
	value *ApiApplicationReportDTOV2
	isSet bool
}

func (v NullableApiApplicationReportDTOV2) Get() *ApiApplicationReportDTOV2 {
	return v.value
}

func (v *NullableApiApplicationReportDTOV2) Set(val *ApiApplicationReportDTOV2) {
	v.value = val
	v.isSet = true
}

func (v NullableApiApplicationReportDTOV2) IsSet() bool {
	return v.isSet
}

func (v *NullableApiApplicationReportDTOV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiApplicationReportDTOV2(val *ApiApplicationReportDTOV2) *NullableApiApplicationReportDTOV2 {
	return &NullableApiApplicationReportDTOV2{value: val, isSet: true}
}

func (v NullableApiApplicationReportDTOV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiApplicationReportDTOV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


