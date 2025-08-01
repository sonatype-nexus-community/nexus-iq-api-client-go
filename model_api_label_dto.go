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

// checks if the ApiLabelDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLabelDTO{}

// ApiLabelDTO struct for ApiLabelDTO
type ApiLabelDTO struct {
	Color *string `json:"color,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	OwnerId *string `json:"ownerId,omitempty"`
	OwnerType *string `json:"ownerType,omitempty"`
}

// NewApiLabelDTO instantiates a new ApiLabelDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLabelDTO() *ApiLabelDTO {
	this := ApiLabelDTO{}
	return &this
}

// NewApiLabelDTOWithDefaults instantiates a new ApiLabelDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLabelDTOWithDefaults() *ApiLabelDTO {
	this := ApiLabelDTO{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *ApiLabelDTO) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabelDTO) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *ApiLabelDTO) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *ApiLabelDTO) SetColor(v string) {
	o.Color = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiLabelDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabelDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiLabelDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiLabelDTO) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiLabelDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabelDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiLabelDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiLabelDTO) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ApiLabelDTO) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabelDTO) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ApiLabelDTO) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ApiLabelDTO) SetLabel(v string) {
	o.Label = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ApiLabelDTO) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabelDTO) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ApiLabelDTO) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ApiLabelDTO) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *ApiLabelDTO) GetOwnerType() string {
	if o == nil || IsNil(o.OwnerType) {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLabelDTO) GetOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerType) {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *ApiLabelDTO) HasOwnerType() bool {
	if o != nil && !IsNil(o.OwnerType) {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *ApiLabelDTO) SetOwnerType(v string) {
	o.OwnerType = &v
}

func (o ApiLabelDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLabelDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.OwnerId) {
		toSerialize["ownerId"] = o.OwnerId
	}
	if !IsNil(o.OwnerType) {
		toSerialize["ownerType"] = o.OwnerType
	}
	return toSerialize, nil
}

type NullableApiLabelDTO struct {
	value *ApiLabelDTO
	isSet bool
}

func (v NullableApiLabelDTO) Get() *ApiLabelDTO {
	return v.value
}

func (v *NullableApiLabelDTO) Set(val *ApiLabelDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLabelDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLabelDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLabelDTO(val *ApiLabelDTO) *NullableApiLabelDTO {
	return &NullableApiLabelDTO{value: val, isSet: true}
}

func (v NullableApiLabelDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLabelDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


