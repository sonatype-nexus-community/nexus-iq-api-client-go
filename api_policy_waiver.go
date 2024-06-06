/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.177.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// PolicyWaiverAPIService PolicyWaiverAPI service
type PolicyWaiverAPIService service

type ApiAddPolicyWaiverRequest struct {
	ctx context.Context
	ApiService *PolicyWaiverAPIService
	policyViolationId string
	ownerType string
	body *string
}

func (r ApiAddPolicyWaiverRequest) Body(body string) ApiAddPolicyWaiverRequest {
	r.body = &body
	return r
}

func (r ApiAddPolicyWaiverRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddPolicyWaiverExecute(r)
}

/*
AddPolicyWaiver Method for AddPolicyWaiver

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param policyViolationId
 @param ownerType
 @return ApiAddPolicyWaiverRequest

Deprecated
*/
func (a *PolicyWaiverAPIService) AddPolicyWaiver(ctx context.Context, policyViolationId string, ownerType string) ApiAddPolicyWaiverRequest {
	return ApiAddPolicyWaiverRequest{
		ApiService: a,
		ctx: ctx,
		policyViolationId: policyViolationId,
		ownerType: ownerType,
	}
}

// Execute executes the request
// Deprecated
func (a *PolicyWaiverAPIService) AddPolicyWaiverExecute(r ApiAddPolicyWaiverRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PolicyWaiverAPIService.AddPolicyWaiver")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/policyWaiver/{policyViolationId}/{ownerType}"
	localVarPath = strings.Replace(localVarPath, "{"+"policyViolationId"+"}", url.PathEscape(parameterValueToString(r.policyViolationId, "policyViolationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"*/*"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
