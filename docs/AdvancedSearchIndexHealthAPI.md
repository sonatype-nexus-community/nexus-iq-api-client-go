# \AdvancedSearchIndexHealthAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Analyze**](AdvancedSearchIndexHealthAPI.md#Analyze) | **Get** /api/v2/search/advanced/index/analyze | 
[**CancelActiveJob**](AdvancedSearchIndexHealthAPI.md#CancelActiveJob) | **Post** /api/v2/search/advanced/index/jobs/cancel | 
[**CompleteUnlock**](AdvancedSearchIndexHealthAPI.md#CompleteUnlock) | **Post** /api/v2/search/advanced/index/unlock/complete | 
[**GetActiveJob**](AdvancedSearchIndexHealthAPI.md#GetActiveJob) | **Get** /api/v2/search/advanced/index/jobs/active | 
[**GetGenerations**](AdvancedSearchIndexHealthAPI.md#GetGenerations) | **Get** /api/v2/search/advanced/index/generations | 
[**GetJob**](AdvancedSearchIndexHealthAPI.md#GetJob) | **Get** /api/v2/search/advanced/index/jobs/{jobId} | 
[**GetJobEvents**](AdvancedSearchIndexHealthAPI.md#GetJobEvents) | **Get** /api/v2/search/advanced/index/jobs/{jobId}/events | 
[**GetJobHistory**](AdvancedSearchIndexHealthAPI.md#GetJobHistory) | **Get** /api/v2/search/advanced/index/jobs/history | 
[**GetSetup**](AdvancedSearchIndexHealthAPI.md#GetSetup) | **Get** /api/v2/search/advanced/index/setup | 
[**MarkUnlockReadyToTest**](AdvancedSearchIndexHealthAPI.md#MarkUnlockReadyToTest) | **Post** /api/v2/search/advanced/index/unlock/ready-to-test | 
[**RunAnalyze**](AdvancedSearchIndexHealthAPI.md#RunAnalyze) | **Post** /api/v2/search/advanced/index/analyze/run | 
[**StartJob**](AdvancedSearchIndexHealthAPI.md#StartJob) | **Post** /api/v2/search/advanced/index/jobs | 
[**StartUnlockIndexing**](AdvancedSearchIndexHealthAPI.md#StartUnlockIndexing) | **Post** /api/v2/search/advanced/index/unlock/start | 



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


## CompleteUnlock

> ApiSearchIndexUnlockStateDTO CompleteUnlock(ctx).Execute()





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
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.CompleteUnlock(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.CompleteUnlock``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CompleteUnlock`: ApiSearchIndexUnlockStateDTO
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.CompleteUnlock`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteUnlockRequest struct via the builder pattern


### Return type

[**ApiSearchIndexUnlockStateDTO**](ApiSearchIndexUnlockStateDTO.md)

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


## GetGenerations

> ApiSearchIndexGenerationsDTO GetGenerations(ctx).Execute()





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
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.GetGenerations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.GetGenerations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGenerations`: ApiSearchIndexGenerationsDTO
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.GetGenerations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenerationsRequest struct via the builder pattern


### Return type

[**ApiSearchIndexGenerationsDTO**](ApiSearchIndexGenerationsDTO.md)

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


## GetJobHistory

> []ApiSearchIndexHistoryEntryDTO GetJobHistory(ctx).Limit(limit).Execute()





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
	limit := int32(56) // int32 |  (optional) (default to 50)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.GetJobHistory(context.Background()).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.GetJobHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetJobHistory`: []ApiSearchIndexHistoryEntryDTO
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.GetJobHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJobHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 50]

### Return type

[**[]ApiSearchIndexHistoryEntryDTO**](ApiSearchIndexHistoryEntryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSetup

> ApiSearchIndexSetupDTO GetSetup(ctx).Execute()





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
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.GetSetup(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.GetSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSetup`: ApiSearchIndexSetupDTO
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.GetSetup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSetupRequest struct via the builder pattern


### Return type

[**ApiSearchIndexSetupDTO**](ApiSearchIndexSetupDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkUnlockReadyToTest

> ApiSearchIndexUnlockStateDTO MarkUnlockReadyToTest(ctx).Execute()





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
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.MarkUnlockReadyToTest(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.MarkUnlockReadyToTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkUnlockReadyToTest`: ApiSearchIndexUnlockStateDTO
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.MarkUnlockReadyToTest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarkUnlockReadyToTestRequest struct via the builder pattern


### Return type

[**ApiSearchIndexUnlockStateDTO**](ApiSearchIndexUnlockStateDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunAnalyze

> ApiSearchIndexAnalyzeDTO RunAnalyze(ctx).Execute()





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
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.RunAnalyze(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.RunAnalyze``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunAnalyze`: ApiSearchIndexAnalyzeDTO
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.RunAnalyze`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRunAnalyzeRequest struct via the builder pattern


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


## StartUnlockIndexing

> SearchIndexJob StartUnlockIndexing(ctx).Execute()





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
	resp, r, err := apiClient.AdvancedSearchIndexHealthAPI.StartUnlockIndexing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchIndexHealthAPI.StartUnlockIndexing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartUnlockIndexing`: SearchIndexJob
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchIndexHealthAPI.StartUnlockIndexing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStartUnlockIndexingRequest struct via the builder pattern


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

