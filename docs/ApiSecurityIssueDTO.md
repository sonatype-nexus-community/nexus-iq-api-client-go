# ApiSecurityIssueDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CvssVector** | Pointer to **string** |  | [optional] 
**CvssVectorSource** | Pointer to **string** |  | [optional] 
**Cwe** | Pointer to **string** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **float32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ThreatCategory** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewApiSecurityIssueDTO

`func NewApiSecurityIssueDTO() *ApiSecurityIssueDTO`

NewApiSecurityIssueDTO instantiates a new ApiSecurityIssueDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSecurityIssueDTOWithDefaults

`func NewApiSecurityIssueDTOWithDefaults() *ApiSecurityIssueDTO`

NewApiSecurityIssueDTOWithDefaults instantiates a new ApiSecurityIssueDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCvssVector

`func (o *ApiSecurityIssueDTO) GetCvssVector() string`

GetCvssVector returns the CvssVector field if non-nil, zero value otherwise.

### GetCvssVectorOk

`func (o *ApiSecurityIssueDTO) GetCvssVectorOk() (*string, bool)`

GetCvssVectorOk returns a tuple with the CvssVector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVector

`func (o *ApiSecurityIssueDTO) SetCvssVector(v string)`

SetCvssVector sets CvssVector field to given value.

### HasCvssVector

`func (o *ApiSecurityIssueDTO) HasCvssVector() bool`

HasCvssVector returns a boolean if a field has been set.

### GetCvssVectorSource

`func (o *ApiSecurityIssueDTO) GetCvssVectorSource() string`

GetCvssVectorSource returns the CvssVectorSource field if non-nil, zero value otherwise.

### GetCvssVectorSourceOk

`func (o *ApiSecurityIssueDTO) GetCvssVectorSourceOk() (*string, bool)`

GetCvssVectorSourceOk returns a tuple with the CvssVectorSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvssVectorSource

`func (o *ApiSecurityIssueDTO) SetCvssVectorSource(v string)`

SetCvssVectorSource sets CvssVectorSource field to given value.

### HasCvssVectorSource

`func (o *ApiSecurityIssueDTO) HasCvssVectorSource() bool`

HasCvssVectorSource returns a boolean if a field has been set.

### GetCwe

`func (o *ApiSecurityIssueDTO) GetCwe() string`

GetCwe returns the Cwe field if non-nil, zero value otherwise.

### GetCweOk

`func (o *ApiSecurityIssueDTO) GetCweOk() (*string, bool)`

GetCweOk returns a tuple with the Cwe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCwe

`func (o *ApiSecurityIssueDTO) SetCwe(v string)`

SetCwe sets Cwe field to given value.

### HasCwe

`func (o *ApiSecurityIssueDTO) HasCwe() bool`

HasCwe returns a boolean if a field has been set.

### GetReference

`func (o *ApiSecurityIssueDTO) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *ApiSecurityIssueDTO) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *ApiSecurityIssueDTO) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *ApiSecurityIssueDTO) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetSeverity

`func (o *ApiSecurityIssueDTO) GetSeverity() float32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ApiSecurityIssueDTO) GetSeverityOk() (*float32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ApiSecurityIssueDTO) SetSeverity(v float32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ApiSecurityIssueDTO) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSource

`func (o *ApiSecurityIssueDTO) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiSecurityIssueDTO) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiSecurityIssueDTO) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiSecurityIssueDTO) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *ApiSecurityIssueDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiSecurityIssueDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiSecurityIssueDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiSecurityIssueDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThreatCategory

`func (o *ApiSecurityIssueDTO) GetThreatCategory() string`

GetThreatCategory returns the ThreatCategory field if non-nil, zero value otherwise.

### GetThreatCategoryOk

`func (o *ApiSecurityIssueDTO) GetThreatCategoryOk() (*string, bool)`

GetThreatCategoryOk returns a tuple with the ThreatCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatCategory

`func (o *ApiSecurityIssueDTO) SetThreatCategory(v string)`

SetThreatCategory sets ThreatCategory field to given value.

### HasThreatCategory

`func (o *ApiSecurityIssueDTO) HasThreatCategory() bool`

HasThreatCategory returns a boolean if a field has been set.

### GetUrl

`func (o *ApiSecurityIssueDTO) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiSecurityIssueDTO) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiSecurityIssueDTO) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApiSecurityIssueDTO) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


