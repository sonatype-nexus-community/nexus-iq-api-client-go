# ApiPolicyViolationDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConstraintViolations** | Pointer to [**[]ApiConstraintViolationDTO**](ApiConstraintViolationDTO.md) |  | [optional] 
**FixTime** | Pointer to **time.Time** |  | [optional] 
**LegacyViolationTime** | Pointer to **time.Time** |  | [optional] 
**OpenTime** | Pointer to **time.Time** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 
**WaiveTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewApiPolicyViolationDTOV2

`func NewApiPolicyViolationDTOV2() *ApiPolicyViolationDTOV2`

NewApiPolicyViolationDTOV2 instantiates a new ApiPolicyViolationDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyViolationDTOV2WithDefaults

`func NewApiPolicyViolationDTOV2WithDefaults() *ApiPolicyViolationDTOV2`

NewApiPolicyViolationDTOV2WithDefaults instantiates a new ApiPolicyViolationDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraintViolations

`func (o *ApiPolicyViolationDTOV2) GetConstraintViolations() []ApiConstraintViolationDTO`

GetConstraintViolations returns the ConstraintViolations field if non-nil, zero value otherwise.

### GetConstraintViolationsOk

`func (o *ApiPolicyViolationDTOV2) GetConstraintViolationsOk() (*[]ApiConstraintViolationDTO, bool)`

GetConstraintViolationsOk returns a tuple with the ConstraintViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintViolations

`func (o *ApiPolicyViolationDTOV2) SetConstraintViolations(v []ApiConstraintViolationDTO)`

SetConstraintViolations sets ConstraintViolations field to given value.

### HasConstraintViolations

`func (o *ApiPolicyViolationDTOV2) HasConstraintViolations() bool`

HasConstraintViolations returns a boolean if a field has been set.

### GetFixTime

`func (o *ApiPolicyViolationDTOV2) GetFixTime() time.Time`

GetFixTime returns the FixTime field if non-nil, zero value otherwise.

### GetFixTimeOk

`func (o *ApiPolicyViolationDTOV2) GetFixTimeOk() (*time.Time, bool)`

GetFixTimeOk returns a tuple with the FixTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixTime

`func (o *ApiPolicyViolationDTOV2) SetFixTime(v time.Time)`

SetFixTime sets FixTime field to given value.

### HasFixTime

`func (o *ApiPolicyViolationDTOV2) HasFixTime() bool`

HasFixTime returns a boolean if a field has been set.

### GetLegacyViolationTime

`func (o *ApiPolicyViolationDTOV2) GetLegacyViolationTime() time.Time`

GetLegacyViolationTime returns the LegacyViolationTime field if non-nil, zero value otherwise.

### GetLegacyViolationTimeOk

`func (o *ApiPolicyViolationDTOV2) GetLegacyViolationTimeOk() (*time.Time, bool)`

GetLegacyViolationTimeOk returns a tuple with the LegacyViolationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyViolationTime

`func (o *ApiPolicyViolationDTOV2) SetLegacyViolationTime(v time.Time)`

SetLegacyViolationTime sets LegacyViolationTime field to given value.

### HasLegacyViolationTime

`func (o *ApiPolicyViolationDTOV2) HasLegacyViolationTime() bool`

HasLegacyViolationTime returns a boolean if a field has been set.

### GetOpenTime

`func (o *ApiPolicyViolationDTOV2) GetOpenTime() time.Time`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *ApiPolicyViolationDTOV2) GetOpenTimeOk() (*time.Time, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *ApiPolicyViolationDTOV2) SetOpenTime(v time.Time)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *ApiPolicyViolationDTOV2) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApiPolicyViolationDTOV2) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApiPolicyViolationDTOV2) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApiPolicyViolationDTOV2) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApiPolicyViolationDTOV2) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiPolicyViolationDTOV2) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiPolicyViolationDTOV2) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiPolicyViolationDTOV2) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiPolicyViolationDTOV2) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *ApiPolicyViolationDTOV2) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *ApiPolicyViolationDTOV2) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *ApiPolicyViolationDTOV2) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *ApiPolicyViolationDTOV2) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiPolicyViolationDTOV2) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiPolicyViolationDTOV2) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiPolicyViolationDTOV2) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiPolicyViolationDTOV2) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.

### GetWaiveTime

`func (o *ApiPolicyViolationDTOV2) GetWaiveTime() time.Time`

GetWaiveTime returns the WaiveTime field if non-nil, zero value otherwise.

### GetWaiveTimeOk

`func (o *ApiPolicyViolationDTOV2) GetWaiveTimeOk() (*time.Time, bool)`

GetWaiveTimeOk returns a tuple with the WaiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiveTime

`func (o *ApiPolicyViolationDTOV2) SetWaiveTime(v time.Time)`

SetWaiveTime sets WaiveTime field to given value.

### HasWaiveTime

`func (o *ApiPolicyViolationDTOV2) HasWaiveTime() bool`

HasWaiveTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


