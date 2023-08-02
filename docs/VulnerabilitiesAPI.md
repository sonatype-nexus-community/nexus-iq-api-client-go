# \VulnerabilitiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSecurityVulnerabilityDetails**](VulnerabilitiesAPI.md#GetSecurityVulnerabilityDetails) | **Get** /api/v2/vulnerabilities/{refId} | 



## GetSecurityVulnerabilityDetails

> SecurityVulnerabilityData GetSecurityVulnerabilityDetails(ctx, refId).ComponentIdentifier(componentIdentifier).IdentificationSource(identificationSource).ScanId(scanId).OwnerType(ownerType).OwnerId(ownerId).Execute()



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
    refId := "refId_example" // string | 
    componentIdentifier := map[string][]sonatypeiq.ComponentIdentifier{ ... } // ComponentIdentifier |  (optional)
    identificationSource := "identificationSource_example" // string |  (optional)
    scanId := "scanId_example" // string |  (optional)
    ownerType := "ownerType_example" // string |  (optional)
    ownerId := "ownerId_example" // string |  (optional)

    configuration := sonatypeiq.NewConfiguration()
    apiClient := sonatypeiq.NewAPIClient(configuration)
    resp, r, err := apiClient.VulnerabilitiesAPI.GetSecurityVulnerabilityDetails(context.Background(), refId).ComponentIdentifier(componentIdentifier).IdentificationSource(identificationSource).ScanId(scanId).OwnerType(ownerType).OwnerId(ownerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VulnerabilitiesAPI.GetSecurityVulnerabilityDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityVulnerabilityDetails`: SecurityVulnerabilityData
    fmt.Fprintf(os.Stdout, "Response from `VulnerabilitiesAPI.GetSecurityVulnerabilityDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityVulnerabilityDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) |  | 
 **identificationSource** | **string** |  | 
 **scanId** | **string** |  | 
 **ownerType** | **string** |  | 
 **ownerId** | **string** |  | 

### Return type

[**SecurityVulnerabilityData**](SecurityVulnerabilityData.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

