# \ApplicationGroupsApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllianceGroupsBulkV1**](ApplicationGroupsApi.md#AllianceGroupsBulkV1) | **Post** /app/v1/alliance-groups | Return groups of multiple alliances.
[**AllianceGroupsV1**](ApplicationGroupsApi.md#AllianceGroupsV1) | **Get** /app/v1/alliance-groups/{aid} | Return groups of the alliance.
[**AllianceGroupsV2**](ApplicationGroupsApi.md#AllianceGroupsV2) | **Get** /app/v2/alliance-groups/{aid} | Return groups of the alliance.
[**CorpGroupsBulkV1**](ApplicationGroupsApi.md#CorpGroupsBulkV1) | **Post** /app/v1/corp-groups | Return groups of multiple corporations.
[**CorpGroupsV1**](ApplicationGroupsApi.md#CorpGroupsV1) | **Get** /app/v1/corp-groups/{cid} | Return groups of the corporation.
[**CorpGroupsV2**](ApplicationGroupsApi.md#CorpGroupsV2) | **Get** /app/v2/corp-groups/{cid} | Return groups of the corporation.
[**GroupsBulkV1**](ApplicationGroupsApi.md#GroupsBulkV1) | **Post** /app/v1/groups | Return groups of multiple players, identified by one of their character IDs.
[**GroupsV1**](ApplicationGroupsApi.md#GroupsV1) | **Get** /app/v1/groups/{cid} | Return groups of the character&#39;s player account.
[**GroupsV2**](ApplicationGroupsApi.md#GroupsV2) | **Get** /app/v2/groups/{cid} | Return groups of the character&#39;s player account.
[**GroupsWithFallbackV1**](ApplicationGroupsApi.md#GroupsWithFallbackV1) | **Get** /app/v1/groups-with-fallback | Returns groups from the character&#39;s account, if available, or the corporation and alliance.



## AllianceGroupsBulkV1

> []Alliance AllianceGroupsBulkV1(ctx, requestBody)

Return groups of multiple alliances.

Needs role: app-groups.<br>      *                  Returns only groups that have been added to the app as well.      *                  Skips alliances that are not found in the local database.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestBody** | [**[]int32**](int32.md)| EVE alliance IDs array. | 

### Return type

[**[]Alliance**](Alliance.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllianceGroupsV1

> []Group AllianceGroupsV1(ctx, aid)

Return groups of the alliance.

Needs role: app-groups.<br>Returns only groups that have been added to the app as well.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aid** | **int32**| EVE alliance ID. | 

### Return type

[**[]Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AllianceGroupsV2

> []Group AllianceGroupsV2(ctx, aid)

Return groups of the alliance.

Needs role: app-groups.<br>Returns only groups that have been added to the app as well.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aid** | **int32**| EVE alliance ID. | 

### Return type

[**[]Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CorpGroupsBulkV1

> []Corporation CorpGroupsBulkV1(ctx, requestBody)

Return groups of multiple corporations.

Needs role: app-groups.<br>      *                  Returns only groups that have been added to the app as well.      *                  Skips corporations that are not found in the local database.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestBody** | [**[]int32**](int32.md)| EVE corporation IDs array. | 

### Return type

[**[]Corporation**](Corporation.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CorpGroupsV1

> []Group CorpGroupsV1(ctx, cid)

Return groups of the corporation.

Needs role: app-groups.<br>Returns only groups that have been added to the app as well.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int32**| EVE corporation ID. | 

### Return type

[**[]Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CorpGroupsV2

> []Group CorpGroupsV2(ctx, cid)

Return groups of the corporation.

Needs role: app-groups.<br>Returns only groups that have been added to the app as well.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int32**| EVE corporation ID. | 

### Return type

[**[]Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsBulkV1

> []CharacterGroups GroupsBulkV1(ctx, requestBody)

Return groups of multiple players, identified by one of their character IDs.

Needs role: app-groups.<br>      *                  Returns only groups that have been added to the app as well.      *                  Skips characters that are not found in the local database.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestBody** | [**[]int32**](int32.md)| EVE character IDs array. | 

### Return type

[**[]CharacterGroups**](CharacterGroups.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsV1

> []Group GroupsV1(ctx, cid)

Return groups of the character's player account.

Needs role: app-groups.<br>Returns only groups that have been added to the app as well.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int32**| EVE character ID. | 

### Return type

[**[]Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsV2

> []Group GroupsV2(ctx, cid)

Return groups of the character's player account.

Needs role: app-groups.<br>Returns only groups that have been added to the app as well.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int32**| EVE character ID. | 

### Return type

[**[]Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsWithFallbackV1

> []Group GroupsWithFallbackV1(ctx, character, corporation, optional)

Returns groups from the character's account, if available, or the corporation and alliance.

Needs role: app-groups.<br>      *                  Returns only groups that have been added to the app as well.<br>      *                  It is not checked if character, corporation and alliance match.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**character** | **int32**| EVE character ID. | 
**corporation** | **int32**| EVE corporation ID. | 
 **optional** | ***GroupsWithFallbackV1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GroupsWithFallbackV1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alliance** | **optional.Int32**| EVE alliance ID. | 

### Return type

[**[]Group**](Group.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

