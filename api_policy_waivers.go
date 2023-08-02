/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 165
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


// PolicyWaiversAPIService PolicyWaiversAPI service
type PolicyWaiversAPIService service

type ApiAddPolicyWaiverByPolicyViolationIdRequest struct {
	ctx context.Context
	ApiService *PolicyWaiversAPIService
	ownerType string
	ownerId string
	policyViolationId string
	apiWaiverOptionsDTO *ApiWaiverOptionsDTO
}

func (r ApiAddPolicyWaiverByPolicyViolationIdRequest) ApiWaiverOptionsDTO(apiWaiverOptionsDTO ApiWaiverOptionsDTO) ApiAddPolicyWaiverByPolicyViolationIdRequest {
	r.apiWaiverOptionsDTO = &apiWaiverOptionsDTO
	return r
}

func (r ApiAddPolicyWaiverByPolicyViolationIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddPolicyWaiverByPolicyViolationIdExecute(r)
}

/*
AddPolicyWaiverByPolicyViolationId Method for AddPolicyWaiverByPolicyViolationId

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @param policyViolationId
 @return ApiAddPolicyWaiverByPolicyViolationIdRequest
*/
func (a *PolicyWaiversAPIService) AddPolicyWaiverByPolicyViolationId(ctx context.Context, ownerType string, ownerId string, policyViolationId string) ApiAddPolicyWaiverByPolicyViolationIdRequest {
	return ApiAddPolicyWaiverByPolicyViolationIdRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
		policyViolationId: policyViolationId,
	}
}

// Execute executes the request
func (a *PolicyWaiversAPIService) AddPolicyWaiverByPolicyViolationIdExecute(r ApiAddPolicyWaiverByPolicyViolationIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PolicyWaiversAPIService.AddPolicyWaiverByPolicyViolationId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/policyWaivers/{ownerType}/{ownerId}/{policyViolationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyViolationId"+"}", url.PathEscape(parameterValueToString(r.policyViolationId, "policyViolationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.apiWaiverOptionsDTO
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

type ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest struct {
	ctx context.Context
	ApiService *PolicyWaiversAPIService
	ownerType string
	ownerId string
	scanId string
	componentIdentifier *ComponentIdentifier
	packageUrl *string
	hash *string
	apiWaiverOptionsDTO *ApiWaiverOptionsDTO
}

func (r ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest) ComponentIdentifier(componentIdentifier ComponentIdentifier) ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest {
	r.componentIdentifier = &componentIdentifier
	return r
}

func (r ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest) PackageUrl(packageUrl string) ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest {
	r.packageUrl = &packageUrl
	return r
}

func (r ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest) Hash(hash string) ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest {
	r.hash = &hash
	return r
}

func (r ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest) ApiWaiverOptionsDTO(apiWaiverOptionsDTO ApiWaiverOptionsDTO) ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest {
	r.apiWaiverOptionsDTO = &apiWaiverOptionsDTO
	return r
}

func (r ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddWaiverToTransitivePolicyViolationsByAppScanComponentExecute(r)
}

/*
AddWaiverToTransitivePolicyViolationsByAppScanComponent Method for AddWaiverToTransitivePolicyViolationsByAppScanComponent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @param scanId
 @return ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest
*/
func (a *PolicyWaiversAPIService) AddWaiverToTransitivePolicyViolationsByAppScanComponent(ctx context.Context, ownerType string, ownerId string, scanId string) ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest {
	return ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
		scanId: scanId,
	}
}

// Execute executes the request
func (a *PolicyWaiversAPIService) AddWaiverToTransitivePolicyViolationsByAppScanComponentExecute(r ApiAddWaiverToTransitivePolicyViolationsByAppScanComponentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PolicyWaiversAPIService.AddWaiverToTransitivePolicyViolationsByAppScanComponent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/policyWaivers/transitive/{ownerType}/{ownerId}/{scanId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scanId"+"}", url.PathEscape(parameterValueToString(r.scanId, "scanId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.componentIdentifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "componentIdentifier", r.componentIdentifier, "")
	}
	if r.packageUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "packageUrl", r.packageUrl, "")
	}
	if r.hash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hash", r.hash, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.apiWaiverOptionsDTO
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

type ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest struct {
	ctx context.Context
	ApiService *PolicyWaiversAPIService
	ownerType string
	ownerId string
	stageId string
	componentIdentifier *ComponentIdentifier
	packageUrl *string
	hash *string
	apiWaiverOptionsDTO *ApiWaiverOptionsDTO
}

func (r ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest) ComponentIdentifier(componentIdentifier ComponentIdentifier) ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest {
	r.componentIdentifier = &componentIdentifier
	return r
}

func (r ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest) PackageUrl(packageUrl string) ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest {
	r.packageUrl = &packageUrl
	return r
}

func (r ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest) Hash(hash string) ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest {
	r.hash = &hash
	return r
}

func (r ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest) ApiWaiverOptionsDTO(apiWaiverOptionsDTO ApiWaiverOptionsDTO) ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest {
	r.apiWaiverOptionsDTO = &apiWaiverOptionsDTO
	return r
}

func (r ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddWaiverToTransitivePolicyViolationsByOwnerStageComponentExecute(r)
}

/*
AddWaiverToTransitivePolicyViolationsByOwnerStageComponent Method for AddWaiverToTransitivePolicyViolationsByOwnerStageComponent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @param stageId
 @return ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest
*/
func (a *PolicyWaiversAPIService) AddWaiverToTransitivePolicyViolationsByOwnerStageComponent(ctx context.Context, ownerType string, ownerId string, stageId string) ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest {
	return ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
		stageId: stageId,
	}
}

// Execute executes the request
func (a *PolicyWaiversAPIService) AddWaiverToTransitivePolicyViolationsByOwnerStageComponentExecute(r ApiAddWaiverToTransitivePolicyViolationsByOwnerStageComponentRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PolicyWaiversAPIService.AddWaiverToTransitivePolicyViolationsByOwnerStageComponent")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/policyWaivers/transitive/{ownerType}/{ownerId}/stages/{stageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stageId"+"}", url.PathEscape(parameterValueToString(r.stageId, "stageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.componentIdentifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "componentIdentifier", r.componentIdentifier, "")
	}
	if r.packageUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "packageUrl", r.packageUrl, "")
	}
	if r.hash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hash", r.hash, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.apiWaiverOptionsDTO
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

type ApiDeletePolicyWaiverRequest struct {
	ctx context.Context
	ApiService *PolicyWaiversAPIService
	ownerType string
	ownerId string
	policyWaiverId string
}

func (r ApiDeletePolicyWaiverRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePolicyWaiverExecute(r)
}

/*
DeletePolicyWaiver Method for DeletePolicyWaiver

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @param policyWaiverId
 @return ApiDeletePolicyWaiverRequest
*/
func (a *PolicyWaiversAPIService) DeletePolicyWaiver(ctx context.Context, ownerType string, ownerId string, policyWaiverId string) ApiDeletePolicyWaiverRequest {
	return ApiDeletePolicyWaiverRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
		policyWaiverId: policyWaiverId,
	}
}

