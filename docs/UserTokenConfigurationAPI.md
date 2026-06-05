# \UserTokenConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration8**](UserTokenConfigurationAPI.md#GetConfiguration8) | **Get** /api/v2/config/userTokens | 
[**ResetConfiguration**](UserTokenConfigurationAPI.md#ResetConfiguration) | **Delete** /api/v2/config/userTokens | 
[**UpdateConfiguration**](UserTokenConfigurationAPI.md#UpdateConfiguration) | **Put** /api/v2/config/userTokens | 



## GetConfiguration8

> ApiUserTokenConfigurationDTO GetConfiguration8(ctx).Execute()





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
	resp, r, err := apiClient.UserTokenConfigurationAPI.GetConfiguration8(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTokenConfigurationAPI.GetConfiguration8``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration8`: ApiUserTokenConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `UserTokenConfigurationAPI.GetConfiguration8`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration8Request struct via the builder pattern


### Return type

[**ApiUserTokenConfigurationDTO**](ApiUserTokenConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetConfiguration

> ApiUserTokenConfigurationDTO ResetConfiguration(ctx).Property(property).Execute()





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
	property := []string{"Inner_example"} // []string |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTokenConfigurationAPI.ResetConfiguration(context.Background()).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTokenConfigurationAPI.ResetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResetConfiguration`: ApiUserTokenConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `UserTokenConfigurationAPI.ResetConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **property** | **[]string** |  | 

### Return type

[**ApiUserTokenConfigurationDTO**](ApiUserTokenConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConfiguration

> ApiUserTokenConfigurationDTO UpdateConfiguration(ctx).ApiUserTokenConfigurationDTO(apiUserTokenConfigurationDTO).Execute()





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
	apiUserTokenConfigurationDTO := *sonatypeiq.NewApiUserTokenConfigurationDTO() // ApiUserTokenConfigurationDTO |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.UserTokenConfigurationAPI.UpdateConfiguration(context.Background()).ApiUserTokenConfigurationDTO(apiUserTokenConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserTokenConfigurationAPI.UpdateConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateConfiguration`: ApiUserTokenConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `UserTokenConfigurationAPI.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiUserTokenConfigurationDTO** | [**ApiUserTokenConfigurationDTO**](ApiUserTokenConfigurationDTO.md) |  | 

### Return type

[**ApiUserTokenConfigurationDTO**](ApiUserTokenConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

