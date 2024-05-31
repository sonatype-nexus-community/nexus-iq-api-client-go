# \SbomDashboardAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSbomsAnalyzedMetrics**](SbomDashboardAPI.md#GetSbomsAnalyzedMetrics) | **Get** /api/v2/sbom/dashboard/sbomsAnalyzed | Gets total of SBOMs analyzed and the threshold in the product license



## GetSbomsAnalyzedMetrics

> GetSbomsAnalyzedMetrics(ctx).Execute()

Gets total of SBOMs analyzed and the threshold in the product license



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
	r, err := apiClient.SbomDashboardAPI.GetSbomsAnalyzedMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomDashboardAPI.GetSbomsAnalyzedMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomsAnalyzedMetricsRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

