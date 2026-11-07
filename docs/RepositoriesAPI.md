# \RepositoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComponent**](RepositoriesAPI.md#GetComponent) | **Get** /api/v2/repositories/{repositoryManagerId}/{repositoryId}/components/{componentId} | 
[**GetComponents**](RepositoriesAPI.md#GetComponents) | **Get** /api/v2/repositories/{repositoryManagerId}/{repositoryId}/components | 
[**GetQuarantinedByPath**](RepositoriesAPI.md#GetQuarantinedByPath) | **Post** /api/v2/repositories/{repositoryManagerInstanceId}/{repositoryPublicId}/components/quarantined/pathnames | 
[**GetQueueStats**](RepositoriesAPI.md#GetQueueStats) | **Get** /api/v2/repositories/{repositoryManagerId}/{repositoryId}/queue/stats | 
[**GetViolations**](RepositoriesAPI.md#GetViolations) | **Get** /api/v2/repositories/{repositoryManagerId}/{repositoryId}/components/{componentId}/violations | 
[**ReleaseQuarantineWithoutReEval**](RepositoriesAPI.md#ReleaseQuarantineWithoutReEval) | **Post** /api/v2/repositories/quarantine/{quarantineId}/release | 
[**UploadScan**](RepositoriesAPI.md#UploadScan) | **Post** /api/v2/repositories/{repositoryManagerId}/{repositoryId}/components | 



## GetComponent

> ApiHostedRepositoryComponentDTO GetComponent(ctx, repositoryManagerId, repositoryId, componentId).Execute()



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
	repositoryManagerId := "repositoryManagerId_example" // string | 
	repositoryId := "repositoryId_example" // string | 
	componentId := "componentId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.GetComponent(context.Background(), repositoryManagerId, repositoryId, componentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.GetComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponent`: ApiHostedRepositoryComponentDTO
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.GetComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerId** | **string** |  | 
**repositoryId** | **string** |  | 
**componentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiHostedRepositoryComponentDTO**](ApiHostedRepositoryComponentDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponents

> ApiHostedRepositoryComponentListDTO GetComponents(ctx, repositoryManagerId, repositoryId).Page(page).PageSize(pageSize).Filter(filter).Execute()



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
	repositoryManagerId := "repositoryManagerId_example" // string | 
	repositoryId := "repositoryId_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 25)
	filter := "filter_example" // string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.GetComponents(context.Background(), repositoryManagerId, repositoryId).Page(page).PageSize(pageSize).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.GetComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponents`: ApiHostedRepositoryComponentListDTO
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.GetComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerId** | **string** |  | 
**repositoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 25]
 **filter** | **string** |  | 

### Return type

[**ApiHostedRepositoryComponentListDTO**](ApiHostedRepositoryComponentListDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuarantinedByPath

> ApiRepositoryPathResponseDTO GetQuarantinedByPath(ctx, repositoryManagerInstanceId, repositoryPublicId).RequestBody(requestBody).Execute()





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
	repositoryManagerInstanceId := "repositoryManagerInstanceId_example" // string | Enter the repository manager instance ID.
	repositoryPublicId := "repositoryPublicId_example" // string | Enter the repository public ID.
	requestBody := []string{"Property_example"} // []string | Specify the pathnames.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.GetQuarantinedByPath(context.Background(), repositoryManagerInstanceId, repositoryPublicId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.GetQuarantinedByPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuarantinedByPath`: ApiRepositoryPathResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.GetQuarantinedByPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerInstanceId** | **string** | Enter the repository manager instance ID. | 
**repositoryPublicId** | **string** | Enter the repository public ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuarantinedByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]string** | Specify the pathnames. | 

### Return type

[**ApiRepositoryPathResponseDTO**](ApiRepositoryPathResponseDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQueueStats

> ApiQueueStatsDTO GetQueueStats(ctx, repositoryManagerId, repositoryId).Execute()



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
	repositoryManagerId := "repositoryManagerId_example" // string | 
	repositoryId := "repositoryId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.GetQueueStats(context.Background(), repositoryManagerId, repositoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.GetQueueStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQueueStats`: ApiQueueStatsDTO
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.GetQueueStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerId** | **string** |  | 
**repositoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiQueueStatsDTO**](ApiQueueStatsDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViolations

> ApiComponentViolationListDTO GetViolations(ctx, repositoryManagerId, repositoryId, componentId).Execute()



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
	repositoryManagerId := "repositoryManagerId_example" // string | 
	repositoryId := "repositoryId_example" // string | 
	componentId := "componentId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.GetViolations(context.Background(), repositoryManagerId, repositoryId, componentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.GetViolations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetViolations`: ApiComponentViolationListDTO
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.GetViolations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerId** | **string** |  | 
**repositoryId** | **string** |  | 
**componentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiComponentViolationListDTO**](ApiComponentViolationListDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseQuarantineWithoutReEval

> ApiComponentReleasedFromQuarantineDTO ReleaseQuarantineWithoutReEval(ctx, quarantineId).Body(body).Execute()





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
	quarantineId := "quarantineId_example" // string | Enter the component `quarantineId`.
	body := "body_example" // string | Enter a waiver comment for releasing the component from quarantine.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.ReleaseQuarantineWithoutReEval(context.Background(), quarantineId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.ReleaseQuarantineWithoutReEval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReleaseQuarantineWithoutReEval`: ApiComponentReleasedFromQuarantineDTO
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.ReleaseQuarantineWithoutReEval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quarantineId** | **string** | Enter the component &#x60;quarantineId&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseQuarantineWithoutReEvalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Enter a waiver comment for releasing the component from quarantine. | 

### Return type

[**ApiComponentReleasedFromQuarantineDTO**](ApiComponentReleasedFromQuarantineDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadScan

> UploadScan(ctx, repositoryManagerId, repositoryId).ComponentId(componentId).PolicyEvaluationStage(policyEvaluationStage).ScanFile(scanFile).Execute()



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
	repositoryManagerId := "repositoryManagerId_example" // string | 
	repositoryId := "repositoryId_example" // string | 
	componentId := "componentId_example" // string |  (optional)
	policyEvaluationStage := "policyEvaluationStage_example" // string |  (optional)
	scanFile := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.RepositoriesAPI.UploadScan(context.Background(), repositoryManagerId, repositoryId).ComponentId(componentId).PolicyEvaluationStage(policyEvaluationStage).ScanFile(scanFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.UploadScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerId** | **string** |  | 
**repositoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **componentId** | **string** |  | 
 **policyEvaluationStage** | **string** |  | 
 **scanFile** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

