/*
Neucore API

Client library of Neucore API

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package neucoreapi

import (
	"encoding/json"
	"time"
)

// checks if the CharacterNameChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CharacterNameChange{}

// CharacterNameChange A previous character name.
type CharacterNameChange struct {
	OldName string `json:"oldName"`
	ChangeDate NullableTime `json:"changeDate"`
}

// NewCharacterNameChange instantiates a new CharacterNameChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCharacterNameChange(oldName string, changeDate NullableTime) *CharacterNameChange {
	this := CharacterNameChange{}
	this.OldName = oldName
	this.ChangeDate = changeDate
	return &this
}

// NewCharacterNameChangeWithDefaults instantiates a new CharacterNameChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCharacterNameChangeWithDefaults() *CharacterNameChange {
	this := CharacterNameChange{}
	return &this
}

// GetOldName returns the OldName field value
func (o *CharacterNameChange) GetOldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OldName
}

// GetOldNameOk returns a tuple with the OldName field value
// and a boolean to check if the value has been set.
func (o *CharacterNameChange) GetOldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldName, true
}

// SetOldName sets field value
func (o *CharacterNameChange) SetOldName(v string) {
	o.OldName = v
}

// GetChangeDate returns the ChangeDate field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *CharacterNameChange) GetChangeDate() time.Time {
	if o == nil || o.ChangeDate.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ChangeDate.Get()
}

// GetChangeDateOk returns a tuple with the ChangeDate field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CharacterNameChange) GetChangeDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ChangeDate.Get(), o.ChangeDate.IsSet()
}

// SetChangeDate sets field value
func (o *CharacterNameChange) SetChangeDate(v time.Time) {
	o.ChangeDate.Set(&v)
}

func (o CharacterNameChange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CharacterNameChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["oldName"] = o.OldName
	toSerialize["changeDate"] = o.ChangeDate.Get()
	return toSerialize, nil
}

type NullableCharacterNameChange struct {
	value *CharacterNameChange
	isSet bool
}

func (v NullableCharacterNameChange) Get() *CharacterNameChange {
	return v.value
}

func (v *NullableCharacterNameChange) Set(val *CharacterNameChange) {
	v.value = val
	v.isSet = true
}

func (v NullableCharacterNameChange) IsSet() bool {
	return v.isSet
}

func (v *NullableCharacterNameChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCharacterNameChange(val *CharacterNameChange) *NullableCharacterNameChange {
	return &NullableCharacterNameChange{value: val, isSet: true}
}

func (v NullableCharacterNameChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCharacterNameChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


