# SearchResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupingByDTOS** | Pointer to [**[]GroupingByDTO**](GroupingByDTO.md) |  | [optional] 
**IsExactTotalNumberOfHits** | Pointer to **bool** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**SearchQuery** | Pointer to **string** |  | [optional] 
**TotalNumberOfHits** | Pointer to **int32** |  | [optional] 

## Methods

### NewSearchResultDTO

`func NewSearchResultDTO() *SearchResultDTO`

NewSearchResultDTO instantiates a new SearchResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultDTOWithDefaults

`func NewSearchResultDTOWithDefaults() *SearchResultDTO`

NewSearchResultDTOWithDefaults instantiates a new SearchResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupingByDTOS

`func (o *SearchResultDTO) GetGroupingByDTOS() []GroupingByDTO`

GetGroupingByDTOS returns the GroupingByDTOS field if non-nil, zero value otherwise.

### GetGroupingByDTOSOk

`func (o *SearchResultDTO) GetGroupingByDTOSOk() (*[]GroupingByDTO, bool)`

GetGroupingByDTOSOk returns a tuple with the GroupingByDTOS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupingByDTOS

`func (o *SearchResultDTO) SetGroupingByDTOS(v []GroupingByDTO)`

SetGroupingByDTOS sets GroupingByDTOS field to given value.

### HasGroupingByDTOS

`func (o *SearchResultDTO) HasGroupingByDTOS() bool`

HasGroupingByDTOS returns a boolean if a field has been set.

### GetIsExactTotalNumberOfHits

`func (o *SearchResultDTO) GetIsExactTotalNumberOfHits() bool`

GetIsExactTotalNumberOfHits returns the IsExactTotalNumberOfHits field if non-nil, zero value otherwise.

### GetIsExactTotalNumberOfHitsOk

`func (o *SearchResultDTO) GetIsExactTotalNumberOfHitsOk() (*bool, bool)`

GetIsExactTotalNumberOfHitsOk returns a tuple with the IsExactTotalNumberOfHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExactTotalNumberOfHits

`func (o *SearchResultDTO) SetIsExactTotalNumberOfHits(v bool)`

SetIsExactTotalNumberOfHits sets IsExactTotalNumberOfHits field to given value.

### HasIsExactTotalNumberOfHits

`func (o *SearchResultDTO) HasIsExactTotalNumberOfHits() bool`

HasIsExactTotalNumberOfHits returns a boolean if a field has been set.

### GetPage

`func (o *SearchResultDTO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SearchResultDTO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SearchResultDTO) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *SearchResultDTO) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *SearchResultDTO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SearchResultDTO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SearchResultDTO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *SearchResultDTO) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSearchQuery

`func (o *SearchResultDTO) GetSearchQuery() string`

GetSearchQuery returns the SearchQuery field if non-nil, zero value otherwise.

### GetSearchQueryOk

`func (o *SearchResultDTO) GetSearchQueryOk() (*string, bool)`

GetSearchQueryOk returns a tuple with the SearchQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchQuery

`func (o *SearchResultDTO) SetSearchQuery(v string)`

SetSearchQuery sets SearchQuery field to given value.

### HasSearchQuery

`func (o *SearchResultDTO) HasSearchQuery() bool`

HasSearchQuery returns a boolean if a field has been set.

### GetTotalNumberOfHits

`func (o *SearchResultDTO) GetTotalNumberOfHits() int32`

GetTotalNumberOfHits returns the TotalNumberOfHits field if non-nil, zero value otherwise.

### GetTotalNumberOfHitsOk

`func (o *SearchResultDTO) GetTotalNumberOfHitsOk() (*int32, bool)`

GetTotalNumberOfHitsOk returns a tuple with the TotalNumberOfHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfHits

`func (o *SearchResultDTO) SetTotalNumberOfHits(v int32)`

SetTotalNumberOfHits sets TotalNumberOfHits field to given value.

### HasTotalNumberOfHits

`func (o *SearchResultDTO) HasTotalNumberOfHits() bool`

HasTotalNumberOfHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


