/*
Sonatype IQ Server

This documents the available APIs into [Sonatype IQ Server](https://www.sonatype.com/products/open-source-security-dependency-management).

API version: 164
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// LabelsAPIService LabelsAPI service
type LabelsAPIService service

type ApiAddLabelRequest struct {
	ctx context.Context
	ApiService *LabelsAPIService
	ownerType string
	ownerId string
	apiLabelDTO *ApiLabelDTO
}

func (r ApiAddLabelRequest) ApiLabelDTO(apiLabelDTO ApiLabelDTO) ApiAddLabelRequest {
	r.apiLabelDTO = &apiLabelDTO
	return r
}

func (r ApiAddLabelRequest) Execute() (*ApiLabelDTO, *http.Response, error) {
	return r.ApiService.AddLabelExecute(r)
}

/*
AddLabel Method for AddLabel

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @return ApiAddLabelRequest
*/
func (a *LabelsAPIService) AddLabel(ctx context.Context, ownerType string, ownerId string) ApiAddLabelRequest {
	return ApiAddLabelRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
	}
}

// Execute executes the request
//  @return ApiLabelDTO
func (a *LabelsAPIService) AddLabelExecute(r ApiAddLabelRequest) (*ApiLabelDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiLabelDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LabelsAPIService.AddLabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/labels/{ownerType}/{ownerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)

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
	localVarPostBody = r.apiLabelDTO
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
			var v ApiLabelDTO
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

type ApiDeleteLabelRequest struct {
	ctx context.Context
	ApiService *LabelsAPIService
	ownerType string
	ownerId string
	labelId string
}

func (r ApiDeleteLabelRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteLabelExecute(r)
}

/*
DeleteLabel Method for DeleteLabel

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @param labelId
 @return ApiDeleteLabelRequest
*/
func (a *LabelsAPIService) DeleteLabel(ctx context.Context, ownerType string, ownerId string, labelId string) ApiDeleteLabelRequest {
	return ApiDeleteLabelRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
		labelId: labelId,
	}
}

// Execute executes the request
func (a *LabelsAPIService) DeleteLabelExecute(r ApiDeleteLabelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LabelsAPIService.DeleteLabel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/labels/{ownerType}/{ownerId}/{labelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"labelId"+"}", url.PathEscape(parameterValueToString(r.labelId, "labelId")), -1)

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

type ApiGetApplicableContextsRequest struct {
	ctx context.Context
	ApiService *LabelsAPIService
	ownerType string
	ownerId string
	labelId string
}

func (r ApiGetApplicableContextsRequest) Execute() (*ApplicableContext, *http.Response, error) {
	return r.ApiService.GetApplicableContextsExecute(r)
}

/*
GetApplicableContexts Method for GetApplicableContexts

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @param labelId
 @return ApiGetApplicableContextsRequest
*/
func (a *LabelsAPIService) GetApplicableContexts(ctx context.Context, ownerType string, ownerId string, labelId string) ApiGetApplicableContextsRequest {
	return ApiGetApplicableContextsRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
		labelId: labelId,
	}
}

// Execute executes the request
//  @return ApplicableContext
func (a *LabelsAPIService) GetApplicableContextsExecute(r ApiGetApplicableContextsRequest) (*ApplicableContext, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicableContext
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LabelsAPIService.GetApplicableContexts")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/labels/{ownerType}/{ownerId}/applicable/context/{labelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"labelId"+"}", url.PathEscape(parameterValueToString(r.labelId, "labelId")), -1)

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
			var v ApplicableContext
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

type ApiGetApplicableLabelsRequest struct {
	ctx context.Context
	ApiService *LabelsAPIService
	ownerType string
	ownerId string
}

func (r ApiGetApplicableLabelsRequest) Execute() (*ApplicableLabels, *http.Response, error) {
	return r.ApiService.GetApplicableLabelsExecute(r)
}

/*
GetApplicableLabels Method for GetApplicableLabels

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @return ApiGetApplicableLabelsRequest
*/
func (a *LabelsAPIService) GetApplicableLabels(ctx context.Context, ownerType string, ownerId string) ApiGetApplicableLabelsRequest {
	return ApiGetApplicableLabelsRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
	}
}

// Execute executes the request
//  @return ApplicableLabels
func (a *LabelsAPIService) GetApplicableLabelsExecute(r ApiGetApplicableLabelsRequest) (*ApplicableLabels, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApplicableLabels
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LabelsAPIService.GetApplicableLabels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/labels/{ownerType}/{ownerId}/applicable"
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
			var v ApplicableLabels
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

type ApiGetLabelsRequest struct {
	ctx context.Context
	ApiService *LabelsAPIService
	ownerType string
	ownerId string
	inherit *bool
}

func (r ApiGetLabelsRequest) Inherit(inherit bool) ApiGetLabelsRequest {
	r.inherit = &inherit
	return r
}

func (r ApiGetLabelsRequest) Execute() ([]ApiLabelDTO, *http.Response, error) {
	return r.ApiService.GetLabelsExecute(r)
}

/*
GetLabels Method for GetLabels

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @return ApiGetLabelsRequest
*/
func (a *LabelsAPIService) GetLabels(ctx context.Context, ownerType string, ownerId string) ApiGetLabelsRequest {
	return ApiGetLabelsRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
	}
}

// Execute executes the request
//  @return []ApiLabelDTO
func (a *LabelsAPIService) GetLabelsExecute(r ApiGetLabelsRequest) ([]ApiLabelDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiLabelDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LabelsAPIService.GetLabels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/labels/{ownerType}/{ownerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.inherit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "inherit", r.inherit, "")
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
			var v []ApiLabelDTO
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

type ApiUpdateLabelRequest struct {
	ctx context.Context
	ApiService *LabelsAPIService
	ownerType string
	ownerId string
	apiLabelDTO *ApiLabelDTO
}

func (r ApiUpdateLabelRequest) ApiLabelDTO(apiLabelDTO ApiLabelDTO) ApiUpdateLabelRequest {
	r.apiLabelDTO = &apiLabelDTO
	return r
}

func (r ApiUpdateLabelRequest) Execute() (*ApiLabelDTO, *http.Response, error) {
	return r.ApiService.UpdateLabelExecute(r)
}

/*
UpdateLabel Method for UpdateLabel

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ownerType
 @param ownerId
 @return ApiUpdateLabelRequest
*/
func (a *LabelsAPIService) UpdateLabel(ctx context.Context, ownerType string, ownerId string) ApiUpdateLabelRequest {
	return ApiUpdateLabelRequest{
		ApiService: a,
		ctx: ctx,
		ownerType: ownerType,
		ownerId: ownerId,
	}
}

// Execute executes the request
//  @return ApiLabelDTO
func (a *LabelsAPIService) UpdateLabelExecute(r ApiUpdateLabelRequest) (*ApiLabelDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiLabelDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LabelsAPIService.UpdateLabel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/labels/{ownerType}/{ownerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ownerType"+"}", url.PathEscape(parameterValueToString(r.ownerType, "ownerType")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"ownerId"+"}", url.PathEscape(parameterValueToString(r.ownerId, "ownerId")), -1)

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
	localVarPostBody = r.apiLabelDTO
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
			var v ApiLabelDTO
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
