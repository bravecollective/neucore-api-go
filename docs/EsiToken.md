# EsiToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EveLoginId** | **int32** | ID of EveLogin | 
**CharacterId** | **int32** | ID of Character | 
**PlayerId** | **int32** | ID of Player | 
**PlayerName** | Pointer to **string** | Name of Player | [optional] 
**Character** | Pointer to [**Character**](Character.md) |  | [optional] 
**ValidToken** | **NullableBool** | Shows if the refresh token is valid or not.  This is null if there is no refresh token (EVE SSOv1 only) or a valid token but without scopes (SSOv2). | 
**ValidTokenTime** | **NullableTime** | Date and time when the valid token property was last changed. | 
**HasRoles** | **NullableBool** | Shows if the EVE character has all required roles for the login.  Null if the login does not require any roles or if the token is invalid. | 
**LastChecked** | Pointer to **NullableTime** | When the refresh token was last checked for validity. | [optional] 

## Methods

### NewEsiToken

`func NewEsiToken(eveLoginId int32, characterId int32, playerId int32, validToken NullableBool, validTokenTime NullableTime, hasRoles NullableBool, ) *EsiToken`

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


### GetCharacterId

`func (o *EsiToken) GetCharacterId() int32`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *EsiToken) GetCharacterIdOk() (*int32, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *EsiToken) SetCharacterId(v int32)`

SetCharacterId sets CharacterId field to given value.


### GetPlayerId

`func (o *EsiToken) GetPlayerId() int32`

GetPlayerId returns the PlayerId field if non-nil, zero value otherwise.

### GetPlayerIdOk

`func (o *EsiToken) GetPlayerIdOk() (*int32, bool)`

GetPlayerIdOk returns a tuple with the PlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerId

`func (o *EsiToken) SetPlayerId(v int32)`

SetPlayerId sets PlayerId field to given value.


### GetPlayerName

`func (o *EsiToken) GetPlayerName() string`

GetPlayerName returns the PlayerName field if non-nil, zero value otherwise.

### GetPlayerNameOk

`func (o *EsiToken) GetPlayerNameOk() (*string, bool)`

GetPlayerNameOk returns a tuple with the PlayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerName

`func (o *EsiToken) SetPlayerName(v string)`

SetPlayerName sets PlayerName field to given value.

### HasPlayerName

`func (o *EsiToken) HasPlayerName() bool`

HasPlayerName returns a boolean if a field has been set.

### GetCharacter

`func (o *EsiToken) GetCharacter() Character`

GetCharacter returns the Character field if non-nil, zero value otherwise.

### GetCharacterOk

`func (o *EsiToken) GetCharacterOk() (*Character, bool)`

GetCharacterOk returns a tuple with the Character field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacter

`func (o *EsiToken) SetCharacter(v Character)`

SetCharacter sets Character field to given value.

### HasCharacter

`func (o *EsiToken) HasCharacter() bool`

HasCharacter returns a boolean if a field has been set.

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
### GetLastChecked

`func (o *EsiToken) GetLastChecked() time.Time`

GetLastChecked returns the LastChecked field if non-nil, zero value otherwise.

### GetLastCheckedOk

`func (o *EsiToken) GetLastCheckedOk() (*time.Time, bool)`

GetLastCheckedOk returns a tuple with the LastChecked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChecked

`func (o *EsiToken) SetLastChecked(v time.Time)`

SetLastChecked sets LastChecked field to given value.

### HasLastChecked

`func (o *EsiToken) HasLastChecked() bool`

HasLastChecked returns a boolean if a field has been set.

### SetLastCheckedNil

`func (o *EsiToken) SetLastCheckedNil(b bool)`

 SetLastCheckedNil sets the value for LastChecked to be an explicit nil

### UnsetLastChecked
`func (o *EsiToken) UnsetLastChecked()`

UnsetLastChecked ensures that no value is present for LastChecked, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


