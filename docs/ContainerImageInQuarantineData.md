# ContainerImageInQuarantineData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**ApplicationName** | Pointer to **string** |  | [optional] 
**ApplicationPublicId** | Pointer to **string** |  | [optional] 
**OpenTime** | Pointer to **time.Time** |  | [optional] 
**PolicyViolationCount** | Pointer to **int64** |  | [optional] 
**RepositoryId** | Pointer to **string** |  | [optional] 
**RepositoryPublicId** | Pointer to **string** |  | [optional] 
**ScanId** | Pointer to **string** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewContainerImageInQuarantineData

`func NewContainerImageInQuarantineData() *ContainerImageInQuarantineData`

NewContainerImageInQuarantineData instantiates a new ContainerImageInQuarantineData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImageInQuarantineDataWithDefaults

`func NewContainerImageInQuarantineDataWithDefaults() *ContainerImageInQuarantineData`

NewContainerImageInQuarantineDataWithDefaults instantiates a new ContainerImageInQuarantineData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ContainerImageInQuarantineData) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ContainerImageInQuarantineData) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ContainerImageInQuarantineData) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ContainerImageInQuarantineData) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationName

`func (o *ContainerImageInQuarantineData) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *ContainerImageInQuarantineData) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *ContainerImageInQuarantineData) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *ContainerImageInQuarantineData) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationPublicId

`func (o *ContainerImageInQuarantineData) GetApplicationPublicId() string`

GetApplicationPublicId returns the ApplicationPublicId field if non-nil, zero value otherwise.

### GetApplicationPublicIdOk

`func (o *ContainerImageInQuarantineData) GetApplicationPublicIdOk() (*string, bool)`

GetApplicationPublicIdOk returns a tuple with the ApplicationPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPublicId

`func (o *ContainerImageInQuarantineData) SetApplicationPublicId(v string)`

SetApplicationPublicId sets ApplicationPublicId field to given value.

### HasApplicationPublicId

`func (o *ContainerImageInQuarantineData) HasApplicationPublicId() bool`

HasApplicationPublicId returns a boolean if a field has been set.

### GetOpenTime

`func (o *ContainerImageInQuarantineData) GetOpenTime() time.Time`

GetOpenTime returns the OpenTime field if non-nil, zero value otherwise.

### GetOpenTimeOk

`func (o *ContainerImageInQuarantineData) GetOpenTimeOk() (*time.Time, bool)`

GetOpenTimeOk returns a tuple with the OpenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenTime

`func (o *ContainerImageInQuarantineData) SetOpenTime(v time.Time)`

SetOpenTime sets OpenTime field to given value.

### HasOpenTime

`func (o *ContainerImageInQuarantineData) HasOpenTime() bool`

HasOpenTime returns a boolean if a field has been set.

### GetPolicyViolationCount

`func (o *ContainerImageInQuarantineData) GetPolicyViolationCount() int64`

GetPolicyViolationCount returns the PolicyViolationCount field if non-nil, zero value otherwise.

### GetPolicyViolationCountOk

`func (o *ContainerImageInQuarantineData) GetPolicyViolationCountOk() (*int64, bool)`

GetPolicyViolationCountOk returns a tuple with the PolicyViolationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationCount

`func (o *ContainerImageInQuarantineData) SetPolicyViolationCount(v int64)`

SetPolicyViolationCount sets PolicyViolationCount field to given value.

### HasPolicyViolationCount

`func (o *ContainerImageInQuarantineData) HasPolicyViolationCount() bool`

HasPolicyViolationCount returns a boolean if a field has been set.

### GetRepositoryId

`func (o *ContainerImageInQuarantineData) GetRepositoryId() string`

GetRepositoryId returns the RepositoryId field if non-nil, zero value otherwise.

### GetRepositoryIdOk

`func (o *ContainerImageInQuarantineData) GetRepositoryIdOk() (*string, bool)`

GetRepositoryIdOk returns a tuple with the RepositoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryId

`func (o *ContainerImageInQuarantineData) SetRepositoryId(v string)`

SetRepositoryId sets RepositoryId field to given value.

### HasRepositoryId

`func (o *ContainerImageInQuarantineData) HasRepositoryId() bool`

HasRepositoryId returns a boolean if a field has been set.

### GetRepositoryPublicId

`func (o *ContainerImageInQuarantineData) GetRepositoryPublicId() string`

GetRepositoryPublicId returns the RepositoryPublicId field if non-nil, zero value otherwise.

### GetRepositoryPublicIdOk

`func (o *ContainerImageInQuarantineData) GetRepositoryPublicIdOk() (*string, bool)`

GetRepositoryPublicIdOk returns a tuple with the RepositoryPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryPublicId

`func (o *ContainerImageInQuarantineData) SetRepositoryPublicId(v string)`

SetRepositoryPublicId sets RepositoryPublicId field to given value.

### HasRepositoryPublicId

`func (o *ContainerImageInQuarantineData) HasRepositoryPublicId() bool`

HasRepositoryPublicId returns a boolean if a field has been set.

### GetScanId

`func (o *ContainerImageInQuarantineData) GetScanId() string`

GetScanId returns the ScanId field if non-nil, zero value otherwise.

### GetScanIdOk

`func (o *ContainerImageInQuarantineData) GetScanIdOk() (*string, bool)`

GetScanIdOk returns a tuple with the ScanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanId

`func (o *ContainerImageInQuarantineData) SetScanId(v string)`

SetScanId sets ScanId field to given value.

### HasScanId

`func (o *ContainerImageInQuarantineData) HasScanId() bool`

HasScanId returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ContainerImageInQuarantineData) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ContainerImageInQuarantineData) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ContainerImageInQuarantineData) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ContainerImageInQuarantineData) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


