/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.182.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
)

// checks if the ApiSamlConfigurationDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSamlConfigurationDTO{}

// ApiSamlConfigurationDTO struct for ApiSamlConfigurationDTO
type ApiSamlConfigurationDTO struct {
	EmailAttributeName *string `json:"emailAttributeName,omitempty"`
	EntityId *string `json:"entityId,omitempty"`
	FirstNameAttributeName *string `json:"firstNameAttributeName,omitempty"`
	GroupsAttributeName *string `json:"groupsAttributeName,omitempty"`
	IdentityProviderName *string `json:"identityProviderName,omitempty"`
	LastNameAttributeName *string `json:"lastNameAttributeName,omitempty"`
	UsernameAttributeName *string `json:"usernameAttributeName,omitempty"`
	ValidateAssertionSignature *bool `json:"validateAssertionSignature,omitempty"`
	ValidateResponseSignature *bool `json:"validateResponseSignature,omitempty"`
}

// NewApiSamlConfigurationDTO instantiates a new ApiSamlConfigurationDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSamlConfigurationDTO() *ApiSamlConfigurationDTO {
	this := ApiSamlConfigurationDTO{}
	return &this
}

// NewApiSamlConfigurationDTOWithDefaults instantiates a new ApiSamlConfigurationDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSamlConfigurationDTOWithDefaults() *ApiSamlConfigurationDTO {
	this := ApiSamlConfigurationDTO{}
	return &this
}

// GetEmailAttributeName returns the EmailAttributeName field value if set, zero value otherwise.
func (o *ApiSamlConfigurationDTO) GetEmailAttributeName() string {
	if o == nil || IsNil(o.EmailAttributeName) {
		var ret string
		return ret
	}
	return *o.EmailAttributeName
}

// GetEmailAttributeNameOk returns a tuple with the EmailAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSamlConfigurationDTO) GetEmailAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAttributeName) {
		return nil, false
	}
	return o.EmailAttributeName, true
}

// HasEmailAttributeName returns a boolean if a field has been set.
func (o *ApiSamlConfigurationDTO) HasEmailAttributeName() bool {
	if o != nil && !IsNil(o.EmailAttributeName) {
		return true
	}

	return false
}

// SetEmailAttributeName gets a reference to the given string and assigns it to the EmailAttributeName field.
func (o *ApiSamlConfigurationDTO) SetEmailAttributeName(v string) {
	o.EmailAttributeName = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *ApiSamlConfigurationDTO) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSamlConfigurationDTO) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *ApiSamlConfigurationDTO) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *ApiSamlConfigurationDTO) SetEntityId(v string) {
	o.EntityId = &v
}

// GetFirstNameAttributeName returns the FirstNameAttributeName field value if set, zero value otherwise.
func (o *ApiSamlConfigurationDTO) GetFirstNameAttributeName() string {
	if o == nil || IsNil(o.FirstNameAttributeName) {
		var ret string
		return ret
	}
	return *o.FirstNameAttributeName
}

// GetFirstNameAttributeNameOk returns a tuple with the FirstNameAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSamlConfigurationDTO) GetFirstNameAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstNameAttributeName) {
		return nil, false
	}
	return o.FirstNameAttributeName, true
}

// HasFirstNameAttributeName returns a boolean if a field has been set.
func (o *ApiSamlConfigurationDTO) HasFirstNameAttributeName() bool {
	if o != nil && !IsNil(o.FirstNameAttributeName) {
		return true
	}

	return false
}

// SetFirstNameAttributeName gets a reference to the given string and assigns it to the FirstNameAttributeName field.
func (o *ApiSamlConfigurationDTO) SetFirstNameAttributeName(v string) {
	o.FirstNameAttributeName = &v
}

// GetGroupsAttributeName returns the GroupsAttributeName field value if set, zero value otherwise.
func (o *ApiSamlConfigurationDTO) GetGroupsAttributeName() string {
	if o == nil || IsNil(o.GroupsAttributeName) {
		var ret string
		return ret
	}
	return *o.GroupsAttributeName
}

