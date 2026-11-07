# \EcosystemConfigurationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddEcosystemConfiguration**](EcosystemConfigurationsAPI.md#AddEcosystemConfiguration) | **Post** /api/v2/ecosystemConfigurations/{ownerType}/{ownerId} | 
[**DeleteEcosystemConfiguration**](EcosystemConfigurationsAPI.md#DeleteEcosystemConfiguration) | **Delete** /api/v2/ecosystemConfigurations/{ownerType}/{ownerId}/{configurationId} | 
[**GetApplicableEcosystemConfigurations**](EcosystemConfigurationsAPI.md#GetApplicableEcosystemConfigurations) | **Get** /api/v2/ecosystemConfigurations/{ownerType}/{ownerId}/applicable | 
[**GetEcosystemConfigurations**](EcosystemConfigurationsAPI.md#GetEcosystemConfigurations) | **Get** /api/v2/ecosystemConfigurations/{ownerType}/{ownerId} | 
[**UpdateEcosystemConfiguration**](EcosystemConfigurationsAPI.md#UpdateEcosystemConfiguration) | **Put** /api/v2/ecosystemConfigurations/{ownerType}/{ownerId} | 



## AddEcosystemConfiguration

> ApiEcosystemConfigurationDTO AddEcosystemConfiguration(ctx, ownerType, ownerId).ApiEcosystemConfigurationDTO(apiEcosystemConfigurationDTO).Execute()





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
	ownerType := "ownerType_example" // string | 
	ownerId := "ownerId_example" // string | 
	apiEcosystemConfigurationDTO := *sonatypeiq.NewApiEcosystemConfigurationDTO() // ApiEcosystemConfigurationDTO | Ecosystem configuration payload. Leave id unset. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.EcosystemConfigurationsAPI.AddEcosystemConfiguration(context.Background(), ownerType, ownerId).ApiEcosystemConfigurationDTO(apiEcosystemConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcosystemConfigurationsAPI.AddEcosystemConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddEcosystemConfiguration`: ApiEcosystemConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `EcosystemConfigurationsAPI.AddEcosystemConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddEcosystemConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiEcosystemConfigurationDTO** | [**ApiEcosystemConfigurationDTO**](ApiEcosystemConfigurationDTO.md) | Ecosystem configuration payload. Leave id unset. | 

### Return type

[**ApiEcosystemConfigurationDTO**](ApiEcosystemConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEcosystemConfiguration

> DeleteEcosystemConfiguration(ctx, ownerType, ownerId, configurationId).Execute()





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
	ownerType := "ownerType_example" // string | 
	ownerId := "ownerId_example" // string | 
	configurationId := "configurationId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.EcosystemConfigurationsAPI.DeleteEcosystemConfiguration(context.Background(), ownerType, ownerId, configurationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcosystemConfigurationsAPI.DeleteEcosystemConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 
**configurationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEcosystemConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicableEcosystemConfigurations

> ApplicableEcosystemConfigurations GetApplicableEcosystemConfigurations(ctx, ownerType, ownerId).Execute()





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
	ownerType := "ownerType_example" // string | 
	ownerId := "ownerId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.EcosystemConfigurationsAPI.GetApplicableEcosystemConfigurations(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcosystemConfigurationsAPI.GetApplicableEcosystemConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicableEcosystemConfigurations`: ApplicableEcosystemConfigurations
	fmt.Fprintf(os.Stdout, "Response from `EcosystemConfigurationsAPI.GetApplicableEcosystemConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableEcosystemConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicableEcosystemConfigurations**](ApplicableEcosystemConfigurations.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEcosystemConfigurations

> []ApiEcosystemConfigurationDTO GetEcosystemConfigurations(ctx, ownerType, ownerId).Execute()





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
	ownerType := "ownerType_example" // string | 
	ownerId := "ownerId_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.EcosystemConfigurationsAPI.GetEcosystemConfigurations(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcosystemConfigurationsAPI.GetEcosystemConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEcosystemConfigurations`: []ApiEcosystemConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `EcosystemConfigurationsAPI.GetEcosystemConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEcosystemConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]ApiEcosystemConfigurationDTO**](ApiEcosystemConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEcosystemConfiguration

> ApiEcosystemConfigurationDTO UpdateEcosystemConfiguration(ctx, ownerType, ownerId).ApiEcosystemConfigurationDTO(apiEcosystemConfigurationDTO).Execute()





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
	ownerType := "ownerType_example" // string | 
	ownerId := "ownerId_example" // string | 
	apiEcosystemConfigurationDTO := *sonatypeiq.NewApiEcosystemConfigurationDTO() // ApiEcosystemConfigurationDTO | Ecosystem configuration payload with id set. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.EcosystemConfigurationsAPI.UpdateEcosystemConfiguration(context.Background(), ownerType, ownerId).ApiEcosystemConfigurationDTO(apiEcosystemConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcosystemConfigurationsAPI.UpdateEcosystemConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEcosystemConfiguration`: ApiEcosystemConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `EcosystemConfigurationsAPI.UpdateEcosystemConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**ownerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEcosystemConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiEcosystemConfigurationDTO** | [**ApiEcosystemConfigurationDTO**](ApiEcosystemConfigurationDTO.md) | Ecosystem configuration payload with id set. | 

### Return type

[**ApiEcosystemConfigurationDTO**](ApiEcosystemConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

