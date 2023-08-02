/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 165
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"encoding/json"
	"time"
)

// checks if the ApiLicenseLegalDataDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiLicenseLegalDataDTO{}

// ApiLicenseLegalDataDTO struct for ApiLicenseLegalDataDTO
type ApiLicenseLegalDataDTO struct {
	Attributions []ComponentObligationAttributionDTO `json:"attributions,omitempty"`
	ComponentCopyrightId *string `json:"componentCopyrightId,omitempty"`
	ComponentCopyrightLastUpdatedAt *time.Time `json:"componentCopyrightLastUpdatedAt,omitempty"`
	ComponentCopyrightLastUpdatedByUsername *string `json:"componentCopyrightLastUpdatedByUsername,omitempty"`
	ComponentCopyrightScopeOwnerId *string `json:"componentCopyrightScopeOwnerId,omitempty"`
	ComponentLicensesId *string `json:"componentLicensesId,omitempty"`
	ComponentLicensesLastUpdatedAt *time.Time `json:"componentLicensesLastUpdatedAt,omitempty"`
	ComponentLicensesLastUpdatedByUsername *string `json:"componentLicensesLastUpdatedByUsername,omitempty"`
	ComponentLicensesScopeOwnerId *string `json:"componentLicensesScopeOwnerId,omitempty"`
	ComponentNoticesId *string `json:"componentNoticesId,omitempty"`
	ComponentNoticesLastUpdatedAt *time.Time `json:"componentNoticesLastUpdatedAt,omitempty"`
	ComponentNoticesLastUpdatedByUsername *string `json:"componentNoticesLastUpdatedByUsername,omitempty"`
	ComponentNoticesScopeOwnerId *string `json:"componentNoticesScopeOwnerId,omitempty"`
	Copyrights []ApiLicenseLegalCopyrightDTO `json:"copyrights,omitempty"`
	DeclaredLicenses []string `json:"declaredLicenses,omitempty"`
	EffectiveLicenseStatus *string `json:"effectiveLicenseStatus,omitempty"`
	EffectiveLicenses []string `json:"effectiveLicenses,omitempty"`
	HighestEffectiveLicenseThreatGroup *ApiLicenseThreatDTOV2 `json:"highestEffectiveLicenseThreatGroup,omitempty"`
	LicenseFiles []ApiLicenseLegalFileDTO `json:"licenseFiles,omitempty"`
	NoticeFiles []ApiLicenseLegalFileDTO `json:"noticeFiles,omitempty"`
	Obligations []ApiLicenseLegalObligationDTO `json:"obligations,omitempty"`
	ObservedLicenses []string `json:"observedLicenses,omitempty"`
	SourceLinks []LegalSourceLinkDTO `json:"sourceLinks,omitempty"`
}

// NewApiLicenseLegalDataDTO instantiates a new ApiLicenseLegalDataDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiLicenseLegalDataDTO() *ApiLicenseLegalDataDTO {
	this := ApiLicenseLegalDataDTO{}
	return &this
}

// NewApiLicenseLegalDataDTOWithDefaults instantiates a new ApiLicenseLegalDataDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiLicenseLegalDataDTOWithDefaults() *ApiLicenseLegalDataDTO {
	this := ApiLicenseLegalDataDTO{}
	return &this
}

// GetAttributions returns the Attributions field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetAttributions() []ComponentObligationAttributionDTO {
	if o == nil || IsNil(o.Attributions) {
		var ret []ComponentObligationAttributionDTO
		return ret
	}
	return o.Attributions
}

// GetAttributionsOk returns a tuple with the Attributions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetAttributionsOk() ([]ComponentObligationAttributionDTO, bool) {
	if o == nil || IsNil(o.Attributions) {
		return nil, false
	}
	return o.Attributions, true
}

// HasAttributions returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasAttributions() bool {
	if o != nil && !IsNil(o.Attributions) {
		return true
	}

	return false
}

