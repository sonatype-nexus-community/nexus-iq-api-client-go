# \SourceControlAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSourceControl**](SourceControlAPI.md#AddSourceControl) | **Post** /api/v2/sourceControl/{ownerType}/{internalOwnerId} | 
[**AutomaticRoleAssignment**](SourceControlAPI.md#AutomaticRoleAssignment) | **Post** /api/v2/sourceControl/automaticRoleAssignment/{publicId} | 
[**DeleteSourceControl**](SourceControlAPI.md#DeleteSourceControl) | **Delete** /api/v2/sourceControl/{ownerType}/{internalOwnerId} | 
[**GetSourceControl1**](SourceControlAPI.md#GetSourceControl1) | **Get** /api/v2/sourceControl/{ownerType}/{internalOwnerId} | 
[**UpdateSourceControl**](SourceControlAPI.md#UpdateSourceControl) | **Put** /api/v2/sourceControl/{ownerType}/{internalOwnerId} | 



## AddSourceControl

> ApiSourceControlDTO AddSourceControl(ctx, ownerType, internalOwnerId).ApiSourceControlDTO(apiSourceControlDTO).Execute()





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
	ownerType := "ownerType_example" // string | Enter the value for ownerType.
	internalOwnerId := "internalOwnerId_example" // string | Enter the value for internal ownerId. Use ROOT_ORGANIZATION_ID for root organization.
	apiSourceControlDTO := *sonatypeiq.NewApiSourceControlDTO() // ApiSourceControlDTO | Specify the SCM settings for the ownerId specified above in the request JSON.<ul><li><code>id</code> is the internal owner ID.</li><li><code>repositoryUrl</code> is the http(s) and ssh urls for the application specified in the ownerId.</li><li><code>username</code> is optional, can be provided for Bitbucket Server and Cloud.</li><li><code>token</code> is optional,if inherited. If provided, this value will override the value inherited from the root organization, organization or application level.<li><code>provider</code> is the name of of the SCM system. Allowed values are <code>azure</code>, <code>github</code>, <code>gitlab</code>, and <code>bitbucket</code>.</li><li><code>baseBranch</code> is required for the root organization. Organizations and applications inherit from the root unless overridden.</li><li><code>enablePullRequests</code> has been deprecated in version 124.</li><li><code>remediationPullRequestsEnabled</code> is optional. Set it to `true` to enable the Automated Pull Requests.</li><li><code>enableStatusChecks</code> has been deprecated in version 124.</li><li><code>statusChecksEnabled</code> is an internal field.</li><li><code>pullRequestCommentingEnabled</code> is optional. Set it to `true` to enable the  Pull Request Commenting feature.</li><li><code>sourceControlEvaluationsEnabled</code> is set to `true` to enable source control evaluations for the continuous risk profile feature.</li><li><code>sourceControlScanTarget</code> is the path inside the repository.</li><li><code>sshEnabled</code> is set to `true` to enable ssh.</li><li><code>commitStatusEnabled</code> is set to `true` if interaction with the commit statuses on the SCM is enabled.</li></ul> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.AddSourceControl(context.Background(), ownerType, internalOwnerId).ApiSourceControlDTO(apiSourceControlDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.AddSourceControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSourceControl`: ApiSourceControlDTO
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.AddSourceControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the value for ownerType. | 
**internalOwnerId** | **string** | Enter the value for internal ownerId. Use ROOT_ORGANIZATION_ID for root organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSourceControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiSourceControlDTO** | [**ApiSourceControlDTO**](ApiSourceControlDTO.md) | Specify the SCM settings for the ownerId specified above in the request JSON.&lt;ul&gt;&lt;li&gt;&lt;code&gt;id&lt;/code&gt; is the internal owner ID.&lt;/li&gt;&lt;li&gt;&lt;code&gt;repositoryUrl&lt;/code&gt; is the http(s) and ssh urls for the application specified in the ownerId.&lt;/li&gt;&lt;li&gt;&lt;code&gt;username&lt;/code&gt; is optional, can be provided for Bitbucket Server and Cloud.&lt;/li&gt;&lt;li&gt;&lt;code&gt;token&lt;/code&gt; is optional,if inherited. If provided, this value will override the value inherited from the root organization, organization or application level.&lt;li&gt;&lt;code&gt;provider&lt;/code&gt; is the name of of the SCM system. Allowed values are &lt;code&gt;azure&lt;/code&gt;, &lt;code&gt;github&lt;/code&gt;, &lt;code&gt;gitlab&lt;/code&gt;, and &lt;code&gt;bitbucket&lt;/code&gt;.&lt;/li&gt;&lt;li&gt;&lt;code&gt;baseBranch&lt;/code&gt; is required for the root organization. Organizations and applications inherit from the root unless overridden.&lt;/li&gt;&lt;li&gt;&lt;code&gt;enablePullRequests&lt;/code&gt; has been deprecated in version 124.&lt;/li&gt;&lt;li&gt;&lt;code&gt;remediationPullRequestsEnabled&lt;/code&gt; is optional. Set it to &#x60;true&#x60; to enable the Automated Pull Requests.&lt;/li&gt;&lt;li&gt;&lt;code&gt;enableStatusChecks&lt;/code&gt; has been deprecated in version 124.&lt;/li&gt;&lt;li&gt;&lt;code&gt;statusChecksEnabled&lt;/code&gt; is an internal field.&lt;/li&gt;&lt;li&gt;&lt;code&gt;pullRequestCommentingEnabled&lt;/code&gt; is optional. Set it to &#x60;true&#x60; to enable the  Pull Request Commenting feature.&lt;/li&gt;&lt;li&gt;&lt;code&gt;sourceControlEvaluationsEnabled&lt;/code&gt; is set to &#x60;true&#x60; to enable source control evaluations for the continuous risk profile feature.&lt;/li&gt;&lt;li&gt;&lt;code&gt;sourceControlScanTarget&lt;/code&gt; is the path inside the repository.&lt;/li&gt;&lt;li&gt;&lt;code&gt;sshEnabled&lt;/code&gt; is set to &#x60;true&#x60; to enable ssh.&lt;/li&gt;&lt;li&gt;&lt;code&gt;commitStatusEnabled&lt;/code&gt; is set to &#x60;true&#x60; if interaction with the commit statuses on the SCM is enabled.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**ApiSourceControlDTO**](ApiSourceControlDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AutomaticRoleAssignment

> []string AutomaticRoleAssignment(ctx, publicId).Execute()





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
	publicId := "publicId_example" // string | Enter the public applicationId for automatic role assignment.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.AutomaticRoleAssignment(context.Background(), publicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.AutomaticRoleAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AutomaticRoleAssignment`: []string
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.AutomaticRoleAssignment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicId** | **string** | Enter the public applicationId for automatic role assignment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAutomaticRoleAssignmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSourceControl

> DeleteSourceControl(ctx, ownerType, internalOwnerId).Execute()





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
	ownerType := "ownerType_example" // string | Enter the value for ownerType.
	internalOwnerId := "internalOwnerId_example" // string | Enter the value for internal ownerId.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.SourceControlAPI.DeleteSourceControl(context.Background(), ownerType, internalOwnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.DeleteSourceControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the value for ownerType. | 
**internalOwnerId** | **string** | Enter the value for internal ownerId. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSourceControlRequest struct via the builder pattern


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


## GetSourceControl1

> ApiSourceControlDTO GetSourceControl1(ctx, ownerType, internalOwnerId).Execute()





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
	ownerType := "ownerType_example" // string | Enter the value for ownerType.
	internalOwnerId := "internalOwnerId_example" // string | Enter the value for internal ownerId. Use ROOT_ORGANIZATION_ID for the root organization

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.GetSourceControl1(context.Background(), ownerType, internalOwnerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.GetSourceControl1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSourceControl1`: ApiSourceControlDTO
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.GetSourceControl1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the value for ownerType. | 
**internalOwnerId** | **string** | Enter the value for internal ownerId. Use ROOT_ORGANIZATION_ID for the root organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceControl1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiSourceControlDTO**](ApiSourceControlDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSourceControl

> ApiSourceControlDTO UpdateSourceControl(ctx, ownerType, internalOwnerId).ApiSourceControlDTO(apiSourceControlDTO).Execute()





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
	ownerType := "ownerType_example" // string | Enter the value for ownerType.
	internalOwnerId := "internalOwnerId_example" // string | Enter the internal ownerId. Use ROOT_ORGANIZATION_ID for the root organization.
	apiSourceControlDTO := *sonatypeiq.NewApiSourceControlDTO() // ApiSourceControlDTO | Specify the SCM settings for the ownerId specified above in the request JSON.<ul><li><code>id</code> is the internal owner ID.</li><li><code>repositoryUrl</code> is the http(s) and ssh urls for the application specified in the ownerId.</li><li><code>username</code> is optional, can be provided for Bitbucket Server and Cloud.</li><li><code>token</code> is optional if inherited. If provided, this value will override the value inherited from the root organization, organization or application level.<li><code>provider</code> is the name of of the SCM system. Allowed values are <code>azure</code>, <code>github</code>, <code>gitlab</code>, and <code>bitbucket</code>.</li><li><code>baseBranch</code> is required for the root organization. Organizations and applications inherit from the root unless overridden.</li><li><code>enablePullRequests</code> has been deprecated in version 124.</li><li><code>remediationPullRequestsEnabled</code> is optional. Set it to `true` to enable the Automated Pull Requests.</li><li><code>enableStatusChecks</code> has been deprecated in version 124.</li><li><code>statusChecksEnabled</code> is an internal field.</li><li><code>pullRequestCommentingEnabled</code> is optional. Set it to `true` to enable the  Pull Request Commenting feature.</li><li><code>sourceControlEvaluationsEnabled</code> is set to `true` to enable source control evaluations for the continuous risk profile feature.</li><li><code>sourceControlScanTarget</code> is the path inside the repository.</li><li><code>sshEnabled</code> is set to `true` to enable ssh.</li><li><code>commitStatusEnabled</code> is set to `true` if interaction with the commit statuses on the SCM is enabled.</li></ul> (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.SourceControlAPI.UpdateSourceControl(context.Background(), ownerType, internalOwnerId).ApiSourceControlDTO(apiSourceControlDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SourceControlAPI.UpdateSourceControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSourceControl`: ApiSourceControlDTO
	fmt.Fprintf(os.Stdout, "Response from `SourceControlAPI.UpdateSourceControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the value for ownerType. | 
**internalOwnerId** | **string** | Enter the internal ownerId. Use ROOT_ORGANIZATION_ID for the root organization. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiSourceControlDTO** | [**ApiSourceControlDTO**](ApiSourceControlDTO.md) | Specify the SCM settings for the ownerId specified above in the request JSON.&lt;ul&gt;&lt;li&gt;&lt;code&gt;id&lt;/code&gt; is the internal owner ID.&lt;/li&gt;&lt;li&gt;&lt;code&gt;repositoryUrl&lt;/code&gt; is the http(s) and ssh urls for the application specified in the ownerId.&lt;/li&gt;&lt;li&gt;&lt;code&gt;username&lt;/code&gt; is optional, can be provided for Bitbucket Server and Cloud.&lt;/li&gt;&lt;li&gt;&lt;code&gt;token&lt;/code&gt; is optional if inherited. If provided, this value will override the value inherited from the root organization, organization or application level.&lt;li&gt;&lt;code&gt;provider&lt;/code&gt; is the name of of the SCM system. Allowed values are &lt;code&gt;azure&lt;/code&gt;, &lt;code&gt;github&lt;/code&gt;, &lt;code&gt;gitlab&lt;/code&gt;, and &lt;code&gt;bitbucket&lt;/code&gt;.&lt;/li&gt;&lt;li&gt;&lt;code&gt;baseBranch&lt;/code&gt; is required for the root organization. Organizations and applications inherit from the root unless overridden.&lt;/li&gt;&lt;li&gt;&lt;code&gt;enablePullRequests&lt;/code&gt; has been deprecated in version 124.&lt;/li&gt;&lt;li&gt;&lt;code&gt;remediationPullRequestsEnabled&lt;/code&gt; is optional. Set it to &#x60;true&#x60; to enable the Automated Pull Requests.&lt;/li&gt;&lt;li&gt;&lt;code&gt;enableStatusChecks&lt;/code&gt; has been deprecated in version 124.&lt;/li&gt;&lt;li&gt;&lt;code&gt;statusChecksEnabled&lt;/code&gt; is an internal field.&lt;/li&gt;&lt;li&gt;&lt;code&gt;pullRequestCommentingEnabled&lt;/code&gt; is optional. Set it to &#x60;true&#x60; to enable the  Pull Request Commenting feature.&lt;/li&gt;&lt;li&gt;&lt;code&gt;sourceControlEvaluationsEnabled&lt;/code&gt; is set to &#x60;true&#x60; to enable source control evaluations for the continuous risk profile feature.&lt;/li&gt;&lt;li&gt;&lt;code&gt;sourceControlScanTarget&lt;/code&gt; is the path inside the repository.&lt;/li&gt;&lt;li&gt;&lt;code&gt;sshEnabled&lt;/code&gt; is set to &#x60;true&#x60; to enable ssh.&lt;/li&gt;&lt;li&gt;&lt;code&gt;commitStatusEnabled&lt;/code&gt; is set to &#x60;true&#x60; if interaction with the commit statuses on the SCM is enabled.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**ApiSourceControlDTO**](ApiSourceControlDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

