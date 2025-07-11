# \ConfigZscalerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllCategories**](ConfigZscalerAPI.md#DeleteAllCategories) | **Delete** /api/v2/config/zscaler/update | 
[**DeleteCategory**](ConfigZscalerAPI.md#DeleteCategory) | **Delete** /api/v2/config/zscaler/update/{format} | 
[**DeleteConfiguration6**](ConfigZscalerAPI.md#DeleteConfiguration6) | **Delete** /api/v2/config/zscaler | 
[**GetConfiguration6**](ConfigZscalerAPI.md#GetConfiguration6) | **Get** /api/v2/config/zscaler | 
[**GetQuota**](ConfigZscalerAPI.md#GetQuota) | **Get** /api/v2/config/zscaler/zscalerLimits | 
[**SetConfiguration6**](ConfigZscalerAPI.md#SetConfiguration6) | **Put** /api/v2/config/zscaler | 
[**TestConfiguration1**](ConfigZscalerAPI.md#TestConfiguration1) | **Post** /api/v2/config/zscaler/testConfig | 
[**TriggerUpdate**](ConfigZscalerAPI.md#TriggerUpdate) | **Post** /api/v2/config/zscaler/update/{format} | 
[**TriggerUpdateAll**](ConfigZscalerAPI.md#TriggerUpdateAll) | **Post** /api/v2/config/zscaler/update | 



## DeleteAllCategories

> DeleteAllCategories(ctx).Execute()





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
	r, err := apiClient.ConfigZscalerAPI.DeleteAllCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigZscalerAPI.DeleteAllCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllCategoriesRequest struct via the builder pattern


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


## DeleteCategory

> DeleteCategory(ctx, format).Execute()





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
	format := "format_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigZscalerAPI.DeleteCategory(context.Background(), format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigZscalerAPI.DeleteCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCategoryRequest struct via the builder pattern


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


## DeleteConfiguration6

> DeleteConfiguration6(ctx).Execute()





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
	r, err := apiClient.ConfigZscalerAPI.DeleteConfiguration6(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigZscalerAPI.DeleteConfiguration6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfiguration6Request struct via the builder pattern


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


## GetConfiguration6

> ApiZScalerConfigurationDTO GetConfiguration6(ctx).Execute()





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
	resp, r, err := apiClient.ConfigZscalerAPI.GetConfiguration6(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigZscalerAPI.GetConfiguration6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration6`: ApiZScalerConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigZscalerAPI.GetConfiguration6`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration6Request struct via the builder pattern


### Return type

[**ApiZScalerConfigurationDTO**](ApiZScalerConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuota

> ApiZScalerQuotaDTO GetQuota(ctx).Execute()



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
	resp, r, err := apiClient.ConfigZscalerAPI.GetQuota(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigZscalerAPI.GetQuota``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuota`: ApiZScalerQuotaDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigZscalerAPI.GetQuota`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuotaRequest struct via the builder pattern


### Return type

[**ApiZScalerQuotaDTO**](ApiZScalerQuotaDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration6

> string SetConfiguration6(ctx).ApiZScalerConfigurationDTO(apiZScalerConfigurationDTO).Execute()





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
	apiZScalerConfigurationDTO := *sonatypeiq.NewApiZScalerConfigurationDTO() // ApiZScalerConfigurationDTO | Provide one or more values for the following in the JSON payload:<ul><li>`username` - is the username for the Zscaler server.</li><li>`password` - is the password for the Zscaler server.</li><li>`hostname` - is the hostname or IP address of the Zscaler server.</li><li>`apiKey` - is the apiKey for the Zscaler Server.</li><li>`eulaAgreed` - is the agreement to the Sonatype's end user license agreement.</li><li>`mavenFormatEnabled` - is the flag to enable or disable the Maven format for Zscaler.</li><li>`npmFormatEnabled` - is the flag to enable or disable the Npm format for Zscaler.</li><li>`pypiFormatEnabled` - is the flag to enable or disable the Pypi format for Zscaler.</li><li>`nugetFormatEnabled` - is the flag to enable or disable the Nuget format for Zscaler.</li></ul> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigZscalerAPI.SetConfiguration6(context.Background()).ApiZScalerConfigurationDTO(apiZScalerConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigZscalerAPI.SetConfiguration6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetConfiguration6`: string
	fmt.Fprintf(os.Stdout, "Response from `ConfigZscalerAPI.SetConfiguration6`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfiguration6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiZScalerConfigurationDTO** | [**ApiZScalerConfigurationDTO**](ApiZScalerConfigurationDTO.md) | Provide one or more values for the following in the JSON payload:&lt;ul&gt;&lt;li&gt;&#x60;username&#x60; - is the username for the Zscaler server.&lt;/li&gt;&lt;li&gt;&#x60;password&#x60; - is the password for the Zscaler server.&lt;/li&gt;&lt;li&gt;&#x60;hostname&#x60; - is the hostname or IP address of the Zscaler server.&lt;/li&gt;&lt;li&gt;&#x60;apiKey&#x60; - is the apiKey for the Zscaler Server.&lt;/li&gt;&lt;li&gt;&#x60;eulaAgreed&#x60; - is the agreement to the Sonatype&#39;s end user license agreement.&lt;/li&gt;&lt;li&gt;&#x60;mavenFormatEnabled&#x60; - is the flag to enable or disable the Maven format for Zscaler.&lt;/li&gt;&lt;li&gt;&#x60;npmFormatEnabled&#x60; - is the flag to enable or disable the Npm format for Zscaler.&lt;/li&gt;&lt;li&gt;&#x60;pypiFormatEnabled&#x60; - is the flag to enable or disable the Pypi format for Zscaler.&lt;/li&gt;&lt;li&gt;&#x60;nugetFormatEnabled&#x60; - is the flag to enable or disable the Nuget format for Zscaler.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConfiguration1

> TestConfiguration1(ctx).ApiZScalerConfigurationDTO(apiZScalerConfigurationDTO).Execute()





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
	apiZScalerConfigurationDTO := *sonatypeiq.NewApiZScalerConfigurationDTO() // ApiZScalerConfigurationDTO | Provide one or more values for the following in the JSON payload:<ul><li>`username` - is the username for the Zscaler server.</li><li>`password` - is the password for the Zscaler server.</li><li>`hostname` - is the hostname or IP address of the Zscaler server.</li><li>`apiKey` - is the apiKey for the Zscaler Server.</li></ul> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigZscalerAPI.TestConfiguration1(context.Background()).ApiZScalerConfigurationDTO(apiZScalerConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigZscalerAPI.TestConfiguration1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestConfiguration1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiZScalerConfigurationDTO** | [**ApiZScalerConfigurationDTO**](ApiZScalerConfigurationDTO.md) | Provide one or more values for the following in the JSON payload:&lt;ul&gt;&lt;li&gt;&#x60;username&#x60; - is the username for the Zscaler server.&lt;/li&gt;&lt;li&gt;&#x60;password&#x60; - is the password for the Zscaler server.&lt;/li&gt;&lt;li&gt;&#x60;hostname&#x60; - is the hostname or IP address of the Zscaler server.&lt;/li&gt;&lt;li&gt;&#x60;apiKey&#x60; - is the apiKey for the Zscaler Server.&lt;/li&gt;&lt;/ul&gt; | 

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


## TriggerUpdate

> TriggerUpdate(ctx, format).Execute()





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
	format := "format_example" // string | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigZscalerAPI.TriggerUpdate(context.Background(), format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigZscalerAPI.TriggerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerUpdateRequest struct via the builder pattern


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


## TriggerUpdateAll

> TriggerUpdateAll(ctx).Execute()





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
	r, err := apiClient.ConfigZscalerAPI.TriggerUpdateAll(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigZscalerAPI.TriggerUpdateAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTriggerUpdateAllRequest struct via the builder pattern


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

