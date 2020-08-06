# Player

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** | A name for the player.  This is the EVE character name of the current main character or of the last main character if there is currently none. | 
**Status** | **string** | Player account status. | [optional] 
**Roles** | [**[]Role**](Role.md) | Roles for authorization. | [optional] 
**Characters** | [**[]Character**](Character.md) |  | [optional] 
**Groups** | [**[]Group**](Group.md) | Group membership. | [optional] 
**ManagerGroups** | [**[]Group**](Group.md) | Manager of groups. | [optional] 
**ManagerApps** | [**[]App**](App.md) | Manager of apps. | [optional] 
**RemovedCharacters** | [**[]RemovedCharacter**](RemovedCharacter.md) | Characters that were removed from a player (API: not included by default). | [optional] 
**IncomingCharacters** | [**[]RemovedCharacter**](RemovedCharacter.md) | Characters that were moved from another player account to this account (API: not included by default). | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


