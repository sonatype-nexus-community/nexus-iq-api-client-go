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


// ScanAPIService ScanAPI service
type ScanAPIService service

type ApiGetIdeUsersOverviewRequest struct {
	ctx context.Context
	ApiService *ScanAPIService
	sinceUtcTimestamp *int64
}

func (r ApiGetIdeUsersOverviewRequest) SinceUtcTimestamp(sinceUtcTimestamp int64) ApiGetIdeUsersOverviewRequest {
	r.sinceUtcTimestamp = &sinceUtcTimestamp
	return r
}

func (r ApiGetIdeUsersOverviewRequest) Execute() (*IdeUsersOverviewDTO, *http.Response, error) {
	return r.ApiService.GetIdeUsersOverviewExecute(r)
}

/*
GetIdeUsersOverview Method for GetIdeUsersOverview

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetIdeUsersOverviewRequest
*/
func (a *ScanAPIService) GetIdeUsersOverview(ctx context.Context) ApiGetIdeUsersOverviewRequest {
	return ApiGetIdeUsersOverviewRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IdeUsersOverviewDTO
func (a *ScanAPIService) GetIdeUsersOverviewExecute(r ApiGetIdeUsersOverviewRequest) (*IdeUsersOverviewDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IdeUsersOverviewDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScanAPIService.GetIdeUsersOverview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scan/applications/ideUser/overview"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sinceUtcTimestamp != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sinceUtcTimestamp", r.sinceUtcTimestamp, "")
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
			var v IdeUsersOverviewDTO
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

type ApiGetScanStatusRequest struct {
	ctx context.Context
	ApiService *ScanAPIService
	applicationId string
	scanRequestId string
}

func (r ApiGetScanStatusRequest) Execute() (*ApiThirdPartyScanResultDTO, *http.Response, error) {
	return r.ApiService.GetScanStatusExecute(r)
}

/*
GetScanStatus Method for GetScanStatus

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId
 @param scanRequestId
 @return ApiGetScanStatusRequest
*/
func (a *ScanAPIService) GetScanStatus(ctx context.Context, applicationId string, scanRequestId string) ApiGetScanStatusRequest {
	return ApiGetScanStatusRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
		scanRequestId: scanRequestId,
	}
}

// Execute executes the request
//  @return ApiThirdPartyScanResultDTO
func (a *ScanAPIService) GetScanStatusExecute(r ApiGetScanStatusRequest) (*ApiThirdPartyScanResultDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiThirdPartyScanResultDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScanAPIService.GetScanStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scan/applications/{applicationId}/status/{scanRequestId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"scanRequestId"+"}", url.PathEscape(parameterValueToString(r.scanRequestId, "scanRequestId")), -1)

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
			var v ApiThirdPartyScanResultDTO
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

type ApiScanComponentsRequest struct {
	ctx context.Context
	ApiService *ScanAPIService
	applicationId string
	source string
	stageId *string
	body *string
}

func (r ApiScanComponentsRequest) StageId(stageId string) ApiScanComponentsRequest {
	r.stageId = &stageId
	return r
}

func (r ApiScanComponentsRequest) Body(body string) ApiScanComponentsRequest {
	r.body = &body
	return r
}

func (r ApiScanComponentsRequest) Execute() (*http.Response, error) {
	return r.ApiService.ScanComponentsExecute(r)
}

/*
ScanComponents Method for ScanComponents

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId
 @param source
 @return ApiScanComponentsRequest
*/
func (a *ScanAPIService) ScanComponents(ctx context.Context, applicationId string, source string) ApiScanComponentsRequest {
	return ApiScanComponentsRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
		source: source,
	}
}

// Execute executes the request
func (a *ScanAPIService) ScanComponentsExecute(r ApiScanComponentsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ScanAPIService.ScanComponents")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/scan/applications/{applicationId}/sources/{source}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"source"+"}", url.PathEscape(parameterValueToString(r.source, "source")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.stageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stageId", r.stageId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

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