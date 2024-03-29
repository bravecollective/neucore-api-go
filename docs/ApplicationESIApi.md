# \ApplicationESIApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsiAccessTokenV1**](ApplicationESIApi.md#EsiAccessTokenV1) | **Get** /app/v1/esi/access-token/{characterId} | Returns an access token for a character and EVE login.
[**EsiEveLoginCharactersV1**](ApplicationESIApi.md#EsiEveLoginCharactersV1) | **Get** /app/v1/esi/eve-login/{name}/characters | Returns character IDs of characters that have an ESI token (including invalid) of an EVE login.
[**EsiEveLoginTokenDataV1**](ApplicationESIApi.md#EsiEveLoginTokenDataV1) | **Get** /app/v1/esi/eve-login/{name}/token-data | Returns data for all valid tokens (roles are also checked if applicable) for an EVE login.
[**EsiPostV1**](ApplicationESIApi.md#EsiPostV1) | **Post** /app/v1/esi | See POST /app/v2/esi
[**EsiPostV2**](ApplicationESIApi.md#EsiPostV2) | **Post** /app/v2/esi | Same as GET /app/v2/esi, but for POST requests.
[**EsiV1**](ApplicationESIApi.md#EsiV1) | **Get** /app/v1/esi | See GET /app/v2/esi
[**EsiV2**](ApplicationESIApi.md#EsiV2) | **Get** /app/v2/esi | Makes an ESI GET request on behalf on an EVE character and returns the result.



## EsiAccessTokenV1

> EsiAccessToken EsiAccessTokenV1(ctx, characterId).EveLoginName(eveLoginName).Execute()

Returns an access token for a character and EVE login.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bravecollective/neucore-api-go"
)

func main() {
    characterId := int32(56) // int32 | The EVE character ID.
    eveLoginName := "eveLoginName_example" // string | Optional EVE login name, defaults to 'core.default'. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationESIApi.EsiAccessTokenV1(context.Background(), characterId).EveLoginName(eveLoginName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationESIApi.EsiAccessTokenV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EsiAccessTokenV1`: EsiAccessToken
    fmt.Fprintf(os.Stdout, "Response from `ApplicationESIApi.EsiAccessTokenV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**characterId** | **int32** | The EVE character ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsiAccessTokenV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **eveLoginName** | **string** | Optional EVE login name, defaults to &#39;core.default&#39;. | 

### Return type

[**EsiAccessToken**](EsiAccessToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsiEveLoginCharactersV1

> []int32 EsiEveLoginCharactersV1(ctx, name).Execute()

Returns character IDs of characters that have an ESI token (including invalid) of an EVE login.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bravecollective/neucore-api-go"
)

func main() {
    name := "name_example" // string | EVE login name.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationESIApi.EsiEveLoginCharactersV1(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationESIApi.EsiEveLoginCharactersV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EsiEveLoginCharactersV1`: []int32
    fmt.Fprintf(os.Stdout, "Response from `ApplicationESIApi.EsiEveLoginCharactersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | EVE login name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsiEveLoginCharactersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]int32**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsiEveLoginTokenDataV1

> []EsiTokenData EsiEveLoginTokenDataV1(ctx, name).Execute()

Returns data for all valid tokens (roles are also checked if applicable) for an EVE login.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bravecollective/neucore-api-go"
)

func main() {
    name := "name_example" // string | EVE login name, 'core.default' is not allowed.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationESIApi.EsiEveLoginTokenDataV1(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationESIApi.EsiEveLoginTokenDataV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EsiEveLoginTokenDataV1`: []EsiTokenData
    fmt.Fprintf(os.Stdout, "Response from `ApplicationESIApi.EsiEveLoginTokenDataV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | EVE login name, &#39;core.default&#39; is not allowed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEsiEveLoginTokenDataV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EsiTokenData**](EsiTokenData.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsiPostV1

> string EsiPostV1(ctx).EsiPathQuery(esiPathQuery).Datasource(datasource).Body(body).NeucoreEveCharacter(neucoreEveCharacter).NeucoreEveLogin(neucoreEveLogin).Execute()

See POST /app/v2/esi

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bravecollective/neucore-api-go"
)

func main() {
    esiPathQuery := "esiPathQuery_example" // string | 
    datasource := "datasource_example" // string | 
    body := "body_example" // string | 
    neucoreEveCharacter := "neucoreEveCharacter_example" // string | The EVE character ID those token should be used. Has priority over the query      *                       parameter 'datasource' (optional)
    neucoreEveLogin := "neucoreEveLogin_example" // string | The EVE login name from which the token should be used, defaults to core.default. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationESIApi.EsiPostV1(context.Background()).EsiPathQuery(esiPathQuery).Datasource(datasource).Body(body).NeucoreEveCharacter(neucoreEveCharacter).NeucoreEveLogin(neucoreEveLogin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationESIApi.EsiPostV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EsiPostV1`: string
    fmt.Fprintf(os.Stdout, "Response from `ApplicationESIApi.EsiPostV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsiPostV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esiPathQuery** | **string** |  | 
 **datasource** | **string** |  | 
 **body** | **string** |  | 
 **neucoreEveCharacter** | **string** | The EVE character ID those token should be used. Has priority over the query      *                       parameter &#39;datasource&#39; | 
 **neucoreEveLogin** | **string** | The EVE login name from which the token should be used, defaults to core.default. | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsiPostV2

> string EsiPostV2(ctx).EsiPathQuery(esiPathQuery).Datasource(datasource).Body(body).NeucoreEveCharacter(neucoreEveCharacter).NeucoreEveLogin(neucoreEveLogin).Execute()

Same as GET /app/v2/esi, but for POST requests.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bravecollective/neucore-api-go"
)

func main() {
    esiPathQuery := "esiPathQuery_example" // string | 
    datasource := "datasource_example" // string | 
    body := "body_example" // string | JSON encoded data.
    neucoreEveCharacter := "neucoreEveCharacter_example" // string | The EVE character ID those token should be used. Has priority over the query      *                       parameter 'datasource' (optional)
    neucoreEveLogin := "neucoreEveLogin_example" // string | The EVE login name from which the token should be used, defaults to core.default. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationESIApi.EsiPostV2(context.Background()).EsiPathQuery(esiPathQuery).Datasource(datasource).Body(body).NeucoreEveCharacter(neucoreEveCharacter).NeucoreEveLogin(neucoreEveLogin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationESIApi.EsiPostV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EsiPostV2`: string
    fmt.Fprintf(os.Stdout, "Response from `ApplicationESIApi.EsiPostV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsiPostV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esiPathQuery** | **string** |  | 
 **datasource** | **string** |  | 
 **body** | **string** | JSON encoded data. | 
 **neucoreEveCharacter** | **string** | The EVE character ID those token should be used. Has priority over the query      *                       parameter &#39;datasource&#39; | 
 **neucoreEveLogin** | **string** | The EVE login name from which the token should be used, defaults to core.default. | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsiV1

> string EsiV1(ctx).EsiPathQuery(esiPathQuery).Datasource(datasource).NeucoreEveCharacter(neucoreEveCharacter).NeucoreEveLogin(neucoreEveLogin).Execute()

See GET /app/v2/esi

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bravecollective/neucore-api-go"
)

func main() {
    esiPathQuery := "esiPathQuery_example" // string | 
    datasource := "datasource_example" // string | 
    neucoreEveCharacter := "neucoreEveCharacter_example" // string | The EVE character ID those token should be used. Has priority over the query      *                       parameter 'datasource' (optional)
    neucoreEveLogin := "neucoreEveLogin_example" // string | The EVE login name from which the token should be used, defaults to core.default. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationESIApi.EsiV1(context.Background()).EsiPathQuery(esiPathQuery).Datasource(datasource).NeucoreEveCharacter(neucoreEveCharacter).NeucoreEveLogin(neucoreEveLogin).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationESIApi.EsiV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EsiV1`: string
    fmt.Fprintf(os.Stdout, "Response from `ApplicationESIApi.EsiV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsiV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esiPathQuery** | **string** |  | 
 **datasource** | **string** |  | 
 **neucoreEveCharacter** | **string** | The EVE character ID those token should be used. Has priority over the query      *                       parameter &#39;datasource&#39; | 
 **neucoreEveLogin** | **string** | The EVE login name from which the token should be used, defaults to core.default. | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EsiV2

> string EsiV2(ctx).EsiPathQuery(esiPathQuery).NeucoreEveCharacter(neucoreEveCharacter).NeucoreEveLogin(neucoreEveLogin).Datasource(datasource).Execute()

Makes an ESI GET request on behalf on an EVE character and returns the result.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/bravecollective/neucore-api-go"
)

func main() {
    esiPathQuery := "esiPathQuery_example" // string | The ESI path and query string (without the datasource parameter).
    neucoreEveCharacter := "neucoreEveCharacter_example" // string | The EVE character ID those token should be used. Has priority over the query                             parameter 'datasource' (optional)
    neucoreEveLogin := "neucoreEveLogin_example" // string | The EVE login name from which the token should be used, defaults to core.default. (optional)
    datasource := "datasource_example" // string | The EVE character ID those token should be used from the default login to make the ESI                             request. Optionally followed by a colon and the name of an EVE login to use an alternative                             ESI token. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationESIApi.EsiV2(context.Background()).EsiPathQuery(esiPathQuery).NeucoreEveCharacter(neucoreEveCharacter).NeucoreEveLogin(neucoreEveLogin).Datasource(datasource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationESIApi.EsiV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EsiV2`: string
    fmt.Fprintf(os.Stdout, "Response from `ApplicationESIApi.EsiV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEsiV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **esiPathQuery** | **string** | The ESI path and query string (without the datasource parameter). | 
 **neucoreEveCharacter** | **string** | The EVE character ID those token should be used. Has priority over the query                             parameter &#39;datasource&#39; | 
 **neucoreEveLogin** | **string** | The EVE login name from which the token should be used, defaults to core.default. | 
 **datasource** | **string** | The EVE character ID those token should be used from the default login to make the ESI                             request. Optionally followed by a colon and the name of an EVE login to use an alternative                             ESI token. | 

### Return type

**string**

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

