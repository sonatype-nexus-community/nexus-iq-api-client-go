# \CompositeSourceControlValidatorAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateSourceControlConfig**](CompositeSourceControlValidatorAPI.md#ValidateSourceControlConfig) | **Get** /api/v2/compositeSourceControlConfigValidator/application/{applicationId} | 



## ValidateSourceControlConfig

> ConfigurationValidationResult ValidateSourceControlConfig(ctx, applicationId).Execute()





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
	applicationId := "applicationId_example" // string | Enter the applicationId for which you want to validate the composite source control configuration.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.CompositeSourceControlValidatorAPI.ValidateSourceControlConfig(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CompositeSourceControlValidatorAPI.ValidateSourceControlConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateSourceControlConfig`: ConfigurationValidationResult
	fmt.Fprintf(os.Stdout, "Response from `CompositeSourceControlValidatorAPI.ValidateSourceControlConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | Enter the applicationId for which you want to validate the composite source control configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateSourceControlConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ConfigurationValidationResult**](ConfigurationValidationResult.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

