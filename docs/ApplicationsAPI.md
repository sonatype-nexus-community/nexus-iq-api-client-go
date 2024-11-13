# \ApplicationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddApplication**](ApplicationsAPI.md#AddApplication) | **Post** /api/v2/applications | 
[**CloneApplication**](ApplicationsAPI.md#CloneApplication) | **Post** /api/v2/applications/{sourceApplicationId}/clone | 
[**DeleteApplication**](ApplicationsAPI.md#DeleteApplication) | **Delete** /api/v2/applications/{applicationId} | 
[**GetApplication**](ApplicationsAPI.md#GetApplication) | **Get** /api/v2/applications/{applicationId} | 
[**GetApplications**](ApplicationsAPI.md#GetApplications) | **Get** /api/v2/applications | 
[**GetApplicationsByOrganizationId**](ApplicationsAPI.md#GetApplicationsByOrganizationId) | **Get** /api/v2/applications/organization/{organizationId} | 
[**MoveApplication**](ApplicationsAPI.md#MoveApplication) | **Post** /api/v2/applications/{applicationId}/move/organization/{organizationId} | 
[**UpdateApplication**](ApplicationsAPI.md#UpdateApplication) | **Put** /api/v2/applications/{applicationId} | 



## AddApplication

> ApiApplicationDTO AddApplication(ctx).ApiApplicationDTO(apiApplicationDTO).Execute()





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
	apiApplicationDTO := *sonatypeiq.NewApiApplicationDTO() // ApiApplicationDTO | Specify the applicationId, application name and the organizationId under which the application should be created. `contactUserName` corresponds to the 'contact' field in the UI and represents the user name. If LDAP is used for authentication, you can use LDAP usernames.`tagId` is the internal identifier for the Application Category that you want to apply to the application. Use the Application Categories REST API for the available categories and the corresponding tagIds. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.AddApplication(context.Background()).ApiApplicationDTO(apiApplicationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.AddApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddApplication`: ApiApplicationDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.AddApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiApplicationDTO** | [**ApiApplicationDTO**](ApiApplicationDTO.md) | Specify the applicationId, application name and the organizationId under which the application should be created. &#x60;contactUserName&#x60; corresponds to the &#39;contact&#39; field in the UI and represents the user name. If LDAP is used for authentication, you can use LDAP usernames.&#x60;tagId&#x60; is the internal identifier for the Application Category that you want to apply to the application. Use the Application Categories REST API for the available categories and the corresponding tagIds. | 

### Return type

[**ApiApplicationDTO**](ApiApplicationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloneApplication

> ApiApplicationDTO CloneApplication(ctx, sourceApplicationId).ClonedApplicationName(clonedApplicationName).ClonedApplicationPublicId(clonedApplicationPublicId).Execute()





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
	sourceApplicationId := "sourceApplicationId_example" // string | Enter the applicationId for the application to be cloned.
	clonedApplicationName := "clonedApplicationName_example" // string | Enter the application name for the new cloned application. (optional)
	clonedApplicationPublicId := "clonedApplicationPublicId_example" // string | Enter the applicationPublicId for the cloned application. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.CloneApplication(context.Background(), sourceApplicationId).ClonedApplicationName(clonedApplicationName).ClonedApplicationPublicId(clonedApplicationPublicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CloneApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CloneApplication`: ApiApplicationDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.CloneApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceApplicationId** | **string** | Enter the applicationId for the application to be cloned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloneApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clonedApplicationName** | **string** | Enter the application name for the new cloned application. | 
 **clonedApplicationPublicId** | **string** | Enter the applicationPublicId for the cloned application. | 

### Return type

[**ApiApplicationDTO**](ApiApplicationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplication

> DeleteApplication(ctx, applicationId).Execute()





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
	applicationId := "applicationId_example" // string | Enter the applicationId to be deleted.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ApplicationsAPI.DeleteApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.DeleteApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the applicationId to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationRequest struct via the builder pattern


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


## GetApplication

> ApiApplicationDTO GetApplication(ctx, applicationId).Execute()





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
	applicationId := "applicationId_example" // string | Enter the applicationId.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.GetApplication(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplication`: ApiApplicationDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the applicationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiApplicationDTO**](ApiApplicationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> ApiApplicationListDTO GetApplications(ctx).PublicId(publicId).IncludeCategories(includeCategories).Execute()





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
	publicId := []string{"Inner_example"} // []string | Enter the applicationId. (optional)
	includeCategories := true // bool | Set this parameter to `true` to obtain the application tags (application categories) in the response. (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.GetApplications(context.Background()).PublicId(publicId).IncludeCategories(includeCategories).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplications`: ApiApplicationListDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **publicId** | **[]string** | Enter the applicationId. | 
 **includeCategories** | **bool** | Set this parameter to &#x60;true&#x60; to obtain the application tags (application categories) in the response. | [default to false]

### Return type

[**ApiApplicationListDTO**](ApiApplicationListDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationsByOrganizationId

> ApiApplicationListDTO GetApplicationsByOrganizationId(ctx, organizationId).Execute()





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
	organizationId := "organizationId_example" // string | Enter the organizationId.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.GetApplicationsByOrganizationId(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetApplicationsByOrganizationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationsByOrganizationId`: ApiApplicationListDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetApplicationsByOrganizationId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Enter the organizationId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsByOrganizationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiApplicationListDTO**](ApiApplicationListDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveApplication

> ApiMoveApplicationResponseDTOV2 MoveApplication(ctx, applicationId, organizationId).Execute()





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
	applicationId := "applicationId_example" // string | Enter the applicationId of the application to be moved.
	organizationId := "organizationId_example" // string | Enter the organizationId of the destination organization.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.MoveApplication(context.Background(), applicationId, organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.MoveApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveApplication`: ApiMoveApplicationResponseDTOV2
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.MoveApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the applicationId of the application to be moved. | 
**organizationId** | **string** | Enter the organizationId of the destination organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiMoveApplicationResponseDTOV2**](ApiMoveApplicationResponseDTOV2.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplication

> ApiApplicationDTO UpdateApplication(ctx, applicationId).ApiApplicationDTO(apiApplicationDTO).Execute()





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
	applicationId := "applicationId_example" // string | 
	apiApplicationDTO := *sonatypeiq.NewApiApplicationDTO() // ApiApplicationDTO | Specify the applicationId, application name and the organizationId under which  the application exists. `contactUserName` corresponds to the 'contact' field in the UI and represents the user name. If LDAP is used for authentication, you can use LDAP usernames.`tagId` is the internal identifier for the Application Category that you want to apply to the application. . Use the Application Categories REST API for the available categories and the corresponding tagIds. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.UpdateApplication(context.Background(), applicationId).ApiApplicationDTO(apiApplicationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.UpdateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplication`: ApiApplicationDTO
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.UpdateApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiApplicationDTO** | [**ApiApplicationDTO**](ApiApplicationDTO.md) | Specify the applicationId, application name and the organizationId under which  the application exists. &#x60;contactUserName&#x60; corresponds to the &#39;contact&#39; field in the UI and represents the user name. If LDAP is used for authentication, you can use LDAP usernames.&#x60;tagId&#x60; is the internal identifier for the Application Category that you want to apply to the application. . Use the Application Categories REST API for the available categories and the corresponding tagIds. | 

### Return type

[**ApiApplicationDTO**](ApiApplicationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

