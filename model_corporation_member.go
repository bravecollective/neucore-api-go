/*
Neucore API

Client library of Neucore API

API version: 2.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
	"time"
)

// checks if the CorporationMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorporationMember{}

// CorporationMember The player property contains only id and name, character does not contain corporation.
type CorporationMember struct {
	Player *Player `json:"player,omitempty"`
	// EVE Character ID.
	Id NullableInt64 `json:"id"`
	// EVE Character name.
	Name NullableString `json:"name"`
	Location *EsiLocation `json:"location,omitempty"`
	LogoffDate NullableTime `json:"logoffDate,omitempty"`
	LogonDate NullableTime `json:"logonDate,omitempty"`
	ShipType *EsiType `json:"shipType,omitempty"`
	StartDate NullableTime `json:"startDate,omitempty"`
	Character *Character `json:"character,omitempty"`
	// Date and time of the last sent mail.
	MissingCharacterMailSentDate NullableTime `json:"missingCharacterMailSentDate,omitempty"`
	// Result of the last sent mail (OK, Blocked, CSPA charge > 0)
	MissingCharacterMailSentResult NullableString `json:"missingCharacterMailSentResult,omitempty"`
	// Number of mails sent, is reset when the character is added.
	MissingCharacterMailSentNumber *int32 `json:"missingCharacterMailSentNumber,omitempty"`
}

// NewCorporationMember instantiates a new CorporationMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorporationMember(id NullableInt64, name NullableString) *CorporationMember {
	this := CorporationMember{}
	this.Id = id
	this.Name = name
	return &this
}

// NewCorporationMemberWithDefaults instantiates a new CorporationMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorporationMemberWithDefaults() *CorporationMember {
	this := CorporationMember{}
	return &this
}

// GetPlayer returns the Player field value if set, zero value otherwise.
func (o *CorporationMember) GetPlayer() Player {
	if o == nil || IsNil(o.Player) {
		var ret Player
		return ret
	}
	return *o.Player
}

// GetPlayerOk returns a tuple with the Player field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporationMember) GetPlayerOk() (*Player, bool) {
	if o == nil || IsNil(o.Player) {
		return nil, false
	}
	return o.Player, true
}

// HasPlayer returns a boolean if a field has been set.
func (o *CorporationMember) HasPlayer() bool {
	if o != nil && !IsNil(o.Player) {
		return true
	}

	return false
}

// SetPlayer gets a reference to the given Player and assigns it to the Player field.
func (o *CorporationMember) SetPlayer(v Player) {
	o.Player = &v
}

// GetId returns the Id field value
// If the value is explicit nil, the zero value for int64 will be returned
func (o *CorporationMember) GetId() int64 {
	if o == nil || o.Id.Get() == nil {
		var ret int64
		return ret
	}

	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporationMember) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// SetId sets field value
func (o *CorporationMember) SetId(v int64) {
	o.Id.Set(&v)
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *CorporationMember) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporationMember) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *CorporationMember) SetName(v string) {
	o.Name.Set(&v)
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *CorporationMember) GetLocation() EsiLocation {
	if o == nil || IsNil(o.Location) {
		var ret EsiLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporationMember) GetLocationOk() (*EsiLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *CorporationMember) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given EsiLocation and assigns it to the Location field.
func (o *CorporationMember) SetLocation(v EsiLocation) {
	o.Location = &v
}

// GetLogoffDate returns the LogoffDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporationMember) GetLogoffDate() time.Time {
	if o == nil || IsNil(o.LogoffDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LogoffDate.Get()
}

// GetLogoffDateOk returns a tuple with the LogoffDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporationMember) GetLogoffDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogoffDate.Get(), o.LogoffDate.IsSet()
}

// HasLogoffDate returns a boolean if a field has been set.
func (o *CorporationMember) HasLogoffDate() bool {
	if o != nil && o.LogoffDate.IsSet() {
		return true
	}

	return false
}

// SetLogoffDate gets a reference to the given NullableTime and assigns it to the LogoffDate field.
func (o *CorporationMember) SetLogoffDate(v time.Time) {
	o.LogoffDate.Set(&v)
}
// SetLogoffDateNil sets the value for LogoffDate to be an explicit nil
func (o *CorporationMember) SetLogoffDateNil() {
	o.LogoffDate.Set(nil)
}

// UnsetLogoffDate ensures that no value is present for LogoffDate, not even an explicit nil
func (o *CorporationMember) UnsetLogoffDate() {
	o.LogoffDate.Unset()
}

// GetLogonDate returns the LogonDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporationMember) GetLogonDate() time.Time {
	if o == nil || IsNil(o.LogonDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LogonDate.Get()
}

// GetLogonDateOk returns a tuple with the LogonDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporationMember) GetLogonDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LogonDate.Get(), o.LogonDate.IsSet()
}

// HasLogonDate returns a boolean if a field has been set.
func (o *CorporationMember) HasLogonDate() bool {
	if o != nil && o.LogonDate.IsSet() {
		return true
	}

	return false
}

// SetLogonDate gets a reference to the given NullableTime and assigns it to the LogonDate field.
func (o *CorporationMember) SetLogonDate(v time.Time) {
	o.LogonDate.Set(&v)
}
// SetLogonDateNil sets the value for LogonDate to be an explicit nil
func (o *CorporationMember) SetLogonDateNil() {
	o.LogonDate.Set(nil)
}

// UnsetLogonDate ensures that no value is present for LogonDate, not even an explicit nil
func (o *CorporationMember) UnsetLogonDate() {
	o.LogonDate.Unset()
}

// GetShipType returns the ShipType field value if set, zero value otherwise.
func (o *CorporationMember) GetShipType() EsiType {
	if o == nil || IsNil(o.ShipType) {
		var ret EsiType
		return ret
	}
	return *o.ShipType
}

// GetShipTypeOk returns a tuple with the ShipType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporationMember) GetShipTypeOk() (*EsiType, bool) {
	if o == nil || IsNil(o.ShipType) {
		return nil, false
	}
	return o.ShipType, true
}

// HasShipType returns a boolean if a field has been set.
func (o *CorporationMember) HasShipType() bool {
	if o != nil && !IsNil(o.ShipType) {
		return true
	}

	return false
}

// SetShipType gets a reference to the given EsiType and assigns it to the ShipType field.
func (o *CorporationMember) SetShipType(v EsiType) {
	o.ShipType = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporationMember) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.StartDate.Get()
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporationMember) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartDate.Get(), o.StartDate.IsSet()
}

// HasStartDate returns a boolean if a field has been set.
func (o *CorporationMember) HasStartDate() bool {
	if o != nil && o.StartDate.IsSet() {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given NullableTime and assigns it to the StartDate field.
func (o *CorporationMember) SetStartDate(v time.Time) {
	o.StartDate.Set(&v)
}
// SetStartDateNil sets the value for StartDate to be an explicit nil
func (o *CorporationMember) SetStartDateNil() {
	o.StartDate.Set(nil)
}

// UnsetStartDate ensures that no value is present for StartDate, not even an explicit nil
func (o *CorporationMember) UnsetStartDate() {
	o.StartDate.Unset()
}

// GetCharacter returns the Character field value if set, zero value otherwise.
func (o *CorporationMember) GetCharacter() Character {
	if o == nil || IsNil(o.Character) {
		var ret Character
		return ret
	}
	return *o.Character
}

// GetCharacterOk returns a tuple with the Character field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporationMember) GetCharacterOk() (*Character, bool) {
	if o == nil || IsNil(o.Character) {
		return nil, false
	}
	return o.Character, true
}

// HasCharacter returns a boolean if a field has been set.
func (o *CorporationMember) HasCharacter() bool {
	if o != nil && !IsNil(o.Character) {
		return true
	}

	return false
}

// SetCharacter gets a reference to the given Character and assigns it to the Character field.
func (o *CorporationMember) SetCharacter(v Character) {
	o.Character = &v
}

// GetMissingCharacterMailSentDate returns the MissingCharacterMailSentDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporationMember) GetMissingCharacterMailSentDate() time.Time {
	if o == nil || IsNil(o.MissingCharacterMailSentDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.MissingCharacterMailSentDate.Get()
}

// GetMissingCharacterMailSentDateOk returns a tuple with the MissingCharacterMailSentDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporationMember) GetMissingCharacterMailSentDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.MissingCharacterMailSentDate.Get(), o.MissingCharacterMailSentDate.IsSet()
}

// HasMissingCharacterMailSentDate returns a boolean if a field has been set.
func (o *CorporationMember) HasMissingCharacterMailSentDate() bool {
	if o != nil && o.MissingCharacterMailSentDate.IsSet() {
		return true
	}

	return false
}

// SetMissingCharacterMailSentDate gets a reference to the given NullableTime and assigns it to the MissingCharacterMailSentDate field.
func (o *CorporationMember) SetMissingCharacterMailSentDate(v time.Time) {
	o.MissingCharacterMailSentDate.Set(&v)
}
// SetMissingCharacterMailSentDateNil sets the value for MissingCharacterMailSentDate to be an explicit nil
func (o *CorporationMember) SetMissingCharacterMailSentDateNil() {
	o.MissingCharacterMailSentDate.Set(nil)
}

// UnsetMissingCharacterMailSentDate ensures that no value is present for MissingCharacterMailSentDate, not even an explicit nil
func (o *CorporationMember) UnsetMissingCharacterMailSentDate() {
	o.MissingCharacterMailSentDate.Unset()
}

// GetMissingCharacterMailSentResult returns the MissingCharacterMailSentResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CorporationMember) GetMissingCharacterMailSentResult() string {
	if o == nil || IsNil(o.MissingCharacterMailSentResult.Get()) {
		var ret string
		return ret
	}
	return *o.MissingCharacterMailSentResult.Get()
}

// GetMissingCharacterMailSentResultOk returns a tuple with the MissingCharacterMailSentResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CorporationMember) GetMissingCharacterMailSentResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MissingCharacterMailSentResult.Get(), o.MissingCharacterMailSentResult.IsSet()
}

// HasMissingCharacterMailSentResult returns a boolean if a field has been set.
func (o *CorporationMember) HasMissingCharacterMailSentResult() bool {
	if o != nil && o.MissingCharacterMailSentResult.IsSet() {
		return true
	}

	return false
}

// SetMissingCharacterMailSentResult gets a reference to the given NullableString and assigns it to the MissingCharacterMailSentResult field.
func (o *CorporationMember) SetMissingCharacterMailSentResult(v string) {
	o.MissingCharacterMailSentResult.Set(&v)
}
// SetMissingCharacterMailSentResultNil sets the value for MissingCharacterMailSentResult to be an explicit nil
func (o *CorporationMember) SetMissingCharacterMailSentResultNil() {
	o.MissingCharacterMailSentResult.Set(nil)
}

// UnsetMissingCharacterMailSentResult ensures that no value is present for MissingCharacterMailSentResult, not even an explicit nil
func (o *CorporationMember) UnsetMissingCharacterMailSentResult() {
	o.MissingCharacterMailSentResult.Unset()
}

// GetMissingCharacterMailSentNumber returns the MissingCharacterMailSentNumber field value if set, zero value otherwise.
func (o *CorporationMember) GetMissingCharacterMailSentNumber() int32 {
	if o == nil || IsNil(o.MissingCharacterMailSentNumber) {
		var ret int32
		return ret
	}
	return *o.MissingCharacterMailSentNumber
}

// GetMissingCharacterMailSentNumberOk returns a tuple with the MissingCharacterMailSentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorporationMember) GetMissingCharacterMailSentNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.MissingCharacterMailSentNumber) {
		return nil, false
	}
	return o.MissingCharacterMailSentNumber, true
}

// HasMissingCharacterMailSentNumber returns a boolean if a field has been set.
func (o *CorporationMember) HasMissingCharacterMailSentNumber() bool {
	if o != nil && !IsNil(o.MissingCharacterMailSentNumber) {
		return true
	}

	return false
}

// SetMissingCharacterMailSentNumber gets a reference to the given int32 and assigns it to the MissingCharacterMailSentNumber field.
func (o *CorporationMember) SetMissingCharacterMailSentNumber(v int32) {
	o.MissingCharacterMailSentNumber = &v
}

func (o CorporationMember) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorporationMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Player) {
		toSerialize["player"] = o.Player
	}
	toSerialize["id"] = o.Id.Get()
	toSerialize["name"] = o.Name.Get()
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if o.LogoffDate.IsSet() {
		toSerialize["logoffDate"] = o.LogoffDate.Get()
	}
	if o.LogonDate.IsSet() {
		toSerialize["logonDate"] = o.LogonDate.Get()
	}
	if !IsNil(o.ShipType) {
		toSerialize["shipType"] = o.ShipType
	}
	if o.StartDate.IsSet() {
		toSerialize["startDate"] = o.StartDate.Get()
	}
	if !IsNil(o.Character) {
		toSerialize["character"] = o.Character
	}
	if o.MissingCharacterMailSentDate.IsSet() {
		toSerialize["missingCharacterMailSentDate"] = o.MissingCharacterMailSentDate.Get()
	}
	if o.MissingCharacterMailSentResult.IsSet() {
		toSerialize["missingCharacterMailSentResult"] = o.MissingCharacterMailSentResult.Get()
	}
	if !IsNil(o.MissingCharacterMailSentNumber) {
		toSerialize["missingCharacterMailSentNumber"] = o.MissingCharacterMailSentNumber
	}
	return toSerialize, nil
}

type NullableCorporationMember struct {
	value *CorporationMember
	isSet bool
}

func (v NullableCorporationMember) Get() *CorporationMember {
	return v.value
}

func (v *NullableCorporationMember) Set(val *CorporationMember) {
	v.value = val
	v.isSet = true
}

func (v NullableCorporationMember) IsSet() bool {
	return v.isSet
}

func (v *NullableCorporationMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorporationMember(val *CorporationMember) *NullableCorporationMember {
	return &NullableCorporationMember{value: val, isSet: true}
}

func (v NullableCorporationMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorporationMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


