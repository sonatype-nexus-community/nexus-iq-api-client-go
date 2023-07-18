# ApiSourceControlEvaluationRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BranchName** | Pointer to **string** |  | [optional] 
**ScanTargets** | Pointer to **[]string** |  | [optional] 
**StageId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiSourceControlEvaluationRequestDTO

`func NewApiSourceControlEvaluationRequestDTO() *ApiSourceControlEvaluationRequestDTO`

NewApiSourceControlEvaluationRequestDTO instantiates a new ApiSourceControlEvaluationRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSourceControlEvaluationRequestDTOWithDefaults

`func NewApiSourceControlEvaluationRequestDTOWithDefaults() *ApiSourceControlEvaluationRequestDTO`

NewApiSourceControlEvaluationRequestDTOWithDefaults instantiates a new ApiSourceControlEvaluationRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranchName

`func (o *ApiSourceControlEvaluationRequestDTO) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *ApiSourceControlEvaluationRequestDTO) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *ApiSourceControlEvaluationRequestDTO) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *ApiSourceControlEvaluationRequestDTO) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetScanTargets

`func (o *ApiSourceControlEvaluationRequestDTO) GetScanTargets() []string`

GetScanTargets returns the ScanTargets field if non-nil, zero value otherwise.

### GetScanTargetsOk

`func (o *ApiSourceControlEvaluationRequestDTO) GetScanTargetsOk() (*[]string, bool)`

GetScanTargetsOk returns a tuple with the ScanTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanTargets

`func (o *ApiSourceControlEvaluationRequestDTO) SetScanTargets(v []string)`

SetScanTargets sets ScanTargets field to given value.

### HasScanTargets

`func (o *ApiSourceControlEvaluationRequestDTO) HasScanTargets() bool`

HasScanTargets returns a boolean if a field has been set.

### GetStageId

`func (o *ApiSourceControlEvaluationRequestDTO) GetStageId() string`

GetStageId returns the StageId field if non-nil, zero value otherwise.

### GetStageIdOk

`func (o *ApiSourceControlEvaluationRequestDTO) GetStageIdOk() (*string, bool)`

GetStageIdOk returns a tuple with the StageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStageId

`func (o *ApiSourceControlEvaluationRequestDTO) SetStageId(v string)`

SetStageId sets StageId field to given value.

### HasStageId

`func (o *ApiSourceControlEvaluationRequestDTO) HasStageId() bool`

HasStageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


