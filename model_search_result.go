/*
Neucore API

Client library of Neucore API

API version: 1.43.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
)

// SearchResult struct for SearchResult
type SearchResult struct {
	CharacterId int32 `json:"characterId"`
	CharacterName string `json:"characterName"`
	PlayerId int32 `json:"playerId"`
	PlayerName string `json:"playerName"`
}

// NewSearchResult instantiates a new SearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchResult(characterId int32, characterName string, playerId int32, playerName string) *SearchResult {
	this := SearchResult{}
	this.CharacterId = characterId
	this.CharacterName = characterName
	this.PlayerId = playerId
	this.PlayerName = playerName
	return &this
}

// NewSearchResultWithDefaults instantiates a new SearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchResultWithDefaults() *SearchResult {
	this := SearchResult{}
	return &this
}

// GetCharacterId returns the CharacterId field value
func (o *SearchResult) GetCharacterId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CharacterId
}

// GetCharacterIdOk returns a tuple with the CharacterId field value
// and a boolean to check if the value has been set.
func (o *SearchResult) GetCharacterIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CharacterId, true
}

// SetCharacterId sets field value
func (o *SearchResult) SetCharacterId(v int32) {
	o.CharacterId = v
}

// GetCharacterName returns the CharacterName field value
func (o *SearchResult) GetCharacterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CharacterName
}

// GetCharacterNameOk returns a tuple with the CharacterName field value
// and a boolean to check if the value has been set.
func (o *SearchResult) GetCharacterNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.CharacterName, true
}

// SetCharacterName sets field value
func (o *SearchResult) SetCharacterName(v string) {
	o.CharacterName = v
}

// GetPlayerId returns the PlayerId field value
func (o *SearchResult) GetPlayerId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PlayerId
}

// GetPlayerIdOk returns a tuple with the PlayerId field value
// and a boolean to check if the value has been set.
func (o *SearchResult) GetPlayerIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlayerId, true
}

// SetPlayerId sets field value
func (o *SearchResult) SetPlayerId(v int32) {
	o.PlayerId = v
}

// GetPlayerName returns the PlayerName field value
func (o *SearchResult) GetPlayerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlayerName
}

// GetPlayerNameOk returns a tuple with the PlayerName field value
// and a boolean to check if the value has been set.
func (o *SearchResult) GetPlayerNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PlayerName, true
}

// SetPlayerName sets field value
func (o *SearchResult) SetPlayerName(v string) {
	o.PlayerName = v
}

func (o SearchResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["characterId"] = o.CharacterId
	}
	if true {
		toSerialize["characterName"] = o.CharacterName
	}
	if true {
		toSerialize["playerId"] = o.PlayerId
	}
	if true {
		toSerialize["playerName"] = o.PlayerName
	}
	return json.Marshal(toSerialize)
}

type NullableSearchResult struct {
	value *SearchResult
	isSet bool
}

func (v NullableSearchResult) Get() *SearchResult {
	return v.value
}

func (v *NullableSearchResult) Set(val *SearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchResult(val *SearchResult) *NullableSearchResult {
	return &NullableSearchResult{value: val, isSet: true}
}

func (v NullableSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


