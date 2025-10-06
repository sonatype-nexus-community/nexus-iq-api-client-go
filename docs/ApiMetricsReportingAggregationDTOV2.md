# ApiMetricsReportingAggregationDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveredCounts** | Pointer to **map[string]map[string]int32** |  | [optional] 
**EvaluationCount** | Pointer to **int32** |  | [optional] 
**FixedCounts** | Pointer to **map[string]map[string]int32** |  | [optional] 
**MttrCriticalThreat** | Pointer to **int64** |  | [optional] 
**MttrLowThreat** | Pointer to **int64** |  | [optional] 
**MttrModerateThreat** | Pointer to **int64** |  | [optional] 
**MttrSevereThreat** | Pointer to **int64** |  | [optional] 
**OpenCountsAtTimePeriodEnd** | Pointer to **map[string]map[string]int32** |  | [optional] 
**TimePeriodStart** | Pointer to **string** |  | [optional] 
**WaivedCounts** | Pointer to **map[string]map[string]int32** |  | [optional] 

## Methods

### NewApiMetricsReportingAggregationDTOV2

`func NewApiMetricsReportingAggregationDTOV2() *ApiMetricsReportingAggregationDTOV2`

NewApiMetricsReportingAggregationDTOV2 instantiates a new ApiMetricsReportingAggregationDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMetricsReportingAggregationDTOV2WithDefaults

`func NewApiMetricsReportingAggregationDTOV2WithDefaults() *ApiMetricsReportingAggregationDTOV2`

NewApiMetricsReportingAggregationDTOV2WithDefaults instantiates a new ApiMetricsReportingAggregationDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveredCounts

`func (o *ApiMetricsReportingAggregationDTOV2) GetDiscoveredCounts() map[string]map[string]int32`

GetDiscoveredCounts returns the DiscoveredCounts field if non-nil, zero value otherwise.

### GetDiscoveredCountsOk

`func (o *ApiMetricsReportingAggregationDTOV2) GetDiscoveredCountsOk() (*map[string]map[string]int32, bool)`

GetDiscoveredCountsOk returns a tuple with the DiscoveredCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredCounts

`func (o *ApiMetricsReportingAggregationDTOV2) SetDiscoveredCounts(v map[string]map[string]int32)`

SetDiscoveredCounts sets DiscoveredCounts field to given value.

### HasDiscoveredCounts

`func (o *ApiMetricsReportingAggregationDTOV2) HasDiscoveredCounts() bool`

HasDiscoveredCounts returns a boolean if a field has been set.

### GetEvaluationCount

`func (o *ApiMetricsReportingAggregationDTOV2) GetEvaluationCount() int32`

GetEvaluationCount returns the EvaluationCount field if non-nil, zero value otherwise.

### GetEvaluationCountOk

`func (o *ApiMetricsReportingAggregationDTOV2) GetEvaluationCountOk() (*int32, bool)`

GetEvaluationCountOk returns a tuple with the EvaluationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationCount

`func (o *ApiMetricsReportingAggregationDTOV2) SetEvaluationCount(v int32)`

SetEvaluationCount sets EvaluationCount field to given value.

### HasEvaluationCount

`func (o *ApiMetricsReportingAggregationDTOV2) HasEvaluationCount() bool`

HasEvaluationCount returns a boolean if a field has been set.

### GetFixedCounts

`func (o *ApiMetricsReportingAggregationDTOV2) GetFixedCounts() map[string]map[string]int32`

GetFixedCounts returns the FixedCounts field if non-nil, zero value otherwise.

### GetFixedCountsOk

`func (o *ApiMetricsReportingAggregationDTOV2) GetFixedCountsOk() (*map[string]map[string]int32, bool)`

GetFixedCountsOk returns a tuple with the FixedCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedCounts

`func (o *ApiMetricsReportingAggregationDTOV2) SetFixedCounts(v map[string]map[string]int32)`

SetFixedCounts sets FixedCounts field to given value.

### HasFixedCounts

`func (o *ApiMetricsReportingAggregationDTOV2) HasFixedCounts() bool`

HasFixedCounts returns a boolean if a field has been set.

### GetMttrCriticalThreat

`func (o *ApiMetricsReportingAggregationDTOV2) GetMttrCriticalThreat() int64`

GetMttrCriticalThreat returns the MttrCriticalThreat field if non-nil, zero value otherwise.

### GetMttrCriticalThreatOk

`func (o *ApiMetricsReportingAggregationDTOV2) GetMttrCriticalThreatOk() (*int64, bool)`

GetMttrCriticalThreatOk returns a tuple with the MttrCriticalThreat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMttrCriticalThreat

`func (o *ApiMetricsReportingAggregationDTOV2) SetMttrCriticalThreat(v int64)`

SetMttrCriticalThreat sets MttrCriticalThreat field to given value.

### HasMttrCriticalThreat

`func (o *ApiMetricsReportingAggregationDTOV2) HasMttrCriticalThreat() bool`

HasMttrCriticalThreat returns a boolean if a field has been set.

### GetMttrLowThreat

`func (o *ApiMetricsReportingAggregationDTOV2) GetMttrLowThreat() int64`

GetMttrLowThreat returns the MttrLowThreat field if non-nil, zero value otherwise.

