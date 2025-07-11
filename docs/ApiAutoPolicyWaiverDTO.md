# ApiAutoPolicyWaiverDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoPolicyWaiverId** | Pointer to **string** |  | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**CreatorId** | Pointer to **string** |  | [optional] 
**CreatorName** | Pointer to **string** |  | [optional] 
**OwnerId** | Pointer to **string** |  | [optional] 
**OwnerName** | Pointer to **string** |  | [optional] 
**OwnerType** | Pointer to **string** |  | [optional] 
**PathForward** | Pointer to **bool** |  | [optional] 
**PublicId** | Pointer to **string** |  | [optional] 
**Reachability** | Pointer to **bool** |  | [optional] 
**ScopesOperatorAny** | Pointer to **bool** |  | [optional] 
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

### GetOwnerName

`func (o *ApiAutoPolicyWaiverDTO) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *ApiAutoPolicyWaiverDTO) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *ApiAutoPolicyWaiverDTO) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *ApiAutoPolicyWaiverDTO) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetOwnerType

`func (o *ApiAutoPolicyWaiverDTO) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *ApiAutoPolicyWaiverDTO) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *ApiAutoPolicyWaiverDTO) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *ApiAutoPolicyWaiverDTO) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

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

### GetPublicId

`func (o *ApiAutoPolicyWaiverDTO) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *ApiAutoPolicyWaiverDTO) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *ApiAutoPolicyWaiverDTO) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *ApiAutoPolicyWaiverDTO) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetReachability

`func (o *ApiAutoPolicyWaiverDTO) GetReachability() bool`

GetReachability returns the Reachability field if non-nil, zero value otherwise.

### GetReachabilityOk

`func (o *ApiAutoPolicyWaiverDTO) GetReachabilityOk() (*bool, bool)`

GetReachabilityOk returns a tuple with the Reachability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachability

`func (o *ApiAutoPolicyWaiverDTO) SetReachability(v bool)`

SetReachability sets Reachability field to given value.

### HasReachability

`func (o *ApiAutoPolicyWaiverDTO) HasReachability() bool`

HasReachability returns a boolean if a field has been set.

### GetScopesOperatorAny

`func (o *ApiAutoPolicyWaiverDTO) GetScopesOperatorAny() bool`

GetScopesOperatorAny returns the ScopesOperatorAny field if non-nil, zero value otherwise.

### GetScopesOperatorAnyOk

`func (o *ApiAutoPolicyWaiverDTO) GetScopesOperatorAnyOk() (*bool, bool)`

GetScopesOperatorAnyOk returns a tuple with the ScopesOperatorAny field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopesOperatorAny

`func (o *ApiAutoPolicyWaiverDTO) SetScopesOperatorAny(v bool)`

SetScopesOperatorAny sets ScopesOperatorAny field to given value.

### HasScopesOperatorAny

`func (o *ApiAutoPolicyWaiverDTO) HasScopesOperatorAny() bool`

HasScopesOperatorAny returns a boolean if a field has been set.

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


