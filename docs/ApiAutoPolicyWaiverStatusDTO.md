# ApiAutoPolicyWaiverStatusDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoPolicyWaiverId** | Pointer to **string** |  | [optional] 
**AutoPolicyWaiverOwnerId** | Pointer to **string** |  | [optional] 
**AutoPolicyWaiverOwnerName** | Pointer to **string** |  | [optional] 
**AutoPolicyWaiverOwnerType** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**HasNoPathForward** | Pointer to **bool** |  | [optional] 
**HasNotReachable** | Pointer to **bool** |  | [optional] 
**IsAutoWaiverEnabled** | Pointer to **bool** |  | [optional] 
**IsInherited** | Pointer to **bool** |  | [optional] 
**ScopesOperatorAny** | Pointer to **bool** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiAutoPolicyWaiverStatusDTO

`func NewApiAutoPolicyWaiverStatusDTO() *ApiAutoPolicyWaiverStatusDTO`

NewApiAutoPolicyWaiverStatusDTO instantiates a new ApiAutoPolicyWaiverStatusDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAutoPolicyWaiverStatusDTOWithDefaults

`func NewApiAutoPolicyWaiverStatusDTOWithDefaults() *ApiAutoPolicyWaiverStatusDTO`

NewApiAutoPolicyWaiverStatusDTOWithDefaults instantiates a new ApiAutoPolicyWaiverStatusDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoPolicyWaiverId

`func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverId() string`

GetAutoPolicyWaiverId returns the AutoPolicyWaiverId field if non-nil, zero value otherwise.

### GetAutoPolicyWaiverIdOk

`func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverIdOk() (*string, bool)`

GetAutoPolicyWaiverIdOk returns a tuple with the AutoPolicyWaiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPolicyWaiverId

`func (o *ApiAutoPolicyWaiverStatusDTO) SetAutoPolicyWaiverId(v string)`

SetAutoPolicyWaiverId sets AutoPolicyWaiverId field to given value.

### HasAutoPolicyWaiverId

`func (o *ApiAutoPolicyWaiverStatusDTO) HasAutoPolicyWaiverId() bool`

HasAutoPolicyWaiverId returns a boolean if a field has been set.

### GetAutoPolicyWaiverOwnerId

`func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverOwnerId() string`

GetAutoPolicyWaiverOwnerId returns the AutoPolicyWaiverOwnerId field if non-nil, zero value otherwise.

### GetAutoPolicyWaiverOwnerIdOk

`func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverOwnerIdOk() (*string, bool)`

GetAutoPolicyWaiverOwnerIdOk returns a tuple with the AutoPolicyWaiverOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPolicyWaiverOwnerId

`func (o *ApiAutoPolicyWaiverStatusDTO) SetAutoPolicyWaiverOwnerId(v string)`

SetAutoPolicyWaiverOwnerId sets AutoPolicyWaiverOwnerId field to given value.

### HasAutoPolicyWaiverOwnerId

`func (o *ApiAutoPolicyWaiverStatusDTO) HasAutoPolicyWaiverOwnerId() bool`

HasAutoPolicyWaiverOwnerId returns a boolean if a field has been set.

### GetAutoPolicyWaiverOwnerName

`func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverOwnerName() string`

GetAutoPolicyWaiverOwnerName returns the AutoPolicyWaiverOwnerName field if non-nil, zero value otherwise.

### GetAutoPolicyWaiverOwnerNameOk

`func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverOwnerNameOk() (*string, bool)`

GetAutoPolicyWaiverOwnerNameOk returns a tuple with the AutoPolicyWaiverOwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPolicyWaiverOwnerName

`func (o *ApiAutoPolicyWaiverStatusDTO) SetAutoPolicyWaiverOwnerName(v string)`

SetAutoPolicyWaiverOwnerName sets AutoPolicyWaiverOwnerName field to given value.

### HasAutoPolicyWaiverOwnerName

`func (o *ApiAutoPolicyWaiverStatusDTO) HasAutoPolicyWaiverOwnerName() bool`

HasAutoPolicyWaiverOwnerName returns a boolean if a field has been set.

### GetAutoPolicyWaiverOwnerType

`func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverOwnerType() string`

GetAutoPolicyWaiverOwnerType returns the AutoPolicyWaiverOwnerType field if non-nil, zero value otherwise.

### GetAutoPolicyWaiverOwnerTypeOk

`func (o *ApiAutoPolicyWaiverStatusDTO) GetAutoPolicyWaiverOwnerTypeOk() (*string, bool)`

GetAutoPolicyWaiverOwnerTypeOk returns a tuple with the AutoPolicyWaiverOwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPolicyWaiverOwnerType

`func (o *ApiAutoPolicyWaiverStatusDTO) SetAutoPolicyWaiverOwnerType(v string)`

SetAutoPolicyWaiverOwnerType sets AutoPolicyWaiverOwnerType field to given value.

### HasAutoPolicyWaiverOwnerType

