/*
Neucore API

Client library of Neucore API

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
	"time"
)

// EsiToken struct for EsiToken
type EsiToken struct {
	// ID of EveLogin
	EveLoginId int32 `json:"eveLoginId"`
	// ID of Character
	CharacterId int32 `json:"characterId"`
	// ID of Player
	PlayerId int32 `json:"playerId"`
	// Name of Player
	PlayerName *string `json:"playerName,omitempty"`
	Character *Character `json:"character,omitempty"`
	// Shows if the refresh token is valid or not.  This is null if there is no refresh token (EVE SSOv1 only) or a valid token but without scopes (SSOv2).
	ValidToken NullableBool `json:"validToken"`
	// Date and time when the valid token property was last changed.
	ValidTokenTime NullableTime `json:"validTokenTime"`
	// Shows if the EVE character has all required roles for the login.  Null if the login does not require any roles or if the token is invalid.
	HasRoles NullableBool `json:"hasRoles"`
	// When the refresh token was last checked for validity.
	LastChecked NullableTime `json:"lastChecked,omitempty"`
}

// NewEsiToken instantiates a new EsiToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsiToken(eveLoginId int32, characterId int32, playerId int32, validToken NullableBool, validTokenTime NullableTime, hasRoles NullableBool) *EsiToken {
	this := EsiToken{}
	this.EveLoginId = eveLoginId
	this.CharacterId = characterId
	this.PlayerId = playerId
	this.ValidToken = validToken
	this.ValidTokenTime = validTokenTime
	this.HasRoles = hasRoles
	return &this
}

// NewEsiTokenWithDefaults instantiates a new EsiToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsiTokenWithDefaults() *EsiToken {
	this := EsiToken{}
	return &this
}

// GetEveLoginId returns the EveLoginId field value
func (o *EsiToken) GetEveLoginId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EveLoginId
}

// GetEveLoginIdOk returns a tuple with the EveLoginId field value
// and a boolean to check if the value has been set.
func (o *EsiToken) GetEveLoginIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.EveLoginId, true
}

// SetEveLoginId sets field value
func (o *EsiToken) SetEveLoginId(v int32) {
	o.EveLoginId = v
}

// GetCharacterId returns the CharacterId field value
func (o *EsiToken) GetCharacterId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CharacterId
}

// GetCharacterIdOk returns a tuple with the CharacterId field value
// and a boolean to check if the value has been set.
func (o *EsiToken) GetCharacterIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CharacterId, true
}

// SetCharacterId sets field value
func (o *EsiToken) SetCharacterId(v int32) {
	o.CharacterId = v
}

// GetPlayerId returns the PlayerId field value
func (o *EsiToken) GetPlayerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlayerId
}

// GetPlayerIdOk returns a tuple with the PlayerId field value
// and a boolean to check if the value has been set.
func (o *EsiToken) GetPlayerIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlayerId, true
}

// SetPlayerId sets field value
func (o *EsiToken) SetPlayerId(v int32) {
	o.PlayerId = v
}

// GetPlayerName returns the PlayerName field value if set, zero value otherwise.
func (o *EsiToken) GetPlayerName() string {
	if o == nil || isNil(o.PlayerName) {
		var ret string
		return ret
	}
	return *o.PlayerName
}

// GetPlayerNameOk returns a tuple with the PlayerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsiToken) GetPlayerNameOk() (*string, bool) {
	if o == nil || isNil(o.PlayerName) {
    return nil, false
	}
	return o.PlayerName, true
}

// HasPlayerName returns a boolean if a field has been set.
func (o *EsiToken) HasPlayerName() bool {
	if o != nil && !isNil(o.PlayerName) {
		return true
	}

	return false
}

// SetPlayerName gets a reference to the given string and assigns it to the PlayerName field.
func (o *EsiToken) SetPlayerName(v string) {
	o.PlayerName = &v
}

// GetCharacter returns the Character field value if set, zero value otherwise.
func (o *EsiToken) GetCharacter() Character {
	if o == nil || isNil(o.Character) {
		var ret Character
		return ret
	}
	return *o.Character
}

// GetCharacterOk returns a tuple with the Character field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EsiToken) GetCharacterOk() (*Character, bool) {
	if o == nil || isNil(o.Character) {
    return nil, false
	}
	return o.Character, true
}

// HasCharacter returns a boolean if a field has been set.
func (o *EsiToken) HasCharacter() bool {
	if o != nil && !isNil(o.Character) {
		return true
	}

	return false
}

// SetCharacter gets a reference to the given Character and assigns it to the Character field.
func (o *EsiToken) SetCharacter(v Character) {
	o.Character = &v
}

// GetValidToken returns the ValidToken field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *EsiToken) GetValidToken() bool {
	if o == nil || o.ValidToken.Get() == nil {
		var ret bool
		return ret
	}

	return *o.ValidToken.Get()
}

// GetValidTokenOk returns a tuple with the ValidToken field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsiToken) GetValidTokenOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.ValidToken.Get(), o.ValidToken.IsSet()
}

// SetValidToken sets field value
func (o *EsiToken) SetValidToken(v bool) {
	o.ValidToken.Set(&v)
}

// GetValidTokenTime returns the ValidTokenTime field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *EsiToken) GetValidTokenTime() time.Time {
	if o == nil || o.ValidTokenTime.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ValidTokenTime.Get()
}

// GetValidTokenTimeOk returns a tuple with the ValidTokenTime field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsiToken) GetValidTokenTimeOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.ValidTokenTime.Get(), o.ValidTokenTime.IsSet()
}

// SetValidTokenTime sets field value
func (o *EsiToken) SetValidTokenTime(v time.Time) {
	o.ValidTokenTime.Set(&v)
}

// GetHasRoles returns the HasRoles field value
// If the value is explicit nil, the zero value for bool will be returned
func (o *EsiToken) GetHasRoles() bool {
	if o == nil || o.HasRoles.Get() == nil {
		var ret bool
		return ret
	}

	return *o.HasRoles.Get()
}

// GetHasRolesOk returns a tuple with the HasRoles field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsiToken) GetHasRolesOk() (*bool, bool) {
	if o == nil {
    return nil, false
	}
	return o.HasRoles.Get(), o.HasRoles.IsSet()
}

// SetHasRoles sets field value
func (o *EsiToken) SetHasRoles(v bool) {
	o.HasRoles.Set(&v)
}

// GetLastChecked returns the LastChecked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EsiToken) GetLastChecked() time.Time {
	if o == nil || isNil(o.LastChecked.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastChecked.Get()
}

// GetLastCheckedOk returns a tuple with the LastChecked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EsiToken) GetLastCheckedOk() (*time.Time, bool) {
	if o == nil {
    return nil, false
	}
	return o.LastChecked.Get(), o.LastChecked.IsSet()
}

// HasLastChecked returns a boolean if a field has been set.
func (o *EsiToken) HasLastChecked() bool {
	if o != nil && o.LastChecked.IsSet() {
		return true
	}

	return false
}

// SetLastChecked gets a reference to the given NullableTime and assigns it to the LastChecked field.
func (o *EsiToken) SetLastChecked(v time.Time) {
	o.LastChecked.Set(&v)
}
// SetLastCheckedNil sets the value for LastChecked to be an explicit nil
func (o *EsiToken) SetLastCheckedNil() {
	o.LastChecked.Set(nil)
}

// UnsetLastChecked ensures that no value is present for LastChecked, not even an explicit nil
func (o *EsiToken) UnsetLastChecked() {
	o.LastChecked.Unset()
}

func (o EsiToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["eveLoginId"] = o.EveLoginId
	}
	if true {
		toSerialize["characterId"] = o.CharacterId
	}
	if true {
		toSerialize["playerId"] = o.PlayerId
	}
	if !isNil(o.PlayerName) {
		toSerialize["playerName"] = o.PlayerName
	}
	if !isNil(o.Character) {
		toSerialize["character"] = o.Character
	}
	if true {
		toSerialize["validToken"] = o.ValidToken.Get()
	}
	if true {
		toSerialize["validTokenTime"] = o.ValidTokenTime.Get()
	}
	if true {
		toSerialize["hasRoles"] = o.HasRoles.Get()
	}
	if o.LastChecked.IsSet() {
		toSerialize["lastChecked"] = o.LastChecked.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEsiToken struct {
	value *EsiToken
	isSet bool
}

func (v NullableEsiToken) Get() *EsiToken {
	return v.value
}

func (v *NullableEsiToken) Set(val *EsiToken) {
	v.value = val
	v.isSet = true
}

func (v NullableEsiToken) IsSet() bool {
	return v.isSet
}

func (v *NullableEsiToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsiToken(val *EsiToken) *NullableEsiToken {
	return &NullableEsiToken{value: val, isSet: true}
}

func (v NullableEsiToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsiToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


