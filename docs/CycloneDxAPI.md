# \CycloneDxAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetByReportId**](CycloneDxAPI.md#GetByReportId) | **Get** /api/v2/cycloneDx/{applicationId}/reports/{reportId} | 
[**GetByReportId1**](CycloneDxAPI.md#GetByReportId1) | **Get** /api/v2/cycloneDx/{cdxVersion}/{applicationId}/reports/{reportId} | 
[**GetLatest**](CycloneDxAPI.md#GetLatest) | **Get** /api/v2/cycloneDx/{applicationId}/stages/{stageId} | 
[**GetLatest1**](CycloneDxAPI.md#GetLatest1) | **Get** /api/v2/cycloneDx/{cdxVersion}/{applicationId}/stages/{stageId} | 



## GetByReportId

> GetByReportId(ctx, applicationId, reportId).Execute()



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
	reportId := "reportId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.CycloneDxAPI.GetByReportId(context.Background(), applicationId, reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CycloneDxAPI.GetByReportId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByReportIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetByReportId1

> GetByReportId1(ctx, applicationId, reportId, cdxVersion).Execute()



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
	reportId := "reportId_example" // string | 
	cdxVersion := "cdxVersion_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.CycloneDxAPI.GetByReportId1(context.Background(), applicationId, reportId, cdxVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CycloneDxAPI.GetByReportId1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**reportId** | **string** |  | 
**cdxVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetByReportId1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatest

> GetLatest(ctx, applicationId, stageId).Execute()



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
	r, err := apiClient.CycloneDxAPI.GetLatest(context.Background(), applicationId, stageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CycloneDxAPI.GetLatest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**stageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatest1

> GetLatest1(ctx, applicationId, stageId, cdxVersion).Execute()



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
	cdxVersion := "cdxVersion_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.CycloneDxAPI.GetLatest1(context.Background(), applicationId, stageId, cdxVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CycloneDxAPI.GetLatest1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**stageId** | **string** |  | 
**cdxVersion** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatest1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

