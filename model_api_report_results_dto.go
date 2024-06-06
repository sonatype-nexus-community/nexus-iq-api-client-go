/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.177.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiReportResultsDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReportResultsDTO{}

// ApiReportResultsDTO struct for ApiReportResultsDTO
type ApiReportResultsDTO struct {
	ApplicationId *string `json:"applicationId,omitempty"`
	CommitHash *string `json:"commitHash,omitempty"`
	EmbeddableReportHtmlUrl *string `json:"embeddableReportHtmlUrl,omitempty"`
	EvaluationDate *time.Time `json:"evaluationDate,omitempty"`
	IsForMonitoring *bool `json:"isForMonitoring,omitempty"`
	IsReevaluation *bool `json:"isReevaluation,omitempty"`
	LatestReportHtmlUrl *string `json:"latestReportHtmlUrl,omitempty"`
	PolicyEvaluationId *string `json:"policyEvaluationId,omitempty"`
	PolicyEvaluationResult *PolicyEvaluationResult `json:"policyEvaluationResult,omitempty"`
	ReportDataUrl *string `json:"reportDataUrl,omitempty"`
	ReportHtmlUrl *string `json:"reportHtmlUrl,omitempty"`
	ReportPdfUrl *string `json:"reportPdfUrl,omitempty"`
	ScanId *string `json:"scanId,omitempty"`
	ScanTriggerType *string `json:"scanTriggerType,omitempty"`
	Stage *string `json:"stage,omitempty"`
}

// NewApiReportResultsDTO instantiates a new ApiReportResultsDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReportResultsDTO() *ApiReportResultsDTO {
	this := ApiReportResultsDTO{}
	return &this
}

// NewApiReportResultsDTOWithDefaults instantiates a new ApiReportResultsDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReportResultsDTOWithDefaults() *ApiReportResultsDTO {
	this := ApiReportResultsDTO{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *ApiReportResultsDTO) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetCommitHash returns the CommitHash field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetCommitHash() string {
	if o == nil || IsNil(o.CommitHash) {
		var ret string
		return ret
	}
	return *o.CommitHash
}

// GetCommitHashOk returns a tuple with the CommitHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetCommitHashOk() (*string, bool) {
	if o == nil || IsNil(o.CommitHash) {
		return nil, false
	}
	return o.CommitHash, true
}

// HasCommitHash returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasCommitHash() bool {
	if o != nil && !IsNil(o.CommitHash) {
		return true
	}

	return false
}

// SetCommitHash gets a reference to the given string and assigns it to the CommitHash field.
func (o *ApiReportResultsDTO) SetCommitHash(v string) {
	o.CommitHash = &v
}

// GetEmbeddableReportHtmlUrl returns the EmbeddableReportHtmlUrl field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetEmbeddableReportHtmlUrl() string {
	if o == nil || IsNil(o.EmbeddableReportHtmlUrl) {
		var ret string
		return ret
	}
	return *o.EmbeddableReportHtmlUrl
}

// GetEmbeddableReportHtmlUrlOk returns a tuple with the EmbeddableReportHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetEmbeddableReportHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EmbeddableReportHtmlUrl) {
		return nil, false
	}
	return o.EmbeddableReportHtmlUrl, true
}

// HasEmbeddableReportHtmlUrl returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasEmbeddableReportHtmlUrl() bool {
	if o != nil && !IsNil(o.EmbeddableReportHtmlUrl) {
		return true
	}

	return false
}

// SetEmbeddableReportHtmlUrl gets a reference to the given string and assigns it to the EmbeddableReportHtmlUrl field.
func (o *ApiReportResultsDTO) SetEmbeddableReportHtmlUrl(v string) {
	o.EmbeddableReportHtmlUrl = &v
}

// GetEvaluationDate returns the EvaluationDate field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetEvaluationDate() time.Time {
	if o == nil || IsNil(o.EvaluationDate) {
		var ret time.Time
		return ret
	}
	return *o.EvaluationDate
}

// GetEvaluationDateOk returns a tuple with the EvaluationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetEvaluationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EvaluationDate) {
		return nil, false
	}
	return o.EvaluationDate, true
}

// HasEvaluationDate returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasEvaluationDate() bool {
	if o != nil && !IsNil(o.EvaluationDate) {
		return true
	}

	return false
}

// SetEvaluationDate gets a reference to the given time.Time and assigns it to the EvaluationDate field.
func (o *ApiReportResultsDTO) SetEvaluationDate(v time.Time) {
	o.EvaluationDate = &v
}

// GetIsForMonitoring returns the IsForMonitoring field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetIsForMonitoring() bool {
	if o == nil || IsNil(o.IsForMonitoring) {
		var ret bool
		return ret
	}
	return *o.IsForMonitoring
}

// GetIsForMonitoringOk returns a tuple with the IsForMonitoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetIsForMonitoringOk() (*bool, bool) {
	if o == nil || IsNil(o.IsForMonitoring) {
		return nil, false
	}
	return o.IsForMonitoring, true
}

