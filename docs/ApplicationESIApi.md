# \ApplicationESIApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsiPostV1**](ApplicationESIApi.md#EsiPostV1) | **Post** /app/v1/esi | Same as GET /app/v1/esi, but for POST requests.
[**EsiV1**](ApplicationESIApi.md#EsiV1) | **Get** /app/v1/esi | Makes an ESI GET request on behalf on an EVE character and returns the result.



## EsiPostV1

> string EsiPostV1(ctx).EsiPathQuery(esiPathQuery).Datasource(datasource).Body(body).Execute()

Same as GET /app/v1/esi, but for POST requests.

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
    datasource := "datasource_example" // string | The EVE character ID those token should be used to make the ESI request
    body := "body_example" // string | JSON encoded data.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationESIApi.EsiPostV1(context.Background()).EsiPathQuery(esiPathQuery).Datasource(datasource).Body(body).Execute()
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
 **esiPathQuery** | **string** | The ESI path and query string (without the datasource parameter). | 
 **datasource** | **string** | The EVE character ID those token should be used to make the ESI request | 
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
    datasource := "datasource_example" // string | The EVE character ID those token should be used to make the ESI request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApplicationESIApi.EsiV1(context.Background()).EsiPathQuery(esiPathQuery).Datasource(datasource).Execute()
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
 **esiPathQuery** | **string** | The ESI path and query string (without the datasource parameter). | 
 **datasource** | **string** | The EVE character ID those token should be used to make the ESI request | 

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

