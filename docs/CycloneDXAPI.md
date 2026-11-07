# \CycloneDXAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetByReportId**](CycloneDXAPI.md#GetByReportId) | **Get** /api/v2/cycloneDx/{cdxVersion}/{applicationId}/reports/{reportId} | 
[**GetByReportIdWithVersion**](CycloneDXAPI.md#GetByReportIdWithVersion) | **Get** /api/v2/cycloneDx/{cdxVersion}/hostedRepositoryComponent/{hrcId}/reports/{reportId} | 
[**GetLatest**](CycloneDXAPI.md#GetLatest) | **Get** /api/v2/cycloneDx/{cdxVersion}/{applicationId}/stages/{stageId} | 
[**GetLatestWithVersion**](CycloneDXAPI.md#GetLatestWithVersion) | **Get** /api/v2/cycloneDx/{cdxVersion}/hostedRepositoryComponent/{hrcId}/stages/{stageId} | 



## GetByReportId

> map[string]interface{} GetByReportId(ctx, applicationId, reportId, cdxVersion).Execute()





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
	applicationId := "applicationId_example" // string | Enter the internal applicationId for the application you want to generate the SBOM. You can also retrieve the applicationId using the Application REST API.
	reportId := "reportId_example" // string | Enter the reportId to generate the SBOM for the application for a specific scan report.
	cdxVersion := "cdxVersion_example" // string | Possible values are 1.1|1.2|1.3|1.4|1.5|1.6|1.7.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.CycloneDXAPI.GetByReportId(context.Background(), applicationId, reportId, cdxVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CycloneDXAPI.GetByReportId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByReportId`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CycloneDXAPI.GetByReportId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the internal applicationId for the application you want to generate the SBOM. You can also retrieve the applicationId using the Application REST API. | 
**reportId** | **string** | Enter the reportId to generate the SBOM for the application for a specific scan report. | 
**cdxVersion** | **string** | Possible values are 1.1|1.2|1.3|1.4|1.5|1.6|1.7. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByReportIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByReportIdWithVersion

> map[string]interface{} GetByReportIdWithVersion(ctx, hrcId, reportId, cdxVersion).Execute()





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
	hrcId := "hrcId_example" // string | Enter the hosted repository component ID (UUID) you want to generate the SBOM for.
	reportId := "reportId_example" // string | Enter the reportId to generate the SBOM for the hosted repository component for a specific scan report.
	cdxVersion := "cdxVersion_example" // string | Possible values are 1.1|1.2|1.3|1.4|1.5|1.6|1.7.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.CycloneDXAPI.GetByReportIdWithVersion(context.Background(), hrcId, reportId, cdxVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CycloneDXAPI.GetByReportIdWithVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetByReportIdWithVersion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CycloneDXAPI.GetByReportIdWithVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hrcId** | **string** | Enter the hosted repository component ID (UUID) you want to generate the SBOM for. | 
**reportId** | **string** | Enter the reportId to generate the SBOM for the hosted repository component for a specific scan report. | 
**cdxVersion** | **string** | Possible values are 1.1|1.2|1.3|1.4|1.5|1.6|1.7. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByReportIdWithVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatest

> map[string]interface{} GetLatest(ctx, applicationId, stageId, cdxVersion).Execute()





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
	applicationId := "applicationId_example" // string | Enter the internal applicationId for the application you want to generate the SBOM. You can also retrieve the applicationId using the Application REST API.
	stageId := "stageId_example" // string | Enter the stageId to generate the SBOM based on the latest application policy evaluation at that stage. Allowed values for stageId are 'develop', 'source', 'build', 'stage-release', 'release', and, 'operate'.
	cdxVersion := "cdxVersion_example" // string | Possible values are 1.1|1.2|1.3|1.4|1.5|1.6|1.7.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.CycloneDXAPI.GetLatest(context.Background(), applicationId, stageId, cdxVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CycloneDXAPI.GetLatest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatest`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CycloneDXAPI.GetLatest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the internal applicationId for the application you want to generate the SBOM. You can also retrieve the applicationId using the Application REST API. | 
**stageId** | **string** | Enter the stageId to generate the SBOM based on the latest application policy evaluation at that stage. Allowed values for stageId are &#39;develop&#39;, &#39;source&#39;, &#39;build&#39;, &#39;stage-release&#39;, &#39;release&#39;, and, &#39;operate&#39;. | 
**cdxVersion** | **string** | Possible values are 1.1|1.2|1.3|1.4|1.5|1.6|1.7. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestWithVersion

> map[string]interface{} GetLatestWithVersion(ctx, hrcId, stageId, cdxVersion).Execute()





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
	hrcId := "hrcId_example" // string | Enter the hosted repository component ID (UUID) you want to generate the SBOM for.
	stageId := "stageId_example" // string | Enter the stageId to generate the SBOM based on the latest hosted repository component policy evaluation at that stage.
	cdxVersion := "cdxVersion_example" // string | Possible values are 1.1|1.2|1.3|1.4|1.5|1.6|1.7.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.CycloneDXAPI.GetLatestWithVersion(context.Background(), hrcId, stageId, cdxVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CycloneDXAPI.GetLatestWithVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLatestWithVersion`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CycloneDXAPI.GetLatestWithVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hrcId** | **string** | Enter the hosted repository component ID (UUID) you want to generate the SBOM for. | 
**stageId** | **string** | Enter the stageId to generate the SBOM based on the latest hosted repository component policy evaluation at that stage. | 
**cdxVersion** | **string** | Possible values are 1.1|1.2|1.3|1.4|1.5|1.6|1.7. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestWithVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

