/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.184.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiPullRequestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPullRequestResult{}

// ApiPullRequestResult struct for ApiPullRequestResult
type ApiPullRequestResult struct {
	ExceptionThrown *bool `json:"exceptionThrown,omitempty"`
	Reasoning *string `json:"reasoning,omitempty"`
	StartTime *time.Time `json:"startTime,omitempty"`
	Successful *bool `json:"successful,omitempty"`
	Title *string `json:"title,omitempty"`
	TotalTime *int64 `json:"totalTime,omitempty"`
}

// NewApiPullRequestResult instantiates a new ApiPullRequestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPullRequestResult() *ApiPullRequestResult {
	this := ApiPullRequestResult{}
	return &this
}

// NewApiPullRequestResultWithDefaults instantiates a new ApiPullRequestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPullRequestResultWithDefaults() *ApiPullRequestResult {
	this := ApiPullRequestResult{}
	return &this
}

// GetExceptionThrown returns the ExceptionThrown field value if set, zero value otherwise.
func (o *ApiPullRequestResult) GetExceptionThrown() bool {
	if o == nil || IsNil(o.ExceptionThrown) {
		var ret bool
		return ret
	}
	return *o.ExceptionThrown
}

// GetExceptionThrownOk returns a tuple with the ExceptionThrown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPullRequestResult) GetExceptionThrownOk() (*bool, bool) {
	if o == nil || IsNil(o.ExceptionThrown) {
		return nil, false
	}
	return o.ExceptionThrown, true
}

// HasExceptionThrown returns a boolean if a field has been set.
func (o *ApiPullRequestResult) HasExceptionThrown() bool {
	if o != nil && !IsNil(o.ExceptionThrown) {
		return true
	}

	return false
}

// SetExceptionThrown gets a reference to the given bool and assigns it to the ExceptionThrown field.
func (o *ApiPullRequestResult) SetExceptionThrown(v bool) {
	o.ExceptionThrown = &v
}

// GetReasoning returns the Reasoning field value if set, zero value otherwise.
func (o *ApiPullRequestResult) GetReasoning() string {
	if o == nil || IsNil(o.Reasoning) {
		var ret string
		return ret
	}
	return *o.Reasoning
}

// GetReasoningOk returns a tuple with the Reasoning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPullRequestResult) GetReasoningOk() (*string, bool) {
	if o == nil || IsNil(o.Reasoning) {
		return nil, false
	}
	return o.Reasoning, true
}

// HasReasoning returns a boolean if a field has been set.
func (o *ApiPullRequestResult) HasReasoning() bool {
	if o != nil && !IsNil(o.Reasoning) {
		return true
	}

	return false
}

// SetReasoning gets a reference to the given string and assigns it to the Reasoning field.
func (o *ApiPullRequestResult) SetReasoning(v string) {
	o.Reasoning = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApiPullRequestResult) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPullRequestResult) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApiPullRequestResult) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApiPullRequestResult) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetSuccessful returns the Successful field value if set, zero value otherwise.
func (o *ApiPullRequestResult) GetSuccessful() bool {
	if o == nil || IsNil(o.Successful) {
		var ret bool
		return ret
	}
	return *o.Successful
}

// GetSuccessfulOk returns a tuple with the Successful field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPullRequestResult) GetSuccessfulOk() (*bool, bool) {
	if o == nil || IsNil(o.Successful) {
		return nil, false
	}
	return o.Successful, true
}

// HasSuccessful returns a boolean if a field has been set.
func (o *ApiPullRequestResult) HasSuccessful() bool {
	if o != nil && !IsNil(o.Successful) {
		return true
	}

	return false
}

// SetSuccessful gets a reference to the given bool and assigns it to the Successful field.
func (o *ApiPullRequestResult) SetSuccessful(v bool) {
	o.Successful = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiPullRequestResult) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPullRequestResult) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiPullRequestResult) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiPullRequestResult) SetTitle(v string) {
	o.Title = &v
}

// GetTotalTime returns the TotalTime field value if set, zero value otherwise.
func (o *ApiPullRequestResult) GetTotalTime() int64 {
	if o == nil || IsNil(o.TotalTime) {
		var ret int64
		return ret
	}
	return *o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPullRequestResult) GetTotalTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalTime) {
		return nil, false
	}
	return o.TotalTime, true
}

// HasTotalTime returns a boolean if a field has been set.
func (o *ApiPullRequestResult) HasTotalTime() bool {
	if o != nil && !IsNil(o.TotalTime) {
		return true
	}

	return false
}

// SetTotalTime gets a reference to the given int64 and assigns it to the TotalTime field.
func (o *ApiPullRequestResult) SetTotalTime(v int64) {
	o.TotalTime = &v
}

func (o ApiPullRequestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPullRequestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExceptionThrown) {
		toSerialize["exceptionThrown"] = o.ExceptionThrown
	}
	if !IsNil(o.Reasoning) {
		toSerialize["reasoning"] = o.Reasoning
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Successful) {
		toSerialize["successful"] = o.Successful
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.TotalTime) {
		toSerialize["totalTime"] = o.TotalTime
	}
	return toSerialize, nil
}

type NullableApiPullRequestResult struct {
	value *ApiPullRequestResult
	isSet bool
}

func (v NullableApiPullRequestResult) Get() *ApiPullRequestResult {
	return v.value
}

func (v *NullableApiPullRequestResult) Set(val *ApiPullRequestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPullRequestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPullRequestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPullRequestResult(val *ApiPullRequestResult) *NullableApiPullRequestResult {
	return &NullableApiPullRequestResult{value: val, isSet: true}
}

func (v NullableApiPullRequestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPullRequestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


