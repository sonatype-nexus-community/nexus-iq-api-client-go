# \FeatureConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableFeature**](FeatureConfigurationAPI.md#DisableFeature) | **Delete** /api/v2/config/features/{feature} | 
[**EnabledFeature**](FeatureConfigurationAPI.md#EnabledFeature) | **Post** /api/v2/config/features/{feature} | 



## DisableFeature

> DisableFeature(ctx, feature).Execute()





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
	feature := "feature_example" // string | Enter the name of the IQ Server feature to be disabled.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.FeatureConfigurationAPI.DisableFeature(context.Background(), feature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureConfigurationAPI.DisableFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feature** | **string** | Enter the name of the IQ Server feature to be disabled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableFeatureRequest struct via the builder pattern


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


## EnabledFeature

> EnabledFeature(ctx, feature).Execute()





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
	feature := "feature_example" // string | Enter the name of the feature to be enabled.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.FeatureConfigurationAPI.EnabledFeature(context.Background(), feature).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureConfigurationAPI.EnabledFeature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**feature** | **string** | Enter the name of the feature to be enabled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnabledFeatureRequest struct via the builder pattern


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

