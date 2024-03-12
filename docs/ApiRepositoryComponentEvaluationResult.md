# ApiRepositoryComponentEvaluationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogDate** | Pointer to **time.Time** |  | [optional] 
**Component** | Pointer to [**ApiRepositoryComponentEvaluationRequest**](ApiRepositoryComponentEvaluationRequest.md) |  | [optional] 
**PolicyViolations** | Pointer to [**[]ApiPolicyViolationDTOV2**](ApiPolicyViolationDTOV2.md) |  | [optional] 
**QuarantineDate** | Pointer to **time.Time** |  | [optional] 
**Quarantined** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiRepositoryComponentEvaluationResult

`func NewApiRepositoryComponentEvaluationResult() *ApiRepositoryComponentEvaluationResult`

NewApiRepositoryComponentEvaluationResult instantiates a new ApiRepositoryComponentEvaluationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRepositoryComponentEvaluationResultWithDefaults

`func NewApiRepositoryComponentEvaluationResultWithDefaults() *ApiRepositoryComponentEvaluationResult`

NewApiRepositoryComponentEvaluationResultWithDefaults instantiates a new ApiRepositoryComponentEvaluationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogDate

`func (o *ApiRepositoryComponentEvaluationResult) GetCatalogDate() time.Time`

GetCatalogDate returns the CatalogDate field if non-nil, zero value otherwise.

### GetCatalogDateOk

`func (o *ApiRepositoryComponentEvaluationResult) GetCatalogDateOk() (*time.Time, bool)`

GetCatalogDateOk returns a tuple with the CatalogDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogDate

`func (o *ApiRepositoryComponentEvaluationResult) SetCatalogDate(v time.Time)`

SetCatalogDate sets CatalogDate field to given value.

### HasCatalogDate

`func (o *ApiRepositoryComponentEvaluationResult) HasCatalogDate() bool`

HasCatalogDate returns a boolean if a field has been set.

### GetComponent

`func (o *ApiRepositoryComponentEvaluationResult) GetComponent() ApiRepositoryComponentEvaluationRequest`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ApiRepositoryComponentEvaluationResult) GetComponentOk() (*ApiRepositoryComponentEvaluationRequest, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ApiRepositoryComponentEvaluationResult) SetComponent(v ApiRepositoryComponentEvaluationRequest)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ApiRepositoryComponentEvaluationResult) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetPolicyViolations

`func (o *ApiRepositoryComponentEvaluationResult) GetPolicyViolations() []ApiPolicyViolationDTOV2`

GetPolicyViolations returns the PolicyViolations field if non-nil, zero value otherwise.

### GetPolicyViolationsOk

`func (o *ApiRepositoryComponentEvaluationResult) GetPolicyViolationsOk() (*[]ApiPolicyViolationDTOV2, bool)`

GetPolicyViolationsOk returns a tuple with the PolicyViolations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolations

`func (o *ApiRepositoryComponentEvaluationResult) SetPolicyViolations(v []ApiPolicyViolationDTOV2)`

SetPolicyViolations sets PolicyViolations field to given value.

### HasPolicyViolations

`func (o *ApiRepositoryComponentEvaluationResult) HasPolicyViolations() bool`

HasPolicyViolations returns a boolean if a field has been set.

### GetQuarantineDate

`func (o *ApiRepositoryComponentEvaluationResult) GetQuarantineDate() time.Time`

GetQuarantineDate returns the QuarantineDate field if non-nil, zero value otherwise.

### GetQuarantineDateOk

`func (o *ApiRepositoryComponentEvaluationResult) GetQuarantineDateOk() (*time.Time, bool)`

GetQuarantineDateOk returns a tuple with the QuarantineDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantineDate

`func (o *ApiRepositoryComponentEvaluationResult) SetQuarantineDate(v time.Time)`

SetQuarantineDate sets QuarantineDate field to given value.

### HasQuarantineDate

`func (o *ApiRepositoryComponentEvaluationResult) HasQuarantineDate() bool`

HasQuarantineDate returns a boolean if a field has been set.

### GetQuarantined

`func (o *ApiRepositoryComponentEvaluationResult) GetQuarantined() bool`

GetQuarantined returns the Quarantined field if non-nil, zero value otherwise.

### GetQuarantinedOk

`func (o *ApiRepositoryComponentEvaluationResult) GetQuarantinedOk() (*bool, bool)`

GetQuarantinedOk returns a tuple with the Quarantined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarantined

`func (o *ApiRepositoryComponentEvaluationResult) SetQuarantined(v bool)`

SetQuarantined sets Quarantined field to given value.

### HasQuarantined

`func (o *ApiRepositoryComponentEvaluationResult) HasQuarantined() bool`

HasQuarantined returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


