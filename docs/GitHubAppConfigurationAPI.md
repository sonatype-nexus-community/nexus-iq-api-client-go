# \GitHubAppConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateManifest**](GitHubAppConfigurationAPI.md#GenerateManifest) | **Post** /api/v2/githubApp/manifest | Generate GitHub App manifest
[**HandleInstallationSetup**](GitHubAppConfigurationAPI.md#HandleInstallationSetup) | **Get** /api/v2/githubApp/setupInstallation | Handle GitHub App installation setup callback with OAuth + PKCE
[**HandleRedirect**](GitHubAppConfigurationAPI.md#HandleRedirect) | **Get** /api/v2/githubApp/redirect | GitHub App registration redirect callback



## GenerateManifest

> GenerateManifest(ctx).OwnerId(ownerId).OrganizationName(organizationName).Execute()

Generate GitHub App manifest



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
	organizationName := "organizationName_example" // string | GitHub organization name (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.GitHubAppConfigurationAPI.GenerateManifest(context.Background()).OwnerId(ownerId).OrganizationName(organizationName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitHubAppConfigurationAPI.GenerateManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ownerId** | **string** | Owner (organization/application) ID | 
 **organizationName** | **string** | GitHub organization name | 

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


## HandleInstallationSetup

> HandleInstallationSetup(ctx).InstallationId(installationId).State(state).Code(code).Execute()

Handle GitHub App installation setup callback with OAuth + PKCE



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
	installationId := int64(789) // int64 | GitHub App installation ID
	state := "state_example" // string | State token for CSRF protection
	code := "code_example" // string | OAuth authorization code

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.GitHubAppConfigurationAPI.HandleInstallationSetup(context.Background()).InstallationId(installationId).State(state).Code(code).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitHubAppConfigurationAPI.HandleInstallationSetup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleInstallationSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **installationId** | **int64** | GitHub App installation ID | 
 **state** | **string** | State token for CSRF protection | 
 **code** | **string** | OAuth authorization code | 

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


## HandleRedirect

> HandleRedirect(ctx).Code(code).State(state).Execute()

GitHub App registration redirect callback



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
	code := "code_example" // string | Temporary manifest conversion code from GitHub
	state := "state_example" // string | OAuth state token for CSRF protection (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.GitHubAppConfigurationAPI.HandleRedirect(context.Background()).Code(code).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitHubAppConfigurationAPI.HandleRedirect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHandleRedirectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **string** | Temporary manifest conversion code from GitHub | 
 **state** | **string** | OAuth state token for CSRF protection | 

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

