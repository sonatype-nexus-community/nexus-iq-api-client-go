# \OrganizationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrganization**](OrganizationsAPI.md#AddOrganization) | **Post** /api/v2/organizations | 
[**GetOrganization**](OrganizationsAPI.md#GetOrganization) | **Get** /api/v2/organizations/{organizationId} | 
[**GetOrganizations**](OrganizationsAPI.md#GetOrganizations) | **Get** /api/v2/organizations | 
[**MoveOrganization**](OrganizationsAPI.md#MoveOrganization) | **Put** /api/v2/organizations/{organizationId}/move/destination/{destinationId} | 



## AddOrganization

> ApiOrganizationDTO AddOrganization(ctx).ApiOrganizationDTO(apiOrganizationDTO).Execute()



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
    apiOrganizationDTO := *iqclient.NewApiOrganizationDTO() // ApiOrganizationDTO |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsAPI.AddOrganization(context.Background()).ApiOrganizationDTO(apiOrganizationDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AddOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOrganization`: ApiOrganizationDTO
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AddOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiOrganizationDTO** | [**ApiOrganizationDTO**](ApiOrganizationDTO.md) |  | 

### Return type

[**ApiOrganizationDTO**](ApiOrganizationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> ApiOrganizationDTO GetOrganization(ctx, organizationId).Execute()



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
    resp, r, err := apiClient.OrganizationsAPI.GetOrganization(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganization`: ApiOrganizationDTO
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiOrganizationDTO**](ApiOrganizationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizations

> ApiOrganizationListDTO GetOrganizations(ctx).OrganizationName(organizationName).Execute()



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
    organizationName := []string{"Inner_example"} // []string |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrganizationsAPI.GetOrganizations(context.Background()).OrganizationName(organizationName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrganizations`: ApiOrganizationListDTO
    fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationName** | **[]string** |  | 

### Return type

[**ApiOrganizationListDTO**](ApiOrganizationListDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveOrganization

> MoveOrganization(ctx, organizationId, destinationId).FailEarlyOnError(failEarlyOnError).Execute()



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
    destinationId := "destinationId_example" // string | 
    failEarlyOnError := true // bool |  (optional) (default to false)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.OrganizationsAPI.MoveOrganization(context.Background(), organizationId, destinationId).FailEarlyOnError(failEarlyOnError).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.MoveOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**destinationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **failEarlyOnError** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

