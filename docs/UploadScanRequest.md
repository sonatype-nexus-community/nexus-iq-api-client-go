# UploadScanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to **string** |  | [optional] 
**ComponentId** | Pointer to **string** |  | [optional] 
**CorrelationId** | Pointer to **string** |  | [optional] 
**EvaluationMode** | Pointer to **string** |  | [optional] 
**PolicyEvaluationStage** | Pointer to **string** |  | [optional] 
**Purl** | Pointer to **string** |  | [optional] 
**RequestedBy** | Pointer to **string** |  | [optional] 
**ScanFile** | Pointer to ***os.File** |  | [optional] 

## Methods

### NewUploadScanRequest

`func NewUploadScanRequest() *UploadScanRequest`

NewUploadScanRequest instantiates a new UploadScanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadScanRequestWithDefaults

`func NewUploadScanRequestWithDefaults() *UploadScanRequest`

NewUploadScanRequestWithDefaults instantiates a new UploadScanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *UploadScanRequest) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *UploadScanRequest) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *UploadScanRequest) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *UploadScanRequest) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetComponentId

`func (o *UploadScanRequest) GetComponentId() string`

GetComponentId returns the ComponentId field if non-nil, zero value otherwise.

### GetComponentIdOk

`func (o *UploadScanRequest) GetComponentIdOk() (*string, bool)`

GetComponentIdOk returns a tuple with the ComponentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentId

`func (o *UploadScanRequest) SetComponentId(v string)`

SetComponentId sets ComponentId field to given value.

### HasComponentId

`func (o *UploadScanRequest) HasComponentId() bool`

HasComponentId returns a boolean if a field has been set.

### GetCorrelationId

`func (o *UploadScanRequest) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *UploadScanRequest) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *UploadScanRequest) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *UploadScanRequest) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetEvaluationMode

`func (o *UploadScanRequest) GetEvaluationMode() string`

GetEvaluationMode returns the EvaluationMode field if non-nil, zero value otherwise.

### GetEvaluationModeOk

`func (o *UploadScanRequest) GetEvaluationModeOk() (*string, bool)`

GetEvaluationModeOk returns a tuple with the EvaluationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationMode

`func (o *UploadScanRequest) SetEvaluationMode(v string)`

SetEvaluationMode sets EvaluationMode field to given value.

### HasEvaluationMode

`func (o *UploadScanRequest) HasEvaluationMode() bool`

HasEvaluationMode returns a boolean if a field has been set.

### GetPolicyEvaluationStage

`func (o *UploadScanRequest) GetPolicyEvaluationStage() string`

GetPolicyEvaluationStage returns the PolicyEvaluationStage field if non-nil, zero value otherwise.

### GetPolicyEvaluationStageOk

`func (o *UploadScanRequest) GetPolicyEvaluationStageOk() (*string, bool)`

GetPolicyEvaluationStageOk returns a tuple with the PolicyEvaluationStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyEvaluationStage

`func (o *UploadScanRequest) SetPolicyEvaluationStage(v string)`

SetPolicyEvaluationStage sets PolicyEvaluationStage field to given value.

### HasPolicyEvaluationStage

`func (o *UploadScanRequest) HasPolicyEvaluationStage() bool`

HasPolicyEvaluationStage returns a boolean if a field has been set.

### GetPurl

`func (o *UploadScanRequest) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *UploadScanRequest) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *UploadScanRequest) SetPurl(v string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *UploadScanRequest) HasPurl() bool`

HasPurl returns a boolean if a field has been set.

### GetRequestedBy

`func (o *UploadScanRequest) GetRequestedBy() string`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *UploadScanRequest) GetRequestedByOk() (*string, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *UploadScanRequest) SetRequestedBy(v string)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *UploadScanRequest) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetScanFile

`func (o *UploadScanRequest) GetScanFile() *os.File`

GetScanFile returns the ScanFile field if non-nil, zero value otherwise.

### GetScanFileOk

`func (o *UploadScanRequest) GetScanFileOk() (**os.File, bool)`

GetScanFileOk returns a tuple with the ScanFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanFile

`func (o *UploadScanRequest) SetScanFile(v *os.File)`

SetScanFile sets ScanFile field to given value.

### HasScanFile

`func (o *UploadScanRequest) HasScanFile() bool`

HasScanFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


