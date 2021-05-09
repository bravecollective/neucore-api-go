# RemovedCharacter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPlayerId** | Pointer to **int32** |  | [optional] 
**NewPlayerName** | Pointer to **string** |  | [optional] 
**Player** | Pointer to [**Player**](Player.md) |  | [optional] 
**CharacterId** | **int64** | EVE character ID. | 
**CharacterName** | **string** | EVE character name. | 
**RemovedDate** | **time.Time** | Date of removal. | 
**Reason** | **string** | How it was removed (deleted or moved to another account). | 
**DeletedBy** | Pointer to [**Player**](Player.md) |  | [optional] 

## Methods

### NewRemovedCharacter

`func NewRemovedCharacter(characterId int64, characterName string, removedDate time.Time, reason string, ) *RemovedCharacter`

NewRemovedCharacter instantiates a new RemovedCharacter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemovedCharacterWithDefaults

`func NewRemovedCharacterWithDefaults() *RemovedCharacter`

NewRemovedCharacterWithDefaults instantiates a new RemovedCharacter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPlayerId

`func (o *RemovedCharacter) GetNewPlayerId() int32`

GetNewPlayerId returns the NewPlayerId field if non-nil, zero value otherwise.

### GetNewPlayerIdOk

`func (o *RemovedCharacter) GetNewPlayerIdOk() (*int32, bool)`

GetNewPlayerIdOk returns a tuple with the NewPlayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlayerId

`func (o *RemovedCharacter) SetNewPlayerId(v int32)`

SetNewPlayerId sets NewPlayerId field to given value.

### HasNewPlayerId

`func (o *RemovedCharacter) HasNewPlayerId() bool`

HasNewPlayerId returns a boolean if a field has been set.

### GetNewPlayerName

`func (o *RemovedCharacter) GetNewPlayerName() string`

GetNewPlayerName returns the NewPlayerName field if non-nil, zero value otherwise.

### GetNewPlayerNameOk

`func (o *RemovedCharacter) GetNewPlayerNameOk() (*string, bool)`

GetNewPlayerNameOk returns a tuple with the NewPlayerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPlayerName

`func (o *RemovedCharacter) SetNewPlayerName(v string)`

SetNewPlayerName sets NewPlayerName field to given value.

### HasNewPlayerName

`func (o *RemovedCharacter) HasNewPlayerName() bool`

HasNewPlayerName returns a boolean if a field has been set.

### GetPlayer

`func (o *RemovedCharacter) GetPlayer() Player`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *RemovedCharacter) GetPlayerOk() (*Player, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *RemovedCharacter) SetPlayer(v Player)`

SetPlayer sets Player field to given value.

### HasPlayer

`func (o *RemovedCharacter) HasPlayer() bool`

HasPlayer returns a boolean if a field has been set.

### GetCharacterId

`func (o *RemovedCharacter) GetCharacterId() int64`

GetCharacterId returns the CharacterId field if non-nil, zero value otherwise.

### GetCharacterIdOk

`func (o *RemovedCharacter) GetCharacterIdOk() (*int64, bool)`

GetCharacterIdOk returns a tuple with the CharacterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterId

`func (o *RemovedCharacter) SetCharacterId(v int64)`

SetCharacterId sets CharacterId field to given value.


### GetCharacterName

`func (o *RemovedCharacter) GetCharacterName() string`

GetCharacterName returns the CharacterName field if non-nil, zero value otherwise.

### GetCharacterNameOk

`func (o *RemovedCharacter) GetCharacterNameOk() (*string, bool)`

GetCharacterNameOk returns a tuple with the CharacterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacterName

`func (o *RemovedCharacter) SetCharacterName(v string)`

SetCharacterName sets CharacterName field to given value.


### GetRemovedDate

`func (o *RemovedCharacter) GetRemovedDate() time.Time`

GetRemovedDate returns the RemovedDate field if non-nil, zero value otherwise.

### GetRemovedDateOk

`func (o *RemovedCharacter) GetRemovedDateOk() (*time.Time, bool)`

GetRemovedDateOk returns a tuple with the RemovedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedDate

`func (o *RemovedCharacter) SetRemovedDate(v time.Time)`

SetRemovedDate sets RemovedDate field to given value.


### GetReason

`func (o *RemovedCharacter) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RemovedCharacter) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RemovedCharacter) SetReason(v string)`

SetReason sets Reason field to given value.


### GetDeletedBy

`func (o *RemovedCharacter) GetDeletedBy() Player`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *RemovedCharacter) GetDeletedByOk() (*Player, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *RemovedCharacter) SetDeletedBy(v Player)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *RemovedCharacter) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


