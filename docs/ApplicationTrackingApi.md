# \ApplicationTrackingApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberTrackingV1**](ApplicationTrackingApi.md#MemberTrackingV1) | **Get** /app/v1/corporation/{id}/member-tracking | Return corporation member tracking data.



## MemberTrackingV1

> []CorporationMember MemberTrackingV1(ctx, id, optional)

Return corporation member tracking data.

Needs role: app-tracking

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32**| EVE corporation ID. | 
 **optional** | ***MemberTrackingV1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a MemberTrackingV1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inactive** | **optional.Int32**| Limit to members who have been inactive for x days or longer. | 
 **active** | **optional.Int32**| Limit to members who were active in the last x days. | 
 **account** | **optional.String**| Limit to members with (true) or without (false) an account. | 

### Return type

[**[]CorporationMember**](CorporationMember.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

