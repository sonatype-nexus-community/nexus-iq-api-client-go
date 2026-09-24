# \ReachabilityEvidenceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReachabilityEvidence**](ReachabilityEvidenceAPI.md#GetReachabilityEvidence) | **Get** /api/v2/applications/{applicationPublicId}/reports/{reportId}/vulnerabilities/{vulnerabilityId}/reachability-evidence | 



## GetReachabilityEvidence

> ApiReachabilityEvidenceResponse GetReachabilityEvidence(ctx, applicationPublicId, reportId, vulnerabilityId).Execute()





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
	applicationPublicId := "applicationPublicId_example" // string | The public ID of the application
	reportId := "reportId_example" // string | The report/scan ID
	vulnerabilityId := "vulnerabilityId_example" // string | The vulnerability ID (e.g., CVE-2023-35116)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ReachabilityEvidenceAPI.GetReachabilityEvidence(context.Background(), applicationPublicId, reportId, vulnerabilityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReachabilityEvidenceAPI.GetReachabilityEvidence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReachabilityEvidence`: ApiReachabilityEvidenceResponse
	fmt.Fprintf(os.Stdout, "Response from `ReachabilityEvidenceAPI.GetReachabilityEvidence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** | The public ID of the application | 
**reportId** | **string** | The report/scan ID | 
**vulnerabilityId** | **string** | The vulnerability ID (e.g., CVE-2023-35116) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReachabilityEvidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiReachabilityEvidenceResponse**](ApiReachabilityEvidenceResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

