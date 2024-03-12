# \ScanAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIdeUsersOverview**](ScanAPI.md#GetIdeUsersOverview) | **Get** /api/v2/scan/applications/ideUser/overview | 
[**GetScanStatus**](ScanAPI.md#GetScanStatus) | **Get** /api/v2/scan/applications/{applicationId}/status/{scanRequestId} | 
[**ScanComponents**](ScanAPI.md#ScanComponents) | **Post** /api/v2/scan/applications/{applicationId}/sources/{source} | 



## GetIdeUsersOverview

> IdeUsersOverviewDTO GetIdeUsersOverview(ctx).SinceUtcTimestamp(sinceUtcTimestamp).Execute()



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
	sinceUtcTimestamp := int64(789) // int64 |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ScanAPI.GetIdeUsersOverview(context.Background()).SinceUtcTimestamp(sinceUtcTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScanAPI.GetIdeUsersOverview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdeUsersOverview`: IdeUsersOverviewDTO
	fmt.Fprintf(os.Stdout, "Response from `ScanAPI.GetIdeUsersOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdeUsersOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sinceUtcTimestamp** | **int64** |  | 

### Return type

[**IdeUsersOverviewDTO**](IdeUsersOverviewDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScanStatus

> ApiThirdPartyScanResultDTO GetScanStatus(ctx, applicationId, scanRequestId).Execute()



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
	scanRequestId := "scanRequestId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ScanAPI.GetScanStatus(context.Background(), applicationId, scanRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScanAPI.GetScanStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScanStatus`: ApiThirdPartyScanResultDTO
	fmt.Fprintf(os.Stdout, "Response from `ScanAPI.GetScanStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**scanRequestId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScanStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiThirdPartyScanResultDTO**](ApiThirdPartyScanResultDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScanComponents

> ApiThirdPartyScanTicketDTO ScanComponents(ctx, applicationId, source).StageId(stageId).Body(body).Execute()



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
	source := "source_example" // string | 
	stageId := "stageId_example" // string |  (optional) (default to "build")
	body := "body_example" // string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ScanAPI.ScanComponents(context.Background(), applicationId, source).StageId(stageId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScanAPI.ScanComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScanComponents`: ApiThirdPartyScanTicketDTO
	fmt.Fprintf(os.Stdout, "Response from `ScanAPI.ScanComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 
**source** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScanComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stageId** | **string** |  | [default to &quot;build&quot;]
 **body** | **string** |  | 

### Return type

[**ApiThirdPartyScanTicketDTO**](ApiThirdPartyScanTicketDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

