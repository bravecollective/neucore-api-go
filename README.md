The version numbers correspond to the Neucore version numbers.

When updating, check the generator version in .openapi-generator/VERSION, 
a new version may break backwards compatibility.

**Breaking changes**

- 1.43.0  
  Player.character_id is now an integer instead of string.

# Go API client for neucoreapi

Client library of Neucore API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.43.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import neucoreapi "github.com/bravecollective/neucore-api-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), neucoreapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), neucoreapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), neucoreapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), neucoreapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApplicationApi* | [**ShowV1**](docs/ApplicationApi.md#showv1) | **Get** /app/v1/show | Show app information.
*ApplicationCharactersApi* | [**CharacterListV1**](docs/ApplicationCharactersApi.md#characterlistv1) | **Post** /app/v1/character-list | Returns all known characters from the parameter list.
*ApplicationCharactersApi* | [**CharactersBulkV1**](docs/ApplicationCharactersApi.md#charactersbulkv1) | **Post** /app/v1/characters | Returns all characters from multiple player accounts identified by character IDs.
*ApplicationCharactersApi* | [**CharactersV1**](docs/ApplicationCharactersApi.md#charactersv1) | **Get** /app/v1/characters/{characterId} | Returns all characters of the player account to which the character ID belongs.
*ApplicationCharactersApi* | [**CorporationCharactersV1**](docs/ApplicationCharactersApi.md#corporationcharactersv1) | **Get** /app/v1/corp-characters/{corporationId} | Returns a list of all known characters from the corporation.
*ApplicationCharactersApi* | [**CorporationPlayersV1**](docs/ApplicationCharactersApi.md#corporationplayersv1) | **Get** /app/v1/corp-players/{corporationId} | Returns a list of all players that have a character in the corporation.
*ApplicationCharactersApi* | [**IncomingCharactersV1**](docs/ApplicationCharactersApi.md#incomingcharactersv1) | **Get** /app/v1/incoming-characters/{characterId} | Returns all characters that were moved from another account to the player account to which the                     ID belongs.
*ApplicationCharactersApi* | [**MainV1**](docs/ApplicationCharactersApi.md#mainv1) | **Get** /app/v1/main/{cid} | Returns the main character of the player account to which the character ID belongs.
*ApplicationCharactersApi* | [**MainV2**](docs/ApplicationCharactersApi.md#mainv2) | **Get** /app/v2/main/{cid} | Returns the main character of the player account to which the character ID belongs.
*ApplicationCharactersApi* | [**PlayerCharactersV1**](docs/ApplicationCharactersApi.md#playercharactersv1) | **Get** /app/v1/player-chars/{playerId} | Returns all characters from the player account.
*ApplicationCharactersApi* | [**PlayerV1**](docs/ApplicationCharactersApi.md#playerv1) | **Get** /app/v1/player/{characterId} | Returns the player account to which the character ID belongs.
*ApplicationCharactersApi* | [**PlayerWithCharactersV1**](docs/ApplicationCharactersApi.md#playerwithcharactersv1) | **Get** /app/v1/player-with-characters/{characterId} | Returns the player account to which the character ID belongs with all characters.
*ApplicationCharactersApi* | [**RemovedCharactersV1**](docs/ApplicationCharactersApi.md#removedcharactersv1) | **Get** /app/v1/removed-characters/{characterId} | Returns all characters that were removed from the player account to which the character ID                     belongs.
*ApplicationESIApi* | [**EsiEveLoginCharactersV1**](docs/ApplicationESIApi.md#esievelogincharactersv1) | **Get** /app/v1/esi/eve-login/{name}/characters | Returns character IDs of characters that have an ESI token (including invalid) of an EVE login.
*ApplicationESIApi* | [**EsiEveLoginTokenDataV1**](docs/ApplicationESIApi.md#esievelogintokendatav1) | **Get** /app/v1/esi/eve-login/{name}/token-data | Returns data for all valid tokens (roles are also checked if applicable) for an EVE login.
*ApplicationESIApi* | [**EsiPostV1**](docs/ApplicationESIApi.md#esipostv1) | **Post** /app/v1/esi | See POST /app/v2/esi
*ApplicationESIApi* | [**EsiPostV2**](docs/ApplicationESIApi.md#esipostv2) | **Post** /app/v2/esi | Same as GET /app/v2/esi, but for POST requests.
*ApplicationESIApi* | [**EsiV1**](docs/ApplicationESIApi.md#esiv1) | **Get** /app/v1/esi | See GET /app/v2/esi
*ApplicationESIApi* | [**EsiV2**](docs/ApplicationESIApi.md#esiv2) | **Get** /app/v2/esi | Makes an ESI GET request on behalf on an EVE character and returns the result.
*ApplicationGroupsApi* | [**AllianceGroupsBulkV1**](docs/ApplicationGroupsApi.md#alliancegroupsbulkv1) | **Post** /app/v1/alliance-groups | Return groups of multiple alliances.
*ApplicationGroupsApi* | [**AllianceGroupsV1**](docs/ApplicationGroupsApi.md#alliancegroupsv1) | **Get** /app/v1/alliance-groups/{aid} | Return groups of the alliance.
*ApplicationGroupsApi* | [**AllianceGroupsV2**](docs/ApplicationGroupsApi.md#alliancegroupsv2) | **Get** /app/v2/alliance-groups/{aid} | Return groups of the alliance.
*ApplicationGroupsApi* | [**CorpGroupsBulkV1**](docs/ApplicationGroupsApi.md#corpgroupsbulkv1) | **Post** /app/v1/corp-groups | Return groups of multiple corporations.
*ApplicationGroupsApi* | [**CorpGroupsV1**](docs/ApplicationGroupsApi.md#corpgroupsv1) | **Get** /app/v1/corp-groups/{cid} | Return groups of the corporation.
*ApplicationGroupsApi* | [**CorpGroupsV2**](docs/ApplicationGroupsApi.md#corpgroupsv2) | **Get** /app/v2/corp-groups/{cid} | Return groups of the corporation.
*ApplicationGroupsApi* | [**GroupMembersV1**](docs/ApplicationGroupsApi.md#groupmembersv1) | **Get** /app/v1/group-members/{groupId} | Returns the main character IDs from all group members.
*ApplicationGroupsApi* | [**GroupsBulkV1**](docs/ApplicationGroupsApi.md#groupsbulkv1) | **Post** /app/v1/groups | Return groups of multiple players, identified by one of their character IDs.
*ApplicationGroupsApi* | [**GroupsV1**](docs/ApplicationGroupsApi.md#groupsv1) | **Get** /app/v1/groups/{cid} | Return groups of the character&#39;s player account.
*ApplicationGroupsApi* | [**GroupsV2**](docs/ApplicationGroupsApi.md#groupsv2) | **Get** /app/v2/groups/{cid} | Return groups of the character&#39;s player account.
*ApplicationGroupsApi* | [**GroupsWithFallbackV1**](docs/ApplicationGroupsApi.md#groupswithfallbackv1) | **Get** /app/v1/groups-with-fallback | Returns groups from the character&#39;s account, if available, or the corporation and alliance.
*ApplicationTrackingApi* | [**MemberTrackingV1**](docs/ApplicationTrackingApi.md#membertrackingv1) | **Get** /app/v1/corporation/{id}/member-tracking | Return corporation member tracking data.


## Documentation For Models

 - [Alliance](docs/Alliance.md)
 - [App](docs/App.md)
 - [Character](docs/Character.md)
 - [CharacterGroups](docs/CharacterGroups.md)
 - [CharacterNameChange](docs/CharacterNameChange.md)
 - [Corporation](docs/Corporation.md)
 - [CorporationMember](docs/CorporationMember.md)
 - [EsiLocation](docs/EsiLocation.md)
 - [EsiToken](docs/EsiToken.md)
 - [EsiTokenData](docs/EsiTokenData.md)
 - [EsiType](docs/EsiType.md)
 - [EveLogin](docs/EveLogin.md)
 - [Group](docs/Group.md)
 - [GroupApplication](docs/GroupApplication.md)
 - [HourlyAppRequests](docs/HourlyAppRequests.md)
 - [MonthlyAppRequests](docs/MonthlyAppRequests.md)
 - [Player](docs/Player.md)
 - [PlayerLoginStatistics](docs/PlayerLoginStatistics.md)
 - [Plugin](docs/Plugin.md)
 - [PluginConfigurationDatabase](docs/PluginConfigurationDatabase.md)
 - [PluginConfigurationFile](docs/PluginConfigurationFile.md)
 - [PluginConfigurationURL](docs/PluginConfigurationURL.md)
 - [RemovedCharacter](docs/RemovedCharacter.md)
 - [Role](docs/Role.md)
 - [SearchResult](docs/SearchResult.md)
 - [ServiceAccount](docs/ServiceAccount.md)
 - [SystemVariable](docs/SystemVariable.md)
 - [TotalDailyAppRequests](docs/TotalDailyAppRequests.md)
 - [TotalMonthlyAppRequests](docs/TotalMonthlyAppRequests.md)
 - [Watchlist](docs/Watchlist.md)


## Documentation For Authorization



### BearerAuth

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