// SetAttributions gets a reference to the given []ComponentObligationAttributionDTO and assigns it to the Attributions field.
func (o *ApiLicenseLegalDataDTO) SetAttributions(v []ComponentObligationAttributionDTO) {
	o.Attributions = v
}

// GetComponentCopyrightId returns the ComponentCopyrightId field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightId() string {
	if o == nil || IsNil(o.ComponentCopyrightId) {
		var ret string
		return ret
	}
	return *o.ComponentCopyrightId
}

// GetComponentCopyrightIdOk returns a tuple with the ComponentCopyrightId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentCopyrightId) {
		return nil, false
	}
	return o.ComponentCopyrightId, true
}

// HasComponentCopyrightId returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasComponentCopyrightId() bool {
	if o != nil && !IsNil(o.ComponentCopyrightId) {
		return true
	}

	return false
}

// SetComponentCopyrightId gets a reference to the given string and assigns it to the ComponentCopyrightId field.
func (o *ApiLicenseLegalDataDTO) SetComponentCopyrightId(v string) {
	o.ComponentCopyrightId = &v
}

// GetComponentCopyrightLastUpdatedAt returns the ComponentCopyrightLastUpdatedAt field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightLastUpdatedAt() time.Time {
	if o == nil || IsNil(o.ComponentCopyrightLastUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.ComponentCopyrightLastUpdatedAt
}

// GetComponentCopyrightLastUpdatedAtOk returns a tuple with the ComponentCopyrightLastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ComponentCopyrightLastUpdatedAt) {
		return nil, false
	}
	return o.ComponentCopyrightLastUpdatedAt, true
}

// HasComponentCopyrightLastUpdatedAt returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasComponentCopyrightLastUpdatedAt() bool {
	if o != nil && !IsNil(o.ComponentCopyrightLastUpdatedAt) {
		return true
	}

	return false
}

// SetComponentCopyrightLastUpdatedAt gets a reference to the given time.Time and assigns it to the ComponentCopyrightLastUpdatedAt field.
func (o *ApiLicenseLegalDataDTO) SetComponentCopyrightLastUpdatedAt(v time.Time) {
	o.ComponentCopyrightLastUpdatedAt = &v
}

// GetComponentCopyrightLastUpdatedByUsername returns the ComponentCopyrightLastUpdatedByUsername field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightLastUpdatedByUsername() string {
	if o == nil || IsNil(o.ComponentCopyrightLastUpdatedByUsername) {
		var ret string
		return ret
	}
	return *o.ComponentCopyrightLastUpdatedByUsername
}

// GetComponentCopyrightLastUpdatedByUsernameOk returns a tuple with the ComponentCopyrightLastUpdatedByUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightLastUpdatedByUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentCopyrightLastUpdatedByUsername) {
		return nil, false
	}
	return o.ComponentCopyrightLastUpdatedByUsername, true
}

// HasComponentCopyrightLastUpdatedByUsername returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasComponentCopyrightLastUpdatedByUsername() bool {
	if o != nil && !IsNil(o.ComponentCopyrightLastUpdatedByUsername) {
		return true
	}

	return false
}

// SetComponentCopyrightLastUpdatedByUsername gets a reference to the given string and assigns it to the ComponentCopyrightLastUpdatedByUsername field.
func (o *ApiLicenseLegalDataDTO) SetComponentCopyrightLastUpdatedByUsername(v string) {
	o.ComponentCopyrightLastUpdatedByUsername = &v
}

// GetComponentCopyrightScopeOwnerId returns the ComponentCopyrightScopeOwnerId field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightScopeOwnerId() string {
	if o == nil || IsNil(o.ComponentCopyrightScopeOwnerId) {
		var ret string
		return ret
	}
	return *o.ComponentCopyrightScopeOwnerId
}

// GetComponentCopyrightScopeOwnerIdOk returns a tuple with the ComponentCopyrightScopeOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetComponentCopyrightScopeOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentCopyrightScopeOwnerId) {
		return nil, false
	}
	return o.ComponentCopyrightScopeOwnerId, true
}

