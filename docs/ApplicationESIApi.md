# \ApplicationESIApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EsiPostV1**](ApplicationESIApi.md#EsiPostV1) | **Post** /app/v1/esi | Same as GET /app/v1/esi, but for POST requests.
[**EsiV1**](ApplicationESIApi.md#EsiV1) | **Get** /app/v1/esi | Makes an ESI GET request on behalf on an EVE character and returns the result.



## EsiPostV1

> string EsiPostV1(ctx, esiPathQuery, datasource, body)

Same as GET /app/v1/esi, but for POST requests.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esiPathQuery** | **string**| The ESI path and query string (without the datasource parameter). | 
**datasource** | **string**| The EVE character ID those token should be used to make the ESI request | 
**body** | **string**| JSON encoded data. | 

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

> string EsiV1(ctx, esiPathQuery, datasource)

Makes an ESI GET request on behalf on an EVE character and returns the result.

Needs role: app-esi<br>      *         Public ESI routes are not allowed.<br>      *         The following headers from ESI are passed through to the response if they exist:                Content-Type Expires X-Esi-Error-Limit-Remain X-Esi-Error-Limit-Reset X-Pages warning, Warning<br>      *         The HTTP status code from ESI is also passed through, so there may be more than the documented ones.<br>      *         The ESI path and query parameters can alternatively be appended to the path of this endpoint,                this allows to use OpenAPI clients that were generated for the ESI API,                see doc/app-esi-examples.php for more.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**esiPathQuery** | **string**| The ESI path and query string (without the datasource parameter). | 
**datasource** | **string**| The EVE character ID those token should be used to make the ESI request | 

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

