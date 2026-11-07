# ApiHostedRepositoryComponentListDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Components** | Pointer to [**[]ApiHostedRepositoryComponentDTO**](ApiHostedRepositoryComponentDTO.md) |  | [optional] 
**HasNextPage** | Pointer to **bool** |  | [optional] 
**HasQueuedScans** | Pointer to **bool** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**RepositoryPublicId** | Pointer to **string** |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiHostedRepositoryComponentListDTO

`func NewApiHostedRepositoryComponentListDTO() *ApiHostedRepositoryComponentListDTO`

NewApiHostedRepositoryComponentListDTO instantiates a new ApiHostedRepositoryComponentListDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHostedRepositoryComponentListDTOWithDefaults

`func NewApiHostedRepositoryComponentListDTOWithDefaults() *ApiHostedRepositoryComponentListDTO`

NewApiHostedRepositoryComponentListDTOWithDefaults instantiates a new ApiHostedRepositoryComponentListDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponents

`func (o *ApiHostedRepositoryComponentListDTO) GetComponents() []ApiHostedRepositoryComponentDTO`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ApiHostedRepositoryComponentListDTO) GetComponentsOk() (*[]ApiHostedRepositoryComponentDTO, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ApiHostedRepositoryComponentListDTO) SetComponents(v []ApiHostedRepositoryComponentDTO)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ApiHostedRepositoryComponentListDTO) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetHasNextPage

`func (o *ApiHostedRepositoryComponentListDTO) GetHasNextPage() bool`

GetHasNextPage returns the HasNextPage field if non-nil, zero value otherwise.

### GetHasNextPageOk

`func (o *ApiHostedRepositoryComponentListDTO) GetHasNextPageOk() (*bool, bool)`

GetHasNextPageOk returns a tuple with the HasNextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNextPage

`func (o *ApiHostedRepositoryComponentListDTO) SetHasNextPage(v bool)`

SetHasNextPage sets HasNextPage field to given value.

### HasHasNextPage

`func (o *ApiHostedRepositoryComponentListDTO) HasHasNextPage() bool`

HasHasNextPage returns a boolean if a field has been set.

### GetHasQueuedScans

`func (o *ApiHostedRepositoryComponentListDTO) GetHasQueuedScans() bool`

GetHasQueuedScans returns the HasQueuedScans field if non-nil, zero value otherwise.

### GetHasQueuedScansOk

`func (o *ApiHostedRepositoryComponentListDTO) GetHasQueuedScansOk() (*bool, bool)`

GetHasQueuedScansOk returns a tuple with the HasQueuedScans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasQueuedScans

`func (o *ApiHostedRepositoryComponentListDTO) SetHasQueuedScans(v bool)`

SetHasQueuedScans sets HasQueuedScans field to given value.

### HasHasQueuedScans

`func (o *ApiHostedRepositoryComponentListDTO) HasHasQueuedScans() bool`

HasHasQueuedScans returns a boolean if a field has been set.

### GetPage

`func (o *ApiHostedRepositoryComponentListDTO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ApiHostedRepositoryComponentListDTO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ApiHostedRepositoryComponentListDTO) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *ApiHostedRepositoryComponentListDTO) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *ApiHostedRepositoryComponentListDTO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ApiHostedRepositoryComponentListDTO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ApiHostedRepositoryComponentListDTO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ApiHostedRepositoryComponentListDTO) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetRepositoryPublicId

`func (o *ApiHostedRepositoryComponentListDTO) GetRepositoryPublicId() string`

GetRepositoryPublicId returns the RepositoryPublicId field if non-nil, zero value otherwise.

### GetRepositoryPublicIdOk

`func (o *ApiHostedRepositoryComponentListDTO) GetRepositoryPublicIdOk() (*string, bool)`

GetRepositoryPublicIdOk returns a tuple with the RepositoryPublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryPublicId

`func (o *ApiHostedRepositoryComponentListDTO) SetRepositoryPublicId(v string)`

SetRepositoryPublicId sets RepositoryPublicId field to given value.

### HasRepositoryPublicId

`func (o *ApiHostedRepositoryComponentListDTO) HasRepositoryPublicId() bool`

HasRepositoryPublicId returns a boolean if a field has been set.

### GetTotalCount

`func (o *ApiHostedRepositoryComponentListDTO) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ApiHostedRepositoryComponentListDTO) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ApiHostedRepositoryComponentListDTO) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ApiHostedRepositoryComponentListDTO) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


