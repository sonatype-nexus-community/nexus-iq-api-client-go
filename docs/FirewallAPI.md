# \FirewallAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddProprietaryComponentNames**](FirewallAPI.md#AddProprietaryComponentNames) | **Post** /api/v2/firewall/namespace_confusion/{format} | 
[**AddRepositoryManager**](FirewallAPI.md#AddRepositoryManager) | **Post** /api/v2/firewall/repositoryManagers | 
[**AddWaiver**](FirewallAPI.md#AddWaiver) | **Post** /api/v2/firewall/container-image/{containerImageId}/policyWaiver | 
[**ConfigureRepositories**](FirewallAPI.md#ConfigureRepositories) | **Post** /api/v2/firewall/repositories/configuration/{repositoryManagerId} | 
[**DeleteContainerImagePolicyWaiver**](FirewallAPI.md#DeleteContainerImagePolicyWaiver) | **Delete** /api/v2/firewall/container-image/{containerImageId}/policyWaiver | 
[**DeleteRepositoryManager**](FirewallAPI.md#DeleteRepositoryManager) | **Delete** /api/v2/firewall/repositoryManagers/{repositoryManagerId} | 
[**EvaluateComponents1**](FirewallAPI.md#EvaluateComponents1) | **Post** /api/v2/firewall/components/{repositoryManagerId}/{repositoryId}/evaluate | 
[**EvaluateMalware**](FirewallAPI.md#EvaluateMalware) | **Post** /api/v2/malware-defense/evaluate | 
[**GetConfiguredRepositories**](FirewallAPI.md#GetConfiguredRepositories) | **Get** /api/v2/firewall/repositories/configuration/{repositoryManagerId} | 
[**GetContainerImagesInQuarantine**](FirewallAPI.md#GetContainerImagesInQuarantine) | **Get** /api/v2/firewall/container-image/policyViolations/quarantined | 
[**GetFirewallAutoUnquarantineConfig**](FirewallAPI.md#GetFirewallAutoUnquarantineConfig) | **Get** /api/v2/firewall/releaseQuarantine/configuration | 
[**GetFirewallMetrics**](FirewallAPI.md#GetFirewallMetrics) | **Get** /api/v2/firewall/metrics/embedded | 
[**GetFirewallUnquarantineSummary**](FirewallAPI.md#GetFirewallUnquarantineSummary) | **Get** /api/v2/firewall/releaseQuarantine/summary | 
[**GetMalwareDefenseMetrics**](FirewallAPI.md#GetMalwareDefenseMetrics) | **Get** /api/v2/malware-defense/metrics | 
[**GetQuarantineList**](FirewallAPI.md#GetQuarantineList) | **Get** /api/v2/firewall/components/quarantined | 
[**GetQuarantineSummary**](FirewallAPI.md#GetQuarantineSummary) | **Get** /api/v2/firewall/quarantine/summary | 
[**GetQuarantinedComponentViewAnonymousAccess**](FirewallAPI.md#GetQuarantinedComponentViewAnonymousAccess) | **Get** /api/v2/firewall/quarantinedComponentView/configuration/anonymousAccess | 
[**GetRepositoryContainer**](FirewallAPI.md#GetRepositoryContainer) | **Get** /api/v2/firewall/repositoryContainer | 
[**GetRepositoryManager**](FirewallAPI.md#GetRepositoryManager) | **Get** /api/v2/firewall/repositoryManagers/{repositoryManagerId} | 
[**GetRepositoryManagers**](FirewallAPI.md#GetRepositoryManagers) | **Get** /api/v2/firewall/repositoryManagers | 
[**GetUnquarantineList**](FirewallAPI.md#GetUnquarantineList) | **Get** /api/v2/firewall/components/autoReleasedFromQuarantine | 
[**GetWaivers**](FirewallAPI.md#GetWaivers) | **Get** /api/v2/firewall/container-image/policyWaiver | 
[**RemoveProprietaryComponentNames**](FirewallAPI.md#RemoveProprietaryComponentNames) | **Delete** /api/v2/firewall/namespace_confusion/{format} | 
[**SetFirewallAutoUnquarantineConfig**](FirewallAPI.md#SetFirewallAutoUnquarantineConfig) | **Put** /api/v2/firewall/releaseQuarantine/configuration | 
[**SetQuarantinedComponentViewAnonymousAccess**](FirewallAPI.md#SetQuarantinedComponentViewAnonymousAccess) | **Put** /api/v2/firewall/quarantinedComponentView/configuration/anonymousAccess/{enabled} | 



## AddProprietaryComponentNames

> AddProprietaryComponentNames(ctx, format).RequestBody(requestBody).Execute()





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
	format := "maven" // string | Format for which the proprietary namespaces are being added.
	requestBody := []string{"Property_example"} // []string | List of namespaces to register as proprietary for this format.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.FirewallAPI.AddProprietaryComponentNames(context.Background(), format).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.AddProprietaryComponentNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** | Format for which the proprietary namespaces are being added. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProprietaryComponentNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** | List of namespaces to register as proprietary for this format. | 

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


## AddRepositoryManager

> ApiRepositoryManagerDTO AddRepositoryManager(ctx).ApiRepositoryManagerDTO(apiRepositoryManagerDTO).Execute()





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
	apiRepositoryManagerDTO := *sonatypeiq.NewApiRepositoryManagerDTO() // ApiRepositoryManagerDTO | Enter values for the new repository manager.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.AddRepositoryManager(context.Background()).ApiRepositoryManagerDTO(apiRepositoryManagerDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.AddRepositoryManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRepositoryManager`: ApiRepositoryManagerDTO
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.AddRepositoryManager`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRepositoryManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiRepositoryManagerDTO** | [**ApiRepositoryManagerDTO**](ApiRepositoryManagerDTO.md) | Enter values for the new repository manager. | 

### Return type

[**ApiRepositoryManagerDTO**](ApiRepositoryManagerDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWaiver

> AddWaiver(ctx, containerImageId).ApiContainerImageWaiverDTO(apiContainerImageWaiverDTO).Execute()





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
	containerImageId := "containerImageId_example" // string | Enter the container image id.
	apiContainerImageWaiverDTO := *sonatypeiq.NewApiContainerImageWaiverDTO() // ApiContainerImageWaiverDTO | The request JSON can include the fields<ol><li>expiryTime (default null): Sets the datetime when the waiver expires.</li><li>waiverReasonId (default null): Sets the specific reason chosen for the waiver.</li><li>comment (default null): Further explanation about the waiver.</li></ol> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.FirewallAPI.AddWaiver(context.Background(), containerImageId).ApiContainerImageWaiverDTO(apiContainerImageWaiverDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.AddWaiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerImageId** | **string** | Enter the container image id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWaiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiContainerImageWaiverDTO** | [**ApiContainerImageWaiverDTO**](ApiContainerImageWaiverDTO.md) | The request JSON can include the fields&lt;ol&gt;&lt;li&gt;expiryTime (default null): Sets the datetime when the waiver expires.&lt;/li&gt;&lt;li&gt;waiverReasonId (default null): Sets the specific reason chosen for the waiver.&lt;/li&gt;&lt;li&gt;comment (default null): Further explanation about the waiver.&lt;/li&gt;&lt;/ol&gt; | 

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


## ConfigureRepositories

> ConfigureRepositories(ctx, repositoryManagerId).ApiRepositoryListDTO(apiRepositoryListDTO).Execute()





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
	repositoryManagerId := "repositoryManagerId_example" // string | Enter the repository manager ID.
	apiRepositoryListDTO := *sonatypeiq.NewApiRepositoryListDTO() // ApiRepositoryListDTO | Enter values for the repository configuration properties to be updated.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.FirewallAPI.ConfigureRepositories(context.Background(), repositoryManagerId).ApiRepositoryListDTO(apiRepositoryListDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.ConfigureRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerId** | **string** | Enter the repository manager ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRepositoryListDTO** | [**ApiRepositoryListDTO**](ApiRepositoryListDTO.md) | Enter values for the repository configuration properties to be updated. | 

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


## DeleteContainerImagePolicyWaiver

> DeleteContainerImagePolicyWaiver(ctx, containerImageId).Execute()





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
	containerImageId := "containerImageId_example" // string | Enter the container id.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.FirewallAPI.DeleteContainerImagePolicyWaiver(context.Background(), containerImageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.DeleteContainerImagePolicyWaiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerImageId** | **string** | Enter the container id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerImagePolicyWaiverRequest struct via the builder pattern


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


## DeleteRepositoryManager

> DeleteRepositoryManager(ctx, repositoryManagerId).Execute()





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
	repositoryManagerId := "repositoryManagerId_example" // string | Enter the repository manager ID.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.FirewallAPI.DeleteRepositoryManager(context.Background(), repositoryManagerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.DeleteRepositoryManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerId** | **string** | Enter the repository manager ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryManagerRequest struct via the builder pattern


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


## EvaluateComponents1

> ApiRepositoryComponentEvaluationResultList EvaluateComponents1(ctx, repositoryManagerId, repositoryId).ApiRepositoryComponentEvaluationRequestList(apiRepositoryComponentEvaluationRequestList).Execute()





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
	repositoryManagerId := "repositoryManagerId_example" // string | Enter the repository manager ID.
	repositoryId := "repositoryId_example" // string | Enter the repository ID.
	apiRepositoryComponentEvaluationRequestList := *sonatypeiq.NewApiRepositoryComponentEvaluationRequestList() // ApiRepositoryComponentEvaluationRequestList | Provide the array of the component identifiers to be evaluated, using the component hash and the (packageUrl or pathname). A maximum of 100 components can be evaluated in one request.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.EvaluateComponents1(context.Background(), repositoryManagerId, repositoryId).ApiRepositoryComponentEvaluationRequestList(apiRepositoryComponentEvaluationRequestList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.EvaluateComponents1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvaluateComponents1`: ApiRepositoryComponentEvaluationResultList
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.EvaluateComponents1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerId** | **string** | Enter the repository manager ID. | 
**repositoryId** | **string** | Enter the repository ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateComponents1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiRepositoryComponentEvaluationRequestList** | [**ApiRepositoryComponentEvaluationRequestList**](ApiRepositoryComponentEvaluationRequestList.md) | Provide the array of the component identifiers to be evaluated, using the component hash and the (packageUrl or pathname). A maximum of 100 components can be evaluated in one request. | 

### Return type

[**ApiRepositoryComponentEvaluationResultList**](ApiRepositoryComponentEvaluationResultList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvaluateMalware

> MalwareDefenseResponseList EvaluateMalware(ctx).ApiMalwareComponentEvaluationRequestList(apiMalwareComponentEvaluationRequestList).Execute()





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
	apiMalwareComponentEvaluationRequestList := *sonatypeiq.NewApiMalwareComponentEvaluationRequestList() // ApiMalwareComponentEvaluationRequestList |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.EvaluateMalware(context.Background()).ApiMalwareComponentEvaluationRequestList(apiMalwareComponentEvaluationRequestList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.EvaluateMalware``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvaluateMalware`: MalwareDefenseResponseList
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.EvaluateMalware`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateMalwareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiMalwareComponentEvaluationRequestList** | [**ApiMalwareComponentEvaluationRequestList**](ApiMalwareComponentEvaluationRequestList.md) |  | 

### Return type

[**MalwareDefenseResponseList**](MalwareDefenseResponseList.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguredRepositories

> ApiRepositoryListDTO GetConfiguredRepositories(ctx, repositoryManagerId).SinceUtcTimestamp(sinceUtcTimestamp).Execute()





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
	repositoryManagerId := "repositoryManagerId_example" // string | Enter the repository manager ID.
	sinceUtcTimestamp := int64(789) // int64 | Enter the epoch time in milliseconds when the repository was last updated. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.GetConfiguredRepositories(context.Background(), repositoryManagerId).SinceUtcTimestamp(sinceUtcTimestamp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetConfiguredRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguredRepositories`: ApiRepositoryListDTO
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetConfiguredRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerId** | **string** | Enter the repository manager ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguredRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sinceUtcTimestamp** | **int64** | Enter the epoch time in milliseconds when the repository was last updated. | 

### Return type

[**ApiRepositoryListDTO**](ApiRepositoryListDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerImagesInQuarantine

> ContainerImageInQuarantineDataResult GetContainerImagesInQuarantine(ctx).Page(page).PageSize(pageSize).Execute()





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
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.GetContainerImagesInQuarantine(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetContainerImagesInQuarantine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContainerImagesInQuarantine`: ContainerImageInQuarantineDataResult
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetContainerImagesInQuarantine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerImagesInQuarantineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**ContainerImageInQuarantineDataResult**](ContainerImageInQuarantineDataResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallAutoUnquarantineConfig

> []ApiFirewallReleaseQuarantineConfigDTO GetFirewallAutoUnquarantineConfig(ctx).Execute()





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
	resp, r, err := apiClient.FirewallAPI.GetFirewallAutoUnquarantineConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetFirewallAutoUnquarantineConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewallAutoUnquarantineConfig`: []ApiFirewallReleaseQuarantineConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetFirewallAutoUnquarantineConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallAutoUnquarantineConfigRequest struct via the builder pattern


### Return type

[**[]ApiFirewallReleaseQuarantineConfigDTO**](ApiFirewallReleaseQuarantineConfigDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallMetrics

> map[string]ApiFirewallMetricsResultDTO GetFirewallMetrics(ctx).Execute()





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
	resp, r, err := apiClient.FirewallAPI.GetFirewallMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetFirewallMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewallMetrics`: map[string]ApiFirewallMetricsResultDTO
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetFirewallMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallMetricsRequest struct via the builder pattern


### Return type

[**map[string]ApiFirewallMetricsResultDTO**](ApiFirewallMetricsResultDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallUnquarantineSummary

> ApiFirewallReleaseQuarantineSummaryDTO GetFirewallUnquarantineSummary(ctx).Execute()





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
	resp, r, err := apiClient.FirewallAPI.GetFirewallUnquarantineSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetFirewallUnquarantineSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewallUnquarantineSummary`: ApiFirewallReleaseQuarantineSummaryDTO
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetFirewallUnquarantineSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallUnquarantineSummaryRequest struct via the builder pattern


### Return type

[**ApiFirewallReleaseQuarantineSummaryDTO**](ApiFirewallReleaseQuarantineSummaryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMalwareDefenseMetrics

> map[string]int64 GetMalwareDefenseMetrics(ctx).Execute()





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
	resp, r, err := apiClient.FirewallAPI.GetMalwareDefenseMetrics(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetMalwareDefenseMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMalwareDefenseMetrics`: map[string]int64
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetMalwareDefenseMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMalwareDefenseMetricsRequest struct via the builder pattern


### Return type

**map[string]int64**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuarantineList

> ApiFirewallQuarantinedComponentDtoResult GetQuarantineList(ctx).Page(page).PageSize(pageSize).PolicyId(policyId).ComponentName(componentName).RepositoryPublicId(repositoryPublicId).QuarantineTime(quarantineTime).SortBy(sortBy).Asc(asc).Execute()





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
	page := int32(56) // int32 | Enter the starting page number for the response. (optional) (default to 1)
	pageSize := int32(56) // int32 | Enter the page size for the response. (optional) (default to 10)
	policyId := []string{"Inner_example"} // []string | Enter the list of policy IDs causing the quarantine. (optional)
	componentName := "componentName_example" // string | Enter the component name. (optional)
	repositoryPublicId := "repositoryPublicId_example" // string | Enter the repository public ID of the quarantined component. (optional)
	quarantineTime := int32(56) // int32 | Enter the quarantine time of the component. (optional)
	sortBy := "sortBy_example" // string | Enter `quarantineTime` to sort the results by quarantine time. (optional)
	asc := true // bool | Select the sort order. (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.GetQuarantineList(context.Background()).Page(page).PageSize(pageSize).PolicyId(policyId).ComponentName(componentName).RepositoryPublicId(repositoryPublicId).QuarantineTime(quarantineTime).SortBy(sortBy).Asc(asc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetQuarantineList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuarantineList`: ApiFirewallQuarantinedComponentDtoResult
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetQuarantineList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuarantineListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Enter the starting page number for the response. | [default to 1]
 **pageSize** | **int32** | Enter the page size for the response. | [default to 10]
 **policyId** | **[]string** | Enter the list of policy IDs causing the quarantine. | 
 **componentName** | **string** | Enter the component name. | 
 **repositoryPublicId** | **string** | Enter the repository public ID of the quarantined component. | 
 **quarantineTime** | **int32** | Enter the quarantine time of the component. | 
 **sortBy** | **string** | Enter &#x60;quarantineTime&#x60; to sort the results by quarantine time. | 
 **asc** | **bool** | Select the sort order. | [default to false]

### Return type

[**ApiFirewallQuarantinedComponentDtoResult**](ApiFirewallQuarantinedComponentDtoResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuarantineSummary

> ApiFirewallQuarantineSummaryDTO GetQuarantineSummary(ctx).Execute()





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
	resp, r, err := apiClient.FirewallAPI.GetQuarantineSummary(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetQuarantineSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuarantineSummary`: ApiFirewallQuarantineSummaryDTO
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetQuarantineSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuarantineSummaryRequest struct via the builder pattern


### Return type

[**ApiFirewallQuarantineSummaryDTO**](ApiFirewallQuarantineSummaryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuarantinedComponentViewAnonymousAccess

> bool GetQuarantinedComponentViewAnonymousAccess(ctx).Execute()





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
	resp, r, err := apiClient.FirewallAPI.GetQuarantinedComponentViewAnonymousAccess(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetQuarantinedComponentViewAnonymousAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuarantinedComponentViewAnonymousAccess`: bool
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetQuarantinedComponentViewAnonymousAccess`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuarantinedComponentViewAnonymousAccessRequest struct via the builder pattern


### Return type

**bool**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryContainer

> ApiRepositoryContainerDTO GetRepositoryContainer(ctx).Execute()





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
	resp, r, err := apiClient.FirewallAPI.GetRepositoryContainer(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetRepositoryContainer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryContainer`: ApiRepositoryContainerDTO
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetRepositoryContainer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryContainerRequest struct via the builder pattern


### Return type

[**ApiRepositoryContainerDTO**](ApiRepositoryContainerDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryManager

> ApiRepositoryManagerDTO GetRepositoryManager(ctx, repositoryManagerId).Execute()





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
	repositoryManagerId := "repositoryManagerId_example" // string | Enter the repository manager ID.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.GetRepositoryManager(context.Background(), repositoryManagerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetRepositoryManager``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryManager`: ApiRepositoryManagerDTO
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetRepositoryManager`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerId** | **string** | Enter the repository manager ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryManagerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiRepositoryManagerDTO**](ApiRepositoryManagerDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositoryManagers

> ApiRepositoryManagerListDTO GetRepositoryManagers(ctx).Execute()





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
	resp, r, err := apiClient.FirewallAPI.GetRepositoryManagers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetRepositoryManagers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositoryManagers`: ApiRepositoryManagerListDTO
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetRepositoryManagers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryManagersRequest struct via the builder pattern


### Return type

[**ApiRepositoryManagerListDTO**](ApiRepositoryManagerListDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnquarantineList

> ApiFirewallComponentDTOResult GetUnquarantineList(ctx).Page(page).PageSize(pageSize).PolicyId(policyId).ComponentName(componentName).SortBy(sortBy).Asc(asc).Execute()





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
	page := int32(56) // int32 | Enter the page number. (optional) (default to 1)
	pageSize := int32(56) // int32 | Enter the number of results to be returned for a page. (optional) (default to 10)
	policyId := "policyId_example" // string | Enter the `policyId`. When provided, the results will include the components that have a policy violation for the policyId. (optional)
	componentName := "componentName_example" // string | Enter the component name. When provided, the results will include components with display names (case-insensitive) that match the given name. (optional)
	sortBy := "sortBy_example" // string | Enter the sort criteria `releaseQuarantineTime` or `quarantineTime`. (optional)
	asc := true // bool | Select `true` to set the sort order to ascending. (optional) (default to true)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.GetUnquarantineList(context.Background()).Page(page).PageSize(pageSize).PolicyId(policyId).ComponentName(componentName).SortBy(sortBy).Asc(asc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetUnquarantineList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnquarantineList`: ApiFirewallComponentDTOResult
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetUnquarantineList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUnquarantineListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Enter the page number. | [default to 1]
 **pageSize** | **int32** | Enter the number of results to be returned for a page. | [default to 10]
 **policyId** | **string** | Enter the &#x60;policyId&#x60;. When provided, the results will include the components that have a policy violation for the policyId. | 
 **componentName** | **string** | Enter the component name. When provided, the results will include components with display names (case-insensitive) that match the given name. | 
 **sortBy** | **string** | Enter the sort criteria &#x60;releaseQuarantineTime&#x60; or &#x60;quarantineTime&#x60;. | 
 **asc** | **bool** | Select &#x60;true&#x60; to set the sort order to ascending. | [default to true]

### Return type

[**ApiFirewallComponentDTOResult**](ApiFirewallComponentDTOResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWaivers

> PolicyContainerWaiverDataResult GetWaivers(ctx).Page(page).PageSize(pageSize).Execute()





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
	page := int32(56) // int32 |  (optional)
	pageSize := int32(56) // int32 |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.GetWaivers(context.Background()).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetWaivers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWaivers`: PolicyContainerWaiverDataResult
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetWaivers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWaiversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | 
 **pageSize** | **int32** |  | 

### Return type

[**PolicyContainerWaiverDataResult**](PolicyContainerWaiverDataResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProprietaryComponentNames

> RemoveProprietaryComponentNames(ctx, format).Execute()





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
	format := "maven" // string | Format for which the proprietary namespaces are being removed.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.FirewallAPI.RemoveProprietaryComponentNames(context.Background(), format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.RemoveProprietaryComponentNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** | Format for which the proprietary namespaces are being removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProprietaryComponentNamesRequest struct via the builder pattern


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


## SetFirewallAutoUnquarantineConfig

> []ApiFirewallReleaseQuarantineConfigDTO SetFirewallAutoUnquarantineConfig(ctx).ApiFirewallReleaseQuarantineConfigDTO(apiFirewallReleaseQuarantineConfigDTO).Execute()





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
	apiFirewallReleaseQuarantineConfigDTO := []sonatypeiq.ApiFirewallReleaseQuarantineConfigDTO{*sonatypeiq.NewApiFirewallReleaseQuarantineConfigDTO()} // []ApiFirewallReleaseQuarantineConfigDTO | Enter value for each repository and the required status for auto-release as `true` or `false`.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallAPI.SetFirewallAutoUnquarantineConfig(context.Background()).ApiFirewallReleaseQuarantineConfigDTO(apiFirewallReleaseQuarantineConfigDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.SetFirewallAutoUnquarantineConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFirewallAutoUnquarantineConfig`: []ApiFirewallReleaseQuarantineConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.SetFirewallAutoUnquarantineConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetFirewallAutoUnquarantineConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiFirewallReleaseQuarantineConfigDTO** | [**[]ApiFirewallReleaseQuarantineConfigDTO**](ApiFirewallReleaseQuarantineConfigDTO.md) | Enter value for each repository and the required status for auto-release as &#x60;true&#x60; or &#x60;false&#x60;. | 

### Return type

[**[]ApiFirewallReleaseQuarantineConfigDTO**](ApiFirewallReleaseQuarantineConfigDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetQuarantinedComponentViewAnonymousAccess

> SetQuarantinedComponentViewAnonymousAccess(ctx, enabled).Execute()





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
	enabled := true // bool | Select `true` or `false` to enable or disable anonymous access.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.FirewallAPI.SetQuarantinedComponentViewAnonymousAccess(context.Background(), enabled).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.SetQuarantinedComponentViewAnonymousAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enabled** | **bool** | Select &#x60;true&#x60; or &#x60;false&#x60; to enable or disable anonymous access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetQuarantinedComponentViewAnonymousAccessRequest struct via the builder pattern


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