// GetGroupsAttributeNameOk returns a tuple with the GroupsAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSamlConfigurationDTO) GetGroupsAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupsAttributeName) {
		return nil, false
	}
	return o.GroupsAttributeName, true
}

// HasGroupsAttributeName returns a boolean if a field has been set.
func (o *ApiSamlConfigurationDTO) HasGroupsAttributeName() bool {
	if o != nil && !IsNil(o.GroupsAttributeName) {
		return true
	}

	return false
}

// SetGroupsAttributeName gets a reference to the given string and assigns it to the GroupsAttributeName field.
func (o *ApiSamlConfigurationDTO) SetGroupsAttributeName(v string) {
	o.GroupsAttributeName = &v
}

// GetIdentityProviderName returns the IdentityProviderName field value if set, zero value otherwise.
func (o *ApiSamlConfigurationDTO) GetIdentityProviderName() string {
	if o == nil || IsNil(o.IdentityProviderName) {
		var ret string
		return ret
	}
	return *o.IdentityProviderName
}

// GetIdentityProviderNameOk returns a tuple with the IdentityProviderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSamlConfigurationDTO) GetIdentityProviderNameOk() (*string, bool) {
	if o == nil || IsNil(o.IdentityProviderName) {
		return nil, false
	}
	return o.IdentityProviderName, true
}

// HasIdentityProviderName returns a boolean if a field has been set.
func (o *ApiSamlConfigurationDTO) HasIdentityProviderName() bool {
	if o != nil && !IsNil(o.IdentityProviderName) {
		return true
	}

	return false
}

// SetIdentityProviderName gets a reference to the given string and assigns it to the IdentityProviderName field.
func (o *ApiSamlConfigurationDTO) SetIdentityProviderName(v string) {
	o.IdentityProviderName = &v
}

// GetLastNameAttributeName returns the LastNameAttributeName field value if set, zero value otherwise.
func (o *ApiSamlConfigurationDTO) GetLastNameAttributeName() string {
	if o == nil || IsNil(o.LastNameAttributeName) {
		var ret string
		return ret
	}
	return *o.LastNameAttributeName
}

// GetLastNameAttributeNameOk returns a tuple with the LastNameAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSamlConfigurationDTO) GetLastNameAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastNameAttributeName) {
		return nil, false
	}
	return o.LastNameAttributeName, true
}

// HasLastNameAttributeName returns a boolean if a field has been set.
func (o *ApiSamlConfigurationDTO) HasLastNameAttributeName() bool {
	if o != nil && !IsNil(o.LastNameAttributeName) {
		return true
	}

	return false
}

// SetLastNameAttributeName gets a reference to the given string and assigns it to the LastNameAttributeName field.
func (o *ApiSamlConfigurationDTO) SetLastNameAttributeName(v string) {
	o.LastNameAttributeName = &v
}

// GetUsernameAttributeName returns the UsernameAttributeName field value if set, zero value otherwise.
func (o *ApiSamlConfigurationDTO) GetUsernameAttributeName() string {
	if o == nil || IsNil(o.UsernameAttributeName) {
		var ret string
		return ret
	}
	return *o.UsernameAttributeName
}

// GetUsernameAttributeNameOk returns a tuple with the UsernameAttributeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSamlConfigurationDTO) GetUsernameAttributeNameOk() (*string, bool) {
	if o == nil || IsNil(o.UsernameAttributeName) {
		return nil, false
	}
	return o.UsernameAttributeName, true
}

// HasUsernameAttributeName returns a boolean if a field has been set.
func (o *ApiSamlConfigurationDTO) HasUsernameAttributeName() bool {
	if o != nil && !IsNil(o.UsernameAttributeName) {
		return true
	}

	return false
}

// SetUsernameAttributeName gets a reference to the given string and assigns it to the UsernameAttributeName field.
func (o *ApiSamlConfigurationDTO) SetUsernameAttributeName(v string) {
	o.UsernameAttributeName = &v
}

