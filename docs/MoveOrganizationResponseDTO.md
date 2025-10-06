# MoveOrganizationResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ValidationError**](ValidationError.md) |  | [optional] 
**Warnings** | Pointer to [**[]ValidationWarning**](ValidationWarning.md) |  | [optional] 

## Methods

### NewMoveOrganizationResponseDTO

`func NewMoveOrganizationResponseDTO() *MoveOrganizationResponseDTO`

NewMoveOrganizationResponseDTO instantiates a new MoveOrganizationResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveOrganizationResponseDTOWithDefaults

`func NewMoveOrganizationResponseDTOWithDefaults() *MoveOrganizationResponseDTO`

NewMoveOrganizationResponseDTOWithDefaults instantiates a new MoveOrganizationResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *MoveOrganizationResponseDTO) GetErrors() []ValidationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MoveOrganizationResponseDTO) GetErrorsOk() (*[]ValidationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MoveOrganizationResponseDTO) SetErrors(v []ValidationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *MoveOrganizationResponseDTO) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetWarnings

`func (o *MoveOrganizationResponseDTO) GetWarnings() []ValidationWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *MoveOrganizationResponseDTO) GetWarningsOk() (*[]ValidationWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *MoveOrganizationResponseDTO) SetWarnings(v []ValidationWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *MoveOrganizationResponseDTO) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


