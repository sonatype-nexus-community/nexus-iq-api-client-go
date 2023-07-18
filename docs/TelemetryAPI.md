# \TelemetryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostExternalTelemetry**](TelemetryAPI.md#PostExternalTelemetry) | **Post** /api/v2/telemetry | 



## PostExternalTelemetry

> PostExternalTelemetry(ctx).UserAgent(userAgent).RequestBody(requestBody).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    iqclient "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {
    userAgent := "userAgent_example" // string |  (optional)
    requestBody := map[string]string{"key": "Inner_example"} // map[string]string |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.TelemetryAPI.PostExternalTelemetry(context.Background()).UserAgent(userAgent).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TelemetryAPI.PostExternalTelemetry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostExternalTelemetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAgent** | **string** |  | 
 **requestBody** | **map[string]string** |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

