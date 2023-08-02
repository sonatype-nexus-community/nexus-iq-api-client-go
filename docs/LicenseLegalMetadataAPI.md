# \LicenseLegalMetadataAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAttributionReportTemplate**](LicenseLegalMetadataAPI.md#DeleteAttributionReportTemplate) | **Delete** /api/v2/licenseLegalMetadata/report-template/{id} | 
[**GetAllAttributionReportTemplates**](LicenseLegalMetadataAPI.md#GetAllAttributionReportTemplates) | **Get** /api/v2/licenseLegalMetadata/report-template | 
[**GetAttributionReportTemplateById**](LicenseLegalMetadataAPI.md#GetAttributionReportTemplateById) | **Get** /api/v2/licenseLegalMetadata/report-template/{id} | 
[**GetLicenseLegalApplicationHTMLReport**](LicenseLegalMetadataAPI.md#GetLicenseLegalApplicationHTMLReport) | **Get** /api/v2/licenseLegalMetadata/application/{applicationId}/stage/{stageId}/report | 
[**GetLicenseLegalApplicationReport**](LicenseLegalMetadataAPI.md#GetLicenseLegalApplicationReport) | **Get** /api/v2/licenseLegalMetadata/application/{applicationId} | 
[**GetLicenseLegalApplicationReport1**](LicenseLegalMetadataAPI.md#GetLicenseLegalApplicationReport1) | **Get** /api/v2/licenseLegalMetadata/application/{applicationId}/stage/{stageId} | 
[**GetLicenseLegalComponentReport**](LicenseLegalMetadataAPI.md#GetLicenseLegalComponentReport) | **Get** /api/v2/licenseLegalMetadata/{ownerType}/{ownerId}/component | 
[**GetLicenseLegalCustomApplicationHTMLReport**](LicenseLegalMetadataAPI.md#GetLicenseLegalCustomApplicationHTMLReport) | **Post** /api/v2/licenseLegalMetadata/application/{applicationId}/stage/{stageId}/report | 
[**GetLicenseLegalCustomApplicationHTMLReport1**](LicenseLegalMetadataAPI.md#GetLicenseLegalCustomApplicationHTMLReport1) | **Post** /api/v2/licenseLegalMetadata/application/{applicationId}/stage/{stageId}/report/templateId/{templateId} | 
[**GetLicenseLegalCustomMultiApplicationHTMLReport1**](LicenseLegalMetadataAPI.md#GetLicenseLegalCustomMultiApplicationHTMLReport1) | **Post** /api/v2/licenseLegalMetadata/multiApplication/report/templateId/{templateId} | 
[**GetLicenseLegalMultiApplicationHTMLReport**](LicenseLegalMetadataAPI.md#GetLicenseLegalMultiApplicationHTMLReport) | **Post** /api/v2/licenseLegalMetadata/multiApplication/report | 
[**GetLicenseLegalMultiApplicationReportFromActiveUserFilter**](LicenseLegalMetadataAPI.md#GetLicenseLegalMultiApplicationReportFromActiveUserFilter) | **Post** /api/v2/licenseLegalMetadata/multiApplication/activeUserFilter/report/templateId/{templateId} | 
[**SaveAttributionReportTemplate**](LicenseLegalMetadataAPI.md#SaveAttributionReportTemplate) | **Post** /api/v2/licenseLegalMetadata/report-template | 



## DeleteAttributionReportTemplate

> DeleteAttributionReportTemplate(ctx, id).Execute()



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
    id := "id_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.LicenseLegalMetadataAPI.DeleteAttributionReportTemplate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.DeleteAttributionReportTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAttributionReportTemplateRequest struct via the builder pattern


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


## GetAllAttributionReportTemplates

> []AttributionReportTemplateDTO GetAllAttributionReportTemplates(ctx).Execute()



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
    resp, r, err := apiClient.LicenseLegalMetadataAPI.GetAllAttributionReportTemplates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetAllAttributionReportTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllAttributionReportTemplates`: []AttributionReportTemplateDTO
    fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataAPI.GetAllAttributionReportTemplates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAttributionReportTemplatesRequest struct via the builder pattern


### Return type

[**[]AttributionReportTemplateDTO**](AttributionReportTemplateDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttributionReportTemplateById

> AttributionReportTemplateDTO GetAttributionReportTemplateById(ctx, id).Execute()



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
    id := "id_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseLegalMetadataAPI.GetAttributionReportTemplateById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetAttributionReportTemplateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttributionReportTemplateById`: AttributionReportTemplateDTO
    fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataAPI.GetAttributionReportTemplateById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttributionReportTemplateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttributionReportTemplateDTO**](AttributionReportTemplateDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalApplicationHTMLReport

> string GetLicenseLegalApplicationHTMLReport(ctx, applicationId, stageId).Execute()



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
    applicationId := "applicationId_example" // string | 
    stageId := "stageId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalApplicationHTMLReport(context.Background(), applicationId, stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalApplicationHTMLReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLegalApplicationHTMLReport`: string
    fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataAPI.GetLicenseLegalApplicationHTMLReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalApplicationHTMLReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalApplicationReport

> ApiLicenseLegalApplicationReportDTO GetLicenseLegalApplicationReport(ctx, applicationId).Execute()



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
    applicationId := "applicationId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalApplicationReport(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalApplicationReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLegalApplicationReport`: ApiLicenseLegalApplicationReportDTO
    fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataAPI.GetLicenseLegalApplicationReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalApplicationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiLicenseLegalApplicationReportDTO**](ApiLicenseLegalApplicationReportDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalApplicationReport1

> ApiLicenseLegalApplicationReportDTO GetLicenseLegalApplicationReport1(ctx, applicationId, stageId).Execute()



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
    applicationId := "applicationId_example" // string | 
    stageId := "stageId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalApplicationReport1(context.Background(), applicationId, stageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalApplicationReport1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLegalApplicationReport1`: ApiLicenseLegalApplicationReportDTO
    fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataAPI.GetLicenseLegalApplicationReport1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalApplicationReport1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiLicenseLegalApplicationReportDTO**](ApiLicenseLegalApplicationReportDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalComponentReport

> ApiLicenseLegalComponentReportDTO GetLicenseLegalComponentReport(ctx, ownerType, ownerId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).IdentificationSource(identificationSource).ScanId(scanId).Execute()



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
    componentIdentifier := map[string][]sonatypeiq.ComponentIdentifier{ ... } // ComponentIdentifier |  (optional)
    packageUrl := "packageUrl_example" // string |  (optional)
    hash := "hash_example" // string |  (optional)
    identificationSource := "identificationSource_example" // string |  (optional)
    scanId := "scanId_example" // string |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalComponentReport(context.Background(), ownerType, ownerId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).IdentificationSource(identificationSource).ScanId(scanId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalComponentReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLegalComponentReport`: ApiLicenseLegalComponentReportDTO
    fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataAPI.GetLicenseLegalComponentReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalComponentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) |  | 
 **packageUrl** | **string** |  | 
 **hash** | **string** |  | 
 **identificationSource** | **string** |  | 
 **scanId** | **string** |  | 

### Return type

[**ApiLicenseLegalComponentReportDTO**](ApiLicenseLegalComponentReportDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalCustomApplicationHTMLReport

> string GetLicenseLegalCustomApplicationHTMLReport(ctx, applicationId, stageId).FormDataMultiPart(formDataMultiPart).Execute()



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
    applicationId := "applicationId_example" // string | 
    stageId := "stageId_example" // string | 
    formDataMultiPart := *sonatypeiq.NewFormDataMultiPart() // FormDataMultiPart |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalCustomApplicationHTMLReport(context.Background(), applicationId, stageId).FormDataMultiPart(formDataMultiPart).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalCustomApplicationHTMLReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLegalCustomApplicationHTMLReport`: string
    fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataAPI.GetLicenseLegalCustomApplicationHTMLReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalCustomApplicationHTMLReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **formDataMultiPart** | [**FormDataMultiPart**](FormDataMultiPart.md) |  | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalCustomApplicationHTMLReport1

> string GetLicenseLegalCustomApplicationHTMLReport1(ctx, applicationId, stageId, templateId).Execute()



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
    applicationId := "applicationId_example" // string | 
    stageId := "stageId_example" // string | 
    templateId := "templateId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalCustomApplicationHTMLReport1(context.Background(), applicationId, stageId, templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalCustomApplicationHTMLReport1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLegalCustomApplicationHTMLReport1`: string
    fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataAPI.GetLicenseLegalCustomApplicationHTMLReport1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**stageId** | **string** |  | 
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalCustomApplicationHTMLReport1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalCustomMultiApplicationHTMLReport1

> string GetLicenseLegalCustomMultiApplicationHTMLReport1(ctx, templateId).Execute()



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
    templateId := "templateId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalCustomMultiApplicationHTMLReport1(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalCustomMultiApplicationHTMLReport1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLegalCustomMultiApplicationHTMLReport1`: string
    fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataAPI.GetLicenseLegalCustomMultiApplicationHTMLReport1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalCustomMultiApplicationHTMLReport1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalMultiApplicationHTMLReport

> string GetLicenseLegalMultiApplicationHTMLReport(ctx).Execute()



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
    resp, r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalMultiApplicationHTMLReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalMultiApplicationHTMLReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLegalMultiApplicationHTMLReport`: string
    fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataAPI.GetLicenseLegalMultiApplicationHTMLReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalMultiApplicationHTMLReportRequest struct via the builder pattern


### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalMultiApplicationReportFromActiveUserFilter

> string GetLicenseLegalMultiApplicationReportFromActiveUserFilter(ctx, templateId).Execute()



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
    templateId := "templateId_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalMultiApplicationReportFromActiveUserFilter(context.Background(), templateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalMultiApplicationReportFromActiveUserFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLicenseLegalMultiApplicationReportFromActiveUserFilter`: string
    fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataAPI.GetLicenseLegalMultiApplicationReportFromActiveUserFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalMultiApplicationReportFromActiveUserFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveAttributionReportTemplate

> AttributionReportTemplateDTO SaveAttributionReportTemplate(ctx).AttributionReportTemplateDTO(attributionReportTemplateDTO).Execute()



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
    attributionReportTemplateDTO := *sonatypeiq.NewAttributionReportTemplateDTO() // AttributionReportTemplateDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.LicenseLegalMetadataAPI.SaveAttributionReportTemplate(context.Background()).AttributionReportTemplateDTO(attributionReportTemplateDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.SaveAttributionReportTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveAttributionReportTemplate`: AttributionReportTemplateDTO
    fmt.Fprintf(os.Stdout, "Response from `LicenseLegalMetadataAPI.SaveAttributionReportTemplate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSaveAttributionReportTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attributionReportTemplateDTO** | [**AttributionReportTemplateDTO**](AttributionReportTemplateDTO.md) |  | 

### Return type

[**AttributionReportTemplateDTO**](AttributionReportTemplateDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

