# SbomComponentListDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]SbomComponentDTO**](SbomComponentDTO.md) |  | [optional] 
**TotalResultsCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewSbomComponentListDTO

`func NewSbomComponentListDTO() *SbomComponentListDTO`

NewSbomComponentListDTO instantiates a new SbomComponentListDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSbomComponentListDTOWithDefaults

`func NewSbomComponentListDTOWithDefaults() *SbomComponentListDTO`

NewSbomComponentListDTOWithDefaults instantiates a new SbomComponentListDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *SbomComponentListDTO) GetResults() []SbomComponentDTO`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SbomComponentListDTO) GetResultsOk() (*[]SbomComponentDTO, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SbomComponentListDTO) SetResults(v []SbomComponentDTO)`

SetResults sets Results field to given value.

### HasResults

`func (o *SbomComponentListDTO) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetTotalResultsCount

`func (o *SbomComponentListDTO) GetTotalResultsCount() int32`

GetTotalResultsCount returns the TotalResultsCount field if non-nil, zero value otherwise.

### GetTotalResultsCountOk

`func (o *SbomComponentListDTO) GetTotalResultsCountOk() (*int32, bool)`

GetTotalResultsCountOk returns a tuple with the TotalResultsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResultsCount

`func (o *SbomComponentListDTO) SetTotalResultsCount(v int32)`

SetTotalResultsCount sets TotalResultsCount field to given value.

### HasTotalResultsCount

`func (o *SbomComponentListDTO) HasTotalResultsCount() bool`

HasTotalResultsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


