# ApiWaiverOptionsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyToAllComponents** | Pointer to **bool** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**ExpireWhenRemediationAvailable** | Pointer to **bool** |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** |  | [optional] 
**MatcherStrategy** | Pointer to **string** |  | [optional] 
**WaiverReasonId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiWaiverOptionsDTO

`func NewApiWaiverOptionsDTO() *ApiWaiverOptionsDTO`

NewApiWaiverOptionsDTO instantiates a new ApiWaiverOptionsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWaiverOptionsDTOWithDefaults

`func NewApiWaiverOptionsDTOWithDefaults() *ApiWaiverOptionsDTO`

NewApiWaiverOptionsDTOWithDefaults instantiates a new ApiWaiverOptionsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyToAllComponents

`func (o *ApiWaiverOptionsDTO) GetApplyToAllComponents() bool`

GetApplyToAllComponents returns the ApplyToAllComponents field if non-nil, zero value otherwise.

### GetApplyToAllComponentsOk

`func (o *ApiWaiverOptionsDTO) GetApplyToAllComponentsOk() (*bool, bool)`

GetApplyToAllComponentsOk returns a tuple with the ApplyToAllComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToAllComponents

`func (o *ApiWaiverOptionsDTO) SetApplyToAllComponents(v bool)`

SetApplyToAllComponents sets ApplyToAllComponents field to given value.

### HasApplyToAllComponents

`func (o *ApiWaiverOptionsDTO) HasApplyToAllComponents() bool`

HasApplyToAllComponents returns a boolean if a field has been set.

### GetComment

`func (o *ApiWaiverOptionsDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiWaiverOptionsDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiWaiverOptionsDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiWaiverOptionsDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpireWhenRemediationAvailable

`func (o *ApiWaiverOptionsDTO) GetExpireWhenRemediationAvailable() bool`

GetExpireWhenRemediationAvailable returns the ExpireWhenRemediationAvailable field if non-nil, zero value otherwise.

### GetExpireWhenRemediationAvailableOk

`func (o *ApiWaiverOptionsDTO) GetExpireWhenRemediationAvailableOk() (*bool, bool)`

GetExpireWhenRemediationAvailableOk returns a tuple with the ExpireWhenRemediationAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireWhenRemediationAvailable

`func (o *ApiWaiverOptionsDTO) SetExpireWhenRemediationAvailable(v bool)`

SetExpireWhenRemediationAvailable sets ExpireWhenRemediationAvailable field to given value.

### HasExpireWhenRemediationAvailable

`func (o *ApiWaiverOptionsDTO) HasExpireWhenRemediationAvailable() bool`

HasExpireWhenRemediationAvailable returns a boolean if a field has been set.

### GetExpiryTime

`func (o *ApiWaiverOptionsDTO) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ApiWaiverOptionsDTO) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ApiWaiverOptionsDTO) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *ApiWaiverOptionsDTO) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetMatcherStrategy

`func (o *ApiWaiverOptionsDTO) GetMatcherStrategy() string`

GetMatcherStrategy returns the MatcherStrategy field if non-nil, zero value otherwise.

### GetMatcherStrategyOk

`func (o *ApiWaiverOptionsDTO) GetMatcherStrategyOk() (*string, bool)`

GetMatcherStrategyOk returns a tuple with the MatcherStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcherStrategy

`func (o *ApiWaiverOptionsDTO) SetMatcherStrategy(v string)`

SetMatcherStrategy sets MatcherStrategy field to given value.

### HasMatcherStrategy

`func (o *ApiWaiverOptionsDTO) HasMatcherStrategy() bool`

HasMatcherStrategy returns a boolean if a field has been set.

### GetWaiverReasonId

`func (o *ApiWaiverOptionsDTO) GetWaiverReasonId() string`

GetWaiverReasonId returns the WaiverReasonId field if non-nil, zero value otherwise.

### GetWaiverReasonIdOk

`func (o *ApiWaiverOptionsDTO) GetWaiverReasonIdOk() (*string, bool)`

GetWaiverReasonIdOk returns a tuple with the WaiverReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiverReasonId

`func (o *ApiWaiverOptionsDTO) SetWaiverReasonId(v string)`

SetWaiverReasonId sets WaiverReasonId field to given value.

### HasWaiverReasonId

`func (o *ApiWaiverOptionsDTO) HasWaiverReasonId() bool`

HasWaiverReasonId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


