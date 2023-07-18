# ApiLicenseLegalMetadataDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsMulti** | Pointer to **bool** |  | [optional] 
**LicenseId** | Pointer to **string** |  | [optional] 
**LicenseName** | Pointer to **string** |  | [optional] 
**LicenseText** | Pointer to **string** |  | [optional] 
**Obligations** | Pointer to [**[]LicenseObligationDTO**](LicenseObligationDTO.md) |  | [optional] 
**SingleLicenseIds** | Pointer to **[]string** |  | [optional] 
**ThreatGroup** | Pointer to [**LicenseThreatGroupDTO**](LicenseThreatGroupDTO.md) |  | [optional] 

## Methods

### NewApiLicenseLegalMetadataDTO

`func NewApiLicenseLegalMetadataDTO() *ApiLicenseLegalMetadataDTO`

NewApiLicenseLegalMetadataDTO instantiates a new ApiLicenseLegalMetadataDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiLicenseLegalMetadataDTOWithDefaults

`func NewApiLicenseLegalMetadataDTOWithDefaults() *ApiLicenseLegalMetadataDTO`

NewApiLicenseLegalMetadataDTOWithDefaults instantiates a new ApiLicenseLegalMetadataDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsMulti

`func (o *ApiLicenseLegalMetadataDTO) GetIsMulti() bool`

GetIsMulti returns the IsMulti field if non-nil, zero value otherwise.

### GetIsMultiOk

`func (o *ApiLicenseLegalMetadataDTO) GetIsMultiOk() (*bool, bool)`

GetIsMultiOk returns a tuple with the IsMulti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMulti

`func (o *ApiLicenseLegalMetadataDTO) SetIsMulti(v bool)`

SetIsMulti sets IsMulti field to given value.

### HasIsMulti

`func (o *ApiLicenseLegalMetadataDTO) HasIsMulti() bool`

HasIsMulti returns a boolean if a field has been set.

### GetLicenseId

`func (o *ApiLicenseLegalMetadataDTO) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *ApiLicenseLegalMetadataDTO) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *ApiLicenseLegalMetadataDTO) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *ApiLicenseLegalMetadataDTO) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetLicenseName

`func (o *ApiLicenseLegalMetadataDTO) GetLicenseName() string`

GetLicenseName returns the LicenseName field if non-nil, zero value otherwise.

### GetLicenseNameOk

`func (o *ApiLicenseLegalMetadataDTO) GetLicenseNameOk() (*string, bool)`

GetLicenseNameOk returns a tuple with the LicenseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseName

`func (o *ApiLicenseLegalMetadataDTO) SetLicenseName(v string)`

SetLicenseName sets LicenseName field to given value.

### HasLicenseName

`func (o *ApiLicenseLegalMetadataDTO) HasLicenseName() bool`

HasLicenseName returns a boolean if a field has been set.

### GetLicenseText

`func (o *ApiLicenseLegalMetadataDTO) GetLicenseText() string`

GetLicenseText returns the LicenseText field if non-nil, zero value otherwise.

### GetLicenseTextOk

`func (o *ApiLicenseLegalMetadataDTO) GetLicenseTextOk() (*string, bool)`

GetLicenseTextOk returns a tuple with the LicenseText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseText

`func (o *ApiLicenseLegalMetadataDTO) SetLicenseText(v string)`

SetLicenseText sets LicenseText field to given value.

### HasLicenseText

`func (o *ApiLicenseLegalMetadataDTO) HasLicenseText() bool`

HasLicenseText returns a boolean if a field has been set.

### GetObligations

`func (o *ApiLicenseLegalMetadataDTO) GetObligations() []LicenseObligationDTO`

GetObligations returns the Obligations field if non-nil, zero value otherwise.

### GetObligationsOk

`func (o *ApiLicenseLegalMetadataDTO) GetObligationsOk() (*[]LicenseObligationDTO, bool)`

GetObligationsOk returns a tuple with the Obligations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObligations

`func (o *ApiLicenseLegalMetadataDTO) SetObligations(v []LicenseObligationDTO)`

SetObligations sets Obligations field to given value.

### HasObligations

`func (o *ApiLicenseLegalMetadataDTO) HasObligations() bool`

HasObligations returns a boolean if a field has been set.

### GetSingleLicenseIds

`func (o *ApiLicenseLegalMetadataDTO) GetSingleLicenseIds() []string`

GetSingleLicenseIds returns the SingleLicenseIds field if non-nil, zero value otherwise.

### GetSingleLicenseIdsOk

`func (o *ApiLicenseLegalMetadataDTO) GetSingleLicenseIdsOk() (*[]string, bool)`

GetSingleLicenseIdsOk returns a tuple with the SingleLicenseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleLicenseIds

`func (o *ApiLicenseLegalMetadataDTO) SetSingleLicenseIds(v []string)`

SetSingleLicenseIds sets SingleLicenseIds field to given value.

### HasSingleLicenseIds

`func (o *ApiLicenseLegalMetadataDTO) HasSingleLicenseIds() bool`

HasSingleLicenseIds returns a boolean if a field has been set.

### GetThreatGroup

`func (o *ApiLicenseLegalMetadataDTO) GetThreatGroup() LicenseThreatGroupDTO`

GetThreatGroup returns the ThreatGroup field if non-nil, zero value otherwise.

### GetThreatGroupOk

`func (o *ApiLicenseLegalMetadataDTO) GetThreatGroupOk() (*LicenseThreatGroupDTO, bool)`

GetThreatGroupOk returns a tuple with the ThreatGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatGroup

`func (o *ApiLicenseLegalMetadataDTO) SetThreatGroup(v LicenseThreatGroupDTO)`

SetThreatGroup sets ThreatGroup field to given value.

### HasThreatGroup

`func (o *ApiLicenseLegalMetadataDTO) HasThreatGroup() bool`

HasThreatGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


