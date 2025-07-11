# ApiLicenseOverrideDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LicenseIds** | Pointer to **[]string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewApiLicenseOverrideDTO

`func NewApiLicenseOverrideDTO() *ApiLicenseOverrideDTO`

NewApiLicenseOverrideDTO instantiates a new ApiLicenseOverrideDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLicenseOverrideDTOWithDefaults

`func NewApiLicenseOverrideDTOWithDefaults() *ApiLicenseOverrideDTO`

NewApiLicenseOverrideDTOWithDefaults instantiates a new ApiLicenseOverrideDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ApiLicenseOverrideDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiLicenseOverrideDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiLicenseOverrideDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiLicenseOverrideDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *ApiLicenseOverrideDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiLicenseOverrideDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiLicenseOverrideDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiLicenseOverrideDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetId

`func (o *ApiLicenseOverrideDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiLicenseOverrideDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiLicenseOverrideDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiLicenseOverrideDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicenseIds

`func (o *ApiLicenseOverrideDTO) GetLicenseIds() []string`

GetLicenseIds returns the LicenseIds field if non-nil, zero value otherwise.

### GetLicenseIdsOk

`func (o *ApiLicenseOverrideDTO) GetLicenseIdsOk() (*[]string, bool)`

GetLicenseIdsOk returns a tuple with the LicenseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseIds

`func (o *ApiLicenseOverrideDTO) SetLicenseIds(v []string)`

SetLicenseIds sets LicenseIds field to given value.

### HasLicenseIds

`func (o *ApiLicenseOverrideDTO) HasLicenseIds() bool`

HasLicenseIds returns a boolean if a field has been set.

### GetOwnerId

`func (o *ApiLicenseOverrideDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ApiLicenseOverrideDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ApiLicenseOverrideDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ApiLicenseOverrideDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetStatus

`func (o *ApiLicenseOverrideDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiLicenseOverrideDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiLicenseOverrideDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiLicenseOverrideDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


