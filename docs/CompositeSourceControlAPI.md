# \CompositeSourceControlAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCompositeSourceControlByOwner**](CompositeSourceControlAPI.md#GetCompositeSourceControlByOwner) | **Get** /api/v2/compositeSourceControl/{ownerType}/{internalOwnerId} | 



## GetCompositeSourceControlByOwner

> ApiCompositeSourceControlDTO GetCompositeSourceControlByOwner(ctx, ownerType, internalOwnerId).Execute()



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
    internalOwnerId := "internalOwnerId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompositeSourceControlAPI.GetCompositeSourceControlByOwner(context.Background(), ownerType, internalOwnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompositeSourceControlAPI.GetCompositeSourceControlByOwner``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompositeSourceControlByOwner`: ApiCompositeSourceControlDTO
    fmt.Fprintf(os.Stdout, "Response from `CompositeSourceControlAPI.GetCompositeSourceControlByOwner`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompositeSourceControlByOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiCompositeSourceControlDTO**](ApiCompositeSourceControlDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

