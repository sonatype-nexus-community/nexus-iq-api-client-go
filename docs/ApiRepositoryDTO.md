# ApiRepositoryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuditEnabled** | Pointer to **bool** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**NamespaceConfusionProtectionEnabled** | Pointer to **bool** |  | [optional] 
**PolicyCompliantComponentSelectionEnabled** | Pointer to **bool** |  | [optional] 
**PublicId** | Pointer to **string** |  | [optional] 
**QuarantineEnabled** | Pointer to **bool** |  | [optional] 
**RepositoryId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewApiRepositoryDTO

`func NewApiRepositoryDTO() *ApiRepositoryDTO`

NewApiRepositoryDTO instantiates a new ApiRepositoryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRepositoryDTOWithDefaults

`func NewApiRepositoryDTOWithDefaults() *ApiRepositoryDTO`

NewApiRepositoryDTOWithDefaults instantiates a new ApiRepositoryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuditEnabled

`func (o *ApiRepositoryDTO) GetAuditEnabled() bool`

GetAuditEnabled returns the AuditEnabled field if non-nil, zero value otherwise.

### GetAuditEnabledOk

`func (o *ApiRepositoryDTO) GetAuditEnabledOk() (*bool, bool)`

GetAuditEnabledOk returns a tuple with the AuditEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditEnabled

`func (o *ApiRepositoryDTO) SetAuditEnabled(v bool)`

SetAuditEnabled sets AuditEnabled field to given value.

### HasAuditEnabled

`func (o *ApiRepositoryDTO) HasAuditEnabled() bool`

HasAuditEnabled returns a boolean if a field has been set.

### GetFormat

`func (o *ApiRepositoryDTO) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ApiRepositoryDTO) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ApiRepositoryDTO) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ApiRepositoryDTO) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetNamespaceConfusionProtectionEnabled

`func (o *ApiRepositoryDTO) GetNamespaceConfusionProtectionEnabled() bool`

GetNamespaceConfusionProtectionEnabled returns the NamespaceConfusionProtectionEnabled field if non-nil, zero value otherwise.

### GetNamespaceConfusionProtectionEnabledOk

`func (o *ApiRepositoryDTO) GetNamespaceConfusionProtectionEnabledOk() (*bool, bool)`

GetNamespaceConfusionProtectionEnabledOk returns a tuple with the NamespaceConfusionProtectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceConfusionProtectionEnabled

`func (o *ApiRepositoryDTO) SetNamespaceConfusionProtectionEnabled(v bool)`

SetNamespaceConfusionProtectionEnabled sets NamespaceConfusionProtectionEnabled field to given value.

### HasNamespaceConfusionProtectionEnabled

`func (o *ApiRepositoryDTO) HasNamespaceConfusionProtectionEnabled() bool`

HasNamespaceConfusionProtectionEnabled returns a boolean if a field has been set.

### GetPolicyCompliantComponentSelectionEnabled

`func (o *ApiRepositoryDTO) GetPolicyCompliantComponentSelectionEnabled() bool`

GetPolicyCompliantComponentSelectionEnabled returns the PolicyCompliantComponentSelectionEnabled field if non-nil, zero value otherwise.

### GetPolicyCompliantComponentSelectionEnabledOk

`func (o *ApiRepositoryDTO) GetPolicyCompliantComponentSelectionEnabledOk() (*bool, bool)`

GetPolicyCompliantComponentSelectionEnabledOk returns a tuple with the PolicyCompliantComponentSelectionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyCompliantComponentSelectionEnabled

`func (o *ApiRepositoryDTO) SetPolicyCompliantComponentSelectionEnabled(v bool)`

SetPolicyCompliantComponentSelectionEnabled sets PolicyCompliantComponentSelectionEnabled field to given value.

### HasPolicyCompliantComponentSelectionEnabled

`func (o *ApiRepositoryDTO) HasPolicyCompliantComponentSelectionEnabled() bool`

HasPolicyCompliantComponentSelectionEnabled returns a boolean if a field has been set.

### GetPublicId

`func (o *ApiRepositoryDTO) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *ApiRepositoryDTO) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *ApiRepositoryDTO) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *ApiRepositoryDTO) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetQuarantineEnabled

`func (o *ApiRepositoryDTO) GetQuarantineEnabled() bool`

GetQuarantineEnabled returns the QuarantineEnabled field if non-nil, zero value otherwise.

### GetQuarantineEnabledOk

`func (o *ApiRepositoryDTO) GetQuarantineEnabledOk() (*bool, bool)`

GetQuarantineEnabledOk returns a tuple with the QuarantineEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineEnabled

`func (o *ApiRepositoryDTO) SetQuarantineEnabled(v bool)`

SetQuarantineEnabled sets QuarantineEnabled field to given value.

### HasQuarantineEnabled

`func (o *ApiRepositoryDTO) HasQuarantineEnabled() bool`

HasQuarantineEnabled returns a boolean if a field has been set.

### GetRepositoryId

`func (o *ApiRepositoryDTO) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *ApiRepositoryDTO) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *ApiRepositoryDTO) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *ApiRepositoryDTO) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetType

`func (o *ApiRepositoryDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiRepositoryDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiRepositoryDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiRepositoryDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


