# \ComponentSearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchComponent**](ComponentSearchAPI.md#SearchComponent) | **Get** /api/v2/search/component | 



## SearchComponent

> ApiSearchResultsDTOV2 SearchComponent(ctx).StageId(stageId).Hash(hash).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Execute()





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
	stageId := "stageId_example" // string | Specify the evaluation report stage.
	hash := "hash_example" // string | Enter the component hash. (optional)
	componentIdentifier := *sonatypeiq.NewComponentIdentifier() // ComponentIdentifier | Specify the componentIdentifier object containing the format and coordinates. (optional)
	packageUrl := "packageUrl_example" // string | Enter the packageUrl. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentSearchAPI.SearchComponent(context.Background()).StageId(stageId).Hash(hash).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentSearchAPI.SearchComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchComponent`: ApiSearchResultsDTOV2
	fmt.Fprintf(os.Stdout, "Response from `ComponentSearchAPI.SearchComponent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stageId** | **string** | Specify the evaluation report stage. | 
 **hash** | **string** | Enter the component hash. | 
 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) | Specify the componentIdentifier object containing the format and coordinates. | 
 **packageUrl** | **string** | Enter the packageUrl. | 

### Return type

[**ApiSearchResultsDTOV2**](ApiSearchResultsDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

