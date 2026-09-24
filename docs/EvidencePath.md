# EvidencePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Segments** | Pointer to [**[]PathSegment**](PathSegment.md) |  | [optional] 

## Methods

### NewEvidencePath

`func NewEvidencePath() *EvidencePath`

NewEvidencePath instantiates a new EvidencePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEvidencePathWithDefaults

`func NewEvidencePathWithDefaults() *EvidencePath`

NewEvidencePathWithDefaults instantiates a new EvidencePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSegments

`func (o *EvidencePath) GetSegments() []PathSegment`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *EvidencePath) GetSegmentsOk() (*[]PathSegment, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *EvidencePath) SetSegments(v []PathSegment)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *EvidencePath) HasSegments() bool`

HasSegments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


