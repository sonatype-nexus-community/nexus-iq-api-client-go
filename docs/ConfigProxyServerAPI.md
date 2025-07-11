# \ConfigProxyServerAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration3**](ConfigProxyServerAPI.md#DeleteConfiguration3) | **Delete** /api/v2/config/httpProxyServer | 
[**GetConfiguration3**](ConfigProxyServerAPI.md#GetConfiguration3) | **Get** /api/v2/config/httpProxyServer | 
[**SetConfiguration3**](ConfigProxyServerAPI.md#SetConfiguration3) | **Put** /api/v2/config/httpProxyServer | 



## DeleteConfiguration3

> DeleteConfiguration3(ctx).Execute()





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
	r, err := apiClient.ConfigProxyServerAPI.DeleteConfiguration3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigProxyServerAPI.DeleteConfiguration3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfiguration3Request struct via the builder pattern


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


## GetConfiguration3

> ApiProxyServerConfigurationDTO GetConfiguration3(ctx).Execute()





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
	resp, r, err := apiClient.ConfigProxyServerAPI.GetConfiguration3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigProxyServerAPI.GetConfiguration3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration3`: ApiProxyServerConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigProxyServerAPI.GetConfiguration3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration3Request struct via the builder pattern


### Return type

[**ApiProxyServerConfigurationDTO**](ApiProxyServerConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration3

> SetConfiguration3(ctx).ApiProxyServerConfigurationDTO(apiProxyServerConfigurationDTO).Execute()





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
	r, err := apiClient.ConfigProxyServerAPI.SetConfiguration3(context.Background()).ApiProxyServerConfigurationDTO(apiProxyServerConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigProxyServerAPI.SetConfiguration3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfiguration3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiProxyServerConfigurationDTO** | [**ApiProxyServerConfigurationDTO**](ApiProxyServerConfigurationDTO.md) | The request JSON could include: &lt;ul&gt;&lt;li&gt;&#x60;hostname&#x60; is host name or IP address of the HTTP proxy server to use for outgoing HTTP connections.&lt;/li&gt;&lt;li&gt;&#x60;port&#x60; is the port number for the HTTP proxy server.&lt;/li&gt;&lt;li&gt;&#x60;username&#x60; is the username used to authenticate with the HTTP proxy server.&lt;/li&gt;&lt;li&gt;&#x60;password&#x60; is the password used for authentication with the HTTP proxy server.&lt;/li&gt;&lt;li&gt;&#x60;passwordIsIncluded&#x60; should be &#x60;true&#x60; if password is included in the request.&lt;ul&gt;&lt;li&gt;If &#x60;true&#x60; but the password is not included the password will be considered as &#x60;null&#x60;.&lt;/li&gt;&lt;li&gt;Can be &#x60;false&#x60; for update operations that do not a require password change. Note that updating the hostname and port requires a password to be provided.&lt;/li&gt; &lt;/ul&gt;&lt;li&gt;&#x60;excludeHosts&#x60; is a list of host names that are to be excluded from using the HTTP proxy server.&lt;/li&gt;&lt;/ul&gt; | 

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

