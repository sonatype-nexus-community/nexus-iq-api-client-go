# \ConfigCrowdAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCrowdConfiguration**](ConfigCrowdAPI.md#DeleteCrowdConfiguration) | **Delete** /api/v2/config/crowd | 
[**GetCrowdConfiguration**](ConfigCrowdAPI.md#GetCrowdConfiguration) | **Get** /api/v2/config/crowd | 
[**InsertOrUpdateCrowdConfiguration**](ConfigCrowdAPI.md#InsertOrUpdateCrowdConfiguration) | **Put** /api/v2/config/crowd | 
[**TestCrowdConfiguration**](ConfigCrowdAPI.md#TestCrowdConfiguration) | **Post** /api/v2/config/crowd/test | 



## DeleteCrowdConfiguration

> DeleteCrowdConfiguration(ctx).Execute()



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
	r, err := apiClient.ConfigCrowdAPI.DeleteCrowdConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigCrowdAPI.DeleteCrowdConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCrowdConfigurationRequest struct via the builder pattern


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


## GetCrowdConfiguration

> ApiCrowdConfigurationDTO GetCrowdConfiguration(ctx).Execute()



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
	resp, r, err := apiClient.ConfigCrowdAPI.GetCrowdConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigCrowdAPI.GetCrowdConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrowdConfiguration`: ApiCrowdConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigCrowdAPI.GetCrowdConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrowdConfigurationRequest struct via the builder pattern


### Return type

[**ApiCrowdConfigurationDTO**](ApiCrowdConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertOrUpdateCrowdConfiguration

> InsertOrUpdateCrowdConfiguration(ctx).ApiCrowdConfigurationDTO(apiCrowdConfigurationDTO).Execute()



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
	apiCrowdConfigurationDTO := *sonatypeiq.NewApiCrowdConfigurationDTO() // ApiCrowdConfigurationDTO |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigCrowdAPI.InsertOrUpdateCrowdConfiguration(context.Background()).ApiCrowdConfigurationDTO(apiCrowdConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigCrowdAPI.InsertOrUpdateCrowdConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsertOrUpdateCrowdConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCrowdConfigurationDTO** | [**ApiCrowdConfigurationDTO**](ApiCrowdConfigurationDTO.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestCrowdConfiguration

> ApiStatusDTO TestCrowdConfiguration(ctx).ApiCrowdConfigurationDTO(apiCrowdConfigurationDTO).Execute()



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
	apiCrowdConfigurationDTO := *sonatypeiq.NewApiCrowdConfigurationDTO() // ApiCrowdConfigurationDTO |  (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigCrowdAPI.TestCrowdConfiguration(context.Background()).ApiCrowdConfigurationDTO(apiCrowdConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigCrowdAPI.TestCrowdConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestCrowdConfiguration`: ApiStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigCrowdAPI.TestCrowdConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestCrowdConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiCrowdConfigurationDTO** | [**ApiCrowdConfigurationDTO**](ApiCrowdConfigurationDTO.md) |  | 

### Return type

[**ApiStatusDTO**](ApiStatusDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

