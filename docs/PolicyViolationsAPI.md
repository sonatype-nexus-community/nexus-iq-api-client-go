# \PolicyViolationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicableWaivers**](PolicyViolationsAPI.md#GetApplicableWaivers) | **Get** /api/v2/policyViolations/{violationId}/applicableWaivers | 
[**GetCrossStagePolicyViolationByConstituentId**](PolicyViolationsAPI.md#GetCrossStagePolicyViolationByConstituentId) | **Get** /api/v2/policyViolations/crossStage | 
[**GetCrossStagePolicyViolationById**](PolicyViolationsAPI.md#GetCrossStagePolicyViolationById) | **Get** /api/v2/policyViolations/crossStage/{violationId} | 
[**GetPolicyViolations**](PolicyViolationsAPI.md#GetPolicyViolations) | **Get** /api/v2/policyViolations | 
[**GetTransitivePolicyViolationsByAppScanComponent**](PolicyViolationsAPI.md#GetTransitivePolicyViolationsByAppScanComponent) | **Get** /api/v2/policyViolations/transitive/{ownerType}/{ownerId}/{scanId} | 
[**GetTransitivePolicyViolationsByOwnerStageComponent**](PolicyViolationsAPI.md#GetTransitivePolicyViolationsByOwnerStageComponent) | **Get** /api/v2/policyViolations/transitive/{ownerType}/{ownerId}/stages/{stageId} | 



## GetApplicableWaivers

> ApiPolicyWaiversApplicableToViolationDTO GetApplicableWaivers(ctx, violationId).Execute()



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
    violationId := "violationId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyViolationsAPI.GetApplicableWaivers(context.Background(), violationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationsAPI.GetApplicableWaivers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicableWaivers`: ApiPolicyWaiversApplicableToViolationDTO
    fmt.Fprintf(os.Stdout, "Response from `PolicyViolationsAPI.GetApplicableWaivers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**violationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableWaiversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiPolicyWaiversApplicableToViolationDTO**](ApiPolicyWaiversApplicableToViolationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossStagePolicyViolationByConstituentId

> ApiCrossStageViolationDTOV2 GetCrossStagePolicyViolationByConstituentId(ctx).ConstituentId(constituentId).Execute()



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
    constituentId := "constituentId_example" // string |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyViolationsAPI.GetCrossStagePolicyViolationByConstituentId(context.Background()).ConstituentId(constituentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationsAPI.GetCrossStagePolicyViolationByConstituentId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrossStagePolicyViolationByConstituentId`: ApiCrossStageViolationDTOV2
    fmt.Fprintf(os.Stdout, "Response from `PolicyViolationsAPI.GetCrossStagePolicyViolationByConstituentId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossStagePolicyViolationByConstituentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **constituentId** | **string** |  | 

### Return type

[**ApiCrossStageViolationDTOV2**](ApiCrossStageViolationDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossStagePolicyViolationById

> ApiCrossStageViolationDTOV2 GetCrossStagePolicyViolationById(ctx, violationId).Execute()



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
    violationId := "violationId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyViolationsAPI.GetCrossStagePolicyViolationById(context.Background(), violationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationsAPI.GetCrossStagePolicyViolationById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCrossStagePolicyViolationById`: ApiCrossStageViolationDTOV2
    fmt.Fprintf(os.Stdout, "Response from `PolicyViolationsAPI.GetCrossStagePolicyViolationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**violationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossStagePolicyViolationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiCrossStageViolationDTOV2**](ApiCrossStageViolationDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyViolations

> ApiApplicationViolationListDTOV2 GetPolicyViolations(ctx).P(p).Execute()



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
    p := []string{"Inner_example"} // []string |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyViolationsAPI.GetPolicyViolations(context.Background()).P(p).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationsAPI.GetPolicyViolations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyViolations`: ApiApplicationViolationListDTOV2
    fmt.Fprintf(os.Stdout, "Response from `PolicyViolationsAPI.GetPolicyViolations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **p** | **[]string** |  | 

### Return type

[**ApiApplicationViolationListDTOV2**](ApiApplicationViolationListDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitivePolicyViolationsByAppScanComponent

> ApiComponentTransitivePolicyViolationsDTO GetTransitivePolicyViolationsByAppScanComponent(ctx, ownerType, ownerId, scanId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()



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
    scanId := "scanId_example" // string | 
    componentIdentifier := map[string][]iqclient.ComponentIdentifier{ ... } // ComponentIdentifier |  (optional)
    packageUrl := "packageUrl_example" // string |  (optional)
    hash := "hash_example" // string |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyViolationsAPI.GetTransitivePolicyViolationsByAppScanComponent(context.Background(), ownerType, ownerId, scanId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationsAPI.GetTransitivePolicyViolationsByAppScanComponent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransitivePolicyViolationsByAppScanComponent`: ApiComponentTransitivePolicyViolationsDTO
    fmt.Fprintf(os.Stdout, "Response from `PolicyViolationsAPI.GetTransitivePolicyViolationsByAppScanComponent`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetTransitivePolicyViolationsByAppScanComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) |  | 
 **packageUrl** | **string** |  | 
 **hash** | **string** |  | 

### Return type

[**ApiComponentTransitivePolicyViolationsDTO**](ApiComponentTransitivePolicyViolationsDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitivePolicyViolationsByOwnerStageComponent

> ApiComponentTransitivePolicyViolationsDTO GetTransitivePolicyViolationsByOwnerStageComponent(ctx, ownerType, ownerId, stageId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()



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
    stageId := "stageId_example" // string | 
    componentIdentifier := map[string][]iqclient.ComponentIdentifier{ ... } // ComponentIdentifier |  (optional)
    packageUrl := "packageUrl_example" // string |  (optional)
    hash := "hash_example" // string |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyViolationsAPI.GetTransitivePolicyViolationsByOwnerStageComponent(context.Background(), ownerType, ownerId, stageId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationsAPI.GetTransitivePolicyViolationsByOwnerStageComponent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransitivePolicyViolationsByOwnerStageComponent`: ApiComponentTransitivePolicyViolationsDTO
    fmt.Fprintf(os.Stdout, "Response from `PolicyViolationsAPI.GetTransitivePolicyViolationsByOwnerStageComponent`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetTransitivePolicyViolationsByOwnerStageComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) |  | 
 **packageUrl** | **string** |  | 
 **hash** | **string** |  | 

### Return type

[**ApiComponentTransitivePolicyViolationsDTO**](ApiComponentTransitivePolicyViolationsDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

