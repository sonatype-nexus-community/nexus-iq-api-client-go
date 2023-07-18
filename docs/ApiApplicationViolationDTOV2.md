# ApiApplicationViolationDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**ApiApplicationBaseDTO**](ApiApplicationBaseDTO.md) |  | [optional] 
**PolicyViolations** | Pointer to [**[]ApiEnhancedPolicyViolationDTOV2**](ApiEnhancedPolicyViolationDTOV2.md) |  | [optional] 

## Methods

### NewApiApplicationViolationDTOV2

`func NewApiApplicationViolationDTOV2() *ApiApplicationViolationDTOV2`

NewApiApplicationViolationDTOV2 instantiates a new ApiApplicationViolationDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiApplicationViolationDTOV2WithDefaults

`func NewApiApplicationViolationDTOV2WithDefaults() *ApiApplicationViolationDTOV2`

NewApiApplicationViolationDTOV2WithDefaults instantiates a new ApiApplicationViolationDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *ApiApplicationViolationDTOV2) GetApplication() ApiApplicationBaseDTO`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApiApplicationViolationDTOV2) GetApplicationOk() (*ApiApplicationBaseDTO, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApiApplicationViolationDTOV2) SetApplication(v ApiApplicationBaseDTO)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApiApplicationViolationDTOV2) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetPolicyViolations

`func (o *ApiApplicationViolationDTOV2) GetPolicyViolations() []ApiEnhancedPolicyViolationDTOV2`

GetPolicyViolations returns the PolicyViolations field if non-nil, zero value otherwise.

### GetPolicyViolationsOk

`func (o *ApiApplicationViolationDTOV2) GetPolicyViolationsOk() (*[]ApiEnhancedPolicyViolationDTOV2, bool)`

GetPolicyViolationsOk returns a tuple with the PolicyViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolations

`func (o *ApiApplicationViolationDTOV2) SetPolicyViolations(v []ApiEnhancedPolicyViolationDTOV2)`

SetPolicyViolations sets PolicyViolations field to given value.

### HasPolicyViolations

`func (o *ApiApplicationViolationDTOV2) HasPolicyViolations() bool`

HasPolicyViolations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


