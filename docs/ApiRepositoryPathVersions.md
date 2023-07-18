# ApiRepositoryPathVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RepositoryComponentPaths** | Pointer to [**[]ApiRepositoryComponentPath**](ApiRepositoryComponentPath.md) |  | [optional] 
**RequestIndex** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiRepositoryPathVersions

`func NewApiRepositoryPathVersions() *ApiRepositoryPathVersions`

NewApiRepositoryPathVersions instantiates a new ApiRepositoryPathVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiRepositoryPathVersionsWithDefaults

`func NewApiRepositoryPathVersionsWithDefaults() *ApiRepositoryPathVersions`

NewApiRepositoryPathVersionsWithDefaults instantiates a new ApiRepositoryPathVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepositoryComponentPaths

`func (o *ApiRepositoryPathVersions) GetRepositoryComponentPaths() []ApiRepositoryComponentPath`

GetRepositoryComponentPaths returns the RepositoryComponentPaths field if non-nil, zero value otherwise.

### GetRepositoryComponentPathsOk

`func (o *ApiRepositoryPathVersions) GetRepositoryComponentPathsOk() (*[]ApiRepositoryComponentPath, bool)`

GetRepositoryComponentPathsOk returns a tuple with the RepositoryComponentPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryComponentPaths

`func (o *ApiRepositoryPathVersions) SetRepositoryComponentPaths(v []ApiRepositoryComponentPath)`

SetRepositoryComponentPaths sets RepositoryComponentPaths field to given value.

### HasRepositoryComponentPaths

`func (o *ApiRepositoryPathVersions) HasRepositoryComponentPaths() bool`

HasRepositoryComponentPaths returns a boolean if a field has been set.

### GetRequestIndex

`func (o *ApiRepositoryPathVersions) GetRequestIndex() int32`

GetRequestIndex returns the RequestIndex field if non-nil, zero value otherwise.

### GetRequestIndexOk

`func (o *ApiRepositoryPathVersions) GetRequestIndexOk() (*int32, bool)`

GetRequestIndexOk returns a tuple with the RequestIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestIndex

`func (o *ApiRepositoryPathVersions) SetRequestIndex(v int32)`

SetRequestIndex sets RequestIndex field to given value.

### HasRequestIndex

`func (o *ApiRepositoryPathVersions) HasRequestIndex() bool`

HasRequestIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


