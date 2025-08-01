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

// checks if the ApiRepositoryComponentEvaluationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryComponentEvaluationRequest{}

// ApiRepositoryComponentEvaluationRequest struct for ApiRepositoryComponentEvaluationRequest
type ApiRepositoryComponentEvaluationRequest struct {
	Hash *string `json:"hash,omitempty"`
	PackageUrl *string `json:"packageUrl,omitempty"`
	Pathname *string `json:"pathname,omitempty"`
}

// NewApiRepositoryComponentEvaluationRequest instantiates a new ApiRepositoryComponentEvaluationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryComponentEvaluationRequest() *ApiRepositoryComponentEvaluationRequest {
	this := ApiRepositoryComponentEvaluationRequest{}
	return &this
}

// NewApiRepositoryComponentEvaluationRequestWithDefaults instantiates a new ApiRepositoryComponentEvaluationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryComponentEvaluationRequestWithDefaults() *ApiRepositoryComponentEvaluationRequest {
	this := ApiRepositoryComponentEvaluationRequest{}
	return &this
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ApiRepositoryComponentEvaluationRequest) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentEvaluationRequest) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ApiRepositoryComponentEvaluationRequest) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ApiRepositoryComponentEvaluationRequest) SetHash(v string) {
	o.Hash = &v
}

// GetPackageUrl returns the PackageUrl field value if set, zero value otherwise.
func (o *ApiRepositoryComponentEvaluationRequest) GetPackageUrl() string {
	if o == nil || IsNil(o.PackageUrl) {
		var ret string
		return ret
	}
	return *o.PackageUrl
}

// GetPackageUrlOk returns a tuple with the PackageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentEvaluationRequest) GetPackageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PackageUrl) {
		return nil, false
	}
	return o.PackageUrl, true
}

// HasPackageUrl returns a boolean if a field has been set.
func (o *ApiRepositoryComponentEvaluationRequest) HasPackageUrl() bool {
	if o != nil && !IsNil(o.PackageUrl) {
		return true
	}

	return false
}

// SetPackageUrl gets a reference to the given string and assigns it to the PackageUrl field.
func (o *ApiRepositoryComponentEvaluationRequest) SetPackageUrl(v string) {
	o.PackageUrl = &v
}

// GetPathname returns the Pathname field value if set, zero value otherwise.
func (o *ApiRepositoryComponentEvaluationRequest) GetPathname() string {
	if o == nil || IsNil(o.Pathname) {
		var ret string
		return ret
	}
	return *o.Pathname
}

// GetPathnameOk returns a tuple with the Pathname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentEvaluationRequest) GetPathnameOk() (*string, bool) {
	if o == nil || IsNil(o.Pathname) {
		return nil, false
	}
	return o.Pathname, true
}

// HasPathname returns a boolean if a field has been set.
func (o *ApiRepositoryComponentEvaluationRequest) HasPathname() bool {
	if o != nil && !IsNil(o.Pathname) {
		return true
	}

	return false
}

// SetPathname gets a reference to the given string and assigns it to the Pathname field.
func (o *ApiRepositoryComponentEvaluationRequest) SetPathname(v string) {
	o.Pathname = &v
}

func (o ApiRepositoryComponentEvaluationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryComponentEvaluationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.PackageUrl) {
		toSerialize["packageUrl"] = o.PackageUrl
	}
	if !IsNil(o.Pathname) {
		toSerialize["pathname"] = o.Pathname
	}
	return toSerialize, nil
}

type NullableApiRepositoryComponentEvaluationRequest struct {
	value *ApiRepositoryComponentEvaluationRequest
	isSet bool
}

func (v NullableApiRepositoryComponentEvaluationRequest) Get() *ApiRepositoryComponentEvaluationRequest {
	return v.value
}

func (v *NullableApiRepositoryComponentEvaluationRequest) Set(val *ApiRepositoryComponentEvaluationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryComponentEvaluationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryComponentEvaluationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryComponentEvaluationRequest(val *ApiRepositoryComponentEvaluationRequest) *NullableApiRepositoryComponentEvaluationRequest {
	return &NullableApiRepositoryComponentEvaluationRequest{value: val, isSet: true}
}

func (v NullableApiRepositoryComponentEvaluationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryComponentEvaluationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


