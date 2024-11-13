# \FirewallAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRepositoryManager**](FirewallAPI.md#AddRepositoryManager) | **Post** /api/v2/firewall/repositoryManagers | 
[**ConfigureRepositories**](FirewallAPI.md#ConfigureRepositories) | **Post** /api/v2/firewall/repositories/configuration/{repositoryManagerId} | 
[**DeleteRepositoryManager**](FirewallAPI.md#DeleteRepositoryManager) | **Delete** /api/v2/firewall/repositoryManagers/{repositoryManagerId} | 
[**EvaluateComponents1**](FirewallAPI.md#EvaluateComponents1) | **Post** /api/v2/firewall/components/{repositoryManagerId}/{repositoryId}/evaluate | 
[**GetConfiguredRepositories**](FirewallAPI.md#GetConfiguredRepositories) | **Get** /api/v2/firewall/repositories/configuration/{repositoryManagerId} | 
[**GetFirewallAutoUnquarantineConfig**](FirewallAPI.md#GetFirewallAutoUnquarantineConfig) | **Get** /api/v2/firewall/releaseQuarantine/configuration | 
[**GetFirewallMetrics**](FirewallAPI.md#GetFirewallMetrics) | **Get** /api/v2/firewall/metrics/embedded | 
[**GetFirewallUnquarantineSummary**](FirewallAPI.md#GetFirewallUnquarantineSummary) | **Get** /api/v2/firewall/releaseQuarantine/summary | 
[**GetQuarantineList**](FirewallAPI.md#GetQuarantineList) | **Get** /api/v2/firewall/components/quarantined | 
[**GetQuarantineSummary**](FirewallAPI.md#GetQuarantineSummary) | **Get** /api/v2/firewall/quarantine/summary | 
[**GetQuarantinedComponentViewAnonymousAccess**](FirewallAPI.md#GetQuarantinedComponentViewAnonymousAccess) | **Get** /api/v2/firewall/quarantinedComponentView/configuration/anonymousAccess | 
[**GetRepositoryContainer**](FirewallAPI.md#GetRepositoryContainer) | **Get** /api/v2/firewall/repositoryContainer | 
[**GetRepositoryManager**](FirewallAPI.md#GetRepositoryManager) | **Get** /api/v2/firewall/repositoryManagers/{repositoryManagerId} | 
[**GetRepositoryManagers**](FirewallAPI.md#GetRepositoryManagers) | **Get** /api/v2/firewall/repositoryManagers | 
[**GetUnquarantineList**](FirewallAPI.md#GetUnquarantineList) | **Get** /api/v2/firewall/components/autoReleasedFromQuarantine | 
[**SetFirewallAutoUnquarantineConfig**](FirewallAPI.md#SetFirewallAutoUnquarantineConfig) | **Put** /api/v2/firewall/releaseQuarantine/configuration | 
[**SetQuarantinedComponentViewAnonymousAccess**](FirewallAPI.md#SetQuarantinedComponentViewAnonymousAccess) | **Put** /api/v2/firewall/quarantinedComponentView/configuration/anonymousAccess/{enabled} | 



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
	apiRepositoryManagerDTO := *sonatypeiq.NewApiRepositoryManagerDTO() // ApiRepositoryManagerDTO |  (optional)

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
 **apiRepositoryManagerDTO** | [**ApiRepositoryManagerDTO**](ApiRepositoryManagerDTO.md) |  | 

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
	repositoryManagerId := "repositoryManagerId_example" // string | 
	apiRepositoryListDTO := *sonatypeiq.NewApiRepositoryListDTO() // ApiRepositoryListDTO |  (optional)

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
**repositoryManagerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfigureRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRepositoryListDTO** | [**ApiRepositoryListDTO**](ApiRepositoryListDTO.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

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
	repositoryManagerId := "repositoryManagerId_example" // string | 

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
**repositoryManagerId** | **string** |  | 

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
- **Accept**: application/json

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
	repositoryManagerId := "repositoryManagerId_example" // string | 
	repositoryId := "repositoryId_example" // string | 
	apiRepositoryComponentEvaluationRequestList := *sonatypeiq.NewApiRepositoryComponentEvaluationRequestList() // ApiRepositoryComponentEvaluationRequestList |  (optional)

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
**repositoryManagerId** | **string** |  | 
**repositoryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateComponents1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiRepositoryComponentEvaluationRequestList** | [**ApiRepositoryComponentEvaluationRequestList**](ApiRepositoryComponentEvaluationRequestList.md) |  | 

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
	repositoryManagerId := "repositoryManagerId_example" // string | 
	sinceUtcTimestamp := int64(789) // int64 |  (optional)

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
**repositoryManagerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguredRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sinceUtcTimestamp** | **int64** |  | 

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


## GetQuarantineList

> GetQuarantineList(ctx).Page(page).PageSize(pageSize).PolicyId(policyId).ComponentName(componentName).RepositoryPublicId(repositoryPublicId).SortBy(sortBy).Asc(asc).Execute()



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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	policyId := []string{"Inner_example"} // []string |  (optional)
	componentName := "componentName_example" // string |  (optional)
	repositoryPublicId := "repositoryPublicId_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	asc := true // bool |  (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.FirewallAPI.GetQuarantineList(context.Background()).Page(page).PageSize(pageSize).PolicyId(policyId).ComponentName(componentName).RepositoryPublicId(repositoryPublicId).SortBy(sortBy).Asc(asc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetQuarantineList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuarantineListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **policyId** | **[]string** |  | 
 **componentName** | **string** |  | 
 **repositoryPublicId** | **string** |  | 
 **sortBy** | **string** |  | 
 **asc** | **bool** |  | [default to false]

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

> GetQuarantinedComponentViewAnonymousAccess(ctx).Execute()



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
	r, err := apiClient.FirewallAPI.GetQuarantinedComponentViewAnonymousAccess(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetQuarantinedComponentViewAnonymousAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuarantinedComponentViewAnonymousAccessRequest struct via the builder pattern


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
	repositoryManagerId := "repositoryManagerId_example" // string | 

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
**repositoryManagerId** | **string** |  | 

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

> GetUnquarantineList(ctx).Page(page).PageSize(pageSize).PolicyId(policyId).ComponentName(componentName).SortBy(sortBy).Asc(asc).Execute()



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
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	policyId := "policyId_example" // string |  (optional)
	componentName := "componentName_example" // string |  (optional)
	sortBy := "sortBy_example" // string |  (optional)
	asc := true // bool |  (optional) (default to true)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.FirewallAPI.GetUnquarantineList(context.Background()).Page(page).PageSize(pageSize).PolicyId(policyId).ComponentName(componentName).SortBy(sortBy).Asc(asc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetUnquarantineList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUnquarantineListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **policyId** | **string** |  | 
 **componentName** | **string** |  | 
 **sortBy** | **string** |  | 
 **asc** | **bool** |  | [default to true]

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
	apiFirewallReleaseQuarantineConfigDTO := []sonatypeiq.ApiFirewallReleaseQuarantineConfigDTO{*sonatypeiq.NewApiFirewallReleaseQuarantineConfigDTO()} // []ApiFirewallReleaseQuarantineConfigDTO |  (optional)

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
 **apiFirewallReleaseQuarantineConfigDTO** | [**[]ApiFirewallReleaseQuarantineConfigDTO**](ApiFirewallReleaseQuarantineConfigDTO.md) |  | 

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
	enabled := true // bool | 

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
**enabled** | **bool** |  | 

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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

