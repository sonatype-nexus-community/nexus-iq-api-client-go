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

// checks if the ComponentFact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentFact{}

// ComponentFact struct for ComponentFact
type ComponentFact struct {
	ComponentIdentifier *ComponentIdentifier `json:"componentIdentifier,omitempty"`
	ConstraintFacts []ConstraintFact `json:"constraintFacts,omitempty"`
	DisplayName *ComponentDisplayName `json:"displayName,omitempty"`
	Hash *string `json:"hash,omitempty"`
	Pathnames []string `json:"pathnames,omitempty"`
}

// NewComponentFact instantiates a new ComponentFact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentFact() *ComponentFact {
	this := ComponentFact{}
	return &this
}

// NewComponentFactWithDefaults instantiates a new ComponentFact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentFactWithDefaults() *ComponentFact {
	this := ComponentFact{}
	return &this
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *ComponentFact) GetComponentIdentifier() ComponentIdentifier {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ComponentIdentifier
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentFact) GetComponentIdentifierOk() (*ComponentIdentifier, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *ComponentFact) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ComponentIdentifier and assigns it to the ComponentIdentifier field.
func (o *ComponentFact) SetComponentIdentifier(v ComponentIdentifier) {
	o.ComponentIdentifier = &v
}

// GetConstraintFacts returns the ConstraintFacts field value if set, zero value otherwise.
func (o *ComponentFact) GetConstraintFacts() []ConstraintFact {
	if o == nil || IsNil(o.ConstraintFacts) {
		var ret []ConstraintFact
		return ret
	}
	return o.ConstraintFacts
}

// GetConstraintFactsOk returns a tuple with the ConstraintFacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentFact) GetConstraintFactsOk() ([]ConstraintFact, bool) {
	if o == nil || IsNil(o.ConstraintFacts) {
		return nil, false
	}
	return o.ConstraintFacts, true
}

// HasConstraintFacts returns a boolean if a field has been set.
func (o *ComponentFact) HasConstraintFacts() bool {
	if o != nil && !IsNil(o.ConstraintFacts) {
		return true
	}

	return false
}

// SetConstraintFacts gets a reference to the given []ConstraintFact and assigns it to the ConstraintFacts field.
func (o *ComponentFact) SetConstraintFacts(v []ConstraintFact) {
	o.ConstraintFacts = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ComponentFact) GetDisplayName() ComponentDisplayName {
	if o == nil || IsNil(o.DisplayName) {
		var ret ComponentDisplayName
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentFact) GetDisplayNameOk() (*ComponentDisplayName, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ComponentFact) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given ComponentDisplayName and assigns it to the DisplayName field.
func (o *ComponentFact) SetDisplayName(v ComponentDisplayName) {
	o.DisplayName = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *ComponentFact) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentFact) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *ComponentFact) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *ComponentFact) SetHash(v string) {
	o.Hash = &v
}

// GetPathnames returns the Pathnames field value if set, zero value otherwise.
func (o *ComponentFact) GetPathnames() []string {
	if o == nil || IsNil(o.Pathnames) {
		var ret []string
		return ret
	}
	return o.Pathnames
}

// GetPathnamesOk returns a tuple with the Pathnames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentFact) GetPathnamesOk() ([]string, bool) {
	if o == nil || IsNil(o.Pathnames) {
		return nil, false
	}
	return o.Pathnames, true
}

// HasPathnames returns a boolean if a field has been set.
func (o *ComponentFact) HasPathnames() bool {
	if o != nil && !IsNil(o.Pathnames) {
		return true
	}

	return false
}

// SetPathnames gets a reference to the given []string and assigns it to the Pathnames field.
func (o *ComponentFact) SetPathnames(v []string) {
	o.Pathnames = v
}

func (o ComponentFact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentFact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.ConstraintFacts) {
		toSerialize["constraintFacts"] = o.ConstraintFacts
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.Pathnames) {
		toSerialize["pathnames"] = o.Pathnames
	}
	return toSerialize, nil
}

type NullableComponentFact struct {
	value *ComponentFact
	isSet bool
}

func (v NullableComponentFact) Get() *ComponentFact {
	return v.value
}

func (v *NullableComponentFact) Set(val *ComponentFact) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentFact) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentFact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentFact(val *ComponentFact) *NullableComponentFact {
	return &NullableComponentFact{value: val, isSet: true}
}

func (v NullableComponentFact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentFact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


