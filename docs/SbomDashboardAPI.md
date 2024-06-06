# \SbomDashboardAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationsHistoryMetric**](SbomDashboardAPI.md#GetApplicationsHistoryMetric) | **Get** /api/v2/sbom/dashboard/sbomsHistoryMetrics | Gets application history metrics
[**GetSbomsAnalyzedMetrics**](SbomDashboardAPI.md#GetSbomsAnalyzedMetrics) | **Get** /api/v2/sbom/dashboard/sbomsAnalyzed | Gets total of SBOMs analyzed and the threshold in the product license
[**GetVulnerabilitiesByThreatLevel**](SbomDashboardAPI.md#GetVulnerabilitiesByThreatLevel) | **Get** /api/v2/sbom/dashboard/vulnerabilitiesByThreatLevel | Gets counters of vulnerabilities and annotations by threat level



## GetApplicationsHistoryMetric

> GetApplicationsHistoryMetric(ctx).Execute()

Gets application history metrics



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
	r, err := apiClient.SbomDashboardAPI.GetApplicationsHistoryMetric(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomDashboardAPI.GetApplicationsHistoryMetric``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsHistoryMetricRequest struct via the builder pattern


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


## GetVulnerabilitiesByThreatLevel

> GetVulnerabilitiesByThreatLevel(ctx).Execute()

Gets counters of vulnerabilities and annotations by threat level



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
	r, err := apiClient.SbomDashboardAPI.GetVulnerabilitiesByThreatLevel(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomDashboardAPI.GetVulnerabilitiesByThreatLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVulnerabilitiesByThreatLevelRequest struct via the builder pattern


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

