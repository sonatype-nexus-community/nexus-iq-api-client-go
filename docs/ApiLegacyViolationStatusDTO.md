# ApiLegacyViolationStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowChange** | Pointer to **bool** | Whether the legacy status can be changed for this owner. | [optional] [readonly] 
**AllowOverride** | Pointer to **bool** | Whether children (orgs and apps) are allowed to override the legacy status. Organization-only. Primitive boolean: an absent or null value in the request body is treated as false. | [optional] 
**Enabled** | Pointer to **bool** | Whether legacy violations are enabled. null means inherit from the parent organization. | [optional] 
**EnabledInParent** | Pointer to **bool** | Whether legacy violations are enabled in the parent organization. null when no parent set this value. | [optional] [readonly] 
**InheritedFromOrganizationName** | Pointer to **string** | The name of the organization the legacy status is inherited from, or null if not inherited. | [optional] [readonly] 

## Methods

### NewApiLegacyViolationStatusDTO

`func NewApiLegacyViolationStatusDTO() *ApiLegacyViolationStatusDTO`

NewApiLegacyViolationStatusDTO instantiates a new ApiLegacyViolationStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLegacyViolationStatusDTOWithDefaults

`func NewApiLegacyViolationStatusDTOWithDefaults() *ApiLegacyViolationStatusDTO`

NewApiLegacyViolationStatusDTOWithDefaults instantiates a new ApiLegacyViolationStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowChange

`func (o *ApiLegacyViolationStatusDTO) GetAllowChange() bool`

GetAllowChange returns the AllowChange field if non-nil, zero value otherwise.

### GetAllowChangeOk

`func (o *ApiLegacyViolationStatusDTO) GetAllowChangeOk() (*bool, bool)`

GetAllowChangeOk returns a tuple with the AllowChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowChange

`func (o *ApiLegacyViolationStatusDTO) SetAllowChange(v bool)`

SetAllowChange sets AllowChange field to given value.

### HasAllowChange

`func (o *ApiLegacyViolationStatusDTO) HasAllowChange() bool`

HasAllowChange returns a boolean if a field has been set.

### GetAllowOverride

`func (o *ApiLegacyViolationStatusDTO) GetAllowOverride() bool`

GetAllowOverride returns the AllowOverride field if non-nil, zero value otherwise.

### GetAllowOverrideOk

`func (o *ApiLegacyViolationStatusDTO) GetAllowOverrideOk() (*bool, bool)`

GetAllowOverrideOk returns a tuple with the AllowOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowOverride

`func (o *ApiLegacyViolationStatusDTO) SetAllowOverride(v bool)`

SetAllowOverride sets AllowOverride field to given value.

### HasAllowOverride

`func (o *ApiLegacyViolationStatusDTO) HasAllowOverride() bool`

HasAllowOverride returns a boolean if a field has been set.

### GetEnabled

`func (o *ApiLegacyViolationStatusDTO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApiLegacyViolationStatusDTO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApiLegacyViolationStatusDTO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApiLegacyViolationStatusDTO) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnabledInParent

`func (o *ApiLegacyViolationStatusDTO) GetEnabledInParent() bool`

GetEnabledInParent returns the EnabledInParent field if non-nil, zero value otherwise.

### GetEnabledInParentOk

`func (o *ApiLegacyViolationStatusDTO) GetEnabledInParentOk() (*bool, bool)`

GetEnabledInParentOk returns a tuple with the EnabledInParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledInParent

`func (o *ApiLegacyViolationStatusDTO) SetEnabledInParent(v bool)`

SetEnabledInParent sets EnabledInParent field to given value.

### HasEnabledInParent

`func (o *ApiLegacyViolationStatusDTO) HasEnabledInParent() bool`

HasEnabledInParent returns a boolean if a field has been set.

### GetInheritedFromOrganizationName

`func (o *ApiLegacyViolationStatusDTO) GetInheritedFromOrganizationName() string`

GetInheritedFromOrganizationName returns the InheritedFromOrganizationName field if non-nil, zero value otherwise.

### GetInheritedFromOrganizationNameOk

`func (o *ApiLegacyViolationStatusDTO) GetInheritedFromOrganizationNameOk() (*string, bool)`

GetInheritedFromOrganizationNameOk returns a tuple with the InheritedFromOrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritedFromOrganizationName

`func (o *ApiLegacyViolationStatusDTO) SetInheritedFromOrganizationName(v string)`

SetInheritedFromOrganizationName sets InheritedFromOrganizationName field to given value.

### HasInheritedFromOrganizationName

`func (o *ApiLegacyViolationStatusDTO) HasInheritedFromOrganizationName() bool`

HasInheritedFromOrganizationName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


