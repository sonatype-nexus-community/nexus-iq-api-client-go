# \ConfigArtifactoryConnectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddArtifactoryConnection**](ConfigArtifactoryConnectionAPI.md#AddArtifactoryConnection) | **Post** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId} | 
[**DeleteArtifactoryConnection**](ConfigArtifactoryConnectionAPI.md#DeleteArtifactoryConnection) | **Delete** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/{artifactoryConnectionId} | 
[**GetArtifactoryConnection**](ConfigArtifactoryConnectionAPI.md#GetArtifactoryConnection) | **Get** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/{artifactoryConnectionId} | 
[**GetOwnerArtifactoryConnection**](ConfigArtifactoryConnectionAPI.md#GetOwnerArtifactoryConnection) | **Get** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId} | 
[**TestArtifactoryConnection**](ConfigArtifactoryConnectionAPI.md#TestArtifactoryConnection) | **Post** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/test | 
[**TestArtifactoryConnection1**](ConfigArtifactoryConnectionAPI.md#TestArtifactoryConnection1) | **Post** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/{artifactoryConnectionId}/test | 
[**UpdateArtifactoryConnection**](ConfigArtifactoryConnectionAPI.md#UpdateArtifactoryConnection) | **Put** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/{artifactoryConnectionId} | 
[**UpdateOwnerArtifactoryConnectionStatus**](ConfigArtifactoryConnectionAPI.md#UpdateOwnerArtifactoryConnectionStatus) | **Put** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId} | 



## AddArtifactoryConnection

> ApiArtifactoryConnectionDTO AddArtifactoryConnection(ctx, ownerType, internalOwnerId).ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO).Execute()



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
    apiArtifactoryConnectionDTO := *sonatypeiq.NewApiArtifactoryConnectionDTO() // ApiArtifactoryConnectionDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigArtifactoryConnectionAPI.AddArtifactoryConnection(context.Background(), ownerType, internalOwnerId).ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigArtifactoryConnectionAPI.AddArtifactoryConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddArtifactoryConnection`: ApiArtifactoryConnectionDTO
    fmt.Fprintf(os.Stdout, "Response from `ConfigArtifactoryConnectionAPI.AddArtifactoryConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddArtifactoryConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiArtifactoryConnectionDTO** | [**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md) |  | 

### Return type

[**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtifactoryConnection

> DeleteArtifactoryConnection(ctx, ownerType, internalOwnerId, artifactoryConnectionId).Execute()



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
    artifactoryConnectionId := "artifactoryConnectionId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.ConfigArtifactoryConnectionAPI.DeleteArtifactoryConnection(context.Background(), ownerType, internalOwnerId, artifactoryConnectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigArtifactoryConnectionAPI.DeleteArtifactoryConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 
**artifactoryConnectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactoryConnectionRequest struct via the builder pattern


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


## GetArtifactoryConnection

> ApiArtifactoryConnectionDTO GetArtifactoryConnection(ctx, ownerType, internalOwnerId, artifactoryConnectionId).Execute()



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
    artifactoryConnectionId := "artifactoryConnectionId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigArtifactoryConnectionAPI.GetArtifactoryConnection(context.Background(), ownerType, internalOwnerId, artifactoryConnectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigArtifactoryConnectionAPI.GetArtifactoryConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetArtifactoryConnection`: ApiArtifactoryConnectionDTO
    fmt.Fprintf(os.Stdout, "Response from `ConfigArtifactoryConnectionAPI.GetArtifactoryConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 
**artifactoryConnectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactoryConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnerArtifactoryConnection

> ApiOwnerArtifactoryConnectionDTO GetOwnerArtifactoryConnection(ctx, ownerType, internalOwnerId).Inherit(inherit).Execute()



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
    inherit := true // bool |  (optional) (default to false)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigArtifactoryConnectionAPI.GetOwnerArtifactoryConnection(context.Background(), ownerType, internalOwnerId).Inherit(inherit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigArtifactoryConnectionAPI.GetOwnerArtifactoryConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOwnerArtifactoryConnection`: ApiOwnerArtifactoryConnectionDTO
    fmt.Fprintf(os.Stdout, "Response from `ConfigArtifactoryConnectionAPI.GetOwnerArtifactoryConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnerArtifactoryConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inherit** | **bool** |  | [default to false]

### Return type

[**ApiOwnerArtifactoryConnectionDTO**](ApiOwnerArtifactoryConnectionDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestArtifactoryConnection

> ApiStatusDTO TestArtifactoryConnection(ctx, ownerType, internalOwnerId).ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO).Execute()



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
    apiArtifactoryConnectionDTO := *sonatypeiq.NewApiArtifactoryConnectionDTO() // ApiArtifactoryConnectionDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigArtifactoryConnectionAPI.TestArtifactoryConnection(context.Background(), ownerType, internalOwnerId).ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigArtifactoryConnectionAPI.TestArtifactoryConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestArtifactoryConnection`: ApiStatusDTO
    fmt.Fprintf(os.Stdout, "Response from `ConfigArtifactoryConnectionAPI.TestArtifactoryConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestArtifactoryConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiArtifactoryConnectionDTO** | [**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md) |  | 

### Return type

[**ApiStatusDTO**](ApiStatusDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestArtifactoryConnection1

> ApiStatusDTO TestArtifactoryConnection1(ctx, ownerType, internalOwnerId, artifactoryConnectionId).Execute()



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
    artifactoryConnectionId := "artifactoryConnectionId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigArtifactoryConnectionAPI.TestArtifactoryConnection1(context.Background(), ownerType, internalOwnerId, artifactoryConnectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigArtifactoryConnectionAPI.TestArtifactoryConnection1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestArtifactoryConnection1`: ApiStatusDTO
    fmt.Fprintf(os.Stdout, "Response from `ConfigArtifactoryConnectionAPI.TestArtifactoryConnection1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 
**artifactoryConnectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestArtifactoryConnection1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiStatusDTO**](ApiStatusDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifactoryConnection

> ApiArtifactoryConnectionDTO UpdateArtifactoryConnection(ctx, ownerType, internalOwnerId, artifactoryConnectionId).ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO).Execute()



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
    artifactoryConnectionId := "artifactoryConnectionId_example" // string | 
    apiArtifactoryConnectionDTO := *sonatypeiq.NewApiArtifactoryConnectionDTO() // ApiArtifactoryConnectionDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigArtifactoryConnectionAPI.UpdateArtifactoryConnection(context.Background(), ownerType, internalOwnerId, artifactoryConnectionId).ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigArtifactoryConnectionAPI.UpdateArtifactoryConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateArtifactoryConnection`: ApiArtifactoryConnectionDTO
    fmt.Fprintf(os.Stdout, "Response from `ConfigArtifactoryConnectionAPI.UpdateArtifactoryConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 
**artifactoryConnectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactoryConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiArtifactoryConnectionDTO** | [**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md) |  | 

### Return type

[**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOwnerArtifactoryConnectionStatus

> UpdateOwnerArtifactoryConnectionStatus(ctx, ownerType, internalOwnerId).ApiArtifactoryConnectionStatusRequestDTO(apiArtifactoryConnectionStatusRequestDTO).Execute()



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
    apiArtifactoryConnectionStatusRequestDTO := *sonatypeiq.NewApiArtifactoryConnectionStatusRequestDTO() // ApiArtifactoryConnectionStatusRequestDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.ConfigArtifactoryConnectionAPI.UpdateOwnerArtifactoryConnectionStatus(context.Background(), ownerType, internalOwnerId).ApiArtifactoryConnectionStatusRequestDTO(apiArtifactoryConnectionStatusRequestDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigArtifactoryConnectionAPI.UpdateOwnerArtifactoryConnectionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOwnerArtifactoryConnectionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiArtifactoryConnectionStatusRequestDTO** | [**ApiArtifactoryConnectionStatusRequestDTO**](ApiArtifactoryConnectionStatusRequestDTO.md) |  | 

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

