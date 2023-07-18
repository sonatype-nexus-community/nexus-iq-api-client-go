# \DataRetentionPoliciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataRetentionPolicies**](DataRetentionPoliciesAPI.md#GetDataRetentionPolicies) | **Get** /api/v2/dataRetentionPolicies/organizations/{organizationId} | 
[**SetDataRetentionPolicies**](DataRetentionPoliciesAPI.md#SetDataRetentionPolicies) | **Put** /api/v2/dataRetentionPolicies/organizations/{organizationId} | 



## GetDataRetentionPolicies

> ApiDataRetentionPoliciesDTO GetDataRetentionPolicies(ctx, organizationId).Execute()



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
    organizationId := "organizationId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataRetentionPoliciesAPI.GetDataRetentionPolicies(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataRetentionPoliciesAPI.GetDataRetentionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataRetentionPolicies`: ApiDataRetentionPoliciesDTO
    fmt.Fprintf(os.Stdout, "Response from `DataRetentionPoliciesAPI.GetDataRetentionPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataRetentionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiDataRetentionPoliciesDTO**](ApiDataRetentionPoliciesDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDataRetentionPolicies

> SetDataRetentionPolicies(ctx, organizationId).ApiDataRetentionPoliciesDTO(apiDataRetentionPoliciesDTO).Execute()



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
    organizationId := "organizationId_example" // string | 
    apiDataRetentionPoliciesDTO := *iqclient.NewApiDataRetentionPoliciesDTO() // ApiDataRetentionPoliciesDTO |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.DataRetentionPoliciesAPI.SetDataRetentionPolicies(context.Background(), organizationId).ApiDataRetentionPoliciesDTO(apiDataRetentionPoliciesDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataRetentionPoliciesAPI.SetDataRetentionPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDataRetentionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiDataRetentionPoliciesDTO** | [**ApiDataRetentionPoliciesDTO**](ApiDataRetentionPoliciesDTO.md) |  | 

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

