# \RepositoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQuarantinedByPath**](RepositoriesAPI.md#GetQuarantinedByPath) | **Post** /api/v2/repositories/{repositoryManagerInstanceId}/{repositoryPublicId}/components/quarantined/pathnames | 
[**ReleaseQuarantineWithoutReEval**](RepositoriesAPI.md#ReleaseQuarantineWithoutReEval) | **Post** /api/v2/repositories/quarantine/{quarantineId}/release | 



## GetQuarantinedByPath

> ApiRepositoryPathResponseDTO GetQuarantinedByPath(ctx, repositoryManagerInstanceId, repositoryPublicId).RequestBody(requestBody).Execute()





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
	repositoryManagerInstanceId := "repositoryManagerInstanceId_example" // string | Enter the repository manager instance ID.
	repositoryPublicId := "repositoryPublicId_example" // string | Enter the repository public ID.
	requestBody := []string{"Property_example"} // []string | Specify the pathnames.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.GetQuarantinedByPath(context.Background(), repositoryManagerInstanceId, repositoryPublicId).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.GetQuarantinedByPath``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuarantinedByPath`: ApiRepositoryPathResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.GetQuarantinedByPath`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryManagerInstanceId** | **string** | Enter the repository manager instance ID. | 
**repositoryPublicId** | **string** | Enter the repository public ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuarantinedByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]string** | Specify the pathnames. | 

### Return type

[**ApiRepositoryPathResponseDTO**](ApiRepositoryPathResponseDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReleaseQuarantineWithoutReEval

> ApiComponentReleasedFromQuarantineDTO ReleaseQuarantineWithoutReEval(ctx, quarantineId).Body(body).Execute()





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
	quarantineId := "quarantineId_example" // string | Enter the component `quarantineId`.
	body := "body_example" // string | Enter a waiver comment for releasing the component from quarantine.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.ReleaseQuarantineWithoutReEval(context.Background(), quarantineId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.ReleaseQuarantineWithoutReEval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReleaseQuarantineWithoutReEval`: ApiComponentReleasedFromQuarantineDTO
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.ReleaseQuarantineWithoutReEval`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quarantineId** | **string** | Enter the component &#x60;quarantineId&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReleaseQuarantineWithoutReEvalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | Enter a waiver comment for releasing the component from quarantine. | 

### Return type

[**ApiComponentReleasedFromQuarantineDTO**](ApiComponentReleasedFromQuarantineDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

