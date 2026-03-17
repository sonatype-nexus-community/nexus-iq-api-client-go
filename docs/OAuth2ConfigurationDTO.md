# OAuth2ConfigurationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailClaim** | Pointer to **string** |  | [optional] 
**ExactMatchClaimsJson** | Pointer to **string** |  | [optional] 
**FirstNameClaim** | Pointer to **string** |  | [optional] 
**GroupsClaim** | Pointer to **string** |  | [optional] 
**IdpIssuer** | Pointer to **string** |  | [optional] 
**IdpJwks** | Pointer to **string** |  | [optional] 
**IdpJwksUrl** | Pointer to **string** |  | [optional] 
**IdpJwsAlgorithm** | Pointer to **string** |  | [optional] 
**LastNameClaim** | Pointer to **string** |  | [optional] 
**UsernameClaim** | Pointer to **string** |  | [optional] 

## Methods

### NewOAuth2ConfigurationDTO

`func NewOAuth2ConfigurationDTO() *OAuth2ConfigurationDTO`

NewOAuth2ConfigurationDTO instantiates a new OAuth2ConfigurationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ConfigurationDTOWithDefaults

`func NewOAuth2ConfigurationDTOWithDefaults() *OAuth2ConfigurationDTO`

NewOAuth2ConfigurationDTOWithDefaults instantiates a new OAuth2ConfigurationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailClaim

`func (o *OAuth2ConfigurationDTO) GetEmailClaim() string`

GetEmailClaim returns the EmailClaim field if non-nil, zero value otherwise.

### GetEmailClaimOk

`func (o *OAuth2ConfigurationDTO) GetEmailClaimOk() (*string, bool)`

GetEmailClaimOk returns a tuple with the EmailClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailClaim

`func (o *OAuth2ConfigurationDTO) SetEmailClaim(v string)`

SetEmailClaim sets EmailClaim field to given value.

### HasEmailClaim

`func (o *OAuth2ConfigurationDTO) HasEmailClaim() bool`

HasEmailClaim returns a boolean if a field has been set.

### GetExactMatchClaimsJson

`func (o *OAuth2ConfigurationDTO) GetExactMatchClaimsJson() string`

GetExactMatchClaimsJson returns the ExactMatchClaimsJson field if non-nil, zero value otherwise.

### GetExactMatchClaimsJsonOk

`func (o *OAuth2ConfigurationDTO) GetExactMatchClaimsJsonOk() (*string, bool)`

GetExactMatchClaimsJsonOk returns a tuple with the ExactMatchClaimsJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactMatchClaimsJson

`func (o *OAuth2ConfigurationDTO) SetExactMatchClaimsJson(v string)`

SetExactMatchClaimsJson sets ExactMatchClaimsJson field to given value.

### HasExactMatchClaimsJson

`func (o *OAuth2ConfigurationDTO) HasExactMatchClaimsJson() bool`

HasExactMatchClaimsJson returns a boolean if a field has been set.

### GetFirstNameClaim

`func (o *OAuth2ConfigurationDTO) GetFirstNameClaim() string`

GetFirstNameClaim returns the FirstNameClaim field if non-nil, zero value otherwise.

### GetFirstNameClaimOk

`func (o *OAuth2ConfigurationDTO) GetFirstNameClaimOk() (*string, bool)`

GetFirstNameClaimOk returns a tuple with the FirstNameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameClaim

`func (o *OAuth2ConfigurationDTO) SetFirstNameClaim(v string)`

SetFirstNameClaim sets FirstNameClaim field to given value.

### HasFirstNameClaim

`func (o *OAuth2ConfigurationDTO) HasFirstNameClaim() bool`

HasFirstNameClaim returns a boolean if a field has been set.

### GetGroupsClaim

`func (o *OAuth2ConfigurationDTO) GetGroupsClaim() string`

GetGroupsClaim returns the GroupsClaim field if non-nil, zero value otherwise.

### GetGroupsClaimOk

`func (o *OAuth2ConfigurationDTO) GetGroupsClaimOk() (*string, bool)`

GetGroupsClaimOk returns a tuple with the GroupsClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupsClaim

`func (o *OAuth2ConfigurationDTO) SetGroupsClaim(v string)`

SetGroupsClaim sets GroupsClaim field to given value.

### HasGroupsClaim

`func (o *OAuth2ConfigurationDTO) HasGroupsClaim() bool`

HasGroupsClaim returns a boolean if a field has been set.

### GetIdpIssuer

`func (o *OAuth2ConfigurationDTO) GetIdpIssuer() string`

