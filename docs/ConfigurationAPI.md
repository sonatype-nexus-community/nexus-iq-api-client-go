# \ConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration**](ConfigurationAPI.md#DeleteConfiguration) | **Delete** /api/v2/config | 
[**GetConfiguration**](ConfigurationAPI.md#GetConfiguration) | **Get** /api/v2/config | 
[**SetConfiguration**](ConfigurationAPI.md#SetConfiguration) | **Put** /api/v2/config | 



## DeleteConfiguration

> DeleteConfiguration(ctx).Property(property).Execute()





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
	r, err := apiClient.ConfigurationAPI.DeleteConfiguration(context.Background()).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.DeleteConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **property** | [**[]SystemConfigProperty**](SystemConfigProperty.md) | Enter the names of the system properties. Values provided for name are case-sensitive. | 

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


## GetConfiguration

> SystemConfig GetConfiguration(ctx).Property(property).Execute()





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
	resp, r, err := apiClient.ConfigurationAPI.GetConfiguration(context.Background()).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.GetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration`: SystemConfig
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **property** | [**[]SystemConfigProperty**](SystemConfigProperty.md) | Enter the names of the system properties. Values provided for name are case-sensitive. | 

### Return type

[**SystemConfig**](SystemConfig.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration

> SetConfiguration(ctx).SystemConfig(systemConfig).Execute()





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
	r, err := apiClient.ConfigurationAPI.SetConfiguration(context.Background()).SystemConfig(systemConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationAPI.SetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **systemConfig** | [**SystemConfig**](SystemConfig.md) |  | 

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

