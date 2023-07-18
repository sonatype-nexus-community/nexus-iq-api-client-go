# ApiSecurityDataDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityIssues** | Pointer to [**[]ApiSecurityIssueDTO**](ApiSecurityIssueDTO.md) |  | [optional] 

## Methods

### NewApiSecurityDataDTO

`func NewApiSecurityDataDTO() *ApiSecurityDataDTO`

NewApiSecurityDataDTO instantiates a new ApiSecurityDataDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSecurityDataDTOWithDefaults

`func NewApiSecurityDataDTOWithDefaults() *ApiSecurityDataDTO`

NewApiSecurityDataDTOWithDefaults instantiates a new ApiSecurityDataDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityIssues

`func (o *ApiSecurityDataDTO) GetSecurityIssues() []ApiSecurityIssueDTO`

GetSecurityIssues returns the SecurityIssues field if non-nil, zero value otherwise.

### GetSecurityIssuesOk

`func (o *ApiSecurityDataDTO) GetSecurityIssuesOk() (*[]ApiSecurityIssueDTO, bool)`

GetSecurityIssuesOk returns a tuple with the SecurityIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityIssues

`func (o *ApiSecurityDataDTO) SetSecurityIssues(v []ApiSecurityIssueDTO)`

SetSecurityIssues sets SecurityIssues field to given value.

### HasSecurityIssues

`func (o *ApiSecurityDataDTO) HasSecurityIssues() bool`

HasSecurityIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


