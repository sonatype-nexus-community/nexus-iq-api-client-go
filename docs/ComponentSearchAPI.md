# \ComponentSearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportComponentSearchReport**](ComponentSearchAPI.md#ExportComponentSearchReport) | **Get** /api/v2/componentSearch/downloadComponentSearchReport | 
[**GetCveAffectedComponents**](ComponentSearchAPI.md#GetCveAffectedComponents) | **Get** /api/v2/componentSearch/cveAffectedComponents | 
[**SearchComponent**](ComponentSearchAPI.md#SearchComponent) | **Get** /api/v2/search/component | 



## ExportComponentSearchReport

> ExportComponentSearchReport(ctx).CveId(cveId).Execute()





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
	cveId := []string{"Inner_example"} // []string | CVE identifier(s). Can be specified multiple times for multiple CVEs. Defaults to CVE-2025-55182 if not specified. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ComponentSearchAPI.ExportComponentSearchReport(context.Background()).CveId(cveId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentSearchAPI.ExportComponentSearchReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportComponentSearchReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cveId** | **[]string** | CVE identifier(s). Can be specified multiple times for multiple CVEs. Defaults to CVE-2025-55182 if not specified. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCveAffectedComponents

> GetCveAffectedComponents(ctx).CveId(cveId).PageNumber(pageNumber).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()





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
	cveId := []string{"Inner_example"} // []string | CVE identifier(s). Can be specified multiple times for multiple CVEs.
	pageNumber := int32(56) // int32 | Page number (1-indexed, minimum: 1, default: 1) (optional) (default to 1)
	pageSize := int32(56) // int32 | Number of items per page (1-1000, default: 10) (optional) (default to 10)
	sortBy := "sortBy_example" // string | Sort field: applicationName, applicationId, componentName, evaluationDate, stage, activeWaiver, violating, cveId. When not specified, sorts by applicationName (asc), then componentName (asc), then cveId (asc) (optional)
	sortOrder := "sortOrder_example" // string | Sort order: asc or desc, default: asc (optional) (default to "asc")

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ComponentSearchAPI.GetCveAffectedComponents(context.Background()).CveId(cveId).PageNumber(pageNumber).PageSize(pageSize).SortBy(sortBy).SortOrder(sortOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentSearchAPI.GetCveAffectedComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCveAffectedComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cveId** | **[]string** | CVE identifier(s). Can be specified multiple times for multiple CVEs. | 
 **pageNumber** | **int32** | Page number (1-indexed, minimum: 1, default: 1) | [default to 1]
 **pageSize** | **int32** | Number of items per page (1-1000, default: 10) | [default to 10]
 **sortBy** | **string** | Sort field: applicationName, applicationId, componentName, evaluationDate, stage, activeWaiver, violating, cveId. When not specified, sorts by applicationName (asc), then componentName (asc), then cveId (asc) | 
 **sortOrder** | **string** | Sort order: asc or desc, default: asc | [default to &quot;asc&quot;]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchComponent

> ApiSearchResultsDTOV2 SearchComponent(ctx).StageId(stageId).Hash(hash).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Execute()





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
	stageId := "stageId_example" // string | Specify the evaluation report stage.
	hash := "hash_example" // string | Enter the component hash. (optional)
	componentIdentifier := *sonatypeiq.NewComponentIdentifier() // ComponentIdentifier | Specify the componentIdentifier object containing the format and coordinates. (optional)
	packageUrl := "packageUrl_example" // string | Enter the packageUrl. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentSearchAPI.SearchComponent(context.Background()).StageId(stageId).Hash(hash).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentSearchAPI.SearchComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchComponent`: ApiSearchResultsDTOV2
	fmt.Fprintf(os.Stdout, "Response from `ComponentSearchAPI.SearchComponent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stageId** | **string** | Specify the evaluation report stage. | 
 **hash** | **string** | Enter the component hash. | 
 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) | Specify the componentIdentifier object containing the format and coordinates. | 
 **packageUrl** | **string** | Enter the packageUrl. | 

### Return type

[**ApiSearchResultsDTOV2**](ApiSearchResultsDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

