# ApiOwnerArtifactoryConnectionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactoryConnection** | Pointer to [**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md) |  | [optional] 
**ArtifactoryConnectionStatus** | Pointer to [**ApiArtifactoryConnectionStatusResponseDTO**](ApiArtifactoryConnectionStatusResponseDTO.md) |  | [optional] 
**OwnerDTO** | Pointer to [**ApiOwnerDTO**](ApiOwnerDTO.md) |  | [optional] 

## Methods

### NewApiOwnerArtifactoryConnectionDTO

`func NewApiOwnerArtifactoryConnectionDTO() *ApiOwnerArtifactoryConnectionDTO`

NewApiOwnerArtifactoryConnectionDTO instantiates a new ApiOwnerArtifactoryConnectionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiOwnerArtifactoryConnectionDTOWithDefaults

`func NewApiOwnerArtifactoryConnectionDTOWithDefaults() *ApiOwnerArtifactoryConnectionDTO`

NewApiOwnerArtifactoryConnectionDTOWithDefaults instantiates a new ApiOwnerArtifactoryConnectionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactoryConnection

`func (o *ApiOwnerArtifactoryConnectionDTO) GetArtifactoryConnection() ApiArtifactoryConnectionDTO`

GetArtifactoryConnection returns the ArtifactoryConnection field if non-nil, zero value otherwise.

### GetArtifactoryConnectionOk

`func (o *ApiOwnerArtifactoryConnectionDTO) GetArtifactoryConnectionOk() (*ApiArtifactoryConnectionDTO, bool)`

GetArtifactoryConnectionOk returns a tuple with the ArtifactoryConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryConnection

`func (o *ApiOwnerArtifactoryConnectionDTO) SetArtifactoryConnection(v ApiArtifactoryConnectionDTO)`

SetArtifactoryConnection sets ArtifactoryConnection field to given value.

### HasArtifactoryConnection

`func (o *ApiOwnerArtifactoryConnectionDTO) HasArtifactoryConnection() bool`

HasArtifactoryConnection returns a boolean if a field has been set.

### GetArtifactoryConnectionStatus

`func (o *ApiOwnerArtifactoryConnectionDTO) GetArtifactoryConnectionStatus() ApiArtifactoryConnectionStatusResponseDTO`

GetArtifactoryConnectionStatus returns the ArtifactoryConnectionStatus field if non-nil, zero value otherwise.

### GetArtifactoryConnectionStatusOk

`func (o *ApiOwnerArtifactoryConnectionDTO) GetArtifactoryConnectionStatusOk() (*ApiArtifactoryConnectionStatusResponseDTO, bool)`

GetArtifactoryConnectionStatusOk returns a tuple with the ArtifactoryConnectionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryConnectionStatus

`func (o *ApiOwnerArtifactoryConnectionDTO) SetArtifactoryConnectionStatus(v ApiArtifactoryConnectionStatusResponseDTO)`

SetArtifactoryConnectionStatus sets ArtifactoryConnectionStatus field to given value.

### HasArtifactoryConnectionStatus

`func (o *ApiOwnerArtifactoryConnectionDTO) HasArtifactoryConnectionStatus() bool`

HasArtifactoryConnectionStatus returns a boolean if a field has been set.

### GetOwnerDTO

`func (o *ApiOwnerArtifactoryConnectionDTO) GetOwnerDTO() ApiOwnerDTO`

GetOwnerDTO returns the OwnerDTO field if non-nil, zero value otherwise.

### GetOwnerDTOOk

`func (o *ApiOwnerArtifactoryConnectionDTO) GetOwnerDTOOk() (*ApiOwnerDTO, bool)`

GetOwnerDTOOk returns a tuple with the OwnerDTO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerDTO

`func (o *ApiOwnerArtifactoryConnectionDTO) SetOwnerDTO(v ApiOwnerDTO)`

SetOwnerDTO sets OwnerDTO field to given value.

### HasOwnerDTO

`func (o *ApiOwnerArtifactoryConnectionDTO) HasOwnerDTO() bool`

HasOwnerDTO returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


