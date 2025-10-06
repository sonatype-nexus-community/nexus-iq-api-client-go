# \PolicyWaiverReasonsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPolicyWaiverReasons**](PolicyWaiverReasonsAPI.md#GetPolicyWaiverReasons) | **Get** /api/v2/policyWaiverReasons | 



## GetPolicyWaiverReasons

> []ApiPolicyWaiverReasonDTO GetPolicyWaiverReasons(ctx).Execute()





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
	resp, r, err := apiClient.PolicyWaiverReasonsAPI.GetPolicyWaiverReasons(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PolicyWaiverReasonsAPI.GetPolicyWaiverReasons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicyWaiverReasons`: []ApiPolicyWaiverReasonDTO
	fmt.Fprintf(os.Stdout, "Response from `PolicyWaiverReasonsAPI.GetPolicyWaiverReasons`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyWaiverReasonsRequest struct via the builder pattern


### Return type

[**[]ApiPolicyWaiverReasonDTO**](ApiPolicyWaiverReasonDTO.md)

### Authorization

[BasicAuth](../README.md#BasicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

