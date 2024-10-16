# \AdvancedSearchAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSearchIndexAsync**](AdvancedSearchAPI.md#CreateSearchIndexAsync) | **Post** /api/v2/search/advanced/index | 
[**GetExportResults**](AdvancedSearchAPI.md#GetExportResults) | **Get** /api/v2/search/advanced/export/csv | 
[**SearchIndex**](AdvancedSearchAPI.md#SearchIndex) | **Get** /api/v2/search/advanced | 



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
	r, err := apiClient.AdvancedSearchAPI.CreateSearchIndexAsync(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchAPI.CreateSearchIndexAsync``: %v\n", err)
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
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExportResults

> GetExportResults(ctx).Query(query).PageSize(pageSize).Page(page).AllComponents(allComponents).Mode(mode).Execute()





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
	query := "query_example" // string | A well-formed search query.
	pageSize := int32(56) // int32 | Enter the no. of results that should be visible per page, unset gives all results (optional)
	page := int32(56) // int32 | Enter the page no. for the page containing results (optional)
	allComponents := true // bool | Set to `true` to retrieve results that include components with no violations. (optional) (default to false)
	mode := "mode_example" // string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.AdvancedSearchAPI.GetExportResults(context.Background()).Query(query).PageSize(pageSize).Page(page).AllComponents(allComponents).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchAPI.GetExportResults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExportResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | A well-formed search query. | 
 **pageSize** | **int32** | Enter the no. of results that should be visible per page, unset gives all results | 
 **page** | **int32** | Enter the page no. for the page containing results | 
 **allComponents** | **bool** | Set to &#x60;true&#x60; to retrieve results that include components with no violations. | [default to false]
 **mode** | **string** |  | 

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


## SearchIndex

> SearchResultDTO SearchIndex(ctx).Query(query).PageSize(pageSize).Page(page).AllComponents(allComponents).Mode(mode).Execute()





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
	query := "query_example" // string | Enter your search query here (optional)
	pageSize := int32(56) // int32 | Enter the no. of results that should be visible per page (optional) (default to 10)
	page := int32(56) // int32 | Enter the page no. for the page containing results (optional)
	allComponents := true // bool | Set to `true` to retrieve results that include components with no violations (optional) (default to false)
	mode := "mode_example" // string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.AdvancedSearchAPI.SearchIndex(context.Background()).Query(query).PageSize(pageSize).Page(page).AllComponents(allComponents).Mode(mode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdvancedSearchAPI.SearchIndex``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchIndex`: SearchResultDTO
	fmt.Fprintf(os.Stdout, "Response from `AdvancedSearchAPI.SearchIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | Enter your search query here | 
 **pageSize** | **int32** | Enter the no. of results that should be visible per page | [default to 10]
 **page** | **int32** | Enter the page no. for the page containing results | 
 **allComponents** | **bool** | Set to &#x60;true&#x60; to retrieve results that include components with no violations | [default to false]
 **mode** | **string** |  | 

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

