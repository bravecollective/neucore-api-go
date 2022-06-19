# \ApplicationTrackingApi

All URIs are relative to *https://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberTrackingV1**](ApplicationTrackingApi.md#MemberTrackingV1) | **Get** /app/v1/corporation/{id}/member-tracking | Return corporation member tracking data.



## MemberTrackingV1

> []CorporationMember MemberTrackingV1(ctx, id).Inactive(inactive).Active(active).Account(account).Execute()

Return corporation member tracking data.



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
    id := int32(56) // int32 | EVE corporation ID.
    inactive := int32(56) // int32 | Limit to members who have been inactive for x days or longer. (optional)
    active := int32(56) // int32 | Limit to members who were active in the last x days. (optional)
    account := "account_example" // string | Limit to members with (true) or without (false) an account. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationTrackingApi.MemberTrackingV1(context.Background(), id).Inactive(inactive).Active(active).Account(account).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationTrackingApi.MemberTrackingV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MemberTrackingV1`: []CorporationMember
    fmt.Fprintf(os.Stdout, "Response from `ApplicationTrackingApi.MemberTrackingV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | EVE corporation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMemberTrackingV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inactive** | **int32** | Limit to members who have been inactive for x days or longer. | 
 **active** | **int32** | Limit to members who were active in the last x days. | 
 **account** | **string** | Limit to members with (true) or without (false) an account. | 

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