// HasComponentCopyrightScopeOwnerId returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasComponentCopyrightScopeOwnerId() bool {
	if o != nil && !IsNil(o.ComponentCopyrightScopeOwnerId) {
		return true
	}

	return false
}

// SetComponentCopyrightScopeOwnerId gets a reference to the given string and assigns it to the ComponentCopyrightScopeOwnerId field.
func (o *ApiLicenseLegalDataDTO) SetComponentCopyrightScopeOwnerId(v string) {
	o.ComponentCopyrightScopeOwnerId = &v
}

// GetComponentLicensesId returns the ComponentLicensesId field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetComponentLicensesId() string {
	if o == nil || IsNil(o.ComponentLicensesId) {
		var ret string
		return ret
	}
	return *o.ComponentLicensesId
}

// GetComponentLicensesIdOk returns a tuple with the ComponentLicensesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetComponentLicensesIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentLicensesId) {
		return nil, false
	}
	return o.ComponentLicensesId, true
}

// HasComponentLicensesId returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasComponentLicensesId() bool {
	if o != nil && !IsNil(o.ComponentLicensesId) {
		return true
	}

	return false
}

// SetComponentLicensesId gets a reference to the given string and assigns it to the ComponentLicensesId field.
func (o *ApiLicenseLegalDataDTO) SetComponentLicensesId(v string) {
	o.ComponentLicensesId = &v
}

// GetComponentLicensesLastUpdatedAt returns the ComponentLicensesLastUpdatedAt field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetComponentLicensesLastUpdatedAt() time.Time {
	if o == nil || IsNil(o.ComponentLicensesLastUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.ComponentLicensesLastUpdatedAt
}

// GetComponentLicensesLastUpdatedAtOk returns a tuple with the ComponentLicensesLastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetComponentLicensesLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ComponentLicensesLastUpdatedAt) {
		return nil, false
	}
	return o.ComponentLicensesLastUpdatedAt, true
}

// HasComponentLicensesLastUpdatedAt returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasComponentLicensesLastUpdatedAt() bool {
	if o != nil && !IsNil(o.ComponentLicensesLastUpdatedAt) {
		return true
	}

	return false
}

// SetComponentLicensesLastUpdatedAt gets a reference to the given time.Time and assigns it to the ComponentLicensesLastUpdatedAt field.
func (o *ApiLicenseLegalDataDTO) SetComponentLicensesLastUpdatedAt(v time.Time) {
	o.ComponentLicensesLastUpdatedAt = &v
}

// GetComponentLicensesLastUpdatedByUsername returns the ComponentLicensesLastUpdatedByUsername field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetComponentLicensesLastUpdatedByUsername() string {
	if o == nil || IsNil(o.ComponentLicensesLastUpdatedByUsername) {
		var ret string
		return ret
	}
	return *o.ComponentLicensesLastUpdatedByUsername
}

// GetComponentLicensesLastUpdatedByUsernameOk returns a tuple with the ComponentLicensesLastUpdatedByUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetComponentLicensesLastUpdatedByUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentLicensesLastUpdatedByUsername) {
		return nil, false
	}
	return o.ComponentLicensesLastUpdatedByUsername, true
}

// HasComponentLicensesLastUpdatedByUsername returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasComponentLicensesLastUpdatedByUsername() bool {
	if o != nil && !IsNil(o.ComponentLicensesLastUpdatedByUsername) {
		return true
	}

	return false
}

// SetComponentLicensesLastUpdatedByUsername gets a reference to the given string and assigns it to the ComponentLicensesLastUpdatedByUsername field.
func (o *ApiLicenseLegalDataDTO) SetComponentLicensesLastUpdatedByUsername(v string) {
	o.ComponentLicensesLastUpdatedByUsername = &v
}

// GetComponentLicensesScopeOwnerId returns the ComponentLicensesScopeOwnerId field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetComponentLicensesScopeOwnerId() string {
	if o == nil || IsNil(o.ComponentLicensesScopeOwnerId) {
		var ret string
		return ret
	}
	return *o.ComponentLicensesScopeOwnerId
}

