# \ApplicationCharactersApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CharacterListV1**](ApplicationCharactersApi.md#CharacterListV1) | **Post** /app/v1/character-list | Returns all known characters from the parameter list.
[**CharactersBulkV1**](ApplicationCharactersApi.md#CharactersBulkV1) | **Post** /app/v1/characters | Returns all characters from multiple player accounts identified by character IDs.
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

> []Character CharacterListV1(ctx).RequestBody(requestBody).Execute()

Returns all known characters from the parameter list.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestBody := []int32{int32(123)} // []int32 | Array with EVE character IDs.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCharactersApi.CharacterListV1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCharactersApi.CharacterListV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CharacterListV1`: []Character
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCharactersApi.CharacterListV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCharacterListV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]int32** | Array with EVE character IDs. | 

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


## CharactersBulkV1

> [][]int32 CharactersBulkV1(ctx).RequestBody(requestBody).Execute()

Returns all characters from multiple player accounts identified by character IDs.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestBody := []int32{int32(123)} // []int32 | EVE character IDs array.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCharactersApi.CharactersBulkV1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCharactersApi.CharactersBulkV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CharactersBulkV1`: [][]int32
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCharactersApi.CharactersBulkV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCharactersBulkV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]int32** | EVE character IDs array. | 

### Return type

[**[][]int32**](array.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CharactersV1

> []Character CharactersV1(ctx, characterId).Execute()

Returns all characters of the player account to which the character ID belongs.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    characterId := int32(56) // int32 | EVE character ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCharactersApi.CharactersV1(context.Background(), characterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCharactersApi.CharactersV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CharactersV1`: []Character
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCharactersApi.CharactersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int32** | EVE character ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCharactersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []Character CorporationCharactersV1(ctx, corporationId).Execute()

Returns a list of all known characters from the corporation.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    corporationId := int32(56) // int32 | EVE corporation ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCharactersApi.CorporationCharactersV1(context.Background(), corporationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCharactersApi.CorporationCharactersV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorporationCharactersV1`: []Character
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCharactersApi.CorporationCharactersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int32** | EVE corporation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCorporationCharactersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []Player CorporationPlayersV1(ctx, corporationId).Execute()

Returns a list of all players that have a character in the corporation.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    corporationId := int32(56) // int32 | EVE corporation ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCharactersApi.CorporationPlayersV1(context.Background(), corporationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCharactersApi.CorporationPlayersV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorporationPlayersV1`: []Player
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCharactersApi.CorporationPlayersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**corporationId** | **int32** | EVE corporation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCorporationPlayersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []RemovedCharacter IncomingCharactersV1(ctx, characterId).Execute()

Returns all characters that were moved from another account to the player account to which the                     ID belongs.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    characterId := int32(56) // int32 | EVE character ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCharactersApi.IncomingCharactersV1(context.Background(), characterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCharactersApi.IncomingCharactersV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomingCharactersV1`: []RemovedCharacter
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCharactersApi.IncomingCharactersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int32** | EVE character ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiIncomingCharactersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> Character MainV1(ctx, cid).Execute()

Returns the main character of the player account to which the character ID belongs.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cid := int32(56) // int32 | EVE character ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCharactersApi.MainV1(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCharactersApi.MainV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MainV1`: Character
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCharactersApi.MainV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int32** | EVE character ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMainV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> Character MainV2(ctx, cid).Execute()

Returns the main character of the player account to which the character ID belongs.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cid := int32(56) // int32 | EVE character ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCharactersApi.MainV2(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCharactersApi.MainV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MainV2`: Character
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCharactersApi.MainV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int32** | EVE character ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMainV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []Character PlayerCharactersV1(ctx, playerId).Execute()

Returns all characters from the player account.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    playerId := int32(56) // int32 | Player ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCharactersApi.PlayerCharactersV1(context.Background(), playerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCharactersApi.PlayerCharactersV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayerCharactersV1`: []Character
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCharactersApi.PlayerCharactersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**playerId** | **int32** | Player ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlayerCharactersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> Player PlayerV1(ctx, characterId).Execute()

Returns the player account to which the character ID belongs.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    characterId := int32(56) // int32 | EVE character ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCharactersApi.PlayerV1(context.Background(), characterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCharactersApi.PlayerV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayerV1`: Player
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCharactersApi.PlayerV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int32** | EVE character ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlayerV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> Player PlayerWithCharactersV1(ctx, characterId).Execute()

Returns the player account to which the character ID belongs with all characters.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    characterId := int32(56) // int32 | EVE character ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCharactersApi.PlayerWithCharactersV1(context.Background(), characterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCharactersApi.PlayerWithCharactersV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PlayerWithCharactersV1`: Player
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCharactersApi.PlayerWithCharactersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int32** | EVE character ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlayerWithCharactersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []RemovedCharacter RemovedCharactersV1(ctx, characterId).Execute()

Returns all characters that were removed from the player account to which the character ID                     belongs.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    characterId := int32(56) // int32 | EVE character ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationCharactersApi.RemovedCharactersV1(context.Background(), characterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationCharactersApi.RemovedCharactersV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemovedCharactersV1`: []RemovedCharacter
    fmt.Fprintf(os.Stdout, "Response from `ApplicationCharactersApi.RemovedCharactersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int32** | EVE character ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovedCharactersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

