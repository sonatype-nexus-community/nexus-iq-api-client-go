# AttributionReportTemplateDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentTitle** | Pointer to **string** |  | [optional] 
**Footer** | Pointer to **string** |  | [optional] 
**Header** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IncludeAppendix** | Pointer to **bool** |  | [optional] 
**IncludeInnerSource** | Pointer to **bool** |  | [optional] 
**IncludeSonatypeSpecialLicenses** | Pointer to **bool** |  | [optional] 
**IncludeStandardLicenseTexts** | Pointer to **bool** |  | [optional] 
**IncludeTableOfContents** | Pointer to **bool** |  | [optional] 
**LastUpdatedAt** | Pointer to **time.Time** |  | [optional] 
**TemplateName** | Pointer to **string** |  | [optional] 

## Methods

### NewAttributionReportTemplateDTO

`func NewAttributionReportTemplateDTO() *AttributionReportTemplateDTO`

NewAttributionReportTemplateDTO instantiates a new AttributionReportTemplateDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttributionReportTemplateDTOWithDefaults

`func NewAttributionReportTemplateDTOWithDefaults() *AttributionReportTemplateDTO`

NewAttributionReportTemplateDTOWithDefaults instantiates a new AttributionReportTemplateDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentTitle

`func (o *AttributionReportTemplateDTO) GetDocumentTitle() string`

GetDocumentTitle returns the DocumentTitle field if non-nil, zero value otherwise.

### GetDocumentTitleOk

`func (o *AttributionReportTemplateDTO) GetDocumentTitleOk() (*string, bool)`

GetDocumentTitleOk returns a tuple with the DocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTitle

`func (o *AttributionReportTemplateDTO) SetDocumentTitle(v string)`

SetDocumentTitle sets DocumentTitle field to given value.

### HasDocumentTitle

`func (o *AttributionReportTemplateDTO) HasDocumentTitle() bool`

HasDocumentTitle returns a boolean if a field has been set.

### GetFooter

`func (o *AttributionReportTemplateDTO) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *AttributionReportTemplateDTO) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *AttributionReportTemplateDTO) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *AttributionReportTemplateDTO) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetHeader

`func (o *AttributionReportTemplateDTO) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *AttributionReportTemplateDTO) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *AttributionReportTemplateDTO) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *AttributionReportTemplateDTO) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetId

