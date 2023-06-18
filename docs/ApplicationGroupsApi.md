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
[**GroupMembersV1**](ApplicationGroupsApi.md#GroupMembersV1) | **Get** /app/v1/group-members/{groupId} | Returns the main character IDs from all group members.
[**GroupsBulkV1**](ApplicationGroupsApi.md#GroupsBulkV1) | **Post** /app/v1/groups | Return groups of multiple players, identified by one of their character IDs.
[**GroupsV1**](ApplicationGroupsApi.md#GroupsV1) | **Get** /app/v1/groups/{cid} | Return groups of the character&#39;s player account.
[**GroupsV2**](ApplicationGroupsApi.md#GroupsV2) | **Get** /app/v2/groups/{cid} | Return groups of the character&#39;s player account.
[**GroupsWithFallbackV1**](ApplicationGroupsApi.md#GroupsWithFallbackV1) | **Get** /app/v1/groups-with-fallback | Returns groups from the character&#39;s account, if available, or the corporation and alliance.



## AllianceGroupsBulkV1

> []Alliance AllianceGroupsBulkV1(ctx).RequestBody(requestBody).Execute()

Return groups of multiple alliances.



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
    requestBody := []int32{int32(123)} // []int32 | EVE alliance IDs array.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsApi.AllianceGroupsBulkV1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsApi.AllianceGroupsBulkV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllianceGroupsBulkV1`: []Alliance
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsApi.AllianceGroupsBulkV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllianceGroupsBulkV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]int32** | EVE alliance IDs array. | 

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

> []Group AllianceGroupsV1(ctx, aid).Execute()

Return groups of the alliance.



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
    aid := int32(56) // int32 | EVE alliance ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsApi.AllianceGroupsV1(context.Background(), aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsApi.AllianceGroupsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllianceGroupsV1`: []Group
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsApi.AllianceGroupsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aid** | **int32** | EVE alliance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllianceGroupsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []Group AllianceGroupsV2(ctx, aid).Execute()

Return groups of the alliance.



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
    aid := int32(56) // int32 | EVE alliance ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsApi.AllianceGroupsV2(context.Background(), aid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsApi.AllianceGroupsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllianceGroupsV2`: []Group
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsApi.AllianceGroupsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**aid** | **int32** | EVE alliance ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAllianceGroupsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []Corporation CorpGroupsBulkV1(ctx).RequestBody(requestBody).Execute()

Return groups of multiple corporations.



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
    requestBody := []int32{int32(123)} // []int32 | EVE corporation IDs array.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsApi.CorpGroupsBulkV1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsApi.CorpGroupsBulkV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorpGroupsBulkV1`: []Corporation
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsApi.CorpGroupsBulkV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorpGroupsBulkV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]int32** | EVE corporation IDs array. | 

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

> []Group CorpGroupsV1(ctx, cid).Execute()

Return groups of the corporation.



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
    cid := int32(56) // int32 | EVE corporation ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsApi.CorpGroupsV1(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsApi.CorpGroupsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorpGroupsV1`: []Group
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsApi.CorpGroupsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int32** | EVE corporation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCorpGroupsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []Group CorpGroupsV2(ctx, cid).Execute()

Return groups of the corporation.



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
    cid := int32(56) // int32 | EVE corporation ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsApi.CorpGroupsV2(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsApi.CorpGroupsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorpGroupsV2`: []Group
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsApi.CorpGroupsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int32** | EVE corporation ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCorpGroupsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GroupMembersV1

> []int32 GroupMembersV1(ctx, groupId).Corporation(corporation).Execute()

Returns the main character IDs from all group members.



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
    groupId := int32(56) // int32 | Group ID.
    corporation := int32(56) // int32 | Limit to characters that are a member of this corporation. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsApi.GroupMembersV1(context.Background(), groupId).Corporation(corporation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsApi.GroupMembersV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupMembersV1`: []int32
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsApi.GroupMembersV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **int32** | Group ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupMembersV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **corporation** | **int32** | Limit to characters that are a member of this corporation. | 

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


## GroupsBulkV1

> []CharacterGroups GroupsBulkV1(ctx).RequestBody(requestBody).Execute()

Return groups of multiple players, identified by one of their character IDs.



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
    requestBody := []int32{int32(123)} // []int32 | EVE character IDs array.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsApi.GroupsBulkV1(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsApi.GroupsBulkV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsBulkV1`: []CharacterGroups
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsApi.GroupsBulkV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsBulkV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **[]int32** | EVE character IDs array. | 

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

> []Group GroupsV1(ctx, cid).Execute()

Return groups of the character's player account.



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
    cid := int32(56) // int32 | EVE character ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsApi.GroupsV1(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsApi.GroupsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsV1`: []Group
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsApi.GroupsV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int32** | EVE character ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []Group GroupsV2(ctx, cid).Execute()

Return groups of the character's player account.



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
    cid := int32(56) // int32 | EVE character ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsApi.GroupsV2(context.Background(), cid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsApi.GroupsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsV2`: []Group
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsApi.GroupsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cid** | **int32** | EVE character ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []Group GroupsWithFallbackV1(ctx).Character(character).Corporation(corporation).Alliance(alliance).Execute()

Returns groups from the character's account, if available, or the corporation and alliance.



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
    character := int32(56) // int32 | EVE character ID.
    corporation := int32(56) // int32 | EVE corporation ID.
    alliance := int32(56) // int32 | EVE alliance ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ApplicationGroupsApi.GroupsWithFallbackV1(context.Background()).Character(character).Corporation(corporation).Alliance(alliance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApplicationGroupsApi.GroupsWithFallbackV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsWithFallbackV1`: []Group
    fmt.Fprintf(os.Stdout, "Response from `ApplicationGroupsApi.GroupsWithFallbackV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGroupsWithFallbackV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **character** | **int32** | EVE character ID. | 
 **corporation** | **int32** | EVE corporation ID. | 
 **alliance** | **int32** | EVE alliance ID. | 

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

