# \ApplicationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplication**](ApplicationsAPI.md#AddApplication) | **Post** /api/v2/applications | 
[**CloneApplication**](ApplicationsAPI.md#CloneApplication) | **Post** /api/v2/applications/{sourceApplicationId}/clone | 
[**DeleteApplication**](ApplicationsAPI.md#DeleteApplication) | **Delete** /api/v2/applications/{applicationId} | 
[**GetApplication**](ApplicationsAPI.md#GetApplication) | **Get** /api/v2/applications/{applicationId} | 
[**GetApplications**](ApplicationsAPI.md#GetApplications) | **Get** /api/v2/applications | 
[**GetApplicationsByOrganizationId**](ApplicationsAPI.md#GetApplicationsByOrganizationId) | **Get** /api/v2/applications/organization/{organizationId} | 
[**GetData**](ApplicationsAPI.md#GetData) | **Get** /api/v2/applications/{applicationPublicId}/reports/{scanId} | 
[**GetDependencyTree**](ApplicationsAPI.md#GetDependencyTree) | **Get** /api/v2/applications/{applicationPublicId}/reports/{scanId}/dependencyTree | 
[**GetPolicyViolationDiff**](ApplicationsAPI.md#GetPolicyViolationDiff) | **Get** /api/v2/applications/{applicationPublicId}/reports/policyViolations/diff | 
[**GetPolicyViolations1**](ApplicationsAPI.md#GetPolicyViolations1) | **Get** /api/v2/applications/{applicationPublicId}/reports/{scanId}/policy | 
[**GetRawData**](ApplicationsAPI.md#GetRawData) | **Get** /api/v2/applications/{applicationPublicId}/reports/{scanId}/raw | 
[**MoveApplication**](ApplicationsAPI.md#MoveApplication) | **Post** /api/v2/applications/{applicationId}/move/organization/{organizationId} | 
[**UpdateApplication**](ApplicationsAPI.md#UpdateApplication) | **Put** /api/v2/applications/{applicationId} | 



## AddApplication

> ApiApplicationDTO AddApplication(ctx).ApiApplicationDTO(apiApplicationDTO).Execute()



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
    apiApplicationDTO := *iqclient.NewApiApplicationDTO() // ApiApplicationDTO |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.AddApplication(context.Background()).ApiApplicationDTO(apiApplicationDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.AddApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddApplication`: ApiApplicationDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.AddApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiApplicationDTO** | [**ApiApplicationDTO**](ApiApplicationDTO.md) |  | 

### Return type

[**ApiApplicationDTO**](ApiApplicationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneApplication

> ApiApplicationDTO CloneApplication(ctx, sourceApplicationId).ClonedApplicationName(clonedApplicationName).ClonedApplicationPublicId(clonedApplicationPublicId).Execute()



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
    sourceApplicationId := "sourceApplicationId_example" // string | 
    clonedApplicationName := "clonedApplicationName_example" // string |  (optional)
    clonedApplicationPublicId := "clonedApplicationPublicId_example" // string |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.CloneApplication(context.Background(), sourceApplicationId).ClonedApplicationName(clonedApplicationName).ClonedApplicationPublicId(clonedApplicationPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CloneApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloneApplication`: ApiApplicationDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.CloneApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceApplicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clonedApplicationName** | **string** |  | 
 **clonedApplicationPublicId** | **string** |  | 

### Return type

[**ApiApplicationDTO**](ApiApplicationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, applicationId).Execute()



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
    r, err := apiClient.ApplicationsAPI.DeleteApplication(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## GetApplication

> ApiApplicationDTO GetApplication(ctx, applicationId).Execute()



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
    resp, r, err := apiClient.ApplicationsAPI.GetApplication(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplication`: ApiApplicationDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiApplicationDTO**](ApiApplicationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> []ApiApplicationCategoryDTO GetApplications(ctx).PublicId(publicId).IncludeCategories(includeCategories).Execute()



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
    publicId := []string{"Inner_example"} // []string |  (optional)
    includeCategories := true // bool |  (optional) (default to false)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.GetApplications(context.Background()).PublicId(publicId).IncludeCategories(includeCategories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplications`: []ApiApplicationCategoryDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicId** | **[]string** |  | 
 **includeCategories** | **bool** |  | [default to false]

### Return type

[**[]ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationsByOrganizationId

> ApiApplicationListDTO GetApplicationsByOrganizationId(ctx, organizationId).Execute()



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
    resp, r, err := apiClient.ApplicationsAPI.GetApplicationsByOrganizationId(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplicationsByOrganizationId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationsByOrganizationId`: ApiApplicationListDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplicationsByOrganizationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsByOrganizationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiApplicationListDTO**](ApiApplicationListDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetData

> GetData(ctx, applicationPublicId, scanId).Execute()



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
    applicationPublicId := "applicationPublicId_example" // string | 
    scanId := "scanId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.ApplicationsAPI.GetData(context.Background(), applicationPublicId, scanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** |  | 
**scanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataRequest struct via the builder pattern


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


## GetDependencyTree

> ApiDependencyTreeResponseDTO GetDependencyTree(ctx, applicationPublicId, scanId).Execute()



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
    applicationPublicId := "applicationPublicId_example" // string | 
    scanId := "scanId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.GetDependencyTree(context.Background(), applicationPublicId, scanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetDependencyTree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDependencyTree`: ApiDependencyTreeResponseDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetDependencyTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** |  | 
**scanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDependencyTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiDependencyTreeResponseDTO**](ApiDependencyTreeResponseDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyViolationDiff

> ApiPolicyViolationDiffDTO GetPolicyViolationDiff(ctx, applicationPublicId).FromCommit(fromCommit).ToCommit(toCommit).FromPolicyEvaluationId(fromPolicyEvaluationId).ToPolicyEvaluationId(toPolicyEvaluationId).Execute()



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
    applicationPublicId := "applicationPublicId_example" // string | 
    fromCommit := "fromCommit_example" // string |  (optional)
    toCommit := "toCommit_example" // string |  (optional)
    fromPolicyEvaluationId := "fromPolicyEvaluationId_example" // string |  (optional)
    toPolicyEvaluationId := "toPolicyEvaluationId_example" // string |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.GetPolicyViolationDiff(context.Background(), applicationPublicId).FromCommit(fromCommit).ToCommit(toCommit).FromPolicyEvaluationId(fromPolicyEvaluationId).ToPolicyEvaluationId(toPolicyEvaluationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetPolicyViolationDiff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyViolationDiff`: ApiPolicyViolationDiffDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetPolicyViolationDiff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyViolationDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromCommit** | **string** |  | 
 **toCommit** | **string** |  | 
 **fromPolicyEvaluationId** | **string** |  | 
 **toPolicyEvaluationId** | **string** |  | 

### Return type

[**ApiPolicyViolationDiffDTO**](ApiPolicyViolationDiffDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyViolations1

> ApiReportPolicyDataDTOV2 GetPolicyViolations1(ctx, applicationPublicId, scanId).Execute()



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
    applicationPublicId := "applicationPublicId_example" // string | 
    scanId := "scanId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.GetPolicyViolations1(context.Background(), applicationPublicId, scanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetPolicyViolations1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyViolations1`: ApiReportPolicyDataDTOV2
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetPolicyViolations1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** |  | 
**scanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyViolations1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiReportPolicyDataDTOV2**](ApiReportPolicyDataDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawData

> ApiReportRawDataDTOV2 GetRawData(ctx, applicationPublicId, scanId).Execute()



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
    applicationPublicId := "applicationPublicId_example" // string | 
    scanId := "scanId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.GetRawData(context.Background(), applicationPublicId, scanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetRawData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRawData`: ApiReportRawDataDTOV2
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetRawData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** |  | 
**scanId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiReportRawDataDTOV2**](ApiReportRawDataDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveApplication

> ApiMoveApplicationResponseDTOV2 MoveApplication(ctx, applicationId, organizationId).Execute()



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
    organizationId := "organizationId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.MoveApplication(context.Background(), applicationId, organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.MoveApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MoveApplication`: ApiMoveApplicationResponseDTOV2
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.MoveApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiMoveApplicationResponseDTOV2**](ApiMoveApplicationResponseDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> ApiApplicationDTO UpdateApplication(ctx, applicationId).ApiApplicationDTO(apiApplicationDTO).Execute()



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
    apiApplicationDTO := *iqclient.NewApiApplicationDTO() // ApiApplicationDTO |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationsAPI.UpdateApplication(context.Background(), applicationId).ApiApplicationDTO(apiApplicationDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.UpdateApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplication`: ApiApplicationDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiApplicationDTO** | [**ApiApplicationDTO**](ApiApplicationDTO.md) |  | 

### Return type

[**ApiApplicationDTO**](ApiApplicationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

