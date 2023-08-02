# \ConfigAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration**](ConfigAPI.md#DeleteConfiguration) | **Delete** /api/v2/config | 
[**DisableFeature**](ConfigAPI.md#DisableFeature) | **Delete** /api/v2/config/features/{feature} | 
[**EnabledFeature**](ConfigAPI.md#EnabledFeature) | **Post** /api/v2/config/features/{feature} | 
[**GetConfiguration**](ConfigAPI.md#GetConfiguration) | **Get** /api/v2/config | 
[**SetConfiguration**](ConfigAPI.md#SetConfiguration) | **Put** /api/v2/config | 



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
    property := []string{"Inner_example"} // []string |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.ConfigAPI.DeleteConfiguration(context.Background()).Property(property).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.DeleteConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **property** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableFeature

> DisableFeature(ctx, feature).Execute()



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
    feature := "feature_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.ConfigAPI.DisableFeature(context.Background(), feature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.DisableFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feature** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnabledFeature

> EnabledFeature(ctx, feature).Execute()



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
    feature := "feature_example" // string | 

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.ConfigAPI.EnabledFeature(context.Background(), feature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.EnabledFeature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feature** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnabledFeatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration

> map[string]map[string]interface{} GetConfiguration(ctx).Property(property).Execute()



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
    property := []string{"Inner_example"} // []string |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigAPI.GetConfiguration(context.Background()).Property(property).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.GetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfiguration`: map[string]map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ConfigAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **property** | **[]string** |  | 

### Return type

**map[string]map[string]interface{}**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration

> SetConfiguration(ctx).RequestBody(requestBody).Execute()



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
    requestBody := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.ConfigAPI.SetConfiguration(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigAPI.SetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

