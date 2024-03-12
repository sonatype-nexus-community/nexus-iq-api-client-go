# \EndpointsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOpenAPI**](EndpointsAPI.md#GetOpenAPI) | **Get** /api/v2/endpoints/{apiType} | 



## GetOpenAPI

> string GetOpenAPI(ctx, apiType).Execute()



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
	apiType := "apiType_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.EndpointsAPI.GetOpenAPI(context.Background(), apiType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EndpointsAPI.GetOpenAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenAPI`: string
	fmt.Fprintf(os.Stdout, "Response from `EndpointsAPI.GetOpenAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