// HasIsForMonitoring returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasIsForMonitoring() bool {
	if o != nil && !IsNil(o.IsForMonitoring) {
		return true
	}

	return false
}

// SetIsForMonitoring gets a reference to the given bool and assigns it to the IsForMonitoring field.
func (o *ApiReportResultsDTO) SetIsForMonitoring(v bool) {
	o.IsForMonitoring = &v
}

// GetIsReevaluation returns the IsReevaluation field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetIsReevaluation() bool {
	if o == nil || IsNil(o.IsReevaluation) {
		var ret bool
		return ret
	}
	return *o.IsReevaluation
}

// GetIsReevaluationOk returns a tuple with the IsReevaluation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetIsReevaluationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsReevaluation) {
		return nil, false
	}
	return o.IsReevaluation, true
}

// HasIsReevaluation returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasIsReevaluation() bool {
	if o != nil && !IsNil(o.IsReevaluation) {
		return true
	}

	return false
}

// SetIsReevaluation gets a reference to the given bool and assigns it to the IsReevaluation field.
func (o *ApiReportResultsDTO) SetIsReevaluation(v bool) {
	o.IsReevaluation = &v
}

// GetLatestReportHtmlUrl returns the LatestReportHtmlUrl field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetLatestReportHtmlUrl() string {
	if o == nil || IsNil(o.LatestReportHtmlUrl) {
		var ret string
		return ret
	}
	return *o.LatestReportHtmlUrl
}

// GetLatestReportHtmlUrlOk returns a tuple with the LatestReportHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetLatestReportHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.LatestReportHtmlUrl) {
		return nil, false
	}
	return o.LatestReportHtmlUrl, true
}

// HasLatestReportHtmlUrl returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasLatestReportHtmlUrl() bool {
	if o != nil && !IsNil(o.LatestReportHtmlUrl) {
		return true
	}

	return false
}

// SetLatestReportHtmlUrl gets a reference to the given string and assigns it to the LatestReportHtmlUrl field.
func (o *ApiReportResultsDTO) SetLatestReportHtmlUrl(v string) {
	o.LatestReportHtmlUrl = &v
}

// GetPolicyEvaluationId returns the PolicyEvaluationId field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetPolicyEvaluationId() string {
	if o == nil || IsNil(o.PolicyEvaluationId) {
		var ret string
		return ret
	}
	return *o.PolicyEvaluationId
}

// GetPolicyEvaluationIdOk returns a tuple with the PolicyEvaluationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetPolicyEvaluationIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyEvaluationId) {
		return nil, false
	}
	return o.PolicyEvaluationId, true
}

// HasPolicyEvaluationId returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasPolicyEvaluationId() bool {
	if o != nil && !IsNil(o.PolicyEvaluationId) {
		return true
	}

	return false
}

// SetPolicyEvaluationId gets a reference to the given string and assigns it to the PolicyEvaluationId field.
func (o *ApiReportResultsDTO) SetPolicyEvaluationId(v string) {
	o.PolicyEvaluationId = &v
}

// GetPolicyEvaluationResult returns the PolicyEvaluationResult field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetPolicyEvaluationResult() PolicyEvaluationResult {
	if o == nil || IsNil(o.PolicyEvaluationResult) {
		var ret PolicyEvaluationResult
		return ret
	}
	return *o.PolicyEvaluationResult
}

// GetPolicyEvaluationResultOk returns a tuple with the PolicyEvaluationResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetPolicyEvaluationResultOk() (*PolicyEvaluationResult, bool) {
	if o == nil || IsNil(o.PolicyEvaluationResult) {
		return nil, false
	}
	return o.PolicyEvaluationResult, true
}

// HasPolicyEvaluationResult returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasPolicyEvaluationResult() bool {
	if o != nil && !IsNil(o.PolicyEvaluationResult) {
		return true
	}

	return false
}

// SetPolicyEvaluationResult gets a reference to the given PolicyEvaluationResult and assigns it to the PolicyEvaluationResult field.
func (o *ApiReportResultsDTO) SetPolicyEvaluationResult(v PolicyEvaluationResult) {
	o.PolicyEvaluationResult = &v
}

// GetReportDataUrl returns the ReportDataUrl field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetReportDataUrl() string {
	if o == nil || IsNil(o.ReportDataUrl) {
		var ret string
		return ret
	}
	return *o.ReportDataUrl
}

// GetReportDataUrlOk returns a tuple with the ReportDataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetReportDataUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportDataUrl) {
		return nil, false
	}
	return o.ReportDataUrl, true
}

// HasReportDataUrl returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasReportDataUrl() bool {
	if o != nil && !IsNil(o.ReportDataUrl) {
		return true
	}

	return false
}

// SetReportDataUrl gets a reference to the given string and assigns it to the ReportDataUrl field.
func (o *ApiReportResultsDTO) SetReportDataUrl(v string) {
	o.ReportDataUrl = &v
}

