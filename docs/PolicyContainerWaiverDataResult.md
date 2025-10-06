# PolicyContainerWaiverDataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | Current page number | [optional] 
**PageCount** | Pointer to **int64** | Total number of pages | [optional] 
**PageSize** | Pointer to **int32** | Number of items per page | [optional] 
**Results** | Pointer to [**[]PolicyContainerWaiverData**](PolicyContainerWaiverData.md) | List of items for the current page | [optional] 
**Total** | Pointer to **int64** | Total number of items | [optional] 

## Methods

### NewPolicyContainerWaiverDataResult

`func NewPolicyContainerWaiverDataResult() *PolicyContainerWaiverDataResult`

NewPolicyContainerWaiverDataResult instantiates a new PolicyContainerWaiverDataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyContainerWaiverDataResultWithDefaults

`func NewPolicyContainerWaiverDataResultWithDefaults() *PolicyContainerWaiverDataResult`

NewPolicyContainerWaiverDataResultWithDefaults instantiates a new PolicyContainerWaiverDataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PolicyContainerWaiverDataResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PolicyContainerWaiverDataResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PolicyContainerWaiverDataResult) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *PolicyContainerWaiverDataResult) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageCount

`func (o *PolicyContainerWaiverDataResult) GetPageCount() int64`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *PolicyContainerWaiverDataResult) GetPageCountOk() (*int64, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *PolicyContainerWaiverDataResult) SetPageCount(v int64)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *PolicyContainerWaiverDataResult) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetPageSize

`func (o *PolicyContainerWaiverDataResult) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *PolicyContainerWaiverDataResult) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *PolicyContainerWaiverDataResult) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *PolicyContainerWaiverDataResult) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetResults

`func (o *PolicyContainerWaiverDataResult) GetResults() []PolicyContainerWaiverData`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PolicyContainerWaiverDataResult) GetResultsOk() (*[]PolicyContainerWaiverData, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PolicyContainerWaiverDataResult) SetResults(v []PolicyContainerWaiverData)`

SetResults sets Results field to given value.

### HasResults

`func (o *PolicyContainerWaiverDataResult) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotal

`func (o *PolicyContainerWaiverDataResult) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PolicyContainerWaiverDataResult) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PolicyContainerWaiverDataResult) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PolicyContainerWaiverDataResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


