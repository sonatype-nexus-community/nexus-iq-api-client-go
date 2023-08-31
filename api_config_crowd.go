/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.166.0-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sonatypeiq

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ConfigCrowdAPIService ConfigCrowdAPI service
type ConfigCrowdAPIService service

type ApiDeleteCrowdConfigurationRequest struct {
	ctx context.Context
	ApiService *ConfigCrowdAPIService
}

func (r ApiDeleteCrowdConfigurationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCrowdConfigurationExecute(r)
}

/*
DeleteCrowdConfiguration Method for DeleteCrowdConfiguration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteCrowdConfigurationRequest
*/
func (a *ConfigCrowdAPIService) DeleteCrowdConfiguration(ctx context.Context) ApiDeleteCrowdConfigurationRequest {
	return ApiDeleteCrowdConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigCrowdAPIService) DeleteCrowdConfigurationExecute(r ApiDeleteCrowdConfigurationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigCrowdAPIService.DeleteCrowdConfiguration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/crowd"

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

type ApiGetCrowdConfigurationRequest struct {
	ctx context.Context
	ApiService *ConfigCrowdAPIService
}

func (r ApiGetCrowdConfigurationRequest) Execute() (*ApiCrowdConfigurationDTO, *http.Response, error) {
	return r.ApiService.GetCrowdConfigurationExecute(r)
}

/*
GetCrowdConfiguration Method for GetCrowdConfiguration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCrowdConfigurationRequest
*/
func (a *ConfigCrowdAPIService) GetCrowdConfiguration(ctx context.Context) ApiGetCrowdConfigurationRequest {
	return ApiGetCrowdConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiCrowdConfigurationDTO
func (a *ConfigCrowdAPIService) GetCrowdConfigurationExecute(r ApiGetCrowdConfigurationRequest) (*ApiCrowdConfigurationDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiCrowdConfigurationDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigCrowdAPIService.GetCrowdConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/crowd"

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
			var v ApiCrowdConfigurationDTO
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

type ApiInsertOrUpdateCrowdConfigurationRequest struct {
	ctx context.Context
	ApiService *ConfigCrowdAPIService
	apiCrowdConfigurationDTO *ApiCrowdConfigurationDTO
}

func (r ApiInsertOrUpdateCrowdConfigurationRequest) ApiCrowdConfigurationDTO(apiCrowdConfigurationDTO ApiCrowdConfigurationDTO) ApiInsertOrUpdateCrowdConfigurationRequest {
	r.apiCrowdConfigurationDTO = &apiCrowdConfigurationDTO
	return r
}

func (r ApiInsertOrUpdateCrowdConfigurationRequest) Execute() (*http.Response, error) {
	return r.ApiService.InsertOrUpdateCrowdConfigurationExecute(r)
}

/*
InsertOrUpdateCrowdConfiguration Method for InsertOrUpdateCrowdConfiguration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiInsertOrUpdateCrowdConfigurationRequest
*/
func (a *ConfigCrowdAPIService) InsertOrUpdateCrowdConfiguration(ctx context.Context) ApiInsertOrUpdateCrowdConfigurationRequest {
	return ApiInsertOrUpdateCrowdConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ConfigCrowdAPIService) InsertOrUpdateCrowdConfigurationExecute(r ApiInsertOrUpdateCrowdConfigurationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigCrowdAPIService.InsertOrUpdateCrowdConfiguration")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/crowd"

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
	localVarPostBody = r.apiCrowdConfigurationDTO
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

type ApiTestCrowdConfigurationRequest struct {
	ctx context.Context
	ApiService *ConfigCrowdAPIService
	apiCrowdConfigurationDTO *ApiCrowdConfigurationDTO
}

func (r ApiTestCrowdConfigurationRequest) ApiCrowdConfigurationDTO(apiCrowdConfigurationDTO ApiCrowdConfigurationDTO) ApiTestCrowdConfigurationRequest {
	r.apiCrowdConfigurationDTO = &apiCrowdConfigurationDTO
	return r
}

func (r ApiTestCrowdConfigurationRequest) Execute() (*ApiStatusDTO, *http.Response, error) {
	return r.ApiService.TestCrowdConfigurationExecute(r)
}

/*
TestCrowdConfiguration Method for TestCrowdConfiguration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTestCrowdConfigurationRequest
*/
func (a *ConfigCrowdAPIService) TestCrowdConfiguration(ctx context.Context) ApiTestCrowdConfigurationRequest {
	return ApiTestCrowdConfigurationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiStatusDTO
func (a *ConfigCrowdAPIService) TestCrowdConfigurationExecute(r ApiTestCrowdConfigurationRequest) (*ApiStatusDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiStatusDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ConfigCrowdAPIService.TestCrowdConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/crowd/test"

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
	localVarPostBody = r.apiCrowdConfigurationDTO
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
			var v ApiStatusDTO
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
