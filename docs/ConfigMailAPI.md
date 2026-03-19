# \ConfigMailAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration3**](ConfigMailAPI.md#DeleteConfiguration3) | **Delete** /api/v2/config/mail | 
[**GetConfiguration3**](ConfigMailAPI.md#GetConfiguration3) | **Get** /api/v2/config/mail | 
[**SetConfiguration3**](ConfigMailAPI.md#SetConfiguration3) | **Put** /api/v2/config/mail | 
[**TestConfiguration**](ConfigMailAPI.md#TestConfiguration) | **Post** /api/v2/config/mail/test/{recipientEmail} | 



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
	r, err := apiClient.ConfigMailAPI.DeleteConfiguration3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigMailAPI.DeleteConfiguration3``: %v\n", err)
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

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration3

> ApiMailConfigurationDTO GetConfiguration3(ctx).Execute()





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
	resp, r, err := apiClient.ConfigMailAPI.GetConfiguration3(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigMailAPI.GetConfiguration3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration3`: ApiMailConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigMailAPI.GetConfiguration3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration3Request struct via the builder pattern


### Return type

[**ApiMailConfigurationDTO**](ApiMailConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration3

> SetConfiguration3(ctx).ApiMailConfigurationDTO(apiMailConfigurationDTO).Execute()





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
	apiMailConfigurationDTO := *sonatypeiq.NewApiMailConfigurationDTO() // ApiMailConfigurationDTO | Provide one or more values for the following in the JSON payload:<ul><li>`hostname` - is the hostname or IP address of the SMTP server used for outgoing mail.</li><li>`port` - is the port number on which the SMTP server accepts email requests.</li><li>`password` - depends upon the value of `passwordIsIncluded`.</li><li>`passwordIsIncluded` - if set to true, value must be provided for `password`, null is allowed.If set to false, the previous value will remain unchanged, provided that `hostname` and `port` are not changed.<li>`sslEnabled` - is a boolean flag indicating if the connection to the SMTP server should use SSL/TLSright from the start.</li><li>`startIsEnabled`- is a boolean flag indicating if the connection to the SMTP server should attempt toupgrade to SSL/TLS using the STARTTLS command.<li>`systemEmail` - is the email address used for the FROM header in emails sent by the IQ Server.</li></ul> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigMailAPI.SetConfiguration3(context.Background()).ApiMailConfigurationDTO(apiMailConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigMailAPI.SetConfiguration3``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfiguration3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiMailConfigurationDTO** | [**ApiMailConfigurationDTO**](ApiMailConfigurationDTO.md) | Provide one or more values for the following in the JSON payload:&lt;ul&gt;&lt;li&gt;&#x60;hostname&#x60; - is the hostname or IP address of the SMTP server used for outgoing mail.&lt;/li&gt;&lt;li&gt;&#x60;port&#x60; - is the port number on which the SMTP server accepts email requests.&lt;/li&gt;&lt;li&gt;&#x60;password&#x60; - depends upon the value of &#x60;passwordIsIncluded&#x60;.&lt;/li&gt;&lt;li&gt;&#x60;passwordIsIncluded&#x60; - if set to true, value must be provided for &#x60;password&#x60;, null is allowed.If set to false, the previous value will remain unchanged, provided that &#x60;hostname&#x60; and &#x60;port&#x60; are not changed.&lt;li&gt;&#x60;sslEnabled&#x60; - is a boolean flag indicating if the connection to the SMTP server should use SSL/TLSright from the start.&lt;/li&gt;&lt;li&gt;&#x60;startIsEnabled&#x60;- is a boolean flag indicating if the connection to the SMTP server should attempt toupgrade to SSL/TLS using the STARTTLS command.&lt;li&gt;&#x60;systemEmail&#x60; - is the email address used for the FROM header in emails sent by the IQ Server.&lt;/li&gt;&lt;/ul&gt; | 

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


## TestConfiguration

> TestConfiguration(ctx, recipientEmail).ApiMailConfigurationDTO(apiMailConfigurationDTO).Execute()





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
	recipientEmail := "recipientEmail_example" // string | Enter the test recipient email address.
	apiMailConfigurationDTO := *sonatypeiq.NewApiMailConfigurationDTO() // ApiMailConfigurationDTO | Provide one or more values for the following in the JSON payload:<ul><li>`hostname` - is the hostname or IP address of the SMTP server used for outgoing mail.</li><li>`port` - is the port number on which the SMTP server accepts email requests.</li><li>`password` - depends upon the value of `passwordIsIncluded`.</li><li>`passwordIsIncluded` - if set to true, value must be provided for `password`, null is allowed.If set to false, the previous value will remain unchanged, provided that `hostname` and `port` are not changed.<li>`sslEnabled` - is a boolean flag indicating if the connection to the SMTP server should use SSL/TLSright from the start.</li><li>`startIsEnabled`- is a boolean flag indicating if the connection to the SMTP server should attempt toupgrade to SSL/TLS using the STARTTLS command.<li>`systemEmail` - is the email address used for the FROM header in emails sent by the IQ Server.</li></ul> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigMailAPI.TestConfiguration(context.Background(), recipientEmail).ApiMailConfigurationDTO(apiMailConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigMailAPI.TestConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recipientEmail** | **string** | Enter the test recipient email address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiMailConfigurationDTO** | [**ApiMailConfigurationDTO**](ApiMailConfigurationDTO.md) | Provide one or more values for the following in the JSON payload:&lt;ul&gt;&lt;li&gt;&#x60;hostname&#x60; - is the hostname or IP address of the SMTP server used for outgoing mail.&lt;/li&gt;&lt;li&gt;&#x60;port&#x60; - is the port number on which the SMTP server accepts email requests.&lt;/li&gt;&lt;li&gt;&#x60;password&#x60; - depends upon the value of &#x60;passwordIsIncluded&#x60;.&lt;/li&gt;&lt;li&gt;&#x60;passwordIsIncluded&#x60; - if set to true, value must be provided for &#x60;password&#x60;, null is allowed.If set to false, the previous value will remain unchanged, provided that &#x60;hostname&#x60; and &#x60;port&#x60; are not changed.&lt;li&gt;&#x60;sslEnabled&#x60; - is a boolean flag indicating if the connection to the SMTP server should use SSL/TLSright from the start.&lt;/li&gt;&lt;li&gt;&#x60;startIsEnabled&#x60;- is a boolean flag indicating if the connection to the SMTP server should attempt toupgrade to SSL/TLS using the STARTTLS command.&lt;li&gt;&#x60;systemEmail&#x60; - is the email address used for the FROM header in emails sent by the IQ Server.&lt;/li&gt;&lt;/ul&gt; | 

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

