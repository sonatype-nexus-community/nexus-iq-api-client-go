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
	ownerType := "ownerType_example" // string | Possible values: application or organization
	internalOwnerId := "internalOwnerId_example" // string | Possible values : applicationId or organizationId
	componentHash := "componentHash_example" // string | Enter the SHA1 hash of the component.
	labelName := "labelName_example" // string | Enter the label name to un-assign from this component.

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
**ownerType** | **string** | Possible values: application or organization | 
**internalOwnerId** | **string** | Possible values : applicationId or organizationId | 
**componentHash** | **string** | Enter the SHA1 hash of the component. | 
**labelName** | **string** | Enter the label name to un-assign from this component. | 

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
- **Accept**: Not defined

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
	apiComponentDetailsRequestDTOV2 := *sonatypeiq.NewApiComponentDetailsRequestDTOV2() // ApiComponentDetailsRequestDTOV2 | You can retrieve component data in any one of the 3 ways via: 1. Component identifier 2. Package URL 3. Hash

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
 **apiComponentDetailsRequestDTOV2** | [**ApiComponentDetailsRequestDTOV2**](ApiComponentDetailsRequestDTOV2.md) | You can retrieve component data in any one of the 3 ways via: 1. Component identifier 2. Package URL 3. Hash | 

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
	apiComponentOrPurlIdentifierDTOV2 := *sonatypeiq.NewApiComponentOrPurlIdentifierDTOV2() // ApiComponentOrPurlIdentifierDTOV2 | Possible values: Component identifier or packageURL (pURL) identifier in the correct format. Use a-name for JavaScript components. (optional)

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
 **apiComponentOrPurlIdentifierDTOV2** | [**ApiComponentOrPurlIdentifierDTOV2**](ApiComponentOrPurlIdentifierDTOV2.md) | Possible values: Component identifier or packageURL (pURL) identifier in the correct format. Use a-name for JavaScript components. | 

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

> GetSuggestedRemediationForComponent(ctx, ownerType, ownerId).StageId(stageId).IdentificationSource(identificationSource).ScanId(scanId).IncludeParentRemediation(includeParentRemediation).ApiComponentDTOV2(apiComponentDTOV2).Execute()





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
	ownerType := "ownerType_example" // string | Possible values: application, organization, repository. 
	ownerId := "ownerId_example" // string | Possible values: applicationId, organizationId or repositoryId.
	stageId := "stageId_example" // string | Enter the stageId to obtain next-non-failing and next-non-failing-with-dependencies remediation types in the response. Possible values are develop, build, stage-release, release and operate. (optional)
	identificationSource := "identificationSource_example" // string | Enter the identification source if you want the remediation result based on third-party scan information (non-Sonatype). The identification source can be obtained from the Component Details Page in the UI. (optional)
	scanId := "scanId_example" // string | Enter the scanId (reportId) if you want the remediation result based on third-party scan information (non-Sonatype). (optional)
	includeParentRemediation := true // bool | Enter true if you want to include parent remediation for transitive dependency in the response based on your application policy scan. (optional) (default to false)
	apiComponentDTOV2 := *sonatypeiq.NewApiComponentDTOV2() // ApiComponentDTOV2 |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ComponentsAPI.GetSuggestedRemediationForComponent(context.Background(), ownerType, ownerId).StageId(stageId).IdentificationSource(identificationSource).ScanId(scanId).IncludeParentRemediation(includeParentRemediation).ApiComponentDTOV2(apiComponentDTOV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetSuggestedRemediationForComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Possible values: application, organization, repository.  | 
**ownerId** | **string** | Possible values: applicationId, organizationId or repositoryId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSuggestedRemediationForComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stageId** | **string** | Enter the stageId to obtain next-non-failing and next-non-failing-with-dependencies remediation types in the response. Possible values are develop, build, stage-release, release and operate. | 
 **identificationSource** | **string** | Enter the identification source if you want the remediation result based on third-party scan information (non-Sonatype). The identification source can be obtained from the Component Details Page in the UI. | 
 **scanId** | **string** | Enter the scanId (reportId) if you want the remediation result based on third-party scan information (non-Sonatype). | 
 **includeParentRemediation** | **bool** | Enter true if you want to include parent remediation for transitive dependency in the response based on your application policy scan. | [default to false]
 **apiComponentDTOV2** | [**ApiComponentDTOV2**](ApiComponentDTOV2.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

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
	ownerType := "ownerType_example" // string | Possible values: application or organization
	internalOwnerId := "internalOwnerId_example" // string | Possible values : applicationId or organizationId
	componentHash := "componentHash_example" // string | Enter the SHA1 hash of the component.
	labelName := "labelName_example" // string | Enter the label name to assign to this component.

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
**ownerType** | **string** | Possible values: application or organization | 
**internalOwnerId** | **string** | Possible values : applicationId or organizationId | 
**componentHash** | **string** | Enter the SHA1 hash of the component. | 
**labelName** | **string** | Enter the label name to assign to this component. | 

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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

