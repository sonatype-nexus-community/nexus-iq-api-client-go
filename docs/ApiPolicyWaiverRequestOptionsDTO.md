# ApiPolicyWaiverRequestOptionsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**ExpireWhenRemediationAvailable** | Pointer to **bool** |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** |  | [optional] 
**MatcherStrategy** | Pointer to **string** |  | [optional] 
**NoteToReviewer** | Pointer to **string** |  | [optional] 
**WaiverReasonId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiPolicyWaiverRequestOptionsDTO

`func NewApiPolicyWaiverRequestOptionsDTO() *ApiPolicyWaiverRequestOptionsDTO`

NewApiPolicyWaiverRequestOptionsDTO instantiates a new ApiPolicyWaiverRequestOptionsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPolicyWaiverRequestOptionsDTOWithDefaults

`func NewApiPolicyWaiverRequestOptionsDTOWithDefaults() *ApiPolicyWaiverRequestOptionsDTO`

NewApiPolicyWaiverRequestOptionsDTOWithDefaults instantiates a new ApiPolicyWaiverRequestOptionsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ApiPolicyWaiverRequestOptionsDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiPolicyWaiverRequestOptionsDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiPolicyWaiverRequestOptionsDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiPolicyWaiverRequestOptionsDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpireWhenRemediationAvailable

`func (o *ApiPolicyWaiverRequestOptionsDTO) GetExpireWhenRemediationAvailable() bool`

GetExpireWhenRemediationAvailable returns the ExpireWhenRemediationAvailable field if non-nil, zero value otherwise.

### GetExpireWhenRemediationAvailableOk

`func (o *ApiPolicyWaiverRequestOptionsDTO) GetExpireWhenRemediationAvailableOk() (*bool, bool)`

GetExpireWhenRemediationAvailableOk returns a tuple with the ExpireWhenRemediationAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWhenRemediationAvailable

`func (o *ApiPolicyWaiverRequestOptionsDTO) SetExpireWhenRemediationAvailable(v bool)`

SetExpireWhenRemediationAvailable sets ExpireWhenRemediationAvailable field to given value.

### HasExpireWhenRemediationAvailable

`func (o *ApiPolicyWaiverRequestOptionsDTO) HasExpireWhenRemediationAvailable() bool`

HasExpireWhenRemediationAvailable returns a boolean if a field has been set.

### GetExpiryTime

`func (o *ApiPolicyWaiverRequestOptionsDTO) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ApiPolicyWaiverRequestOptionsDTO) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ApiPolicyWaiverRequestOptionsDTO) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *ApiPolicyWaiverRequestOptionsDTO) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetMatcherStrategy

`func (o *ApiPolicyWaiverRequestOptionsDTO) GetMatcherStrategy() string`

GetMatcherStrategy returns the MatcherStrategy field if non-nil, zero value otherwise.

### GetMatcherStrategyOk

`func (o *ApiPolicyWaiverRequestOptionsDTO) GetMatcherStrategyOk() (*string, bool)`

GetMatcherStrategyOk returns a tuple with the MatcherStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcherStrategy

`func (o *ApiPolicyWaiverRequestOptionsDTO) SetMatcherStrategy(v string)`

SetMatcherStrategy sets MatcherStrategy field to given value.

### HasMatcherStrategy

`func (o *ApiPolicyWaiverRequestOptionsDTO) HasMatcherStrategy() bool`

HasMatcherStrategy returns a boolean if a field has been set.

### GetNoteToReviewer

`func (o *ApiPolicyWaiverRequestOptionsDTO) GetNoteToReviewer() string`

GetNoteToReviewer returns the NoteToReviewer field if non-nil, zero value otherwise.

### GetNoteToReviewerOk

`func (o *ApiPolicyWaiverRequestOptionsDTO) GetNoteToReviewerOk() (*string, bool)`

GetNoteToReviewerOk returns a tuple with the NoteToReviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteToReviewer

`func (o *ApiPolicyWaiverRequestOptionsDTO) SetNoteToReviewer(v string)`

SetNoteToReviewer sets NoteToReviewer field to given value.

### HasNoteToReviewer

`func (o *ApiPolicyWaiverRequestOptionsDTO) HasNoteToReviewer() bool`

HasNoteToReviewer returns a boolean if a field has been set.

### GetWaiverReasonId

`func (o *ApiPolicyWaiverRequestOptionsDTO) GetWaiverReasonId() string`

GetWaiverReasonId returns the WaiverReasonId field if non-nil, zero value otherwise.

### GetWaiverReasonIdOk

`func (o *ApiPolicyWaiverRequestOptionsDTO) GetWaiverReasonIdOk() (*string, bool)`

GetWaiverReasonIdOk returns a tuple with the WaiverReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiverReasonId

`func (o *ApiPolicyWaiverRequestOptionsDTO) SetWaiverReasonId(v string)`

SetWaiverReasonId sets WaiverReasonId field to given value.

### HasWaiverReasonId

`func (o *ApiPolicyWaiverRequestOptionsDTO) HasWaiverReasonId() bool`

HasWaiverReasonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


