# \SourceControlAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSourceControl**](SourceControlAPI.md#AddSourceControl) | **Post** /api/v2/sourceControl/{ownerType}/{internalOwnerId} | 
[**AutomaticRoleAssignment**](SourceControlAPI.md#AutomaticRoleAssignment) | **Post** /api/v2/sourceControl/automaticRoleAssignment/{publicId} | 
[**DeleteSourceControl**](SourceControlAPI.md#DeleteSourceControl) | **Delete** /api/v2/sourceControl/{ownerType}/{internalOwnerId} | 
[**GetSourceControl1**](SourceControlAPI.md#GetSourceControl1) | **Get** /api/v2/sourceControl/{ownerType}/{internalOwnerId} | 
[**UpdateSourceControl**](SourceControlAPI.md#UpdateSourceControl) | **Put** /api/v2/sourceControl/{ownerType}/{internalOwnerId} | 



## AddSourceControl

> ApiSourceControlDTO AddSourceControl(ctx, ownerType, internalOwnerId).ApiSourceControlDTO(apiSourceControlDTO).Execute()



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
	internalOwnerId := "internalOwnerId_example" // string | 
	apiSourceControlDTO := *sonatypeiq.NewApiSourceControlDTO() // ApiSourceControlDTO |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.AddSourceControl(context.Background(), ownerType, internalOwnerId).ApiSourceControlDTO(apiSourceControlDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.AddSourceControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSourceControl`: ApiSourceControlDTO
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.AddSourceControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSourceControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiSourceControlDTO** | [**ApiSourceControlDTO**](ApiSourceControlDTO.md) |  | 

### Return type

[**ApiSourceControlDTO**](ApiSourceControlDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomaticRoleAssignment

> []string AutomaticRoleAssignment(ctx, publicId).Execute()



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
	publicId := "publicId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.AutomaticRoleAssignment(context.Background(), publicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.AutomaticRoleAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomaticRoleAssignment`: []string
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.AutomaticRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomaticRoleAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSourceControl

> DeleteSourceControl(ctx, ownerType, internalOwnerId).Execute()



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
	internalOwnerId := "internalOwnerId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SourceControlAPI.DeleteSourceControl(context.Background(), ownerType, internalOwnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.DeleteSourceControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourceControl1

> ApiSourceControlDTO GetSourceControl1(ctx, ownerType, internalOwnerId).Execute()



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
	internalOwnerId := "internalOwnerId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.GetSourceControl1(context.Background(), ownerType, internalOwnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.GetSourceControl1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceControl1`: ApiSourceControlDTO
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.GetSourceControl1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceControl1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiSourceControlDTO**](ApiSourceControlDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceControl

> ApiSourceControlDTO UpdateSourceControl(ctx, ownerType, internalOwnerId).ApiSourceControlDTO(apiSourceControlDTO).Execute()



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
	internalOwnerId := "internalOwnerId_example" // string | 
	apiSourceControlDTO := *sonatypeiq.NewApiSourceControlDTO() // ApiSourceControlDTO |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.UpdateSourceControl(context.Background(), ownerType, internalOwnerId).ApiSourceControlDTO(apiSourceControlDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.UpdateSourceControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSourceControl`: ApiSourceControlDTO
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.UpdateSourceControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiSourceControlDTO** | [**ApiSourceControlDTO**](ApiSourceControlDTO.md) |  | 

### Return type

[**ApiSourceControlDTO**](ApiSourceControlDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

