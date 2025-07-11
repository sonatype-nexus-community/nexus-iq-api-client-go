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
- **Accept**: Not defined

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
	apiCrowdConfigurationDTO := *sonatypeiq.NewApiCrowdConfigurationDTO() // ApiCrowdConfigurationDTO | The request JSON should include the `serverUrl`, `applicationName`, and the `applicationPassword` which will be used for authentication against the Atlassian Crowd Server.  If updating the `serverUrl`, the `applicationPassword` field is required. (optional)

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
 **apiCrowdConfigurationDTO** | [**ApiCrowdConfigurationDTO**](ApiCrowdConfigurationDTO.md) | The request JSON should include the &#x60;serverUrl&#x60;, &#x60;applicationName&#x60;, and the &#x60;applicationPassword&#x60; which will be used for authentication against the Atlassian Crowd Server.  If updating the &#x60;serverUrl&#x60;, the &#x60;applicationPassword&#x60; field is required. | 

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
	apiCrowdConfigurationDTO := *sonatypeiq.NewApiCrowdConfigurationDTO() // ApiCrowdConfigurationDTO | To test an existing configuration, the request body is not required.  To test a new configuration, provide the `serverURl`, `applicationName`, and `applicationPassword` for the configuration. (optional)

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
 **apiCrowdConfigurationDTO** | [**ApiCrowdConfigurationDTO**](ApiCrowdConfigurationDTO.md) | To test an existing configuration, the request body is not required.  To test a new configuration, provide the &#x60;serverURl&#x60;, &#x60;applicationName&#x60;, and &#x60;applicationPassword&#x60; for the configuration. | 

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

