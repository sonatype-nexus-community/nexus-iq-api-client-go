# \PolicyWaiverAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPolicyWaiver**](PolicyWaiverAPI.md#AddPolicyWaiver) | **Post** /api/v2/policyWaiver/{policyViolationId}/{ownerType} | 



## AddPolicyWaiver

> AddPolicyWaiver(ctx, policyViolationId, ownerType).Body(body).Execute()



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
	policyViolationId := "policyViolationId_example" // string | 
	ownerType := "ownerType_example" // string | 
	body := "body_example" // string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.PolicyWaiverAPI.AddPolicyWaiver(context.Background(), policyViolationId, ownerType).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiverAPI.AddPolicyWaiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyViolationId** | **string** |  | 
**ownerType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPolicyWaiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

