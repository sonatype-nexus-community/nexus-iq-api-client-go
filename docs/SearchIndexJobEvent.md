# SearchIndexJobEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**EventCode** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**SearchIndexJobId** | Pointer to **string** |  | [optional] 
**Seq** | Pointer to **int64** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 

## Methods

### NewSearchIndexJobEvent

`func NewSearchIndexJobEvent() *SearchIndexJobEvent`

NewSearchIndexJobEvent instantiates a new SearchIndexJobEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchIndexJobEventWithDefaults

`func NewSearchIndexJobEventWithDefaults() *SearchIndexJobEvent`

NewSearchIndexJobEventWithDefaults instantiates a new SearchIndexJobEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SearchIndexJobEvent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SearchIndexJobEvent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SearchIndexJobEvent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SearchIndexJobEvent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEventCode

`func (o *SearchIndexJobEvent) GetEventCode() string`

GetEventCode returns the EventCode field if non-nil, zero value otherwise.

### GetEventCodeOk

`func (o *SearchIndexJobEvent) GetEventCodeOk() (*string, bool)`

GetEventCodeOk returns a tuple with the EventCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCode

`func (o *SearchIndexJobEvent) SetEventCode(v string)`

SetEventCode sets EventCode field to given value.

### HasEventCode

`func (o *SearchIndexJobEvent) HasEventCode() bool`

HasEventCode returns a boolean if a field has been set.

### GetId

`func (o *SearchIndexJobEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchIndexJobEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchIndexJobEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchIndexJobEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *SearchIndexJobEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SearchIndexJobEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SearchIndexJobEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SearchIndexJobEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSearchIndexJobId

`func (o *SearchIndexJobEvent) GetSearchIndexJobId() string`

GetSearchIndexJobId returns the SearchIndexJobId field if non-nil, zero value otherwise.

### GetSearchIndexJobIdOk

`func (o *SearchIndexJobEvent) GetSearchIndexJobIdOk() (*string, bool)`

GetSearchIndexJobIdOk returns a tuple with the SearchIndexJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchIndexJobId

`func (o *SearchIndexJobEvent) SetSearchIndexJobId(v string)`

SetSearchIndexJobId sets SearchIndexJobId field to given value.

### HasSearchIndexJobId

`func (o *SearchIndexJobEvent) HasSearchIndexJobId() bool`

HasSearchIndexJobId returns a boolean if a field has been set.

### GetSeq

`func (o *SearchIndexJobEvent) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *SearchIndexJobEvent) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *SearchIndexJobEvent) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *SearchIndexJobEvent) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetSeverity

`func (o *SearchIndexJobEvent) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SearchIndexJobEvent) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SearchIndexJobEvent) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *SearchIndexJobEvent) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


