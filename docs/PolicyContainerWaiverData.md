# PolicyContainerWaiverData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationScope** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**ExpiryTime** | Pointer to **time.Time** |  | [optional] 
**MaxThreatLevel** | Pointer to **int32** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**PolicyWaiverId** | Pointer to **string** |  | [optional] 
**UniqueComponentCount** | Pointer to **int64** |  | [optional] 
**UniquePolicyCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewPolicyContainerWaiverData

`func NewPolicyContainerWaiverData() *PolicyContainerWaiverData`

NewPolicyContainerWaiverData instantiates a new PolicyContainerWaiverData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyContainerWaiverDataWithDefaults

`func NewPolicyContainerWaiverDataWithDefaults() *PolicyContainerWaiverData`

NewPolicyContainerWaiverDataWithDefaults instantiates a new PolicyContainerWaiverData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationScope

`func (o *PolicyContainerWaiverData) GetApplicationScope() string`

GetApplicationScope returns the ApplicationScope field if non-nil, zero value otherwise.

### GetApplicationScopeOk

`func (o *PolicyContainerWaiverData) GetApplicationScopeOk() (*string, bool)`

GetApplicationScopeOk returns a tuple with the ApplicationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationScope

`func (o *PolicyContainerWaiverData) SetApplicationScope(v string)`

SetApplicationScope sets ApplicationScope field to given value.

### HasApplicationScope

`func (o *PolicyContainerWaiverData) HasApplicationScope() bool`

HasApplicationScope returns a boolean if a field has been set.

### GetCreateTime

`func (o *PolicyContainerWaiverData) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *PolicyContainerWaiverData) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *PolicyContainerWaiverData) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *PolicyContainerWaiverData) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetExpiryTime

`func (o *PolicyContainerWaiverData) GetExpiryTime() time.Time`

GetExpiryTime returns the ExpiryTime field if non-nil, zero value otherwise.

### GetExpiryTimeOk

`func (o *PolicyContainerWaiverData) GetExpiryTimeOk() (*time.Time, bool)`

GetExpiryTimeOk returns a tuple with the ExpiryTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryTime

`func (o *PolicyContainerWaiverData) SetExpiryTime(v time.Time)`

SetExpiryTime sets ExpiryTime field to given value.

### HasExpiryTime

`func (o *PolicyContainerWaiverData) HasExpiryTime() bool`

HasExpiryTime returns a boolean if a field has been set.

### GetMaxThreatLevel

`func (o *PolicyContainerWaiverData) GetMaxThreatLevel() int32`

GetMaxThreatLevel returns the MaxThreatLevel field if non-nil, zero value otherwise.

### GetMaxThreatLevelOk

`func (o *PolicyContainerWaiverData) GetMaxThreatLevelOk() (*int32, bool)`

GetMaxThreatLevelOk returns a tuple with the MaxThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThreatLevel

`func (o *PolicyContainerWaiverData) SetMaxThreatLevel(v int32)`

SetMaxThreatLevel sets MaxThreatLevel field to given value.

### HasMaxThreatLevel

`func (o *PolicyContainerWaiverData) HasMaxThreatLevel() bool`

HasMaxThreatLevel returns a boolean if a field has been set.

### GetOwnerId

`func (o *PolicyContainerWaiverData) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *PolicyContainerWaiverData) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *PolicyContainerWaiverData) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *PolicyContainerWaiverData) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPolicyWaiverId

`func (o *PolicyContainerWaiverData) GetPolicyWaiverId() string`

GetPolicyWaiverId returns the PolicyWaiverId field if non-nil, zero value otherwise.

### GetPolicyWaiverIdOk

`func (o *PolicyContainerWaiverData) GetPolicyWaiverIdOk() (*string, bool)`

GetPolicyWaiverIdOk returns a tuple with the PolicyWaiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyWaiverId

`func (o *PolicyContainerWaiverData) SetPolicyWaiverId(v string)`

SetPolicyWaiverId sets PolicyWaiverId field to given value.

### HasPolicyWaiverId

`func (o *PolicyContainerWaiverData) HasPolicyWaiverId() bool`

HasPolicyWaiverId returns a boolean if a field has been set.

### GetUniqueComponentCount

`func (o *PolicyContainerWaiverData) GetUniqueComponentCount() int64`

GetUniqueComponentCount returns the UniqueComponentCount field if non-nil, zero value otherwise.

### GetUniqueComponentCountOk

`func (o *PolicyContainerWaiverData) GetUniqueComponentCountOk() (*int64, bool)`

GetUniqueComponentCountOk returns a tuple with the UniqueComponentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueComponentCount

`func (o *PolicyContainerWaiverData) SetUniqueComponentCount(v int64)`

SetUniqueComponentCount sets UniqueComponentCount field to given value.

### HasUniqueComponentCount

`func (o *PolicyContainerWaiverData) HasUniqueComponentCount() bool`

HasUniqueComponentCount returns a boolean if a field has been set.

### GetUniquePolicyCount

`func (o *PolicyContainerWaiverData) GetUniquePolicyCount() int64`

GetUniquePolicyCount returns the UniquePolicyCount field if non-nil, zero value otherwise.

### GetUniquePolicyCountOk

`func (o *PolicyContainerWaiverData) GetUniquePolicyCountOk() (*int64, bool)`

GetUniquePolicyCountOk returns a tuple with the UniquePolicyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniquePolicyCount

`func (o *PolicyContainerWaiverData) SetUniquePolicyCount(v int64)`

SetUniquePolicyCount sets UniquePolicyCount field to given value.

### HasUniquePolicyCount

`func (o *PolicyContainerWaiverData) HasUniquePolicyCount() bool`

HasUniquePolicyCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


