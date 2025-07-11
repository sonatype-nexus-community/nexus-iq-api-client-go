# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Add**](UsersAPI.md#Add) | **Post** /api/v2/users | 
[**Delete1**](UsersAPI.md#Delete1) | **Delete** /api/v2/users/{username} | 
[**Get1**](UsersAPI.md#Get1) | **Get** /api/v2/users/{username} | 
[**GetAll2**](UsersAPI.md#GetAll2) | **Get** /api/v2/users | 
[**Update**](UsersAPI.md#Update) | **Put** /api/v2/users/{username} | 



## Add

> Add(ctx).ApiUserDTO(apiUserDTO).Execute()





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
	apiUserDTO := *sonatypeiq.NewApiUserDTO() // ApiUserDTO | Specify the user details for the new user to be created. All fields except `realm` are required. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.Add(context.Background()).ApiUserDTO(apiUserDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.Add``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiUserDTO** | [**ApiUserDTO**](ApiUserDTO.md) | Specify the user details for the new user to be created. All fields except &#x60;realm&#x60; are required. | 

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


## Delete1

> Delete1(ctx, username).Realm(realm).Execute()





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
	username := "username_example" // string | Enter the username to be deleted.
	realm := "realm_example" // string | Enter the `realm`. Allowed values are `Internal`,`OAUTH2`, and `SAML`. (optional) (default to "Internal")

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.Delete1(context.Background(), username).Realm(realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.Delete1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Enter the username to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **realm** | **string** | Enter the &#x60;realm&#x60;. Allowed values are &#x60;Internal&#x60;,&#x60;OAUTH2&#x60;, and &#x60;SAML&#x60;. | [default to &quot;Internal&quot;]

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


## Get1

> ApiUserDTO Get1(ctx, username).Realm(realm).Execute()





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
	username := "username_example" // string | Enter the username.
	realm := "realm_example" // string | Enter the `realm`. Allowed values are `Internal`,`OAUTH2`, and `SAML`. (optional) (default to "Internal")

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.Get1(context.Background(), username).Realm(realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.Get1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get1`: ApiUserDTO
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.Get1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Enter the username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **realm** | **string** | Enter the &#x60;realm&#x60;. Allowed values are &#x60;Internal&#x60;,&#x60;OAUTH2&#x60;, and &#x60;SAML&#x60;. | [default to &quot;Internal&quot;]

### Return type

[**ApiUserDTO**](ApiUserDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll2

> ApiUserListDTO GetAll2(ctx).Realm(realm).Execute()





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
	realm := "realm_example" // string | Enter the `realm`. Allowed values are `Internal`,`OAUTH2`, and `SAML`. (optional) (default to "Internal")

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetAll2(context.Background()).Realm(realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetAll2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAll2`: ApiUserListDTO
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetAll2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAll2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **realm** | **string** | Enter the &#x60;realm&#x60;. Allowed values are &#x60;Internal&#x60;,&#x60;OAUTH2&#x60;, and &#x60;SAML&#x60;. | [default to &quot;Internal&quot;]

### Return type

[**ApiUserListDTO**](ApiUserListDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> ApiUserDTO Update(ctx, username).ApiUserDTO(apiUserDTO).Execute()





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
	username := "username_example" // string | Enter the username.
	apiUserDTO := *sonatypeiq.NewApiUserDTO() // ApiUserDTO | Specify the user details to be updated. Any unspecified field will remain unchanged. Username, password, and realm cannot be updated. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.Update(context.Background(), username).ApiUserDTO(apiUserDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: ApiUserDTO
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | Enter the username. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiUserDTO** | [**ApiUserDTO**](ApiUserDTO.md) | Specify the user details to be updated. Any unspecified field will remain unchanged. Username, password, and realm cannot be updated. | 

### Return type

[**ApiUserDTO**](ApiUserDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

