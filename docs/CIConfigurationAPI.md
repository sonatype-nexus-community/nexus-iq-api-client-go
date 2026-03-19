# \CIConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration**](CIConfigurationAPI.md#DeleteConfiguration) | **Delete** /api/v2/config/ci/{ownerType}/{ownerId} | 
[**GetConfiguration**](CIConfigurationAPI.md#GetConfiguration) | **Get** /api/v2/config/ci/{ownerType}/{ownerId} | 
[**SetConfiguration**](CIConfigurationAPI.md#SetConfiguration) | **Put** /api/v2/config/ci/{ownerType}/{ownerId} | 



## DeleteConfiguration

> DeleteConfiguration(ctx, ownerType, ownerId).Execute()





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
	r, err := apiClient.CIConfigurationAPI.DeleteConfiguration(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIConfigurationAPI.DeleteConfiguration``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteConfigurationRequest struct via the builder pattern


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


## GetConfiguration

> ApiCiConfigurationResponseDto GetConfiguration(ctx, ownerType, ownerId).Direct(direct).Execute()





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
	direct := true // bool | Set to true to retrieve only direct configuration, false (default) to retrieve merged configuration from hierarchy (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.CIConfigurationAPI.GetConfiguration(context.Background(), ownerType, ownerId).Direct(direct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIConfigurationAPI.GetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration`: ApiCiConfigurationResponseDto
	fmt.Fprintf(os.Stdout, "Response from `CIConfigurationAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | The owner type (application or organization) | 
**ownerId** | **string** | The internal ID of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **direct** | **bool** | Set to true to retrieve only direct configuration, false (default) to retrieve merged configuration from hierarchy | [default to false]

### Return type

[**ApiCiConfigurationResponseDto**](ApiCiConfigurationResponseDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration

> ApiCiConfigurationDto SetConfiguration(ctx, ownerType, ownerId).ApiCiConfigurationDto(apiCiConfigurationDto).Execute()





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
	apiCiConfigurationDto := *sonatypeiq.NewApiCiConfigurationDto() // ApiCiConfigurationDto | Provide the CI integration configuration as a JSON object. The structure supports different CI systems. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.CIConfigurationAPI.SetConfiguration(context.Background(), ownerType, ownerId).ApiCiConfigurationDto(apiCiConfigurationDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CIConfigurationAPI.SetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetConfiguration`: ApiCiConfigurationDto
	fmt.Fprintf(os.Stdout, "Response from `CIConfigurationAPI.SetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | The owner type (application or organization) | 
**ownerId** | **string** | The internal ID of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiCiConfigurationDto** | [**ApiCiConfigurationDto**](ApiCiConfigurationDto.md) | Provide the CI integration configuration as a JSON object. The structure supports different CI systems. | 

### Return type

[**ApiCiConfigurationDto**](ApiCiConfigurationDto.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

