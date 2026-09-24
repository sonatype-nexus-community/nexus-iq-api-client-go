# ApiReachabilityEvidenceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | Pointer to [**[]EvidencePath**](EvidencePath.md) |  | [optional] 
**Truncated** | Pointer to **bool** |  | [optional] 
**VulnerabilityId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiReachabilityEvidenceResponse

`func NewApiReachabilityEvidenceResponse() *ApiReachabilityEvidenceResponse`

NewApiReachabilityEvidenceResponse instantiates a new ApiReachabilityEvidenceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiReachabilityEvidenceResponseWithDefaults

`func NewApiReachabilityEvidenceResponseWithDefaults() *ApiReachabilityEvidenceResponse`

NewApiReachabilityEvidenceResponseWithDefaults instantiates a new ApiReachabilityEvidenceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *ApiReachabilityEvidenceResponse) GetPaths() []EvidencePath`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *ApiReachabilityEvidenceResponse) GetPathsOk() (*[]EvidencePath, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *ApiReachabilityEvidenceResponse) SetPaths(v []EvidencePath)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *ApiReachabilityEvidenceResponse) HasPaths() bool`

HasPaths returns a boolean if a field has been set.

### GetTruncated

`func (o *ApiReachabilityEvidenceResponse) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *ApiReachabilityEvidenceResponse) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *ApiReachabilityEvidenceResponse) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *ApiReachabilityEvidenceResponse) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetVulnerabilityId

`func (o *ApiReachabilityEvidenceResponse) GetVulnerabilityId() string`

GetVulnerabilityId returns the VulnerabilityId field if non-nil, zero value otherwise.

### GetVulnerabilityIdOk

`func (o *ApiReachabilityEvidenceResponse) GetVulnerabilityIdOk() (*string, bool)`

GetVulnerabilityIdOk returns a tuple with the VulnerabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityId

`func (o *ApiReachabilityEvidenceResponse) SetVulnerabilityId(v string)`

SetVulnerabilityId sets VulnerabilityId field to given value.

### HasVulnerabilityId

`func (o *ApiReachabilityEvidenceResponse) HasVulnerabilityId() bool`

HasVulnerabilityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


