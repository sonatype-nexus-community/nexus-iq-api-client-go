# ApiHashComponentIdentifierDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**PackageUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewApiHashComponentIdentifierDTO

`func NewApiHashComponentIdentifierDTO() *ApiHashComponentIdentifierDTO`

NewApiHashComponentIdentifierDTO instantiates a new ApiHashComponentIdentifierDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHashComponentIdentifierDTOWithDefaults

`func NewApiHashComponentIdentifierDTOWithDefaults() *ApiHashComponentIdentifierDTO`

NewApiHashComponentIdentifierDTOWithDefaults instantiates a new ApiHashComponentIdentifierDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ApiHashComponentIdentifierDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiHashComponentIdentifierDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiHashComponentIdentifierDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiHashComponentIdentifierDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *ApiHashComponentIdentifierDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiHashComponentIdentifierDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiHashComponentIdentifierDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiHashComponentIdentifierDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetCreateTime

`func (o *ApiHashComponentIdentifierDTO) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ApiHashComponentIdentifierDTO) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ApiHashComponentIdentifierDTO) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ApiHashComponentIdentifierDTO) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetHash

`func (o *ApiHashComponentIdentifierDTO) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiHashComponentIdentifierDTO) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiHashComponentIdentifierDTO) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiHashComponentIdentifierDTO) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetPackageUrl

`func (o *ApiHashComponentIdentifierDTO) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *ApiHashComponentIdentifierDTO) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *ApiHashComponentIdentifierDTO) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *ApiHashComponentIdentifierDTO) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


