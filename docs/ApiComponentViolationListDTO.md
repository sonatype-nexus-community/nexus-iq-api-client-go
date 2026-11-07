# ApiComponentViolationListDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxThreatLevel** | Pointer to **int32** |  | [optional] 
**TotalViolations** | Pointer to **int32** |  | [optional] 
**Violations** | Pointer to [**[]ApiComponentViolationDTO**](ApiComponentViolationDTO.md) |  | [optional] 

## Methods

### NewApiComponentViolationListDTO

`func NewApiComponentViolationListDTO() *ApiComponentViolationListDTO`

NewApiComponentViolationListDTO instantiates a new ApiComponentViolationListDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiComponentViolationListDTOWithDefaults

`func NewApiComponentViolationListDTOWithDefaults() *ApiComponentViolationListDTO`

NewApiComponentViolationListDTOWithDefaults instantiates a new ApiComponentViolationListDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxThreatLevel

`func (o *ApiComponentViolationListDTO) GetMaxThreatLevel() int32`

GetMaxThreatLevel returns the MaxThreatLevel field if non-nil, zero value otherwise.

### GetMaxThreatLevelOk

`func (o *ApiComponentViolationListDTO) GetMaxThreatLevelOk() (*int32, bool)`

GetMaxThreatLevelOk returns a tuple with the MaxThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThreatLevel

`func (o *ApiComponentViolationListDTO) SetMaxThreatLevel(v int32)`

SetMaxThreatLevel sets MaxThreatLevel field to given value.

### HasMaxThreatLevel

`func (o *ApiComponentViolationListDTO) HasMaxThreatLevel() bool`

HasMaxThreatLevel returns a boolean if a field has been set.

### GetTotalViolations

`func (o *ApiComponentViolationListDTO) GetTotalViolations() int32`

GetTotalViolations returns the TotalViolations field if non-nil, zero value otherwise.

### GetTotalViolationsOk

`func (o *ApiComponentViolationListDTO) GetTotalViolationsOk() (*int32, bool)`

GetTotalViolationsOk returns a tuple with the TotalViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalViolations

`func (o *ApiComponentViolationListDTO) SetTotalViolations(v int32)`

SetTotalViolations sets TotalViolations field to given value.

### HasTotalViolations

`func (o *ApiComponentViolationListDTO) HasTotalViolations() bool`

HasTotalViolations returns a boolean if a field has been set.

### GetViolations

`func (o *ApiComponentViolationListDTO) GetViolations() []ApiComponentViolationDTO`

GetViolations returns the Violations field if non-nil, zero value otherwise.

### GetViolationsOk

`func (o *ApiComponentViolationListDTO) GetViolationsOk() (*[]ApiComponentViolationDTO, bool)`

GetViolationsOk returns a tuple with the Violations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolations

`func (o *ApiComponentViolationListDTO) SetViolations(v []ApiComponentViolationDTO)`

SetViolations sets Violations field to given value.

### HasViolations

`func (o *ApiComponentViolationListDTO) HasViolations() bool`

HasViolations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


