# \SbomAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSbomVersion**](SbomAPI.md#DeleteSbomVersion) | **Delete** /api/v2/sbom/applications/{applicationId}/versions/{version} | Delete sbom version
[**GetActiveSbomVersionListByApplication**](SbomAPI.md#GetActiveSbomVersionListByApplication) | **Get** /api/v2/sbom/applications/{applicationId}/versions | Gets a list of active sbom versions by application id
[**GetImportStatus**](SbomAPI.md#GetImportStatus) | **Get** /api/v2/sbom/applications/{applicationId}/status/{importRequestId} | Get sbom import status
[**GetSbomComponents**](SbomAPI.md#GetSbomComponents) | **Get** /api/v2/sbom/applications/{applicationId}/versions/{version}/components | Gets the components found in a specific sbom version
[**GetSbomMetadataSummaryForApplication**](SbomAPI.md#GetSbomMetadataSummaryForApplication) | **Get** /api/v2/sbom/applications/{applicationId} | Gets a paginated list of SBOMs for an application
[**GetSbomVersion**](SbomAPI.md#GetSbomVersion) | **Get** /api/v2/sbom/applications/{applicationId}/versions/{version} | Gets a sbom version
[**ImportSbom**](SbomAPI.md#ImportSbom) | **Post** /api/v2/sbom/import | Import a new sbom version
[**SaveVulnerabilityAnalysis**](SbomAPI.md#SaveVulnerabilityAnalysis) | **Put** /api/v2/sbom/applications/{applicationId}/versions/{version}/vulnerability/{refId}/analysis | Updates a vulnerability analysis annotation for a specific SBOM vulnerability



## DeleteSbomVersion

> DeleteSbomVersion(ctx, applicationId, version).Execute()

Delete sbom version



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
	applicationId := "applicationId_example" // string | The internal id of the application
	version := "version_example" // string | URL Encoded version value of the sbom to be deleted

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SbomAPI.DeleteSbomVersion(context.Background(), applicationId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.DeleteSbomVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The internal id of the application | 
**version** | **string** | URL Encoded version value of the sbom to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSbomVersionRequest struct via the builder pattern


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


## GetActiveSbomVersionListByApplication

> GetActiveSbomVersionListByApplication(ctx, applicationId).Execute()

Gets a list of active sbom versions by application id



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
	applicationId := "applicationId_example" // string | The internal id of the application

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SbomAPI.GetActiveSbomVersionListByApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.GetActiveSbomVersionListByApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The internal id of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveSbomVersionListByApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImportStatus

> GetImportStatus(ctx, applicationId, importRequestId).Execute()

Get sbom import status



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
	applicationId := "applicationId_example" // string | The internal id of the application
	importRequestId := "importRequestId_example" // string | The id of the import request

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SbomAPI.GetImportStatus(context.Background(), applicationId, importRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.GetImportStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The internal id of the application | 
**importRequestId** | **string** | The id of the import request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSbomComponents

> GetSbomComponents(ctx, applicationId, version).VulnerabilityThreatLevels(vulnerabilityThreatLevels).DependencyTypes(dependencyTypes).SortBy(sortBy).Asc(asc).Page(page).PageSize(pageSize).Execute()

Gets the components found in a specific sbom version



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
	applicationId := "applicationId_example" // string | The internal id of the application
	version := "version_example" // string | URL Encoded version value of the sbom to query its components
	vulnerabilityThreatLevels := []string{"VulnerabilityThreatLevels_example"} // []string | If provided, filter components by the given threat level on their vulnerabilities (optional)
	dependencyTypes := []string{"DependencyTypes_example"} // []string | If provided, filter components by the given dependency types (optional)
	sortBy := "sortBy_example" // string | Criteria to sort the results. default = VULNERABILITIES (optional) (default to "VULNERABILITIES")
	asc := true // bool | Order mode ASC=true or DESC=false. default = false (optional) (default to false)
	page := int32(56) // int32 | Current page number. default = 1 (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items to return by page. default = 50 (optional) (default to 50)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SbomAPI.GetSbomComponents(context.Background(), applicationId, version).VulnerabilityThreatLevels(vulnerabilityThreatLevels).DependencyTypes(dependencyTypes).SortBy(sortBy).Asc(asc).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.GetSbomComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The internal id of the application | 
**version** | **string** | URL Encoded version value of the sbom to query its components | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **vulnerabilityThreatLevels** | **[]string** | If provided, filter components by the given threat level on their vulnerabilities | 
 **dependencyTypes** | **[]string** | If provided, filter components by the given dependency types | 
 **sortBy** | **string** | Criteria to sort the results. default &#x3D; VULNERABILITIES | [default to &quot;VULNERABILITIES&quot;]
 **asc** | **bool** | Order mode ASC&#x3D;true or DESC&#x3D;false. default &#x3D; false | [default to false]
 **page** | **int32** | Current page number. default &#x3D; 1 | [default to 1]
 **pageSize** | **int32** | Number of items to return by page. default &#x3D; 50 | [default to 50]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSbomMetadataSummaryForApplication

> GetSbomMetadataSummaryForApplication(ctx, applicationId).SortByDate(sortByDate).PageSize(pageSize).Page(page).Execute()

Gets a paginated list of SBOMs for an application



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
	applicationId := "applicationId_example" // string | The internal id of the application
	sortByDate := "sortByDate_example" // string | Sort results by import date. Allowed values [asc|desc]. default = asc (optional) (default to "asc")
	pageSize := int32(56) // int32 | Number of items to return by page. default = 10 (optional) (default to 10)
	page := int32(56) // int32 | Current page number. default = 1 (optional) (default to 1)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SbomAPI.GetSbomMetadataSummaryForApplication(context.Background(), applicationId).SortByDate(sortByDate).PageSize(pageSize).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.GetSbomMetadataSummaryForApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The internal id of the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomMetadataSummaryForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortByDate** | **string** | Sort results by import date. Allowed values [asc|desc]. default &#x3D; asc | [default to &quot;asc&quot;]
 **pageSize** | **int32** | Number of items to return by page. default &#x3D; 10 | [default to 10]
 **page** | **int32** | Current page number. default &#x3D; 1 | [default to 1]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSbomVersion

> GetSbomVersion(ctx, applicationId, version).State(state).Specification(specification).Accept(accept).Execute()

Gets a sbom version



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
	applicationId := "applicationId_example" // string | The internal id of the application
	version := "version_example" // string | URL Encoded version value of the sbom
	state := "state_example" // string | The state of the sbom version. Allowed values [original|current]. default = current (optional) (default to "current")
	specification := "specification_example" // string | Target specification of the sbom. Allowed values [cyclonedx1.5|spdx2.3]. default = cyclonedx1.5 (optional) (default to "cyclonedx1.5")
	accept := "accept_example" // string | Output format(json/xml) of the sbom. Changing the output format only applicable when downloading the current form of the SBOM. The original sbom will always return in the original form that it was ingested. When requesting `current` form and if this header value is not present the sbom will be returned in its original ingested format. Allowed values {'application/json'|'application/xml'}. default = null (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SbomAPI.GetSbomVersion(context.Background(), applicationId, version).State(state).Specification(specification).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.GetSbomVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The internal id of the application | 
**version** | **string** | URL Encoded version value of the sbom | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **state** | **string** | The state of the sbom version. Allowed values [original|current]. default &#x3D; current | [default to &quot;current&quot;]
 **specification** | **string** | Target specification of the sbom. Allowed values [cyclonedx1.5|spdx2.3]. default &#x3D; cyclonedx1.5 | [default to &quot;cyclonedx1.5&quot;]
 **accept** | **string** | Output format(json/xml) of the sbom. Changing the output format only applicable when downloading the current form of the SBOM. The original sbom will always return in the original form that it was ingested. When requesting &#x60;current&#x60; form and if this header value is not present the sbom will be returned in its original ingested format. Allowed values {&#39;application/json&#39;|&#39;application/xml&#39;}. default &#x3D; null | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json|application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImportSbom

> ImportSbom(ctx).ImportSbomRequest(importSbomRequest).Execute()

Import a new sbom version



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
	importSbomRequest := *sonatypeiq.NewImportSbomRequest("ApplicationId_example") // ImportSbomRequest |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SbomAPI.ImportSbom(context.Background()).ImportSbomRequest(importSbomRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.ImportSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **importSbomRequest** | [**ImportSbomRequest**](ImportSbomRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveVulnerabilityAnalysis

> SaveVulnerabilityAnalysis(ctx, applicationId, version, refId).SBOMVulnerabilityAnalysisRequest(sBOMVulnerabilityAnalysisRequest).Execute()

Updates a vulnerability analysis annotation for a specific SBOM vulnerability



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
	applicationId := "applicationId_example" // string | The internal id of the application
	version := "version_example" // string | The version for a specific SBOM where the vulnerability is present
	refId := "refId_example" // string | The vulnerability id of a vulnerability
	sBOMVulnerabilityAnalysisRequest := *sonatypeiq.NewSBOMVulnerabilityAnalysisRequest(*sonatypeiq.NewComponentLocator(), *sonatypeiq.NewVulnerabilityAnalysis("Detail_example", "Justification_example", "Response_example", "State_example")) // SBOMVulnerabilityAnalysisRequest | Vulnerability analysis details with component information

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SbomAPI.SaveVulnerabilityAnalysis(context.Background(), applicationId, version, refId).SBOMVulnerabilityAnalysisRequest(sBOMVulnerabilityAnalysisRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.SaveVulnerabilityAnalysis``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The internal id of the application | 
**version** | **string** | The version for a specific SBOM where the vulnerability is present | 
**refId** | **string** | The vulnerability id of a vulnerability | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveVulnerabilityAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **sBOMVulnerabilityAnalysisRequest** | [**SBOMVulnerabilityAnalysisRequest**](SBOMVulnerabilityAnalysisRequest.md) | Vulnerability analysis details with component information | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

