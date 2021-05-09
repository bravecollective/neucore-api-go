/*
 * Neucore API
 *
 * Client library of Neucore API
 *
 * API version: 1.21.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
	"time"
)

// RemovedCharacter struct for RemovedCharacter
type RemovedCharacter struct {
	NewPlayerId *int32 `json:"newPlayerId,omitempty"`
	NewPlayerName *string `json:"newPlayerName,omitempty"`
	Player *Player `json:"player,omitempty"`
	// EVE character ID.
	CharacterId int64 `json:"characterId"`
	// EVE character name.
	CharacterName string `json:"characterName"`
	// Date of removal.
	RemovedDate time.Time `json:"removedDate"`
	// How it was removed (deleted or moved to another account).
	Reason string `json:"reason"`
	DeletedBy *Player `json:"deletedBy,omitempty"`
}

// NewRemovedCharacter instantiates a new RemovedCharacter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemovedCharacter(characterId int64, characterName string, removedDate time.Time, reason string) *RemovedCharacter {
	this := RemovedCharacter{}
	this.CharacterId = characterId
	this.CharacterName = characterName
	this.RemovedDate = removedDate
	this.Reason = reason
	return &this
}

// NewRemovedCharacterWithDefaults instantiates a new RemovedCharacter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemovedCharacterWithDefaults() *RemovedCharacter {
	this := RemovedCharacter{}
	return &this
}

// GetNewPlayerId returns the NewPlayerId field value if set, zero value otherwise.
func (o *RemovedCharacter) GetNewPlayerId() int32 {
	if o == nil || o.NewPlayerId == nil {
		var ret int32
		return ret
	}
	return *o.NewPlayerId
}

// GetNewPlayerIdOk returns a tuple with the NewPlayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemovedCharacter) GetNewPlayerIdOk() (*int32, bool) {
	if o == nil || o.NewPlayerId == nil {
		return nil, false
	}
	return o.NewPlayerId, true
}

// HasNewPlayerId returns a boolean if a field has been set.
func (o *RemovedCharacter) HasNewPlayerId() bool {
	if o != nil && o.NewPlayerId != nil {
		return true
	}

	return false
}

// SetNewPlayerId gets a reference to the given int32 and assigns it to the NewPlayerId field.
func (o *RemovedCharacter) SetNewPlayerId(v int32) {
	o.NewPlayerId = &v
}

// GetNewPlayerName returns the NewPlayerName field value if set, zero value otherwise.
func (o *RemovedCharacter) GetNewPlayerName() string {
	if o == nil || o.NewPlayerName == nil {
		var ret string
		return ret
	}
	return *o.NewPlayerName
}

// GetNewPlayerNameOk returns a tuple with the NewPlayerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemovedCharacter) GetNewPlayerNameOk() (*string, bool) {
	if o == nil || o.NewPlayerName == nil {
		return nil, false
	}
	return o.NewPlayerName, true
}

// HasNewPlayerName returns a boolean if a field has been set.
func (o *RemovedCharacter) HasNewPlayerName() bool {
	if o != nil && o.NewPlayerName != nil {
		return true
	}

	return false
}

// SetNewPlayerName gets a reference to the given string and assigns it to the NewPlayerName field.
func (o *RemovedCharacter) SetNewPlayerName(v string) {
	o.NewPlayerName = &v
}

// GetPlayer returns the Player field value if set, zero value otherwise.
func (o *RemovedCharacter) GetPlayer() Player {
	if o == nil || o.Player == nil {
		var ret Player
		return ret
	}
	return *o.Player
}

// GetPlayerOk returns a tuple with the Player field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemovedCharacter) GetPlayerOk() (*Player, bool) {
	if o == nil || o.Player == nil {
		return nil, false
	}
	return o.Player, true
}

// HasPlayer returns a boolean if a field has been set.
func (o *RemovedCharacter) HasPlayer() bool {
	if o != nil && o.Player != nil {
		return true
	}

	return false
}

// SetPlayer gets a reference to the given Player and assigns it to the Player field.
func (o *RemovedCharacter) SetPlayer(v Player) {
	o.Player = &v
}

// GetCharacterId returns the CharacterId field value
func (o *RemovedCharacter) GetCharacterId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CharacterId
}

// GetCharacterIdOk returns a tuple with the CharacterId field value
// and a boolean to check if the value has been set.
func (o *RemovedCharacter) GetCharacterIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CharacterId, true
}

// SetCharacterId sets field value
func (o *RemovedCharacter) SetCharacterId(v int64) {
	o.CharacterId = v
}

// GetCharacterName returns the CharacterName field value
func (o *RemovedCharacter) GetCharacterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CharacterName
}

// GetCharacterNameOk returns a tuple with the CharacterName field value
// and a boolean to check if the value has been set.
func (o *RemovedCharacter) GetCharacterNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CharacterName, true
}

// SetCharacterName sets field value
func (o *RemovedCharacter) SetCharacterName(v string) {
	o.CharacterName = v
}

// GetRemovedDate returns the RemovedDate field value
func (o *RemovedCharacter) GetRemovedDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.RemovedDate
}

// GetRemovedDateOk returns a tuple with the RemovedDate field value
// and a boolean to check if the value has been set.
func (o *RemovedCharacter) GetRemovedDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RemovedDate, true
}

// SetRemovedDate sets field value
func (o *RemovedCharacter) SetRemovedDate(v time.Time) {
	o.RemovedDate = v
}

// GetReason returns the Reason field value
func (o *RemovedCharacter) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *RemovedCharacter) GetReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *RemovedCharacter) SetReason(v string) {
	o.Reason = v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise.
func (o *RemovedCharacter) GetDeletedBy() Player {
	if o == nil || o.DeletedBy == nil {
		var ret Player
		return ret
	}
	return *o.DeletedBy
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemovedCharacter) GetDeletedByOk() (*Player, bool) {
	if o == nil || o.DeletedBy == nil {
		return nil, false
	}
	return o.DeletedBy, true
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *RemovedCharacter) HasDeletedBy() bool {
	if o != nil && o.DeletedBy != nil {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given Player and assigns it to the DeletedBy field.
func (o *RemovedCharacter) SetDeletedBy(v Player) {
	o.DeletedBy = &v
}

func (o RemovedCharacter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NewPlayerId != nil {
		toSerialize["newPlayerId"] = o.NewPlayerId
	}
	if o.NewPlayerName != nil {
		toSerialize["newPlayerName"] = o.NewPlayerName
	}
	if o.Player != nil {
		toSerialize["player"] = o.Player
	}
	if true {
		toSerialize["characterId"] = o.CharacterId
	}
	if true {
		toSerialize["characterName"] = o.CharacterName
	}
	if true {
		toSerialize["removedDate"] = o.RemovedDate
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	if o.DeletedBy != nil {
		toSerialize["deletedBy"] = o.DeletedBy
	}
	return json.Marshal(toSerialize)
}

type NullableRemovedCharacter struct {
	value *RemovedCharacter
	isSet bool
}

func (v NullableRemovedCharacter) Get() *RemovedCharacter {
	return v.value
}

func (v *NullableRemovedCharacter) Set(val *RemovedCharacter) {
	v.value = val
	v.isSet = true
}

func (v NullableRemovedCharacter) IsSet() bool {
	return v.isSet
}

func (v *NullableRemovedCharacter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemovedCharacter(val *RemovedCharacter) *NullableRemovedCharacter {
	return &NullableRemovedCharacter{value: val, isSet: true}
}

func (v NullableRemovedCharacter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemovedCharacter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


