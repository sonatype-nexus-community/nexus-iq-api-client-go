# \RepositoryComponentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteComponents**](RepositoryComponentsAPI.md#DeleteComponents) | **Delete** /api/v2/repositories/{repositoryManagerInstanceId}/components | 
[**DeleteRepositoryComponents**](RepositoryComponentsAPI.md#DeleteRepositoryComponents) | **Delete** /api/v2/repositories/{repositoryManagerInstanceId}/repositories | 



## DeleteComponents

> DeleteComponents(ctx, repositoryManagerInstanceId).RequestBody(requestBody).Execute()





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
	repositoryManagerInstanceId := "repositoryManagerInstanceId_example" // string | 
	requestBody := []string{"Property_example"} // []string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.RepositoryComponentsAPI.DeleteComponents(context.Background(), repositoryManagerInstanceId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryComponentsAPI.DeleteComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerInstanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRepositoryComponents

> DeleteRepositoryComponents(ctx, repositoryManagerInstanceId).RequestBody(requestBody).Execute()





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
	repositoryManagerInstanceId := "repositoryManagerInstanceId_example" // string | 
	requestBody := []string{"Property_example"} // []string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.RepositoryComponentsAPI.DeleteRepositoryComponents(context.Background(), repositoryManagerInstanceId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoryComponentsAPI.DeleteRepositoryComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerInstanceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

