# SCMUserMatchingResultDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchedUsers** | Pointer to **[]string** |  | [optional] 
**SuccessfulMapping** | Pointer to [**UserMapping**](UserMapping.md) |  | [optional] 

## Methods

### NewSCMUserMatchingResultDTO

`func NewSCMUserMatchingResultDTO() *SCMUserMatchingResultDTO`

NewSCMUserMatchingResultDTO instantiates a new SCMUserMatchingResultDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSCMUserMatchingResultDTOWithDefaults

`func NewSCMUserMatchingResultDTOWithDefaults() *SCMUserMatchingResultDTO`

NewSCMUserMatchingResultDTOWithDefaults instantiates a new SCMUserMatchingResultDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchedUsers

`func (o *SCMUserMatchingResultDTO) GetMatchedUsers() []string`

GetMatchedUsers returns the MatchedUsers field if non-nil, zero value otherwise.

### GetMatchedUsersOk

`func (o *SCMUserMatchingResultDTO) GetMatchedUsersOk() (*[]string, bool)`

GetMatchedUsersOk returns a tuple with the MatchedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchedUsers

`func (o *SCMUserMatchingResultDTO) SetMatchedUsers(v []string)`

SetMatchedUsers sets MatchedUsers field to given value.

### HasMatchedUsers

`func (o *SCMUserMatchingResultDTO) HasMatchedUsers() bool`

HasMatchedUsers returns a boolean if a field has been set.

### GetSuccessfulMapping

`func (o *SCMUserMatchingResultDTO) GetSuccessfulMapping() UserMapping`

GetSuccessfulMapping returns the SuccessfulMapping field if non-nil, zero value otherwise.

### GetSuccessfulMappingOk

`func (o *SCMUserMatchingResultDTO) GetSuccessfulMappingOk() (*UserMapping, bool)`

GetSuccessfulMappingOk returns a tuple with the SuccessfulMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulMapping

`func (o *SCMUserMatchingResultDTO) SetSuccessfulMapping(v UserMapping)`

SetSuccessfulMapping sets SuccessfulMapping field to given value.

### HasSuccessfulMapping

`func (o *SCMUserMatchingResultDTO) HasSuccessfulMapping() bool`

HasSuccessfulMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


