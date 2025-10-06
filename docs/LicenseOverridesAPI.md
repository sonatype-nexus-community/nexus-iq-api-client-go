# \LicenseOverridesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLicenseOverride**](LicenseOverridesAPI.md#AddLicenseOverride) | **Post** /api/v2/licenseOverrides/{ownerType}/{ownerId} | 
[**DeleteLicenseOverride**](LicenseOverridesAPI.md#DeleteLicenseOverride) | **Delete** /api/v2/licenseOverrides/{ownerType}/{ownerId}/{licenseOverrideId} | 
[**GetAppliedLicenseOverrides**](LicenseOverridesAPI.md#GetAppliedLicenseOverrides) | **Get** /api/v2/licenseOverrides/{ownerType}/{ownerId} | 
[**GetAppliedLicenseOverridesForLegalReviewer**](LicenseOverridesAPI.md#GetAppliedLicenseOverridesForLegalReviewer) | **Get** /api/v2/licenseOverrides/{ownerType}/{ownerId}/legalReviewer | 



## AddLicenseOverride

> ApiLicenseOverrideDTO AddLicenseOverride(ctx, ownerType, ownerId).ApiLicenseOverrideDTO(apiLicenseOverrideDTO).Where(where).Execute()





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
	ownerType := "ownerType_example" // string | Select the `ownerType` scope for which you want to add or update a license override
	ownerId := "ownerId_example" // string | Enter the id of the application, organization or the repository.
	apiLicenseOverrideDTO := *sonatypeiq.NewApiLicenseOverrideDTO() // ApiLicenseOverrideDTO | Enter the license override details to add or update a license override for a component. The request body should contain the following fields:  - `ownerId`: Enter the id of the application, organization or the repository.  - `comment`: Enter a comment for the license override.  - `licenseIds`: Enter the license ids for the license override.  - `componentIdentifier`: Enter the componentIdentifier consisting of format and coordinates.  - `status`: Enter the status of the license override. The possible values are `OPEN`, `ACKNOWLEDGED`, `OVERRIDDEN`, `SELECTED`, and `CONFIRMED`.
	where := "where_example" // string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseOverridesAPI.AddLicenseOverride(context.Background(), ownerType, ownerId).ApiLicenseOverrideDTO(apiLicenseOverrideDTO).Where(where).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseOverridesAPI.AddLicenseOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLicenseOverride`: ApiLicenseOverrideDTO
	fmt.Fprintf(os.Stdout, "Response from `LicenseOverridesAPI.AddLicenseOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the &#x60;ownerType&#x60; scope for which you want to add or update a license override | 
**ownerId** | **string** | Enter the id of the application, organization or the repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLicenseOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiLicenseOverrideDTO** | [**ApiLicenseOverrideDTO**](ApiLicenseOverrideDTO.md) | Enter the license override details to add or update a license override for a component. The request body should contain the following fields:  - &#x60;ownerId&#x60;: Enter the id of the application, organization or the repository.  - &#x60;comment&#x60;: Enter a comment for the license override.  - &#x60;licenseIds&#x60;: Enter the license ids for the license override.  - &#x60;componentIdentifier&#x60;: Enter the componentIdentifier consisting of format and coordinates.  - &#x60;status&#x60;: Enter the status of the license override. The possible values are &#x60;OPEN&#x60;, &#x60;ACKNOWLEDGED&#x60;, &#x60;OVERRIDDEN&#x60;, &#x60;SELECTED&#x60;, and &#x60;CONFIRMED&#x60;. | 
 **where** | **string** |  | 

### Return type

[**ApiLicenseOverrideDTO**](ApiLicenseOverrideDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLicenseOverride

> DeleteLicenseOverride(ctx, ownerType, ownerId, licenseOverrideId).Where(where).Execute()





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
	ownerType := "ownerType_example" // string | Select the `ownerType` scope for which you want to delete license override
	ownerId := "ownerId_example" // string | Enter the id of the application, organization or the repository.
	licenseOverrideId := "licenseOverrideId_example" // string | Enter the id of the license override you want to delete.
	where := "where_example" // string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.LicenseOverridesAPI.DeleteLicenseOverride(context.Background(), ownerType, ownerId, licenseOverrideId).Where(where).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseOverridesAPI.DeleteLicenseOverride``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the &#x60;ownerType&#x60; scope for which you want to delete license override | 
**ownerId** | **string** | Enter the id of the application, organization or the repository. | 
**licenseOverrideId** | **string** | Enter the id of the license override you want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLicenseOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **where** | **string** |  | 

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


## GetAppliedLicenseOverrides

> ApiAppliedLicenseOverridesDTO GetAppliedLicenseOverrides(ctx, ownerType, ownerId).ComponentIdentifier(componentIdentifier).Execute()





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
	ownerType := "ownerType_example" // string | Select the `ownerType` for which you want to retrieve the applied license overrides.
	ownerId := "ownerId_example" // string | Enter the id of the application, organization or the repository.
	componentIdentifier := *sonatypeiq.NewComponentIdentifier() // ComponentIdentifier | Enter the componentIdentifier consisting of format and coordinates as a JSON e.g., `?componentIdentifier={\"format\":\"maven\",\"coordinates\":\"{...}}\"}

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseOverridesAPI.GetAppliedLicenseOverrides(context.Background(), ownerType, ownerId).ComponentIdentifier(componentIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseOverridesAPI.GetAppliedLicenseOverrides``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppliedLicenseOverrides`: ApiAppliedLicenseOverridesDTO
	fmt.Fprintf(os.Stdout, "Response from `LicenseOverridesAPI.GetAppliedLicenseOverrides`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the &#x60;ownerType&#x60; for which you want to retrieve the applied license overrides. | 
**ownerId** | **string** | Enter the id of the application, organization or the repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppliedLicenseOverridesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) | Enter the componentIdentifier consisting of format and coordinates as a JSON e.g., &#x60;?componentIdentifier&#x3D;{\&quot;format\&quot;:\&quot;maven\&quot;,\&quot;coordinates\&quot;:\&quot;{...}}\&quot;} | 

### Return type

[**ApiAppliedLicenseOverridesDTO**](ApiAppliedLicenseOverridesDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppliedLicenseOverridesForLegalReviewer

> ApiAppliedLicenseOverridesDTO GetAppliedLicenseOverridesForLegalReviewer(ctx, ownerType, ownerId).ComponentIdentifier(componentIdentifier).Execute()





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
	ownerType := "ownerType_example" // string | Select the `ownerType` for which you want to retrieve the applied license overrides.
	ownerId := "ownerId_example" // string | Enter the id of the owner.
	componentIdentifier := *sonatypeiq.NewComponentIdentifier() // ComponentIdentifier | Enter the component format and coordinates.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseOverridesAPI.GetAppliedLicenseOverridesForLegalReviewer(context.Background(), ownerType, ownerId).ComponentIdentifier(componentIdentifier).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseOverridesAPI.GetAppliedLicenseOverridesForLegalReviewer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppliedLicenseOverridesForLegalReviewer`: ApiAppliedLicenseOverridesDTO
	fmt.Fprintf(os.Stdout, "Response from `LicenseOverridesAPI.GetAppliedLicenseOverridesForLegalReviewer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the &#x60;ownerType&#x60; for which you want to retrieve the applied license overrides. | 
**ownerId** | **string** | Enter the id of the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppliedLicenseOverridesForLegalReviewerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) | Enter the component format and coordinates. | 

### Return type

[**ApiAppliedLicenseOverridesDTO**](ApiAppliedLicenseOverridesDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

