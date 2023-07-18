# ApiLicenseLegalComponentReportDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Component** | Pointer to [**ApiLicenseLegalComponentDTO**](ApiLicenseLegalComponentDTO.md) |  | [optional] 
**LicenseLegalMetadata** | Pointer to [**[]ApiLicenseLegalMetadataDTO**](ApiLicenseLegalMetadataDTO.md) |  | [optional] 

## Methods

### NewApiLicenseLegalComponentReportDTO

`func NewApiLicenseLegalComponentReportDTO() *ApiLicenseLegalComponentReportDTO`

NewApiLicenseLegalComponentReportDTO instantiates a new ApiLicenseLegalComponentReportDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLicenseLegalComponentReportDTOWithDefaults

`func NewApiLicenseLegalComponentReportDTOWithDefaults() *ApiLicenseLegalComponentReportDTO`

NewApiLicenseLegalComponentReportDTOWithDefaults instantiates a new ApiLicenseLegalComponentReportDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponent

`func (o *ApiLicenseLegalComponentReportDTO) GetComponent() ApiLicenseLegalComponentDTO`

GetComponent returns the Component field if non-nil, zero value otherwise.

### GetComponentOk

`func (o *ApiLicenseLegalComponentReportDTO) GetComponentOk() (*ApiLicenseLegalComponentDTO, bool)`

GetComponentOk returns a tuple with the Component field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponent

`func (o *ApiLicenseLegalComponentReportDTO) SetComponent(v ApiLicenseLegalComponentDTO)`

SetComponent sets Component field to given value.

### HasComponent

`func (o *ApiLicenseLegalComponentReportDTO) HasComponent() bool`

HasComponent returns a boolean if a field has been set.

### GetLicenseLegalMetadata

`func (o *ApiLicenseLegalComponentReportDTO) GetLicenseLegalMetadata() []ApiLicenseLegalMetadataDTO`

GetLicenseLegalMetadata returns the LicenseLegalMetadata field if non-nil, zero value otherwise.

### GetLicenseLegalMetadataOk

`func (o *ApiLicenseLegalComponentReportDTO) GetLicenseLegalMetadataOk() (*[]ApiLicenseLegalMetadataDTO, bool)`

GetLicenseLegalMetadataOk returns a tuple with the LicenseLegalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseLegalMetadata

`func (o *ApiLicenseLegalComponentReportDTO) SetLicenseLegalMetadata(v []ApiLicenseLegalMetadataDTO)`

SetLicenseLegalMetadata sets LicenseLegalMetadata field to given value.

### HasLicenseLegalMetadata

`func (o *ApiLicenseLegalComponentReportDTO) HasLicenseLegalMetadata() bool`

HasLicenseLegalMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


