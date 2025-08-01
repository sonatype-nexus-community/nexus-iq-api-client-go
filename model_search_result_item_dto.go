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

// checks if the SearchResultItemDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchResultItemDTO{}

// SearchResultItemDTO struct for SearchResultItemDTO
type SearchResultItemDTO struct {
	ApplicationCategoryColor *string `json:"applicationCategoryColor,omitempty"`
	ApplicationCategoryDescription *string `json:"applicationCategoryDescription,omitempty"`
	ApplicationCategoryId *string `json:"applicationCategoryId,omitempty"`
	ApplicationCategoryName *string `json:"applicationCategoryName,omitempty"`
	ApplicationId *string `json:"applicationId,omitempty"`
	ApplicationName *string `json:"applicationName,omitempty"`
	ApplicationPublicId *string `json:"applicationPublicId,omitempty"`
	ApplicationVersion *string `json:"applicationVersion,omitempty"`
	ComponentHash *string `json:"componentHash,omitempty"`
	ComponentIdentifier *ApiComponentIdentifierDTOV2 `json:"componentIdentifier,omitempty"`
	ComponentLabelColor *string `json:"componentLabelColor,omitempty"`
	ComponentLabelDescription *string `json:"componentLabelDescription,omitempty"`
	ComponentLabelId *string `json:"componentLabelId,omitempty"`
	ComponentLabelName *string `json:"componentLabelName,omitempty"`
	ComponentName *string `json:"componentName,omitempty"`
	ItemType *string `json:"itemType,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
	OrganizationName *string `json:"organizationName,omitempty"`
	PolicyEvaluationStage *string `json:"policyEvaluationStage,omitempty"`
	PolicyId *string `json:"policyId,omitempty"`
	PolicyName *string `json:"policyName,omitempty"`
	PolicyThreatCategory *string `json:"policyThreatCategory,omitempty"`
	PolicyThreatLevel *int32 `json:"policyThreatLevel,omitempty"`
	ReportId *string `json:"reportId,omitempty"`
	ResultIndex *int32 `json:"resultIndex,omitempty"`
	SbomSpecification *string `json:"sbomSpecification,omitempty"`
	VulnerabilityDescription *string `json:"vulnerabilityDescription,omitempty"`
	VulnerabilityId *string `json:"vulnerabilityId,omitempty"`
	VulnerabilityStatus *string `json:"vulnerabilityStatus,omitempty"`
}

// NewSearchResultItemDTO instantiates a new SearchResultItemDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResultItemDTO() *SearchResultItemDTO {
	this := SearchResultItemDTO{}
	return &this
}

// NewSearchResultItemDTOWithDefaults instantiates a new SearchResultItemDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultItemDTOWithDefaults() *SearchResultItemDTO {
	this := SearchResultItemDTO{}
	return &this
}

// GetApplicationCategoryColor returns the ApplicationCategoryColor field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetApplicationCategoryColor() string {
	if o == nil || IsNil(o.ApplicationCategoryColor) {
		var ret string
		return ret
	}
	return *o.ApplicationCategoryColor
}

// GetApplicationCategoryColorOk returns a tuple with the ApplicationCategoryColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetApplicationCategoryColorOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationCategoryColor) {
		return nil, false
	}
	return o.ApplicationCategoryColor, true
}

// HasApplicationCategoryColor returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasApplicationCategoryColor() bool {
	if o != nil && !IsNil(o.ApplicationCategoryColor) {
		return true
	}

	return false
}

// SetApplicationCategoryColor gets a reference to the given string and assigns it to the ApplicationCategoryColor field.
func (o *SearchResultItemDTO) SetApplicationCategoryColor(v string) {
	o.ApplicationCategoryColor = &v
}

// GetApplicationCategoryDescription returns the ApplicationCategoryDescription field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetApplicationCategoryDescription() string {
	if o == nil || IsNil(o.ApplicationCategoryDescription) {
		var ret string
		return ret
	}
	return *o.ApplicationCategoryDescription
}

// GetApplicationCategoryDescriptionOk returns a tuple with the ApplicationCategoryDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetApplicationCategoryDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationCategoryDescription) {
		return nil, false
	}
	return o.ApplicationCategoryDescription, true
}

// HasApplicationCategoryDescription returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasApplicationCategoryDescription() bool {
	if o != nil && !IsNil(o.ApplicationCategoryDescription) {
		return true
	}

	return false
}

// SetApplicationCategoryDescription gets a reference to the given string and assigns it to the ApplicationCategoryDescription field.
func (o *SearchResultItemDTO) SetApplicationCategoryDescription(v string) {
	o.ApplicationCategoryDescription = &v
}

// GetApplicationCategoryId returns the ApplicationCategoryId field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetApplicationCategoryId() string {
	if o == nil || IsNil(o.ApplicationCategoryId) {
		var ret string
		return ret
	}
	return *o.ApplicationCategoryId
}

// GetApplicationCategoryIdOk returns a tuple with the ApplicationCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetApplicationCategoryIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationCategoryId) {
		return nil, false
	}
	return o.ApplicationCategoryId, true
}

// HasApplicationCategoryId returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasApplicationCategoryId() bool {
	if o != nil && !IsNil(o.ApplicationCategoryId) {
		return true
	}

	return false
}

// SetApplicationCategoryId gets a reference to the given string and assigns it to the ApplicationCategoryId field.
func (o *SearchResultItemDTO) SetApplicationCategoryId(v string) {
	o.ApplicationCategoryId = &v
}

// GetApplicationCategoryName returns the ApplicationCategoryName field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetApplicationCategoryName() string {
	if o == nil || IsNil(o.ApplicationCategoryName) {
		var ret string
		return ret
	}
	return *o.ApplicationCategoryName
}

// GetApplicationCategoryNameOk returns a tuple with the ApplicationCategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetApplicationCategoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationCategoryName) {
		return nil, false
	}
	return o.ApplicationCategoryName, true
}

// HasApplicationCategoryName returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasApplicationCategoryName() bool {
	if o != nil && !IsNil(o.ApplicationCategoryName) {
		return true
	}

	return false
}

// SetApplicationCategoryName gets a reference to the given string and assigns it to the ApplicationCategoryName field.
func (o *SearchResultItemDTO) SetApplicationCategoryName(v string) {
	o.ApplicationCategoryName = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *SearchResultItemDTO) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetApplicationName() string {
	if o == nil || IsNil(o.ApplicationName) {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetApplicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationName) {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasApplicationName() bool {
	if o != nil && !IsNil(o.ApplicationName) {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *SearchResultItemDTO) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetApplicationPublicId returns the ApplicationPublicId field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetApplicationPublicId() string {
	if o == nil || IsNil(o.ApplicationPublicId) {
		var ret string
		return ret
	}
	return *o.ApplicationPublicId
}

// GetApplicationPublicIdOk returns a tuple with the ApplicationPublicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetApplicationPublicIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationPublicId) {
		return nil, false
	}
	return o.ApplicationPublicId, true
}

// HasApplicationPublicId returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasApplicationPublicId() bool {
	if o != nil && !IsNil(o.ApplicationPublicId) {
		return true
	}

	return false
}

// SetApplicationPublicId gets a reference to the given string and assigns it to the ApplicationPublicId field.
func (o *SearchResultItemDTO) SetApplicationPublicId(v string) {
	o.ApplicationPublicId = &v
}

// GetApplicationVersion returns the ApplicationVersion field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetApplicationVersion() string {
	if o == nil || IsNil(o.ApplicationVersion) {
		var ret string
		return ret
	}
	return *o.ApplicationVersion
}

// GetApplicationVersionOk returns a tuple with the ApplicationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetApplicationVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationVersion) {
		return nil, false
	}
	return o.ApplicationVersion, true
}

// HasApplicationVersion returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasApplicationVersion() bool {
	if o != nil && !IsNil(o.ApplicationVersion) {
		return true
	}

	return false
}

// SetApplicationVersion gets a reference to the given string and assigns it to the ApplicationVersion field.
func (o *SearchResultItemDTO) SetApplicationVersion(v string) {
	o.ApplicationVersion = &v
}

// GetComponentHash returns the ComponentHash field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetComponentHash() string {
	if o == nil || IsNil(o.ComponentHash) {
		var ret string
		return ret
	}
	return *o.ComponentHash
}

// GetComponentHashOk returns a tuple with the ComponentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetComponentHashOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentHash) {
		return nil, false
	}
	return o.ComponentHash, true
}

// HasComponentHash returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasComponentHash() bool {
	if o != nil && !IsNil(o.ComponentHash) {
		return true
	}

	return false
}

// SetComponentHash gets a reference to the given string and assigns it to the ComponentHash field.
func (o *SearchResultItemDTO) SetComponentHash(v string) {
	o.ComponentHash = &v
}

// GetComponentIdentifier returns the ComponentIdentifier field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2 {
	if o == nil || IsNil(o.ComponentIdentifier) {
		var ret ApiComponentIdentifierDTOV2
		return ret
	}
	return *o.ComponentIdentifier
}

// GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool) {
	if o == nil || IsNil(o.ComponentIdentifier) {
		return nil, false
	}
	return o.ComponentIdentifier, true
}

// HasComponentIdentifier returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasComponentIdentifier() bool {
	if o != nil && !IsNil(o.ComponentIdentifier) {
		return true
	}

	return false
}

// SetComponentIdentifier gets a reference to the given ApiComponentIdentifierDTOV2 and assigns it to the ComponentIdentifier field.
func (o *SearchResultItemDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2) {
	o.ComponentIdentifier = &v
}

// GetComponentLabelColor returns the ComponentLabelColor field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetComponentLabelColor() string {
	if o == nil || IsNil(o.ComponentLabelColor) {
		var ret string
		return ret
	}
	return *o.ComponentLabelColor
}

// GetComponentLabelColorOk returns a tuple with the ComponentLabelColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetComponentLabelColorOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentLabelColor) {
		return nil, false
	}
	return o.ComponentLabelColor, true
}

// HasComponentLabelColor returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasComponentLabelColor() bool {
	if o != nil && !IsNil(o.ComponentLabelColor) {
		return true
	}

	return false
}

// SetComponentLabelColor gets a reference to the given string and assigns it to the ComponentLabelColor field.
func (o *SearchResultItemDTO) SetComponentLabelColor(v string) {
	o.ComponentLabelColor = &v
}

// GetComponentLabelDescription returns the ComponentLabelDescription field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetComponentLabelDescription() string {
	if o == nil || IsNil(o.ComponentLabelDescription) {
		var ret string
		return ret
	}
	return *o.ComponentLabelDescription
}

// GetComponentLabelDescriptionOk returns a tuple with the ComponentLabelDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetComponentLabelDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentLabelDescription) {
		return nil, false
	}
	return o.ComponentLabelDescription, true
}

// HasComponentLabelDescription returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasComponentLabelDescription() bool {
	if o != nil && !IsNil(o.ComponentLabelDescription) {
		return true
	}

	return false
}

// SetComponentLabelDescription gets a reference to the given string and assigns it to the ComponentLabelDescription field.
func (o *SearchResultItemDTO) SetComponentLabelDescription(v string) {
	o.ComponentLabelDescription = &v
}

// GetComponentLabelId returns the ComponentLabelId field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetComponentLabelId() string {
	if o == nil || IsNil(o.ComponentLabelId) {
		var ret string
		return ret
	}
	return *o.ComponentLabelId
}

// GetComponentLabelIdOk returns a tuple with the ComponentLabelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetComponentLabelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentLabelId) {
		return nil, false
	}
	return o.ComponentLabelId, true
}

// HasComponentLabelId returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasComponentLabelId() bool {
	if o != nil && !IsNil(o.ComponentLabelId) {
		return true
	}

	return false
}

// SetComponentLabelId gets a reference to the given string and assigns it to the ComponentLabelId field.
func (o *SearchResultItemDTO) SetComponentLabelId(v string) {
	o.ComponentLabelId = &v
}

// GetComponentLabelName returns the ComponentLabelName field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetComponentLabelName() string {
	if o == nil || IsNil(o.ComponentLabelName) {
		var ret string
		return ret
	}
	return *o.ComponentLabelName
}

// GetComponentLabelNameOk returns a tuple with the ComponentLabelName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetComponentLabelNameOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentLabelName) {
		return nil, false
	}
	return o.ComponentLabelName, true
}

// HasComponentLabelName returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasComponentLabelName() bool {
	if o != nil && !IsNil(o.ComponentLabelName) {
		return true
	}

	return false
}

// SetComponentLabelName gets a reference to the given string and assigns it to the ComponentLabelName field.
func (o *SearchResultItemDTO) SetComponentLabelName(v string) {
	o.ComponentLabelName = &v
}

// GetComponentName returns the ComponentName field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetComponentName() string {
	if o == nil || IsNil(o.ComponentName) {
		var ret string
		return ret
	}
	return *o.ComponentName
}

// GetComponentNameOk returns a tuple with the ComponentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetComponentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentName) {
		return nil, false
	}
	return o.ComponentName, true
}

// HasComponentName returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasComponentName() bool {
	if o != nil && !IsNil(o.ComponentName) {
		return true
	}

	return false
}

// SetComponentName gets a reference to the given string and assigns it to the ComponentName field.
func (o *SearchResultItemDTO) SetComponentName(v string) {
	o.ComponentName = &v
}

// GetItemType returns the ItemType field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetItemType() string {
	if o == nil || IsNil(o.ItemType) {
		var ret string
		return ret
	}
	return *o.ItemType
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetItemTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ItemType) {
		return nil, false
	}
	return o.ItemType, true
}

// HasItemType returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasItemType() bool {
	if o != nil && !IsNil(o.ItemType) {
		return true
	}

	return false
}

// SetItemType gets a reference to the given string and assigns it to the ItemType field.
func (o *SearchResultItemDTO) SetItemType(v string) {
	o.ItemType = &v
}

// GetOrganizationId returns the OrganizationId field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetOrganizationId() string {
	if o == nil || IsNil(o.OrganizationId) {
		var ret string
		return ret
	}
	return *o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetOrganizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationId) {
		return nil, false
	}
	return o.OrganizationId, true
}

// HasOrganizationId returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasOrganizationId() bool {
	if o != nil && !IsNil(o.OrganizationId) {
		return true
	}

	return false
}

// SetOrganizationId gets a reference to the given string and assigns it to the OrganizationId field.
func (o *SearchResultItemDTO) SetOrganizationId(v string) {
	o.OrganizationId = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetOrganizationName() string {
	if o == nil || IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetOrganizationNameOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasOrganizationName() bool {
	if o != nil && !IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *SearchResultItemDTO) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetPolicyEvaluationStage returns the PolicyEvaluationStage field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetPolicyEvaluationStage() string {
	if o == nil || IsNil(o.PolicyEvaluationStage) {
		var ret string
		return ret
	}
	return *o.PolicyEvaluationStage
}

// GetPolicyEvaluationStageOk returns a tuple with the PolicyEvaluationStage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetPolicyEvaluationStageOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyEvaluationStage) {
		return nil, false
	}
	return o.PolicyEvaluationStage, true
}

// HasPolicyEvaluationStage returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasPolicyEvaluationStage() bool {
	if o != nil && !IsNil(o.PolicyEvaluationStage) {
		return true
	}

	return false
}

// SetPolicyEvaluationStage gets a reference to the given string and assigns it to the PolicyEvaluationStage field.
func (o *SearchResultItemDTO) SetPolicyEvaluationStage(v string) {
	o.PolicyEvaluationStage = &v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetPolicyId() string {
	if o == nil || IsNil(o.PolicyId) {
		var ret string
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetPolicyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasPolicyId() bool {
	if o != nil && !IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given string and assigns it to the PolicyId field.
func (o *SearchResultItemDTO) SetPolicyId(v string) {
	o.PolicyId = &v
}

// GetPolicyName returns the PolicyName field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetPolicyName() string {
	if o == nil || IsNil(o.PolicyName) {
		var ret string
		return ret
	}
	return *o.PolicyName
}

// GetPolicyNameOk returns a tuple with the PolicyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetPolicyNameOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyName) {
		return nil, false
	}
	return o.PolicyName, true
}

// HasPolicyName returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasPolicyName() bool {
	if o != nil && !IsNil(o.PolicyName) {
		return true
	}

	return false
}

// SetPolicyName gets a reference to the given string and assigns it to the PolicyName field.
func (o *SearchResultItemDTO) SetPolicyName(v string) {
	o.PolicyName = &v
}

// GetPolicyThreatCategory returns the PolicyThreatCategory field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetPolicyThreatCategory() string {
	if o == nil || IsNil(o.PolicyThreatCategory) {
		var ret string
		return ret
	}
	return *o.PolicyThreatCategory
}

// GetPolicyThreatCategoryOk returns a tuple with the PolicyThreatCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetPolicyThreatCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyThreatCategory) {
		return nil, false
	}
	return o.PolicyThreatCategory, true
}

// HasPolicyThreatCategory returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasPolicyThreatCategory() bool {
	if o != nil && !IsNil(o.PolicyThreatCategory) {
		return true
	}

	return false
}

// SetPolicyThreatCategory gets a reference to the given string and assigns it to the PolicyThreatCategory field.
func (o *SearchResultItemDTO) SetPolicyThreatCategory(v string) {
	o.PolicyThreatCategory = &v
}

// GetPolicyThreatLevel returns the PolicyThreatLevel field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetPolicyThreatLevel() int32 {
	if o == nil || IsNil(o.PolicyThreatLevel) {
		var ret int32
		return ret
	}
	return *o.PolicyThreatLevel
}

// GetPolicyThreatLevelOk returns a tuple with the PolicyThreatLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetPolicyThreatLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.PolicyThreatLevel) {
		return nil, false
	}
	return o.PolicyThreatLevel, true
}

// HasPolicyThreatLevel returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasPolicyThreatLevel() bool {
	if o != nil && !IsNil(o.PolicyThreatLevel) {
		return true
	}

	return false
}

// SetPolicyThreatLevel gets a reference to the given int32 and assigns it to the PolicyThreatLevel field.
func (o *SearchResultItemDTO) SetPolicyThreatLevel(v int32) {
	o.PolicyThreatLevel = &v
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetReportId() string {
	if o == nil || IsNil(o.ReportId) {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetReportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReportId) {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasReportId() bool {
	if o != nil && !IsNil(o.ReportId) {
		return true
	}

	return false
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *SearchResultItemDTO) SetReportId(v string) {
	o.ReportId = &v
}

// GetResultIndex returns the ResultIndex field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetResultIndex() int32 {
	if o == nil || IsNil(o.ResultIndex) {
		var ret int32
		return ret
	}
	return *o.ResultIndex
}

// GetResultIndexOk returns a tuple with the ResultIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetResultIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.ResultIndex) {
		return nil, false
	}
	return o.ResultIndex, true
}

// HasResultIndex returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasResultIndex() bool {
	if o != nil && !IsNil(o.ResultIndex) {
		return true
	}

	return false
}

// SetResultIndex gets a reference to the given int32 and assigns it to the ResultIndex field.
func (o *SearchResultItemDTO) SetResultIndex(v int32) {
	o.ResultIndex = &v
}

// GetSbomSpecification returns the SbomSpecification field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetSbomSpecification() string {
	if o == nil || IsNil(o.SbomSpecification) {
		var ret string
		return ret
	}
	return *o.SbomSpecification
}

// GetSbomSpecificationOk returns a tuple with the SbomSpecification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetSbomSpecificationOk() (*string, bool) {
	if o == nil || IsNil(o.SbomSpecification) {
		return nil, false
	}
	return o.SbomSpecification, true
}

// HasSbomSpecification returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasSbomSpecification() bool {
	if o != nil && !IsNil(o.SbomSpecification) {
		return true
	}

	return false
}

// SetSbomSpecification gets a reference to the given string and assigns it to the SbomSpecification field.
func (o *SearchResultItemDTO) SetSbomSpecification(v string) {
	o.SbomSpecification = &v
}

// GetVulnerabilityDescription returns the VulnerabilityDescription field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetVulnerabilityDescription() string {
	if o == nil || IsNil(o.VulnerabilityDescription) {
		var ret string
		return ret
	}
	return *o.VulnerabilityDescription
}

// GetVulnerabilityDescriptionOk returns a tuple with the VulnerabilityDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetVulnerabilityDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.VulnerabilityDescription) {
		return nil, false
	}
	return o.VulnerabilityDescription, true
}

// HasVulnerabilityDescription returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasVulnerabilityDescription() bool {
	if o != nil && !IsNil(o.VulnerabilityDescription) {
		return true
	}

	return false
}

// SetVulnerabilityDescription gets a reference to the given string and assigns it to the VulnerabilityDescription field.
func (o *SearchResultItemDTO) SetVulnerabilityDescription(v string) {
	o.VulnerabilityDescription = &v
}

// GetVulnerabilityId returns the VulnerabilityId field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetVulnerabilityId() string {
	if o == nil || IsNil(o.VulnerabilityId) {
		var ret string
		return ret
	}
	return *o.VulnerabilityId
}

// GetVulnerabilityIdOk returns a tuple with the VulnerabilityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetVulnerabilityIdOk() (*string, bool) {
	if o == nil || IsNil(o.VulnerabilityId) {
		return nil, false
	}
	return o.VulnerabilityId, true
}

// HasVulnerabilityId returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasVulnerabilityId() bool {
	if o != nil && !IsNil(o.VulnerabilityId) {
		return true
	}

	return false
}

// SetVulnerabilityId gets a reference to the given string and assigns it to the VulnerabilityId field.
func (o *SearchResultItemDTO) SetVulnerabilityId(v string) {
	o.VulnerabilityId = &v
}

// GetVulnerabilityStatus returns the VulnerabilityStatus field value if set, zero value otherwise.
func (o *SearchResultItemDTO) GetVulnerabilityStatus() string {
	if o == nil || IsNil(o.VulnerabilityStatus) {
		var ret string
		return ret
	}
	return *o.VulnerabilityStatus
}

// GetVulnerabilityStatusOk returns a tuple with the VulnerabilityStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchResultItemDTO) GetVulnerabilityStatusOk() (*string, bool) {
	if o == nil || IsNil(o.VulnerabilityStatus) {
		return nil, false
	}
	return o.VulnerabilityStatus, true
}

// HasVulnerabilityStatus returns a boolean if a field has been set.
func (o *SearchResultItemDTO) HasVulnerabilityStatus() bool {
	if o != nil && !IsNil(o.VulnerabilityStatus) {
		return true
	}

	return false
}

// SetVulnerabilityStatus gets a reference to the given string and assigns it to the VulnerabilityStatus field.
func (o *SearchResultItemDTO) SetVulnerabilityStatus(v string) {
	o.VulnerabilityStatus = &v
}

func (o SearchResultItemDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchResultItemDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationCategoryColor) {
		toSerialize["applicationCategoryColor"] = o.ApplicationCategoryColor
	}
	if !IsNil(o.ApplicationCategoryDescription) {
		toSerialize["applicationCategoryDescription"] = o.ApplicationCategoryDescription
	}
	if !IsNil(o.ApplicationCategoryId) {
		toSerialize["applicationCategoryId"] = o.ApplicationCategoryId
	}
	if !IsNil(o.ApplicationCategoryName) {
		toSerialize["applicationCategoryName"] = o.ApplicationCategoryName
	}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.ApplicationName) {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if !IsNil(o.ApplicationPublicId) {
		toSerialize["applicationPublicId"] = o.ApplicationPublicId
	}
	if !IsNil(o.ApplicationVersion) {
		toSerialize["applicationVersion"] = o.ApplicationVersion
	}
	if !IsNil(o.ComponentHash) {
		toSerialize["componentHash"] = o.ComponentHash
	}
	if !IsNil(o.ComponentIdentifier) {
		toSerialize["componentIdentifier"] = o.ComponentIdentifier
	}
	if !IsNil(o.ComponentLabelColor) {
		toSerialize["componentLabelColor"] = o.ComponentLabelColor
	}
	if !IsNil(o.ComponentLabelDescription) {
		toSerialize["componentLabelDescription"] = o.ComponentLabelDescription
	}
	if !IsNil(o.ComponentLabelId) {
		toSerialize["componentLabelId"] = o.ComponentLabelId
	}
	if !IsNil(o.ComponentLabelName) {
		toSerialize["componentLabelName"] = o.ComponentLabelName
	}
	if !IsNil(o.ComponentName) {
		toSerialize["componentName"] = o.ComponentName
	}
	if !IsNil(o.ItemType) {
		toSerialize["itemType"] = o.ItemType
	}
	if !IsNil(o.OrganizationId) {
		toSerialize["organizationId"] = o.OrganizationId
	}
	if !IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !IsNil(o.PolicyEvaluationStage) {
		toSerialize["policyEvaluationStage"] = o.PolicyEvaluationStage
	}
	if !IsNil(o.PolicyId) {
		toSerialize["policyId"] = o.PolicyId
	}
	if !IsNil(o.PolicyName) {
		toSerialize["policyName"] = o.PolicyName
	}
	if !IsNil(o.PolicyThreatCategory) {
		toSerialize["policyThreatCategory"] = o.PolicyThreatCategory
	}
	if !IsNil(o.PolicyThreatLevel) {
		toSerialize["policyThreatLevel"] = o.PolicyThreatLevel
	}
	if !IsNil(o.ReportId) {
		toSerialize["reportId"] = o.ReportId
	}
	if !IsNil(o.ResultIndex) {
		toSerialize["resultIndex"] = o.ResultIndex
	}
	if !IsNil(o.SbomSpecification) {
		toSerialize["sbomSpecification"] = o.SbomSpecification
	}
	if !IsNil(o.VulnerabilityDescription) {
		toSerialize["vulnerabilityDescription"] = o.VulnerabilityDescription
	}
	if !IsNil(o.VulnerabilityId) {
		toSerialize["vulnerabilityId"] = o.VulnerabilityId
	}
	if !IsNil(o.VulnerabilityStatus) {
		toSerialize["vulnerabilityStatus"] = o.VulnerabilityStatus
	}
	return toSerialize, nil
}

type NullableSearchResultItemDTO struct {
	value *SearchResultItemDTO
	isSet bool
}

func (v NullableSearchResultItemDTO) Get() *SearchResultItemDTO {
	return v.value
}

func (v *NullableSearchResultItemDTO) Set(val *SearchResultItemDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResultItemDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResultItemDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResultItemDTO(val *SearchResultItemDTO) *NullableSearchResultItemDTO {
	return &NullableSearchResultItemDTO{value: val, isSet: true}
}

func (v NullableSearchResultItemDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResultItemDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