// GetComponentLicensesScopeOwnerIdOk returns a tuple with the ComponentLicensesScopeOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetComponentLicensesScopeOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentLicensesScopeOwnerId) {
		return nil, false
	}
	return o.ComponentLicensesScopeOwnerId, true
}

// HasComponentLicensesScopeOwnerId returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasComponentLicensesScopeOwnerId() bool {
	if o != nil && !IsNil(o.ComponentLicensesScopeOwnerId) {
		return true
	}

	return false
}

// SetComponentLicensesScopeOwnerId gets a reference to the given string and assigns it to the ComponentLicensesScopeOwnerId field.
func (o *ApiLicenseLegalDataDTO) SetComponentLicensesScopeOwnerId(v string) {
	o.ComponentLicensesScopeOwnerId = &v
}

// GetComponentNoticesId returns the ComponentNoticesId field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetComponentNoticesId() string {
	if o == nil || IsNil(o.ComponentNoticesId) {
		var ret string
		return ret
	}
	return *o.ComponentNoticesId
}

// GetComponentNoticesIdOk returns a tuple with the ComponentNoticesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetComponentNoticesIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentNoticesId) {
		return nil, false
	}
	return o.ComponentNoticesId, true
}

// HasComponentNoticesId returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasComponentNoticesId() bool {
	if o != nil && !IsNil(o.ComponentNoticesId) {
		return true
	}

	return false
}

// SetComponentNoticesId gets a reference to the given string and assigns it to the ComponentNoticesId field.
func (o *ApiLicenseLegalDataDTO) SetComponentNoticesId(v string) {
	o.ComponentNoticesId = &v
}

// GetComponentNoticesLastUpdatedAt returns the ComponentNoticesLastUpdatedAt field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetComponentNoticesLastUpdatedAt() time.Time {
	if o == nil || IsNil(o.ComponentNoticesLastUpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.ComponentNoticesLastUpdatedAt
}

// GetComponentNoticesLastUpdatedAtOk returns a tuple with the ComponentNoticesLastUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetComponentNoticesLastUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ComponentNoticesLastUpdatedAt) {
		return nil, false
	}
	return o.ComponentNoticesLastUpdatedAt, true
}

// HasComponentNoticesLastUpdatedAt returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasComponentNoticesLastUpdatedAt() bool {
	if o != nil && !IsNil(o.ComponentNoticesLastUpdatedAt) {
		return true
	}

	return false
}

// SetComponentNoticesLastUpdatedAt gets a reference to the given time.Time and assigns it to the ComponentNoticesLastUpdatedAt field.
func (o *ApiLicenseLegalDataDTO) SetComponentNoticesLastUpdatedAt(v time.Time) {
	o.ComponentNoticesLastUpdatedAt = &v
}

// GetComponentNoticesLastUpdatedByUsername returns the ComponentNoticesLastUpdatedByUsername field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetComponentNoticesLastUpdatedByUsername() string {
	if o == nil || IsNil(o.ComponentNoticesLastUpdatedByUsername) {
		var ret string
		return ret
	}
	return *o.ComponentNoticesLastUpdatedByUsername
}

// GetComponentNoticesLastUpdatedByUsernameOk returns a tuple with the ComponentNoticesLastUpdatedByUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetComponentNoticesLastUpdatedByUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentNoticesLastUpdatedByUsername) {
		return nil, false
	}
	return o.ComponentNoticesLastUpdatedByUsername, true
}

// HasComponentNoticesLastUpdatedByUsername returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasComponentNoticesLastUpdatedByUsername() bool {
	if o != nil && !IsNil(o.ComponentNoticesLastUpdatedByUsername) {
		return true
	}

	return false
}

// SetComponentNoticesLastUpdatedByUsername gets a reference to the given string and assigns it to the ComponentNoticesLastUpdatedByUsername field.
func (o *ApiLicenseLegalDataDTO) SetComponentNoticesLastUpdatedByUsername(v string) {
	o.ComponentNoticesLastUpdatedByUsername = &v
}

