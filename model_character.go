/*
 * Neucore API
 *
 * Client library of Neucore API
 *
 * API version: 1.22.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
	"time"
)

// Character struct for Character
type Character struct {
	// EVE character ID.
	Id int64 `json:"id"`
	// EVE character name.
	Name string `json:"name"`
	Main *bool `json:"main,omitempty"`
	// Shows if character's refresh token is valid or not.  This is null if there is no refresh token (EVE SSOv1 only) or a valid token but without scopes (SSOv2).
	ValidToken NullableBool `json:"validToken,omitempty"`
	// Date and time when that valid token property was last changed.
	ValidTokenTime NullableTime `json:"validTokenTime,omitempty"`
	Created NullableTime `json:"created,omitempty"`
	// Last ESI update.
	LastUpdate NullableTime `json:"lastUpdate,omitempty"`
	Corporation *Corporation `json:"corporation,omitempty"`
	// List of previous character names (API: not included by default).
	CharacterNameChanges *[]CharacterNameChange `json:"characterNameChanges,omitempty"`
}

// NewCharacter instantiates a new Character object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharacter(id int64, name string) *Character {
	this := Character{}
	this.Id = id
	this.Name = name
	return &this
}

// NewCharacterWithDefaults instantiates a new Character object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharacterWithDefaults() *Character {
	this := Character{}
	return &this
}

// GetId returns the Id field value
func (o *Character) GetId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Character) GetIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Character) SetId(v int64) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Character) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Character) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Character) SetName(v string) {
	o.Name = v
}

// GetMain returns the Main field value if set, zero value otherwise.
func (o *Character) GetMain() bool {
	if o == nil || o.Main == nil {
		var ret bool
		return ret
	}
	return *o.Main
}

// GetMainOk returns a tuple with the Main field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Character) GetMainOk() (*bool, bool) {
	if o == nil || o.Main == nil {
		return nil, false
	}
	return o.Main, true
}

// HasMain returns a boolean if a field has been set.
func (o *Character) HasMain() bool {
	if o != nil && o.Main != nil {
		return true
	}

	return false
}

// SetMain gets a reference to the given bool and assigns it to the Main field.
func (o *Character) SetMain(v bool) {
	o.Main = &v
}

// GetValidToken returns the ValidToken field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Character) GetValidToken() bool {
	if o == nil || o.ValidToken.Get() == nil {
		var ret bool
		return ret
	}
	return *o.ValidToken.Get()
}

// GetValidTokenOk returns a tuple with the ValidToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Character) GetValidTokenOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ValidToken.Get(), o.ValidToken.IsSet()
}

// HasValidToken returns a boolean if a field has been set.
func (o *Character) HasValidToken() bool {
	if o != nil && o.ValidToken.IsSet() {
		return true
	}

	return false
}

// SetValidToken gets a reference to the given NullableBool and assigns it to the ValidToken field.
func (o *Character) SetValidToken(v bool) {
	o.ValidToken.Set(&v)
}
// SetValidTokenNil sets the value for ValidToken to be an explicit nil
func (o *Character) SetValidTokenNil() {
	o.ValidToken.Set(nil)
}

// UnsetValidToken ensures that no value is present for ValidToken, not even an explicit nil
func (o *Character) UnsetValidToken() {
	o.ValidToken.Unset()
}

// GetValidTokenTime returns the ValidTokenTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Character) GetValidTokenTime() time.Time {
	if o == nil || o.ValidTokenTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ValidTokenTime.Get()
}

// GetValidTokenTimeOk returns a tuple with the ValidTokenTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Character) GetValidTokenTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ValidTokenTime.Get(), o.ValidTokenTime.IsSet()
}

// HasValidTokenTime returns a boolean if a field has been set.
func (o *Character) HasValidTokenTime() bool {
	if o != nil && o.ValidTokenTime.IsSet() {
		return true
	}

	return false
}

// SetValidTokenTime gets a reference to the given NullableTime and assigns it to the ValidTokenTime field.
func (o *Character) SetValidTokenTime(v time.Time) {
	o.ValidTokenTime.Set(&v)
}
// SetValidTokenTimeNil sets the value for ValidTokenTime to be an explicit nil
func (o *Character) SetValidTokenTimeNil() {
	o.ValidTokenTime.Set(nil)
}

// UnsetValidTokenTime ensures that no value is present for ValidTokenTime, not even an explicit nil
func (o *Character) UnsetValidTokenTime() {
	o.ValidTokenTime.Unset()
}

// GetCreated returns the Created field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Character) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Character) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// HasCreated returns a boolean if a field has been set.
func (o *Character) HasCreated() bool {
	if o != nil && o.Created.IsSet() {
		return true
	}

	return false
}

// SetCreated gets a reference to the given NullableTime and assigns it to the Created field.
func (o *Character) SetCreated(v time.Time) {
	o.Created.Set(&v)
}
// SetCreatedNil sets the value for Created to be an explicit nil
func (o *Character) SetCreatedNil() {
	o.Created.Set(nil)
}

// UnsetCreated ensures that no value is present for Created, not even an explicit nil
func (o *Character) UnsetCreated() {
	o.Created.Unset()
}

// GetLastUpdate returns the LastUpdate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Character) GetLastUpdate() time.Time {
	if o == nil || o.LastUpdate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdate.Get()
}

// GetLastUpdateOk returns a tuple with the LastUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Character) GetLastUpdateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUpdate.Get(), o.LastUpdate.IsSet()
}

// HasLastUpdate returns a boolean if a field has been set.
func (o *Character) HasLastUpdate() bool {
	if o != nil && o.LastUpdate.IsSet() {
		return true
	}

	return false
}

// SetLastUpdate gets a reference to the given NullableTime and assigns it to the LastUpdate field.
func (o *Character) SetLastUpdate(v time.Time) {
	o.LastUpdate.Set(&v)
}
// SetLastUpdateNil sets the value for LastUpdate to be an explicit nil
func (o *Character) SetLastUpdateNil() {
	o.LastUpdate.Set(nil)
}

// UnsetLastUpdate ensures that no value is present for LastUpdate, not even an explicit nil
func (o *Character) UnsetLastUpdate() {
	o.LastUpdate.Unset()
}

// GetCorporation returns the Corporation field value if set, zero value otherwise.
func (o *Character) GetCorporation() Corporation {
	if o == nil || o.Corporation == nil {
		var ret Corporation
		return ret
	}
	return *o.Corporation
}

// GetCorporationOk returns a tuple with the Corporation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Character) GetCorporationOk() (*Corporation, bool) {
	if o == nil || o.Corporation == nil {
		return nil, false
	}
	return o.Corporation, true
}

// HasCorporation returns a boolean if a field has been set.
func (o *Character) HasCorporation() bool {
	if o != nil && o.Corporation != nil {
		return true
	}

	return false
}

// SetCorporation gets a reference to the given Corporation and assigns it to the Corporation field.
func (o *Character) SetCorporation(v Corporation) {
	o.Corporation = &v
}

// GetCharacterNameChanges returns the CharacterNameChanges field value if set, zero value otherwise.
func (o *Character) GetCharacterNameChanges() []CharacterNameChange {
	if o == nil || o.CharacterNameChanges == nil {
		var ret []CharacterNameChange
		return ret
	}
	return *o.CharacterNameChanges
}

// GetCharacterNameChangesOk returns a tuple with the CharacterNameChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Character) GetCharacterNameChangesOk() (*[]CharacterNameChange, bool) {
	if o == nil || o.CharacterNameChanges == nil {
		return nil, false
	}
	return o.CharacterNameChanges, true
}

// HasCharacterNameChanges returns a boolean if a field has been set.
func (o *Character) HasCharacterNameChanges() bool {
	if o != nil && o.CharacterNameChanges != nil {
		return true
	}

	return false
}

// SetCharacterNameChanges gets a reference to the given []CharacterNameChange and assigns it to the CharacterNameChanges field.
func (o *Character) SetCharacterNameChanges(v []CharacterNameChange) {
	o.CharacterNameChanges = &v
}

func (o Character) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Main != nil {
		toSerialize["main"] = o.Main
	}
	if o.ValidToken.IsSet() {
		toSerialize["validToken"] = o.ValidToken.Get()
	}
	if o.ValidTokenTime.IsSet() {
		toSerialize["validTokenTime"] = o.ValidTokenTime.Get()
	}
	if o.Created.IsSet() {
		toSerialize["created"] = o.Created.Get()
	}
	if o.LastUpdate.IsSet() {
		toSerialize["lastUpdate"] = o.LastUpdate.Get()
	}
	if o.Corporation != nil {
		toSerialize["corporation"] = o.Corporation
	}
	if o.CharacterNameChanges != nil {
		toSerialize["characterNameChanges"] = o.CharacterNameChanges
	}
	return json.Marshal(toSerialize)
}

type NullableCharacter struct {
	value *Character
	isSet bool
}

func (v NullableCharacter) Get() *Character {
	return v.value
}

func (v *NullableCharacter) Set(val *Character) {
	v.value = val
	v.isSet = true
}

func (v NullableCharacter) IsSet() bool {
	return v.isSet
}

func (v *NullableCharacter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharacter(val *Character) *NullableCharacter {
	return &NullableCharacter{value: val, isSet: true}
}

func (v NullableCharacter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharacter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


