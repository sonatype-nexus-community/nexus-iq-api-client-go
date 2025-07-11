# \ComponentLabelsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLabel**](ComponentLabelsAPI.md#AddLabel) | **Post** /api/v2/labels/{ownerType}/{ownerId} | 
[**DeleteLabel**](ComponentLabelsAPI.md#DeleteLabel) | **Delete** /api/v2/labels/{ownerType}/{ownerId}/{labelId} | 
[**GetApplicableContexts**](ComponentLabelsAPI.md#GetApplicableContexts) | **Get** /api/v2/labels/{ownerType}/{ownerId}/applicable/context/{labelId} | 
[**GetApplicableLabels**](ComponentLabelsAPI.md#GetApplicableLabels) | **Get** /api/v2/labels/{ownerType}/{ownerId}/applicable | 
[**GetLabels**](ComponentLabelsAPI.md#GetLabels) | **Get** /api/v2/labels/{ownerType}/{ownerId} | 
[**UpdateLabel**](ComponentLabelsAPI.md#UpdateLabel) | **Put** /api/v2/labels/{ownerType}/{ownerId} | 



## AddLabel

> ApiLabelDTO AddLabel(ctx, ownerType, ownerId).ApiLabelDTO(apiLabelDTO).Execute()





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
	ownerType := "ownerType_example" // string | Select the ownerType to which the label will be assigned.
	ownerId := "ownerId_example" // string | Enter the id for the selected ownerType.
	apiLabelDTO := *sonatypeiq.NewApiLabelDTO() // ApiLabelDTO | Specify a label name, description and color for the label. Valid values for color are `light-red` , `light-green` , `light-blue` , `light-purple`, `dark-red` , `dark-green` , `dark-blue` , `dark-purple` , `orange` , `yellow`. Do not enter value for the `id` field. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentLabelsAPI.AddLabel(context.Background(), ownerType, ownerId).ApiLabelDTO(apiLabelDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentLabelsAPI.AddLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddLabel`: ApiLabelDTO
	fmt.Fprintf(os.Stdout, "Response from `ComponentLabelsAPI.AddLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the ownerType to which the label will be assigned. | 
**ownerId** | **string** | Enter the id for the selected ownerType. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiLabelDTO** | [**ApiLabelDTO**](ApiLabelDTO.md) | Specify a label name, description and color for the label. Valid values for color are &#x60;light-red&#x60; , &#x60;light-green&#x60; , &#x60;light-blue&#x60; , &#x60;light-purple&#x60;, &#x60;dark-red&#x60; , &#x60;dark-green&#x60; , &#x60;dark-blue&#x60; , &#x60;dark-purple&#x60; , &#x60;orange&#x60; , &#x60;yellow&#x60;. Do not enter value for the &#x60;id&#x60; field. | 

### Return type

[**ApiLabelDTO**](ApiLabelDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLabel

> DeleteLabel(ctx, ownerType, ownerId, labelId).Execute()





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
	ownerType := "ownerType_example" // string | Select the ownerType for which the label will be deleted.
	ownerId := "ownerId_example" // string | Enter the id for the selected ownerType.
	labelId := "labelId_example" // string | Enter the id for the label to be deleted.

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	r, err := apiClient.ComponentLabelsAPI.DeleteLabel(context.Background(), ownerType, ownerId, labelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentLabelsAPI.DeleteLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the ownerType for which the label will be deleted. | 
**ownerId** | **string** | Enter the id for the selected ownerType. | 
**labelId** | **string** | Enter the id for the label to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLabelRequest struct via the builder pattern


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


## GetApplicableContexts

> ApplicableContext GetApplicableContexts(ctx, ownerType, ownerId, labelId).Execute()





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
	ownerType := "ownerType_example" // string | Select the ownerType.
	ownerId := "ownerId_example" // string | Enter the ownerId
	labelId := "labelId_example" // string | Enter the labelId

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentLabelsAPI.GetApplicableContexts(context.Background(), ownerType, ownerId, labelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentLabelsAPI.GetApplicableContexts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicableContexts`: ApplicableContext
	fmt.Fprintf(os.Stdout, "Response from `ComponentLabelsAPI.GetApplicableContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the ownerType. | 
**ownerId** | **string** | Enter the ownerId | 
**labelId** | **string** | Enter the labelId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ApplicableContext**](ApplicableContext.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicableLabels

> ApplicableLabels GetApplicableLabels(ctx, ownerType, ownerId).Execute()





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
	ownerType := "ownerType_example" // string | Select the ownerType to retrieve the component label information for.
	ownerId := "ownerId_example" // string | Enter the id for the application, organization or repository

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentLabelsAPI.GetApplicableLabels(context.Background(), ownerType, ownerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentLabelsAPI.GetApplicableLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicableLabels`: ApplicableLabels
	fmt.Fprintf(os.Stdout, "Response from `ComponentLabelsAPI.GetApplicableLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the ownerType to retrieve the component label information for. | 
**ownerId** | **string** | Enter the id for the application, organization or repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicableLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApplicableLabels**](ApplicableLabels.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLabels

> []ApiLabelDTO GetLabels(ctx, ownerType, ownerId).Inherit(inherit).Execute()





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
	ownerType := "ownerType_example" // string | Select the `ownerType` for which you want to retrieve the component label information.
	ownerId := "ownerId_example" // string | Enter the id of the application, organization or the repository.
	inherit := true // bool | Set to `true` to retrieve inherited component labels. (optional) (default to false)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentLabelsAPI.GetLabels(context.Background(), ownerType, ownerId).Inherit(inherit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentLabelsAPI.GetLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabels`: []ApiLabelDTO
	fmt.Fprintf(os.Stdout, "Response from `ComponentLabelsAPI.GetLabels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the &#x60;ownerType&#x60; for which you want to retrieve the component label information. | 
**ownerId** | **string** | Enter the id of the application, organization or the repository. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **inherit** | **bool** | Set to &#x60;true&#x60; to retrieve inherited component labels. | [default to false]

### Return type

[**[]ApiLabelDTO**](ApiLabelDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLabel

> ApiLabelDTO UpdateLabel(ctx, ownerType, ownerId).ApiLabelDTO(apiLabelDTO).Execute()





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
	ownerType := "ownerType_example" // string | Select the ownerType for which the label will be updated.
	ownerId := "ownerId_example" // string | Enter the id for the selected ownerType.
	apiLabelDTO := *sonatypeiq.NewApiLabelDTO() // ApiLabelDTO | Specify the new values for label name, description, color and the corresponding label id for the component label to be updated. Valid values for color are `light-red` , `light-green` , `light-blue` , `light--purple`, `dark-red` , `dark-green` ,`dark-blue` , `dark-purple` ,`orange` , `yellow`. (optional)

	configuration := sonatypeiq.NewConfiguration()
	apiClient := sonatypeiq.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentLabelsAPI.UpdateLabel(context.Background(), ownerType, ownerId).ApiLabelDTO(apiLabelDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentLabelsAPI.UpdateLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLabel`: ApiLabelDTO
	fmt.Fprintf(os.Stdout, "Response from `ComponentLabelsAPI.UpdateLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerType** | **string** | Select the ownerType for which the label will be updated. | 
**ownerId** | **string** | Enter the id for the selected ownerType. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiLabelDTO** | [**ApiLabelDTO**](ApiLabelDTO.md) | Specify the new values for label name, description, color and the corresponding label id for the component label to be updated. Valid values for color are &#x60;light-red&#x60; , &#x60;light-green&#x60; , &#x60;light-blue&#x60; , &#x60;light--purple&#x60;, &#x60;dark-red&#x60; , &#x60;dark-green&#x60; ,&#x60;dark-blue&#x60; , &#x60;dark-purple&#x60; ,&#x60;orange&#x60; , &#x60;yellow&#x60;. | 

### Return type

[**ApiLabelDTO**](ApiLabelDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

