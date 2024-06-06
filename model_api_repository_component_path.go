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

// checks if the ApiRepositoryComponentPath type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiRepositoryComponentPath{}

// ApiRepositoryComponentPath struct for ApiRepositoryComponentPath
type ApiRepositoryComponentPath struct {
	Pathname *string `json:"pathname,omitempty"`
	Quarantine *bool `json:"quarantine,omitempty"`
}

// NewApiRepositoryComponentPath instantiates a new ApiRepositoryComponentPath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRepositoryComponentPath() *ApiRepositoryComponentPath {
	this := ApiRepositoryComponentPath{}
	return &this
}

// NewApiRepositoryComponentPathWithDefaults instantiates a new ApiRepositoryComponentPath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRepositoryComponentPathWithDefaults() *ApiRepositoryComponentPath {
	this := ApiRepositoryComponentPath{}
	return &this
}

// GetPathname returns the Pathname field value if set, zero value otherwise.
func (o *ApiRepositoryComponentPath) GetPathname() string {
	if o == nil || IsNil(o.Pathname) {
		var ret string
		return ret
	}
	return *o.Pathname
}

// GetPathnameOk returns a tuple with the Pathname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentPath) GetPathnameOk() (*string, bool) {
	if o == nil || IsNil(o.Pathname) {
		return nil, false
	}
	return o.Pathname, true
}

// HasPathname returns a boolean if a field has been set.
func (o *ApiRepositoryComponentPath) HasPathname() bool {
	if o != nil && !IsNil(o.Pathname) {
		return true
	}

	return false
}

// SetPathname gets a reference to the given string and assigns it to the Pathname field.
func (o *ApiRepositoryComponentPath) SetPathname(v string) {
	o.Pathname = &v
}

// GetQuarantine returns the Quarantine field value if set, zero value otherwise.
func (o *ApiRepositoryComponentPath) GetQuarantine() bool {
	if o == nil || IsNil(o.Quarantine) {
		var ret bool
		return ret
	}
	return *o.Quarantine
}

// GetQuarantineOk returns a tuple with the Quarantine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRepositoryComponentPath) GetQuarantineOk() (*bool, bool) {
	if o == nil || IsNil(o.Quarantine) {
		return nil, false
	}
	return o.Quarantine, true
}

// HasQuarantine returns a boolean if a field has been set.
func (o *ApiRepositoryComponentPath) HasQuarantine() bool {
	if o != nil && !IsNil(o.Quarantine) {
		return true
	}

	return false
}

// SetQuarantine gets a reference to the given bool and assigns it to the Quarantine field.
func (o *ApiRepositoryComponentPath) SetQuarantine(v bool) {
	o.Quarantine = &v
}

func (o ApiRepositoryComponentPath) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRepositoryComponentPath) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pathname) {
		toSerialize["pathname"] = o.Pathname
	}
	if !IsNil(o.Quarantine) {
		toSerialize["quarantine"] = o.Quarantine
	}
	return toSerialize, nil
}

type NullableApiRepositoryComponentPath struct {
	value *ApiRepositoryComponentPath
	isSet bool
}

func (v NullableApiRepositoryComponentPath) Get() *ApiRepositoryComponentPath {
	return v.value
}

func (v *NullableApiRepositoryComponentPath) Set(val *ApiRepositoryComponentPath) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRepositoryComponentPath) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRepositoryComponentPath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRepositoryComponentPath(val *ApiRepositoryComponentPath) *NullableApiRepositoryComponentPath {
	return &NullableApiRepositoryComponentPath{value: val, isSet: true}
}

func (v NullableApiRepositoryComponentPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRepositoryComponentPath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


