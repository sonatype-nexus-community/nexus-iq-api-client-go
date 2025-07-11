# \ConfigJiraAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration1**](ConfigJiraAPI.md#DeleteConfiguration1) | **Delete** /api/v2/config/jira | 
[**GetConfiguration1**](ConfigJiraAPI.md#GetConfiguration1) | **Get** /api/v2/config/jira | 
[**SetConfiguration1**](ConfigJiraAPI.md#SetConfiguration1) | **Put** /api/v2/config/jira | 



## DeleteConfiguration1

> DeleteConfiguration1(ctx).Execute()





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
	r, err := apiClient.ConfigJiraAPI.DeleteConfiguration1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigJiraAPI.DeleteConfiguration1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfiguration1Request struct via the builder pattern


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


## GetConfiguration1

> ApiJiraConfigurationDTO GetConfiguration1(ctx).Execute()





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
	resp, r, err := apiClient.ConfigJiraAPI.GetConfiguration1(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigJiraAPI.GetConfiguration1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration1`: ApiJiraConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigJiraAPI.GetConfiguration1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration1Request struct via the builder pattern


### Return type

[**ApiJiraConfigurationDTO**](ApiJiraConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration1

> SetConfiguration1(ctx).ApiJiraConfigurationDTO(apiJiraConfigurationDTO).Execute()





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
	apiJiraConfigurationDTO := *sonatypeiq.NewApiJiraConfigurationDTO() // ApiJiraConfigurationDTO | Enter the Jira configuration details here. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigJiraAPI.SetConfiguration1(context.Background()).ApiJiraConfigurationDTO(apiJiraConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigJiraAPI.SetConfiguration1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfiguration1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiJiraConfigurationDTO** | [**ApiJiraConfigurationDTO**](ApiJiraConfigurationDTO.md) | Enter the Jira configuration details here. | 

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

