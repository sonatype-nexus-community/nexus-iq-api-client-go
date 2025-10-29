# \ConfigSAMLAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSamlConfiguration**](ConfigSAMLAPI.md#DeleteSamlConfiguration) | **Delete** /api/v2/config/saml | 
[**GetMetadata**](ConfigSAMLAPI.md#GetMetadata) | **Get** /api/v2/config/saml/metadata | 
[**GetSamlConfiguration**](ConfigSAMLAPI.md#GetSamlConfiguration) | **Get** /api/v2/config/saml | 
[**InsertOrUpdateSamlConfiguration**](ConfigSAMLAPI.md#InsertOrUpdateSamlConfiguration) | **Put** /api/v2/config/saml | 



## DeleteSamlConfiguration

> DeleteSamlConfiguration(ctx).Execute()





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
	r, err := apiClient.ConfigSAMLAPI.DeleteSamlConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigSAMLAPI.DeleteSamlConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSamlConfigurationRequest struct via the builder pattern


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


## GetMetadata

> string GetMetadata(ctx).Execute()





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
	resp, r, err := apiClient.ConfigSAMLAPI.GetMetadata(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigSAMLAPI.GetMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetadata`: string
	fmt.Fprintf(os.Stdout, "Response from `ConfigSAMLAPI.GetMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataRequest struct via the builder pattern


### Return type

**string**

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSamlConfiguration

> ApiSamlConfigurationResponseDTO GetSamlConfiguration(ctx).Execute()





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
	resp, r, err := apiClient.ConfigSAMLAPI.GetSamlConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigSAMLAPI.GetSamlConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSamlConfiguration`: ApiSamlConfigurationResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `ConfigSAMLAPI.GetSamlConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSamlConfigurationRequest struct via the builder pattern


### Return type

[**ApiSamlConfigurationResponseDTO**](ApiSamlConfigurationResponseDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InsertOrUpdateSamlConfiguration

> InsertOrUpdateSamlConfiguration(ctx).IdentityProviderXml(identityProviderXml).SamlConfiguration(samlConfiguration).Execute()





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
	identityProviderXml := os.NewFile(1234, "some_file") // *os.File | Enter the SAML metadata XML of your IdP. Refer to the IdP documentation to obtain this metadata.
	samlConfiguration := *sonatypeiq.NewApiSamlConfigurationDTO() // ApiSamlConfigurationDTO | 

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ConfigSAMLAPI.InsertOrUpdateSamlConfiguration(context.Background()).IdentityProviderXml(identityProviderXml).SamlConfiguration(samlConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigSAMLAPI.InsertOrUpdateSamlConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInsertOrUpdateSamlConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityProviderXml** | ***os.File** | Enter the SAML metadata XML of your IdP. Refer to the IdP documentation to obtain this metadata. | 
 **samlConfiguration** | [**ApiSamlConfigurationDTO**](ApiSamlConfigurationDTO.md) |  | 

### Return type

 (empty response body)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

