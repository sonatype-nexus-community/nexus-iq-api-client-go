# \PolicyWaiverRequestsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPolicyWaiverRequestByPolicyViolationId**](PolicyWaiverRequestsAPI.md#AddPolicyWaiverRequestByPolicyViolationId) | **Post** /api/v2/policyWaiverRequests/{ownerType}/{ownerId}/policyViolation/{policyViolationId} | 
[**GetPolicyWaiverRequest**](PolicyWaiverRequestsAPI.md#GetPolicyWaiverRequest) | **Get** /api/v2/policyWaiverRequests/{ownerType}/{ownerId}/{policyWaiverRequestId} | 
[**ReviewPolicyWaiverRequest**](PolicyWaiverRequestsAPI.md#ReviewPolicyWaiverRequest) | **Post** /api/v2/policyWaiverRequests/{ownerType}/{ownerId}/review/{policyWaiverRequestId} | 
[**UpdatePolicyWaiverRequest**](PolicyWaiverRequestsAPI.md#UpdatePolicyWaiverRequest) | **Put** /api/v2/policyWaiverRequests/{ownerType}/{ownerId}/{policyWaiverRequestId} | 



## AddPolicyWaiverRequestByPolicyViolationId

> ApiPolicyWaiverRequestDTO AddPolicyWaiverRequestByPolicyViolationId(ctx, ownerType, ownerId, policyViolationId).ApiPolicyWaiverRequestOptionsDTO(apiPolicyWaiverRequestOptionsDTO).Execute()





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
	ownerType := "ownerType_example" // string | The scope of the policy waiver request. Possible values are application, organization, repository, repository_manager, repository_container.
	ownerId := "ownerId_example" // string | The id for the ownerType provided above. E.g. applicationId if the ownerType is application.
	policyViolationId := "policyViolationId_example" // string | The policyViolationId for the policy violation on which you want to create a policy waiver request. Use the Policy Violation REST API or Reports REST API to obtain the policyViolationId.
	apiPolicyWaiverRequestOptionsDTO := *sonatypeiq.NewApiPolicyWaiverRequestOptionsDTO() // ApiPolicyWaiverRequestOptionsDTO | The request JSON can include the fields<ol><li>comment (optional, to indicate the reason of the waiver) default value is null</li><li>matcherStrategy (enumeration, required) can have values DEFAULT, EXACT_COMPONENT, ALL_COMPONENTS, ALL_VERSIONS. DEFAULT will match all components if no hash is provided.</li><li>expiryTime (default null) to set the datetime when the waiver expires.</li><li>expireWhenRemediationAvailable (default false) to expire the waiver when a remediation is available.</li><li>noteToReviewer (optional) to add a note to the reviewer</li></ol>

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyWaiverRequestsAPI.AddPolicyWaiverRequestByPolicyViolationId(context.Background(), ownerType, ownerId, policyViolationId).ApiPolicyWaiverRequestOptionsDTO(apiPolicyWaiverRequestOptionsDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiverRequestsAPI.AddPolicyWaiverRequestByPolicyViolationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPolicyWaiverRequestByPolicyViolationId`: ApiPolicyWaiverRequestDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyWaiverRequestsAPI.AddPolicyWaiverRequestByPolicyViolationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | The scope of the policy waiver request. Possible values are application, organization, repository, repository_manager, repository_container. | 
**ownerId** | **string** | The id for the ownerType provided above. E.g. applicationId if the ownerType is application. | 
**policyViolationId** | **string** | The policyViolationId for the policy violation on which you want to create a policy waiver request. Use the Policy Violation REST API or Reports REST API to obtain the policyViolationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPolicyWaiverRequestByPolicyViolationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiPolicyWaiverRequestOptionsDTO** | [**ApiPolicyWaiverRequestOptionsDTO**](ApiPolicyWaiverRequestOptionsDTO.md) | The request JSON can include the fields&lt;ol&gt;&lt;li&gt;comment (optional, to indicate the reason of the waiver) default value is null&lt;/li&gt;&lt;li&gt;matcherStrategy (enumeration, required) can have values DEFAULT, EXACT_COMPONENT, ALL_COMPONENTS, ALL_VERSIONS. DEFAULT will match all components if no hash is provided.&lt;/li&gt;&lt;li&gt;expiryTime (default null) to set the datetime when the waiver expires.&lt;/li&gt;&lt;li&gt;expireWhenRemediationAvailable (default false) to expire the waiver when a remediation is available.&lt;/li&gt;&lt;li&gt;noteToReviewer (optional) to add a note to the reviewer&lt;/li&gt;&lt;/ol&gt; | 

### Return type

[**ApiPolicyWaiverRequestDTO**](ApiPolicyWaiverRequestDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyWaiverRequest

> ApiPolicyWaiverRequestDTO GetPolicyWaiverRequest(ctx, ownerType, ownerId, policyWaiverRequestId).Execute()





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
	ownerType := "ownerType_example" // string | The scope of the policy waiver request. Possible values are application, organization, repository, repository_manager, repository_container.
	ownerId := "ownerId_example" // string | The id for the ownerType provided above.
	policyWaiverRequestId := "policyWaiverRequestId_example" // string | The policyWaiverRequestId for which you want to retrieve the details.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyWaiverRequestsAPI.GetPolicyWaiverRequest(context.Background(), ownerType, ownerId, policyWaiverRequestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiverRequestsAPI.GetPolicyWaiverRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyWaiverRequest`: ApiPolicyWaiverRequestDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyWaiverRequestsAPI.GetPolicyWaiverRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | The scope of the policy waiver request. Possible values are application, organization, repository, repository_manager, repository_container. | 
**ownerId** | **string** | The id for the ownerType provided above. | 
**policyWaiverRequestId** | **string** | The policyWaiverRequestId for which you want to retrieve the details. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyWaiverRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiPolicyWaiverRequestDTO**](ApiPolicyWaiverRequestDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReviewPolicyWaiverRequest

> ApiPolicyWaiverRequestDTO ReviewPolicyWaiverRequest(ctx, ownerType, ownerId, policyWaiverRequestId).ApiPolicyWaiverRequestReviewDTO(apiPolicyWaiverRequestReviewDTO).Execute()





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
	ownerType := "ownerType_example" // string | The scope of the policy waiver request. Possible values are application, organization, repository, repository_manager, repository_container.
	ownerId := "ownerId_example" // string | The id for the ownerType provided above. E.g. applicationId if the ownerType is application.
	policyWaiverRequestId := "policyWaiverRequestId_example" // string | The policyWaiverRequestId for the policy waiver request to be approved or rejected.
	apiPolicyWaiverRequestReviewDTO := *sonatypeiq.NewApiPolicyWaiverRequestReviewDTO() // ApiPolicyWaiverRequestReviewDTO | The request JSON can include the fields<ol><li>status. Can be APPROVED or REJECTED</li><li>rejectionReason (optional). A text explaining the reason for the rejection., <li>comment (optional, to indicate the reason of the waiver) default value is null</li><li>matcherStrategy (enumeration, required) can have values DEFAULT, EXACT_COMPONENT, ALL_COMPONENTS, ALL_VERSIONS. DEFAULT will match all components if no hash is provided.</li><li>expiryTime (default null) to set the datetime when the waiver expires.</li></ol><li>expireWhenRemediationAvailable (default false) to expire the waiver when a remediation is available.</li>

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyWaiverRequestsAPI.ReviewPolicyWaiverRequest(context.Background(), ownerType, ownerId, policyWaiverRequestId).ApiPolicyWaiverRequestReviewDTO(apiPolicyWaiverRequestReviewDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiverRequestsAPI.ReviewPolicyWaiverRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReviewPolicyWaiverRequest`: ApiPolicyWaiverRequestDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyWaiverRequestsAPI.ReviewPolicyWaiverRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | The scope of the policy waiver request. Possible values are application, organization, repository, repository_manager, repository_container. | 
**ownerId** | **string** | The id for the ownerType provided above. E.g. applicationId if the ownerType is application. | 
**policyWaiverRequestId** | **string** | The policyWaiverRequestId for the policy waiver request to be approved or rejected. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReviewPolicyWaiverRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiPolicyWaiverRequestReviewDTO** | [**ApiPolicyWaiverRequestReviewDTO**](ApiPolicyWaiverRequestReviewDTO.md) | The request JSON can include the fields&lt;ol&gt;&lt;li&gt;status. Can be APPROVED or REJECTED&lt;/li&gt;&lt;li&gt;rejectionReason (optional). A text explaining the reason for the rejection., &lt;li&gt;comment (optional, to indicate the reason of the waiver) default value is null&lt;/li&gt;&lt;li&gt;matcherStrategy (enumeration, required) can have values DEFAULT, EXACT_COMPONENT, ALL_COMPONENTS, ALL_VERSIONS. DEFAULT will match all components if no hash is provided.&lt;/li&gt;&lt;li&gt;expiryTime (default null) to set the datetime when the waiver expires.&lt;/li&gt;&lt;/ol&gt;&lt;li&gt;expireWhenRemediationAvailable (default false) to expire the waiver when a remediation is available.&lt;/li&gt; | 

### Return type

[**ApiPolicyWaiverRequestDTO**](ApiPolicyWaiverRequestDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePolicyWaiverRequest

> ApiPolicyWaiverRequestDTO UpdatePolicyWaiverRequest(ctx, ownerType, ownerId, policyWaiverRequestId).ApiPolicyWaiverRequestOptionsDTO(apiPolicyWaiverRequestOptionsDTO).Execute()





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
	ownerType := "ownerType_example" // string | The scope of the policy waiver request. Possible values are application, organization, repository, repository_manager, repository_container.
	ownerId := "ownerId_example" // string | The id for the ownerType provided above. E.g. applicationId if the ownerType is application.
	policyWaiverRequestId := "policyWaiverRequestId_example" // string | The id of the policy waiver request to be updated.
	apiPolicyWaiverRequestOptionsDTO := *sonatypeiq.NewApiPolicyWaiverRequestOptionsDTO() // ApiPolicyWaiverRequestOptionsDTO | The request JSON can include the fields<ol><li>comment (optional, to indicate the reason of the waiver) default value is null</li><li>matcherStrategy (enumeration, required) can have values DEFAULT, EXACT_COMPONENT, ALL_COMPONENTS, ALL_VERSIONS. DEFAULT will match all components if no hash is provided.</li><li>expiryTime (default null) to set the datetime when the waiver expires.</li><li>expireWhenRemediationAvailable (default false) to expire the waiver when a remediation is available.</li><li>noteToReviewer (optional) to add a note to the reviewer</li></ol>

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyWaiverRequestsAPI.UpdatePolicyWaiverRequest(context.Background(), ownerType, ownerId, policyWaiverRequestId).ApiPolicyWaiverRequestOptionsDTO(apiPolicyWaiverRequestOptionsDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiverRequestsAPI.UpdatePolicyWaiverRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePolicyWaiverRequest`: ApiPolicyWaiverRequestDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyWaiverRequestsAPI.UpdatePolicyWaiverRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | The scope of the policy waiver request. Possible values are application, organization, repository, repository_manager, repository_container. | 
**ownerId** | **string** | The id for the ownerType provided above. E.g. applicationId if the ownerType is application. | 
**policyWaiverRequestId** | **string** | The id of the policy waiver request to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyWaiverRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiPolicyWaiverRequestOptionsDTO** | [**ApiPolicyWaiverRequestOptionsDTO**](ApiPolicyWaiverRequestOptionsDTO.md) | The request JSON can include the fields&lt;ol&gt;&lt;li&gt;comment (optional, to indicate the reason of the waiver) default value is null&lt;/li&gt;&lt;li&gt;matcherStrategy (enumeration, required) can have values DEFAULT, EXACT_COMPONENT, ALL_COMPONENTS, ALL_VERSIONS. DEFAULT will match all components if no hash is provided.&lt;/li&gt;&lt;li&gt;expiryTime (default null) to set the datetime when the waiver expires.&lt;/li&gt;&lt;li&gt;expireWhenRemediationAvailable (default false) to expire the waiver when a remediation is available.&lt;/li&gt;&lt;li&gt;noteToReviewer (optional) to add a note to the reviewer&lt;/li&gt;&lt;/ol&gt; | 

### Return type

[**ApiPolicyWaiverRequestDTO**](ApiPolicyWaiverRequestDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

