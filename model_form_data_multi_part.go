/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 164
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FormDataMultiPart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FormDataMultiPart{}

// FormDataMultiPart struct for FormDataMultiPart
type FormDataMultiPart struct {
	BodyParts []BodyPart `json:"bodyParts,omitempty"`
	ContentDisposition *ContentDisposition `json:"contentDisposition,omitempty"`
	Entity map[string]interface{} `json:"entity,omitempty"`
	Fields *map[string][]FormDataBodyPart `json:"fields,omitempty"`
	Headers *map[string][]string `json:"headers,omitempty"`
	MediaType *BodyPartMediaType `json:"mediaType,omitempty"`
	MessageBodyWorkers map[string]interface{} `json:"messageBodyWorkers,omitempty"`
	ParameterizedHeaders *map[string][]ParameterizedHeader `json:"parameterizedHeaders,omitempty"`
	Parent *MultiPart `json:"parent,omitempty"`
	Providers map[string]interface{} `json:"providers,omitempty"`
}

// NewFormDataMultiPart instantiates a new FormDataMultiPart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormDataMultiPart() *FormDataMultiPart {
	this := FormDataMultiPart{}
	return &this
}

// NewFormDataMultiPartWithDefaults instantiates a new FormDataMultiPart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormDataMultiPartWithDefaults() *FormDataMultiPart {
	this := FormDataMultiPart{}
	return &this
}

// GetBodyParts returns the BodyParts field value if set, zero value otherwise.
func (o *FormDataMultiPart) GetBodyParts() []BodyPart {
	if o == nil || IsNil(o.BodyParts) {
		var ret []BodyPart
		return ret
	}
	return o.BodyParts
}

// GetBodyPartsOk returns a tuple with the BodyParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataMultiPart) GetBodyPartsOk() ([]BodyPart, bool) {
	if o == nil || IsNil(o.BodyParts) {
		return nil, false
	}
	return o.BodyParts, true
}

// HasBodyParts returns a boolean if a field has been set.
func (o *FormDataMultiPart) HasBodyParts() bool {
	if o != nil && !IsNil(o.BodyParts) {
		return true
	}

	return false
}

// SetBodyParts gets a reference to the given []BodyPart and assigns it to the BodyParts field.
func (o *FormDataMultiPart) SetBodyParts(v []BodyPart) {
	o.BodyParts = v
}

// GetContentDisposition returns the ContentDisposition field value if set, zero value otherwise.
func (o *FormDataMultiPart) GetContentDisposition() ContentDisposition {
	if o == nil || IsNil(o.ContentDisposition) {
		var ret ContentDisposition
		return ret
	}
	return *o.ContentDisposition
}

// GetContentDispositionOk returns a tuple with the ContentDisposition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataMultiPart) GetContentDispositionOk() (*ContentDisposition, bool) {
	if o == nil || IsNil(o.ContentDisposition) {
		return nil, false
	}
	return o.ContentDisposition, true
}

// HasContentDisposition returns a boolean if a field has been set.
func (o *FormDataMultiPart) HasContentDisposition() bool {
	if o != nil && !IsNil(o.ContentDisposition) {
		return true
	}

	return false
}

// SetContentDisposition gets a reference to the given ContentDisposition and assigns it to the ContentDisposition field.
func (o *FormDataMultiPart) SetContentDisposition(v ContentDisposition) {
	o.ContentDisposition = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *FormDataMultiPart) GetEntity() map[string]interface{} {
	if o == nil || IsNil(o.Entity) {
		var ret map[string]interface{}
		return ret
	}
	return o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataMultiPart) GetEntityOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Entity) {
		return map[string]interface{}{}, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *FormDataMultiPart) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given map[string]interface{} and assigns it to the Entity field.