// GetComponentNoticesScopeOwnerId returns the ComponentNoticesScopeOwnerId field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetComponentNoticesScopeOwnerId() string {
	if o == nil || IsNil(o.ComponentNoticesScopeOwnerId) {
		var ret string
		return ret
	}
	return *o.ComponentNoticesScopeOwnerId
}

// GetComponentNoticesScopeOwnerIdOk returns a tuple with the ComponentNoticesScopeOwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetComponentNoticesScopeOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentNoticesScopeOwnerId) {
		return nil, false
	}
	return o.ComponentNoticesScopeOwnerId, true
}

// HasComponentNoticesScopeOwnerId returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasComponentNoticesScopeOwnerId() bool {
	if o != nil && !IsNil(o.ComponentNoticesScopeOwnerId) {
		return true
	}

	return false
}

// SetComponentNoticesScopeOwnerId gets a reference to the given string and assigns it to the ComponentNoticesScopeOwnerId field.
func (o *ApiLicenseLegalDataDTO) SetComponentNoticesScopeOwnerId(v string) {
	o.ComponentNoticesScopeOwnerId = &v
}

// GetCopyrights returns the Copyrights field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetCopyrights() []ApiLicenseLegalCopyrightDTO {
	if o == nil || IsNil(o.Copyrights) {
		var ret []ApiLicenseLegalCopyrightDTO
		return ret
	}
	return o.Copyrights
}

// GetCopyrightsOk returns a tuple with the Copyrights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetCopyrightsOk() ([]ApiLicenseLegalCopyrightDTO, bool) {
	if o == nil || IsNil(o.Copyrights) {
		return nil, false
	}
	return o.Copyrights, true
}

// HasCopyrights returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasCopyrights() bool {
	if o != nil && !IsNil(o.Copyrights) {
		return true
	}

	return false
}

// SetCopyrights gets a reference to the given []ApiLicenseLegalCopyrightDTO and assigns it to the Copyrights field.
func (o *ApiLicenseLegalDataDTO) SetCopyrights(v []ApiLicenseLegalCopyrightDTO) {
	o.Copyrights = v
}

// GetDeclaredLicenses returns the DeclaredLicenses field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetDeclaredLicenses() []string {
	if o == nil || IsNil(o.DeclaredLicenses) {
		var ret []string
		return ret
	}
	return o.DeclaredLicenses
}

// GetDeclaredLicensesOk returns a tuple with the DeclaredLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetDeclaredLicensesOk() ([]string, bool) {
	if o == nil || IsNil(o.DeclaredLicenses) {
		return nil, false
	}
	return o.DeclaredLicenses, true
}

// HasDeclaredLicenses returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasDeclaredLicenses() bool {
	if o != nil && !IsNil(o.DeclaredLicenses) {
		return true
	}

	return false
}

// SetDeclaredLicenses gets a reference to the given []string and assigns it to the DeclaredLicenses field.
func (o *ApiLicenseLegalDataDTO) SetDeclaredLicenses(v []string) {
	o.DeclaredLicenses = v
}

// GetEffectiveLicenseStatus returns the EffectiveLicenseStatus field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetEffectiveLicenseStatus() string {
	if o == nil || IsNil(o.EffectiveLicenseStatus) {
		var ret string
		return ret
	}
	return *o.EffectiveLicenseStatus
}

// GetEffectiveLicenseStatusOk returns a tuple with the EffectiveLicenseStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetEffectiveLicenseStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveLicenseStatus) {
		return nil, false
	}
	return o.EffectiveLicenseStatus, true
}

// HasEffectiveLicenseStatus returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasEffectiveLicenseStatus() bool {
	if o != nil && !IsNil(o.EffectiveLicenseStatus) {
		return true
	}

	return false
}

// SetEffectiveLicenseStatus gets a reference to the given string and assigns it to the EffectiveLicenseStatus field.
func (o *ApiLicenseLegalDataDTO) SetEffectiveLicenseStatus(v string) {
	o.EffectiveLicenseStatus = &v
}

