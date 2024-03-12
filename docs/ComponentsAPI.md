# \ComponentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteComponentLabel**](ComponentsAPI.md#DeleteComponentLabel) | **Delete** /api/v2/components/{componentHash}/labels/{labelName}/{ownerType}s/{internalOwnerId} | 
[**GetComponentDetails**](ComponentsAPI.md#GetComponentDetails) | **Post** /api/v2/components/details | 
[**GetComponentVersions**](ComponentsAPI.md#GetComponentVersions) | **Post** /api/v2/components/versions | 
[**GetSuggestedRemediationForComponent**](ComponentsAPI.md#GetSuggestedRemediationForComponent) | **Post** /api/v2/components/remediation/{ownerType}/{ownerId} | 
[**SetComponentLabel**](ComponentsAPI.md#SetComponentLabel) | **Post** /api/v2/components/{componentHash}/labels/{labelName}/{ownerType}s/{internalOwnerId} | 



## DeleteComponentLabel

> DeleteComponentLabel(ctx, ownerType, internalOwnerId, componentHash, labelName).Execute()



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
	componentHash := "componentHash_example" // string | 
	labelName := "labelName_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ComponentsAPI.DeleteComponentLabel(context.Background(), ownerType, internalOwnerId, componentHash, labelName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.DeleteComponentLabel``: %v\n", err)
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
**componentHash** | **string** |  | 
**labelName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComponentLabelRequest struct via the builder pattern


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


## GetComponentDetails

> ApiComponentDetailsResultDTOV2 GetComponentDetails(ctx).ApiComponentDetailsRequestDTOV2(apiComponentDetailsRequestDTOV2).Execute()



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
	apiComponentDetailsRequestDTOV2 := *sonatypeiq.NewApiComponentDetailsRequestDTOV2() // ApiComponentDetailsRequestDTOV2 |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.GetComponentDetails(context.Background()).ApiComponentDetailsRequestDTOV2(apiComponentDetailsRequestDTOV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetComponentDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponentDetails`: ApiComponentDetailsResultDTOV2
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetComponentDetails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiComponentDetailsRequestDTOV2** | [**ApiComponentDetailsRequestDTOV2**](ApiComponentDetailsRequestDTOV2.md) |  | 

### Return type

[**ApiComponentDetailsResultDTOV2**](ApiComponentDetailsResultDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentVersions

> []string GetComponentVersions(ctx).ApiComponentOrPurlIdentifierDTOV2(apiComponentOrPurlIdentifierDTOV2).Execute()



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
	apiComponentOrPurlIdentifierDTOV2 := *sonatypeiq.NewApiComponentOrPurlIdentifierDTOV2() // ApiComponentOrPurlIdentifierDTOV2 |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.GetComponentVersions(context.Background()).ApiComponentOrPurlIdentifierDTOV2(apiComponentOrPurlIdentifierDTOV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetComponentVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponentVersions`: []string
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetComponentVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiComponentOrPurlIdentifierDTOV2** | [**ApiComponentOrPurlIdentifierDTOV2**](ApiComponentOrPurlIdentifierDTOV2.md) |  | 

### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSuggestedRemediationForComponent

> ApiComponentRemediationDTO GetSuggestedRemediationForComponent(ctx, ownerType, ownerId).StageId(stageId).IdentificationSource(identificationSource).ScanId(scanId).ApiComponentDTOV2(apiComponentDTOV2).Execute()



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
	stageId := "stageId_example" // string |  (optional)
	identificationSource := "identificationSource_example" // string |  (optional)
	scanId := "scanId_example" // string |  (optional)
	apiComponentDTOV2 := *sonatypeiq.NewApiComponentDTOV2() // ApiComponentDTOV2 |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.GetSuggestedRemediationForComponent(context.Background(), ownerType, ownerId).StageId(stageId).IdentificationSource(identificationSource).ScanId(scanId).ApiComponentDTOV2(apiComponentDTOV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetSuggestedRemediationForComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSuggestedRemediationForComponent`: ApiComponentRemediationDTO
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetSuggestedRemediationForComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestedRemediationForComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stageId** | **string** |  | 
 **identificationSource** | **string** |  | 
 **scanId** | **string** |  | 
 **apiComponentDTOV2** | [**ApiComponentDTOV2**](ApiComponentDTOV2.md) |  | 

### Return type

[**ApiComponentRemediationDTO**](ApiComponentRemediationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetComponentLabel

> SetComponentLabel(ctx, ownerType, internalOwnerId, componentHash, labelName).Execute()



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
	componentHash := "componentHash_example" // string | 
	labelName := "labelName_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ComponentsAPI.SetComponentLabel(context.Background(), ownerType, internalOwnerId, componentHash, labelName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.SetComponentLabel``: %v\n", err)
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
**componentHash** | **string** |  | 
**labelName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetComponentLabelRequest struct via the builder pattern


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

