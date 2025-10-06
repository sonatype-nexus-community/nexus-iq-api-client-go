# \OrganizationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOrganization**](OrganizationsAPI.md#AddOrganization) | **Post** /api/v2/organizations | 
[**DeleteOrganization**](OrganizationsAPI.md#DeleteOrganization) | **Delete** /api/v2/organizations/{organizationId} | 
[**GetOrganization**](OrganizationsAPI.md#GetOrganization) | **Get** /api/v2/organizations/{organizationId} | 
[**GetOrganizations**](OrganizationsAPI.md#GetOrganizations) | **Get** /api/v2/organizations | 
[**MoveOrganization**](OrganizationsAPI.md#MoveOrganization) | **Put** /api/v2/organizations/{organizationId}/move/destination/{destinationId} | 



## AddOrganization

> ApiOrganizationDTO AddOrganization(ctx).ApiOrganizationDTO(apiOrganizationDTO).Execute()





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
	apiOrganizationDTO := *sonatypeiq.NewApiOrganizationDTO() // ApiOrganizationDTO | The request JSON should include the name of the organization (should be unique), name of the parent organization and tags containing additional organization details. If the parent organization is not specified, this organization will be created under the root organization. Tags represent identifying characteristics of an application. They are created at the organization level and then applied to applications under the organization. The tags can be used to decide which applications will be evaluated against a selected policy. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.AddOrganization(context.Background()).ApiOrganizationDTO(apiOrganizationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.AddOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddOrganization`: ApiOrganizationDTO
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.AddOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiOrganizationDTO** | [**ApiOrganizationDTO**](ApiOrganizationDTO.md) | The request JSON should include the name of the organization (should be unique), name of the parent organization and tags containing additional organization details. If the parent organization is not specified, this organization will be created under the root organization. Tags represent identifying characteristics of an application. They are created at the organization level and then applied to applications under the organization. The tags can be used to decide which applications will be evaluated against a selected policy. | 

### Return type

[**ApiOrganizationDTO**](ApiOrganizationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> DeleteOrganization(ctx, organizationId).Execute()





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
	organizationId := "organizationId_example" // string | Enter the organization id to be deleted.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.OrganizationsAPI.DeleteOrganization(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.DeleteOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Enter the organization id to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationRequest struct via the builder pattern


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


## GetOrganization

> ApiOrganizationDTO GetOrganization(ctx, organizationId).Execute()





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
	organizationId := "organizationId_example" // string | Enter the organization id.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganization(context.Background(), organizationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganization`: ApiOrganizationDTO
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Enter the organization id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiOrganizationDTO**](ApiOrganizationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizations

> ApiOrganizationListDTO GetOrganizations(ctx).OrganizationName(organizationName).Execute()





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
	organizationName := []string{"Inner_example"} // []string | Enter the organization names. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.GetOrganizations(context.Background()).OrganizationName(organizationName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.GetOrganizations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrganizations`: ApiOrganizationListDTO
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.GetOrganizations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrganizationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationName** | **[]string** | Enter the organization names. | 

### Return type

[**ApiOrganizationListDTO**](ApiOrganizationListDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveOrganization

> MoveOrganizationResponseDTO MoveOrganization(ctx, organizationId, destinationId).FailEarlyOnError(failEarlyOnError).Execute()





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
	organizationId := "organizationId_example" // string | Enter the id for the organization to be moved under the new parent.
	destinationId := "destinationId_example" // string | Enter the id for the new parent organization.
	failEarlyOnError := true // bool |  (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationsAPI.MoveOrganization(context.Background(), organizationId, destinationId).FailEarlyOnError(failEarlyOnError).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationsAPI.MoveOrganization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveOrganization`: MoveOrganizationResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `OrganizationsAPI.MoveOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**organizationId** | **string** | Enter the id for the organization to be moved under the new parent. | 
**destinationId** | **string** | Enter the id for the new parent organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **failEarlyOnError** | **bool** |  | [default to false]

### Return type

[**MoveOrganizationResponseDTO**](MoveOrganizationResponseDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

