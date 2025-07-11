# \ThirdPartyAnalysisAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetScanStatus**](ThirdPartyAnalysisAPI.md#GetScanStatus) | **Get** /api/v2/scan/applications/{applicationId}/status/{scanRequestId} | 
[**ScanComponents**](ThirdPartyAnalysisAPI.md#ScanComponents) | **Post** /api/v2/scan/applications/{applicationId}/sources/{source} | 



## GetScanStatus

> ApiThirdPartyScanResultDTO GetScanStatus(ctx, applicationId, scanRequestId).Execute()





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
	applicationId := "applicationId_example" // string | Enter the application internal id for the SBOM to be evaluated.
	scanRequestId := "scanRequestId_example" // string | Enter the statusId from the statusUrl generated as a response to the POST request to perform the evaluation.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ThirdPartyAnalysisAPI.GetScanStatus(context.Background(), applicationId, scanRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyAnalysisAPI.GetScanStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScanStatus`: ApiThirdPartyScanResultDTO
	fmt.Fprintf(os.Stdout, "Response from `ThirdPartyAnalysisAPI.GetScanStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the application internal id for the SBOM to be evaluated. | 
**scanRequestId** | **string** | Enter the statusId from the statusUrl generated as a response to the POST request to perform the evaluation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScanStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiThirdPartyScanResultDTO**](ApiThirdPartyScanResultDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScanComponents

> ApiThirdPartyScanTicketDTO ScanComponents(ctx, applicationId, source).StageId(stageId).Body(body).Execute()





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
	applicationId := "applicationId_example" // string | Enter the application internal id. Use the Applications REST API to retrieve theapplication internal id.
	source := "source_example" // string | Specify the specification name of the SBOM to be evaluated. Allowed values are `cyclonedx` and `spdx`
	stageId := "stageId_example" // string | Enter the stageId to run the evaluation for. The policy actions will be determined by the stage selected. Allowed values are `develop`, `build`, `stage-release`, `release` and `operate` (optional) (default to "build")
	body := "body_example" // string | Select the request header content-type from the dropdown, depending on whether the SBOM is in XML or JSON format.  Embed the contents of the SBOM here or enter the file path for the SBOM. A component in the SBOM can be identified by: <ol><li>packageUrl</li><li>Component hash</li><li>Component name and version</li></ol>  Any SPE and SWID tags for the component will be preserved in the evaluation report. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ThirdPartyAnalysisAPI.ScanComponents(context.Background(), applicationId, source).StageId(stageId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThirdPartyAnalysisAPI.ScanComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScanComponents`: ApiThirdPartyScanTicketDTO
	fmt.Fprintf(os.Stdout, "Response from `ThirdPartyAnalysisAPI.ScanComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the application internal id. Use the Applications REST API to retrieve theapplication internal id. | 
**source** | **string** | Specify the specification name of the SBOM to be evaluated. Allowed values are &#x60;cyclonedx&#x60; and &#x60;spdx&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiScanComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stageId** | **string** | Enter the stageId to run the evaluation for. The policy actions will be determined by the stage selected. Allowed values are &#x60;develop&#x60;, &#x60;build&#x60;, &#x60;stage-release&#x60;, &#x60;release&#x60; and &#x60;operate&#x60; | [default to &quot;build&quot;]
 **body** | **string** | Select the request header content-type from the dropdown, depending on whether the SBOM is in XML or JSON format.  Embed the contents of the SBOM here or enter the file path for the SBOM. A component in the SBOM can be identified by: &lt;ol&gt;&lt;li&gt;packageUrl&lt;/li&gt;&lt;li&gt;Component hash&lt;/li&gt;&lt;li&gt;Component name and version&lt;/li&gt;&lt;/ol&gt;  Any SPE and SWID tags for the component will be preserved in the evaluation report. | 

### Return type

[**ApiThirdPartyScanTicketDTO**](ApiThirdPartyScanTicketDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