// GetEffectiveLicenses returns the EffectiveLicenses field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetEffectiveLicenses() []string {
	if o == nil || IsNil(o.EffectiveLicenses) {
		var ret []string
		return ret
	}
	return o.EffectiveLicenses
}

// GetEffectiveLicensesOk returns a tuple with the EffectiveLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetEffectiveLicensesOk() ([]string, bool) {
	if o == nil || IsNil(o.EffectiveLicenses) {
		return nil, false
	}
	return o.EffectiveLicenses, true
}

// HasEffectiveLicenses returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasEffectiveLicenses() bool {
	if o != nil && !IsNil(o.EffectiveLicenses) {
		return true
	}

	return false
}

// SetEffectiveLicenses gets a reference to the given []string and assigns it to the EffectiveLicenses field.
func (o *ApiLicenseLegalDataDTO) SetEffectiveLicenses(v []string) {
	o.EffectiveLicenses = v
}

// GetHighestEffectiveLicenseThreatGroup returns the HighestEffectiveLicenseThreatGroup field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetHighestEffectiveLicenseThreatGroup() ApiLicenseThreatDTOV2 {
	if o == nil || IsNil(o.HighestEffectiveLicenseThreatGroup) {
		var ret ApiLicenseThreatDTOV2
		return ret
	}
	return *o.HighestEffectiveLicenseThreatGroup
}

// GetHighestEffectiveLicenseThreatGroupOk returns a tuple with the HighestEffectiveLicenseThreatGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetHighestEffectiveLicenseThreatGroupOk() (*ApiLicenseThreatDTOV2, bool) {
	if o == nil || IsNil(o.HighestEffectiveLicenseThreatGroup) {
		return nil, false
	}
	return o.HighestEffectiveLicenseThreatGroup, true
}

// HasHighestEffectiveLicenseThreatGroup returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasHighestEffectiveLicenseThreatGroup() bool {
	if o != nil && !IsNil(o.HighestEffectiveLicenseThreatGroup) {
		return true
	}

	return false
}

// SetHighestEffectiveLicenseThreatGroup gets a reference to the given ApiLicenseThreatDTOV2 and assigns it to the HighestEffectiveLicenseThreatGroup field.
func (o *ApiLicenseLegalDataDTO) SetHighestEffectiveLicenseThreatGroup(v ApiLicenseThreatDTOV2) {
	o.HighestEffectiveLicenseThreatGroup = &v
}

// GetLicenseFiles returns the LicenseFiles field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetLicenseFiles() []ApiLicenseLegalFileDTO {
	if o == nil || IsNil(o.LicenseFiles) {
		var ret []ApiLicenseLegalFileDTO
		return ret
	}
	return o.LicenseFiles
}

// GetLicenseFilesOk returns a tuple with the LicenseFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetLicenseFilesOk() ([]ApiLicenseLegalFileDTO, bool) {
	if o == nil || IsNil(o.LicenseFiles) {
		return nil, false
	}
	return o.LicenseFiles, true
}

// HasLicenseFiles returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasLicenseFiles() bool {
	if o != nil && !IsNil(o.LicenseFiles) {
		return true
	}

	return false
}

// SetLicenseFiles gets a reference to the given []ApiLicenseLegalFileDTO and assigns it to the LicenseFiles field.
func (o *ApiLicenseLegalDataDTO) SetLicenseFiles(v []ApiLicenseLegalFileDTO) {
	o.LicenseFiles = v
}

// GetNoticeFiles returns the NoticeFiles field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetNoticeFiles() []ApiLicenseLegalFileDTO {
	if o == nil || IsNil(o.NoticeFiles) {
		var ret []ApiLicenseLegalFileDTO
		return ret
	}
	return o.NoticeFiles
}

