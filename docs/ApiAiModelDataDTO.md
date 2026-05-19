# ApiAiModelDataDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTypes** | Pointer to [**[]ApiAiModelContentTypeDTO**](ApiAiModelContentTypeDTO.md) |  | [optional] 
**DerivedFromComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**DerivedFromSimilarityScore** | Pointer to **float64** |  | [optional] 

## Methods

### NewApiAiModelDataDTO

`func NewApiAiModelDataDTO() *ApiAiModelDataDTO`

NewApiAiModelDataDTO instantiates a new ApiAiModelDataDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAiModelDataDTOWithDefaults

`func NewApiAiModelDataDTOWithDefaults() *ApiAiModelDataDTO`

NewApiAiModelDataDTOWithDefaults instantiates a new ApiAiModelDataDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTypes

`func (o *ApiAiModelDataDTO) GetContentTypes() []ApiAiModelContentTypeDTO`

GetContentTypes returns the ContentTypes field if non-nil, zero value otherwise.

### GetContentTypesOk

`func (o *ApiAiModelDataDTO) GetContentTypesOk() (*[]ApiAiModelContentTypeDTO, bool)`

GetContentTypesOk returns a tuple with the ContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTypes

`func (o *ApiAiModelDataDTO) SetContentTypes(v []ApiAiModelContentTypeDTO)`

SetContentTypes sets ContentTypes field to given value.

### HasContentTypes

`func (o *ApiAiModelDataDTO) HasContentTypes() bool`

HasContentTypes returns a boolean if a field has been set.

### GetDerivedFromComponentIdentifier

`func (o *ApiAiModelDataDTO) GetDerivedFromComponentIdentifier() ApiComponentIdentifierDTOV2`

GetDerivedFromComponentIdentifier returns the DerivedFromComponentIdentifier field if non-nil, zero value otherwise.

### GetDerivedFromComponentIdentifierOk

`func (o *ApiAiModelDataDTO) GetDerivedFromComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetDerivedFromComponentIdentifierOk returns a tuple with the DerivedFromComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFromComponentIdentifier

`func (o *ApiAiModelDataDTO) SetDerivedFromComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetDerivedFromComponentIdentifier sets DerivedFromComponentIdentifier field to given value.

### HasDerivedFromComponentIdentifier

`func (o *ApiAiModelDataDTO) HasDerivedFromComponentIdentifier() bool`

HasDerivedFromComponentIdentifier returns a boolean if a field has been set.

### GetDerivedFromSimilarityScore

`func (o *ApiAiModelDataDTO) GetDerivedFromSimilarityScore() float64`

GetDerivedFromSimilarityScore returns the DerivedFromSimilarityScore field if non-nil, zero value otherwise.

### GetDerivedFromSimilarityScoreOk

`func (o *ApiAiModelDataDTO) GetDerivedFromSimilarityScoreOk() (*float64, bool)`

GetDerivedFromSimilarityScoreOk returns a tuple with the DerivedFromSimilarityScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDerivedFromSimilarityScore

`func (o *ApiAiModelDataDTO) SetDerivedFromSimilarityScore(v float64)`

SetDerivedFromSimilarityScore sets DerivedFromSimilarityScore field to given value.

### HasDerivedFromSimilarityScore

`func (o *ApiAiModelDataDTO) HasDerivedFromSimilarityScore() bool`

HasDerivedFromSimilarityScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


