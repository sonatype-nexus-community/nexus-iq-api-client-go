# \ClaimAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](ClaimAPI.md#Delete) | **Delete** /api/v2/claim/components/{hash} | 
[**Get**](ClaimAPI.md#Get) | **Get** /api/v2/claim/components/{hash} | 
[**GetAll**](ClaimAPI.md#GetAll) | **Get** /api/v2/claim/components | 
[**Set**](ClaimAPI.md#Set) | **Post** /api/v2/claim/components | 



## Delete

> Delete(ctx, hash).Execute()



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
    hash := "hash_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.ClaimAPI.Delete(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClaimAPI.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


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


## Get

> ApiHashComponentIdentifierDTO Get(ctx, hash).Execute()



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
    hash := "hash_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ClaimAPI.Get(context.Background(), hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClaimAPI.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: ApiHashComponentIdentifierDTO
    fmt.Fprintf(os.Stdout, "Response from `ClaimAPI.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiHashComponentIdentifierDTO**](ApiHashComponentIdentifierDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> ApiHashComponentIdentifiersDTO GetAll(ctx).Execute()



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
    resp, r, err := apiClient.ClaimAPI.GetAll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClaimAPI.GetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAll`: ApiHashComponentIdentifiersDTO
    fmt.Fprintf(os.Stdout, "Response from `ClaimAPI.GetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


### Return type

[**ApiHashComponentIdentifiersDTO**](ApiHashComponentIdentifiersDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Set

> ApiHashComponentIdentifierDTO Set(ctx).ApiHashComponentIdentifierDTO(apiHashComponentIdentifierDTO).Execute()



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
    apiHashComponentIdentifierDTO := *sonatypeiq.NewApiHashComponentIdentifierDTO() // ApiHashComponentIdentifierDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ClaimAPI.Set(context.Background()).ApiHashComponentIdentifierDTO(apiHashComponentIdentifierDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClaimAPI.Set``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Set`: ApiHashComponentIdentifierDTO
    fmt.Fprintf(os.Stdout, "Response from `ClaimAPI.Set`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiHashComponentIdentifierDTO** | [**ApiHashComponentIdentifierDTO**](ApiHashComponentIdentifierDTO.md) |  | 

### Return type

[**ApiHashComponentIdentifierDTO**](ApiHashComponentIdentifierDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