// GetNoticeFilesOk returns a tuple with the NoticeFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetNoticeFilesOk() ([]ApiLicenseLegalFileDTO, bool) {
	if o == nil || IsNil(o.NoticeFiles) {
		return nil, false
	}
	return o.NoticeFiles, true
}

// HasNoticeFiles returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasNoticeFiles() bool {
	if o != nil && !IsNil(o.NoticeFiles) {
		return true
	}

	return false
}

// SetNoticeFiles gets a reference to the given []ApiLicenseLegalFileDTO and assigns it to the NoticeFiles field.
func (o *ApiLicenseLegalDataDTO) SetNoticeFiles(v []ApiLicenseLegalFileDTO) {
	o.NoticeFiles = v
}

// GetObligations returns the Obligations field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetObligations() []ApiLicenseLegalObligationDTO {
	if o == nil || IsNil(o.Obligations) {
		var ret []ApiLicenseLegalObligationDTO
		return ret
	}
	return o.Obligations
}

// GetObligationsOk returns a tuple with the Obligations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetObligationsOk() ([]ApiLicenseLegalObligationDTO, bool) {
	if o == nil || IsNil(o.Obligations) {
		return nil, false
	}
	return o.Obligations, true
}

// HasObligations returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasObligations() bool {
	if o != nil && !IsNil(o.Obligations) {
		return true
	}

	return false
}

// SetObligations gets a reference to the given []ApiLicenseLegalObligationDTO and assigns it to the Obligations field.
func (o *ApiLicenseLegalDataDTO) SetObligations(v []ApiLicenseLegalObligationDTO) {
	o.Obligations = v
}

// GetObservedLicenses returns the ObservedLicenses field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetObservedLicenses() []string {
	if o == nil || IsNil(o.ObservedLicenses) {
		var ret []string
		return ret
	}
	return o.ObservedLicenses
}

// GetObservedLicensesOk returns a tuple with the ObservedLicenses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetObservedLicensesOk() ([]string, bool) {
	if o == nil || IsNil(o.ObservedLicenses) {
		return nil, false
	}
	return o.ObservedLicenses, true
}

// HasObservedLicenses returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasObservedLicenses() bool {
	if o != nil && !IsNil(o.ObservedLicenses) {
		return true
	}

	return false
}

// SetObservedLicenses gets a reference to the given []string and assigns it to the ObservedLicenses field.
func (o *ApiLicenseLegalDataDTO) SetObservedLicenses(v []string) {
	o.ObservedLicenses = v
}

// GetSourceLinks returns the SourceLinks field value if set, zero value otherwise.
func (o *ApiLicenseLegalDataDTO) GetSourceLinks() []LegalSourceLinkDTO {
	if o == nil || IsNil(o.SourceLinks) {
		var ret []LegalSourceLinkDTO
		return ret
	}
	return o.SourceLinks
}

// GetSourceLinksOk returns a tuple with the SourceLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiLicenseLegalDataDTO) GetSourceLinksOk() ([]LegalSourceLinkDTO, bool) {
	if o == nil || IsNil(o.SourceLinks) {
		return nil, false
	}
	return o.SourceLinks, true
}

// HasSourceLinks returns a boolean if a field has been set.
func (o *ApiLicenseLegalDataDTO) HasSourceLinks() bool {
	if o != nil && !IsNil(o.SourceLinks) {
		return true
	}

	return false
}

// SetSourceLinks gets a reference to the given []LegalSourceLinkDTO and assigns it to the SourceLinks field.
func (o *ApiLicenseLegalDataDTO) SetSourceLinks(v []LegalSourceLinkDTO) {
	o.SourceLinks = v
}

func (o ApiLicenseLegalDataDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiLicenseLegalDataDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributions) {
		toSerialize["attributions"] = o.Attributions
	}
	if !IsNil(o.ComponentCopyrightId) {
		toSerialize["componentCopyrightId"] = o.ComponentCopyrightId
	}
	if !IsNil(o.ComponentCopyrightLastUpdatedAt) {
		toSerialize["componentCopyrightLastUpdatedAt"] = o.ComponentCopyrightLastUpdatedAt
	}
	if !IsNil(o.ComponentCopyrightLastUpdatedByUsername) {
		toSerialize["componentCopyrightLastUpdatedByUsername"] = o.ComponentCopyrightLastUpdatedByUsername
	}
	if !IsNil(o.ComponentCopyrightScopeOwnerId) {
		toSerialize["componentCopyrightScopeOwnerId"] = o.ComponentCopyrightScopeOwnerId
	}
	if !IsNil(o.ComponentLicensesId) {
		toSerialize["componentLicensesId"] = o.ComponentLicensesId
	}
	if !IsNil(o.ComponentLicensesLastUpdatedAt) {
		toSerialize["componentLicensesLastUpdatedAt"] = o.ComponentLicensesLastUpdatedAt
	}
	if !IsNil(o.ComponentLicensesLastUpdatedByUsername) {
		toSerialize["componentLicensesLastUpdatedByUsername"] = o.ComponentLicensesLastUpdatedByUsername
	}
	if !IsNil(o.ComponentLicensesScopeOwnerId) {
		toSerialize["componentLicensesScopeOwnerId"] = o.ComponentLicensesScopeOwnerId
	}
	if !IsNil(o.ComponentNoticesId) {
		toSerialize["componentNoticesId"] = o.ComponentNoticesId
	}
	if !IsNil(o.ComponentNoticesLastUpdatedAt) {
		toSerialize["componentNoticesLastUpdatedAt"] = o.ComponentNoticesLastUpdatedAt
	}
	if !IsNil(o.ComponentNoticesLastUpdatedByUsername) {
		toSerialize["componentNoticesLastUpdatedByUsername"] = o.ComponentNoticesLastUpdatedByUsername
	}
	if !IsNil(o.ComponentNoticesScopeOwnerId) {
		toSerialize["componentNoticesScopeOwnerId"] = o.ComponentNoticesScopeOwnerId
	}
	if !IsNil(o.Copyrights) {
		toSerialize["copyrights"] = o.Copyrights
	}
	if !IsNil(o.DeclaredLicenses) {
		toSerialize["declaredLicenses"] = o.DeclaredLicenses
	}
	if !IsNil(o.EffectiveLicenseStatus) {
		toSerialize["effectiveLicenseStatus"] = o.EffectiveLicenseStatus
	}
	if !IsNil(o.EffectiveLicenses) {
		toSerialize["effectiveLicenses"] = o.EffectiveLicenses
	}
	if !IsNil(o.HighestEffectiveLicenseThreatGroup) {
		toSerialize["highestEffectiveLicenseThreatGroup"] = o.HighestEffectiveLicenseThreatGroup
	}
	if !IsNil(o.LicenseFiles) {
		toSerialize["licenseFiles"] = o.LicenseFiles
	}
	if !IsNil(o.NoticeFiles) {
		toSerialize["noticeFiles"] = o.NoticeFiles
	}
	if !IsNil(o.Obligations) {
		toSerialize["obligations"] = o.Obligations
	}
	if !IsNil(o.ObservedLicenses) {
		toSerialize["observedLicenses"] = o.ObservedLicenses
	}
	if !IsNil(o.SourceLinks) {
		toSerialize["sourceLinks"] = o.SourceLinks
	}
	return toSerialize, nil
}

type NullableApiLicenseLegalDataDTO struct {
	value *ApiLicenseLegalDataDTO
	isSet bool
}

func (v NullableApiLicenseLegalDataDTO) Get() *ApiLicenseLegalDataDTO {
	return v.value
}

func (v *NullableApiLicenseLegalDataDTO) Set(val *ApiLicenseLegalDataDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableApiLicenseLegalDataDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableApiLicenseLegalDataDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiLicenseLegalDataDTO(val *ApiLicenseLegalDataDTO) *NullableApiLicenseLegalDataDTO {
	return &NullableApiLicenseLegalDataDTO{value: val, isSet: true}
}

func (v NullableApiLicenseLegalDataDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiLicenseLegalDataDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