### GetMttrLowThreatOk

`func (o *ApiMetricsReportingAggregationDTOV2) GetMttrLowThreatOk() (*int64, bool)`

GetMttrLowThreatOk returns a tuple with the MttrLowThreat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMttrLowThreat

`func (o *ApiMetricsReportingAggregationDTOV2) SetMttrLowThreat(v int64)`

SetMttrLowThreat sets MttrLowThreat field to given value.

### HasMttrLowThreat

`func (o *ApiMetricsReportingAggregationDTOV2) HasMttrLowThreat() bool`

HasMttrLowThreat returns a boolean if a field has been set.

### GetMttrModerateThreat

`func (o *ApiMetricsReportingAggregationDTOV2) GetMttrModerateThreat() int64`

GetMttrModerateThreat returns the MttrModerateThreat field if non-nil, zero value otherwise.

### GetMttrModerateThreatOk

`func (o *ApiMetricsReportingAggregationDTOV2) GetMttrModerateThreatOk() (*int64, bool)`

GetMttrModerateThreatOk returns a tuple with the MttrModerateThreat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMttrModerateThreat

`func (o *ApiMetricsReportingAggregationDTOV2) SetMttrModerateThreat(v int64)`

SetMttrModerateThreat sets MttrModerateThreat field to given value.

### HasMttrModerateThreat

`func (o *ApiMetricsReportingAggregationDTOV2) HasMttrModerateThreat() bool`

HasMttrModerateThreat returns a boolean if a field has been set.

### GetMttrSevereThreat

`func (o *ApiMetricsReportingAggregationDTOV2) GetMttrSevereThreat() int64`

GetMttrSevereThreat returns the MttrSevereThreat field if non-nil, zero value otherwise.

### GetMttrSevereThreatOk

`func (o *ApiMetricsReportingAggregationDTOV2) GetMttrSevereThreatOk() (*int64, bool)`

GetMttrSevereThreatOk returns a tuple with the MttrSevereThreat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMttrSevereThreat

`func (o *ApiMetricsReportingAggregationDTOV2) SetMttrSevereThreat(v int64)`

SetMttrSevereThreat sets MttrSevereThreat field to given value.

### HasMttrSevereThreat

`func (o *ApiMetricsReportingAggregationDTOV2) HasMttrSevereThreat() bool`

HasMttrSevereThreat returns a boolean if a field has been set.

### GetOpenCountsAtTimePeriodEnd

`func (o *ApiMetricsReportingAggregationDTOV2) GetOpenCountsAtTimePeriodEnd() map[string]map[string]int32`

GetOpenCountsAtTimePeriodEnd returns the OpenCountsAtTimePeriodEnd field if non-nil, zero value otherwise.

### GetOpenCountsAtTimePeriodEndOk

`func (o *ApiMetricsReportingAggregationDTOV2) GetOpenCountsAtTimePeriodEndOk() (*map[string]map[string]int32, bool)`

GetOpenCountsAtTimePeriodEndOk returns a tuple with the OpenCountsAtTimePeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenCountsAtTimePeriodEnd

`func (o *ApiMetricsReportingAggregationDTOV2) SetOpenCountsAtTimePeriodEnd(v map[string]map[string]int32)`

SetOpenCountsAtTimePeriodEnd sets OpenCountsAtTimePeriodEnd field to given value.

### HasOpenCountsAtTimePeriodEnd

`func (o *ApiMetricsReportingAggregationDTOV2) HasOpenCountsAtTimePeriodEnd() bool`

HasOpenCountsAtTimePeriodEnd returns a boolean if a field has been set.

### GetTimePeriodStart

`func (o *ApiMetricsReportingAggregationDTOV2) GetTimePeriodStart() string`

GetTimePeriodStart returns the TimePeriodStart field if non-nil, zero value otherwise.

### GetTimePeriodStartOk

`func (o *ApiMetricsReportingAggregationDTOV2) GetTimePeriodStartOk() (*string, bool)`

GetTimePeriodStartOk returns a tuple with the TimePeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimePeriodStart

`func (o *ApiMetricsReportingAggregationDTOV2) SetTimePeriodStart(v string)`

SetTimePeriodStart sets TimePeriodStart field to given value.

### HasTimePeriodStart

`func (o *ApiMetricsReportingAggregationDTOV2) HasTimePeriodStart() bool`

HasTimePeriodStart returns a boolean if a field has been set.

### GetWaivedCounts

`func (o *ApiMetricsReportingAggregationDTOV2) GetWaivedCounts() map[string]map[string]int32`

GetWaivedCounts returns the WaivedCounts field if non-nil, zero value otherwise.

### GetWaivedCountsOk

`func (o *ApiMetricsReportingAggregationDTOV2) GetWaivedCountsOk() (*map[string]map[string]int32, bool)`

GetWaivedCountsOk returns a tuple with the WaivedCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaivedCounts

`func (o *ApiMetricsReportingAggregationDTOV2) SetWaivedCounts(v map[string]map[string]int32)`

SetWaivedCounts sets WaivedCounts field to given value.

### HasWaivedCounts

`func (o *ApiMetricsReportingAggregationDTOV2) HasWaivedCounts() bool`

HasWaivedCounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


