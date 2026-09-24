# \LegacyViolationsConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfig**](LegacyViolationsConfigurationAPI.md#GetConfig) | **Get** /api/v2/config/legacyViolations/{ownerType}/{ownerId} | 
[**SetConfig**](LegacyViolationsConfigurationAPI.md#SetConfig) | **Put** /api/v2/config/legacyViolations/{ownerType}/{ownerId} | 



## GetConfig

> ApiLegacyViolationStatusDTO GetConfig(ctx, ownerType, ownerId).Execute()





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
	ownerType := "ownerType_example" // string | Owner type. Allowed values: `application`, `organization`.
	ownerId := "ownerId_example" // string | Public id of the application, or id of the organization.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacyViolationsConfigurationAPI.GetConfig(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacyViolationsConfigurationAPI.GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfig`: ApiLegacyViolationStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `LegacyViolationsConfigurationAPI.GetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Owner type. Allowed values: &#x60;application&#x60;, &#x60;organization&#x60;. | 
**ownerId** | **string** | Public id of the application, or id of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiLegacyViolationStatusDTO**](ApiLegacyViolationStatusDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfig

> ApiLegacyViolationStatusDTO SetConfig(ctx, ownerType, ownerId).ApiLegacyViolationStatusDTO(apiLegacyViolationStatusDTO).Execute()





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
	ownerType := "ownerType_example" // string | Owner type. Allowed values: `application`, `organization`.
	ownerId := "ownerId_example" // string | Public id of the application, or id of the organization.
	apiLegacyViolationStatusDTO := *sonatypeiq.NewApiLegacyViolationStatusDTO() // ApiLegacyViolationStatusDTO |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacyViolationsConfigurationAPI.SetConfig(context.Background(), ownerType, ownerId).ApiLegacyViolationStatusDTO(apiLegacyViolationStatusDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacyViolationsConfigurationAPI.SetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetConfig`: ApiLegacyViolationStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `LegacyViolationsConfigurationAPI.SetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Owner type. Allowed values: &#x60;application&#x60;, &#x60;organization&#x60;. | 
**ownerId** | **string** | Public id of the application, or id of the organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiLegacyViolationStatusDTO** | [**ApiLegacyViolationStatusDTO**](ApiLegacyViolationStatusDTO.md) |  | 

### Return type

[**ApiLegacyViolationStatusDTO**](ApiLegacyViolationStatusDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

