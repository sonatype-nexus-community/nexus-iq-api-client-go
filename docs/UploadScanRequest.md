# UploadScanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentId** | Pointer to **string** |  | [optional] 
**PolicyEvaluationStage** | Pointer to **string** |  | [optional] 
**Purl** | Pointer to **string** |  | [optional] 
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


