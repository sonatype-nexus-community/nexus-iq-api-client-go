# ApiConstraintViolationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConstraintId** | Pointer to **string** |  | [optional] 
**ConstraintName** | Pointer to **string** |  | [optional] 
**Reasons** | Pointer to [**[]ApiConstraintViolationReasonDTO**](ApiConstraintViolationReasonDTO.md) |  | [optional] 

## Methods

### NewApiConstraintViolationDTO

`func NewApiConstraintViolationDTO() *ApiConstraintViolationDTO`

NewApiConstraintViolationDTO instantiates a new ApiConstraintViolationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiConstraintViolationDTOWithDefaults

`func NewApiConstraintViolationDTOWithDefaults() *ApiConstraintViolationDTO`

NewApiConstraintViolationDTOWithDefaults instantiates a new ApiConstraintViolationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraintId

`func (o *ApiConstraintViolationDTO) GetConstraintId() string`

GetConstraintId returns the ConstraintId field if non-nil, zero value otherwise.

### GetConstraintIdOk

`func (o *ApiConstraintViolationDTO) GetConstraintIdOk() (*string, bool)`

GetConstraintIdOk returns a tuple with the ConstraintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintId

`func (o *ApiConstraintViolationDTO) SetConstraintId(v string)`

SetConstraintId sets ConstraintId field to given value.

### HasConstraintId

`func (o *ApiConstraintViolationDTO) HasConstraintId() bool`

HasConstraintId returns a boolean if a field has been set.

### GetConstraintName

`func (o *ApiConstraintViolationDTO) GetConstraintName() string`

GetConstraintName returns the ConstraintName field if non-nil, zero value otherwise.

### GetConstraintNameOk

`func (o *ApiConstraintViolationDTO) GetConstraintNameOk() (*string, bool)`

GetConstraintNameOk returns a tuple with the ConstraintName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintName

`func (o *ApiConstraintViolationDTO) SetConstraintName(v string)`

SetConstraintName sets ConstraintName field to given value.

### HasConstraintName

`func (o *ApiConstraintViolationDTO) HasConstraintName() bool`

HasConstraintName returns a boolean if a field has been set.

### GetReasons

`func (o *ApiConstraintViolationDTO) GetReasons() []ApiConstraintViolationReasonDTO`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *ApiConstraintViolationDTO) GetReasonsOk() (*[]ApiConstraintViolationReasonDTO, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *ApiConstraintViolationDTO) SetReasons(v []ApiConstraintViolationReasonDTO)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *ApiConstraintViolationDTO) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


