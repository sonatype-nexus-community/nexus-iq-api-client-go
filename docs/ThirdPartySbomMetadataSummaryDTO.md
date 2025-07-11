# ThirdPartySbomMetadataSummaryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationVersion** | Pointer to **string** |  | [optional] 
**Critical** | Pointer to **int32** |  | [optional] 
**High** | Pointer to **int32** |  | [optional] 
**ImportDate** | Pointer to **time.Time** |  | [optional] 
**IsValid** | Pointer to **bool** |  | [optional] 
**Low** | Pointer to **int32** |  | [optional] 
**Medium** | Pointer to **int32** |  | [optional] 
**None** | Pointer to **int32** |  | [optional] 
**ReleaseStatusPercentage** | Pointer to **float64** |  | [optional] 
**Spec** | Pointer to **string** |  | [optional] 
**SpecVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewThirdPartySbomMetadataSummaryDTO

`func NewThirdPartySbomMetadataSummaryDTO() *ThirdPartySbomMetadataSummaryDTO`

NewThirdPartySbomMetadataSummaryDTO instantiates a new ThirdPartySbomMetadataSummaryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThirdPartySbomMetadataSummaryDTOWithDefaults

`func NewThirdPartySbomMetadataSummaryDTOWithDefaults() *ThirdPartySbomMetadataSummaryDTO`

NewThirdPartySbomMetadataSummaryDTOWithDefaults instantiates a new ThirdPartySbomMetadataSummaryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationVersion

`func (o *ThirdPartySbomMetadataSummaryDTO) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *ThirdPartySbomMetadataSummaryDTO) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *ThirdPartySbomMetadataSummaryDTO) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *ThirdPartySbomMetadataSummaryDTO) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### GetCritical

`func (o *ThirdPartySbomMetadataSummaryDTO) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *ThirdPartySbomMetadataSummaryDTO) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *ThirdPartySbomMetadataSummaryDTO) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *ThirdPartySbomMetadataSummaryDTO) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetHigh

`func (o *ThirdPartySbomMetadataSummaryDTO) GetHigh() int32`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *ThirdPartySbomMetadataSummaryDTO) GetHighOk() (*int32, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *ThirdPartySbomMetadataSummaryDTO) SetHigh(v int32)`

SetHigh sets High field to given value.

### HasHigh

`func (o *ThirdPartySbomMetadataSummaryDTO) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetImportDate

`func (o *ThirdPartySbomMetadataSummaryDTO) GetImportDate() time.Time`

GetImportDate returns the ImportDate field if non-nil, zero value otherwise.

### GetImportDateOk

`func (o *ThirdPartySbomMetadataSummaryDTO) GetImportDateOk() (*time.Time, bool)`

GetImportDateOk returns a tuple with the ImportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportDate

`func (o *ThirdPartySbomMetadataSummaryDTO) SetImportDate(v time.Time)`

SetImportDate sets ImportDate field to given value.

### HasImportDate

`func (o *ThirdPartySbomMetadataSummaryDTO) HasImportDate() bool`

HasImportDate returns a boolean if a field has been set.

### GetIsValid

`func (o *ThirdPartySbomMetadataSummaryDTO) GetIsValid() bool`

GetIsValid returns the IsValid field if non-nil, zero value otherwise.

### GetIsValidOk

`func (o *ThirdPartySbomMetadataSummaryDTO) GetIsValidOk() (*bool, bool)`

GetIsValidOk returns a tuple with the IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValid

`func (o *ThirdPartySbomMetadataSummaryDTO) SetIsValid(v bool)`

SetIsValid sets IsValid field to given value.

### HasIsValid

`func (o *ThirdPartySbomMetadataSummaryDTO) HasIsValid() bool`

HasIsValid returns a boolean if a field has been set.

### GetLow

`func (o *ThirdPartySbomMetadataSummaryDTO) GetLow() int32`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *ThirdPartySbomMetadataSummaryDTO) GetLowOk() (*int32, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *ThirdPartySbomMetadataSummaryDTO) SetLow(v int32)`

SetLow sets Low field to given value.

### HasLow

`func (o *ThirdPartySbomMetadataSummaryDTO) HasLow() bool`

HasLow returns a boolean if a field has been set.

### GetMedium

`func (o *ThirdPartySbomMetadataSummaryDTO) GetMedium() int32`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *ThirdPartySbomMetadataSummaryDTO) GetMediumOk() (*int32, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *ThirdPartySbomMetadataSummaryDTO) SetMedium(v int32)`

SetMedium sets Medium field to given value.

### HasMedium

`func (o *ThirdPartySbomMetadataSummaryDTO) HasMedium() bool`

HasMedium returns a boolean if a field has been set.

### GetNone

`func (o *ThirdPartySbomMetadataSummaryDTO) GetNone() int32`

GetNone returns the None field if non-nil, zero value otherwise.

### GetNoneOk

`func (o *ThirdPartySbomMetadataSummaryDTO) GetNoneOk() (*int32, bool)`

GetNoneOk returns a tuple with the None field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNone

`func (o *ThirdPartySbomMetadataSummaryDTO) SetNone(v int32)`

SetNone sets None field to given value.

### HasNone

`func (o *ThirdPartySbomMetadataSummaryDTO) HasNone() bool`

HasNone returns a boolean if a field has been set.

### GetReleaseStatusPercentage

`func (o *ThirdPartySbomMetadataSummaryDTO) GetReleaseStatusPercentage() float64`

GetReleaseStatusPercentage returns the ReleaseStatusPercentage field if non-nil, zero value otherwise.

### GetReleaseStatusPercentageOk

`func (o *ThirdPartySbomMetadataSummaryDTO) GetReleaseStatusPercentageOk() (*float64, bool)`

GetReleaseStatusPercentageOk returns a tuple with the ReleaseStatusPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseStatusPercentage

`func (o *ThirdPartySbomMetadataSummaryDTO) SetReleaseStatusPercentage(v float64)`

SetReleaseStatusPercentage sets ReleaseStatusPercentage field to given value.

### HasReleaseStatusPercentage

`func (o *ThirdPartySbomMetadataSummaryDTO) HasReleaseStatusPercentage() bool`

HasReleaseStatusPercentage returns a boolean if a field has been set.

### GetSpec

`func (o *ThirdPartySbomMetadataSummaryDTO) GetSpec() string`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *ThirdPartySbomMetadataSummaryDTO) GetSpecOk() (*string, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *ThirdPartySbomMetadataSummaryDTO) SetSpec(v string)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *ThirdPartySbomMetadataSummaryDTO) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

### GetSpecVersion

`func (o *ThirdPartySbomMetadataSummaryDTO) GetSpecVersion() string`

GetSpecVersion returns the SpecVersion field if non-nil, zero value otherwise.

### GetSpecVersionOk

`func (o *ThirdPartySbomMetadataSummaryDTO) GetSpecVersionOk() (*string, bool)`

GetSpecVersionOk returns a tuple with the SpecVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecVersion

`func (o *ThirdPartySbomMetadataSummaryDTO) SetSpecVersion(v string)`

SetSpecVersion sets SpecVersion field to given value.

### HasSpecVersion

`func (o *ThirdPartySbomMetadataSummaryDTO) HasSpecVersion() bool`

HasSpecVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


