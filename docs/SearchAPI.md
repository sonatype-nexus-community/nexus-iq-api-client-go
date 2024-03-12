# \SearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSearchIndexAsync**](SearchAPI.md#CreateSearchIndexAsync) | **Post** /api/v2/search/advanced/index | 
[**GetExportResults**](SearchAPI.md#GetExportResults) | **Get** /api/v2/search/advanced/export/csv | 
[**SearchComponent**](SearchAPI.md#SearchComponent) | **Get** /api/v2/search/component | 
[**SearchIndex**](SearchAPI.md#SearchIndex) | **Get** /api/v2/search/advanced | 



## CreateSearchIndexAsync

> CreateSearchIndexAsync(ctx).Execute()



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

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SearchAPI.CreateSearchIndexAsync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.CreateSearchIndexAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSearchIndexAsyncRequest struct via the builder pattern


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


## GetExportResults

> GetExportResults(ctx).Query(query).AllComponents(allComponents).Execute()



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
	query := "query_example" // string |  (optional)
	allComponents := true // bool |  (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SearchAPI.GetExportResults(context.Background()).Query(query).AllComponents(allComponents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.GetExportResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExportResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **allComponents** | **bool** |  | [default to false]

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
	stageId := "stageId_example" // string |  (optional)
	hash := "hash_example" // string |  (optional)
	componentIdentifier := *sonatypeiq.NewComponentIdentifier() // ComponentIdentifier |  (optional)
	packageUrl := "packageUrl_example" // string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchComponent(context.Background()).StageId(stageId).Hash(hash).ComponentIdentifier(componentIdentifier).PackageUrl(packageUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchComponent`: ApiSearchResultsDTOV2
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchComponent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stageId** | **string** |  | 
 **hash** | **string** |  | 
 **componentIdentifier** | [**ComponentIdentifier**](ComponentIdentifier.md) |  | 
 **packageUrl** | **string** |  | 

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


## SearchIndex

> SearchResultDTO SearchIndex(ctx).Query(query).PageSize(pageSize).Page(page).AllComponents(allComponents).Execute()



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
	query := "query_example" // string |  (optional)
	pageSize := int32(56) // int32 |  (optional) (default to 10)
	page := int32(56) // int32 |  (optional)
	allComponents := true // bool |  (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SearchAPI.SearchIndex(context.Background()).Query(query).PageSize(pageSize).Page(page).AllComponents(allComponents).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SearchAPI.SearchIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchIndex`: SearchResultDTO
	fmt.Fprintf(os.Stdout, "Response from `SearchAPI.SearchIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** |  | 
 **pageSize** | **int32** |  | [default to 10]
 **page** | **int32** |  | 
 **allComponents** | **bool** |  | [default to false]

### Return type

[**SearchResultDTO**](SearchResultDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

