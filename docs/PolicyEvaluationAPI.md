# \PolicyEvaluationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EvaluateComponents**](PolicyEvaluationAPI.md#EvaluateComponents) | **Post** /api/v2/evaluation/applications/{applicationId} | 
[**EvaluateSourceControl**](PolicyEvaluationAPI.md#EvaluateSourceControl) | **Post** /api/v2/evaluation/applications/{applicationId}/sourceControlEvaluation | 
[**GetApplicationEvaluationStatus**](PolicyEvaluationAPI.md#GetApplicationEvaluationStatus) | **Get** /api/v2/evaluation/applications/{applicationId}/status/{statusId} | 
[**GetComponentEvaluation**](PolicyEvaluationAPI.md#GetComponentEvaluation) | **Get** /api/v2/evaluation/applications/{applicationId}/results/{resultId} | 
[**PromoteScan**](PolicyEvaluationAPI.md#PromoteScan) | **Post** /api/v2/evaluation/applications/{applicationId}/promoteScan | 



## EvaluateComponents

> ApiComponentEvaluationTicketDTOV2 EvaluateComponents(ctx, applicationId).ApiComponentEvaluationRequestDTOV2(apiComponentEvaluationRequestDTOV2).Execute()





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
	applicationId := "applicationId_example" // string | Enter the internal applicationId. Use the Applications REST API to retrieve the internal applicationId.
	apiComponentEvaluationRequestDTOV2 := *sonatypeiq.NewApiComponentEvaluationRequestDTOV2() // ApiComponentEvaluationRequestDTOV2 | The request JSON should contain component coordinates or the hash (SHA1) for each component. You can provide the packageURL instead of component information or hash. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyEvaluationAPI.EvaluateComponents(context.Background(), applicationId).ApiComponentEvaluationRequestDTOV2(apiComponentEvaluationRequestDTOV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyEvaluationAPI.EvaluateComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvaluateComponents`: ApiComponentEvaluationTicketDTOV2
	fmt.Fprintf(os.Stdout, "Response from `PolicyEvaluationAPI.EvaluateComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the internal applicationId. Use the Applications REST API to retrieve the internal applicationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiComponentEvaluationRequestDTOV2** | [**ApiComponentEvaluationRequestDTOV2**](ApiComponentEvaluationRequestDTOV2.md) | The request JSON should contain component coordinates or the hash (SHA1) for each component. You can provide the packageURL instead of component information or hash. | 

### Return type

[**ApiComponentEvaluationTicketDTOV2**](ApiComponentEvaluationTicketDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvaluateSourceControl

> ApiApplicationEvaluationStatusDTOV2 EvaluateSourceControl(ctx, applicationId).ApiSourceControlEvaluationRequestDTO(apiSourceControlEvaluationRequestDTO).Execute()





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
	applicationId := "applicationId_example" // string | Enter the internal applicationId. Use the Applications REST API to retrieve the internal applicationId.
	apiSourceControlEvaluationRequestDTO := *sonatypeiq.NewApiSourceControlEvaluationRequestDTO() // ApiSourceControlEvaluationRequestDTO | The request JSON should include the 1. branch name (name of the target branch in the source control repository, 2. stageId (recommended values are 'develop' for feature branches, and 'source' for default branches. Other stageIds that can be used are 'build', 'stage-release', 'release', 'operate' but are not recommended), 3. scanTargets (optional, specify one or more paths inside the repository. If not specified, the entire repository will be evaluated by default). Ensure that the repository paths are not relative and do not contain '../' or '..\\'. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyEvaluationAPI.EvaluateSourceControl(context.Background(), applicationId).ApiSourceControlEvaluationRequestDTO(apiSourceControlEvaluationRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyEvaluationAPI.EvaluateSourceControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvaluateSourceControl`: ApiApplicationEvaluationStatusDTOV2
	fmt.Fprintf(os.Stdout, "Response from `PolicyEvaluationAPI.EvaluateSourceControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the internal applicationId. Use the Applications REST API to retrieve the internal applicationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateSourceControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiSourceControlEvaluationRequestDTO** | [**ApiSourceControlEvaluationRequestDTO**](ApiSourceControlEvaluationRequestDTO.md) | The request JSON should include the 1. branch name (name of the target branch in the source control repository, 2. stageId (recommended values are &#39;develop&#39; for feature branches, and &#39;source&#39; for default branches. Other stageIds that can be used are &#39;build&#39;, &#39;stage-release&#39;, &#39;release&#39;, &#39;operate&#39; but are not recommended), 3. scanTargets (optional, specify one or more paths inside the repository. If not specified, the entire repository will be evaluated by default). Ensure that the repository paths are not relative and do not contain &#39;../&#39; or &#39;..\\&#39;. | 

### Return type

[**ApiApplicationEvaluationStatusDTOV2**](ApiApplicationEvaluationStatusDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationEvaluationStatus

> ApiApplicationEvaluationResultDTOV2 GetApplicationEvaluationStatus(ctx, applicationId, statusId).Execute()





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
	applicationId := "applicationId_example" // string | Enter the applicationId, for the which policy evaluation was requested.
	statusId := "statusId_example" // string | Enter the statusId value obtained as response of the POST call in step 1.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyEvaluationAPI.GetApplicationEvaluationStatus(context.Background(), applicationId, statusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyEvaluationAPI.GetApplicationEvaluationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationEvaluationStatus`: ApiApplicationEvaluationResultDTOV2
	fmt.Fprintf(os.Stdout, "Response from `PolicyEvaluationAPI.GetApplicationEvaluationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the applicationId, for the which policy evaluation was requested. | 
**statusId** | **string** | Enter the statusId value obtained as response of the POST call in step 1. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationEvaluationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiApplicationEvaluationResultDTOV2**](ApiApplicationEvaluationResultDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentEvaluation

> ApiComponentEvaluationResultDTOV2 GetComponentEvaluation(ctx, applicationId, resultId).Execute()





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
	applicationId := "applicationId_example" // string | Enter the internal applicationId (same as that sent in the POST request (step 1))
	resultId := "resultId_example" // string | Enter the resultId obtained from the POST response (step 1) used for component evaluation.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyEvaluationAPI.GetComponentEvaluation(context.Background(), applicationId, resultId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyEvaluationAPI.GetComponentEvaluation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponentEvaluation`: ApiComponentEvaluationResultDTOV2
	fmt.Fprintf(os.Stdout, "Response from `PolicyEvaluationAPI.GetComponentEvaluation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the internal applicationId (same as that sent in the POST request (step 1)) | 
**resultId** | **string** | Enter the resultId obtained from the POST response (step 1) used for component evaluation. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentEvaluationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiComponentEvaluationResultDTOV2**](ApiComponentEvaluationResultDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PromoteScan

> ApiApplicationEvaluationStatusDTOV2 PromoteScan(ctx, applicationId).ApiPromoteScanRequestDTOV2(apiPromoteScanRequestDTOV2).Execute()





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
	applicationId := "applicationId_example" // string | Enter the internal applicationId. Use the Applications REST API to retrieve the internal applicationId.
	apiPromoteScanRequestDTOV2 := *sonatypeiq.NewApiPromoteScanRequestDTOV2() // ApiPromoteScanRequestDTOV2 | You can provide either the scanId (reportId) of the previous scan OR the source stageId (possible values are 'build', 'stage-release', 'release' or 'operate'). When using the stageId, the latest scanId for the application will be used. Enter the targetStageId for the new stage you want your scan to be promoted to (possible values are 'build', 'stage-release', 'release' or 'operate'). Using the same value for source and target stage will resubmit the latest scan report. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyEvaluationAPI.PromoteScan(context.Background(), applicationId).ApiPromoteScanRequestDTOV2(apiPromoteScanRequestDTOV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyEvaluationAPI.PromoteScan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromoteScan`: ApiApplicationEvaluationStatusDTOV2
	fmt.Fprintf(os.Stdout, "Response from `PolicyEvaluationAPI.PromoteScan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the internal applicationId. Use the Applications REST API to retrieve the internal applicationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromoteScanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiPromoteScanRequestDTOV2** | [**ApiPromoteScanRequestDTOV2**](ApiPromoteScanRequestDTOV2.md) | You can provide either the scanId (reportId) of the previous scan OR the source stageId (possible values are &#39;build&#39;, &#39;stage-release&#39;, &#39;release&#39; or &#39;operate&#39;). When using the stageId, the latest scanId for the application will be used. Enter the targetStageId for the new stage you want your scan to be promoted to (possible values are &#39;build&#39;, &#39;stage-release&#39;, &#39;release&#39; or &#39;operate&#39;). Using the same value for source and target stage will resubmit the latest scan report. | 

### Return type

[**ApiApplicationEvaluationStatusDTOV2**](ApiApplicationEvaluationStatusDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