// Execute executes the request
func (a *PolicyWaiversAPIService) DeletePolicyWaiverExecute(r ApiDeletePolicyWaiverRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PolicyWaiversAPIService.DeletePolicyWaiver")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/policyWaivers/{ownerType}/{ownerId}/{policyWaiverId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyWaiverId"+"}", url.PathEscape(parameterValueToString(r.policyWaiverId, "policyWaiverId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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

type ApiGetPolicyWaiverRequest struct {
	ctx context.Context
	ApiService *PolicyWaiversAPIService
	ownerType string
	ownerId string
	policyWaiverId string
}

func (r ApiGetPolicyWaiverRequest) Execute() (*ApiPolicyWaiverDTO, *http.Response, error) {
	return r.ApiService.GetPolicyWaiverExecute(r)
}

/*
GetPolicyWaiver Method for GetPolicyWaiver

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @param policyWaiverId
 @return ApiGetPolicyWaiverRequest
*/
func (a *PolicyWaiversAPIService) GetPolicyWaiver(ctx context.Context, ownerType string, ownerId string, policyWaiverId string) ApiGetPolicyWaiverRequest {
	return ApiGetPolicyWaiverRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
		policyWaiverId: policyWaiverId,
	}
}

// Execute executes the request
//  @return ApiPolicyWaiverDTO
func (a *PolicyWaiversAPIService) GetPolicyWaiverExecute(r ApiGetPolicyWaiverRequest) (*ApiPolicyWaiverDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiPolicyWaiverDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PolicyWaiversAPIService.GetPolicyWaiver")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/policyWaivers/{ownerType}/{ownerId}/{policyWaiverId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"policyWaiverId"+"}", url.PathEscape(parameterValueToString(r.policyWaiverId, "policyWaiverId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ApiPolicyWaiverDTO
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetPolicyWaiversRequest struct {
	ctx context.Context
	ApiService *PolicyWaiversAPIService
	ownerType string
	ownerId string
}

func (r ApiGetPolicyWaiversRequest) Execute() ([]ApiPolicyWaiverDTO, *http.Response, error) {
	return r.ApiService.GetPolicyWaiversExecute(r)
}

/*
GetPolicyWaivers Method for GetPolicyWaivers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @return ApiGetPolicyWaiversRequest
*/
func (a *PolicyWaiversAPIService) GetPolicyWaivers(ctx context.Context, ownerType string, ownerId string) ApiGetPolicyWaiversRequest {
	return ApiGetPolicyWaiversRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
	}
}

// Execute executes the request
//  @return []ApiPolicyWaiverDTO
func (a *PolicyWaiversAPIService) GetPolicyWaiversExecute(r ApiGetPolicyWaiversRequest) ([]ApiPolicyWaiverDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiPolicyWaiverDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PolicyWaiversAPIService.GetPolicyWaivers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/policyWaivers/{ownerType}/{ownerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v []ApiPolicyWaiverDTO
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTransitivePolicyWaiversByAppScanComponentRequest struct {
	ctx context.Context
	ApiService *PolicyWaiversAPIService
	ownerType string
	ownerId string
	scanId string
	componentIdentifier *ComponentIdentifier
	packageUrl *string
	hash *string
}

func (r ApiGetTransitivePolicyWaiversByAppScanComponentRequest) ComponentIdentifier(componentIdentifier ComponentIdentifier) ApiGetTransitivePolicyWaiversByAppScanComponentRequest {
	r.componentIdentifier = &componentIdentifier
	return r
}

func (r ApiGetTransitivePolicyWaiversByAppScanComponentRequest) PackageUrl(packageUrl string) ApiGetTransitivePolicyWaiversByAppScanComponentRequest {
	r.packageUrl = &packageUrl
	return r
}

func (r ApiGetTransitivePolicyWaiversByAppScanComponentRequest) Hash(hash string) ApiGetTransitivePolicyWaiversByAppScanComponentRequest {
	r.hash = &hash
	return r
}

func (r ApiGetTransitivePolicyWaiversByAppScanComponentRequest) Execute() (*ApiComponentPolicyWaiversDTO, *http.Response, error) {
	return r.ApiService.GetTransitivePolicyWaiversByAppScanComponentExecute(r)
}

/*
GetTransitivePolicyWaiversByAppScanComponent Method for GetTransitivePolicyWaiversByAppScanComponent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @param scanId
 @return ApiGetTransitivePolicyWaiversByAppScanComponentRequest
*/
func (a *PolicyWaiversAPIService) GetTransitivePolicyWaiversByAppScanComponent(ctx context.Context, ownerType string, ownerId string, scanId string) ApiGetTransitivePolicyWaiversByAppScanComponentRequest {
	return ApiGetTransitivePolicyWaiversByAppScanComponentRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
		scanId: scanId,
	}
}

// Execute executes the request
//  @return ApiComponentPolicyWaiversDTO
func (a *PolicyWaiversAPIService) GetTransitivePolicyWaiversByAppScanComponentExecute(r ApiGetTransitivePolicyWaiversByAppScanComponentRequest) (*ApiComponentPolicyWaiversDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiComponentPolicyWaiversDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PolicyWaiversAPIService.GetTransitivePolicyWaiversByAppScanComponent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/policyWaivers/transitive/{ownerType}/{ownerId}/{scanId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scanId"+"}", url.PathEscape(parameterValueToString(r.scanId, "scanId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.componentIdentifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "componentIdentifier", r.componentIdentifier, "")
	}
	if r.packageUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "packageUrl", r.packageUrl, "")
	}
	if r.hash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hash", r.hash, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ApiComponentPolicyWaiversDTO
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}