// GetValidateAssertionSignature returns the ValidateAssertionSignature field value if set, zero value otherwise.
func (o *ApiSamlConfigurationDTO) GetValidateAssertionSignature() bool {
	if o == nil || IsNil(o.ValidateAssertionSignature) {
		var ret bool
		return ret
	}
	return *o.ValidateAssertionSignature
}

// GetValidateAssertionSignatureOk returns a tuple with the ValidateAssertionSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSamlConfigurationDTO) GetValidateAssertionSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidateAssertionSignature) {
		return nil, false
	}
	return o.ValidateAssertionSignature, true
}

// HasValidateAssertionSignature returns a boolean if a field has been set.
func (o *ApiSamlConfigurationDTO) HasValidateAssertionSignature() bool {
	if o != nil && !IsNil(o.ValidateAssertionSignature) {
		return true
	}

	return false
}

// SetValidateAssertionSignature gets a reference to the given bool and assigns it to the ValidateAssertionSignature field.
func (o *ApiSamlConfigurationDTO) SetValidateAssertionSignature(v bool) {
	o.ValidateAssertionSignature = &v
}

// GetValidateResponseSignature returns the ValidateResponseSignature field value if set, zero value otherwise.
func (o *ApiSamlConfigurationDTO) GetValidateResponseSignature() bool {
	if o == nil || IsNil(o.ValidateResponseSignature) {
		var ret bool
		return ret
	}
	return *o.ValidateResponseSignature
}

// GetValidateResponseSignatureOk returns a tuple with the ValidateResponseSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSamlConfigurationDTO) GetValidateResponseSignatureOk() (*bool, bool) {
	if o == nil || IsNil(o.ValidateResponseSignature) {
		return nil, false
	}
	return o.ValidateResponseSignature, true
}

// HasValidateResponseSignature returns a boolean if a field has been set.
func (o *ApiSamlConfigurationDTO) HasValidateResponseSignature() bool {
	if o != nil && !IsNil(o.ValidateResponseSignature) {
		return true
	}

	return false
}

// SetValidateResponseSignature gets a reference to the given bool and assigns it to the ValidateResponseSignature field.
func (o *ApiSamlConfigurationDTO) SetValidateResponseSignature(v bool) {
	o.ValidateResponseSignature = &v
}

func (o ApiSamlConfigurationDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSamlConfigurationDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailAttributeName) {
		toSerialize["emailAttributeName"] = o.EmailAttributeName
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.FirstNameAttributeName) {
		toSerialize["firstNameAttributeName"] = o.FirstNameAttributeName
	}
	if !IsNil(o.GroupsAttributeName) {
		toSerialize["groupsAttributeName"] = o.GroupsAttributeName
	}
	if !IsNil(o.IdentityProviderName) {
		toSerialize["identityProviderName"] = o.IdentityProviderName
	}
	if !IsNil(o.LastNameAttributeName) {
		toSerialize["lastNameAttributeName"] = o.LastNameAttributeName
	}
	if !IsNil(o.UsernameAttributeName) {
		toSerialize["usernameAttributeName"] = o.UsernameAttributeName
	}
	if !IsNil(o.ValidateAssertionSignature) {
		toSerialize["validateAssertionSignature"] = o.ValidateAssertionSignature
	}
	if !IsNil(o.ValidateResponseSignature) {
		toSerialize["validateResponseSignature"] = o.ValidateResponseSignature
	}
	return toSerialize, nil
}

type NullableApiSamlConfigurationDTO struct {
	value *ApiSamlConfigurationDTO
	isSet bool
}

func (v NullableApiSamlConfigurationDTO) Get() *ApiSamlConfigurationDTO {
	return v.value
}

func (v *NullableApiSamlConfigurationDTO) Set(val *ApiSamlConfigurationDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSamlConfigurationDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSamlConfigurationDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSamlConfigurationDTO(val *ApiSamlConfigurationDTO) *NullableApiSamlConfigurationDTO {
	return &NullableApiSamlConfigurationDTO{value: val, isSet: true}
}

func (v NullableApiSamlConfigurationDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSamlConfigurationDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


