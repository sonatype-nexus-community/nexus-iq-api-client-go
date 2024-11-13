# \AutoPolicyWaiverRevocationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAutoPolicyWaiverRevocation**](AutoPolicyWaiverRevocationsAPI.md#AddAutoPolicyWaiverRevocation) | **Post** /api/v2/autoPolicyWaiverRevocations/{ownerType}/{ownerId} | 
[**DeleteAutoPolicyWaiverRevocation**](AutoPolicyWaiverRevocationsAPI.md#DeleteAutoPolicyWaiverRevocation) | **Delete** /api/v2/autoPolicyWaiverRevocations/{ownerType}/{ownerId}/{autoPolicyWaiverRevocationId} | 



## AddAutoPolicyWaiverRevocation

> ApiAutoPolicyWaiverRevocationDTO AddAutoPolicyWaiverRevocation(ctx, ownerType, ownerId).ApiAutoPolicyWaiverRevocationDTO(apiAutoPolicyWaiverRevocationDTO).Execute()





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
	ownerType := "ownerType_example" // string | Enter the ownerType to specify the scope.
	ownerId := "ownerId_example" // string | Enter the corresponding id for the ownerType specified above.
	apiAutoPolicyWaiverRevocationDTO := *sonatypeiq.NewApiAutoPolicyWaiverRevocationDTO() // ApiAutoPolicyWaiverRevocationDTO | The request JSON can include the fields<ol><li>autoPolicyWaiverId</li><li>hash</li><li>scanId</li></ol>

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoPolicyWaiverRevocationsAPI.AddAutoPolicyWaiverRevocation(context.Background(), ownerType, ownerId).ApiAutoPolicyWaiverRevocationDTO(apiAutoPolicyWaiverRevocationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPolicyWaiverRevocationsAPI.AddAutoPolicyWaiverRevocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAutoPolicyWaiverRevocation`: ApiAutoPolicyWaiverRevocationDTO
	fmt.Fprintf(os.Stdout, "Response from `AutoPolicyWaiverRevocationsAPI.AddAutoPolicyWaiverRevocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAutoPolicyWaiverRevocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAutoPolicyWaiverRevocationDTO** | [**ApiAutoPolicyWaiverRevocationDTO**](ApiAutoPolicyWaiverRevocationDTO.md) | The request JSON can include the fields&lt;ol&gt;&lt;li&gt;autoPolicyWaiverId&lt;/li&gt;&lt;li&gt;hash&lt;/li&gt;&lt;li&gt;scanId&lt;/li&gt;&lt;/ol&gt; | 

### Return type

[**ApiAutoPolicyWaiverRevocationDTO**](ApiAutoPolicyWaiverRevocationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoPolicyWaiverRevocation

> DeleteAutoPolicyWaiverRevocation(ctx, ownerType, ownerId, autoPolicyWaiverRevocationId).Execute()





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
	ownerType := "ownerType_example" // string | Enter the ownerType to specify the scope. A waiver revocation corresponding to the autoPolicyWaiverRevocationId provided and within the scope specified will be deleted.
	ownerId := "ownerId_example" // string | Enter the corresponding id for the ownerType specified above.
	autoPolicyWaiverRevocationId := "autoPolicyWaiverRevocationId_example" // string | Enter the autoPolicyWaiverId to be deleted

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.AutoPolicyWaiverRevocationsAPI.DeleteAutoPolicyWaiverRevocation(context.Background(), ownerType, ownerId, autoPolicyWaiverRevocationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPolicyWaiverRevocationsAPI.DeleteAutoPolicyWaiverRevocation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. A waiver revocation corresponding to the autoPolicyWaiverRevocationId provided and within the scope specified will be deleted. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 
**autoPolicyWaiverRevocationId** | **string** | Enter the autoPolicyWaiverId to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoPolicyWaiverRevocationRequest struct via the builder pattern


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

