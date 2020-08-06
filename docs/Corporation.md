# Corporation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | EVE corporation ID. | 
**Name** | Pointer to **string** | EVE corporation name. | 
**Ticker** | Pointer to **string** | Corporation ticker. | 
**Alliance** | [**Alliance**](Alliance.md) |  | [optional] 
**Groups** | [**[]Group**](Group.md) | Groups for automatic assignment (API: not included by default). | [optional] 
**TrackingLastUpdate** | Pointer to [**time.Time**](time.Time.md) | Last update of corporation member tracking data (API: not included by default). | [optional] 
**AutoAllowlist** | **bool** | True if this corporation was automatically placed on the allowlist of a watchlist (API: not included by default). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


