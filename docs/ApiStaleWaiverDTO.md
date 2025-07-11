# ApiStaleWaiverDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**ConstraintFacts** | Pointer to [**[]ApiConstraintFactDTO**](ApiConstraintFactDTO.md) |  | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**CreatorName** | Pointer to **string** |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyWaiverReasonId** | Pointer to **string** |  | [optional] 
**ReasonText** | Pointer to **string** |  | [optional] 
**ScopeOwnerId** | Pointer to **string** |  | [optional] 
**ScopeOwnerName** | Pointer to **string** |  | [optional] 
**ScopeOwnerType** | Pointer to **string** |  | [optional] 
**StaleEvaluations** | Pointer to [**ApiStaleEvaluationsDTO**](ApiStaleEvaluationsDTO.md) |  | [optional] 
**WaiverId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiStaleWaiverDTO

`func NewApiStaleWaiverDTO() *ApiStaleWaiverDTO`

NewApiStaleWaiverDTO instantiates a new ApiStaleWaiverDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStaleWaiverDTOWithDefaults

`func NewApiStaleWaiverDTOWithDefaults() *ApiStaleWaiverDTO`

NewApiStaleWaiverDTOWithDefaults instantiates a new ApiStaleWaiverDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ApiStaleWaiverDTO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ApiStaleWaiverDTO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ApiStaleWaiverDTO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ApiStaleWaiverDTO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConstraintFacts

`func (o *ApiStaleWaiverDTO) GetConstraintFacts() []ApiConstraintFactDTO`

GetConstraintFacts returns the ConstraintFacts field if non-nil, zero value otherwise.

### GetConstraintFactsOk

`func (o *ApiStaleWaiverDTO) GetConstraintFactsOk() (*[]ApiConstraintFactDTO, bool)`

GetConstraintFactsOk returns a tuple with the ConstraintFacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintFacts

`func (o *ApiStaleWaiverDTO) SetConstraintFacts(v []ApiConstraintFactDTO)`

SetConstraintFacts sets ConstraintFacts field to given value.

### HasConstraintFacts

`func (o *ApiStaleWaiverDTO) HasConstraintFacts() bool`

HasConstraintFacts returns a boolean if a field has been set.

### GetCreateTime

`func (o *ApiStaleWaiverDTO) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ApiStaleWaiverDTO) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ApiStaleWaiverDTO) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ApiStaleWaiverDTO) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *ApiStaleWaiverDTO) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ApiStaleWaiverDTO) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ApiStaleWaiverDTO) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *ApiStaleWaiverDTO) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCreatorName

`func (o *ApiStaleWaiverDTO) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *ApiStaleWaiverDTO) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *ApiStaleWaiverDTO) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *ApiStaleWaiverDTO) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetExpiryTime

