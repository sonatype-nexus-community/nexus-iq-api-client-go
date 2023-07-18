# ComponentObligationAttributionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdatedByUsername** | Pointer to **string** |  | [optional] 
**ObligationName** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**PackageUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewComponentObligationAttributionDTO

`func NewComponentObligationAttributionDTO() *ComponentObligationAttributionDTO`

NewComponentObligationAttributionDTO instantiates a new ComponentObligationAttributionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentObligationAttributionDTOWithDefaults

`func NewComponentObligationAttributionDTOWithDefaults() *ComponentObligationAttributionDTO`

NewComponentObligationAttributionDTOWithDefaults instantiates a new ComponentObligationAttributionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentIdentifier

`func (o *ComponentObligationAttributionDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ComponentObligationAttributionDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ComponentObligationAttributionDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ComponentObligationAttributionDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetContent

`func (o *ComponentObligationAttributionDTO) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ComponentObligationAttributionDTO) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ComponentObligationAttributionDTO) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ComponentObligationAttributionDTO) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetId

`func (o *ComponentObligationAttributionDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentObligationAttributionDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentObligationAttributionDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentObligationAttributionDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *ComponentObligationAttributionDTO) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *ComponentObligationAttributionDTO) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *ComponentObligationAttributionDTO) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *ComponentObligationAttributionDTO) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetLastUpdatedByUsername

`func (o *ComponentObligationAttributionDTO) GetLastUpdatedByUsername() string`

GetLastUpdatedByUsername returns the LastUpdatedByUsername field if non-nil, zero value otherwise.

### GetLastUpdatedByUsernameOk

`func (o *ComponentObligationAttributionDTO) GetLastUpdatedByUsernameOk() (*string, bool)`

GetLastUpdatedByUsernameOk returns a tuple with the LastUpdatedByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedByUsername

`func (o *ComponentObligationAttributionDTO) SetLastUpdatedByUsername(v string)`

SetLastUpdatedByUsername sets LastUpdatedByUsername field to given value.

### HasLastUpdatedByUsername

`func (o *ComponentObligationAttributionDTO) HasLastUpdatedByUsername() bool`

HasLastUpdatedByUsername returns a boolean if a field has been set.

### GetObligationName

`func (o *ComponentObligationAttributionDTO) GetObligationName() string`

GetObligationName returns the ObligationName field if non-nil, zero value otherwise.

### GetObligationNameOk

`func (o *ComponentObligationAttributionDTO) GetObligationNameOk() (*string, bool)`

GetObligationNameOk returns a tuple with the ObligationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObligationName

`func (o *ComponentObligationAttributionDTO) SetObligationName(v string)`

SetObligationName sets ObligationName field to given value.

### HasObligationName

`func (o *ComponentObligationAttributionDTO) HasObligationName() bool`

HasObligationName returns a boolean if a field has been set.

### GetOwnerId

`func (o *ComponentObligationAttributionDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ComponentObligationAttributionDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ComponentObligationAttributionDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ComponentObligationAttributionDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPackageUrl

`func (o *ComponentObligationAttributionDTO) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *ComponentObligationAttributionDTO) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *ComponentObligationAttributionDTO) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *ComponentObligationAttributionDTO) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


