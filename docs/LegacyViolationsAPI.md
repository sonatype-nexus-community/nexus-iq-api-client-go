# \LegacyViolationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Grant**](LegacyViolationsAPI.md#Grant) | **Post** /api/v2/legacyViolations/application/{applicationPublicId}/grant | 
[**ListLegacyViolations**](LegacyViolationsAPI.md#ListLegacyViolations) | **Get** /api/v2/legacyViolations/application/{applicationPublicId} | 
[**Revoke**](LegacyViolationsAPI.md#Revoke) | **Post** /api/v2/legacyViolations/application/{applicationPublicId}/revoke | 



## Grant

> ApiLegacyViolationChangeResponseDTO Grant(ctx, applicationPublicId).Execute()





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
	applicationPublicId := "applicationPublicId_example" // string | The public id of the application.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacyViolationsAPI.Grant(context.Background(), applicationPublicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacyViolationsAPI.Grant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Grant`: ApiLegacyViolationChangeResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `LegacyViolationsAPI.Grant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** | The public id of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiLegacyViolationChangeResponseDTO**](ApiLegacyViolationChangeResponseDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLegacyViolations

> []ApiPolicyViolationDTOV2 ListLegacyViolations(ctx, applicationPublicId).PolicyId(policyId).ComponentIdentifier(componentIdentifier).Execute()





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
	applicationPublicId := "applicationPublicId_example" // string | The public id of the application.
	policyId := "policyId_example" // string | Optional policy id filter. (optional)
	componentIdentifier := "componentIdentifier_example" // string | Optional component identifier filter, expressed as a package URL (purl). (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacyViolationsAPI.ListLegacyViolations(context.Background(), applicationPublicId).PolicyId(policyId).ComponentIdentifier(componentIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacyViolationsAPI.ListLegacyViolations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLegacyViolations`: []ApiPolicyViolationDTOV2
	fmt.Fprintf(os.Stdout, "Response from `LegacyViolationsAPI.ListLegacyViolations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** | The public id of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLegacyViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **policyId** | **string** | Optional policy id filter. | 
 **componentIdentifier** | **string** | Optional component identifier filter, expressed as a package URL (purl). | 

### Return type

[**[]ApiPolicyViolationDTOV2**](ApiPolicyViolationDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Revoke

> ApiLegacyViolationChangeResponseDTO Revoke(ctx, applicationPublicId).Execute()





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
	applicationPublicId := "applicationPublicId_example" // string | The public id of the application.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LegacyViolationsAPI.Revoke(context.Background(), applicationPublicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LegacyViolationsAPI.Revoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Revoke`: ApiLegacyViolationChangeResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `LegacyViolationsAPI.Revoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** | The public id of the application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiLegacyViolationChangeResponseDTO**](ApiLegacyViolationChangeResponseDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

