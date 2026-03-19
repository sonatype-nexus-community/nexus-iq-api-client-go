# \ConfigProxyServerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration4**](ConfigProxyServerAPI.md#DeleteConfiguration4) | **Delete** /api/v2/config/httpProxyServer | 
[**GetConfiguration4**](ConfigProxyServerAPI.md#GetConfiguration4) | **Get** /api/v2/config/httpProxyServer | 
[**SetConfiguration4**](ConfigProxyServerAPI.md#SetConfiguration4) | **Put** /api/v2/config/httpProxyServer | 



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
	r, err := apiClient.ConfigProxyServerAPI.DeleteConfiguration4(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigProxyServerAPI.DeleteConfiguration4``: %v\n", err)
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

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration4

> ApiProxyServerConfigurationDTO GetConfiguration4(ctx).Execute()





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
	resp, r, err := apiClient.ConfigProxyServerAPI.GetConfiguration4(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigProxyServerAPI.GetConfiguration4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration4`: ApiProxyServerConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigProxyServerAPI.GetConfiguration4`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration4Request struct via the builder pattern


### Return type

[**ApiProxyServerConfigurationDTO**](ApiProxyServerConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration4

> SetConfiguration4(ctx).ApiProxyServerConfigurationDTO(apiProxyServerConfigurationDTO).Execute()





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
	apiProxyServerConfigurationDTO := *sonatypeiq.NewApiProxyServerConfigurationDTO() // ApiProxyServerConfigurationDTO | The request JSON could include: <ul><li>`hostname` is host name or IP address of the HTTP proxy server to use for outgoing HTTP connections.</li><li>`port` is the port number for the HTTP proxy server.</li><li>`username` is the username used to authenticate with the HTTP proxy server.</li><li>`password` is the password used for authentication with the HTTP proxy server.</li><li>`passwordIsIncluded` should be `true` if password is included in the request.<ul><li>If `true` but the password is not included the password will be considered as `null`.</li><li>Can be `false` for update operations that do not a require password change. Note that updating the hostname and port requires a password to be provided.</li> </ul><li>`excludeHosts` is a list of host names that are to be excluded from using the HTTP proxy server.</li></ul> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigProxyServerAPI.SetConfiguration4(context.Background()).ApiProxyServerConfigurationDTO(apiProxyServerConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigProxyServerAPI.SetConfiguration4``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfiguration4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiProxyServerConfigurationDTO** | [**ApiProxyServerConfigurationDTO**](ApiProxyServerConfigurationDTO.md) | The request JSON could include: &lt;ul&gt;&lt;li&gt;&#x60;hostname&#x60; is host name or IP address of the HTTP proxy server to use for outgoing HTTP connections.&lt;/li&gt;&lt;li&gt;&#x60;port&#x60; is the port number for the HTTP proxy server.&lt;/li&gt;&lt;li&gt;&#x60;username&#x60; is the username used to authenticate with the HTTP proxy server.&lt;/li&gt;&lt;li&gt;&#x60;password&#x60; is the password used for authentication with the HTTP proxy server.&lt;/li&gt;&lt;li&gt;&#x60;passwordIsIncluded&#x60; should be &#x60;true&#x60; if password is included in the request.&lt;ul&gt;&lt;li&gt;If &#x60;true&#x60; but the password is not included the password will be considered as &#x60;null&#x60;.&lt;/li&gt;&lt;li&gt;Can be &#x60;false&#x60; for update operations that do not a require password change. Note that updating the hostname and port requires a password to be provided.&lt;/li&gt; &lt;/ul&gt;&lt;li&gt;&#x60;excludeHosts&#x60; is a list of host names that are to be excluded from using the HTTP proxy server.&lt;/li&gt;&lt;/ul&gt; | 

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

