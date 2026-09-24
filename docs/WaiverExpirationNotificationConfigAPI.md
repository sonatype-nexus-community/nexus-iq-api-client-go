# \WaiverExpirationNotificationConfigAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfig1**](WaiverExpirationNotificationConfigAPI.md#GetConfig1) | **Get** /api/v2/waiverExpirationNotificationConfig/{ownerType}/{ownerId} | 
[**SaveConfig**](WaiverExpirationNotificationConfigAPI.md#SaveConfig) | **Put** /api/v2/waiverExpirationNotificationConfig/{ownerType}/{ownerId} | 



## GetConfig1

> ApiWaiverExpirationNotificationConfigDTO GetConfig1(ctx, ownerType, ownerId).Execute()





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
	ownerType := "ownerType_example" // string | The owner type (organization, repository_manager, or repository_container).
	ownerId := "ownerId_example" // string | The internal owner ID assigned by IQ Server.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.WaiverExpirationNotificationConfigAPI.GetConfig1(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WaiverExpirationNotificationConfigAPI.GetConfig1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfig1`: ApiWaiverExpirationNotificationConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `WaiverExpirationNotificationConfigAPI.GetConfig1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | The owner type (organization, repository_manager, or repository_container). | 
**ownerId** | **string** | The internal owner ID assigned by IQ Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfig1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiWaiverExpirationNotificationConfigDTO**](ApiWaiverExpirationNotificationConfigDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveConfig

> SaveConfig(ctx, ownerType, ownerId).ApiWaiverExpirationNotificationConfigDTO(apiWaiverExpirationNotificationConfigDTO).Execute()





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
	ownerType := "ownerType_example" // string | The owner type (organization, repository_manager, or repository_container).
	ownerId := "ownerId_example" // string | The internal owner ID assigned by IQ Server.
	apiWaiverExpirationNotificationConfigDTO := *sonatypeiq.NewApiWaiverExpirationNotificationConfigDTO() // ApiWaiverExpirationNotificationConfigDTO | The notification configuration to save.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.WaiverExpirationNotificationConfigAPI.SaveConfig(context.Background(), ownerType, ownerId).ApiWaiverExpirationNotificationConfigDTO(apiWaiverExpirationNotificationConfigDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WaiverExpirationNotificationConfigAPI.SaveConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | The owner type (organization, repository_manager, or repository_container). | 
**ownerId** | **string** | The internal owner ID assigned by IQ Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiWaiverExpirationNotificationConfigDTO** | [**ApiWaiverExpirationNotificationConfigDTO**](ApiWaiverExpirationNotificationConfigDTO.md) | The notification configuration to save. | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