// GetReportHtmlUrl returns the ReportHtmlUrl field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetReportHtmlUrl() string {
	if o == nil || IsNil(o.ReportHtmlUrl) {
		var ret string
		return ret
	}
	return *o.ReportHtmlUrl
}

// GetReportHtmlUrlOk returns a tuple with the ReportHtmlUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetReportHtmlUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportHtmlUrl) {
		return nil, false
	}
	return o.ReportHtmlUrl, true
}

// HasReportHtmlUrl returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasReportHtmlUrl() bool {
	if o != nil && !IsNil(o.ReportHtmlUrl) {
		return true
	}

	return false
}

// SetReportHtmlUrl gets a reference to the given string and assigns it to the ReportHtmlUrl field.
func (o *ApiReportResultsDTO) SetReportHtmlUrl(v string) {
	o.ReportHtmlUrl = &v
}

// GetReportPdfUrl returns the ReportPdfUrl field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetReportPdfUrl() string {
	if o == nil || IsNil(o.ReportPdfUrl) {
		var ret string
		return ret
	}
	return *o.ReportPdfUrl
}

// GetReportPdfUrlOk returns a tuple with the ReportPdfUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetReportPdfUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ReportPdfUrl) {
		return nil, false
	}
	return o.ReportPdfUrl, true
}

// HasReportPdfUrl returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasReportPdfUrl() bool {
	if o != nil && !IsNil(o.ReportPdfUrl) {
		return true
	}

	return false
}

// SetReportPdfUrl gets a reference to the given string and assigns it to the ReportPdfUrl field.
func (o *ApiReportResultsDTO) SetReportPdfUrl(v string) {
	o.ReportPdfUrl = &v
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetScanId() string {
	if o == nil || IsNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasScanId() bool {
	if o != nil && !IsNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *ApiReportResultsDTO) SetScanId(v string) {
	o.ScanId = &v
}

// GetScanTriggerType returns the ScanTriggerType field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetScanTriggerType() string {
	if o == nil || IsNil(o.ScanTriggerType) {
		var ret string
		return ret
	}
	return *o.ScanTriggerType
}

// GetScanTriggerTypeOk returns a tuple with the ScanTriggerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetScanTriggerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ScanTriggerType) {
		return nil, false
	}
	return o.ScanTriggerType, true
}

// HasScanTriggerType returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasScanTriggerType() bool {
	if o != nil && !IsNil(o.ScanTriggerType) {
		return true
	}

	return false
}

// SetScanTriggerType gets a reference to the given string and assigns it to the ScanTriggerType field.
func (o *ApiReportResultsDTO) SetScanTriggerType(v string) {
	o.ScanTriggerType = &v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *ApiReportResultsDTO) GetStage() string {
	if o == nil || IsNil(o.Stage) {
		var ret string
		return ret
	}
	return *o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportResultsDTO) GetStageOk() (*string, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *ApiReportResultsDTO) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given string and assigns it to the Stage field.
func (o *ApiReportResultsDTO) SetStage(v string) {
	o.Stage = &v
}

func (o ApiReportResultsDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReportResultsDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.CommitHash) {
		toSerialize["commitHash"] = o.CommitHash
	}
	if !IsNil(o.EmbeddableReportHtmlUrl) {
		toSerialize["embeddableReportHtmlUrl"] = o.EmbeddableReportHtmlUrl
	}
	if !IsNil(o.EvaluationDate) {
		toSerialize["evaluationDate"] = o.EvaluationDate
	}
	if !IsNil(o.IsForMonitoring) {
		toSerialize["isForMonitoring"] = o.IsForMonitoring
	}
	if !IsNil(o.IsReevaluation) {
		toSerialize["isReevaluation"] = o.IsReevaluation
	}
	if !IsNil(o.LatestReportHtmlUrl) {
		toSerialize["latestReportHtmlUrl"] = o.LatestReportHtmlUrl
	}
	if !IsNil(o.PolicyEvaluationId) {
		toSerialize["policyEvaluationId"] = o.PolicyEvaluationId
	}
	if !IsNil(o.PolicyEvaluationResult) {
		toSerialize["policyEvaluationResult"] = o.PolicyEvaluationResult
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
	if !IsNil(o.ScanId) {
		toSerialize["scanId"] = o.ScanId
	}
	if !IsNil(o.ScanTriggerType) {
		toSerialize["scanTriggerType"] = o.ScanTriggerType
	}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	return toSerialize, nil
}

type NullableApiReportResultsDTO struct {
	value *ApiReportResultsDTO
	isSet bool
}

func (v NullableApiReportResultsDTO) Get() *ApiReportResultsDTO {
	return v.value
}

func (v *NullableApiReportResultsDTO) Set(val *ApiReportResultsDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReportResultsDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReportResultsDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReportResultsDTO(val *ApiReportResultsDTO) *NullableApiReportResultsDTO {
	return &NullableApiReportResultsDTO{value: val, isSet: true}
}

func (v NullableApiReportResultsDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReportResultsDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


