# ApiReportPolicyViolationDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to [**[]ApiReportConstraintViolationDTOV2**](ApiReportConstraintViolationDTOV2.md) |  | [optional] 
**Grandfathered** | Pointer to **bool** |  | [optional] 
**LegacyViolation** | Pointer to **bool** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyThreatCategory** | Pointer to **string** |  | [optional] 
**PolicyThreatLevel** | Pointer to **int32** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**Waived** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiReportPolicyViolationDTOV2

`func NewApiReportPolicyViolationDTOV2() *ApiReportPolicyViolationDTOV2`

NewApiReportPolicyViolationDTOV2 instantiates a new ApiReportPolicyViolationDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiReportPolicyViolationDTOV2WithDefaults

`func NewApiReportPolicyViolationDTOV2WithDefaults() *ApiReportPolicyViolationDTOV2`

NewApiReportPolicyViolationDTOV2WithDefaults instantiates a new ApiReportPolicyViolationDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *ApiReportPolicyViolationDTOV2) GetConstraints() []ApiReportConstraintViolationDTOV2`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *ApiReportPolicyViolationDTOV2) GetConstraintsOk() (*[]ApiReportConstraintViolationDTOV2, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *ApiReportPolicyViolationDTOV2) SetConstraints(v []ApiReportConstraintViolationDTOV2)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *ApiReportPolicyViolationDTOV2) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetGrandfathered

`func (o *ApiReportPolicyViolationDTOV2) GetGrandfathered() bool`

GetGrandfathered returns the Grandfathered field if non-nil, zero value otherwise.

### GetGrandfatheredOk

`func (o *ApiReportPolicyViolationDTOV2) GetGrandfatheredOk() (*bool, bool)`

GetGrandfatheredOk returns a tuple with the Grandfathered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrandfathered

`func (o *ApiReportPolicyViolationDTOV2) SetGrandfathered(v bool)`

SetGrandfathered sets Grandfathered field to given value.

### HasGrandfathered

`func (o *ApiReportPolicyViolationDTOV2) HasGrandfathered() bool`

HasGrandfathered returns a boolean if a field has been set.

### GetLegacyViolation

`func (o *ApiReportPolicyViolationDTOV2) GetLegacyViolation() bool`

GetLegacyViolation returns the LegacyViolation field if non-nil, zero value otherwise.

### GetLegacyViolationOk

`func (o *ApiReportPolicyViolationDTOV2) GetLegacyViolationOk() (*bool, bool)`

GetLegacyViolationOk returns a tuple with the LegacyViolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyViolation

`func (o *ApiReportPolicyViolationDTOV2) SetLegacyViolation(v bool)`

SetLegacyViolation sets LegacyViolation field to given value.

### HasLegacyViolation

`func (o *ApiReportPolicyViolationDTOV2) HasLegacyViolation() bool`

HasLegacyViolation returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApiReportPolicyViolationDTOV2) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApiReportPolicyViolationDTOV2) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApiReportPolicyViolationDTOV2) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApiReportPolicyViolationDTOV2) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiReportPolicyViolationDTOV2) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiReportPolicyViolationDTOV2) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiReportPolicyViolationDTOV2) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiReportPolicyViolationDTOV2) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyThreatCategory

`func (o *ApiReportPolicyViolationDTOV2) GetPolicyThreatCategory() string`

GetPolicyThreatCategory returns the PolicyThreatCategory field if non-nil, zero value otherwise.

### GetPolicyThreatCategoryOk

`func (o *ApiReportPolicyViolationDTOV2) GetPolicyThreatCategoryOk() (*string, bool)`

GetPolicyThreatCategoryOk returns a tuple with the PolicyThreatCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyThreatCategory

`func (o *ApiReportPolicyViolationDTOV2) SetPolicyThreatCategory(v string)`

SetPolicyThreatCategory sets PolicyThreatCategory field to given value.

### HasPolicyThreatCategory

`func (o *ApiReportPolicyViolationDTOV2) HasPolicyThreatCategory() bool`

HasPolicyThreatCategory returns a boolean if a field has been set.

### GetPolicyThreatLevel

`func (o *ApiReportPolicyViolationDTOV2) GetPolicyThreatLevel() int32`

GetPolicyThreatLevel returns the PolicyThreatLevel field if non-nil, zero value otherwise.

### GetPolicyThreatLevelOk

`func (o *ApiReportPolicyViolationDTOV2) GetPolicyThreatLevelOk() (*int32, bool)`

GetPolicyThreatLevelOk returns a tuple with the PolicyThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyThreatLevel

`func (o *ApiReportPolicyViolationDTOV2) SetPolicyThreatLevel(v int32)`

SetPolicyThreatLevel sets PolicyThreatLevel field to given value.

### HasPolicyThreatLevel

`func (o *ApiReportPolicyViolationDTOV2) HasPolicyThreatLevel() bool`

HasPolicyThreatLevel returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *ApiReportPolicyViolationDTOV2) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *ApiReportPolicyViolationDTOV2) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *ApiReportPolicyViolationDTOV2) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *ApiReportPolicyViolationDTOV2) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetWaived

`func (o *ApiReportPolicyViolationDTOV2) GetWaived() bool`

GetWaived returns the Waived field if non-nil, zero value otherwise.

### GetWaivedOk

`func (o *ApiReportPolicyViolationDTOV2) GetWaivedOk() (*bool, bool)`

GetWaivedOk returns a tuple with the Waived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaived

`func (o *ApiReportPolicyViolationDTOV2) SetWaived(v bool)`

SetWaived sets Waived field to given value.

### HasWaived

`func (o *ApiReportPolicyViolationDTOV2) HasWaived() bool`

HasWaived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


