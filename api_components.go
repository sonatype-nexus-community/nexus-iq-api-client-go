/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.184.0-01
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


// ComponentsAPIService ComponentsAPI service
type ComponentsAPIService service

type ApiDeleteComponentLabelRequest struct {
	ctx context.Context
	ApiService *ComponentsAPIService
	ownerType string
	internalOwnerId string
	componentHash string
	labelName string
}

func (r ApiDeleteComponentLabelRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteComponentLabelExecute(r)
}

/*
DeleteComponentLabel Method for DeleteComponentLabel

Use this method to un-assign a label from a component.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType Possible values: application or organization
 @param internalOwnerId Possible values : applicationId or organizationId
 @param componentHash Enter the SHA1 hash of the component.
 @param labelName Enter the label name to un-assign from this component.
 @return ApiDeleteComponentLabelRequest
*/
func (a *ComponentsAPIService) DeleteComponentLabel(ctx context.Context, ownerType string, internalOwnerId string, componentHash string, labelName string) ApiDeleteComponentLabelRequest {
	return ApiDeleteComponentLabelRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
		componentHash: componentHash,
		labelName: labelName,
	}
}

// Execute executes the request
func (a *ComponentsAPIService) DeleteComponentLabelExecute(r ApiDeleteComponentLabelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentsAPIService.DeleteComponentLabel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/components/{componentHash}/labels/{labelName}/{ownerType}s/{internalOwnerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"componentHash"+"}", url.PathEscape(parameterValueToString(r.componentHash, "componentHash")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"labelName"+"}", url.PathEscape(parameterValueToString(r.labelName, "labelName")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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

type ApiGetComponentDetailsRequest struct {
	ctx context.Context
	ApiService *ComponentsAPIService
	apiComponentDetailsRequestDTOV2 *ApiComponentDetailsRequestDTOV2
}

// You can retrieve component data in any one of the 3 ways via: 1. Component identifier 2. Package URL 3. Hash
func (r ApiGetComponentDetailsRequest) ApiComponentDetailsRequestDTOV2(apiComponentDetailsRequestDTOV2 ApiComponentDetailsRequestDTOV2) ApiGetComponentDetailsRequest {
	r.apiComponentDetailsRequestDTOV2 = &apiComponentDetailsRequestDTOV2
	return r
}

func (r ApiGetComponentDetailsRequest) Execute() (*ApiComponentDetailsResultDTOV2, *http.Response, error) {
	return r.ApiService.GetComponentDetailsExecute(r)
}

/*
GetComponentDetails Method for GetComponentDetails

Use this method to retrieve data related to a component.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetComponentDetailsRequest
*/
func (a *ComponentsAPIService) GetComponentDetails(ctx context.Context) ApiGetComponentDetailsRequest {
	return ApiGetComponentDetailsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiComponentDetailsResultDTOV2
func (a *ComponentsAPIService) GetComponentDetailsExecute(r ApiGetComponentDetailsRequest) (*ApiComponentDetailsResultDTOV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiComponentDetailsResultDTOV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentsAPIService.GetComponentDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/components/details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiComponentDetailsRequestDTOV2 == nil {
		return localVarReturnValue, nil, reportError("apiComponentDetailsRequestDTOV2 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.apiComponentDetailsRequestDTOV2
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

type ApiGetComponentVersionsRequest struct {
	ctx context.Context
	ApiService *ComponentsAPIService
	apiComponentOrPurlIdentifierDTOV2 *ApiComponentOrPurlIdentifierDTOV2
}

// Possible values: Component identifier or packageURL (pURL) identifier in the correct format. Use a-name for JavaScript components.
func (r ApiGetComponentVersionsRequest) ApiComponentOrPurlIdentifierDTOV2(apiComponentOrPurlIdentifierDTOV2 ApiComponentOrPurlIdentifierDTOV2) ApiGetComponentVersionsRequest {
	r.apiComponentOrPurlIdentifierDTOV2 = &apiComponentOrPurlIdentifierDTOV2
	return r
}

func (r ApiGetComponentVersionsRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.GetComponentVersionsExecute(r)
}

/*
GetComponentVersions Method for GetComponentVersions

Use this method to retrieve all known versions of a component.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetComponentVersionsRequest
*/
func (a *ComponentsAPIService) GetComponentVersions(ctx context.Context) ApiGetComponentVersionsRequest {
	return ApiGetComponentVersionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []string
func (a *ComponentsAPIService) GetComponentVersionsExecute(r ApiGetComponentVersionsRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentsAPIService.GetComponentVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/components/versions"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiComponentOrPurlIdentifierDTOV2
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

type ApiGetSuggestedRemediationForComponentRequest struct {
	ctx context.Context
	ApiService *ComponentsAPIService
	ownerType string
	ownerId string
	stageId *string
	identificationSource *string
	scanId *string
	includeParentRemediation *bool
	apiComponentDTOV2 *ApiComponentDTOV2
}

// Enter the stageId to obtain next-non-failing and next-non-failing-with-dependencies remediation types in the response. Possible values are develop, build, stage-release, release and operate.
func (r ApiGetSuggestedRemediationForComponentRequest) StageId(stageId string) ApiGetSuggestedRemediationForComponentRequest {
	r.stageId = &stageId
	return r
}

// Enter the identification source if you want the remediation result based on third-party scan information (non-Sonatype). The identification source can be obtained from the Component Details Page in the UI.
func (r ApiGetSuggestedRemediationForComponentRequest) IdentificationSource(identificationSource string) ApiGetSuggestedRemediationForComponentRequest {
	r.identificationSource = &identificationSource
	return r
}

// Enter the scanId (reportId) if you want the remediation result based on third-party scan information (non-Sonatype).
func (r ApiGetSuggestedRemediationForComponentRequest) ScanId(scanId string) ApiGetSuggestedRemediationForComponentRequest {
	r.scanId = &scanId
	return r
}

// Enter true if you want to include parent remediation for transitive dependency in the response based on your application policy scan.
func (r ApiGetSuggestedRemediationForComponentRequest) IncludeParentRemediation(includeParentRemediation bool) ApiGetSuggestedRemediationForComponentRequest {
	r.includeParentRemediation = &includeParentRemediation
	return r
}

func (r ApiGetSuggestedRemediationForComponentRequest) ApiComponentDTOV2(apiComponentDTOV2 ApiComponentDTOV2) ApiGetSuggestedRemediationForComponentRequest {
	r.apiComponentDTOV2 = &apiComponentDTOV2
	return r
}

func (r ApiGetSuggestedRemediationForComponentRequest) Execute() (*GetSuggestedRemediationForComponent200Response, *http.Response, error) {
	return r.ApiService.GetSuggestedRemediationForComponentExecute(r)
}

/*
GetSuggestedRemediationForComponent Method for GetSuggestedRemediationForComponent

Use this method to obtain remediation suggestions for policy violations on a component basis. Remediations obtained from this method are same as those appearing on the Component Details Page in the UI.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType Possible values: application, organization, repository. 
 @param ownerId Possible values: applicationId, organizationId or repositoryId.
 @return ApiGetSuggestedRemediationForComponentRequest
*/
func (a *ComponentsAPIService) GetSuggestedRemediationForComponent(ctx context.Context, ownerType string, ownerId string) ApiGetSuggestedRemediationForComponentRequest {
	return ApiGetSuggestedRemediationForComponentRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
	}
}

// Execute executes the request
//  @return GetSuggestedRemediationForComponent200Response
func (a *ComponentsAPIService) GetSuggestedRemediationForComponentExecute(r ApiGetSuggestedRemediationForComponentRequest) (*GetSuggestedRemediationForComponent200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSuggestedRemediationForComponent200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentsAPIService.GetSuggestedRemediationForComponent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/components/remediation/{ownerType}/{ownerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.stageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stageId", r.stageId, "form", "")
	}
	if r.identificationSource != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identificationSource", r.identificationSource, "form", "")
	}
	if r.scanId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scanId", r.scanId, "form", "")
	}
	if r.includeParentRemediation != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeParentRemediation", r.includeParentRemediation, "form", "")
	} else {
		var defaultValue bool = false
		r.includeParentRemediation = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.apiComponentDTOV2
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

type ApiSetComponentLabelRequest struct {
	ctx context.Context
	ApiService *ComponentsAPIService
	ownerType string
	internalOwnerId string
	componentHash string
	labelName string
}

func (r ApiSetComponentLabelRequest) Execute() (*http.Response, error) {
	return r.ApiService.SetComponentLabelExecute(r)
}

/*
SetComponentLabel Method for SetComponentLabel

Use this method to assign an existing label to a component.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType Possible values: application or organization
 @param internalOwnerId Possible values : applicationId or organizationId
 @param componentHash Enter the SHA1 hash of the component.
 @param labelName Enter the label name to assign to this component.
 @return ApiSetComponentLabelRequest
*/
func (a *ComponentsAPIService) SetComponentLabel(ctx context.Context, ownerType string, internalOwnerId string, componentHash string, labelName string) ApiSetComponentLabelRequest {
	return ApiSetComponentLabelRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		internalOwnerId: internalOwnerId,
		componentHash: componentHash,
		labelName: labelName,
	}
}

// Execute executes the request
func (a *ComponentsAPIService) SetComponentLabelExecute(r ApiSetComponentLabelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ComponentsAPIService.SetComponentLabel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/components/{componentHash}/labels/{labelName}/{ownerType}s/{internalOwnerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"internalOwnerId"+"}", url.PathEscape(parameterValueToString(r.internalOwnerId, "internalOwnerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"componentHash"+"}", url.PathEscape(parameterValueToString(r.componentHash, "componentHash")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"labelName"+"}", url.PathEscape(parameterValueToString(r.labelName, "labelName")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
