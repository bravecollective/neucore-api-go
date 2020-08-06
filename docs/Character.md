# Character

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | EVE character ID. | 
**Name** | **string** | EVE character name. | 
**Main** | **bool** |  | [optional] 
**ValidToken** | Pointer to **bool** | Shows if character&#39;s refresh token is valid or not.  This is null if there is no refresh token (EVE SSOv1 only) or a valid token but without scopes (SSOv2). | [optional] 
**ValidTokenTime** | Pointer to [**time.Time**](time.Time.md) | Date and time when that valid token property was last changed. | [optional] 
**Created** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LastUpdate** | Pointer to [**time.Time**](time.Time.md) | Last ESI update. | [optional] 
**Corporation** | [**Corporation**](Corporation.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


