# \LabelsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLabel**](LabelsAPI.md#AddLabel) | **Post** /api/v2/labels/{ownerType}/{ownerId} | 
[**DeleteLabel**](LabelsAPI.md#DeleteLabel) | **Delete** /api/v2/labels/{ownerType}/{ownerId}/{labelId} | 
[**GetApplicableContexts**](LabelsAPI.md#GetApplicableContexts) | **Get** /api/v2/labels/{ownerType}/{ownerId}/applicable/context/{labelId} | 
[**GetApplicableLabels**](LabelsAPI.md#GetApplicableLabels) | **Get** /api/v2/labels/{ownerType}/{ownerId}/applicable | 
[**GetLabels**](LabelsAPI.md#GetLabels) | **Get** /api/v2/labels/{ownerType}/{ownerId} | 
[**UpdateLabel**](LabelsAPI.md#UpdateLabel) | **Put** /api/v2/labels/{ownerType}/{ownerId} | 



## AddLabel

> ApiLabelDTO AddLabel(ctx, ownerType, ownerId).ApiLabelDTO(apiLabelDTO).Execute()



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
    ownerType := "ownerType_example" // string | 
    ownerId := "ownerId_example" // string | 
    apiLabelDTO := *iqclient.NewApiLabelDTO() // ApiLabelDTO |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsAPI.AddLabel(context.Background(), ownerType, ownerId).ApiLabelDTO(apiLabelDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.AddLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLabel`: ApiLabelDTO
    fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.AddLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiLabelDTO** | [**ApiLabelDTO**](ApiLabelDTO.md) |  | 

### Return type

[**ApiLabelDTO**](ApiLabelDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLabel

> DeleteLabel(ctx, ownerType, ownerId, labelId).Execute()



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
    ownerType := "ownerType_example" // string | 
    ownerId := "ownerId_example" // string | 
    labelId := "labelId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.LabelsAPI.DeleteLabel(context.Background(), ownerType, ownerId, labelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.DeleteLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 
**labelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicableContexts

> ApplicableContext GetApplicableContexts(ctx, ownerType, ownerId, labelId).Execute()



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
    ownerType := "ownerType_example" // string | 
    ownerId := "ownerId_example" // string | 
    labelId := "labelId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsAPI.GetApplicableContexts(context.Background(), ownerType, ownerId, labelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.GetApplicableContexts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicableContexts`: ApplicableContext
    fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.GetApplicableContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 
**labelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApplicableContext**](ApplicableContext.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicableLabels

> ApplicableLabels GetApplicableLabels(ctx, ownerType, ownerId).Execute()



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
    ownerType := "ownerType_example" // string | 
    ownerId := "ownerId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsAPI.GetApplicableLabels(context.Background(), ownerType, ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.GetApplicableLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicableLabels`: ApplicableLabels
    fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.GetApplicableLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicableLabels**](ApplicableLabels.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabels

> []ApiLabelDTO GetLabels(ctx, ownerType, ownerId).Inherit(inherit).Execute()



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
    ownerType := "ownerType_example" // string | 
    ownerId := "ownerId_example" // string | 
    inherit := true // bool |  (optional) (default to false)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsAPI.GetLabels(context.Background(), ownerType, ownerId).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.GetLabels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLabels`: []ApiLabelDTO
    fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.GetLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inherit** | **bool** |  | [default to false]

### Return type

[**[]ApiLabelDTO**](ApiLabelDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLabel

> ApiLabelDTO UpdateLabel(ctx, ownerType, ownerId).ApiLabelDTO(apiLabelDTO).Execute()



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
    ownerType := "ownerType_example" // string | 
    ownerId := "ownerId_example" // string | 
    apiLabelDTO := *iqclient.NewApiLabelDTO() // ApiLabelDTO |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LabelsAPI.UpdateLabel(context.Background(), ownerType, ownerId).ApiLabelDTO(apiLabelDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.UpdateLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLabel`: ApiLabelDTO
    fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.UpdateLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiLabelDTO** | [**ApiLabelDTO**](ApiLabelDTO.md) |  | 

### Return type

[**ApiLabelDTO**](ApiLabelDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

