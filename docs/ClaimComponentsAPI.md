# \ClaimComponentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](ClaimComponentsAPI.md#Delete) | **Delete** /api/v2/claim/components/{hash} | 
[**Get**](ClaimComponentsAPI.md#Get) | **Get** /api/v2/claim/components/{hash} | 
[**GetAll**](ClaimComponentsAPI.md#GetAll) | **Get** /api/v2/claim/components | 
[**Set**](ClaimComponentsAPI.md#Set) | **Post** /api/v2/claim/components | 



## Delete

> Delete(ctx, hash).Execute()





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
	hash := "hash_example" // string | Enter the SHA1 hash for the component.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ClaimComponentsAPI.Delete(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaimComponentsAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | Enter the SHA1 hash for the component. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


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


## Get

> ApiHashComponentIdentifierDTO Get(ctx, hash).Execute()





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
	hash := "hash_example" // string | The hash of the claimed component.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ClaimComponentsAPI.Get(context.Background(), hash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaimComponentsAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ApiHashComponentIdentifierDTO
	fmt.Fprintf(os.Stdout, "Response from `ClaimComponentsAPI.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hash** | **string** | The hash of the claimed component. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiHashComponentIdentifierDTO**](ApiHashComponentIdentifierDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> ApiHashComponentIdentifiersDTO GetAll(ctx).Execute()





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
	resp, r, err := apiClient.ClaimComponentsAPI.GetAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaimComponentsAPI.GetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAll`: ApiHashComponentIdentifiersDTO
	fmt.Fprintf(os.Stdout, "Response from `ClaimComponentsAPI.GetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


### Return type

[**ApiHashComponentIdentifiersDTO**](ApiHashComponentIdentifiersDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Set

> ApiHashComponentIdentifierDTO Set(ctx).ApiHashComponentIdentifierDTO(apiHashComponentIdentifierDTO).Execute()





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
	apiHashComponentIdentifierDTO := *sonatypeiq.NewApiHashComponentIdentifierDTO() // ApiHashComponentIdentifierDTO | Specify the hash (required), comment (optional), createTime (optional), and the component identifier/package URL (required) with non-null/non-empty format and coordinates,  for the component to be claimed.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ClaimComponentsAPI.Set(context.Background()).ApiHashComponentIdentifierDTO(apiHashComponentIdentifierDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClaimComponentsAPI.Set``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Set`: ApiHashComponentIdentifierDTO
	fmt.Fprintf(os.Stdout, "Response from `ClaimComponentsAPI.Set`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiHashComponentIdentifierDTO** | [**ApiHashComponentIdentifierDTO**](ApiHashComponentIdentifierDTO.md) | Specify the hash (required), comment (optional), createTime (optional), and the component identifier/package URL (required) with non-null/non-empty format and coordinates,  for the component to be claimed. | 

### Return type

[**ApiHashComponentIdentifierDTO**](ApiHashComponentIdentifierDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

