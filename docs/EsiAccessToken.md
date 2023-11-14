# EsiAccessToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Scopes** | **[]string** |  | 
**Expires** | **int32** |  | 

## Methods

### NewEsiAccessToken

`func NewEsiAccessToken(token string, scopes []string, expires int32, ) *EsiAccessToken`

NewEsiAccessToken instantiates a new EsiAccessToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsiAccessTokenWithDefaults

`func NewEsiAccessTokenWithDefaults() *EsiAccessToken`

NewEsiAccessTokenWithDefaults instantiates a new EsiAccessToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *EsiAccessToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EsiAccessToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EsiAccessToken) SetToken(v string)`

SetToken sets Token field to given value.


### GetScopes

`func (o *EsiAccessToken) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *EsiAccessToken) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *EsiAccessToken) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetExpires

`func (o *EsiAccessToken) GetExpires() int32`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *EsiAccessToken) GetExpiresOk() (*int32, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *EsiAccessToken) SetExpires(v int32)`

SetExpires sets Expires field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


