# \ApplicationReportDataAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetData**](ApplicationReportDataAPI.md#GetData) | **Get** /api/v2/applications/{applicationPublicId}/reports/{scanId} | 
[**GetDependencyTree**](ApplicationReportDataAPI.md#GetDependencyTree) | **Get** /api/v2/applications/{applicationPublicId}/reports/{scanId}/dependencyTree | 
[**GetPolicyViolationDiff**](ApplicationReportDataAPI.md#GetPolicyViolationDiff) | **Get** /api/v2/applications/{applicationPublicId}/reports/policyViolations/diff | 
[**GetPolicyViolations1**](ApplicationReportDataAPI.md#GetPolicyViolations1) | **Get** /api/v2/applications/{applicationPublicId}/reports/{scanId}/policy | 
[**GetRawData**](ApplicationReportDataAPI.md#GetRawData) | **Get** /api/v2/applications/{applicationPublicId}/reports/{scanId}/raw | 



## GetData

> GetData(ctx, applicationPublicId, scanId).Execute()





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
	applicationPublicId := "applicationPublicId_example" // string | Enter the applicationPublicId for the evaluated application.
	scanId := "scanId_example" // string | Enter the scanId (reportId) of the application report created after the evaluation. 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ApplicationReportDataAPI.GetData(context.Background(), applicationPublicId, scanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationReportDataAPI.GetData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** | Enter the applicationPublicId for the evaluated application. | 
**scanId** | **string** | Enter the scanId (reportId) of the application report created after the evaluation.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataRequest struct via the builder pattern


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


## GetDependencyTree

> ApiDependencyTreeResponseDTO GetDependencyTree(ctx, applicationPublicId, scanId).Execute()





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
	applicationPublicId := "applicationPublicId_example" // string | Enter the applicationPublicId created at the time of creating the application.
	scanId := "scanId_example" // string |  Enter the reportId (scanId) created at the time of evaluating the application.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationReportDataAPI.GetDependencyTree(context.Background(), applicationPublicId, scanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationReportDataAPI.GetDependencyTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDependencyTree`: ApiDependencyTreeResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationReportDataAPI.GetDependencyTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** | Enter the applicationPublicId created at the time of creating the application. | 
**scanId** | **string** |  Enter the reportId (scanId) created at the time of evaluating the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDependencyTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiDependencyTreeResponseDTO**](ApiDependencyTreeResponseDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyViolationDiff

> ApiPolicyViolationDiffDTO GetPolicyViolationDiff(ctx, applicationPublicId).FromCommit(fromCommit).ToCommit(toCommit).FromPolicyEvaluationId(fromPolicyEvaluationId).ToPolicyEvaluationId(toPolicyEvaluationId).IncludeViolationTimes(includeViolationTimes).Execute()





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
	applicationPublicId := "applicationPublicId_example" // string | Enter the applicationPublicId, created at the time of creating the application
	fromCommit := "fromCommit_example" // string | Enter the commit hash linked to the earlier policy evaluation.
	toCommit := "toCommit_example" // string | Enter the commit hash linked to the other (later) policy evaluation to compare.
	fromPolicyEvaluationId := "fromPolicyEvaluationId_example" // string | Enter the policy evaluation Id linked to the earlier policy evaluation to compare (optional)
	toPolicyEvaluationId := "toPolicyEvaluationId_example" // string | Enter the policy evaluation Id linked to the other (later) policy evaluation to compare (optional)
	includeViolationTimes := true // bool | Set to true to include policy violation times (open, legacy, waived, fixed) in the response if set. (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationReportDataAPI.GetPolicyViolationDiff(context.Background(), applicationPublicId).FromCommit(fromCommit).ToCommit(toCommit).FromPolicyEvaluationId(fromPolicyEvaluationId).ToPolicyEvaluationId(toPolicyEvaluationId).IncludeViolationTimes(includeViolationTimes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationReportDataAPI.GetPolicyViolationDiff``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyViolationDiff`: ApiPolicyViolationDiffDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationReportDataAPI.GetPolicyViolationDiff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** | Enter the applicationPublicId, created at the time of creating the application | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyViolationDiffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromCommit** | **string** | Enter the commit hash linked to the earlier policy evaluation. | 
 **toCommit** | **string** | Enter the commit hash linked to the other (later) policy evaluation to compare. | 
 **fromPolicyEvaluationId** | **string** | Enter the policy evaluation Id linked to the earlier policy evaluation to compare | 
 **toPolicyEvaluationId** | **string** | Enter the policy evaluation Id linked to the other (later) policy evaluation to compare | 
 **includeViolationTimes** | **bool** | Set to true to include policy violation times (open, legacy, waived, fixed) in the response if set. | [default to false]

### Return type

[**ApiPolicyViolationDiffDTO**](ApiPolicyViolationDiffDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyViolations1

> ApiReportPolicyDataDTOV2 GetPolicyViolations1(ctx, applicationPublicId, scanId).IncludeViolationTimes(includeViolationTimes).Execute()





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
	applicationPublicId := "applicationPublicId_example" // string | Enter the applicationPublicId created at the time of creating the application.
	scanId := "scanId_example" // string | Enter the reportId (scanId) created at the time of evaluating the application.
	includeViolationTimes := true // bool | Set to true to include policy violation times (open, legacy, waived, fixed) in the response if set. (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationReportDataAPI.GetPolicyViolations1(context.Background(), applicationPublicId, scanId).IncludeViolationTimes(includeViolationTimes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationReportDataAPI.GetPolicyViolations1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyViolations1`: ApiReportPolicyDataDTOV2
	fmt.Fprintf(os.Stdout, "Response from `ApplicationReportDataAPI.GetPolicyViolations1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** | Enter the applicationPublicId created at the time of creating the application. | 
**scanId** | **string** | Enter the reportId (scanId) created at the time of evaluating the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyViolations1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeViolationTimes** | **bool** | Set to true to include policy violation times (open, legacy, waived, fixed) in the response if set. | [default to false]

### Return type

[**ApiReportPolicyDataDTOV2**](ApiReportPolicyDataDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRawData

> ApiReportRawDataDTOV2 GetRawData(ctx, applicationPublicId, scanId).Execute()





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
	applicationPublicId := "applicationPublicId_example" // string | Enter the applicationPublicId (assigned at the time of creating a new application.) 
	scanId := "scanId_example" // string | Enter the reportId (scanId) created at the time of evaluating the application. application.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationReportDataAPI.GetRawData(context.Background(), applicationPublicId, scanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationReportDataAPI.GetRawData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRawData`: ApiReportRawDataDTOV2
	fmt.Fprintf(os.Stdout, "Response from `ApplicationReportDataAPI.GetRawData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** | Enter the applicationPublicId (assigned at the time of creating a new application.)  | 
**scanId** | **string** | Enter the reportId (scanId) created at the time of evaluating the application. application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRawDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiReportRawDataDTOV2**](ApiReportRawDataDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

