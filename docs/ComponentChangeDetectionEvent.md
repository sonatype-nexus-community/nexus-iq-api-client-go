# ComponentChangeDetectionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedTime** | Pointer to **time.Time** |  | [optional] 
**ComponentEvaluationData** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Purl** | Pointer to **string** |  | [optional] 

## Methods

### NewComponentChangeDetectionEvent

`func NewComponentChangeDetectionEvent() *ComponentChangeDetectionEvent`

NewComponentChangeDetectionEvent instantiates a new ComponentChangeDetectionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentChangeDetectionEventWithDefaults

`func NewComponentChangeDetectionEventWithDefaults() *ComponentChangeDetectionEvent`

NewComponentChangeDetectionEventWithDefaults instantiates a new ComponentChangeDetectionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedTime

`func (o *ComponentChangeDetectionEvent) GetAddedTime() time.Time`

GetAddedTime returns the AddedTime field if non-nil, zero value otherwise.

### GetAddedTimeOk

`func (o *ComponentChangeDetectionEvent) GetAddedTimeOk() (*time.Time, bool)`

GetAddedTimeOk returns a tuple with the AddedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedTime

`func (o *ComponentChangeDetectionEvent) SetAddedTime(v time.Time)`

SetAddedTime sets AddedTime field to given value.

### HasAddedTime

`func (o *ComponentChangeDetectionEvent) HasAddedTime() bool`

HasAddedTime returns a boolean if a field has been set.

### GetComponentEvaluationData

`func (o *ComponentChangeDetectionEvent) GetComponentEvaluationData() string`

GetComponentEvaluationData returns the ComponentEvaluationData field if non-nil, zero value otherwise.

### GetComponentEvaluationDataOk

`func (o *ComponentChangeDetectionEvent) GetComponentEvaluationDataOk() (*string, bool)`

GetComponentEvaluationDataOk returns a tuple with the ComponentEvaluationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentEvaluationData

`func (o *ComponentChangeDetectionEvent) SetComponentEvaluationData(v string)`

SetComponentEvaluationData sets ComponentEvaluationData field to given value.

### HasComponentEvaluationData

`func (o *ComponentChangeDetectionEvent) HasComponentEvaluationData() bool`

HasComponentEvaluationData returns a boolean if a field has been set.

### GetId

`func (o *ComponentChangeDetectionEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ComponentChangeDetectionEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ComponentChangeDetectionEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ComponentChangeDetectionEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPurl

`func (o *ComponentChangeDetectionEvent) GetPurl() string`

GetPurl returns the Purl field if non-nil, zero value otherwise.

### GetPurlOk

`func (o *ComponentChangeDetectionEvent) GetPurlOk() (*string, bool)`

GetPurlOk returns a tuple with the Purl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurl

`func (o *ComponentChangeDetectionEvent) SetPurl(v string)`

SetPurl sets Purl field to given value.

### HasPurl

`func (o *ComponentChangeDetectionEvent) HasPurl() bool`

HasPurl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


