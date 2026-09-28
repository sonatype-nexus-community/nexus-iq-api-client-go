# \AdvancedSearchIndexHealthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Analyze**](AdvancedSearchIndexHealthAPI.md#Analyze) | **Get** /api/v2/search/advanced/index/analyze | 
[**CancelActiveJob**](AdvancedSearchIndexHealthAPI.md#CancelActiveJob) | **Post** /api/v2/search/advanced/index/jobs/cancel | 
[**GetActiveJob**](AdvancedSearchIndexHealthAPI.md#GetActiveJob) | **Get** /api/v2/search/advanced/index/jobs/active | 
[**GetJob**](AdvancedSearchIndexHealthAPI.md#GetJob) | **Get** /api/v2/search/advanced/index/jobs/{jobId} | 
[**GetJobEvents**](AdvancedSearchIndexHealthAPI.md#GetJobEvents) | **Get** /api/v2/search/advanced/index/jobs/{jobId}/events | 
[**StartJob**](AdvancedSearchIndexHealthAPI.md#StartJob) | **Post** /api/v2/search/advanced/index/jobs | 



## Analyze

> ApiSearchIndexAnalyzeDTO Analyze(ctx).Execute()





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
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.Analyze(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.Analyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Analyze`: ApiSearchIndexAnalyzeDTO
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.Analyze`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAnalyzeRequest struct via the builder pattern


### Return type

[**ApiSearchIndexAnalyzeDTO**](ApiSearchIndexAnalyzeDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelActiveJob

> SearchIndexJob CancelActiveJob(ctx).Execute()





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
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.CancelActiveJob(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.CancelActiveJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelActiveJob`: SearchIndexJob
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.CancelActiveJob`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCancelActiveJobRequest struct via the builder pattern


### Return type

[**SearchIndexJob**](SearchIndexJob.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveJob

> SearchIndexJob GetActiveJob(ctx).Execute()





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
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.GetActiveJob(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.GetActiveJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActiveJob`: SearchIndexJob
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.GetActiveJob`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveJobRequest struct via the builder pattern


### Return type

[**SearchIndexJob**](SearchIndexJob.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJob

> SearchIndexJob GetJob(ctx, jobId).Execute()





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
	jobId := "jobId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJob`: SearchIndexJob
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SearchIndexJob**](SearchIndexJob.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJobEvents

> []SearchIndexJobEvent GetJobEvents(ctx, jobId).Limit(limit).Execute()





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
	jobId := "jobId_example" // string | 
	limit := int32(56) // int32 |  (optional) (default to 100)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.GetJobEvents(context.Background(), jobId).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.GetJobEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobEvents`: []SearchIndexJobEvent
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.GetJobEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJobEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 100]

### Return type

[**[]SearchIndexJobEvent**](SearchIndexJobEvent.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartJob

> SearchIndexJob StartJob(ctx).ApiSearchIndexJobRequestDTO(apiSearchIndexJobRequestDTO).Execute()





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
	apiSearchIndexJobRequestDTO := *sonatypeiq.NewApiSearchIndexJobRequestDTO() // ApiSearchIndexJobRequestDTO |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.StartJob(context.Background()).ApiSearchIndexJobRequestDTO(apiSearchIndexJobRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.StartJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartJob`: SearchIndexJob
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.StartJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStartJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiSearchIndexJobRequestDTO** | [**ApiSearchIndexJobRequestDTO**](ApiSearchIndexJobRequestDTO.md) |  | 

### Return type

[**SearchIndexJob**](SearchIndexJob.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

