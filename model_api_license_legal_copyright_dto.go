/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.184.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiLicenseLegalCopyrightDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicenseLegalCopyrightDTO{}

// ApiLicenseLegalCopyrightDTO struct for ApiLicenseLegalCopyrightDTO
type ApiLicenseLegalCopyrightDTO struct {
	Content *string `json:"content,omitempty"`
	Id *string `json:"id,omitempty"`
	OriginalContentHash *string `json:"originalContentHash,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewApiLicenseLegalCopyrightDTO instantiates a new ApiLicenseLegalCopyrightDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicenseLegalCopyrightDTO() *ApiLicenseLegalCopyrightDTO {
	this := ApiLicenseLegalCopyrightDTO{}
	return &this
}

// NewApiLicenseLegalCopyrightDTOWithDefaults instantiates a new ApiLicenseLegalCopyrightDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicenseLegalCopyrightDTOWithDefaults() *ApiLicenseLegalCopyrightDTO {
	this := ApiLicenseLegalCopyrightDTO{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *ApiLicenseLegalCopyrightDTO) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalCopyrightDTO) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *ApiLicenseLegalCopyrightDTO) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *ApiLicenseLegalCopyrightDTO) SetContent(v string) {
	o.Content = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiLicenseLegalCopyrightDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalCopyrightDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiLicenseLegalCopyrightDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiLicenseLegalCopyrightDTO) SetId(v string) {
	o.Id = &v
}

// GetOriginalContentHash returns the OriginalContentHash field value if set, zero value otherwise.
func (o *ApiLicenseLegalCopyrightDTO) GetOriginalContentHash() string {
	if o == nil || IsNil(o.OriginalContentHash) {
		var ret string
		return ret
	}
	return *o.OriginalContentHash
}

// GetOriginalContentHashOk returns a tuple with the OriginalContentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalCopyrightDTO) GetOriginalContentHashOk() (*string, bool) {
	if o == nil || IsNil(o.OriginalContentHash) {
		return nil, false
	}
	return o.OriginalContentHash, true
}

// HasOriginalContentHash returns a boolean if a field has been set.
func (o *ApiLicenseLegalCopyrightDTO) HasOriginalContentHash() bool {
	if o != nil && !IsNil(o.OriginalContentHash) {
		return true
	}

	return false
}

// SetOriginalContentHash gets a reference to the given string and assigns it to the OriginalContentHash field.
func (o *ApiLicenseLegalCopyrightDTO) SetOriginalContentHash(v string) {
	o.OriginalContentHash = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiLicenseLegalCopyrightDTO) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalCopyrightDTO) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiLicenseLegalCopyrightDTO) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiLicenseLegalCopyrightDTO) SetStatus(v string) {
	o.Status = &v
}

func (o ApiLicenseLegalCopyrightDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicenseLegalCopyrightDTO) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableApiLicenseLegalCopyrightDTO struct {
	value *ApiLicenseLegalCopyrightDTO
	isSet bool
}

func (v NullableApiLicenseLegalCopyrightDTO) Get() *ApiLicenseLegalCopyrightDTO {
	return v.value
}

func (v *NullableApiLicenseLegalCopyrightDTO) Set(val *ApiLicenseLegalCopyrightDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicenseLegalCopyrightDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicenseLegalCopyrightDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicenseLegalCopyrightDTO(val *ApiLicenseLegalCopyrightDTO) *NullableApiLicenseLegalCopyrightDTO {
	return &NullableApiLicenseLegalCopyrightDTO{value: val, isSet: true}
}

func (v NullableApiLicenseLegalCopyrightDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicenseLegalCopyrightDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