`func (o *AttributionReportTemplateDTO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AttributionReportTemplateDTO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AttributionReportTemplateDTO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AttributionReportTemplateDTO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncludeAppendix

`func (o *AttributionReportTemplateDTO) GetIncludeAppendix() bool`

GetIncludeAppendix returns the IncludeAppendix field if non-nil, zero value otherwise.

### GetIncludeAppendixOk

`func (o *AttributionReportTemplateDTO) GetIncludeAppendixOk() (*bool, bool)`

GetIncludeAppendixOk returns a tuple with the IncludeAppendix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAppendix

`func (o *AttributionReportTemplateDTO) SetIncludeAppendix(v bool)`

SetIncludeAppendix sets IncludeAppendix field to given value.

### HasIncludeAppendix

`func (o *AttributionReportTemplateDTO) HasIncludeAppendix() bool`

HasIncludeAppendix returns a boolean if a field has been set.

### GetIncludeInnerSource

`func (o *AttributionReportTemplateDTO) GetIncludeInnerSource() bool`

GetIncludeInnerSource returns the IncludeInnerSource field if non-nil, zero value otherwise.

### GetIncludeInnerSourceOk

`func (o *AttributionReportTemplateDTO) GetIncludeInnerSourceOk() (*bool, bool)`

GetIncludeInnerSourceOk returns a tuple with the IncludeInnerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInnerSource

`func (o *AttributionReportTemplateDTO) SetIncludeInnerSource(v bool)`

SetIncludeInnerSource sets IncludeInnerSource field to given value.

### HasIncludeInnerSource

`func (o *AttributionReportTemplateDTO) HasIncludeInnerSource() bool`

HasIncludeInnerSource returns a boolean if a field has been set.

### GetIncludeSonatypeSpecialLicenses

`func (o *AttributionReportTemplateDTO) GetIncludeSonatypeSpecialLicenses() bool`

GetIncludeSonatypeSpecialLicenses returns the IncludeSonatypeSpecialLicenses field if non-nil, zero value otherwise.

### GetIncludeSonatypeSpecialLicensesOk

`func (o *AttributionReportTemplateDTO) GetIncludeSonatypeSpecialLicensesOk() (*bool, bool)`

GetIncludeSonatypeSpecialLicensesOk returns a tuple with the IncludeSonatypeSpecialLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSonatypeSpecialLicenses

`func (o *AttributionReportTemplateDTO) SetIncludeSonatypeSpecialLicenses(v bool)`

SetIncludeSonatypeSpecialLicenses sets IncludeSonatypeSpecialLicenses field to given value.

### HasIncludeSonatypeSpecialLicenses

`func (o *AttributionReportTemplateDTO) HasIncludeSonatypeSpecialLicenses() bool`

HasIncludeSonatypeSpecialLicenses returns a boolean if a field has been set.

### GetIncludeStandardLicenseTexts

`func (o *AttributionReportTemplateDTO) GetIncludeStandardLicenseTexts() bool`

GetIncludeStandardLicenseTexts returns the IncludeStandardLicenseTexts field if non-nil, zero value otherwise.

### GetIncludeStandardLicenseTextsOk

`func (o *AttributionReportTemplateDTO) GetIncludeStandardLicenseTextsOk() (*bool, bool)`

GetIncludeStandardLicenseTextsOk returns a tuple with the IncludeStandardLicenseTexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeStandardLicenseTexts

`func (o *AttributionReportTemplateDTO) SetIncludeStandardLicenseTexts(v bool)`

SetIncludeStandardLicenseTexts sets IncludeStandardLicenseTexts field to given value.

### HasIncludeStandardLicenseTexts

`func (o *AttributionReportTemplateDTO) HasIncludeStandardLicenseTexts() bool`

HasIncludeStandardLicenseTexts returns a boolean if a field has been set.

### GetIncludeTableOfContents

`func (o *AttributionReportTemplateDTO) GetIncludeTableOfContents() bool`

GetIncludeTableOfContents returns the IncludeTableOfContents field if non-nil, zero value otherwise.

### GetIncludeTableOfContentsOk

`func (o *AttributionReportTemplateDTO) GetIncludeTableOfContentsOk() (*bool, bool)`

GetIncludeTableOfContentsOk returns a tuple with the IncludeTableOfContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTableOfContents

`func (o *AttributionReportTemplateDTO) SetIncludeTableOfContents(v bool)`

SetIncludeTableOfContents sets IncludeTableOfContents field to given value.

### HasIncludeTableOfContents

`func (o *AttributionReportTemplateDTO) HasIncludeTableOfContents() bool`

HasIncludeTableOfContents returns a boolean if a field has been set.

### GetLastUpdatedAt

`func (o *AttributionReportTemplateDTO) GetLastUpdatedAt() time.Time`

GetLastUpdatedAt returns the LastUpdatedAt field if non-nil, zero value otherwise.

### GetLastUpdatedAtOk

`func (o *AttributionReportTemplateDTO) GetLastUpdatedAtOk() (*time.Time, bool)`

GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedAt

`func (o *AttributionReportTemplateDTO) SetLastUpdatedAt(v time.Time)`

SetLastUpdatedAt sets LastUpdatedAt field to given value.

### HasLastUpdatedAt

`func (o *AttributionReportTemplateDTO) HasLastUpdatedAt() bool`

HasLastUpdatedAt returns a boolean if a field has been set.

### GetTemplateName

`func (o *AttributionReportTemplateDTO) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *AttributionReportTemplateDTO) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *AttributionReportTemplateDTO) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *AttributionReportTemplateDTO) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