`func (o *ApiStaleWaiverDTO) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *ApiStaleWaiverDTO) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *ApiStaleWaiverDTO) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *ApiStaleWaiverDTO) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApiStaleWaiverDTO) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApiStaleWaiverDTO) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApiStaleWaiverDTO) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApiStaleWaiverDTO) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiStaleWaiverDTO) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiStaleWaiverDTO) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiStaleWaiverDTO) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiStaleWaiverDTO) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyWaiverReasonId

`func (o *ApiStaleWaiverDTO) GetPolicyWaiverReasonId() string`

GetPolicyWaiverReasonId returns the PolicyWaiverReasonId field if non-nil, zero value otherwise.

### GetPolicyWaiverReasonIdOk

`func (o *ApiStaleWaiverDTO) GetPolicyWaiverReasonIdOk() (*string, bool)`

GetPolicyWaiverReasonIdOk returns a tuple with the PolicyWaiverReasonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverReasonId

`func (o *ApiStaleWaiverDTO) SetPolicyWaiverReasonId(v string)`

SetPolicyWaiverReasonId sets PolicyWaiverReasonId field to given value.

### HasPolicyWaiverReasonId

`func (o *ApiStaleWaiverDTO) HasPolicyWaiverReasonId() bool`

HasPolicyWaiverReasonId returns a boolean if a field has been set.

### GetReasonText

`func (o *ApiStaleWaiverDTO) GetReasonText() string`

GetReasonText returns the ReasonText field if non-nil, zero value otherwise.

### GetReasonTextOk

`func (o *ApiStaleWaiverDTO) GetReasonTextOk() (*string, bool)`

GetReasonTextOk returns a tuple with the ReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonText

`func (o *ApiStaleWaiverDTO) SetReasonText(v string)`

SetReasonText sets ReasonText field to given value.

### HasReasonText

`func (o *ApiStaleWaiverDTO) HasReasonText() bool`

HasReasonText returns a boolean if a field has been set.

### GetScopeOwnerId

`func (o *ApiStaleWaiverDTO) GetScopeOwnerId() string`

GetScopeOwnerId returns the ScopeOwnerId field if non-nil, zero value otherwise.

### GetScopeOwnerIdOk

`func (o *ApiStaleWaiverDTO) GetScopeOwnerIdOk() (*string, bool)`

GetScopeOwnerIdOk returns a tuple with the ScopeOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOwnerId

`func (o *ApiStaleWaiverDTO) SetScopeOwnerId(v string)`

SetScopeOwnerId sets ScopeOwnerId field to given value.

### HasScopeOwnerId

`func (o *ApiStaleWaiverDTO) HasScopeOwnerId() bool`

HasScopeOwnerId returns a boolean if a field has been set.

### GetScopeOwnerName

`func (o *ApiStaleWaiverDTO) GetScopeOwnerName() string`

GetScopeOwnerName returns the ScopeOwnerName field if non-nil, zero value otherwise.

### GetScopeOwnerNameOk

`func (o *ApiStaleWaiverDTO) GetScopeOwnerNameOk() (*string, bool)`

GetScopeOwnerNameOk returns a tuple with the ScopeOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOwnerName

`func (o *ApiStaleWaiverDTO) SetScopeOwnerName(v string)`

SetScopeOwnerName sets ScopeOwnerName field to given value.

### HasScopeOwnerName

`func (o *ApiStaleWaiverDTO) HasScopeOwnerName() bool`

HasScopeOwnerName returns a boolean if a field has been set.

### GetScopeOwnerType

`func (o *ApiStaleWaiverDTO) GetScopeOwnerType() string`

GetScopeOwnerType returns the ScopeOwnerType field if non-nil, zero value otherwise.

### GetScopeOwnerTypeOk

`func (o *ApiStaleWaiverDTO) GetScopeOwnerTypeOk() (*string, bool)`

GetScopeOwnerTypeOk returns a tuple with the ScopeOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeOwnerType

`func (o *ApiStaleWaiverDTO) SetScopeOwnerType(v string)`

SetScopeOwnerType sets ScopeOwnerType field to given value.

### HasScopeOwnerType

`func (o *ApiStaleWaiverDTO) HasScopeOwnerType() bool`

HasScopeOwnerType returns a boolean if a field has been set.

### GetStaleEvaluations

`func (o *ApiStaleWaiverDTO) GetStaleEvaluations() ApiStaleEvaluationsDTO`

GetStaleEvaluations returns the StaleEvaluations field if non-nil, zero value otherwise.

### GetStaleEvaluationsOk

`func (o *ApiStaleWaiverDTO) GetStaleEvaluationsOk() (*ApiStaleEvaluationsDTO, bool)`

GetStaleEvaluationsOk returns a tuple with the StaleEvaluations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleEvaluations

`func (o *ApiStaleWaiverDTO) SetStaleEvaluations(v ApiStaleEvaluationsDTO)`

SetStaleEvaluations sets StaleEvaluations field to given value.

### HasStaleEvaluations

`func (o *ApiStaleWaiverDTO) HasStaleEvaluations() bool`

HasStaleEvaluations returns a boolean if a field has been set.

### GetWaiverId

`func (o *ApiStaleWaiverDTO) GetWaiverId() string`

GetWaiverId returns the WaiverId field if non-nil, zero value otherwise.

### GetWaiverIdOk

`func (o *ApiStaleWaiverDTO) GetWaiverIdOk() (*string, bool)`

GetWaiverIdOk returns a tuple with the WaiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiverId

`func (o *ApiStaleWaiverDTO) SetWaiverId(v string)`

SetWaiverId sets WaiverId field to given value.

### HasWaiverId

`func (o *ApiStaleWaiverDTO) HasWaiverId() bool`

HasWaiverId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


