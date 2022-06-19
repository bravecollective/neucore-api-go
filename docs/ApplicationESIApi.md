# \ApplicationESIApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsiEveLoginCharactersV1**](ApplicationESIApi.md#EsiEveLoginCharactersV1) | **Get** /app/v1/esi/eve-login/{name}/characters | Returns character IDs of characters that have a valid ESI token of the specified EVE login.
[**EsiPostV1**](ApplicationESIApi.md#EsiPostV1) | **Post** /app/v1/esi | See POST /app/v2/esi
[**EsiPostV2**](ApplicationESIApi.md#EsiPostV2) | **Post** /app/v2/esi | Same as GET /app/v2/esi, but for POST requests.
[**EsiV1**](ApplicationESIApi.md#EsiV1) | **Get** /app/v1/esi | See GET /app/v2/esi
[**EsiV2**](ApplicationESIApi.md#EsiV2) | **Get** /app/v2/esi | Makes an ESI GET request on behalf on an EVE character and returns the result.



## EsiEveLoginCharactersV1

> []int32 EsiEveLoginCharactersV1(ctx, name).Execute()

Returns character IDs of characters that have a valid ESI token of the specified EVE login.



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
    name := "name_example" // string | EVE login name, 'core.default' is not allowed.

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
**name** | **string** | EVE login name, &#39;core.default&#39; is not allowed. | 

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


## EsiPostV1

> string EsiPostV1(ctx).EsiPathQuery(esiPathQuery).Datasource(datasource).Body(body).Execute()

See POST /app/v2/esi

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
    esiPathQuery := "esiPathQuery_example" // string | 
    datasource := "datasource_example" // string | 
    body := "body_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationESIApi.EsiPostV1(context.Background()).EsiPathQuery(esiPathQuery).Datasource(datasource).Body(body).Execute()
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

> string EsiPostV2(ctx).EsiPathQuery(esiPathQuery).Datasource(datasource).Body(body).Execute()

Same as GET /app/v2/esi, but for POST requests.

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
    esiPathQuery := "esiPathQuery_example" // string | 
    datasource := "datasource_example" // string | 
    body := "body_example" // string | JSON encoded data.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationESIApi.EsiPostV2(context.Background()).EsiPathQuery(esiPathQuery).Datasource(datasource).Body(body).Execute()
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

> string EsiV1(ctx).EsiPathQuery(esiPathQuery).Datasource(datasource).Execute()

See GET /app/v2/esi

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
    esiPathQuery := "esiPathQuery_example" // string | 
    datasource := "datasource_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationESIApi.EsiV1(context.Background()).EsiPathQuery(esiPathQuery).Datasource(datasource).Execute()
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

> string EsiV2(ctx).EsiPathQuery(esiPathQuery).Datasource(datasource).Execute()

Makes an ESI GET request on behalf on an EVE character and returns the result.



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
    esiPathQuery := "esiPathQuery_example" // string | The ESI path and query string (without the datasource parameter).
    datasource := "datasource_example" // string | The EVE character ID those token should be used to make the ESI request. Optionally                             followed by a colon and the name of an EVE login to use an alternative ESI token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationESIApi.EsiV2(context.Background()).EsiPathQuery(esiPathQuery).Datasource(datasource).Execute()
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
 **datasource** | **string** | The EVE character ID those token should be used to make the ESI request. Optionally                             followed by a colon and the name of an EVE login to use an alternative ESI token. | 

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

