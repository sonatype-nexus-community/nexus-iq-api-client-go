# \AutoPolicyWaiversAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAutoPolicyWaiver**](AutoPolicyWaiversAPI.md#AddAutoPolicyWaiver) | **Post** /api/v2/autoPolicyWaivers/{ownerType}/{ownerId} | 
[**DeleteAutoPolicyWaiver**](AutoPolicyWaiversAPI.md#DeleteAutoPolicyWaiver) | **Delete** /api/v2/autoPolicyWaivers/{ownerType}/{ownerId}/{autoPolicyWaiverId} | 
[**GetAutoPolicyWaiver**](AutoPolicyWaiversAPI.md#GetAutoPolicyWaiver) | **Get** /api/v2/autoPolicyWaivers/{ownerType}/{ownerId}/{autoPolicyWaiverId} | 
[**GetAutoPolicyWaiverStatus**](AutoPolicyWaiversAPI.md#GetAutoPolicyWaiverStatus) | **Get** /api/v2/autoPolicyWaivers/{ownerType}/{ownerId}/status | 
[**GetAutoPolicyWaivers**](AutoPolicyWaiversAPI.md#GetAutoPolicyWaivers) | **Get** /api/v2/autoPolicyWaivers/{ownerType}/{ownerId} | 
[**UpdateAutoPolicyWaiver**](AutoPolicyWaiversAPI.md#UpdateAutoPolicyWaiver) | **Put** /api/v2/autoPolicyWaivers/{ownerType}/{ownerId}/{autoPolicyWaiverId} | 



## AddAutoPolicyWaiver

> ApiAutoPolicyWaiverDTO AddAutoPolicyWaiver(ctx, ownerType, ownerId).ApiAutoPolicyWaiverDTO(apiAutoPolicyWaiverDTO).Execute()





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
	apiAutoPolicyWaiverDTO := *sonatypeiq.NewApiAutoPolicyWaiverDTO() // ApiAutoPolicyWaiverDTO | The request JSON can include the fields<ol><li>threatLevel</li><li>pathForward</li><li>reachable</li><li>durationInDays</li></ol>

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoPolicyWaiversAPI.AddAutoPolicyWaiver(context.Background(), ownerType, ownerId).ApiAutoPolicyWaiverDTO(apiAutoPolicyWaiverDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPolicyWaiversAPI.AddAutoPolicyWaiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAutoPolicyWaiver`: ApiAutoPolicyWaiverDTO
	fmt.Fprintf(os.Stdout, "Response from `AutoPolicyWaiversAPI.AddAutoPolicyWaiver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. The response will contain the details for waivers within the scope. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAutoPolicyWaiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiAutoPolicyWaiverDTO** | [**ApiAutoPolicyWaiverDTO**](ApiAutoPolicyWaiverDTO.md) | The request JSON can include the fields&lt;ol&gt;&lt;li&gt;threatLevel&lt;/li&gt;&lt;li&gt;pathForward&lt;/li&gt;&lt;li&gt;reachable&lt;/li&gt;&lt;li&gt;durationInDays&lt;/li&gt;&lt;/ol&gt; | 

### Return type

[**ApiAutoPolicyWaiverDTO**](ApiAutoPolicyWaiverDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAutoPolicyWaiver

> DeleteAutoPolicyWaiver(ctx, ownerType, ownerId, autoPolicyWaiverId).Execute()





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
	ownerType := "ownerType_example" // string | Enter the ownerType to specify the scope. A waiver corresponding to the autoPolicyWaiverId provided and within the scope specified will be deleted.
	ownerId := "ownerId_example" // string | Enter the corresponding id for the ownerType specified above.
	autoPolicyWaiverId := "autoPolicyWaiverId_example" // string | Enter the autoPolicyWaiverId to be deleted

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.AutoPolicyWaiversAPI.DeleteAutoPolicyWaiver(context.Background(), ownerType, ownerId, autoPolicyWaiverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPolicyWaiversAPI.DeleteAutoPolicyWaiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. A waiver corresponding to the autoPolicyWaiverId provided and within the scope specified will be deleted. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 
**autoPolicyWaiverId** | **string** | Enter the autoPolicyWaiverId to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoPolicyWaiverRequest struct via the builder pattern


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


## GetAutoPolicyWaiver

> ApiAutoPolicyWaiverDTO GetAutoPolicyWaiver(ctx, ownerType, ownerId, autoPolicyWaiverId).Execute()





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
	autoPolicyWaiverId := "autoPolicyWaiverId_example" // string | Enter the autoPolicyWaiverId for which you want to retrieve the auto policy waiver details.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoPolicyWaiversAPI.GetAutoPolicyWaiver(context.Background(), ownerType, ownerId, autoPolicyWaiverId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPolicyWaiversAPI.GetAutoPolicyWaiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoPolicyWaiver`: ApiAutoPolicyWaiverDTO
	fmt.Fprintf(os.Stdout, "Response from `AutoPolicyWaiversAPI.GetAutoPolicyWaiver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. The response will contain the details for waivers within the scope. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 
**autoPolicyWaiverId** | **string** | Enter the autoPolicyWaiverId for which you want to retrieve the auto policy waiver details. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoPolicyWaiverRequest struct via the builder pattern


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


## GetAutoPolicyWaiverStatus

> ApiAutoPolicyWaiverStatusDTO GetAutoPolicyWaiverStatus(ctx, ownerType, ownerId).Execute()





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
	ownerType := "ownerType_example" // string | Enter the ownerType to specify the scope. The response will contain status details for the active auto policy waiver, if any, that is within the scope specified.
	ownerId := "ownerId_example" // string | Enter the corresponding id for the ownerType specified above.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoPolicyWaiversAPI.GetAutoPolicyWaiverStatus(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPolicyWaiversAPI.GetAutoPolicyWaiverStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoPolicyWaiverStatus`: ApiAutoPolicyWaiverStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `AutoPolicyWaiversAPI.GetAutoPolicyWaiverStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. The response will contain status details for the active auto policy waiver, if any, that is within the scope specified. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoPolicyWaiverStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiAutoPolicyWaiverStatusDTO**](ApiAutoPolicyWaiverStatusDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoPolicyWaivers

> []ApiAutoPolicyWaiverDTO GetAutoPolicyWaivers(ctx, ownerType, ownerId).Execute()





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
	resp, r, err := apiClient.AutoPolicyWaiversAPI.GetAutoPolicyWaivers(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPolicyWaiversAPI.GetAutoPolicyWaivers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoPolicyWaivers`: []ApiAutoPolicyWaiverDTO
	fmt.Fprintf(os.Stdout, "Response from `AutoPolicyWaiversAPI.GetAutoPolicyWaivers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. The response will contain waivers that are within the scope specified. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoPolicyWaiversRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ApiAutoPolicyWaiverDTO**](ApiAutoPolicyWaiverDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoPolicyWaiver

> ApiAutoPolicyWaiverDTO UpdateAutoPolicyWaiver(ctx, ownerType, ownerId, autoPolicyWaiverId).ApiAutoPolicyWaiverDTO(apiAutoPolicyWaiverDTO).Execute()





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
	autoPolicyWaiverId := "autoPolicyWaiverId_example" // string | Enter the autoPolicyWaiverId to be updated.
	apiAutoPolicyWaiverDTO := *sonatypeiq.NewApiAutoPolicyWaiverDTO() // ApiAutoPolicyWaiverDTO | The request JSON can include the fields<ol><li>autoPolicyWaiverId</li><li>threatLevel</li><li>pathForward</li><li>reachable</li><li>durationInDays</li></ol>

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.AutoPolicyWaiversAPI.UpdateAutoPolicyWaiver(context.Background(), ownerType, ownerId, autoPolicyWaiverId).ApiAutoPolicyWaiverDTO(apiAutoPolicyWaiverDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AutoPolicyWaiversAPI.UpdateAutoPolicyWaiver``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAutoPolicyWaiver`: ApiAutoPolicyWaiverDTO
	fmt.Fprintf(os.Stdout, "Response from `AutoPolicyWaiversAPI.UpdateAutoPolicyWaiver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the ownerType to specify the scope. The response will contain the details for waivers within the scope. | 
**ownerId** | **string** | Enter the corresponding id for the ownerType specified above. | 
**autoPolicyWaiverId** | **string** | Enter the autoPolicyWaiverId to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoPolicyWaiverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiAutoPolicyWaiverDTO** | [**ApiAutoPolicyWaiverDTO**](ApiAutoPolicyWaiverDTO.md) | The request JSON can include the fields&lt;ol&gt;&lt;li&gt;autoPolicyWaiverId&lt;/li&gt;&lt;li&gt;threatLevel&lt;/li&gt;&lt;li&gt;pathForward&lt;/li&gt;&lt;li&gt;reachable&lt;/li&gt;&lt;li&gt;durationInDays&lt;/li&gt;&lt;/ol&gt; | 

### Return type

[**ApiAutoPolicyWaiverDTO**](ApiAutoPolicyWaiverDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