GetIdpIssuer returns the IdpIssuer field if non-nil, zero value otherwise.

### GetIdpIssuerOk

`func (o *OAuth2ConfigurationDTO) GetIdpIssuerOk() (*string, bool)`

GetIdpIssuerOk returns a tuple with the IdpIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpIssuer

`func (o *OAuth2ConfigurationDTO) SetIdpIssuer(v string)`

SetIdpIssuer sets IdpIssuer field to given value.

### HasIdpIssuer

`func (o *OAuth2ConfigurationDTO) HasIdpIssuer() bool`

HasIdpIssuer returns a boolean if a field has been set.

### GetIdpJwks

`func (o *OAuth2ConfigurationDTO) GetIdpJwks() string`

GetIdpJwks returns the IdpJwks field if non-nil, zero value otherwise.

### GetIdpJwksOk

`func (o *OAuth2ConfigurationDTO) GetIdpJwksOk() (*string, bool)`

GetIdpJwksOk returns a tuple with the IdpJwks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpJwks

`func (o *OAuth2ConfigurationDTO) SetIdpJwks(v string)`

SetIdpJwks sets IdpJwks field to given value.

### HasIdpJwks

`func (o *OAuth2ConfigurationDTO) HasIdpJwks() bool`

HasIdpJwks returns a boolean if a field has been set.

### GetIdpJwksUrl

`func (o *OAuth2ConfigurationDTO) GetIdpJwksUrl() string`

GetIdpJwksUrl returns the IdpJwksUrl field if non-nil, zero value otherwise.

### GetIdpJwksUrlOk

`func (o *OAuth2ConfigurationDTO) GetIdpJwksUrlOk() (*string, bool)`

GetIdpJwksUrlOk returns a tuple with the IdpJwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpJwksUrl

`func (o *OAuth2ConfigurationDTO) SetIdpJwksUrl(v string)`

SetIdpJwksUrl sets IdpJwksUrl field to given value.

### HasIdpJwksUrl

`func (o *OAuth2ConfigurationDTO) HasIdpJwksUrl() bool`

HasIdpJwksUrl returns a boolean if a field has been set.

### GetIdpJwsAlgorithm

`func (o *OAuth2ConfigurationDTO) GetIdpJwsAlgorithm() string`

GetIdpJwsAlgorithm returns the IdpJwsAlgorithm field if non-nil, zero value otherwise.

### GetIdpJwsAlgorithmOk

`func (o *OAuth2ConfigurationDTO) GetIdpJwsAlgorithmOk() (*string, bool)`

GetIdpJwsAlgorithmOk returns a tuple with the IdpJwsAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpJwsAlgorithm

`func (o *OAuth2ConfigurationDTO) SetIdpJwsAlgorithm(v string)`

SetIdpJwsAlgorithm sets IdpJwsAlgorithm field to given value.

### HasIdpJwsAlgorithm

`func (o *OAuth2ConfigurationDTO) HasIdpJwsAlgorithm() bool`

HasIdpJwsAlgorithm returns a boolean if a field has been set.

### GetLastNameClaim

`func (o *OAuth2ConfigurationDTO) GetLastNameClaim() string`

GetLastNameClaim returns the LastNameClaim field if non-nil, zero value otherwise.

### GetLastNameClaimOk

`func (o *OAuth2ConfigurationDTO) GetLastNameClaimOk() (*string, bool)`

GetLastNameClaimOk returns a tuple with the LastNameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameClaim

`func (o *OAuth2ConfigurationDTO) SetLastNameClaim(v string)`

SetLastNameClaim sets LastNameClaim field to given value.

### HasLastNameClaim

`func (o *OAuth2ConfigurationDTO) HasLastNameClaim() bool`

HasLastNameClaim returns a boolean if a field has been set.

### GetUsernameClaim

`func (o *OAuth2ConfigurationDTO) GetUsernameClaim() string`

GetUsernameClaim returns the UsernameClaim field if non-nil, zero value otherwise.

### GetUsernameClaimOk

`func (o *OAuth2ConfigurationDTO) GetUsernameClaimOk() (*string, bool)`

GetUsernameClaimOk returns a tuple with the UsernameClaim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameClaim

`func (o *OAuth2ConfigurationDTO) SetUsernameClaim(v string)`

SetUsernameClaim sets UsernameClaim field to given value.

### HasUsernameClaim

`func (o *OAuth2ConfigurationDTO) HasUsernameClaim() bool`

HasUsernameClaim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


