# \ReportsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAll1**](ReportsAPI.md#GetAll1) | **Get** /api/v2/reports/applications | 
[**GetByApplicationId**](ReportsAPI.md#GetByApplicationId) | **Get** /api/v2/reports/applications/{applicationId} | 
[**GetComponentsInQuarantine**](ReportsAPI.md#GetComponentsInQuarantine) | **Get** /api/v2/reports/components/quarantined | 
[**GetComponentsWithWaivers**](ReportsAPI.md#GetComponentsWithWaivers) | **Get** /api/v2/reports/components/waivers | 
[**GetMetrics**](ReportsAPI.md#GetMetrics) | **Post** /api/v2/reports/metrics | 
[**GetReportHistoryForApplication**](ReportsAPI.md#GetReportHistoryForApplication) | **Get** /api/v2/reports/applications/{applicationId}/history | 
[**GetStaleWaivers**](ReportsAPI.md#GetStaleWaivers) | **Get** /api/v2/reports/waivers/stale | 



## GetAll1

> []ApiApplicationReportDTOV2 GetAll1(ctx).Execute()



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

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.GetAll1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetAll1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAll1`: []ApiApplicationReportDTOV2
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetAll1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAll1Request struct via the builder pattern


### Return type

[**[]ApiApplicationReportDTOV2**](ApiApplicationReportDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByApplicationId

> []ApiApplicationReportDTOV2 GetByApplicationId(ctx, applicationId).Execute()



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
    applicationId := "applicationId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.GetByApplicationId(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetByApplicationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetByApplicationId`: []ApiApplicationReportDTOV2
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetByApplicationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApiApplicationReportDTOV2**](ApiApplicationReportDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentsInQuarantine

> ApiComponentsInQuarantineDTO GetComponentsInQuarantine(ctx).Execute()



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

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.GetComponentsInQuarantine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetComponentsInQuarantine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentsInQuarantine`: ApiComponentsInQuarantineDTO
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetComponentsInQuarantine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentsInQuarantineRequest struct via the builder pattern


### Return type

[**ApiComponentsInQuarantineDTO**](ApiComponentsInQuarantineDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentsWithWaivers

> ApiComponentWaiversDTO GetComponentsWithWaivers(ctx).Format(format).Execute()



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
    format := "format_example" // string |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.GetComponentsWithWaivers(context.Background()).Format(format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetComponentsWithWaivers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentsWithWaivers`: ApiComponentWaiversDTO
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetComponentsWithWaivers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentsWithWaiversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | 

### Return type

[**ApiComponentWaiversDTO**](ApiComponentWaiversDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetrics

> GetMetrics(ctx).ApiMetricsReportingQueryDTOV2(apiMetricsReportingQueryDTOV2).Execute()



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
    apiMetricsReportingQueryDTOV2 := *iqclient.NewApiMetricsReportingQueryDTOV2() // ApiMetricsReportingQueryDTOV2 |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.ReportsAPI.GetMetrics(context.Background()).ApiMetricsReportingQueryDTOV2(apiMetricsReportingQueryDTOV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiMetricsReportingQueryDTOV2** | [**ApiMetricsReportingQueryDTOV2**](ApiMetricsReportingQueryDTOV2.md) |  | 

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


## GetReportHistoryForApplication

> ApiReportHistoryDTO GetReportHistoryForApplication(ctx, applicationId).Stage(stage).Limit(limit).Execute()



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
    applicationId := "applicationId_example" // string | 
    stage := "stage_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.GetReportHistoryForApplication(context.Background(), applicationId).Stage(stage).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportHistoryForApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportHistoryForApplication`: ApiReportHistoryDTO
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportHistoryForApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportHistoryForApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stage** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

[**ApiReportHistoryDTO**](ApiReportHistoryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStaleWaivers

> ApiStaleWaiversResponseDTO GetStaleWaivers(ctx).Execute()



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

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.GetStaleWaivers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetStaleWaivers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStaleWaivers`: ApiStaleWaiversResponseDTO
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetStaleWaivers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStaleWaiversRequest struct via the builder pattern


### Return type

[**ApiStaleWaiversResponseDTO**](ApiStaleWaiversResponseDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

