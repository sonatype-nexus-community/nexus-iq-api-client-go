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
)


// SearchAPIService SearchAPI service
type SearchAPIService service

type ApiSearchComponentRequest struct {
	ctx context.Context
	ApiService *SearchAPIService
	stageId *string
	hash *string
	componentIdentifier *ComponentIdentifier
	packageUrl *string
}

func (r ApiSearchComponentRequest) StageId(stageId string) ApiSearchComponentRequest {
	r.stageId = &stageId
	return r
}

func (r ApiSearchComponentRequest) Hash(hash string) ApiSearchComponentRequest {
	r.hash = &hash
	return r
}

func (r ApiSearchComponentRequest) ComponentIdentifier(componentIdentifier ComponentIdentifier) ApiSearchComponentRequest {
	r.componentIdentifier = &componentIdentifier
	return r
}

func (r ApiSearchComponentRequest) PackageUrl(packageUrl string) ApiSearchComponentRequest {
	r.packageUrl = &packageUrl
	return r
}

func (r ApiSearchComponentRequest) Execute() (*ApiSearchResultsDTOV2, *http.Response, error) {
	return r.ApiService.SearchComponentExecute(r)
}

/*
SearchComponent Method for SearchComponent

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchComponentRequest
*/
func (a *SearchAPIService) SearchComponent(ctx context.Context) ApiSearchComponentRequest {
	return ApiSearchComponentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiSearchResultsDTOV2
func (a *SearchAPIService) SearchComponentExecute(r ApiSearchComponentRequest) (*ApiSearchResultsDTOV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiSearchResultsDTOV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchAPIService.SearchComponent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/search/component"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.stageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "stageId", r.stageId, "form", "")
	}
	if r.hash != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "hash", r.hash, "form", "")
	}
	if r.componentIdentifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "componentIdentifier", r.componentIdentifier, "form", "")
	}
	if r.packageUrl != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "packageUrl", r.packageUrl, "form", "")
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
			var v ApiSearchResultsDTOV2
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
