# \CallFlowAnalysisAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCallFlowAnalysisConfig**](CallFlowAnalysisAPI.md#DeleteCallFlowAnalysisConfig) | **Delete** /api/v2/callFlowAnalysis/configuration/{ownerType}/{ownerId} | 
[**GetCallFlowAnalysisConfig**](CallFlowAnalysisAPI.md#GetCallFlowAnalysisConfig) | **Get** /api/v2/callFlowAnalysis/configuration/{ownerType}/{ownerId} | 
[**GetCallFlowAnalysisConfigByPublicId**](CallFlowAnalysisAPI.md#GetCallFlowAnalysisConfigByPublicId) | **Get** /api/v2/callFlowAnalysis/configuration/{ownerType}/{ownerId}/publicId | 
[**UpsertCallFlowAnalysisConfig**](CallFlowAnalysisAPI.md#UpsertCallFlowAnalysisConfig) | **Put** /api/v2/callFlowAnalysis/configuration/{ownerType}/{ownerId} | 



## DeleteCallFlowAnalysisConfig

> DeleteCallFlowAnalysisConfig(ctx, ownerType, ownerId).Execute()



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
	ownerType := "ownerType_example" // string | 
	ownerId := "ownerId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.CallFlowAnalysisAPI.DeleteCallFlowAnalysisConfig(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallFlowAnalysisAPI.DeleteCallFlowAnalysisConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallFlowAnalysisConfigRequest struct via the builder pattern


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


## GetCallFlowAnalysisConfig

> ApiCallFlowAnalysisConfigDTO GetCallFlowAnalysisConfig(ctx, ownerType, ownerId).Execute()



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
	ownerType := "ownerType_example" // string | 
	ownerId := "ownerId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.CallFlowAnalysisAPI.GetCallFlowAnalysisConfig(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallFlowAnalysisAPI.GetCallFlowAnalysisConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallFlowAnalysisConfig`: ApiCallFlowAnalysisConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `CallFlowAnalysisAPI.GetCallFlowAnalysisConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallFlowAnalysisConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiCallFlowAnalysisConfigDTO**](ApiCallFlowAnalysisConfigDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCallFlowAnalysisConfigByPublicId

> ApiCallFlowAnalysisConfigDTO GetCallFlowAnalysisConfigByPublicId(ctx, ownerType, ownerId).Execute()



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
	ownerType := "ownerType_example" // string | 
	ownerId := "ownerId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.CallFlowAnalysisAPI.GetCallFlowAnalysisConfigByPublicId(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallFlowAnalysisAPI.GetCallFlowAnalysisConfigByPublicId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCallFlowAnalysisConfigByPublicId`: ApiCallFlowAnalysisConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `CallFlowAnalysisAPI.GetCallFlowAnalysisConfigByPublicId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallFlowAnalysisConfigByPublicIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiCallFlowAnalysisConfigDTO**](ApiCallFlowAnalysisConfigDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertCallFlowAnalysisConfig

> ApiCallFlowAnalysisConfigDTO UpsertCallFlowAnalysisConfig(ctx, ownerType, ownerId).ApiCallFlowAnalysisConfigDTO(apiCallFlowAnalysisConfigDTO).Execute()



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
	ownerType := "ownerType_example" // string | 
	ownerId := "ownerId_example" // string | 
	apiCallFlowAnalysisConfigDTO := *sonatypeiq.NewApiCallFlowAnalysisConfigDTO() // ApiCallFlowAnalysisConfigDTO |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.CallFlowAnalysisAPI.UpsertCallFlowAnalysisConfig(context.Background(), ownerType, ownerId).ApiCallFlowAnalysisConfigDTO(apiCallFlowAnalysisConfigDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CallFlowAnalysisAPI.UpsertCallFlowAnalysisConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpsertCallFlowAnalysisConfig`: ApiCallFlowAnalysisConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `CallFlowAnalysisAPI.UpsertCallFlowAnalysisConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertCallFlowAnalysisConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiCallFlowAnalysisConfigDTO** | [**ApiCallFlowAnalysisConfigDTO**](ApiCallFlowAnalysisConfigDTO.md) |  | 

### Return type

[**ApiCallFlowAnalysisConfigDTO**](ApiCallFlowAnalysisConfigDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

