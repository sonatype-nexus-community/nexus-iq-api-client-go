# ImportSbomRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | The internal id of the application. | 
**ApplicationVersion** | Pointer to **string** | The SBOM version. | [optional] 
**File** | ***os.File** | Your SBOM. | 

## Methods

### NewImportSbomRequest

`func NewImportSbomRequest(applicationId string, file *os.File, ) *ImportSbomRequest`

NewImportSbomRequest instantiates a new ImportSbomRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportSbomRequestWithDefaults

`func NewImportSbomRequestWithDefaults() *ImportSbomRequest`

NewImportSbomRequestWithDefaults instantiates a new ImportSbomRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ImportSbomRequest) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ImportSbomRequest) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ImportSbomRequest) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetApplicationVersion

`func (o *ImportSbomRequest) GetApplicationVersion() string`

GetApplicationVersion returns the ApplicationVersion field if non-nil, zero value otherwise.

### GetApplicationVersionOk

`func (o *ImportSbomRequest) GetApplicationVersionOk() (*string, bool)`

GetApplicationVersionOk returns a tuple with the ApplicationVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationVersion

`func (o *ImportSbomRequest) SetApplicationVersion(v string)`

SetApplicationVersion sets ApplicationVersion field to given value.

### HasApplicationVersion

`func (o *ImportSbomRequest) HasApplicationVersion() bool`

HasApplicationVersion returns a boolean if a field has been set.

### GetFile

`func (o *ImportSbomRequest) GetFile() *os.File`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ImportSbomRequest) GetFileOk() (**os.File, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ImportSbomRequest) SetFile(v *os.File)`

SetFile sets File field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


