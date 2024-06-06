# \DataRetentionPoliciesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDataRetentionPolicies**](DataRetentionPoliciesAPI.md#GetDataRetentionPolicies) | **Get** /api/v2/dataRetentionPolicies/organizations/{organizationId} | 
[**GetParentDataRetentionPolicies**](DataRetentionPoliciesAPI.md#GetParentDataRetentionPolicies) | **Get** /api/v2/dataRetentionPolicies/organizations/{organizationId}/parent | 
[**SetDataRetentionPolicies**](DataRetentionPoliciesAPI.md#SetDataRetentionPolicies) | **Put** /api/v2/dataRetentionPolicies/organizations/{organizationId} | 



## GetDataRetentionPolicies

> ApiDataRetentionPoliciesDTO GetDataRetentionPolicies(ctx, organizationId).Execute()





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
	organizationId := "organizationId_example" // string | The organizationId assigned by IQ Server. Use the organization REST API to retrieve the organizationId.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.DataRetentionPoliciesAPI.GetDataRetentionPolicies(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataRetentionPoliciesAPI.GetDataRetentionPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDataRetentionPolicies`: ApiDataRetentionPoliciesDTO
	fmt.Fprintf(os.Stdout, "Response from `DataRetentionPoliciesAPI.GetDataRetentionPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The organizationId assigned by IQ Server. Use the organization REST API to retrieve the organizationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataRetentionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiDataRetentionPoliciesDTO**](ApiDataRetentionPoliciesDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParentDataRetentionPolicies

> ApiDataRetentionPoliciesDTO GetParentDataRetentionPolicies(ctx, organizationId).Execute()





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
	organizationId := "organizationId_example" // string | The organizationId assigned by IQ Server. Use the organization REST API to retrieve the parent organizationId

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.DataRetentionPoliciesAPI.GetParentDataRetentionPolicies(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataRetentionPoliciesAPI.GetParentDataRetentionPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParentDataRetentionPolicies`: ApiDataRetentionPoliciesDTO
	fmt.Fprintf(os.Stdout, "Response from `DataRetentionPoliciesAPI.GetParentDataRetentionPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The organizationId assigned by IQ Server. Use the organization REST API to retrieve the parent organizationId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParentDataRetentionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiDataRetentionPoliciesDTO**](ApiDataRetentionPoliciesDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDataRetentionPolicies

> SetDataRetentionPolicies(ctx, organizationId).ApiDataRetentionPoliciesDTO(apiDataRetentionPoliciesDTO).Execute()





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
	organizationId := "organizationId_example" // string | The organizationId for the organization you want to set the data retention policy. Use the organization REST API to retrieve the organizationId.
	apiDataRetentionPoliciesDTO := *sonatypeiq.NewApiDataRetentionPoliciesDTO() // ApiDataRetentionPoliciesDTO | The request JSON should include the retention policy settings for both application reports and success metrics.  Policy settings for application reports can be specified for each stage of development represented in the example below by additionalProp1.  Example values for additionalProp1 are develop, build, stage-release, release, operate & continuous monitoring. For application reports created during continuous monitoring use the key continuous-monitoring instead of the stage name. <ul><li>inheritPolicy IS a boolean flag indicating whether the policy is inherited from a parent organization.</li><li>enablePurging IS a boolean flag indicating enabled or disabled status for automatic purging. </li><li>maxCount IS the maximum no. of reports to retain.</li><li>maxAge IS the maximum age that a report is allowed to reach before it is purged. Possible values are days, weeks, months, years.</li></ul>

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.DataRetentionPoliciesAPI.SetDataRetentionPolicies(context.Background(), organizationId).ApiDataRetentionPoliciesDTO(apiDataRetentionPoliciesDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DataRetentionPoliciesAPI.SetDataRetentionPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The organizationId for the organization you want to set the data retention policy. Use the organization REST API to retrieve the organizationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDataRetentionPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiDataRetentionPoliciesDTO** | [**ApiDataRetentionPoliciesDTO**](ApiDataRetentionPoliciesDTO.md) | The request JSON should include the retention policy settings for both application reports and success metrics.  Policy settings for application reports can be specified for each stage of development represented in the example below by additionalProp1.  Example values for additionalProp1 are develop, build, stage-release, release, operate &amp; continuous monitoring. For application reports created during continuous monitoring use the key continuous-monitoring instead of the stage name. &lt;ul&gt;&lt;li&gt;inheritPolicy IS a boolean flag indicating whether the policy is inherited from a parent organization.&lt;/li&gt;&lt;li&gt;enablePurging IS a boolean flag indicating enabled or disabled status for automatic purging. &lt;/li&gt;&lt;li&gt;maxCount IS the maximum no. of reports to retain.&lt;/li&gt;&lt;li&gt;maxAge IS the maximum age that a report is allowed to reach before it is purged. Possible values are days, weeks, months, years.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

