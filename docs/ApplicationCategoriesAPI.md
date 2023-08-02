# \ApplicationCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTag**](ApplicationCategoriesAPI.md#AddTag) | **Post** /api/v2/applicationCategories/organization/{organizationId} | 
[**DeleteTag**](ApplicationCategoriesAPI.md#DeleteTag) | **Delete** /api/v2/applicationCategories/organization/{organizationId}/{tagId} | 
[**GetApplicableTags**](ApplicationCategoriesAPI.md#GetApplicableTags) | **Get** /api/v2/applicationCategories/organization/{organizationId}/applicable | 
[**GetApplicableTagsByApplicationPublicId**](ApplicationCategoriesAPI.md#GetApplicableTagsByApplicationPublicId) | **Get** /api/v2/applicationCategories/application/{applicationPublicId}/applicable | 
[**GetApplicationApplicableTags**](ApplicationCategoriesAPI.md#GetApplicationApplicableTags) | **Get** /api/v2/applicationCategories/application/{applicationPublicId} | 
[**GetAppliedPolicyTags**](ApplicationCategoriesAPI.md#GetAppliedPolicyTags) | **Get** /api/v2/applicationCategories/organization/{organizationId}/policy | 
[**GetAppliedTags**](ApplicationCategoriesAPI.md#GetAppliedTags) | **Get** /api/v2/applicationCategories/organization/{organizationId}/applied | 
[**GetTags**](ApplicationCategoriesAPI.md#GetTags) | **Get** /api/v2/applicationCategories/organization/{organizationId} | 
[**GetTagsUsedByApplications**](ApplicationCategoriesAPI.md#GetTagsUsedByApplications) | **Get** /api/v2/applicationCategories/application | 
[**UpdateTag**](ApplicationCategoriesAPI.md#UpdateTag) | **Put** /api/v2/applicationCategories/organization/{organizationId} | 



## AddTag

> ApiApplicationCategoryDTO AddTag(ctx, organizationId).ApiApplicationCategoryDTO(apiApplicationCategoryDTO).Execute()



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
    organizationId := "organizationId_example" // string | 
    apiApplicationCategoryDTO := *sonatypeiq.NewApiApplicationCategoryDTO() // ApiApplicationCategoryDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCategoriesAPI.AddTag(context.Background(), organizationId).ApiApplicationCategoryDTO(apiApplicationCategoryDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.AddTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddTag`: ApiApplicationCategoryDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.AddTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiApplicationCategoryDTO** | [**ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md) |  | 

### Return type

[**ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteTag(ctx, organizationId, tagId).Execute()



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
    organizationId := "organizationId_example" // string | 
    tagId := "tagId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.ApplicationCategoriesAPI.DeleteTag(context.Background(), organizationId, tagId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.DeleteTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 
**tagId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


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


## GetApplicableTags

> ApplicableTagsDTO GetApplicableTags(ctx, organizationId).Execute()



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
    organizationId := "organizationId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCategoriesAPI.GetApplicableTags(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetApplicableTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicableTags`: ApplicableTagsDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetApplicableTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicableTagsDTO**](ApplicableTagsDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicableTagsByApplicationPublicId

> []ApiApplicationCategoryDTO GetApplicableTagsByApplicationPublicId(ctx, applicationPublicId).Execute()



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
    applicationPublicId := "applicationPublicId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCategoriesAPI.GetApplicableTagsByApplicationPublicId(context.Background(), applicationPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetApplicableTagsByApplicationPublicId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicableTagsByApplicationPublicId`: []ApiApplicationCategoryDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetApplicableTagsByApplicationPublicId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableTagsByApplicationPublicIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetApplicationApplicableTags

> ApplicableTagsDTO GetApplicationApplicableTags(ctx, applicationPublicId).Execute()



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
    applicationPublicId := "applicationPublicId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCategoriesAPI.GetApplicationApplicableTags(context.Background(), applicationPublicId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetApplicationApplicableTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationApplicableTags`: ApplicableTagsDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetApplicationApplicableTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationApplicableTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicableTagsDTO**](ApplicableTagsDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppliedPolicyTags

> []PolicyTag GetAppliedPolicyTags(ctx, organizationId).Execute()



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
    organizationId := "organizationId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCategoriesAPI.GetAppliedPolicyTags(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetAppliedPolicyTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppliedPolicyTags`: []PolicyTag
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetAppliedPolicyTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppliedPolicyTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PolicyTag**](PolicyTag.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppliedTags

> AppliedTagsDTO GetAppliedTags(ctx, organizationId).Execute()



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
    organizationId := "organizationId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCategoriesAPI.GetAppliedTags(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetAppliedTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppliedTags`: AppliedTagsDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetAppliedTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppliedTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppliedTagsDTO**](AppliedTagsDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> []ApiApplicationCategoryDTO GetTags(ctx, organizationId).Execute()



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
    organizationId := "organizationId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCategoriesAPI.GetTags(context.Background(), organizationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTags`: []ApiApplicationCategoryDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetTagsUsedByApplications

> []ApiApplicationCategoryDTO GetTagsUsedByApplications(ctx).Execute()



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
    resp, r, err := apiClient.ApplicationCategoriesAPI.GetTagsUsedByApplications(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetTagsUsedByApplications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTagsUsedByApplications`: []ApiApplicationCategoryDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetTagsUsedByApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsUsedByApplicationsRequest struct via the builder pattern


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


## UpdateTag

> ApiApplicationCategoryDTO UpdateTag(ctx, organizationId).ApiApplicationCategoryDTO(apiApplicationCategoryDTO).Execute()



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
    organizationId := "organizationId_example" // string | 
    apiApplicationCategoryDTO := *sonatypeiq.NewApiApplicationCategoryDTO() // ApiApplicationCategoryDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCategoriesAPI.UpdateTag(context.Background(), organizationId).ApiApplicationCategoryDTO(apiApplicationCategoryDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.UpdateTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTag`: ApiApplicationCategoryDTO
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.UpdateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiApplicationCategoryDTO** | [**ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md) |  | 

### Return type

[**ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

