# \PolicyExportAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportPolicies**](PolicyExportAPI.md#ExportPolicies) | **Get** /api/v2/policy/{ownerType}/{ownerId}/export | Export policy configuration



## ExportPolicies

> PolicyExportResult ExportPolicies(ctx, ownerType, ownerId).IncludeInherited(includeInherited).Execute()

Export policy configuration



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
	ownerType := "ownerType_example" // string | Type of owner (organization, application, or repository)
	ownerId := "ownerId_example" // string | Internal ID of the owner
	includeInherited := true // bool | If true, include policies from parent levels in the hierarchy. For repositories, includes policies from the repository and parent organization (RepositoryManager and RepositoryContainer levels are skipped as they do not define policies). For applications, includes policies from the application and parent organization(s). For organizations, includes policies from the organization and any parent organizations. Default: false (direct policies only) (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.PolicyExportAPI.ExportPolicies(context.Background(), ownerType, ownerId).IncludeInherited(includeInherited).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyExportAPI.ExportPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportPolicies`: PolicyExportResult
	fmt.Fprintf(os.Stdout, "Response from `PolicyExportAPI.ExportPolicies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Type of owner (organization, application, or repository) | 
**ownerId** | **string** | Internal ID of the owner | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeInherited** | **bool** | If true, include policies from parent levels in the hierarchy. For repositories, includes policies from the repository and parent organization (RepositoryManager and RepositoryContainer levels are skipped as they do not define policies). For applications, includes policies from the application and parent organization(s). For organizations, includes policies from the organization and any parent organizations. Default: false (direct policies only) | [default to false]

### Return type

[**PolicyExportResult**](PolicyExportResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

