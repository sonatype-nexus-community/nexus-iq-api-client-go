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

// checks if the ApiRepositoryComponentEvaluationResultList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryComponentEvaluationResultList{}

// ApiRepositoryComponentEvaluationResultList struct for ApiRepositoryComponentEvaluationResultList
type ApiRepositoryComponentEvaluationResultList struct {
	RepositoryId *string `json:"repositoryId,omitempty"`
	RepositoryManagerId *string `json:"repositoryManagerId,omitempty"`
	RepositoryPublicId *string `json:"repositoryPublicId,omitempty"`
	RepositoryType *string `json:"repositoryType,omitempty"`
	Results []ApiRepositoryComponentEvaluationResult `json:"results,omitempty"`
}

// NewApiRepositoryComponentEvaluationResultList instantiates a new ApiRepositoryComponentEvaluationResultList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryComponentEvaluationResultList() *ApiRepositoryComponentEvaluationResultList {
	this := ApiRepositoryComponentEvaluationResultList{}
	return &this
}

// NewApiRepositoryComponentEvaluationResultListWithDefaults instantiates a new ApiRepositoryComponentEvaluationResultList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryComponentEvaluationResultListWithDefaults() *ApiRepositoryComponentEvaluationResultList {
	this := ApiRepositoryComponentEvaluationResultList{}
	return &this
}

// GetRepositoryId returns the RepositoryId field value if set, zero value otherwise.
func (o *ApiRepositoryComponentEvaluationResultList) GetRepositoryId() string {
	if o == nil || IsNil(o.RepositoryId) {
		var ret string
		return ret
	}
	return *o.RepositoryId
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentEvaluationResultList) GetRepositoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryId) {
		return nil, false
	}
	return o.RepositoryId, true
}

// HasRepositoryId returns a boolean if a field has been set.
func (o *ApiRepositoryComponentEvaluationResultList) HasRepositoryId() bool {
	if o != nil && !IsNil(o.RepositoryId) {
		return true
	}

	return false
}

// SetRepositoryId gets a reference to the given string and assigns it to the RepositoryId field.
func (o *ApiRepositoryComponentEvaluationResultList) SetRepositoryId(v string) {
	o.RepositoryId = &v
}

// GetRepositoryManagerId returns the RepositoryManagerId field value if set, zero value otherwise.
func (o *ApiRepositoryComponentEvaluationResultList) GetRepositoryManagerId() string {
	if o == nil || IsNil(o.RepositoryManagerId) {
		var ret string
		return ret
	}
	return *o.RepositoryManagerId
}

// GetRepositoryManagerIdOk returns a tuple with the RepositoryManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentEvaluationResultList) GetRepositoryManagerIdOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryManagerId) {
		return nil, false
	}
	return o.RepositoryManagerId, true
}

// HasRepositoryManagerId returns a boolean if a field has been set.
func (o *ApiRepositoryComponentEvaluationResultList) HasRepositoryManagerId() bool {
	if o != nil && !IsNil(o.RepositoryManagerId) {
		return true
	}

	return false
}

// SetRepositoryManagerId gets a reference to the given string and assigns it to the RepositoryManagerId field.
func (o *ApiRepositoryComponentEvaluationResultList) SetRepositoryManagerId(v string) {
	o.RepositoryManagerId = &v
}

// GetRepositoryPublicId returns the RepositoryPublicId field value if set, zero value otherwise.
func (o *ApiRepositoryComponentEvaluationResultList) GetRepositoryPublicId() string {
	if o == nil || IsNil(o.RepositoryPublicId) {
		var ret string
		return ret
	}
	return *o.RepositoryPublicId
}

// GetRepositoryPublicIdOk returns a tuple with the RepositoryPublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentEvaluationResultList) GetRepositoryPublicIdOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryPublicId) {
		return nil, false
	}
	return o.RepositoryPublicId, true
}

// HasRepositoryPublicId returns a boolean if a field has been set.
func (o *ApiRepositoryComponentEvaluationResultList) HasRepositoryPublicId() bool {
	if o != nil && !IsNil(o.RepositoryPublicId) {
		return true
	}

	return false
}

// SetRepositoryPublicId gets a reference to the given string and assigns it to the RepositoryPublicId field.
func (o *ApiRepositoryComponentEvaluationResultList) SetRepositoryPublicId(v string) {
	o.RepositoryPublicId = &v
}

// GetRepositoryType returns the RepositoryType field value if set, zero value otherwise.
func (o *ApiRepositoryComponentEvaluationResultList) GetRepositoryType() string {
	if o == nil || IsNil(o.RepositoryType) {
		var ret string
		return ret
	}
	return *o.RepositoryType
}

// GetRepositoryTypeOk returns a tuple with the RepositoryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentEvaluationResultList) GetRepositoryTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RepositoryType) {
		return nil, false
	}
	return o.RepositoryType, true
}

// HasRepositoryType returns a boolean if a field has been set.
func (o *ApiRepositoryComponentEvaluationResultList) HasRepositoryType() bool {
	if o != nil && !IsNil(o.RepositoryType) {
		return true
	}

	return false
}

// SetRepositoryType gets a reference to the given string and assigns it to the RepositoryType field.
func (o *ApiRepositoryComponentEvaluationResultList) SetRepositoryType(v string) {
	o.RepositoryType = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ApiRepositoryComponentEvaluationResultList) GetResults() []ApiRepositoryComponentEvaluationResult {
	if o == nil || IsNil(o.Results) {
		var ret []ApiRepositoryComponentEvaluationResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentEvaluationResultList) GetResultsOk() ([]ApiRepositoryComponentEvaluationResult, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ApiRepositoryComponentEvaluationResultList) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ApiRepositoryComponentEvaluationResult and assigns it to the Results field.
func (o *ApiRepositoryComponentEvaluationResultList) SetResults(v []ApiRepositoryComponentEvaluationResult) {
	o.Results = v
}

func (o ApiRepositoryComponentEvaluationResultList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryComponentEvaluationResultList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RepositoryId) {
		toSerialize["repositoryId"] = o.RepositoryId
	}
	if !IsNil(o.RepositoryManagerId) {
		toSerialize["repositoryManagerId"] = o.RepositoryManagerId
	}
	if !IsNil(o.RepositoryPublicId) {
		toSerialize["repositoryPublicId"] = o.RepositoryPublicId
	}
	if !IsNil(o.RepositoryType) {
		toSerialize["repositoryType"] = o.RepositoryType
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableApiRepositoryComponentEvaluationResultList struct {
	value *ApiRepositoryComponentEvaluationResultList
	isSet bool
}

func (v NullableApiRepositoryComponentEvaluationResultList) Get() *ApiRepositoryComponentEvaluationResultList {
	return v.value
}

func (v *NullableApiRepositoryComponentEvaluationResultList) Set(val *ApiRepositoryComponentEvaluationResultList) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryComponentEvaluationResultList) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryComponentEvaluationResultList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryComponentEvaluationResultList(val *ApiRepositoryComponentEvaluationResultList) *NullableApiRepositoryComponentEvaluationResultList {
	return &NullableApiRepositoryComponentEvaluationResultList{value: val, isSet: true}
}

func (v NullableApiRepositoryComponentEvaluationResultList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryComponentEvaluationResultList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


