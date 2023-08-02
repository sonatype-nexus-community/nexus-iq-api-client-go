# \PolicyWaiversAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPolicyWaiverByPolicyViolationId**](PolicyWaiversAPI.md#AddPolicyWaiverByPolicyViolationId) | **Post** /api/v2/policyWaivers/{ownerType}/{ownerId}/{policyViolationId} | 
[**AddWaiverToTransitivePolicyViolationsByAppScanComponent**](PolicyWaiversAPI.md#AddWaiverToTransitivePolicyViolationsByAppScanComponent) | **Post** /api/v2/policyWaivers/transitive/{ownerType}/{ownerId}/{scanId} | 
[**AddWaiverToTransitivePolicyViolationsByOwnerStageComponent**](PolicyWaiversAPI.md#AddWaiverToTransitivePolicyViolationsByOwnerStageComponent) | **Post** /api/v2/policyWaivers/transitive/{ownerType}/{ownerId}/stages/{stageId} | 
[**DeletePolicyWaiver**](PolicyWaiversAPI.md#DeletePolicyWaiver) | **Delete** /api/v2/policyWaivers/{ownerType}/{ownerId}/{policyWaiverId} | 
[**GetPolicyWaiver**](PolicyWaiversAPI.md#GetPolicyWaiver) | **Get** /api/v2/policyWaivers/{ownerType}/{ownerId}/{policyWaiverId} | 
[**GetPolicyWaivers**](PolicyWaiversAPI.md#GetPolicyWaivers) | **Get** /api/v2/policyWaivers/{ownerType}/{ownerId} | 
[**GetTransitivePolicyWaiversByAppScanComponent**](PolicyWaiversAPI.md#GetTransitivePolicyWaiversByAppScanComponent) | **Get** /api/v2/policyWaivers/transitive/{ownerType}/{ownerId}/{scanId} | 



## AddPolicyWaiverByPolicyViolationId

> AddPolicyWaiverByPolicyViolationId(ctx, ownerType, ownerId, policyViolationId).ApiWaiverOptionsDTO(apiWaiverOptionsDTO).Execute()



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
    ownerId := "ownerId_example" // string | 
    policyViolationId := "policyViolationId_example" // string | 
    apiWaiverOptionsDTO := *sonatypeiq.NewApiWaiverOptionsDTO() // ApiWaiverOptionsDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.PolicyWaiversAPI.AddPolicyWaiverByPolicyViolationId(context.Background(), ownerType, ownerId, policyViolationId).ApiWaiverOptionsDTO(apiWaiverOptionsDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.AddPolicyWaiverByPolicyViolationId``: %v\n", err)
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
**policyViolationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPolicyWaiverByPolicyViolationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiWaiverOptionsDTO** | [**ApiWaiverOptionsDTO**](ApiWaiverOptionsDTO.md) |  | 

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


## AddWaiverToTransitivePolicyViolationsByAppScanComponent

> AddWaiverToTransitivePolicyViolationsByAppScanComponent(ctx, ownerType, ownerId, scanId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).ApiWaiverOptionsDTO(apiWaiverOptionsDTO).Execute()



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
    ownerId := "ownerId_example" // string | 
    scanId := "scanId_example" // string | 
    componentIdentifier := map[string][]sonatypeiq.ComponentIdentifier{ ... } // ComponentIdentifier |  (optional)
    packageUrl := "packageUrl_example" // string |  (optional)
    hash := "hash_example" // string |  (optional)
    apiWaiverOptionsDTO := *sonatypeiq.NewApiWaiverOptionsDTO() // ApiWaiverOptionsDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.PolicyWaiversAPI.AddWaiverToTransitivePolicyViolationsByAppScanComponent(context.Background(), ownerType, ownerId, scanId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).ApiWaiverOptionsDTO(apiWaiverOptionsDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.AddWaiverToTransitivePolicyViolationsByAppScanComponent``: %v\n", err)
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
**scanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) |  | 
 **packageUrl** | **string** |  | 
 **hash** | **string** |  | 
 **apiWaiverOptionsDTO** | [**ApiWaiverOptionsDTO**](ApiWaiverOptionsDTO.md) |  | 

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


## AddWaiverToTransitivePolicyViolationsByOwnerStageComponent

> AddWaiverToTransitivePolicyViolationsByOwnerStageComponent(ctx, ownerType, ownerId, stageId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).ApiWaiverOptionsDTO(apiWaiverOptionsDTO).Execute()



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
    ownerId := "ownerId_example" // string | 
    stageId := "stageId_example" // string | 
    componentIdentifier := map[string][]sonatypeiq.ComponentIdentifier{ ... } // ComponentIdentifier |  (optional)
    packageUrl := "packageUrl_example" // string |  (optional)
    hash := "hash_example" // string |  (optional)
    apiWaiverOptionsDTO := *sonatypeiq.NewApiWaiverOptionsDTO() // ApiWaiverOptionsDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.PolicyWaiversAPI.AddWaiverToTransitivePolicyViolationsByOwnerStageComponent(context.Background(), ownerType, ownerId, stageId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).ApiWaiverOptionsDTO(apiWaiverOptionsDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.AddWaiverToTransitivePolicyViolationsByOwnerStageComponent``: %v\n", err)
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
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) |  | 
 **packageUrl** | **string** |  | 
 **hash** | **string** |  | 
 **apiWaiverOptionsDTO** | [**ApiWaiverOptionsDTO**](ApiWaiverOptionsDTO.md) |  | 

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


## DeletePolicyWaiver

> DeletePolicyWaiver(ctx, ownerType, ownerId, policyWaiverId).Execute()



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
    ownerId := "ownerId_example" // string | 
    policyWaiverId := "policyWaiverId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.PolicyWaiversAPI.DeletePolicyWaiver(context.Background(), ownerType, ownerId, policyWaiverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.DeletePolicyWaiver``: %v\n", err)
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
**policyWaiverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyWaiverRequest struct via the builder pattern


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


## GetPolicyWaiver

> ApiPolicyWaiverDTO GetPolicyWaiver(ctx, ownerType, ownerId, policyWaiverId).Execute()



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
    ownerId := "ownerId_example" // string | 
    policyWaiverId := "policyWaiverId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyWaiversAPI.GetPolicyWaiver(context.Background(), ownerType, ownerId, policyWaiverId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.GetPolicyWaiver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyWaiver`: ApiPolicyWaiverDTO
    fmt.Fprintf(os.Stdout, "Response from `PolicyWaiversAPI.GetPolicyWaiver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 
**policyWaiverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyWaiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiPolicyWaiverDTO**](ApiPolicyWaiverDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyWaivers

> []ApiPolicyWaiverDTO GetPolicyWaivers(ctx, ownerType, ownerId).Execute()



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
    ownerId := "ownerId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyWaiversAPI.GetPolicyWaivers(context.Background(), ownerType, ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.GetPolicyWaivers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyWaivers`: []ApiPolicyWaiverDTO
    fmt.Fprintf(os.Stdout, "Response from `PolicyWaiversAPI.GetPolicyWaivers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyWaiversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ApiPolicyWaiverDTO**](ApiPolicyWaiverDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitivePolicyWaiversByAppScanComponent

> ApiComponentPolicyWaiversDTO GetTransitivePolicyWaiversByAppScanComponent(ctx, ownerType, ownerId, scanId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()



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
    ownerId := "ownerId_example" // string | 
    scanId := "scanId_example" // string | 
    componentIdentifier := map[string][]sonatypeiq.ComponentIdentifier{ ... } // ComponentIdentifier |  (optional)
    packageUrl := "packageUrl_example" // string |  (optional)
    hash := "hash_example" // string |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyWaiversAPI.GetTransitivePolicyWaiversByAppScanComponent(context.Background(), ownerType, ownerId, scanId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.GetTransitivePolicyWaiversByAppScanComponent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransitivePolicyWaiversByAppScanComponent`: ApiComponentPolicyWaiversDTO
    fmt.Fprintf(os.Stdout, "Response from `PolicyWaiversAPI.GetTransitivePolicyWaiversByAppScanComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 
**scanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitivePolicyWaiversByAppScanComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) |  | 
 **packageUrl** | **string** |  | 
 **hash** | **string** |  | 

### Return type

[**ApiComponentPolicyWaiversDTO**](ApiComponentPolicyWaiversDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

