# \ConfigMailAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration2**](ConfigMailAPI.md#DeleteConfiguration2) | **Delete** /api/v2/config/mail | 
[**GetConfiguration2**](ConfigMailAPI.md#GetConfiguration2) | **Get** /api/v2/config/mail | 
[**SetConfiguration2**](ConfigMailAPI.md#SetConfiguration2) | **Put** /api/v2/config/mail | 
[**TestConfiguration**](ConfigMailAPI.md#TestConfiguration) | **Post** /api/v2/config/mail/test/{recipientEmail} | 



## DeleteConfiguration2

> DeleteConfiguration2(ctx).Execute()



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
    r, err := apiClient.ConfigMailAPI.DeleteConfiguration2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigMailAPI.DeleteConfiguration2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfiguration2Request struct via the builder pattern


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


## GetConfiguration2

> ApiMailConfigurationDTO GetConfiguration2(ctx).Execute()



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
    resp, r, err := apiClient.ConfigMailAPI.GetConfiguration2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigMailAPI.GetConfiguration2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfiguration2`: ApiMailConfigurationDTO
    fmt.Fprintf(os.Stdout, "Response from `ConfigMailAPI.GetConfiguration2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration2Request struct via the builder pattern


### Return type

[**ApiMailConfigurationDTO**](ApiMailConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration2

> SetConfiguration2(ctx).ApiMailConfigurationDTO(apiMailConfigurationDTO).Execute()



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
    apiMailConfigurationDTO := *sonatypeiq.NewApiMailConfigurationDTO() // ApiMailConfigurationDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.ConfigMailAPI.SetConfiguration2(context.Background()).ApiMailConfigurationDTO(apiMailConfigurationDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigMailAPI.SetConfiguration2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfiguration2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiMailConfigurationDTO** | [**ApiMailConfigurationDTO**](ApiMailConfigurationDTO.md) |  | 

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


## TestConfiguration

> TestConfiguration(ctx, recipientEmail).ApiMailConfigurationDTO(apiMailConfigurationDTO).Execute()



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
    recipientEmail := "recipientEmail_example" // string | 
    apiMailConfigurationDTO := *sonatypeiq.NewApiMailConfigurationDTO() // ApiMailConfigurationDTO |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    r, err := apiClient.ConfigMailAPI.TestConfiguration(context.Background(), recipientEmail).ApiMailConfigurationDTO(apiMailConfigurationDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigMailAPI.TestConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recipientEmail** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiMailConfigurationDTO** | [**ApiMailConfigurationDTO**](ApiMailConfigurationDTO.md) |  | 

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

