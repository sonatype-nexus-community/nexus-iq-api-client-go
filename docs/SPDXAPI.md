# \SPDXAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetByScanId**](SPDXAPI.md#GetByScanId) | **Get** /api/v2/spdx/{applicationId}/reports/{scanId} | 
[**GetLatestForStage**](SPDXAPI.md#GetLatestForStage) | **Get** /api/v2/spdx/{applicationId}/stages/{stageId} | 



## GetByScanId

> string GetByScanId(ctx, applicationId, scanId).Format(format).GenerateCycloneDx(generateCycloneDx).SpdxVersion(spdxVersion).Execute()





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
	applicationId := "applicationId_example" // string | Enter the applicationId for the application you want to generate the SBOM(s).
	scanId := "scanId_example" // string | Enter the scanId of the application scan.
	format := "format_example" // string | Enter the format for the SBOM(s) to be generated. (optional) (default to "json")
	generateCycloneDx := true // bool | Set to `true` to generate an equivalent CycloneDx SBOM. Both SBOMs will be combined as a tar.gz archive. (optional) (default to false)
	spdxVersion := "spdxVersion_example" // string | Enter the desired SPDX version, possible values are 2.2|2.3 (optional) (default to "2.3")

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SPDXAPI.GetByScanId(context.Background(), applicationId, scanId).Format(format).GenerateCycloneDx(generateCycloneDx).SpdxVersion(spdxVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SPDXAPI.GetByScanId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByScanId`: string
	fmt.Fprintf(os.Stdout, "Response from `SPDXAPI.GetByScanId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the applicationId for the application you want to generate the SBOM(s). | 
**scanId** | **string** | Enter the scanId of the application scan. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByScanIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** | Enter the format for the SBOM(s) to be generated. | [default to &quot;json&quot;]
 **generateCycloneDx** | **bool** | Set to &#x60;true&#x60; to generate an equivalent CycloneDx SBOM. Both SBOMs will be combined as a tar.gz archive. | [default to false]
 **spdxVersion** | **string** | Enter the desired SPDX version, possible values are 2.2|2.3 | [default to &quot;2.3&quot;]

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestForStage

> string GetLatestForStage(ctx, applicationId, stageId).Format(format).GenerateCycloneDx(generateCycloneDx).SpdxVersion(spdxVersion).Execute()





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
	applicationId := "applicationId_example" // string | Enter the applicationId for the application you want to generate the SBOM(s).
	stageId := "stageId_example" // string | Specify the stageId for the application evaluation. Allowed values are `develop`, `build`, `stage-release`, `release` and `operate`.
	format := "format_example" // string | Enter the format for the SBOM(s) to be generated. (optional) (default to "json")
	generateCycloneDx := true // bool | Set to `true` to generate an equivalent CycloneDx SBOM. Both SBOMs will be combined as a tar.gz archive. (optional) (default to false)
	spdxVersion := "spdxVersion_example" // string | Enter the desired SPDX version, possible values are 2.2|2.3 (optional) (default to "2.3")

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SPDXAPI.GetLatestForStage(context.Background(), applicationId, stageId).Format(format).GenerateCycloneDx(generateCycloneDx).SpdxVersion(spdxVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SPDXAPI.GetLatestForStage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestForStage`: string
	fmt.Fprintf(os.Stdout, "Response from `SPDXAPI.GetLatestForStage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the applicationId for the application you want to generate the SBOM(s). | 
**stageId** | **string** | Specify the stageId for the application evaluation. Allowed values are &#x60;develop&#x60;, &#x60;build&#x60;, &#x60;stage-release&#x60;, &#x60;release&#x60; and &#x60;operate&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestForStageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **format** | **string** | Enter the format for the SBOM(s) to be generated. | [default to &quot;json&quot;]
 **generateCycloneDx** | **bool** | Set to &#x60;true&#x60; to generate an equivalent CycloneDx SBOM. Both SBOMs will be combined as a tar.gz archive. | [default to false]
 **spdxVersion** | **string** | Enter the desired SPDX version, possible values are 2.2|2.3 | [default to &quot;2.3&quot;]

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/octet-stream, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

