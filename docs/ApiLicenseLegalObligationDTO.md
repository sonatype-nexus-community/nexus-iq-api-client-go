# ApiLicenseLegalObligationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**LastUpdatedByUsername** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**PackageUrl** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewApiLicenseLegalObligationDTO

`func NewApiLicenseLegalObligationDTO() *ApiLicenseLegalObligationDTO`

NewApiLicenseLegalObligationDTO instantiates a new ApiLicenseLegalObligationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLicenseLegalObligationDTOWithDefaults

`func NewApiLicenseLegalObligationDTOWithDefaults() *ApiLicenseLegalObligationDTO`

NewApiLicenseLegalObligationDTOWithDefaults instantiates a new ApiLicenseLegalObligationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ApiLicenseLegalObligationDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiLicenseLegalObligationDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiLicenseLegalObligationDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiLicenseLegalObligationDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *ApiLicenseLegalObligationDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiLicenseLegalObligationDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiLicenseLegalObligationDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiLicenseLegalObligationDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetId

`func (o *ApiLicenseLegalObligationDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiLicenseLegalObligationDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiLicenseLegalObligationDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiLicenseLegalObligationDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *ApiLicenseLegalObligationDTO) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *ApiLicenseLegalObligationDTO) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *ApiLicenseLegalObligationDTO) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *ApiLicenseLegalObligationDTO) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetLastUpdatedByUsername

`func (o *ApiLicenseLegalObligationDTO) GetLastUpdatedByUsername() string`

GetLastUpdatedByUsername returns the LastUpdatedByUsername field if non-nil, zero value otherwise.

### GetLastUpdatedByUsernameOk

`func (o *ApiLicenseLegalObligationDTO) GetLastUpdatedByUsernameOk() (*string, bool)`

GetLastUpdatedByUsernameOk returns a tuple with the LastUpdatedByUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedByUsername

`func (o *ApiLicenseLegalObligationDTO) SetLastUpdatedByUsername(v string)`

SetLastUpdatedByUsername sets LastUpdatedByUsername field to given value.

### HasLastUpdatedByUsername

`func (o *ApiLicenseLegalObligationDTO) HasLastUpdatedByUsername() bool`

HasLastUpdatedByUsername returns a boolean if a field has been set.

### GetName

`func (o *ApiLicenseLegalObligationDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiLicenseLegalObligationDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiLicenseLegalObligationDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiLicenseLegalObligationDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwnerId

`func (o *ApiLicenseLegalObligationDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ApiLicenseLegalObligationDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ApiLicenseLegalObligationDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ApiLicenseLegalObligationDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPackageUrl

`func (o *ApiLicenseLegalObligationDTO) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *ApiLicenseLegalObligationDTO) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *ApiLicenseLegalObligationDTO) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *ApiLicenseLegalObligationDTO) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.

### GetStatus

`func (o *ApiLicenseLegalObligationDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiLicenseLegalObligationDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiLicenseLegalObligationDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiLicenseLegalObligationDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


