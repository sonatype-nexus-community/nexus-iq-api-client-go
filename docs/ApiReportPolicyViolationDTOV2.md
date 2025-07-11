# ApiReportPolicyViolationDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to [**[]ApiReportConstraintViolationDTOV2**](ApiReportConstraintViolationDTOV2.md) |  | [optional] 
**FixTime** | Pointer to **time.Time** |  | [optional] 
**Grandfathered** | Pointer to **bool** |  | [optional] 
**LegacyViolation** | Pointer to **bool** |  | [optional] 
**LegacyViolationTime** | Pointer to **time.Time** |  | [optional] 
**OpenTime** | Pointer to **time.Time** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyThreatCategory** | Pointer to **string** |  | [optional] 
**PolicyThreatLevel** | Pointer to **int32** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**WaiveTime** | Pointer to **time.Time** |  | [optional] 
**Waived** | Pointer to **bool** |  | [optional] 
**WaivedWithAutoWaiver** | Pointer to **bool** |  | [optional] 

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

### GetFixTime

`func (o *ApiReportPolicyViolationDTOV2) GetFixTime() time.Time`

GetFixTime returns the FixTime field if non-nil, zero value otherwise.

### GetFixTimeOk

`func (o *ApiReportPolicyViolationDTOV2) GetFixTimeOk() (*time.Time, bool)`

GetFixTimeOk returns a tuple with the FixTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixTime

`func (o *ApiReportPolicyViolationDTOV2) SetFixTime(v time.Time)`

SetFixTime sets FixTime field to given value.

### HasFixTime

`func (o *ApiReportPolicyViolationDTOV2) HasFixTime() bool`

HasFixTime returns a boolean if a field has been set.

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

### GetLegacyViolationTime

`func (o *ApiReportPolicyViolationDTOV2) GetLegacyViolationTime() time.Time`

GetLegacyViolationTime returns the LegacyViolationTime field if non-nil, zero value otherwise.

### GetLegacyViolationTimeOk

`func (o *ApiReportPolicyViolationDTOV2) GetLegacyViolationTimeOk() (*time.Time, bool)`

GetLegacyViolationTimeOk returns a tuple with the LegacyViolationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyViolationTime

`func (o *ApiReportPolicyViolationDTOV2) SetLegacyViolationTime(v time.Time)`

SetLegacyViolationTime sets LegacyViolationTime field to given value.

### HasLegacyViolationTime

`func (o *ApiReportPolicyViolationDTOV2) HasLegacyViolationTime() bool`

HasLegacyViolationTime returns a boolean if a field has been set.

### GetOpenTime

`func (o *ApiReportPolicyViolationDTOV2) GetOpenTime() time.Time`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *ApiReportPolicyViolationDTOV2) GetOpenTimeOk() (*time.Time, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *ApiReportPolicyViolationDTOV2) SetOpenTime(v time.Time)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *ApiReportPolicyViolationDTOV2) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

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

### GetWaiveTime

`func (o *ApiReportPolicyViolationDTOV2) GetWaiveTime() time.Time`

GetWaiveTime returns the WaiveTime field if non-nil, zero value otherwise.

### GetWaiveTimeOk

`func (o *ApiReportPolicyViolationDTOV2) GetWaiveTimeOk() (*time.Time, bool)`

GetWaiveTimeOk returns a tuple with the WaiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiveTime

`func (o *ApiReportPolicyViolationDTOV2) SetWaiveTime(v time.Time)`

SetWaiveTime sets WaiveTime field to given value.

### HasWaiveTime

`func (o *ApiReportPolicyViolationDTOV2) HasWaiveTime() bool`

HasWaiveTime returns a boolean if a field has been set.

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

### GetWaivedWithAutoWaiver

`func (o *ApiReportPolicyViolationDTOV2) GetWaivedWithAutoWaiver() bool`

GetWaivedWithAutoWaiver returns the WaivedWithAutoWaiver field if non-nil, zero value otherwise.

### GetWaivedWithAutoWaiverOk

`func (o *ApiReportPolicyViolationDTOV2) GetWaivedWithAutoWaiverOk() (*bool, bool)`

GetWaivedWithAutoWaiverOk returns a tuple with the WaivedWithAutoWaiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaivedWithAutoWaiver

`func (o *ApiReportPolicyViolationDTOV2) SetWaivedWithAutoWaiver(v bool)`

SetWaivedWithAutoWaiver sets WaivedWithAutoWaiver field to given value.

### HasWaivedWithAutoWaiver

`func (o *ApiReportPolicyViolationDTOV2) HasWaivedWithAutoWaiver() bool`

HasWaivedWithAutoWaiver returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


