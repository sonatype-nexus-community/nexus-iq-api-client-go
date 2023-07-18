# ApiLicenseDataDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeclaredLicenses** | Pointer to [**[]ApiLicenseDTO**](ApiLicenseDTO.md) |  | [optional] 
**EffectiveLicenseThreats** | Pointer to [**[]ApiLicenseThreatDTOV2**](ApiLicenseThreatDTOV2.md) |  | [optional] 
**EffectiveLicenses** | Pointer to [**[]ApiLicenseDTO**](ApiLicenseDTO.md) |  | [optional] 
**ObservedLicenses** | Pointer to [**[]ApiLicenseDTO**](ApiLicenseDTO.md) |  | [optional] 
**OverriddenLicenses** | Pointer to [**[]ApiLicenseDTO**](ApiLicenseDTO.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewApiLicenseDataDTOV2

`func NewApiLicenseDataDTOV2() *ApiLicenseDataDTOV2`

NewApiLicenseDataDTOV2 instantiates a new ApiLicenseDataDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLicenseDataDTOV2WithDefaults

`func NewApiLicenseDataDTOV2WithDefaults() *ApiLicenseDataDTOV2`

NewApiLicenseDataDTOV2WithDefaults instantiates a new ApiLicenseDataDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeclaredLicenses

`func (o *ApiLicenseDataDTOV2) GetDeclaredLicenses() []ApiLicenseDTO`

GetDeclaredLicenses returns the DeclaredLicenses field if non-nil, zero value otherwise.

### GetDeclaredLicensesOk

`func (o *ApiLicenseDataDTOV2) GetDeclaredLicensesOk() (*[]ApiLicenseDTO, bool)`

GetDeclaredLicensesOk returns a tuple with the DeclaredLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredLicenses

`func (o *ApiLicenseDataDTOV2) SetDeclaredLicenses(v []ApiLicenseDTO)`

SetDeclaredLicenses sets DeclaredLicenses field to given value.

### HasDeclaredLicenses

`func (o *ApiLicenseDataDTOV2) HasDeclaredLicenses() bool`

HasDeclaredLicenses returns a boolean if a field has been set.

### GetEffectiveLicenseThreats

`func (o *ApiLicenseDataDTOV2) GetEffectiveLicenseThreats() []ApiLicenseThreatDTOV2`

GetEffectiveLicenseThreats returns the EffectiveLicenseThreats field if non-nil, zero value otherwise.

### GetEffectiveLicenseThreatsOk

`func (o *ApiLicenseDataDTOV2) GetEffectiveLicenseThreatsOk() (*[]ApiLicenseThreatDTOV2, bool)`

GetEffectiveLicenseThreatsOk returns a tuple with the EffectiveLicenseThreats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveLicenseThreats

`func (o *ApiLicenseDataDTOV2) SetEffectiveLicenseThreats(v []ApiLicenseThreatDTOV2)`

SetEffectiveLicenseThreats sets EffectiveLicenseThreats field to given value.

### HasEffectiveLicenseThreats

`func (o *ApiLicenseDataDTOV2) HasEffectiveLicenseThreats() bool`

HasEffectiveLicenseThreats returns a boolean if a field has been set.

### GetEffectiveLicenses

`func (o *ApiLicenseDataDTOV2) GetEffectiveLicenses() []ApiLicenseDTO`

GetEffectiveLicenses returns the EffectiveLicenses field if non-nil, zero value otherwise.

### GetEffectiveLicensesOk

`func (o *ApiLicenseDataDTOV2) GetEffectiveLicensesOk() (*[]ApiLicenseDTO, bool)`

GetEffectiveLicensesOk returns a tuple with the EffectiveLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveLicenses

`func (o *ApiLicenseDataDTOV2) SetEffectiveLicenses(v []ApiLicenseDTO)`

SetEffectiveLicenses sets EffectiveLicenses field to given value.

### HasEffectiveLicenses

`func (o *ApiLicenseDataDTOV2) HasEffectiveLicenses() bool`

HasEffectiveLicenses returns a boolean if a field has been set.

### GetObservedLicenses

`func (o *ApiLicenseDataDTOV2) GetObservedLicenses() []ApiLicenseDTO`

GetObservedLicenses returns the ObservedLicenses field if non-nil, zero value otherwise.

### GetObservedLicensesOk

`func (o *ApiLicenseDataDTOV2) GetObservedLicensesOk() (*[]ApiLicenseDTO, bool)`

GetObservedLicensesOk returns a tuple with the ObservedLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedLicenses

`func (o *ApiLicenseDataDTOV2) SetObservedLicenses(v []ApiLicenseDTO)`

SetObservedLicenses sets ObservedLicenses field to given value.

### HasObservedLicenses

`func (o *ApiLicenseDataDTOV2) HasObservedLicenses() bool`

HasObservedLicenses returns a boolean if a field has been set.

### GetOverriddenLicenses

`func (o *ApiLicenseDataDTOV2) GetOverriddenLicenses() []ApiLicenseDTO`

GetOverriddenLicenses returns the OverriddenLicenses field if non-nil, zero value otherwise.

### GetOverriddenLicensesOk

`func (o *ApiLicenseDataDTOV2) GetOverriddenLicensesOk() (*[]ApiLicenseDTO, bool)`

GetOverriddenLicensesOk returns a tuple with the OverriddenLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenLicenses

`func (o *ApiLicenseDataDTOV2) SetOverriddenLicenses(v []ApiLicenseDTO)`

SetOverriddenLicenses sets OverriddenLicenses field to given value.

### HasOverriddenLicenses

`func (o *ApiLicenseDataDTOV2) HasOverriddenLicenses() bool`

HasOverriddenLicenses returns a boolean if a field has been set.

### GetStatus

`func (o *ApiLicenseDataDTOV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiLicenseDataDTOV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiLicenseDataDTOV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiLicenseDataDTOV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


