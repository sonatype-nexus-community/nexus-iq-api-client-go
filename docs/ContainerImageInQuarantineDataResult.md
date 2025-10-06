# ContainerImageInQuarantineDataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | Current page number | [optional] 
**PageCount** | Pointer to **int64** | Total number of pages | [optional] 
**PageSize** | Pointer to **int32** | Number of items per page | [optional] 
**Results** | Pointer to [**[]ContainerImageInQuarantineData**](ContainerImageInQuarantineData.md) | List of items for the current page | [optional] 
**Total** | Pointer to **int64** | Total number of items | [optional] 

## Methods

### NewContainerImageInQuarantineDataResult

`func NewContainerImageInQuarantineDataResult() *ContainerImageInQuarantineDataResult`

NewContainerImageInQuarantineDataResult instantiates a new ContainerImageInQuarantineDataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerImageInQuarantineDataResultWithDefaults

`func NewContainerImageInQuarantineDataResultWithDefaults() *ContainerImageInQuarantineDataResult`

NewContainerImageInQuarantineDataResultWithDefaults instantiates a new ContainerImageInQuarantineDataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ContainerImageInQuarantineDataResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ContainerImageInQuarantineDataResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ContainerImageInQuarantineDataResult) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ContainerImageInQuarantineDataResult) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageCount

`func (o *ContainerImageInQuarantineDataResult) GetPageCount() int64`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *ContainerImageInQuarantineDataResult) GetPageCountOk() (*int64, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *ContainerImageInQuarantineDataResult) SetPageCount(v int64)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *ContainerImageInQuarantineDataResult) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetPageSize

`func (o *ContainerImageInQuarantineDataResult) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ContainerImageInQuarantineDataResult) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ContainerImageInQuarantineDataResult) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ContainerImageInQuarantineDataResult) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetResults

`func (o *ContainerImageInQuarantineDataResult) GetResults() []ContainerImageInQuarantineData`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ContainerImageInQuarantineDataResult) GetResultsOk() (*[]ContainerImageInQuarantineData, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ContainerImageInQuarantineDataResult) SetResults(v []ContainerImageInQuarantineData)`

SetResults sets Results field to given value.

### HasResults

`func (o *ContainerImageInQuarantineDataResult) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotal

`func (o *ContainerImageInQuarantineDataResult) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ContainerImageInQuarantineDataResult) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ContainerImageInQuarantineDataResult) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ContainerImageInQuarantineDataResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


