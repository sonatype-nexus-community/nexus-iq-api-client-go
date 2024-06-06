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
)


// TelemetryAPIService TelemetryAPI service
type TelemetryAPIService service

type ApiPostExternalTelemetryRequest struct {
	ctx context.Context
	ApiService *TelemetryAPIService
	userAgent *string
	requestBody *map[string]string
}

func (r ApiPostExternalTelemetryRequest) UserAgent(userAgent string) ApiPostExternalTelemetryRequest {
	r.userAgent = &userAgent
	return r
}

func (r ApiPostExternalTelemetryRequest) RequestBody(requestBody map[string]string) ApiPostExternalTelemetryRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiPostExternalTelemetryRequest) Execute() (*http.Response, error) {
	return r.ApiService.PostExternalTelemetryExecute(r)
}

/*
PostExternalTelemetry Method for PostExternalTelemetry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPostExternalTelemetryRequest
*/
func (a *TelemetryAPIService) PostExternalTelemetry(ctx context.Context) ApiPostExternalTelemetryRequest {
	return ApiPostExternalTelemetryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TelemetryAPIService) PostExternalTelemetryExecute(r ApiPostExternalTelemetryRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelemetryAPIService.PostExternalTelemetry")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/telemetry"

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
	if r.userAgent != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "user-agent", r.userAgent, "")
	}
	// body params
	localVarPostBody = r.requestBody
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
