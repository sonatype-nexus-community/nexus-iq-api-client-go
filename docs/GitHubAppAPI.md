# \GitHubAppAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGitHubApp**](GitHubAppAPI.md#DeleteGitHubApp) | **Delete** /api/v2/githubApp/{githubAppId} | Delete a GitHub App
[**ListGitHubApps**](GitHubAppAPI.md#ListGitHubApps) | **Get** /api/v2/githubApp | List GitHub Apps for an owner



## DeleteGitHubApp

> DeleteGitHubApp(ctx, githubAppId).OwnerId(ownerId).Execute()

Delete a GitHub App



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
	githubAppId := "githubAppId_example" // string | GitHub App ID
	ownerId := "ownerId_example" // string | Owner (organization/application) ID

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.GitHubAppAPI.DeleteGitHubApp(context.Background(), githubAppId).OwnerId(ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitHubAppAPI.DeleteGitHubApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubAppId** | **string** | GitHub App ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGitHubAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ownerId** | **string** | Owner (organization/application) ID | 

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


## ListGitHubApps

> ListGitHubApps(ctx).OwnerId(ownerId).Execute()

List GitHub Apps for an owner



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
	ownerId := "ownerId_example" // string | Owner (organization/application) ID

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.GitHubAppAPI.ListGitHubApps(context.Background()).OwnerId(ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitHubAppAPI.ListGitHubApps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGitHubAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | Owner (organization/application) ID | 

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

