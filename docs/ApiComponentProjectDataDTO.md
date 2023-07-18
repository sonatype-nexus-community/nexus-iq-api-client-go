# ApiComponentProjectDataDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstReleaseDate** | Pointer to **time.Time** |  | [optional] 
**LastReleaseDate** | Pointer to **time.Time** |  | [optional] 
**ProjectMetadata** | Pointer to [**ApiComponentProjectMetadataDTO**](ApiComponentProjectMetadataDTO.md) |  | [optional] 
**SourceControlManagement** | Pointer to [**ApiComponentProjectScmDTO**](ApiComponentProjectScmDTO.md) |  | [optional] 

## Methods

### NewApiComponentProjectDataDTO

`func NewApiComponentProjectDataDTO() *ApiComponentProjectDataDTO`

NewApiComponentProjectDataDTO instantiates a new ApiComponentProjectDataDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiComponentProjectDataDTOWithDefaults

`func NewApiComponentProjectDataDTOWithDefaults() *ApiComponentProjectDataDTO`

NewApiComponentProjectDataDTOWithDefaults instantiates a new ApiComponentProjectDataDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstReleaseDate

`func (o *ApiComponentProjectDataDTO) GetFirstReleaseDate() time.Time`

GetFirstReleaseDate returns the FirstReleaseDate field if non-nil, zero value otherwise.

### GetFirstReleaseDateOk

`func (o *ApiComponentProjectDataDTO) GetFirstReleaseDateOk() (*time.Time, bool)`

GetFirstReleaseDateOk returns a tuple with the FirstReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstReleaseDate

`func (o *ApiComponentProjectDataDTO) SetFirstReleaseDate(v time.Time)`

SetFirstReleaseDate sets FirstReleaseDate field to given value.

### HasFirstReleaseDate

`func (o *ApiComponentProjectDataDTO) HasFirstReleaseDate() bool`

HasFirstReleaseDate returns a boolean if a field has been set.

### GetLastReleaseDate

`func (o *ApiComponentProjectDataDTO) GetLastReleaseDate() time.Time`

GetLastReleaseDate returns the LastReleaseDate field if non-nil, zero value otherwise.

### GetLastReleaseDateOk

`func (o *ApiComponentProjectDataDTO) GetLastReleaseDateOk() (*time.Time, bool)`

GetLastReleaseDateOk returns a tuple with the LastReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReleaseDate

`func (o *ApiComponentProjectDataDTO) SetLastReleaseDate(v time.Time)`

SetLastReleaseDate sets LastReleaseDate field to given value.

### HasLastReleaseDate

`func (o *ApiComponentProjectDataDTO) HasLastReleaseDate() bool`

HasLastReleaseDate returns a boolean if a field has been set.

### GetProjectMetadata

`func (o *ApiComponentProjectDataDTO) GetProjectMetadata() ApiComponentProjectMetadataDTO`

GetProjectMetadata returns the ProjectMetadata field if non-nil, zero value otherwise.

### GetProjectMetadataOk

`func (o *ApiComponentProjectDataDTO) GetProjectMetadataOk() (*ApiComponentProjectMetadataDTO, bool)`

GetProjectMetadataOk returns a tuple with the ProjectMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectMetadata

`func (o *ApiComponentProjectDataDTO) SetProjectMetadata(v ApiComponentProjectMetadataDTO)`

SetProjectMetadata sets ProjectMetadata field to given value.

### HasProjectMetadata

`func (o *ApiComponentProjectDataDTO) HasProjectMetadata() bool`

HasProjectMetadata returns a boolean if a field has been set.

### GetSourceControlManagement

`func (o *ApiComponentProjectDataDTO) GetSourceControlManagement() ApiComponentProjectScmDTO`

GetSourceControlManagement returns the SourceControlManagement field if non-nil, zero value otherwise.

### GetSourceControlManagementOk

`func (o *ApiComponentProjectDataDTO) GetSourceControlManagementOk() (*ApiComponentProjectScmDTO, bool)`

GetSourceControlManagementOk returns a tuple with the SourceControlManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceControlManagement

`func (o *ApiComponentProjectDataDTO) SetSourceControlManagement(v ApiComponentProjectScmDTO)`

SetSourceControlManagement sets SourceControlManagement field to given value.

### HasSourceControlManagement

`func (o *ApiComponentProjectDataDTO) HasSourceControlManagement() bool`

HasSourceControlManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


