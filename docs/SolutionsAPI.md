# \SolutionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLicensedSolutions**](SolutionsAPI.md#GetLicensedSolutions) | **Get** /api/v2/solutions/licensed | 



## GetLicensedSolutions

> []ApiLicensedSolutionDTO GetLicensedSolutions(ctx).AllowRelativeUrls(allowRelativeUrls).Execute()



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
	allowRelativeUrls := true // bool |  (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SolutionsAPI.GetLicensedSolutions(context.Background()).AllowRelativeUrls(allowRelativeUrls).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SolutionsAPI.GetLicensedSolutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicensedSolutions`: []ApiLicensedSolutionDTO
	fmt.Fprintf(os.Stdout, "Response from `SolutionsAPI.GetLicensedSolutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLicensedSolutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowRelativeUrls** | **bool** |  | [default to false]

### Return type

[**[]ApiLicensedSolutionDTO**](ApiLicensedSolutionDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

