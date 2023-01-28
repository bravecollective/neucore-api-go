# Player

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccounts** | Pointer to [**[]ServiceAccount**](ServiceAccount.md) | External service accounts (API: not included by default) | [optional] 
**CharacterId** | Pointer to **int32** | ID of main character (API: not included by default) | [optional] 
**CorporationName** | Pointer to **string** | Corporation of main character (API: not included by default) | [optional] 
**AllianceName** | Pointer to **string** | Alliance of main character (API: not included by default) | [optional] 
**Id** | **NullableInt32** |  | 
**Name** | **string** | A name for the player.  This is the EVE character name of the current main character or of the last main character if there is currently none. | 
**Status** | Pointer to **string** | Player account status. | [optional] 
**Roles** | Pointer to [**[]Role**](Role.md) | Roles for authorization. | [optional] 
**Characters** | Pointer to [**[]Character**](Character.md) |  | [optional] 
**Groups** | Pointer to [**[]Group**](Group.md) | Group membership. | [optional] 
**ManagerGroups** | Pointer to [**[]Group**](Group.md) | Manager of groups. | [optional] 
**ManagerApps** | Pointer to [**[]App**](App.md) | Manager of apps. | [optional] 
**RemovedCharacters** | Pointer to [**[]RemovedCharacter**](RemovedCharacter.md) | Characters that were removed from a player (API: not included by default). | [optional] 
**IncomingCharacters** | Pointer to [**[]RemovedCharacter**](RemovedCharacter.md) | Characters that were moved from another player account to this account (API: not included by default). | [optional] 

## Methods

### NewPlayer

`func NewPlayer(id NullableInt32, name string, ) *Player`

NewPlayer instantiates a new Player object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayerWithDefaults

`func NewPlayerWithDefaults() *Player`

NewPlayerWithDefaults instantiates a new Player object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccounts

`func (o *Player) GetServiceAccounts() []ServiceAccount`

GetServiceAccounts returns the ServiceAccounts field if non-nil, zero value otherwise.

### GetServiceAccountsOk

`func (o *Player) GetServiceAccountsOk() (*[]ServiceAccount, bool)`

GetServiceAccountsOk returns a tuple with the ServiceAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccounts

`func (o *Player) SetServiceAccounts(v []ServiceAccount)`

SetServiceAccounts sets ServiceAccounts field to given value.

### HasServiceAccounts

`func (o *Player) HasServiceAccounts() bool`

HasServiceAccounts returns a boolean if a field has been set.

### GetCharacterId

`func (o *Player) GetCharacterId() int32`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *Player) GetCharacterIdOk() (*int32, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *Player) SetCharacterId(v int32)`

SetCharacterId sets CharacterId field to given value.

### HasCharacterId

`func (o *Player) HasCharacterId() bool`

HasCharacterId returns a boolean if a field has been set.

### GetCorporationName

`func (o *Player) GetCorporationName() string`

GetCorporationName returns the CorporationName field if non-nil, zero value otherwise.

### GetCorporationNameOk

`func (o *Player) GetCorporationNameOk() (*string, bool)`

GetCorporationNameOk returns a tuple with the CorporationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporationName

`func (o *Player) SetCorporationName(v string)`

SetCorporationName sets CorporationName field to given value.

### HasCorporationName

`func (o *Player) HasCorporationName() bool`

HasCorporationName returns a boolean if a field has been set.

### GetAllianceName

`func (o *Player) GetAllianceName() string`

GetAllianceName returns the AllianceName field if non-nil, zero value otherwise.

### GetAllianceNameOk

`func (o *Player) GetAllianceNameOk() (*string, bool)`

GetAllianceNameOk returns a tuple with the AllianceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllianceName

`func (o *Player) SetAllianceName(v string)`

SetAllianceName sets AllianceName field to given value.

### HasAllianceName

`func (o *Player) HasAllianceName() bool`

HasAllianceName returns a boolean if a field has been set.

### GetId

`func (o *Player) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Player) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Player) SetId(v int32)`

SetId sets Id field to given value.


### SetIdNil

`func (o *Player) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *Player) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *Player) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Player) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Player) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *Player) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Player) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Player) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Player) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRoles

`func (o *Player) GetRoles() []Role`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Player) GetRolesOk() (*[]Role, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Player) SetRoles(v []Role)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Player) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetCharacters

`func (o *Player) GetCharacters() []Character`

GetCharacters returns the Characters field if non-nil, zero value otherwise.

### GetCharactersOk

`func (o *Player) GetCharactersOk() (*[]Character, bool)`

GetCharactersOk returns a tuple with the Characters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacters

`func (o *Player) SetCharacters(v []Character)`

SetCharacters sets Characters field to given value.

### HasCharacters

`func (o *Player) HasCharacters() bool`

HasCharacters returns a boolean if a field has been set.

### GetGroups

`func (o *Player) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Player) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Player) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *Player) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetManagerGroups

`func (o *Player) GetManagerGroups() []Group`

GetManagerGroups returns the ManagerGroups field if non-nil, zero value otherwise.

### GetManagerGroupsOk

`func (o *Player) GetManagerGroupsOk() (*[]Group, bool)`

GetManagerGroupsOk returns a tuple with the ManagerGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerGroups

`func (o *Player) SetManagerGroups(v []Group)`

SetManagerGroups sets ManagerGroups field to given value.

### HasManagerGroups

`func (o *Player) HasManagerGroups() bool`

HasManagerGroups returns a boolean if a field has been set.

### GetManagerApps

`func (o *Player) GetManagerApps() []App`

GetManagerApps returns the ManagerApps field if non-nil, zero value otherwise.

### GetManagerAppsOk

`func (o *Player) GetManagerAppsOk() (*[]App, bool)`

GetManagerAppsOk returns a tuple with the ManagerApps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerApps

`func (o *Player) SetManagerApps(v []App)`

SetManagerApps sets ManagerApps field to given value.

### HasManagerApps

`func (o *Player) HasManagerApps() bool`

HasManagerApps returns a boolean if a field has been set.

### GetRemovedCharacters

`func (o *Player) GetRemovedCharacters() []RemovedCharacter`

GetRemovedCharacters returns the RemovedCharacters field if non-nil, zero value otherwise.

### GetRemovedCharactersOk

`func (o *Player) GetRemovedCharactersOk() (*[]RemovedCharacter, bool)`

GetRemovedCharactersOk returns a tuple with the RemovedCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedCharacters

`func (o *Player) SetRemovedCharacters(v []RemovedCharacter)`

SetRemovedCharacters sets RemovedCharacters field to given value.

### HasRemovedCharacters

`func (o *Player) HasRemovedCharacters() bool`

HasRemovedCharacters returns a boolean if a field has been set.

### GetIncomingCharacters

`func (o *Player) GetIncomingCharacters() []RemovedCharacter`

GetIncomingCharacters returns the IncomingCharacters field if non-nil, zero value otherwise.

### GetIncomingCharactersOk

`func (o *Player) GetIncomingCharactersOk() (*[]RemovedCharacter, bool)`

GetIncomingCharactersOk returns a tuple with the IncomingCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomingCharacters

`func (o *Player) SetIncomingCharacters(v []RemovedCharacter)`

SetIncomingCharacters sets IncomingCharacters field to given value.

### HasIncomingCharacters

`func (o *Player) HasIncomingCharacters() bool`

HasIncomingCharacters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


