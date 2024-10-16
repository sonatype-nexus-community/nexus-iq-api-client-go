# \PolicyWaiversAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPolicyWaiverByPolicyViolationId**](PolicyWaiversAPI.md#AddPolicyWaiverByPolicyViolationId) | **Post** /api/v2/policyWaivers/{ownerType}/{ownerId}/{policyViolationId} | 
[**AddWaiverToTransitivePolicyViolationsByAppScanComponent**](PolicyWaiversAPI.md#AddWaiverToTransitivePolicyViolationsByAppScanComponent) | **Post** /api/v2/policyWaivers/transitive/{ownerType}/{ownerId}/{scanId} | 
[**AddWaiverToTransitivePolicyViolationsByOwnerStageComponent**](PolicyWaiversAPI.md#AddWaiverToTransitivePolicyViolationsByOwnerStageComponent) | **Post** /api/v2/policyWaivers/transitive/{ownerType}/{ownerId}/stages/{stageId} | 
[**DeletePolicyWaiver**](PolicyWaiversAPI.md#DeletePolicyWaiver) | **Delete** /api/v2/policyWaivers/{ownerType}/{ownerId}/{policyWaiverId} | 
[**GetPolicyWaiver**](PolicyWaiversAPI.md#GetPolicyWaiver) | **Get** /api/v2/policyWaivers/{ownerType}/{ownerId}/{policyWaiverId} | 
[**GetPolicyWaivers**](PolicyWaiversAPI.md#GetPolicyWaivers) | **Get** /api/v2/policyWaivers/{ownerType}/{ownerId} | 
[**GetTransitivePolicyWaiversByAppScanComponent**](PolicyWaiversAPI.md#GetTransitivePolicyWaiversByAppScanComponent) | **Get** /api/v2/policyWaivers/transitive/{ownerType}/{ownerId}/{scanId} | 
[**RequestPolicyWaiver**](PolicyWaiversAPI.md#RequestPolicyWaiver) | **Post** /api/v2/policyWaivers/waiverRequests/{policyViolationId} | 



## AddPolicyWaiverByPolicyViolationId

> AddPolicyWaiverByPolicyViolationId(ctx, ownerType, ownerId, policyViolationId).ApiWaiverOptionsDTO(apiWaiverOptionsDTO).Execute()





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
	ownerType := "ownerType_example" // string | Indicates the scope of the waiver. Possible values are application, organization, repository, repository_manager, repository_container, global.
	ownerId := "ownerId_example" // string | Enter the id for the ownerType provided above. E.g. applicationId if the ownerType is application.
	policyViolationId := "policyViolationId_example" // string | Enter the policyViolationId for the policy on which you want to create a waiver. Use the Policy Violation REST API or Reports REST API to obtain the policyViolationId.
	apiWaiverOptionsDTO := *sonatypeiq.NewApiWaiverOptionsDTO() // ApiWaiverOptionsDTO | The request JSON can include the fields<ol><li>comment (optional, to indicate the reason of the waiver) default value is null</li><li>applyToAllComponents (boolean, default 'false'),deprecated in favor of matcherStrategy. If matcherStrategy is not set, 'true' means this will apply the waiver to all components, 'false' means this will apply to a specific component.</li><li>matcherStrategy (enumeration, required) can have values DEFAULT, EXACT_COMPONENT, ALL_COMPONENTS, ALL_VERSIONS. DEFAULT will match all components if no hash is provided.</li><li>expiryTime (default null) to set the datetime when the waiver expires.</li></ol>

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.PolicyWaiversAPI.AddPolicyWaiverByPolicyViolationId(context.Background(), ownerType, ownerId, policyViolationId).ApiWaiverOptionsDTO(apiWaiverOptionsDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.AddPolicyWaiverByPolicyViolationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Indicates the scope of the waiver. Possible values are application, organization, repository, repository_manager, repository_container, global. | 
**ownerId** | **string** | Enter the id for the ownerType provided above. E.g. applicationId if the ownerType is application. | 
**policyViolationId** | **string** | Enter the policyViolationId for the policy on which you want to create a waiver. Use the Policy Violation REST API or Reports REST API to obtain the policyViolationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPolicyWaiverByPolicyViolationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiWaiverOptionsDTO** | [**ApiWaiverOptionsDTO**](ApiWaiverOptionsDTO.md) | The request JSON can include the fields&lt;ol&gt;&lt;li&gt;comment (optional, to indicate the reason of the waiver) default value is null&lt;/li&gt;&lt;li&gt;applyToAllComponents (boolean, default &#39;false&#39;),deprecated in favor of matcherStrategy. If matcherStrategy is not set, &#39;true&#39; means this will apply the waiver to all components, &#39;false&#39; means this will apply to a specific component.&lt;/li&gt;&lt;li&gt;matcherStrategy (enumeration, required) can have values DEFAULT, EXACT_COMPONENT, ALL_COMPONENTS, ALL_VERSIONS. DEFAULT will match all components if no hash is provided.&lt;/li&gt;&lt;li&gt;expiryTime (default null) to set the datetime when the waiver expires.&lt;/li&gt;&lt;/ol&gt; | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWaiverToTransitivePolicyViolationsByAppScanComponent

> AddWaiverToTransitivePolicyViolationsByAppScanComponent(ctx, ownerType, ownerId, scanId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).ApiWaiverOptionsDTO(apiWaiverOptionsDTO).Execute()





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
	ownerType := "ownerType_example" // string | Indicates the scope of the waiver that will be created.
	ownerId := "ownerId_example" // string | Enter the corresponding id for the ownerType specified above.
	scanId := "scanId_example" // string | Enter the scanId (reportId) of the evaluation report that shows the transitive component.
	componentIdentifier := *sonatypeiq.NewComponentIdentifier() // ComponentIdentifier | Enter the component identifier of the transitive component on which you want to create a policy waiver. (optional)
	packageUrl := "packageUrl_example" // string | Enter the package URL of the transitive component on which you want to create a policy waiver. (optional)
	hash := "hash_example" // string | Enter the hash of the transitive component on which you want to create a policy waiver. (optional)
	apiWaiverOptionsDTO := *sonatypeiq.NewApiWaiverOptionsDTO() // ApiWaiverOptionsDTO | The request JSON can include the fields<ol><li>comment (optional, to indicate the reason of the waiver) default value is null</li><li>applyToAllComponents (boolean, default 'false'),deprecated in favor of matcherStrategy. If matcherStrategy is not set, 'true' means this will apply the waiver to all components, 'false' means this will apply to a specific component.</li><li>matcherStrategy (enumeration, required) can have values DEFAULT, EXACT_COMPONENT, ALL_COMPONENTS, ALL_VERSIONS. DEFAULT will match all components if no hash is provided.</li><li>expiryTime (default null) to set the datetime when the waiver expires.</li></ol> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.PolicyWaiversAPI.AddWaiverToTransitivePolicyViolationsByAppScanComponent(context.Background(), ownerType, ownerId, scanId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).ApiWaiverOptionsDTO(apiWaiverOptionsDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.AddWaiverToTransitivePolicyViolationsByAppScanComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Indicates the scope of the waiver that will be created. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 
**scanId** | **string** | Enter the scanId (reportId) of the evaluation report that shows the transitive component. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) | Enter the component identifier of the transitive component on which you want to create a policy waiver. | 
 **packageUrl** | **string** | Enter the package URL of the transitive component on which you want to create a policy waiver. | 
 **hash** | **string** | Enter the hash of the transitive component on which you want to create a policy waiver. | 
 **apiWaiverOptionsDTO** | [**ApiWaiverOptionsDTO**](ApiWaiverOptionsDTO.md) | The request JSON can include the fields&lt;ol&gt;&lt;li&gt;comment (optional, to indicate the reason of the waiver) default value is null&lt;/li&gt;&lt;li&gt;applyToAllComponents (boolean, default &#39;false&#39;),deprecated in favor of matcherStrategy. If matcherStrategy is not set, &#39;true&#39; means this will apply the waiver to all components, &#39;false&#39; means this will apply to a specific component.&lt;/li&gt;&lt;li&gt;matcherStrategy (enumeration, required) can have values DEFAULT, EXACT_COMPONENT, ALL_COMPONENTS, ALL_VERSIONS. DEFAULT will match all components if no hash is provided.&lt;/li&gt;&lt;li&gt;expiryTime (default null) to set the datetime when the waiver expires.&lt;/li&gt;&lt;/ol&gt; | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddWaiverToTransitivePolicyViolationsByOwnerStageComponent

> AddWaiverToTransitivePolicyViolationsByOwnerStageComponent(ctx, ownerType, ownerId, stageId).ApiWaiverOptionsDTO(apiWaiverOptionsDTO).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()





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
	ownerType := "ownerType_example" // string | Indicates the scope of the waiver that will be created.
	ownerId := "ownerId_example" // string | Enter the corresponding id for the ownerType specified above. E.g. applicationId for ownerType 'application' or organizationId for ownerType 'organization'.
	stageId := "stageId_example" // string | Enter the stageId corresponding to the evaluation stage at which you want to create a waiver. Possible values are 'develop', 'source', 'build', 'stage-release', 'release' and 'operate'.
	apiWaiverOptionsDTO := *sonatypeiq.NewApiWaiverOptionsDTO() // ApiWaiverOptionsDTO | <ol><li>comment (optional, to indicate the reason of the waiver) default value is null</li><li>applyToAllComponents (boolean, default 'false'),deprecated in favor of matcherStrategy. If matcherStrategy is not set, 'true' means this will apply the waiver to all components, 'false' means this will apply to a specific component.</li><li>matcherStrategy (enumeration, required) can have values DEFAULT, EXACT_COMPONENT, ALL_COMPONENTS, ALL_VERSIONS. DEFAULT will match all components if no hash is provided.</li><li>expiryTime (default null) to set the datetime when the waiver expires.</li></ol>
	componentIdentifier := *sonatypeiq.NewComponentIdentifier() // ComponentIdentifier | Enter the component identifier and coordinates of the component for which you want to waive the transitive violations. (optional)
	packageUrl := "packageUrl_example" // string | Enter the package URL of the component for which you want to waive the transitive violations. (optional)
	hash := "hash_example" // string | Enter the hash for the component for which you want to waive the transitive violations  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.PolicyWaiversAPI.AddWaiverToTransitivePolicyViolationsByOwnerStageComponent(context.Background(), ownerType, ownerId, stageId).ApiWaiverOptionsDTO(apiWaiverOptionsDTO).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.AddWaiverToTransitivePolicyViolationsByOwnerStageComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Indicates the scope of the waiver that will be created. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. E.g. applicationId for ownerType &#39;application&#39; or organizationId for ownerType &#39;organization&#39;. | 
**stageId** | **string** | Enter the stageId corresponding to the evaluation stage at which you want to create a waiver. Possible values are &#39;develop&#39;, &#39;source&#39;, &#39;build&#39;, &#39;stage-release&#39;, &#39;release&#39; and &#39;operate&#39;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiWaiverOptionsDTO** | [**ApiWaiverOptionsDTO**](ApiWaiverOptionsDTO.md) | &lt;ol&gt;&lt;li&gt;comment (optional, to indicate the reason of the waiver) default value is null&lt;/li&gt;&lt;li&gt;applyToAllComponents (boolean, default &#39;false&#39;),deprecated in favor of matcherStrategy. If matcherStrategy is not set, &#39;true&#39; means this will apply the waiver to all components, &#39;false&#39; means this will apply to a specific component.&lt;/li&gt;&lt;li&gt;matcherStrategy (enumeration, required) can have values DEFAULT, EXACT_COMPONENT, ALL_COMPONENTS, ALL_VERSIONS. DEFAULT will match all components if no hash is provided.&lt;/li&gt;&lt;li&gt;expiryTime (default null) to set the datetime when the waiver expires.&lt;/li&gt;&lt;/ol&gt; | 
 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) | Enter the component identifier and coordinates of the component for which you want to waive the transitive violations. | 
 **packageUrl** | **string** | Enter the package URL of the component for which you want to waive the transitive violations. | 
 **hash** | **string** | Enter the hash for the component for which you want to waive the transitive violations  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePolicyWaiver

> DeletePolicyWaiver(ctx, ownerType, ownerId, policyWaiverId).Execute()





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
	ownerType := "ownerType_example" // string | Enter the ownerType to specify the scope. A waiver corresponding to the policyWaiverId provided and within the scope specified will be deleted.
	ownerId := "ownerId_example" // string | Enter the corresponding id for the ownerType specified above.
	policyWaiverId := "policyWaiverId_example" // string | Enter the policyWaiverId to be deleted.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.PolicyWaiversAPI.DeletePolicyWaiver(context.Background(), ownerType, ownerId, policyWaiverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.DeletePolicyWaiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. A waiver corresponding to the policyWaiverId provided and within the scope specified will be deleted. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 
**policyWaiverId** | **string** | Enter the policyWaiverId to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyWaiverRequest struct via the builder pattern


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


## GetPolicyWaiver

> ApiPolicyWaiverDTO GetPolicyWaiver(ctx, ownerType, ownerId, policyWaiverId).Execute()





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
	ownerType := "ownerType_example" // string | Enter the ownerType to specify the scope. The response will contain the details for waivers within the scope.
	ownerId := "ownerId_example" // string | Enter the corresponding id for the ownerType specified above.
	policyWaiverId := "policyWaiverId_example" // string | Enter the policyWaiverId for which you want to retrieve the waiver details.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyWaiversAPI.GetPolicyWaiver(context.Background(), ownerType, ownerId, policyWaiverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.GetPolicyWaiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyWaiver`: ApiPolicyWaiverDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyWaiversAPI.GetPolicyWaiver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. The response will contain the details for waivers within the scope. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 
**policyWaiverId** | **string** | Enter the policyWaiverId for which you want to retrieve the waiver details. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyWaiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiPolicyWaiverDTO**](ApiPolicyWaiverDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyWaivers

> []ApiPolicyWaiverDTO GetPolicyWaivers(ctx, ownerType, ownerId).Execute()





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
	ownerType := "ownerType_example" // string | Enter the ownerType to specify the scope. The response will contain waivers that are within the scope specified.
	ownerId := "ownerId_example" // string | Enter the corresponding id for the ownerType specified above.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyWaiversAPI.GetPolicyWaivers(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.GetPolicyWaivers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyWaivers`: []ApiPolicyWaiverDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyWaiversAPI.GetPolicyWaivers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. The response will contain waivers that are within the scope specified. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyWaiversRequest struct via the builder pattern


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


## GetTransitivePolicyWaiversByAppScanComponent

> ApiComponentPolicyWaiversDTO GetTransitivePolicyWaiversByAppScanComponent(ctx, ownerType, ownerId, scanId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()





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
	ownerType := "ownerType_example" // string | Enter the ownerType to specify the scope. The response will contain the policy violations that are within the scope specified.
	ownerId := "ownerId_example" // string | Enter the corresponding id for the ownerType specified above.
	scanId := "scanId_example" // string | Enter the scanId (reportId) of the scan for which you want to retrieve the waivers on transitive policy violations occurring due the dependencies of a component.
	componentIdentifier := *sonatypeiq.NewComponentIdentifier() // ComponentIdentifier | Enter the component identifier for the component for which you want to retrieve the waivers on transitive policy violations, for the specified scanId. (optional)
	packageUrl := "packageUrl_example" // string | Enter the package URL for the component for which you want to retrieve the waivers on transitive policy violations, for the specified scanId. (optional)
	hash := "hash_example" // string | Enter the hash for the component for which you want to retrieve the waivers on transitive policy violations, for the specified scanId. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyWaiversAPI.GetTransitivePolicyWaiversByAppScanComponent(context.Background(), ownerType, ownerId, scanId).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Hash(hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.GetTransitivePolicyWaiversByAppScanComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransitivePolicyWaiversByAppScanComponent`: ApiComponentPolicyWaiversDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyWaiversAPI.GetTransitivePolicyWaiversByAppScanComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. The response will contain the policy violations that are within the scope specified. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 
**scanId** | **string** | Enter the scanId (reportId) of the scan for which you want to retrieve the waivers on transitive policy violations occurring due the dependencies of a component. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitivePolicyWaiversByAppScanComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) | Enter the component identifier for the component for which you want to retrieve the waivers on transitive policy violations, for the specified scanId. | 
 **packageUrl** | **string** | Enter the package URL for the component for which you want to retrieve the waivers on transitive policy violations, for the specified scanId. | 
 **hash** | **string** | Enter the hash for the component for which you want to retrieve the waivers on transitive policy violations, for the specified scanId. | 

### Return type

[**ApiComponentPolicyWaiversDTO**](ApiComponentPolicyWaiversDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestPolicyWaiver

> RequestPolicyWaiver(ctx, policyViolationId).ApiRequestPolicyWaiverDTO(apiRequestPolicyWaiverDTO).Execute()





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
	policyViolationId := "policyViolationId_example" // string | Enter the policyViolationId for which you want to trigger the waiver request event.
	apiRequestPolicyWaiverDTO := *sonatypeiq.NewApiRequestPolicyWaiverDTO() // ApiRequestPolicyWaiverDTO | The request JSON should contain<ol><li>comment (optional, default null) to indicate the waiver request reason</li><li>policyViolationLink (link to the policy violation page in the Lifecycle UI)</li><li>addWaiverLink (link to the Add Waiver page in the Lifecycle UI)</li></ol> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.PolicyWaiversAPI.RequestPolicyWaiver(context.Background(), policyViolationId).ApiRequestPolicyWaiverDTO(apiRequestPolicyWaiverDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiversAPI.RequestPolicyWaiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyViolationId** | **string** | Enter the policyViolationId for which you want to trigger the waiver request event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestPolicyWaiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiRequestPolicyWaiverDTO** | [**ApiRequestPolicyWaiverDTO**](ApiRequestPolicyWaiverDTO.md) | The request JSON should contain&lt;ol&gt;&lt;li&gt;comment (optional, default null) to indicate the waiver request reason&lt;/li&gt;&lt;li&gt;policyViolationLink (link to the policy violation page in the Lifecycle UI)&lt;/li&gt;&lt;li&gt;addWaiverLink (link to the Add Waiver page in the Lifecycle UI)&lt;/li&gt;&lt;/ol&gt; | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

