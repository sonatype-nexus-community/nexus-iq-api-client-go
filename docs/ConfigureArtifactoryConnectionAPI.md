# \ConfigureArtifactoryConnectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddArtifactoryConnection**](ConfigureArtifactoryConnectionAPI.md#AddArtifactoryConnection) | **Post** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId} | 
[**DeleteArtifactoryConnection**](ConfigureArtifactoryConnectionAPI.md#DeleteArtifactoryConnection) | **Delete** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/{artifactoryConnectionId} | 
[**GetArtifactoryConnection**](ConfigureArtifactoryConnectionAPI.md#GetArtifactoryConnection) | **Get** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/{artifactoryConnectionId} | 
[**GetOwnerArtifactoryConnection**](ConfigureArtifactoryConnectionAPI.md#GetOwnerArtifactoryConnection) | **Get** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId} | 
[**TestArtifactoryConnection**](ConfigureArtifactoryConnectionAPI.md#TestArtifactoryConnection) | **Post** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/test | 
[**TestArtifactoryConnection1**](ConfigureArtifactoryConnectionAPI.md#TestArtifactoryConnection1) | **Post** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/{artifactoryConnectionId}/test | 
[**UpdateArtifactoryConnection**](ConfigureArtifactoryConnectionAPI.md#UpdateArtifactoryConnection) | **Put** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId}/{artifactoryConnectionId} | 
[**UpdateOwnerArtifactoryConnectionStatus**](ConfigureArtifactoryConnectionAPI.md#UpdateOwnerArtifactoryConnectionStatus) | **Put** /api/v2/config/artifactoryConnection/{ownerType}/{internalOwnerId} | 



## AddArtifactoryConnection

> ApiArtifactoryConnectionDTO AddArtifactoryConnection(ctx, ownerType, internalOwnerId).ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO).Execute()





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
	ownerType := "ownerType_example" // string | Select the owner type.
	internalOwnerId := "internalOwnerId_example" // string | Enter the internal ID of the owner.
	apiArtifactoryConnectionDTO := *sonatypeiq.NewApiArtifactoryConnectionDTO() // ApiArtifactoryConnectionDTO | Enter values for the new Artifactory connection.<ul><li>`isAnonymous` indicates if the connection is anonymous.</li><li>`baseUrl` is the baseURL of the Artifactory instance.</li><li>`username` and `password` to authenticate the Artifactory connection.</li></ul>

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureArtifactoryConnectionAPI.AddArtifactoryConnection(context.Background(), ownerType, internalOwnerId).ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureArtifactoryConnectionAPI.AddArtifactoryConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddArtifactoryConnection`: ApiArtifactoryConnectionDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigureArtifactoryConnectionAPI.AddArtifactoryConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the owner type. | 
**internalOwnerId** | **string** | Enter the internal ID of the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddArtifactoryConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiArtifactoryConnectionDTO** | [**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md) | Enter values for the new Artifactory connection.&lt;ul&gt;&lt;li&gt;&#x60;isAnonymous&#x60; indicates if the connection is anonymous.&lt;/li&gt;&lt;li&gt;&#x60;baseUrl&#x60; is the baseURL of the Artifactory instance.&lt;/li&gt;&lt;li&gt;&#x60;username&#x60; and &#x60;password&#x60; to authenticate the Artifactory connection.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtifactoryConnection

> DeleteArtifactoryConnection(ctx, ownerType, internalOwnerId, artifactoryConnectionId).Execute()





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
	ownerType := "ownerType_example" // string | Select the owner type.
	internalOwnerId := "internalOwnerId_example" // string | Enter the internal ID of the owner.
	artifactoryConnectionId := "artifactoryConnectionId_example" // string | Enter the Artifactory connection ID.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigureArtifactoryConnectionAPI.DeleteArtifactoryConnection(context.Background(), ownerType, internalOwnerId, artifactoryConnectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureArtifactoryConnectionAPI.DeleteArtifactoryConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the owner type. | 
**internalOwnerId** | **string** | Enter the internal ID of the owner. | 
**artifactoryConnectionId** | **string** | Enter the Artifactory connection ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtifactoryConnectionRequest struct via the builder pattern


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


## GetArtifactoryConnection

> ApiArtifactoryConnectionDTO GetArtifactoryConnection(ctx, ownerType, internalOwnerId, artifactoryConnectionId).Execute()





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
	ownerType := "ownerType_example" // string | Select the owner type.
	internalOwnerId := "internalOwnerId_example" // string | Enter the internal ID of the owner.
	artifactoryConnectionId := "artifactoryConnectionId_example" // string | Enter the Artifactory connection ID.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureArtifactoryConnectionAPI.GetArtifactoryConnection(context.Background(), ownerType, internalOwnerId, artifactoryConnectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureArtifactoryConnectionAPI.GetArtifactoryConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifactoryConnection`: ApiArtifactoryConnectionDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigureArtifactoryConnectionAPI.GetArtifactoryConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the owner type. | 
**internalOwnerId** | **string** | Enter the internal ID of the owner. | 
**artifactoryConnectionId** | **string** | Enter the Artifactory connection ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactoryConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOwnerArtifactoryConnection

> ApiOwnerArtifactoryConnectionDTO GetOwnerArtifactoryConnection(ctx, ownerType, internalOwnerId).Inherit(inherit).Execute()





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
	ownerType := "ownerType_example" // string | Select the owner type.
	internalOwnerId := "internalOwnerId_example" // string | Enter the internal ID of the owner.
	inherit := true // bool | Specify whether to include details from an inherited Artifactory connection. (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureArtifactoryConnectionAPI.GetOwnerArtifactoryConnection(context.Background(), ownerType, internalOwnerId).Inherit(inherit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureArtifactoryConnectionAPI.GetOwnerArtifactoryConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOwnerArtifactoryConnection`: ApiOwnerArtifactoryConnectionDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigureArtifactoryConnectionAPI.GetOwnerArtifactoryConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the owner type. | 
**internalOwnerId** | **string** | Enter the internal ID of the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOwnerArtifactoryConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inherit** | **bool** | Specify whether to include details from an inherited Artifactory connection. | [default to false]

### Return type

[**ApiOwnerArtifactoryConnectionDTO**](ApiOwnerArtifactoryConnectionDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestArtifactoryConnection

> ApiStatusDTO TestArtifactoryConnection(ctx, ownerType, internalOwnerId).ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO).Execute()





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
	ownerType := "ownerType_example" // string | Select the owner type.
	internalOwnerId := "internalOwnerId_example" // string | Enter the internal ID of the owner.
	apiArtifactoryConnectionDTO := *sonatypeiq.NewApiArtifactoryConnectionDTO() // ApiArtifactoryConnectionDTO | Enter values for the Artifactory connection.<ul><li>`baseUrl` is the baseURL of the Artifactory instance.</li><li>`username` and `password` to authenticate the Artifactory connection.</li></ul>

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureArtifactoryConnectionAPI.TestArtifactoryConnection(context.Background(), ownerType, internalOwnerId).ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureArtifactoryConnectionAPI.TestArtifactoryConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestArtifactoryConnection`: ApiStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigureArtifactoryConnectionAPI.TestArtifactoryConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the owner type. | 
**internalOwnerId** | **string** | Enter the internal ID of the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestArtifactoryConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiArtifactoryConnectionDTO** | [**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md) | Enter values for the Artifactory connection.&lt;ul&gt;&lt;li&gt;&#x60;baseUrl&#x60; is the baseURL of the Artifactory instance.&lt;/li&gt;&lt;li&gt;&#x60;username&#x60; and &#x60;password&#x60; to authenticate the Artifactory connection.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**ApiStatusDTO**](ApiStatusDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestArtifactoryConnection1

> ApiStatusDTO TestArtifactoryConnection1(ctx, ownerType, internalOwnerId, artifactoryConnectionId).Execute()





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
	ownerType := "ownerType_example" // string | Enter the owner type.
	internalOwnerId := "internalOwnerId_example" // string | Enter the internal ID of the owner.
	artifactoryConnectionId := "artifactoryConnectionId_example" // string | Enter the Artifactory connection ID.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureArtifactoryConnectionAPI.TestArtifactoryConnection1(context.Background(), ownerType, internalOwnerId, artifactoryConnectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureArtifactoryConnectionAPI.TestArtifactoryConnection1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TestArtifactoryConnection1`: ApiStatusDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigureArtifactoryConnectionAPI.TestArtifactoryConnection1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Enter the owner type. | 
**internalOwnerId** | **string** | Enter the internal ID of the owner. | 
**artifactoryConnectionId** | **string** | Enter the Artifactory connection ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestArtifactoryConnection1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApiStatusDTO**](ApiStatusDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateArtifactoryConnection

> ApiArtifactoryConnectionDTO UpdateArtifactoryConnection(ctx, ownerType, internalOwnerId, artifactoryConnectionId).ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO).Execute()





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
	ownerType := "ownerType_example" // string | Specify the owner type.
	internalOwnerId := "internalOwnerId_example" // string | Enter the internal ID of the owner.
	artifactoryConnectionId := "artifactoryConnectionId_example" // string | Enter the Artifactory connection ID.
	apiArtifactoryConnectionDTO := *sonatypeiq.NewApiArtifactoryConnectionDTO() // ApiArtifactoryConnectionDTO | Enter values for the new Artifactory connection.<ul><li>`isAnonymous` indicates if the connection is anonymous.</li><li>`baseUrl` is the baseURL of the Artifactory instance.</li><li>`username` and `password` to authenticate the Artifactory connection.</li></ul>

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigureArtifactoryConnectionAPI.UpdateArtifactoryConnection(context.Background(), ownerType, internalOwnerId, artifactoryConnectionId).ApiArtifactoryConnectionDTO(apiArtifactoryConnectionDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureArtifactoryConnectionAPI.UpdateArtifactoryConnection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateArtifactoryConnection`: ApiArtifactoryConnectionDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigureArtifactoryConnectionAPI.UpdateArtifactoryConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Specify the owner type. | 
**internalOwnerId** | **string** | Enter the internal ID of the owner. | 
**artifactoryConnectionId** | **string** | Enter the Artifactory connection ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateArtifactoryConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **apiArtifactoryConnectionDTO** | [**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md) | Enter values for the new Artifactory connection.&lt;ul&gt;&lt;li&gt;&#x60;isAnonymous&#x60; indicates if the connection is anonymous.&lt;/li&gt;&lt;li&gt;&#x60;baseUrl&#x60; is the baseURL of the Artifactory instance.&lt;/li&gt;&lt;li&gt;&#x60;username&#x60; and &#x60;password&#x60; to authenticate the Artifactory connection.&lt;/li&gt;&lt;/ul&gt; | 

### Return type

[**ApiArtifactoryConnectionDTO**](ApiArtifactoryConnectionDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOwnerArtifactoryConnectionStatus

> UpdateOwnerArtifactoryConnectionStatus(ctx, ownerType, internalOwnerId).ApiArtifactoryConnectionStatusRequestDTO(apiArtifactoryConnectionStatusRequestDTO).Execute()





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
	ownerType := "ownerType_example" // string | Select the owner type.
	internalOwnerId := "internalOwnerId_example" // string | Enter the internal ID of the owner.
	apiArtifactoryConnectionStatusRequestDTO := *sonatypeiq.NewApiArtifactoryConnectionStatusRequestDTO() // ApiArtifactoryConnectionStatusRequestDTO | Set values for the connection properties `enabled` and `allowOverride`.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigureArtifactoryConnectionAPI.UpdateOwnerArtifactoryConnectionStatus(context.Background(), ownerType, internalOwnerId).ApiArtifactoryConnectionStatusRequestDTO(apiArtifactoryConnectionStatusRequestDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigureArtifactoryConnectionAPI.UpdateOwnerArtifactoryConnectionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the owner type. | 
**internalOwnerId** | **string** | Enter the internal ID of the owner. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOwnerArtifactoryConnectionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiArtifactoryConnectionStatusRequestDTO** | [**ApiArtifactoryConnectionStatusRequestDTO**](ApiArtifactoryConnectionStatusRequestDTO.md) | Set values for the connection properties &#x60;enabled&#x60; and &#x60;allowOverride&#x60;. | 

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