func (o *FormDataMultiPart) SetEntity(v map[string]interface{}) {
	o.Entity = v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *FormDataMultiPart) GetFields() map[string][]FormDataBodyPart {
	if o == nil || IsNil(o.Fields) {
		var ret map[string][]FormDataBodyPart
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataMultiPart) GetFieldsOk() (*map[string][]FormDataBodyPart, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *FormDataMultiPart) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given map[string][]FormDataBodyPart and assigns it to the Fields field.
func (o *FormDataMultiPart) SetFields(v map[string][]FormDataBodyPart) {
	o.Fields = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *FormDataMultiPart) GetHeaders() map[string][]string {
	if o == nil || IsNil(o.Headers) {
		var ret map[string][]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataMultiPart) GetHeadersOk() (*map[string][]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *FormDataMultiPart) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string][]string and assigns it to the Headers field.
func (o *FormDataMultiPart) SetHeaders(v map[string][]string) {
	o.Headers = &v
}

// GetMediaType returns the MediaType field value if set, zero value otherwise.
func (o *FormDataMultiPart) GetMediaType() BodyPartMediaType {
	if o == nil || IsNil(o.MediaType) {
		var ret BodyPartMediaType
		return ret
	}
	return *o.MediaType
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataMultiPart) GetMediaTypeOk() (*BodyPartMediaType, bool) {
	if o == nil || IsNil(o.MediaType) {
		return nil, false
	}
	return o.MediaType, true
}

// HasMediaType returns a boolean if a field has been set.
func (o *FormDataMultiPart) HasMediaType() bool {
	if o != nil && !IsNil(o.MediaType) {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given BodyPartMediaType and assigns it to the MediaType field.
func (o *FormDataMultiPart) SetMediaType(v BodyPartMediaType) {
	o.MediaType = &v
}

// GetMessageBodyWorkers returns the MessageBodyWorkers field value if set, zero value otherwise.
func (o *FormDataMultiPart) GetMessageBodyWorkers() map[string]interface{} {
	if o == nil || IsNil(o.MessageBodyWorkers) {
		var ret map[string]interface{}
		return ret
	}
	return o.MessageBodyWorkers
}

// GetMessageBodyWorkersOk returns a tuple with the MessageBodyWorkers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataMultiPart) GetMessageBodyWorkersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.MessageBodyWorkers) {
		return map[string]interface{}{}, false
	}
	return o.MessageBodyWorkers, true
}

// HasMessageBodyWorkers returns a boolean if a field has been set.
func (o *FormDataMultiPart) HasMessageBodyWorkers() bool {
	if o != nil && !IsNil(o.MessageBodyWorkers) {
		return true
	}

	return false
}

// SetMessageBodyWorkers gets a reference to the given map[string]interface{} and assigns it to the MessageBodyWorkers field.
func (o *FormDataMultiPart) SetMessageBodyWorkers(v map[string]interface{}) {
	o.MessageBodyWorkers = v
}

// GetParameterizedHeaders returns the ParameterizedHeaders field value if set, zero value otherwise.
func (o *FormDataMultiPart) GetParameterizedHeaders() map[string][]ParameterizedHeader {
	if o == nil || IsNil(o.ParameterizedHeaders) {
		var ret map[string][]ParameterizedHeader
		return ret
	}
	return *o.ParameterizedHeaders
}

// GetParameterizedHeadersOk returns a tuple with the ParameterizedHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataMultiPart) GetParameterizedHeadersOk() (*map[string][]ParameterizedHeader, bool) {
	if o == nil || IsNil(o.ParameterizedHeaders) {
		return nil, false
	}
	return o.ParameterizedHeaders, true
}

// HasParameterizedHeaders returns a boolean if a field has been set.
func (o *FormDataMultiPart) HasParameterizedHeaders() bool {
	if o != nil && !IsNil(o.ParameterizedHeaders) {
		return true
	}

	return false
}

// SetParameterizedHeaders gets a reference to the given map[string][]ParameterizedHeader and assigns it to the ParameterizedHeaders field.
func (o *FormDataMultiPart) SetParameterizedHeaders(v map[string][]ParameterizedHeader) {
	o.ParameterizedHeaders = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *FormDataMultiPart) GetParent() MultiPart {
	if o == nil || IsNil(o.Parent) {
		var ret MultiPart
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataMultiPart) GetParentOk() (*MultiPart, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *FormDataMultiPart) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given MultiPart and assigns it to the Parent field.
func (o *FormDataMultiPart) SetParent(v MultiPart) {
	o.Parent = &v
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *FormDataMultiPart) GetProviders() map[string]interface{} {
	if o == nil || IsNil(o.Providers) {
		var ret map[string]interface{}
		return ret
	}
	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormDataMultiPart) GetProvidersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Providers) {
		return map[string]interface{}{}, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *FormDataMultiPart) HasProviders() bool {
	if o != nil && !IsNil(o.Providers) {
		return true
	}

	return false
}

// SetProviders gets a reference to the given map[string]interface{} and assigns it to the Providers field.
func (o *FormDataMultiPart) SetProviders(v map[string]interface{}) {
	o.Providers = v
}

func (o FormDataMultiPart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FormDataMultiPart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BodyParts) {
		toSerialize["bodyParts"] = o.BodyParts
	}
	if !IsNil(o.ContentDisposition) {
		toSerialize["contentDisposition"] = o.ContentDisposition
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.MediaType) {
		toSerialize["mediaType"] = o.MediaType
	}
	if !IsNil(o.MessageBodyWorkers) {
		toSerialize["messageBodyWorkers"] = o.MessageBodyWorkers
	}
	if !IsNil(o.ParameterizedHeaders) {
		toSerialize["parameterizedHeaders"] = o.ParameterizedHeaders
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Providers) {
		toSerialize["providers"] = o.Providers
	}
	return toSerialize, nil
}

type NullableFormDataMultiPart struct {
	value *FormDataMultiPart
	isSet bool
}

func (v NullableFormDataMultiPart) Get() *FormDataMultiPart {
	return v.value
}

func (v *NullableFormDataMultiPart) Set(val *FormDataMultiPart) {
	v.value = val
	v.isSet = true
}

func (v NullableFormDataMultiPart) IsSet() bool {
	return v.isSet
}

func (v *NullableFormDataMultiPart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormDataMultiPart(val *FormDataMultiPart) *NullableFormDataMultiPart {
	return &NullableFormDataMultiPart{value: val, isSet: true}
}

func (v NullableFormDataMultiPart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormDataMultiPart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


