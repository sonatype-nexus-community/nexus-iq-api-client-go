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


// VulnerabilitiesAPIService VulnerabilitiesAPI service
type VulnerabilitiesAPIService service

type ApiGetSecurityVulnerabilityDetailsRequest struct {
	ctx context.Context
	ApiService *VulnerabilitiesAPIService
	refId string
	componentIdentifier *ComponentIdentifier
	identificationSource *string
	scanId *string
	ownerType *string
	ownerId *string
}

func (r ApiGetSecurityVulnerabilityDetailsRequest) ComponentIdentifier(componentIdentifier ComponentIdentifier) ApiGetSecurityVulnerabilityDetailsRequest {
	r.componentIdentifier = &componentIdentifier
	return r
}

func (r ApiGetSecurityVulnerabilityDetailsRequest) IdentificationSource(identificationSource string) ApiGetSecurityVulnerabilityDetailsRequest {
	r.identificationSource = &identificationSource
	return r
}

func (r ApiGetSecurityVulnerabilityDetailsRequest) ScanId(scanId string) ApiGetSecurityVulnerabilityDetailsRequest {
	r.scanId = &scanId
	return r
}

func (r ApiGetSecurityVulnerabilityDetailsRequest) OwnerType(ownerType string) ApiGetSecurityVulnerabilityDetailsRequest {
	r.ownerType = &ownerType
	return r
}

func (r ApiGetSecurityVulnerabilityDetailsRequest) OwnerId(ownerId string) ApiGetSecurityVulnerabilityDetailsRequest {
	r.ownerId = &ownerId
	return r
}

func (r ApiGetSecurityVulnerabilityDetailsRequest) Execute() (*SecurityVulnerabilityData, *http.Response, error) {
	return r.ApiService.GetSecurityVulnerabilityDetailsExecute(r)
}

/*
GetSecurityVulnerabilityDetails Method for GetSecurityVulnerabilityDetails

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param refId
 @return ApiGetSecurityVulnerabilityDetailsRequest
*/
func (a *VulnerabilitiesAPIService) GetSecurityVulnerabilityDetails(ctx context.Context, refId string) ApiGetSecurityVulnerabilityDetailsRequest {
	return ApiGetSecurityVulnerabilityDetailsRequest{
		ApiService: a,
		ctx: ctx,
		refId: refId,
	}
}

// Execute executes the request
//  @return SecurityVulnerabilityData
func (a *VulnerabilitiesAPIService) GetSecurityVulnerabilityDetailsExecute(r ApiGetSecurityVulnerabilityDetailsRequest) (*SecurityVulnerabilityData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SecurityVulnerabilityData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VulnerabilitiesAPIService.GetSecurityVulnerabilityDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/vulnerabilities/{refId}"
	localVarPath = strings.Replace(localVarPath, "{"+"refId"+"}", url.PathEscape(parameterValueToString(r.refId, "refId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.componentIdentifier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "componentIdentifier", r.componentIdentifier, "form", "")
	}
	if r.identificationSource != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "identificationSource", r.identificationSource, "form", "")
	}
	if r.scanId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scanId", r.scanId, "form", "")
	}
	if r.ownerType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ownerType", r.ownerType, "form", "")
	}
	if r.ownerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ownerId", r.ownerId, "form", "")
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
			var v SecurityVulnerabilityData
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
