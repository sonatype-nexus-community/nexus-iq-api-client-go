# \AuditLogsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditLogs**](AuditLogsAPI.md#GetAuditLogs) | **Get** /api/v2/auditLogs | 



## GetAuditLogs

> GetAuditLogs(ctx).StartUtcDate(startUtcDate).EndUtcDate(endUtcDate).Execute()



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
	startUtcDate := "startUtcDate_example" // string |  (optional)
	endUtcDate := "endUtcDate_example" // string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.AuditLogsAPI.GetAuditLogs(context.Background()).StartUtcDate(startUtcDate).EndUtcDate(endUtcDate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.GetAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startUtcDate** | **string** |  | 
 **endUtcDate** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

