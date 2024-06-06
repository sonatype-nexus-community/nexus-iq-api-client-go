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

// checks if the ApiLicenseLegalFileDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicenseLegalFileDTO{}

// ApiLicenseLegalFileDTO struct for ApiLicenseLegalFileDTO
type ApiLicenseLegalFileDTO struct {
	Content *string `json:"content,omitempty"`
	Id *string `json:"id,omitempty"`
	OriginalContentHash *string `json:"originalContentHash,omitempty"`
	RelPath *string `json:"relPath,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewApiLicenseLegalFileDTO instantiates a new ApiLicenseLegalFileDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicenseLegalFileDTO() *ApiLicenseLegalFileDTO {
	this := ApiLicenseLegalFileDTO{}
	return &this
}

// NewApiLicenseLegalFileDTOWithDefaults instantiates a new ApiLicenseLegalFileDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicenseLegalFileDTOWithDefaults() *ApiLicenseLegalFileDTO {
	this := ApiLicenseLegalFileDTO{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ApiLicenseLegalFileDTO) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalFileDTO) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ApiLicenseLegalFileDTO) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ApiLicenseLegalFileDTO) SetContent(v string) {
	o.Content = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiLicenseLegalFileDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalFileDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiLicenseLegalFileDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiLicenseLegalFileDTO) SetId(v string) {
	o.Id = &v
}

// GetOriginalContentHash returns the OriginalContentHash field value if set, zero value otherwise.
func (o *ApiLicenseLegalFileDTO) GetOriginalContentHash() string {
	if o == nil || IsNil(o.OriginalContentHash) {
		var ret string
		return ret
	}
	return *o.OriginalContentHash
}

// GetOriginalContentHashOk returns a tuple with the OriginalContentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalFileDTO) GetOriginalContentHashOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalContentHash) {
		return nil, false
	}
	return o.OriginalContentHash, true
}

// HasOriginalContentHash returns a boolean if a field has been set.
func (o *ApiLicenseLegalFileDTO) HasOriginalContentHash() bool {
	if o != nil && !IsNil(o.OriginalContentHash) {
		return true
	}

	return false
}

// SetOriginalContentHash gets a reference to the given string and assigns it to the OriginalContentHash field.
func (o *ApiLicenseLegalFileDTO) SetOriginalContentHash(v string) {
	o.OriginalContentHash = &v
}

// GetRelPath returns the RelPath field value if set, zero value otherwise.
func (o *ApiLicenseLegalFileDTO) GetRelPath() string {
	if o == nil || IsNil(o.RelPath) {
		var ret string
		return ret
	}
	return *o.RelPath
}

// GetRelPathOk returns a tuple with the RelPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalFileDTO) GetRelPathOk() (*string, bool) {
	if o == nil || IsNil(o.RelPath) {
		return nil, false
	}
	return o.RelPath, true
}

// HasRelPath returns a boolean if a field has been set.
func (o *ApiLicenseLegalFileDTO) HasRelPath() bool {
	if o != nil && !IsNil(o.RelPath) {
		return true
	}

	return false
}

// SetRelPath gets a reference to the given string and assigns it to the RelPath field.
func (o *ApiLicenseLegalFileDTO) SetRelPath(v string) {
	o.RelPath = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiLicenseLegalFileDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalFileDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiLicenseLegalFileDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiLicenseLegalFileDTO) SetStatus(v string) {
	o.Status = &v
}

func (o ApiLicenseLegalFileDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicenseLegalFileDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.OriginalContentHash) {
		toSerialize["originalContentHash"] = o.OriginalContentHash
	}
	if !IsNil(o.RelPath) {
		toSerialize["relPath"] = o.RelPath
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableApiLicenseLegalFileDTO struct {
	value *ApiLicenseLegalFileDTO
	isSet bool
}

func (v NullableApiLicenseLegalFileDTO) Get() *ApiLicenseLegalFileDTO {
	return v.value
}

func (v *NullableApiLicenseLegalFileDTO) Set(val *ApiLicenseLegalFileDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicenseLegalFileDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicenseLegalFileDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicenseLegalFileDTO(val *ApiLicenseLegalFileDTO) *NullableApiLicenseLegalFileDTO {
	return &NullableApiLicenseLegalFileDTO{value: val, isSet: true}
}

func (v NullableApiLicenseLegalFileDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicenseLegalFileDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


