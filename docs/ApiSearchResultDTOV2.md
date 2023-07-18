# ApiSearchResultDTOV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **string** |  | [optional] 
**ApplicationName** | Pointer to **string** |  | [optional] 
**ComponentIdentifier** | Pointer to [**ApiComponentIdentifierDTOV2**](ApiComponentIdentifierDTOV2.md) |  | [optional] 
**DependencyData** | Pointer to [**ApiDependencyDataDTO**](ApiDependencyDataDTO.md) |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**PackageUrl** | Pointer to **string** |  | [optional] 
**ReportHtmlUrl** | Pointer to **string** |  | [optional] 
**ReportUrl** | Pointer to **string** |  | [optional] 
**ThreatLevel** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiSearchResultDTOV2

`func NewApiSearchResultDTOV2() *ApiSearchResultDTOV2`

NewApiSearchResultDTOV2 instantiates a new ApiSearchResultDTOV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSearchResultDTOV2WithDefaults

`func NewApiSearchResultDTOV2WithDefaults() *ApiSearchResultDTOV2`

NewApiSearchResultDTOV2WithDefaults instantiates a new ApiSearchResultDTOV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApiSearchResultDTOV2) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApiSearchResultDTOV2) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApiSearchResultDTOV2) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApiSearchResultDTOV2) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationName

`func (o *ApiSearchResultDTOV2) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *ApiSearchResultDTOV2) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *ApiSearchResultDTOV2) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *ApiSearchResultDTOV2) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetComponentIdentifier

`func (o *ApiSearchResultDTOV2) GetComponentIdentifier() ApiComponentIdentifierDTOV2`

GetComponentIdentifier returns the ComponentIdentifier field if non-nil, zero value otherwise.

### GetComponentIdentifierOk

`func (o *ApiSearchResultDTOV2) GetComponentIdentifierOk() (*ApiComponentIdentifierDTOV2, bool)`

GetComponentIdentifierOk returns a tuple with the ComponentIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentIdentifier

`func (o *ApiSearchResultDTOV2) SetComponentIdentifier(v ApiComponentIdentifierDTOV2)`

SetComponentIdentifier sets ComponentIdentifier field to given value.

### HasComponentIdentifier

`func (o *ApiSearchResultDTOV2) HasComponentIdentifier() bool`

HasComponentIdentifier returns a boolean if a field has been set.

### GetDependencyData

`func (o *ApiSearchResultDTOV2) GetDependencyData() ApiDependencyDataDTO`

GetDependencyData returns the DependencyData field if non-nil, zero value otherwise.

### GetDependencyDataOk

`func (o *ApiSearchResultDTOV2) GetDependencyDataOk() (*ApiDependencyDataDTO, bool)`

GetDependencyDataOk returns a tuple with the DependencyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyData

`func (o *ApiSearchResultDTOV2) SetDependencyData(v ApiDependencyDataDTO)`

SetDependencyData sets DependencyData field to given value.

### HasDependencyData

`func (o *ApiSearchResultDTOV2) HasDependencyData() bool`

HasDependencyData returns a boolean if a field has been set.

### GetHash

`func (o *ApiSearchResultDTOV2) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *ApiSearchResultDTOV2) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *ApiSearchResultDTOV2) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *ApiSearchResultDTOV2) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetPackageUrl

`func (o *ApiSearchResultDTOV2) GetPackageUrl() string`

GetPackageUrl returns the PackageUrl field if non-nil, zero value otherwise.

### GetPackageUrlOk

`func (o *ApiSearchResultDTOV2) GetPackageUrlOk() (*string, bool)`

GetPackageUrlOk returns a tuple with the PackageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageUrl

`func (o *ApiSearchResultDTOV2) SetPackageUrl(v string)`

SetPackageUrl sets PackageUrl field to given value.

### HasPackageUrl

`func (o *ApiSearchResultDTOV2) HasPackageUrl() bool`

HasPackageUrl returns a boolean if a field has been set.

### GetReportHtmlUrl

`func (o *ApiSearchResultDTOV2) GetReportHtmlUrl() string`

GetReportHtmlUrl returns the ReportHtmlUrl field if non-nil, zero value otherwise.

### GetReportHtmlUrlOk

`func (o *ApiSearchResultDTOV2) GetReportHtmlUrlOk() (*string, bool)`

GetReportHtmlUrlOk returns a tuple with the ReportHtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportHtmlUrl

`func (o *ApiSearchResultDTOV2) SetReportHtmlUrl(v string)`

SetReportHtmlUrl sets ReportHtmlUrl field to given value.

### HasReportHtmlUrl

`func (o *ApiSearchResultDTOV2) HasReportHtmlUrl() bool`

HasReportHtmlUrl returns a boolean if a field has been set.

### GetReportUrl

`func (o *ApiSearchResultDTOV2) GetReportUrl() string`

GetReportUrl returns the ReportUrl field if non-nil, zero value otherwise.

### GetReportUrlOk

`func (o *ApiSearchResultDTOV2) GetReportUrlOk() (*string, bool)`

GetReportUrlOk returns a tuple with the ReportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportUrl

`func (o *ApiSearchResultDTOV2) SetReportUrl(v string)`

SetReportUrl sets ReportUrl field to given value.

### HasReportUrl

`func (o *ApiSearchResultDTOV2) HasReportUrl() bool`

HasReportUrl returns a boolean if a field has been set.

### GetThreatLevel

`func (o *ApiSearchResultDTOV2) GetThreatLevel() int32`

GetThreatLevel returns the ThreatLevel field if non-nil, zero value otherwise.

### GetThreatLevelOk

`func (o *ApiSearchResultDTOV2) GetThreatLevelOk() (*int32, bool)`

GetThreatLevelOk returns a tuple with the ThreatLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatLevel

`func (o *ApiSearchResultDTOV2) SetThreatLevel(v int32)`

SetThreatLevel sets ThreatLevel field to given value.

### HasThreatLevel

`func (o *ApiSearchResultDTOV2) HasThreatLevel() bool`

HasThreatLevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


