# \ApplicationCharactersApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CharacterListV1**](ApplicationCharactersApi.md#CharacterListV1) | **Post** /app/v1/character-list | Returns all known characters from the parameter list.
[**CharactersV1**](ApplicationCharactersApi.md#CharactersV1) | **Get** /app/v1/characters/{characterId} | Returns all characters of the player account to which the character ID belongs.
[**CorporationCharactersV1**](ApplicationCharactersApi.md#CorporationCharactersV1) | **Get** /app/v1/corp-characters/{corporationId} | Returns a list of all known characters from the corporation.
[**CorporationPlayersV1**](ApplicationCharactersApi.md#CorporationPlayersV1) | **Get** /app/v1/corp-players/{corporationId} | Returns a list of all players that have a character in the corporation.
[**IncomingCharactersV1**](ApplicationCharactersApi.md#IncomingCharactersV1) | **Get** /app/v1/incoming-characters/{characterId} | Returns all characters that were moved from another account to the player account to which the                     ID belongs.
[**MainV1**](ApplicationCharactersApi.md#MainV1) | **Get** /app/v1/main/{cid} | Returns the main character of the player account to which the character ID belongs.
[**MainV2**](ApplicationCharactersApi.md#MainV2) | **Get** /app/v2/main/{cid} | Returns the main character of the player account to which the character ID belongs.
[**PlayerCharactersV1**](ApplicationCharactersApi.md#PlayerCharactersV1) | **Get** /app/v1/player-chars/{playerId} | Returns all characters from the player account.
[**PlayerV1**](ApplicationCharactersApi.md#PlayerV1) | **Get** /app/v1/player/{characterId} | Returns the player account to which the character ID belongs.
[**PlayerWithCharactersV1**](ApplicationCharactersApi.md#PlayerWithCharactersV1) | **Get** /app/v1/player-with-characters/{characterId} | Returns the player account to which the character ID belongs with all characters.
[**RemovedCharactersV1**](ApplicationCharactersApi.md#RemovedCharactersV1) | **Get** /app/v1/removed-characters/{characterId} | Returns all characters that were removed from the player account to which the character ID                     belongs.



## CharacterListV1

> []Character CharacterListV1(ctx, requestBody)

Returns all known characters from the parameter list.

Needs role: app-chars.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestBody** | [**[]int32**](int32.md)| Array with EVE character IDs. | 

### Return type

[**[]Character**](Character.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CharactersV1

> []Character CharactersV1(ctx, characterId)

Returns all characters of the player account to which the character ID belongs.

Needs role: app-chars.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int32**| EVE character ID. | 

### Return type

[**[]Character**](Character.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CorporationCharactersV1

> []Character CorporationCharactersV1(ctx, corporationId)

Returns a list of all known characters from the corporation.

Needs role: app-chars.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int32**| EVE corporation ID. | 

### Return type

[**[]Character**](Character.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CorporationPlayersV1

> []Player CorporationPlayersV1(ctx, corporationId)

Returns a list of all players that have a character in the corporation.

Needs role: app-chars.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int32**| EVE corporation ID. | 

### Return type

[**[]Player**](Player.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomingCharactersV1

> []RemovedCharacter IncomingCharactersV1(ctx, characterId)

Returns all characters that were moved from another account to the player account to which the                     ID belongs.

Needs role: app-chars.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int32**| EVE character ID. | 

### Return type

[**[]RemovedCharacter**](RemovedCharacter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MainV1

> Character MainV1(ctx, cid)

Returns the main character of the player account to which the character ID belongs.

Needs role: app-chars.<br>It is possible that an account has no main character.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int32**| EVE character ID. | 

### Return type

[**Character**](Character.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MainV2

> Character MainV2(ctx, cid)

Returns the main character of the player account to which the character ID belongs.

Needs role: app-chars.<br>It is possible that an account has no main character.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int32**| EVE character ID. | 

### Return type

[**Character**](Character.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayerCharactersV1

> []Character PlayerCharactersV1(ctx, playerId)

Returns all characters from the player account.

Needs role: app-chars.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **int32**| Player ID. | 

### Return type

[**[]Character**](Character.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayerV1

> Player PlayerV1(ctx, characterId)

Returns the player account to which the character ID belongs.

Needs role: app-chars.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int32**| EVE character ID. | 

### Return type

[**Player**](Player.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlayerWithCharactersV1

> Player PlayerWithCharactersV1(ctx, characterId)

Returns the player account to which the character ID belongs with all characters.

Needs role: app-chars.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int32**| EVE character ID. | 

### Return type

[**Player**](Player.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovedCharactersV1

> []RemovedCharacter RemovedCharactersV1(ctx, characterId)

Returns all characters that were removed from the player account to which the character ID                     belongs.

Needs role: app-chars.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int32**| EVE character ID. | 

### Return type

[**[]RemovedCharacter**](RemovedCharacter.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

