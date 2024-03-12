# \SourceControlMetricsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSourceControl**](SourceControlMetricsAPI.md#GetSourceControl) | **Get** /api/v2/sourceControlMetrics/{ownerType}/{internalOwnerId} | 



## GetSourceControl

> ApiPullRequestResults GetSourceControl(ctx, ownerType, internalOwnerId).Execute()



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
	ownerType := "ownerType_example" // string | 
	internalOwnerId := "internalOwnerId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlMetricsAPI.GetSourceControl(context.Background(), ownerType, internalOwnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlMetricsAPI.GetSourceControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceControl`: ApiPullRequestResults
	fmt.Fprintf(os.Stdout, "Response from `SourceControlMetricsAPI.GetSourceControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiPullRequestResults**](ApiPullRequestResults.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

