# CorporationMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Player** | [**Player**](Player.md) |  | [optional] 
**Id** | **int64** | EVE Character ID. | 
**Name** | Pointer to **string** | EVE Character name. | 
**Location** | [**EsiLocation**](EsiLocation.md) |  | [optional] 
**LogoffDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**LogonDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**ShipType** | [**EsiType**](EsiType.md) |  | [optional] 
**StartDate** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Character** | [**Character**](Character.md) |  | [optional] 
**MissingCharacterMailSentDate** | Pointer to [**time.Time**](time.Time.md) | Date and time of the last sent mail. | [optional] 
**MissingCharacterMailSentResult** | Pointer to **string** | Result of the last sent mail (OK, Blocked, CSPA charge &gt; 0) | [optional] 
**MissingCharacterMailSentNumber** | **int32** | Number of mails sent, is reset when the character is added. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


