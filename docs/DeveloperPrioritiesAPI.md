# \DeveloperPrioritiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPriorities**](DeveloperPrioritiesAPI.md#GetPriorities) | **Get** /api/v2/developer/priorities/{applicationId}/{scanId} | 
[**GetPrioritiesExport**](DeveloperPrioritiesAPI.md#GetPrioritiesExport) | **Get** /api/v2/developer/priorities/{applicationId}/{scanId}/export | 



## GetPriorities

> DevelopmentPrioritizationResults GetPriorities(ctx, applicationId, scanId).IncludeRemediation(includeRemediation).Page(page).PageSize(pageSize).OptionalComponentNameFilter(optionalComponentNameFilter).Execute()





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
	applicationId := "applicationId_example" // string | 
	scanId := "scanId_example" // string | 
	includeRemediation := true // bool |  (optional) (default to false)
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	optionalComponentNameFilter := "optionalComponentNameFilter_example" // string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.DeveloperPrioritiesAPI.GetPriorities(context.Background(), applicationId, scanId).IncludeRemediation(includeRemediation).Page(page).PageSize(pageSize).OptionalComponentNameFilter(optionalComponentNameFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeveloperPrioritiesAPI.GetPriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriorities`: DevelopmentPrioritizationResults
	fmt.Fprintf(os.Stdout, "Response from `DeveloperPrioritiesAPI.GetPriorities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**scanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeRemediation** | **bool** |  | [default to false]
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **optionalComponentNameFilter** | **string** |  | 

### Return type

[**DevelopmentPrioritizationResults**](DevelopmentPrioritizationResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrioritiesExport

> GetPrioritiesExport(ctx, applicationId, scanId).Execute()





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
	applicationId := "applicationId_example" // string | 
	scanId := "scanId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.DeveloperPrioritiesAPI.GetPrioritiesExport(context.Background(), applicationId, scanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeveloperPrioritiesAPI.GetPrioritiesExport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**scanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrioritiesExportRequest struct via the builder pattern


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

