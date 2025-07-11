# \LicenseLegalMetadataAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLicenseLegalApplicationHTMLReport**](LicenseLegalMetadataAPI.md#GetLicenseLegalApplicationHTMLReport) | **Get** /api/v2/licenseLegalMetadata/application/{applicationId}/stage/{stageId}/report | 
[**GetLicenseLegalApplicationReport**](LicenseLegalMetadataAPI.md#GetLicenseLegalApplicationReport) | **Get** /api/v2/licenseLegalMetadata/application/{applicationId} | 
[**GetLicenseLegalApplicationReport1**](LicenseLegalMetadataAPI.md#GetLicenseLegalApplicationReport1) | **Get** /api/v2/licenseLegalMetadata/application/{applicationId}/stage/{stageId} | 
[**GetLicenseLegalComponentReport**](LicenseLegalMetadataAPI.md#GetLicenseLegalComponentReport) | **Get** /api/v2/licenseLegalMetadata/{ownerType}/{ownerId}/component | 
[**GetLicenseLegalCustomApplicationHTMLReport**](LicenseLegalMetadataAPI.md#GetLicenseLegalCustomApplicationHTMLReport) | **Post** /api/v2/licenseLegalMetadata/application/{applicationId}/stage/{stageId}/report | 
[**GetLicenseLegalCustomApplicationHTMLReport1**](LicenseLegalMetadataAPI.md#GetLicenseLegalCustomApplicationHTMLReport1) | **Post** /api/v2/licenseLegalMetadata/application/{applicationId}/stage/{stageId}/report/templateId/{templateId} | 
[**GetLicenseLegalCustomMultiApplicationHTMLReport1**](LicenseLegalMetadataAPI.md#GetLicenseLegalCustomMultiApplicationHTMLReport1) | **Post** /api/v2/licenseLegalMetadata/multiApplication/report/templateId/{templateId} | 
[**GetLicenseLegalMultiApplicationHTMLReport**](LicenseLegalMetadataAPI.md#GetLicenseLegalMultiApplicationHTMLReport) | **Post** /api/v2/licenseLegalMetadata/multiApplication/report | 
[**GetLicenseLegalMultiApplicationReportFromActiveUserFilter**](LicenseLegalMetadataAPI.md#GetLicenseLegalMultiApplicationReportFromActiveUserFilter) | **Post** /api/v2/licenseLegalMetadata/multiApplication/activeUserFilter/report/templateId/{templateId} | 



## GetLicenseLegalApplicationHTMLReport

> GetLicenseLegalApplicationHTMLReport(ctx, applicationId, stageId).Execute()





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
	applicationId := "applicationId_example" // string | Enter the application id or public id.
	stageId := "stageId_example" // string | Enter the stageId.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalApplicationHTMLReport(context.Background(), applicationId, stageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalApplicationHTMLReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the application id or public id. | 
**stageId** | **string** | Enter the stageId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalApplicationHTMLReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

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
	applicationId := "applicationId_example" // string | Enter the application id or public id.

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
**applicationId** | **string** | Enter the application id or public id. | 

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
	applicationId := "applicationId_example" // string | Enter the application id or public id.
	stageId := "stageId_example" // string | Enter the stageId.

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
**applicationId** | **string** | Enter the application id or public id. | 
**stageId** | **string** | Enter the stageId. | 

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
	ownerType := "ownerType_example" // string | Enter the ownerType
	ownerId := "ownerId_example" // string | Enter the ownerId corresponding to the ownerType.
	componentIdentifier := *sonatypeiq.NewComponentIdentifier() // ComponentIdentifier | Enter the componentIdentifier consisting of format and coordinates. (optional)
	packageUrl := "packageUrl_example" // string | Enter the package URL. (optional)
	hash := "hash_example" // string | Enter the component hash. (optional)
	identificationSource := "identificationSource_example" // string | Enter the identification source if a third party scan is used. (optional)
	scanId := "scanId_example" // string | Enter the scanId for the report where the component was identified (required if identified by third party scan). (optional)

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
**ownerType** | **string** | Enter the ownerType | 
**ownerId** | **string** | Enter the ownerId corresponding to the ownerType. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalComponentReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) | Enter the componentIdentifier consisting of format and coordinates. | 
 **packageUrl** | **string** | Enter the package URL. | 
 **hash** | **string** | Enter the component hash. | 
 **identificationSource** | **string** | Enter the identification source if a third party scan is used. | 
 **scanId** | **string** | Enter the scanId for the report where the component was identified (required if identified by third party scan). | 

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

> GetLicenseLegalCustomApplicationHTMLReport(ctx, applicationId, stageId).Execute()





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
	applicationId := "applicationId_example" // string | Enter the application id or public id.
	stageId := "stageId_example" // string | Enter the stageId.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalCustomApplicationHTMLReport(context.Background(), applicationId, stageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalCustomApplicationHTMLReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the application id or public id. | 
**stageId** | **string** | Enter the stageId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalCustomApplicationHTMLReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalCustomApplicationHTMLReport1

> GetLicenseLegalCustomApplicationHTMLReport1(ctx, applicationId, stageId, templateId).Execute()





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
	applicationId := "applicationId_example" // string | Enter the application id or public id.
	stageId := "stageId_example" // string | Enter the stageId.
	templateId := "templateId_example" // string | Enter the templateId for the HTML report format.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalCustomApplicationHTMLReport1(context.Background(), applicationId, stageId, templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalCustomApplicationHTMLReport1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the application id or public id. | 
**stageId** | **string** | Enter the stageId. | 
**templateId** | **string** | Enter the templateId for the HTML report format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalCustomApplicationHTMLReport1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalCustomMultiApplicationHTMLReport1

> GetLicenseLegalCustomMultiApplicationHTMLReport1(ctx, templateId).Execute()





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
	templateId := "templateId_example" // string | Enter the `templateId` for the HTML report.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalCustomMultiApplicationHTMLReport1(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalCustomMultiApplicationHTMLReport1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Enter the &#x60;templateId&#x60; for the HTML report. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalCustomMultiApplicationHTMLReport1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalMultiApplicationHTMLReport

> GetLicenseLegalMultiApplicationHTMLReport(ctx).Execute()





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
	r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalMultiApplicationHTMLReport(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalMultiApplicationHTMLReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalMultiApplicationHTMLReportRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicenseLegalMultiApplicationReportFromActiveUserFilter

> GetLicenseLegalMultiApplicationReportFromActiveUserFilter(ctx, templateId).Execute()





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
	templateId := "templateId_example" // string | Enter the templateId for the license legal data.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.LicenseLegalMetadataAPI.GetLicenseLegalMultiApplicationReportFromActiveUserFilter(context.Background(), templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseLegalMetadataAPI.GetLicenseLegalMultiApplicationReportFromActiveUserFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **string** | Enter the templateId for the license legal data. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseLegalMultiApplicationReportFromActiveUserFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

