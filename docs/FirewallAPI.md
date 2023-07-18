# \FirewallAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFirewallAutoUnquarantineConfig**](FirewallAPI.md#GetFirewallAutoUnquarantineConfig) | **Get** /api/v2/firewall/releaseQuarantine/configuration | 
[**GetFirewallUnquarantineSummary**](FirewallAPI.md#GetFirewallUnquarantineSummary) | **Get** /api/v2/firewall/releaseQuarantine/summary | 
[**GetQuarantineList**](FirewallAPI.md#GetQuarantineList) | **Get** /api/v2/firewall/components/quarantined | 
[**GetQuarantineSummary**](FirewallAPI.md#GetQuarantineSummary) | **Get** /api/v2/firewall/quarantine/summary | 
[**GetQuarantinedComponentViewAnonymousAccess**](FirewallAPI.md#GetQuarantinedComponentViewAnonymousAccess) | **Get** /api/v2/firewall/quarantinedComponentView/configuration/anonymousAccess | 
[**GetUnquarantineList**](FirewallAPI.md#GetUnquarantineList) | **Get** /api/v2/firewall/components/autoReleasedFromQuarantine | 
[**SetFirewallAutoUnquarantineConfig**](FirewallAPI.md#SetFirewallAutoUnquarantineConfig) | **Put** /api/v2/firewall/releaseQuarantine/configuration | 
[**SetQuarantinedComponentViewAnonymousAccess**](FirewallAPI.md#SetQuarantinedComponentViewAnonymousAccess) | **Put** /api/v2/firewall/quarantinedComponentView/configuration/anonymousAccess/{enabled} | 



## GetFirewallAutoUnquarantineConfig

> []ApiFirewallReleaseQuarantineConfigDTO GetFirewallAutoUnquarantineConfig(ctx).Execute()



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
    resp, r, err := apiClient.FirewallAPI.GetFirewallAutoUnquarantineConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetFirewallAutoUnquarantineConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallAutoUnquarantineConfig`: []ApiFirewallReleaseQuarantineConfigDTO
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetFirewallAutoUnquarantineConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallAutoUnquarantineConfigRequest struct via the builder pattern


### Return type

[**[]ApiFirewallReleaseQuarantineConfigDTO**](ApiFirewallReleaseQuarantineConfigDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallUnquarantineSummary

> ApiFirewallReleaseQuarantineSummaryDTO GetFirewallUnquarantineSummary(ctx).Execute()



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
    resp, r, err := apiClient.FirewallAPI.GetFirewallUnquarantineSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetFirewallUnquarantineSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFirewallUnquarantineSummary`: ApiFirewallReleaseQuarantineSummaryDTO
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetFirewallUnquarantineSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallUnquarantineSummaryRequest struct via the builder pattern


### Return type

[**ApiFirewallReleaseQuarantineSummaryDTO**](ApiFirewallReleaseQuarantineSummaryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuarantineList

> GetQuarantineList(ctx).Page(page).PageSize(pageSize).PolicyId(policyId).ComponentName(componentName).SortBy(sortBy).Asc(asc).Execute()



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
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 10)
    policyId := "policyId_example" // string |  (optional)
    componentName := "componentName_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    asc := true // bool |  (optional) (default to true)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallAPI.GetQuarantineList(context.Background()).Page(page).PageSize(pageSize).PolicyId(policyId).ComponentName(componentName).SortBy(sortBy).Asc(asc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetQuarantineList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetQuarantineListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **policyId** | **string** |  | 
 **componentName** | **string** |  | 
 **sortBy** | **string** |  | 
 **asc** | **bool** |  | [default to true]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuarantineSummary

> ApiFirewallQuarantineSummaryDTO GetQuarantineSummary(ctx).Execute()



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
    resp, r, err := apiClient.FirewallAPI.GetQuarantineSummary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetQuarantineSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetQuarantineSummary`: ApiFirewallQuarantineSummaryDTO
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.GetQuarantineSummary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuarantineSummaryRequest struct via the builder pattern


### Return type

[**ApiFirewallQuarantineSummaryDTO**](ApiFirewallQuarantineSummaryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuarantinedComponentViewAnonymousAccess

> GetQuarantinedComponentViewAnonymousAccess(ctx).Execute()



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
    r, err := apiClient.FirewallAPI.GetQuarantinedComponentViewAnonymousAccess(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetQuarantinedComponentViewAnonymousAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuarantinedComponentViewAnonymousAccessRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnquarantineList

> GetUnquarantineList(ctx).Page(page).PageSize(pageSize).PolicyId(policyId).ComponentName(componentName).SortBy(sortBy).Asc(asc).Execute()



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
    page := int32(56) // int32 |  (optional) (default to 1)
    pageSize := int32(56) // int32 |  (optional) (default to 10)
    policyId := "policyId_example" // string |  (optional)
    componentName := "componentName_example" // string |  (optional)
    sortBy := "sortBy_example" // string |  (optional)
    asc := true // bool |  (optional) (default to true)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallAPI.GetUnquarantineList(context.Background()).Page(page).PageSize(pageSize).PolicyId(policyId).ComponentName(componentName).SortBy(sortBy).Asc(asc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.GetUnquarantineList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUnquarantineListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]
 **policyId** | **string** |  | 
 **componentName** | **string** |  | 
 **sortBy** | **string** |  | 
 **asc** | **bool** |  | [default to true]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetFirewallAutoUnquarantineConfig

> []ApiFirewallReleaseQuarantineConfigDTO SetFirewallAutoUnquarantineConfig(ctx).ApiFirewallReleaseQuarantineConfigDTO(apiFirewallReleaseQuarantineConfigDTO).Execute()



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
    apiFirewallReleaseQuarantineConfigDTO := []iqclient.ApiFirewallReleaseQuarantineConfigDTO{*iqclient.NewApiFirewallReleaseQuarantineConfigDTO()} // []ApiFirewallReleaseQuarantineConfigDTO |  (optional)

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FirewallAPI.SetFirewallAutoUnquarantineConfig(context.Background()).ApiFirewallReleaseQuarantineConfigDTO(apiFirewallReleaseQuarantineConfigDTO).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.SetFirewallAutoUnquarantineConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetFirewallAutoUnquarantineConfig`: []ApiFirewallReleaseQuarantineConfigDTO
    fmt.Fprintf(os.Stdout, "Response from `FirewallAPI.SetFirewallAutoUnquarantineConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetFirewallAutoUnquarantineConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiFirewallReleaseQuarantineConfigDTO** | [**[]ApiFirewallReleaseQuarantineConfigDTO**](ApiFirewallReleaseQuarantineConfigDTO.md) |  | 

### Return type

[**[]ApiFirewallReleaseQuarantineConfigDTO**](ApiFirewallReleaseQuarantineConfigDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetQuarantinedComponentViewAnonymousAccess

> SetQuarantinedComponentViewAnonymousAccess(ctx, enabled).Execute()



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
    enabled := true // bool | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.FirewallAPI.SetQuarantinedComponentViewAnonymousAccess(context.Background(), enabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FirewallAPI.SetQuarantinedComponentViewAnonymousAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**enabled** | **bool** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetQuarantinedComponentViewAnonymousAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

