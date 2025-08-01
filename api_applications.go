/*
Sonatype Lifecycle Public REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.193.0-01
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
	"reflect"
)


// ApplicationsAPIService ApplicationsAPI service
type ApplicationsAPIService service

type ApiAddApplicationRequest struct {
	ctx context.Context
	ApiService *ApplicationsAPIService
	apiApplicationDTO *ApiApplicationDTO
}

// Specify the applicationId, application name and the organizationId under which the application should be created. &#x60;contactUserName&#x60; corresponds to the &#39;contact&#39; field in the UI and represents the user name. If LDAP is used for authentication, you can use LDAP usernames.&#x60;tagId&#x60; is the internal identifier for the Application Category that you want to apply to the application. Use the Application Categories REST API for the available categories and the corresponding tagIds.
func (r ApiAddApplicationRequest) ApiApplicationDTO(apiApplicationDTO ApiApplicationDTO) ApiAddApplicationRequest {
	r.apiApplicationDTO = &apiApplicationDTO
	return r
}

func (r ApiAddApplicationRequest) Execute() (*ApiApplicationDTO, *http.Response, error) {
	return r.ApiService.AddApplicationExecute(r)
}

/*
AddApplication Method for AddApplication

Use this method to create an application under an organization. Use the Organization REST API to obtain organizationId.

Permissions required: Add Application (on parent organization)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddApplicationRequest
*/
func (a *ApplicationsAPIService) AddApplication(ctx context.Context) ApiAddApplicationRequest {
	return ApiAddApplicationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiApplicationDTO
func (a *ApplicationsAPIService) AddApplicationExecute(r ApiAddApplicationRequest) (*ApiApplicationDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiApplicationDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsAPIService.AddApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/applications"

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
	localVarPostBody = r.apiApplicationDTO
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

type ApiCloneApplicationRequest struct {
	ctx context.Context
	ApiService *ApplicationsAPIService
	sourceApplicationId string
	clonedApplicationName *string
	clonedApplicationPublicId *string
}

// Enter the application name for the new cloned application.
func (r ApiCloneApplicationRequest) ClonedApplicationName(clonedApplicationName string) ApiCloneApplicationRequest {
	r.clonedApplicationName = &clonedApplicationName
	return r
}

// Enter the applicationPublicId for the cloned application.
func (r ApiCloneApplicationRequest) ClonedApplicationPublicId(clonedApplicationPublicId string) ApiCloneApplicationRequest {
	r.clonedApplicationPublicId = &clonedApplicationPublicId
	return r
}

func (r ApiCloneApplicationRequest) Execute() (*ApiApplicationDTO, *http.Response, error) {
	return r.ApiService.CloneApplicationExecute(r)
}

/*
CloneApplication Method for CloneApplication

Use this method to clone an existing application.

Permissions required: Add Application (on the parent organization)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param sourceApplicationId Enter the applicationId for the application to be cloned.
 @return ApiCloneApplicationRequest
*/
func (a *ApplicationsAPIService) CloneApplication(ctx context.Context, sourceApplicationId string) ApiCloneApplicationRequest {
	return ApiCloneApplicationRequest{
		ApiService: a,
		ctx: ctx,
		sourceApplicationId: sourceApplicationId,
	}
}

// Execute executes the request
//  @return ApiApplicationDTO
func (a *ApplicationsAPIService) CloneApplicationExecute(r ApiCloneApplicationRequest) (*ApiApplicationDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiApplicationDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsAPIService.CloneApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/applications/{sourceApplicationId}/clone"
	localVarPath = strings.Replace(localVarPath, "{"+"sourceApplicationId"+"}", url.PathEscape(parameterValueToString(r.sourceApplicationId, "sourceApplicationId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.clonedApplicationName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clonedApplicationName", r.clonedApplicationName, "form", "")
	}
	if r.clonedApplicationPublicId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clonedApplicationPublicId", r.clonedApplicationPublicId, "form", "")
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

type ApiDeleteApplicationRequest struct {
	ctx context.Context
	ApiService *ApplicationsAPIService
	applicationId string
}

func (r ApiDeleteApplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteApplicationExecute(r)
}

/*
DeleteApplication Method for DeleteApplication

Use this method to permanently delete an existing application and all data associated with it. This action cannot be un-done. Before deleting, confirm that the application being deleted does not impact any integrations that could depend on it.

Permissions required: Edit IQ Elements

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Enter the applicationId to be deleted.
 @return ApiDeleteApplicationRequest
*/
func (a *ApplicationsAPIService) DeleteApplication(ctx context.Context, applicationId string) ApiDeleteApplicationRequest {
	return ApiDeleteApplicationRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
func (a *ApplicationsAPIService) DeleteApplicationExecute(r ApiDeleteApplicationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsAPIService.DeleteApplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/applications/{applicationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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

type ApiGetApplicationRequest struct {
	ctx context.Context
	ApiService *ApplicationsAPIService
	applicationId string
}

func (r ApiGetApplicationRequest) Execute() (*ApiApplicationDTO, *http.Response, error) {
	return r.ApiService.GetApplicationExecute(r)
}

/*
GetApplication Method for GetApplication

Use this method to retrieve the application details, by providing the applicationId.

Permissions required: View IQ Elements

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Enter the applicationId.
 @return ApiGetApplicationRequest
*/
func (a *ApplicationsAPIService) GetApplication(ctx context.Context, applicationId string) ApiGetApplicationRequest {
	return ApiGetApplicationRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return ApiApplicationDTO
func (a *ApplicationsAPIService) GetApplicationExecute(r ApiGetApplicationRequest) (*ApiApplicationDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiApplicationDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsAPIService.GetApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/applications/{applicationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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

type ApiGetApplicationsRequest struct {
	ctx context.Context
	ApiService *ApplicationsAPIService
	publicId *[]string
	includeCategories *bool
}

// Enter the applicationId.
func (r ApiGetApplicationsRequest) PublicId(publicId []string) ApiGetApplicationsRequest {
	r.publicId = &publicId
	return r
}

// Set this parameter to &#x60;true&#x60; to obtain the application tags (application categories) in the response.
func (r ApiGetApplicationsRequest) IncludeCategories(includeCategories bool) ApiGetApplicationsRequest {
	r.includeCategories = &includeCategories
	return r
}

func (r ApiGetApplicationsRequest) Execute() (*ApiApplicationListDTO, *http.Response, error) {
	return r.ApiService.GetApplicationsExecute(r)
}

/*
GetApplications Method for GetApplications

Use this method to retrieve the application details for the applicationId(s) provided.

Permissions required: View IQ Elements

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetApplicationsRequest
*/
func (a *ApplicationsAPIService) GetApplications(ctx context.Context) ApiGetApplicationsRequest {
	return ApiGetApplicationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiApplicationListDTO
func (a *ApplicationsAPIService) GetApplicationsExecute(r ApiGetApplicationsRequest) (*ApiApplicationListDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiApplicationListDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsAPIService.GetApplications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/applications"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.publicId != nil {
		t := *r.publicId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "publicId", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "publicId", t, "form", "multi")
		}
	}
	if r.includeCategories != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCategories", r.includeCategories, "form", "")
	} else {
		var defaultValue bool = false
		r.includeCategories = &defaultValue
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
			var v ApiApplicationListDTO
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

type ApiGetApplicationsByOrganizationIdRequest struct {
	ctx context.Context
	ApiService *ApplicationsAPIService
	organizationId string
}

func (r ApiGetApplicationsByOrganizationIdRequest) Execute() (*ApiApplicationListDTO, *http.Response, error) {
	return r.ApiService.GetApplicationsByOrganizationIdExecute(r)
}

/*
GetApplicationsByOrganizationId Method for GetApplicationsByOrganizationId

Use this method to retrieve application details for all applications under the organizationId provided.

Permissions required: View IQ Elements

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId Enter the organizationId.
 @return ApiGetApplicationsByOrganizationIdRequest
*/
func (a *ApplicationsAPIService) GetApplicationsByOrganizationId(ctx context.Context, organizationId string) ApiGetApplicationsByOrganizationIdRequest {
	return ApiGetApplicationsByOrganizationIdRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return ApiApplicationListDTO
func (a *ApplicationsAPIService) GetApplicationsByOrganizationIdExecute(r ApiGetApplicationsByOrganizationIdRequest) (*ApiApplicationListDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiApplicationListDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsAPIService.GetApplicationsByOrganizationId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/applications/organization/{organizationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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

type ApiMoveApplicationRequest struct {
	ctx context.Context
	ApiService *ApplicationsAPIService
	applicationId string
	organizationId string
}

func (r ApiMoveApplicationRequest) Execute() (*ApiMoveApplicationResponseDTOV2, *http.Response, error) {
	return r.ApiService.MoveApplicationExecute(r)
}

/*
MoveApplication Method for MoveApplication

Use this method to move an application from one organization to another.

Permissions required: Edit IQ Elements

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId Enter the applicationId of the application to be moved.
 @param organizationId Enter the organizationId of the destination organization.
 @return ApiMoveApplicationRequest
*/
func (a *ApplicationsAPIService) MoveApplication(ctx context.Context, applicationId string, organizationId string) ApiMoveApplicationRequest {
	return ApiMoveApplicationRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return ApiMoveApplicationResponseDTOV2
func (a *ApplicationsAPIService) MoveApplicationExecute(r ApiMoveApplicationRequest) (*ApiMoveApplicationResponseDTOV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiMoveApplicationResponseDTOV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsAPIService.MoveApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/applications/{applicationId}/move/organization/{organizationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterValueToString(r.organizationId, "organizationId")), -1)

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
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiMoveApplicationResponseDTOV2
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiUpdateApplicationRequest struct {
	ctx context.Context
	ApiService *ApplicationsAPIService
	applicationId string
	apiApplicationDTO *ApiApplicationDTO
}

// Specify the applicationId, application name and the organizationId under which  the application exists. &#x60;contactUserName&#x60; corresponds to the &#39;contact&#39; field in the UI and represents the user name. If LDAP is used for authentication, you can use LDAP usernames.&#x60;tagId&#x60; is the internal identifier for the Application Category that you want to apply to the application. . Use the Application Categories REST API for the available categories and the corresponding tagIds.
func (r ApiUpdateApplicationRequest) ApiApplicationDTO(apiApplicationDTO ApiApplicationDTO) ApiUpdateApplicationRequest {
	r.apiApplicationDTO = &apiApplicationDTO
	return r
}

func (r ApiUpdateApplicationRequest) Execute() (*ApiApplicationDTO, *http.Response, error) {
	return r.ApiService.UpdateApplicationExecute(r)
}

/*
UpdateApplication Method for UpdateApplication

Use this method to update the application name, application tags or the contact user name for an existing application by providing the applicationId. 

NOTE: This method cannot be used to change the organizationId of an application.

Permissions required: Edit IQ Elements

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param applicationId
 @return ApiUpdateApplicationRequest
*/
func (a *ApplicationsAPIService) UpdateApplication(ctx context.Context, applicationId string) ApiUpdateApplicationRequest {
	return ApiUpdateApplicationRequest{
		ApiService: a,
		ctx: ctx,
		applicationId: applicationId,
	}
}

// Execute executes the request
//  @return ApiApplicationDTO
func (a *ApplicationsAPIService) UpdateApplicationExecute(r ApiUpdateApplicationRequest) (*ApiApplicationDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiApplicationDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApplicationsAPIService.UpdateApplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/applications/{applicationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"applicationId"+"}", url.PathEscape(parameterValueToString(r.applicationId, "applicationId")), -1)

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
	localVarPostBody = r.apiApplicationDTO
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
