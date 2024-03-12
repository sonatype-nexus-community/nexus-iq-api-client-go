# \SbomAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSbomVersion**](SbomAPI.md#DeleteSbomVersion) | **Delete** /api/v2/sbom/{applicationId}/version/{sbomVersion} | Delete sbom version
[**GetSbomVersion**](SbomAPI.md#GetSbomVersion) | **Get** /api/v2/sbom/{applicationId}/version/{sbomVersion} | Gets a sbom version



## DeleteSbomVersion

> DeleteSbomVersion(ctx, applicationId, sbomVersion).Execute()

Delete sbom version



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
	applicationId := "applicationId_example" // string | The internal id of the application
	sbomVersion := "sbomVersion_example" // string | URL Encoded version value of the sbom to be deleted

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SbomAPI.DeleteSbomVersion(context.Background(), applicationId, sbomVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.DeleteSbomVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The internal id of the application | 
**sbomVersion** | **string** | URL Encoded version value of the sbom to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSbomVersionRequest struct via the builder pattern


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


## GetSbomVersion

> GetSbomVersion(ctx, applicationId, sbomVersion).Form(form).Accept(accept).Execute()

Gets a sbom version



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
	applicationId := "applicationId_example" // string | The internal id of the application
	sbomVersion := "sbomVersion_example" // string | URL Encoded version value of the sbom to be deleted
	form := "form_example" // string | The form of the sbom version. Allowed values [original|current]. default = current (optional) (default to "current")
	accept := "accept_example" // string | Output format(json/xml) of the sbom. Changing the output format only applicable when downloading the current form of the SBOM. The original sbom will always return in the original form that it was ingested. When requesting `current` form and if this header value is not present the sbom will be returned in its original ingested format. Allowed values {'application/json'|'application/xml'}. default = null (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SbomAPI.GetSbomVersion(context.Background(), applicationId, sbomVersion).Form(form).Accept(accept).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SbomAPI.GetSbomVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The internal id of the application | 
**sbomVersion** | **string** | URL Encoded version value of the sbom to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSbomVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **form** | **string** | The form of the sbom version. Allowed values [original|current]. default &#x3D; current | [default to &quot;current&quot;]
 **accept** | **string** | Output format(json/xml) of the sbom. Changing the output format only applicable when downloading the current form of the SBOM. The original sbom will always return in the original form that it was ingested. When requesting &#x60;current&#x60; form and if this header value is not present the sbom will be returned in its original ingested format. Allowed values {&#39;application/json&#39;|&#39;application/xml&#39;}. default &#x3D; null | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json|application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

