# RenewWaiversRequestDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**NewExpiryTime** | Pointer to **time.Time** |  | [optional] 
**ReasonId** | Pointer to **string** |  | [optional] 
**WaiverIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRenewWaiversRequestDTO

`func NewRenewWaiversRequestDTO() *RenewWaiversRequestDTO`

NewRenewWaiversRequestDTO instantiates a new RenewWaiversRequestDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenewWaiversRequestDTOWithDefaults

`func NewRenewWaiversRequestDTOWithDefaults() *RenewWaiversRequestDTO`

NewRenewWaiversRequestDTOWithDefaults instantiates a new RenewWaiversRequestDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *RenewWaiversRequestDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RenewWaiversRequestDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RenewWaiversRequestDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RenewWaiversRequestDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetNewExpiryTime

`func (o *RenewWaiversRequestDTO) GetNewExpiryTime() time.Time`

GetNewExpiryTime returns the NewExpiryTime field if non-nil, zero value otherwise.

### GetNewExpiryTimeOk

`func (o *RenewWaiversRequestDTO) GetNewExpiryTimeOk() (*time.Time, bool)`

GetNewExpiryTimeOk returns a tuple with the NewExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewExpiryTime

`func (o *RenewWaiversRequestDTO) SetNewExpiryTime(v time.Time)`

SetNewExpiryTime sets NewExpiryTime field to given value.

### HasNewExpiryTime

`func (o *RenewWaiversRequestDTO) HasNewExpiryTime() bool`

HasNewExpiryTime returns a boolean if a field has been set.

### GetReasonId

`func (o *RenewWaiversRequestDTO) GetReasonId() string`

GetReasonId returns the ReasonId field if non-nil, zero value otherwise.

### GetReasonIdOk

`func (o *RenewWaiversRequestDTO) GetReasonIdOk() (*string, bool)`

GetReasonIdOk returns a tuple with the ReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonId

`func (o *RenewWaiversRequestDTO) SetReasonId(v string)`

SetReasonId sets ReasonId field to given value.

### HasReasonId

`func (o *RenewWaiversRequestDTO) HasReasonId() bool`

HasReasonId returns a boolean if a field has been set.

### GetWaiverIds

`func (o *RenewWaiversRequestDTO) GetWaiverIds() []string`

GetWaiverIds returns the WaiverIds field if non-nil, zero value otherwise.

### GetWaiverIdsOk

`func (o *RenewWaiversRequestDTO) GetWaiverIdsOk() (*[]string, bool)`

GetWaiverIdsOk returns a tuple with the WaiverIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiverIds

`func (o *RenewWaiversRequestDTO) SetWaiverIds(v []string)`

SetWaiverIds sets WaiverIds field to given value.

### HasWaiverIds

`func (o *RenewWaiversRequestDTO) HasWaiverIds() bool`

HasWaiverIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


