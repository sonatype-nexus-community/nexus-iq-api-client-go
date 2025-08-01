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
)


// UserTokensAPIService UserTokensAPI service
type UserTokensAPIService service

type ApiCreateUserTokenRequest struct {
	ctx context.Context
	ApiService *UserTokensAPIService
}

func (r ApiCreateUserTokenRequest) Execute() (*ApiUserTokenDTO, *http.Response, error) {
	return r.ApiService.CreateUserTokenExecute(r)
}

/*
CreateUserToken Method for CreateUserToken

Use this method to generate a user token for the currently logged in user.

Permissions required: None

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateUserTokenRequest
*/
func (a *UserTokensAPIService) CreateUserToken(ctx context.Context) ApiCreateUserTokenRequest {
	return ApiCreateUserTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiUserTokenDTO
func (a *UserTokensAPIService) CreateUserTokenExecute(r ApiCreateUserTokenRequest) (*ApiUserTokenDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiUserTokenDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserTokensAPIService.CreateUserToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/userTokens/currentUser"

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

type ApiDeleteCurrentUserTokenRequest struct {
	ctx context.Context
	ApiService *UserTokensAPIService
}

func (r ApiDeleteCurrentUserTokenRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteCurrentUserTokenExecute(r)
}

/*
DeleteCurrentUserToken Method for DeleteCurrentUserToken

Use this method to delete an existing user token for the currently logged in user.

Permissions required: None

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiDeleteCurrentUserTokenRequest
*/
func (a *UserTokensAPIService) DeleteCurrentUserToken(ctx context.Context) ApiDeleteCurrentUserTokenRequest {
	return ApiDeleteCurrentUserTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserTokensAPIService) DeleteCurrentUserTokenExecute(r ApiDeleteCurrentUserTokenRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserTokensAPIService.DeleteCurrentUserToken")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/userTokens/currentUser"

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

type ApiDeleteUserTokenByUserCodeRequest struct {
	ctx context.Context
	ApiService *UserTokensAPIService
	userCode string
}

func (r ApiDeleteUserTokenByUserCodeRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteUserTokenByUserCodeExecute(r)
}

/*
DeleteUserTokenByUserCode Method for DeleteUserTokenByUserCode

Use this method to delete an existing user token by specifying the userCode.

Permissions required: Edit System Configuration and Users

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userCode Enter the `userCode` to be deleted.
 @return ApiDeleteUserTokenByUserCodeRequest
*/
func (a *UserTokensAPIService) DeleteUserTokenByUserCode(ctx context.Context, userCode string) ApiDeleteUserTokenByUserCodeRequest {
	return ApiDeleteUserTokenByUserCodeRequest{
		ApiService: a,
		ctx: ctx,
		userCode: userCode,
	}
}

// Execute executes the request
func (a *UserTokensAPIService) DeleteUserTokenByUserCodeExecute(r ApiDeleteUserTokenByUserCodeRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserTokensAPIService.DeleteUserTokenByUserCode")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/userTokens/userCode/{userCode}"
	localVarPath = strings.Replace(localVarPath, "{"+"userCode"+"}", url.PathEscape(parameterValueToString(r.userCode, "userCode")), -1)

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

type ApiGetUserTokenByUsernameAndRealmIdRequest struct {
	ctx context.Context
	ApiService *UserTokensAPIService
	username string
	realm *string
}

// Enter the realmId. Possible values are &#x60;Internal&#x60;, &#x60;SAML&#x60; , &#x60;OAUTH2&#x60; , and &#x60;Crowd&#x60;.
func (r ApiGetUserTokenByUsernameAndRealmIdRequest) Realm(realm string) ApiGetUserTokenByUsernameAndRealmIdRequest {
	r.realm = &realm
	return r
}

func (r ApiGetUserTokenByUsernameAndRealmIdRequest) Execute() (*ApiUserTokenDTO, *http.Response, error) {
	return r.ApiService.GetUserTokenByUsernameAndRealmIdExecute(r)
}

/*
GetUserTokenByUsernameAndRealmId Method for GetUserTokenByUsernameAndRealmId

Use this method to retrieve a user token by specifying a username and realmId.

Permissions required: Edit System Configuration and Users

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username Enter the username.
 @return ApiGetUserTokenByUsernameAndRealmIdRequest
*/
func (a *UserTokensAPIService) GetUserTokenByUsernameAndRealmId(ctx context.Context, username string) ApiGetUserTokenByUsernameAndRealmIdRequest {
	return ApiGetUserTokenByUsernameAndRealmIdRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return ApiUserTokenDTO
func (a *UserTokensAPIService) GetUserTokenByUsernameAndRealmIdExecute(r ApiGetUserTokenByUsernameAndRealmIdRequest) (*ApiUserTokenDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiUserTokenDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserTokensAPIService.GetUserTokenByUsernameAndRealmId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/userTokens/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.realm != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "realm", r.realm, "form", "")
	} else {
		var defaultValue string = "Internal"
		r.realm = &defaultValue
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

type ApiGetUserTokenExistsForCurrentUserRequest struct {
	ctx context.Context
	ApiService *UserTokensAPIService
}

func (r ApiGetUserTokenExistsForCurrentUserRequest) Execute() (*ApiUserTokenExistsDTO, *http.Response, error) {
	return r.ApiService.GetUserTokenExistsForCurrentUserExecute(r)
}

/*
GetUserTokenExistsForCurrentUser Method for GetUserTokenExistsForCurrentUser

Use this method to check if a user token has been issued to the logged in user.

Permissions required: None 

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserTokenExistsForCurrentUserRequest
*/
func (a *UserTokensAPIService) GetUserTokenExistsForCurrentUser(ctx context.Context) ApiGetUserTokenExistsForCurrentUserRequest {
	return ApiGetUserTokenExistsForCurrentUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiUserTokenExistsDTO
func (a *UserTokensAPIService) GetUserTokenExistsForCurrentUserExecute(r ApiGetUserTokenExistsForCurrentUserRequest) (*ApiUserTokenExistsDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiUserTokenExistsDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserTokensAPIService.GetUserTokenExistsForCurrentUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/userTokens/currentUser/hasToken"

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

type ApiGetUserTokensByCreatedBetweenAndRealmIdRequest struct {
	ctx context.Context
	ApiService *UserTokensAPIService
	createdAfter *string
	createdBefore *string
	realm *string
}

// Enter the start date for the date range in &#x60;yyyy-mm-dd&#x60; format.
func (r ApiGetUserTokensByCreatedBetweenAndRealmIdRequest) CreatedAfter(createdAfter string) ApiGetUserTokensByCreatedBetweenAndRealmIdRequest {
	r.createdAfter = &createdAfter
	return r
}

// Enter the end date for the date range in &#x60;yyyy-mm-dd&#x60; format.
func (r ApiGetUserTokensByCreatedBetweenAndRealmIdRequest) CreatedBefore(createdBefore string) ApiGetUserTokensByCreatedBetweenAndRealmIdRequest {
	r.createdBefore = &createdBefore
	return r
}

// Enter the &#x60;realmId&#x60;. Possible values are &#x60;Internal&#x60;, &#x60;SAML&#x60; , &#x60;OAUTH2&#x60;, and &#x60;Crowd&#x60;.
func (r ApiGetUserTokensByCreatedBetweenAndRealmIdRequest) Realm(realm string) ApiGetUserTokensByCreatedBetweenAndRealmIdRequest {
	r.realm = &realm
	return r
}

func (r ApiGetUserTokensByCreatedBetweenAndRealmIdRequest) Execute() ([]ApiUserTokenDTO, *http.Response, error) {
	return r.ApiService.GetUserTokensByCreatedBetweenAndRealmIdExecute(r)
}

/*
GetUserTokensByCreatedBetweenAndRealmId Method for GetUserTokensByCreatedBetweenAndRealmId

Use this method to retrieve user tokens created within a date range, in the supported IQ Server realms.

Permissions required: Edit System Configuration and Users.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserTokensByCreatedBetweenAndRealmIdRequest
*/
func (a *UserTokensAPIService) GetUserTokensByCreatedBetweenAndRealmId(ctx context.Context) ApiGetUserTokensByCreatedBetweenAndRealmIdRequest {
	return ApiGetUserTokensByCreatedBetweenAndRealmIdRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ApiUserTokenDTO
func (a *UserTokensAPIService) GetUserTokensByCreatedBetweenAndRealmIdExecute(r ApiGetUserTokensByCreatedBetweenAndRealmIdRequest) ([]ApiUserTokenDTO, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiUserTokenDTO
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserTokensAPIService.GetUserTokensByCreatedBetweenAndRealmId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/userTokens"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.createdAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdAfter", r.createdAfter, "form", "")
	}
	if r.createdBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdBefore", r.createdBefore, "form", "")
	}
	if r.realm != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "realm", r.realm, "form", "")
	} else {
		var defaultValue string = "Internal"
		r.realm = &defaultValue
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

type ApiPurgeUserTokensRequest struct {
	ctx context.Context
	ApiService *UserTokensAPIService
}

func (r ApiPurgeUserTokensRequest) Execute() (*http.Response, error) {
	return r.ApiService.PurgeUserTokensExecute(r)
}

/*
PurgeUserTokens Method for PurgeUserTokens

Use this method to delete all existing LDAP user tokens.

Permissions required: Edit System Configuration and Users

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPurgeUserTokensRequest
*/
func (a *UserTokensAPIService) PurgeUserTokens(ctx context.Context) ApiPurgeUserTokensRequest {
	return ApiPurgeUserTokensRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *UserTokensAPIService) PurgeUserTokensExecute(r ApiPurgeUserTokensRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserTokensAPIService.PurgeUserTokens")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/userTokens/purge"

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
