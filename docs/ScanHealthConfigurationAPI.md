# \ScanHealthConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration6**](ScanHealthConfigurationAPI.md#DeleteConfiguration6) | **Delete** /api/v2/config/scanHealth/{ownerType}/{ownerId} | 
[**GetConfiguration6**](ScanHealthConfigurationAPI.md#GetConfiguration6) | **Get** /api/v2/config/scanHealth/{ownerType}/{ownerId} | 
[**SaveConfiguration**](ScanHealthConfigurationAPI.md#SaveConfiguration) | **Put** /api/v2/config/scanHealth/{ownerType}/{ownerId} | 



## DeleteConfiguration6

> DeleteConfiguration6(ctx, ownerType, ownerId).Execute()





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
	ownerType := "ownerType_example" // string | The owner type (application or organization)
	ownerId := "ownerId_example" // string | The internal ID of the owner

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ScanHealthConfigurationAPI.DeleteConfiguration6(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScanHealthConfigurationAPI.DeleteConfiguration6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | The owner type (application or organization) | 
**ownerId** | **string** | The internal ID of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfiguration6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration6

> ScanHealthConfigDTO GetConfiguration6(ctx, ownerType, ownerId).Execute()





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
	ownerType := "ownerType_example" // string | The owner type (application or organization)
	ownerId := "ownerId_example" // string | The internal ID of the owner

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ScanHealthConfigurationAPI.GetConfiguration6(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScanHealthConfigurationAPI.GetConfiguration6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration6`: ScanHealthConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `ScanHealthConfigurationAPI.GetConfiguration6`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | The owner type (application or organization) | 
**ownerId** | **string** | The internal ID of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ScanHealthConfigDTO**](ScanHealthConfigDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveConfiguration

> ScanHealthConfigDTO SaveConfiguration(ctx, ownerType, ownerId).ScanHealthConfigDTO(scanHealthConfigDTO).Execute()





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
	ownerType := "ownerType_example" // string | The owner type (application or organization)
	ownerId := "ownerId_example" // string | The internal ID of the owner
	scanHealthConfigDTO := *sonatypeiq.NewScanHealthConfigDTO() // ScanHealthConfigDTO | The Scan Health configuration to save. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ScanHealthConfigurationAPI.SaveConfiguration(context.Background(), ownerType, ownerId).ScanHealthConfigDTO(scanHealthConfigDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScanHealthConfigurationAPI.SaveConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SaveConfiguration`: ScanHealthConfigDTO
	fmt.Fprintf(os.Stdout, "Response from `ScanHealthConfigurationAPI.SaveConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | The owner type (application or organization) | 
**ownerId** | **string** | The internal ID of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scanHealthConfigDTO** | [**ScanHealthConfigDTO**](ScanHealthConfigDTO.md) | The Scan Health configuration to save. | 

### Return type

[**ScanHealthConfigDTO**](ScanHealthConfigDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

