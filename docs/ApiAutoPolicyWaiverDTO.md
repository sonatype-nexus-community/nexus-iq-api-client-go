# ApiAutoPolicyWaiverDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoPolicyWaiverId** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**CreatorName** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**PathForward** | Pointer to **bool** |  | [optional] 
**Reachable** | Pointer to **bool** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiAutoPolicyWaiverDTO

`func NewApiAutoPolicyWaiverDTO() *ApiAutoPolicyWaiverDTO`

NewApiAutoPolicyWaiverDTO instantiates a new ApiAutoPolicyWaiverDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAutoPolicyWaiverDTOWithDefaults

`func NewApiAutoPolicyWaiverDTOWithDefaults() *ApiAutoPolicyWaiverDTO`

NewApiAutoPolicyWaiverDTOWithDefaults instantiates a new ApiAutoPolicyWaiverDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoPolicyWaiverId

`func (o *ApiAutoPolicyWaiverDTO) GetAutoPolicyWaiverId() string`

GetAutoPolicyWaiverId returns the AutoPolicyWaiverId field if non-nil, zero value otherwise.

### GetAutoPolicyWaiverIdOk

`func (o *ApiAutoPolicyWaiverDTO) GetAutoPolicyWaiverIdOk() (*string, bool)`

GetAutoPolicyWaiverIdOk returns a tuple with the AutoPolicyWaiverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPolicyWaiverId

`func (o *ApiAutoPolicyWaiverDTO) SetAutoPolicyWaiverId(v string)`

SetAutoPolicyWaiverId sets AutoPolicyWaiverId field to given value.

### HasAutoPolicyWaiverId

`func (o *ApiAutoPolicyWaiverDTO) HasAutoPolicyWaiverId() bool`

HasAutoPolicyWaiverId returns a boolean if a field has been set.

### GetCreateTime

`func (o *ApiAutoPolicyWaiverDTO) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ApiAutoPolicyWaiverDTO) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ApiAutoPolicyWaiverDTO) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ApiAutoPolicyWaiverDTO) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetCreatorId

`func (o *ApiAutoPolicyWaiverDTO) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *ApiAutoPolicyWaiverDTO) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *ApiAutoPolicyWaiverDTO) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *ApiAutoPolicyWaiverDTO) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### GetCreatorName

`func (o *ApiAutoPolicyWaiverDTO) GetCreatorName() string`

GetCreatorName returns the CreatorName field if non-nil, zero value otherwise.

### GetCreatorNameOk

`func (o *ApiAutoPolicyWaiverDTO) GetCreatorNameOk() (*string, bool)`

GetCreatorNameOk returns a tuple with the CreatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorName

`func (o *ApiAutoPolicyWaiverDTO) SetCreatorName(v string)`

SetCreatorName sets CreatorName field to given value.

### HasCreatorName

`func (o *ApiAutoPolicyWaiverDTO) HasCreatorName() bool`

HasCreatorName returns a boolean if a field has been set.

### GetOwnerId

`func (o *ApiAutoPolicyWaiverDTO) GetOwnerId() string`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ApiAutoPolicyWaiverDTO) GetOwnerIdOk() (*string, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ApiAutoPolicyWaiverDTO) SetOwnerId(v string)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ApiAutoPolicyWaiverDTO) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetPathForward

`func (o *ApiAutoPolicyWaiverDTO) GetPathForward() bool`

GetPathForward returns the PathForward field if non-nil, zero value otherwise.

### GetPathForwardOk

`func (o *ApiAutoPolicyWaiverDTO) GetPathForwardOk() (*bool, bool)`

GetPathForwardOk returns a tuple with the PathForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathForward

`func (o *ApiAutoPolicyWaiverDTO) SetPathForward(v bool)`

SetPathForward sets PathForward field to given value.

### HasPathForward

`func (o *ApiAutoPolicyWaiverDTO) HasPathForward() bool`

HasPathForward returns a boolean if a field has been set.

### GetReachable

`func (o *ApiAutoPolicyWaiverDTO) GetReachable() bool`

GetReachable returns the Reachable field if non-nil, zero value otherwise.

### GetReachableOk

`func (o *ApiAutoPolicyWaiverDTO) GetReachableOk() (*bool, bool)`

GetReachableOk returns a tuple with the Reachable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachable

`func (o *ApiAutoPolicyWaiverDTO) SetReachable(v bool)`

SetReachable sets Reachable field to given value.

### HasReachable

`func (o *ApiAutoPolicyWaiverDTO) HasReachable() bool`

HasReachable returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiAutoPolicyWaiverDTO) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiAutoPolicyWaiverDTO) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiAutoPolicyWaiverDTO) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiAutoPolicyWaiverDTO) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


