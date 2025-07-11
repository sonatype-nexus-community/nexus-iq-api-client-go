# \SBOMAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSbomVersion**](SBOMAPI.md#DeleteSbomVersion) | **Delete** /api/v2/sbom/applications/{applicationId}/versions/{version} | Delete sbom version
[**DeleteVulnerabilityAnalysis**](SBOMAPI.md#DeleteVulnerabilityAnalysis) | **Delete** /api/v2/sbom/applications/{applicationId}/versions/{version}/vulnerability/{refId}/analysis | Deletes a Vulnerability analysis for a given component.
[**GetActiveSbomVersionListByApplication**](SBOMAPI.md#GetActiveSbomVersionListByApplication) | **Get** /api/v2/sbom/applications/{applicationId}/versions | Gets a list of active sbom versions by application id
[**GetImportStatus**](SBOMAPI.md#GetImportStatus) | **Get** /api/v2/sbom/applications/{applicationId}/status/{importRequestId} | Get sbom import status
[**GetSbomComponents**](SBOMAPI.md#GetSbomComponents) | **Get** /api/v2/sbom/applications/{applicationId}/versions/{version}/components | Gets the components found in a specific sbom version
[**GetSbomMetadataSummaryForApplication**](SBOMAPI.md#GetSbomMetadataSummaryForApplication) | **Get** /api/v2/sbom/applications/{applicationId} | Gets a paginated list of SBOMs for an application
[**GetSbomVersion**](SBOMAPI.md#GetSbomVersion) | **Get** /api/v2/sbom/applications/{applicationId}/versions/{version} | Gets a sbom version
[**ImportSbom**](SBOMAPI.md#ImportSbom) | **Post** /api/v2/sbom/import | Import a new sbom version
[**SaveVulnerabilityAnalysis**](SBOMAPI.md#SaveVulnerabilityAnalysis) | **Put** /api/v2/sbom/applications/{applicationId}/versions/{version}/vulnerability/{refId}/analysis | Updates a vulnerability analysis annotation for a specific SBOM vulnerability



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
	r, err := apiClient.SBOMAPI.DeleteSbomVersion(context.Background(), applicationId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.DeleteSbomVersion``: %v\n", err)
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


## DeleteVulnerabilityAnalysis

> DeleteVulnerabilityAnalysis(ctx, applicationId, version, refId).ComponentLocator(componentLocator).Execute()

Deletes a Vulnerability analysis for a given component.



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
	componentLocator := *sonatypeiq.NewComponentLocator() // ComponentLocator | Hash or packageUrl to identify the component

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SBOMAPI.DeleteVulnerabilityAnalysis(context.Background(), applicationId, version, refId).ComponentLocator(componentLocator).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.DeleteVulnerabilityAnalysis``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteVulnerabilityAnalysisRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentLocator** | [**ComponentLocator**](ComponentLocator.md) | Hash or packageUrl to identify the component | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
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
	r, err := apiClient.SBOMAPI.GetActiveSbomVersionListByApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.GetActiveSbomVersionListByApplication``: %v\n", err)
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
	r, err := apiClient.SBOMAPI.GetImportStatus(context.Background(), applicationId, importRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.GetImportStatus``: %v\n", err)
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

> GetSbomComponents(ctx, applicationId, version).VulnerabilityThreatLevels(vulnerabilityThreatLevels).DependencyTypes(dependencyTypes).Filter(filter).SortBy(sortBy).Asc(asc).Page(page).PageSize(pageSize).Execute()

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
	filter := "filter_example" // string | If provided, filter components by the given search criteria (optional)
	sortBy := "sortBy_example" // string | Criteria to sort the results. default = VULNERABILITIES (optional) (default to "VULNERABILITIES")
	asc := true // bool | Order mode ASC=true or DESC=false. default = false (optional) (default to false)
	page := int32(56) // int32 | Current page number. default = 1 (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items to return by page. default = 50 (optional) (default to 50)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SBOMAPI.GetSbomComponents(context.Background(), applicationId, version).VulnerabilityThreatLevels(vulnerabilityThreatLevels).DependencyTypes(dependencyTypes).Filter(filter).SortBy(sortBy).Asc(asc).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.GetSbomComponents``: %v\n", err)
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
 **filter** | **string** | If provided, filter components by the given search criteria | 
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

> GetSbomMetadataSummaryForApplication(ctx, applicationId).SortByDate(sortByDate).PageSize(pageSize).Page(page).SortBy(sortBy).Asc(asc).Execute()

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
	sortByDate := "sortByDate_example" // string | Deprecated, use sortBy and asc instead. Sort results by import date. Allowed values [asc|desc]. default = asc (optional) (default to "asc")
	pageSize := int32(56) // int32 | Number of items to return by page. default = 10 (optional) (default to 10)
	page := int32(56) // int32 | Current page number. default = 1 (optional) (default to 1)
	sortBy := "sortBy_example" // string | Criteria to sort the results. default = IMPORT_DATE, when used sortByDate is ignored (optional) (default to "IMPORT_DATE")
	asc := true // bool | Order mode ASC=true or DESC=false. default = true (optional) (default to true)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SBOMAPI.GetSbomMetadataSummaryForApplication(context.Background(), applicationId).SortByDate(sortByDate).PageSize(pageSize).Page(page).SortBy(sortBy).Asc(asc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.GetSbomMetadataSummaryForApplication``: %v\n", err)
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

 **sortByDate** | **string** | Deprecated, use sortBy and asc instead. Sort results by import date. Allowed values [asc|desc]. default &#x3D; asc | [default to &quot;asc&quot;]
 **pageSize** | **int32** | Number of items to return by page. default &#x3D; 10 | [default to 10]
 **page** | **int32** | Current page number. default &#x3D; 1 | [default to 1]
 **sortBy** | **string** | Criteria to sort the results. default &#x3D; IMPORT_DATE, when used sortByDate is ignored | [default to &quot;IMPORT_DATE&quot;]
 **asc** | **bool** | Order mode ASC&#x3D;true or DESC&#x3D;false. default &#x3D; true | [default to true]

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
	specification := "specification_example" // string | Target specification of the sbom. Allowed values [cyclonedx1.6|cyclonedx1.5|spdx2.2|spdx2.3]. default = cyclonedx1.6 (optional) (default to "cyclonedx1.6")
	accept := "accept_example" // string | Output format(json/xml) of the sbom. Changing the output format only applicable when downloading the current form of the SBOM. The original sbom will always return in the original form that it was ingested. When requesting `current` form and if this header value is not present the sbom will be returned in its original ingested format. Allowed values {'application/json'|'application/xml'}. default = null (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SBOMAPI.GetSbomVersion(context.Background(), applicationId, version).State(state).Specification(specification).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.GetSbomVersion``: %v\n", err)
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
 **specification** | **string** | Target specification of the sbom. Allowed values [cyclonedx1.6|cyclonedx1.5|spdx2.2|spdx2.3]. default &#x3D; cyclonedx1.6 | [default to &quot;cyclonedx1.6&quot;]
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

> ImportSbom(ctx).EnableBinaryImport(enableBinaryImport).IgnoreValidationError(ignoreValidationError).ImportSbomRequest(importSbomRequest).Execute()

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
	enableBinaryImport := true // bool | Enable importing as a binary file. (optional) (default to false)
	ignoreValidationError := true // bool | Skip the SBOM validation if an error occurs. (optional) (default to false)
	importSbomRequest := *sonatypeiq.NewImportSbomRequest("ApplicationId_example") // ImportSbomRequest |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SBOMAPI.ImportSbom(context.Background()).EnableBinaryImport(enableBinaryImport).IgnoreValidationError(ignoreValidationError).ImportSbomRequest(importSbomRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.ImportSbom``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiImportSbomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **enableBinaryImport** | **bool** | Enable importing as a binary file. | [default to false]
 **ignoreValidationError** | **bool** | Skip the SBOM validation if an error occurs. | [default to false]
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

> SaveVulnerabilityAnalysis(ctx, applicationId, version, refId).SbomVulnerabilityAnalysisRequest(sbomVulnerabilityAnalysisRequest).Execute()

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
	sbomVulnerabilityAnalysisRequest := *sonatypeiq.NewSbomVulnerabilityAnalysisRequest(*sonatypeiq.NewComponentLocator(), *sonatypeiq.NewVulnerabilityAnalysis("Detail_example", "Justification_example", "Response_example", "State_example")) // SbomVulnerabilityAnalysisRequest | Vulnerability analysis details with component information

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SBOMAPI.SaveVulnerabilityAnalysis(context.Background(), applicationId, version, refId).SbomVulnerabilityAnalysisRequest(sbomVulnerabilityAnalysisRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SBOMAPI.SaveVulnerabilityAnalysis``: %v\n", err)
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



 **sbomVulnerabilityAnalysisRequest** | [**SbomVulnerabilityAnalysisRequest**](SbomVulnerabilityAnalysisRequest.md) | Vulnerability analysis details with component information | 

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

