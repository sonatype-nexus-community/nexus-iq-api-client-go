/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.182.0-01
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


// SpdxAPIService SpdxAPI service
type SpdxAPIService service

type ApiGetByScanIdRequest struct {
	ctx context.Context
	ApiService *SpdxAPIService
	applicationId string
	scanId string
	format *string
	generateCycloneDx *bool
	spdxVersion *string
}

func (r ApiGetByScanIdRequest) Format(format string) ApiGetByScanIdRequest {
	r.format = &format
	return r
}

func (r ApiGetByScanIdRequest) GenerateCycloneDx(generateCycloneDx bool) ApiGetByScanIdRequest {
	r.generateCycloneDx = &generateCycloneDx
	return r
}

func (r ApiGetByScanIdRequest) SpdxVersion(spdxVersion string) ApiGetByScanIdRequest {
	r.spdxVersion = &spdxVersion
	return r
}

func (r ApiGetByScanIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetByScanIdExecute(r)
}

/*
GetByScanId Method for GetByScanId

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId
 @param scanId
 @return ApiGetByScanIdRequest
*/
func (a *SpdxAPIService) GetByScanId(ctx context.Context, applicationId string, scanId string) ApiGetByScanIdRequest {
	return ApiGetByScanIdRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
		scanId: scanId,
	}
}

// Execute executes the request
func (a *SpdxAPIService) GetByScanIdExecute(r ApiGetByScanIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SpdxAPIService.GetByScanId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/spdx/{applicationId}/reports/{scanId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scanId"+"}", url.PathEscape(parameterValueToString(r.scanId, "scanId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
		var defaultValue string = "json"
		r.format = &defaultValue
	}
	if r.generateCycloneDx != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "generateCycloneDx", r.generateCycloneDx, "form", "")
	} else {
		var defaultValue bool = false
		r.generateCycloneDx = &defaultValue
	}
	if r.spdxVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "spdxVersion", r.spdxVersion, "form", "")
	} else {
		var defaultValue string = "2.3"
		r.spdxVersion = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/octet-stream", "application/xml"}

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

type ApiGetLatestForStageRequest struct {
	ctx context.Context
	ApiService *SpdxAPIService
	applicationId string
	stageId string
	format *string
	generateCycloneDx *bool
	spdxVersion *string
}

func (r ApiGetLatestForStageRequest) Format(format string) ApiGetLatestForStageRequest {
	r.format = &format
	return r
}

func (r ApiGetLatestForStageRequest) GenerateCycloneDx(generateCycloneDx bool) ApiGetLatestForStageRequest {
	r.generateCycloneDx = &generateCycloneDx
	return r
}

func (r ApiGetLatestForStageRequest) SpdxVersion(spdxVersion string) ApiGetLatestForStageRequest {
	r.spdxVersion = &spdxVersion
	return r
}

func (r ApiGetLatestForStageRequest) Execute() (*http.Response, error) {
	return r.ApiService.GetLatestForStageExecute(r)
}

/*
GetLatestForStage Method for GetLatestForStage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId
 @param stageId
 @return ApiGetLatestForStageRequest
*/
func (a *SpdxAPIService) GetLatestForStage(ctx context.Context, applicationId string, stageId string) ApiGetLatestForStageRequest {
	return ApiGetLatestForStageRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
		stageId: stageId,
	}
}

// Execute executes the request
func (a *SpdxAPIService) GetLatestForStageExecute(r ApiGetLatestForStageRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SpdxAPIService.GetLatestForStage")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/spdx/{applicationId}/stages/{stageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stageId"+"}", url.PathEscape(parameterValueToString(r.stageId, "stageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.format != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "format", r.format, "form", "")
	} else {
		var defaultValue string = "json"
		r.format = &defaultValue
	}
	if r.generateCycloneDx != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "generateCycloneDx", r.generateCycloneDx, "form", "")
	} else {
		var defaultValue bool = false
		r.generateCycloneDx = &defaultValue
	}
	if r.spdxVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "spdxVersion", r.spdxVersion, "form", "")
	} else {
		var defaultValue string = "2.3"
		r.spdxVersion = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/octet-stream", "application/xml"}

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