`func (o *ApiAutoPolicyWaiverStatusDTO) HasAutoPolicyWaiverOwnerType() bool`

HasAutoPolicyWaiverOwnerType returns a boolean if a field has been set.

### GetCreateTime

`func (o *ApiAutoPolicyWaiverStatusDTO) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ApiAutoPolicyWaiverStatusDTO) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ApiAutoPolicyWaiverStatusDTO) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ApiAutoPolicyWaiverStatusDTO) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetHasNoPathForward

`func (o *ApiAutoPolicyWaiverStatusDTO) GetHasNoPathForward() bool`

GetHasNoPathForward returns the HasNoPathForward field if non-nil, zero value otherwise.

### GetHasNoPathForwardOk

`func (o *ApiAutoPolicyWaiverStatusDTO) GetHasNoPathForwardOk() (*bool, bool)`

GetHasNoPathForwardOk returns a tuple with the HasNoPathForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNoPathForward

`func (o *ApiAutoPolicyWaiverStatusDTO) SetHasNoPathForward(v bool)`

SetHasNoPathForward sets HasNoPathForward field to given value.

### HasHasNoPathForward

`func (o *ApiAutoPolicyWaiverStatusDTO) HasHasNoPathForward() bool`

HasHasNoPathForward returns a boolean if a field has been set.

### GetHasNotReachable

`func (o *ApiAutoPolicyWaiverStatusDTO) GetHasNotReachable() bool`

GetHasNotReachable returns the HasNotReachable field if non-nil, zero value otherwise.

### GetHasNotReachableOk

`func (o *ApiAutoPolicyWaiverStatusDTO) GetHasNotReachableOk() (*bool, bool)`

GetHasNotReachableOk returns a tuple with the HasNotReachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNotReachable

`func (o *ApiAutoPolicyWaiverStatusDTO) SetHasNotReachable(v bool)`

SetHasNotReachable sets HasNotReachable field to given value.

### HasHasNotReachable

`func (o *ApiAutoPolicyWaiverStatusDTO) HasHasNotReachable() bool`

HasHasNotReachable returns a boolean if a field has been set.

### GetIsAutoWaiverEnabled

`func (o *ApiAutoPolicyWaiverStatusDTO) GetIsAutoWaiverEnabled() bool`

GetIsAutoWaiverEnabled returns the IsAutoWaiverEnabled field if non-nil, zero value otherwise.

### GetIsAutoWaiverEnabledOk

`func (o *ApiAutoPolicyWaiverStatusDTO) GetIsAutoWaiverEnabledOk() (*bool, bool)`

GetIsAutoWaiverEnabledOk returns a tuple with the IsAutoWaiverEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoWaiverEnabled

`func (o *ApiAutoPolicyWaiverStatusDTO) SetIsAutoWaiverEnabled(v bool)`

SetIsAutoWaiverEnabled sets IsAutoWaiverEnabled field to given value.

### HasIsAutoWaiverEnabled

`func (o *ApiAutoPolicyWaiverStatusDTO) HasIsAutoWaiverEnabled() bool`

HasIsAutoWaiverEnabled returns a boolean if a field has been set.

### GetIsInherited

`func (o *ApiAutoPolicyWaiverStatusDTO) GetIsInherited() bool`

GetIsInherited returns the IsInherited field if non-nil, zero value otherwise.

### GetIsInheritedOk

`func (o *ApiAutoPolicyWaiverStatusDTO) GetIsInheritedOk() (*bool, bool)`

GetIsInheritedOk returns a tuple with the IsInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInherited

`func (o *ApiAutoPolicyWaiverStatusDTO) SetIsInherited(v bool)`

SetIsInherited sets IsInherited field to given value.

### HasIsInherited

`func (o *ApiAutoPolicyWaiverStatusDTO) HasIsInherited() bool`

HasIsInherited returns a boolean if a field has been set.

### GetScopesOperatorAny

`func (o *ApiAutoPolicyWaiverStatusDTO) GetScopesOperatorAny() bool`

GetScopesOperatorAny returns the ScopesOperatorAny field if non-nil, zero value otherwise.

### GetScopesOperatorAnyOk

`func (o *ApiAutoPolicyWaiverStatusDTO) GetScopesOperatorAnyOk() (*bool, bool)`

GetScopesOperatorAnyOk returns a tuple with the ScopesOperatorAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesOperatorAny

`func (o *ApiAutoPolicyWaiverStatusDTO) SetScopesOperatorAny(v bool)`

SetScopesOperatorAny sets ScopesOperatorAny field to given value.

### HasScopesOperatorAny

`func (o *ApiAutoPolicyWaiverStatusDTO) HasScopesOperatorAny() bool`

HasScopesOperatorAny returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiAutoPolicyWaiverStatusDTO) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiAutoPolicyWaiverStatusDTO) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiAutoPolicyWaiverStatusDTO) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiAutoPolicyWaiverStatusDTO) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


