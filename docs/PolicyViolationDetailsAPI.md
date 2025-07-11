# \PolicyViolationDetailsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicableAutoWaiver**](PolicyViolationDetailsAPI.md#GetApplicableAutoWaiver) | **Get** /api/v2/policyViolations/{violationId}/applicableAutoWaiver | 
[**GetApplicableWaiverRequests**](PolicyViolationDetailsAPI.md#GetApplicableWaiverRequests) | **Get** /api/v2/policyViolations/{violationId}/applicableWaiverRequests | 
[**GetApplicableWaivers**](PolicyViolationDetailsAPI.md#GetApplicableWaivers) | **Get** /api/v2/policyViolations/{violationId}/applicableWaivers | 
[**GetCrossStagePolicyViolationByConstituentId**](PolicyViolationDetailsAPI.md#GetCrossStagePolicyViolationByConstituentId) | **Get** /api/v2/policyViolations/crossStage | 
[**GetCrossStagePolicyViolationById**](PolicyViolationDetailsAPI.md#GetCrossStagePolicyViolationById) | **Get** /api/v2/policyViolations/crossStage/{violationId} | 
[**GetPolicyViolations**](PolicyViolationDetailsAPI.md#GetPolicyViolations) | **Get** /api/v2/policyViolations | 
[**GetSimilarWaivers**](PolicyViolationDetailsAPI.md#GetSimilarWaivers) | **Get** /api/v2/policyViolations/{violationId}/similarWaivers | 
[**GetTransitivePolicyViolationsByAppScanComponent**](PolicyViolationDetailsAPI.md#GetTransitivePolicyViolationsByAppScanComponent) | **Get** /api/v2/policyViolations/transitive/{ownerType}/{ownerId}/{scanId} | 
[**GetTransitivePolicyViolationsByOwnerStageComponent**](PolicyViolationDetailsAPI.md#GetTransitivePolicyViolationsByOwnerStageComponent) | **Get** /api/v2/policyViolations/transitive/{ownerType}/{ownerId}/stages/{stageId} | 



## GetApplicableAutoWaiver

> ApiAutoPolicyWaiverDTO GetApplicableAutoWaiver(ctx, violationId).Execute()





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
	violationId := "violationId_example" // string | Enter the policy violationId for which you want to obtain the applicable auto policy waiver 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyViolationDetailsAPI.GetApplicableAutoWaiver(context.Background(), violationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationDetailsAPI.GetApplicableAutoWaiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicableAutoWaiver`: ApiAutoPolicyWaiverDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyViolationDetailsAPI.GetApplicableAutoWaiver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**violationId** | **string** | Enter the policy violationId for which you want to obtain the applicable auto policy waiver  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableAutoWaiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiAutoPolicyWaiverDTO**](ApiAutoPolicyWaiverDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicableWaiverRequests

> ApiPolicyWaiverRequestsApplicableToViolationDTO GetApplicableWaiverRequests(ctx, violationId).Execute()





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
	violationId := "violationId_example" // string | Enter the policy violationId for which you want to obtain the applicable waiver requests.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyViolationDetailsAPI.GetApplicableWaiverRequests(context.Background(), violationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationDetailsAPI.GetApplicableWaiverRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicableWaiverRequests`: ApiPolicyWaiverRequestsApplicableToViolationDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyViolationDetailsAPI.GetApplicableWaiverRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**violationId** | **string** | Enter the policy violationId for which you want to obtain the applicable waiver requests. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableWaiverRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiPolicyWaiverRequestsApplicableToViolationDTO**](ApiPolicyWaiverRequestsApplicableToViolationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicableWaivers

> ApiPolicyWaiversApplicableToViolationDTO GetApplicableWaivers(ctx, violationId).Execute()





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
	violationId := "violationId_example" // string | Enter the policy violationId for which you want to obtain the applicable waivers.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyViolationDetailsAPI.GetApplicableWaivers(context.Background(), violationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationDetailsAPI.GetApplicableWaivers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicableWaivers`: ApiPolicyWaiversApplicableToViolationDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyViolationDetailsAPI.GetApplicableWaivers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**violationId** | **string** | Enter the policy violationId for which you want to obtain the applicable waivers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableWaiversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiPolicyWaiversApplicableToViolationDTO**](ApiPolicyWaiversApplicableToViolationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossStagePolicyViolationByConstituentId

> ApiCrossStageViolationDTOV2 GetCrossStagePolicyViolationByConstituentId(ctx).ConstituentId(constituentId).Execute()





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
	constituentId := "constituentId_example" // string | Enter the violationId. Use the GET method described for the endpoint /api/v2/policyViolations to obtain the policy violationId.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyViolationDetailsAPI.GetCrossStagePolicyViolationByConstituentId(context.Background()).ConstituentId(constituentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationDetailsAPI.GetCrossStagePolicyViolationByConstituentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrossStagePolicyViolationByConstituentId`: ApiCrossStageViolationDTOV2
	fmt.Fprintf(os.Stdout, "Response from `PolicyViolationDetailsAPI.GetCrossStagePolicyViolationByConstituentId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossStagePolicyViolationByConstituentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **constituentId** | **string** | Enter the violationId. Use the GET method described for the endpoint /api/v2/policyViolations to obtain the policy violationId. | 

### Return type

[**ApiCrossStageViolationDTOV2**](ApiCrossStageViolationDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrossStagePolicyViolationById

> ApiCrossStageViolationDTOV2 GetCrossStagePolicyViolationById(ctx, violationId).Execute()





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
	violationId := "violationId_example" // string | Enter the policy `violationId`. Use the GET method described for the endpoint /api/v2/policyViolations to obtain the policy violationId. 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyViolationDetailsAPI.GetCrossStagePolicyViolationById(context.Background(), violationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationDetailsAPI.GetCrossStagePolicyViolationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrossStagePolicyViolationById`: ApiCrossStageViolationDTOV2
	fmt.Fprintf(os.Stdout, "Response from `PolicyViolationDetailsAPI.GetCrossStagePolicyViolationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**violationId** | **string** | Enter the policy &#x60;violationId&#x60;. Use the GET method described for the endpoint /api/v2/policyViolations to obtain the policy violationId.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrossStagePolicyViolationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiCrossStageViolationDTOV2**](ApiCrossStageViolationDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyViolations

> ApiApplicationViolationListDTOV2 GetPolicyViolations(ctx).P(p).OpenTimeAfter(openTimeAfter).OpenTimeBefore(openTimeBefore).Type_(type_).Execute()





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
	p := []string{"Inner_example"} // []string | Enter the policyIds to obtain the corresponding violation details
	openTimeAfter := "openTimeAfter_example" // string | Enter the date (format YYYY-MM-DD) from which you want to retrieve the violation details (optional)
	openTimeBefore := "openTimeBefore_example" // string | Enter the date (format YYYY-MM-DD) until which you want to retrieve the violation details (optional)
	type_ := []string{"Type_example"} // []string | Set one or more policy violation type (active, legacy, waived) to include (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyViolationDetailsAPI.GetPolicyViolations(context.Background()).P(p).OpenTimeAfter(openTimeAfter).OpenTimeBefore(openTimeBefore).Type_(type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationDetailsAPI.GetPolicyViolations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyViolations`: ApiApplicationViolationListDTOV2
	fmt.Fprintf(os.Stdout, "Response from `PolicyViolationDetailsAPI.GetPolicyViolations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyViolationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **p** | **[]string** | Enter the policyIds to obtain the corresponding violation details | 
 **openTimeAfter** | **string** | Enter the date (format YYYY-MM-DD) from which you want to retrieve the violation details | 
 **openTimeBefore** | **string** | Enter the date (format YYYY-MM-DD) until which you want to retrieve the violation details | 
 **type_** | **[]string** | Set one or more policy violation type (active, legacy, waived) to include | 

### Return type

[**ApiApplicationViolationListDTOV2**](ApiApplicationViolationListDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimilarWaivers

> []ApiPolicyWaiverDTO GetSimilarWaivers(ctx, violationId).Execute()



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
	violationId := "violationId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyViolationDetailsAPI.GetSimilarWaivers(context.Background(), violationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationDetailsAPI.GetSimilarWaivers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSimilarWaivers`: []ApiPolicyWaiverDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyViolationDetailsAPI.GetSimilarWaivers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**violationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimilarWaiversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApiPolicyWaiverDTO**](ApiPolicyWaiverDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitivePolicyViolationsByAppScanComponent

> ApiComponentTransitivePolicyViolationsDTO GetTransitivePolicyViolationsByAppScanComponent(ctx, ownerType, ownerId, scanId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()





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
	ownerType := "ownerType_example" // string | Enter the scope for this violation. Possible values are 'application'
	ownerId := "ownerId_example" // string | Enter the identifier for the scope specified above. E.g. applicationId
	scanId := "scanId_example" // string | Enter the scanId/reportId corresponding to the scan.
	componentIdentifier := *sonatypeiq.NewComponentIdentifier() // ComponentIdentifier | Enter the component identifier and the coordinates of the component for which you want to retrieve the transitive policy violations. This is optional, not required if package URL or hash value is provided. (optional)
	packageUrl := "packageUrl_example" // string | Enter the package URL for the component for which you want to retrieve the transitive policy violations in the specific scan. (optional)
	hash := "hash_example" // string | Enter the hash value for the component for which you want to retrieve the transitive policy violations in the specific scan. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyViolationDetailsAPI.GetTransitivePolicyViolationsByAppScanComponent(context.Background(), ownerType, ownerId, scanId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationDetailsAPI.GetTransitivePolicyViolationsByAppScanComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransitivePolicyViolationsByAppScanComponent`: ApiComponentTransitivePolicyViolationsDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyViolationDetailsAPI.GetTransitivePolicyViolationsByAppScanComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the scope for this violation. Possible values are &#39;application&#39; | 
**ownerId** | **string** | Enter the identifier for the scope specified above. E.g. applicationId | 
**scanId** | **string** | Enter the scanId/reportId corresponding to the scan. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitivePolicyViolationsByAppScanComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) | Enter the component identifier and the coordinates of the component for which you want to retrieve the transitive policy violations. This is optional, not required if package URL or hash value is provided. | 
 **packageUrl** | **string** | Enter the package URL for the component for which you want to retrieve the transitive policy violations in the specific scan. | 
 **hash** | **string** | Enter the hash value for the component for which you want to retrieve the transitive policy violations in the specific scan. | 

### Return type

[**ApiComponentTransitivePolicyViolationsDTO**](ApiComponentTransitivePolicyViolationsDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitivePolicyViolationsByOwnerStageComponent

> ApiComponentTransitivePolicyViolationsDTO GetTransitivePolicyViolationsByOwnerStageComponent(ctx, ownerType, ownerId, stageId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()





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
	ownerType := "ownerType_example" // string | Possible values are 'application' or 'organization'
	ownerId := "ownerId_example" // string | Possible values are applicationId, organizationId
	stageId := "stageId_example" // string | Possible values are 'develop', 'source', 'build', 'stage-release', 'release', and, 'operate'.
	componentIdentifier := *sonatypeiq.NewComponentIdentifier() // ComponentIdentifier | Enter the component identifier and the coordinates of the component for which you want to obtain the transitive violations. This is optional, not required if package URL or hash value is provided. (optional)
	packageUrl := "packageUrl_example" // string | Enter the package URL of the component. This is optional, not required if component identifier or hash value is provided. (optional)
	hash := "hash_example" // string | Enter the hash value of the component. This is optional, not required if component identifier or package URL is provided. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyViolationDetailsAPI.GetTransitivePolicyViolationsByOwnerStageComponent(context.Background(), ownerType, ownerId, stageId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyViolationDetailsAPI.GetTransitivePolicyViolationsByOwnerStageComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransitivePolicyViolationsByOwnerStageComponent`: ApiComponentTransitivePolicyViolationsDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyViolationDetailsAPI.GetTransitivePolicyViolationsByOwnerStageComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Possible values are &#39;application&#39; or &#39;organization&#39; | 
**ownerId** | **string** | Possible values are applicationId, organizationId | 
**stageId** | **string** | Possible values are &#39;develop&#39;, &#39;source&#39;, &#39;build&#39;, &#39;stage-release&#39;, &#39;release&#39;, and, &#39;operate&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitivePolicyViolationsByOwnerStageComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) | Enter the component identifier and the coordinates of the component for which you want to obtain the transitive violations. This is optional, not required if package URL or hash value is provided. | 
 **packageUrl** | **string** | Enter the package URL of the component. This is optional, not required if component identifier or hash value is provided. | 
 **hash** | **string** | Enter the hash value of the component. This is optional, not required if component identifier or package URL is provided. | 

### Return type

[**ApiComponentTransitivePolicyViolationsDTO**](ApiComponentTransitivePolicyViolationsDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

