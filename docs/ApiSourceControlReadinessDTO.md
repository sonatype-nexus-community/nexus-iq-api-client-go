# ApiSourceControlReadinessDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbmState** | Pointer to **string** | Default branch monitoring state. ACTIVE means nothing is suppressing monitoring; it is not a guarantee that a scan is imminent | [optional] 
**HasCompletedInitialSourceEvaluation** | Pointer to **bool** | Whether a Source-stage evaluation has completed for this repository at least once | [optional] 
**InternalOwnerId** | Pointer to **string** | Internal id of the owner | [optional] 
**LastSourceEvaluationTime** | Pointer to **time.Time** | When the most recent Source-stage evaluation ran, if any | [optional] 
**OwnerType** | Pointer to **string** | Owner type of the configuration; always &#39;application&#39; | [optional] 
**PrEvaluationEligible** | Pointer to **bool** | Whether configuration permits pull request (Develop-stage) evaluation. Describes configuration only, not whether any pull request currently exists | [optional] 
**PublicId** | Pointer to **string** | Public id of the application | [optional] 
**Reasons** | Pointer to **[]string** | Stable reason codes explaining why the repository is not fully ready. Empty when the repository is ready, so callers may gate on this list being empty | [optional] 

## Methods

### NewApiSourceControlReadinessDTO

`func NewApiSourceControlReadinessDTO() *ApiSourceControlReadinessDTO`

NewApiSourceControlReadinessDTO instantiates a new ApiSourceControlReadinessDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSourceControlReadinessDTOWithDefaults

`func NewApiSourceControlReadinessDTOWithDefaults() *ApiSourceControlReadinessDTO`

NewApiSourceControlReadinessDTOWithDefaults instantiates a new ApiSourceControlReadinessDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbmState

`func (o *ApiSourceControlReadinessDTO) GetDbmState() string`

GetDbmState returns the DbmState field if non-nil, zero value otherwise.

### GetDbmStateOk

`func (o *ApiSourceControlReadinessDTO) GetDbmStateOk() (*string, bool)`

GetDbmStateOk returns a tuple with the DbmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbmState

`func (o *ApiSourceControlReadinessDTO) SetDbmState(v string)`

SetDbmState sets DbmState field to given value.

### HasDbmState

`func (o *ApiSourceControlReadinessDTO) HasDbmState() bool`

HasDbmState returns a boolean if a field has been set.

### GetHasCompletedInitialSourceEvaluation

`func (o *ApiSourceControlReadinessDTO) GetHasCompletedInitialSourceEvaluation() bool`

GetHasCompletedInitialSourceEvaluation returns the HasCompletedInitialSourceEvaluation field if non-nil, zero value otherwise.

### GetHasCompletedInitialSourceEvaluationOk

`func (o *ApiSourceControlReadinessDTO) GetHasCompletedInitialSourceEvaluationOk() (*bool, bool)`

GetHasCompletedInitialSourceEvaluationOk returns a tuple with the HasCompletedInitialSourceEvaluation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCompletedInitialSourceEvaluation

`func (o *ApiSourceControlReadinessDTO) SetHasCompletedInitialSourceEvaluation(v bool)`

SetHasCompletedInitialSourceEvaluation sets HasCompletedInitialSourceEvaluation field to given value.

### HasHasCompletedInitialSourceEvaluation

`func (o *ApiSourceControlReadinessDTO) HasHasCompletedInitialSourceEvaluation() bool`

HasHasCompletedInitialSourceEvaluation returns a boolean if a field has been set.

### GetInternalOwnerId

`func (o *ApiSourceControlReadinessDTO) GetInternalOwnerId() string`

GetInternalOwnerId returns the InternalOwnerId field if non-nil, zero value otherwise.

### GetInternalOwnerIdOk

`func (o *ApiSourceControlReadinessDTO) GetInternalOwnerIdOk() (*string, bool)`

GetInternalOwnerIdOk returns a tuple with the InternalOwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalOwnerId

`func (o *ApiSourceControlReadinessDTO) SetInternalOwnerId(v string)`

SetInternalOwnerId sets InternalOwnerId field to given value.

### HasInternalOwnerId

`func (o *ApiSourceControlReadinessDTO) HasInternalOwnerId() bool`

HasInternalOwnerId returns a boolean if a field has been set.

### GetLastSourceEvaluationTime

`func (o *ApiSourceControlReadinessDTO) GetLastSourceEvaluationTime() time.Time`

GetLastSourceEvaluationTime returns the LastSourceEvaluationTime field if non-nil, zero value otherwise.

### GetLastSourceEvaluationTimeOk

`func (o *ApiSourceControlReadinessDTO) GetLastSourceEvaluationTimeOk() (*time.Time, bool)`

GetLastSourceEvaluationTimeOk returns a tuple with the LastSourceEvaluationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSourceEvaluationTime

`func (o *ApiSourceControlReadinessDTO) SetLastSourceEvaluationTime(v time.Time)`

SetLastSourceEvaluationTime sets LastSourceEvaluationTime field to given value.

### HasLastSourceEvaluationTime

`func (o *ApiSourceControlReadinessDTO) HasLastSourceEvaluationTime() bool`

HasLastSourceEvaluationTime returns a boolean if a field has been set.

### GetOwnerType

`func (o *ApiSourceControlReadinessDTO) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *ApiSourceControlReadinessDTO) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *ApiSourceControlReadinessDTO) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *ApiSourceControlReadinessDTO) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetPrEvaluationEligible

`func (o *ApiSourceControlReadinessDTO) GetPrEvaluationEligible() bool`

GetPrEvaluationEligible returns the PrEvaluationEligible field if non-nil, zero value otherwise.

### GetPrEvaluationEligibleOk

`func (o *ApiSourceControlReadinessDTO) GetPrEvaluationEligibleOk() (*bool, bool)`

GetPrEvaluationEligibleOk returns a tuple with the PrEvaluationEligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrEvaluationEligible

`func (o *ApiSourceControlReadinessDTO) SetPrEvaluationEligible(v bool)`

SetPrEvaluationEligible sets PrEvaluationEligible field to given value.

### HasPrEvaluationEligible

`func (o *ApiSourceControlReadinessDTO) HasPrEvaluationEligible() bool`

HasPrEvaluationEligible returns a boolean if a field has been set.

### GetPublicId

`func (o *ApiSourceControlReadinessDTO) GetPublicId() string`

GetPublicId returns the PublicId field if non-nil, zero value otherwise.

### GetPublicIdOk

`func (o *ApiSourceControlReadinessDTO) GetPublicIdOk() (*string, bool)`

GetPublicIdOk returns a tuple with the PublicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicId

`func (o *ApiSourceControlReadinessDTO) SetPublicId(v string)`

SetPublicId sets PublicId field to given value.

### HasPublicId

`func (o *ApiSourceControlReadinessDTO) HasPublicId() bool`

HasPublicId returns a boolean if a field has been set.

### GetReasons

`func (o *ApiSourceControlReadinessDTO) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *ApiSourceControlReadinessDTO) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *ApiSourceControlReadinessDTO) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *ApiSourceControlReadinessDTO) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


