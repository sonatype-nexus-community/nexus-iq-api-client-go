# \LicenseLegalMetadataTemplatesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAttributionReportTemplate**](LicenseLegalMetadataTemplatesAPI.md#DeleteAttributionReportTemplate) | **Delete** /api/v2/licenseLegalMetadata/report-template/{id} | 
[**GetAllAttributionReportTemplates**](LicenseLegalMetadataTemplatesAPI.md#GetAllAttributionReportTemplates) | **Get** /api/v2/licenseLegalMetadata/report-template | 
[**GetAttributionReportTemplateById**](LicenseLegalMetadataTemplatesAPI.md#GetAttributionReportTemplateById) | **Get** /api/v2/licenseLegalMetadata/report-template/{id} | 
[**SaveAttributionReportTemplate**](LicenseLegalMetadataTemplatesAPI.md#SaveAttributionReportTemplate) | **Post** /api/v2/licenseLegalMetadata/report-template | 



## DeleteAttributionReportTemplate

> DeleteAttributionReportTemplate(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatypeiq "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {
	id := "id_example" // string | Enter the template id for the template to be deleted.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.LicenseLegalMetadataTemplatesAPI.DeleteAttributionReportTemplate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataTemplatesAPI.DeleteAttributionReportTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Enter the template id for the template to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttributionReportTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAttributionReportTemplates

> []AttributionReportTemplateDTO GetAllAttributionReportTemplates(ctx).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatypeiq "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseLegalMetadataTemplatesAPI.GetAllAttributionReportTemplates(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataTemplatesAPI.GetAllAttributionReportTemplates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAttributionReportTemplates`: []AttributionReportTemplateDTO
	fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataTemplatesAPI.GetAllAttributionReportTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAttributionReportTemplatesRequest struct via the builder pattern


### Return type

[**[]AttributionReportTemplateDTO**](AttributionReportTemplateDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributionReportTemplateById

> AttributionReportTemplateDTO GetAttributionReportTemplateById(ctx, id).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatypeiq "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {
	id := "id_example" // string | Enter the templateId for the report.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseLegalMetadataTemplatesAPI.GetAttributionReportTemplateById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataTemplatesAPI.GetAttributionReportTemplateById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttributionReportTemplateById`: AttributionReportTemplateDTO
	fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataTemplatesAPI.GetAttributionReportTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Enter the templateId for the report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributionReportTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttributionReportTemplateDTO**](AttributionReportTemplateDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAttributionReportTemplate

> AttributionReportTemplateDTO SaveAttributionReportTemplate(ctx).AttributionReportTemplateDTO(attributionReportTemplateDTO).Execute()





### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	sonatypeiq "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {
	attributionReportTemplateDTO := *sonatypeiq.NewAttributionReportTemplateDTO() // AttributionReportTemplateDTO | Specify the details for the template as:<ul><li>`id` is the template id.</li><li>`templateName` indicates the name of the stored template.</li><li>`documentTitle` is the title that is displayed at the top of the report.</li><li>`header` is the text that will be displayed above the `documentTitle`.</li><li>`footer` is the text that will be displayed at the bottom of the report.<li><li>`includeTableOfContents` is `true` if a table of contents containing links to the components and their licenses will be added to the report.<li>`includeAppendix` is `true` if standard license text will be grouped in the report appendix.</li><li>`includeStandardLicenseTexts` is `true` if the standard license text will be displayed for components with no license files.</li><li>`includeSonatypeSpecialLicenses` is `true` if Sonatype Special Licenses (e.g. Generic-Copyleft-Clause, Generic-Liberal-Clause, See-License-Clause, Identity-Clause etc.) will be included in the report.</li><li>`includeInnerSource` is `true` if InnerSource components will be included in the report.</li></ul> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseLegalMetadataTemplatesAPI.SaveAttributionReportTemplate(context.Background()).AttributionReportTemplateDTO(attributionReportTemplateDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataTemplatesAPI.SaveAttributionReportTemplate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveAttributionReportTemplate`: AttributionReportTemplateDTO
	fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataTemplatesAPI.SaveAttributionReportTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveAttributionReportTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributionReportTemplateDTO** | [**AttributionReportTemplateDTO**](AttributionReportTemplateDTO.md) | Specify the details for the template as:&lt;ul&gt;&lt;li&gt;&#x60;id&#x60; is the template id.&lt;/li&gt;&lt;li&gt;&#x60;templateName&#x60; indicates the name of the stored template.&lt;/li&gt;&lt;li&gt;&#x60;documentTitle&#x60; is the title that is displayed at the top of the report.&lt;/li&gt;&lt;li&gt;&#x60;header&#x60; is the text that will be displayed above the &#x60;documentTitle&#x60;.&lt;/li&gt;&lt;li&gt;&#x60;footer&#x60; is the text that will be displayed at the bottom of the report.&lt;li&gt;&lt;li&gt;&#x60;includeTableOfContents&#x60; is &#x60;true&#x60; if a table of contents containing links to the components and their licenses will be added to the report.&lt;li&gt;&#x60;includeAppendix&#x60; is &#x60;true&#x60; if standard license text will be grouped in the report appendix.&lt;/li&gt;&lt;li&gt;&#x60;includeStandardLicenseTexts&#x60; is &#x60;true&#x60; if the standard license text will be displayed for components with no license files.&lt;/li&gt;&lt;li&gt;&#x60;includeSonatypeSpecialLicenses&#x60; is &#x60;true&#x60; if Sonatype Special Licenses (e.g. Generic-Copyleft-Clause, Generic-Liberal-Clause, See-License-Clause, Identity-Clause etc.) will be included in the report.&lt;/li&gt;&lt;li&gt;&#x60;includeInnerSource&#x60; is &#x60;true&#x60; if InnerSource components will be included in the report.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**AttributionReportTemplateDTO**](AttributionReportTemplateDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

