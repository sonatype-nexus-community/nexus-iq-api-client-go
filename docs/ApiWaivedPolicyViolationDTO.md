# ApiWaivedPolicyViolationDTO

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
**PolicyWaiver** | Pointer to [**ApiPolicyWaiverDTO**](ApiPolicyWaiverDTO.md) |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 
**WaiveTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewApiWaivedPolicyViolationDTO

`func NewApiWaivedPolicyViolationDTO() *ApiWaivedPolicyViolationDTO`

NewApiWaivedPolicyViolationDTO instantiates a new ApiWaivedPolicyViolationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWaivedPolicyViolationDTOWithDefaults

`func NewApiWaivedPolicyViolationDTOWithDefaults() *ApiWaivedPolicyViolationDTO`

NewApiWaivedPolicyViolationDTOWithDefaults instantiates a new ApiWaivedPolicyViolationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraintViolations

`func (o *ApiWaivedPolicyViolationDTO) GetConstraintViolations() []ApiConstraintViolationDTO`

GetConstraintViolations returns the ConstraintViolations field if non-nil, zero value otherwise.

### GetConstraintViolationsOk

`func (o *ApiWaivedPolicyViolationDTO) GetConstraintViolationsOk() (*[]ApiConstraintViolationDTO, bool)`

GetConstraintViolationsOk returns a tuple with the ConstraintViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintViolations

`func (o *ApiWaivedPolicyViolationDTO) SetConstraintViolations(v []ApiConstraintViolationDTO)`

SetConstraintViolations sets ConstraintViolations field to given value.

### HasConstraintViolations

`func (o *ApiWaivedPolicyViolationDTO) HasConstraintViolations() bool`

HasConstraintViolations returns a boolean if a field has been set.

### GetFixTime

`func (o *ApiWaivedPolicyViolationDTO) GetFixTime() time.Time`

GetFixTime returns the FixTime field if non-nil, zero value otherwise.

### GetFixTimeOk

`func (o *ApiWaivedPolicyViolationDTO) GetFixTimeOk() (*time.Time, bool)`

GetFixTimeOk returns a tuple with the FixTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixTime

`func (o *ApiWaivedPolicyViolationDTO) SetFixTime(v time.Time)`

SetFixTime sets FixTime field to given value.

### HasFixTime

`func (o *ApiWaivedPolicyViolationDTO) HasFixTime() bool`

HasFixTime returns a boolean if a field has been set.

### GetLegacyViolationTime

`func (o *ApiWaivedPolicyViolationDTO) GetLegacyViolationTime() time.Time`

GetLegacyViolationTime returns the LegacyViolationTime field if non-nil, zero value otherwise.

### GetLegacyViolationTimeOk

`func (o *ApiWaivedPolicyViolationDTO) GetLegacyViolationTimeOk() (*time.Time, bool)`

GetLegacyViolationTimeOk returns a tuple with the LegacyViolationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyViolationTime

`func (o *ApiWaivedPolicyViolationDTO) SetLegacyViolationTime(v time.Time)`

SetLegacyViolationTime sets LegacyViolationTime field to given value.

### HasLegacyViolationTime

`func (o *ApiWaivedPolicyViolationDTO) HasLegacyViolationTime() bool`

HasLegacyViolationTime returns a boolean if a field has been set.

### GetOpenTime

`func (o *ApiWaivedPolicyViolationDTO) GetOpenTime() time.Time`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *ApiWaivedPolicyViolationDTO) GetOpenTimeOk() (*time.Time, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *ApiWaivedPolicyViolationDTO) SetOpenTime(v time.Time)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *ApiWaivedPolicyViolationDTO) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApiWaivedPolicyViolationDTO) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApiWaivedPolicyViolationDTO) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApiWaivedPolicyViolationDTO) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApiWaivedPolicyViolationDTO) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiWaivedPolicyViolationDTO) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiWaivedPolicyViolationDTO) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiWaivedPolicyViolationDTO) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiWaivedPolicyViolationDTO) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *ApiWaivedPolicyViolationDTO) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *ApiWaivedPolicyViolationDTO) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *ApiWaivedPolicyViolationDTO) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *ApiWaivedPolicyViolationDTO) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetPolicyWaiver

`func (o *ApiWaivedPolicyViolationDTO) GetPolicyWaiver() ApiPolicyWaiverDTO`

GetPolicyWaiver returns the PolicyWaiver field if non-nil, zero value otherwise.

### GetPolicyWaiverOk

`func (o *ApiWaivedPolicyViolationDTO) GetPolicyWaiverOk() (*ApiPolicyWaiverDTO, bool)`

GetPolicyWaiverOk returns a tuple with the PolicyWaiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiver

`func (o *ApiWaivedPolicyViolationDTO) SetPolicyWaiver(v ApiPolicyWaiverDTO)`

SetPolicyWaiver sets PolicyWaiver field to given value.

### HasPolicyWaiver

`func (o *ApiWaivedPolicyViolationDTO) HasPolicyWaiver() bool`

HasPolicyWaiver returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiWaivedPolicyViolationDTO) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiWaivedPolicyViolationDTO) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiWaivedPolicyViolationDTO) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiWaivedPolicyViolationDTO) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.

### GetWaiveTime

`func (o *ApiWaivedPolicyViolationDTO) GetWaiveTime() time.Time`

GetWaiveTime returns the WaiveTime field if non-nil, zero value otherwise.

### GetWaiveTimeOk

`func (o *ApiWaivedPolicyViolationDTO) GetWaiveTimeOk() (*time.Time, bool)`

GetWaiveTimeOk returns a tuple with the WaiveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiveTime

`func (o *ApiWaivedPolicyViolationDTO) SetWaiveTime(v time.Time)`

SetWaiveTime sets WaiveTime field to given value.

### HasWaiveTime

`func (o *ApiWaivedPolicyViolationDTO) HasWaiveTime() bool`

HasWaiveTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


