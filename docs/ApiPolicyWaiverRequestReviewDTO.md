# ApiPolicyWaiverRequestReviewDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**ExpireWhenRemediationAvailable** | Pointer to **bool** |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** |  | [optional] 
**MatcherStrategy** | Pointer to **string** |  | [optional] 
**RejectionReason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**WaiverReasonId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiPolicyWaiverRequestReviewDTO

`func NewApiPolicyWaiverRequestReviewDTO() *ApiPolicyWaiverRequestReviewDTO`

NewApiPolicyWaiverRequestReviewDTO instantiates a new ApiPolicyWaiverRequestReviewDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyWaiverRequestReviewDTOWithDefaults

`func NewApiPolicyWaiverRequestReviewDTOWithDefaults() *ApiPolicyWaiverRequestReviewDTO`

NewApiPolicyWaiverRequestReviewDTOWithDefaults instantiates a new ApiPolicyWaiverRequestReviewDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ApiPolicyWaiverRequestReviewDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiPolicyWaiverRequestReviewDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiPolicyWaiverRequestReviewDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiPolicyWaiverRequestReviewDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpireWhenRemediationAvailable

`func (o *ApiPolicyWaiverRequestReviewDTO) GetExpireWhenRemediationAvailable() bool`

GetExpireWhenRemediationAvailable returns the ExpireWhenRemediationAvailable field if non-nil, zero value otherwise.

### GetExpireWhenRemediationAvailableOk

`func (o *ApiPolicyWaiverRequestReviewDTO) GetExpireWhenRemediationAvailableOk() (*bool, bool)`

GetExpireWhenRemediationAvailableOk returns a tuple with the ExpireWhenRemediationAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWhenRemediationAvailable

`func (o *ApiPolicyWaiverRequestReviewDTO) SetExpireWhenRemediationAvailable(v bool)`

SetExpireWhenRemediationAvailable sets ExpireWhenRemediationAvailable field to given value.

### HasExpireWhenRemediationAvailable

`func (o *ApiPolicyWaiverRequestReviewDTO) HasExpireWhenRemediationAvailable() bool`

HasExpireWhenRemediationAvailable returns a boolean if a field has been set.

### GetExpiryTime

`func (o *ApiPolicyWaiverRequestReviewDTO) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ApiPolicyWaiverRequestReviewDTO) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ApiPolicyWaiverRequestReviewDTO) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *ApiPolicyWaiverRequestReviewDTO) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetMatcherStrategy

`func (o *ApiPolicyWaiverRequestReviewDTO) GetMatcherStrategy() string`

GetMatcherStrategy returns the MatcherStrategy field if non-nil, zero value otherwise.

### GetMatcherStrategyOk

`func (o *ApiPolicyWaiverRequestReviewDTO) GetMatcherStrategyOk() (*string, bool)`

GetMatcherStrategyOk returns a tuple with the MatcherStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcherStrategy

`func (o *ApiPolicyWaiverRequestReviewDTO) SetMatcherStrategy(v string)`

SetMatcherStrategy sets MatcherStrategy field to given value.

### HasMatcherStrategy

`func (o *ApiPolicyWaiverRequestReviewDTO) HasMatcherStrategy() bool`

HasMatcherStrategy returns a boolean if a field has been set.

### GetRejectionReason

`func (o *ApiPolicyWaiverRequestReviewDTO) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *ApiPolicyWaiverRequestReviewDTO) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *ApiPolicyWaiverRequestReviewDTO) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *ApiPolicyWaiverRequestReviewDTO) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetStatus

`func (o *ApiPolicyWaiverRequestReviewDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiPolicyWaiverRequestReviewDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiPolicyWaiverRequestReviewDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiPolicyWaiverRequestReviewDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWaiverReasonId

`func (o *ApiPolicyWaiverRequestReviewDTO) GetWaiverReasonId() string`

GetWaiverReasonId returns the WaiverReasonId field if non-nil, zero value otherwise.

### GetWaiverReasonIdOk

`func (o *ApiPolicyWaiverRequestReviewDTO) GetWaiverReasonIdOk() (*string, bool)`

GetWaiverReasonIdOk returns a tuple with the WaiverReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiverReasonId

`func (o *ApiPolicyWaiverRequestReviewDTO) SetWaiverReasonId(v string)`

SetWaiverReasonId sets WaiverReasonId field to given value.

### HasWaiverReasonId

`func (o *ApiPolicyWaiverRequestReviewDTO) HasWaiverReasonId() bool`

HasWaiverReasonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


