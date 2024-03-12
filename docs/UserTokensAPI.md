# \UserTokensAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserToken**](UserTokensAPI.md#CreateUserToken) | **Post** /api/v2/userTokens/currentUser | 
[**DeleteCurrentUserToken**](UserTokensAPI.md#DeleteCurrentUserToken) | **Delete** /api/v2/userTokens/currentUser | 
[**DeleteUserTokenByUserCode**](UserTokensAPI.md#DeleteUserTokenByUserCode) | **Delete** /api/v2/userTokens/userCode/{userCode} | 
[**GetUserTokenByUsernameAndRealmId**](UserTokensAPI.md#GetUserTokenByUsernameAndRealmId) | **Get** /api/v2/userTokens/{username} | 
[**GetUserTokenExistsForCurrentUser**](UserTokensAPI.md#GetUserTokenExistsForCurrentUser) | **Get** /api/v2/userTokens/currentUser/hasToken | 
[**GetUserTokensByCreatedBetweenAndRealmId**](UserTokensAPI.md#GetUserTokensByCreatedBetweenAndRealmId) | **Get** /api/v2/userTokens | 
[**PurgeUserTokens**](UserTokensAPI.md#PurgeUserTokens) | **Delete** /api/v2/userTokens/purge | 



## CreateUserToken

> ApiUserTokenDTO CreateUserToken(ctx).Execute()



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
	resp, r, err := apiClient.UserTokensAPI.CreateUserToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTokensAPI.CreateUserToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserToken`: ApiUserTokenDTO
	fmt.Fprintf(os.Stdout, "Response from `UserTokensAPI.CreateUserToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserTokenRequest struct via the builder pattern


### Return type

[**ApiUserTokenDTO**](ApiUserTokenDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCurrentUserToken

> DeleteCurrentUserToken(ctx).Execute()



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
	r, err := apiClient.UserTokensAPI.DeleteCurrentUserToken(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTokensAPI.DeleteCurrentUserToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCurrentUserTokenRequest struct via the builder pattern


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


## DeleteUserTokenByUserCode

> DeleteUserTokenByUserCode(ctx, userCode).Execute()



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
	userCode := "userCode_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.UserTokensAPI.DeleteUserTokenByUserCode(context.Background(), userCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTokensAPI.DeleteUserTokenByUserCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userCode** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserTokenByUserCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetUserTokenByUsernameAndRealmId

> ApiUserTokenDTO GetUserTokenByUsernameAndRealmId(ctx, username).Realm(realm).Execute()



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
	username := "username_example" // string | 
	realm := "realm_example" // string |  (optional) (default to "Internal")

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTokensAPI.GetUserTokenByUsernameAndRealmId(context.Background(), username).Realm(realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTokensAPI.GetUserTokenByUsernameAndRealmId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTokenByUsernameAndRealmId`: ApiUserTokenDTO
	fmt.Fprintf(os.Stdout, "Response from `UserTokensAPI.GetUserTokenByUsernameAndRealmId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTokenByUsernameAndRealmIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **realm** | **string** |  | [default to &quot;Internal&quot;]

### Return type

[**ApiUserTokenDTO**](ApiUserTokenDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTokenExistsForCurrentUser

> ApiUserTokenExistsDTO GetUserTokenExistsForCurrentUser(ctx).Execute()



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
	resp, r, err := apiClient.UserTokensAPI.GetUserTokenExistsForCurrentUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTokensAPI.GetUserTokenExistsForCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTokenExistsForCurrentUser`: ApiUserTokenExistsDTO
	fmt.Fprintf(os.Stdout, "Response from `UserTokensAPI.GetUserTokenExistsForCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTokenExistsForCurrentUserRequest struct via the builder pattern


### Return type

[**ApiUserTokenExistsDTO**](ApiUserTokenExistsDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTokensByCreatedBetweenAndRealmId

> []ApiUserTokenDTO GetUserTokensByCreatedBetweenAndRealmId(ctx).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Realm(realm).Execute()



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
	createdAfter := "createdAfter_example" // string |  (optional)
	createdBefore := "createdBefore_example" // string |  (optional)
	realm := "realm_example" // string |  (optional) (default to "Internal")

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTokensAPI.GetUserTokensByCreatedBetweenAndRealmId(context.Background()).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Realm(realm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTokensAPI.GetUserTokensByCreatedBetweenAndRealmId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserTokensByCreatedBetweenAndRealmId`: []ApiUserTokenDTO
	fmt.Fprintf(os.Stdout, "Response from `UserTokensAPI.GetUserTokensByCreatedBetweenAndRealmId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTokensByCreatedBetweenAndRealmIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createdAfter** | **string** |  | 
 **createdBefore** | **string** |  | 
 **realm** | **string** |  | [default to &quot;Internal&quot;]

### Return type

[**[]ApiUserTokenDTO**](ApiUserTokenDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PurgeUserTokens

> PurgeUserTokens(ctx).Execute()



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
	r, err := apiClient.UserTokensAPI.PurgeUserTokens(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTokensAPI.PurgeUserTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeUserTokensRequest struct via the builder pattern


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

