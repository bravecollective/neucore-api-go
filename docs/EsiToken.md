# EsiToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EveLoginId** | **int32** | ID of EveLogin | 
**ValidToken** | **NullableBool** | Shows if the refresh token is valid or not.  This is null if there is no refresh token (EVE SSOv1 only) or a valid token but without scopes (SSOv2). | 
**ValidTokenTime** | **NullableTime** | Date and time when the valid token property was last changed. | 
**HasRoles** | **NullableBool** | Shows if the EVE character has all required roles for the login.  Null if the login does not require any roles or if the token is invalid. | 

## Methods

### NewEsiToken

`func NewEsiToken(eveLoginId int32, validToken NullableBool, validTokenTime NullableTime, hasRoles NullableBool, ) *EsiToken`

NewEsiToken instantiates a new EsiToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEsiTokenWithDefaults

`func NewEsiTokenWithDefaults() *EsiToken`

NewEsiTokenWithDefaults instantiates a new EsiToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEveLoginId

`func (o *EsiToken) GetEveLoginId() int32`

GetEveLoginId returns the EveLoginId field if non-nil, zero value otherwise.

### GetEveLoginIdOk

`func (o *EsiToken) GetEveLoginIdOk() (*int32, bool)`

GetEveLoginIdOk returns a tuple with the EveLoginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEveLoginId

`func (o *EsiToken) SetEveLoginId(v int32)`

SetEveLoginId sets EveLoginId field to given value.


### GetValidToken

`func (o *EsiToken) GetValidToken() bool`

GetValidToken returns the ValidToken field if non-nil, zero value otherwise.

### GetValidTokenOk

`func (o *EsiToken) GetValidTokenOk() (*bool, bool)`

GetValidTokenOk returns a tuple with the ValidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidToken

`func (o *EsiToken) SetValidToken(v bool)`

SetValidToken sets ValidToken field to given value.


### SetValidTokenNil

`func (o *EsiToken) SetValidTokenNil(b bool)`

 SetValidTokenNil sets the value for ValidToken to be an explicit nil

### UnsetValidToken
`func (o *EsiToken) UnsetValidToken()`

UnsetValidToken ensures that no value is present for ValidToken, not even an explicit nil
### GetValidTokenTime

`func (o *EsiToken) GetValidTokenTime() time.Time`

GetValidTokenTime returns the ValidTokenTime field if non-nil, zero value otherwise.

### GetValidTokenTimeOk

`func (o *EsiToken) GetValidTokenTimeOk() (*time.Time, bool)`

GetValidTokenTimeOk returns a tuple with the ValidTokenTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTokenTime

`func (o *EsiToken) SetValidTokenTime(v time.Time)`

SetValidTokenTime sets ValidTokenTime field to given value.


### SetValidTokenTimeNil

`func (o *EsiToken) SetValidTokenTimeNil(b bool)`

 SetValidTokenTimeNil sets the value for ValidTokenTime to be an explicit nil

### UnsetValidTokenTime
`func (o *EsiToken) UnsetValidTokenTime()`

UnsetValidTokenTime ensures that no value is present for ValidTokenTime, not even an explicit nil
### GetHasRoles

`func (o *EsiToken) GetHasRoles() bool`

GetHasRoles returns the HasRoles field if non-nil, zero value otherwise.

### GetHasRolesOk

`func (o *EsiToken) GetHasRolesOk() (*bool, bool)`

GetHasRolesOk returns a tuple with the HasRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRoles

`func (o *EsiToken) SetHasRoles(v bool)`

SetHasRoles sets HasRoles field to given value.


### SetHasRolesNil

`func (o *EsiToken) SetHasRolesNil(b bool)`

 SetHasRolesNil sets the value for HasRoles to be an explicit nil

### UnsetHasRoles
`func (o *EsiToken) UnsetHasRoles()`

UnsetHasRoles ensures that no value is present for HasRoles, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


