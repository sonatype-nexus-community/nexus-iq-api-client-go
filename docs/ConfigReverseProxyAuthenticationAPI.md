# \ConfigReverseProxyAuthenticationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration4**](ConfigReverseProxyAuthenticationAPI.md#DeleteConfiguration4) | **Delete** /api/v2/config/reverseProxyAuthentication | 
[**GetConfiguration4**](ConfigReverseProxyAuthenticationAPI.md#GetConfiguration4) | **Get** /api/v2/config/reverseProxyAuthentication | 
[**SetConfiguration4**](ConfigReverseProxyAuthenticationAPI.md#SetConfiguration4) | **Put** /api/v2/config/reverseProxyAuthentication | 



## DeleteConfiguration4

> DeleteConfiguration4(ctx).Execute()





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
	r, err := apiClient.ConfigReverseProxyAuthenticationAPI.DeleteConfiguration4(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigReverseProxyAuthenticationAPI.DeleteConfiguration4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfiguration4Request struct via the builder pattern


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


## GetConfiguration4

> ApiReverseProxyAuthenticationConfigurationDTO GetConfiguration4(ctx).Execute()





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
	resp, r, err := apiClient.ConfigReverseProxyAuthenticationAPI.GetConfiguration4(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigReverseProxyAuthenticationAPI.GetConfiguration4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration4`: ApiReverseProxyAuthenticationConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigReverseProxyAuthenticationAPI.GetConfiguration4`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration4Request struct via the builder pattern


### Return type

[**ApiReverseProxyAuthenticationConfigurationDTO**](ApiReverseProxyAuthenticationConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration4

> SetConfiguration4(ctx).ApiReverseProxyAuthenticationConfigurationDTO(apiReverseProxyAuthenticationConfigurationDTO).Execute()





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
	apiReverseProxyAuthenticationConfigurationDTO := *sonatypeiq.NewApiReverseProxyAuthenticationConfigurationDTO() // ApiReverseProxyAuthenticationConfigurationDTO | The request JSON could include: <ul><li>`enabled` indicates if the configuration is enabled.</li><li>`usernameHeader` is the name of the HTTP request header field that contains the username. The default value is `REMOTE_USER`.</li><li>`csrfProtectionDisabled` indicates if Cross-Site Request Forgery (CSRF) protection is disabled. Used for backward compatibility with old client plugins.</li><li>`logoutUrl` is the redirect URL when a user logs out. If set to `null` the user will not be redirected.</li></ul> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigReverseProxyAuthenticationAPI.SetConfiguration4(context.Background()).ApiReverseProxyAuthenticationConfigurationDTO(apiReverseProxyAuthenticationConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigReverseProxyAuthenticationAPI.SetConfiguration4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfiguration4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiReverseProxyAuthenticationConfigurationDTO** | [**ApiReverseProxyAuthenticationConfigurationDTO**](ApiReverseProxyAuthenticationConfigurationDTO.md) | The request JSON could include: &lt;ul&gt;&lt;li&gt;&#x60;enabled&#x60; indicates if the configuration is enabled.&lt;/li&gt;&lt;li&gt;&#x60;usernameHeader&#x60; is the name of the HTTP request header field that contains the username. The default value is &#x60;REMOTE_USER&#x60;.&lt;/li&gt;&lt;li&gt;&#x60;csrfProtectionDisabled&#x60; indicates if Cross-Site Request Forgery (CSRF) protection is disabled. Used for backward compatibility with old client plugins.&lt;/li&gt;&lt;li&gt;&#x60;logoutUrl&#x60; is the redirect URL when a user logs out. If set to &#x60;null&#x60; the user will not be redirected.&lt;/li&gt;&lt;/ul&gt; | 

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

