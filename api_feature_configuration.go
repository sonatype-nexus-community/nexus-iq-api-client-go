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


// FeatureConfigurationAPIService FeatureConfigurationAPI service
type FeatureConfigurationAPIService service

type ApiDisableFeatureRequest struct {
	ctx context.Context
	ApiService *FeatureConfigurationAPIService
	feature string
}

func (r ApiDisableFeatureRequest) Execute() (*http.Response, error) {
	return r.ApiService.DisableFeatureExecute(r)
}

/*
DisableFeature Method for DisableFeature

Use this method to disable an IQ Server feature.

Permissions required: Edit System Configuration and Users

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param feature Enter the name of the IQ Server feature to be disabled.
 @return ApiDisableFeatureRequest
*/
func (a *FeatureConfigurationAPIService) DisableFeature(ctx context.Context, feature string) ApiDisableFeatureRequest {
	return ApiDisableFeatureRequest{
		ApiService: a,
		ctx: ctx,
		feature: feature,
	}
}

// Execute executes the request
func (a *FeatureConfigurationAPIService) DisableFeatureExecute(r ApiDisableFeatureRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureConfigurationAPIService.DisableFeature")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/features/{feature}"
	localVarPath = strings.Replace(localVarPath, "{"+"feature"+"}", url.PathEscape(parameterValueToString(r.feature, "feature")), -1)

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

type ApiEnabledFeatureRequest struct {
	ctx context.Context
	ApiService *FeatureConfigurationAPIService
	feature string
}

func (r ApiEnabledFeatureRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnabledFeatureExecute(r)
}

/*
EnabledFeature Method for EnabledFeature

Use this method to enable an IQ Server feature.

Permissions required: Edit System Configuration and Users

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param feature Enter the name of the feature to be enabled.
 @return ApiEnabledFeatureRequest
*/
func (a *FeatureConfigurationAPIService) EnabledFeature(ctx context.Context, feature string) ApiEnabledFeatureRequest {
	return ApiEnabledFeatureRequest{
		ApiService: a,
		ctx: ctx,
		feature: feature,
	}
}

// Execute executes the request
func (a *FeatureConfigurationAPIService) EnabledFeatureExecute(r ApiEnabledFeatureRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FeatureConfigurationAPIService.EnabledFeature")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/config/features/{feature}"
	localVarPath = strings.Replace(localVarPath, "{"+"feature"+"}", url.PathEscape(parameterValueToString(r.feature, "feature")), -1)

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
