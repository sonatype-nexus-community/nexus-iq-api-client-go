# ApiWaiverOptionsDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | Reason for waiving the violation(s). Must be non-blank. | 
**ExpireWhenRemediationAvailable** | Pointer to **bool** | If true, the waiver will automatically expire when a remediation becomes available. Can only be set to true when matcherStrategy is EXACT_COMPONENT. | [optional] 
**ExpiryTime** | Pointer to **time.Time** | Optional expiration date/time for the waiver in ISO 8601 format. Must be in the future if provided. | [optional] 
**MatcherStrategy** | **string** | Component matching strategy. For Firewall bulk waivers, only EXACT_COMPONENT and ALL_VERSIONS are supported. | 
**WaiverReasonId** | Pointer to **string** | Optional reference to a pre-defined waiver reason ID | [optional] 

## Methods

### NewApiWaiverOptionsDTO

`func NewApiWaiverOptionsDTO(comment string, matcherStrategy string, ) *ApiWaiverOptionsDTO`

NewApiWaiverOptionsDTO instantiates a new ApiWaiverOptionsDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWaiverOptionsDTOWithDefaults

`func NewApiWaiverOptionsDTOWithDefaults() *ApiWaiverOptionsDTO`

NewApiWaiverOptionsDTOWithDefaults instantiates a new ApiWaiverOptionsDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


