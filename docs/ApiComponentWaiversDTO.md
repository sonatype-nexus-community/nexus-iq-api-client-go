# ApiComponentWaiversDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationWaivers** | Pointer to [**[]ApiApplicationWaiverDTO**](ApiApplicationWaiverDTO.md) |  | [optional] 
**RepositoryWaivers** | Pointer to [**[]ApiRepositoryWaiverDTO**](ApiRepositoryWaiverDTO.md) |  | [optional] 

## Methods

### NewApiComponentWaiversDTO

`func NewApiComponentWaiversDTO() *ApiComponentWaiversDTO`

NewApiComponentWaiversDTO instantiates a new ApiComponentWaiversDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiComponentWaiversDTOWithDefaults

`func NewApiComponentWaiversDTOWithDefaults() *ApiComponentWaiversDTO`

NewApiComponentWaiversDTOWithDefaults instantiates a new ApiComponentWaiversDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationWaivers

`func (o *ApiComponentWaiversDTO) GetApplicationWaivers() []ApiApplicationWaiverDTO`

GetApplicationWaivers returns the ApplicationWaivers field if non-nil, zero value otherwise.

### GetApplicationWaiversOk

`func (o *ApiComponentWaiversDTO) GetApplicationWaiversOk() (*[]ApiApplicationWaiverDTO, bool)`

GetApplicationWaiversOk returns a tuple with the ApplicationWaivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationWaivers

`func (o *ApiComponentWaiversDTO) SetApplicationWaivers(v []ApiApplicationWaiverDTO)`

SetApplicationWaivers sets ApplicationWaivers field to given value.

### HasApplicationWaivers

`func (o *ApiComponentWaiversDTO) HasApplicationWaivers() bool`

HasApplicationWaivers returns a boolean if a field has been set.

### GetRepositoryWaivers

`func (o *ApiComponentWaiversDTO) GetRepositoryWaivers() []ApiRepositoryWaiverDTO`

GetRepositoryWaivers returns the RepositoryWaivers field if non-nil, zero value otherwise.

### GetRepositoryWaiversOk

`func (o *ApiComponentWaiversDTO) GetRepositoryWaiversOk() (*[]ApiRepositoryWaiverDTO, bool)`

GetRepositoryWaiversOk returns a tuple with the RepositoryWaivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryWaivers

`func (o *ApiComponentWaiversDTO) SetRepositoryWaivers(v []ApiRepositoryWaiverDTO)`

SetRepositoryWaivers sets RepositoryWaivers field to given value.

### HasRepositoryWaivers

`func (o *ApiComponentWaiversDTO) HasRepositoryWaivers() bool`

HasRepositoryWaivers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


