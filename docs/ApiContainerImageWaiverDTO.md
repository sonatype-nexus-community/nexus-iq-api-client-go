# ApiContainerImageWaiverDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** |  | [optional] 
**WaiverReasonId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiContainerImageWaiverDTO

`func NewApiContainerImageWaiverDTO() *ApiContainerImageWaiverDTO`

NewApiContainerImageWaiverDTO instantiates a new ApiContainerImageWaiverDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiContainerImageWaiverDTOWithDefaults

`func NewApiContainerImageWaiverDTOWithDefaults() *ApiContainerImageWaiverDTO`

NewApiContainerImageWaiverDTOWithDefaults instantiates a new ApiContainerImageWaiverDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ApiContainerImageWaiverDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiContainerImageWaiverDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiContainerImageWaiverDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiContainerImageWaiverDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpiryTime

`func (o *ApiContainerImageWaiverDTO) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ApiContainerImageWaiverDTO) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ApiContainerImageWaiverDTO) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *ApiContainerImageWaiverDTO) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetWaiverReasonId

`func (o *ApiContainerImageWaiverDTO) GetWaiverReasonId() string`

GetWaiverReasonId returns the WaiverReasonId field if non-nil, zero value otherwise.

### GetWaiverReasonIdOk

`func (o *ApiContainerImageWaiverDTO) GetWaiverReasonIdOk() (*string, bool)`

GetWaiverReasonIdOk returns a tuple with the WaiverReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiverReasonId

`func (o *ApiContainerImageWaiverDTO) SetWaiverReasonId(v string)`

SetWaiverReasonId sets WaiverReasonId field to given value.

### HasWaiverReasonId

`func (o *ApiContainerImageWaiverDTO) HasWaiverReasonId() bool`

HasWaiverReasonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


