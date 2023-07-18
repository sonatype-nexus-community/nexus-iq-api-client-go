# \ConfigProxyServerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration3**](ConfigProxyServerAPI.md#DeleteConfiguration3) | **Delete** /api/v2/config/httpProxyServer | 
[**GetConfiguration3**](ConfigProxyServerAPI.md#GetConfiguration3) | **Get** /api/v2/config/httpProxyServer | 
[**SetConfiguration3**](ConfigProxyServerAPI.md#SetConfiguration3) | **Put** /api/v2/config/httpProxyServer | 



## DeleteConfiguration3

> DeleteConfiguration3(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    iqclient "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigProxyServerAPI.DeleteConfiguration3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigProxyServerAPI.DeleteConfiguration3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfiguration3Request struct via the builder pattern


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


## GetConfiguration3

> ApiProxyServerConfigurationDTO GetConfiguration3(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    iqclient "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConfigProxyServerAPI.GetConfiguration3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigProxyServerAPI.GetConfiguration3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfiguration3`: ApiProxyServerConfigurationDTO
    fmt.Fprintf(os.Stdout, "Response from `ConfigProxyServerAPI.GetConfiguration3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration3Request struct via the builder pattern


### Return type

[**ApiProxyServerConfigurationDTO**](ApiProxyServerConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration3

> SetConfiguration3(ctx).ApiProxyServerConfigurationDTO(apiProxyServerConfigurationDTO).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    iqclient "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {
    apiProxyServerConfigurationDTO := *iqclient.NewApiProxyServerConfigurationDTO() // ApiProxyServerConfigurationDTO |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.ConfigProxyServerAPI.SetConfiguration3(context.Background()).ApiProxyServerConfigurationDTO(apiProxyServerConfigurationDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConfigProxyServerAPI.SetConfiguration3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfiguration3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiProxyServerConfigurationDTO** | [**ApiProxyServerConfigurationDTO**](ApiProxyServerConfigurationDTO.md) |  | 

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

