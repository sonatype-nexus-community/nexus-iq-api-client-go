# ApiStagePolicyViolationComponentDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**PackageUrl** | Pointer to **string** |  | [optional] 
**PolicyId** | Pointer to **string** |  | [optional] 
**PolicyName** | Pointer to **string** |  | [optional] 
**PolicyViolationId** | Pointer to **string** |  | [optional] 
**ThreatCategory** | Pointer to **string** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiStagePolicyViolationComponentDTO

`func NewApiStagePolicyViolationComponentDTO() *ApiStagePolicyViolationComponentDTO`

NewApiStagePolicyViolationComponentDTO instantiates a new ApiStagePolicyViolationComponentDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStagePolicyViolationComponentDTOWithDefaults

`func NewApiStagePolicyViolationComponentDTOWithDefaults() *ApiStagePolicyViolationComponentDTO`

NewApiStagePolicyViolationComponentDTOWithDefaults instantiates a new ApiStagePolicyViolationComponentDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ApiStagePolicyViolationComponentDTO) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ApiStagePolicyViolationComponentDTO) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ApiStagePolicyViolationComponentDTO) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ApiStagePolicyViolationComponentDTO) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *ApiStagePolicyViolationComponentDTO) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiStagePolicyViolationComponentDTO) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiStagePolicyViolationComponentDTO) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiStagePolicyViolationComponentDTO) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApiStagePolicyViolationComponentDTO) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApiStagePolicyViolationComponentDTO) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApiStagePolicyViolationComponentDTO) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApiStagePolicyViolationComponentDTO) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHash

`func (o *ApiStagePolicyViolationComponentDTO) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiStagePolicyViolationComponentDTO) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiStagePolicyViolationComponentDTO) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiStagePolicyViolationComponentDTO) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetPackageUrl

`func (o *ApiStagePolicyViolationComponentDTO) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *ApiStagePolicyViolationComponentDTO) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *ApiStagePolicyViolationComponentDTO) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *ApiStagePolicyViolationComponentDTO) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.

### GetPolicyId

`func (o *ApiStagePolicyViolationComponentDTO) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *ApiStagePolicyViolationComponentDTO) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *ApiStagePolicyViolationComponentDTO) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *ApiStagePolicyViolationComponentDTO) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *ApiStagePolicyViolationComponentDTO) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *ApiStagePolicyViolationComponentDTO) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *ApiStagePolicyViolationComponentDTO) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *ApiStagePolicyViolationComponentDTO) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyViolationId

`func (o *ApiStagePolicyViolationComponentDTO) GetPolicyViolationId() string`

GetPolicyViolationId returns the PolicyViolationId field if non-nil, zero value otherwise.

### GetPolicyViolationIdOk

`func (o *ApiStagePolicyViolationComponentDTO) GetPolicyViolationIdOk() (*string, bool)`

GetPolicyViolationIdOk returns a tuple with the PolicyViolationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyViolationId

`func (o *ApiStagePolicyViolationComponentDTO) SetPolicyViolationId(v string)`

SetPolicyViolationId sets PolicyViolationId field to given value.

### HasPolicyViolationId

`func (o *ApiStagePolicyViolationComponentDTO) HasPolicyViolationId() bool`

HasPolicyViolationId returns a boolean if a field has been set.

### GetThreatCategory

`func (o *ApiStagePolicyViolationComponentDTO) GetThreatCategory() string`

GetThreatCategory returns the ThreatCategory field if non-nil, zero value otherwise.

### GetThreatCategoryOk

`func (o *ApiStagePolicyViolationComponentDTO) GetThreatCategoryOk() (*string, bool)`

GetThreatCategoryOk returns a tuple with the ThreatCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCategory

`func (o *ApiStagePolicyViolationComponentDTO) SetThreatCategory(v string)`

SetThreatCategory sets ThreatCategory field to given value.

### HasThreatCategory

`func (o *ApiStagePolicyViolationComponentDTO) HasThreatCategory() bool`

HasThreatCategory returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiStagePolicyViolationComponentDTO) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiStagePolicyViolationComponentDTO) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiStagePolicyViolationComponentDTO) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiStagePolicyViolationComponentDTO) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


