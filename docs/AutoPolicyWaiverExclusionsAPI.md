# \AutoPolicyWaiverExclusionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAutoPolicyWaiveExclusion**](AutoPolicyWaiverExclusionsAPI.md#AddAutoPolicyWaiveExclusion) | **Post** /api/v2/autoPolicyWaiverExclusions/{ownerType}/{ownerId} | 
[**DeleteAutoPolicyWaiverExclusion**](AutoPolicyWaiverExclusionsAPI.md#DeleteAutoPolicyWaiverExclusion) | **Delete** /api/v2/autoPolicyWaiverExclusions/{ownerType}/{ownerId}/{autoPolicyWaiverId}/{autoPolicyWaiverExclusionId} | 
[**GetAutoPolicyWaiverExclusions**](AutoPolicyWaiverExclusionsAPI.md#GetAutoPolicyWaiverExclusions) | **Get** /api/v2/autoPolicyWaiverExclusions/{ownerType}/{ownerId}/{autoPolicyWaiverId} | 



## AddAutoPolicyWaiveExclusion

> ApiAutoPolicyWaiverExclusionResponseDTO AddAutoPolicyWaiveExclusion(ctx, ownerType, ownerId).ApiAutoPolicyWaiverExclusionRequestDTO(apiAutoPolicyWaiverExclusionRequestDTO).Execute()





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
	ownerType := "ownerType_example" // string | Enter the ownerType to specify which resource type owns the auto waiver you want to apply a exclusion to. Possible values are application, organization.
	ownerId := "ownerId_example" // string | Enter the corresponding id for the ownerType specified above.
	apiAutoPolicyWaiverExclusionRequestDTO := *sonatypeiq.NewApiAutoPolicyWaiverExclusionRequestDTO() // ApiAutoPolicyWaiverExclusionRequestDTO | The request JSON can include the fields<ol><li>applicationPublicId</li><li>ownerId - ID of the application or organization which will own the auto waiver exclusion</li><li>policyViolationId - ID of the policy violation which the exclusion will apply to</li><li>autoPolicyWaiverId - ID of the auto waiver you want to apply a exclusion to</li><li>scanId - ID of the scan which the violation being waived appeared in</li><li>matchStrategy (enumeration, required) can have values EXACT_COMPONENT, ALL_VERSIONS, POLICY_VIOLATION. </li></ol>

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoPolicyWaiverExclusionsAPI.AddAutoPolicyWaiveExclusion(context.Background(), ownerType, ownerId).ApiAutoPolicyWaiverExclusionRequestDTO(apiAutoPolicyWaiverExclusionRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPolicyWaiverExclusionsAPI.AddAutoPolicyWaiveExclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAutoPolicyWaiveExclusion`: ApiAutoPolicyWaiverExclusionResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `AutoPolicyWaiverExclusionsAPI.AddAutoPolicyWaiveExclusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify which resource type owns the auto waiver you want to apply a exclusion to. Possible values are application, organization. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAutoPolicyWaiveExclusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAutoPolicyWaiverExclusionRequestDTO** | [**ApiAutoPolicyWaiverExclusionRequestDTO**](ApiAutoPolicyWaiverExclusionRequestDTO.md) | The request JSON can include the fields&lt;ol&gt;&lt;li&gt;applicationPublicId&lt;/li&gt;&lt;li&gt;ownerId - ID of the application or organization which will own the auto waiver exclusion&lt;/li&gt;&lt;li&gt;policyViolationId - ID of the policy violation which the exclusion will apply to&lt;/li&gt;&lt;li&gt;autoPolicyWaiverId - ID of the auto waiver you want to apply a exclusion to&lt;/li&gt;&lt;li&gt;scanId - ID of the scan which the violation being waived appeared in&lt;/li&gt;&lt;li&gt;matchStrategy (enumeration, required) can have values EXACT_COMPONENT, ALL_VERSIONS, POLICY_VIOLATION. &lt;/li&gt;&lt;/ol&gt; | 

### Return type

[**ApiAutoPolicyWaiverExclusionResponseDTO**](ApiAutoPolicyWaiverExclusionResponseDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoPolicyWaiverExclusion

> DeleteAutoPolicyWaiverExclusion(ctx, ownerType, ownerId, autoPolicyWaiverId, autoPolicyWaiverExclusionId).Execute()





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
	ownerType := "ownerType_example" // string | Enter the ownerType to specify the scope. A waiver exclusion corresponding to the autoPolicyWaiverExclusionId provided and within the scope specified will be deleted.
	ownerId := "ownerId_example" // string | Enter the corresponding id for the ownerType specified above.
	autoPolicyWaiverId := "autoPolicyWaiverId_example" // string | Enter the relevant Auto Policy Waiver ID.
	autoPolicyWaiverExclusionId := "autoPolicyWaiverExclusionId_example" // string | Enter the autoPolicyWaiverId to be deleted

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.AutoPolicyWaiverExclusionsAPI.DeleteAutoPolicyWaiverExclusion(context.Background(), ownerType, ownerId, autoPolicyWaiverId, autoPolicyWaiverExclusionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPolicyWaiverExclusionsAPI.DeleteAutoPolicyWaiverExclusion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. A waiver exclusion corresponding to the autoPolicyWaiverExclusionId provided and within the scope specified will be deleted. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 
**autoPolicyWaiverId** | **string** | Enter the relevant Auto Policy Waiver ID. | 
**autoPolicyWaiverExclusionId** | **string** | Enter the autoPolicyWaiverId to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoPolicyWaiverExclusionRequest struct via the builder pattern


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


## GetAutoPolicyWaiverExclusions

> []ApiAutoPolicyWaiverExclusionResponseDTO GetAutoPolicyWaiverExclusions(ctx, ownerType, ownerId, autoPolicyWaiverId).Page(page).PageSize(pageSize).Execute()



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
	ownerType := "ownerType_example" // string | 
	ownerId := "ownerId_example" // string | 
	autoPolicyWaiverId := "autoPolicyWaiverId_example" // string | 
	page := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 10)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoPolicyWaiverExclusionsAPI.GetAutoPolicyWaiverExclusions(context.Background(), ownerType, ownerId, autoPolicyWaiverId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPolicyWaiverExclusionsAPI.GetAutoPolicyWaiverExclusions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoPolicyWaiverExclusions`: []ApiAutoPolicyWaiverExclusionResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `AutoPolicyWaiverExclusionsAPI.GetAutoPolicyWaiverExclusions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 
**autoPolicyWaiverId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoPolicyWaiverExclusionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 10]

### Return type

[**[]ApiAutoPolicyWaiverExclusionResponseDTO**](ApiAutoPolicyWaiverExclusionResponseDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

