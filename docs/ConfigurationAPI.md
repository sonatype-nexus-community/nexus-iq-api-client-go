# \ConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration1**](ConfigurationAPI.md#DeleteConfiguration1) | **Delete** /api/v2/config | 
[**GetConfiguration1**](ConfigurationAPI.md#GetConfiguration1) | **Get** /api/v2/config | 
[**InvalidateCache**](ConfigurationAPI.md#InvalidateCache) | **Delete** /api/v2/config/integrationVersions/cache | 
[**SetConfiguration1**](ConfigurationAPI.md#SetConfiguration1) | **Put** /api/v2/config | 



## DeleteConfiguration1

> DeleteConfiguration1(ctx).Property(property).Execute()





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
	property := []sonatypeiq.SystemConfigProperty{sonatypeiq.SystemConfigProperty("baseUrl")} // []SystemConfigProperty | Enter the names of the system properties. Values provided for name are case-sensitive. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationAPI.DeleteConfiguration1(context.Background()).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.DeleteConfiguration1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfiguration1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **property** | [**[]SystemConfigProperty**](SystemConfigProperty.md) | Enter the names of the system properties. Values provided for name are case-sensitive. | 

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


## GetConfiguration1

> SystemConfig GetConfiguration1(ctx).Property(property).Execute()





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
	property := []sonatypeiq.SystemConfigProperty{sonatypeiq.SystemConfigProperty("baseUrl")} // []SystemConfigProperty | Enter the names of the system properties. Values provided for name are case-sensitive. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationAPI.GetConfiguration1(context.Background()).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetConfiguration1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration1`: SystemConfig
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetConfiguration1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **property** | [**[]SystemConfigProperty**](SystemConfigProperty.md) | Enter the names of the system properties. Values provided for name are case-sensitive. | 

### Return type

[**SystemConfig**](SystemConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvalidateCache

> CacheInvalidationResponse InvalidateCache(ctx).Execute()





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
	resp, r, err := apiClient.ConfigurationAPI.InvalidateCache(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.InvalidateCache``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvalidateCache`: CacheInvalidationResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.InvalidateCache`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiInvalidateCacheRequest struct via the builder pattern


### Return type

[**CacheInvalidationResponse**](CacheInvalidationResponse.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration1

> SetConfiguration1(ctx).SystemConfig(systemConfig).Execute()





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
	systemConfig := *sonatypeiq.NewSystemConfig() // SystemConfig |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigurationAPI.SetConfiguration1(context.Background()).SystemConfig(systemConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.SetConfiguration1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfiguration1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemConfig** | [**SystemConfig**](SystemConfig.md) |  | 

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

