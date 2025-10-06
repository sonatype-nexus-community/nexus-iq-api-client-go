# ApiFirewallComponentDTOResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | Current page number | [optional] 
**PageCount** | Pointer to **int64** | Total number of pages | [optional] 
**PageSize** | Pointer to **int32** | Number of items per page | [optional] 
**Results** | Pointer to [**[]ApiFirewallComponentDTO**](ApiFirewallComponentDTO.md) | List of items for the current page | [optional] 
**Total** | Pointer to **int64** | Total number of items | [optional] 

## Methods

### NewApiFirewallComponentDTOResult

`func NewApiFirewallComponentDTOResult() *ApiFirewallComponentDTOResult`

NewApiFirewallComponentDTOResult instantiates a new ApiFirewallComponentDTOResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiFirewallComponentDTOResultWithDefaults

`func NewApiFirewallComponentDTOResultWithDefaults() *ApiFirewallComponentDTOResult`

NewApiFirewallComponentDTOResultWithDefaults instantiates a new ApiFirewallComponentDTOResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ApiFirewallComponentDTOResult) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ApiFirewallComponentDTOResult) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ApiFirewallComponentDTOResult) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ApiFirewallComponentDTOResult) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageCount

`func (o *ApiFirewallComponentDTOResult) GetPageCount() int64`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *ApiFirewallComponentDTOResult) GetPageCountOk() (*int64, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *ApiFirewallComponentDTOResult) SetPageCount(v int64)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *ApiFirewallComponentDTOResult) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetPageSize

`func (o *ApiFirewallComponentDTOResult) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ApiFirewallComponentDTOResult) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ApiFirewallComponentDTOResult) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ApiFirewallComponentDTOResult) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetResults

`func (o *ApiFirewallComponentDTOResult) GetResults() []ApiFirewallComponentDTO`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ApiFirewallComponentDTOResult) GetResultsOk() (*[]ApiFirewallComponentDTO, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ApiFirewallComponentDTOResult) SetResults(v []ApiFirewallComponentDTO)`

SetResults sets Results field to given value.

### HasResults

`func (o *ApiFirewallComponentDTOResult) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotal

`func (o *ApiFirewallComponentDTOResult) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ApiFirewallComponentDTOResult) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ApiFirewallComponentDTOResult) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ApiFirewallComponentDTOResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


