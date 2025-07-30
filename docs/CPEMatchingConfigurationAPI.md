# \CPEMatchingConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCpeMatchingConfiguration**](CPEMatchingConfigurationAPI.md#GetCpeMatchingConfiguration) | **Get** /api/v2/{ownerType}/{internalOwnerId}/configuration/publicSource/cpe | 
[**UpdateCpeMatchingConfiguration**](CPEMatchingConfigurationAPI.md#UpdateCpeMatchingConfiguration) | **Put** /api/v2/{ownerType}/{internalOwnerId}/configuration/publicSource/cpe | 



## GetCpeMatchingConfiguration

> GetCpeMatchingConfiguration(ctx, ownerType, internalOwnerId).Execute()





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
	internalOwnerId := "internalOwnerId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.CPEMatchingConfigurationAPI.GetCpeMatchingConfiguration(context.Background(), ownerType, internalOwnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CPEMatchingConfigurationAPI.GetCpeMatchingConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCpeMatchingConfigurationRequest struct via the builder pattern


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


## UpdateCpeMatchingConfiguration

> UpdateCpeMatchingConfiguration(ctx, ownerType, internalOwnerId).CpeMatchingConfigurationRequest(cpeMatchingConfigurationRequest).Execute()





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
	internalOwnerId := "internalOwnerId_example" // string | 
	cpeMatchingConfigurationRequest := *sonatypeiq.NewCpeMatchingConfigurationRequest() // CpeMatchingConfigurationRequest |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.CPEMatchingConfigurationAPI.UpdateCpeMatchingConfiguration(context.Background(), ownerType, internalOwnerId).CpeMatchingConfigurationRequest(cpeMatchingConfigurationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CPEMatchingConfigurationAPI.UpdateCpeMatchingConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCpeMatchingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cpeMatchingConfigurationRequest** | [**CpeMatchingConfigurationRequest**](CpeMatchingConfigurationRequest.md) |  | 

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

