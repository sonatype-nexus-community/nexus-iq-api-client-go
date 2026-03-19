# \ConfigReverseProxyAuthenticationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration5**](ConfigReverseProxyAuthenticationAPI.md#DeleteConfiguration5) | **Delete** /api/v2/config/reverseProxyAuthentication | 
[**GetConfiguration5**](ConfigReverseProxyAuthenticationAPI.md#GetConfiguration5) | **Get** /api/v2/config/reverseProxyAuthentication | 
[**SetConfiguration5**](ConfigReverseProxyAuthenticationAPI.md#SetConfiguration5) | **Put** /api/v2/config/reverseProxyAuthentication | 



## DeleteConfiguration5

> DeleteConfiguration5(ctx).Execute()





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
	r, err := apiClient.ConfigReverseProxyAuthenticationAPI.DeleteConfiguration5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigReverseProxyAuthenticationAPI.DeleteConfiguration5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfiguration5Request struct via the builder pattern


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


## GetConfiguration5

> ApiReverseProxyAuthenticationConfigurationDTO GetConfiguration5(ctx).Execute()





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
	resp, r, err := apiClient.ConfigReverseProxyAuthenticationAPI.GetConfiguration5(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigReverseProxyAuthenticationAPI.GetConfiguration5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration5`: ApiReverseProxyAuthenticationConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigReverseProxyAuthenticationAPI.GetConfiguration5`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration5Request struct via the builder pattern


### Return type

[**ApiReverseProxyAuthenticationConfigurationDTO**](ApiReverseProxyAuthenticationConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration5

> SetConfiguration5(ctx).ApiReverseProxyAuthenticationConfigurationDTO(apiReverseProxyAuthenticationConfigurationDTO).Execute()





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
	r, err := apiClient.ConfigReverseProxyAuthenticationAPI.SetConfiguration5(context.Background()).ApiReverseProxyAuthenticationConfigurationDTO(apiReverseProxyAuthenticationConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigReverseProxyAuthenticationAPI.SetConfiguration5``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfiguration5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiReverseProxyAuthenticationConfigurationDTO** | [**ApiReverseProxyAuthenticationConfigurationDTO**](ApiReverseProxyAuthenticationConfigurationDTO.md) | The request JSON could include: &lt;ul&gt;&lt;li&gt;&#x60;enabled&#x60; indicates if the configuration is enabled.&lt;/li&gt;&lt;li&gt;&#x60;usernameHeader&#x60; is the name of the HTTP request header field that contains the username. The default value is &#x60;REMOTE_USER&#x60;.&lt;/li&gt;&lt;li&gt;&#x60;csrfProtectionDisabled&#x60; indicates if Cross-Site Request Forgery (CSRF) protection is disabled. Used for backward compatibility with old client plugins.&lt;/li&gt;&lt;li&gt;&#x60;logoutUrl&#x60; is the redirect URL when a user logs out. If set to &#x60;null&#x60; the user will not be redirected.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

