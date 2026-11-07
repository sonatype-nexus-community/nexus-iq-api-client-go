# \HostedRepositoryComponentReportDataAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDependencyTree**](HostedRepositoryComponentReportDataAPI.md#GetDependencyTree) | **Get** /api/v2/hostedRepositoryComponent/{hrcId}/reports/{scanId}/dependencyTree | 
[**GetPolicyViolations**](HostedRepositoryComponentReportDataAPI.md#GetPolicyViolations) | **Get** /api/v2/hostedRepositoryComponent/{hrcId}/reports/{scanId}/policy | 
[**GetRawData**](HostedRepositoryComponentReportDataAPI.md#GetRawData) | **Get** /api/v2/hostedRepositoryComponent/{hrcId}/reports/{scanId}/raw | 



## GetDependencyTree

> ApiDependencyTreeResponseDTO GetDependencyTree(ctx, hrcId, scanId).Execute()





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
	hrcId := "hrcId_example" // string | Enter the hosted repository component ID (UUID).
	scanId := "scanId_example" // string |  Enter the reportId (scanId) created at the time of evaluation.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedRepositoryComponentReportDataAPI.GetDependencyTree(context.Background(), hrcId, scanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryComponentReportDataAPI.GetDependencyTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDependencyTree`: ApiDependencyTreeResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `HostedRepositoryComponentReportDataAPI.GetDependencyTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hrcId** | **string** | Enter the hosted repository component ID (UUID). | 
**scanId** | **string** |  Enter the reportId (scanId) created at the time of evaluation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDependencyTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiDependencyTreeResponseDTO**](ApiDependencyTreeResponseDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyViolations

> ApiReportPolicyDataDTOV2 GetPolicyViolations(ctx, hrcId, scanId).IncludeViolationTimes(includeViolationTimes).Page(page).PageSize(pageSize).Execute()





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
	hrcId := "hrcId_example" // string | Enter the hosted repository component ID (UUID).
	scanId := "scanId_example" // string | Enter the reportId (scanId) created at the time of evaluation.
	includeViolationTimes := true // bool | Set to true to include policy violation times (open, legacy, waived, fixed) in the response if set. (optional) (default to false)
	page := int32(56) // int32 | Page number (1-indexed). Must be provided together with 'pageSize'. When omitted, all components are returned. (optional)
	pageSize := int32(56) // int32 | Number of components per page (1-500). Must be provided together with 'page'. When omitted, all components are returned. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedRepositoryComponentReportDataAPI.GetPolicyViolations(context.Background(), hrcId, scanId).IncludeViolationTimes(includeViolationTimes).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryComponentReportDataAPI.GetPolicyViolations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyViolations`: ApiReportPolicyDataDTOV2
	fmt.Fprintf(os.Stdout, "Response from `HostedRepositoryComponentReportDataAPI.GetPolicyViolations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hrcId** | **string** | Enter the hosted repository component ID (UUID). | 
**scanId** | **string** | Enter the reportId (scanId) created at the time of evaluation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeViolationTimes** | **bool** | Set to true to include policy violation times (open, legacy, waived, fixed) in the response if set. | [default to false]
 **page** | **int32** | Page number (1-indexed). Must be provided together with &#39;pageSize&#39;. When omitted, all components are returned. | 
 **pageSize** | **int32** | Number of components per page (1-500). Must be provided together with &#39;page&#39;. When omitted, all components are returned. | 

### Return type

[**ApiReportPolicyDataDTOV2**](ApiReportPolicyDataDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawData

> ApiReportRawDataDTOV2 GetRawData(ctx, hrcId, scanId).IncludeCustomSecurityVulnerabilityData(includeCustomSecurityVulnerabilityData).Execute()





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
	hrcId := "hrcId_example" // string | Enter the hosted repository component ID (UUID).
	scanId := "scanId_example" // string | Enter the reportId (scanId) created at the time of evaluating the hosted repository component.
	includeCustomSecurityVulnerabilityData := true // bool | Set to true to include security vulnerability custom data (remediation, cweId, cvssVector, cvssSeverity) for each securityIssue. Defaults to false. (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.HostedRepositoryComponentReportDataAPI.GetRawData(context.Background(), hrcId, scanId).IncludeCustomSecurityVulnerabilityData(includeCustomSecurityVulnerabilityData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostedRepositoryComponentReportDataAPI.GetRawData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawData`: ApiReportRawDataDTOV2
	fmt.Fprintf(os.Stdout, "Response from `HostedRepositoryComponentReportDataAPI.GetRawData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hrcId** | **string** | Enter the hosted repository component ID (UUID). | 
**scanId** | **string** | Enter the reportId (scanId) created at the time of evaluating the hosted repository component. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeCustomSecurityVulnerabilityData** | **bool** | Set to true to include security vulnerability custom data (remediation, cweId, cvssVector, cvssSeverity) for each securityIssue. Defaults to false. | [default to false]

### Return type

[**ApiReportRawDataDTOV2**](ApiReportRawDataDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

