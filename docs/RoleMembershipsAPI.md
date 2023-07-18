# \RoleMembershipsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRoleMembershipsApplicationOrOrganization**](RoleMembershipsAPI.md#GetRoleMembershipsApplicationOrOrganization) | **Get** /api/v2/roleMemberships/{ownerType}/{internalOwnerId} | 
[**GetRoleMembershipsGlobalOrRepositoryContainer**](RoleMembershipsAPI.md#GetRoleMembershipsGlobalOrRepositoryContainer) | **Get** /api/v2/roleMemberships/{ownerType} | 
[**GrantRoleMembershipApplicationOrOrganization**](RoleMembershipsAPI.md#GrantRoleMembershipApplicationOrOrganization) | **Put** /api/v2/roleMemberships/{ownerType}/{internalOwnerId}/role/{roleId}/{memberType}/{memberName} | 
[**GrantRoleMembershipGlobalOrRepositoryContainer**](RoleMembershipsAPI.md#GrantRoleMembershipGlobalOrRepositoryContainer) | **Put** /api/v2/roleMemberships/{ownerType}/role/{roleId}/{memberType}/{memberName} | 
[**RevokeRoleMembershipApplicationOrOrganization**](RoleMembershipsAPI.md#RevokeRoleMembershipApplicationOrOrganization) | **Delete** /api/v2/roleMemberships/{ownerType}/{internalOwnerId}/role/{roleId}/{memberType}/{memberName} | 
[**RevokeRoleMembershipGlobalOrRepositoryContainer**](RoleMembershipsAPI.md#RevokeRoleMembershipGlobalOrRepositoryContainer) | **Delete** /api/v2/roleMemberships/{ownerType}/role/{roleId}/{memberType}/{memberName} | 



## GetRoleMembershipsApplicationOrOrganization

> ApiRoleMemberMappingListDTO GetRoleMembershipsApplicationOrOrganization(ctx, ownerType, internalOwnerId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    iqclient "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {
    ownerType := "ownerType_example" // string | 
    internalOwnerId := "internalOwnerId_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleMembershipsAPI.GetRoleMembershipsApplicationOrOrganization(context.Background(), ownerType, internalOwnerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleMembershipsAPI.GetRoleMembershipsApplicationOrOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleMembershipsApplicationOrOrganization`: ApiRoleMemberMappingListDTO
    fmt.Fprintf(os.Stdout, "Response from `RoleMembershipsAPI.GetRoleMembershipsApplicationOrOrganization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleMembershipsApplicationOrOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiRoleMemberMappingListDTO**](ApiRoleMemberMappingListDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRoleMembershipsGlobalOrRepositoryContainer

> ApiRoleMemberMappingListDTO GetRoleMembershipsGlobalOrRepositoryContainer(ctx, ownerType).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    iqclient "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {
    ownerType := "ownerType_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleMembershipsAPI.GetRoleMembershipsGlobalOrRepositoryContainer(context.Background(), ownerType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleMembershipsAPI.GetRoleMembershipsGlobalOrRepositoryContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRoleMembershipsGlobalOrRepositoryContainer`: ApiRoleMemberMappingListDTO
    fmt.Fprintf(os.Stdout, "Response from `RoleMembershipsAPI.GetRoleMembershipsGlobalOrRepositoryContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRoleMembershipsGlobalOrRepositoryContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiRoleMemberMappingListDTO**](ApiRoleMemberMappingListDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantRoleMembershipApplicationOrOrganization

> GrantRoleMembershipApplicationOrOrganization(ctx, ownerType, internalOwnerId, roleId, memberType, memberName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    iqclient "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {
    ownerType := "ownerType_example" // string | 
    internalOwnerId := "internalOwnerId_example" // string | 
    roleId := "roleId_example" // string | 
    memberType := "memberType_example" // string | 
    memberName := "memberName_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.RoleMembershipsAPI.GrantRoleMembershipApplicationOrOrganization(context.Background(), ownerType, internalOwnerId, roleId, memberType, memberName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleMembershipsAPI.GrantRoleMembershipApplicationOrOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 
**roleId** | **string** |  | 
**memberType** | **string** |  | 
**memberName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantRoleMembershipApplicationOrOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantRoleMembershipGlobalOrRepositoryContainer

> GrantRoleMembershipGlobalOrRepositoryContainer(ctx, ownerType, roleId, memberType, memberName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    iqclient "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {
    ownerType := "ownerType_example" // string | 
    roleId := "roleId_example" // string | 
    memberType := "memberType_example" // string | 
    memberName := "memberName_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.RoleMembershipsAPI.GrantRoleMembershipGlobalOrRepositoryContainer(context.Background(), ownerType, roleId, memberType, memberName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleMembershipsAPI.GrantRoleMembershipGlobalOrRepositoryContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**roleId** | **string** |  | 
**memberType** | **string** |  | 
**memberName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGrantRoleMembershipGlobalOrRepositoryContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeRoleMembershipApplicationOrOrganization

> RevokeRoleMembershipApplicationOrOrganization(ctx, ownerType, internalOwnerId, roleId, memberType, memberName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    iqclient "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {
    ownerType := "ownerType_example" // string | 
    internalOwnerId := "internalOwnerId_example" // string | 
    roleId := "roleId_example" // string | 
    memberType := "memberType_example" // string | 
    memberName := "memberName_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.RoleMembershipsAPI.RevokeRoleMembershipApplicationOrOrganization(context.Background(), ownerType, internalOwnerId, roleId, memberType, memberName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleMembershipsAPI.RevokeRoleMembershipApplicationOrOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**internalOwnerId** | **string** |  | 
**roleId** | **string** |  | 
**memberType** | **string** |  | 
**memberName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeRoleMembershipApplicationOrOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeRoleMembershipGlobalOrRepositoryContainer

> RevokeRoleMembershipGlobalOrRepositoryContainer(ctx, ownerType, roleId, memberType, memberName).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    iqclient "github.com/sonatype-nexus-community/nexus-iq-api-client-go"
)

func main() {
    ownerType := "ownerType_example" // string | 
    roleId := "roleId_example" // string | 
    memberType := "memberType_example" // string | 
    memberName := "memberName_example" // string | 

    configuration := iqclient.NewConfiguration()
    apiClient := iqclient.NewAPIClient(configuration)
    r, err := apiClient.RoleMembershipsAPI.RevokeRoleMembershipGlobalOrRepositoryContainer(context.Background(), ownerType, roleId, memberType, memberName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleMembershipsAPI.RevokeRoleMembershipGlobalOrRepositoryContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** |  | 
**roleId** | **string** |  | 
**memberType** | **string** |  | 
**memberName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeRoleMembershipGlobalOrRepositoryContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

