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

// checks if the ApiPullRequestResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPullRequestResults{}

// ApiPullRequestResults struct for ApiPullRequestResults
type ApiPullRequestResults struct {
	Results []ApiPullRequestResult `json:"results,omitempty"`
}

// NewApiPullRequestResults instantiates a new ApiPullRequestResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPullRequestResults() *ApiPullRequestResults {
	this := ApiPullRequestResults{}
	return &this
}

// NewApiPullRequestResultsWithDefaults instantiates a new ApiPullRequestResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPullRequestResultsWithDefaults() *ApiPullRequestResults {
	this := ApiPullRequestResults{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ApiPullRequestResults) GetResults() []ApiPullRequestResult {
	if o == nil || IsNil(o.Results) {
		var ret []ApiPullRequestResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPullRequestResults) GetResultsOk() ([]ApiPullRequestResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ApiPullRequestResults) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ApiPullRequestResult and assigns it to the Results field.
func (o *ApiPullRequestResults) SetResults(v []ApiPullRequestResult) {
	o.Results = v
}

func (o ApiPullRequestResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPullRequestResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableApiPullRequestResults struct {
	value *ApiPullRequestResults
	isSet bool
}

func (v NullableApiPullRequestResults) Get() *ApiPullRequestResults {
	return v.value
}

func (v *NullableApiPullRequestResults) Set(val *ApiPullRequestResults) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPullRequestResults) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPullRequestResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPullRequestResults(val *ApiPullRequestResults) *NullableApiPullRequestResults {
	return &NullableApiPullRequestResults{value: val, isSet: true}
}

func (v NullableApiPullRequestResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPullRequestResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


