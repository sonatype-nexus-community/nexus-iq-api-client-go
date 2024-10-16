# \ApplicationCategoriesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTag**](ApplicationCategoriesAPI.md#AddTag) | **Post** /api/v2/applicationCategories/organization/{organizationId} | 
[**DeleteTag**](ApplicationCategoriesAPI.md#DeleteTag) | **Delete** /api/v2/applicationCategories/organization/{organizationId}/{tagId} | 
[**GetApplicableTags**](ApplicationCategoriesAPI.md#GetApplicableTags) | **Get** /api/v2/applicationCategories/organization/{organizationId}/applicable | 
[**GetApplicableTagsByApplicationPublicId**](ApplicationCategoriesAPI.md#GetApplicableTagsByApplicationPublicId) | **Get** /api/v2/applicationCategories/application/{applicationPublicId}/applicable | 
[**GetApplicationApplicableTags**](ApplicationCategoriesAPI.md#GetApplicationApplicableTags) | **Get** /api/v2/applicationCategories/application/{applicationPublicId} | 
[**GetAppliedPolicyTags**](ApplicationCategoriesAPI.md#GetAppliedPolicyTags) | **Get** /api/v2/applicationCategories/organization/{organizationId}/policy | 
[**GetAppliedTags**](ApplicationCategoriesAPI.md#GetAppliedTags) | **Get** /api/v2/applicationCategories/organization/{organizationId}/applied | 
[**GetTags**](ApplicationCategoriesAPI.md#GetTags) | **Get** /api/v2/applicationCategories/organization/{organizationId} | 
[**GetTagsUsedByApplications**](ApplicationCategoriesAPI.md#GetTagsUsedByApplications) | **Get** /api/v2/applicationCategories/application | 
[**UpdateTag**](ApplicationCategoriesAPI.md#UpdateTag) | **Put** /api/v2/applicationCategories/organization/{organizationId} | 



## AddTag

> ApiApplicationCategoryDTO AddTag(ctx, organizationId).ApiApplicationCategoryDTO(apiApplicationCategoryDTO).Execute()





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
	organizationId := "organizationId_example" // string | The organizationId assigned by IQ Server, for which you want to create the application category.
	apiApplicationCategoryDTO := *sonatypeiq.NewApiApplicationCategoryDTO() // ApiApplicationCategoryDTO | Specify the the name, description and color for the new application category to be  created. The application category id is not required to create a new application category  and should not be included.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCategoriesAPI.AddTag(context.Background(), organizationId).ApiApplicationCategoryDTO(apiApplicationCategoryDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.AddTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTag`: ApiApplicationCategoryDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.AddTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The organizationId assigned by IQ Server, for which you want to create the application category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiApplicationCategoryDTO** | [**ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md) | Specify the the name, description and color for the new application category to be  created. The application category id is not required to create a new application category  and should not be included. | 

### Return type

[**ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteTag(ctx, organizationId, tagId).Execute()





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
	organizationId := "organizationId_example" // string | The organizationId assigned by IQ Server, corresponding to the application category tag you want to delete.
	tagId := "tagId_example" // string | The application category ID assigned by IQ Server, to be deleted.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ApplicationCategoriesAPI.DeleteTag(context.Background(), organizationId, tagId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.DeleteTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The organizationId assigned by IQ Server, corresponding to the application category tag you want to delete. | 
**tagId** | **string** | The application category ID assigned by IQ Server, to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


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


## GetApplicableTags

> ApplicableTagsDTO GetApplicableTags(ctx, organizationId).Execute()





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
	organizationId := "organizationId_example" // string | The organizationId assigned by IQ Server, for which you want to retrieve the applicable tags or application categories.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCategoriesAPI.GetApplicableTags(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetApplicableTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicableTags`: ApplicableTagsDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetApplicableTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The organizationId assigned by IQ Server, for which you want to retrieve the applicable tags or application categories. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicableTagsDTO**](ApplicableTagsDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicableTagsByApplicationPublicId

> []ApiApplicationCategoryDTO GetApplicableTagsByApplicationPublicId(ctx, applicationPublicId).Execute()





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
	applicationPublicId := "applicationPublicId_example" // string | Provide the application public ID assigned by IQ Server.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCategoriesAPI.GetApplicableTagsByApplicationPublicId(context.Background(), applicationPublicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetApplicableTagsByApplicationPublicId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicableTagsByApplicationPublicId`: []ApiApplicationCategoryDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetApplicableTagsByApplicationPublicId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** | Provide the application public ID assigned by IQ Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableTagsByApplicationPublicIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationApplicableTags

> ApplicableTagsDTO GetApplicationApplicableTags(ctx, applicationPublicId).Execute()





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
	applicationPublicId := "applicationPublicId_example" // string | The application public ID 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCategoriesAPI.GetApplicationApplicableTags(context.Background(), applicationPublicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetApplicationApplicableTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationApplicableTags`: ApplicableTagsDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetApplicationApplicableTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationPublicId** | **string** | The application public ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationApplicableTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicableTagsDTO**](ApplicableTagsDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppliedPolicyTags

> []PolicyTag GetAppliedPolicyTags(ctx, organizationId).Execute()





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
	organizationId := "organizationId_example" // string | The organizationId assigned by IQ Server.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCategoriesAPI.GetAppliedPolicyTags(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetAppliedPolicyTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppliedPolicyTags`: []PolicyTag
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetAppliedPolicyTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The organizationId assigned by IQ Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppliedPolicyTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]PolicyTag**](PolicyTag.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAppliedTags

> AppliedTagsDTO GetAppliedTags(ctx, organizationId).Execute()





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
	organizationId := "organizationId_example" // string | The organizationId assigned by IQ Server.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCategoriesAPI.GetAppliedTags(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetAppliedTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppliedTags`: AppliedTagsDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetAppliedTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The organizationId assigned by IQ Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppliedTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AppliedTagsDTO**](AppliedTagsDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTags

> []ApiApplicationCategoryDTO GetTags(ctx, organizationId).Execute()





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
	organizationId := "organizationId_example" // string | The organizationId assigned by IQ Server.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCategoriesAPI.GetTags(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTags`: []ApiApplicationCategoryDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The organizationId assigned by IQ Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTagsUsedByApplications

> []ApiApplicationCategoryDTO GetTagsUsedByApplications(ctx).Execute()





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
	resp, r, err := apiClient.ApplicationCategoriesAPI.GetTagsUsedByApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.GetTagsUsedByApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagsUsedByApplications`: []ApiApplicationCategoryDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.GetTagsUsedByApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagsUsedByApplicationsRequest struct via the builder pattern


### Return type

[**[]ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> ApiApplicationCategoryDTO UpdateTag(ctx, organizationId).ApiApplicationCategoryDTO(apiApplicationCategoryDTO).Execute()





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
	organizationId := "organizationId_example" // string | The organizationId assigned by IQ Server.
	apiApplicationCategoryDTO := *sonatypeiq.NewApiApplicationCategoryDTO() // ApiApplicationCategoryDTO | Specify the id (application category id) and id of the organization that owns this  application category, to update the name, description and color.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationCategoriesAPI.UpdateTag(context.Background(), organizationId).ApiApplicationCategoryDTO(apiApplicationCategoryDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCategoriesAPI.UpdateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTag`: ApiApplicationCategoryDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationCategoriesAPI.UpdateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | The organizationId assigned by IQ Server. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiApplicationCategoryDTO** | [**ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md) | Specify the id (application category id) and id of the organization that owns this  application category, to update the name, description and color. | 

### Return type

[**ApiApplicationCategoryDTO**](ApiApplicationCategoryDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

