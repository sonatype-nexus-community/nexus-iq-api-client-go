# \EvaluationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeprecatedManifestEvaluation**](EvaluationAPI.md#DeprecatedManifestEvaluation) | **Post** /api/v2/evaluation/applications/{applicationId}/manifestEvaluation | 
[**EvaluateComponents**](EvaluationAPI.md#EvaluateComponents) | **Post** /api/v2/evaluation/applications/{applicationId} | 
[**EvaluateSourceControl**](EvaluationAPI.md#EvaluateSourceControl) | **Post** /api/v2/evaluation/applications/{applicationId}/sourceControlEvaluation | 
[**GetApplicationEvaluationStatus**](EvaluationAPI.md#GetApplicationEvaluationStatus) | **Get** /api/v2/evaluation/applications/{applicationId}/status/{statusId} | 
[**GetComponentEvaluation**](EvaluationAPI.md#GetComponentEvaluation) | **Get** /api/v2/evaluation/applications/{applicationId}/results/{resultId} | 
[**PromoteScan**](EvaluationAPI.md#PromoteScan) | **Post** /api/v2/evaluation/applications/{applicationId}/promoteScan | 



## DeprecatedManifestEvaluation

> ApiApplicationEvaluationStatusDTOV2 DeprecatedManifestEvaluation(ctx, applicationId).ApiSourceControlEvaluationRequestDTO(apiSourceControlEvaluationRequestDTO).Execute()



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
    apiSourceControlEvaluationRequestDTO := *iqclient.NewApiSourceControlEvaluationRequestDTO() // ApiSourceControlEvaluationRequestDTO |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EvaluationAPI.DeprecatedManifestEvaluation(context.Background(), applicationId).ApiSourceControlEvaluationRequestDTO(apiSourceControlEvaluationRequestDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EvaluationAPI.DeprecatedManifestEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeprecatedManifestEvaluation`: ApiApplicationEvaluationStatusDTOV2
    fmt.Fprintf(os.Stdout, "Response from `EvaluationAPI.DeprecatedManifestEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeprecatedManifestEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiSourceControlEvaluationRequestDTO** | [**ApiSourceControlEvaluationRequestDTO**](ApiSourceControlEvaluationRequestDTO.md) |  | 

### Return type

[**ApiApplicationEvaluationStatusDTOV2**](ApiApplicationEvaluationStatusDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvaluateComponents

> ApiComponentEvaluationTicketDTOV2 EvaluateComponents(ctx, applicationId).ApiComponentEvaluationRequestDTOV2(apiComponentEvaluationRequestDTOV2).Execute()



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
    apiComponentEvaluationRequestDTOV2 := *iqclient.NewApiComponentEvaluationRequestDTOV2() // ApiComponentEvaluationRequestDTOV2 |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EvaluationAPI.EvaluateComponents(context.Background(), applicationId).ApiComponentEvaluationRequestDTOV2(apiComponentEvaluationRequestDTOV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EvaluationAPI.EvaluateComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EvaluateComponents`: ApiComponentEvaluationTicketDTOV2
    fmt.Fprintf(os.Stdout, "Response from `EvaluationAPI.EvaluateComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiComponentEvaluationRequestDTOV2** | [**ApiComponentEvaluationRequestDTOV2**](ApiComponentEvaluationRequestDTOV2.md) |  | 

### Return type

[**ApiComponentEvaluationTicketDTOV2**](ApiComponentEvaluationTicketDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvaluateSourceControl

> ApiApplicationEvaluationStatusDTOV2 EvaluateSourceControl(ctx, applicationId).ApiSourceControlEvaluationRequestDTO(apiSourceControlEvaluationRequestDTO).Execute()



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
    apiSourceControlEvaluationRequestDTO := *iqclient.NewApiSourceControlEvaluationRequestDTO() // ApiSourceControlEvaluationRequestDTO |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EvaluationAPI.EvaluateSourceControl(context.Background(), applicationId).ApiSourceControlEvaluationRequestDTO(apiSourceControlEvaluationRequestDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EvaluationAPI.EvaluateSourceControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EvaluateSourceControl`: ApiApplicationEvaluationStatusDTOV2
    fmt.Fprintf(os.Stdout, "Response from `EvaluationAPI.EvaluateSourceControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateSourceControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiSourceControlEvaluationRequestDTO** | [**ApiSourceControlEvaluationRequestDTO**](ApiSourceControlEvaluationRequestDTO.md) |  | 

### Return type

[**ApiApplicationEvaluationStatusDTOV2**](ApiApplicationEvaluationStatusDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEvaluationStatus

> ApiApplicationEvaluationResultDTOV2 GetApplicationEvaluationStatus(ctx, applicationId, statusId).Execute()



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
    statusId := "statusId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EvaluationAPI.GetApplicationEvaluationStatus(context.Background(), applicationId, statusId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EvaluationAPI.GetApplicationEvaluationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationEvaluationStatus`: ApiApplicationEvaluationResultDTOV2
    fmt.Fprintf(os.Stdout, "Response from `EvaluationAPI.GetApplicationEvaluationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**statusId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEvaluationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiApplicationEvaluationResultDTOV2**](ApiApplicationEvaluationResultDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentEvaluation

> ApiComponentEvaluationResultDTOV2 GetComponentEvaluation(ctx, applicationId, resultId).Execute()



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
    resultId := "resultId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EvaluationAPI.GetComponentEvaluation(context.Background(), applicationId, resultId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EvaluationAPI.GetComponentEvaluation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetComponentEvaluation`: ApiComponentEvaluationResultDTOV2
    fmt.Fprintf(os.Stdout, "Response from `EvaluationAPI.GetComponentEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**resultId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiComponentEvaluationResultDTOV2**](ApiComponentEvaluationResultDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromoteScan

> ApiApplicationEvaluationStatusDTOV2 PromoteScan(ctx, applicationId).ApiPromoteScanRequestDTOV2(apiPromoteScanRequestDTOV2).Execute()



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
    apiPromoteScanRequestDTOV2 := *iqclient.NewApiPromoteScanRequestDTOV2() // ApiPromoteScanRequestDTOV2 |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EvaluationAPI.PromoteScan(context.Background(), applicationId).ApiPromoteScanRequestDTOV2(apiPromoteScanRequestDTOV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EvaluationAPI.PromoteScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PromoteScan`: ApiApplicationEvaluationStatusDTOV2
    fmt.Fprintf(os.Stdout, "Response from `EvaluationAPI.PromoteScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromoteScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiPromoteScanRequestDTOV2** | [**ApiPromoteScanRequestDTOV2**](ApiPromoteScanRequestDTOV2.md) |  | 

### Return type

[**ApiApplicationEvaluationStatusDTOV2**](ApiApplicationEvaluationStatusDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

