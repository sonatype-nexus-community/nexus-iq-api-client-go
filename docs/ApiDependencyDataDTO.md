# ApiDependencyDataDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DirectDependency** | Pointer to **bool** |  | [optional] 
**InnerSource** | Pointer to **bool** |  | [optional] 
**InnerSourceData** | Pointer to [**[]InnerSourceData**](InnerSourceData.md) |  | [optional] 
**ParentComponentPurls** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiDependencyDataDTO

`func NewApiDependencyDataDTO() *ApiDependencyDataDTO`

NewApiDependencyDataDTO instantiates a new ApiDependencyDataDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDependencyDataDTOWithDefaults

`func NewApiDependencyDataDTOWithDefaults() *ApiDependencyDataDTO`

NewApiDependencyDataDTOWithDefaults instantiates a new ApiDependencyDataDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirectDependency

`func (o *ApiDependencyDataDTO) GetDirectDependency() bool`

GetDirectDependency returns the DirectDependency field if non-nil, zero value otherwise.

### GetDirectDependencyOk

`func (o *ApiDependencyDataDTO) GetDirectDependencyOk() (*bool, bool)`

GetDirectDependencyOk returns a tuple with the DirectDependency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDependency

`func (o *ApiDependencyDataDTO) SetDirectDependency(v bool)`

SetDirectDependency sets DirectDependency field to given value.

### HasDirectDependency

`func (o *ApiDependencyDataDTO) HasDirectDependency() bool`

HasDirectDependency returns a boolean if a field has been set.

### GetInnerSource

`func (o *ApiDependencyDataDTO) GetInnerSource() bool`

GetInnerSource returns the InnerSource field if non-nil, zero value otherwise.

### GetInnerSourceOk

`func (o *ApiDependencyDataDTO) GetInnerSourceOk() (*bool, bool)`

GetInnerSourceOk returns a tuple with the InnerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerSource

`func (o *ApiDependencyDataDTO) SetInnerSource(v bool)`

SetInnerSource sets InnerSource field to given value.

### HasInnerSource

`func (o *ApiDependencyDataDTO) HasInnerSource() bool`

HasInnerSource returns a boolean if a field has been set.

### GetInnerSourceData

`func (o *ApiDependencyDataDTO) GetInnerSourceData() []InnerSourceData`

GetInnerSourceData returns the InnerSourceData field if non-nil, zero value otherwise.

### GetInnerSourceDataOk

`func (o *ApiDependencyDataDTO) GetInnerSourceDataOk() (*[]InnerSourceData, bool)`

GetInnerSourceDataOk returns a tuple with the InnerSourceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerSourceData

`func (o *ApiDependencyDataDTO) SetInnerSourceData(v []InnerSourceData)`

SetInnerSourceData sets InnerSourceData field to given value.

### HasInnerSourceData

`func (o *ApiDependencyDataDTO) HasInnerSourceData() bool`

HasInnerSourceData returns a boolean if a field has been set.

### GetParentComponentPurls

`func (o *ApiDependencyDataDTO) GetParentComponentPurls() []string`

GetParentComponentPurls returns the ParentComponentPurls field if non-nil, zero value otherwise.

### GetParentComponentPurlsOk

`func (o *ApiDependencyDataDTO) GetParentComponentPurlsOk() (*[]string, bool)`

GetParentComponentPurlsOk returns a tuple with the ParentComponentPurls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentComponentPurls

`func (o *ApiDependencyDataDTO) SetParentComponentPurls(v []string)`

SetParentComponentPurls sets ParentComponentPurls field to given value.

### HasParentComponentPurls

`func (o *ApiDependencyDataDTO) HasParentComponentPurls() bool`

HasParentComponentPurls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


