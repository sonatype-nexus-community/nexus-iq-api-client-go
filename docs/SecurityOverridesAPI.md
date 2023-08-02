# \SecurityOverridesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecurityVulnerabilityOverrides**](SecurityOverridesAPI.md#GetSecurityVulnerabilityOverrides) | **Get** /api/v2/securityOverrides | 



## GetSecurityVulnerabilityOverrides

> ApiSecurityVulnerabilityOverrideResponseDTOV2 GetSecurityVulnerabilityOverrides(ctx).RefId(refId).ComponentPurl(componentPurl).OwnerId(ownerId).Execute()



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
    refId := "refId_example" // string |  (optional)
    componentPurl := "componentPurl_example" // string |  (optional)
    ownerId := "ownerId_example" // string |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityOverridesAPI.GetSecurityVulnerabilityOverrides(context.Background()).RefId(refId).ComponentPurl(componentPurl).OwnerId(ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityOverridesAPI.GetSecurityVulnerabilityOverrides``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityVulnerabilityOverrides`: ApiSecurityVulnerabilityOverrideResponseDTOV2
    fmt.Fprintf(os.Stdout, "Response from `SecurityOverridesAPI.GetSecurityVulnerabilityOverrides`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityVulnerabilityOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **refId** | **string** |  | 
 **componentPurl** | **string** |  | 
 **ownerId** | **string** |  | 

### Return type

[**ApiSecurityVulnerabilityOverrideResponseDTOV2**](ApiSecurityVulnerabilityOverrideResponseDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

