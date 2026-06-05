# \ConfigSourceControlAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteConfiguration7**](ConfigSourceControlAPI.md#DeleteConfiguration7) | **Delete** /api/v2/config/sourceControl | 
[**GetConfiguration7**](ConfigSourceControlAPI.md#GetConfiguration7) | **Get** /api/v2/config/sourceControl | 
[**SetConfiguration6**](ConfigSourceControlAPI.md#SetConfiguration6) | **Put** /api/v2/config/sourceControl | 



## DeleteConfiguration7

> DeleteConfiguration7(ctx).Execute()





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
	r, err := apiClient.ConfigSourceControlAPI.DeleteConfiguration7(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigSourceControlAPI.DeleteConfiguration7``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConfiguration7Request struct via the builder pattern


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


## GetConfiguration7

> ApiSourceControlConfigurationDTO GetConfiguration7(ctx).Execute()





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
	resp, r, err := apiClient.ConfigSourceControlAPI.GetConfiguration7(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigSourceControlAPI.GetConfiguration7``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration7`: ApiSourceControlConfigurationDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigSourceControlAPI.GetConfiguration7`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfiguration7Request struct via the builder pattern


### Return type

[**ApiSourceControlConfigurationDTO**](ApiSourceControlConfigurationDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetConfiguration6

> SetConfiguration6(ctx).ApiSourceControlConfigurationDTO(apiSourceControlConfigurationDTO).Execute()





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
	apiSourceControlConfigurationDTO := *sonatypeiq.NewApiSourceControlConfigurationDTO() // ApiSourceControlConfigurationDTO | Provide the settings for the SCM configuration as below: <ul><li>`cloneDirectory` is the location of the cloned repository that will be used by the IQ server. If a relative path is provided, then that path will be created inside the  `sonatype-work` directory and your repository will be created within this. A return value `source-control` indicates that this setting is not configured.</li><li>`gitImplementation` will have the value `java` for JGit or `native` for a native git client.</li><li>`prCommentPurgeWindow` is the number of days until the comments of a Pull Request (PR) are allowed to be purged.</li><li>`prEventPurgeWindow` is the number of days until PR events are allowed to be purged.</li><li>`gitExecutable` is the absolute path to a native client. No value indicates the native git client is on the system path.</li>`gitTimeoutSeconds` is the number of seconds a git command can execute before timing out.</li>`commitUsername` is the username that will be used for the SCM features. The value `NexusIQ` indicates the default value.</li>`commitEmail` is the commit email that will be used for the SCM features.`useUsernameInRepositoryCloneUrl` indicates if the username will be added to the URL for the clonedrepository. This can be used in conjunction with `commitEmail` to support the 'Verified Committer' feature of Bitbucket.</li>`defaultBranchMonitoringStartTime` has a default value between 00:00 and 00:10. It is the time at which the default branch monitoring will start for the first time.</li>`defaultBranchMonitoringIntervalHours` is the number of hours elapsed between the executions of default branch monitoring by the IQ Server. The default value is 24 hours.</li><li>`pullRequestMonitoringIntervalSeconds` is the time in seconds between consecutive execution of PR monitoring. The default value is 60 seconds.</li></ul> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigSourceControlAPI.SetConfiguration6(context.Background()).ApiSourceControlConfigurationDTO(apiSourceControlConfigurationDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigSourceControlAPI.SetConfiguration6``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetConfiguration6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiSourceControlConfigurationDTO** | [**ApiSourceControlConfigurationDTO**](ApiSourceControlConfigurationDTO.md) | Provide the settings for the SCM configuration as below: &lt;ul&gt;&lt;li&gt;&#x60;cloneDirectory&#x60; is the location of the cloned repository that will be used by the IQ server. If a relative path is provided, then that path will be created inside the  &#x60;sonatype-work&#x60; directory and your repository will be created within this. A return value &#x60;source-control&#x60; indicates that this setting is not configured.&lt;/li&gt;&lt;li&gt;&#x60;gitImplementation&#x60; will have the value &#x60;java&#x60; for JGit or &#x60;native&#x60; for a native git client.&lt;/li&gt;&lt;li&gt;&#x60;prCommentPurgeWindow&#x60; is the number of days until the comments of a Pull Request (PR) are allowed to be purged.&lt;/li&gt;&lt;li&gt;&#x60;prEventPurgeWindow&#x60; is the number of days until PR events are allowed to be purged.&lt;/li&gt;&lt;li&gt;&#x60;gitExecutable&#x60; is the absolute path to a native client. No value indicates the native git client is on the system path.&lt;/li&gt;&#x60;gitTimeoutSeconds&#x60; is the number of seconds a git command can execute before timing out.&lt;/li&gt;&#x60;commitUsername&#x60; is the username that will be used for the SCM features. The value &#x60;NexusIQ&#x60; indicates the default value.&lt;/li&gt;&#x60;commitEmail&#x60; is the commit email that will be used for the SCM features.&#x60;useUsernameInRepositoryCloneUrl&#x60; indicates if the username will be added to the URL for the clonedrepository. This can be used in conjunction with &#x60;commitEmail&#x60; to support the &#39;Verified Committer&#39; feature of Bitbucket.&lt;/li&gt;&#x60;defaultBranchMonitoringStartTime&#x60; has a default value between 00:00 and 00:10. It is the time at which the default branch monitoring will start for the first time.&lt;/li&gt;&#x60;defaultBranchMonitoringIntervalHours&#x60; is the number of hours elapsed between the executions of default branch monitoring by the IQ Server. The default value is 24 hours.&lt;/li&gt;&lt;li&gt;&#x60;pullRequestMonitoringIntervalSeconds&#x60; is the time in seconds between consecutive execution of PR monitoring. The default value is 60 seconds.&lt;/li&gt;&lt;/ul&gt; | 

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

