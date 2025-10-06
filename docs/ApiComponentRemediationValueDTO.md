# ApiComponentRemediationValueDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuggestedVersionChange** | Pointer to [**ApiSuggestedVersionChangeOptionDTO**](ApiSuggestedVersionChangeOptionDTO.md) |  | [optional] 
**VersionChanges** | Pointer to [**[]ApiVersionChangeOptionDTO**](ApiVersionChangeOptionDTO.md) |  | [optional] 

## Methods

### NewApiComponentRemediationValueDTO

`func NewApiComponentRemediationValueDTO() *ApiComponentRemediationValueDTO`

NewApiComponentRemediationValueDTO instantiates a new ApiComponentRemediationValueDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiComponentRemediationValueDTOWithDefaults

`func NewApiComponentRemediationValueDTOWithDefaults() *ApiComponentRemediationValueDTO`

NewApiComponentRemediationValueDTOWithDefaults instantiates a new ApiComponentRemediationValueDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestedVersionChange

`func (o *ApiComponentRemediationValueDTO) GetSuggestedVersionChange() ApiSuggestedVersionChangeOptionDTO`

GetSuggestedVersionChange returns the SuggestedVersionChange field if non-nil, zero value otherwise.

### GetSuggestedVersionChangeOk

`func (o *ApiComponentRemediationValueDTO) GetSuggestedVersionChangeOk() (*ApiSuggestedVersionChangeOptionDTO, bool)`

GetSuggestedVersionChangeOk returns a tuple with the SuggestedVersionChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedVersionChange

`func (o *ApiComponentRemediationValueDTO) SetSuggestedVersionChange(v ApiSuggestedVersionChangeOptionDTO)`

SetSuggestedVersionChange sets SuggestedVersionChange field to given value.

### HasSuggestedVersionChange

`func (o *ApiComponentRemediationValueDTO) HasSuggestedVersionChange() bool`

HasSuggestedVersionChange returns a boolean if a field has been set.

### GetVersionChanges

`func (o *ApiComponentRemediationValueDTO) GetVersionChanges() []ApiVersionChangeOptionDTO`

GetVersionChanges returns the VersionChanges field if non-nil, zero value otherwise.

### GetVersionChangesOk

`func (o *ApiComponentRemediationValueDTO) GetVersionChangesOk() (*[]ApiVersionChangeOptionDTO, bool)`

GetVersionChangesOk returns a tuple with the VersionChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionChanges

`func (o *ApiComponentRemediationValueDTO) SetVersionChanges(v []ApiVersionChangeOptionDTO)`

SetVersionChanges sets VersionChanges field to given value.

### HasVersionChanges

`func (o *ApiComponentRemediationValueDTO) HasVersionChanges() bool`

HasVersionChanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


